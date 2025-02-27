// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/channel/v1/upgrade.proto

package v10

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

// Upgrade is a verifiable type which contains the relevant information
// for an attempted upgrade. It provides the proposed changes to the channel
// end, the timeout for this upgrade attempt and the next packet sequence
// which allows the counterparty to efficiently know the highest sequence it has received.
// The next sequence send is used for pruning and upgrading from unordered to ordered channels.
type Upgrade struct {
	Fields           UpgradeFields `protobuf:"bytes,1,opt,name=fields,proto3" json:"fields"`
	Timeout          Timeout       `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout"`
	NextSequenceSend uint64        `protobuf:"varint,3,opt,name=next_sequence_send,json=nextSequenceSend,proto3" json:"next_sequence_send,omitempty"`
}

func (m *Upgrade) Reset()         { *m = Upgrade{} }
func (m *Upgrade) String() string { return proto.CompactTextString(m) }
func (*Upgrade) ProtoMessage()    {}
func (*Upgrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb1cef68588848b2, []int{0}
}
func (m *Upgrade) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Upgrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Upgrade.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Upgrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Upgrade.Merge(m, src)
}
func (m *Upgrade) XXX_Size() int {
	return m.Size()
}
func (m *Upgrade) XXX_DiscardUnknown() {
	xxx_messageInfo_Upgrade.DiscardUnknown(m)
}

var xxx_messageInfo_Upgrade proto.InternalMessageInfo

// UpgradeFields are the fields in a channel end which may be changed
// during a channel upgrade.
type UpgradeFields struct {
	Ordering       Order    `protobuf:"varint,1,opt,name=ordering,proto3,enum=ibc.core.channel.v1.Order" json:"ordering,omitempty"`
	ConnectionHops []string `protobuf:"bytes,2,rep,name=connection_hops,json=connectionHops,proto3" json:"connection_hops,omitempty"`
	Version        string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *UpgradeFields) Reset()         { *m = UpgradeFields{} }
func (m *UpgradeFields) String() string { return proto.CompactTextString(m) }
func (*UpgradeFields) ProtoMessage()    {}
func (*UpgradeFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb1cef68588848b2, []int{1}
}
func (m *UpgradeFields) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpgradeFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpgradeFields.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpgradeFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeFields.Merge(m, src)
}
func (m *UpgradeFields) XXX_Size() int {
	return m.Size()
}
func (m *UpgradeFields) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeFields.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeFields proto.InternalMessageInfo

// ErrorReceipt defines a type which encapsulates the upgrade sequence and error associated with the
// upgrade handshake failure. When a channel upgrade handshake is aborted both chains are expected to increment to the
// next sequence.
type ErrorReceipt struct {
	// the channel upgrade sequence
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// the error message detailing the cause of failure
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *ErrorReceipt) Reset()         { *m = ErrorReceipt{} }
func (m *ErrorReceipt) String() string { return proto.CompactTextString(m) }
func (*ErrorReceipt) ProtoMessage()    {}
func (*ErrorReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb1cef68588848b2, []int{2}
}
func (m *ErrorReceipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrorReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrorReceipt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrorReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorReceipt.Merge(m, src)
}
func (m *ErrorReceipt) XXX_Size() int {
	return m.Size()
}
func (m *ErrorReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorReceipt proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Upgrade)(nil), "ibc.core.channel.v1.Upgrade")
	proto.RegisterType((*UpgradeFields)(nil), "ibc.core.channel.v1.UpgradeFields")
	proto.RegisterType((*ErrorReceipt)(nil), "ibc.core.channel.v1.ErrorReceipt")
}

func init() { proto.RegisterFile("ibc/core/channel/v1/upgrade.proto", fileDescriptor_fb1cef68588848b2) }

