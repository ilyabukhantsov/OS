package kernel

import (
	"Lab3/frame"
	"Lab3/memoryManagementUnit"
	"Lab3/process"
	"fmt"
)

type Kernel struct{
			Process *[]process.Process
			MemoryManagementUnit memoryManagementUnit.MemoryManagementUnit
			PhysicalFrames *[]frame.Frame
}

func NewKernel(p process.Process, mmu memoryManagementUnit.MemoryManagementUnit, frameNumber int) *Kernel {
	frames := []frame.Frame{}
	for index := range frameNumber {
   	frames = append(frames, frame.NewFrame(index))
	}
	return &Kernel{
		Process: &[]process.Process{p},
		MemoryManagementUnit: mmu,
		PhysicalFrames: &frames,
	}
}

func (k *Kernel) Start() int{
	for true{
		fmt.Println("Working with New Process")
		
		p := (*k.Process)[0]
		p.Work()

	}
	return 0
}
