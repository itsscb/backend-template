// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: service_backend.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_backend_proto protoreflect.FileDescriptor

var file_service_backend_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x72, 0x70, 0x63, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72, 0x70, 0x63, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x72,
	0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x70, 0x63, 0x5f,
	0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x72, 0x70, 0x63, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72, 0x70, 0x63, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x27, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x62, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x72, 0x70, 0x63,
	0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xfa, 0x1a, 0x0a, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x42, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x68, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22,
	0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0xa4, 0x01, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x61, 0x92, 0x41, 0x2f, 0x12, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x20, 0x62, 0x79, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x92, 0x01, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x92, 0x41, 0x27, 0x12, 0x13, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x20, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44,
	0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x32, 0x1a, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x92, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x92, 0x41, 0x2d, 0x12, 0x19, 0x47, 0x65, 0x74, 0x20,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x62, 0x79, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x96, 0x01, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x17, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x53, 0x92, 0x41, 0x2e, 0x12, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x20, 0x5b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x20, 0x6f, 0x6e, 0x6c, 0x79,
	0x5d, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x7f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x92, 0x41, 0x10,
	0x12, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x91, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x92,
	0x41, 0x22, 0x12, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x32, 0x1b, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0xbf, 0x01, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x64, 0x92, 0x41, 0x33, 0x12, 0x1f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x20, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x20, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x62, 0x10, 0x0a, 0x0e,
	0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x32, 0x23, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x8b, 0x01, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x48, 0x92, 0x41, 0x21, 0x12, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22,
	0x19, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x8b, 0x01, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x48,
	0x92, 0x41, 0x21, 0x12, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x32, 0x19, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x84, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x4a, 0x92, 0x41, 0x24, 0x12, 0x10, 0x47, 0x65, 0x74, 0x20, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x65, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x93, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x50, 0x92, 0x41, 0x27, 0x12, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x20, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x62, 0x10, 0x0a,
	0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x20, 0x2a, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x9e, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x92, 0x41, 0x2e, 0x12, 0x1a, 0x4c, 0x69, 0x73,
	0x74, 0x20, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x20, 0x62, 0x79, 0x20, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12,
	0x25, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x91, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x92,
	0x41, 0x22, 0x12, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x8a, 0x01, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x92, 0x41, 0x25, 0x12, 0x11, 0x47,
	0x65, 0x74, 0x20, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44,
	0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x99, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53,
	0x92, 0x41, 0x28, 0x12, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x22, 0x2a, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0xa4, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61, 0x92, 0x41, 0x2f, 0x12, 0x1b, 0x4c, 0x69,
	0x73, 0x74, 0x20, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x62, 0x79, 0x20, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x29, 0x12, 0x27, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x91, 0x01, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4b, 0x92, 0x41, 0x22, 0x12, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01,
	0x2a, 0x32, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0xb0,
	0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x4c, 0x6f,
	0x67, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x67, 0x92, 0x41, 0x30, 0x12, 0x1c, 0x4c,
	0x69, 0x73, 0x74, 0x20, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x4c, 0x6f, 0x67, 0x20, 0x62,
	0x79, 0x20, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x62, 0x10, 0x0a, 0x0e, 0x0a,
	0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2e, 0x12, 0x2c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73,
	0x5f, 0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x2f, 0x7b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0xca, 0x02, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x80, 0x02, 0x92, 0x41,
	0xe0, 0x01, 0x12, 0x1b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x20, 0x5b, 0x6f, 0x6e, 0x6c, 0x79, 0x20, 0x48, 0x54, 0x54, 0x50, 0x5d, 0x1a,
	0xae, 0x01, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x76, 0x69, 0x61, 0x20, 0x73, 0x77,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x70, 0x6f, 0x73,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x2e, 0x20, 0x54, 0x72, 0x79, 0x20, 0x60, 0x60, 0x60, 0x63, 0x75,
	0x72, 0x6c, 0x20, 0x2d, 0x58, 0x20, 0x50, 0x4f, 0x53, 0x54, 0x20, 0x2d, 0x48, 0x20, 0x22, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x20, 0x7b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x7d, 0x22, 0x20, 0x2d, 0x46,
	0x20, 0x22, 0x66, 0x69, 0x6c, 0x65, 0x3d, 0x40, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x74, 0x6f,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x20, 0x2d, 0x46, 0x20, 0x22, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x3d, 0x31, 0x22, 0x20, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f,
	0x7b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x52, 0x49, 0x7d, 0x2f, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x60, 0x60, 0x60,
	0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x9f,
	0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x92, 0x41, 0x29, 0x12, 0x15, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x62,
	0x79, 0x20, 0x49, 0x44, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x2a, 0x22, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x1a, 0x0c, 0x92, 0x41, 0x09, 0x12, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x42, 0xd1,
	0x01, 0x92, 0x41, 0xa6, 0x01, 0x12, 0x57, 0x0a, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x20, 0x41, 0x50, 0x49, 0x22, 0x43, 0x0a, 0x06, 0x69, 0x74, 0x73, 0x73, 0x63, 0x62, 0x12, 0x2a,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x73, 0x73, 0x63, 0x62, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x0d, 0x64, 0x65, 0x76, 0x40,
	0x69, 0x74, 0x73, 0x73, 0x63, 0x62, 0x2e, 0x64, 0x65, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02,
	0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x23, 0x0a, 0x21, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x73, 0x73, 0x63, 0x62, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_backend_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),                 // 0: pb.LoginRequest
	(*RefreshTokenRequest)(nil),          // 1: pb.RefreshTokenRequest
	(*ListSessionsRequest)(nil),          // 2: pb.ListSessionsRequest
	(*BlockSessionRequest)(nil),          // 3: pb.BlockSessionRequest
	(*GetAccountRequest)(nil),            // 4: pb.GetAccountRequest
	(*ListAccountsRequest)(nil),          // 5: pb.ListAccountsRequest
	(*CreateAccountRequest)(nil),         // 6: pb.CreateAccountRequest
	(*UpdateAccountRequest)(nil),         // 7: pb.UpdateAccountRequest
	(*UpdateAccountPrivacyRequest)(nil),  // 8: pb.UpdateAccountPrivacyRequest
	(*CreatePersonRequest)(nil),          // 9: pb.CreatePersonRequest
	(*UpdatePersonRequest)(nil),          // 10: pb.UpdatePersonRequest
	(*GetPersonRequest)(nil),             // 11: pb.GetPersonRequest
	(*DeletePersonRequest)(nil),          // 12: pb.DeletePersonRequest
	(*ListPersonsRequest)(nil),           // 13: pb.ListPersonsRequest
	(*CreatePaymentRequest)(nil),         // 14: pb.CreatePaymentRequest
	(*GetPaymentRequest)(nil),            // 15: pb.GetPaymentRequest
	(*DeletePaymentRequest)(nil),         // 16: pb.DeletePaymentRequest
	(*ListPaymentsRequest)(nil),          // 17: pb.ListPaymentsRequest
	(*UpdatePaymentRequest)(nil),         // 18: pb.UpdatePaymentRequest
	(*ListReturnsLogRequest)(nil),        // 19: pb.ListReturnsLogRequest
	(*UploadDocumentRequest)(nil),        // 20: pb.UploadDocumentRequest
	(*DeleteDocumentRequest)(nil),        // 21: pb.DeleteDocumentRequest
	(*LoginResponse)(nil),                // 22: pb.LoginResponse
	(*RefreshTokenResponse)(nil),         // 23: pb.RefreshTokenResponse
	(*ListSessionsResponse)(nil),         // 24: pb.ListSessionsResponse
	(*BlockSessionResponse)(nil),         // 25: pb.BlockSessionResponse
	(*GetAccountResponse)(nil),           // 26: pb.GetAccountResponse
	(*ListAccountsResponse)(nil),         // 27: pb.ListAccountsResponse
	(*CreateAccountResponse)(nil),        // 28: pb.CreateAccountResponse
	(*UpdateAccountResponse)(nil),        // 29: pb.UpdateAccountResponse
	(*UpdateAccountPrivacyResponse)(nil), // 30: pb.UpdateAccountPrivacyResponse
	(*CreatePersonResponse)(nil),         // 31: pb.CreatePersonResponse
	(*UpdatePersonResponse)(nil),         // 32: pb.UpdatePersonResponse
	(*GetPersonResponse)(nil),            // 33: pb.GetPersonResponse
	(*DeletePersonResponse)(nil),         // 34: pb.DeletePersonResponse
	(*ListPersonsResponse)(nil),          // 35: pb.ListPersonsResponse
	(*CreatePaymentResponse)(nil),        // 36: pb.CreatePaymentResponse
	(*GetPaymentResponse)(nil),           // 37: pb.GetPaymentResponse
	(*DeletePaymentResponse)(nil),        // 38: pb.DeletePaymentResponse
	(*ListPaymentsResponse)(nil),         // 39: pb.ListPaymentsResponse
	(*UpdatePaymentResponse)(nil),        // 40: pb.UpdatePaymentResponse
	(*ListReturnsLogResponse)(nil),       // 41: pb.ListReturnsLogResponse
	(*UploadDocumentResponse)(nil),       // 42: pb.UploadDocumentResponse
	(*DeleteDocumentResponse)(nil),       // 43: pb.DeleteDocumentResponse
}
var file_service_backend_proto_depIdxs = []int32{
	0,  // 0: pb.backend.Login:input_type -> pb.LoginRequest
	1,  // 1: pb.backend.RefreshToken:input_type -> pb.RefreshTokenRequest
	2,  // 2: pb.backend.ListSessions:input_type -> pb.ListSessionsRequest
	3,  // 3: pb.backend.BlockSession:input_type -> pb.BlockSessionRequest
	4,  // 4: pb.backend.GetAccount:input_type -> pb.GetAccountRequest
	5,  // 5: pb.backend.ListAccounts:input_type -> pb.ListAccountsRequest
	6,  // 6: pb.backend.CreateAccount:input_type -> pb.CreateAccountRequest
	7,  // 7: pb.backend.UpdateAccount:input_type -> pb.UpdateAccountRequest
	8,  // 8: pb.backend.UpdateAccountPrivacy:input_type -> pb.UpdateAccountPrivacyRequest
	9,  // 9: pb.backend.CreatePerson:input_type -> pb.CreatePersonRequest
	10, // 10: pb.backend.UpdatePerson:input_type -> pb.UpdatePersonRequest
	11, // 11: pb.backend.GetPerson:input_type -> pb.GetPersonRequest
	12, // 12: pb.backend.DeletePerson:input_type -> pb.DeletePersonRequest
	13, // 13: pb.backend.ListPersons:input_type -> pb.ListPersonsRequest
	14, // 14: pb.backend.CreatePayment:input_type -> pb.CreatePaymentRequest
	15, // 15: pb.backend.GetPayment:input_type -> pb.GetPaymentRequest
	16, // 16: pb.backend.DeletePayment:input_type -> pb.DeletePaymentRequest
	17, // 17: pb.backend.ListPayments:input_type -> pb.ListPaymentsRequest
	18, // 18: pb.backend.UpdatePayment:input_type -> pb.UpdatePaymentRequest
	19, // 19: pb.backend.ListReturnsLog:input_type -> pb.ListReturnsLogRequest
	20, // 20: pb.backend.UploadDocument:input_type -> pb.UploadDocumentRequest
	21, // 21: pb.backend.DeleteDocument:input_type -> pb.DeleteDocumentRequest
	22, // 22: pb.backend.Login:output_type -> pb.LoginResponse
	23, // 23: pb.backend.RefreshToken:output_type -> pb.RefreshTokenResponse
	24, // 24: pb.backend.ListSessions:output_type -> pb.ListSessionsResponse
	25, // 25: pb.backend.BlockSession:output_type -> pb.BlockSessionResponse
	26, // 26: pb.backend.GetAccount:output_type -> pb.GetAccountResponse
	27, // 27: pb.backend.ListAccounts:output_type -> pb.ListAccountsResponse
	28, // 28: pb.backend.CreateAccount:output_type -> pb.CreateAccountResponse
	29, // 29: pb.backend.UpdateAccount:output_type -> pb.UpdateAccountResponse
	30, // 30: pb.backend.UpdateAccountPrivacy:output_type -> pb.UpdateAccountPrivacyResponse
	31, // 31: pb.backend.CreatePerson:output_type -> pb.CreatePersonResponse
	32, // 32: pb.backend.UpdatePerson:output_type -> pb.UpdatePersonResponse
	33, // 33: pb.backend.GetPerson:output_type -> pb.GetPersonResponse
	34, // 34: pb.backend.DeletePerson:output_type -> pb.DeletePersonResponse
	35, // 35: pb.backend.ListPersons:output_type -> pb.ListPersonsResponse
	36, // 36: pb.backend.CreatePayment:output_type -> pb.CreatePaymentResponse
	37, // 37: pb.backend.GetPayment:output_type -> pb.GetPaymentResponse
	38, // 38: pb.backend.DeletePayment:output_type -> pb.DeletePaymentResponse
	39, // 39: pb.backend.ListPayments:output_type -> pb.ListPaymentsResponse
	40, // 40: pb.backend.UpdatePayment:output_type -> pb.UpdatePaymentResponse
	41, // 41: pb.backend.ListReturnsLog:output_type -> pb.ListReturnsLogResponse
	42, // 42: pb.backend.UploadDocument:output_type -> pb.UploadDocumentResponse
	43, // 43: pb.backend.DeleteDocument:output_type -> pb.DeleteDocumentResponse
	22, // [22:44] is the sub-list for method output_type
	0,  // [0:22] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_backend_proto_init() }
func file_service_backend_proto_init() {
	if File_service_backend_proto != nil {
		return
	}
	file_rpc_create_payment_proto_init()
	file_rpc_get_payment_proto_init()
	file_rpc_list_payments_proto_init()
	file_rpc_update_payment_proto_init()
	file_rpc_delete_payment_proto_init()
	file_rpc_create_person_proto_init()
	file_rpc_get_person_proto_init()
	file_rpc_list_persons_proto_init()
	file_rpc_update_person_proto_init()
	file_rpc_delete_person_proto_init()
	file_rpc_create_account_proto_init()
	file_rpc_get_account_proto_init()
	file_rpc_list_accounts_proto_init()
	file_rpc_update_account_proto_init()
	file_rpc_update_account_privacy_proto_init()
	file_rpc_login_proto_init()
	file_rpc_list_sessions_proto_init()
	file_rpc_refresh_token_proto_init()
	file_rpc_block_session_proto_init()
	file_rpc_list_returns_log_by_person_id_proto_init()
	file_rpc_upload_document_proto_init()
	file_rpc_delete_document_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_backend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_backend_proto_goTypes,
		DependencyIndexes: file_service_backend_proto_depIdxs,
	}.Build()
	File_service_backend_proto = out.File
	file_service_backend_proto_rawDesc = nil
	file_service_backend_proto_goTypes = nil
	file_service_backend_proto_depIdxs = nil
}
