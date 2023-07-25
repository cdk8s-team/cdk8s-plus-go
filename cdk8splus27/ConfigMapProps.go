package cdk8splus27

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Properties for initialization of `ConfigMap`.
type ConfigMapProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// BinaryData contains the binary data.
	//
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	// BinaryData can contain byte sequences that are not in the UTF-8 range. The
	// keys stored in BinaryData must not overlap with the ones in the Data field,
	// this is enforced during validation process.
	//
	// You can also add binary data using `configMap.addBinaryData()`.
	BinaryData *map[string]*string `field:"optional" json:"binaryData" yaml:"binaryData"`
	// Data contains the configuration data.
	//
	// Each key must consist of alphanumeric characters, '-', '_' or '.'. Values
	// with non-UTF-8 byte sequences must use the BinaryData field. The keys
	// stored in Data must not overlap with the keys in the BinaryData field, this
	// is enforced during validation process.
	//
	// You can also add data using `configMap.addData()`.
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	// If set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified).
	//
	// If not set to true, the field can be modified at any time.
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
}

