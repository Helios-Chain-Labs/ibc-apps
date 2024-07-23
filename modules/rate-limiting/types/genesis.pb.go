// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ratelimit/v1/genesis.proto

package types

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

// GenesisState defines the ratelimit module's genesis state.
type GenesisState struct {
	Params                           Params                   `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params"`
	RateLimits                       []RateLimit              `protobuf:"bytes,2,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits" yaml:"rate_limits"`
	WhitelistedAddressPairs          []WhitelistedAddressPair `protobuf:"bytes,3,rep,name=whitelisted_address_pairs,json=whitelistedAddressPairs,proto3" json:"whitelisted_address_pairs" yaml:"whitelisted_address_pairs"`
	BlacklistedDenoms                []string                 `protobuf:"bytes,4,rep,name=blacklisted_denoms,json=blacklistedDenoms,proto3" json:"blacklisted_denoms,omitempty"`
	PendingSendPacketSequenceNumbers []string                 `protobuf:"bytes,5,rep,name=pending_send_packet_sequence_numbers,json=pendingSendPacketSequenceNumbers,proto3" json:"pending_send_packet_sequence_numbers,omitempty"`
	HourEpoch                        HourEpoch                `protobuf:"bytes,6,opt,name=hour_epoch,json=hourEpoch,proto3" json:"hour_epoch" yaml:"hour_epoch"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbb08d6119688a03, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetRateLimits() []RateLimit {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

func (m *GenesisState) GetWhitelistedAddressPairs() []WhitelistedAddressPair {
	if m != nil {
		return m.WhitelistedAddressPairs
	}
	return nil
}

func (m *GenesisState) GetBlacklistedDenoms() []string {
	if m != nil {
		return m.BlacklistedDenoms
	}
	return nil
}

func (m *GenesisState) GetPendingSendPacketSequenceNumbers() []string {
	if m != nil {
		return m.PendingSendPacketSequenceNumbers
	}
	return nil
}

func (m *GenesisState) GetHourEpoch() HourEpoch {
	if m != nil {
		return m.HourEpoch
	}
	return HourEpoch{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ratelimit.v1.GenesisState")
}

func init() { proto.RegisterFile("ratelimit/v1/genesis.proto", fileDescriptor_fbb08d6119688a03) }

var fileDescriptor_fbb08d6119688a03 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xb1, 0x6e, 0xd3, 0x40,
	0x1c, 0xc6, 0x63, 0x52, 0x22, 0xf5, 0x52, 0x86, 0x5a, 0x45, 0x75, 0x2c, 0xe4, 0x5a, 0x56, 0x87,
	0x2c, 0xb1, 0xd5, 0x32, 0x20, 0xd8, 0x08, 0x20, 0x18, 0x50, 0x15, 0x1c, 0x24, 0x24, 0x16, 0xeb,
	0x6c, 0xff, 0x65, 0x9f, 0x1a, 0xfb, 0x8e, 0xfb, 0x9f, 0x53, 0xf5, 0x0d, 0x10, 0x13, 0x8f, 0xd5,
	0xb1, 0x23, 0x53, 0x85, 0x92, 0x37, 0xe0, 0x09, 0x90, 0xef, 0x4c, 0x93, 0x20, 0xd8, 0x6c, 0x7d,
	0xbf, 0xdf, 0xf7, 0xe9, 0xec, 0x23, 0xae, 0xa4, 0x0a, 0x16, 0xac, 0x62, 0x2a, 0x5a, 0x9e, 0x45,
	0x05, 0xd4, 0x80, 0x0c, 0x43, 0x21, 0xb9, 0xe2, 0xf6, 0xc1, 0x7d, 0x16, 0x2e, 0xcf, 0xdc, 0xa3,
	0x82, 0x17, 0x5c, 0x07, 0x51, 0xfb, 0x64, 0x18, 0x77, 0xb4, 0xe3, 0x0b, 0x2a, 0x69, 0xd5, 0xe9,
	0xee, 0x93, 0x9d, 0x68, 0xd3, 0xa5, 0xd3, 0xe0, 0xeb, 0x1e, 0x39, 0x78, 0x6b, 0xe6, 0xe6, 0x8a,
	0x2a, 0xb0, 0x5f, 0x91, 0x81, 0xd1, 0x1d, 0xcb, 0xb7, 0xc6, 0xc3, 0xf3, 0xa3, 0x70, 0x7b, 0x3e,
	0x9c, 0xe9, 0x6c, 0xfa, 0xf8, 0xe6, 0xee, 0xa4, 0xf7, 0xeb, 0xee, 0xe4, 0xd1, 0x35, 0xad, 0x16,
	0x2f, 0x02, 0x63, 0x04, 0x71, 0xa7, 0xda, 0x1f, 0xc9, 0xb0, 0xb5, 0x12, 0xad, 0xa1, 0xf3, 0xc0,
	0xef, 0x8f, 0x87, 0xe7, 0xc7, 0xbb, 0x4d, 0x31, 0x55, 0xf0, 0xbe, 0x7d, 0x99, 0xba, 0x5d, 0x99,
	0x6d, 0xca, 0xb6, 0xcc, 0x20, 0x26, 0xf2, 0x0f, 0x86, 0xf6, 0x37, 0x8b, 0x8c, 0xae, 0x4a, 0xd6,
	0x76, 0xa0, 0x82, 0x3c, 0xa1, 0x79, 0x2e, 0x01, 0x31, 0x11, 0x94, 0x49, 0x74, 0xfa, 0x7a, 0xe4,
	0x74, 0x77, 0xe4, 0xd3, 0x06, 0x7f, 0x69, 0xe8, 0x19, 0x65, 0x72, 0x3a, 0xee, 0x16, 0x7d, 0xb3,
	0xf8, 0xdf, 0xd2, 0x20, 0x3e, 0xbe, 0xfa, 0x67, 0x03, 0xda, 0x13, 0x62, 0xa7, 0x0b, 0x9a, 0x5d,
	0x76, 0x5a, 0x0e, 0x35, 0xaf, 0xd0, 0xd9, 0xf3, 0xfb, 0xe3, 0xfd, 0xf8, 0x70, 0x2b, 0x79, 0xad,
	0x03, 0xfb, 0x82, 0x9c, 0x0a, 0xa8, 0x73, 0x56, 0x17, 0x09, 0x42, 0x9d, 0x27, 0x82, 0x66, 0x97,
	0xa0, 0x12, 0x84, 0x2f, 0x0d, 0xd4, 0x19, 0x24, 0x75, 0x53, 0xa5, 0x20, 0xd1, 0x79, 0xa8, 0x0b,
	0xfc, 0x8e, 0x9d, 0x43, 0x9d, 0xcf, 0x34, 0x39, 0xef, 0xc0, 0x0b, 0xc3, 0xd9, 0x1f, 0x08, 0x29,
	0x79, 0x23, 0x13, 0x10, 0x3c, 0x2b, 0x9d, 0x81, 0xfe, 0x55, 0x7f, 0x7d, 0xe0, 0x77, 0xbc, 0x91,
	0x6f, 0xda, 0x78, 0x3a, 0xea, 0x8e, 0x7b, 0x68, 0x8e, 0xbb, 0x11, 0x83, 0x78, 0xbf, 0xbc, 0xa7,
	0xe6, 0x37, 0x2b, 0xcf, 0xba, 0x5d, 0x79, 0xd6, 0xcf, 0x95, 0x67, 0x7d, 0x5f, 0x7b, 0xbd, 0xdb,
	0xb5, 0xd7, 0xfb, 0xb1, 0xf6, 0x7a, 0x9f, 0x9f, 0x17, 0x4c, 0x95, 0x4d, 0x1a, 0x66, 0xbc, 0x8a,
	0x32, 0x8e, 0x15, 0xc7, 0x88, 0xa5, 0xd9, 0x84, 0x0a, 0x81, 0x51, 0xc5, 0xf3, 0x66, 0x01, 0xa8,
	0x2f, 0xd6, 0x44, 0x6f, 0xb3, 0xba, 0x88, 0x96, 0xcf, 0x22, 0x75, 0x2d, 0x00, 0xd3, 0x81, 0xbe,
	0x66, 0x4f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x77, 0xf2, 0xc9, 0xad, 0xe1, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.HourEpoch.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.PendingSendPacketSequenceNumbers) > 0 {
		for iNdEx := len(m.PendingSendPacketSequenceNumbers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PendingSendPacketSequenceNumbers[iNdEx])
			copy(dAtA[i:], m.PendingSendPacketSequenceNumbers[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.PendingSendPacketSequenceNumbers[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.BlacklistedDenoms) > 0 {
		for iNdEx := len(m.BlacklistedDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BlacklistedDenoms[iNdEx])
			copy(dAtA[i:], m.BlacklistedDenoms[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.BlacklistedDenoms[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.WhitelistedAddressPairs) > 0 {
		for iNdEx := len(m.WhitelistedAddressPairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WhitelistedAddressPairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RateLimits) > 0 {
		for iNdEx := len(m.RateLimits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RateLimits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.RateLimits) > 0 {
		for _, e := range m.RateLimits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.WhitelistedAddressPairs) > 0 {
		for _, e := range m.WhitelistedAddressPairs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BlacklistedDenoms) > 0 {
		for _, s := range m.BlacklistedDenoms {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PendingSendPacketSequenceNumbers) > 0 {
		for _, s := range m.PendingSendPacketSequenceNumbers {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.HourEpoch.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RateLimits = append(m.RateLimits, RateLimit{})
			if err := m.RateLimits[len(m.RateLimits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhitelistedAddressPairs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhitelistedAddressPairs = append(m.WhitelistedAddressPairs, WhitelistedAddressPair{})
			if err := m.WhitelistedAddressPairs[len(m.WhitelistedAddressPairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlacklistedDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlacklistedDenoms = append(m.BlacklistedDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingSendPacketSequenceNumbers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PendingSendPacketSequenceNumbers = append(m.PendingSendPacketSequenceNumbers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HourEpoch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HourEpoch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
