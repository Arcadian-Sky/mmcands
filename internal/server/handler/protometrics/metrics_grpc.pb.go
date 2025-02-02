//protoc -I $GOPATH/pkg/mod/github.com/googleapis/googleapis@latest/ -I . --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative service.proto
//protoc -I proto/googleapi/ -I proto/. --grpc-gateway_out=. --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative metrics.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: metrics.proto

package protometrics

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
	MetricsService_GetMetric_FullMethodName         = "/metrics.MetricsService/GetMetric"
	MetricsService_GetMetrics_FullMethodName        = "/metrics.MetricsService/GetMetrics"
	MetricsService_GetMetricsRoot_FullMethodName    = "/metrics.MetricsService/GetMetricsRoot"
	MetricsService_GetJSONMetrics_FullMethodName    = "/metrics.MetricsService/GetJSONMetrics"
	MetricsService_UpdateJSONMetrics_FullMethodName = "/metrics.MetricsService/UpdateJSONMetrics"
	MetricsService_UpdateJSONMetric_FullMethodName  = "/metrics.MetricsService/UpdateJSONMetric"
	MetricsService_UpdateMetric_FullMethodName      = "/metrics.MetricsService/UpdateMetric"
	MetricsService_PingDB_FullMethodName            = "/metrics.MetricsService/PingDB"
)

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsServiceClient interface {
	// Получает метрику
	GetMetric(ctx context.Context, in *Metric, opts ...grpc.CallOption) (*GetMetricResponse, error)
	// Получает метрики
	GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
	GetMetricsRoot(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
	// Получает метрики
	GetJSONMetrics(ctx context.Context, in *Metric, opts ...grpc.CallOption) (*GetJSONMetricsResponse, error)
	// Обновляет метрики через JSON
	UpdateJSONMetrics(ctx context.Context, in *UpdateJSONMetricsRequest, opts ...grpc.CallOption) (*UpdateJSONMetricsResponse, error)
	// Обновляет метрику (для маршрута /update/)
	UpdateJSONMetric(ctx context.Context, in *Metric, opts ...grpc.CallOption) (*Metric, error)
	// Обновляет метрику
	UpdateMetric(ctx context.Context, in *UpdateMetricRequest, opts ...grpc.CallOption) (*UpdateMetricResponse, error)
	// Проверяет соединение с базой данных
	PingDB(ctx context.Context, in *PingDBRequest, opts ...grpc.CallOption) (*PingDBResponse, error)
}

type metricsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsServiceClient(cc grpc.ClientConnInterface) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) GetMetric(ctx context.Context, in *Metric, opts ...grpc.CallOption) (*GetMetricResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMetricResponse)
	err := c.cc.Invoke(ctx, MetricsService_GetMetric_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, MetricsService_GetMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) GetMetricsRoot(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, MetricsService_GetMetricsRoot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) GetJSONMetrics(ctx context.Context, in *Metric, opts ...grpc.CallOption) (*GetJSONMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetJSONMetricsResponse)
	err := c.cc.Invoke(ctx, MetricsService_GetJSONMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) UpdateJSONMetrics(ctx context.Context, in *UpdateJSONMetricsRequest, opts ...grpc.CallOption) (*UpdateJSONMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateJSONMetricsResponse)
	err := c.cc.Invoke(ctx, MetricsService_UpdateJSONMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) UpdateJSONMetric(ctx context.Context, in *Metric, opts ...grpc.CallOption) (*Metric, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Metric)
	err := c.cc.Invoke(ctx, MetricsService_UpdateJSONMetric_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) UpdateMetric(ctx context.Context, in *UpdateMetricRequest, opts ...grpc.CallOption) (*UpdateMetricResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMetricResponse)
	err := c.cc.Invoke(ctx, MetricsService_UpdateMetric_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) PingDB(ctx context.Context, in *PingDBRequest, opts ...grpc.CallOption) (*PingDBResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingDBResponse)
	err := c.cc.Invoke(ctx, MetricsService_PingDB_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
// All implementations must embed UnimplementedMetricsServiceServer
// for forward compatibility.
type MetricsServiceServer interface {
	// Получает метрику
	GetMetric(context.Context, *Metric) (*GetMetricResponse, error)
	// Получает метрики
	GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
	GetMetricsRoot(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
	// Получает метрики
	GetJSONMetrics(context.Context, *Metric) (*GetJSONMetricsResponse, error)
	// Обновляет метрики через JSON
	UpdateJSONMetrics(context.Context, *UpdateJSONMetricsRequest) (*UpdateJSONMetricsResponse, error)
	// Обновляет метрику (для маршрута /update/)
	UpdateJSONMetric(context.Context, *Metric) (*Metric, error)
	// Обновляет метрику
	UpdateMetric(context.Context, *UpdateMetricRequest) (*UpdateMetricResponse, error)
	// Проверяет соединение с базой данных
	PingDB(context.Context, *PingDBRequest) (*PingDBResponse, error)
	mustEmbedUnimplementedMetricsServiceServer()
}

// UnimplementedMetricsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMetricsServiceServer struct{}

func (UnimplementedMetricsServiceServer) GetMetric(context.Context, *Metric) (*GetMetricResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetric not implemented")
}
func (UnimplementedMetricsServiceServer) GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}
func (UnimplementedMetricsServiceServer) GetMetricsRoot(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricsRoot not implemented")
}
func (UnimplementedMetricsServiceServer) GetJSONMetrics(context.Context, *Metric) (*GetJSONMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJSONMetrics not implemented")
}
func (UnimplementedMetricsServiceServer) UpdateJSONMetrics(context.Context, *UpdateJSONMetricsRequest) (*UpdateJSONMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJSONMetrics not implemented")
}
func (UnimplementedMetricsServiceServer) UpdateJSONMetric(context.Context, *Metric) (*Metric, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJSONMetric not implemented")
}
func (UnimplementedMetricsServiceServer) UpdateMetric(context.Context, *UpdateMetricRequest) (*UpdateMetricResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMetric not implemented")
}
func (UnimplementedMetricsServiceServer) PingDB(context.Context, *PingDBRequest) (*PingDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingDB not implemented")
}
func (UnimplementedMetricsServiceServer) mustEmbedUnimplementedMetricsServiceServer() {}
func (UnimplementedMetricsServiceServer) testEmbeddedByValue()                        {}

// UnsafeMetricsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServiceServer will
// result in compilation errors.
type UnsafeMetricsServiceServer interface {
	mustEmbedUnimplementedMetricsServiceServer()
}

func RegisterMetricsServiceServer(s grpc.ServiceRegistrar, srv MetricsServiceServer) {
	// If the following call pancis, it indicates UnimplementedMetricsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MetricsService_ServiceDesc, srv)
}

func _MetricsService_GetMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Metric)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_GetMetric_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetMetric(ctx, req.(*Metric))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_GetMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetMetrics(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_GetMetricsRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetMetricsRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_GetMetricsRoot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetMetricsRoot(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_GetJSONMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Metric)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetJSONMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_GetJSONMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetJSONMetrics(ctx, req.(*Metric))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_UpdateJSONMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJSONMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).UpdateJSONMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_UpdateJSONMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).UpdateJSONMetrics(ctx, req.(*UpdateJSONMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_UpdateJSONMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Metric)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).UpdateJSONMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_UpdateJSONMetric_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).UpdateJSONMetric(ctx, req.(*Metric))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_UpdateMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).UpdateMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_UpdateMetric_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).UpdateMetric(ctx, req.(*UpdateMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_PingDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).PingDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_PingDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).PingDB(ctx, req.(*PingDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricsService_ServiceDesc is the grpc.ServiceDesc for MetricsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metrics.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetric",
			Handler:    _MetricsService_GetMetric_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _MetricsService_GetMetrics_Handler,
		},
		{
			MethodName: "GetMetricsRoot",
			Handler:    _MetricsService_GetMetricsRoot_Handler,
		},
		{
			MethodName: "GetJSONMetrics",
			Handler:    _MetricsService_GetJSONMetrics_Handler,
		},
		{
			MethodName: "UpdateJSONMetrics",
			Handler:    _MetricsService_UpdateJSONMetrics_Handler,
		},
		{
			MethodName: "UpdateJSONMetric",
			Handler:    _MetricsService_UpdateJSONMetric_Handler,
		},
		{
			MethodName: "UpdateMetric",
			Handler:    _MetricsService_UpdateMetric_Handler,
		},
		{
			MethodName: "PingDB",
			Handler:    _MetricsService_PingDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metrics.proto",
}
