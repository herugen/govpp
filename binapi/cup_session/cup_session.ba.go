// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.9.0-dev
//  VPP:              22.10.1-1~g1b93285ce
// source: plugins/cup_session.api.json

// Package cup_session contains generated bindings for API file cup_session.api.
//
// Contents:
// -  2 enums
// - 12 structs
// -  1 union
// -  6 messages
package cup_session

import (
	"strconv"

	api "go.fd.io/govpp/api"
	_ "go.fd.io/govpp/binapi/acl_types"
	cup_event "go.fd.io/govpp/binapi/cup_event"
	cup_ipfilter "go.fd.io/govpp/binapi/cup_ipfilter"
	_ "go.fd.io/govpp/binapi/ethernet_types"
	_ "go.fd.io/govpp/binapi/interface_types"
	ip_types "go.fd.io/govpp/binapi/ip_types"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "cup_session"
	APIVersion = "0.0.1"
	VersionCrc = 0x5b704b54
)

// Exhausted defines enum 'exhausted'.
type Exhausted uint8

const (
	CUP_EX_ACT_NONE Exhausted = 0
	CUP_EX_ACT_DROP Exhausted = 1
)

var (
	Exhausted_name = map[uint8]string{
		0: "CUP_EX_ACT_NONE",
		1: "CUP_EX_ACT_DROP",
	}
	Exhausted_value = map[string]uint8{
		"CUP_EX_ACT_NONE": 0,
		"CUP_EX_ACT_DROP": 1,
	}
)

func (x Exhausted) String() string {
	s, ok := Exhausted_name[uint8(x)]
	if ok {
		return s
	}
	return "Exhausted(" + strconv.Itoa(int(x)) + ")"
}

// PolType defines enum 'pol_type'.
type PolType uint8

const (
	CUP_POL_NONE       PolType = 0
	CUP_POL_EXT_FILTER PolType = 1
	CUP_POL_CM_FILTER  PolType = 2
	CUP_POL_GTPU_ENCAP PolType = 3
	CUP_POL_GTPU_DECAP PolType = 4
	CUP_POL_ROUTE      PolType = 5
	CUP_POL_ADV_ROUTE  PolType = 6
	CUP_POL_DROP       PolType = 7
	CUP_POL_SHAPING    PolType = 8
	CUP_POL_MARK       PolType = 9
	CUP_POL_MEASURE    PolType = 10
	CUP_POL_HOLD       PolType = 11
	CUP_POL_PIPELINE   PolType = 12
)

var (
	PolType_name = map[uint8]string{
		0:  "CUP_POL_NONE",
		1:  "CUP_POL_EXT_FILTER",
		2:  "CUP_POL_CM_FILTER",
		3:  "CUP_POL_GTPU_ENCAP",
		4:  "CUP_POL_GTPU_DECAP",
		5:  "CUP_POL_ROUTE",
		6:  "CUP_POL_ADV_ROUTE",
		7:  "CUP_POL_DROP",
		8:  "CUP_POL_SHAPING",
		9:  "CUP_POL_MARK",
		10: "CUP_POL_MEASURE",
		11: "CUP_POL_HOLD",
		12: "CUP_POL_PIPELINE",
	}
	PolType_value = map[string]uint8{
		"CUP_POL_NONE":       0,
		"CUP_POL_EXT_FILTER": 1,
		"CUP_POL_CM_FILTER":  2,
		"CUP_POL_GTPU_ENCAP": 3,
		"CUP_POL_GTPU_DECAP": 4,
		"CUP_POL_ROUTE":      5,
		"CUP_POL_ADV_ROUTE":  6,
		"CUP_POL_DROP":       7,
		"CUP_POL_SHAPING":    8,
		"CUP_POL_MARK":       9,
		"CUP_POL_MEASURE":    10,
		"CUP_POL_HOLD":       11,
		"CUP_POL_PIPELINE":   12,
	}
)

func (x PolType) String() string {
	s, ok := PolType_name[uint8(x)]
	if ok {
		return s
	}
	return "PolType(" + strconv.Itoa(int(x)) + ")"
}

// AdvRoute defines type 'adv_route'.
type AdvRoute struct {
	ID    uint16 `binapi:"u16,name=id" json:"id,omitempty"`
	VrfID uint32 `binapi:"u32,name=vrf_id" json:"vrf_id,omitempty"`
}

