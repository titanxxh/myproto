package api

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"reflect"
)

type Producer func() proto.Message

type TypeMap struct {
	producers     map[string]Producer
	protoTypes    map[string]reflect.Type
	revProtoTypes map[reflect.Type]string
}

func (m *TypeMap) GetProducer(t string) Producer {
	if p, ok := m.producers[t]; ok {
		return p
	} else {
		return nil
	}
}

func (m *TypeMap) registerType(x proto.Message, name string) {
	if _, ok := m.protoTypes[name]; ok {
		panic(fmt.Sprintf("proto: duplicate proto type registered: %s", name))
		return
	}
	t := reflect.TypeOf(x)
	m.protoTypes[name] = t
	m.revProtoTypes[t] = name
	m.producers[name] = func() proto.Message {
		return reflect.New(t.Elem()).Interface().(proto.Message)
	}
}

func TypeMapSingleton() *TypeMap {
	return &TypeMap{
		producers:     make(map[string]Producer),
		protoTypes:    make(map[string]reflect.Type),
		revProtoTypes: make(map[reflect.Type]string),
	}
}

var GlobalTypeMap = TypeMapSingleton()

func init() {
	GlobalTypeMap.registerType((*ProtoBufA)(nil), "api.ProtoBufA")
}
