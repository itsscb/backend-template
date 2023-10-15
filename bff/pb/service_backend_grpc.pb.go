// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: service_backend.proto

package pb

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

const (
	Backend_Login_FullMethodName                = "/pb.backend/Login"
	Backend_RefreshToken_FullMethodName         = "/pb.backend/RefreshToken"
	Backend_ListSessions_FullMethodName         = "/pb.backend/ListSessions"
	Backend_BlockSession_FullMethodName         = "/pb.backend/BlockSession"
	Backend_GetAccount_FullMethodName           = "/pb.backend/GetAccount"
	Backend_ListAccounts_FullMethodName         = "/pb.backend/ListAccounts"
	Backend_CreateAccount_FullMethodName        = "/pb.backend/CreateAccount"
	Backend_UpdateAccount_FullMethodName        = "/pb.backend/UpdateAccount"
	Backend_UpdateAccountPrivacy_FullMethodName = "/pb.backend/UpdateAccountPrivacy"
	Backend_CreatePerson_FullMethodName         = "/pb.backend/CreatePerson"
	Backend_UpdatePerson_FullMethodName         = "/pb.backend/UpdatePerson"
	Backend_GetPerson_FullMethodName            = "/pb.backend/GetPerson"
	Backend_DeletePerson_FullMethodName         = "/pb.backend/DeletePerson"
	Backend_ListPersons_FullMethodName          = "/pb.backend/ListPersons"
	Backend_CreatePayment_FullMethodName        = "/pb.backend/CreatePayment"
	Backend_GetPayment_FullMethodName           = "/pb.backend/GetPayment"
	Backend_DeletePayment_FullMethodName        = "/pb.backend/DeletePayment"
	Backend_ListPayments_FullMethodName         = "/pb.backend/ListPayments"
	Backend_UpdatePayment_FullMethodName        = "/pb.backend/UpdatePayment"
	Backend_ListReturnsLog_FullMethodName       = "/pb.backend/ListReturnsLog"
	Backend_UploadDocument_FullMethodName       = "/pb.backend/UploadDocument"
	Backend_DeleteDocument_FullMethodName       = "/pb.backend/DeleteDocument"
)

// BackendClient is the client API for Backend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackendClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	ListSessions(ctx context.Context, in *ListSessionsRequest, opts ...grpc.CallOption) (*ListSessionsResponse, error)
	BlockSession(ctx context.Context, in *BlockSessionRequest, opts ...grpc.CallOption) (*BlockSessionResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error)
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error)
	UpdateAccountPrivacy(ctx context.Context, in *UpdateAccountPrivacyRequest, opts ...grpc.CallOption) (*UpdateAccountPrivacyResponse, error)
	CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*CreatePersonResponse, error)
	UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*UpdatePersonResponse, error)
	GetPerson(ctx context.Context, in *GetPersonRequest, opts ...grpc.CallOption) (*GetPersonResponse, error)
	DeletePerson(ctx context.Context, in *DeletePersonRequest, opts ...grpc.CallOption) (*DeletePersonResponse, error)
	ListPersons(ctx context.Context, in *ListPersonsRequest, opts ...grpc.CallOption) (*ListPersonsResponse, error)
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
	GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*GetPaymentResponse, error)
	DeletePayment(ctx context.Context, in *DeletePaymentRequest, opts ...grpc.CallOption) (*DeletePaymentResponse, error)
	ListPayments(ctx context.Context, in *ListPaymentsRequest, opts ...grpc.CallOption) (*ListPaymentsResponse, error)
	UpdatePayment(ctx context.Context, in *UpdatePaymentRequest, opts ...grpc.CallOption) (*UpdatePaymentResponse, error)
	ListReturnsLog(ctx context.Context, in *ListReturnsLogRequest, opts ...grpc.CallOption) (*ListReturnsLogResponse, error)
	UploadDocument(ctx context.Context, in *UploadDocumentRequest, opts ...grpc.CallOption) (*UploadDocumentResponse, error)
	DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*DeleteDocumentResponse, error)
}

type backendClient struct {
	cc grpc.ClientConnInterface
}

