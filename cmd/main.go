package main

import (
	"context"

	baseservice "github.com/yeencloud/lib-base"
	rpc2 "github.com/yeencloud/lib-rpc"
	rpcConfig "github.com/yeencloud/lib-rpc/domain/config"
	sharedConfig "github.com/yeencloud/lib-shared/config"
	contract "github.com/yeencloud/svc-identity/contract/proto"
	"github.com/yeencloud/svc-identity/internal/adapters/database"
	"github.com/yeencloud/svc-identity/internal/adapters/http"
	"github.com/yeencloud/svc-identity/internal/adapters/rpc"
	"github.com/yeencloud/svc-identity/internal/service"
)

func main() {
	baseservice.Run("identity", baseservice.Options{
		UseDatabase: true,
		UseEvents:   true,
	}, func(ctx context.Context, svc *baseservice.BaseService) error {
		dbEngine, err := svc.GetDatabase()
		if err != nil {
			return err
		}

		db, err := database.NewDatabase(ctx, dbEngine.Gorm)
		if err != nil {
			return err
		}

		httpServer, err := svc.GetHttpServer()
		if err != nil {
			return err
		}

		mqPublisher, err := svc.GetMqPublisher()
		if err != nil {
			return err
		}

		rpcCfg, err := sharedConfig.FetchConfig[rpcConfig.Config]()
		if err != nil {
			return err
		}

		usecases := service.NewUsecases(database.NewUserRepo(), mqPublisher)

		http.NewHTTPServer(httpServer, usecases, db.Gorm)

		rpcServer := rpc2.NewRPCServer(rpcCfg)
		rpcResponder := rpc.NewRPCServer(usecases)
		contract.RegisterIdentityServiceServer(rpcServer.RpcServer, rpcResponder)
		rpcServer.Start(ctx)

		return nil
	})
}
