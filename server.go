package astro

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpSrv *http.Server
}

func (s *Server) Run(port string, h http.Handler) error {
	s.httpSrv = &http.Server{
		Addr:           ":" + port,
		Handler:        h,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpSrv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpSrv.Shutdown(ctx)
}
