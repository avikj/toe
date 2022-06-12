// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: toe/next_game_id.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type NextGameId struct {
	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *NextGameId) Reset()         { *m = NextGameId{} }
func (m *NextGameId) String() string { return proto.CompactTextString(m) }
func (*NextGameId) ProtoMessage()    {}
func (*NextGameId) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3cc993593a7857, []int{0}
}
func (m *NextGameId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NextGameId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NextGameId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NextGameId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NextGameId.Merge(m, src)
}
func (m *NextGameId) XXX_Size() int {
	return m.Size()
}
func (m *NextGameId) XXX_DiscardUnknown() {
	xxx_messageInfo_NextGameId.DiscardUnknown(m)
}

var xxx_messageInfo_NextGameId proto.InternalMessageInfo

func (m *NextGameId) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*NextGameId)(nil), "avikj.toe.toe.NextGameId")
}

func init() { proto.RegisterFile("toe/next_game_id.proto", fileDescriptor_dd3cc993593a7857) }

var fileDescriptor_dd3cc993593a7857 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0xc9, 0x4f, 0xd5,
	0xcf, 0x4b, 0xad, 0x28, 0x89, 0x4f, 0x4f, 0xcc, 0x4d, 0x8d, 0xcf, 0x4c, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x4d, 0x2c, 0xcb, 0xcc, 0xce, 0xd2, 0x2b, 0xc9, 0x4f, 0x05, 0x61, 0x25,
	0x25, 0x2e, 0x2e, 0xbf, 0xd4, 0x8a, 0x12, 0xf7, 0xc4, 0xdc, 0x54, 0xcf, 0x14, 0x21, 0x11, 0x2e,
	0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x08, 0xc7, 0xc9,
	0xea, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58,
	0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x14, 0xd2, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xc1, 0xe6, 0xea, 0x83, 0x6c, 0xad, 0x00, 0x93, 0x25,
	0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x5b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7e,
	0x7d, 0x7f, 0x30, 0x8f, 0x00, 0x00, 0x00,
}

func (m *NextGameId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NextGameId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NextGameId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintNextGameId(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNextGameId(dAtA []byte, offset int, v uint64) int {
	offset -= sovNextGameId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NextGameId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovNextGameId(uint64(m.Value))
	}
	return n
}

func sovNextGameId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNextGameId(x uint64) (n int) {
	return sovNextGameId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NextGameId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNextGameId
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
			return fmt.Errorf("proto: NextGameId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NextGameId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNextGameId
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNextGameId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNextGameId
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
func skipNextGameId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNextGameId
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
					return 0, ErrIntOverflowNextGameId
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
					return 0, ErrIntOverflowNextGameId
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
				return 0, ErrInvalidLengthNextGameId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNextGameId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNextGameId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNextGameId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNextGameId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNextGameId = fmt.Errorf("proto: unexpected end of group")
)