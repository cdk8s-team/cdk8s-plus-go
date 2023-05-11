package cdk8splus25

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus25/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// ConfigMap holds configuration data for pods to consume.
type ConfigMap interface {
	Resource
	IConfigMap
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// The binary data associated with this config map.
	//
	// Returns a copy. To add data records, use `addBinaryData()` or `addData()`.
	BinaryData() *map[string]*string
	// The data associated with this config map.
	//
	// Returns an copy. To add data records, use `addData()` or `addBinaryData()`.
	Data() *map[string]*string
	// Whether or not this config map is immutable.
	Immutable() *bool
	// The object kind (e.g. "Deployment").
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// The tree node.
	Node() constructs.Node
	Permissions() ResourcePermissions
	// The unique, namespace-global, name of an object inside the Kubernetes cluster.
	//
	// If this is omitted, the ApiResource should represent all objects of the given type.
	ResourceName() *string
	// The name of a resource type as it appears in the relevant API endpoint.
	ResourceType() *string
	// Adds a binary data entry to the config map.
	//
	// BinaryData can contain byte
	// sequences that are not in the UTF-8 range.
	AddBinaryData(key *string, value *string)
	// Adds a data entry to the config map.
	AddData(key *string, value *string)
	// Adds a directory to the ConfigMap.
	AddDirectory(localDir *string, options *AddDirectoryOptions)
	// Adds a file to the ConfigMap.
	AddFile(localFile *string, key *string)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ConfigMap
type jsiiProxy_ConfigMap struct {
	jsiiProxy_Resource
	jsiiProxy_IConfigMap
}

func (j *jsiiProxy_ConfigMap) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) BinaryData() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"binaryData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) Data() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) Immutable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"immutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewConfigMap(scope constructs.Construct, id *string, props *ConfigMapProps) ConfigMap {
	_init_.Initialize()

	if err := validateNewConfigMapParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigMap{}

	_jsii_.Create(
		"cdk8s-plus-25.ConfigMap",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewConfigMap_Override(c ConfigMap, scope constructs.Construct, id *string, props *ConfigMapProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-25.ConfigMap",
		[]interface{}{scope, id, props},
		c,
	)
}

// Represents a ConfigMap created elsewhere.
func ConfigMap_FromConfigMapName(scope constructs.Construct, id *string, name *string) IConfigMap {
	_init_.Initialize()

	if err := validateConfigMap_FromConfigMapNameParameters(scope, id, name); err != nil {
		panic(err)
	}
	var returns IConfigMap

	_jsii_.StaticInvoke(
		"cdk8s-plus-25.ConfigMap",
		"fromConfigMapName",
		[]interface{}{scope, id, name},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ConfigMap_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConfigMap_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-25.ConfigMap",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigMap) AddBinaryData(key *string, value *string) {
	if err := c.validateAddBinaryDataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addBinaryData",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_ConfigMap) AddData(key *string, value *string) {
	if err := c.validateAddDataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addData",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_ConfigMap) AddDirectory(localDir *string, options *AddDirectoryOptions) {
	if err := c.validateAddDirectoryParameters(localDir, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDirectory",
		[]interface{}{localDir, options},
	)
}

func (c *jsiiProxy_ConfigMap) AddFile(localFile *string, key *string) {
	if err := c.validateAddFileParameters(localFile); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addFile",
		[]interface{}{localFile, key},
	)
}

func (c *jsiiProxy_ConfigMap) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		c,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigMap) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

