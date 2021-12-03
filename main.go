package main

import (
	"fmt"
	"io/ioutil"

	"github.com/jwmwalrus/protobuf-go/complexpb"
	"github.com/jwmwalrus/protobuf-go/enumpb"
	"github.com/jwmwalrus/protobuf-go/simplepb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("Hello world!")

	msg := doSimple()

	readAndWriteDemo(msg)

	jsonDemo(msg)

	doEnum()

	doComplex()
}

func doComplex() {
	msg := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third message",
			},
		},
	}

	fmt.Println(msg.String())
}

func doEnum() {
	msg := enumpb.EnumMessage{
		Id:           32,
		DayOfTheWeek: enumpb.DayOfTheWeek_MONDAY,
	}

	fmt.Println("Enum message:", msg.String())

}

func jsonDemo(msg proto.Message) {
	j, err := toJSON(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("toJSON:", j)

	d := &simplepb.SimpleMessage{}
	if err = fromJSON(j, d); err != nil {
		panic(err)
	}

	fmt.Println("fromJSON:", d.String())
}

func fromJSON(str string, msg proto.Message) (err error) {
	err = protojson.Unmarshal([]byte(str), msg)
	if err != nil {
		return
	}
	return
}

func toJSON(msg proto.Message) (str string, err error) {
	bv, err := protojson.Marshal(msg)
	if err != nil {
		return
	}
	str = string(bv)
	return
}

func readAndWriteDemo(msg proto.Message) {

	if err := writeToFile("simple.bin", msg); err != nil {
		panic(err)
	}

	newMsg := &simplepb.SimpleMessage{}
	err := readFromFile("simple.bin", newMsg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message read is: \n%v\n", newMsg.String())
}

func readFromFile(filename string, msg proto.Message) (err error) {
	bv, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	if err = proto.Unmarshal(bv, msg); err != nil {
		return
	}

	fmt.Println("File read")
	return
}

func writeToFile(filename string, msg proto.Message) (err error) {
	bv, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(filename, bv, 0644); err != nil {
		return err
	}

	fmt.Println("File written")

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	msg := simplepb.SimpleMessage{
		Id:         2,
		IsSimple:   true,
		Name:       "Myself",
		SampleList: []int32{1, 2, 3},
	}

	fmt.Println("Simple message")
	fmt.Println(msg.String())

	msg.Name = "I renamed you"
	fmt.Println(msg.String())

	fmt.Printf("The ID is: %v\n", msg.GetId())

	return &msg
}
