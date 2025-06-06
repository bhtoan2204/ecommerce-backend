package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"payment/package/logger"
	"strconv"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Server provides a graceful shutdown
type Server struct {
	ip       string
	port     string
	listener net.Listener
}

// New create a new server listening on the provided port. It will starts the listener but
// does not start the server.
func New(port int) (*Server, error) {
	// addr := ":0" // service discovery in the future
	addr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Server{
		ip:       listener.Addr().(*net.TCPAddr).IP.String(),
		port:     strconv.Itoa(listener.Addr().(*net.TCPAddr).Port),
		listener: listener,
	}, nil
}

// ServeHTTP start the server and block until the context is closed.
func (s *Server) ServeHTTP(ctx context.Context, srv *http.Server) error {
	log := logger.FromContext(ctx)

	errCh := make(chan error, 1)
	go func() {
		<-ctx.Done()
		log.Debug("ServeHTTP: context is closed")

		shutdownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		log.Info("ServeHTTP: shutting down")
		errCh <- srv.Shutdown(shutdownCtx)
	}()

	// This will block until the context is closed.
	err := srv.Serve(s.listener)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to serve: %w", err)
	}

	log.Info("ServeHTTP: serve stopped")

	err = <-errCh
	return err
}

// ServeHTTPHandle is a wrapper of ServeHTTP. It create HTTP server by provided http.Handler.
func (s *Server) ServeHTTPHandler(ctx context.Context, handler http.Handler) error {
	return s.ServeHTTP(ctx, &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		Handler:           handler,
	})
}

// ServeGRPC starts the server and blocks until the provided context is closed.
func (s *Server) ServeGRPC(ctx context.Context, srv *grpc.Server) error {
	log := logger.FromContext(ctx)

	go func() {
		<-ctx.Done()

		log.Info("ServeGRPC: shutting down")
		srv.GracefulStop()
	}()

	log.Info("ServeGRPC: gRPC server started", zap.String("addr", s.listener.Addr().String()))
	// Run the server. This will block until the provided context is closed.
	if err := srv.Serve(s.listener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		return fmt.Errorf("failed to serve: %w", err)
	}

	log.Info("ServeGPRC: serve stopped")

	return nil
}

func (s *Server) Addr() string {
	return net.JoinHostPort(s.ip, s.port)
}

func (s *Server) IP() string {
	return s.ip
}

func (s *Server) Port() string {
	return s.port
}
