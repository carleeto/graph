package main

import "fmt"

type PortType int

const (
	String = iota
	Int
)

//Everything needs a name and a description
type Identifiable interface {
	Name() string
	Description() string
}

//A port also has a type that specifies the type of data it carries
type Port interface {
	Identifiable
	Type() PortType
}

//An InPort, when used, provides a means of input
//The Input method takes a certain type.
//For now, we don't care what that type is
type InPort interface {
	Port
	Input(interface{}) error
}

//An OutPort, when used, provides a means of output
//The Output method takes a certain type.
//For now, we don't care what that type is
type OutPort interface {
	Port
	Output() (interface{}, error)
}

// A source only has outputs
type Source interface{
  Outputs() []OutPort
}

// A sink only has inputs
type Sink interface{
  Inputs() []InPort
}

//A Link has outputs and inputs
type Link interface{
  Identifiable
  Source
  Sink
}

func connect(output OutPort, to_input InPort) error {
	if output.Type() == to_input.Type() {
		fmt.Println("connected", output.Name(), "to", to_input.Name())
    if o, err := output.Output(); err == nil{
      to_input.Input(o)
    }
		return nil
	}
	return fmt.Errorf("port types dont match. wont connect")
}

// Some sample implementations
type SourceInt struct{}
func (_ SourceInt) Name() string        { return "int source" }
func (_ SourceInt) Description() string { return "source of integers" }
func (_ SourceInt) Type() PortType      { return Int }
func (i SourceInt) Output() (interface{},error) {
	fmt.Println(i.Name(), "output triggered")
	return 5, nil
}

type SinkInt struct{}
func (_ SinkInt) Name() string        { return "int sink" }
func (_ SinkInt) Description() string { return "source of integers" }
func (_ SinkInt) Type() PortType      { return Int }
func (i SinkInt) Input(d interface{}) error {
  fmt.Println(i.Name(), "input triggered. data:",d)
	return nil
}

//A sink has a least one input
type SinkStr struct{}
func (_ SinkStr) Name() string        { return "string sink" }
func (_ SinkStr) Description() string { return "source of integers" }
func (_ SinkStr) Type() PortType      { return String }
func (s SinkStr) Input(interface{}) error {
	fmt.Println(s.Name(), "output triggered")
	return nil
}

func main() {
	var a1 SourceInt
  var a2 SinkInt
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
