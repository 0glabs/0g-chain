// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zgc/dasigners/v1/dasigners.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

type Signer struct {
	// account defines the hex address of signer without 0x
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// socket defines the da node socket address
	Socket string `protobuf:"bytes,2,opt,name=socket,proto3" json:"socket,omitempty"`
	// pubkey_g1 defines the public key on bn254 G1
	PubkeyG1 []byte `protobuf:"bytes,3,opt,name=pubkey_g1,json=pubkeyG1,proto3" json:"pubkey_g1,omitempty"`
	// pubkey_g1 defines the public key on bn254 G2
	PubkeyG2 []byte `protobuf:"bytes,4,opt,name=pubkey_g2,json=pubkeyG2,proto3" json:"pubkey_g2,omitempty"`
}

func (m *Signer) Reset()         { *m = Signer{} }
func (m *Signer) String() string { return proto.CompactTextString(m) }
func (*Signer) ProtoMessage()    {}
func (*Signer) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7328dc8ffac059e, []int{0}
}
func (m *Signer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Signer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Signer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Signer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signer.Merge(m, src)
}
func (m *Signer) XXX_Size() int {
	return m.Size()
}
func (m *Signer) XXX_DiscardUnknown() {
	xxx_messageInfo_Signer.DiscardUnknown(m)
}

var xxx_messageInfo_Signer proto.InternalMessageInfo

type EpochSignerSet struct {
	Signers []string `protobuf:"bytes,1,rep,name=signers,proto3" json:"signers,omitempty"`
}

func (m *EpochSignerSet) Reset()         { *m = EpochSignerSet{} }
func (m *EpochSignerSet) String() string { return proto.CompactTextString(m) }
func (*EpochSignerSet) ProtoMessage()    {}
func (*EpochSignerSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7328dc8ffac059e, []int{1}
}
func (m *EpochSignerSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochSignerSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochSignerSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochSignerSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochSignerSet.Merge(m, src)
}
func (m *EpochSignerSet) XXX_Size() int {
	return m.Size()
}
func (m *EpochSignerSet) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochSignerSet.DiscardUnknown(m)
}

var xxx_messageInfo_EpochSignerSet proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Signer)(nil), "zgc.dasigners.v1.Signer")
	proto.RegisterType((*EpochSignerSet)(nil), "zgc.dasigners.v1.EpochSignerSet")
}

func init() { proto.RegisterFile("zgc/dasigners/v1/dasigners.proto", fileDescriptor_b7328dc8ffac059e) }

var fileDescriptor_b7328dc8ffac059e = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x1b, 0x27, 0xd3, 0x05, 0x11, 0x29, 0x22, 0xd9, 0x84, 0x50, 0x76, 0x1a, 0x82, 0xcd,
	0x3a, 0xdf, 0x40, 0x10, 0x4f, 0x5e, 0xb6, 0x9b, 0x97, 0x91, 0x66, 0x31, 0x2d, 0xdb, 0xfa, 0x2f,
	0x4d, 0x3a, 0xec, 0x9e, 0xc2, 0xc7, 0xda, 0x71, 0x47, 0x8f, 0xda, 0xbe, 0x88, 0xb4, 0xa9, 0xcc,
	0x79, 0xcb, 0xef, 0xfb, 0x05, 0x3e, 0xfe, 0x1f, 0xf6, 0xb6, 0x4a, 0xb0, 0x05, 0xd7, 0xb1, 0x4a,
	0x64, 0xa6, 0xd9, 0x26, 0x38, 0x80, 0x9f, 0x66, 0x60, 0xc0, 0xbd, 0xda, 0x2a, 0xe1, 0x1f, 0xc2,
	0x4d, 0x30, 0xe8, 0x0b, 0xd0, 0x6b, 0xd0, 0xf3, 0xc6, 0x33, 0x0b, 0xf6, 0xf3, 0xe0, 0x5a, 0x81,
	0x02, 0x9b, 0xd7, 0xaf, 0x36, 0xed, 0x2b, 0x00, 0xb5, 0x92, 0xac, 0xa1, 0x30, 0x7f, 0x63, 0x3c,
	0x29, 0x5a, 0x45, 0xff, 0xab, 0x45, 0x9e, 0x71, 0x13, 0x43, 0x62, 0xfd, 0xd0, 0xe0, 0xee, 0xac,
	0x69, 0x76, 0x09, 0x3e, 0xe3, 0x42, 0x40, 0x9e, 0x18, 0x82, 0x3c, 0x34, 0xea, 0x4d, 0x7f, 0xd1,
	0xbd, 0xc1, 0x5d, 0x0d, 0x62, 0x29, 0x0d, 0x39, 0x69, 0x44, 0x4b, 0xee, 0x2d, 0xee, 0xa5, 0x79,
	0xb8, 0x94, 0xc5, 0x5c, 0x05, 0xa4, 0xe3, 0xa1, 0xd1, 0xc5, 0xf4, 0xdc, 0x06, 0xcf, 0xc1, 0x5f,
	0x39, 0x21, 0xa7, 0x47, 0x72, 0x32, 0xbc, 0xc3, 0x97, 0x4f, 0x29, 0x88, 0xc8, 0x56, 0xcf, 0xa4,
	0xa9, 0xdb, 0xdb, 0x05, 0x08, 0xf2, 0x3a, 0x75, 0x7b, 0x8b, 0x8f, 0x2f, 0xbb, 0x6f, 0xea, 0xec,
	0x4a, 0x8a, 0xf6, 0x25, 0x45, 0x5f, 0x25, 0x45, 0x1f, 0x15, 0x75, 0xf6, 0x15, 0x75, 0x3e, 0x2b,
	0xea, 0xbc, 0x32, 0x15, 0x9b, 0x28, 0x0f, 0x7d, 0x01, 0x6b, 0x36, 0x56, 0x2b, 0x1e, 0x6a, 0x36,
	0x56, 0xf7, 0x22, 0xe2, 0x71, 0xc2, 0xde, 0x8f, 0x87, 0x37, 0x45, 0x2a, 0x75, 0xd8, 0x6d, 0xee,
	0x7e, 0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0x77, 0x51, 0x09, 0xd9, 0x99, 0x01, 0x00, 0x00,
}

