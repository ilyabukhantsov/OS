package core

import (
	"Lab3/process"
	"io"
)

type Core struct{
			Process *[]process.Process
			Log io.Writer
}

func NewCore(p process.Process, logger io.Writer) *Core {
	return &Core{
		Process: &[]process.Process{p},
		Log:     logger,
	}
}

func (*Core) Start(algorithm string, filename string) string{
	if algorithm == "random"{
		
	} else{

	}
	return "1"
}
