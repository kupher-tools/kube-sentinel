package utils

import (
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetKubeClient() (client.Client, error) {
	cfg, err := GetKubeConfig()
	if err != nil {
		ctrl.Log.Error(err, "Error getting kube config for client")
		return nil, err
	}
	c, err := client.New(cfg, client.Options{})
	if err != nil {
		ctrl.Log.Error(err, "Error creating Kubernetes client")
		return nil, err
	}
	return c, nil
}
