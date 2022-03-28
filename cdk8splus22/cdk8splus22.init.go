package cdk8splus22

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.AddDeploymentOptions",
		reflect.TypeOf((*AddDeploymentOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.AddDirectoryOptions",
		reflect.TypeOf((*AddDirectoryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.CommandProbeOptions",
		reflect.TypeOf((*CommandProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.ConfigMap",
		reflect.TypeOf((*ConfigMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBinaryData", GoMethod: "AddBinaryData"},
			_jsii_.MemberMethod{JsiiMethod: "addData", GoMethod: "AddData"},
			_jsii_.MemberMethod{JsiiMethod: "addDirectory", GoMethod: "AddDirectory"},
			_jsii_.MemberMethod{JsiiMethod: "addFile", GoMethod: "AddFile"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "binaryData", GoGetter: "BinaryData"},
			_jsii_.MemberProperty{JsiiProperty: "data", GoGetter: "Data"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
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
		"cdk8s-plus-22.ConfigMapProps",
		reflect.TypeOf((*ConfigMapProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.ConfigMapVolumeOptions",
		reflect.TypeOf((*ConfigMapVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.Container",
		reflect.TypeOf((*Container)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEnv", GoMethod: "AddEnv"},
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
		"cdk8s-plus-22.ContainerLifecycle",
		reflect.TypeOf((*ContainerLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.ContainerProps",
		reflect.TypeOf((*ContainerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.ContainerSecurityContext",
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
		"cdk8s-plus-22.ContainerSecurityContextProps",
		reflect.TypeOf((*ContainerSecurityContextProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.Cpu",
		reflect.TypeOf((*Cpu)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "amount", GoGetter: "Amount"},
		},
		func() interface{} {
			return &jsiiProxy_Cpu{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.CpuResources",
		reflect.TypeOf((*CpuResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.Deployment",
		reflect.TypeOf((*Deployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addHostAlias", GoMethod: "AddHostAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberMethod{JsiiMethod: "exposeViaIngress", GoMethod: "ExposeViaIngress"},
			_jsii_.MemberMethod{JsiiMethod: "exposeViaService", GoMethod: "ExposeViaService"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "labelSelector", GoGetter: "LabelSelector"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "replicas", GoGetter: "Replicas"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberMethod{JsiiMethod: "selectByLabel", GoMethod: "SelectByLabel"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_Deployment{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPodTemplate)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.DeploymentProps",
		reflect.TypeOf((*DeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-22.EmptyDirMedium",
		reflect.TypeOf((*EmptyDirMedium)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": EmptyDirMedium_DEFAULT,
			"MEMORY": EmptyDirMedium_MEMORY,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.EmptyDirVolumeOptions",
		reflect.TypeOf((*EmptyDirVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-22.EnvFieldPaths",
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
		"cdk8s-plus-22.EnvValue",
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
		"cdk8s-plus-22.EnvValueFromConfigMapOptions",
		reflect.TypeOf((*EnvValueFromConfigMapOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.EnvValueFromFieldRefOptions",
		reflect.TypeOf((*EnvValueFromFieldRefOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.EnvValueFromProcessOptions",
		reflect.TypeOf((*EnvValueFromProcessOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.EnvValueFromResourceOptions",
		reflect.TypeOf((*EnvValueFromResourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.EnvValueFromSecretOptions",
		reflect.TypeOf((*EnvValueFromSecretOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.ExposeDeploymentViaIngressOptions",
		reflect.TypeOf((*ExposeDeploymentViaIngressOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.ExposeDeploymentViaServiceOptions",
		reflect.TypeOf((*ExposeDeploymentViaServiceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.ExposeServiceViaIngressOptions",
		reflect.TypeOf((*ExposeServiceViaIngressOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-22.FsGroupChangePolicy",
		reflect.TypeOf((*FsGroupChangePolicy)(nil)).Elem(),
		map[string]interface{}{
			"ON_ROOT_MISMATCH": FsGroupChangePolicy_ON_ROOT_MISMATCH,
			"ALWAYS": FsGroupChangePolicy_ALWAYS,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.Handler",
		reflect.TypeOf((*Handler)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Handler{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.HandlerFromHttpGetOptions",
		reflect.TypeOf((*HandlerFromHttpGetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.HandlerFromTcpSocketOptions",
		reflect.TypeOf((*HandlerFromTcpSocketOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.HostAlias",
		reflect.TypeOf((*HostAlias)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.HttpGetProbeOptions",
		reflect.TypeOf((*HttpGetProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-22.HttpIngressPathType",
		reflect.TypeOf((*HttpIngressPathType)(nil)).Elem(),
		map[string]interface{}{
			"PREFIX": HttpIngressPathType_PREFIX,
			"EXACT": HttpIngressPathType_EXACT,
			"IMPLEMENTATION_SPECIFIC": HttpIngressPathType_IMPLEMENTATION_SPECIFIC,
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-22.IConfigMap",
		reflect.TypeOf((*IConfigMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_IConfigMap{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-22.IPodSpec",
		reflect.TypeOf((*IPodSpec)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			return &jsiiProxy_IPodSpec{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-22.IPodTemplate",
		reflect.TypeOf((*IPodTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_IPodTemplate{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPodSpec)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-22.IResource",
		reflect.TypeOf((*IResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_IResource{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-22.ISecret",
		reflect.TypeOf((*ISecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_ISecret{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-22.IServiceAccount",
		reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_IServiceAccount{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-22.ImagePullPolicy",
		reflect.TypeOf((*ImagePullPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": ImagePullPolicy_ALWAYS,
			"IF_NOT_PRESENT": ImagePullPolicy_IF_NOT_PRESENT,
			"NEVER": ImagePullPolicy_NEVER,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.Ingress",
		reflect.TypeOf((*Ingress)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultBackend", GoMethod: "AddDefaultBackend"},
			_jsii_.MemberMethod{JsiiMethod: "addHostDefaultBackend", GoMethod: "AddHostDefaultBackend"},
			_jsii_.MemberMethod{JsiiMethod: "addHostRule", GoMethod: "AddHostRule"},
			_jsii_.MemberMethod{JsiiMethod: "addRule", GoMethod: "AddRule"},
			_jsii_.MemberMethod{JsiiMethod: "addRules", GoMethod: "AddRules"},
			_jsii_.MemberMethod{JsiiMethod: "addTls", GoMethod: "AddTls"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ingress{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.IngressBackend",
		reflect.TypeOf((*IngressBackend)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IngressBackend{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.IngressProps",
		reflect.TypeOf((*IngressProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.IngressRule",
		reflect.TypeOf((*IngressRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.IngressTls",
		reflect.TypeOf((*IngressTls)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.Job",
		reflect.TypeOf((*Job)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activeDeadline", GoGetter: "ActiveDeadline"},
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addHostAlias", GoMethod: "AddHostAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "backoffLimit", GoGetter: "BackoffLimit"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "ttlAfterFinished", GoGetter: "TtlAfterFinished"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_Job{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPodTemplate)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.JobProps",
		reflect.TypeOf((*JobProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.MemoryResources",
		reflect.TypeOf((*MemoryResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.MountOptions",
		reflect.TypeOf((*MountOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-22.MountPropagation",
		reflect.TypeOf((*MountPropagation)(nil)).Elem(),
		map[string]interface{}{
			"NONE": MountPropagation_NONE,
			"HOST_TO_CONTAINER": MountPropagation_HOST_TO_CONTAINER,
			"BIDIRECTIONAL": MountPropagation_BIDIRECTIONAL,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.PathMapping",
		reflect.TypeOf((*PathMapping)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.Pod",
		reflect.TypeOf((*Pod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addHostAlias", GoMethod: "AddHostAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_Pod{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPodSpec)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-22.PodManagementPolicy",
		reflect.TypeOf((*PodManagementPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ORDERED_READY": PodManagementPolicy_ORDERED_READY,
			"PARALLEL": PodManagementPolicy_PARALLEL,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.PodProps",
		reflect.TypeOf((*PodProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.PodSecurityContext",
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
		"cdk8s-plus-22.PodSecurityContextProps",
		reflect.TypeOf((*PodSecurityContextProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.PodSpec",
		reflect.TypeOf((*PodSpec)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addHostAlias", GoMethod: "AddHostAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_PodSpec{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPodSpec)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.PodSpecProps",
		reflect.TypeOf((*PodSpecProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.PodTemplate",
		reflect.TypeOf((*PodTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addHostAlias", GoMethod: "AddHostAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_PodTemplate{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PodSpec)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPodTemplate)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.PodTemplateProps",
		reflect.TypeOf((*PodTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.Probe",
		reflect.TypeOf((*Probe)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Probe{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.ProbeOptions",
		reflect.TypeOf((*ProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-22.Protocol",
		reflect.TypeOf((*Protocol)(nil)).Elem(),
		map[string]interface{}{
			"TCP": Protocol_TCP,
			"UDP": Protocol_UDP,
			"SCTP": Protocol_SCTP,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.Resource",
		reflect.TypeOf((*Resource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Resource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-22.ResourceFieldPaths",
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
		"cdk8s-plus-22.ResourceProps",
		reflect.TypeOf((*ResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.Resources",
		reflect.TypeOf((*Resources)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-22.RestartPolicy",
		reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RestartPolicy_ALWAYS,
			"ON_FAILURE": RestartPolicy_ON_FAILURE,
			"NEVER": RestartPolicy_NEVER,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.Secret",
		reflect.TypeOf((*Secret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStringData", GoMethod: "AddStringData"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberMethod{JsiiMethod: "getStringData", GoMethod: "GetStringData"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
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
		"cdk8s-plus-22.SecretProps",
		reflect.TypeOf((*SecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.SecretValue",
		reflect.TypeOf((*SecretValue)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.SecretVolumeOptions",
		reflect.TypeOf((*SecretVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.Service",
		reflect.TypeOf((*Service)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeployment", GoMethod: "AddDeployment"},
			_jsii_.MemberMethod{JsiiMethod: "addSelector", GoMethod: "AddSelector"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIP", GoGetter: "ClusterIP"},
			_jsii_.MemberMethod{JsiiMethod: "exposeViaIngress", GoMethod: "ExposeViaIngress"},
			_jsii_.MemberProperty{JsiiProperty: "externalName", GoGetter: "ExternalName"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "ports", GoGetter: "Ports"},
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
		"cdk8s-plus-22.ServiceAccount",
		reflect.TypeOf((*ServiceAccount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSecret", GoMethod: "AddSecret"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "secrets", GoGetter: "Secrets"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceAccount{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServiceAccount)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.ServiceAccountProps",
		reflect.TypeOf((*ServiceAccountProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.ServiceIngressBackendOptions",
		reflect.TypeOf((*ServiceIngressBackendOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.ServicePort",
		reflect.TypeOf((*ServicePort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.ServicePortOptions",
		reflect.TypeOf((*ServicePortOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.ServiceProps",
		reflect.TypeOf((*ServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-22.ServiceType",
		reflect.TypeOf((*ServiceType)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER_IP": ServiceType_CLUSTER_IP,
			"NODE_PORT": ServiceType_NODE_PORT,
			"LOAD_BALANCER": ServiceType_LOAD_BALANCER,
			"EXTERNAL_NAME": ServiceType_EXTERNAL_NAME,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.StatefulSet",
		reflect.TypeOf((*StatefulSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addHostAlias", GoMethod: "AddHostAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addInitContainer", GoMethod: "AddInitContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "apiObject", GoGetter: "ApiObject"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "hostAliases", GoGetter: "HostAliases"},
			_jsii_.MemberProperty{JsiiProperty: "initContainers", GoGetter: "InitContainers"},
			_jsii_.MemberProperty{JsiiProperty: "labelSelector", GoGetter: "LabelSelector"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "podManagementPolicy", GoGetter: "PodManagementPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "podMetadata", GoGetter: "PodMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "replicas", GoGetter: "Replicas"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityContext", GoGetter: "SecurityContext"},
			_jsii_.MemberMethod{JsiiMethod: "selectByLabel", GoMethod: "SelectByLabel"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_StatefulSet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPodTemplate)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.StatefulSetProps",
		reflect.TypeOf((*StatefulSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.Sysctl",
		reflect.TypeOf((*Sysctl)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.TcpSocketProbeOptions",
		reflect.TypeOf((*TcpSocketProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-22.Volume",
		reflect.TypeOf((*Volume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_Volume{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-22.VolumeMount",
		reflect.TypeOf((*VolumeMount)(nil)).Elem(),
	)
}
