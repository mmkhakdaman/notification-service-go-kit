package server

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
	"net"
	"notification_service/config"
	"notification_service/database"
	"notification_service/endpoints"
	"notification_service/pb"
	"notification_service/services"
	transport "notification_service/transports"
	"os"
	"os/signal"
	"syscall"
)

// GRPCServer represents a gRPC server for the NotificationService
type GRPCServer struct {
	server *grpc.Server
}

// NewGRPCServer returns a new instance of GRPCServer
func NewGRPCServer() *GRPCServer {
	return &GRPCServer{server: grpc.NewServer()}
}

// Start starts the gRPC server and listens for incoming requests
func (s *GRPCServer) Start() error {

	logger := s.logger()

	addService := services.NewService(logger, database.GetDB())
	addEndpoint := endpoints.MakeEndpoints(addService)
	grpcServer := transport.NewGRPCServer(addEndpoint, logger)

	serverPort := config.GetEnv("SERVER_PORT")
	grpcListener, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go server(grpcServer, logger, grpcListener)

	errs := errHandler()
	level.Error(logger).Log("exit", <-errs)
	return nil
}

func server(grpcServer pb.NotificationServiceServer, logger log.Logger, grpcListener net.Listener) {
	baseServer := grpc.NewServer()
	pb.RegisterNotificationServiceServer(baseServer, grpcServer)
	level.Info(logger).Log("msg", "Server started successfully 🚀")
	baseServer.Serve(grpcListener)
}

func errHandler() chan error {
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	return errs
}

func (s *GRPCServer) logger() log.Logger {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	return logger
}

// Stop stops the gRPC server and closes all connections
func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
}