func (m *Signer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Signer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Signer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubkeyG2) > 0 {
		i -= len(m.PubkeyG2)
		copy(dAtA[i:], m.PubkeyG2)
		i = encodeVarintDasigners(dAtA, i, uint64(len(m.PubkeyG2)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PubkeyG1) > 0 {
		i -= len(m.PubkeyG1)
		copy(dAtA[i:], m.PubkeyG1)
		i = encodeVarintDasigners(dAtA, i, uint64(len(m.PubkeyG1)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Socket) > 0 {
		i -= len(m.Socket)
		copy(dAtA[i:], m.Socket)
		i = encodeVarintDasigners(dAtA, i, uint64(len(m.Socket)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintDasigners(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EpochSignerSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochSignerSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochSignerSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintDasigners(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDasigners(dAtA []byte, offset int, v uint64) int {
	offset -= sovDasigners(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Signer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovDasigners(uint64(l))
	}
	l = len(m.Socket)
	if l > 0 {
		n += 1 + l + sovDasigners(uint64(l))
	}
	l = len(m.PubkeyG1)
	if l > 0 {
		n += 1 + l + sovDasigners(uint64(l))
	}
	l = len(m.PubkeyG2)
	if l > 0 {
		n += 1 + l + sovDasigners(uint64(l))
	}
	return n
}

func (m *EpochSignerSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovDasigners(uint64(l))
		}
	}
	return n
}

func sovDasigners(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDasigners(x uint64) (n int) {
	return sovDasigners(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Signer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDasigners
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
			return fmt.Errorf("proto: Signer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Signer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDasigners
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
				return ErrInvalidLengthDasigners
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDasigners
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Socket", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDasigners
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
				return ErrInvalidLengthDasigners
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDasigners
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Socket = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubkeyG1", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDasigners
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDasigners
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDasigners
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubkeyG1 = append(m.PubkeyG1[:0], dAtA[iNdEx:postIndex]...)
			if m.PubkeyG1 == nil {
				m.PubkeyG1 = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubkeyG2", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDasigners
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDasigners
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDasigners
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubkeyG2 = append(m.PubkeyG2[:0], dAtA[iNdEx:postIndex]...)
			if m.PubkeyG2 == nil {
				m.PubkeyG2 = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDasigners(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDasigners
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
func (m *EpochSignerSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDasigners
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
			return fmt.Errorf("proto: EpochSignerSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochSignerSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDasigners
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
				return ErrInvalidLengthDasigners
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDasigners
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDasigners(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDasigners
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
func skipDasigners(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDasigners
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
					return 0, ErrIntOverflowDasigners
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
					return 0, ErrIntOverflowDasigners
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
				return 0, ErrInvalidLengthDasigners
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDasigners
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDasigners
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDasigners        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDasigners          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDasigners = fmt.Errorf("proto: unexpected end of group")
)
