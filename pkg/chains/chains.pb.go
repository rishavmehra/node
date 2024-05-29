// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/pkg/chains/chains.proto

package chains

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ReceiveStatus represents the status of an outbound
// TODO: Rename and move
// https://github.com/zeta-chain/node/issues/2257
type ReceiveStatus int32

const (
	// Created is used for inbounds
	ReceiveStatus_created ReceiveStatus = 0
	ReceiveStatus_success ReceiveStatus = 1
	ReceiveStatus_failed  ReceiveStatus = 2
)

var ReceiveStatus_name = map[int32]string{
	0: "created",
	1: "success",
	2: "failed",
}

var ReceiveStatus_value = map[string]int32{
	"created": 0,
	"success": 1,
	"failed":  2,
}

func (x ReceiveStatus) String() string {
	return proto.EnumName(ReceiveStatus_name, int32(x))
}

func (ReceiveStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_236b85e7bff6130d, []int{0}
}

// ChainName represents the name of the chain
type ChainName int32

const (
	ChainName_empty            ChainName = 0
	ChainName_eth_mainnet      ChainName = 1
	ChainName_zeta_mainnet     ChainName = 2
	ChainName_btc_mainnet      ChainName = 3
	ChainName_polygon_mainnet  ChainName = 4
	ChainName_bsc_mainnet      ChainName = 5
	ChainName_goerli_testnet   ChainName = 6
	ChainName_mumbai_testnet   ChainName = 7
	ChainName_bsc_testnet      ChainName = 10
	ChainName_zeta_testnet     ChainName = 11
	ChainName_btc_testnet      ChainName = 12
	ChainName_sepolia_testnet  ChainName = 13
	ChainName_goerli_localnet  ChainName = 14
	ChainName_btc_regtest      ChainName = 15
	ChainName_amoy_testnet     ChainName = 16
	ChainName_optimism_mainnet ChainName = 17
	ChainName_optimism_sepolia ChainName = 18
	ChainName_base_mainnet     ChainName = 19
	ChainName_base_sepolia     ChainName = 20
)

var ChainName_name = map[int32]string{
	0:  "empty",
	1:  "eth_mainnet",
	2:  "zeta_mainnet",
	3:  "btc_mainnet",
	4:  "polygon_mainnet",
	5:  "bsc_mainnet",
	6:  "goerli_testnet",
	7:  "mumbai_testnet",
	10: "bsc_testnet",
	11: "zeta_testnet",
	12: "btc_testnet",
	13: "sepolia_testnet",
	14: "goerli_localnet",
	15: "btc_regtest",
	16: "amoy_testnet",
	17: "optimism_mainnet",
	18: "optimism_sepolia",
	19: "base_mainnet",
	20: "base_sepolia",
}

var ChainName_value = map[string]int32{
	"empty":            0,
	"eth_mainnet":      1,
	"zeta_mainnet":     2,
	"btc_mainnet":      3,
	"polygon_mainnet":  4,
	"bsc_mainnet":      5,
	"goerli_testnet":   6,
	"mumbai_testnet":   7,
	"bsc_testnet":      10,
	"zeta_testnet":     11,
	"btc_testnet":      12,
	"sepolia_testnet":  13,
	"goerli_localnet":  14,
	"btc_regtest":      15,
	"amoy_testnet":     16,
	"optimism_mainnet": 17,
	"optimism_sepolia": 18,
	"base_mainnet":     19,
	"base_sepolia":     20,
}

func (x ChainName) String() string {
	return proto.EnumName(ChainName_name, int32(x))
}

func (ChainName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_236b85e7bff6130d, []int{1}
}

// Network represents the network of the chain
// there is a single instance of the network on mainnet
// then the network can have eventual testnets or devnets
type Network int32

const (
	Network_eth      Network = 0
	Network_zeta     Network = 1
	Network_btc      Network = 2
	Network_polygon  Network = 3
	Network_bsc      Network = 4
	Network_optimism Network = 5
	Network_base     Network = 6
)

var Network_name = map[int32]string{
	0: "eth",
	1: "zeta",
	2: "btc",
	3: "polygon",
	4: "bsc",
	5: "optimism",
	6: "base",
}

