// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.7.1
// source: siswaservice.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TemaList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemaList []string `protobuf:"bytes,1,rep,name=TemaList,proto3" json:"TemaList,omitempty"`
}

func (x *TemaList) Reset() {
	*x = TemaList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_siswaservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemaList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemaList) ProtoMessage() {}

func (x *TemaList) ProtoReflect() protoreflect.Message {
	mi := &file_siswaservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemaList.ProtoReflect.Descriptor instead.
func (*TemaList) Descriptor() ([]byte, []int) {
	return file_siswaservice_proto_rawDescGZIP(), []int{0}
}

func (x *TemaList) GetTemaList() []string {
	if x != nil {
		return x.TemaList
	}
	return nil
}

type PilihTema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tingkat string `protobuf:"bytes,1,opt,name=Tingkat,proto3" json:"Tingkat,omitempty"`
	Kelas   string `protobuf:"bytes,2,opt,name=Kelas,proto3" json:"Kelas,omitempty"`
	Matpel  string `protobuf:"bytes,3,opt,name=Matpel,proto3" json:"Matpel,omitempty"`
}

func (x *PilihTema) Reset() {
	*x = PilihTema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_siswaservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PilihTema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PilihTema) ProtoMessage() {}

func (x *PilihTema) ProtoReflect() protoreflect.Message {
	mi := &file_siswaservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PilihTema.ProtoReflect.Descriptor instead.
func (*PilihTema) Descriptor() ([]byte, []int) {
	return file_siswaservice_proto_rawDescGZIP(), []int{1}
}

func (x *PilihTema) GetTingkat() string {
	if x != nil {
		return x.Tingkat
	}
	return ""
}

func (x *PilihTema) GetKelas() string {
	if x != nil {
		return x.Kelas
	}
	return ""
}

func (x *PilihTema) GetMatpel() string {
	if x != nil {
		return x.Matpel
	}
	return ""
}

//ini seperti objek
type ServerGuruReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ServerGuruReply) Reset() {
	*x = ServerGuruReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_siswaservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerGuruReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerGuruReply) ProtoMessage() {}

