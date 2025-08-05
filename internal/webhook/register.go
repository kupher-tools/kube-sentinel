package webhook

import (
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func Register(mgr manager.Manager) error {

	klog.Info("Registering Kuper's kube-sentinel : A mutating admission controller")

	server := webhook.NewServer(webhook.Options{
		Port:    443,
		CertDir: "test",
	})

	server.Register("/kube-sentinel", &admission.Webhook{
		Handler: &KubeSentinel{},
	})
	return nil

}
