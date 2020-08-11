// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConnectClient is the client API for Connect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectClient interface {
	ConnectToServer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ServerReply, error)
	Soal(ctx context.Context, in *PilihanSoal, opts ...grpc.CallOption) (*SoalList, error)
	Hasil(ctx context.Context, in *KumpulanJawaban, opts ...grpc.CallOption) (*Result, error)
	AmbilTema(ctx context.Context, in *PilihTema, opts ...grpc.CallOption) (*TemaList, error)
	AmbilMapel(ctx context.Context, in *PilihMapel, opts ...grpc.CallOption) (*MapelList, error)
	AmbilKelas(ctx context.Context, in *PilihKelas, opts ...grpc.CallOption) (*KelasList, error)
}

type connectClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectClient(cc grpc.ClientConnInterface) ConnectClient {
	return &connectClient{cc}
}

func (c *connectClient) ConnectToServer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ServerReply, error) {
	out := new(ServerReply)
	err := c.cc.Invoke(ctx, "/SiswaService.Connect/ConnectToServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) Soal(ctx context.Context, in *PilihanSoal, opts ...grpc.CallOption) (*SoalList, error) {
	out := new(SoalList)
	err := c.cc.Invoke(ctx, "/SiswaService.Connect/Soal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) Hasil(ctx context.Context, in *KumpulanJawaban, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/SiswaService.Connect/Hasil", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) AmbilTema(ctx context.Context, in *PilihTema, opts ...grpc.CallOption) (*TemaList, error) {
	out := new(TemaList)
	err := c.cc.Invoke(ctx, "/SiswaService.Connect/AmbilTema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) AmbilMapel(ctx context.Context, in *PilihMapel, opts ...grpc.CallOption) (*MapelList, error) {
	out := new(MapelList)
	err := c.cc.Invoke(ctx, "/SiswaService.Connect/AmbilMapel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) AmbilKelas(ctx context.Context, in *PilihKelas, opts ...grpc.CallOption) (*KelasList, error) {
	out := new(KelasList)
	err := c.cc.Invoke(ctx, "/SiswaService.Connect/AmbilKelas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectServer is the server API for Connect service.
// All implementations must embed UnimplementedConnectServer
// for forward compatibility
type ConnectServer interface {
	ConnectToServer(context.Context, *empty.Empty) (*ServerReply, error)
	Soal(context.Context, *PilihanSoal) (*SoalList, error)
	Hasil(context.Context, *KumpulanJawaban) (*Result, error)
	AmbilTema(context.Context, *PilihTema) (*TemaList, error)
	AmbilMapel(context.Context, *PilihMapel) (*MapelList, error)
	AmbilKelas(context.Context, *PilihKelas) (*KelasList, error)
	mustEmbedUnimplementedConnectServer()
}

// UnimplementedConnectServer must be embedded to have forward compatible implementations.
type UnimplementedConnectServer struct {
}

func (*UnimplementedConnectServer) ConnectToServer(context.Context, *empty.Empty) (*ServerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectToServer not implemented")
}
func (*UnimplementedConnectServer) Soal(context.Context, *PilihanSoal) (*SoalList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Soal not implemented")
}
func (*UnimplementedConnectServer) Hasil(context.Context, *KumpulanJawaban) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hasil not implemented")
}
func (*UnimplementedConnectServer) AmbilTema(context.Context, *PilihTema) (*TemaList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AmbilTema not implemented")
}
func (*UnimplementedConnectServer) AmbilMapel(context.Context, *PilihMapel) (*MapelList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AmbilMapel not implemented")
}
func (*UnimplementedConnectServer) AmbilKelas(context.Context, *PilihKelas) (*KelasList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AmbilKelas not implemented")
}
func (*UnimplementedConnectServer) mustEmbedUnimplementedConnectServer() {}

func RegisterConnectServer(s *grpc.Server, srv ConnectServer) {
	s.RegisterService(&_Connect_serviceDesc, srv)
}

func _Connect_ConnectToServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).ConnectToServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiswaService.Connect/ConnectToServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).ConnectToServer(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_Soal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PilihanSoal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).Soal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiswaService.Connect/Soal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).Soal(ctx, req.(*PilihanSoal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_Hasil_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KumpulanJawaban)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).Hasil(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiswaService.Connect/Hasil",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).Hasil(ctx, req.(*KumpulanJawaban))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_AmbilTema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PilihTema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).AmbilTema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiswaService.Connect/AmbilTema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).AmbilTema(ctx, req.(*PilihTema))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_AmbilMapel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PilihMapel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).AmbilMapel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiswaService.Connect/AmbilMapel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).AmbilMapel(ctx, req.(*PilihMapel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_AmbilKelas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PilihKelas)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).AmbilKelas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiswaService.Connect/AmbilKelas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).AmbilKelas(ctx, req.(*PilihKelas))
	}
	return interceptor(ctx, in, info, handler)
}

var _Connect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SiswaService.Connect",
	HandlerType: (*ConnectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConnectToServer",
			Handler:    _Connect_ConnectToServer_Handler,
		},
		{
			MethodName: "Soal",
			Handler:    _Connect_Soal_Handler,
		},
		{
			MethodName: "Hasil",
			Handler:    _Connect_Hasil_Handler,
		},
		{
			MethodName: "AmbilTema",
			Handler:    _Connect_AmbilTema_Handler,
		},
		{
			MethodName: "AmbilMapel",
			Handler:    _Connect_AmbilMapel_Handler,
		},
		{
			MethodName: "AmbilKelas",
			Handler:    _Connect_AmbilKelas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "siswaservice.proto",
}

// GuruClient is the client API for Guru service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GuruClient interface {
	KirimSoalKeServer(ctx context.Context, in *PaketSoal, opts ...grpc.CallOption) (*ServerGuruReply, error)
}

