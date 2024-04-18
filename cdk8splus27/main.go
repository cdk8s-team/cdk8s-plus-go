// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-27 synthesizes Kubernetes manifests for Kubernetes 1.27.0
package cdk8splus27

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk8s-plus-27.AbstractPod",
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
			_jsii_.MemberMethod{JsiiMethod: "attachContainer", GoMethod: "AttachContainer"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "hostNetwork", GoGetter: "HostNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "isolate", GoGetter: "Isolate"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "terminationGracePeriod", GoGetter: "TerminationGracePeriod"},
			_jsii_.MemberMethod{JsiiMethod: "toNetworkPolicyPeerConfig", GoMethod: "ToNetworkPolicyPeerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelector", GoMethod: "ToPodSelector"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelectorConfig", GoMethod: "ToPodSelectorConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toSubjectConfiguration", GoMethod: "ToSubjectConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_AbstractPod{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INetworkPolicyPeer)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPodSelector)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISubject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.AbstractPodProps",
		reflect.TypeOf((*AbstractPodProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.AddDeploymentOptions",
		reflect.TypeOf((*AddDeploymentOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.AddDirectoryOptions",
		reflect.TypeOf((*AddDirectoryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.ApiResource",
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
		"cdk8s-plus-27.ApiResourceOptions",
		reflect.TypeOf((*ApiResourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.AwsElasticBlockStorePersistentVolume",
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
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.AwsElasticBlockStorePersistentVolumeProps",
		reflect.TypeOf((*AwsElasticBlockStorePersistentVolumeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.AwsElasticBlockStoreVolumeOptions",
		reflect.TypeOf((*AwsElasticBlockStoreVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.AzureDiskPersistentVolume",
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
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.AzureDiskPersistentVolumeCachingMode",
		reflect.TypeOf((*AzureDiskPersistentVolumeCachingMode)(nil)).Elem(),
		map[string]interface{}{
			"NONE": AzureDiskPersistentVolumeCachingMode_NONE,
			"READ_ONLY": AzureDiskPersistentVolumeCachingMode_READ_ONLY,
			"READ_WRITE": AzureDiskPersistentVolumeCachingMode_READ_WRITE,
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.AzureDiskPersistentVolumeKind",
		reflect.TypeOf((*AzureDiskPersistentVolumeKind)(nil)).Elem(),
		map[string]interface{}{
			"SHARED": AzureDiskPersistentVolumeKind_SHARED,
			"DEDICATED": AzureDiskPersistentVolumeKind_DEDICATED,
			"MANAGED": AzureDiskPersistentVolumeKind_MANAGED,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.AzureDiskPersistentVolumeProps",
		reflect.TypeOf((*AzureDiskPersistentVolumeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.AzureDiskVolumeOptions",
		reflect.TypeOf((*AzureDiskVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.BasicAuthSecret",
		reflect.TypeOf((*BasicAuthSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStringData", GoMethod: "AddStringData"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "envValue", GoMethod: "EnvValue"},
			_jsii_.MemberMethod{JsiiMethod: "getStringData", GoMethod: "GetStringData"},
			_jsii_.MemberProperty{JsiiProperty: "immutable", GoGetter: "Immutable"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.BasicAuthSecretProps",
		reflect.TypeOf((*BasicAuthSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.Capability",
		reflect.TypeOf((*Capability)(nil)).Elem(),
		map[string]interface{}{
			"AUDIT_CONTROL": Capability_AUDIT_CONTROL,
			"AUDIT_READ": Capability_AUDIT_READ,
			"AUDIT_WRITE": Capability_AUDIT_WRITE,
			"BLOCK_SUSPEND": Capability_BLOCK_SUSPEND,
			"BPF": Capability_BPF,
			"CHECKPOINT_RESTORE": Capability_CHECKPOINT_RESTORE,
			"CHOWN": Capability_CHOWN,
			"DAC_OVERRIDE": Capability_DAC_OVERRIDE,
			"DAC_READ_SEARCH": Capability_DAC_READ_SEARCH,
			"FOWNER": Capability_FOWNER,
			"FSETID": Capability_FSETID,
			"IPC_LOCK": Capability_IPC_LOCK,
			"IPC_OWNER": Capability_IPC_OWNER,
			"KILL": Capability_KILL,
			"LEASE": Capability_LEASE,
			"LINUX_IMMUTABLE": Capability_LINUX_IMMUTABLE,
			"MAC_ADMIN": Capability_MAC_ADMIN,
			"MAC_OVERRIDE": Capability_MAC_OVERRIDE,
			"MKNOD": Capability_MKNOD,
			"NET_ADMIN": Capability_NET_ADMIN,
			"NET_BIND_SERVICE": Capability_NET_BIND_SERVICE,
			"NET_BROADCAST": Capability_NET_BROADCAST,
			"NET_RAW": Capability_NET_RAW,
			"PERFMON": Capability_PERFMON,
			"SETGID": Capability_SETGID,
			"SETFCAP": Capability_SETFCAP,
			"SETPCAP": Capability_SETPCAP,
			"SETUID": Capability_SETUID,
			"SYS_ADMIN": Capability_SYS_ADMIN,
			"SYS_BOOT": Capability_SYS_BOOT,
			"SYS_CHROOT": Capability_SYS_CHROOT,
			"SYS_MODULE": Capability_SYS_MODULE,
			"SYS_NICE": Capability_SYS_NICE,
			"SYS_PACCT": Capability_SYS_PACCT,
			"SYS_PTRACE": Capability_SYS_PTRACE,
			"SYS_RAWIO": Capability_SYS_RAWIO,
			"SYS_RESOURCE": Capability_SYS_RESOURCE,
			"SYS_TIME": Capability_SYS_TIME,
			"SYS_TTY_CONFIG": Capability_SYS_TTY_CONFIG,
			"SYSLOG": Capability_SYSLOG,
			"WAKE_ALARM": Capability_WAKE_ALARM,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.ClusterRole",
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
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.ClusterRoleBinding",
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
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.ClusterRoleBindingProps",
		reflect.TypeOf((*ClusterRoleBindingProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ClusterRolePolicyRule",
		reflect.TypeOf((*ClusterRolePolicyRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ClusterRoleProps",
		reflect.TypeOf((*ClusterRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.CommandProbeOptions",
		reflect.TypeOf((*CommandProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.CommonSecretProps",
		reflect.TypeOf((*CommonSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.ConcurrencyPolicy",
		reflect.TypeOf((*ConcurrencyPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ALLOW": ConcurrencyPolicy_ALLOW,
			"FORBID": ConcurrencyPolicy_FORBID,
			"REPLACE": ConcurrencyPolicy_REPLACE,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.ConfigMap",
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
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.ConfigMapProps",
		reflect.TypeOf((*ConfigMapProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ConfigMapVolumeOptions",
		reflect.TypeOf((*ConfigMapVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.ConnectionScheme",
		reflect.TypeOf((*ConnectionScheme)(nil)).Elem(),
		map[string]interface{}{
			"HTTP": ConnectionScheme_HTTP,
			"HTTPS": ConnectionScheme_HTTPS,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Container",
		reflect.TypeOf((*Container)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPort", GoMethod: "AddPort"},
			_jsii_.MemberProperty{JsiiProperty: "args", GoGetter: "Args"},
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "imagePullPolicy", GoGetter: "ImagePullPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "mount", GoMethod: "Mount"},
			_jsii_.MemberProperty{JsiiProperty: "mounts", GoGetter: "Mounts"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portNumber", GoGetter: "PortNumber"},
			_jsii_.MemberProperty{JsiiProperty: "ports", GoGetter: "Ports"},
			_jsii_.MemberProperty{JsiiProperty: "resources", GoGetter: "Resources"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberProperty{JsiiProperty: "workingDir", GoGetter: "WorkingDir"},
		},
		func() interface{} {
			return &jsiiProxy_Container{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ContainerLifecycle",
		reflect.TypeOf((*ContainerLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ContainerOpts",
		reflect.TypeOf((*ContainerOpts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ContainerPort",
		reflect.TypeOf((*ContainerPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ContainerProps",
		reflect.TypeOf((*ContainerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ContainerResources",
		reflect.TypeOf((*ContainerResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.ContainerSecurityContext",
		reflect.TypeOf((*ContainerSecurityContext)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPrivilegeEscalation", GoGetter: "AllowPrivilegeEscalation"},
			_jsii_.MemberProperty{JsiiProperty: "capabilities", GoGetter: "Capabilities"},
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
		"cdk8s-plus-27.ContainerSecurityContextProps",
		reflect.TypeOf((*ContainerSecurityContextProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ContainerSecutiryContextCapabilities",
		reflect.TypeOf((*ContainerSecutiryContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Cpu",
		reflect.TypeOf((*Cpu)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "amount", GoGetter: "Amount"},
		},
		func() interface{} {
			return &jsiiProxy_Cpu{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.CpuResources",
		reflect.TypeOf((*CpuResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.CronJob",
		reflect.TypeOf((*CronJob)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "attachContainer", GoMethod: "AttachContainer"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "concurrencyPolicy", GoGetter: "ConcurrencyPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberProperty{JsiiProperty: "failedJobsRetained", GoGetter: "FailedJobsRetained"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "hostNetwork", GoGetter: "HostNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "isolate", GoGetter: "Isolate"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "matchExpressions", GoGetter: "MatchExpressions"},
			_jsii_.MemberProperty{JsiiProperty: "matchLabels", GoGetter: "MatchLabels"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "schedule", GoGetter: "Schedule"},
			_jsii_.MemberProperty{JsiiProperty: "scheduling", GoGetter: "Scheduling"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberMethod{JsiiMethod: "select", GoMethod: "Select"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "startingDeadline", GoGetter: "StartingDeadline"},
			_jsii_.MemberProperty{JsiiProperty: "successfulJobsRetained", GoGetter: "SuccessfulJobsRetained"},
			_jsii_.MemberProperty{JsiiProperty: "suspend", GoGetter: "Suspend"},
			_jsii_.MemberProperty{JsiiProperty: "terminationGracePeriod", GoGetter: "TerminationGracePeriod"},
			_jsii_.MemberProperty{JsiiProperty: "timeZone", GoGetter: "TimeZone"},
			_jsii_.MemberMethod{JsiiMethod: "toNetworkPolicyPeerConfig", GoMethod: "ToNetworkPolicyPeerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelector", GoMethod: "ToPodSelector"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelectorConfig", GoMethod: "ToPodSelectorConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toSubjectConfiguration", GoMethod: "ToSubjectConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_CronJob{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Workload)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.CronJobProps",
		reflect.TypeOf((*CronJobProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.CsiVolumeOptions",
		reflect.TypeOf((*CsiVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.DaemonSet",
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
			_jsii_.MemberMethod{JsiiMethod: "attachContainer", GoMethod: "AttachContainer"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "hostNetwork", GoGetter: "HostNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "isolate", GoGetter: "Isolate"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "matchExpressions", GoGetter: "MatchExpressions"},
			_jsii_.MemberProperty{JsiiProperty: "matchLabels", GoGetter: "MatchLabels"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "minReadySeconds", GoGetter: "MinReadySeconds"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "scheduling", GoGetter: "Scheduling"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberMethod{JsiiMethod: "select", GoMethod: "Select"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "terminationGracePeriod", GoGetter: "TerminationGracePeriod"},
			_jsii_.MemberMethod{JsiiMethod: "toNetworkPolicyPeerConfig", GoMethod: "ToNetworkPolicyPeerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelector", GoMethod: "ToPodSelector"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelectorConfig", GoMethod: "ToPodSelectorConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toSubjectConfiguration", GoMethod: "ToSubjectConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_DaemonSet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Workload)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.DaemonSetProps",
		reflect.TypeOf((*DaemonSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Deployment",
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
			_jsii_.MemberMethod{JsiiMethod: "attachContainer", GoMethod: "AttachContainer"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberMethod{JsiiMethod: "exposeViaIngress", GoMethod: "ExposeViaIngress"},
			_jsii_.MemberMethod{JsiiMethod: "exposeViaService", GoMethod: "ExposeViaService"},
			_jsii_.MemberProperty{JsiiProperty: "hasAutoscaler", GoGetter: "HasAutoscaler"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "hostNetwork", GoGetter: "HostNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "isolate", GoGetter: "Isolate"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberMethod{JsiiMethod: "markHasAutoscaler", GoMethod: "MarkHasAutoscaler"},
			_jsii_.MemberProperty{JsiiProperty: "matchExpressions", GoGetter: "MatchExpressions"},
			_jsii_.MemberProperty{JsiiProperty: "matchLabels", GoGetter: "MatchLabels"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "minReady", GoGetter: "MinReady"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "progressDeadline", GoGetter: "ProgressDeadline"},
			_jsii_.MemberProperty{JsiiProperty: "replicas", GoGetter: "Replicas"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "scheduling", GoGetter: "Scheduling"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberMethod{JsiiMethod: "select", GoMethod: "Select"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberProperty{JsiiProperty: "terminationGracePeriod", GoGetter: "TerminationGracePeriod"},
			_jsii_.MemberMethod{JsiiMethod: "toNetworkPolicyPeerConfig", GoMethod: "ToNetworkPolicyPeerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelector", GoMethod: "ToPodSelector"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelectorConfig", GoMethod: "ToPodSelectorConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toScalingTarget", GoMethod: "ToScalingTarget"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toSubjectConfiguration", GoMethod: "ToSubjectConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_Deployment{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Workload)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IScalable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.DeploymentExposeViaServiceOptions",
		reflect.TypeOf((*DeploymentExposeViaServiceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.DeploymentProps",
		reflect.TypeOf((*DeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.DeploymentStrategy",
		reflect.TypeOf((*DeploymentStrategy)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DeploymentStrategy{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.DeploymentStrategyRollingUpdateOptions",
		reflect.TypeOf((*DeploymentStrategyRollingUpdateOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.DnsOption",
		reflect.TypeOf((*DnsOption)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.DnsPolicy",
		reflect.TypeOf((*DnsPolicy)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER_FIRST": DnsPolicy_CLUSTER_FIRST,
			"CLUSTER_FIRST_WITH_HOST_NET": DnsPolicy_CLUSTER_FIRST_WITH_HOST_NET,
			"DEFAULT": DnsPolicy_DEFAULT,
			"NONE": DnsPolicy_NONE,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.DockerConfigSecret",
		reflect.TypeOf((*DockerConfigSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStringData", GoMethod: "AddStringData"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "envValue", GoMethod: "EnvValue"},
			_jsii_.MemberMethod{JsiiMethod: "getStringData", GoMethod: "GetStringData"},
			_jsii_.MemberProperty{JsiiProperty: "immutable", GoGetter: "Immutable"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.DockerConfigSecretProps",
		reflect.TypeOf((*DockerConfigSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.EmptyDirMedium",
		reflect.TypeOf((*EmptyDirMedium)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": EmptyDirMedium_DEFAULT,
			"MEMORY": EmptyDirMedium_MEMORY,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.EmptyDirVolumeOptions",
		reflect.TypeOf((*EmptyDirVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Env",
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
		"cdk8s-plus-27.EnvFieldPaths",
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
		"cdk8s-plus-27.EnvFrom",
		reflect.TypeOf((*EnvFrom)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EnvFrom{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.EnvValue",
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
		"cdk8s-plus-27.EnvValueFromConfigMapOptions",
		reflect.TypeOf((*EnvValueFromConfigMapOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.EnvValueFromFieldRefOptions",
		reflect.TypeOf((*EnvValueFromFieldRefOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.EnvValueFromProcessOptions",
		reflect.TypeOf((*EnvValueFromProcessOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.EnvValueFromResourceOptions",
		reflect.TypeOf((*EnvValueFromResourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.EnvValueFromSecretOptions",
		reflect.TypeOf((*EnvValueFromSecretOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.EphemeralStorageResources",
		reflect.TypeOf((*EphemeralStorageResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ExposeDeploymentViaIngressOptions",
		reflect.TypeOf((*ExposeDeploymentViaIngressOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ExposeServiceViaIngressOptions",
		reflect.TypeOf((*ExposeServiceViaIngressOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.FromServiceAccountNameOptions",
		reflect.TypeOf((*FromServiceAccountNameOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.FsGroupChangePolicy",
		reflect.TypeOf((*FsGroupChangePolicy)(nil)).Elem(),
		map[string]interface{}{
			"ON_ROOT_MISMATCH": FsGroupChangePolicy_ON_ROOT_MISMATCH,
			"ALWAYS": FsGroupChangePolicy_ALWAYS,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.GCEPersistentDiskPersistentVolume",
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
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.GCEPersistentDiskPersistentVolumeProps",
		reflect.TypeOf((*GCEPersistentDiskPersistentVolumeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.GCEPersistentDiskVolumeOptions",
		reflect.TypeOf((*GCEPersistentDiskVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Group",
		reflect.TypeOf((*Group)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toSubjectConfiguration", GoMethod: "ToSubjectConfiguration"},
		},
		func() interface{} {
			j := jsiiProxy_Group{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISubject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Handler",
		reflect.TypeOf((*Handler)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Handler{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.HandlerFromHttpGetOptions",
		reflect.TypeOf((*HandlerFromHttpGetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.HandlerFromTcpSocketOptions",
		reflect.TypeOf((*HandlerFromTcpSocketOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.HorizontalPodAutoscaler",
		reflect.TypeOf((*HorizontalPodAutoscaler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "maxReplicas", GoGetter: "MaxReplicas"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metrics", GoGetter: "Metrics"},
			_jsii_.MemberProperty{JsiiProperty: "minReplicas", GoGetter: "MinReplicas"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "scaleDown", GoGetter: "ScaleDown"},
			_jsii_.MemberProperty{JsiiProperty: "scaleUp", GoGetter: "ScaleUp"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_HorizontalPodAutoscaler{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.HorizontalPodAutoscalerProps",
		reflect.TypeOf((*HorizontalPodAutoscalerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.HostAlias",
		reflect.TypeOf((*HostAlias)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.HostPathVolumeOptions",
		reflect.TypeOf((*HostPathVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.HostPathVolumeType",
		reflect.TypeOf((*HostPathVolumeType)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": HostPathVolumeType_DEFAULT,
			"DIRECTORY_OR_CREATE": HostPathVolumeType_DIRECTORY_OR_CREATE,
			"DIRECTORY": HostPathVolumeType_DIRECTORY,
			"FILE_OR_CREATE": HostPathVolumeType_FILE_OR_CREATE,
			"FILE": HostPathVolumeType_FILE,
			"SOCKET": HostPathVolumeType_SOCKET,
			"CHAR_DEVICE": HostPathVolumeType_CHAR_DEVICE,
			"BLOCK_DEVICE": HostPathVolumeType_BLOCK_DEVICE,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.HttpGetProbeOptions",
		reflect.TypeOf((*HttpGetProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.HttpIngressPathType",
		reflect.TypeOf((*HttpIngressPathType)(nil)).Elem(),
		map[string]interface{}{
			"PREFIX": HttpIngressPathType_PREFIX,
			"EXACT": HttpIngressPathType_EXACT,
			"IMPLEMENTATION_SPECIFIC": HttpIngressPathType_IMPLEMENTATION_SPECIFIC,
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-27.IApiEndpoint",
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
		"cdk8s-plus-27.IApiResource",
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
		"cdk8s-plus-27.IClusterRole",
		reflect.TypeOf((*IClusterRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
		},
		func() interface{} {
			j := jsiiProxy_IClusterRole{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-27.IConfigMap",
		reflect.TypeOf((*IConfigMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
		},
		func() interface{} {
			j := jsiiProxy_IConfigMap{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-27.INamespaceSelector",
		reflect.TypeOf((*INamespaceSelector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toNamespaceSelectorConfig", GoMethod: "ToNamespaceSelectorConfig"},
		},
		func() interface{} {
			j := jsiiProxy_INamespaceSelector{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-27.INetworkPolicyPeer",
		reflect.TypeOf((*INetworkPolicyPeer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toNetworkPolicyPeerConfig", GoMethod: "ToNetworkPolicyPeerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelector", GoMethod: "ToPodSelector"},
		},
		func() interface{} {
			j := jsiiProxy_INetworkPolicyPeer{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-27.IPersistentVolume",
		reflect.TypeOf((*IPersistentVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
		},
		func() interface{} {
			j := jsiiProxy_IPersistentVolume{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-27.IPersistentVolumeClaim",
		reflect.TypeOf((*IPersistentVolumeClaim)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
		},
		func() interface{} {
			j := jsiiProxy_IPersistentVolumeClaim{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-27.IPodSelector",
		reflect.TypeOf((*IPodSelector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelectorConfig", GoMethod: "ToPodSelectorConfig"},
		},
		func() interface{} {
			j := jsiiProxy_IPodSelector{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-27.IResource",
		reflect.TypeOf((*IResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
		},
		func() interface{} {
			j := jsiiProxy_IResource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApiResource)
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-27.IRole",
		reflect.TypeOf((*IRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
		},
		func() interface{} {
			j := jsiiProxy_IRole{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-27.IScalable",
		reflect.TypeOf((*IScalable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "hasAutoscaler", GoGetter: "HasAutoscaler"},
			_jsii_.MemberMethod{JsiiMethod: "markHasAutoscaler", GoMethod: "MarkHasAutoscaler"},
			_jsii_.MemberMethod{JsiiMethod: "toScalingTarget", GoMethod: "ToScalingTarget"},
		},
		func() interface{} {
			return &jsiiProxy_IScalable{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-27.ISecret",
		reflect.TypeOf((*ISecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "envValue", GoMethod: "EnvValue"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
		},
		func() interface{} {
			j := jsiiProxy_ISecret{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-27.IServiceAccount",
		reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toSubjectConfiguration", GoMethod: "ToSubjectConfiguration"},
		},
		func() interface{} {
			j := jsiiProxy_IServiceAccount{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISubject)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-27.IStorage",
		reflect.TypeOf((*IStorage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "asVolume", GoMethod: "AsVolume"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IStorage{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-27.ISubject",
		reflect.TypeOf((*ISubject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toSubjectConfiguration", GoMethod: "ToSubjectConfiguration"},
		},
		func() interface{} {
			j := jsiiProxy_ISubject{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.ImagePullPolicy",
		reflect.TypeOf((*ImagePullPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": ImagePullPolicy_ALWAYS,
			"IF_NOT_PRESENT": ImagePullPolicy_IF_NOT_PRESENT,
			"NEVER": ImagePullPolicy_NEVER,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Ingress",
		reflect.TypeOf((*Ingress)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ingress{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.IngressBackend",
		reflect.TypeOf((*IngressBackend)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IngressBackend{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.IngressProps",
		reflect.TypeOf((*IngressProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.IngressRule",
		reflect.TypeOf((*IngressRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.IngressTls",
		reflect.TypeOf((*IngressTls)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Job",
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
			_jsii_.MemberMethod{JsiiMethod: "attachContainer", GoMethod: "AttachContainer"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "backoffLimit", GoGetter: "BackoffLimit"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "hostNetwork", GoGetter: "HostNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "isolate", GoGetter: "Isolate"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "matchExpressions", GoGetter: "MatchExpressions"},
			_jsii_.MemberProperty{JsiiProperty: "matchLabels", GoGetter: "MatchLabels"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "scheduling", GoGetter: "Scheduling"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberMethod{JsiiMethod: "select", GoMethod: "Select"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "terminationGracePeriod", GoGetter: "TerminationGracePeriod"},
			_jsii_.MemberMethod{JsiiMethod: "toNetworkPolicyPeerConfig", GoMethod: "ToNetworkPolicyPeerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelector", GoMethod: "ToPodSelector"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelectorConfig", GoMethod: "ToPodSelectorConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toSubjectConfiguration", GoMethod: "ToSubjectConfiguration"},
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
		"cdk8s-plus-27.JobProps",
		reflect.TypeOf((*JobProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.LabelExpression",
		reflect.TypeOf((*LabelExpression)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
		},
		func() interface{} {
			return &jsiiProxy_LabelExpression{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.LabelSelector",
		reflect.TypeOf((*LabelSelector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "isEmpty", GoMethod: "IsEmpty"},
		},
		func() interface{} {
			return &jsiiProxy_LabelSelector{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.LabelSelectorOptions",
		reflect.TypeOf((*LabelSelectorOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.LabelSelectorRequirement",
		reflect.TypeOf((*LabelSelectorRequirement)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.LabeledNode",
		reflect.TypeOf((*LabeledNode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "labelSelector", GoGetter: "LabelSelector"},
		},
		func() interface{} {
			return &jsiiProxy_LabeledNode{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.MemoryResources",
		reflect.TypeOf((*MemoryResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Metric",
		reflect.TypeOf((*Metric)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_Metric{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.MetricContainerResourceOptions",
		reflect.TypeOf((*MetricContainerResourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.MetricObjectOptions",
		reflect.TypeOf((*MetricObjectOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.MetricOptions",
		reflect.TypeOf((*MetricOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.MetricTarget",
		reflect.TypeOf((*MetricTarget)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MetricTarget{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.MountOptions",
		reflect.TypeOf((*MountOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.MountPropagation",
		reflect.TypeOf((*MountPropagation)(nil)).Elem(),
		map[string]interface{}{
			"NONE": MountPropagation_NONE,
			"HOST_TO_CONTAINER": MountPropagation_HOST_TO_CONTAINER,
			"BIDIRECTIONAL": MountPropagation_BIDIRECTIONAL,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.NamedNode",
		reflect.TypeOf((*NamedNode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_NamedNode{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Namespace",
		reflect.TypeOf((*Namespace)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toNamespaceSelectorConfig", GoMethod: "ToNamespaceSelectorConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toNetworkPolicyPeerConfig", GoMethod: "ToNetworkPolicyPeerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelector", GoMethod: "ToPodSelector"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Namespace{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INamespaceSelector)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INetworkPolicyPeer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.NamespaceProps",
		reflect.TypeOf((*NamespaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.NamespaceSelectorConfig",
		reflect.TypeOf((*NamespaceSelectorConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Namespaces",
		reflect.TypeOf((*Namespaces)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toNamespaceSelectorConfig", GoMethod: "ToNamespaceSelectorConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toNetworkPolicyPeerConfig", GoMethod: "ToNetworkPolicyPeerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelector", GoMethod: "ToPodSelector"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Namespaces{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INamespaceSelector)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INetworkPolicyPeer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.NamespacesSelectOptions",
		reflect.TypeOf((*NamespacesSelectOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.NetworkPolicy",
		reflect.TypeOf((*NetworkPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEgressRule", GoMethod: "AddEgressRule"},
			_jsii_.MemberMethod{JsiiMethod: "addIngressRule", GoMethod: "AddIngressRule"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkPolicy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.NetworkPolicyAddEgressRuleOptions",
		reflect.TypeOf((*NetworkPolicyAddEgressRuleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.NetworkPolicyIpBlock",
		reflect.TypeOf((*NetworkPolicyIpBlock)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cidr", GoGetter: "Cidr"},
			_jsii_.MemberProperty{JsiiProperty: "except", GoGetter: "Except"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toNetworkPolicyPeerConfig", GoMethod: "ToNetworkPolicyPeerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelector", GoMethod: "ToPodSelector"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkPolicyIpBlock{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INetworkPolicyPeer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.NetworkPolicyPeerConfig",
		reflect.TypeOf((*NetworkPolicyPeerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.NetworkPolicyPort",
		reflect.TypeOf((*NetworkPolicyPort)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_NetworkPolicyPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.NetworkPolicyPortProps",
		reflect.TypeOf((*NetworkPolicyPortProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.NetworkPolicyProps",
		reflect.TypeOf((*NetworkPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.NetworkPolicyRule",
		reflect.TypeOf((*NetworkPolicyRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.NetworkPolicyTraffic",
		reflect.TypeOf((*NetworkPolicyTraffic)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.NetworkPolicyTrafficDefault",
		reflect.TypeOf((*NetworkPolicyTrafficDefault)(nil)).Elem(),
		map[string]interface{}{
			"DENY": NetworkPolicyTrafficDefault_DENY,
			"ALLOW": NetworkPolicyTrafficDefault_ALLOW,
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.NetworkProtocol",
		reflect.TypeOf((*NetworkProtocol)(nil)).Elem(),
		map[string]interface{}{
			"TCP": NetworkProtocol_TCP,
			"UDP": NetworkProtocol_UDP,
			"SCTP": NetworkProtocol_SCTP,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.NfsVolumeOptions",
		reflect.TypeOf((*NfsVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Node",
		reflect.TypeOf((*Node)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Node{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.NodeLabelQuery",
		reflect.TypeOf((*NodeLabelQuery)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_NodeLabelQuery{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.NodeTaintQuery",
		reflect.TypeOf((*NodeTaintQuery)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_NodeTaintQuery{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.NodeTaintQueryOptions",
		reflect.TypeOf((*NodeTaintQueryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.NonApiResource",
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
		"cdk8s-plus-27.PathMapping",
		reflect.TypeOf((*PathMapping)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.PercentOrAbsolute",
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
		"cdk8s-plus-27.PersistentVolume",
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
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.PersistentVolumeAccessMode",
		reflect.TypeOf((*PersistentVolumeAccessMode)(nil)).Elem(),
		map[string]interface{}{
			"READ_WRITE_ONCE": PersistentVolumeAccessMode_READ_WRITE_ONCE,
			"READ_ONLY_MANY": PersistentVolumeAccessMode_READ_ONLY_MANY,
			"READ_WRITE_MANY": PersistentVolumeAccessMode_READ_WRITE_MANY,
			"READ_WRITE_ONCE_POD": PersistentVolumeAccessMode_READ_WRITE_ONCE_POD,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.PersistentVolumeClaim",
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
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.PersistentVolumeClaimProps",
		reflect.TypeOf((*PersistentVolumeClaimProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.PersistentVolumeClaimVolumeOptions",
		reflect.TypeOf((*PersistentVolumeClaimVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.PersistentVolumeMode",
		reflect.TypeOf((*PersistentVolumeMode)(nil)).Elem(),
		map[string]interface{}{
			"FILE_SYSTEM": PersistentVolumeMode_FILE_SYSTEM,
			"BLOCK": PersistentVolumeMode_BLOCK,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.PersistentVolumeProps",
		reflect.TypeOf((*PersistentVolumeProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.PersistentVolumeReclaimPolicy",
		reflect.TypeOf((*PersistentVolumeReclaimPolicy)(nil)).Elem(),
		map[string]interface{}{
			"RETAIN": PersistentVolumeReclaimPolicy_RETAIN,
			"DELETE": PersistentVolumeReclaimPolicy_DELETE,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Pod",
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
			_jsii_.MemberMethod{JsiiMethod: "attachContainer", GoMethod: "AttachContainer"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "hostNetwork", GoGetter: "HostNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "isolate", GoGetter: "Isolate"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "scheduling", GoGetter: "Scheduling"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "terminationGracePeriod", GoGetter: "TerminationGracePeriod"},
			_jsii_.MemberMethod{JsiiMethod: "toNetworkPolicyPeerConfig", GoMethod: "ToNetworkPolicyPeerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelector", GoMethod: "ToPodSelector"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelectorConfig", GoMethod: "ToPodSelectorConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toSubjectConfiguration", GoMethod: "ToSubjectConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_Pod{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AbstractPod)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.PodConnections",
		reflect.TypeOf((*PodConnections)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allowFrom", GoMethod: "AllowFrom"},
			_jsii_.MemberMethod{JsiiMethod: "allowTo", GoMethod: "AllowTo"},
			_jsii_.MemberProperty{JsiiProperty: "instance", GoGetter: "Instance"},
			_jsii_.MemberMethod{JsiiMethod: "isolate", GoMethod: "Isolate"},
		},
		func() interface{} {
			return &jsiiProxy_PodConnections{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.PodConnectionsAllowFromOptions",
		reflect.TypeOf((*PodConnectionsAllowFromOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.PodConnectionsAllowToOptions",
		reflect.TypeOf((*PodConnectionsAllowToOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.PodConnectionsIsolation",
		reflect.TypeOf((*PodConnectionsIsolation)(nil)).Elem(),
		map[string]interface{}{
			"POD": PodConnectionsIsolation_POD,
			"PEER": PodConnectionsIsolation_PEER,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.PodDns",
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
		"cdk8s-plus-27.PodDnsProps",
		reflect.TypeOf((*PodDnsProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.PodManagementPolicy",
		reflect.TypeOf((*PodManagementPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ORDERED_READY": PodManagementPolicy_ORDERED_READY,
			"PARALLEL": PodManagementPolicy_PARALLEL,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.PodProps",
		reflect.TypeOf((*PodProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.PodScheduling",
		reflect.TypeOf((*PodScheduling)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assign", GoMethod: "Assign"},
			_jsii_.MemberMethod{JsiiMethod: "attract", GoMethod: "Attract"},
			_jsii_.MemberMethod{JsiiMethod: "colocate", GoMethod: "Colocate"},
			_jsii_.MemberProperty{JsiiProperty: "instance", GoGetter: "Instance"},
			_jsii_.MemberMethod{JsiiMethod: "separate", GoMethod: "Separate"},
			_jsii_.MemberMethod{JsiiMethod: "tolerate", GoMethod: "Tolerate"},
		},
		func() interface{} {
			return &jsiiProxy_PodScheduling{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.PodSchedulingAttractOptions",
		reflect.TypeOf((*PodSchedulingAttractOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.PodSchedulingColocateOptions",
		reflect.TypeOf((*PodSchedulingColocateOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.PodSchedulingSeparateOptions",
		reflect.TypeOf((*PodSchedulingSeparateOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.PodSecurityContext",
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
		"cdk8s-plus-27.PodSecurityContextProps",
		reflect.TypeOf((*PodSecurityContextProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.PodSelectorConfig",
		reflect.TypeOf((*PodSelectorConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Pods",
		reflect.TypeOf((*Pods)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toNetworkPolicyPeerConfig", GoMethod: "ToNetworkPolicyPeerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelector", GoMethod: "ToPodSelector"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelectorConfig", GoMethod: "ToPodSelectorConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Pods{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPodSelector)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.PodsAllOptions",
		reflect.TypeOf((*PodsAllOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.PodsSelectOptions",
		reflect.TypeOf((*PodsSelectOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Probe",
		reflect.TypeOf((*Probe)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Probe{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ProbeOptions",
		reflect.TypeOf((*ProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.Protocol",
		reflect.TypeOf((*Protocol)(nil)).Elem(),
		map[string]interface{}{
			"TCP": Protocol_TCP,
			"UDP": Protocol_UDP,
			"SCTP": Protocol_SCTP,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Replicas",
		reflect.TypeOf((*Replicas)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Replicas{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Resource",
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
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.ResourceFieldPaths",
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
	_jsii_.RegisterClass(
		"cdk8s-plus-27.ResourcePermissions",
		reflect.TypeOf((*ResourcePermissions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberProperty{JsiiProperty: "instance", GoGetter: "Instance"},
		},
		func() interface{} {
			return &jsiiProxy_ResourcePermissions{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ResourceProps",
		reflect.TypeOf((*ResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.RestartPolicy",
		reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RestartPolicy_ALWAYS,
			"ON_FAILURE": RestartPolicy_ON_FAILURE,
			"NEVER": RestartPolicy_NEVER,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Role",
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
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.RoleBinding",
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
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.RoleBindingProps",
		reflect.TypeOf((*RoleBindingProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.RolePolicyRule",
		reflect.TypeOf((*RolePolicyRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.RoleProps",
		reflect.TypeOf((*RoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ScalingPolicy",
		reflect.TypeOf((*ScalingPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ScalingRules",
		reflect.TypeOf((*ScalingRules)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.ScalingStrategy",
		reflect.TypeOf((*ScalingStrategy)(nil)).Elem(),
		map[string]interface{}{
			"MAX_CHANGE": ScalingStrategy_MAX_CHANGE,
			"MIN_CHANGE": ScalingStrategy_MIN_CHANGE,
			"DISABLED": ScalingStrategy_DISABLED,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ScalingTarget",
		reflect.TypeOf((*ScalingTarget)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Secret",
		reflect.TypeOf((*Secret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStringData", GoMethod: "AddStringData"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "envValue", GoMethod: "EnvValue"},
			_jsii_.MemberMethod{JsiiMethod: "getStringData", GoMethod: "GetStringData"},
			_jsii_.MemberProperty{JsiiProperty: "immutable", GoGetter: "Immutable"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.SecretProps",
		reflect.TypeOf((*SecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.SecretValue",
		reflect.TypeOf((*SecretValue)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.SecretVolumeOptions",
		reflect.TypeOf((*SecretVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Service",
		reflect.TypeOf((*Service)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIP", GoGetter: "ClusterIP"},
			_jsii_.MemberMethod{JsiiMethod: "exposeViaIngress", GoMethod: "ExposeViaIngress"},
			_jsii_.MemberProperty{JsiiProperty: "externalName", GoGetter: "ExternalName"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "ports", GoGetter: "Ports"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "select", GoMethod: "Select"},
			_jsii_.MemberMethod{JsiiMethod: "selectLabel", GoMethod: "SelectLabel"},
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
		"cdk8s-plus-27.ServiceAccount",
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
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "secrets", GoGetter: "Secrets"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toSubjectConfiguration", GoMethod: "ToSubjectConfiguration"},
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
		"cdk8s-plus-27.ServiceAccountProps",
		reflect.TypeOf((*ServiceAccountProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.ServiceAccountTokenSecret",
		reflect.TypeOf((*ServiceAccountTokenSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStringData", GoMethod: "AddStringData"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "envValue", GoMethod: "EnvValue"},
			_jsii_.MemberMethod{JsiiMethod: "getStringData", GoMethod: "GetStringData"},
			_jsii_.MemberProperty{JsiiProperty: "immutable", GoGetter: "Immutable"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.ServiceAccountTokenSecretProps",
		reflect.TypeOf((*ServiceAccountTokenSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ServiceBindOptions",
		reflect.TypeOf((*ServiceBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ServiceIngressBackendOptions",
		reflect.TypeOf((*ServiceIngressBackendOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ServicePort",
		reflect.TypeOf((*ServicePort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.ServiceProps",
		reflect.TypeOf((*ServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.ServiceType",
		reflect.TypeOf((*ServiceType)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER_IP": ServiceType_CLUSTER_IP,
			"NODE_PORT": ServiceType_NODE_PORT,
			"LOAD_BALANCER": ServiceType_LOAD_BALANCER,
			"EXTERNAL_NAME": ServiceType_EXTERNAL_NAME,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.SshAuthSecret",
		reflect.TypeOf((*SshAuthSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStringData", GoMethod: "AddStringData"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "envValue", GoMethod: "EnvValue"},
			_jsii_.MemberMethod{JsiiMethod: "getStringData", GoMethod: "GetStringData"},
			_jsii_.MemberProperty{JsiiProperty: "immutable", GoGetter: "Immutable"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.SshAuthSecretProps",
		reflect.TypeOf((*SshAuthSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.StatefulSet",
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
			_jsii_.MemberMethod{JsiiMethod: "attachContainer", GoMethod: "AttachContainer"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberProperty{JsiiProperty: "hasAutoscaler", GoGetter: "HasAutoscaler"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "hostNetwork", GoGetter: "HostNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "isolate", GoGetter: "Isolate"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberMethod{JsiiMethod: "markHasAutoscaler", GoMethod: "MarkHasAutoscaler"},
			_jsii_.MemberProperty{JsiiProperty: "matchExpressions", GoGetter: "MatchExpressions"},
			_jsii_.MemberProperty{JsiiProperty: "matchLabels", GoGetter: "MatchLabels"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "minReady", GoGetter: "MinReady"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "podManagementPolicy", GoGetter: "PodManagementPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "replicas", GoGetter: "Replicas"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "scheduling", GoGetter: "Scheduling"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberMethod{JsiiMethod: "select", GoMethod: "Select"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberProperty{JsiiProperty: "terminationGracePeriod", GoGetter: "TerminationGracePeriod"},
			_jsii_.MemberMethod{JsiiMethod: "toNetworkPolicyPeerConfig", GoMethod: "ToNetworkPolicyPeerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelector", GoMethod: "ToPodSelector"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelectorConfig", GoMethod: "ToPodSelectorConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toScalingTarget", GoMethod: "ToScalingTarget"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toSubjectConfiguration", GoMethod: "ToSubjectConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_StatefulSet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Workload)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IScalable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.StatefulSetProps",
		reflect.TypeOf((*StatefulSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.StatefulSetUpdateStrategy",
		reflect.TypeOf((*StatefulSetUpdateStrategy)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_StatefulSetUpdateStrategy{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.StatefulSetUpdateStrategyRollingUpdateOptions",
		reflect.TypeOf((*StatefulSetUpdateStrategyRollingUpdateOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.SubjectConfiguration",
		reflect.TypeOf((*SubjectConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.Sysctl",
		reflect.TypeOf((*Sysctl)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-27.TaintEffect",
		reflect.TypeOf((*TaintEffect)(nil)).Elem(),
		map[string]interface{}{
			"NO_SCHEDULE": TaintEffect_NO_SCHEDULE,
			"PREFER_NO_SCHEDULE": TaintEffect_PREFER_NO_SCHEDULE,
			"NO_EXECUTE": TaintEffect_NO_EXECUTE,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.TaintedNode",
		reflect.TypeOf((*TaintedNode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "taintSelector", GoGetter: "TaintSelector"},
		},
		func() interface{} {
			return &jsiiProxy_TaintedNode{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.TcpSocketProbeOptions",
		reflect.TypeOf((*TcpSocketProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.TlsSecret",
		reflect.TypeOf((*TlsSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStringData", GoMethod: "AddStringData"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "asApiResource", GoMethod: "AsApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "asNonApiResource", GoMethod: "AsNonApiResource"},
			_jsii_.MemberMethod{JsiiMethod: "envValue", GoMethod: "EnvValue"},
			_jsii_.MemberMethod{JsiiMethod: "getStringData", GoMethod: "GetStringData"},
			_jsii_.MemberProperty{JsiiProperty: "immutable", GoGetter: "Immutable"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
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
		"cdk8s-plus-27.TlsSecretProps",
		reflect.TypeOf((*TlsSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Topology",
		reflect.TypeOf((*Topology)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
		},
		func() interface{} {
			return &jsiiProxy_Topology{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.User",
		reflect.TypeOf((*User)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toSubjectConfiguration", GoMethod: "ToSubjectConfiguration"},
		},
		func() interface{} {
			j := jsiiProxy_User{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISubject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Volume",
		reflect.TypeOf((*Volume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "asVolume", GoMethod: "AsVolume"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Volume{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IStorage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.VolumeMount",
		reflect.TypeOf((*VolumeMount)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.Workload",
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
			_jsii_.MemberMethod{JsiiMethod: "attachContainer", GoMethod: "AttachContainer"},
			_jsii_.MemberProperty{JsiiProperty: "automountServiceAccountToken", GoGetter: "AutomountServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryAuth", GoGetter: "DockerRegistryAuth"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "hostNetwork", GoGetter: "HostNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "isolate", GoGetter: "Isolate"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "matchExpressions", GoGetter: "MatchExpressions"},
			_jsii_.MemberProperty{JsiiProperty: "matchLabels", GoGetter: "MatchLabels"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "scheduling", GoGetter: "Scheduling"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberMethod{JsiiMethod: "select", GoMethod: "Select"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "terminationGracePeriod", GoGetter: "TerminationGracePeriod"},
			_jsii_.MemberMethod{JsiiMethod: "toNetworkPolicyPeerConfig", GoMethod: "ToNetworkPolicyPeerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelector", GoMethod: "ToPodSelector"},
			_jsii_.MemberMethod{JsiiMethod: "toPodSelectorConfig", GoMethod: "ToPodSelectorConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toSubjectConfiguration", GoMethod: "ToSubjectConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_Workload{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AbstractPod)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.WorkloadProps",
		reflect.TypeOf((*WorkloadProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-27.WorkloadScheduling",
		reflect.TypeOf((*WorkloadScheduling)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assign", GoMethod: "Assign"},
			_jsii_.MemberMethod{JsiiMethod: "attract", GoMethod: "Attract"},
			_jsii_.MemberMethod{JsiiMethod: "colocate", GoMethod: "Colocate"},
			_jsii_.MemberProperty{JsiiProperty: "instance", GoGetter: "Instance"},
			_jsii_.MemberMethod{JsiiMethod: "separate", GoMethod: "Separate"},
			_jsii_.MemberMethod{JsiiMethod: "spread", GoMethod: "Spread"},
			_jsii_.MemberMethod{JsiiMethod: "tolerate", GoMethod: "Tolerate"},
		},
		func() interface{} {
			j := jsiiProxy_WorkloadScheduling{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PodScheduling)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-27.WorkloadSchedulingSpreadOptions",
		reflect.TypeOf((*WorkloadSchedulingSpreadOptions)(nil)).Elem(),
	)
}