var Network_value = map[string]int32{
	"eth":      0,
	"zeta":     1,
	"btc":      2,
	"polygon":  3,
	"bsc":      4,
	"optimism": 5,
	"base":     6,
}

func (x Network) String() string {
	return proto.EnumName(Network_name, int32(x))
}

func (Network) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_236b85e7bff6130d, []int{2}
}

// NetworkType represents the network type of the chain
// Mainnet, Testnet, Privnet, Devnet
type NetworkType int32

const (
	NetworkType_mainnet NetworkType = 0
	NetworkType_testnet NetworkType = 1
	NetworkType_privnet NetworkType = 2
	NetworkType_devnet  NetworkType = 3
)

var NetworkType_name = map[int32]string{
	0: "mainnet",
	1: "testnet",
	2: "privnet",
	3: "devnet",
}

var NetworkType_value = map[string]int32{
	"mainnet": 0,
	"testnet": 1,
	"privnet": 2,
	"devnet":  3,
}

func (x NetworkType) String() string {
	return proto.EnumName(NetworkType_name, int32(x))
}

func (NetworkType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_236b85e7bff6130d, []int{3}
}

// Vm represents the virtual machine type of the chain to support smart
// contracts
type Vm int32

const (
	Vm_no_vm Vm = 0
	Vm_evm   Vm = 1
)

var Vm_name = map[int32]string{
	0: "no_vm",
	1: "evm",
}

var Vm_value = map[string]int32{
	"no_vm": 0,
	"evm":   1,
}

func (x Vm) String() string {
	return proto.EnumName(Vm_name, int32(x))
}

func (Vm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_236b85e7bff6130d, []int{4}
}

// Consensus represents the consensus algorithm used by the chain
// this can represent the consensus of a L1
// this can also represent the solution of a L2
type Consensus int32

const (
	Consensus_ethereum   Consensus = 0
	Consensus_tendermint Consensus = 1
	Consensus_bitcoin    Consensus = 2
	Consensus_op_stack   Consensus = 3
)

var Consensus_name = map[int32]string{
	0: "ethereum",
	1: "tendermint",
	2: "bitcoin",
	3: "op_stack",
}

var Consensus_value = map[string]int32{
	"ethereum":   0,
	"tendermint": 1,
	"bitcoin":    2,
	"op_stack":   3,
}

func (x Consensus) String() string {
	return proto.EnumName(Consensus_name, int32(x))
}

func (Consensus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_236b85e7bff6130d, []int{5}
}

// CCTXGateway describes for the chain the gateway used to handle CCTX outbounds
type CCTXGateway int32

const (
	// zevm is the internal CCTX gateway to process outbound on the ZEVM and read
	// inbound events from the ZEVM only used for ZetaChain chains
	CCTXGateway_zevm CCTXGateway = 0
	// observers is the CCTX gateway for chains relying on the observer set to
	// observe inbounds and TSS for outbounds
	CCTXGateway_observers CCTXGateway = 1
)

var CCTXGateway_name = map[int32]string{
	0: "zevm",
	1: "observers",
}

var CCTXGateway_value = map[string]int32{
	"zevm":      0,
	"observers": 1,
}

func (x CCTXGateway) String() string {
	return proto.EnumName(CCTXGateway_name, int32(x))
}

func (CCTXGateway) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_236b85e7bff6130d, []int{6}
}

