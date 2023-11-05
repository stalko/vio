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
	countBulkInsert  int
}

func NewCSVImporter(s storage.Storage, countBulkInsert int) Importer {
	return &csvImporter{
		s:               s,
		countBulkInsert: countBulkInsert,
	}
}

func (s *csvImporter) worker(output <-chan []string) {
	var queue []storage.InsertIPLocation

	for record := range output {

		model, err := RecordToModel(record)
		if err != nil {
			s.addDiscardedEntries(1)
			continue
		}

		queue = append(queue, storage.InsertIPLocation{
			IPAddress:    model.IPAddress,
			CountryName:  model.Country,
			CountryCode:  model.CountryCode,
			City:         model.City,
			Latitude:     model.Latitude,
			Longitude:    model.Longitude,
			MysteryValue: model.MysteryValue,
		})

		if len(queue) > s.countBulkInsert {
			//insert into storage
			err = s.s.BulkInsertIPLocation(context.Background(), queue)
			if err != nil {
				s.addDiscardedEntries(len(queue))
			} else {
				s.addAcceptedEntries(len(queue))
			}

			queue = []storage.InsertIPLocation{}
		}
	}

	if len(queue) > 0 {
		//insert into storage
		err := s.s.BulkInsertIPLocation(context.Background(), queue)
		if err != nil {
			s.addDiscardedEntries(len(queue))
		} else {
			s.addAcceptedEntries(len(queue))
		}
	}
}

func (s *csvImporter) addAcceptedEntries(i int) {
	s.m.Lock()
	s.acceptedEntries += i
	s.m.Unlock()
}

func (s *csvImporter) addDiscardedEntries(i int) {
	s.m.Lock()
	s.discardedEntries += i
	s.m.Unlock()
}

func (s *csvImporter) Import(filePath string, countGoRoutine int) (*Output, error) {
	if countGoRoutine <= 0 {
		return nil, fmt.Errorf("count of the go routine can't be 0 or negative")
	}

	start := time.Now()

	// Create a channel to store CSV records
	output := make(chan []string)

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

	close(output) //notify all channels that file read is finished

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