func NewBackendClient(cc grpc.ClientConnInterface) BackendClient {
	return &backendClient{cc}
}

func (c *backendClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Backend_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, Backend_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) ListSessions(ctx context.Context, in *ListSessionsRequest, opts ...grpc.CallOption) (*ListSessionsResponse, error) {
	out := new(ListSessionsResponse)
	err := c.cc.Invoke(ctx, Backend_ListSessions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) BlockSession(ctx context.Context, in *BlockSessionRequest, opts ...grpc.CallOption) (*BlockSessionResponse, error) {
	out := new(BlockSessionResponse)
	err := c.cc.Invoke(ctx, Backend_BlockSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, Backend_GetAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error) {
	out := new(ListAccountsResponse)
	err := c.cc.Invoke(ctx, Backend_ListAccounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, Backend_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error) {
	out := new(UpdateAccountResponse)
	err := c.cc.Invoke(ctx, Backend_UpdateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) UpdateAccountPrivacy(ctx context.Context, in *UpdateAccountPrivacyRequest, opts ...grpc.CallOption) (*UpdateAccountPrivacyResponse, error) {
	out := new(UpdateAccountPrivacyResponse)
	err := c.cc.Invoke(ctx, Backend_UpdateAccountPrivacy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*CreatePersonResponse, error) {
	out := new(CreatePersonResponse)
	err := c.cc.Invoke(ctx, Backend_CreatePerson_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*UpdatePersonResponse, error) {
	out := new(UpdatePersonResponse)
	err := c.cc.Invoke(ctx, Backend_UpdatePerson_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) GetPerson(ctx context.Context, in *GetPersonRequest, opts ...grpc.CallOption) (*GetPersonResponse, error) {
	out := new(GetPersonResponse)
	err := c.cc.Invoke(ctx, Backend_GetPerson_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) DeletePerson(ctx context.Context, in *DeletePersonRequest, opts ...grpc.CallOption) (*DeletePersonResponse, error) {
	out := new(DeletePersonResponse)
	err := c.cc.Invoke(ctx, Backend_DeletePerson_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) ListPersons(ctx context.Context, in *ListPersonsRequest, opts ...grpc.CallOption) (*ListPersonsResponse, error) {
	out := new(ListPersonsResponse)
	err := c.cc.Invoke(ctx, Backend_ListPersons_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, Backend_CreatePayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*GetPaymentResponse, error) {
	out := new(GetPaymentResponse)
	err := c.cc.Invoke(ctx, Backend_GetPayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) DeletePayment(ctx context.Context, in *DeletePaymentRequest, opts ...grpc.CallOption) (*DeletePaymentResponse, error) {
	out := new(DeletePaymentResponse)
	err := c.cc.Invoke(ctx, Backend_DeletePayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) ListPayments(ctx context.Context, in *ListPaymentsRequest, opts ...grpc.CallOption) (*ListPaymentsResponse, error) {
	out := new(ListPaymentsResponse)
	err := c.cc.Invoke(ctx, Backend_ListPayments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) UpdatePayment(ctx context.Context, in *UpdatePaymentRequest, opts ...grpc.CallOption) (*UpdatePaymentResponse, error) {
	out := new(UpdatePaymentResponse)
	err := c.cc.Invoke(ctx, Backend_UpdatePayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) ListReturnsLog(ctx context.Context, in *ListReturnsLogRequest, opts ...grpc.CallOption) (*ListReturnsLogResponse, error) {
	out := new(ListReturnsLogResponse)
	err := c.cc.Invoke(ctx, Backend_ListReturnsLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) UploadDocument(ctx context.Context, in *UploadDocumentRequest, opts ...grpc.CallOption) (*UploadDocumentResponse, error) {
	out := new(UploadDocumentResponse)
	err := c.cc.Invoke(ctx, Backend_UploadDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*DeleteDocumentResponse, error) {
	out := new(DeleteDocumentResponse)
	err := c.cc.Invoke(ctx, Backend_DeleteDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendServer is the server API for Backend service.
// All implementations must embed UnimplementedBackendServer
// for forward compatibility
type BackendServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	ListSessions(context.Context, *ListSessionsRequest) (*ListSessionsResponse, error)
	BlockSession(context.Context, *BlockSessionRequest) (*BlockSessionResponse, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error)
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error)
	UpdateAccountPrivacy(context.Context, *UpdateAccountPrivacyRequest) (*UpdateAccountPrivacyResponse, error)
	CreatePerson(context.Context, *CreatePersonRequest) (*CreatePersonResponse, error)
	UpdatePerson(context.Context, *UpdatePersonRequest) (*UpdatePersonResponse, error)
	GetPerson(context.Context, *GetPersonRequest) (*GetPersonResponse, error)
	DeletePerson(context.Context, *DeletePersonRequest) (*DeletePersonResponse, error)
	ListPersons(context.Context, *ListPersonsRequest) (*ListPersonsResponse, error)
	CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
	GetPayment(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error)
	DeletePayment(context.Context, *DeletePaymentRequest) (*DeletePaymentResponse, error)
	ListPayments(context.Context, *ListPaymentsRequest) (*ListPaymentsResponse, error)
	UpdatePayment(context.Context, *UpdatePaymentRequest) (*UpdatePaymentResponse, error)
	ListReturnsLog(context.Context, *ListReturnsLogRequest) (*ListReturnsLogResponse, error)
	UploadDocument(context.Context, *UploadDocumentRequest) (*UploadDocumentResponse, error)
	DeleteDocument(context.Context, *DeleteDocumentRequest) (*DeleteDocumentResponse, error)
	mustEmbedUnimplementedBackendServer()
}

// UnimplementedBackendServer must be embedded to have forward compatible implementations.
type UnimplementedBackendServer struct {
}

func (UnimplementedBackendServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedBackendServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedBackendServer) ListSessions(context.Context, *ListSessionsRequest) (*ListSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSessions not implemented")
}
func (UnimplementedBackendServer) BlockSession(context.Context, *BlockSessionRequest) (*BlockSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockSession not implemented")
}
func (UnimplementedBackendServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedBackendServer) ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}
func (UnimplementedBackendServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedBackendServer) UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedBackendServer) UpdateAccountPrivacy(context.Context, *UpdateAccountPrivacyRequest) (*UpdateAccountPrivacyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccountPrivacy not implemented")
}
func (UnimplementedBackendServer) CreatePerson(context.Context, *CreatePersonRequest) (*CreatePersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePerson not implemented")
}
func (UnimplementedBackendServer) UpdatePerson(context.Context, *UpdatePersonRequest) (*UpdatePersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerson not implemented")
}
func (UnimplementedBackendServer) GetPerson(context.Context, *GetPersonRequest) (*GetPersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (UnimplementedBackendServer) DeletePerson(context.Context, *DeletePersonRequest) (*DeletePersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePerson not implemented")
}
func (UnimplementedBackendServer) ListPersons(context.Context, *ListPersonsRequest) (*ListPersonsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPersons not implemented")
}
func (UnimplementedBackendServer) CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedBackendServer) GetPayment(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayment not implemented")
}
func (UnimplementedBackendServer) DeletePayment(context.Context, *DeletePaymentRequest) (*DeletePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePayment not implemented")
}
func (UnimplementedBackendServer) ListPayments(context.Context, *ListPaymentsRequest) (*ListPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPayments not implemented")
}
func (UnimplementedBackendServer) UpdatePayment(context.Context, *UpdatePaymentRequest) (*UpdatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayment not implemented")
}
func (UnimplementedBackendServer) ListReturnsLog(context.Context, *ListReturnsLogRequest) (*ListReturnsLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReturnsLog not implemented")
}
func (UnimplementedBackendServer) UploadDocument(context.Context, *UploadDocumentRequest) (*UploadDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadDocument not implemented")
}
func (UnimplementedBackendServer) DeleteDocument(context.Context, *DeleteDocumentRequest) (*DeleteDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDocument not implemented")
}
func (UnimplementedBackendServer) mustEmbedUnimplementedBackendServer() {}

// UnsafeBackendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackendServer will
// result in compilation errors.
type UnsafeBackendServer interface {
	mustEmbedUnimplementedBackendServer()
}

func RegisterBackendServer(s grpc.ServiceRegistrar, srv BackendServer) {
	s.RegisterService(&Backend_ServiceDesc, srv)
}

func _Backend_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_ListSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).ListSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_ListSessions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).ListSessions(ctx, req.(*ListSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_BlockSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).BlockSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_BlockSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).BlockSession(ctx, req.(*BlockSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_ListAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).ListAccounts(ctx, req.(*ListAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_UpdateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).UpdateAccount(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_UpdateAccountPrivacy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountPrivacyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).UpdateAccountPrivacy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_UpdateAccountPrivacy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).UpdateAccountPrivacy(ctx, req.(*UpdateAccountPrivacyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_CreatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).CreatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_CreatePerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).CreatePerson(ctx, req.(*CreatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_UpdatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).UpdatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_UpdatePerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).UpdatePerson(ctx, req.(*UpdatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_GetPerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetPerson(ctx, req.(*GetPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_DeletePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).DeletePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_DeletePerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).DeletePerson(ctx, req.(*DeletePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_ListPersons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPersonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).ListPersons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_ListPersons_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).ListPersons(ctx, req.(*ListPersonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_CreatePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_GetPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_GetPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetPayment(ctx, req.(*GetPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_DeletePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).DeletePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_DeletePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).DeletePayment(ctx, req.(*DeletePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_ListPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).ListPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_ListPayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).ListPayments(ctx, req.(*ListPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_UpdatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).UpdatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_UpdatePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).UpdatePayment(ctx, req.(*UpdatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_ListReturnsLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReturnsLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).ListReturnsLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_ListReturnsLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).ListReturnsLog(ctx, req.(*ListReturnsLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_UploadDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).UploadDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_UploadDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).UploadDocument(ctx, req.(*UploadDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_DeleteDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).DeleteDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_DeleteDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).DeleteDocument(ctx, req.(*DeleteDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Backend_ServiceDesc is the grpc.ServiceDesc for Backend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Backend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.backend",
	HandlerType: (*BackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Backend_Login_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _Backend_RefreshToken_Handler,
		},
		{
			MethodName: "ListSessions",
			Handler:    _Backend_ListSessions_Handler,
		},
		{
			MethodName: "BlockSession",
			Handler:    _Backend_BlockSession_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Backend_GetAccount_Handler,
		},
		{
			MethodName: "ListAccounts",
			Handler:    _Backend_ListAccounts_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Backend_CreateAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _Backend_UpdateAccount_Handler,
		},
		{
			MethodName: "UpdateAccountPrivacy",
			Handler:    _Backend_UpdateAccountPrivacy_Handler,
		},
		{
			MethodName: "CreatePerson",
			Handler:    _Backend_CreatePerson_Handler,
		},
		{
			MethodName: "UpdatePerson",
			Handler:    _Backend_UpdatePerson_Handler,
		},
		{
			MethodName: "GetPerson",
			Handler:    _Backend_GetPerson_Handler,
		},
		{
			MethodName: "DeletePerson",
			Handler:    _Backend_DeletePerson_Handler,
		},
		{
			MethodName: "ListPersons",
			Handler:    _Backend_ListPersons_Handler,
		},
		{
			MethodName: "CreatePayment",
			Handler:    _Backend_CreatePayment_Handler,
		},
		{
			MethodName: "GetPayment",
			Handler:    _Backend_GetPayment_Handler,
		},
		{
			MethodName: "DeletePayment",
			Handler:    _Backend_DeletePayment_Handler,
		},
		{
			MethodName: "ListPayments",
			Handler:    _Backend_ListPayments_Handler,
		},
		{
			MethodName: "UpdatePayment",
			Handler:    _Backend_UpdatePayment_Handler,
		},
		{
			MethodName: "ListReturnsLog",
			Handler:    _Backend_ListReturnsLog_Handler,
		},
		{
			MethodName: "UploadDocument",
			Handler:    _Backend_UploadDocument_Handler,
		},
		{
			MethodName: "DeleteDocument",
			Handler:    _Backend_DeleteDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_backend.proto",
}