package kernel

import (
	"Lab3/process"
	"Lab3/memoryManagementUnit"
	"Lab3/frame"
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
		MemoryManagementUnit: memoryManagementUnit.MemoryManagementUnit{},
		PhysicalFrames: &frames,
	}
}

func (*Kernel) Start(algorithm string, filename string) string{
	if algorithm == "random"{
		
	} else{

	}
	return "1"
}
