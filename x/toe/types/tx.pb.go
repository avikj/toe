// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: toe/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgNewGame struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgNewGame) Reset()         { *m = MsgNewGame{} }
func (m *MsgNewGame) String() string { return proto.CompactTextString(m) }
func (*MsgNewGame) ProtoMessage()    {}
func (*MsgNewGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_9899a8de15c32647, []int{0}
}
func (m *MsgNewGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNewGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNewGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNewGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNewGame.Merge(m, src)
}
func (m *MsgNewGame) XXX_Size() int {
	return m.Size()
}
func (m *MsgNewGame) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNewGame.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNewGame proto.InternalMessageInfo

func (m *MsgNewGame) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgNewGameResponse struct {
	GameId uint64 `protobuf:"varint,1,opt,name=gameId,proto3" json:"gameId,omitempty"`
}

func (m *MsgNewGameResponse) Reset()         { *m = MsgNewGameResponse{} }
func (m *MsgNewGameResponse) String() string { return proto.CompactTextString(m) }
func (*MsgNewGameResponse) ProtoMessage()    {}
func (*MsgNewGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9899a8de15c32647, []int{1}
}
func (m *MsgNewGameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNewGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNewGameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNewGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNewGameResponse.Merge(m, src)
}
func (m *MsgNewGameResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgNewGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNewGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNewGameResponse proto.InternalMessageInfo

func (m *MsgNewGameResponse) GetGameId() uint64 {
	if m != nil {
		return m.GameId
	}
	return 0
}

type MsgJoinGame struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	GameId  uint64 `protobuf:"varint,2,opt,name=gameId,proto3" json:"gameId,omitempty"`
}

func (m *MsgJoinGame) Reset()         { *m = MsgJoinGame{} }
func (m *MsgJoinGame) String() string { return proto.CompactTextString(m) }
func (*MsgJoinGame) ProtoMessage()    {}
func (*MsgJoinGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_9899a8de15c32647, []int{2}
}
func (m *MsgJoinGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgJoinGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgJoinGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgJoinGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgJoinGame.Merge(m, src)
}
func (m *MsgJoinGame) XXX_Size() int {
	return m.Size()
}
func (m *MsgJoinGame) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgJoinGame.DiscardUnknown(m)
}

var xxx_messageInfo_MsgJoinGame proto.InternalMessageInfo

func (m *MsgJoinGame) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgJoinGame) GetGameId() uint64 {
	if m != nil {
		return m.GameId
	}
	return 0
}

type MsgJoinGameResponse struct {
	PlayerX string `protobuf:"bytes,1,opt,name=playerX,proto3" json:"playerX,omitempty"`
	PlayerO string `protobuf:"bytes,2,opt,name=playerO,proto3" json:"playerO,omitempty"`
}

