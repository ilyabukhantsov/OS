package memoryManagementUnit

import (
	"Lab3/pageTableEntries"
	"Lab3/process"
	"fmt"
)

type KernelFaultHandler interface {
	HandlePageFault(pid int, pageTable []*pageTableEntries.PageTableEntries, vpage int) error
}

type MemoryManagementUnit struct {
	Kernel KernelFaultHandler
}

func NewMemoryManagementUnit() *MemoryManagementUnit {
	return &MemoryManagementUnit{}
}

func (mmu *MemoryManagementUnit) SetKernel(kernel KernelFaultHandler) {
	mmu.Kernel = kernel
}

func (mmu *MemoryManagementUnit) Access(pid int, pageTable []*pageTableEntries.PageTableEntries, vpage int, accessType process.AccessType) error {
	if vpage < 0 || vpage >= len(pageTable) {
		return fmt.Errorf("out of bounds virtual page %d", vpage)
	}

	pte := pageTable[vpage]

	if !pte.Present {
		fmt.Printf("[MMU] PAGE FAULT! Process %d, Virtual Page %d\n", pid, vpage)
		err := mmu.Kernel.HandlePageFault(pid, pageTable, vpage)
		if err != nil {
			return err
		}
	}

	pte.Referenced = true
	if accessType == process.Write {
		pte.Modified = true
	}

	return nil
}