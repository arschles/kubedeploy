package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	client "github.com/jchauncey/kubeclient/client/unversioned"
)

const (
	appName         = "app_name"
	deployedVersion = "version"
)

type deployResp struct {
	Status string `json:"status"`
}

// Deploy is a handlers.Handler that knows how to do deploys
type Deploy struct {
	cl *client.Client
}

// NewDeploy creates a new deployment handler
func NewDeploy(cl *client.Client) *Deploy {
	return &Deploy{cl: cl}
}

// ServeHTTP is the http.Handler interface implementation
func (d *Deploy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(deployResp{Status: "deployed"}); err != nil {
		http.Error(w, "error encoding JSON", http.StatusInternalServerError)
	}
}

// PathInfo is the handlers.Handler interface implementation
func (d *Deploy) PathInfo() (HTTPMethod, string) {
	return HTTPPost, fmt.Sprintf("/deployments/{%s}/{%s}", appName, deployedVersion)
}