// CmFilter defines type 'cm_filter'.
type CmFilter struct {
	ID         uint16                `binapi:"u16,name=id" json:"id,omitempty"`
	Ipf        cup_ipfilter.IPFilter `binapi:"ip_filter,name=ipf" json:"ipf,omitempty"`
	Precedence uint32                `binapi:"u32,name=precedence" json:"precedence,omitempty"`
	MatchTeid  bool                  `binapi:"bool,name=match_teid" json:"match_teid,omitempty"`
	Teid       uint32                `binapi:"u32,name=teid" json:"teid,omitempty"`
	MatchQfi   bool                  `binapi:"bool,name=match_qfi" json:"match_qfi,omitempty"`
	Qfi        uint8                 `binapi:"u8,name=qfi" json:"qfi,omitempty"`
	PID        uint16                `binapi:"u16,name=pid" json:"pid,omitempty"`
}

// ExtFilter defines type 'ext_filter'.
type ExtFilter struct {
	ID         uint16 `binapi:"u16,name=id" json:"id,omitempty"`
	LinkID     uint64 `binapi:"u64,name=link_id" json:"link_id,omitempty"`
	Precedence uint32 `binapi:"u32,name=precedence" json:"precedence,omitempty"`
	MatchTeid  bool   `binapi:"bool,name=match_teid" json:"match_teid,omitempty"`
	Teid       uint32 `binapi:"u32,name=teid" json:"teid,omitempty"`
	MatchQfi   bool   `binapi:"bool,name=match_qfi" json:"match_qfi,omitempty"`
	Qfi        uint8  `binapi:"u8,name=qfi" json:"qfi,omitempty"`
	PID        uint16 `binapi:"u16,name=pid" json:"pid,omitempty"`
}

// GtpuDecap defines type 'gtpu_decap'.
type GtpuDecap struct {
	ID         uint16           `binapi:"u16,name=id" json:"id,omitempty"`
	DstAddress ip_types.Address `binapi:"address,name=dst_address" json:"dst_address,omitempty"`
	VrfID      uint32           `binapi:"u32,name=vrf_id" json:"vrf_id,omitempty"`
	Teid       uint32           `binapi:"u32,name=teid" json:"teid,omitempty"`
}

// GtpuEncap defines type 'gtpu_encap'.
type GtpuEncap struct {
	ID         uint16           `binapi:"u16,name=id" json:"id,omitempty"`
	SrcAddress ip_types.Address `binapi:"address,name=src_address" json:"src_address,omitempty"`
	DstAddress ip_types.Address `binapi:"address,name=dst_address" json:"dst_address,omitempty"`
	SrcPort    uint16           `binapi:"u16,name=src_port" json:"src_port,omitempty"`
	DstPort    uint16           `binapi:"u16,name=dst_port" json:"dst_port,omitempty"`
	VrfID      uint32           `binapi:"u32,name=vrf_id" json:"vrf_id,omitempty"`
	Teid       uint32           `binapi:"u32,name=teid" json:"teid,omitempty"`
	Qfi        uint8            `binapi:"u8,name=qfi" json:"qfi,omitempty"`
	Dscp       uint8            `binapi:"u8,name=dscp" json:"dscp,omitempty"`
	Ppi        uint8            `binapi:"u8,name=ppi" json:"ppi,omitempty"`
	Rqi        uint8            `binapi:"u8,name=rqi" json:"rqi,omitempty"`
}

// Hold defines type 'hold'.
type Hold struct {
	ID  uint16 `binapi:"u16,name=id" json:"id,omitempty"`
	HDl uint8  `binapi:"u8,name=h_dl" json:"h_dl,omitempty"`
}

// Mark defines type 'mark'.
type Mark struct {
	ID   uint16 `binapi:"u16,name=id" json:"id,omitempty"`
	Dscp uint8  `binapi:"u8,name=dscp" json:"dscp,omitempty"`
}

// Measure defines type 'measure'.
type Measure struct {
	ID        uint16    `binapi:"u16,name=id" json:"id,omitempty"`
	LinkedIds []uint16  `binapi:"u16[3],name=linked_ids" json:"linked_ids,omitempty"`
	Act       Exhausted `binapi:"exhausted,name=act" json:"act,omitempty"`
	Thresh    uint64    `binapi:"u64,name=thresh" json:"thresh,omitempty"`
	Quota     uint64    `binapi:"u64,name=quota" json:"quota,omitempty"`
}

// Pipeline defines type 'pipeline'.
type Pipeline struct {
	ID    uint16      `binapi:"u16,name=id" json:"id,omitempty"`
	Ptype [16]PolType `binapi:"pol_type[16],name=ptype" json:"ptype,omitempty"`
	PID   []uint16    `binapi:"u16[16],name=pid" json:"pid,omitempty"`
}

// Policy defines type 'policy'.
type Policy struct {
	Type  PolType       `binapi:"pol_type,name=type" json:"type,omitempty"`
	Value PolValueUnion `binapi:"pol_value,name=value" json:"value,omitempty"`
}

// SessionPrefix defines type 'session_prefix'.
type SessionPrefix struct {
	VrfID  uint32          `binapi:"u32,name=vrf_id" json:"vrf_id,omitempty"`
	Prefix ip_types.Prefix `binapi:"prefix,name=prefix" json:"prefix,omitempty"`
}

