// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: product.proto

package product

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProductCatalogService_ListProducts_FullMethodName   = "/product.ProductCatalogService/ListProducts"
	ProductCatalogService_GetProduct_FullMethodName     = "/product.ProductCatalogService/GetProduct"
	ProductCatalogService_SearchProducts_FullMethodName = "/product.ProductCatalogService/SearchProducts"
	ProductCatalogService_CreateProduct_FullMethodName  = "/product.ProductCatalogService/CreateProduct"
	ProductCatalogService_ModifyProduct_FullMethodName  = "/product.ProductCatalogService/ModifyProduct"
	ProductCatalogService_DeleteProduct_FullMethodName  = "/product.ProductCatalogService/DeleteProduct"
)

// ProductCatalogServiceClient is the client API for ProductCatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductCatalogServiceClient interface {
	ListProducts(ctx context.Context, in *ListProductsReq, opts ...grpc.CallOption) (*ListProductsResp, error)
	GetProduct(ctx context.Context, in *GetProductReq, opts ...grpc.CallOption) (*GetProductResp, error)
	SearchProducts(ctx context.Context, in *SearchProductsReq, opts ...grpc.CallOption) (*SearchProductsResp, error)
	CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error)
	ModifyProduct(ctx context.Context, in *ModifyProductReq, opts ...grpc.CallOption) (*ModifyProductResp, error)
	DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductResp, error)
}

type productCatalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCatalogServiceClient(cc grpc.ClientConnInterface) ProductCatalogServiceClient {
	return &productCatalogServiceClient{cc}
}

func (c *productCatalogServiceClient) ListProducts(ctx context.Context, in *ListProductsReq, opts ...grpc.CallOption) (*ListProductsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProductsResp)
	err := c.cc.Invoke(ctx, ProductCatalogService_ListProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) GetProduct(ctx context.Context, in *GetProductReq, opts ...grpc.CallOption) (*GetProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductResp)
	err := c.cc.Invoke(ctx, ProductCatalogService_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) SearchProducts(ctx context.Context, in *SearchProductsReq, opts ...grpc.CallOption) (*SearchProductsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchProductsResp)
	err := c.cc.Invoke(ctx, ProductCatalogService_SearchProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProductResp)
	err := c.cc.Invoke(ctx, ProductCatalogService_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) ModifyProduct(ctx context.Context, in *ModifyProductReq, opts ...grpc.CallOption) (*ModifyProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ModifyProductResp)
	err := c.cc.Invoke(ctx, ProductCatalogService_ModifyProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProductResp)
	err := c.cc.Invoke(ctx, ProductCatalogService_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCatalogServiceServer is the server API for ProductCatalogService service.
// All implementations must embed UnimplementedProductCatalogServiceServer
// for forward compatibility.
type ProductCatalogServiceServer interface {
	ListProducts(context.Context, *ListProductsReq) (*ListProductsResp, error)
	GetProduct(context.Context, *GetProductReq) (*GetProductResp, error)
	SearchProducts(context.Context, *SearchProductsReq) (*SearchProductsResp, error)
	CreateProduct(context.Context, *CreateProductReq) (*CreateProductResp, error)
	ModifyProduct(context.Context, *ModifyProductReq) (*ModifyProductResp, error)
	DeleteProduct(context.Context, *DeleteProductReq) (*DeleteProductResp, error)
	mustEmbedUnimplementedProductCatalogServiceServer()
}

// UnimplementedProductCatalogServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductCatalogServiceServer struct{}

func (UnimplementedProductCatalogServiceServer) ListProducts(context.Context, *ListProductsReq) (*ListProductsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedProductCatalogServiceServer) GetProduct(context.Context, *GetProductReq) (*GetProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductCatalogServiceServer) SearchProducts(context.Context, *SearchProductsReq) (*SearchProductsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}
func (UnimplementedProductCatalogServiceServer) CreateProduct(context.Context, *CreateProductReq) (*CreateProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductCatalogServiceServer) ModifyProduct(context.Context, *ModifyProductReq) (*ModifyProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyProduct not implemented")
}
func (UnimplementedProductCatalogServiceServer) DeleteProduct(context.Context, *DeleteProductReq) (*DeleteProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductCatalogServiceServer) mustEmbedUnimplementedProductCatalogServiceServer() {}
func (UnimplementedProductCatalogServiceServer) testEmbeddedByValue()                               {}

// UnsafeProductCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCatalogServiceServer will
// result in compilation errors.
type UnsafeProductCatalogServiceServer interface {
	mustEmbedUnimplementedProductCatalogServiceServer()
}

func RegisterProductCatalogServiceServer(s grpc.ServiceRegistrar, srv ProductCatalogServiceServer) {
	// If the following call pancis, it indicates UnimplementedProductCatalogServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProductCatalogService_ServiceDesc, srv)
}

func _ProductCatalogService_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_ListProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).ListProducts(ctx, req.(*ListProductsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).GetProduct(ctx, req.(*GetProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_SearchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProductsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).SearchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_SearchProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).SearchProducts(ctx, req.(*SearchProductsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).CreateProduct(ctx, req.(*CreateProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_ModifyProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).ModifyProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_ModifyProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).ModifyProduct(ctx, req.(*ModifyProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).DeleteProduct(ctx, req.(*DeleteProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductCatalogService_ServiceDesc is the grpc.ServiceDesc for ProductCatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductCatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductCatalogService",
	HandlerType: (*ProductCatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProducts",
			Handler:    _ProductCatalogService_ListProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductCatalogService_GetProduct_Handler,
		},
		{
			MethodName: "SearchProducts",
			Handler:    _ProductCatalogService_SearchProducts_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _ProductCatalogService_CreateProduct_Handler,
		},
		{
			MethodName: "ModifyProduct",
			Handler:    _ProductCatalogService_ModifyProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductCatalogService_DeleteProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
