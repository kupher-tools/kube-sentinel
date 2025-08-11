package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kupher-tools/kube-sentinel/internal/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type KubeSentinel struct {
	//decoder Decoder
}

func (m *KubeSentinel) Handle(ctx context.Context, req admission.Request) admission.Response {
	//var decoder admission.Decoder
	decoder := admission.NewDecoder(scheme.Scheme)
	var pod corev1.Pod

	if decoder == nil {
		ctrl.Log.Error(nil, "Decoder not injected")
		return admission.Errored(500, fmt.Errorf("decoder not injected"))
	}

	if err := decoder.Decode(req, &pod); err != nil {
		ctrl.Log.Error(err, "Failed to decode pod")
		return admission.Errored(400, err)
	}

	ctrl.Log.Info("Admission review for pod: %s/%s", pod.Namespace, pod.Name)

	client, err := utils.GetKubeClient()
	if err != nil {
		ctrl.Log.Error(err, "Failed to get Kubernetes client")
		return admission.Errored(400, err)
	}

	podSecCtx, containerSecCtx, err := readSecurityContextConfig(ctx, client, pod.Namespace, "security-context-config")
	if err != nil {
		ctrl.Log.Error(err, "Failed to get Kubernetes client")
		return admission.Errored(500, err)
	}

	// Now you have podSecCtx and containerSecCtx structs ready to use in your patch logic

	// Mutate container security context
	for i := range pod.Spec.Containers {
		if pod.Spec.Containers[i].SecurityContext == nil {
			pod.Spec.Containers[i].SecurityContext = &corev1.SecurityContext{}
		}
		setContainerSecurityContext(containerSecCtx, pod.Spec.Containers[i].SecurityContext)
		ctrl.Log.Info("Mutated container security context for container", "container", pod.Spec.Containers[i].SecurityContext)
	}

	if pod.Spec.SecurityContext == nil {
		pod.Spec.SecurityContext = &corev1.PodSecurityContext{}
	}

	setPodSecurityContext(podSecCtx, pod.Spec.SecurityContext)

	// Return the mutated pod object as a patch
	marshaledPod, err := json.Marshal(pod)
	if err != nil {
		ctrl.Log.Error(err, "Failed to marshal mutated pod")
		return admission.Errored(500, err)
	}

	ctrl.Log.Info("Pod mutated successfully")
	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledPod)
}

/*
func (m *KubeSentinel) InjectDecoder(d *admission.Decoder) error {
	m.decoder = d
	return nil
}*/

// Helper to avoid nil dereference
func setBoolIfNil(field **bool, val bool) {
	if *field == nil {
		*field = &val
	}
}

func setContainerSecurityContext(src *ContainerSecurityContext, dst *corev1.SecurityContext) {
	if src == nil {
		return
	}

	if dst == nil {
		dst = &corev1.SecurityContext{}
	}

	if dst.RunAsNonRoot == nil {
		dst.RunAsNonRoot = new(bool)
		*dst.RunAsNonRoot = src.RunAsNonRoot
	}

	if dst.AllowPrivilegeEscalation == nil {
		dst.AllowPrivilegeEscalation = new(bool)
		*dst.AllowPrivilegeEscalation = src.AllowPrivilegeEscalation
	}

	if dst.ReadOnlyRootFilesystem == nil {
		dst.ReadOnlyRootFilesystem = new(bool)
		*dst.ReadOnlyRootFilesystem = src.ReadOnlyRootFilesystem
	}

	if dst.Capabilities == nil && (len(src.Capabilities.Add) > 0 || len(src.Capabilities.Drop) > 0) {
		dst.Capabilities = &corev1.Capabilities{
			Add:  convertStringSliceToCapabilitySlice(src.Capabilities.Add),
			Drop: convertStringSliceToCapabilitySlice(src.Capabilities.Drop),
		}
	}

	if dst.SeccompProfile == nil && src.SeccompProfile.Type != "" {
		dst.SeccompProfile = &corev1.SeccompProfile{
			Type: corev1.SeccompProfileType(src.SeccompProfile.Type),
		}
	}
}

func setPodSecurityContext(src *PodSecurityContext, dst *corev1.PodSecurityContext) {
	if src == nil {
		return
	}
	if dst == nil {
		dst = &corev1.PodSecurityContext{}
	}

	if dst.RunAsUser == nil && src.RunAsUser != nil {
		dst.RunAsUser = new(int64)
		*dst.RunAsUser = *src.RunAsUser
	}

	if dst.RunAsGroup == nil && src.RunAsGroup != nil {
		dst.RunAsGroup = new(int64)
		*dst.RunAsGroup = *src.RunAsGroup
	}

	if dst.FSGroup == nil && src.FSGroup != nil {
		dst.FSGroup = new(int64)
		*dst.FSGroup = *src.FSGroup
	}

	if dst.RunAsNonRoot == nil {
		dst.RunAsNonRoot = new(bool)
		*dst.RunAsNonRoot = src.RunAsNonRoot
	}

	if dst.SeccompProfile == nil && src.SeccompProfile.Type != "" {
		dst.SeccompProfile = &corev1.SeccompProfile{
			Type: corev1.SeccompProfileType(src.SeccompProfile.Type),
		}
	}
}

func convertStringSliceToCapabilitySlice(strs []string) []corev1.Capability {
	caps := make([]corev1.Capability, len(strs))
	for i, s := range strs {
		caps[i] = corev1.Capability(s)
	}
	return caps
}
