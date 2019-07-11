package model

// Credentials stores the space cloud credentials
type Credentials struct {
	User string `json:"user,omitempty" yaml:"user,omitempty"`
	Pass string `json:"pass,omitempty" yaml:"pass,omitempty"`
}