// Shaping defines type 'shaping'.
type Shaping struct {
	ID  uint16 `binapi:"u16,name=id" json:"id,omitempty"`
	Mbr uint64 `binapi:"u64,name=mbr" json:"mbr,omitempty"`
}

// PolValueUnion defines union 'pol_value'.
type PolValueUnion struct {
	// PolValueUnion can be one of:
	// - ExtFilter *ExtFilter
	// - CmFilter *CmFilter
	// - GtpuEncap *GtpuEncap
	// - GtpuDecap *GtpuDecap
	// - AdvRoute *AdvRoute
	// - Shaping *Shaping
	// - Mark *Mark
	// - Measure *Measure
	// - Hold *Hold
	// - Pipeline *Pipeline
	XXX_UnionData [68]byte
}

func PolValueUnionExtFilter(a ExtFilter) (u PolValueUnion) {
	u.SetExtFilter(a)
	return
}
func (u *PolValueUnion) SetExtFilter(a ExtFilter) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint16(a.ID)
	buf.EncodeUint64(a.LinkID)
	buf.EncodeUint32(a.Precedence)
	buf.EncodeBool(a.MatchTeid)
	buf.EncodeUint32(a.Teid)
	buf.EncodeBool(a.MatchQfi)
	buf.EncodeUint8(a.Qfi)
	buf.EncodeUint16(a.PID)
}
func (u *PolValueUnion) GetExtFilter() (a ExtFilter) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.ID = buf.DecodeUint16()
	a.LinkID = buf.DecodeUint64()
	a.Precedence = buf.DecodeUint32()
	a.MatchTeid = buf.DecodeBool()
	a.Teid = buf.DecodeUint32()
	a.MatchQfi = buf.DecodeBool()
	a.Qfi = buf.DecodeUint8()
	a.PID = buf.DecodeUint16()
	return
}

func PolValueUnionCmFilter(a CmFilter) (u PolValueUnion) {
	u.SetCmFilter(a)
	return
}
func (u *PolValueUnion) SetCmFilter(a CmFilter) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint16(a.ID)
	buf.EncodeUint64(a.Ipf.ID)
	buf.EncodeUint8(a.Ipf.Proto)
	buf.EncodeUint8(uint8(a.Ipf.SrcPrefix.Address.Af))
	buf.EncodeBytes(a.Ipf.SrcPrefix.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(a.Ipf.SrcPrefix.Len)
	buf.EncodeUint8(uint8(a.Ipf.DstPrefix.Address.Af))
	buf.EncodeBytes(a.Ipf.DstPrefix.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(a.Ipf.DstPrefix.Len)
	buf.EncodeUint16(a.Ipf.SrcportFirst)
	buf.EncodeUint16(a.Ipf.SrcportLast)
	buf.EncodeUint16(a.Ipf.DstportFirst)
	buf.EncodeUint16(a.Ipf.DstportLast)
	buf.EncodeUint32(a.Precedence)
	buf.EncodeBool(a.MatchTeid)
	buf.EncodeUint32(a.Teid)
	buf.EncodeBool(a.MatchQfi)
	buf.EncodeUint8(a.Qfi)
	buf.EncodeUint16(a.PID)
}
func (u *PolValueUnion) GetCmFilter() (a CmFilter) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.ID = buf.DecodeUint16()
	a.Ipf.ID = buf.DecodeUint64()
	a.Ipf.Proto = buf.DecodeUint8()
	a.Ipf.SrcPrefix.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(a.Ipf.SrcPrefix.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	a.Ipf.SrcPrefix.Len = buf.DecodeUint8()
	a.Ipf.DstPrefix.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(a.Ipf.DstPrefix.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	a.Ipf.DstPrefix.Len = buf.DecodeUint8()
	a.Ipf.SrcportFirst = buf.DecodeUint16()
	a.Ipf.SrcportLast = buf.DecodeUint16()
	a.Ipf.DstportFirst = buf.DecodeUint16()
	a.Ipf.DstportLast = buf.DecodeUint16()
	a.Precedence = buf.DecodeUint32()
	a.MatchTeid = buf.DecodeBool()
	a.Teid = buf.DecodeUint32()
	a.MatchQfi = buf.DecodeBool()
	a.Qfi = buf.DecodeUint8()
	a.PID = buf.DecodeUint16()
	return
}

func PolValueUnionGtpuEncap(a GtpuEncap) (u PolValueUnion) {
	u.SetGtpuEncap(a)
	return
}
func (u *PolValueUnion) SetGtpuEncap(a GtpuEncap) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint16(a.ID)
	buf.EncodeUint8(uint8(a.SrcAddress.Af))
	buf.EncodeBytes(a.SrcAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(a.DstAddress.Af))
	buf.EncodeBytes(a.DstAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint16(a.SrcPort)
	buf.EncodeUint16(a.DstPort)
	buf.EncodeUint32(a.VrfID)
	buf.EncodeUint32(a.Teid)
	buf.EncodeUint8(a.Qfi)
	buf.EncodeUint8(a.Dscp)
	buf.EncodeUint8(a.Ppi)
	buf.EncodeUint8(a.Rqi)
}
func (u *PolValueUnion) GetGtpuEncap() (a GtpuEncap) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.ID = buf.DecodeUint16()
	a.SrcAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(a.SrcAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	a.DstAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(a.DstAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	a.SrcPort = buf.DecodeUint16()
	a.DstPort = buf.DecodeUint16()
	a.VrfID = buf.DecodeUint32()
	a.Teid = buf.DecodeUint32()
	a.Qfi = buf.DecodeUint8()
	a.Dscp = buf.DecodeUint8()
	a.Ppi = buf.DecodeUint8()
	a.Rqi = buf.DecodeUint8()
	return
}

func PolValueUnionGtpuDecap(a GtpuDecap) (u PolValueUnion) {
	u.SetGtpuDecap(a)
	return
}
func (u *PolValueUnion) SetGtpuDecap(a GtpuDecap) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint16(a.ID)
	buf.EncodeUint8(uint8(a.DstAddress.Af))
	buf.EncodeBytes(a.DstAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(a.VrfID)
	buf.EncodeUint32(a.Teid)
}
func (u *PolValueUnion) GetGtpuDecap() (a GtpuDecap) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.ID = buf.DecodeUint16()
	a.DstAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(a.DstAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	a.VrfID = buf.DecodeUint32()
	a.Teid = buf.DecodeUint32()
	return
}

func PolValueUnionAdvRoute(a AdvRoute) (u PolValueUnion) {
	u.SetAdvRoute(a)
	return
}
func (u *PolValueUnion) SetAdvRoute(a AdvRoute) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint16(a.ID)
	buf.EncodeUint32(a.VrfID)
}
func (u *PolValueUnion) GetAdvRoute() (a AdvRoute) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.ID = buf.DecodeUint16()
	a.VrfID = buf.DecodeUint32()
	return
}

func PolValueUnionShaping(a Shaping) (u PolValueUnion) {
	u.SetShaping(a)
	return
}
func (u *PolValueUnion) SetShaping(a Shaping) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint16(a.ID)
	buf.EncodeUint64(a.Mbr)
}
func (u *PolValueUnion) GetShaping() (a Shaping) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.ID = buf.DecodeUint16()
	a.Mbr = buf.DecodeUint64()
	return
}

