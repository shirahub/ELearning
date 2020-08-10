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

// ConnectServer is the server API for Connect service.
// All implementations must embed UnimplementedConnectServer
// for forward compatibility
type ConnectServer interface {
	ConnectToServer(context.Context, *empty.Empty) (*ServerReply, error)
	Soal(context.Context, *PilihanSoal) (*SoalList, error)
	Hasil(context.Context, *KumpulanJawaban) (*Result, error)
	AmbilTema(context.Context, *PilihTema) (*TemaList, error)
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
