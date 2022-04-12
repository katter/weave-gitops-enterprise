// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package profiles

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProfilesClient is the client API for Profiles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfilesClient interface {
	// GetProfiles returns a list of profiles from the cluster.
	GetProfiles(ctx context.Context, in *GetProfilesRequest, opts ...grpc.CallOption) (*GetProfilesResponse, error)
	// GetProfileValues returns a list of values for a given version of a profile from the cluster.
	GetProfileValues(ctx context.Context, in *GetProfileValuesRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
}

type profilesClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilesClient(cc grpc.ClientConnInterface) ProfilesClient {
	return &profilesClient{cc}
}

func (c *profilesClient) GetProfiles(ctx context.Context, in *GetProfilesRequest, opts ...grpc.CallOption) (*GetProfilesResponse, error) {
	out := new(GetProfilesResponse)
	err := c.cc.Invoke(ctx, "/weave_gitops_profiles.v1.Profiles/GetProfiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) GetProfileValues(ctx context.Context, in *GetProfileValuesRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/weave_gitops_profiles.v1.Profiles/GetProfileValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilesServer is the server API for Profiles service.
// All implementations must embed UnimplementedProfilesServer
// for forward compatibility
type ProfilesServer interface {
	// GetProfiles returns a list of profiles from the cluster.
	GetProfiles(context.Context, *GetProfilesRequest) (*GetProfilesResponse, error)
	// GetProfileValues returns a list of values for a given version of a profile from the cluster.
	GetProfileValues(context.Context, *GetProfileValuesRequest) (*httpbody.HttpBody, error)
	mustEmbedUnimplementedProfilesServer()
}

// UnimplementedProfilesServer must be embedded to have forward compatible implementations.
type UnimplementedProfilesServer struct {
}

func (UnimplementedProfilesServer) GetProfiles(context.Context, *GetProfilesRequest) (*GetProfilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfiles not implemented")
}
func (UnimplementedProfilesServer) GetProfileValues(context.Context, *GetProfileValuesRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileValues not implemented")
}
func (UnimplementedProfilesServer) mustEmbedUnimplementedProfilesServer() {}

// UnsafeProfilesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfilesServer will
// result in compilation errors.
type UnsafeProfilesServer interface {
	mustEmbedUnimplementedProfilesServer()
}

func RegisterProfilesServer(s grpc.ServiceRegistrar, srv ProfilesServer) {
	s.RegisterService(&Profiles_ServiceDesc, srv)
}

func _Profiles_GetProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).GetProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weave_gitops_profiles.v1.Profiles/GetProfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).GetProfiles(ctx, req.(*GetProfilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profiles_GetProfileValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).GetProfileValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weave_gitops_profiles.v1.Profiles/GetProfileValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).GetProfileValues(ctx, req.(*GetProfileValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Profiles_ServiceDesc is the grpc.ServiceDesc for Profiles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profiles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "weave_gitops_profiles.v1.Profiles",
	HandlerType: (*ProfilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfiles",
			Handler:    _Profiles_GetProfiles_Handler,
		},
		{
			MethodName: "GetProfileValues",
			Handler:    _Profiles_GetProfileValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profiles/profiles.proto",
}
