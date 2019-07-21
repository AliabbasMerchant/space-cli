package utils

// BuildVersion is the current version of Space Cli
const BuildVersion = "0.1.1"

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

const (
	// Python3Image is the type used for Python3Image
	Python3Image string = "spaceuptech/runtime-python:latest"

	// Java11Image is the type used for Java11Image
	Java11Image string = "spaceuptech/runtime-java:latest"

	// GolangImage is the type used for GolangImage
	GolangImage string = "spaceuptech/runtime-golang:latest"

	// NodeJSImage is the type used for NodeJSImage
	NodeJSImage string = "spaceuptech/runtime-node:latest"

	// WebAppImage is the type used for WebAppImage
	WebAppImage string = "spaceuptech/runtime-web-app:latest"
)

const (
	// KindService is the kind used for a Service
	KindService string = "service"

	// KindWebApp is the kind used for a WebApp
	KindWebApp string = "web-app"
)

// DefaultConfigFilePath is the default path to load / store the config file
const DefaultConfigFilePath string = "deployer.yaml"

// ZipName is the name of the temporary zip file created
const ZipName string = "SPACEzip.zip"
