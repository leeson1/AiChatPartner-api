// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: db.proto

package server

import (
	"context"

	"AiChatPartner/rpc/db/db"
	"AiChatPartner/rpc/db/internal/logic/databaseservice"
	"AiChatPartner/rpc/db/internal/svc"
)

type DatabaseServiceServer struct {
	svcCtx *svc.ServiceContext
	db.UnimplementedDatabaseServiceServer
}

func NewDatabaseServiceServer(svcCtx *svc.ServiceContext) *DatabaseServiceServer {
	return &DatabaseServiceServer{
		svcCtx: svcCtx,
	}
}

// 连接管理
func (s *DatabaseServiceServer) Connect(ctx context.Context, in *db.ConnectRequest) (*db.ConnectResponse, error) {
	l := databaseservicelogic.NewConnectLogic(ctx, s.svcCtx)
	return l.Connect(in)
}

func (s *DatabaseServiceServer) Disconnect(ctx context.Context, in *db.DisconnectRequest) (*db.DisconnectResponse, error) {
	l := databaseservicelogic.NewDisconnectLogic(ctx, s.svcCtx)
	return l.Disconnect(in)
}

// CRUD 操作
func (s *DatabaseServiceServer) Create(ctx context.Context, in *db.CreateRequest) (*db.CreateResponse, error) {
	l := databaseservicelogic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *DatabaseServiceServer) Read(ctx context.Context, in *db.ReadRequest) (*db.ReadResponse, error) {
	l := databaseservicelogic.NewReadLogic(ctx, s.svcCtx)
	return l.Read(in)
}

func (s *DatabaseServiceServer) Update(ctx context.Context, in *db.UpdateRequest) (*db.UpdateResponse, error) {
	l := databaseservicelogic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *DatabaseServiceServer) Delete(ctx context.Context, in *db.DeleteRequest) (*db.DeleteResponse, error) {
	l := databaseservicelogic.NewDeleteLogic(ctx, s.svcCtx)
	return l.Delete(in)
}

// 事务管理
func (s *DatabaseServiceServer) BeginTransaction(ctx context.Context, in *db.BeginTransactionRequest) (*db.BeginTransactionResponse, error) {
	l := databaseservicelogic.NewBeginTransactionLogic(ctx, s.svcCtx)
	return l.BeginTransaction(in)
}

func (s *DatabaseServiceServer) CommitTransaction(ctx context.Context, in *db.CommitTransactionRequest) (*db.CommitTransactionResponse, error) {
	l := databaseservicelogic.NewCommitTransactionLogic(ctx, s.svcCtx)
	return l.CommitTransaction(in)
}

func (s *DatabaseServiceServer) RollbackTransaction(ctx context.Context, in *db.RollbackTransactionRequest) (*db.RollbackTransactionResponse, error) {
	l := databaseservicelogic.NewRollbackTransactionLogic(ctx, s.svcCtx)
	return l.RollbackTransaction(in)
}
