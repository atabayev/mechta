package service

import (
	"sync"
	"testing"
)

func TestCalculateSum(t *testing.T) {
	s := &Service{
		data: Data{
			{A: 1, B: 2},
			{A: 3, B: 4},
			{A: 5, B: 6},
			{A: 7, B: 8},
		},
	}
	goroutineCount := 2

	result, err := s.CalculateSum(goroutineCount)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != 36 {
		t.Errorf("Expected %d, but got %d", 36, result)
	}
}

func TestCalculateForSlice(t *testing.T) {
	data := Data{
		{A: 1, B: 2},
		{A: 3, B: 4},
		{A: 5, B: 6},
	}

	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	service := Service{}
	go service.calculateForSlice(data, ch, &wg)

	result := <-ch
	wg.Wait()
	close(ch)

	if result != 21 {
		t.Errorf("Expected %d, but got %d", 21, result)
	}

}

func TestCalcGenSum(t *testing.T) {
	ch := make(chan int)
	res := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		service := Service{}
		service.calcGenSum(ch, res)
		close(res)
	}()

	result := <-res

	if result != 10 {
		t.Errorf("Expected sum to be 10, got %d", result)
	}
}