func (m *MsgJoinGameResponse) Reset()         { *m = MsgJoinGameResponse{} }
func (m *MsgJoinGameResponse) String() string { return proto.CompactTextString(m) }
func (*MsgJoinGameResponse) ProtoMessage()    {}
func (*MsgJoinGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9899a8de15c32647, []int{3}
}
func (m *MsgJoinGameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgJoinGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgJoinGameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgJoinGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgJoinGameResponse.Merge(m, src)
}
func (m *MsgJoinGameResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgJoinGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgJoinGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgJoinGameResponse proto.InternalMessageInfo

func (m *MsgJoinGameResponse) GetPlayerX() string {
	if m != nil {
		return m.PlayerX
	}
	return ""
}

func (m *MsgJoinGameResponse) GetPlayerO() string {
	if m != nil {
		return m.PlayerO
	}
	return ""
}

type MsgPlaceMarker struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	GameId  uint64 `protobuf:"varint,2,opt,name=gameId,proto3" json:"gameId,omitempty"`
	Pos     uint64 `protobuf:"varint,3,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (m *MsgPlaceMarker) Reset()         { *m = MsgPlaceMarker{} }
func (m *MsgPlaceMarker) String() string { return proto.CompactTextString(m) }
func (*MsgPlaceMarker) ProtoMessage()    {}
func (*MsgPlaceMarker) Descriptor() ([]byte, []int) {
	return fileDescriptor_9899a8de15c32647, []int{4}
}
func (m *MsgPlaceMarker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlaceMarker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlaceMarker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlaceMarker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlaceMarker.Merge(m, src)
}
func (m *MsgPlaceMarker) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlaceMarker) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlaceMarker.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlaceMarker proto.InternalMessageInfo

func (m *MsgPlaceMarker) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgPlaceMarker) GetGameId() uint64 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *MsgPlaceMarker) GetPos() uint64 {
	if m != nil {
		return m.Pos
	}
	return 0
}

type MsgPlaceMarkerResponse struct {
	GameStatus string `protobuf:"bytes,1,opt,name=gameStatus,proto3" json:"gameStatus,omitempty"`
}

func (m *MsgPlaceMarkerResponse) Reset()         { *m = MsgPlaceMarkerResponse{} }
func (m *MsgPlaceMarkerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPlaceMarkerResponse) ProtoMessage()    {}
func (*MsgPlaceMarkerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9899a8de15c32647, []int{5}
}
func (m *MsgPlaceMarkerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlaceMarkerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlaceMarkerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlaceMarkerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlaceMarkerResponse.Merge(m, src)
}
func (m *MsgPlaceMarkerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlaceMarkerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlaceMarkerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlaceMarkerResponse proto.InternalMessageInfo

func (m *MsgPlaceMarkerResponse) GetGameStatus() string {
	if m != nil {
		return m.GameStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgNewGame)(nil), "avikj.toe.toe.MsgNewGame")
	proto.RegisterType((*MsgNewGameResponse)(nil), "avikj.toe.toe.MsgNewGameResponse")
	proto.RegisterType((*MsgJoinGame)(nil), "avikj.toe.toe.MsgJoinGame")
	proto.RegisterType((*MsgJoinGameResponse)(nil), "avikj.toe.toe.MsgJoinGameResponse")
	proto.RegisterType((*MsgPlaceMarker)(nil), "avikj.toe.toe.MsgPlaceMarker")
	proto.RegisterType((*MsgPlaceMarkerResponse)(nil), "avikj.toe.toe.MsgPlaceMarkerResponse")
}

func init() { proto.RegisterFile("toe/tx.proto", fileDescriptor_9899a8de15c32647) }

var fileDescriptor_9899a8de15c32647 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc9, 0x4f, 0xd5,
	0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4d, 0x2c, 0xcb, 0xcc, 0xce, 0xd2,
	0x2b, 0xc9, 0x4f, 0x05, 0x61, 0x25, 0x35, 0x2e, 0x2e, 0xdf, 0xe2, 0x74, 0xbf, 0xd4, 0x72, 0xf7,
	0xc4, 0xdc, 0x54, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xa2, 0xd4, 0xc4, 0x92, 0xfc, 0x22, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x49, 0x87, 0x4b, 0x08, 0xa1, 0x2e, 0x28, 0xb5, 0xb8,
	0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8c, 0x8b, 0x2d, 0x3d, 0x31, 0x37, 0xd5, 0x33, 0x05, 0xac,
	0x9c, 0x25, 0x08, 0xca, 0x53, 0xb2, 0xe7, 0xe2, 0xf6, 0x2d, 0x4e, 0xf7, 0xca, 0xcf, 0xcc, 0xc3,
	0x6f, 0x2c, 0x92, 0x01, 0x4c, 0x28, 0x06, 0x78, 0x72, 0x09, 0x23, 0x19, 0x00, 0xb7, 0x4f, 0x82,
	0x8b, 0xbd, 0x20, 0x27, 0xb1, 0x32, 0xb5, 0x28, 0x02, 0x66, 0x10, 0x94, 0x8b, 0x90, 0xf1, 0x07,
	0x9b, 0x04, 0x97, 0xf1, 0x57, 0x0a, 0xe1, 0xe2, 0xf3, 0x2d, 0x4e, 0x0f, 0xc8, 0x49, 0x4c, 0x4e,
	0xf5, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0x22, 0xdd, 0x39, 0x42, 0x02, 0x5c, 0xcc, 0x05, 0xf9, 0xc5,
	0x12, 0xcc, 0x60, 0x41, 0x10, 0x53, 0xc9, 0x82, 0x4b, 0x0c, 0xd5, 0x54, 0xb8, 0x1b, 0xe5, 0xb8,
	0xb8, 0x40, 0xba, 0x82, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0xa1, 0x16, 0x20, 0x89, 0x18, 0xbd, 0x67,
	0xe4, 0x62, 0xf6, 0x2d, 0x4e, 0x17, 0x72, 0xe7, 0x62, 0x87, 0x05, 0xbb, 0xa4, 0x1e, 0x4a, 0xa4,
	0xe8, 0x21, 0x42, 0x5a, 0x4a, 0x11, 0xa7, 0x14, 0xdc, 0x42, 0x2f, 0x2e, 0x0e, 0x78, 0x48, 0x4b,
	0x61, 0x2a, 0x87, 0xc9, 0x49, 0x29, 0xe1, 0x96, 0x83, 0x9b, 0x15, 0xcc, 0xc5, 0x8d, 0x1c, 0x52,
	0xb2, 0x98, 0x5a, 0x90, 0xa4, 0xa5, 0x54, 0xf1, 0x4a, 0xc3, 0x0c, 0x75, 0xb2, 0x3a, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0,
	0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x85, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd,
	0xe4, 0xfc, 0x5c, 0x7d, 0xb0, 0x51, 0xfa, 0xa0, 0xb4, 0x5a, 0x01, 0x26, 0x4b, 0x2a, 0x0b, 0x52,
	0x8b, 0x93, 0xd8, 0xc0, 0xa9, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x90, 0x3b, 0x58, 0x3a,
	0xc5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	NewGame(ctx context.Context, in *MsgNewGame, opts ...grpc.CallOption) (*MsgNewGameResponse, error)
	JoinGame(ctx context.Context, in *MsgJoinGame, opts ...grpc.CallOption) (*MsgJoinGameResponse, error)
	PlaceMarker(ctx context.Context, in *MsgPlaceMarker, opts ...grpc.CallOption) (*MsgPlaceMarkerResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) NewGame(ctx context.Context, in *MsgNewGame, opts ...grpc.CallOption) (*MsgNewGameResponse, error) {
	out := new(MsgNewGameResponse)
	err := c.cc.Invoke(ctx, "/avikj.toe.toe.Msg/NewGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) JoinGame(ctx context.Context, in *MsgJoinGame, opts ...grpc.CallOption) (*MsgJoinGameResponse, error) {
	out := new(MsgJoinGameResponse)
	err := c.cc.Invoke(ctx, "/avikj.toe.toe.Msg/JoinGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) PlaceMarker(ctx context.Context, in *MsgPlaceMarker, opts ...grpc.CallOption) (*MsgPlaceMarkerResponse, error) {
	out := new(MsgPlaceMarkerResponse)
	err := c.cc.Invoke(ctx, "/avikj.toe.toe.Msg/PlaceMarker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	NewGame(context.Context, *MsgNewGame) (*MsgNewGameResponse, error)
	JoinGame(context.Context, *MsgJoinGame) (*MsgJoinGameResponse, error)
	PlaceMarker(context.Context, *MsgPlaceMarker) (*MsgPlaceMarkerResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) NewGame(ctx context.Context, req *MsgNewGame) (*MsgNewGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewGame not implemented")
}
func (*UnimplementedMsgServer) JoinGame(ctx context.Context, req *MsgJoinGame) (*MsgJoinGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGame not implemented")
}
func (*UnimplementedMsgServer) PlaceMarker(ctx context.Context, req *MsgPlaceMarker) (*MsgPlaceMarkerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceMarker not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_NewGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewGame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/avikj.toe.toe.Msg/NewGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewGame(ctx, req.(*MsgNewGame))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_JoinGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgJoinGame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).JoinGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/avikj.toe.toe.Msg/JoinGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).JoinGame(ctx, req.(*MsgJoinGame))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_PlaceMarker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPlaceMarker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PlaceMarker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/avikj.toe.toe.Msg/PlaceMarker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PlaceMarker(ctx, req.(*MsgPlaceMarker))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "avikj.toe.toe.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewGame",
			Handler:    _Msg_NewGame_Handler,
		},
		{
			MethodName: "JoinGame",
			Handler:    _Msg_JoinGame_Handler,
		},
		{
			MethodName: "PlaceMarker",
			Handler:    _Msg_PlaceMarker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "toe/tx.proto",
}

func (m *MsgNewGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNewGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNewGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgNewGameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNewGameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNewGameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GameId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GameId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgJoinGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgJoinGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgJoinGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GameId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GameId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgJoinGameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgJoinGameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgJoinGameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PlayerO) > 0 {
		i -= len(m.PlayerO)
		copy(dAtA[i:], m.PlayerO)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PlayerO)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PlayerX) > 0 {
		i -= len(m.PlayerX)
		copy(dAtA[i:], m.PlayerX)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PlayerX)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPlaceMarker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlaceMarker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlaceMarker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pos != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Pos))
		i--
		dAtA[i] = 0x18
	}
	if m.GameId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GameId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPlaceMarkerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlaceMarkerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlaceMarkerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GameStatus) > 0 {
		i -= len(m.GameStatus)
		copy(dAtA[i:], m.GameStatus)
		i = encodeVarintTx(dAtA, i, uint64(len(m.GameStatus)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgNewGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgNewGameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GameId != 0 {
		n += 1 + sovTx(uint64(m.GameId))
	}
	return n
}

func (m *MsgJoinGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GameId != 0 {
		n += 1 + sovTx(uint64(m.GameId))
	}
	return n
}

func (m *MsgJoinGameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PlayerX)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PlayerO)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgPlaceMarker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GameId != 0 {
		n += 1 + sovTx(uint64(m.GameId))
	}
	if m.Pos != 0 {
		n += 1 + sovTx(uint64(m.Pos))
	}
	return n
}

func (m *MsgPlaceMarkerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GameStatus)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgNewGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgNewGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNewGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgNewGameResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgNewGameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNewGameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			m.GameId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgJoinGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgJoinGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgJoinGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			m.GameId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgJoinGameResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgJoinGameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgJoinGameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerX", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerX = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerO", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerO = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgPlaceMarker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPlaceMarker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlaceMarker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			m.GameId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pos", wireType)
			}
			m.Pos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pos |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgPlaceMarkerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPlaceMarkerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlaceMarkerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GameStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
