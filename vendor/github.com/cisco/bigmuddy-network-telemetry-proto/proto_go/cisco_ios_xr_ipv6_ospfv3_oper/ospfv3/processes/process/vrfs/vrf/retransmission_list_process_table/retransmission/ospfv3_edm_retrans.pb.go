// Code generated by protoc-gen-go.
// source: ospfv3_edm_retrans.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_retransmission_list_process_table_retransmission is a generated protocol buffer package.

It is generated from these files:
	ospfv3_edm_retrans.proto

It has these top-level messages:
	Ospfv3EdmRetrans_KEYS
	Ospfv3EdmRetrans
	Ospfv3EdmLsaSum
*/
package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_retransmission_list_process_table_retransmission

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// OSPFv3 retransmission list information
type Ospfv3EdmRetrans_KEYS struct {
	ProcessName     string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	VrfName         string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	InterfaceName   string `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	NeighborAddress string `protobuf:"bytes,4,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
}

func (m *Ospfv3EdmRetrans_KEYS) Reset()                    { *m = Ospfv3EdmRetrans_KEYS{} }
func (m *Ospfv3EdmRetrans_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmRetrans_KEYS) ProtoMessage()               {}
func (*Ospfv3EdmRetrans_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ospfv3EdmRetrans_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmRetrans_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ospfv3EdmRetrans_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ospfv3EdmRetrans_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type Ospfv3EdmRetrans struct {
	// Neighbor IP address
	RetransmissionNeighborAddress string `protobuf:"bytes,50,opt,name=retransmission_neighbor_address,json=retransmissionNeighborAddress" json:"retransmission_neighbor_address,omitempty"`
	// If true, virtual link is retransmitted
	IsRetransmissionVirtualLink bool `protobuf:"varint,51,opt,name=is_retransmission_virtual_link,json=isRetransmissionVirtualLink" json:"is_retransmission_virtual_link,omitempty"`
	// Retransmission virtual link ID
	RetransmissionVirtualLinkId uint32 `protobuf:"varint,52,opt,name=retransmission_virtual_link_id,json=retransmissionVirtualLinkId" json:"retransmission_virtual_link_id,omitempty"`
	// If true, sham link is retransmitted
	IsRetransmissionShamLink bool `protobuf:"varint,53,opt,name=is_retransmission_sham_link,json=isRetransmissionShamLink" json:"is_retransmission_sham_link,omitempty"`
	// Retransmission sham link ID
	RetransmissionShamLinkId uint32 `protobuf:"varint,54,opt,name=retransmission_sham_link_id,json=retransmissionShamLinkId" json:"retransmission_sham_link_id,omitempty"`
	// Amount of time remaining on retransmission timer (ms)
	RetransmissionTimer uint32 `protobuf:"varint,55,opt,name=retransmission_timer,json=retransmissionTimer" json:"retransmission_timer,omitempty"`
	// Retransmission queue length
	RetransmissionLength uint32 `protobuf:"varint,56,opt,name=retransmission_length,json=retransmissionLength" json:"retransmission_length,omitempty"`
	// List of virtual link scope entries
	RetransmissionVirtualLinkDbList []*Ospfv3EdmLsaSum `protobuf:"bytes,57,rep,name=retransmission_virtual_link_db_list,json=retransmissionVirtualLinkDbList" json:"retransmission_virtual_link_db_list,omitempty"`
	// List of area scope entries
	RetransmissionAreaDbList []*Ospfv3EdmLsaSum `protobuf:"bytes,58,rep,name=retransmission_area_db_list,json=retransmissionAreaDbList" json:"retransmission_area_db_list,omitempty"`
	// List of AS scope entries
	RetransmissionAsdbList []*Ospfv3EdmLsaSum `protobuf:"bytes,59,rep,name=retransmission_asdb_list,json=retransmissionAsdbList" json:"retransmission_asdb_list,omitempty"`
}

func (m *Ospfv3EdmRetrans) Reset()                    { *m = Ospfv3EdmRetrans{} }
func (m *Ospfv3EdmRetrans) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmRetrans) ProtoMessage()               {}
func (*Ospfv3EdmRetrans) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ospfv3EdmRetrans) GetRetransmissionNeighborAddress() string {
	if m != nil {
		return m.RetransmissionNeighborAddress
	}
	return ""
}

