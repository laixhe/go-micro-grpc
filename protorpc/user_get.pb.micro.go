// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user_get.proto

package protorpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	// 定义 GetUser 方法 - 获取某个 user 数据
	GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error)
	// 定义 GetUserList 方法 - 获取 user 所有数据
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...client.CallOption) (*UserListResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetUser", in)
	out := new(GetUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...client.CallOption) (*UserListResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetUserList", in)
	out := new(UserListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	// 定义 GetUser 方法 - 获取某个 user 数据
	GetUser(context.Context, *GetUserRequest, *GetUserResponse) error
	// 定义 GetUserList 方法 - 获取 user 所有数据
	GetUserList(context.Context, *GetUserListRequest, *UserListResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error
		GetUserList(ctx context.Context, in *GetUserListRequest, out *UserListResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error {
	return h.UserHandler.GetUser(ctx, in, out)
}

func (h *userHandler) GetUserList(ctx context.Context, in *GetUserListRequest, out *UserListResponse) error {
	return h.UserHandler.GetUserList(ctx, in, out)
}
