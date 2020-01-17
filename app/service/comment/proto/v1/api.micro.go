// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api.proto

package gmicro_srv_comment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PostCommentSrv service

type PostCommentSrvService interface {
	AddComment(ctx context.Context, in *AddCommentRequest, opts ...client.CallOption) (*AddCommentRespond, error)
	DelComment(ctx context.Context, in *DelCommentRequest, opts ...client.CallOption) (*DelCommentRespond, error)
	ListComment(ctx context.Context, in *ListCommentRequest, opts ...client.CallOption) (*ListCommentRespond, error)
}

type postCommentSrvService struct {
	c    client.Client
	name string
}

func NewPostCommentSrvService(name string, c client.Client) PostCommentSrvService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gmicro.srv.comment"
	}
	return &postCommentSrvService{
		c:    c,
		name: name,
	}
}

func (c *postCommentSrvService) AddComment(ctx context.Context, in *AddCommentRequest, opts ...client.CallOption) (*AddCommentRespond, error) {
	req := c.c.NewRequest(c.name, "PostCommentSrv.AddComment", in)
	out := new(AddCommentRespond)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postCommentSrvService) DelComment(ctx context.Context, in *DelCommentRequest, opts ...client.CallOption) (*DelCommentRespond, error) {
	req := c.c.NewRequest(c.name, "PostCommentSrv.DelComment", in)
	out := new(DelCommentRespond)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postCommentSrvService) ListComment(ctx context.Context, in *ListCommentRequest, opts ...client.CallOption) (*ListCommentRespond, error) {
	req := c.c.NewRequest(c.name, "PostCommentSrv.ListComment", in)
	out := new(ListCommentRespond)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PostCommentSrv service

type PostCommentSrvHandler interface {
	AddComment(context.Context, *AddCommentRequest, *AddCommentRespond) error
	DelComment(context.Context, *DelCommentRequest, *DelCommentRespond) error
	ListComment(context.Context, *ListCommentRequest, *ListCommentRespond) error
}

func RegisterPostCommentSrvHandler(s server.Server, hdlr PostCommentSrvHandler, opts ...server.HandlerOption) error {
	type postCommentSrv interface {
		AddComment(ctx context.Context, in *AddCommentRequest, out *AddCommentRespond) error
		DelComment(ctx context.Context, in *DelCommentRequest, out *DelCommentRespond) error
		ListComment(ctx context.Context, in *ListCommentRequest, out *ListCommentRespond) error
	}
	type PostCommentSrv struct {
		postCommentSrv
	}
	h := &postCommentSrvHandler{hdlr}
	return s.Handle(s.NewHandler(&PostCommentSrv{h}, opts...))
}

type postCommentSrvHandler struct {
	PostCommentSrvHandler
}

func (h *postCommentSrvHandler) AddComment(ctx context.Context, in *AddCommentRequest, out *AddCommentRespond) error {
	return h.PostCommentSrvHandler.AddComment(ctx, in, out)
}

func (h *postCommentSrvHandler) DelComment(ctx context.Context, in *DelCommentRequest, out *DelCommentRespond) error {
	return h.PostCommentSrvHandler.DelComment(ctx, in, out)
}

func (h *postCommentSrvHandler) ListComment(ctx context.Context, in *ListCommentRequest, out *ListCommentRespond) error {
	return h.PostCommentSrvHandler.ListComment(ctx, in, out)
}
