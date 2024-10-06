package main

import (
	"fmt"
	sw "github.com/ZPI-2024-25/KubernetesUserManager/go/api"
	"github.com/ZPI-2024-25/KubernetesUserManager/go/common"
	"log"
	"net/http"
)

var clientsetSingleton *common.ClientSetSingleton

func main() {
	log.Printf("Server started")
	router := sw.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

/*
	TODO
    This code should be removed as well as deploy folder and Deployment Instructions.md when logic under api endpoints
    is implemented.
*/

func queryPods(w http.ResponseWriter, r *http.Request) {
	pods, err := clientsetSingleton.QueryPods("default")
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	fmt.Fprintf(w, "Pods in default namespace:\n")

	for _, pod := range pods {
		fmt.Fprintf(w, "- %s\n", pod)
	}

	fmt.Fprintf(w, "\n There are %d pods in the cluster\n", len(pods))
}

func queryDeployments(w http.ResponseWriter, r *http.Request) {
	deployments, err := clientsetSingleton.QueryDeployments("default")
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	fmt.Fprintf(w, "Deployments in default namespace:\n")

	for _, deployment := range deployments {
		fmt.Fprintf(w, "- %s\n", deployment)
	}

	fmt.Fprintf(w, "\nThere are %d deployments in the cluster\n", len(deployments))
}

func queryServices(w http.ResponseWriter, r *http.Request) {
	service, err := clientsetSingleton.QueryServices("default")
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	fmt.Fprintf(w, "Services in default namespace:\n")

	for _, service := range service {
		fmt.Fprintf(w, "- %s\n", service)
	}

	fmt.Fprintf(w, "\nThere are %d services in the cluster\n", len(service))
}
