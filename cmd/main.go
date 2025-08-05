package main

import (
	"flag"
	"os"

	"github.com/kupher-tools/kube-sentinel/internal/manager"
	"github.com/kupher-tools/kube-sentinel/internal/webhook"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	klog.InitFlags(nil)
	flag.Set("v", "2") // Default verbosity
	flag.Parse()
	defer klog.Flush()

	klog.Info("Starting kube-sentinel webhook server")

	mgr, err := manager.NewManager("/tmp/k8s-webhook-server/serving-certs", 9443)
	if err != nil {
		//Error out and exit
		klog.Fatalf("Unable to initialize manager : %v", err)
	}

	err = webhook.Register(mgr)
	if err != nil {
		klog.Fatalf("Unable to register webhook : %v", err)
	}

	klog.Info("kube-sentinel is ready and listening on /kube-sentinel")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		klog.Fatalf("Manager stopped with error: %v", err)
		os.Exit(1)
	}
}
