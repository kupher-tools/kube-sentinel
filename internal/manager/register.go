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
		Port:    8443,
		CertDir: "/tls",
	})

	sentinel := &handler.KubeSentinel{}

	server.Register("/kube-sentinel", &admission.Webhook{
		Handler: sentinel,
	})

	if err := mgr.Add(server); err != nil {
		ctrl.Log.Error(err, "Unable to add webhook server to manager:")
		return err
	}
	ctrl.Log.Info("Server Registerted to manager successfully")
	return nil

}
