// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.9.0-dev
//  VPP:              22.10.1-1~g1b93285ce
// source: plugins/cup_handoff.api.json

// Package cup_handoff contains generated bindings for API file cup_handoff.api.
//
// Contents:
// -  4 messages
package cup_handoff

import (
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "cup_handoff"
	APIVersion = "0.0.1"
	VersionCrc = 0x148af2d2
)

// CupSetOutputs defines message 'cup_set_outputs'.
type CupSetOutputs struct {
	Count   uint32   `binapi:"u32,name=count" json:"-"`
	Outputs []uint32 `binapi:"u32[count],name=outputs" json:"outputs,omitempty"`
}

func (m *CupSetOutputs) Reset()               { *m = CupSetOutputs{} }
func (*CupSetOutputs) GetMessageName() string { return "cup_set_outputs" }
func (*CupSetOutputs) GetCrcString() string   { return "6964ac78" }
func (*CupSetOutputs) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CupSetOutputs) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                  // m.Count
	size += 4 * len(m.Outputs) // m.Outputs
	return size
}
func (m *CupSetOutputs) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(len(m.Outputs)))
	for i := 0; i < len(m.Outputs); i++ {
		var x uint32
		if i < len(m.Outputs) {
			x = uint32(m.Outputs[i])
		}
		buf.EncodeUint32(x)
	}
	return buf.Bytes(), nil
}
func (m *CupSetOutputs) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Count = buf.DecodeUint32()
	m.Outputs = make([]uint32, m.Count)
	for i := 0; i < len(m.Outputs); i++ {
		m.Outputs[i] = buf.DecodeUint32()
	}
	return nil
}

// CupSetOutputsReply defines message 'cup_set_outputs_reply'.
type CupSetOutputsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *CupSetOutputsReply) Reset()               { *m = CupSetOutputsReply{} }
func (*CupSetOutputsReply) GetMessageName() string { return "cup_set_outputs_reply" }
func (*CupSetOutputsReply) GetCrcString() string   { return "e8d4e804" }
func (*CupSetOutputsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CupSetOutputsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *CupSetOutputsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *CupSetOutputsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// CupSetWorkers defines message 'cup_set_workers'.
type CupSetWorkers struct {
	Count   uint32   `binapi:"u32,name=count" json:"-"`
	Workers []uint32 `binapi:"u32[count],name=workers" json:"workers,omitempty"`
}

func (m *CupSetWorkers) Reset()               { *m = CupSetWorkers{} }
func (*CupSetWorkers) GetMessageName() string { return "cup_set_workers" }
func (*CupSetWorkers) GetCrcString() string   { return "47030c3d" }
func (*CupSetWorkers) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CupSetWorkers) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                  // m.Count
	size += 4 * len(m.Workers) // m.Workers
	return size
}
func (m *CupSetWorkers) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(len(m.Workers)))
	for i := 0; i < len(m.Workers); i++ {
		var x uint32
		if i < len(m.Workers) {
			x = uint32(m.Workers[i])
		}
		buf.EncodeUint32(x)
	}
	return buf.Bytes(), nil
}
func (m *CupSetWorkers) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Count = buf.DecodeUint32()
	m.Workers = make([]uint32, m.Count)
	for i := 0; i < len(m.Workers); i++ {
		m.Workers[i] = buf.DecodeUint32()
	}
	return nil
}

// CupSetWorkersReply defines message 'cup_set_workers_reply'.
type CupSetWorkersReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *CupSetWorkersReply) Reset()               { *m = CupSetWorkersReply{} }
func (*CupSetWorkersReply) GetMessageName() string { return "cup_set_workers_reply" }
func (*CupSetWorkersReply) GetCrcString() string   { return "e8d4e804" }
func (*CupSetWorkersReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CupSetWorkersReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *CupSetWorkersReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *CupSetWorkersReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_cup_handoff_binapi_init() }
func file_cup_handoff_binapi_init() {
	api.RegisterMessage((*CupSetOutputs)(nil), "cup_set_outputs_6964ac78")
	api.RegisterMessage((*CupSetOutputsReply)(nil), "cup_set_outputs_reply_e8d4e804")
	api.RegisterMessage((*CupSetWorkers)(nil), "cup_set_workers_47030c3d")
	api.RegisterMessage((*CupSetWorkersReply)(nil), "cup_set_workers_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*CupSetOutputs)(nil),
		(*CupSetOutputsReply)(nil),
		(*CupSetWorkers)(nil),
		(*CupSetWorkersReply)(nil),
	}
}
