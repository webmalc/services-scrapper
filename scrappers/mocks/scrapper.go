package mocks

import (
	"sync"

	"github.com/stretchr/testify/mock"
)

// Scrapper is mock object
type Scrapper struct {
	mock.Mock
}

// Error is a method mock
func (s *Scrapper) Scrap(wg *sync.WaitGroup) {
	s.Called(wg)
	wg.Done()
}
