package main

import (
	"fmt"
	"net/http"

	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var clientset *kubernetes.Clientset

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Starting server on port 8080...")

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Server is running")
	}))

	mux.HandleFunc("/pods", queryPods)

	mux.HandleFunc("/deployments", queryDeployments)

	mux.HandleFunc("/services", queryServices)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}

func queryPods(w http.ResponseWriter, r *http.Request) {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			fmt.Fprintf(w, "Pods not found")
		} else {
			fmt.Fprintf(w, "Error: %s", err.Error())
		}
		return
	}

	fmt.Fprintf(w, "Pods:\n")

	for _, pod := range pods.Items {
		fmt.Fprintf(w, "- %s\n", pod.Name)
	}

	fmt.Fprintf(w, "\n There are %d pods in the cluster\n", len(pods.Items))
}

func queryDeployments(w http.ResponseWriter, r *http.Request) {
	deployments, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			fmt.Fprintf(w, "Deployments not found")
		} else {
			fmt.Fprintf(w, "Error: %s", err.Error())
		}
		return
	}

	fmt.Fprintf(w, "Deployments:\n")

	for _, deployment := range deployments.Items {
		fmt.Fprintf(w, "- %s\n", deployment.Name)
	}

	fmt.Fprintf(w, "\nThere are %d deployments in the cluster\n", len(deployments.Items))
}

func queryServices(w http.ResponseWriter, r *http.Request) {
	service, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			fmt.Fprintf(w, "Services not found")
		} else {
			fmt.Fprintf(w, "Error: %s", err.Error())
		}
		return
	}

	fmt.Fprintf(w, "Services:\n")

	for _, service := range service.Items {
		fmt.Fprintf(w, "- %s\n", service.Name)
	}

	fmt.Fprintf(w, "\nThere are %d services in the cluster\n", len(service.Items))
}
