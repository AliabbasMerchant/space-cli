package utils

// BuildVersion is the current version of Space Cli
const BuildVersion = "0.1.0"

const (
	// Python3 is the type used for Python3
	Python3 string = "Python3"

	// Java11 is the type used for Java11
	Java11 string = "Java11"

	// Golang is the type used for Golang
	Golang string = "Golang"

	// NodeJS is the type used for NodeJS
	NodeJS string = "NodeJS"
)

// DefaultConfigFilePath is the default path to load / store the config file
const DefaultConfigFilePath string = "deployer.yaml"

// ZipName is the name of the temporary zip file created
const ZipName string = "SPACEzip.zip"

const Temporary string = "/v1/api/deploy"
