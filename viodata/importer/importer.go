package importer

import "time"

type Importer interface {
	Import(filePath string, countGoRoutine int) (*Output, error)
}

type Output struct {
	Duration         time.Duration
	AcceptedEntries  int
	DiscardedEntries int
}
