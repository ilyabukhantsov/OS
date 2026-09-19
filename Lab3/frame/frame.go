package frame

type Frame struct {
	Number            int
	OwnerPID          int
	VirtualPageNumber int
}

func NewFrame(number int) Frame {
	return Frame{
		Number:            number,
		OwnerPID:          -1,
		VirtualPageNumber: -1,
	}
}

func (f *Frame) IsFree() bool {
	return f.OwnerPID == -1
}
