// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetadataServiceClient is the client API for MetadataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetadataServiceClient interface {
	GetConnectionInformation(ctx context.Context, in *EmptyParams, opts ...grpc.CallOption) (*ConnectionInformation, error)
	RegisterObjectStore(ctx context.Context, in *ObjectStoreConfig, opts ...grpc.CallOption) (*StatusResponse, error)
	ListObjectStores(ctx context.Context, in *EmptyParams, opts ...grpc.CallOption) (*AvailableObjectStoreConfigs, error)
	CreateBucket(ctx context.Context, in *Bucket, opts ...grpc.CallOption) (*StatusResponse, error)
	DeleteBucket(ctx context.Context, in *Bucket, opts ...grpc.CallOption) (*StatusResponse, error)
	ListBuckets(ctx context.Context, in *EmptyParams, opts ...grpc.CallOption) (*BucketListResponse, error)
	LookupBucket(ctx context.Context, in *Bucket, opts ...grpc.CallOption) (*StatusResponse, error)
	Create(ctx context.Context, in *Object, opts ...grpc.CallOption) (*StatusResponse, error)
	Update(ctx context.Context, in *Object, opts ...grpc.CallOption) (*StatusResponse, error)
	Delete(ctx context.Context, in *ObjectID, opts ...grpc.CallOption) (*StatusResponse, error)
	DeletePrefix(ctx context.Context, in *ObjectID, opts ...grpc.CallOption) (*StatusResponse, error)
	Lookup(ctx context.Context, in *ObjectID, opts ...grpc.CallOption) (*ObjectResponse, error)
	List(ctx context.Context, in *ObjectListRequest, opts ...grpc.CallOption) (*ObjectListResponse, error)
	TestRPC(ctx context.Context, in *ConnectionInformation, opts ...grpc.CallOption) (*ConnectionInformation, error)
	SubscribeObjects(ctx context.Context, in *ObjectEventSubscription, opts ...grpc.CallOption) (MetadataService_SubscribeObjectsClient, error)
}

type metadataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataServiceClient(cc grpc.ClientConnInterface) MetadataServiceClient {
	return &metadataServiceClient{cc}
}

