// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/employee.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Employee struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname            string   `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Position             string   `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
	Salary               float32  `protobuf:"fixed32,5,opt,name=salary,proto3" json:"salary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Employee) Reset()         { *m = Employee{} }
func (m *Employee) String() string { return proto.CompactTextString(m) }
func (*Employee) ProtoMessage()    {}
func (*Employee) Descriptor() ([]byte, []int) {
	return fileDescriptor_2884e02695323283, []int{0}
}

func (m *Employee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Employee.Unmarshal(m, b)
}
func (m *Employee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Employee.Marshal(b, m, deterministic)
}
func (m *Employee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Employee.Merge(m, src)
}
func (m *Employee) XXX_Size() int {
	return xxx_messageInfo_Employee.Size(m)
}
func (m *Employee) XXX_DiscardUnknown() {
	xxx_messageInfo_Employee.DiscardUnknown(m)
}

var xxx_messageInfo_Employee proto.InternalMessageInfo

func (m *Employee) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Employee) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *Employee) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *Employee) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Employee) GetSalary() float32 {
	if m != nil {
		return m.Salary
	}
	return 0
}

type EmployeeResponse struct {
	Status               string    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Created              *Employee `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EmployeeResponse) Reset()         { *m = EmployeeResponse{} }
func (m *EmployeeResponse) String() string { return proto.CompactTextString(m) }
func (*EmployeeResponse) ProtoMessage()    {}
func (*EmployeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2884e02695323283, []int{1}
}

