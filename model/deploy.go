package model

// Deploy holds the config for the deployment
type Deploy struct {
	// Deployment specific config
	Name        string            `json:"name" yaml:"name"`
	Project     string            `json:"project" yaml:"project"`
	Kind        string            `json:"kind" yaml:"kind"`
	Runtime     *Runtime          `json:"runtime" yaml:"runtime"`
	Env         map[string]string `json:"env" yaml:"env"`
	Constraints *Constraints      `json:"constraints" yaml:"constraints"`
	Ports       []*Port           `json:"ports,omitempty" yaml:"ports,omitempty"`
	Expose      []*Expose         `json:"expose,omitempty" yaml:"expose,omitempty"`

	// CLI specific config
	WorkingDir string            `json:"workingDir" yaml:"workingDir"`
	Ignore     string            `json:"ignore" yaml:"ignore"`
	Clusters   map[string]string `json:"clusters" yaml:"clusters"`
}

// Runtime holds the runtime information
type Runtime struct {
	Name    string `json:"name" yaml:"name"`
	Install string `json:"install" yaml:"install"`
	Run     string `json:"run" yaml:"run"`
}

// Constraints holds the constraints information
type Constraints struct {
	Replicas int32    `json:"replicas" yaml:"replicas"`
	CPU      *float32 `json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Memory   *int64   `json:"memory,omitempty" yaml:"memory,omitempty"`
}

// Port holds the container ports information
type Port struct {
	Name     *string `json:"name,omitempty" yaml:"name,omitempty"`
	Port     int32   `json:"port" yaml:"port"`
	Protocol *string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}

// Expose holds the information about the ports to expose
type Expose struct {
	Prefix string `json:"prefix" yaml:"prefix"`
	Host   string `json:"host" yaml:"host"`
	Proxy  string `json:"proxy" yaml:"proxy"`
}
