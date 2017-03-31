// Code generated by protoc-gen-go.
// source: isis_sh_ipv6_topo_entry.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_topology_levels_topology_level_ipv6_link_topologies_ipv6_link_topology is a generated protocol buffer package.

It is generated from these files:
	isis_sh_ipv6_topo_entry.proto

It has these top-level messages:
	IsisShIpv6TopoEntry_KEYS
	IsisShIpv6TopoEntry
	IsisNodeIdType
	IsisSnpaType
	IsisPerPriorityCounts
	IsisShRepEl
	IsisShIpv6FrrBackup
	IsisShIpv6Path
	IsisShTopoNeighbor
	IsisShIpv6TopoReachableDetails
	IsisShIpv6TopoReachableStatus
*/
package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_topology_levels_topology_level_ipv6_link_topologies_ipv6_link_topology

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

// IPv6 IS Link Topology Entry
type IsisShIpv6TopoEntry_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	AfName       string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName      string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	TopologyName string `protobuf:"bytes,4,opt,name=topology_name,json=topologyName" json:"topology_name,omitempty"`
	Level        string `protobuf:"bytes,5,opt,name=level" json:"level,omitempty"`
	SystemId     string `protobuf:"bytes,6,opt,name=system_id,json=systemId" json:"system_id,omitempty"`
}

func (m *IsisShIpv6TopoEntry_KEYS) Reset()                    { *m = IsisShIpv6TopoEntry_KEYS{} }
func (m *IsisShIpv6TopoEntry_KEYS) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv6TopoEntry_KEYS) ProtoMessage()               {}
func (*IsisShIpv6TopoEntry_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IsisShIpv6TopoEntry_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShIpv6TopoEntry_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisShIpv6TopoEntry_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *IsisShIpv6TopoEntry_KEYS) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

func (m *IsisShIpv6TopoEntry_KEYS) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *IsisShIpv6TopoEntry_KEYS) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