func (x *ServerGuruReply) ProtoReflect() protoreflect.Message {
	mi := &file_siswaservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerGuruReply.ProtoReflect.Descriptor instead.
func (*ServerGuruReply) Descriptor() ([]byte, []int) {
	return file_siswaservice_proto_rawDescGZIP(), []int{2}
}

func (x *ServerGuruReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PaketSoal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tingkat string  `protobuf:"bytes,1,opt,name=Tingkat,proto3" json:"Tingkat,omitempty"`
	Kelas   string  `protobuf:"bytes,2,opt,name=Kelas,proto3" json:"Kelas,omitempty"`
	Matpel  string  `protobuf:"bytes,3,opt,name=Matpel,proto3" json:"Matpel,omitempty"`
	Tema    string  `protobuf:"bytes,4,opt,name=Tema,proto3" json:"Tema,omitempty"`
	Soal    []*Soal `protobuf:"bytes,5,rep,name=Soal,proto3" json:"Soal,omitempty"`
}

func (x *PaketSoal) Reset() {
	*x = PaketSoal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_siswaservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaketSoal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaketSoal) ProtoMessage() {}

func (x *PaketSoal) ProtoReflect() protoreflect.Message {
	mi := &file_siswaservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaketSoal.ProtoReflect.Descriptor instead.
func (*PaketSoal) Descriptor() ([]byte, []int) {
	return file_siswaservice_proto_rawDescGZIP(), []int{3}
}

func (x *PaketSoal) GetTingkat() string {
	if x != nil {
		return x.Tingkat
	}
	return ""
}

func (x *PaketSoal) GetKelas() string {
	if x != nil {
		return x.Kelas
	}
	return ""
}

func (x *PaketSoal) GetMatpel() string {
	if x != nil {
		return x.Matpel
	}
	return ""
}

func (x *PaketSoal) GetTema() string {
	if x != nil {
		return x.Tema
	}
	return ""
}

func (x *PaketSoal) GetSoal() []*Soal {
	if x != nil {
		return x.Soal
	}
	return nil
}

type ServerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ServerReply) Reset() {
	*x = ServerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_siswaservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerReply) ProtoMessage() {}

func (x *ServerReply) ProtoReflect() protoreflect.Message {
	mi := &file_siswaservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerReply.ProtoReflect.Descriptor instead.
func (*ServerReply) Descriptor() ([]byte, []int) {
	return file_siswaservice_proto_rawDescGZIP(), []int{4}
}

func (x *ServerReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PilihanSoal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tingkat  string `protobuf:"bytes,1,opt,name=Tingkat,proto3" json:"Tingkat,omitempty"`
	Kelas    string `protobuf:"bytes,2,opt,name=Kelas,proto3" json:"Kelas,omitempty"`
	Mapel    string `protobuf:"bytes,3,opt,name=Mapel,proto3" json:"Mapel,omitempty"`
	TemaSoal string `protobuf:"bytes,4,opt,name=TemaSoal,proto3" json:"TemaSoal,omitempty"`
}

func (x *PilihanSoal) Reset() {
	*x = PilihanSoal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_siswaservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PilihanSoal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PilihanSoal) ProtoMessage() {}

func (x *PilihanSoal) ProtoReflect() protoreflect.Message {
	mi := &file_siswaservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PilihanSoal.ProtoReflect.Descriptor instead.
func (*PilihanSoal) Descriptor() ([]byte, []int) {
	return file_siswaservice_proto_rawDescGZIP(), []int{5}
}

func (x *PilihanSoal) GetTingkat() string {
	if x != nil {
		return x.Tingkat
	}
	return ""
}

func (x *PilihanSoal) GetKelas() string {
	if x != nil {
		return x.Kelas
	}
	return ""
}

func (x *PilihanSoal) GetMapel() string {
	if x != nil {
		return x.Mapel
	}
	return ""
}

func (x *PilihanSoal) GetTemaSoal() string {
	if x != nil {
		return x.TemaSoal
	}
	return ""
}

type SoalList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SoalList []*Soal `protobuf:"bytes,1,rep,name=SoalList,proto3" json:"SoalList,omitempty"`
}

func (x *SoalList) Reset() {
	*x = SoalList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_siswaservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoalList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoalList) ProtoMessage() {}

func (x *SoalList) ProtoReflect() protoreflect.Message {
	mi := &file_siswaservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoalList.ProtoReflect.Descriptor instead.
func (*SoalList) Descriptor() ([]byte, []int) {
	return file_siswaservice_proto_rawDescGZIP(), []int{6}
}

func (x *SoalList) GetSoalList() []*Soal {
	if x != nil {
		return x.SoalList
	}
	return nil
}

type Soal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pertanyaan string   `protobuf:"bytes,1,opt,name=Pertanyaan,proto3" json:"Pertanyaan,omitempty"`
	Jawaban    string   `protobuf:"bytes,2,opt,name=Jawaban,proto3" json:"Jawaban,omitempty"`
	Pilihan    []string `protobuf:"bytes,3,rep,name=Pilihan,proto3" json:"Pilihan,omitempty"`
}

func (x *Soal) Reset() {
	*x = Soal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_siswaservice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Soal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Soal) ProtoMessage() {}

func (x *Soal) ProtoReflect() protoreflect.Message {
	mi := &file_siswaservice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Soal.ProtoReflect.Descriptor instead.
func (*Soal) Descriptor() ([]byte, []int) {
	return file_siswaservice_proto_rawDescGZIP(), []int{7}
}

func (x *Soal) GetPertanyaan() string {
	if x != nil {
		return x.Pertanyaan
	}
	return ""
}

func (x *Soal) GetJawaban() string {
	if x != nil {
		return x.Jawaban
	}
	return ""
}

func (x *Soal) GetPilihan() []string {
	if x != nil {
		return x.Pilihan
	}
	return nil
}

type KumpulanJawaban struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tingkat     string   `protobuf:"bytes,1,opt,name=Tingkat,proto3" json:"Tingkat,omitempty"`
	Kelas       string   `protobuf:"bytes,2,opt,name=Kelas,proto3" json:"Kelas,omitempty"`
	Mapel       string   `protobuf:"bytes,3,opt,name=Mapel,proto3" json:"Mapel,omitempty"`
	TemaSoal    string   `protobuf:"bytes,4,opt,name=TemaSoal,proto3" json:"TemaSoal,omitempty"`
	JawabanList []string `protobuf:"bytes,5,rep,name=JawabanList,proto3" json:"JawabanList,omitempty"`
}

func (x *KumpulanJawaban) Reset() {
	*x = KumpulanJawaban{}
	if protoimpl.UnsafeEnabled {
		mi := &file_siswaservice_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KumpulanJawaban) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KumpulanJawaban) ProtoMessage() {}

func (x *KumpulanJawaban) ProtoReflect() protoreflect.Message {
	mi := &file_siswaservice_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KumpulanJawaban.ProtoReflect.Descriptor instead.
func (*KumpulanJawaban) Descriptor() ([]byte, []int) {
	return file_siswaservice_proto_rawDescGZIP(), []int{8}
}

func (x *KumpulanJawaban) GetTingkat() string {
	if x != nil {
		return x.Tingkat
	}
	return ""
}

func (x *KumpulanJawaban) GetKelas() string {
	if x != nil {
		return x.Kelas
	}
	return ""
}

func (x *KumpulanJawaban) GetMapel() string {
	if x != nil {
		return x.Mapel
	}
	return ""
}

func (x *KumpulanJawaban) GetTemaSoal() string {
	if x != nil {
		return x.TemaSoal
	}
	return ""
}

func (x *KumpulanJawaban) GetJawabanList() []string {
	if x != nil {
		return x.JawabanList
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_siswaservice_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_siswaservice_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_siswaservice_proto_rawDescGZIP(), []int{9}
}

func (x *Result) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_siswaservice_proto protoreflect.FileDescriptor

var file_siswaservice_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x69, 0x73, 0x77, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x53, 0x69, 0x73, 0x77, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x26, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54,
	0x65, 0x6d, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x54,
	0x65, 0x6d, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x09, 0x50, 0x69, 0x6c, 0x69, 0x68,
	0x54, 0x65, 0x6d, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x4b, 0x65, 0x6c, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4b,
	0x65, 0x6c, 0x61, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x61, 0x74, 0x70, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x61, 0x74, 0x70, 0x65, 0x6c, 0x22, 0x2b, 0x0a, 0x0f,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x72, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x09, 0x50, 0x61,
	0x6b, 0x65, 0x74, 0x53, 0x6f, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6e, 0x67, 0x6b,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x69, 0x6e, 0x67, 0x6b, 0x61,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x4b, 0x65, 0x6c, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x4b, 0x65, 0x6c, 0x61, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x61, 0x74, 0x70, 0x65,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x61, 0x74, 0x70, 0x65, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x65, 0x6d, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x65, 0x6d, 0x61, 0x12, 0x26, 0x0a, 0x04, 0x53, 0x6f, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x53, 0x69, 0x73, 0x77, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x6f, 0x61, 0x6c, 0x52, 0x04, 0x53, 0x6f, 0x61, 0x6c, 0x22, 0x27, 0x0a, 0x0b, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x6f, 0x0a, 0x0b, 0x50, 0x69, 0x6c, 0x69, 0x68, 0x61, 0x6e, 0x53,
	0x6f, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x4b, 0x65, 0x6c, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4b, 0x65,
	0x6c, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x70, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x4d, 0x61, 0x70, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x6d,
	0x61, 0x53, 0x6f, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x6d,
	0x61, 0x53, 0x6f, 0x61, 0x6c, 0x22, 0x3a, 0x0a, 0x08, 0x53, 0x6f, 0x61, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x08, 0x53, 0x6f, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x69, 0x73, 0x77, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x6f, 0x61, 0x6c, 0x52, 0x08, 0x53, 0x6f, 0x61, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x5a, 0x0a, 0x04, 0x53, 0x6f, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x65, 0x72,
	0x74, 0x61, 0x6e, 0x79, 0x61, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50,
	0x65, 0x72, 0x74, 0x61, 0x6e, 0x79, 0x61, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x4a, 0x61, 0x77,
	0x61, 0x62, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4a, 0x61, 0x77, 0x61,
	0x62, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x69, 0x6c, 0x69, 0x68, 0x61, 0x6e, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x50, 0x69, 0x6c, 0x69, 0x68, 0x61, 0x6e, 0x22, 0x95, 0x01,
	0x0a, 0x0f, 0x4b, 0x75, 0x6d, 0x70, 0x75, 0x6c, 0x61, 0x6e, 0x4a, 0x61, 0x77, 0x61, 0x62, 0x61,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x54, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4b,
	0x65, 0x6c, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4b, 0x65, 0x6c, 0x61,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x70, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x4d, 0x61, 0x70, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x61, 0x53,
	0x6f, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x6d, 0x61, 0x53,
	0x6f, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x61, 0x77, 0x61, 0x62, 0x61, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x4a, 0x61, 0x77, 0x61, 0x62, 0x61,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x20, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x8e, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x46, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19,
	0x2e, 0x53, 0x69, 0x73, 0x77, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x53,
	0x6f, 0x61, 0x6c, 0x12, 0x19, 0x2e, 0x53, 0x69, 0x73, 0x77, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x69, 0x6c, 0x69, 0x68, 0x61, 0x6e, 0x53, 0x6f, 0x61, 0x6c, 0x1a, 0x16,
	0x2e, 0x53, 0x69, 0x73, 0x77, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x6f,
	0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x05, 0x48, 0x61, 0x73, 0x69,
	0x6c, 0x12, 0x1d, 0x2e, 0x53, 0x69, 0x73, 0x77, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4b, 0x75, 0x6d, 0x70, 0x75, 0x6c, 0x61, 0x6e, 0x4a, 0x61, 0x77, 0x61, 0x62, 0x61, 0x6e,
	0x1a, 0x14, 0x2e, 0x53, 0x69, 0x73, 0x77, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x41, 0x6d, 0x62, 0x69,
	0x6c, 0x54, 0x65, 0x6d, 0x61, 0x12, 0x17, 0x2e, 0x53, 0x69, 0x73, 0x77, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x69, 0x6c, 0x69, 0x68, 0x54, 0x65, 0x6d, 0x61, 0x1a, 0x16,
	0x2e, 0x53, 0x69, 0x73, 0x77, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65,
	0x6d, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x32, 0x55, 0x0a, 0x04, 0x47, 0x75, 0x72, 0x75,
	0x12, 0x4d, 0x0a, 0x11, 0x4b, 0x69, 0x72, 0x69, 0x6d, 0x53, 0x6f, 0x61, 0x6c, 0x4b, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x53, 0x69, 0x73, 0x77, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x6b, 0x65, 0x74, 0x53, 0x6f, 0x61, 0x6c, 0x1a, 0x1d,
	0x2e, 0x53, 0x69, 0x73, 0x77, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x72, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_siswaservice_proto_rawDescOnce sync.Once
	file_siswaservice_proto_rawDescData = file_siswaservice_proto_rawDesc
)

func file_siswaservice_proto_rawDescGZIP() []byte {
	file_siswaservice_proto_rawDescOnce.Do(func() {
		file_siswaservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_siswaservice_proto_rawDescData)
	})
	return file_siswaservice_proto_rawDescData
}

var file_siswaservice_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_siswaservice_proto_goTypes = []interface{}{
	(*TemaList)(nil),        // 0: SiswaService.TemaList
	(*PilihTema)(nil),       // 1: SiswaService.PilihTema
	(*ServerGuruReply)(nil), // 2: SiswaService.ServerGuruReply
	(*PaketSoal)(nil),       // 3: SiswaService.PaketSoal
	(*ServerReply)(nil),     // 4: SiswaService.ServerReply
	(*PilihanSoal)(nil),     // 5: SiswaService.PilihanSoal
	(*SoalList)(nil),        // 6: SiswaService.SoalList
	(*Soal)(nil),            // 7: SiswaService.Soal
	(*KumpulanJawaban)(nil), // 8: SiswaService.KumpulanJawaban
	(*Result)(nil),          // 9: SiswaService.Result
	(*empty.Empty)(nil),     // 10: google.protobuf.Empty
}
var file_siswaservice_proto_depIdxs = []int32{
	7,  // 0: SiswaService.PaketSoal.Soal:type_name -> SiswaService.Soal
	7,  // 1: SiswaService.SoalList.SoalList:type_name -> SiswaService.Soal
	10, // 2: SiswaService.Connect.ConnectToServer:input_type -> google.protobuf.Empty
	5,  // 3: SiswaService.Connect.Soal:input_type -> SiswaService.PilihanSoal
	8,  // 4: SiswaService.Connect.Hasil:input_type -> SiswaService.KumpulanJawaban
	1,  // 5: SiswaService.Connect.AmbilTema:input_type -> SiswaService.PilihTema
	3,  // 6: SiswaService.Guru.KirimSoalKeServer:input_type -> SiswaService.PaketSoal
	4,  // 7: SiswaService.Connect.ConnectToServer:output_type -> SiswaService.ServerReply
	6,  // 8: SiswaService.Connect.Soal:output_type -> SiswaService.SoalList
	9,  // 9: SiswaService.Connect.Hasil:output_type -> SiswaService.Result
	0,  // 10: SiswaService.Connect.AmbilTema:output_type -> SiswaService.TemaList
	2,  // 11: SiswaService.Guru.KirimSoalKeServer:output_type -> SiswaService.ServerGuruReply
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_siswaservice_proto_init() }
func file_siswaservice_proto_init() {
	if File_siswaservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_siswaservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemaList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_siswaservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PilihTema); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_siswaservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerGuruReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_siswaservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaketSoal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_siswaservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_siswaservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PilihanSoal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_siswaservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoalList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_siswaservice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Soal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_siswaservice_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KumpulanJawaban); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_siswaservice_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_siswaservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_siswaservice_proto_goTypes,
		DependencyIndexes: file_siswaservice_proto_depIdxs,
		MessageInfos:      file_siswaservice_proto_msgTypes,
	}.Build()
	File_siswaservice_proto = out.File
	file_siswaservice_proto_rawDesc = nil
	file_siswaservice_proto_goTypes = nil
	file_siswaservice_proto_depIdxs = nil
}
