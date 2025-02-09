package avro

import (
	"fmt"
	"github.com/linkedin/goavro/v2"
)

type avro struct {
	schema string
	codec  *goavro.Codec
}

type Avro interface {
	Serializer(data interface{}) ([]byte, error)
	Deserializer(record []byte)
}

func NewAvro(schema string) (Avro, error) {

	codec, err := goavro.NewCodec(schema)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("error in initializing codec, err : %s", err))
		return nil, err
	}
	avroObj := avro{
		codec: codec,
	}
	return &avroObj, nil
}

func (a avro) Serializer(data interface{}) ([]byte, error) {

	record, err := a.codec.BinaryFromNative(nil, data)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("error in serialization in avro, err : %s", err))
		return nil, err
	}

	fmt.Println(fmt.Sprintf("serialized record : %s ", record))
	return record, nil
}

func (a avro) Deserializer(record []byte) {
	native, _, err := a.codec.NativeFromBinary(record)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("error in deserialization in avro, err : %s", err))
		return
	}

	fmt.Println(fmt.Sprintf("deserialized record : %s ", native))
	return
}
