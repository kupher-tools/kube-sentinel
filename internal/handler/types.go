package handler

type SeccompProfile struct {
	Type string `yaml:"type"`
}

type PodSecurityContext struct {
	RunAsUser      *int64         `yaml:"runAsUser,omitempty"`
	RunAsGroup     *int64         `yaml:"runAsGroup,omitempty"`
	FSGroup        *int64         `yaml:"fsGroup,omitempty"`
	RunAsNonRoot   bool           `yaml:"runAsNonRoot"`
	SeccompProfile SeccompProfile `yaml:"seccompProfile"`
}

type Capabilities struct {
	Add  []string `yaml:"add,omitempty"`
	Drop []string `yaml:"drop,omitempty"`
}

type ContainerSecurityContext struct {
	AllowPrivilegeEscalation bool           `yaml:"allowPrivilegeEscalation"`
	ReadOnlyRootFilesystem   bool           `yaml:"readOnlyRootFilesystem"`
	RunAsNonRoot             bool           `yaml:"runAsNonRoot"`
	Capabilities             Capabilities   `yaml:"capabilities,omitempty"`
	SeccompProfile           SeccompProfile `yaml:"seccompProfile"`
}
