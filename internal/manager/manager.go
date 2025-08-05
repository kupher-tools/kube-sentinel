package manager

import (
	"fmt"

	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func NewManager(certDir string, port int) (manager.Manager, error) {

	cfg := ctrl.GetConfigOrDie()

	options := manager.Options{}

	mgr, err := ctrl.NewManager(cfg, options)
	if err != nil {
		klog.Errorf("Failed to create controller manager: %v", err)
		return nil, fmt.Errorf("failed to create controller manager: %w", err)
	}
	klog.Info("Controller Manager created successfully")
	return mgr, nil
}
