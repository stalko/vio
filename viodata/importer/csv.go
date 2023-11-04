package importer

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/stalko/viodata/storage"
)

type csvImporter struct {
	s                storage.Storage
	acceptedEntries  int
	discardedEntries int
	m                sync.Mutex
	workerTimeout    time.Duration
}

func NewCSVImporter(s storage.Storage, workerTimeout time.Duration) Importer {
	return &csvImporter{
		s:             s,
		workerTimeout: workerTimeout,
	}
}

func (s *csvImporter) worker(output <-chan []string) {
	for record := range output {
		ctx, cancelCtx := context.WithTimeout(context.Background(), s.workerTimeout)

		model, err := RecordToModel(record)
		if err != nil {
			s.incrementDiscardedEntries()
			cancelCtx()
			continue
		}

		//insert into storage
		err = s.s.InsertIPLocation(ctx,
			model.IPAddress,
			model.Country,
			model.CountryCode,
			model.City,
			model.Latitude,
			model.Longitude,
			model.MysteryValue,
		)
		if err != nil {
			s.incrementDiscardedEntries()
			cancelCtx()
			continue
		}

		s.incrementAcceptedEntries()
		cancelCtx()
	}
}

func (s *csvImporter) incrementAcceptedEntries() {
	s.m.Lock()
	s.acceptedEntries++
	s.m.Unlock()
}

func (s *csvImporter) incrementDiscardedEntries() {
	s.m.Lock()
	s.discardedEntries++
	s.m.Unlock()
}

func (s *csvImporter) Import(filePath string, countGoRoutine int) (*Output, error) {
	if countGoRoutine <= 0 {
		return nil, fmt.Errorf("count of the go routine can't be 0 or negative")
	}

	start := time.Now()

	// Create a channel to store CSV records
	output := make(chan []string)
	defer close(output)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//create a group of goroutine
	for i := 0; i < countGoRoutine; i++ {
		go s.worker(output)
	}
	//all workers are created

	reader := csv.NewReader(file)

	_, err = reader.Read() //skip headers
	if err != nil {
		return nil, err
	}

	totalRows := 0
	for {
		record, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		totalRows++
		output <- record
	}

	for {
		s.m.Lock()
		isAllWorkerDone := totalRows == s.acceptedEntries+s.discardedEntries
		s.m.Unlock()

		if isAllWorkerDone {
			break
		}
		time.Sleep(time.Second) // waiting for all workers to be done
	}

	return &Output{
		Duration:         time.Since(start),
		AcceptedEntries:  s.acceptedEntries,
		DiscardedEntries: s.discardedEntries,
	}, nil
}
