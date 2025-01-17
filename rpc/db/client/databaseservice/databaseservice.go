// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: db.proto

package databaseservice

import (
	"context"

	"AiChatPartner/rpc/db/db"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BeginTransactionRequest     = db.BeginTransactionRequest
	BeginTransactionResponse    = db.BeginTransactionResponse
	CommitTransactionRequest    = db.CommitTransactionRequest
	CommitTransactionResponse   = db.CommitTransactionResponse
	ConnectRequest              = db.ConnectRequest
	ConnectResponse             = db.ConnectResponse
	CreateRequest               = db.CreateRequest
	CreateResponse              = db.CreateResponse
	DelRequest                  = db.DelRequest
	DelResponse                 = db.DelResponse
	DeleteRequest               = db.DeleteRequest
	DeleteResponse              = db.DeleteResponse
	DisconnectRequest           = db.DisconnectRequest
	DisconnectResponse          = db.DisconnectResponse
	GetRequest                  = db.GetRequest
	GetResponse                 = db.GetResponse
	ReadRequest                 = db.ReadRequest
	ReadResponse                = db.ReadResponse
	RollbackTransactionRequest  = db.RollbackTransactionRequest
	RollbackTransactionResponse = db.RollbackTransactionResponse
	SetRequest                  = db.SetRequest
	SetResponse                 = db.SetResponse
	UpdateRequest               = db.UpdateRequest
	UpdateResponse              = db.UpdateResponse

	DatabaseService interface {
		// 连接管理
		Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
		Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
		// CRUD 操作
		Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
		Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
		Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
		Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
		// 事务管理
		BeginTransaction(ctx context.Context, in *BeginTransactionRequest, opts ...grpc.CallOption) (*BeginTransactionResponse, error)
		CommitTransaction(ctx context.Context, in *CommitTransactionRequest, opts ...grpc.CallOption) (*CommitTransactionResponse, error)
		RollbackTransaction(ctx context.Context, in *RollbackTransactionRequest, opts ...grpc.CallOption) (*RollbackTransactionResponse, error)
	}

	defaultDatabaseService struct {
		cli zrpc.Client
	}
)

func NewDatabaseService(cli zrpc.Client) DatabaseService {
	return &defaultDatabaseService{
		cli: cli,
	}
}

// 连接管理
func (m *defaultDatabaseService) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	client := db.NewDatabaseServiceClient(m.cli.Conn())
	return client.Connect(ctx, in, opts...)
}

func (m *defaultDatabaseService) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	client := db.NewDatabaseServiceClient(m.cli.Conn())
	return client.Disconnect(ctx, in, opts...)
}

// CRUD 操作
func (m *defaultDatabaseService) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	client := db.NewDatabaseServiceClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultDatabaseService) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	client := db.NewDatabaseServiceClient(m.cli.Conn())
	return client.Read(ctx, in, opts...)
}

func (m *defaultDatabaseService) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	client := db.NewDatabaseServiceClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}

func (m *defaultDatabaseService) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	client := db.NewDatabaseServiceClient(m.cli.Conn())
	return client.Delete(ctx, in, opts...)
}

// 事务管理
func (m *defaultDatabaseService) BeginTransaction(ctx context.Context, in *BeginTransactionRequest, opts ...grpc.CallOption) (*BeginTransactionResponse, error) {
	client := db.NewDatabaseServiceClient(m.cli.Conn())
	return client.BeginTransaction(ctx, in, opts...)
}

func (m *defaultDatabaseService) CommitTransaction(ctx context.Context, in *CommitTransactionRequest, opts ...grpc.CallOption) (*CommitTransactionResponse, error) {
	client := db.NewDatabaseServiceClient(m.cli.Conn())
	return client.CommitTransaction(ctx, in, opts...)
}

func (m *defaultDatabaseService) RollbackTransaction(ctx context.Context, in *RollbackTransactionRequest, opts ...grpc.CallOption) (*RollbackTransactionResponse, error) {
	client := db.NewDatabaseServiceClient(m.cli.Conn())
	return client.RollbackTransaction(ctx, in, opts...)
}