type IsisShIpv6TopoEntry struct {
	// Source Address
	SourceAddress string `protobuf:"bytes,50,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	// Does the IS participate in the topology?
	IsParticipant bool `protobuf:"varint,51,opt,name=is_participant,json=isParticipant" json:"is_participant,omitempty"`
	// Is the IS overloaded?
	IsOverloaded bool `protobuf:"varint,52,opt,name=is_overloaded,json=isOverloaded" json:"is_overloaded,omitempty"`
	// Is the IS attached?
	IsAttached bool `protobuf:"varint,53,opt,name=is_attached,json=isAttached" json:"is_attached,omitempty"`
	// Is the IS reachable, and, if so, its status within the SPT
	ReachabilityStatus *IsisShIpv6TopoReachableStatus `protobuf:"bytes,54,opt,name=reachability_status,json=reachabilityStatus" json:"reachability_status,omitempty"`
	// Per-priority counts of prefix items advertised by the IS
	AdvertisedPrefixItemCounts *IsisPerPriorityCounts `protobuf:"bytes,55,opt,name=advertised_prefix_item_counts,json=advertisedPrefixItemCounts" json:"advertised_prefix_item_counts,omitempty"`
}

func (m *IsisShIpv6TopoEntry) Reset()                    { *m = IsisShIpv6TopoEntry{} }
func (m *IsisShIpv6TopoEntry) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv6TopoEntry) ProtoMessage()               {}
func (*IsisShIpv6TopoEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IsisShIpv6TopoEntry) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *IsisShIpv6TopoEntry) GetIsParticipant() bool {
	if m != nil {
		return m.IsParticipant
	}
	return false
}

func (m *IsisShIpv6TopoEntry) GetIsOverloaded() bool {
	if m != nil {
		return m.IsOverloaded
	}
	return false
}

func (m *IsisShIpv6TopoEntry) GetIsAttached() bool {
	if m != nil {
		return m.IsAttached
	}
	return false
}

func (m *IsisShIpv6TopoEntry) GetReachabilityStatus() *IsisShIpv6TopoReachableStatus {
	if m != nil {
		return m.ReachabilityStatus
	}
	return nil
}

func (m *IsisShIpv6TopoEntry) GetAdvertisedPrefixItemCounts() *IsisPerPriorityCounts {
	if m != nil {
		return m.AdvertisedPrefixItemCounts
	}
	return nil
}

type IsisNodeIdType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IsisNodeIdType) Reset()                    { *m = IsisNodeIdType{} }
func (m *IsisNodeIdType) String() string            { return proto.CompactTextString(m) }
func (*IsisNodeIdType) ProtoMessage()               {}
func (*IsisNodeIdType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IsisNodeIdType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type IsisSnpaType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IsisSnpaType) Reset()                    { *m = IsisSnpaType{} }
func (m *IsisSnpaType) String() string            { return proto.CompactTextString(m) }
func (*IsisSnpaType) ProtoMessage()               {}
func (*IsisSnpaType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IsisSnpaType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Per-priority counts
type IsisPerPriorityCounts struct {
	// Critical priority
	Critical uint32 `protobuf:"varint,1,opt,name=critical" json:"critical,omitempty"`
	// High priority
	High uint32 `protobuf:"varint,2,opt,name=high" json:"high,omitempty"`
	// Medium priority
	Medium uint32 `protobuf:"varint,3,opt,name=medium" json:"medium,omitempty"`
	// Low priority
	Low uint32 `protobuf:"varint,4,opt,name=low" json:"low,omitempty"`
}

func (m *IsisPerPriorityCounts) Reset()                    { *m = IsisPerPriorityCounts{} }
func (m *IsisPerPriorityCounts) String() string            { return proto.CompactTextString(m) }
func (*IsisPerPriorityCounts) ProtoMessage()               {}
func (*IsisPerPriorityCounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IsisPerPriorityCounts) GetCritical() uint32 {
	if m != nil {
		return m.Critical
	}
	return 0
}

func (m *IsisPerPriorityCounts) GetHigh() uint32 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *IsisPerPriorityCounts) GetMedium() uint32 {
	if m != nil {
		return m.Medium
	}
	return 0
}

func (m *IsisPerPriorityCounts) GetLow() uint32 {
	if m != nil {
		return m.Low
	}
	return 0
}

// OSPF Repair Element
type IsisShRepEl struct {
	// RepairElementNodeID
	RepairElementNodeId string `protobuf:"bytes,1,opt,name=repair_element_node_id,json=repairElementNodeId" json:"repair_element_node_id,omitempty"`
	// RepairIPv4Addr
	RepairIpv4Addr string `protobuf:"bytes,2,opt,name=repair_ipv4_addr,json=repairIpv4Addr" json:"repair_ipv4_addr,omitempty"`
	// RepairIPv6Addr
	RepairIpv6Addr string `protobuf:"bytes,3,opt,name=repair_ipv6_addr,json=repairIpv6Addr" json:"repair_ipv6_addr,omitempty"`
	// Repair Label
	RepairLabel uint32 `protobuf:"varint,4,opt,name=repair_label,json=repairLabel" json:"repair_label,omitempty"`
	// Repair Element Type
	RepairElementType uint32 `protobuf:"varint,5,opt,name=repair_element_type,json=repairElementType" json:"repair_element_type,omitempty"`
}

func (m *IsisShRepEl) Reset()                    { *m = IsisShRepEl{} }
func (m *IsisShRepEl) String() string            { return proto.CompactTextString(m) }
func (*IsisShRepEl) ProtoMessage()               {}
func (*IsisShRepEl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IsisShRepEl) GetRepairElementNodeId() string {
	if m != nil {
		return m.RepairElementNodeId
	}
	return ""
}

func (m *IsisShRepEl) GetRepairIpv4Addr() string {
	if m != nil {
		return m.RepairIpv4Addr
	}
	return ""
}

func (m *IsisShRepEl) GetRepairIpv6Addr() string {
	if m != nil {
		return m.RepairIpv6Addr
	}
	return ""
}

func (m *IsisShRepEl) GetRepairLabel() uint32 {
	if m != nil {
		return m.RepairLabel
	}
	return 0
}

func (m *IsisShRepEl) GetRepairElementType() uint32 {
	if m != nil {
		return m.RepairElementType
	}
	return 0
}

// FRR backup path
type IsisShIpv6FrrBackup struct {
	// Next hop neighbor ID
	NeighborId string `protobuf:"bytes,1,opt,name=neighbor_id,json=neighborId" json:"neighbor_id,omitempty"`
	// Interface to send the packet out of
	EgressInterface string `protobuf:"bytes,2,opt,name=egress_interface,json=egressInterface" json:"egress_interface,omitempty"`
	// Next hop neighbor's forwarding address
	NeighborAddress string `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
	// Tunnel interface to send the packet out of
	TunnelEgressInterface string `protobuf:"bytes,4,opt,name=tunnel_egress_interface,json=tunnelEgressInterface" json:"tunnel_egress_interface,omitempty"`
	// Next hop neighbor's SNPA
	NeighborSnpa *IsisSnpaType `protobuf:"bytes,5,opt,name=neighbor_snpa,json=neighborSnpa" json:"neighbor_snpa,omitempty"`
	// Remote LFA PQ Node's ID
	RemoteLfaSystemId string `protobuf:"bytes,6,opt,name=remote_lfa_system_id,json=remoteLfaSystemId" json:"remote_lfa_system_id,omitempty"`
	// Remote LFA Router ID
	RemoteLfaRouterId string `protobuf:"bytes,7,opt,name=remote_lfa_router_id,json=remoteLfaRouterId" json:"remote_lfa_router_id,omitempty"`
	// Remote LFA PQ Node's ID
	RemoteLfaSystemPid string `protobuf:"bytes,8,opt,name=remote_lfa_system_pid,json=remoteLfaSystemPid" json:"remote_lfa_system_pid,omitempty"`
	// Remote LFA Router ID
	RemoteLfaRouterPid string `protobuf:"bytes,9,opt,name=remote_lfa_router_pid,json=remoteLfaRouterPid" json:"remote_lfa_router_pid,omitempty"`
	// Distance to the network via this backup path
	TotalBackupDistance uint32 `protobuf:"varint,10,opt,name=total_backup_distance,json=totalBackupDistance" json:"total_backup_distance,omitempty"`
	// Segment routing sid value received from first hop
	SegmentRoutingSidValue uint32 `protobuf:"varint,11,opt,name=segment_routing_sid_value,json=segmentRoutingSidValue" json:"segment_routing_sid_value,omitempty"`
	// Number of SIDs in TI-LFA/rLFA
	NumSid uint32 `protobuf:"varint,12,opt,name=num_sid,json=numSid" json:"num_sid,omitempty"`
	// Segment routing sid values for TI-LFA/rLFA
	SegmentRoutingSidValues []uint32 `protobuf:"varint,13,rep,packed,name=segment_routing_sid_values,json=segmentRoutingSidValues" json:"segment_routing_sid_values,omitempty"`
	// Backup Repair List Size
	BackupRepairListSize uint32 `protobuf:"varint,14,opt,name=backup_repair_list_size,json=backupRepairListSize" json:"backup_repair_list_size,omitempty"`
	// Ti LFA computation which provided backup path
	TilfaComputation string `protobuf:"bytes,15,opt,name=tilfa_computation,json=tilfaComputation" json:"tilfa_computation,omitempty"`
	// BAckup Repair List
	BackupRepairList []*IsisShRepEl `protobuf:"bytes,16,rep,name=backup_repair_list,json=backupRepairList" json:"backup_repair_list,omitempty"`
	// Is the backup path via downstream node?
	IsDownstream bool `protobuf:"varint,17,opt,name=is_downstream,json=isDownstream" json:"is_downstream,omitempty"`
	// Is the backup path line card disjoint with primary?
	IsLcDisjoint bool `protobuf:"varint,18,opt,name=is_lc_disjoint,json=isLcDisjoint" json:"is_lc_disjoint,omitempty"`
	// Is the backup path node protecting?
	IsNodeProtecting bool `protobuf:"varint,19,opt,name=is_node_protecting,json=isNodeProtecting" json:"is_node_protecting,omitempty"`
	// Is the backup path an ECMP to the network?
	IsPrimaryPath bool `protobuf:"varint,20,opt,name=is_primary_path,json=isPrimaryPath" json:"is_primary_path,omitempty"`
	// Is the backup path SRLG disjoint with primary?
	IsSrlgDisjoint bool `protobuf:"varint,21,opt,name=is_srlg_disjoint,json=isSrlgDisjoint" json:"is_srlg_disjoint,omitempty"`
	// Is the backup path via a Remote LFA?
	IsRemoteLfa bool `protobuf:"varint,22,opt,name=is_remote_lfa,json=isRemoteLfa" json:"is_remote_lfa,omitempty"`
	// Is the backup path via a TI-LFA?
	IsEpcfrrLfa bool `protobuf:"varint,23,opt,name=is_epcfrr_lfa,json=isEpcfrrLfa" json:"is_epcfrr_lfa,omitempty"`
	// Is SR TE tunnel requested
	IsTunnelRequested bool `protobuf:"varint,24,opt,name=is_tunnel_requested,json=isTunnelRequested" json:"is_tunnel_requested,omitempty"`
	// Weight configured on the interface
	Weight uint32 `protobuf:"varint,25,opt,name=weight" json:"weight,omitempty"`
}