func (c *metadataServiceClient) GetConnectionInformation(ctx context.Context, in *EmptyParams, opts ...grpc.CallOption) (*ConnectionInformation, error) {
	out := new(ConnectionInformation)
	err := c.cc.Invoke(ctx, "/geds.rpc.MetadataService/GetConnectionInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) RegisterObjectStore(ctx context.Context, in *ObjectStoreConfig, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/geds.rpc.MetadataService/RegisterObjectStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) ListObjectStores(ctx context.Context, in *EmptyParams, opts ...grpc.CallOption) (*AvailableObjectStoreConfigs, error) {
	out := new(AvailableObjectStoreConfigs)
	err := c.cc.Invoke(ctx, "/geds.rpc.MetadataService/ListObjectStores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) CreateBucket(ctx context.Context, in *Bucket, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/geds.rpc.MetadataService/CreateBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) DeleteBucket(ctx context.Context, in *Bucket, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/geds.rpc.MetadataService/DeleteBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) ListBuckets(ctx context.Context, in *EmptyParams, opts ...grpc.CallOption) (*BucketListResponse, error) {
	out := new(BucketListResponse)
	err := c.cc.Invoke(ctx, "/geds.rpc.MetadataService/ListBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) LookupBucket(ctx context.Context, in *Bucket, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/geds.rpc.MetadataService/LookupBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) Create(ctx context.Context, in *Object, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/geds.rpc.MetadataService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) Update(ctx context.Context, in *Object, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/geds.rpc.MetadataService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) Delete(ctx context.Context, in *ObjectID, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/geds.rpc.MetadataService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) DeletePrefix(ctx context.Context, in *ObjectID, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/geds.rpc.MetadataService/DeletePrefix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) Lookup(ctx context.Context, in *ObjectID, opts ...grpc.CallOption) (*ObjectResponse, error) {
	out := new(ObjectResponse)
	err := c.cc.Invoke(ctx, "/geds.rpc.MetadataService/Lookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) List(ctx context.Context, in *ObjectListRequest, opts ...grpc.CallOption) (*ObjectListResponse, error) {
	out := new(ObjectListResponse)
	err := c.cc.Invoke(ctx, "/geds.rpc.MetadataService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) TestRPC(ctx context.Context, in *ConnectionInformation, opts ...grpc.CallOption) (*ConnectionInformation, error) {
	out := new(ConnectionInformation)
	err := c.cc.Invoke(ctx, "/geds.rpc.MetadataService/TestRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) SubscribeObjects(ctx context.Context, in *ObjectEventSubscription, opts ...grpc.CallOption) (MetadataService_SubscribeObjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MetadataService_ServiceDesc.Streams[0], "/geds.rpc.MetadataService/SubscribeObjects", opts...)
	if err != nil {
		return nil, err
	}
	x := &metadataServiceSubscribeObjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MetadataService_SubscribeObjectsClient interface {
	Recv() (*ObjectSubscription, error)
	grpc.ClientStream
}

type metadataServiceSubscribeObjectsClient struct {
	grpc.ClientStream
}

func (x *metadataServiceSubscribeObjectsClient) Recv() (*ObjectSubscription, error) {
	m := new(ObjectSubscription)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetadataServiceServer is the server API for MetadataService service.
// All implementations should embed UnimplementedMetadataServiceServer
// for forward compatibility
type MetadataServiceServer interface {
	GetConnectionInformation(context.Context, *EmptyParams) (*ConnectionInformation, error)
	RegisterObjectStore(context.Context, *ObjectStoreConfig) (*StatusResponse, error)
	ListObjectStores(context.Context, *EmptyParams) (*AvailableObjectStoreConfigs, error)
	CreateBucket(context.Context, *Bucket) (*StatusResponse, error)
	DeleteBucket(context.Context, *Bucket) (*StatusResponse, error)
	ListBuckets(context.Context, *EmptyParams) (*BucketListResponse, error)
	LookupBucket(context.Context, *Bucket) (*StatusResponse, error)
	Create(context.Context, *Object) (*StatusResponse, error)
	Update(context.Context, *Object) (*StatusResponse, error)
	Delete(context.Context, *ObjectID) (*StatusResponse, error)
	DeletePrefix(context.Context, *ObjectID) (*StatusResponse, error)
	Lookup(context.Context, *ObjectID) (*ObjectResponse, error)
	List(context.Context, *ObjectListRequest) (*ObjectListResponse, error)
	TestRPC(context.Context, *ConnectionInformation) (*ConnectionInformation, error)
	SubscribeObjects(*ObjectEventSubscription, MetadataService_SubscribeObjectsServer) error
}

// UnimplementedMetadataServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMetadataServiceServer struct {
}

func (UnimplementedMetadataServiceServer) GetConnectionInformation(context.Context, *EmptyParams) (*ConnectionInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectionInformation not implemented")
}
func (UnimplementedMetadataServiceServer) RegisterObjectStore(context.Context, *ObjectStoreConfig) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterObjectStore not implemented")
}
func (UnimplementedMetadataServiceServer) ListObjectStores(context.Context, *EmptyParams) (*AvailableObjectStoreConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObjectStores not implemented")
}
func (UnimplementedMetadataServiceServer) CreateBucket(context.Context, *Bucket) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBucket not implemented")
}
func (UnimplementedMetadataServiceServer) DeleteBucket(context.Context, *Bucket) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBucket not implemented")
}
func (UnimplementedMetadataServiceServer) ListBuckets(context.Context, *EmptyParams) (*BucketListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuckets not implemented")
}
func (UnimplementedMetadataServiceServer) LookupBucket(context.Context, *Bucket) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupBucket not implemented")
}
func (UnimplementedMetadataServiceServer) Create(context.Context, *Object) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMetadataServiceServer) Update(context.Context, *Object) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMetadataServiceServer) Delete(context.Context, *ObjectID) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMetadataServiceServer) DeletePrefix(context.Context, *ObjectID) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePrefix not implemented")
}
func (UnimplementedMetadataServiceServer) Lookup(context.Context, *ObjectID) (*ObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lookup not implemented")
}
func (UnimplementedMetadataServiceServer) List(context.Context, *ObjectListRequest) (*ObjectListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMetadataServiceServer) TestRPC(context.Context, *ConnectionInformation) (*ConnectionInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestRPC not implemented")
}
func (UnimplementedMetadataServiceServer) SubscribeObjects(*ObjectEventSubscription, MetadataService_SubscribeObjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeObjects not implemented")
}

// UnsafeMetadataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetadataServiceServer will
// result in compilation errors.
type UnsafeMetadataServiceServer interface {
	mustEmbedUnimplementedMetadataServiceServer()
}

func RegisterMetadataServiceServer(s grpc.ServiceRegistrar, srv MetadataServiceServer) {
	s.RegisterService(&MetadataService_ServiceDesc, srv)
}

