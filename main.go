package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/arschles/kubedeploy/handlers"
	"github.com/gorilla/mux"
	client "github.com/jchauncey/kubeclient/client/unversioned"
)

const (
	hostEnvVar  = "SERVE_HOST"
	portEnvVar  = "SERVE_PORT"
	defaultHost = "0.0.0.0"
	defaultPort = 8080
)

func main() {
	cl, err := client.NewInCluster()
	if err != nil {
		log.Printf("Error creating Kubernetes client (%s)", err)
		os.Exit(1)
	}

	router := mux.NewRouter()

	deployHandler := handlers.NewDeploy(cl)
	method, path := deployHandler.PathInfo()
	router.Handle(path, deployHandler).Methods(method.String())

	host := os.Getenv(hostEnvVar)
	if host == "" {
		host = defaultHost
	}
	port, err := strconv.Atoi(os.Getenv(portEnvVar))
	if err != nil {
		port = defaultPort
	}
	hostStr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Serving on %s", hostStr)
	if err := http.ListenAndServe(hostStr, router); err != nil {
		log.Printf("Error serving (%s)", err)
		os.Exit(1)
	}
}
