package handler

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

func readSecurityContextConfig(ctx context.Context, c client.Client, namespace, name string) (*PodSecurityContext, *ContainerSecurityContext, error) {
	cm := &corev1.ConfigMap{}
	err := c.Get(ctx, client.ObjectKey{Namespace: namespace, Name: name}, cm)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get ConfigMap: %w", err)
	}

	podSecCtxYaml, ok := cm.Data["podSecurityContext"]
	if !ok {
		return nil, nil, fmt.Errorf("podSecurityContext key missing in ConfigMap")
	}
	containerSecCtxYaml, ok := cm.Data["containerSecurityContext"]
	if !ok {
		return nil, nil, fmt.Errorf("containerSecurityContext key missing in ConfigMap")
	}

	var podSecCtx PodSecurityContext
	if err := yaml.Unmarshal([]byte(podSecCtxYaml), &podSecCtx); err != nil {
		return nil, nil, fmt.Errorf("failed to parse podSecurityContext yaml: %w", err)
	}

	var containerSecCtx ContainerSecurityContext
	if err := yaml.Unmarshal([]byte(containerSecCtxYaml), &containerSecCtx); err != nil {
		return nil, nil, fmt.Errorf("failed to parse containerSecurityContext yaml: %w", err)
	}

	return &podSecCtx, &containerSecCtx, nil
}
