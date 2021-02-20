// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: tasks/internal/proto-files/service/tasks-service.proto

package service

import (
	context "context"
	messages "github.com/MaxPolarfox/tasks/pkg/grpc/messages"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AddRepositoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddedTask *messages.Task `protobuf:"bytes,1,opt,name=addedTask,proto3" json:"addedTask,omitempty"`
	Error     *Error         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddRepositoryResponse) Reset() {
	*x = AddRepositoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_internal_proto_files_service_tasks_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRepositoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRepositoryResponse) ProtoMessage() {}

func (x *AddRepositoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_internal_proto_files_service_tasks_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRepositoryResponse.ProtoReflect.Descriptor instead.
func (*AddRepositoryResponse) Descriptor() ([]byte, []int) {
	return file_tasks_internal_proto_files_service_tasks_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddRepositoryResponse) GetAddedTask() *messages.Task {
	if x != nil {
		return x.AddedTask
	}
	return nil
}

func (x *AddRepositoryResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type GetAllTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*messages.Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	Error *Error           `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetAllTasksResponse) Reset() {
	*x = GetAllTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_internal_proto_files_service_tasks_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTasksResponse) ProtoMessage() {}

func (x *GetAllTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_internal_proto_files_service_tasks_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTasksResponse.ProtoReflect.Descriptor instead.
func (*GetAllTasksResponse) Descriptor() ([]byte, []int) {
	return file_tasks_internal_proto_files_service_tasks_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllTasksResponse) GetTasks() []*messages.Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *GetAllTasksResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task  *messages.Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Error *Error         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_internal_proto_files_service_tasks_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_internal_proto_files_service_tasks_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_tasks_internal_proto_files_service_tasks_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteResponse) GetTask() *messages.Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *DeleteResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_internal_proto_files_service_tasks_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_internal_proto_files_service_tasks_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_tasks_internal_proto_files_service_tasks_service_proto_rawDescGZIP(), []int{3}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_tasks_internal_proto_files_service_tasks_service_proto protoreflect.FileDescriptor

var file_tasks_internal_proto_files_service_tasks_service_proto_rawDesc = []byte{
	0x0a, 0x36, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x09,
	0x61, 0x64, 0x64, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x24, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x61, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x24, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x5a, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x24, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x35,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xba, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x12, 0x0e, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x1e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06,
	0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tasks_internal_proto_files_service_tasks_service_proto_rawDescOnce sync.Once
	file_tasks_internal_proto_files_service_tasks_service_proto_rawDescData = file_tasks_internal_proto_files_service_tasks_service_proto_rawDesc
)

func file_tasks_internal_proto_files_service_tasks_service_proto_rawDescGZIP() []byte {
	file_tasks_internal_proto_files_service_tasks_service_proto_rawDescOnce.Do(func() {
		file_tasks_internal_proto_files_service_tasks_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_tasks_internal_proto_files_service_tasks_service_proto_rawDescData)
	})
	return file_tasks_internal_proto_files_service_tasks_service_proto_rawDescData
}

var file_tasks_internal_proto_files_service_tasks_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tasks_internal_proto_files_service_tasks_service_proto_goTypes = []interface{}{
	(*AddRepositoryResponse)(nil),  // 0: service.AddRepositoryResponse
	(*GetAllTasksResponse)(nil),    // 1: service.GetAllTasksResponse
	(*DeleteResponse)(nil),         // 2: service.DeleteResponse
	(*Error)(nil),                  // 3: service.Error
	(*messages.Task)(nil),          // 4: messages.Task
	(*messages.Action)(nil),        // 5: messages.Action
	(*messages.DeleteRequest)(nil), // 6: messages.DeleteRequest
}
var file_tasks_internal_proto_files_service_tasks_service_proto_depIdxs = []int32{
	4, // 0: service.AddRepositoryResponse.addedTask:type_name -> messages.Task
	3, // 1: service.AddRepositoryResponse.error:type_name -> service.Error
	4, // 2: service.GetAllTasksResponse.tasks:type_name -> messages.Task
	3, // 3: service.GetAllTasksResponse.error:type_name -> service.Error
	4, // 4: service.DeleteResponse.task:type_name -> messages.Task
	3, // 5: service.DeleteResponse.error:type_name -> service.Error
	4, // 6: service.TaskService.add:input_type -> messages.Task
	5, // 7: service.TaskService.getAll:input_type -> messages.Action
	6, // 8: service.TaskService.delete:input_type -> messages.DeleteRequest
	0, // 9: service.TaskService.add:output_type -> service.AddRepositoryResponse
	1, // 10: service.TaskService.getAll:output_type -> service.GetAllTasksResponse
	2, // 11: service.TaskService.delete:output_type -> service.DeleteResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tasks_internal_proto_files_service_tasks_service_proto_init() }
func file_tasks_internal_proto_files_service_tasks_service_proto_init() {
	if File_tasks_internal_proto_files_service_tasks_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tasks_internal_proto_files_service_tasks_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRepositoryResponse); i {
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
		file_tasks_internal_proto_files_service_tasks_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTasksResponse); i {
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
		file_tasks_internal_proto_files_service_tasks_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_tasks_internal_proto_files_service_tasks_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_tasks_internal_proto_files_service_tasks_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tasks_internal_proto_files_service_tasks_service_proto_goTypes,
		DependencyIndexes: file_tasks_internal_proto_files_service_tasks_service_proto_depIdxs,
		MessageInfos:      file_tasks_internal_proto_files_service_tasks_service_proto_msgTypes,
	}.Build()
	File_tasks_internal_proto_files_service_tasks_service_proto = out.File
	file_tasks_internal_proto_files_service_tasks_service_proto_rawDesc = nil
	file_tasks_internal_proto_files_service_tasks_service_proto_goTypes = nil
	file_tasks_internal_proto_files_service_tasks_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskServiceClient interface {
	Add(ctx context.Context, in *messages.Task, opts ...grpc.CallOption) (*AddRepositoryResponse, error)
	GetAll(ctx context.Context, in *messages.Action, opts ...grpc.CallOption) (*GetAllTasksResponse, error)
	Delete(ctx context.Context, in *messages.DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) Add(ctx context.Context, in *messages.Task, opts ...grpc.CallOption) (*AddRepositoryResponse, error) {
	out := new(AddRepositoryResponse)
	err := c.cc.Invoke(ctx, "/service.TaskService/add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetAll(ctx context.Context, in *messages.Action, opts ...grpc.CallOption) (*GetAllTasksResponse, error) {
	out := new(GetAllTasksResponse)
	err := c.cc.Invoke(ctx, "/service.TaskService/getAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Delete(ctx context.Context, in *messages.DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/service.TaskService/delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
type TaskServiceServer interface {
	Add(context.Context, *messages.Task) (*AddRepositoryResponse, error)
	GetAll(context.Context, *messages.Action) (*GetAllTasksResponse, error)
	Delete(context.Context, *messages.DeleteRequest) (*DeleteResponse, error)
}

// UnimplementedTaskServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (*UnimplementedTaskServiceServer) Add(context.Context, *messages.Task) (*AddRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedTaskServiceServer) GetAll(context.Context, *messages.Action) (*GetAllTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedTaskServiceServer) Delete(context.Context, *messages.DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(messages.Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.TaskService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Add(ctx, req.(*messages.Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(messages.Action)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.TaskService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetAll(ctx, req.(*messages.Action))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(messages.DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.TaskService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Delete(ctx, req.(*messages.DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _TaskService_Add_Handler,
		},
		{
			MethodName: "getAll",
			Handler:    _TaskService_GetAll_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _TaskService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tasks/internal/proto-files/service/tasks-service.proto",
}
