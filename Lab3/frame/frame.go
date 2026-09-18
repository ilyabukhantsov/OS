package frame

import "Lab3/process"

type Frame struct {
    Number            int
    Owner             *process.Process
    VirtualPageNumber int
}

func NewFrame(number int) Frame {
    return Frame{
        Number:            number,
        Owner:             nil,
        VirtualPageNumber: -1,
    }
}