func (m *Ospfv3EdmRetrans) GetIsRetransmissionVirtualLink() bool {
	if m != nil {
		return m.IsRetransmissionVirtualLink
	}
	return false
}

func (m *Ospfv3EdmRetrans) GetRetransmissionVirtualLinkId() uint32 {
	if m != nil {
		return m.RetransmissionVirtualLinkId
	}
	return 0
}

func (m *Ospfv3EdmRetrans) GetIsRetransmissionShamLink() bool {
	if m != nil {
		return m.IsRetransmissionShamLink
	}
	return false
}

func (m *Ospfv3EdmRetrans) GetRetransmissionShamLinkId() uint32 {
	if m != nil {
		return m.RetransmissionShamLinkId
	}
	return 0
}

func (m *Ospfv3EdmRetrans) GetRetransmissionTimer() uint32 {
	if m != nil {
		return m.RetransmissionTimer
	}
	return 0
}

func (m *Ospfv3EdmRetrans) GetRetransmissionLength() uint32 {
	if m != nil {
		return m.RetransmissionLength
	}
	return 0
}

func (m *Ospfv3EdmRetrans) GetRetransmissionVirtualLinkDbList() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.RetransmissionVirtualLinkDbList
	}
	return nil
}

func (m *Ospfv3EdmRetrans) GetRetransmissionAreaDbList() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.RetransmissionAreaDbList
	}
	return nil
}

func (m *Ospfv3EdmRetrans) GetRetransmissionAsdbList() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.RetransmissionAsdbList
	}
	return nil
}

// LSA summary entry
type Ospfv3EdmLsaSum struct {
	// LSA type
	HeaderLsaType string `protobuf:"bytes,1,opt,name=header_lsa_type,json=headerLsaType" json:"header_lsa_type,omitempty"`
	// Age of the LSA (seconds)
	HeaderLsaAge uint32 `protobuf:"varint,2,opt,name=header_lsa_age,json=headerLsaAge" json:"header_lsa_age,omitempty"`
	// LSA ID
	HeaderLsaId string `protobuf:"bytes,3,opt,name=header_lsa_id,json=headerLsaId" json:"header_lsa_id,omitempty"`
	// Router ID of the advertising router
	HeaderAdvertisingRouter string `protobuf:"bytes,4,opt,name=header_advertising_router,json=headerAdvertisingRouter" json:"header_advertising_router,omitempty"`
	// Current LSA sequence number
	HeaderSequenceNumber int32 `protobuf:"zigzag32,5,opt,name=header_sequence_number,json=headerSequenceNumber" json:"header_sequence_number,omitempty"`
}

