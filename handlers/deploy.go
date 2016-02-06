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

type Deploy struct {
	cl *client.Client
}

func NewDeploy(cl *client.Client) *Deploy {
	return &Deploy{cl: cl}
}

func (d *Deploy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(deployResp{Status: "deployed"}); err != nil {
		http.Error(w, "error encoding JSON", http.StatusInternalServerError)
	}
}

func (d *Deploy) PathInfo() (HTTPMethod, string) {
	return HTTPPost, fmt.Sprintf("/deployments/{%s}/{%s}", appName, deployedVersion)
}
