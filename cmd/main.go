package main

import (
	"github.com/kupher-tools/kube-sentinel/internal/manager"
	"github.com/kupher-tools/kube-sentinel/internal/utils"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {

	utils.InitLogger()
	ctrl.Log.Info("Starting kube-sentinel webhook server")

	mgr, err := manager.NewManager()
	if err != nil {
		//Error out and exit
		ctrl.Log.Error(err, "Unable to initialize manager ")
	}

	err = manager.Register(mgr)
	if err != nil {
		ctrl.Log.Error(err, "Unable to register webhook ")
	}

	ctrl.Log.Info("kube-sentinel is ready and listening on /kube-sentinel")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		ctrl.Log.Error(err, "Manager stopped with error")

	}
}
