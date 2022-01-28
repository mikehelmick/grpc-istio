package server

import (
	"context"
	"testing"

	cpb "github.com/mikehelmick/grpc-istio/pkg/counter/pb"
)

func TestIncrement(t *testing.T) {
	t.Parallel()

	srv := NewServer()
	ctx := context.Background()

	result, err := srv.Increment(ctx, &cpb.IncrementRequest{
		Name: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	if result.Value != 1 {
		t.Fatalf("wrong value want: 1 got: %v", result.Value)
	}
}