func PolValueUnionMark(a Mark) (u PolValueUnion) {
	u.SetMark(a)
	return
}
func (u *PolValueUnion) SetMark(a Mark) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint16(a.ID)
	buf.EncodeUint8(a.Dscp)
}
func (u *PolValueUnion) GetMark() (a Mark) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.ID = buf.DecodeUint16()
	a.Dscp = buf.DecodeUint8()
	return
}

func PolValueUnionMeasure(a Measure) (u PolValueUnion) {
	u.SetMeasure(a)
	return
}
func (u *PolValueUnion) SetMeasure(a Measure) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint16(a.ID)
	for i := 0; i < 3; i++ {
		var x uint16
		if i < len(a.LinkedIds) {
			x = uint16(a.LinkedIds[i])
		}
		buf.EncodeUint16(x)
	}
	buf.EncodeUint8(uint8(a.Act))
	buf.EncodeUint64(a.Thresh)
	buf.EncodeUint64(a.Quota)
}
func (u *PolValueUnion) GetMeasure() (a Measure) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.ID = buf.DecodeUint16()
	a.LinkedIds = make([]uint16, 3)
	for i := 0; i < len(a.LinkedIds); i++ {
		a.LinkedIds[i] = buf.DecodeUint16()
	}
	a.Act = Exhausted(buf.DecodeUint8())
	a.Thresh = buf.DecodeUint64()
	a.Quota = buf.DecodeUint64()
	return
}

func PolValueUnionHold(a Hold) (u PolValueUnion) {
	u.SetHold(a)
	return
}
func (u *PolValueUnion) SetHold(a Hold) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint16(a.ID)
	buf.EncodeUint8(a.HDl)
}
func (u *PolValueUnion) GetHold() (a Hold) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.ID = buf.DecodeUint16()
	a.HDl = buf.DecodeUint8()
	return
}

