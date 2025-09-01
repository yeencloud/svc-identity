package main

import (
	"context"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	baseservice "github.com/yeencloud/lib-base"
	contract "github.com/yeencloud/svc-identity/contract/proto"
	"github.com/yeencloud/svc-identity/internal/adapters/database"
	"github.com/yeencloud/svc-identity/internal/adapters/http"
	"github.com/yeencloud/svc-identity/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	contract.IdentityServiceServer
}

func (s *server) Register(_ context.Context, in *contract.RegisterObject) (*contract.RegisterResponse, error) {
	println("CALLED")
	return &contract.RegisterResponse{Id: in.Mail}, nil
}

func unaryInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {

	m, err := handler(ctx, req)
	if err != nil {
		logger := log.Error
		logger("RPC failed with error: %v", err)
	}
	return m, err
}

func main() {
	baseservice.Run("base-service", baseservice.Options{
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

		grpcPort := 6042
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
		if err != nil {
			return err
		}

		s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))

		reflection.Register(s)
		contract.RegisterIdentityServiceServer(s, &server{})
		go func() {
			log.Info("Start service")
			if err := s.Serve(lis); err != nil {
				log.Fatalln(err)
			}
		}()

		usecases := service.NewUsecases(database.NewViewOriginRepo(), mqPublisher)
		http.NewHTTPServer(httpServer, usecases, db.Gorm)

		return nil
	})
}
