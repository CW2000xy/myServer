// Code generated by protoc-gen-go.
// source: console_broadcast.proto
// DO NOT EDIT!

package protocol

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type BroadcastToSingleUser struct {
	Cmd              *int32  `protobuf:"varint,1,opt" json:"Cmd,omitempty"`
	Targets          []int32 `protobuf:"varint,2,rep,name=targets" json:"targets,omitempty"`
	PackBinary       []byte  `protobuf:"bytes,3,opt" json:"PackBinary,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BroadcastToSingleUser) Reset()         { *m = BroadcastToSingleUser{} }
func (m *BroadcastToSingleUser) String() string { return proto.CompactTextString(m) }
func (*BroadcastToSingleUser) ProtoMessage()    {}

func (m *BroadcastToSingleUser) GetCmd() int32 {
	if m != nil && m.Cmd != nil {
		return *m.Cmd
	}
	return 0
}

func (m *BroadcastToSingleUser) GetTargets() []int32 {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *BroadcastToSingleUser) GetPackBinary() []byte {
	if m != nil {
		return m.PackBinary
	}
	return nil
}

func init() {
}