package main

import "fmt"

type PortType int

const (
	String = iota
	Int
)

type Identifiable interface {
	Name() string
	Description() string
}

type Port interface {
	Identifiable
	Type() PortType
}

type InPort interface {
	Port
	Input(interface{}) error
}

type OutPort interface {
	Port
	Output() (interface{}, error)
}

type Component interface {
	Identifiable
	Inputs() []InPort
	Outputs() []OutPort
}

func connect(output, to_input Port) error {
	if output.Type() == to_input.Type() {
		fmt.Println("connected", output.Name(), "to", to_input.Name())
		return nil
	}
	return fmt.Errorf("port types dont match. wont connect")
}

type SourceInt struct{}

func (_ SourceInt) Name() string        { return "int source" }
func (_ SourceInt) Description() string { return "source of integers" }
func (_ SourceInt) Type() PortType      { return Int }
func (i SourceInt) Input(interface{}) error {
	fmt.Println(i.Name(), "input triggered")
	return nil
}

type SinkStr struct{}

func (_ SinkStr) Name() string        { return "string sink" }
func (_ SinkStr) Description() string { return "source of integers" }
func (_ SinkStr) Type() PortType      { return String }
func (s SinkStr) Output() (interface{}, error) {
	fmt.Println(s.Name(), "output triggered")
	return 5, nil
}

func main() {
	var a1, a2 SourceInt
	if connect(a1, a2) == nil {
		println("source and sink port type matches")
	}

	//Try and connect a source of ints to a sink expecting a string input.
	//Doesn't work
	var b SinkStr
	if err := connect(a1, b); err != nil {
		fmt.Println(err)
	}
}
