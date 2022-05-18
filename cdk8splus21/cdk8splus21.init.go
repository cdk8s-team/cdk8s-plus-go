package cdk8splus21

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk8s-plus-21.AbstractPod",
		reflect.TypeOf((*AbstractPod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addHostAlias", GoMethod: "AddHostAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_AbstractPod{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.AbstractPodProps",
		reflect.TypeOf((*AbstractPodProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.AddDeploymentOptions",
		reflect.TypeOf((*AddDeploymentOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.AddDirectoryOptions",
		reflect.TypeOf((*AddDirectoryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.ApiResource",
		reflect.TypeOf((*ApiResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
		},
		func() interface{} {
			j := jsiiProxy_ApiResource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApiEndpoint)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApiResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ApiResourceOptions",
		reflect.TypeOf((*ApiResourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.AwsElasticBlockStorePersistentVolume",
		reflect.TypeOf((*AwsElasticBlockStorePersistentVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessModes", GoGetter: "AccessModes"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asVolume", GoMethod: "AsVolume"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "claim", GoGetter: "Claim"},
			_jsii_.MemberProperty{JsiiProperty: "fsType", GoGetter: "FsType"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
			_jsii_.MemberProperty{JsiiProperty: "mountOptions", GoGetter: "MountOptions"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "partition", GoGetter: "Partition"},
			_jsii_.MemberProperty{JsiiProperty: "readOnly", GoGetter: "ReadOnly"},
			_jsii_.MemberProperty{JsiiProperty: "reclaimPolicy", GoGetter: "ReclaimPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "reserve", GoMethod: "Reserve"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "storage", GoGetter: "Storage"},
			_jsii_.MemberProperty{JsiiProperty: "storageClassName", GoGetter: "StorageClassName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumeId", GoGetter: "VolumeId"},
		},
		func() interface{} {
			j := jsiiProxy_AwsElasticBlockStorePersistentVolume{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PersistentVolume)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.AwsElasticBlockStorePersistentVolumeProps",
		reflect.TypeOf((*AwsElasticBlockStorePersistentVolumeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.AwsElasticBlockStoreVolumeOptions",
		reflect.TypeOf((*AwsElasticBlockStoreVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.AzureDiskPersistentVolume",
		reflect.TypeOf((*AzureDiskPersistentVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessModes", GoGetter: "AccessModes"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asVolume", GoMethod: "AsVolume"},
			_jsii_.MemberProperty{JsiiProperty: "azureKind", GoGetter: "AzureKind"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "cachingMode", GoGetter: "CachingMode"},
			_jsii_.MemberProperty{JsiiProperty: "claim", GoGetter: "Claim"},
			_jsii_.MemberProperty{JsiiProperty: "diskName", GoGetter: "DiskName"},
			_jsii_.MemberProperty{JsiiProperty: "diskUri", GoGetter: "DiskUri"},
			_jsii_.MemberProperty{JsiiProperty: "fsType", GoGetter: "FsType"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
			_jsii_.MemberProperty{JsiiProperty: "mountOptions", GoGetter: "MountOptions"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "readOnly", GoGetter: "ReadOnly"},
			_jsii_.MemberProperty{JsiiProperty: "reclaimPolicy", GoGetter: "ReclaimPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "reserve", GoMethod: "Reserve"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "storage", GoGetter: "Storage"},
			_jsii_.MemberProperty{JsiiProperty: "storageClassName", GoGetter: "StorageClassName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AzureDiskPersistentVolume{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PersistentVolume)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.AzureDiskPersistentVolumeCachingMode",
		reflect.TypeOf((*AzureDiskPersistentVolumeCachingMode)(nil)).Elem(),
		map[string]interface{}{
			"NONE": AzureDiskPersistentVolumeCachingMode_NONE,
			"READ_ONLY": AzureDiskPersistentVolumeCachingMode_READ_ONLY,
			"READ_WRITE": AzureDiskPersistentVolumeCachingMode_READ_WRITE,
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.AzureDiskPersistentVolumeKind",
		reflect.TypeOf((*AzureDiskPersistentVolumeKind)(nil)).Elem(),
		map[string]interface{}{
			"SHARED": AzureDiskPersistentVolumeKind_SHARED,
			"DEDICATED": AzureDiskPersistentVolumeKind_DEDICATED,
			"MANAGED": AzureDiskPersistentVolumeKind_MANAGED,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.AzureDiskPersistentVolumeProps",
		reflect.TypeOf((*AzureDiskPersistentVolumeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.AzureDiskVolumeOptions",
		reflect.TypeOf((*AzureDiskVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.BasicAuthSecret",
		reflect.TypeOf((*BasicAuthSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStringData", GoMethod: "AddStringData"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "getStringData", GoMethod: "GetStringData"},
			_jsii_.MemberProperty{JsiiProperty: "immutable", GoGetter: "Immutable"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BasicAuthSecret{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Secret)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.BasicAuthSecretProps",
		reflect.TypeOf((*BasicAuthSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.ClusterRole",
		reflect.TypeOf((*ClusterRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "aggregate", GoMethod: "Aggregate"},
			_jsii_.MemberMethod{JsiiMethod: "allow", GoMethod: "Allow"},
			_jsii_.MemberMethod{JsiiMethod: "allowCreate", GoMethod: "AllowCreate"},
			_jsii_.MemberMethod{JsiiMethod: "allowDelete", GoMethod: "AllowDelete"},
			_jsii_.MemberMethod{JsiiMethod: "allowDeleteCollection", GoMethod: "AllowDeleteCollection"},
			_jsii_.MemberMethod{JsiiMethod: "allowGet", GoMethod: "AllowGet"},
			_jsii_.MemberMethod{JsiiMethod: "allowList", GoMethod: "AllowList"},
			_jsii_.MemberMethod{JsiiMethod: "allowPatch", GoMethod: "AllowPatch"},
			_jsii_.MemberMethod{JsiiMethod: "allowRead", GoMethod: "AllowRead"},
			_jsii_.MemberMethod{JsiiMethod: "allowReadWrite", GoMethod: "AllowReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "allowUpdate", GoMethod: "AllowUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "allowWatch", GoMethod: "AllowWatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindInNamespace", GoMethod: "BindInNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "combine", GoMethod: "Combine"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterRole{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IClusterRole)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRole)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.ClusterRoleBinding",
		reflect.TypeOf((*ClusterRoleBinding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSubjects", GoMethod: "AddSubjects"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "subjects", GoGetter: "Subjects"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterRoleBinding{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ClusterRoleBindingProps",
		reflect.TypeOf((*ClusterRoleBindingProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ClusterRolePolicyRule",
		reflect.TypeOf((*ClusterRolePolicyRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ClusterRoleProps",
		reflect.TypeOf((*ClusterRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.CommandProbeOptions",
		reflect.TypeOf((*CommandProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.CommonSecretProps",
		reflect.TypeOf((*CommonSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.ConfigMap",
		reflect.TypeOf((*ConfigMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBinaryData", GoMethod: "AddBinaryData"},
			_jsii_.MemberMethod{JsiiMethod: "addData", GoMethod: "AddData"},
			_jsii_.MemberMethod{JsiiMethod: "addDirectory", GoMethod: "AddDirectory"},
			_jsii_.MemberMethod{JsiiMethod: "addFile", GoMethod: "AddFile"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "binaryData", GoGetter: "BinaryData"},
			_jsii_.MemberProperty{JsiiProperty: "data", GoGetter: "Data"},
			_jsii_.MemberProperty{JsiiProperty: "immutable", GoGetter: "Immutable"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConfigMap{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConfigMap)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ConfigMapProps",
		reflect.TypeOf((*ConfigMapProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ConfigMapVolumeOptions",
		reflect.TypeOf((*ConfigMapVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Container",
		reflect.TypeOf((*Container)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "args", GoGetter: "Args"},
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "imagePullPolicy", GoGetter: "ImagePullPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "mount", GoMethod: "Mount"},
			_jsii_.MemberProperty{JsiiProperty: "mounts", GoGetter: "Mounts"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "resources", GoGetter: "Resources"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberProperty{JsiiProperty: "workingDir", GoGetter: "WorkingDir"},
		},
		func() interface{} {
			return &jsiiProxy_Container{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ContainerLifecycle",
		reflect.TypeOf((*ContainerLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ContainerProps",
		reflect.TypeOf((*ContainerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ContainerResources",
		reflect.TypeOf((*ContainerResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.ContainerSecurityContext",
		reflect.TypeOf((*ContainerSecurityContext)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ensureNonRoot", GoGetter: "EnsureNonRoot"},
			_jsii_.MemberProperty{JsiiProperty: "group", GoGetter: "Group"},
			_jsii_.MemberProperty{JsiiProperty: "privileged", GoGetter: "Privileged"},
			_jsii_.MemberProperty{JsiiProperty: "readOnlyRootFilesystem", GoGetter: "ReadOnlyRootFilesystem"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
		},
		func() interface{} {
			return &jsiiProxy_ContainerSecurityContext{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ContainerSecurityContextProps",
		reflect.TypeOf((*ContainerSecurityContextProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Cpu",
		reflect.TypeOf((*Cpu)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "amount", GoGetter: "Amount"},
		},
		func() interface{} {
			return &jsiiProxy_Cpu{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.CpuResources",
		reflect.TypeOf((*CpuResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.DaemonSet",
		reflect.TypeOf((*DaemonSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addHostAlias", GoMethod: "AddHostAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "matchExpressions", GoGetter: "MatchExpressions"},
			_jsii_.MemberProperty{JsiiProperty: "matchLabels", GoGetter: "MatchLabels"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "minReadySeconds", GoGetter: "MinReadySeconds"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberMethod{JsiiMethod: "select", GoMethod: "Select"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_DaemonSet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Workload)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.DaemonSetProps",
		reflect.TypeOf((*DaemonSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Deployment",
		reflect.TypeOf((*Deployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addHostAlias", GoMethod: "AddHostAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberMethod{JsiiMethod: "exposeViaIngress", GoMethod: "ExposeViaIngress"},
			_jsii_.MemberMethod{JsiiMethod: "exposeViaService", GoMethod: "ExposeViaService"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "matchExpressions", GoGetter: "MatchExpressions"},
			_jsii_.MemberProperty{JsiiProperty: "matchLabels", GoGetter: "MatchLabels"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "minReady", GoGetter: "MinReady"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "progressDeadline", GoGetter: "ProgressDeadline"},
			_jsii_.MemberProperty{JsiiProperty: "replicas", GoGetter: "Replicas"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberMethod{JsiiMethod: "select", GoMethod: "Select"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_Deployment{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Workload)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.DeploymentProps",
		reflect.TypeOf((*DeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.DeploymentStrategy",
		reflect.TypeOf((*DeploymentStrategy)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DeploymentStrategy{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.DeploymentStrategyRollingUpdateOptions",
		reflect.TypeOf((*DeploymentStrategyRollingUpdateOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.DnsOption",
		reflect.TypeOf((*DnsOption)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.DnsPolicy",
		reflect.TypeOf((*DnsPolicy)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER_FIRST": DnsPolicy_CLUSTER_FIRST,
			"CLUSTER_FIRST_WITH_HOST_NET": DnsPolicy_CLUSTER_FIRST_WITH_HOST_NET,
			"DEFAULT": DnsPolicy_DEFAULT,
			"NONE": DnsPolicy_NONE,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.DockerConfigSecret",
		reflect.TypeOf((*DockerConfigSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStringData", GoMethod: "AddStringData"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "getStringData", GoMethod: "GetStringData"},
			_jsii_.MemberProperty{JsiiProperty: "immutable", GoGetter: "Immutable"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DockerConfigSecret{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Secret)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.DockerConfigSecretProps",
		reflect.TypeOf((*DockerConfigSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.EmptyDirMedium",
		reflect.TypeOf((*EmptyDirMedium)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": EmptyDirMedium_DEFAULT,
			"MEMORY": EmptyDirMedium_MEMORY,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.EmptyDirVolumeOptions",
		reflect.TypeOf((*EmptyDirVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Env",
		reflect.TypeOf((*Env)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVariable", GoMethod: "AddVariable"},
			_jsii_.MemberMethod{JsiiMethod: "copyFrom", GoMethod: "CopyFrom"},
			_jsii_.MemberProperty{JsiiProperty: "sources", GoGetter: "Sources"},
			_jsii_.MemberProperty{JsiiProperty: "variables", GoGetter: "Variables"},
		},
		func() interface{} {
			return &jsiiProxy_Env{}
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.EnvFieldPaths",
		reflect.TypeOf((*EnvFieldPaths)(nil)).Elem(),
		map[string]interface{}{
			"POD_NAME": EnvFieldPaths_POD_NAME,
			"POD_NAMESPACE": EnvFieldPaths_POD_NAMESPACE,
			"POD_UID": EnvFieldPaths_POD_UID,
			"POD_LABEL": EnvFieldPaths_POD_LABEL,
			"POD_ANNOTATION": EnvFieldPaths_POD_ANNOTATION,
			"POD_IP": EnvFieldPaths_POD_IP,
			"SERVICE_ACCOUNT_NAME": EnvFieldPaths_SERVICE_ACCOUNT_NAME,
			"NODE_NAME": EnvFieldPaths_NODE_NAME,
			"NODE_IP": EnvFieldPaths_NODE_IP,
			"POD_IPS": EnvFieldPaths_POD_IPS,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.EnvFrom",
		reflect.TypeOf((*EnvFrom)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EnvFrom{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.EnvValue",
		reflect.TypeOf((*EnvValue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueFrom", GoGetter: "ValueFrom"},
		},
		func() interface{} {
			return &jsiiProxy_EnvValue{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.EnvValueFromConfigMapOptions",
		reflect.TypeOf((*EnvValueFromConfigMapOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.EnvValueFromFieldRefOptions",
		reflect.TypeOf((*EnvValueFromFieldRefOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.EnvValueFromProcessOptions",
		reflect.TypeOf((*EnvValueFromProcessOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.EnvValueFromResourceOptions",
		reflect.TypeOf((*EnvValueFromResourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.EnvValueFromSecretOptions",
		reflect.TypeOf((*EnvValueFromSecretOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ExposeDeploymentViaIngressOptions",
		reflect.TypeOf((*ExposeDeploymentViaIngressOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ExposeDeploymentViaServiceOptions",
		reflect.TypeOf((*ExposeDeploymentViaServiceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ExposeServiceViaIngressOptions",
		reflect.TypeOf((*ExposeServiceViaIngressOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.FsGroupChangePolicy",
		reflect.TypeOf((*FsGroupChangePolicy)(nil)).Elem(),
		map[string]interface{}{
			"ON_ROOT_MISMATCH": FsGroupChangePolicy_ON_ROOT_MISMATCH,
			"ALWAYS": FsGroupChangePolicy_ALWAYS,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.GCEPersistentDiskPersistentVolume",
		reflect.TypeOf((*GCEPersistentDiskPersistentVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessModes", GoGetter: "AccessModes"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asVolume", GoMethod: "AsVolume"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "claim", GoGetter: "Claim"},
			_jsii_.MemberProperty{JsiiProperty: "fsType", GoGetter: "FsType"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
			_jsii_.MemberProperty{JsiiProperty: "mountOptions", GoGetter: "MountOptions"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "partition", GoGetter: "Partition"},
			_jsii_.MemberProperty{JsiiProperty: "pdName", GoGetter: "PdName"},
			_jsii_.MemberProperty{JsiiProperty: "readOnly", GoGetter: "ReadOnly"},
			_jsii_.MemberProperty{JsiiProperty: "reclaimPolicy", GoGetter: "ReclaimPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "reserve", GoMethod: "Reserve"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "storage", GoGetter: "Storage"},
			_jsii_.MemberProperty{JsiiProperty: "storageClassName", GoGetter: "StorageClassName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GCEPersistentDiskPersistentVolume{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PersistentVolume)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.GCEPersistentDiskPersistentVolumeProps",
		reflect.TypeOf((*GCEPersistentDiskPersistentVolumeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.GCEPersistentDiskVolumeOptions",
		reflect.TypeOf((*GCEPersistentDiskVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Group",
		reflect.TypeOf((*Group)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_Group{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISubject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.GroupProps",
		reflect.TypeOf((*GroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Handler",
		reflect.TypeOf((*Handler)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Handler{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.HandlerFromHttpGetOptions",
		reflect.TypeOf((*HandlerFromHttpGetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.HandlerFromTcpSocketOptions",
		reflect.TypeOf((*HandlerFromTcpSocketOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.HostAlias",
		reflect.TypeOf((*HostAlias)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.HttpGetProbeOptions",
		reflect.TypeOf((*HttpGetProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-21.IApiEndpoint",
		reflect.TypeOf((*IApiEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
		},
		func() interface{} {
			return &jsiiProxy_IApiEndpoint{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-21.IApiResource",
		reflect.TypeOf((*IApiResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
		},
		func() interface{} {
			return &jsiiProxy_IApiResource{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-21.IClusterRole",
		reflect.TypeOf((*IClusterRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_IClusterRole{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-21.IConfigMap",
		reflect.TypeOf((*IConfigMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_IConfigMap{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-21.IPersistentVolume",
		reflect.TypeOf((*IPersistentVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_IPersistentVolume{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-21.IPersistentVolumeClaim",
		reflect.TypeOf((*IPersistentVolumeClaim)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_IPersistentVolumeClaim{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-21.IResource",
		reflect.TypeOf((*IResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_IResource{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-21.IRole",
		reflect.TypeOf((*IRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_IRole{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-21.ISecret",
		reflect.TypeOf((*ISecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_ISecret{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-21.IServiceAccount",
		reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_IServiceAccount{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-21.IStorage",
		reflect.TypeOf((*IStorage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "asVolume", GoMethod: "AsVolume"},
		},
		func() interface{} {
			return &jsiiProxy_IStorage{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-21.ISubject",
		reflect.TypeOf((*ISubject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
		},
		func() interface{} {
			return &jsiiProxy_ISubject{}
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.ImagePullPolicy",
		reflect.TypeOf((*ImagePullPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": ImagePullPolicy_ALWAYS,
			"IF_NOT_PRESENT": ImagePullPolicy_IF_NOT_PRESENT,
			"NEVER": ImagePullPolicy_NEVER,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.IngressV1Beta1",
		reflect.TypeOf((*IngressV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultBackend", GoMethod: "AddDefaultBackend"},
			_jsii_.MemberMethod{JsiiMethod: "addHostDefaultBackend", GoMethod: "AddHostDefaultBackend"},
			_jsii_.MemberMethod{JsiiMethod: "addHostRule", GoMethod: "AddHostRule"},
			_jsii_.MemberMethod{JsiiMethod: "addRule", GoMethod: "AddRule"},
			_jsii_.MemberMethod{JsiiMethod: "addRules", GoMethod: "AddRules"},
			_jsii_.MemberMethod{JsiiMethod: "addTls", GoMethod: "AddTls"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IngressV1Beta1{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.IngressV1Beta1Backend",
		reflect.TypeOf((*IngressV1Beta1Backend)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IngressV1Beta1Backend{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.IngressV1Beta1Props",
		reflect.TypeOf((*IngressV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.IngressV1Beta1Rule",
		reflect.TypeOf((*IngressV1Beta1Rule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.IngressV1Beta1Tls",
		reflect.TypeOf((*IngressV1Beta1Tls)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Job",
		reflect.TypeOf((*Job)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activeDeadline", GoGetter: "ActiveDeadline"},
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addHostAlias", GoMethod: "AddHostAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "backoffLimit", GoGetter: "BackoffLimit"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "matchExpressions", GoGetter: "MatchExpressions"},
			_jsii_.MemberProperty{JsiiProperty: "matchLabels", GoGetter: "MatchLabels"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberMethod{JsiiMethod: "select", GoMethod: "Select"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "ttlAfterFinished", GoGetter: "TtlAfterFinished"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_Job{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Workload)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.JobProps",
		reflect.TypeOf((*JobProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.LabelSelector",
		reflect.TypeOf((*LabelSelector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applyToTemplate", GoGetter: "ApplyToTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
		},
		func() interface{} {
			return &jsiiProxy_LabelSelector{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.LabelSelectorRequirement",
		reflect.TypeOf((*LabelSelectorRequirement)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.LabelSelectorRequirementOperator",
		reflect.TypeOf((*LabelSelectorRequirementOperator)(nil)).Elem(),
		map[string]interface{}{
			"IN": LabelSelectorRequirementOperator_IN,
			"NOT_IN": LabelSelectorRequirementOperator_NOT_IN,
			"EXISTS": LabelSelectorRequirementOperator_EXISTS,
			"DOES_NOT_EXIST": LabelSelectorRequirementOperator_DOES_NOT_EXIST,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.MemoryResources",
		reflect.TypeOf((*MemoryResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.MountOptions",
		reflect.TypeOf((*MountOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.MountPropagation",
		reflect.TypeOf((*MountPropagation)(nil)).Elem(),
		map[string]interface{}{
			"NONE": MountPropagation_NONE,
			"HOST_TO_CONTAINER": MountPropagation_HOST_TO_CONTAINER,
			"BIDIRECTIONAL": MountPropagation_BIDIRECTIONAL,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.NonApiResource",
		reflect.TypeOf((*NonApiResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
		},
		func() interface{} {
			j := jsiiProxy_NonApiResource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApiEndpoint)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.PathMapping",
		reflect.TypeOf((*PathMapping)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.PercentOrAbsolute",
		reflect.TypeOf((*PercentOrAbsolute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "isZero", GoMethod: "IsZero"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PercentOrAbsolute{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.PersistentVolume",
		reflect.TypeOf((*PersistentVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessModes", GoGetter: "AccessModes"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asVolume", GoMethod: "AsVolume"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "claim", GoGetter: "Claim"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
			_jsii_.MemberProperty{JsiiProperty: "mountOptions", GoGetter: "MountOptions"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "reclaimPolicy", GoGetter: "ReclaimPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "reserve", GoMethod: "Reserve"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "storage", GoGetter: "Storage"},
			_jsii_.MemberProperty{JsiiProperty: "storageClassName", GoGetter: "StorageClassName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PersistentVolume{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPersistentVolume)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IStorage)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.PersistentVolumeAccessMode",
		reflect.TypeOf((*PersistentVolumeAccessMode)(nil)).Elem(),
		map[string]interface{}{
			"READ_WRITE_ONCE": PersistentVolumeAccessMode_READ_WRITE_ONCE,
			"READ_ONLY_MANY": PersistentVolumeAccessMode_READ_ONLY_MANY,
			"READ_WRITE_MANY": PersistentVolumeAccessMode_READ_WRITE_MANY,
			"READ_WRITE_ONCE_POD": PersistentVolumeAccessMode_READ_WRITE_ONCE_POD,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.PersistentVolumeClaim",
		reflect.TypeOf((*PersistentVolumeClaim)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessModes", GoGetter: "AccessModes"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "storage", GoGetter: "Storage"},
			_jsii_.MemberProperty{JsiiProperty: "storageClassName", GoGetter: "StorageClassName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volume", GoGetter: "Volume"},
			_jsii_.MemberProperty{JsiiProperty: "volumeMode", GoGetter: "VolumeMode"},
		},
		func() interface{} {
			j := jsiiProxy_PersistentVolumeClaim{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPersistentVolumeClaim)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.PersistentVolumeClaimProps",
		reflect.TypeOf((*PersistentVolumeClaimProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.PersistentVolumeClaimVolumeOptions",
		reflect.TypeOf((*PersistentVolumeClaimVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.PersistentVolumeMode",
		reflect.TypeOf((*PersistentVolumeMode)(nil)).Elem(),
		map[string]interface{}{
			"FILE_SYSTEM": PersistentVolumeMode_FILE_SYSTEM,
			"BLOCK": PersistentVolumeMode_BLOCK,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.PersistentVolumeProps",
		reflect.TypeOf((*PersistentVolumeProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.PersistentVolumeReclaimPolicy",
		reflect.TypeOf((*PersistentVolumeReclaimPolicy)(nil)).Elem(),
		map[string]interface{}{
			"RETAIN": PersistentVolumeReclaimPolicy_RETAIN,
			"DELETE": PersistentVolumeReclaimPolicy_DELETE,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Pod",
		reflect.TypeOf((*Pod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addHostAlias", GoMethod: "AddHostAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_Pod{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AbstractPod)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.PodDns",
		reflect.TypeOf((*PodDns)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addNameserver", GoMethod: "AddNameserver"},
			_jsii_.MemberMethod{JsiiMethod: "addOption", GoMethod: "AddOption"},
			_jsii_.MemberMethod{JsiiMethod: "addSearch", GoMethod: "AddSearch"},
			_jsii_.MemberProperty{JsiiProperty: "hostname", GoGetter: "Hostname"},
			_jsii_.MemberProperty{JsiiProperty: "hostnameAsFQDN", GoGetter: "HostnameAsFQDN"},
			_jsii_.MemberProperty{JsiiProperty: "nameservers", GoGetter: "Nameservers"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "searches", GoGetter: "Searches"},
			_jsii_.MemberProperty{JsiiProperty: "subdomain", GoGetter: "Subdomain"},
		},
		func() interface{} {
			return &jsiiProxy_PodDns{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.PodDnsProps",
		reflect.TypeOf((*PodDnsProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.PodManagementPolicy",
		reflect.TypeOf((*PodManagementPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ORDERED_READY": PodManagementPolicy_ORDERED_READY,
			"PARALLEL": PodManagementPolicy_PARALLEL,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.PodProps",
		reflect.TypeOf((*PodProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.PodSecurityContext",
		reflect.TypeOf((*PodSecurityContext)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ensureNonRoot", GoGetter: "EnsureNonRoot"},
			_jsii_.MemberProperty{JsiiProperty: "fsGroup", GoGetter: "FsGroup"},
			_jsii_.MemberProperty{JsiiProperty: "fsGroupChangePolicy", GoGetter: "FsGroupChangePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "group", GoGetter: "Group"},
			_jsii_.MemberProperty{JsiiProperty: "sysctls", GoGetter: "Sysctls"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
		},
		func() interface{} {
			return &jsiiProxy_PodSecurityContext{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.PodSecurityContextProps",
		reflect.TypeOf((*PodSecurityContextProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Probe",
		reflect.TypeOf((*Probe)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Probe{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ProbeOptions",
		reflect.TypeOf((*ProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.Protocol",
		reflect.TypeOf((*Protocol)(nil)).Elem(),
		map[string]interface{}{
			"TCP": Protocol_TCP,
			"UDP": Protocol_UDP,
			"SCTP": Protocol_SCTP,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Resource",
		reflect.TypeOf((*Resource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Resource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApiEndpoint)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApiResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.ResourceFieldPaths",
		reflect.TypeOf((*ResourceFieldPaths)(nil)).Elem(),
		map[string]interface{}{
			"CPU_LIMIT": ResourceFieldPaths_CPU_LIMIT,
			"MEMORY_LIMIT": ResourceFieldPaths_MEMORY_LIMIT,
			"CPU_REQUEST": ResourceFieldPaths_CPU_REQUEST,
			"MEMORY_REQUEST": ResourceFieldPaths_MEMORY_REQUEST,
			"STORAGE_LIMIT": ResourceFieldPaths_STORAGE_LIMIT,
			"STORAGE_REQUEST": ResourceFieldPaths_STORAGE_REQUEST,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ResourceProps",
		reflect.TypeOf((*ResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.RestartPolicy",
		reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RestartPolicy_ALWAYS,
			"ON_FAILURE": RestartPolicy_ON_FAILURE,
			"NEVER": RestartPolicy_NEVER,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Role",
		reflect.TypeOf((*Role)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allow", GoMethod: "Allow"},
			_jsii_.MemberMethod{JsiiMethod: "allowCreate", GoMethod: "AllowCreate"},
			_jsii_.MemberMethod{JsiiMethod: "allowDelete", GoMethod: "AllowDelete"},
			_jsii_.MemberMethod{JsiiMethod: "allowDeleteCollection", GoMethod: "AllowDeleteCollection"},
			_jsii_.MemberMethod{JsiiMethod: "allowGet", GoMethod: "AllowGet"},
			_jsii_.MemberMethod{JsiiMethod: "allowList", GoMethod: "AllowList"},
			_jsii_.MemberMethod{JsiiMethod: "allowPatch", GoMethod: "AllowPatch"},
			_jsii_.MemberMethod{JsiiMethod: "allowRead", GoMethod: "AllowRead"},
			_jsii_.MemberMethod{JsiiMethod: "allowReadWrite", GoMethod: "AllowReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "allowUpdate", GoMethod: "AllowUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "allowWatch", GoMethod: "AllowWatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Role{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRole)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.RoleBinding",
		reflect.TypeOf((*RoleBinding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSubjects", GoMethod: "AddSubjects"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "subjects", GoGetter: "Subjects"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RoleBinding{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.RoleBindingProps",
		reflect.TypeOf((*RoleBindingProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.RolePolicyRule",
		reflect.TypeOf((*RolePolicyRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.RoleProps",
		reflect.TypeOf((*RoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Secret",
		reflect.TypeOf((*Secret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStringData", GoMethod: "AddStringData"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "getStringData", GoMethod: "GetStringData"},
			_jsii_.MemberProperty{JsiiProperty: "immutable", GoGetter: "Immutable"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Secret{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISecret)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.SecretProps",
		reflect.TypeOf((*SecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.SecretValue",
		reflect.TypeOf((*SecretValue)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.SecretVolumeOptions",
		reflect.TypeOf((*SecretVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Service",
		reflect.TypeOf((*Service)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeployment", GoMethod: "AddDeployment"},
			_jsii_.MemberMethod{JsiiMethod: "addSelector", GoMethod: "AddSelector"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIP", GoGetter: "ClusterIP"},
			_jsii_.MemberMethod{JsiiMethod: "exposeViaIngress", GoMethod: "ExposeViaIngress"},
			_jsii_.MemberProperty{JsiiProperty: "externalName", GoGetter: "ExternalName"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ports", GoGetter: "Ports"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "selector", GoGetter: "Selector"},
			_jsii_.MemberMethod{JsiiMethod: "serve", GoMethod: "Serve"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_Service{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.ServiceAccount",
		reflect.TypeOf((*ServiceAccount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSecret", GoMethod: "AddSecret"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "automountToken", GoGetter: "AutomountToken"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "secrets", GoGetter: "Secrets"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceAccount{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServiceAccount)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISubject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ServiceAccountProps",
		reflect.TypeOf((*ServiceAccountProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.ServiceAccountTokenSecret",
		reflect.TypeOf((*ServiceAccountTokenSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStringData", GoMethod: "AddStringData"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "getStringData", GoMethod: "GetStringData"},
			_jsii_.MemberProperty{JsiiProperty: "immutable", GoGetter: "Immutable"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceAccountTokenSecret{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Secret)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ServiceAccountTokenSecretProps",
		reflect.TypeOf((*ServiceAccountTokenSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ServiceIngressV1BetaBackendOptions",
		reflect.TypeOf((*ServiceIngressV1BetaBackendOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ServicePort",
		reflect.TypeOf((*ServicePort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ServicePortOptions",
		reflect.TypeOf((*ServicePortOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.ServiceProps",
		reflect.TypeOf((*ServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-21.ServiceType",
		reflect.TypeOf((*ServiceType)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER_IP": ServiceType_CLUSTER_IP,
			"NODE_PORT": ServiceType_NODE_PORT,
			"LOAD_BALANCER": ServiceType_LOAD_BALANCER,
			"EXTERNAL_NAME": ServiceType_EXTERNAL_NAME,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.SshAuthSecret",
		reflect.TypeOf((*SshAuthSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStringData", GoMethod: "AddStringData"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "getStringData", GoMethod: "GetStringData"},
			_jsii_.MemberProperty{JsiiProperty: "immutable", GoGetter: "Immutable"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SshAuthSecret{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Secret)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.SshAuthSecretProps",
		reflect.TypeOf((*SshAuthSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.StatefulSet",
		reflect.TypeOf((*StatefulSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addHostAlias", GoMethod: "AddHostAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "matchExpressions", GoGetter: "MatchExpressions"},
			_jsii_.MemberProperty{JsiiProperty: "matchLabels", GoGetter: "MatchLabels"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "podManagementPolicy", GoGetter: "PodManagementPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "replicas", GoGetter: "Replicas"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberMethod{JsiiMethod: "select", GoMethod: "Select"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_StatefulSet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Workload)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.StatefulSetProps",
		reflect.TypeOf((*StatefulSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.StatefulSetUpdateStrategy",
		reflect.TypeOf((*StatefulSetUpdateStrategy)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_StatefulSetUpdateStrategy{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.StatefulSetUpdateStrategyRollingUpdateOptions",
		reflect.TypeOf((*StatefulSetUpdateStrategyRollingUpdateOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.Sysctl",
		reflect.TypeOf((*Sysctl)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.TcpSocketProbeOptions",
		reflect.TypeOf((*TcpSocketProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.TlsSecret",
		reflect.TypeOf((*TlsSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStringData", GoMethod: "AddStringData"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "getStringData", GoMethod: "GetStringData"},
			_jsii_.MemberProperty{JsiiProperty: "immutable", GoGetter: "Immutable"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TlsSecret{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Secret)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.TlsSecretProps",
		reflect.TypeOf((*TlsSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.User",
		reflect.TypeOf((*User)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_User{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISubject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.UserProps",
		reflect.TypeOf((*UserProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Volume",
		reflect.TypeOf((*Volume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "asVolume", GoMethod: "AsVolume"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_Volume{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IStorage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.VolumeMount",
		reflect.TypeOf((*VolumeMount)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-21.Workload",
		reflect.TypeOf((*Workload)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addHostAlias", GoMethod: "AddHostAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "matchExpressions", GoGetter: "MatchExpressions"},
			_jsii_.MemberProperty{JsiiProperty: "matchLabels", GoGetter: "MatchLabels"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberMethod{JsiiMethod: "select", GoMethod: "Select"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_Workload{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AbstractPod)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-21.WorkloadProps",
		reflect.TypeOf((*WorkloadProps)(nil)).Elem(),
	)
}
