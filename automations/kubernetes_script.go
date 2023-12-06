package main

import (
	"flag"
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/wait"
	"k8s.io/client-go/util/wait/waitgroup"
	"k8s.io/client-go/util/yaml"
)

func main() {
	var kubeconfig *string
	var namespacesFile *string
	var outputFilePath *string

	home := homedir.HomeDir()
	kubeconfig = flag.String("kubeconfig", home+"/.kube/config", "absolute path to the kubeconfig file")
	namespacesFile = flag.String("namespaces-file", "namespaces.txt", "file containing the list of namespaces")
	outputFilePath = flag.String("output-file", "deployments.yaml", "file to dump deployments information")

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error building kubeconfig: %v\n", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %v\n", err)
		os.Exit(1)
	}

	namespaces, err := readNamespacesFromFile(*namespacesFile)
	if err != nil {
		fmt.Printf("Error reading namespaces from file: %v\n", err)
		os.Exit(1)
	}

	err = getDeploymentsInNamespaces(clientset, namespaces, *outputFilePath)
	if err != nil {
		fmt.Printf("Error getting deployments: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Deployments information written to %s\n", *outputFilePath)
}

func readNamespacesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var namespaces []string
	decoder := yaml.NewYAMLToJSONDecoder(file)
	for {
		var namespace string
		if err := decoder.Decode(&namespace); err != nil {
			break
		}
		namespaces = append(namespaces, namespace)
	}

	return namespaces, nil
}

func getDeploymentsInNamespaces(clientset *kubernetes.Clientset, namespaces []string, outputFilePath string) error {
	var wg waitgroup.Group

	for _, namespace := range namespaces {
		namespace := namespace
		wg.StartWithChannel(func() {
			defer wg.Done()
			deployments, err := clientset.AppsV1().Deployments(namespace).List(wait.NeverStop)
			if err != nil {
				fmt.Printf("Error getting deployments in namespace %s: %v\n", namespace, err)
				return
			}

			file, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				fmt.Printf("Error opening file %s: %v\n", outputFilePath, err)
				return
			}
			defer file.Close()

			for _, deployment := range deployments.Items {
				// Convert deployment to YAML
				deploymentYAML, err := yaml.Marshal(deployment)
				if err != nil {
					fmt.Printf("Error marshaling deployment to YAML: %v\n", err)
					return
				}

				// Write deployment YAML to file
				_, err = file.Write(deploymentYAML)
				if err != nil {
					fmt.Printf("Error writing deployment to file: %v\n", err)
					return
				}

				// Add separator between deployments
				_, err = file.WriteString("---\n")
				if err != nil {
					fmt.Printf("Error writing separator to file: %v\n", err)
					return
				}
			}
		})
	}

	wg.Wait()
	return nil
}
