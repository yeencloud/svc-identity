package main

import (
	"context"

	baseservice "github.com/yeencloud/lib-base"
	rpc2 "github.com/yeencloud/lib-rpc"
	rpcConfig "github.com/yeencloud/lib-rpc/domain/config"
	sharedConfig "github.com/yeencloud/lib-shared/config"
	contract "github.com/yeencloud/svc-identity/contract/proto/generated"
	"github.com/yeencloud/svc-identity/internal/adapters/database"
	"github.com/yeencloud/svc-identity/internal/adapters/http"
	"github.com/yeencloud/svc-identity/internal/adapters/rpc"
	"github.com/yeencloud/svc-identity/internal/domain"
	"github.com/yeencloud/svc-identity/internal/domain/config"
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

		adminCfg, err := sharedConfig.FetchConfig[config.Admin]()
		if err != nil {
			return err
		}

		authenticationCfg, err := sharedConfig.FetchConfig[config.Authentication]()
		if err != nil {
			return err
		}

		registrationCfg, err := sharedConfig.FetchConfig[config.Registration]()
		if err != nil {
			return err
		}

		appConfig := domain.AppConfig{
			Admin:          *adminCfg,
			Authentication: *authenticationCfg,
			Registration:   *registrationCfg,
		}

		usecases := service.NewUsecases(database.NewUserRepo(), appConfig, mqPublisher, db.Gorm)

		http.NewHTTPServer(httpServer, usecases, db.Gorm)

		rpcServer := rpc2.NewRPCServer(rpcCfg)
		rpcHandler := rpc.NewRPCHandler(usecases, svc.Validator)
		contract.RegisterIdentityServiceServer(rpcServer.RpcServer, rpcHandler)
		rpcServer.Start(ctx)

		return nil
	})
}
