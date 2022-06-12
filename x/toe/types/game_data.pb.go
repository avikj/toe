// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: toe/game_data.proto

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

type GameData struct {
	Index      string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	PlayerX    string `protobuf:"bytes,2,opt,name=playerX,proto3" json:"playerX,omitempty"`
	PlayerO    string `protobuf:"bytes,3,opt,name=playerO,proto3" json:"playerO,omitempty"`
	BoardState string `protobuf:"bytes,4,opt,name=boardState,proto3" json:"boardState,omitempty"`
	GameId     uint64 `protobuf:"varint,5,opt,name=gameId,proto3" json:"gameId,omitempty"`
}

func (m *GameData) Reset()         { *m = GameData{} }
func (m *GameData) String() string { return proto.CompactTextString(m) }
func (*GameData) ProtoMessage()    {}
func (*GameData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f9980694d09f96, []int{0}
}
func (m *GameData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GameData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GameData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GameData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameData.Merge(m, src)
}
func (m *GameData) XXX_Size() int {
	return m.Size()
}
func (m *GameData) XXX_DiscardUnknown() {
	xxx_messageInfo_GameData.DiscardUnknown(m)
}

var xxx_messageInfo_GameData proto.InternalMessageInfo

func (m *GameData) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *GameData) GetPlayerX() string {
	if m != nil {
		return m.PlayerX
	}
	return ""
}

func (m *GameData) GetPlayerO() string {
	if m != nil {
		return m.PlayerO
	}
	return ""
}

func (m *GameData) GetBoardState() string {
	if m != nil {
		return m.BoardState
	}
	return ""
}

func (m *GameData) GetGameId() uint64 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func init() {
	proto.RegisterType((*GameData)(nil), "avikj.toe.toe.GameData")
}

func init() { proto.RegisterFile("toe/game_data.proto", fileDescriptor_f3f9980694d09f96) }

var fileDescriptor_f3f9980694d09f96 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xc9, 0x4f, 0xd5,
	0x4f, 0x4f, 0xcc, 0x4d, 0x8d, 0x4f, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4d, 0x2c, 0xcb, 0xcc, 0xce, 0xd2, 0x2b, 0xc9, 0x4f, 0x05, 0x61, 0xa5, 0x1e, 0x46, 0x2e,
	0x0e, 0xf7, 0xc4, 0xdc, 0x54, 0x97, 0xc4, 0x92, 0x44, 0x21, 0x11, 0x2e, 0xd6, 0xcc, 0xbc, 0x94,
	0xd4, 0x0a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x82, 0x8b, 0xbd, 0x20,
	0x27, 0xb1, 0x32, 0xb5, 0x28, 0x42, 0x82, 0x09, 0x2c, 0x0e, 0xe3, 0x22, 0x64, 0xfc, 0x25, 0x98,
	0x91, 0x65, 0xfc, 0x85, 0xe4, 0xb8, 0xb8, 0x92, 0xf2, 0x13, 0x8b, 0x52, 0x82, 0x4b, 0x12, 0x4b,
	0x52, 0x25, 0x58, 0xc0, 0x92, 0x48, 0x22, 0x42, 0x62, 0x5c, 0x6c, 0x20, 0x87, 0x79, 0xa6, 0x48,
	0xb0, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0x41, 0x79, 0x4e, 0x56, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78,
	0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc,
	0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x90, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab,
	0x0f, 0xf6, 0x82, 0x3e, 0xc8, 0x77, 0x15, 0x60, 0xb2, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d,
	0xec, 0x41, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x9f, 0x55, 0x53, 0xf7, 0x00, 0x00,
	0x00,
}

func (m *GameData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GameData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GameData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GameId != 0 {
		i = encodeVarintGameData(dAtA, i, uint64(m.GameId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.BoardState) > 0 {
		i -= len(m.BoardState)
		copy(dAtA[i:], m.BoardState)
		i = encodeVarintGameData(dAtA, i, uint64(len(m.BoardState)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PlayerO) > 0 {
		i -= len(m.PlayerO)
		copy(dAtA[i:], m.PlayerO)
		i = encodeVarintGameData(dAtA, i, uint64(len(m.PlayerO)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PlayerX) > 0 {
		i -= len(m.PlayerX)
		copy(dAtA[i:], m.PlayerX)
		i = encodeVarintGameData(dAtA, i, uint64(len(m.PlayerX)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintGameData(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGameData(dAtA []byte, offset int, v uint64) int {
	offset -= sovGameData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GameData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovGameData(uint64(l))
	}
	l = len(m.PlayerX)
	if l > 0 {
		n += 1 + l + sovGameData(uint64(l))
	}
	l = len(m.PlayerO)
	if l > 0 {
		n += 1 + l + sovGameData(uint64(l))
	}
	l = len(m.BoardState)
	if l > 0 {
		n += 1 + l + sovGameData(uint64(l))
	}
	if m.GameId != 0 {
		n += 1 + sovGameData(uint64(m.GameId))
	}
	return n
}

func sovGameData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGameData(x uint64) (n int) {
	return sovGameData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GameData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGameData
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
			return fmt.Errorf("proto: GameData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GameData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGameData
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
				return ErrInvalidLengthGameData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGameData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerX", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGameData
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
				return ErrInvalidLengthGameData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGameData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerX = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerO", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGameData
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
				return ErrInvalidLengthGameData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGameData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerO = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoardState", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGameData
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
				return ErrInvalidLengthGameData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGameData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BoardState = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			m.GameId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGameData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGameData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGameData
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
func skipGameData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGameData
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
					return 0, ErrIntOverflowGameData
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
					return 0, ErrIntOverflowGameData
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
				return 0, ErrInvalidLengthGameData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGameData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGameData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGameData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGameData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGameData = fmt.Errorf("proto: unexpected end of group")
)
