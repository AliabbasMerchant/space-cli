package model

// GlobalConfig stores the global config
type GlobalConfig struct {
	Clusters Clusters `json:"clusters"`
}

// Clusters is a map of clusters
type Clusters map[string]*Cluster // cluster_name -> cluster

// Cluster is an entity holding a single cluster
type Cluster struct {
	URL   string `json:"url"`
	Token string `json:"token,omitempty"`
}
