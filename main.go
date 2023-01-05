package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {

	var kubeconfig *string
	//1) Use the flag package to build the location to the local kubeconfig
	home := homedir.HomeDir()
	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "absolute path to kubeconfig")
	flag.Parse()
	//fmt.Println(*kubeconfig)

	//2) Create a config using the kubeconfig path
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error)
	}

	//3) Now, interact with K8S cluster. Create clientset using the config
	// Q1: Why do we need the config variable to create the clienset, why no directly from the kubeconfig path ?
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error)
	}

	//4) List all pods in the cluster
	// Q2: What is the context package - congext.Background() and the v1.ListOptions magic ?
	// Q3: The clienset has access to the CoreV1 group (GVK) and the other groups ?
	podlist, err := clientset.CoreV1().Pods("").List(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err.Error)
	}
	// fmt.Println(podlist)

	for _, pod := range podlist.Items {
		fmt.Printf("Pod Name : %s \t\t| Namespace : %s \n", pod.Name, pod.Namespace)
	}
}