type guruClient struct {
	cc grpc.ClientConnInterface
}

func NewGuruClient(cc grpc.ClientConnInterface) GuruClient {
	return &guruClient{cc}
}

func (c *guruClient) KirimSoalKeServer(ctx context.Context, in *PaketSoal, opts ...grpc.CallOption) (*ServerGuruReply, error) {
	out := new(ServerGuruReply)
	err := c.cc.Invoke(ctx, "/SiswaService.Guru/KirimSoalKeServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuruServer is the server API for Guru service.
// All implementations must embed UnimplementedGuruServer
// for forward compatibility
type GuruServer interface {
	KirimSoalKeServer(context.Context, *PaketSoal) (*ServerGuruReply, error)
	mustEmbedUnimplementedGuruServer()
}

// UnimplementedGuruServer must be embedded to have forward compatible implementations.
type UnimplementedGuruServer struct {
}

func (*UnimplementedGuruServer) KirimSoalKeServer(context.Context, *PaketSoal) (*ServerGuruReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KirimSoalKeServer not implemented")
}
func (*UnimplementedGuruServer) mustEmbedUnimplementedGuruServer() {}

func RegisterGuruServer(s *grpc.Server, srv GuruServer) {
	s.RegisterService(&_Guru_serviceDesc, srv)
}

func _Guru_KirimSoalKeServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaketSoal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuruServer).KirimSoalKeServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiswaService.Guru/KirimSoalKeServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuruServer).KirimSoalKeServer(ctx, req.(*PaketSoal))
	}
	return interceptor(ctx, in, info, handler)
}

var _Guru_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SiswaService.Guru",
	HandlerType: (*GuruServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KirimSoalKeServer",
			Handler:    _Guru_KirimSoalKeServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "siswaservice.proto",
}

// DataClient is the client API for Data service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataClient interface {
	CheckLogin(ctx context.Context, in *InputLogin, opts ...grpc.CallOption) (*UserData, error)
	NewSignUp(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserData, error)
}

type dataClient struct {
	cc grpc.ClientConnInterface
}

func NewDataClient(cc grpc.ClientConnInterface) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) CheckLogin(ctx context.Context, in *InputLogin, opts ...grpc.CallOption) (*UserData, error) {
	out := new(UserData)
	err := c.cc.Invoke(ctx, "/SiswaService.Data/CheckLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) NewSignUp(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserData, error) {
	out := new(UserData)
	err := c.cc.Invoke(ctx, "/SiswaService.Data/NewSignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServer is the server API for Data service.
// All implementations must embed UnimplementedDataServer
// for forward compatibility
type DataServer interface {
	CheckLogin(context.Context, *InputLogin) (*UserData, error)
	NewSignUp(context.Context, *User) (*UserData, error)
	mustEmbedUnimplementedDataServer()
}

// UnimplementedDataServer must be embedded to have forward compatible implementations.
type UnimplementedDataServer struct {
}

func (*UnimplementedDataServer) CheckLogin(context.Context, *InputLogin) (*UserData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckLogin not implemented")
}
func (*UnimplementedDataServer) NewSignUp(context.Context, *User) (*UserData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewSignUp not implemented")
}
func (*UnimplementedDataServer) mustEmbedUnimplementedDataServer() {}

func RegisterDataServer(s *grpc.Server, srv DataServer) {
	s.RegisterService(&_Data_serviceDesc, srv)
}

func _Data_CheckLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).CheckLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiswaService.Data/CheckLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).CheckLogin(ctx, req.(*InputLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_NewSignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).NewSignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiswaService.Data/NewSignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).NewSignUp(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _Data_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SiswaService.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckLogin",
			Handler:    _Data_CheckLogin_Handler,
		},
		{
			MethodName: "NewSignUp",
			Handler:    _Data_NewSignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "siswaservice.proto",
}
