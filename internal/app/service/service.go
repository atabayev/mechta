// Package service reads an array of data from a file,
// deserializes it into a structure, and calculates
// the sum of the slice elements by dividing the work
// among goroutines. It utilizes the 'easyjson' library for
// JSON deserialization.
package service

import (
	"os"
	"sync"

	"github.com/mailru/easyjson"
)

//easyjson:json
type Data []struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Service struct {
	data Data
}

func New() *Service {
	return &Service{}
}

// Unmarshal read data from a file and deserialize it into
// a structure
func (s *Service) Unmarshal(filename string) error {
	body, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return easyjson.Unmarshal(body, &s.data)
}

// CalculateSum divides a large array of data evenly among goroutines,
// then starts each goroutine with its respective slice portion.
func (s *Service) CalculateSum(goroutineCount int) (int, error) {
	ch := make(chan int)
	res := make(chan int)
	var wg sync.WaitGroup

	// calculate portion of array for one goroutine
	// p contains the number of elements processed by one goroutine.
	p := 0
	count := 0
	if len(s.data) > goroutineCount {
		p = len(s.data) / goroutineCount
		count = goroutineCount
	} else {
		p = 1
		count = len(s.data)
	}
	for i := 0; i < count; i++ {
		wg.Add(1)
		// startPos holds the initial index, while
		// finishPos holds the final index of the allocated slice.
		// P.S. Defined in separate variables for clarity and code readability.
		startPos := p * i
		finishPos := (p * i) + p
		if i == count-1 {
			// if last lap
			go s.calculateForSlice(s.data[startPos:], ch, &wg)
		} else {
			go s.calculateForSlice(s.data[startPos:finishPos], ch, &wg)
		}
	}

	go s.calcGenSum(ch, res)

	wg.Wait()
	close(ch)

	return <-res, nil
}

// calculateForSlice calculates the sum of elements in the allocated slice.
func (s *Service) calculateForSlice(data Data, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, d := range data {
		sum = sum + d.A + d.B
	}
	ch <- sum
}

// calcGenSum collects and sums up the results of
// goroutine execution
func (s *Service) calcGenSum(ch, res chan int) {
	sum := 0
	for c := range ch {
		sum += c
	}
	res <- sum
}
