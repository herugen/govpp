// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: tap.api.json

/*
 Package tap is a generated from VPP binary API module 'tap'.

 It contains following objects:
	  8 messages
	  4 services

*/
package tap

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

/* Messages */

// TapConnect represents the VPP binary API message 'tap_connect'.
//
//            "tap_connect",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "use_random_mac"
//            ],
//            [
//                "u8",
//                "tap_name",
//                64
//            ],
//            [
//                "u8",
//                "mac_address",
//                6
//            ],
//            [
//                "u8",
//                "renumber"
//            ],
//            [
//                "u32",
//                "custom_dev_instance"
//            ],
//            [
//                "u8",
//                "ip4_address_set"
//            ],
//            [
//                "u8",
//                "ip4_address",
//                4
//            ],
//            [
//                "u8",
//                "ip4_mask_width"
//            ],
//            [
//                "u8",
//                "ip6_address_set"
//            ],
//            [
//                "u8",
//                "ip6_address",
//                16
//            ],
//            [
//                "u8",
//                "ip6_mask_width"
//            ],
//            [
//                "u8",
//                "tag",
//                64
//            ],
//            {
//                "crc": "0x9b9c396f"
//            }
//
type TapConnect struct {
	UseRandomMac      uint8
	TapName           []byte `struc:"[64]byte"`
	MacAddress        []byte `struc:"[6]byte"`
	Renumber          uint8
	CustomDevInstance uint32
	IP4AddressSet     uint8
	IP4Address        []byte `struc:"[4]byte"`
	IP4MaskWidth      uint8
	IP6AddressSet     uint8
	IP6Address        []byte `struc:"[16]byte"`
	IP6MaskWidth      uint8
	Tag               []byte `struc:"[64]byte"`
}

func (*TapConnect) GetMessageName() string {
	return "tap_connect"
}
func (*TapConnect) GetCrcString() string {
	return "9b9c396f"
}
func (*TapConnect) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapConnectReply represents the VPP binary API message 'tap_connect_reply'.
//
//            "tap_connect_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            {
//                "crc": "0xfda5941f"
//            }
//
type TapConnectReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*TapConnectReply) GetMessageName() string {
	return "tap_connect_reply"
}
func (*TapConnectReply) GetCrcString() string {
	return "fda5941f"
}
func (*TapConnectReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// TapModify represents the VPP binary API message 'tap_modify'.
//
//            "tap_modify",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u8",
//                "use_random_mac"
//            ],
//            [
//                "u8",
//                "tap_name",
//                64
//            ],
//            [
//                "u8",
//                "mac_address",
//                6
//            ],
//            [
//                "u8",
//                "renumber"
//            ],
//            [
//                "u32",
//                "custom_dev_instance"
//            ],
//            {
//                "crc": "0x8047ae5c"
//            }
//
type TapModify struct {
	SwIfIndex         uint32
	UseRandomMac      uint8
	TapName           []byte `struc:"[64]byte"`
	MacAddress        []byte `struc:"[6]byte"`
	Renumber          uint8
	CustomDevInstance uint32
}

func (*TapModify) GetMessageName() string {
	return "tap_modify"
}
func (*TapModify) GetCrcString() string {
	return "8047ae5c"
}
func (*TapModify) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapModifyReply represents the VPP binary API message 'tap_modify_reply'.
//
//            "tap_modify_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            {
//                "crc": "0xfda5941f"
//            }
//
type TapModifyReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*TapModifyReply) GetMessageName() string {
	return "tap_modify_reply"
}
func (*TapModifyReply) GetCrcString() string {
	return "fda5941f"
}
func (*TapModifyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// TapDelete represents the VPP binary API message 'tap_delete'.
//
//            "tap_delete",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            {
//                "crc": "0x529cb13f"
//            }
//
type TapDelete struct {
	SwIfIndex uint32
}

func (*TapDelete) GetMessageName() string {
	return "tap_delete"
}
func (*TapDelete) GetCrcString() string {
	return "529cb13f"
}
func (*TapDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapDeleteReply represents the VPP binary API message 'tap_delete_reply'.
//
//            "tap_delete_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type TapDeleteReply struct {
	Retval int32
}

func (*TapDeleteReply) GetMessageName() string {
	return "tap_delete_reply"
}
func (*TapDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*TapDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwInterfaceTapDump represents the VPP binary API message 'sw_interface_tap_dump'.
//
//            "sw_interface_tap_dump",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type SwInterfaceTapDump struct{}

func (*SwInterfaceTapDump) GetMessageName() string {
	return "sw_interface_tap_dump"
}
func (*SwInterfaceTapDump) GetCrcString() string {
	return "51077d14"
}
func (*SwInterfaceTapDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SwInterfaceTapDetails represents the VPP binary API message 'sw_interface_tap_details'.
//
//            "sw_interface_tap_details",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u8",
//                "dev_name",
//                64
//            ],
//            {
//                "crc": "0x76229a57"
//            }
//
type SwInterfaceTapDetails struct {
	SwIfIndex uint32
	DevName   []byte `struc:"[64]byte"`
}

func (*SwInterfaceTapDetails) GetMessageName() string {
	return "sw_interface_tap_details"
}
func (*SwInterfaceTapDetails) GetCrcString() string {
	return "76229a57"
}
func (*SwInterfaceTapDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

/* Services */

type Services interface {
	DumpSwInterfaceTap(*SwInterfaceTapDump) (*SwInterfaceTapDetails, error)
	TapConnect(*TapConnect) (*TapConnectReply, error)
	TapDelete(*TapDelete) (*TapDeleteReply, error)
	TapModify(*TapModify) (*TapModifyReply, error)
}

func init() {
	api.RegisterMessage((*TapConnect)(nil), "tap.TapConnect")
	api.RegisterMessage((*TapConnectReply)(nil), "tap.TapConnectReply")
	api.RegisterMessage((*TapModify)(nil), "tap.TapModify")
	api.RegisterMessage((*TapModifyReply)(nil), "tap.TapModifyReply")
	api.RegisterMessage((*TapDelete)(nil), "tap.TapDelete")
	api.RegisterMessage((*TapDeleteReply)(nil), "tap.TapDeleteReply")
	api.RegisterMessage((*SwInterfaceTapDump)(nil), "tap.SwInterfaceTapDump")
	api.RegisterMessage((*SwInterfaceTapDetails)(nil), "tap.SwInterfaceTapDetails")
}
