package internalhttp

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct { // TODO
	ctx context.Context
}

type MyHandler struct {
	// some useful field
}

//type Logger interface { // TODO
//	handler http.Handler
//}
type Logger struct {
}
type Application interface { // TODO
}

func NewServer(logger Logger, app Application) *Server {
	return &Server{}
}

//func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("kurwa")
//	serverHandler(w, r)
//	fmt.Println("kurwa2")
//}
func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("kurwa3")
	switch r.URL.Path {
	case "/hello":
		w.Write([]byte("hello"))
	default:
		//resp.Error.Message = fmt.Sprintf("uri %s not found", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)

	}
}
func (s *Server) Start(ctx context.Context) error {
	// TODO
	handlerHttp := &MyHandler{}
	//http.Handle("/hello", handlerHttp)
	server := &http.Server{
		Addr: ":8080",
		//Handler:      http.HandlerFunc(serverHandler),
		Handler:      loggingMiddleware(handlerHttp),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	ctx.Done()
	// TODO
	return nil
}
