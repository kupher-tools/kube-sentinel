package handler

import (
	"context"
	"encoding/json"

	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type KubeSentinel struct {
	decoder admission.Decoder
}

func (m *KubeSentinel) Handle(_ context.Context, req admission.Request) admission.Response {
	pod := &corev1.Pod{}

	if err := m.decoder.Decode(req, pod); err != nil {
		ctrl.Log.Error(err, "Failed to decode pod")
		return admission.Errored(400, err)
	}

	ctrl.Log.Info("Admission review for pod: %s/%s", pod.Namespace, pod.Name)

	// Mutate container security context
	for i := range pod.Spec.Containers {
		c := &pod.Spec.Containers[i]
		if c.SecurityContext == nil {
			c.SecurityContext = &corev1.SecurityContext{}
		}
		setBoolIfNil(&c.SecurityContext.RunAsNonRoot, true)
		setBoolIfNil(&c.SecurityContext.AllowPrivilegeEscalation, false)
		setBoolIfNil(&c.SecurityContext.ReadOnlyRootFilesystem, true)
	}

	// Mutate pod-level security context
	if pod.Spec.SecurityContext == nil {
		pod.Spec.SecurityContext = &corev1.PodSecurityContext{}
	}
	if pod.Spec.SecurityContext.SeccompProfile == nil {
		pod.Spec.SecurityContext.SeccompProfile = &corev1.SeccompProfile{
			Type: corev1.SeccompProfileTypeRuntimeDefault,
		}
	}

	// Return the mutated pod object as a patch
	marshaledPod, err := json.Marshal(pod)
	if err != nil {
		ctrl.Log.Error(err, "Failed to marshal mutated pod")
		return admission.Errored(500, err)
	}

	ctrl.Log.Info("Pod mutated successfully")
	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledPod)
}

func (m *KubeSentinel) InjectDecoder(d admission.Decoder) error {
	m.decoder = d
	return nil
}

// Helper to avoid nil dereference
func setBoolIfNil(field **bool, val bool) {
	if *field == nil {
		*field = &val
	}
}
