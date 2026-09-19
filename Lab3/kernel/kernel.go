package kernel

import (
	"Lab3/frame"
	"Lab3/memoryManagementUnit"
	"Lab3/pageTableEntries"
	"Lab3/process"
	"fmt"
	"math/rand"
)

type Kernel struct {
	Processes            []*process.Process
	MemoryManagementUnit *memoryManagementUnit.MemoryManagementUnit
	PhysicalFrames       []*frame.Frame
	Algorithm            string
	PageFaultCount       int
	TotalAccesses        int
}

func NewKernel(mmu *memoryManagementUnit.MemoryManagementUnit, frameNumber int, algorithm string) *Kernel {
	frames := make([]*frame.Frame, frameNumber)
	for i := 0; i < frameNumber; i++ {
		f := frame.NewFrame(i)
		frames[i] = &f
	}

	k := &Kernel{
		Processes:            []*process.Process{},
		MemoryManagementUnit: mmu,
		PhysicalFrames:       frames,
		Algorithm:            algorithm,
	}

	mmu.SetKernel(k)
	return k
}

func (k *Kernel) AddProcess(p *process.Process) {
	k.Processes = append(k.Processes, p)
}

func (k *Kernel) Start() int {
	quantum := 5

	for len(k.Processes) > 0 {
		p := k.Processes[0]

		for i := 0; i < quantum && p.RemainingLife > 0; i++ {
			k.TotalAccesses++
			err := p.Work()
			if err != nil {
				fmt.Printf("Error during execution of PID %d: %v\n", p.ID, err)
			}
		}

		if p.RemainingLife <= 0 {
			fmt.Printf("--- Process %d FINISHED ---\n", p.ID)
			k.freeProcessMemory(p.ID)
			k.Processes = k.Processes[1:]
		} else {
			k.Processes = k.Processes[1:]
			k.Processes = append(k.Processes, p)
		}
	}

	fmt.Printf("\n=== SIMULATION COMPLETED ===\nTotal Accesses: %d, Page Faults: %d (%.2f%%)\n",
		k.TotalAccesses, k.PageFaultCount, float64(k.PageFaultCount)/float64(k.TotalAccesses)*100)

	return 0
}

func (k *Kernel) HandlePageFault(pid int, pageTable []*pageTableEntries.PageTableEntries, vpage int) error {
	k.PageFaultCount++

	var targetFrame *frame.Frame
	for _, f := range k.PhysicalFrames {
		if f.IsFree() {
			targetFrame = f
			break
		}
	}

	if targetFrame == nil {
		targetFrame = k.selectFrameToEvict(pageTable)
		k.evictFrame(targetFrame)
	}

	targetFrame.OwnerPID = pid
	targetFrame.VirtualPageNumber = vpage

	pte := pageTable[vpage]
	pte.Present = true
	pte.PhysicalFrameNumber = targetFrame.Number

	fmt.Printf("[KERNEL] Loaded PID %d VPage %d -> Frame %d\n", pid, vpage, targetFrame.Number)
	return nil
}

func (k *Kernel) evictFrame(f *frame.Frame) {
	for _, p := range k.Processes {
		if p.ID == f.OwnerPID {
			for _, pte := range p.Table {
				if pte.PhysicalFrameNumber == f.Number && pte.Present {
					pte.Present = false
					pte.PhysicalFrameNumber = -1
					if pte.Modified {
						pte.InFS = true
					}
					fmt.Printf("[KERNEL] Evicted PID %d VPage %d from Frame %d (Was Modified: %v)\n",
						p.ID, pte.VirtualPageNumber, f.Number, pte.Modified)
					break
				}
			}
		}
	}
	f.OwnerPID = -1
	f.VirtualPageNumber = -1
}

func (k *Kernel) selectFrameToEvict(currentPageTable []*pageTableEntries.PageTableEntries) *frame.Frame {
	switch k.Algorithm {
	case "random":
		idx := rand.Intn(len(k.PhysicalFrames))
		return k.PhysicalFrames[idx]

	case "nru":
		var classes [4][]*frame.Frame

		for _, f := range k.PhysicalFrames {
			pte := k.getPTEForFrame(f)
			if pte == nil {
				continue
			}
			classIdx := 0
			if pte.Referenced {
				classIdx += 2
			}
			if pte.Modified {
				classIdx += 1
			}
			classes[classIdx] = append(classes[classIdx], f)
		}

		for i := 0; i < 4; i++ {
			if len(classes[i]) > 0 {
				return classes[i][rand.Intn(len(classes[i]))]
			}
		}
		return k.PhysicalFrames[0]

	default:
		return k.PhysicalFrames[rand.Intn(len(k.PhysicalFrames))]
	}
}

func (k *Kernel) getPTEForFrame(f *frame.Frame) *pageTableEntries.PageTableEntries {
	for _, p := range k.Processes {
		if p.ID == f.OwnerPID {
			for _, pte := range p.Table {
				if pte.PhysicalFrameNumber == f.Number && pte.Present {
					return pte
				}
			}
		}
	}
	return nil
}

func (k *Kernel) freeProcessMemory(pid int) {
	for _, f := range k.PhysicalFrames {
		if f.OwnerPID == pid {
			f.OwnerPID = -1
			f.VirtualPageNumber = -1
		}
	}
}
