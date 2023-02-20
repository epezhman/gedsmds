package mdsprocessor

type Processor struct {
	uuid string
}

func InitProcessor() *Processor {
	tempProcessor := &Processor{}

	return tempProcessor
}
