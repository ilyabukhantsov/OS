package kernel

import (
	"Lab3/frame"
	"Lab3/memoryManagementUnit"
	"Lab3/process"
	"fmt"
)

type Kernel struct{
			Process []*process.Process
			MemoryManagementUnit memoryManagementUnit.MemoryManagementUnit
			PhysicalFrames []*frame.Frame
}

func NewKernel(p *process.Process, mmu memoryManagementUnit.MemoryManagementUnit, frameNumber int) *Kernel {
	frames := make([]*frame.Frame, 0, frameNumber)
	for index := range frameNumber {
		f := frame.NewFrame(index)
		frames = append(frames, &f)
	}

	processes := []*process.Process{}
	if p != nil {
		processes = append(processes, p)
	}

	return &Kernel{
		Process:              processes,
		MemoryManagementUnit: mmu,
		PhysicalFrames:       frames,
	}
}

func (k *Kernel) Start() int{
	for len(k.Process) > 0{
		fmt.Println("Working with New Process")
		
		p := (k.Process)[0]
		p.Work()
		if p.RemainingLife <= 0{
			k.Process = k.Process[1:]
			k.Process = append(k.Process, p)
		}
		fmt.Println("We moved to next Process!")
	}

	fmt.Println("No processes left")

	return 0
}
