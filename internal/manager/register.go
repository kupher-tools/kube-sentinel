package manager

import (
	"github.com/kupher-tools/kube-sentinel/internal/handler"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func Register(mgr manager.Manager) error {

	ctrl.Log.Info("Registering Kuper's kube-sentinel : A mutating admission controller")

	server := webhook.NewServer(webhook.Options{
		Port:    443,
		CertDir: "test",
	})

	server.Register("/kube-sentinel", &admission.Webhook{
		Handler: &(handler.KubeSentinel{}),
	})
	return nil

}
