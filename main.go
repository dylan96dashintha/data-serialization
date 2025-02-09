package main

import (
	"fmt"
	avroObj "github.com/data-serialization/avro"
	"github.com/data-serialization/protobuf"
	"google.golang.org/protobuf/proto"
)

var schema = `{
	  "type": "record",
	  "name": "Person",
	  "fields": [
	    { "name": "name", "type": "string" },
	    { "name": "age", "type": "int" },
		 { 
      "name": "address", 
      "type": {
        "type": "array",
        "items": {
          "type": "record",
          "name": "Address",
          "fields": [
            { "name": "street", "type": "string" },
            { "name": "city", "type": "string" },
            { "name": "province", "type": "string" },
            { "name": "postalCode", "type": "int" }
          ]
        }
      }
    }
	]
	}`

func main() {

	// avro related serialization and deserialization
	avro, err := avroObj.NewAvro(schema)
	if err != nil {
		panic(err)
	}

	person := map[string]interface{}{
		"name": "Dilan",
		"age":  28,
		"address": []interface{}{
			map[string]interface{}{
				"street":     "Mampitiya",
				"city":       "Galle",
				"province":   "Southern",
				"postalCode": 80000,
			},
			map[string]interface{}{
				"street":     "ranpokuna",
				"city":       "Pitakotte",
				"province":   "Colombo",
				"postalCode": 10100,
			},
		},
	}

	// serialization
	record, err := avro.Serializer(person)
	if err != nil {
		panic(err)
	}
	// deserialization
	avro.Deserializer(record)

	// protobuff related serialization and deserialization
	ads := &protobuf.Address{
		Street:     "Mampitiya",
		City:       "Galle",
		Province:   "Southern",
		PostalCode: 80000,
	}

	secondAds := &protobuf.Address{
		Street:     "ranpokuna",
		City:       "Pitakotte",
		Province:   "Colombo",
		PostalCode: 10100,
	}
	personProtoBuf := protobuf.Person{
		Name:    "Dilan",
		Age:     28,
		Address: []*protobuf.Address{ads, secondAds},
	}
	// Serialize to binary
	data, err := proto.Marshal(&personProtoBuf)
	if err != nil {
		panic(err)
	}
	fmt.Println("Serialized ProtoBuf Data:", data)

	// Deserialize from binary
	newPerson := &protobuf.Person{}
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deserialized ProtoBuf Data:", newPerson)
}
