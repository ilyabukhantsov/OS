package main

import (
	"Lab3/kernel"
	"Lab3/memoryManagementUnit"
	"Lab3/process"
	"fmt"
)

func main() {
	fmt.Println("--- Starting OS Virtual Memory Simulator ---")

	mmu := memoryManagementUnit.NewMemoryManagementUnit()

	sysKernel := kernel.NewKernel(mmu, 4, "nru")

	p1 := process.NewProcess(1, 50, 16, mmu)
	p2 := process.NewProcess(2, 50, 16, mmu)

	sysKernel.AddProcess(p1)
	sysKernel.AddProcess(p2)

	sysKernel.Start()
}

