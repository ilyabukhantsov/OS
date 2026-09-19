package process

import (
	"Lab3/pageTableEntries"
	"fmt"
	"math/rand"
)

type AccessType int

const (
	Read AccessType = iota
	Write
)

type MMUInterface interface {
	Access(pid int, pageTable []*pageTableEntries.PageTableEntries, vpage int, accessType AccessType) error
}

type Process struct {
	ID            int
	RemainingLife int
	Table         []*pageTableEntries.PageTableEntries
	WorkingSet    []int
	MMU           MMUInterface
	StepCount     int
}

func NewProcess(id int, remainingLife int, numberOfPages int, mmu MMUInterface) *Process {
	table := make([]*pageTableEntries.PageTableEntries, numberOfPages)
	for i := 0; i < numberOfPages; i++ {
		table[i] = pageTableEntries.NewPageTableEntries(i)
	}

	p := &Process{
		ID:            id,
		RemainingLife: remainingLife,
		Table:         table,
		WorkingSet:    []int{},
		MMU:           mmu,
	}
	p.initWorkingSet(4)
	return p
}

func (p *Process) initWorkingSet(size int) {
	if size > len(p.Table) {
		size = len(p.Table)
	}
	p.WorkingSet = make([]int, size)
	for i := 0; i < size; i++ {
		p.WorkingSet[i] = i
	}
}

func (p *Process) Work() error {
	if p.RemainingLife <= 0 {
		return nil
	}

	p.RemainingLife--
	p.StepCount++

	var vpage int
	if rand.Float64() < 0.90 && len(p.WorkingSet) > 0 {
		vpage = p.WorkingSet[rand.Intn(len(p.WorkingSet))]
	} else {
		vpage = rand.Intn(len(p.Table))
	}

	accessType := Read
	if rand.Float64() < 0.30 {
		accessType = Write
	}

	err := p.MMU.Access(p.ID, p.Table, vpage, accessType)
	if err != nil {
		return fmt.Errorf("process %d access fault: %w", p.ID, err)
	}

	return nil
}