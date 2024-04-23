package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fullgukbap/protobuf/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	// p := pb.Person{
	// 	Id:    1234,
	// 	Name:  "Jone Dae",
	// 	Email: "jode@example.com",
	// 	Phones: []*pb.Person_PhoneNumber{
	// 		{Number: "555-4251", Type: pb.PhoneType_PHONE_TYPE_MOBILE},
	// 	},
	// }

	// out, err := proto.Marshal(&p)
	// if err != nil {
	// 	log.Fatalln("Failed to encode address book:", err)
	// }

	// if err := os.WriteFile("res", out, 0644); err != nil {
	// 	log.Fatalln("Failed to write address book:", err)
	// }

	// Read the existing address book.
	in, err := os.ReadFile("res")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	inPerson := &pb.Person{}
	if err := proto.Unmarshal(in, inPerson); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println("inPerson: ", inPerson)
}
