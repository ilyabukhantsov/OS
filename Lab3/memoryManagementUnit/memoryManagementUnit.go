package memoryManagementUnit

type MemoryManagementUnit struct{
		algorithm string
}

func NewMemoryManagementUnit(algorithm string) *MemoryManagementUnit{
	return &MemoryManagementUnit{
		algorithm: algorithm,
	}
}