func (m *IsisShIpv6FrrBackup) Reset()                    { *m = IsisShIpv6FrrBackup{} }
func (m *IsisShIpv6FrrBackup) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv6FrrBackup) ProtoMessage()               {}
func (*IsisShIpv6FrrBackup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *IsisShIpv6FrrBackup) GetNeighborId() string {
	if m != nil {
		return m.NeighborId
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetEgressInterface() string {
	if m != nil {
		return m.EgressInterface
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetTunnelEgressInterface() string {
	if m != nil {
		return m.TunnelEgressInterface
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetNeighborSnpa() *IsisSnpaType {
	if m != nil {
		return m.NeighborSnpa
	}
	return nil
}

func (m *IsisShIpv6FrrBackup) GetRemoteLfaSystemId() string {
	if m != nil {
		return m.RemoteLfaSystemId
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetRemoteLfaRouterId() string {
	if m != nil {
		return m.RemoteLfaRouterId
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetRemoteLfaSystemPid() string {
	if m != nil {
		return m.RemoteLfaSystemPid
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetRemoteLfaRouterPid() string {
	if m != nil {
		return m.RemoteLfaRouterPid
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetTotalBackupDistance() uint32 {
	if m != nil {
		return m.TotalBackupDistance
	}
	return 0
}

func (m *IsisShIpv6FrrBackup) GetSegmentRoutingSidValue() uint32 {
	if m != nil {
		return m.SegmentRoutingSidValue
	}
	return 0
}

func (m *IsisShIpv6FrrBackup) GetNumSid() uint32 {
	if m != nil {
		return m.NumSid
	}
	return 0
}

func (m *IsisShIpv6FrrBackup) GetSegmentRoutingSidValues() []uint32 {
	if m != nil {
		return m.SegmentRoutingSidValues
	}
	return nil
}

func (m *IsisShIpv6FrrBackup) GetBackupRepairListSize() uint32 {
	if m != nil {
		return m.BackupRepairListSize
	}
	return 0
}

func (m *IsisShIpv6FrrBackup) GetTilfaComputation() string {
	if m != nil {
		return m.TilfaComputation
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetBackupRepairList() []*IsisShRepEl {
	if m != nil {
		return m.BackupRepairList
	}
	return nil
}

func (m *IsisShIpv6FrrBackup) GetIsDownstream() bool {
	if m != nil {
		return m.IsDownstream
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsLcDisjoint() bool {
	if m != nil {
		return m.IsLcDisjoint
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsNodeProtecting() bool {
	if m != nil {
		return m.IsNodeProtecting
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsPrimaryPath() bool {
	if m != nil {
		return m.IsPrimaryPath
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsSrlgDisjoint() bool {
	if m != nil {
		return m.IsSrlgDisjoint
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsRemoteLfa() bool {
	if m != nil {
		return m.IsRemoteLfa
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsEpcfrrLfa() bool {
	if m != nil {
		return m.IsEpcfrrLfa
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsTunnelRequested() bool {
	if m != nil {
		return m.IsTunnelRequested
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

// IPv6 path to a destination
type IsisShIpv6Path struct {
	// Next hop neighbor ID
	NeighborId string `protobuf:"bytes,1,opt,name=neighbor_id,json=neighborId" json:"neighbor_id,omitempty"`
	// Interface to send the packet out of
	EgressInterface string `protobuf:"bytes,2,opt,name=egress_interface,json=egressInterface" json:"egress_interface,omitempty"`
	// Next hop neighbor's forwarding address
	NeighborAddress string `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
	// Next hop neighbor's SNPA
	NeighborSnpa *IsisSnpaType `protobuf:"bytes,4,opt,name=neighbor_snpa,json=neighborSnpa" json:"neighbor_snpa,omitempty"`
	// Tag associated with the path
	Tag uint32 `protobuf:"varint,5,opt,name=tag" json:"tag,omitempty"`
	// FRR backup for this path
	FrrBackup *IsisShIpv6FrrBackup `protobuf:"bytes,6,opt,name=frr_backup,json=frrBackup" json:"frr_backup,omitempty"`
	// Uloop Explicit List
	UloopExplicitList []*IsisShRepEl `protobuf:"bytes,7,rep,name=uloop_explicit_list,json=uloopExplicitList" json:"uloop_explicit_list,omitempty"`
	// Explicit path tunnel interface
	TunnelInterface string `protobuf:"bytes,8,opt,name=tunnel_interface,json=tunnelInterface" json:"tunnel_interface,omitempty"`
	// Segment routing sid value received from first hop
	SegmentRoutingSidValue uint32 `protobuf:"varint,9,opt,name=segment_routing_sid_value,json=segmentRoutingSidValue" json:"segment_routing_sid_value,omitempty"`
	// Weight configured on the interface
	Weight uint32 `protobuf:"varint,10,opt,name=weight" json:"weight,omitempty"`
}

func (m *IsisShIpv6Path) Reset()                    { *m = IsisShIpv6Path{} }
func (m *IsisShIpv6Path) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv6Path) ProtoMessage()               {}
func (*IsisShIpv6Path) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IsisShIpv6Path) GetNeighborId() string {
	if m != nil {
		return m.NeighborId
	}
	return ""
}

func (m *IsisShIpv6Path) GetEgressInterface() string {
	if m != nil {
		return m.EgressInterface
	}
	return ""
}

func (m *IsisShIpv6Path) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *IsisShIpv6Path) GetNeighborSnpa() *IsisSnpaType {
	if m != nil {
		return m.NeighborSnpa
	}
	return nil
}

func (m *IsisShIpv6Path) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *IsisShIpv6Path) GetFrrBackup() *IsisShIpv6FrrBackup {
	if m != nil {
		return m.FrrBackup
	}
	return nil
}

func (m *IsisShIpv6Path) GetUloopExplicitList() []*IsisShRepEl {
	if m != nil {
		return m.UloopExplicitList
	}
	return nil
}

func (m *IsisShIpv6Path) GetTunnelInterface() string {
	if m != nil {
		return m.TunnelInterface
	}
	return ""
}

func (m *IsisShIpv6Path) GetSegmentRoutingSidValue() uint32 {
	if m != nil {
		return m.SegmentRoutingSidValue
	}
	return 0
}

func (m *IsisShIpv6Path) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

// SPT Neighbor
type IsisShTopoNeighbor struct {
	// Neighbor ID
	NeighborId string `protobuf:"bytes,1,opt,name=neighbor_id,json=neighborId" json:"neighbor_id,omitempty"`
	// Pseudonode between system and its neighbor
	IntermediatePseudonode *IsisNodeIdType `protobuf:"bytes,2,opt,name=intermediate_pseudonode,json=intermediatePseudonode" json:"intermediate_pseudonode,omitempty"`
}

func (m *IsisShTopoNeighbor) Reset()                    { *m = IsisShTopoNeighbor{} }
func (m *IsisShTopoNeighbor) String() string            { return proto.CompactTextString(m) }
func (*IsisShTopoNeighbor) ProtoMessage()               {}
func (*IsisShTopoNeighbor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *IsisShTopoNeighbor) GetNeighborId() string {
	if m != nil {
		return m.NeighborId
	}
	return ""
}

func (m *IsisShTopoNeighbor) GetIntermediatePseudonode() *IsisNodeIdType {
	if m != nil {
		return m.IntermediatePseudonode
	}
	return nil
}

// Status of a reachable IPv6 IS
type IsisShIpv6TopoReachableDetails struct {
	// Distance to the IS
	RootDistance uint32 `protobuf:"varint,1,opt,name=root_distance,json=rootDistance" json:"root_distance,omitempty"`
	// Distance to the IS
	MulticastRootDistance uint32 `protobuf:"varint,2,opt,name=multicast_root_distance,json=multicastRootDistance" json:"multicast_root_distance,omitempty"`
	// First hops towards the IS
	Paths []*IsisShIpv6Path `protobuf:"bytes,3,rep,name=paths" json:"paths,omitempty"`
	// Multicast intact first hops towards the IS
	MulticastPaths []*IsisShIpv6Path `protobuf:"bytes,4,rep,name=multicast_paths,json=multicastPaths" json:"multicast_paths,omitempty"`
	// Parents of the IS within the SPT
	Parents []*IsisShTopoNeighbor `protobuf:"bytes,5,rep,name=parents" json:"parents,omitempty"`
	// Children of the IS within the SPT
	Children []*IsisShTopoNeighbor `protobuf:"bytes,6,rep,name=children" json:"children,omitempty"`
}

func (m *IsisShIpv6TopoReachableDetails) Reset()                    { *m = IsisShIpv6TopoReachableDetails{} }
func (m *IsisShIpv6TopoReachableDetails) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv6TopoReachableDetails) ProtoMessage()               {}
func (*IsisShIpv6TopoReachableDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *IsisShIpv6TopoReachableDetails) GetRootDistance() uint32 {
	if m != nil {
		return m.RootDistance
	}
	return 0
}

func (m *IsisShIpv6TopoReachableDetails) GetMulticastRootDistance() uint32 {
	if m != nil {
		return m.MulticastRootDistance
	}
	return 0
}

func (m *IsisShIpv6TopoReachableDetails) GetPaths() []*IsisShIpv6Path {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *IsisShIpv6TopoReachableDetails) GetMulticastPaths() []*IsisShIpv6Path {
	if m != nil {
		return m.MulticastPaths
	}
	return nil
}

func (m *IsisShIpv6TopoReachableDetails) GetParents() []*IsisShTopoNeighbor {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *IsisShIpv6TopoReachableDetails) GetChildren() []*IsisShTopoNeighbor {
	if m != nil {
		return m.Children
	}
	return nil
}

// Reachability status of an IPv6 IS
type IsisShIpv6TopoReachableStatus struct {
	ReachableStatus string `protobuf:"bytes,1,opt,name=reachable_status,json=reachableStatus" json:"reachable_status,omitempty"`
	// Status of the IS within the SPT
	ReachableDetails *IsisShIpv6TopoReachableDetails `protobuf:"bytes,2,opt,name=reachable_details,json=reachableDetails" json:"reachable_details,omitempty"`
}

func (m *IsisShIpv6TopoReachableStatus) Reset()                    { *m = IsisShIpv6TopoReachableStatus{} }
func (m *IsisShIpv6TopoReachableStatus) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv6TopoReachableStatus) ProtoMessage()               {}
func (*IsisShIpv6TopoReachableStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *IsisShIpv6TopoReachableStatus) GetReachableStatus() string {
	if m != nil {
		return m.ReachableStatus
	}
	return ""
}

func (m *IsisShIpv6TopoReachableStatus) GetReachableDetails() *IsisShIpv6TopoReachableDetails {
	if m != nil {
		return m.ReachableDetails
	}
	return nil
}

func init() {
	proto.RegisterType((*IsisShIpv6TopoEntry_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_ipv6_topo_entry_KEYS")
	proto.RegisterType((*IsisShIpv6TopoEntry)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_ipv6_topo_entry")
	proto.RegisterType((*IsisNodeIdType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_node_id_type")
	proto.RegisterType((*IsisSnpaType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_snpa_type")
	proto.RegisterType((*IsisPerPriorityCounts)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_per_priority_counts")
	proto.RegisterType((*IsisShRepEl)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_rep_el")
	proto.RegisterType((*IsisShIpv6FrrBackup)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_ipv6_frr_backup")
	proto.RegisterType((*IsisShIpv6Path)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_ipv6_path")
	proto.RegisterType((*IsisShTopoNeighbor)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_topo_neighbor")
	proto.RegisterType((*IsisShIpv6TopoReachableDetails)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_ipv6_topo_reachable_details")
	proto.RegisterType((*IsisShIpv6TopoReachableStatus)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_ipv6_topo_reachable_status")
}

func init() { proto.RegisterFile("isis_sh_ipv6_topo_entry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x58, 0xcd, 0x6f, 0x1c, 0xc5,
	0x12, 0xd7, 0xc4, 0xdf, 0xb5, 0x5e, 0x7b, 0xb7, 0xfd, 0xb1, 0x93, 0xbc, 0x17, 0x3d, 0xbf, 0xcd,
	0x7b, 0x91, 0x23, 0xd0, 0x22, 0x9c, 0xc4, 0x08, 0x71, 0x0a, 0xb1, 0x0f, 0x16, 0x51, 0xb0, 0x66,
	0x23, 0x24, 0x4e, 0xad, 0xf6, 0x4c, 0xef, 0x6e, 0x93, 0x99, 0xe9, 0xa1, 0xbb, 0xc7, 0xc9, 0xe6,
	0xca, 0x9d, 0x0b, 0x12, 0x47, 0x24, 0x0e, 0x08, 0x24, 0x6e, 0xe4, 0xc2, 0x11, 0x21, 0x24, 0x0e,
	0x08, 0xe5, 0x0f, 0xe0, 0x84, 0xc4, 0x31, 0xff, 0x02, 0x12, 0xea, 0xea, 0xf9, 0xb0, 0xd7, 0xd8,
	0x39, 0xb2, 0xbe, 0xf5, 0xfc, 0x7e, 0x55, 0xd3, 0x55, 0x3d, 0x55, 0xbf, 0xae, 0x5d, 0xb8, 0x2e,
	0xb4, 0xd0, 0x54, 0x8f, 0xa8, 0xc8, 0x8e, 0x77, 0xa9, 0x91, 0x99, 0xa4, 0x3c, 0x35, 0x6a, 0xdc,
	0xcb, 0x94, 0x34, 0x92, 0x7c, 0xea, 0x85, 0x42, 0x87, 0x92, 0x0a, 0xa9, 0xe9, 0x53, 0x45, 0xc3,
	0x38, 0xd5, 0x14, 0x3d, 0x64, 0xc6, 0x55, 0xcf, 0xae, 0x7a, 0x22, 0xd5, 0x86, 0xa5, 0x21, 0xaf,
	0x57, 0x3d, 0xfb, 0x9a, 0x58, 0x0e, 0x05, 0xd7, 0xe5, 0x72, 0x5c, 0x2d, 0x68, 0xcc, 0x8f, 0x79,
	0xac, 0x27, 0x9e, 0x7b, 0xb8, 0x7d, 0x2c, 0xd2, 0xc7, 0xf4, 0x84, 0xf3, 0x19, 0x70, 0xdc, 0x7d,
	0xe1, 0xc1, 0xbf, 0xcf, 0x09, 0x99, 0xbe, 0xb7, 0xff, 0x61, 0x9f, 0xdc, 0x80, 0x66, 0x19, 0x07,
	0x4d, 0x59, 0xc2, 0x7d, 0x6f, 0xcb, 0xdb, 0x5e, 0x0a, 0x96, 0x4b, 0xf0, 0x21, 0x4b, 0x38, 0xe9,
	0xc0, 0x02, 0x1b, 0x38, 0xfa, 0x0a, 0xd2, 0xf3, 0x6c, 0x80, 0xc4, 0x55, 0x58, 0xd4, 0x25, 0x33,
	0x83, 0xcc, 0x82, 0x2e, 0xa8, 0x1b, 0xd0, 0xac, 0x62, 0x46, 0x7e, 0xd6, 0xbd, 0xb8, 0x04, 0xd1,
	0x68, 0x1d, 0xe6, 0x30, 0x1f, 0x7f, 0x0e, 0x49, 0xf7, 0x40, 0xfe, 0x05, 0x4b, 0x7a, 0xac, 0x0d,
	0x4f, 0xa8, 0x88, 0xfc, 0x79, 0x64, 0x16, 0x1d, 0x70, 0x10, 0x75, 0x3f, 0x99, 0x83, 0xce, 0x39,
	0x19, 0x91, 0xff, 0xc3, 0x8a, 0x96, 0xb9, 0x0a, 0x39, 0x65, 0x51, 0xa4, 0xb8, 0xd6, 0xfe, 0x0e,
	0x7a, 0x37, 0x1d, 0x7a, 0xcf, 0x81, 0xd6, 0x4c, 0x68, 0x9a, 0x31, 0x65, 0x44, 0x28, 0x32, 0x96,
	0x1a, 0xff, 0xf6, 0x96, 0xb7, 0xbd, 0x18, 0x34, 0x85, 0x3e, 0xac, 0x41, 0x3c, 0x1a, 0x4d, 0xe5,
	0x31, 0x57, 0xb1, 0x64, 0x11, 0x8f, 0xfc, 0x3b, 0x68, 0xb5, 0x2c, 0xf4, 0xfb, 0x15, 0x46, 0xfe,
	0x03, 0x0d, 0xa1, 0x29, 0x33, 0x86, 0x85, 0x23, 0x1e, 0xf9, 0x77, 0xd1, 0x04, 0x84, 0xbe, 0x57,
	0x20, 0xe4, 0x0f, 0x0f, 0xd6, 0x14, 0x67, 0xe1, 0x88, 0x1d, 0x89, 0x58, 0x98, 0x31, 0xd5, 0x86,
	0x99, 0x5c, 0xfb, 0xbb, 0x5b, 0xde, 0x76, 0x63, 0xe7, 0x3b, 0xaf, 0x37, 0x5d, 0x15, 0xd3, 0x3b,
	0x7b, 0xb6, 0x45, 0xf4, 0x31, 0x2f, 0x42, 0x0f, 0xc8, 0xc9, 0x7c, 0xfa, 0x88, 0x91, 0x97, 0x1e,
	0x5c, 0x67, 0xd1, 0x31, 0x57, 0x46, 0x68, 0x1e, 0xd1, 0x4c, 0xf1, 0x81, 0x78, 0x4a, 0x85, 0xfd,
	0x88, 0xa1, 0xcc, 0x53, 0xa3, 0xfd, 0xb7, 0x30, 0xe1, 0x6f, 0xa6, 0x33, 0xe1, 0x8c, 0x2b, 0x9a,
	0x29, 0x21, 0x95, 0xfd, 0x42, 0x2e, 0xe0, 0xe0, 0x5a, 0x9d, 0xcf, 0x21, 0xa6, 0x73, 0x60, 0x78,
	0x72, 0x1f, 0xb9, 0xee, 0x2d, 0x68, 0xa3, 0x5f, 0x2a, 0x23, 0x4e, 0x45, 0x44, 0xcd, 0x38, 0xc3,
	0x6a, 0x3e, 0x66, 0x71, 0x5e, 0xf6, 0x90, 0x7b, 0xe8, 0xde, 0xb4, 0xd5, 0x66, 0xcf, 0x34, 0xcd,
	0xd8, 0x45, 0x76, 0x06, 0xfc, 0xf3, 0x42, 0x21, 0xd7, 0x60, 0x31, 0x54, 0xc2, 0x88, 0x90, 0xc5,
	0xe8, 0xd4, 0x0c, 0xaa, 0x67, 0x42, 0x60, 0x76, 0x24, 0x86, 0x23, 0xec, 0xcc, 0x66, 0x80, 0x6b,
	0xb2, 0x09, 0xf3, 0x09, 0x8f, 0x44, 0x9e, 0x60, 0x57, 0x36, 0x83, 0xe2, 0x89, 0xb4, 0x60, 0x26,
	0x96, 0x4f, 0xb0, 0x15, 0x9b, 0x81, 0x5d, 0x76, 0x5f, 0x7a, 0x65, 0x78, 0x23, 0xaa, 0x78, 0x46,
	0x79, 0x4c, 0x6e, 0xc3, 0xa6, 0xe2, 0x19, 0x13, 0x8a, 0xf2, 0x98, 0x27, 0x3c, 0x35, 0x65, 0x96,
	0x45, 0xbc, 0x6b, 0x8e, 0xdd, 0x77, 0xe4, 0x43, 0x19, 0xf1, 0x83, 0x88, 0x6c, 0x43, 0xab, 0x70,
	0x12, 0xd9, 0xf1, 0x1d, 0xec, 0xbf, 0x42, 0x2b, 0x56, 0x1c, 0x7e, 0x90, 0x1d, 0xdf, 0xb1, 0x0d,
	0x78, 0xda, 0x72, 0xd7, 0x59, 0xce, 0x4c, 0x58, 0xee, 0xa2, 0xe5, 0x7f, 0x61, 0xb9, 0xb0, 0x8c,
	0xd9, 0x11, 0x8f, 0x8b, 0xb0, 0x1b, 0x0e, 0x7b, 0x60, 0x21, 0xd2, 0x83, 0xb5, 0x89, 0x58, 0xed,
	0x09, 0xa3, 0x9c, 0x34, 0x83, 0xf6, 0xa9, 0x40, 0x1f, 0x8d, 0x33, 0xde, 0xfd, 0xb6, 0x31, 0xa1,
	0x1e, 0x03, 0xa5, 0xe8, 0x11, 0x0b, 0x1f, 0xe7, 0x99, 0x6d, 0xe5, 0x94, 0x8b, 0xe1, 0xe8, 0x48,
	0xaa, 0x3a, 0x59, 0x28, 0xa1, 0x83, 0x88, 0xdc, 0x82, 0x16, 0x1f, 0x5a, 0x05, 0xa1, 0x22, 0x35,
	0x5c, 0x0d, 0x58, 0x58, 0xea, 0xe1, 0xaa, 0xc3, 0x0f, 0x4a, 0xd8, 0x9a, 0x56, 0xef, 0x2a, 0xb5,
	0xc8, 0x25, 0xb9, 0x5a, 0xe2, 0xa5, 0x1a, 0xed, 0x42, 0xc7, 0xe4, 0x69, 0xca, 0x63, 0x7a, 0xe6,
	0xe5, 0x4e, 0x32, 0x37, 0x1c, 0xbd, 0x3f, 0xb1, 0xc5, 0x4f, 0x1e, 0x34, 0xab, 0x3d, 0x6c, 0x71,
	0x61, 0xd6, 0x8d, 0x9d, 0x2f, 0xa6, 0x54, 0x52, 0xca, 0xf2, 0x0f, 0x96, 0xcb, 0xa8, 0xfb, 0x69,
	0xc6, 0xc8, 0x1b, 0xb0, 0xae, 0x78, 0x22, 0x0d, 0xa7, 0xf1, 0x80, 0xd1, 0x49, 0xdd, 0x6f, 0x3b,
	0xee, 0xc1, 0x80, 0xf5, 0x8b, 0x0b, 0x60, 0xc2, 0x41, 0xc9, 0xdc, 0x70, 0xfc, 0x5e, 0x0b, 0x13,
	0x0e, 0x01, 0x32, 0x07, 0x11, 0x79, 0x13, 0x36, 0xce, 0xee, 0x90, 0x89, 0xc8, 0x5f, 0x44, 0x0f,
	0x32, 0xb1, 0xc5, 0xa1, 0x98, 0x74, 0x29, 0xf6, 0xb0, 0x2e, 0x4b, 0x13, 0x2e, 0x6e, 0x13, 0xeb,
	0xb2, 0x03, 0x1b, 0x46, 0x1a, 0x16, 0x17, 0xd5, 0x44, 0x23, 0xe1, 0x0e, 0xd6, 0x07, 0xac, 0xc5,
	0x35, 0x24, 0xdf, 0x45, 0x6e, 0xaf, 0xa0, 0xc8, 0xdb, 0x70, 0x55, 0xf3, 0x21, 0x96, 0xad, 0xdd,
	0x43, 0xa4, 0x43, 0xaa, 0x45, 0x44, 0x9d, 0x38, 0x34, 0xd0, 0x6f, 0xb3, 0x30, 0x08, 0x1c, 0xdf,
	0x17, 0xd1, 0x07, 0x96, 0xb5, 0x57, 0x72, 0x9a, 0x27, 0xd6, 0xdc, 0x5f, 0x76, 0x2d, 0x9e, 0xe6,
	0x49, 0x5f, 0x44, 0xe4, 0x1d, 0xb8, 0x76, 0xee, 0x3b, 0xb5, 0xdf, 0xdc, 0x9a, 0xd9, 0x6e, 0x06,
	0x9d, 0xbf, 0x7f, 0xa9, 0x26, 0x77, 0xa1, 0x53, 0x84, 0x5f, 0x36, 0x9e, 0xd0, 0x86, 0x6a, 0xf1,
	0x8c, 0xfb, 0x2b, 0xb8, 0xcb, 0xba, 0xa3, 0x03, 0xd7, 0x82, 0x42, 0x9b, 0xbe, 0x78, 0xc6, 0xc9,
	0x6b, 0xd0, 0x36, 0xc2, 0x9e, 0x54, 0x28, 0x93, 0x2c, 0x37, 0xcc, 0x08, 0x99, 0xfa, 0xab, 0x78,
	0x54, 0x2d, 0x24, 0xee, 0xd7, 0x38, 0xf9, 0xd5, 0x03, 0x72, 0x76, 0x13, 0xbf, 0xb5, 0x35, 0x33,
	0xc5, 0xc5, 0x5b, 0x8a, 0x63, 0xd0, 0x9a, 0x3c, 0x80, 0x62, 0x4c, 0x88, 0xe4, 0x93, 0x54, 0x1b,
	0xc5, 0x59, 0xe2, 0xb7, 0xcb, 0x31, 0x61, 0xaf, 0xc2, 0xc8, 0xff, 0x70, 0xe4, 0x88, 0x43, 0x5b,
	0x16, 0x1f, 0x49, 0x91, 0x1a, 0x9f, 0x94, 0x56, 0x0f, 0xc2, 0xbd, 0x02, 0x23, 0xaf, 0x03, 0x29,
	0xef, 0x14, 0x3b, 0x4f, 0xf2, 0xd0, 0x7e, 0x1d, 0x7f, 0x0d, 0x2d, 0x5b, 0x42, 0x5b, 0xa9, 0x3d,
	0xac, 0x70, 0x72, 0x13, 0x56, 0xed, 0x75, 0xa1, 0x44, 0xc2, 0xd4, 0x98, 0x66, 0xcc, 0x8c, 0xfc,
	0xf5, 0x6a, 0x8e, 0x71, 0xe8, 0x21, 0x33, 0x23, 0x2b, 0xb8, 0x36, 0x05, 0x15, 0x0f, 0xeb, 0xdd,
	0x37, 0xd0, 0x70, 0x45, 0xe8, 0xbe, 0x8a, 0x87, 0xd5, 0xfe, 0x5d, 0x4c, 0xa5, 0xae, 0x7c, 0x7f,
	0x13, 0xcd, 0x1a, 0x42, 0x07, 0x65, 0xc1, 0x17, 0x36, 0x3c, 0x0b, 0xad, 0x72, 0x5a, 0x9b, 0x4e,
	0x69, 0xb3, 0x8f, 0x98, 0xb5, 0xe9, 0xc1, 0x9a, 0xd0, 0xb4, 0x50, 0x35, 0xc5, 0x3f, 0xce, 0xb9,
	0x36, 0x3c, 0xf2, 0x7d, 0xb4, 0x6c, 0x0b, 0xfd, 0x08, 0x99, 0xa0, 0x24, 0xec, 0x75, 0xf5, 0xc4,
	0x6a, 0x82, 0xf1, 0xaf, 0xba, 0x5a, 0x76, 0x4f, 0xdd, 0x3f, 0xe7, 0x8b, 0x6b, 0xb6, 0x54, 0x6b,
	0x9b, 0xe4, 0x3f, 0xa5, 0xd3, 0x67, 0xf5, 0x76, 0xf6, 0x32, 0xea, 0x6d, 0x0b, 0x66, 0x0c, 0x1b,
	0x16, 0x37, 0xa4, 0x5d, 0x92, 0x9f, 0x3d, 0x80, 0xfa, 0x1a, 0x44, 0xe1, 0x6d, 0xec, 0x7c, 0x3d,
	0xdd, 0x83, 0x69, 0x1d, 0x6f, 0xb0, 0x34, 0x50, 0xca, 0xe9, 0x2a, 0x79, 0xe1, 0xc1, 0x5a, 0x1e,
	0x4b, 0x99, 0x51, 0xfe, 0x34, 0x8b, 0x45, 0x28, 0x8c, 0xd3, 0x96, 0x85, 0xcb, 0xa1, 0x2d, 0x6d,
	0x8c, 0x7d, 0xbf, 0x08, 0x1d, 0xc5, 0xe5, 0x16, 0xb4, 0x8a, 0x36, 0xaa, 0x4b, 0xd9, 0x5d, 0x5b,
	0xab, 0x0e, 0xaf, 0x4b, 0xf9, 0xc2, 0xcb, 0x64, 0xe9, 0xc2, 0xcb, 0xa4, 0xee, 0x3f, 0x38, 0xd5,
	0x7f, 0x9f, 0x5f, 0x81, 0x8d, 0x32, 0x46, 0xfc, 0x29, 0x50, 0x56, 0xd2, 0xab, 0x7b, 0xf0, 0x37,
	0x0f, 0x3a, 0x18, 0xb2, 0x9d, 0x3c, 0x99, 0xe1, 0x34, 0xd3, 0x3c, 0x8f, 0xa4, 0xd5, 0x36, 0xec,
	0xc5, 0xc6, 0xce, 0x97, 0xd3, 0xf9, 0x39, 0x4e, 0x4e, 0xf4, 0xc1, 0xe6, 0xc9, 0x14, 0x0e, 0xab,
	0x0c, 0xba, 0x9f, 0x2d, 0xc0, 0x8d, 0x8b, 0x7e, 0x28, 0x45, 0xdc, 0x30, 0x11, 0x6b, 0x7b, 0x37,
	0x28, 0x29, 0x4d, 0x3d, 0x0c, 0xb8, 0xe1, 0x7d, 0xd9, 0x82, 0xd5, 0x14, 0xb0, 0x0b, 0x9d, 0x24,
	0x8f, 0xed, 0x30, 0xaf, 0xed, 0xa7, 0x3b, 0x69, 0xee, 0x66, 0xfa, 0x8d, 0x8a, 0x0e, 0x4e, 0xfa,
	0x7d, 0xef, 0xc1, 0x9c, 0x15, 0x44, 0xab, 0x58, 0x33, 0xd3, 0x7b, 0xa0, 0x27, 0xb5, 0x3b, 0x70,
	0x01, 0x93, 0x5f, 0x3c, 0x58, 0xad, 0x73, 0x76, 0x49, 0xcc, 0x5e, 0x9a, 0x24, 0x56, 0xaa, 0xd0,
	0x0f, 0x31, 0x9b, 0x1f, 0x3c, 0x58, 0xc8, 0x98, 0xe2, 0xf6, 0x57, 0xee, 0x1c, 0x66, 0xf1, 0xd5,
	0xd4, 0x66, 0x71, 0xaa, 0x8d, 0x83, 0x32, 0x6c, 0xf2, 0xa3, 0x07, 0x8b, 0xe1, 0x48, 0xc4, 0x91,
	0xe2, 0xa9, 0x3f, 0x7f, 0xa9, 0x72, 0xa8, 0xe2, 0xee, 0x3e, 0xbf, 0x02, 0xdd, 0x57, 0xff, 0x7d,
	0x61, 0x35, 0x75, 0x12, 0x2b, 0x04, 0x6c, 0xb5, 0xc2, 0x8b, 0x7f, 0x35, 0x7e, 0xf7, 0xa0, 0x7d,
	0xa6, 0xab, 0x0b, 0xfd, 0x7a, 0x7e, 0xa9, 0xfe, 0xba, 0x29, 0x62, 0x0f, 0xea, 0xd4, 0xf7, 0x1c,
	0x72, 0x34, 0x8f, 0x7f, 0x5d, 0xde, 0xfe, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x6e, 0xdc, 0x65,
	0xdb, 0x14, 0x00, 0x00,
}