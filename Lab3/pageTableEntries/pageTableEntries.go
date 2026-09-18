package pageTableEntries

type PageTableEntries struct {
    Presence bool
    Reference bool
    Modification bool
    PhysicalPageNumber  int
}

func NewPageTableEntries(PhysicalPageNumber int) *PageTableEntries{
	return &PageTableEntries{
		Presence: false,
		Reference: false,
		Modification: false,
		PhysicalPageNumber: PhysicalPageNumber,
	}
}