func (m *EmployeeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmployeeResponse.Unmarshal(m, b)
}
func (m *EmployeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmployeeResponse.Marshal(b, m, deterministic)
}
func (m *EmployeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmployeeResponse.Merge(m, src)
}
func (m *EmployeeResponse) XXX_Size() int {
	return xxx_messageInfo_EmployeeResponse.Size(m)
}
func (m *EmployeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmployeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmployeeResponse proto.InternalMessageInfo

func (m *EmployeeResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *EmployeeResponse) GetCreated() *Employee {
	if m != nil {
		return m.Created
	}
	return nil
}

type EmployeeRequest struct {
	Empid                string   `protobuf:"bytes,1,opt,name=empid,proto3" json:"empid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmployeeRequest) Reset()         { *m = EmployeeRequest{} }
func (m *EmployeeRequest) String() string { return proto.CompactTextString(m) }
func (*EmployeeRequest) ProtoMessage()    {}
func (*EmployeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2884e02695323283, []int{2}
}

func (m *EmployeeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmployeeRequest.Unmarshal(m, b)
}
func (m *EmployeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmployeeRequest.Marshal(b, m, deterministic)
}
func (m *EmployeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmployeeRequest.Merge(m, src)
}
func (m *EmployeeRequest) XXX_Size() int {
	return xxx_messageInfo_EmployeeRequest.Size(m)
}
func (m *EmployeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmployeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmployeeRequest proto.InternalMessageInfo

func (m *EmployeeRequest) GetEmpid() string {
	if m != nil {
		return m.Empid
	}
	return ""
}

type GetAllEmployes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllEmployes) Reset()         { *m = GetAllEmployes{} }
func (m *GetAllEmployes) String() string { return proto.CompactTextString(m) }
func (*GetAllEmployes) ProtoMessage()    {}
func (*GetAllEmployes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2884e02695323283, []int{3}
}

func (m *GetAllEmployes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllEmployes.Unmarshal(m, b)
}
func (m *GetAllEmployes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllEmployes.Marshal(b, m, deterministic)
}
func (m *GetAllEmployes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllEmployes.Merge(m, src)
}
func (m *GetAllEmployes) XXX_Size() int {
	return xxx_messageInfo_GetAllEmployes.Size(m)
}
func (m *GetAllEmployes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllEmployes.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllEmployes proto.InternalMessageInfo

type DeleteEmloyeReponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEmloyeReponse) Reset()         { *m = DeleteEmloyeReponse{} }
func (m *DeleteEmloyeReponse) String() string { return proto.CompactTextString(m) }
func (*DeleteEmloyeReponse) ProtoMessage()    {}
func (*DeleteEmloyeReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2884e02695323283, []int{4}
}

func (m *DeleteEmloyeReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEmloyeReponse.Unmarshal(m, b)
}
func (m *DeleteEmloyeReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEmloyeReponse.Marshal(b, m, deterministic)
}
func (m *DeleteEmloyeReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEmloyeReponse.Merge(m, src)
}
func (m *DeleteEmloyeReponse) XXX_Size() int {
	return xxx_messageInfo_DeleteEmloyeReponse.Size(m)
}
func (m *DeleteEmloyeReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEmloyeReponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEmloyeReponse proto.InternalMessageInfo

func (m *DeleteEmloyeReponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type UpdateEmployeeRequest struct {
	Emp                  *Employee `protobuf:"bytes,1,opt,name=emp,proto3" json:"emp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateEmployeeRequest) Reset()         { *m = UpdateEmployeeRequest{} }
func (m *UpdateEmployeeRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateEmployeeRequest) ProtoMessage()    {}
func (*UpdateEmployeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2884e02695323283, []int{5}
}

func (m *UpdateEmployeeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEmployeeRequest.Unmarshal(m, b)
}
func (m *UpdateEmployeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEmployeeRequest.Marshal(b, m, deterministic)
}
func (m *UpdateEmployeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEmployeeRequest.Merge(m, src)
}
func (m *UpdateEmployeeRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateEmployeeRequest.Size(m)
}
func (m *UpdateEmployeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEmployeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEmployeeRequest proto.InternalMessageInfo

func (m *UpdateEmployeeRequest) GetEmp() *Employee {
	if m != nil {
		return m.Emp
	}
	return nil
}

func init() {
	proto.RegisterType((*Employee)(nil), "EmployeeRpc.Employee")
	proto.RegisterType((*EmployeeResponse)(nil), "EmployeeRpc.EmployeeResponse")
	proto.RegisterType((*EmployeeRequest)(nil), "EmployeeRpc.EmployeeRequest")
	proto.RegisterType((*GetAllEmployes)(nil), "EmployeeRpc.GetAllEmployes")
	proto.RegisterType((*DeleteEmloyeReponse)(nil), "EmployeeRpc.DeleteEmloyeReponse")
	proto.RegisterType((*UpdateEmployeeRequest)(nil), "EmployeeRpc.UpdateEmployeeRequest")
}

func init() { proto.RegisterFile("pb/employee.proto", fileDescriptor_2884e02695323283) }

var fileDescriptor_2884e02695323283 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x4e, 0xea, 0x50,
	0x10, 0xa5, 0xed, 0x83, 0x07, 0xc3, 0x4b, 0x1f, 0x6f, 0x9e, 0x98, 0x06, 0x31, 0x21, 0x77, 0x03,
	0x1b, 0xc1, 0xe0, 0x0f, 0xa8, 0x88, 0xb8, 0x71, 0x53, 0xe3, 0x46, 0x57, 0x85, 0x8e, 0x49, 0x93,
	0x96, 0x5e, 0x7b, 0x2f, 0x26, 0xfc, 0x81, 0x89, 0x3f, 0x6d, 0x7a, 0xb9, 0x6d, 0x6d, 0xd2, 0xb2,
	0xeb, 0x99, 0x73, 0x7a, 0xe6, 0xcc, 0x4c, 0x2e, 0xfc, 0xe3, 0xeb, 0x19, 0x45, 0x3c, 0x8c, 0xf7,
	0x44, 0x53, 0x9e, 0xc4, 0x32, 0xc6, 0xee, 0x52, 0x63, 0x97, 0x6f, 0xd8, 0xa7, 0x01, 0xed, 0x0c,
	0xa3, 0x0d, 0x66, 0xe0, 0x3b, 0xc6, 0xc8, 0x98, 0x74, 0x5c, 0x33, 0xf0, 0x71, 0x08, 0x9d, 0xb7,
	0x20, 0x11, 0x72, 0xeb, 0x45, 0xe4, 0x98, 0xaa, 0x5c, 0x14, 0x70, 0x00, 0xed, 0xd0, 0xd3, 0xa4,
	0xa5, 0xc8, 0x1c, 0xa7, 0x1c, 0x8f, 0x45, 0x20, 0x83, 0x78, 0xeb, 0xfc, 0x3a, 0x70, 0x19, 0xc6,
	0x53, 0x68, 0x09, 0x2f, 0xf4, 0x92, 0xbd, 0xd3, 0x1c, 0x19, 0x13, 0xd3, 0xd5, 0x88, 0xbd, 0x42,
	0x2f, 0x4f, 0x46, 0x82, 0xc7, 0x5b, 0x41, 0x4a, 0x2b, 0x3d, 0xb9, 0x13, 0x3a, 0x95, 0x46, 0x38,
	0x83, 0xdf, 0x9b, 0x84, 0x3c, 0x49, 0xbe, 0xca, 0xd5, 0x9d, 0xf7, 0xa7, 0x3f, 0xa6, 0x2a, 0xbe,
	0x33, 0x15, 0x1b, 0xc3, 0xdf, 0xc2, 0xfc, 0x7d, 0x47, 0x42, 0xe2, 0x09, 0x34, 0x29, 0xe2, 0xf9,
	0xc0, 0x07, 0xc0, 0x7a, 0x60, 0xaf, 0x48, 0xde, 0x84, 0xa1, 0x96, 0x0b, 0x76, 0x01, 0xff, 0xef,
	0x28, 0x24, 0x49, 0xcb, 0x28, 0x2d, 0xb8, 0x74, 0x34, 0x1a, 0xbb, 0x86, 0xfe, 0x33, 0xf7, 0xbd,
	0x54, 0x5e, 0xee, 0x37, 0x06, 0x8b, 0x22, 0xae, 0xd4, 0xb5, 0x79, 0x53, 0xc5, 0xfc, 0xcb, 0x2a,
	0xc2, 0x3e, 0x51, 0xf2, 0x11, 0x6c, 0x08, 0x1f, 0xc0, 0x5e, 0xa8, 0x51, 0xf2, 0x63, 0x55, 0x3b,
	0x0c, 0xce, 0xab, 0x8d, 0xf5, 0x42, 0x59, 0x03, 0x17, 0x00, 0x2b, 0x92, 0x9a, 0xc0, 0x61, 0x8d,
	0x5c, 0x45, 0x1e, 0x54, 0xf7, 0x60, 0x0d, 0xbc, 0x87, 0x3f, 0x85, 0x09, 0x09, 0x3c, 0x2b, 0x09,
	0xcb, 0x0b, 0xac, 0x75, 0xb9, 0x34, 0xd0, 0x05, 0x3b, 0xdb, 0xad, 0x1e, 0xeb, 0x78, 0xa0, 0x51,
	0x89, 0xad, 0x38, 0x0b, 0x6b, 0xe0, 0x23, 0xd8, 0xe5, 0x03, 0x20, 0x2b, 0xfd, 0x55, 0x79, 0x9d,
	0xda, 0x90, 0xb7, 0xcd, 0x17, 0x6b, 0xc6, 0xd7, 0xeb, 0x96, 0x7a, 0x3c, 0x57, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x16, 0x9a, 0xaa, 0xdf, 0x51, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	CreateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*EmployeeResponse, error)
	GetEmploye(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*Employee, error)
	GetEmployees(ctx context.Context, in *GetAllEmployes, opts ...grpc.CallOption) (EmployeeService_GetEmployeesClient, error)
	DeleteEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*DeleteEmloyeReponse, error)
	UpdateEmployee(ctx context.Context, in *UpdateEmployeeRequest, opts ...grpc.CallOption) (*Employee, error)
}

type employeeServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmployeeServiceClient(cc *grpc.ClientConn) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) CreateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := c.cc.Invoke(ctx, "/EmployeeRpc.EmployeeService/CreateEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetEmploye(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*Employee, error) {
	out := new(Employee)
	err := c.cc.Invoke(ctx, "/EmployeeRpc.EmployeeService/GetEmploye", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetEmployees(ctx context.Context, in *GetAllEmployes, opts ...grpc.CallOption) (EmployeeService_GetEmployeesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EmployeeService_serviceDesc.Streams[0], "/EmployeeRpc.EmployeeService/GetEmployees", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceGetEmployeesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EmployeeService_GetEmployeesClient interface {
	Recv() (*Employee, error)
	grpc.ClientStream
}

type employeeServiceGetEmployeesClient struct {
	grpc.ClientStream
}

func (x *employeeServiceGetEmployeesClient) Recv() (*Employee, error) {
	m := new(Employee)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *employeeServiceClient) DeleteEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*DeleteEmloyeReponse, error) {
	out := new(DeleteEmloyeReponse)
	err := c.cc.Invoke(ctx, "/EmployeeRpc.EmployeeService/DeleteEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) UpdateEmployee(ctx context.Context, in *UpdateEmployeeRequest, opts ...grpc.CallOption) (*Employee, error) {
	out := new(Employee)
	err := c.cc.Invoke(ctx, "/EmployeeRpc.EmployeeService/UpdateEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
type EmployeeServiceServer interface {
	CreateEmployee(context.Context, *Employee) (*EmployeeResponse, error)
	GetEmploye(context.Context, *EmployeeRequest) (*Employee, error)
	GetEmployees(*GetAllEmployes, EmployeeService_GetEmployeesServer) error
	DeleteEmployee(context.Context, *EmployeeRequest) (*DeleteEmloyeReponse, error)
	UpdateEmployee(context.Context, *UpdateEmployeeRequest) (*Employee, error)
}

// UnimplementedEmployeeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmployeeServiceServer struct {
}

func (*UnimplementedEmployeeServiceServer) CreateEmployee(ctx context.Context, req *Employee) (*EmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmployee not implemented")
}
func (*UnimplementedEmployeeServiceServer) GetEmploye(ctx context.Context, req *EmployeeRequest) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmploye not implemented")
}
func (*UnimplementedEmployeeServiceServer) GetEmployees(req *GetAllEmployes, srv EmployeeService_GetEmployeesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEmployees not implemented")
}
func (*UnimplementedEmployeeServiceServer) DeleteEmployee(ctx context.Context, req *EmployeeRequest) (*DeleteEmloyeReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployee not implemented")
}
func (*UnimplementedEmployeeServiceServer) UpdateEmployee(ctx context.Context, req *UpdateEmployeeRequest) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployee not implemented")
}

func RegisterEmployeeServiceServer(s *grpc.Server, srv EmployeeServiceServer) {
	s.RegisterService(&_EmployeeService_serviceDesc, srv)
}

func _EmployeeService_CreateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).CreateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeRpc.EmployeeService/CreateEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).CreateEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetEmploye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetEmploye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeRpc.EmployeeService/GetEmploye",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetEmploye(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetEmployees_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllEmployes)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmployeeServiceServer).GetEmployees(m, &employeeServiceGetEmployeesServer{stream})
}

type EmployeeService_GetEmployeesServer interface {
	Send(*Employee) error
	grpc.ServerStream
}

type employeeServiceGetEmployeesServer struct {
	grpc.ServerStream
}

func (x *employeeServiceGetEmployeesServer) Send(m *Employee) error {
	return x.ServerStream.SendMsg(m)
}

func _EmployeeService_DeleteEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).DeleteEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeRpc.EmployeeService/DeleteEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).DeleteEmployee(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_UpdateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).UpdateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeRpc.EmployeeService/UpdateEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).UpdateEmployee(ctx, req.(*UpdateEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmployeeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "EmployeeRpc.EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmployee",
			Handler:    _EmployeeService_CreateEmployee_Handler,
		},
		{
			MethodName: "GetEmploye",
			Handler:    _EmployeeService_GetEmploye_Handler,
		},
		{
			MethodName: "DeleteEmployee",
			Handler:    _EmployeeService_DeleteEmployee_Handler,
		},
		{
			MethodName: "UpdateEmployee",
			Handler:    _EmployeeService_UpdateEmployee_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEmployees",
			Handler:       _EmployeeService_GetEmployees_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/employee.proto",
}