// Chain represents static data about a blockchain network
// it is identified by a unique chain ID
type Chain struct {
	// ChainId is the unique identifier of the chain
	ChainId int64 `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// ChainName is the name of the chain
	ChainName ChainName `protobuf:"varint,1,opt,name=chain_name,json=chainName,proto3,enum=zetachain.zetacore.pkg.chains.ChainName" json:"chain_name,omitempty"`
	// Network is the network of the chain
	Network Network `protobuf:"varint,3,opt,name=network,proto3,enum=zetachain.zetacore.pkg.chains.Network" json:"network,omitempty"`
	// NetworkType is the network type of the chain: mainnet, testnet, etc..
	NetworkType NetworkType `protobuf:"varint,4,opt,name=network_type,json=networkType,proto3,enum=zetachain.zetacore.pkg.chains.NetworkType" json:"network_type,omitempty"`
	// Vm is the virtual machine used in the chain
	Vm Vm `protobuf:"varint,5,opt,name=vm,proto3,enum=zetachain.zetacore.pkg.chains.Vm" json:"vm,omitempty"`
	// Consensus is the underlying consensus algorithm used by the chain
	Consensus Consensus `protobuf:"varint,6,opt,name=consensus,proto3,enum=zetachain.zetacore.pkg.chains.Consensus" json:"consensus,omitempty"`
	// IsExternal describe if the chain is ZetaChain or external
	IsExternal bool `protobuf:"varint,7,opt,name=is_external,json=isExternal,proto3" json:"is_external,omitempty"`
	// CCTXGateway is the gateway used to handle CCTX outbounds
	CctxGateway CCTXGateway `protobuf:"varint,8,opt,name=cctx_gateway,json=cctxGateway,proto3,enum=zetachain.zetacore.pkg.chains.CCTXGateway" json:"cctx_gateway,omitempty"`
}

func (m *Chain) Reset()         { *m = Chain{} }
func (m *Chain) String() string { return proto.CompactTextString(m) }
func (*Chain) ProtoMessage()    {}
func (*Chain) Descriptor() ([]byte, []int) {
	return fileDescriptor_236b85e7bff6130d, []int{0}
}
func (m *Chain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Chain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Chain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Chain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chain.Merge(m, src)
}
func (m *Chain) XXX_Size() int {
	return m.Size()
}
func (m *Chain) XXX_DiscardUnknown() {
	xxx_messageInfo_Chain.DiscardUnknown(m)
}

var xxx_messageInfo_Chain proto.InternalMessageInfo

func (m *Chain) GetChainId() int64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *Chain) GetChainName() ChainName {
	if m != nil {
		return m.ChainName
	}
	return ChainName_empty
}

func (m *Chain) GetNetwork() Network {
	if m != nil {
		return m.Network
	}
	return Network_eth
}

func (m *Chain) GetNetworkType() NetworkType {
	if m != nil {
		return m.NetworkType
	}
	return NetworkType_mainnet
}

func (m *Chain) GetVm() Vm {
	if m != nil {
		return m.Vm
	}
	return Vm_no_vm
}

func (m *Chain) GetConsensus() Consensus {
	if m != nil {
		return m.Consensus
	}
	return Consensus_ethereum
}

func (m *Chain) GetIsExternal() bool {
	if m != nil {
		return m.IsExternal
	}
	return false
}

func (m *Chain) GetCctxGateway() CCTXGateway {
	if m != nil {
		return m.CctxGateway
	}
	return CCTXGateway_zevm
}

func init() {
	proto.RegisterEnum("zetachain.zetacore.pkg.chains.ReceiveStatus", ReceiveStatus_name, ReceiveStatus_value)
	proto.RegisterEnum("zetachain.zetacore.pkg.chains.ChainName", ChainName_name, ChainName_value)
	proto.RegisterEnum("zetachain.zetacore.pkg.chains.Network", Network_name, Network_value)
	proto.RegisterEnum("zetachain.zetacore.pkg.chains.NetworkType", NetworkType_name, NetworkType_value)
	proto.RegisterEnum("zetachain.zetacore.pkg.chains.Vm", Vm_name, Vm_value)
	proto.RegisterEnum("zetachain.zetacore.pkg.chains.Consensus", Consensus_name, Consensus_value)
	proto.RegisterEnum("zetachain.zetacore.pkg.chains.CCTXGateway", CCTXGateway_name, CCTXGateway_value)
	proto.RegisterType((*Chain)(nil), "zetachain.zetacore.pkg.chains.Chain")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/pkg/chains/chains.proto", fileDescriptor_236b85e7bff6130d)
}

var fileDescriptor_236b85e7bff6130d = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x3d, 0x8f, 0x2b, 0x35,
	0x14, 0xcd, 0x4c, 0xbe, 0xef, 0x64, 0xb3, 0xc6, 0x6f, 0x8b, 0xe1, 0x49, 0x0c, 0x0b, 0x05, 0x0a,
	0x11, 0x24, 0x02, 0x4a, 0x1a, 0x44, 0xc4, 0x7b, 0x02, 0x89, 0x57, 0x0c, 0xab, 0x15, 0xd0, 0x8c,
	0x3c, 0xce, 0x65, 0x62, 0x25, 0x1e, 0x8f, 0xc6, 0x4e, 0x76, 0xc3, 0xaf, 0xe0, 0x47, 0x50, 0xf0,
	0x53, 0x28, 0xb7, 0xa4, 0x44, 0xbb, 0x05, 0x25, 0x7f, 0x01, 0xd9, 0xf3, 0x91, 0xa5, 0x80, 0xdd,
	0x2a, 0xd7, 0x67, 0xce, 0x39, 0xf7, 0xd8, 0xbe, 0x31, 0xcc, 0x7f, 0x46, 0xc3, 0xf8, 0x86, 0x89,
	0x7c, 0xe9, 0x2a, 0x55, 0xe2, 0xb2, 0xd8, 0x66, 0x4b, 0x07, 0xe9, 0xfa, 0x67, 0x51, 0x94, 0xca,
	0x28, 0xfa, 0x4e, 0xcb, 0x5d, 0x34, 0xdc, 0x45, 0xb1, 0xcd, 0x16, 0x15, 0xe9, 0xe5, 0x45, 0xa6,
	0x32, 0xe5, 0x98, 0x4b, 0x5b, 0x55, 0xa2, 0xf7, 0xff, 0xea, 0x42, 0x7f, 0x65, 0x09, 0xf4, 0x6d,
	0x18, 0x39, 0x66, 0x22, 0xd6, 0xa1, 0x7f, 0xe9, 0xcd, 0xba, 0xf1, 0xd0, 0xad, 0xbf, 0x5e, 0xd3,
	0xd7, 0x00, 0xd5, 0xa7, 0x9c, 0x49, 0x0c, 0xbd, 0x4b, 0x6f, 0x36, 0xfd, 0x74, 0xb6, 0xf8, 0xdf,
	0x76, 0x0b, 0x67, 0xfa, 0x86, 0x49, 0x8c, 0xc7, 0xbc, 0x29, 0xe9, 0x17, 0x30, 0xcc, 0xd1, 0xdc,
	0xa8, 0x72, 0x1b, 0x76, 0x9d, 0xcb, 0x07, 0x4f, 0xb8, 0xbc, 0xa9, 0xd8, 0x71, 0x23, 0xa3, 0xdf,
	0xc2, 0xa4, 0x2e, 0x13, 0x73, 0x2c, 0x30, 0xec, 0x39, 0x9b, 0xf9, 0xf3, 0x6c, 0xae, 0x8e, 0x05,
	0xc6, 0x41, 0x7e, 0x5a, 0xd0, 0x4f, 0xc0, 0x3f, 0xc8, 0xb0, 0xef, 0x4c, 0xde, 0x7b, 0xc2, 0xe4,
	0x5a, 0xc6, 0xfe, 0x41, 0xd2, 0x57, 0x30, 0xe6, 0x2a, 0xd7, 0x98, 0xeb, 0xbd, 0x0e, 0x07, 0xcf,
	0x3b, 0x8b, 0x86, 0x1f, 0x9f, 0xa4, 0xf4, 0x5d, 0x08, 0x84, 0x4e, 0xf0, 0xd6, 0x60, 0x99, 0xb3,
	0x5d, 0x38, 0xbc, 0xf4, 0x66, 0xa3, 0x18, 0x84, 0xfe, 0xaa, 0x46, 0xec, 0x56, 0x39, 0x37, 0xb7,
	0x49, 0xc6, 0x0c, 0xde, 0xb0, 0x63, 0x38, 0x7a, 0xd6, 0x56, 0x57, 0xab, 0xab, 0xef, 0x5f, 0x57,
	0x8a, 0x38, 0xb0, 0xfa, 0x7a, 0x31, 0xff, 0x1c, 0xce, 0x62, 0xe4, 0x28, 0x0e, 0xf8, 0x9d, 0x61,
	0x66, 0xaf, 0x69, 0x00, 0x43, 0x5e, 0x22, 0x33, 0xb8, 0x26, 0x1d, 0xbb, 0xd0, 0x7b, 0xce, 0x51,
	0x6b, 0xe2, 0x51, 0x80, 0xc1, 0x4f, 0x4c, 0xec, 0x70, 0x4d, 0xfc, 0x97, 0xbd, 0xdf, 0x7e, 0x8d,
	0xbc, 0xf9, 0xdf, 0x3e, 0x8c, 0xdb, 0x1b, 0xa5, 0x63, 0xe8, 0xa3, 0x2c, 0xcc, 0x91, 0x74, 0xe8,
	0x39, 0x04, 0x68, 0x36, 0x89, 0x64, 0x22, 0xcf, 0xd1, 0x10, 0x8f, 0x12, 0x98, 0xd8, 0x58, 0x2d,
	0xe2, 0x5b, 0x4a, 0x6a, 0x78, 0x0b, 0x74, 0xe9, 0x0b, 0x38, 0x2f, 0xd4, 0xee, 0x98, 0xa9, 0xbc,
	0x05, 0x7b, 0x8e, 0xa5, 0x4f, 0xac, 0x3e, 0xa5, 0x30, 0xcd, 0x14, 0x96, 0x3b, 0x91, 0x18, 0xd4,
	0xc6, 0x62, 0x03, 0x8b, 0xc9, 0xbd, 0x4c, 0xd9, 0x09, 0x1b, 0x36, 0xc2, 0x06, 0x80, 0x36, 0x41,
	0x83, 0x04, 0x4d, 0x82, 0x06, 0x98, 0xd8, 0x04, 0x1a, 0x0b, 0xb5, 0x13, 0x27, 0xd6, 0x99, 0x05,
	0xeb, 0x86, 0x3b, 0xc5, 0xd9, 0xce, 0x82, 0xd3, 0x46, 0x5a, 0x62, 0x66, 0x89, 0xe4, 0xdc, 0xba,
	0x33, 0xa9, 0x8e, 0xad, 0x8e, 0xd0, 0x0b, 0x20, 0xaa, 0x30, 0x42, 0x0a, 0x2d, 0xdb, 0xf8, 0x6f,
	0xfd, 0x0b, 0xad, 0x7b, 0x11, 0x6a, 0xd5, 0x29, 0xd3, 0xd8, 0xf2, 0x5e, 0xb4, 0x48, 0xc3, 0xb9,
	0xa8, 0x4f, 0xfc, 0x07, 0x18, 0xd6, 0x53, 0x4b, 0x87, 0xd0, 0x45, 0xb3, 0x21, 0x1d, 0x3a, 0x82,
	0x9e, 0xdd, 0x19, 0xf1, 0x2c, 0x94, 0x1a, 0x4e, 0x7c, 0x7b, 0x6f, 0xf5, 0x59, 0x92, 0xae, 0x43,
	0x35, 0x27, 0x3d, 0x3a, 0x81, 0x51, 0xd3, 0x9c, 0xf4, 0xad, 0xcc, 0xb6, 0x20, 0x83, 0xda, 0xfa,
	0x15, 0x04, 0x8f, 0xfe, 0x10, 0xd6, 0xa2, 0x89, 0xe3, 0xe6, 0xa0, 0xd9, 0x99, 0xe7, 0xcc, 0x4b,
	0x71, 0xa8, 0xae, 0x11, 0x60, 0xb0, 0x46, 0x57, 0x77, 0x6b, 0x9f, 0x08, 0xfc, 0x6b, 0x69, 0x87,
	0x21, 0x57, 0xc9, 0x41, 0x92, 0x8e, 0x0b, 0x7a, 0x90, 0xc4, 0xab, 0xbf, 0x7f, 0x03, 0xe3, 0x76,
	0xf2, 0x6d, 0x24, 0x34, 0x1b, 0x2c, 0x71, 0x6f, 0x99, 0x53, 0x00, 0x83, 0xf9, 0x1a, 0x4b, 0x29,
	0xf2, 0xba, 0x53, 0x2a, 0x0c, 0x57, 0x22, 0x27, 0x7e, 0x95, 0x3e, 0xd1, 0x86, 0xf1, 0x6d, 0xdb,
	0xeb, 0x23, 0x08, 0x1e, 0x4d, 0x76, 0x75, 0x12, 0xae, 0xe7, 0x19, 0x8c, 0x55, 0xaa, 0xb1, 0x3c,
	0x60, 0xa9, 0x9b, 0xce, 0x5f, 0xae, 0x7e, 0xbf, 0x8f, 0xbc, 0xbb, 0xfb, 0xc8, 0xfb, 0xf3, 0x3e,
	0xf2, 0x7e, 0x79, 0x88, 0x3a, 0x77, 0x0f, 0x51, 0xe7, 0x8f, 0x87, 0xa8, 0xf3, 0xe3, 0x87, 0x99,
	0x30, 0x9b, 0x7d, 0xba, 0xe0, 0x4a, 0xba, 0x17, 0xf5, 0xe3, 0xff, 0x7c, 0x5c, 0xd3, 0x81, 0x7b,
	0x21, 0x3f, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x05, 0xa7, 0x7f, 0x40, 0x84, 0x05, 0x00, 0x00,
}

func (m *Chain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Chain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Chain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CctxGateway != 0 {
		i = encodeVarintChains(dAtA, i, uint64(m.CctxGateway))
		i--
		dAtA[i] = 0x40
	}
	if m.IsExternal {
		i--
		if m.IsExternal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.Consensus != 0 {
		i = encodeVarintChains(dAtA, i, uint64(m.Consensus))
		i--
		dAtA[i] = 0x30
	}
	if m.Vm != 0 {
		i = encodeVarintChains(dAtA, i, uint64(m.Vm))
		i--
		dAtA[i] = 0x28
	}
	if m.NetworkType != 0 {
		i = encodeVarintChains(dAtA, i, uint64(m.NetworkType))
		i--
		dAtA[i] = 0x20
	}
	if m.Network != 0 {
		i = encodeVarintChains(dAtA, i, uint64(m.Network))
		i--
		dAtA[i] = 0x18
	}
	if m.ChainId != 0 {
		i = encodeVarintChains(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x10
	}
	if m.ChainName != 0 {
		i = encodeVarintChains(dAtA, i, uint64(m.ChainName))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintChains(dAtA []byte, offset int, v uint64) int {
	offset -= sovChains(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Chain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainName != 0 {
		n += 1 + sovChains(uint64(m.ChainName))
	}
	if m.ChainId != 0 {
		n += 1 + sovChains(uint64(m.ChainId))
	}
	if m.Network != 0 {
		n += 1 + sovChains(uint64(m.Network))
	}
	if m.NetworkType != 0 {
		n += 1 + sovChains(uint64(m.NetworkType))
	}
	if m.Vm != 0 {
		n += 1 + sovChains(uint64(m.Vm))
	}
	if m.Consensus != 0 {
		n += 1 + sovChains(uint64(m.Consensus))
	}
	if m.IsExternal {
		n += 2
	}
	if m.CctxGateway != 0 {
		n += 1 + sovChains(uint64(m.CctxGateway))
	}
	return n
}

func sovChains(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChains(x uint64) (n int) {
	return sovChains(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Chain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChains
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Chain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Chain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			m.ChainName = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainName |= ChainName(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			m.Network = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Network |= Network(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkType", wireType)
			}
			m.NetworkType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NetworkType |= NetworkType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vm", wireType)
			}
			m.Vm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vm |= Vm(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consensus", wireType)
			}
			m.Consensus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Consensus |= Consensus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsExternal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsExternal = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CctxGateway", wireType)
			}
			m.CctxGateway = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CctxGateway |= CCTXGateway(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChains(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChains
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipChains(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChains
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowChains
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowChains
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthChains
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChains
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChains
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChains        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChains          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChains = fmt.Errorf("proto: unexpected end of group")
)
