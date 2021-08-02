package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"saskara/protobuf-go/proto-sandbox/using-go/src/complexpb"
	"saskara/protobuf-go/proto-sandbox/using-go/src/enum_example"
	"saskara/protobuf-go/proto-sandbox/using-go/src/example"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()

	readAndWriteDemo(sm)
	JSONDemo(sm)
	doEnum()
	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "Rad first message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Rad second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Rad third message",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {
	ep := enum_example.EnumMessage{
		Id:           42,
		DayOfTheWeek: enum_example.DayOfTheWeek_MONDAY,
	}
	ep.DayOfTheWeek = enum_example.DayOfTheWeek_SATURDAY

	fmt.Println(ep)
}

func JSONDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &example.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Cant convert to JSON", err)
		return ""
	}

	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Coulnt unmarshall", err)
	}
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &example.SimpleMessage{}
	readFromFile("simple.bin", sm2)

	fmt.Println("Read from content", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Cant write", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Cant write", err)
		return err
	}
	fmt.Println("Data has been written ")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Cant read file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("couldnt put the bytes into protobuf struct", err)
		return err2
	}

	return nil
}

func doSimple() *example.SimpleMessage {
	sm := example.SimpleMessage{
		Id:         123124,
		IsSimple:   true,
		Name:       "My simple message",
		SampleList: []int32{1, 4, 5, 6, 4},
	}
	fmt.Println(sm.Name)
	sm.Name = "Rad renamed you"

	fmt.Println(sm.Name)

	return &sm
}