var fileDescriptor_fb1cef68588848b2 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe3, 0xad, 0x5a, 0x57, 0x03, 0x03, 0x19, 0x0e, 0x51, 0x85, 0xb2, 0xd2, 0x0b, 0x3d,
	0xb0, 0x78, 0x1d, 0x88, 0xc3, 0xc4, 0x01, 0x4d, 0x02, 0x21, 0x2e, 0x48, 0xde, 0xb8, 0x70, 0xa9,
	0x1a, 0xe7, 0x91, 0x5a, 0x6a, 0xfc, 0x82, 0xed, 0x44, 0xf0, 0x0d, 0x38, 0xee, 0x23, 0xf0, 0x45,
	0xb8, 0xef, 0xb8, 0x23, 0x27, 0x84, 0xda, 0x2f, 0x82, 0xe2, 0x24, 0x45, 0x48, 0xb9, 0xf9, 0xf9,
	0xfd, 0xfe, 0x7f, 0xff, 0x9f, 0x1f, 0x7d, 0xa2, 0x12, 0xc9, 0x25, 0x1a, 0xe0, 0x72, 0xb5, 0xd4,
	0x1a, 0xd6, 0xbc, 0x9a, 0xf3, 0xb2, 0xc8, 0xcc, 0x32, 0x85, 0xb8, 0x30, 0xe8, 0x90, 0x3d, 0x54,
	0x89, 0x8c, 0x6b, 0x24, 0x6e, 0x91, 0xb8, 0x9a, 0x8f, 0x1f, 0x65, 0x98, 0xa1, 0xef, 0xf3, 0xfa,
	0xd4, 0xa0, 0xe3, 0x5e, 0xb7, 0x4e, 0xe5, 0x91, 0xe9, 0x4f, 0x42, 0x87, 0x1f, 0x1b, 0x7f, 0xf6,
	0x9a, 0x1e, 0x7c, 0x56, 0xb0, 0x4e, 0x6d, 0x48, 0x26, 0x64, 0x76, 0xe7, 0x6c, 0x1a, 0xf7, 0x3c,
	0x15, 0xb7, 0xf4, 0x5b, 0x4f, 0x5e, 0x0c, 0x6e, 0x7e, 0x1f, 0x07, 0xa2, 0xd5, 0xb1, 0x57, 0x74,
	0xe8, 0x54, 0x0e, 0x58, 0xba, 0x70, 0xcf, 0x5b, 0x3c, 0xee, 0xb5, 0xb8, 0x6a, 0x98, 0x56, 0xdc,
	0x49, 0xd8, 0x33, 0xca, 0x34, 0x7c, 0x75, 0x0b, 0x0b, 0x5f, 0x4a, 0xd0, 0x12, 0x16, 0x16, 0x74,
	0x1a, 0xee, 0x4f, 0xc8, 0x6c, 0x20, 0x1e, 0xd4, 0x9d, 0xcb, 0xb6, 0x71, 0x09, 0x3a, 0x3d, 0x1f,
	0x7c, 0xff, 0x71, 0x1c, 0x4c, 0xaf, 0x09, 0xbd, 0xf7, 0x5f, 0x22, 0xf6, 0x92, 0x1e, 0xa2, 0x49,
	0xc1, 0x28, 0x9d, 0xf9, 0x39, 0x8e, 0xce, 0xc6, 0xbd, 0x21, 0x3e, 0xd4, 0x90, 0xd8, 0xb1, 0xec,
	0x29, 0xbd, 0x2f, 0x51, 0x6b, 0x90, 0x4e, 0xa1, 0x5e, 0xac, 0xb0, 0xb0, 0xe1, 0xde, 0x64, 0x7f,
	0x36, 0x12, 0x47, 0xff, 0xae, 0xdf, 0x61, 0x61, 0x59, 0x48, 0x87, 0x15, 0x18, 0xab, 0x50, 0xfb,
	0x6c, 0x23, 0xd1, 0x95, 0x6d, 0xa4, 0xf7, 0xf4, 0xee, 0x1b, 0x63, 0xd0, 0x08, 0x90, 0xa0, 0x0a,
	0xc7, 0xc6, 0xf4, 0xb0, 0x9b, 0xc8, 0x07, 0x1a, 0x88, 0x5d, 0x5d, 0x7b, 0xe5, 0x60, 0xed, 0x32,
	0x03, 0xff, 0x61, 0x23, 0xd1, 0x95, 0x8d, 0xd7, 0xc5, 0xd5, 0xcd, 0x26, 0x22, 0xb7, 0x9b, 0x88,
	0xfc, 0xd9, 0x44, 0xe4, 0x7a, 0x1b, 0x05, 0xb7, 0xdb, 0x28, 0xf8, 0xb5, 0x8d, 0x82, 0x4f, 0xe7,
	0x99, 0x72, 0xab, 0x32, 0x89, 0x25, 0xe6, 0x5c, 0xa2, 0xcd, 0xd1, 0x72, 0x95, 0xc8, 0x93, 0x0c,
	0x79, 0x35, 0x3f, 0xe5, 0x39, 0xa6, 0xe5, 0x1a, 0x6c, 0xb3, 0xfc, 0xd3, 0x17, 0x27, 0xdd, 0xfe,
	0xdd, 0xb7, 0x02, 0x6c, 0x72, 0xe0, 0x77, 0xff, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90,
	0x11, 0x85, 0x24, 0x6e, 0x02, 0x00, 0x00,
}