func (m *Ospfv3EdmLsaSum) Reset()                    { *m = Ospfv3EdmLsaSum{} }
func (m *Ospfv3EdmLsaSum) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmLsaSum) ProtoMessage()               {}
func (*Ospfv3EdmLsaSum) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ospfv3EdmLsaSum) GetHeaderLsaType() string {
	if m != nil {
		return m.HeaderLsaType
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderLsaAge() uint32 {
	if m != nil {
		return m.HeaderLsaAge
	}
	return 0
}

func (m *Ospfv3EdmLsaSum) GetHeaderLsaId() string {
	if m != nil {
		return m.HeaderLsaId
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderAdvertisingRouter() string {
	if m != nil {
		return m.HeaderAdvertisingRouter
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderSequenceNumber() int32 {
	if m != nil {
		return m.HeaderSequenceNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*Ospfv3EdmRetrans_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.retransmission_list_process_table.retransmission.ospfv3_edm_retrans_KEYS")
	proto.RegisterType((*Ospfv3EdmRetrans)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.retransmission_list_process_table.retransmission.ospfv3_edm_retrans")
	proto.RegisterType((*Ospfv3EdmLsaSum)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.retransmission_list_process_table.retransmission.ospfv3_edm_lsa_sum")
}

func init() { proto.RegisterFile("ospfv3_edm_retrans.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xdd, 0x6e, 0x12, 0x41,
	0x14, 0xc7, 0xb3, 0xd6, 0x8f, 0x3a, 0x94, 0x56, 0xc7, 0xda, 0x4e, 0x43, 0xb4, 0x88, 0x1f, 0xc1,
	0x9b, 0x4d, 0x2c, 0xb5, 0x6a, 0x4d, 0x2f, 0x88, 0x1f, 0x09, 0x91, 0x70, 0xb1, 0x34, 0x26, 0x5e,
	0x4d, 0x66, 0x99, 0x03, 0x4c, 0xca, 0x7e, 0x38, 0x67, 0xd8, 0xd8, 0x7b, 0x5f, 0xc6, 0xf8, 0x10,
	0xde, 0xf8, 0x06, 0xbe, 0x8c, 0x97, 0x86, 0xd9, 0x01, 0x61, 0x91, 0x5e, 0x9a, 0xde, 0x90, 0xcd,
	0xf9, 0xff, 0xce, 0xc9, 0xff, 0xbf, 0x1c, 0x0e, 0x84, 0x25, 0x98, 0xf6, 0xb3, 0x06, 0x07, 0x19,
	0x71, 0x0d, 0x46, 0x8b, 0x18, 0xfd, 0x54, 0x27, 0x26, 0xa1, 0x69, 0x4f, 0x61, 0x2f, 0xe1, 0x2a,
	0x41, 0xfe, 0x45, 0x73, 0x95, 0x66, 0x47, 0xdc, 0xb1, 0x49, 0x0a, 0xda, 0xcf, 0x9f, 0x27, 0x6c,
	0x0f, 0x10, 0x01, 0xa7, 0x4f, 0x7e, 0xa6, 0xfb, 0xf6, 0xc3, 0x77, 0xe3, 0x22, 0x85, 0xa8, 0x92,
	0x98, 0x8f, 0x14, 0x1a, 0xee, 0x20, 0x6e, 0x44, 0x38, 0x82, 0x02, 0x51, 0xfb, 0xe6, 0x91, 0xdd,
	0x65, 0x3b, 0xfc, 0xc3, 0xbb, 0x4f, 0x5d, 0xfa, 0x80, 0x6c, 0x4c, 0x7b, 0x63, 0x11, 0x01, 0xf3,
	0xaa, 0x5e, 0xfd, 0x66, 0x50, 0x72, 0xb5, 0x8e, 0x88, 0x80, 0xee, 0x91, 0xf5, 0x4c, 0xf7, 0x73,
	0xf9, 0x8a, 0x95, 0x6f, 0x64, 0xba, 0x6f, 0xa5, 0xc7, 0x64, 0x53, 0xc5, 0x06, 0x74, 0x5f, 0xf4,
	0x20, 0x07, 0xd6, 0x2c, 0x50, 0x9e, 0x55, 0x2d, 0xf6, 0x94, 0xdc, 0x8a, 0x41, 0x0d, 0x86, 0x61,
	0xa2, 0xb9, 0x90, 0x52, 0x03, 0x22, 0xbb, 0x6a, 0xc1, 0xad, 0x69, 0xbd, 0x99, 0x97, 0x6b, 0xdf,
	0xd7, 0x09, 0x5d, 0xf6, 0x4a, 0xdf, 0x93, 0xfd, 0x42, 0xec, 0xa5, 0x81, 0x07, 0x76, 0xe0, 0xbd,
	0x45, 0xac, 0xb3, 0x38, 0x9e, 0xbe, 0x21, 0xf7, 0x15, 0xf2, 0xc2, 0xa8, 0x4c, 0x69, 0x33, 0x16,
	0x23, 0x3e, 0x52, 0xf1, 0x19, 0x6b, 0x54, 0xbd, 0xfa, 0x7a, 0x50, 0x51, 0x18, 0x2c, 0x40, 0x1f,
	0x73, 0xa6, 0xad, 0xe2, 0xb3, 0xc9, 0x90, 0x0b, 0x26, 0x70, 0x25, 0xd9, 0x61, 0xd5, 0xab, 0x97,
	0x83, 0x8a, 0x5e, 0x35, 0xa2, 0x25, 0xe9, 0x09, 0xa9, 0x2c, 0x3b, 0xc1, 0xa1, 0x88, 0x72, 0x1b,
	0xcf, 0xad, 0x0d, 0x56, 0xb4, 0xd1, 0x1d, 0x8a, 0xc8, 0x7a, 0x38, 0x21, 0x95, 0x55, 0xbd, 0x13,
	0x03, 0x47, 0xd6, 0x00, 0xd3, 0xff, 0x6c, 0x6e, 0x49, 0xfa, 0x8c, 0x6c, 0x17, 0xda, 0x8d, 0x8a,
	0x40, 0xb3, 0x17, 0xb6, 0xef, 0xce, 0xa2, 0x76, 0x3a, 0x91, 0x68, 0x83, 0xdc, 0x2d, 0x6e, 0x1e,
	0xc4, 0x03, 0x33, 0x64, 0x2f, 0x6d, 0x4f, 0x61, 0x5e, 0xdb, 0x6a, 0xf4, 0x97, 0x47, 0x1e, 0x5e,
	0xf4, 0xae, 0x64, 0x68, 0xf7, 0x97, 0xbd, 0xaa, 0xae, 0xd5, 0x4b, 0x07, 0x5f, 0x3d, 0xff, 0x7f,
	0xff, 0x38, 0xfc, 0xb9, 0x65, 0x1b, 0xa1, 0xe0, 0x38, 0x8e, 0x82, 0xfd, 0x95, 0xdf, 0xdb, 0xdb,
	0xb0, 0xad, 0xd0, 0xd0, 0x9f, 0xde, 0xd2, 0xdb, 0x17, 0x1a, 0xc4, 0x2c, 0xcd, 0xf1, 0x65, 0x4a,
	0x53, 0x58, 0x82, 0xa6, 0x06, 0xe1, 0x62, 0xfc, 0xf0, 0x08, 0x2b, 0xc6, 0xc0, 0x69, 0x86, 0xd7,
	0x97, 0x29, 0xc3, 0x4e, 0x21, 0x03, 0x4a, 0x9b, 0xa0, 0xf6, 0xdb, 0x5b, 0xb8, 0x16, 0x0e, 0xa7,
	0x4f, 0xc8, 0xd6, 0x10, 0x84, 0x04, 0x6d, 0x2b, 0xe6, 0x3c, 0x9d, 0xde, 0xb5, 0x72, 0x5e, 0x6e,
	0xa3, 0x38, 0x3d, 0x4f, 0x81, 0x3e, 0x22, 0x9b, 0x73, 0x9c, 0x18, 0xe4, 0xf7, 0xad, 0x1c, 0x6c,
	0xcc, 0xb0, 0xe6, 0x00, 0x68, 0x8d, 0x94, 0xe7, 0x28, 0x25, 0xdd, 0x8d, 0x2b, 0xcd, 0xa0, 0x96,
	0xa4, 0xc7, 0x64, 0xcf, 0x31, 0x42, 0x66, 0xa0, 0x8d, 0x42, 0x15, 0x0f, 0xb8, 0x4e, 0xc6, 0x06,
	0xb4, 0x3b, 0x75, 0xbb, 0x39, 0xd0, 0xfc, 0xab, 0x07, 0x56, 0xa6, 0x87, 0x64, 0xc7, 0xf5, 0x22,
	0x7c, 0x1e, 0x43, 0x3c, 0x39, 0xa5, 0xe3, 0x28, 0x04, 0xcd, 0xae, 0x55, 0xbd, 0xfa, 0xed, 0x60,
	0x3b, 0x57, 0xbb, 0x4e, 0xec, 0x58, 0x2d, 0xbc, 0x6e, 0xff, 0x4d, 0x1a, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x5f, 0x6b, 0xbe, 0xae, 0x69, 0x06, 0x00, 0x00,
}
