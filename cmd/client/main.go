package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"github.com/gorilla/mux"
	cpb "github.com/mikehelmick/grpc-istio/pkg/counter/pb"
)

var (
	port       = flag.Int("port", 8080, "port for http server")
	serverAddr = flag.String("addr", "127.0.0.1:3232", "address of server")
)

func clientHandler(client cpb.EchoClient) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// see if there is a counter query param
		param := req.URL.Query()["counter"]
		name := "counter"
		if len(param) > 0 {
			name = param[0] // just take the first
		}

		request := &cpb.IncrementRequest{
			Name: name,
		}
		result, err := client.Increment(req.Context(), request)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error calling increment service: %v", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("counter named: %q has value: %v", name, result.Value)))
	}
}

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := cpb.NewEchoClient(conn)

	r := mux.NewRouter()
	r.HandleFunc("/", clientHandler(client)).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("127.0.0.1:%v", *port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
