package utils

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	ctrl "sigs.k8s.io/controller-runtime"
)

func GetKubeConfig() (*rest.Config, error) {
	cfg, err := rest.InClusterConfig()

	if err == nil {
		return cfg, nil
	}

	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	if env := os.Getenv("KUBECONFIG"); env != "" {
		kubeconfig = env
	}

	ctrl.Log.Info("Using out-of-cluster config", "kubeconfig", kubeconfig)

	cfg, err = clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		ctrl.Log.Error(err, "Error loading kube config file")
		return nil, err
	}

	return cfg, nil
}
