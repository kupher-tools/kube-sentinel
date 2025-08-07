package manager

import (
	"fmt"

	"github.com/kupher-tools/kube-sentinel/internal/utils"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

func NewManager() (manager.Manager, error) {

	cfg, err := utils.GetKubeConfig()
	if err != nil {
		ctrl.Log.Error(err, "Error getting kube config")
		return nil, err
	}

	options := ctrl.Options{
		LeaderElection: false,
		Metrics: server.Options{
			BindAddress: "0", //disables metrics
		},
	}

	mgr, err := ctrl.NewManager(cfg, options)
	if err != nil {
		ctrl.Log.Error(err, "Failed to create controller manager")
		return nil, fmt.Errorf("failed to create controller manager: %w", err)
	}
	ctrl.Log.Info("Controller Manager created successfully")
	return mgr, nil
}
