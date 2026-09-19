package pageTableEntries

type PageTableEntries struct {
	VirtualPageNumber   int
	PhysicalFrameNumber int
	Present             bool
	Referenced          bool
	Modified            bool
	InFS                bool // Is page saved
}

func NewPageTableEntries(vpage int) *PageTableEntries {
	return &PageTableEntries{
		VirtualPageNumber:   vpage,
		PhysicalFrameNumber: -1,
		Present:             false,
		Referenced:          false,
		Modified:            false,
		InFS:                false,
	}
}