func _MetadataService_GetConnectionInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetConnectionInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.MetadataService/GetConnectionInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetConnectionInformation(ctx, req.(*EmptyParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_RegisterObjectStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectStoreConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).RegisterObjectStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.MetadataService/RegisterObjectStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).RegisterObjectStore(ctx, req.(*ObjectStoreConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_ListObjectStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).ListObjectStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.MetadataService/ListObjectStores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).ListObjectStores(ctx, req.(*EmptyParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_CreateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).CreateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.MetadataService/CreateBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).CreateBucket(ctx, req.(*Bucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_DeleteBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).DeleteBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.MetadataService/DeleteBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).DeleteBucket(ctx, req.(*Bucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_ListBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).ListBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.MetadataService/ListBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).ListBuckets(ctx, req.(*EmptyParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_LookupBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).LookupBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.MetadataService/LookupBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).LookupBucket(ctx, req.(*Bucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.MetadataService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).Create(ctx, req.(*Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.MetadataService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).Update(ctx, req.(*Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.MetadataService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).Delete(ctx, req.(*ObjectID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_DeletePrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).DeletePrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.MetadataService/DeletePrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).DeletePrefix(ctx, req.(*ObjectID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.MetadataService/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).Lookup(ctx, req.(*ObjectID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.MetadataService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).List(ctx, req.(*ObjectListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_TestRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionInformation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).TestRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.MetadataService/TestRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).TestRPC(ctx, req.(*ConnectionInformation))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_SubscribeObjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ObjectEventSubscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetadataServiceServer).SubscribeObjects(m, &metadataServiceSubscribeObjectsServer{stream})
}

type MetadataService_SubscribeObjectsServer interface {
	Send(*ObjectSubscription) error
	grpc.ServerStream
}

type metadataServiceSubscribeObjectsServer struct {
	grpc.ServerStream
}

func (x *metadataServiceSubscribeObjectsServer) Send(m *ObjectSubscription) error {
	return x.ServerStream.SendMsg(m)
}

// MetadataService_ServiceDesc is the grpc.ServiceDesc for MetadataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetadataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "geds.rpc.MetadataService",
	HandlerType: (*MetadataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConnectionInformation",
			Handler:    _MetadataService_GetConnectionInformation_Handler,
		},
		{
			MethodName: "RegisterObjectStore",
			Handler:    _MetadataService_RegisterObjectStore_Handler,
		},
		{
			MethodName: "ListObjectStores",
			Handler:    _MetadataService_ListObjectStores_Handler,
		},
		{
			MethodName: "CreateBucket",
			Handler:    _MetadataService_CreateBucket_Handler,
		},
		{
			MethodName: "DeleteBucket",
			Handler:    _MetadataService_DeleteBucket_Handler,
		},
		{
			MethodName: "ListBuckets",
			Handler:    _MetadataService_ListBuckets_Handler,
		},
		{
			MethodName: "LookupBucket",
			Handler:    _MetadataService_LookupBucket_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _MetadataService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MetadataService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MetadataService_Delete_Handler,
		},
		{
			MethodName: "DeletePrefix",
			Handler:    _MetadataService_DeletePrefix_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _MetadataService_Lookup_Handler,
		},
		{
			MethodName: "List",
			Handler:    _MetadataService_List_Handler,
		},
		{
			MethodName: "TestRPC",
			Handler:    _MetadataService_TestRPC_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeObjects",
			Handler:       _MetadataService_SubscribeObjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "geds.proto",
}

// GEDSServiceClient is the client API for GEDSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GEDSServiceClient interface {
	GetAvailEndpoints(ctx context.Context, in *EmptyParams, opts ...grpc.CallOption) (*AvailTransportEndpoints, error)
}

type gEDSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGEDSServiceClient(cc grpc.ClientConnInterface) GEDSServiceClient {
	return &gEDSServiceClient{cc}
}

func (c *gEDSServiceClient) GetAvailEndpoints(ctx context.Context, in *EmptyParams, opts ...grpc.CallOption) (*AvailTransportEndpoints, error) {
	out := new(AvailTransportEndpoints)
	err := c.cc.Invoke(ctx, "/geds.rpc.GEDSService/GetAvailEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GEDSServiceServer is the server API for GEDSService service.
// All implementations should embed UnimplementedGEDSServiceServer
// for forward compatibility
type GEDSServiceServer interface {
	GetAvailEndpoints(context.Context, *EmptyParams) (*AvailTransportEndpoints, error)
}

// UnimplementedGEDSServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGEDSServiceServer struct {
}

func (UnimplementedGEDSServiceServer) GetAvailEndpoints(context.Context, *EmptyParams) (*AvailTransportEndpoints, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailEndpoints not implemented")
}

// UnsafeGEDSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GEDSServiceServer will
// result in compilation errors.
type UnsafeGEDSServiceServer interface {
	mustEmbedUnimplementedGEDSServiceServer()
}

func RegisterGEDSServiceServer(s grpc.ServiceRegistrar, srv GEDSServiceServer) {
	s.RegisterService(&GEDSService_ServiceDesc, srv)
}

func _GEDSService_GetAvailEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GEDSServiceServer).GetAvailEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geds.rpc.GEDSService/GetAvailEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GEDSServiceServer).GetAvailEndpoints(ctx, req.(*EmptyParams))
	}
	return interceptor(ctx, in, info, handler)
}

// GEDSService_ServiceDesc is the grpc.ServiceDesc for GEDSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GEDSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "geds.rpc.GEDSService",
	HandlerType: (*GEDSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvailEndpoints",
			Handler:    _GEDSService_GetAvailEndpoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geds.proto",
}
