package car

import (
	"net/http"
	"time"
	"strconv"
)

type Server interface {
	Start() error
	Wait()
	Stop()
}

type server struct {
	stopListener  chan struct{}
	port          int
	server        *http.Server
	staticHandler http.Handler
	steering      Steering
	mux           map[string]func(*server, http.ResponseWriter, *http.Request)
}

func NewServer(port int, steering Steering, staticPath string) Server {
	var srv server

	s := &http.Server{
		Addr:           ":" + strconv.Itoa(port),
		Handler:        &srv,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	srv = server{
		stopListener: make(chan struct{}),
		port: port,
		server: s,
		staticHandler: http.FileServer(http.Dir(staticPath)),
		steering: steering,
		mux :make(map[string]func(*server, http.ResponseWriter, *http.Request)),
	}

	srv.mux["/"] = static
	srv.mux["/left"] = left
	srv.mux["/right"] = right
	srv.mux["/forward"] = forward
	srv.mux["/backward"] = backward
	srv.mux["/resetSteering"] = resetSteering
	srv.mux["/resetThrust"] = resetThrust

	return &srv
}

func static(s *server, w http.ResponseWriter, r *http.Request) {
	s.staticHandler.ServeHTTP(w, r)

}

func left(s *server, w http.ResponseWriter, r *http.Request) {
	s.steering.Left()
}

func right(s *server, w http.ResponseWriter, r *http.Request) {
	s.steering.Right()
}

func forward(s *server, w http.ResponseWriter, r *http.Request) {
	s.steering.Forward()
}

func backward(s *server, w http.ResponseWriter, r *http.Request) {
	s.steering.Backward()
}

func resetSteering(s *server, w http.ResponseWriter, r *http.Request) {
	s.steering.ResetSteering()
}

func resetThrust(s *server, w http.ResponseWriter, r *http.Request) {
	s.steering.ResetThrusting()
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := s.mux[r.URL.String()]; ok {
		h(s, w, r)
		return
	}
	s.mux["/"](s, w, r)
}

func (s *server) Start() error {
	return s.server.ListenAndServe()
}

func (s *server) Wait() {
	<-s.stopListener
}

func (s *server) Stop() {
	s.stopListener <- struct{}{}
}