func PolValueUnionPipeline(a Pipeline) (u PolValueUnion) {
	u.SetPipeline(a)
	return
}
func (u *PolValueUnion) SetPipeline(a Pipeline) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint16(a.ID)
	for j1 := 0; j1 < 16; j1++ {
		buf.EncodeUint8(uint8(a.Ptype[j1]))
	}
	for i := 0; i < 16; i++ {
		var x uint16
		if i < len(a.PID) {
			x = uint16(a.PID[i])
		}
		buf.EncodeUint16(x)
	}
}
func (u *PolValueUnion) GetPipeline() (a Pipeline) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.ID = buf.DecodeUint16()
	for j1 := 0; j1 < 16; j1++ {
		a.Ptype[j1] = PolType(buf.DecodeUint8())
	}
	a.PID = make([]uint16, 16)
	for i := 0; i < len(a.PID); i++ {
		a.PID[i] = buf.DecodeUint16()
	}
	return
}

// CupAddDelSession defines message 'cup_add_del_session'.
type CupAddDelSession struct {
	IsAdd    bool             `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	ID       uint64           `binapi:"u64,name=id" json:"id,omitempty"`
	Prefix   [2]SessionPrefix `binapi:"session_prefix[2],name=prefix" json:"prefix,omitempty"`
	Count    uint8            `binapi:"u8,name=count" json:"-"`
	Policies []Policy         `binapi:"policy[count],name=policies" json:"policies,omitempty"`
}

func (m *CupAddDelSession) Reset()               { *m = CupAddDelSession{} }
func (*CupAddDelSession) GetMessageName() string { return "cup_add_del_session" }
func (*CupAddDelSession) GetCrcString() string   { return "27815055" }
func (*CupAddDelSession) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CupAddDelSession) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 8 // m.ID
	for j1 := 0; j1 < 2; j1++ {
		size += 4      // m.Prefix[j1].VrfID
		size += 1      // m.Prefix[j1].Prefix.Address.Af
		size += 1 * 16 // m.Prefix[j1].Prefix.Address.Un
		size += 1      // m.Prefix[j1].Prefix.Len
	}
	size += 1 // m.Count
	for j1 := 0; j1 < len(m.Policies); j1++ {
		var s1 Policy
		_ = s1
		if j1 < len(m.Policies) {
			s1 = m.Policies[j1]
		}
		size += 1      // s1.Type
		size += 1 * 68 // s1.Value
	}
	return size
}
func (m *CupAddDelSession) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint64(m.ID)
	for j0 := 0; j0 < 2; j0++ {
		buf.EncodeUint32(m.Prefix[j0].VrfID)
		buf.EncodeUint8(uint8(m.Prefix[j0].Prefix.Address.Af))
		buf.EncodeBytes(m.Prefix[j0].Prefix.Address.Un.XXX_UnionData[:], 16)
		buf.EncodeUint8(m.Prefix[j0].Prefix.Len)
	}
	buf.EncodeUint8(uint8(len(m.Policies)))
	for j0 := 0; j0 < len(m.Policies); j0++ {
		var v0 Policy // Policies
		if j0 < len(m.Policies) {
			v0 = m.Policies[j0]
		}
		buf.EncodeUint8(uint8(v0.Type))
		buf.EncodeBytes(v0.Value.XXX_UnionData[:], 68)
	}
	return buf.Bytes(), nil
}
func (m *CupAddDelSession) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.ID = buf.DecodeUint64()
	for j0 := 0; j0 < 2; j0++ {
		m.Prefix[j0].VrfID = buf.DecodeUint32()
		m.Prefix[j0].Prefix.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
		copy(m.Prefix[j0].Prefix.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
		m.Prefix[j0].Prefix.Len = buf.DecodeUint8()
	}
	m.Count = buf.DecodeUint8()
	m.Policies = make([]Policy, m.Count)
	for j0 := 0; j0 < len(m.Policies); j0++ {
		m.Policies[j0].Type = PolType(buf.DecodeUint8())
		copy(m.Policies[j0].Value.XXX_UnionData[:], buf.DecodeBytes(68))
	}
	return nil
}

// CupAddDelSessionReply defines message 'cup_add_del_session_reply'.
type CupAddDelSessionReply struct {
	Retval      int32                   `binapi:"i32,name=retval" json:"retval,omitempty"`
	Count       uint16                  `binapi:"u16,name=count" json:"-"`
	Measurement []cup_event.Measurement `binapi:"measurement[count],name=measurement" json:"measurement,omitempty"`
}

func (m *CupAddDelSessionReply) Reset()               { *m = CupAddDelSessionReply{} }
func (*CupAddDelSessionReply) GetMessageName() string { return "cup_add_del_session_reply" }
func (*CupAddDelSessionReply) GetCrcString() string   { return "1207ffaa" }
func (*CupAddDelSessionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CupAddDelSessionReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 2 // m.Count
	for j1 := 0; j1 < len(m.Measurement); j1++ {
		var s1 cup_event.Measurement
		_ = s1
		if j1 < len(m.Measurement) {
			s1 = m.Measurement[j1]
		}
		size += 2 // s1.ID
		size += 1 // s1.Tr
		size += 8 // s1.Sid
		size += 8 // s1.Ul.Bytes
		size += 8 // s1.Ul.Pkts
		size += 8 // s1.Dl.Bytes
		size += 8 // s1.Dl.Pkts
		size += 8 // s1.Total.Bytes
		size += 8 // s1.Total.Pkts
		size += 8 // s1.FirstTime
		size += 8 // s1.LastTime
	}
	return size
}
func (m *CupAddDelSessionReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint16(uint16(len(m.Measurement)))
	for j0 := 0; j0 < len(m.Measurement); j0++ {
		var v0 cup_event.Measurement // Measurement
		if j0 < len(m.Measurement) {
			v0 = m.Measurement[j0]
		}
		buf.EncodeUint16(v0.ID)
		buf.EncodeUint8(uint8(v0.Tr))
		buf.EncodeUint64(v0.Sid)
		buf.EncodeUint64(v0.Ul.Bytes)
		buf.EncodeUint64(v0.Ul.Pkts)
		buf.EncodeUint64(v0.Dl.Bytes)
		buf.EncodeUint64(v0.Dl.Pkts)
		buf.EncodeUint64(v0.Total.Bytes)
		buf.EncodeUint64(v0.Total.Pkts)
		buf.EncodeFloat64(v0.FirstTime)
		buf.EncodeFloat64(v0.LastTime)
	}
	return buf.Bytes(), nil
}
func (m *CupAddDelSessionReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Count = buf.DecodeUint16()
	m.Measurement = make([]cup_event.Measurement, m.Count)
	for j0 := 0; j0 < len(m.Measurement); j0++ {
		m.Measurement[j0].ID = buf.DecodeUint16()
		m.Measurement[j0].Tr = cup_event.Trigger(buf.DecodeUint8())
		m.Measurement[j0].Sid = buf.DecodeUint64()
		m.Measurement[j0].Ul.Bytes = buf.DecodeUint64()
		m.Measurement[j0].Ul.Pkts = buf.DecodeUint64()
		m.Measurement[j0].Dl.Bytes = buf.DecodeUint64()
		m.Measurement[j0].Dl.Pkts = buf.DecodeUint64()
		m.Measurement[j0].Total.Bytes = buf.DecodeUint64()
		m.Measurement[j0].Total.Pkts = buf.DecodeUint64()
		m.Measurement[j0].FirstTime = buf.DecodeFloat64()
		m.Measurement[j0].LastTime = buf.DecodeFloat64()
	}
	return nil
}

// CupDumpMeasurement defines message 'cup_dump_measurement'.
type CupDumpMeasurement struct {
	ID    uint64   `binapi:"u64,name=id" json:"id,omitempty"`
	Count uint16   `binapi:"u16,name=count" json:"-"`
	Mid   []uint16 `binapi:"u16[count],name=mid" json:"mid,omitempty"`
}

func (m *CupDumpMeasurement) Reset()               { *m = CupDumpMeasurement{} }
func (*CupDumpMeasurement) GetMessageName() string { return "cup_dump_measurement" }
func (*CupDumpMeasurement) GetCrcString() string   { return "6a4edced" }
func (*CupDumpMeasurement) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CupDumpMeasurement) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 8              // m.ID
	size += 2              // m.Count
	size += 2 * len(m.Mid) // m.Mid
	return size
}
func (m *CupDumpMeasurement) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint64(m.ID)
	buf.EncodeUint16(uint16(len(m.Mid)))
	for i := 0; i < len(m.Mid); i++ {
		var x uint16
		if i < len(m.Mid) {
			x = uint16(m.Mid[i])
		}
		buf.EncodeUint16(x)
	}
	return buf.Bytes(), nil
}
func (m *CupDumpMeasurement) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ID = buf.DecodeUint64()
	m.Count = buf.DecodeUint16()
	m.Mid = make([]uint16, m.Count)
	for i := 0; i < len(m.Mid); i++ {
		m.Mid[i] = buf.DecodeUint16()
	}
	return nil
}

// CupDumpMeasurementReply defines message 'cup_dump_measurement_reply'.
type CupDumpMeasurementReply struct {
	Retval      int32                   `binapi:"i32,name=retval" json:"retval,omitempty"`
	Count       uint16                  `binapi:"u16,name=count" json:"-"`
	Measurement []cup_event.Measurement `binapi:"measurement[count],name=measurement" json:"measurement,omitempty"`
}

func (m *CupDumpMeasurementReply) Reset()               { *m = CupDumpMeasurementReply{} }
func (*CupDumpMeasurementReply) GetMessageName() string { return "cup_dump_measurement_reply" }
func (*CupDumpMeasurementReply) GetCrcString() string   { return "1207ffaa" }
func (*CupDumpMeasurementReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CupDumpMeasurementReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 2 // m.Count
	for j1 := 0; j1 < len(m.Measurement); j1++ {
		var s1 cup_event.Measurement
		_ = s1
		if j1 < len(m.Measurement) {
			s1 = m.Measurement[j1]
		}
		size += 2 // s1.ID
		size += 1 // s1.Tr
		size += 8 // s1.Sid
		size += 8 // s1.Ul.Bytes
		size += 8 // s1.Ul.Pkts
		size += 8 // s1.Dl.Bytes
		size += 8 // s1.Dl.Pkts
		size += 8 // s1.Total.Bytes
		size += 8 // s1.Total.Pkts
		size += 8 // s1.FirstTime
		size += 8 // s1.LastTime
	}
	return size
}
func (m *CupDumpMeasurementReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint16(uint16(len(m.Measurement)))
	for j0 := 0; j0 < len(m.Measurement); j0++ {
		var v0 cup_event.Measurement // Measurement
		if j0 < len(m.Measurement) {
			v0 = m.Measurement[j0]
		}
		buf.EncodeUint16(v0.ID)
		buf.EncodeUint8(uint8(v0.Tr))
		buf.EncodeUint64(v0.Sid)
		buf.EncodeUint64(v0.Ul.Bytes)
		buf.EncodeUint64(v0.Ul.Pkts)
		buf.EncodeUint64(v0.Dl.Bytes)
		buf.EncodeUint64(v0.Dl.Pkts)
		buf.EncodeUint64(v0.Total.Bytes)
		buf.EncodeUint64(v0.Total.Pkts)
		buf.EncodeFloat64(v0.FirstTime)
		buf.EncodeFloat64(v0.LastTime)
	}
	return buf.Bytes(), nil
}
func (m *CupDumpMeasurementReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Count = buf.DecodeUint16()
	m.Measurement = make([]cup_event.Measurement, m.Count)
	for j0 := 0; j0 < len(m.Measurement); j0++ {
		m.Measurement[j0].ID = buf.DecodeUint16()
		m.Measurement[j0].Tr = cup_event.Trigger(buf.DecodeUint8())
		m.Measurement[j0].Sid = buf.DecodeUint64()
		m.Measurement[j0].Ul.Bytes = buf.DecodeUint64()
		m.Measurement[j0].Ul.Pkts = buf.DecodeUint64()
		m.Measurement[j0].Dl.Bytes = buf.DecodeUint64()
		m.Measurement[j0].Dl.Pkts = buf.DecodeUint64()
		m.Measurement[j0].Total.Bytes = buf.DecodeUint64()
		m.Measurement[j0].Total.Pkts = buf.DecodeUint64()
		m.Measurement[j0].FirstTime = buf.DecodeFloat64()
		m.Measurement[j0].LastTime = buf.DecodeFloat64()
	}
	return nil
}

// CupUpdateSession defines message 'cup_update_session'.
type CupUpdateSession struct {
	ID       uint64   `binapi:"u64,name=id" json:"id,omitempty"`
	Count    uint8    `binapi:"u8,name=count" json:"-"`
	Policies []Policy `binapi:"policy[count],name=policies" json:"policies,omitempty"`
}

func (m *CupUpdateSession) Reset()               { *m = CupUpdateSession{} }
func (*CupUpdateSession) GetMessageName() string { return "cup_update_session" }
func (*CupUpdateSession) GetCrcString() string   { return "4522a415" }
func (*CupUpdateSession) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CupUpdateSession) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 8 // m.ID
	size += 1 // m.Count
	for j1 := 0; j1 < len(m.Policies); j1++ {
		var s1 Policy
		_ = s1
		if j1 < len(m.Policies) {
			s1 = m.Policies[j1]
		}
		size += 1      // s1.Type
		size += 1 * 68 // s1.Value
	}
	return size
}
func (m *CupUpdateSession) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint64(m.ID)
	buf.EncodeUint8(uint8(len(m.Policies)))
	for j0 := 0; j0 < len(m.Policies); j0++ {
		var v0 Policy // Policies
		if j0 < len(m.Policies) {
			v0 = m.Policies[j0]
		}
		buf.EncodeUint8(uint8(v0.Type))
		buf.EncodeBytes(v0.Value.XXX_UnionData[:], 68)
	}
	return buf.Bytes(), nil
}
func (m *CupUpdateSession) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ID = buf.DecodeUint64()
	m.Count = buf.DecodeUint8()
	m.Policies = make([]Policy, m.Count)
	for j0 := 0; j0 < len(m.Policies); j0++ {
		m.Policies[j0].Type = PolType(buf.DecodeUint8())
		copy(m.Policies[j0].Value.XXX_UnionData[:], buf.DecodeBytes(68))
	}
	return nil
}

// CupUpdateSessionReply defines message 'cup_update_session_reply'.
type CupUpdateSessionReply struct {
	Retval      int32                   `binapi:"i32,name=retval" json:"retval,omitempty"`
	Count       uint16                  `binapi:"u16,name=count" json:"-"`
	Measurement []cup_event.Measurement `binapi:"measurement[count],name=measurement" json:"measurement,omitempty"`
}

func (m *CupUpdateSessionReply) Reset()               { *m = CupUpdateSessionReply{} }
func (*CupUpdateSessionReply) GetMessageName() string { return "cup_update_session_reply" }
func (*CupUpdateSessionReply) GetCrcString() string   { return "1207ffaa" }
func (*CupUpdateSessionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CupUpdateSessionReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 2 // m.Count
	for j1 := 0; j1 < len(m.Measurement); j1++ {
		var s1 cup_event.Measurement
		_ = s1
		if j1 < len(m.Measurement) {
			s1 = m.Measurement[j1]
		}
		size += 2 // s1.ID
		size += 1 // s1.Tr
		size += 8 // s1.Sid
		size += 8 // s1.Ul.Bytes
		size += 8 // s1.Ul.Pkts
		size += 8 // s1.Dl.Bytes
		size += 8 // s1.Dl.Pkts
		size += 8 // s1.Total.Bytes
		size += 8 // s1.Total.Pkts
		size += 8 // s1.FirstTime
		size += 8 // s1.LastTime
	}
	return size
}
func (m *CupUpdateSessionReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint16(uint16(len(m.Measurement)))
	for j0 := 0; j0 < len(m.Measurement); j0++ {
		var v0 cup_event.Measurement // Measurement
		if j0 < len(m.Measurement) {
			v0 = m.Measurement[j0]
		}
		buf.EncodeUint16(v0.ID)
		buf.EncodeUint8(uint8(v0.Tr))
		buf.EncodeUint64(v0.Sid)
		buf.EncodeUint64(v0.Ul.Bytes)
		buf.EncodeUint64(v0.Ul.Pkts)
		buf.EncodeUint64(v0.Dl.Bytes)
		buf.EncodeUint64(v0.Dl.Pkts)
		buf.EncodeUint64(v0.Total.Bytes)
		buf.EncodeUint64(v0.Total.Pkts)
		buf.EncodeFloat64(v0.FirstTime)
		buf.EncodeFloat64(v0.LastTime)
	}
	return buf.Bytes(), nil
}
func (m *CupUpdateSessionReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Count = buf.DecodeUint16()
	m.Measurement = make([]cup_event.Measurement, m.Count)
	for j0 := 0; j0 < len(m.Measurement); j0++ {
		m.Measurement[j0].ID = buf.DecodeUint16()
		m.Measurement[j0].Tr = cup_event.Trigger(buf.DecodeUint8())
		m.Measurement[j0].Sid = buf.DecodeUint64()
		m.Measurement[j0].Ul.Bytes = buf.DecodeUint64()
		m.Measurement[j0].Ul.Pkts = buf.DecodeUint64()
		m.Measurement[j0].Dl.Bytes = buf.DecodeUint64()
		m.Measurement[j0].Dl.Pkts = buf.DecodeUint64()
		m.Measurement[j0].Total.Bytes = buf.DecodeUint64()
		m.Measurement[j0].Total.Pkts = buf.DecodeUint64()
		m.Measurement[j0].FirstTime = buf.DecodeFloat64()
		m.Measurement[j0].LastTime = buf.DecodeFloat64()
	}
	return nil
}

func init() { file_cup_session_binapi_init() }
func file_cup_session_binapi_init() {
	api.RegisterMessage((*CupAddDelSession)(nil), "cup_add_del_session_27815055")
	api.RegisterMessage((*CupAddDelSessionReply)(nil), "cup_add_del_session_reply_1207ffaa")
	api.RegisterMessage((*CupDumpMeasurement)(nil), "cup_dump_measurement_6a4edced")
	api.RegisterMessage((*CupDumpMeasurementReply)(nil), "cup_dump_measurement_reply_1207ffaa")
	api.RegisterMessage((*CupUpdateSession)(nil), "cup_update_session_4522a415")
	api.RegisterMessage((*CupUpdateSessionReply)(nil), "cup_update_session_reply_1207ffaa")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*CupAddDelSession)(nil),
		(*CupAddDelSessionReply)(nil),
		(*CupDumpMeasurement)(nil),
		(*CupDumpMeasurementReply)(nil),
		(*CupUpdateSession)(nil),
		(*CupUpdateSessionReply)(nil),
	}
}
