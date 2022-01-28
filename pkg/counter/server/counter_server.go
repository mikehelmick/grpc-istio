package server

import (
	"context"
	"log"
	"sync"

	cpb "github.com/mikehelmick/grpc-istio/pkg/counter/pb"
)

type counterServer struct {
	state map[string]int64
	mu    sync.Mutex

	cpb.UnimplementedEchoServer
}

func NewServer() cpb.EchoServer {
	return &counterServer{
		state: make(map[string]int64),
	}
}

func (cs *counterServer) Increment(ctx context.Context, request *cpb.IncrementRequest) (*cpb.IncrementResponse, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if _, ok := cs.state[request.Name]; !ok {
		cs.state[request.Name] = 0
	}

	value := cs.state[request.Name] + 1
	cs.state[request.Name] = value

	log.Printf("Echo.Increment: name: %q value: %v", request.Name, value)

	return &cpb.IncrementResponse{
		Name:  request.Name,
		Value: value,
	}, nil
}
