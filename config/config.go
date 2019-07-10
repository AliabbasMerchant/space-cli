package config

// Config holds the entire configuration
type Config struct {
	Name        string            `json:"name" yaml:"name"`
	Project     string            `json:"project" yaml:"project"`
	WorkingDir  string            `json:"workingDir" yaml:"workingDir"`
	Ignore      string            `json:"ignore" yaml:"ignore"`
	Runtime     *Runtime          `json:"runtime" yaml:"runtime"`
	Constraints *Constraints      `json:"constraints" yaml:"constraints"`
	Ports       []string         `json:"ports" yaml:"ports"`
	Env         map[string]string `json:"env" yaml:"env"`
	Clusters    map[string]string `json:"clusters" yaml:"clusters"`
}

// Runtime holds the runtime information
type Runtime struct {
	Name    string `json:"name" yaml:"name"`
	Install string `json:"install" yaml:"install"`
	Run     string `json:"run" yaml:"run"`
}

// Constraints holds the constraints information
type Constraints struct {
	Replicas int      `json:"replicas" yaml:"replicas"`
	CPU      float32  `json:"cpu" yaml:"cpu"`
	Memory   string   `json:"memory" yaml:"memory"`
}