func (m *Upgrade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Upgrade) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Upgrade) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NextSequenceSend != 0 {
		i = encodeVarintUpgrade(dAtA, i, uint64(m.NextSequenceSend))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.Timeout.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintUpgrade(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Fields.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintUpgrade(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *UpgradeFields) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpgradeFields) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpgradeFields) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionHops) > 0 {
		for iNdEx := len(m.ConnectionHops) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ConnectionHops[iNdEx])
			copy(dAtA[i:], m.ConnectionHops[iNdEx])
			i = encodeVarintUpgrade(dAtA, i, uint64(len(m.ConnectionHops[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Ordering != 0 {
		i = encodeVarintUpgrade(dAtA, i, uint64(m.Ordering))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ErrorReceipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorReceipt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrorReceipt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.Sequence != 0 {
		i = encodeVarintUpgrade(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintUpgrade(dAtA []byte, offset int, v uint64) int {
	offset -= sovUpgrade(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Upgrade) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Fields.Size()
	n += 1 + l + sovUpgrade(uint64(l))
	l = m.Timeout.Size()
	n += 1 + l + sovUpgrade(uint64(l))
	if m.NextSequenceSend != 0 {
		n += 1 + sovUpgrade(uint64(m.NextSequenceSend))
	}
	return n
}

func (m *UpgradeFields) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ordering != 0 {
		n += 1 + sovUpgrade(uint64(m.Ordering))
	}
	if len(m.ConnectionHops) > 0 {
		for _, s := range m.ConnectionHops {
			l = len(s)
			n += 1 + l + sovUpgrade(uint64(l))
		}
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	return n
}

func (m *ErrorReceipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovUpgrade(uint64(m.Sequence))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	return n
}

func sovUpgrade(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUpgrade(x uint64) (n int) {
	return sovUpgrade(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Upgrade) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpgrade
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
			return fmt.Errorf("proto: Upgrade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Upgrade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fields", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fields.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Timeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextSequenceSend", wireType)
			}
			m.NextSequenceSend = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextSequenceSend |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpgrade
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
func (m *UpgradeFields) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpgrade
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
			return fmt.Errorf("proto: UpgradeFields: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpgradeFields: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ordering", wireType)
			}
			m.Ordering = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ordering |= Order(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionHops", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionHops = append(m.ConnectionHops, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpgrade
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
func (m *ErrorReceipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpgrade
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
			return fmt.Errorf("proto: ErrorReceipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorReceipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpgrade
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
func skipUpgrade(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUpgrade
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
					return 0, ErrIntOverflowUpgrade
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
					return 0, ErrIntOverflowUpgrade
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
				return 0, ErrInvalidLengthUpgrade
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUpgrade
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUpgrade
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUpgrade        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUpgrade          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUpgrade = fmt.Errorf("proto: unexpected end of group")
)
