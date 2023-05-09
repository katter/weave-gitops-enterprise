// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: api/pipelines/pipelines.proto

package api

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListPipelinesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ListPipelinesRequest) Reset() {
	*x = ListPipelinesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pipelines_pipelines_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPipelinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPipelinesRequest) ProtoMessage() {}

func (x *ListPipelinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pipelines_pipelines_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPipelinesRequest.ProtoReflect.Descriptor instead.
func (*ListPipelinesRequest) Descriptor() ([]byte, []int) {
	return file_api_pipelines_pipelines_proto_rawDescGZIP(), []int{0}
}

func (x *ListPipelinesRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type ListPipelinesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pipelines []*Pipeline  `protobuf:"bytes,1,rep,name=pipelines,proto3" json:"pipelines,omitempty"`
	Errors    []*ListError `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ListPipelinesResponse) Reset() {
	*x = ListPipelinesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pipelines_pipelines_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPipelinesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPipelinesResponse) ProtoMessage() {}

func (x *ListPipelinesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_pipelines_pipelines_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPipelinesResponse.ProtoReflect.Descriptor instead.
func (*ListPipelinesResponse) Descriptor() ([]byte, []int) {
	return file_api_pipelines_pipelines_proto_rawDescGZIP(), []int{1}
}

func (x *ListPipelinesResponse) GetPipelines() []*Pipeline {
	if x != nil {
		return x.Pipelines
	}
	return nil
}

func (x *ListPipelinesResponse) GetErrors() []*ListError {
	if x != nil {
		return x.Errors
	}
	return nil
}

type GetPipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *GetPipelineRequest) Reset() {
	*x = GetPipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pipelines_pipelines_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPipelineRequest) ProtoMessage() {}

func (x *GetPipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pipelines_pipelines_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPipelineRequest.ProtoReflect.Descriptor instead.
func (*GetPipelineRequest) Descriptor() ([]byte, []int) {
	return file_api_pipelines_pipelines_proto_rawDescGZIP(), []int{2}
}

func (x *GetPipelineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPipelineRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type GetPipelineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	Errors   []string  `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *GetPipelineResponse) Reset() {
	*x = GetPipelineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pipelines_pipelines_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPipelineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPipelineResponse) ProtoMessage() {}

func (x *GetPipelineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_pipelines_pipelines_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPipelineResponse.ProtoReflect.Descriptor instead.
func (*GetPipelineResponse) Descriptor() ([]byte, []int) {
	return file_api_pipelines_pipelines_proto_rawDescGZIP(), []int{3}
}

func (x *GetPipelineResponse) GetPipeline() *Pipeline {
	if x != nil {
		return x.Pipeline
	}
	return nil
}

func (x *GetPipelineResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

type ApprovePromotionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Env       string `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"`
	Revision  string `protobuf:"bytes,4,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *ApprovePromotionRequest) Reset() {
	*x = ApprovePromotionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pipelines_pipelines_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovePromotionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovePromotionRequest) ProtoMessage() {}

func (x *ApprovePromotionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pipelines_pipelines_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovePromotionRequest.ProtoReflect.Descriptor instead.
func (*ApprovePromotionRequest) Descriptor() ([]byte, []int) {
	return file_api_pipelines_pipelines_proto_rawDescGZIP(), []int{4}
}

func (x *ApprovePromotionRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ApprovePromotionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApprovePromotionRequest) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *ApprovePromotionRequest) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

type ApprovePromotionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PullRequestURL string `protobuf:"bytes,1,opt,name=pullRequestURL,proto3" json:"pullRequestURL,omitempty"`
}

func (x *ApprovePromotionResponse) Reset() {
	*x = ApprovePromotionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pipelines_pipelines_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovePromotionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovePromotionResponse) ProtoMessage() {}

func (x *ApprovePromotionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_pipelines_pipelines_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovePromotionResponse.ProtoReflect.Descriptor instead.
func (*ApprovePromotionResponse) Descriptor() ([]byte, []int) {
	return file_api_pipelines_pipelines_proto_rawDescGZIP(), []int{5}
}

func (x *ApprovePromotionResponse) GetPullRequestURL() string {
	if x != nil {
		return x.PullRequestURL
	}
	return ""
}

type ListError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ListError) Reset() {
	*x = ListError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pipelines_pipelines_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListError) ProtoMessage() {}

func (x *ListError) ProtoReflect() protoreflect.Message {
	mi := &file_api_pipelines_pipelines_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListError.ProtoReflect.Descriptor instead.
func (*ListError) Descriptor() ([]byte, []int) {
	return file_api_pipelines_pipelines_proto_rawDescGZIP(), []int{6}
}

func (x *ListError) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListPullRequestsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipelineName      string `protobuf:"bytes,1,opt,name=pipelineName,proto3" json:"pipelineName,omitempty"`
	PipelineNamespace string `protobuf:"bytes,2,opt,name=pipelineNamespace,proto3" json:"pipelineNamespace,omitempty"`
}

func (x *ListPullRequestsRequest) Reset() {
	*x = ListPullRequestsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pipelines_pipelines_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPullRequestsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPullRequestsRequest) ProtoMessage() {}

func (x *ListPullRequestsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pipelines_pipelines_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPullRequestsRequest.ProtoReflect.Descriptor instead.
func (*ListPullRequestsRequest) Descriptor() ([]byte, []int) {
	return file_api_pipelines_pipelines_proto_rawDescGZIP(), []int{7}
}

func (x *ListPullRequestsRequest) GetPipelineName() string {
	if x != nil {
		return x.PipelineName
	}
	return ""
}

func (x *ListPullRequestsRequest) GetPipelineNamespace() string {
	if x != nil {
		return x.PipelineNamespace
	}
	return ""
}

type ListPullRequestsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PullRequests map[string]string `protobuf:"bytes,1,rep,name=pullRequests,proto3" json:"pullRequests,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListPullRequestsResponse) Reset() {
	*x = ListPullRequestsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pipelines_pipelines_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPullRequestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPullRequestsResponse) ProtoMessage() {}

func (x *ListPullRequestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_pipelines_pipelines_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPullRequestsResponse.ProtoReflect.Descriptor instead.
func (*ListPullRequestsResponse) Descriptor() ([]byte, []int) {
	return file_api_pipelines_pipelines_proto_rawDescGZIP(), []int{8}
}

func (x *ListPullRequestsResponse) GetPullRequests() map[string]string {
	if x != nil {
		return x.PullRequests
	}
	return nil
}

var File_api_pipelines_pipelines_proto protoreflect.FileDescriptor

var file_api_pipelines_pipelines_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2f,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x7e, 0x0a, 0x15,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x09, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x46, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x22, 0x61, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x79, 0x0a, 0x17, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x18, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x70, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x52, 0x4c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x22, 0x43, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6b, 0x0a, 0x17, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0c, 0x70, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x70, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0x91, 0x04, 0x0a, 0x09, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x73, 0x12, 0x6f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x12, 0x70, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x20, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12,
	0x14, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x8a, 0x01, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x12, 0x93, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x25, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01,
	0x2a, 0x22, 0x25, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x73, 0x2f, 0x7b, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0xbe, 0x01, 0x92, 0x41, 0x7e, 0x12, 0x58,
	0x0a, 0x1a, 0x57, 0x65, 0x61, 0x76, 0x65, 0x20, 0x47, 0x69, 0x74, 0x4f, 0x70, 0x73, 0x20, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x20, 0x41, 0x50, 0x49, 0x12, 0x35, 0x54, 0x68,
	0x65, 0x20, 0x41, 0x50, 0x49, 0x20, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x20, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x57, 0x65, 0x61,
	0x76, 0x65, 0x20, 0x47, 0x69, 0x74, 0x4f, 0x70, 0x73, 0x20, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2d, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73,
	0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_pipelines_pipelines_proto_rawDescOnce sync.Once
	file_api_pipelines_pipelines_proto_rawDescData = file_api_pipelines_pipelines_proto_rawDesc
)

func file_api_pipelines_pipelines_proto_rawDescGZIP() []byte {
	file_api_pipelines_pipelines_proto_rawDescOnce.Do(func() {
		file_api_pipelines_pipelines_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_pipelines_pipelines_proto_rawDescData)
	})
	return file_api_pipelines_pipelines_proto_rawDescData
}

var file_api_pipelines_pipelines_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_pipelines_pipelines_proto_goTypes = []interface{}{
	(*ListPipelinesRequest)(nil),     // 0: pipelines.v1.ListPipelinesRequest
	(*ListPipelinesResponse)(nil),    // 1: pipelines.v1.ListPipelinesResponse
	(*GetPipelineRequest)(nil),       // 2: pipelines.v1.GetPipelineRequest
	(*GetPipelineResponse)(nil),      // 3: pipelines.v1.GetPipelineResponse
	(*ApprovePromotionRequest)(nil),  // 4: pipelines.v1.ApprovePromotionRequest
	(*ApprovePromotionResponse)(nil), // 5: pipelines.v1.ApprovePromotionResponse
	(*ListError)(nil),                // 6: pipelines.v1.ListError
	(*ListPullRequestsRequest)(nil),  // 7: pipelines.v1.ListPullRequestsRequest
	(*ListPullRequestsResponse)(nil), // 8: pipelines.v1.ListPullRequestsResponse
	nil,                              // 9: pipelines.v1.ListPullRequestsResponse.PullRequestsEntry
	(*Pipeline)(nil),                 // 10: pipelines.v1.Pipeline
}
var file_api_pipelines_pipelines_proto_depIdxs = []int32{
	10, // 0: pipelines.v1.ListPipelinesResponse.pipelines:type_name -> pipelines.v1.Pipeline
	6,  // 1: pipelines.v1.ListPipelinesResponse.errors:type_name -> pipelines.v1.ListError
	10, // 2: pipelines.v1.GetPipelineResponse.pipeline:type_name -> pipelines.v1.Pipeline
	9,  // 3: pipelines.v1.ListPullRequestsResponse.pullRequests:type_name -> pipelines.v1.ListPullRequestsResponse.PullRequestsEntry
	0,  // 4: pipelines.v1.Pipelines.ListPipelines:input_type -> pipelines.v1.ListPipelinesRequest
	2,  // 5: pipelines.v1.Pipelines.GetPipeline:input_type -> pipelines.v1.GetPipelineRequest
	4,  // 6: pipelines.v1.Pipelines.ApprovePromotion:input_type -> pipelines.v1.ApprovePromotionRequest
	7,  // 7: pipelines.v1.Pipelines.ListPullRequests:input_type -> pipelines.v1.ListPullRequestsRequest
	1,  // 8: pipelines.v1.Pipelines.ListPipelines:output_type -> pipelines.v1.ListPipelinesResponse
	3,  // 9: pipelines.v1.Pipelines.GetPipeline:output_type -> pipelines.v1.GetPipelineResponse
	5,  // 10: pipelines.v1.Pipelines.ApprovePromotion:output_type -> pipelines.v1.ApprovePromotionResponse
	8,  // 11: pipelines.v1.Pipelines.ListPullRequests:output_type -> pipelines.v1.ListPullRequestsResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_pipelines_pipelines_proto_init() }
func file_api_pipelines_pipelines_proto_init() {
	if File_api_pipelines_pipelines_proto != nil {
		return
	}
	file_api_pipelines_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_pipelines_pipelines_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPipelinesRequest); i {
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
		file_api_pipelines_pipelines_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPipelinesResponse); i {
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
		file_api_pipelines_pipelines_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPipelineRequest); i {
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
		file_api_pipelines_pipelines_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPipelineResponse); i {
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
		file_api_pipelines_pipelines_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovePromotionRequest); i {
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
		file_api_pipelines_pipelines_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovePromotionResponse); i {
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
		file_api_pipelines_pipelines_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListError); i {
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
		file_api_pipelines_pipelines_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPullRequestsRequest); i {
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
		file_api_pipelines_pipelines_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPullRequestsResponse); i {
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
			RawDescriptor: file_api_pipelines_pipelines_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_pipelines_pipelines_proto_goTypes,
		DependencyIndexes: file_api_pipelines_pipelines_proto_depIdxs,
		MessageInfos:      file_api_pipelines_pipelines_proto_msgTypes,
	}.Build()
	File_api_pipelines_pipelines_proto = out.File
	file_api_pipelines_pipelines_proto_rawDesc = nil
	file_api_pipelines_pipelines_proto_goTypes = nil
	file_api_pipelines_pipelines_proto_depIdxs = nil
}
