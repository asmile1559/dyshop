// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: order_page.proto

package order_page

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
	OrderService_ListOrders_FullMethodName    = "/order_page.OrderService/ListOrders"
	OrderService_PlaceOrder_FullMethodName    = "/order_page.OrderService/PlaceOrder"
	OrderService_GetOrder_FullMethodName      = "/order_page.OrderService/GetOrder"
	OrderService_ModifyOrder_FullMethodName   = "/order_page.OrderService/ModifyOrder"
	OrderService_CancelOrder_FullMethodName   = "/order_page.OrderService/CancelOrder"
	OrderService_MarkOrderPaid_FullMethodName = "/order_page.OrderService/MarkOrderPaid"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	ListOrders(ctx context.Context, in *ListOrdersReq, opts ...grpc.CallOption) (*ListOrdersResp, error)
	PlaceOrder(ctx context.Context, in *PlaceOrderReq, opts ...grpc.CallOption) (*PlaceOrderResp, error)
	GetOrder(ctx context.Context, in *GetOrderReq, opts ...grpc.CallOption) (*GetOrderResp, error)
	ModifyOrder(ctx context.Context, in *ModifyOrderReq, opts ...grpc.CallOption) (*ModifyOrderResp, error)
	CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...grpc.CallOption) (*CancelOrderResp, error)
	MarkOrderPaid(ctx context.Context, in *MarkOrderPaidReq, opts ...grpc.CallOption) (*MarkOrderPaidResp, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) ListOrders(ctx context.Context, in *ListOrdersReq, opts ...grpc.CallOption) (*ListOrdersResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOrdersResp)
	err := c.cc.Invoke(ctx, OrderService_ListOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) PlaceOrder(ctx context.Context, in *PlaceOrderReq, opts ...grpc.CallOption) (*PlaceOrderResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlaceOrderResp)
	err := c.cc.Invoke(ctx, OrderService_PlaceOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *GetOrderReq, opts ...grpc.CallOption) (*GetOrderResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderResp)
	err := c.cc.Invoke(ctx, OrderService_GetOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ModifyOrder(ctx context.Context, in *ModifyOrderReq, opts ...grpc.CallOption) (*ModifyOrderResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ModifyOrderResp)
	err := c.cc.Invoke(ctx, OrderService_ModifyOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...grpc.CallOption) (*CancelOrderResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelOrderResp)
	err := c.cc.Invoke(ctx, OrderService_CancelOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) MarkOrderPaid(ctx context.Context, in *MarkOrderPaidReq, opts ...grpc.CallOption) (*MarkOrderPaidResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarkOrderPaidResp)
	err := c.cc.Invoke(ctx, OrderService_MarkOrderPaid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility.
type OrderServiceServer interface {
	ListOrders(context.Context, *ListOrdersReq) (*ListOrdersResp, error)
	PlaceOrder(context.Context, *PlaceOrderReq) (*PlaceOrderResp, error)
	GetOrder(context.Context, *GetOrderReq) (*GetOrderResp, error)
	ModifyOrder(context.Context, *ModifyOrderReq) (*ModifyOrderResp, error)
	CancelOrder(context.Context, *CancelOrderReq) (*CancelOrderResp, error)
	MarkOrderPaid(context.Context, *MarkOrderPaidReq) (*MarkOrderPaidResp, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderServiceServer struct{}

func (UnimplementedOrderServiceServer) ListOrders(context.Context, *ListOrdersReq) (*ListOrdersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedOrderServiceServer) PlaceOrder(context.Context, *PlaceOrderReq) (*PlaceOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrder(context.Context, *GetOrderReq) (*GetOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServiceServer) ModifyOrder(context.Context, *ModifyOrderReq) (*ModifyOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyOrder not implemented")
}
func (UnimplementedOrderServiceServer) CancelOrder(context.Context, *CancelOrderReq) (*CancelOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrderServiceServer) MarkOrderPaid(context.Context, *MarkOrderPaidReq) (*MarkOrderPaidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkOrderPaid not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}
func (UnimplementedOrderServiceServer) testEmbeddedByValue()                      {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ListOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOrders(ctx, req.(*ListOrdersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_PlaceOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PlaceOrder(ctx, req.(*PlaceOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*GetOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ModifyOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ModifyOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ModifyOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ModifyOrder(ctx, req.(*ModifyOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CancelOrder(ctx, req.(*CancelOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_MarkOrderPaid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkOrderPaidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).MarkOrderPaid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_MarkOrderPaid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).MarkOrderPaid(ctx, req.(*MarkOrderPaidReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order_page.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOrders",
			Handler:    _OrderService_ListOrders_Handler,
		},
		{
			MethodName: "PlaceOrder",
			Handler:    _OrderService_PlaceOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
		{
			MethodName: "ModifyOrder",
			Handler:    _OrderService_ModifyOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrderService_CancelOrder_Handler,
		},
		{
			MethodName: "MarkOrderPaid",
			Handler:    _OrderService_MarkOrderPaid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order_page.proto",
}
