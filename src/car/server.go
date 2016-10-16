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
	stopListener chan struct{}
	port         int
	server       *http.Server
	steering     Steering
	mux          map[string]func(Steering, http.ResponseWriter, *http.Request)
}

func NewServer(port int, steering Steering) Server {
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
		steering: steering,
		mux :make(map[string]func(Steering, http.ResponseWriter, *http.Request)),
	}

	srv.mux["/left"] = left
	srv.mux["/right"] = right
	srv.mux["/forward"] = forward
	srv.mux["/backward"] = backward
	srv.mux["/resetSteering"] = resetSteering
	srv.mux["/resetThrust"] = resetThrust

	return &srv
}

func left(s Steering, w http.ResponseWriter, r *http.Request) {
	s.Left()
}

func right(s Steering, w http.ResponseWriter, r *http.Request) {
	s.Right()
}

func forward(s Steering, w http.ResponseWriter, r *http.Request) {
	s.Forward()
}

func backward(s Steering, w http.ResponseWriter, r *http.Request) {
	s.Backward()
}

func resetSteering(s Steering, w http.ResponseWriter, r *http.Request) {
	s.ResetSteering()
}

func resetThrust(s Steering, w http.ResponseWriter, r *http.Request) {
	s.ResetThrusting()
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := s.mux[r.URL.String()]; ok {
		h(s.steering, w, r)
		return
	}
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
