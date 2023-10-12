package mechta

import "mechta/internal/app/service"

type Service interface {
	Unmarshal(filename string) error
	CalculateSum(goroutineCount int) (int, error)
}

type Mechta struct {
	filename string
	gc       int
	svc      Service
}

func New(filename string, gc int) *Mechta {
	m := &Mechta{}
	m.gc = gc
	m.filename = filename
	m.svc = service.New()
	return m
}

func (m *Mechta) Run() error {
	err := m.svc.Unmarshal(m.filename)
	if err != nil {
		return err
	}

	sum, err := m.svc.CalculateSum(m.gc)
	if err != nil {
		return err
	}

	println(sum)
	return nil
}
