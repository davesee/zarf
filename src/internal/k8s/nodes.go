package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetNodes returns a list of nodes from the k8s cluster.
func GetNodes() (*corev1.NodeList, error) {
	clientset, err := getClientset()
	if err != nil {
		return nil, err
	}

	metaOptions := metav1.ListOptions{}
	return clientset.CoreV1().Nodes().List(context.TODO(), metaOptions)
}
