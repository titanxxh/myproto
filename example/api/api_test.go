package api

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeMap_GetProducer(t *testing.T) {
	p := GlobalTypeMap.GetProducer("api.ProtoBufA")
	origin := &ProtoBufA{
		Name:     "hello world",
		BirthDay: 20170925,
		Phone:    "secret",
		Siblings: 1,
		Spouse:   true,
		Money:    0.37,
	}
	buf, _ := proto.Marshal(origin)
	protoMsg := p()
	proto.Unmarshal(buf, protoMsg)
	fmt.Printf("%+v\n", protoMsg)
	assert.Equal(t, origin, protoMsg)
}

func BenchmarkTypeMap_GetProducer_UseProducer(b *testing.B) {
	p := GlobalTypeMap.GetProducer("api.ProtoBufA")
	origin := &ProtoBufA{
		Name:     "hello world",
		BirthDay: 20170925,
		Phone:    "secret",
		Siblings: 1,
		Spouse:   true,
		Money:    0.37,
	}
	buf, _ := proto.Marshal(origin)
	for i := 0; i < b.N; i++ {
		protoMsg := p()
		proto.Unmarshal(buf, protoMsg)
	}
}

func BenchmarkTypeMap_GetProducer_UseStruct(b *testing.B) {
	origin := &ProtoBufA{
		Name:     "hello world",
		BirthDay: 20170925,
		Phone:    "secret",
		Siblings: 1,
		Spouse:   true,
		Money:    0.37,
	}
	buf, _ := proto.Marshal(origin)
	for i := 0; i < b.N; i++ {
		protoMsg := &ProtoBufA{}
		proto.Unmarshal(buf, protoMsg)
	}
}
