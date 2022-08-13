// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Options for `Handler.fromTcpSocket`.
type HandlerFromTcpSocketOptions struct {
	// The host name to connect to on the container.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The TCP port to connect to on the container.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

