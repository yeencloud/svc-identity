package http

import (
	"github.com/yeencloud/lib-base/transaction"
	httpserver "github.com/yeencloud/lib-httpserver"
	"github.com/yeencloud/svc-identity/internal/ports"
)

type HTTPServer struct {
	server *httpserver.HttpServer

	usecases ports.Usecases

	trx transaction.TransactionInterface
}

func NewHTTPServer(server *httpserver.HttpServer, usecases ports.Usecases, trx transaction.TransactionInterface) *HTTPServer {
	httpServer := &HTTPServer{
		server:   server,
		usecases: usecases,
		trx:      trx,
	}

	return httpServer
}
