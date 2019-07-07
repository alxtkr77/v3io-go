// Code generated by capnpc-go. DO NOT EDIT.

package vncapnp

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type VnObjectAttributeKeyMap struct{ capnp.Struct }

// VnObjectAttributeKeyMap_TypeID is the unique identifier for the type VnObjectAttributeKeyMap.
const VnObjectAttributeKeyMap_TypeID = 0x8f34d3b1223bf828

func NewVnObjectAttributeKeyMap(s *capnp.Segment) (VnObjectAttributeKeyMap, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return VnObjectAttributeKeyMap{st}, err
}

func NewRootVnObjectAttributeKeyMap(s *capnp.Segment) (VnObjectAttributeKeyMap, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return VnObjectAttributeKeyMap{st}, err
}

func ReadRootVnObjectAttributeKeyMap(msg *capnp.Message) (VnObjectAttributeKeyMap, error) {
	root, err := msg.RootPtr()
	return VnObjectAttributeKeyMap{root.Struct()}, err
}

func (s VnObjectAttributeKeyMap) String() string {
	str, _ := text.Marshal(0x8f34d3b1223bf828, s.Struct)
	return str
}

func (s VnObjectAttributeKeyMap) Names() (StringWrapperList, error) {
	p, err := s.Struct.Ptr(0)
	return StringWrapperList{Struct: p.Struct()}, err
}

func (s VnObjectAttributeKeyMap) HasNames() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VnObjectAttributeKeyMap) SetNames(v StringWrapperList) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewNames sets the names field to a newly
// allocated StringWrapperList struct, preferring placement in s's segment.
func (s VnObjectAttributeKeyMap) NewNames() (StringWrapperList, error) {
	ss, err := NewStringWrapperList(s.Struct.Segment())
	if err != nil {
		return StringWrapperList{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// VnObjectAttributeKeyMap_List is a list of VnObjectAttributeKeyMap.
type VnObjectAttributeKeyMap_List struct{ capnp.List }

// NewVnObjectAttributeKeyMap creates a new list of VnObjectAttributeKeyMap.
func NewVnObjectAttributeKeyMap_List(s *capnp.Segment, sz int32) (VnObjectAttributeKeyMap_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return VnObjectAttributeKeyMap_List{l}, err
}

func (s VnObjectAttributeKeyMap_List) At(i int) VnObjectAttributeKeyMap {
	return VnObjectAttributeKeyMap{s.List.Struct(i)}
}

func (s VnObjectAttributeKeyMap_List) Set(i int, v VnObjectAttributeKeyMap) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s VnObjectAttributeKeyMap_List) String() string {
	str, _ := text.MarshalList(0x8f34d3b1223bf828, s.List)
	return str
}

// VnObjectAttributeKeyMap_Promise is a wrapper for a VnObjectAttributeKeyMap promised by a client call.
type VnObjectAttributeKeyMap_Promise struct{ *capnp.Pipeline }

func (p VnObjectAttributeKeyMap_Promise) Struct() (VnObjectAttributeKeyMap, error) {
	s, err := p.Pipeline.Struct()
	return VnObjectAttributeKeyMap{s}, err
}

func (p VnObjectAttributeKeyMap_Promise) Names() StringWrapperList_Promise {
	return StringWrapperList_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

const schema_986bf57944c8b89f = "x\xda\x12\x98\xee\xc0d\xc8z\x9e\x99\x81!P\x87\x95" +
	"\xed\xbf\xc6\x0fk\xa5\x8d\x97M\xfa\x19\x04\x0d\x18\xff\xcf" +
	"\xdfq\xc2\xa5\xf2k\xf6\x0c\x06VFv\x06\x06\xe3\xb3" +
	"\x8cNL\xc2oAL\xe1\x97\x8c\xf6\x0c\xf9\xff\x93\x13" +
	"\x0b\xf2\x0a\xf4\xf3\xf29RR\xf5\x93\xf3ss\xf3\xf3" +
	"\xf4\xc3\xf2\xfc\x93\xb2R\x93K\x1cKJ\x8a2\x93J" +
	"KR\xbdS+}\x13\x0b\xf4\xc0J\xadp\xc82\x04" +
	"02\x06\xb20\xb300\xb0020\x08\xf2\x1a1" +
	"0\x04r03\x06\x8a01\xca\xe7%\xe6\xa6\x163" +
	"\x0a\xfc\xdf\x10\xc1\xd1{\xdce\xcd\x1b\x06\x06FF\x01" +
	"\x06F@\x00\x00\x00\xff\xff\x08<70"

func init() {
	schemas.Register(schema_986bf57944c8b89f,
		0x8f34d3b1223bf828)
}
