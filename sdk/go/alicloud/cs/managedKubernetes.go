// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource will help you to manage a ManagedKubernetes Cluster in Alibaba Cloud Kubernetes Service.
//
// > **NOTE:** Available since v1.26.0.
//
// > **NOTE:** It is recommended to create a cluster with zero worker nodes, and then use a node pool to manage the cluster nodes.
//
// > **NOTE:** Kubernetes cluster only supports VPC network and it can access internet while creating kubernetes cluster.
// A Nat Gateway and configuring a SNAT for it can ensure one VPC network access internet. If there is no nat gateway in the
// VPC, you can set `newNatGateway` to "true" to create one automatically.
//
// > **NOTE:** Creating kubernetes cluster need to install several packages and it will cost about 15 minutes. Please be patient.
//
// > **NOTE:** From version 1.9.4, the provider supports to download kube config, client certificate, client key and cluster ca certificate
// after creating cluster successfully, and you can put them into the specified location, like '~/.kube/config'.
//
// > **NOTE:** From version 1.20.0, the provider supports disabling internet load balancer for API Server by setting `false` to `slbInternetEnabled`.
//
// > **NOTE:** If you want to manage Kubernetes, you can use Kubernetes Provider.
//
// > **NOTE:** You need to activate several other products and confirm Authorization Policy used by Container Service before using this resource.
// Please refer to the `Authorization management` and `Cluster management` sections in the [Document Center](https://www.alibabacloud.com/help/doc-detail/86488.htm).
//
// > **NOTE:** From version 1.72.0, Some parameters have been removed from resource,You can check them below and re-import the cluster if necessary.
//
// > **NOTE:** From version 1.120.0, Support for cluster migration from Standard cluster to professional.
//
// > **NOTE:** From version 1.177.0+, `runtime`,`enableSsh`,`rdsInstances`,`excludeAutoscalerNodes`,`workerNumber`,`workerInstanceTypes`,`password`,`keyName`,`kmsEncryptedPassword`,`kmsEncryptionContext`,`workerInstanceChargeType`,`workerPeriod`,`workerPeriodUnit`,`workerAutoRenew`,`workerAutoRenewPeriod`,`workerDiskCategory`,`workerDiskSize`,`workerDataDisks`,`nodeNameMode`,`nodePortRange`,`osType`,`platform`,`imageId`,`cpuPolicy`,`userData`,`taints`,`workerDiskPerformanceLevel`,`workerDiskSnapshotPolicyId`,`installCloudMonitor` are deprecated.
// We Suggest you using resource **`cs.NodePool`** to manage your cluster worker nodes.
//
// > **NOTE:** From version 1.212.0, `runtime`,`enableSsh`,`rdsInstances`,`excludeAutoscalerNodes`,`workerNumber`,`workerInstanceTypes`,`password`,`keyName`,`kmsEncryptedPassword`,`kmsEncryptionContext`,`workerInstanceChargeType`,`workerPeriod`,`workerPeriodUnit`,`workerAutoRenew`,`workerAutoRenewPeriod`,`workerDiskCategory`,`workerDiskSize`,`workerDataDisks`,`nodeNameMode`,`nodePortRange`,`osType`,`platform`,`imageId`,`cpuPolicy`,`userData`,`taints`,`workerDiskPerformanceLevel`,`workerDiskSnapshotPolicyId`,`installCloudMonitor`,`kubeConfig`,`availabilityZone` are removed.
// Please use resource **`cs.NodePool`** to manage your cluster worker nodes.
//
// ## Import
//
// Kubernetes managed cluster can be imported using the id, e.g. Then complete the main.tf accords to the result of `pulumi preview`.
//
// ```sh
// $ pulumi import alicloud:cs/managedKubernetes:ManagedKubernetes main cluster_id
// ```
type ManagedKubernetes struct {
	pulumi.CustomResourceState

	Addons       ManagedKubernetesAddonArrayOutput `pulumi:"addons"`
	ApiAudiences pulumi.StringArrayOutput          `pulumi:"apiAudiences"`
	// (Available in 1.105.0+) Nested attribute containing certificate authority data for your cluster.
	CertificateAuthority ManagedKubernetesCertificateAuthorityOutput `pulumi:"certificateAuthority"`
	// The base64 encoded client certificate data required to communicate with your cluster. Add this to the client-certificate-data section of the kubeconfig file for your cluster.
	ClientCert pulumi.StringPtrOutput `pulumi:"clientCert"`
	// The base64 encoded client key data required to communicate with your cluster. Add this to the client-key-data section of the kubeconfig file for your cluster.
	ClientKey     pulumi.StringPtrOutput `pulumi:"clientKey"`
	ClusterCaCert pulumi.StringPtrOutput `pulumi:"clusterCaCert"`
	// cluster local domain
	ClusterDomain pulumi.StringPtrOutput `pulumi:"clusterDomain"`
	ClusterSpec   pulumi.StringOutput    `pulumi:"clusterSpec"`
	// Map of kubernetes cluster connection information.
	Connections               ManagedKubernetesConnectionsOutput       `pulumi:"connections"`
	ControlPlaneLogComponents pulumi.StringArrayOutput                 `pulumi:"controlPlaneLogComponents"`
	ControlPlaneLogProject    pulumi.StringOutput                      `pulumi:"controlPlaneLogProject"`
	ControlPlaneLogTtl        pulumi.StringOutput                      `pulumi:"controlPlaneLogTtl"`
	CustomSan                 pulumi.StringPtrOutput                   `pulumi:"customSan"`
	DeleteOptions             ManagedKubernetesDeleteOptionArrayOutput `pulumi:"deleteOptions"`
	DeletionProtection        pulumi.BoolPtrOutput                     `pulumi:"deletionProtection"`
	EnableRrsa                pulumi.BoolPtrOutput                     `pulumi:"enableRrsa"`
	// disk encryption key, only in ack-pro
	EncryptionProviderKey     pulumi.StringPtrOutput                   `pulumi:"encryptionProviderKey"`
	IsEnterpriseSecurityGroup pulumi.BoolOutput                        `pulumi:"isEnterpriseSecurityGroup"`
	LoadBalancerSpec          pulumi.StringPtrOutput                   `pulumi:"loadBalancerSpec"`
	MaintenanceWindow         ManagedKubernetesMaintenanceWindowOutput `pulumi:"maintenanceWindow"`
	// Node name.
	Name       pulumi.StringOutput    `pulumi:"name"`
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// The ID of nat gateway used to launch kubernetes cluster.
	NatGatewayId    pulumi.StringOutput      `pulumi:"natGatewayId"`
	NewNatGateway   pulumi.BoolPtrOutput     `pulumi:"newNatGateway"`
	NodeCidrMask    pulumi.IntPtrOutput      `pulumi:"nodeCidrMask"`
	PodCidr         pulumi.StringPtrOutput   `pulumi:"podCidr"`
	PodVswitchIds   pulumi.StringArrayOutput `pulumi:"podVswitchIds"`
	ProxyMode       pulumi.StringPtrOutput   `pulumi:"proxyMode"`
	ResourceGroupId pulumi.StringOutput      `pulumi:"resourceGroupId"`
	RetainResources pulumi.StringArrayOutput `pulumi:"retainResources"`
	// (Optional, Available in v1.185.0+) Nested attribute containing RRSA related data for your cluster.
	RrsaMetadata         ManagedKubernetesRrsaMetadataOutput `pulumi:"rrsaMetadata"`
	SecurityGroupId      pulumi.StringOutput                 `pulumi:"securityGroupId"`
	ServiceAccountIssuer pulumi.StringPtrOutput              `pulumi:"serviceAccountIssuer"`
	ServiceCidr          pulumi.StringPtrOutput              `pulumi:"serviceCidr"`
	// The ID of APIServer load balancer.
	SlbId pulumi.StringOutput `pulumi:"slbId"`
	// The public ip of load balancer.
	SlbInternet        pulumi.StringOutput  `pulumi:"slbInternet"`
	SlbInternetEnabled pulumi.BoolPtrOutput `pulumi:"slbInternetEnabled"`
	// The ID of private load balancer where the current cluster master node is located.
	SlbIntranet pulumi.StringOutput    `pulumi:"slbIntranet"`
	Tags        pulumi.MapOutput       `pulumi:"tags"`
	Timezone    pulumi.StringPtrOutput `pulumi:"timezone"`
	UserCa      pulumi.StringPtrOutput `pulumi:"userCa"`
	Version     pulumi.StringOutput    `pulumi:"version"`
	// The ID of VPC where the current cluster is located.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The RamRole Name attached to worker node.
	WorkerRamRoleName pulumi.StringOutput      `pulumi:"workerRamRoleName"`
	WorkerVswitchIds  pulumi.StringArrayOutput `pulumi:"workerVswitchIds"`
}

// NewManagedKubernetes registers a new resource with the given unique name, arguments, and options.
func NewManagedKubernetes(ctx *pulumi.Context,
	name string, args *ManagedKubernetesArgs, opts ...pulumi.ResourceOption) (*ManagedKubernetes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.WorkerVswitchIds == nil {
		return nil, errors.New("invalid value for required argument 'WorkerVswitchIds'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ManagedKubernetes
	err := ctx.RegisterResource("alicloud:cs/managedKubernetes:ManagedKubernetes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedKubernetes gets an existing ManagedKubernetes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedKubernetes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedKubernetesState, opts ...pulumi.ResourceOption) (*ManagedKubernetes, error) {
	var resource ManagedKubernetes
	err := ctx.ReadResource("alicloud:cs/managedKubernetes:ManagedKubernetes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedKubernetes resources.
type managedKubernetesState struct {
	Addons       []ManagedKubernetesAddon `pulumi:"addons"`
	ApiAudiences []string                 `pulumi:"apiAudiences"`
	// (Available in 1.105.0+) Nested attribute containing certificate authority data for your cluster.
	CertificateAuthority *ManagedKubernetesCertificateAuthority `pulumi:"certificateAuthority"`
	// The base64 encoded client certificate data required to communicate with your cluster. Add this to the client-certificate-data section of the kubeconfig file for your cluster.
	ClientCert *string `pulumi:"clientCert"`
	// The base64 encoded client key data required to communicate with your cluster. Add this to the client-key-data section of the kubeconfig file for your cluster.
	ClientKey     *string `pulumi:"clientKey"`
	ClusterCaCert *string `pulumi:"clusterCaCert"`
	// cluster local domain
	ClusterDomain *string `pulumi:"clusterDomain"`
	ClusterSpec   *string `pulumi:"clusterSpec"`
	// Map of kubernetes cluster connection information.
	Connections               *ManagedKubernetesConnections   `pulumi:"connections"`
	ControlPlaneLogComponents []string                        `pulumi:"controlPlaneLogComponents"`
	ControlPlaneLogProject    *string                         `pulumi:"controlPlaneLogProject"`
	ControlPlaneLogTtl        *string                         `pulumi:"controlPlaneLogTtl"`
	CustomSan                 *string                         `pulumi:"customSan"`
	DeleteOptions             []ManagedKubernetesDeleteOption `pulumi:"deleteOptions"`
	DeletionProtection        *bool                           `pulumi:"deletionProtection"`
	EnableRrsa                *bool                           `pulumi:"enableRrsa"`
	// disk encryption key, only in ack-pro
	EncryptionProviderKey     *string                             `pulumi:"encryptionProviderKey"`
	IsEnterpriseSecurityGroup *bool                               `pulumi:"isEnterpriseSecurityGroup"`
	LoadBalancerSpec          *string                             `pulumi:"loadBalancerSpec"`
	MaintenanceWindow         *ManagedKubernetesMaintenanceWindow `pulumi:"maintenanceWindow"`
	// Node name.
	Name       *string `pulumi:"name"`
	NamePrefix *string `pulumi:"namePrefix"`
	// The ID of nat gateway used to launch kubernetes cluster.
	NatGatewayId    *string  `pulumi:"natGatewayId"`
	NewNatGateway   *bool    `pulumi:"newNatGateway"`
	NodeCidrMask    *int     `pulumi:"nodeCidrMask"`
	PodCidr         *string  `pulumi:"podCidr"`
	PodVswitchIds   []string `pulumi:"podVswitchIds"`
	ProxyMode       *string  `pulumi:"proxyMode"`
	ResourceGroupId *string  `pulumi:"resourceGroupId"`
	RetainResources []string `pulumi:"retainResources"`
	// (Optional, Available in v1.185.0+) Nested attribute containing RRSA related data for your cluster.
	RrsaMetadata         *ManagedKubernetesRrsaMetadata `pulumi:"rrsaMetadata"`
	SecurityGroupId      *string                        `pulumi:"securityGroupId"`
	ServiceAccountIssuer *string                        `pulumi:"serviceAccountIssuer"`
	ServiceCidr          *string                        `pulumi:"serviceCidr"`
	// The ID of APIServer load balancer.
	SlbId *string `pulumi:"slbId"`
	// The public ip of load balancer.
	SlbInternet        *string `pulumi:"slbInternet"`
	SlbInternetEnabled *bool   `pulumi:"slbInternetEnabled"`
	// The ID of private load balancer where the current cluster master node is located.
	SlbIntranet *string                `pulumi:"slbIntranet"`
	Tags        map[string]interface{} `pulumi:"tags"`
	Timezone    *string                `pulumi:"timezone"`
	UserCa      *string                `pulumi:"userCa"`
	Version     *string                `pulumi:"version"`
	// The ID of VPC where the current cluster is located.
	VpcId *string `pulumi:"vpcId"`
	// The RamRole Name attached to worker node.
	WorkerRamRoleName *string  `pulumi:"workerRamRoleName"`
	WorkerVswitchIds  []string `pulumi:"workerVswitchIds"`
}

type ManagedKubernetesState struct {
	Addons       ManagedKubernetesAddonArrayInput
	ApiAudiences pulumi.StringArrayInput
	// (Available in 1.105.0+) Nested attribute containing certificate authority data for your cluster.
	CertificateAuthority ManagedKubernetesCertificateAuthorityPtrInput
	// The base64 encoded client certificate data required to communicate with your cluster. Add this to the client-certificate-data section of the kubeconfig file for your cluster.
	ClientCert pulumi.StringPtrInput
	// The base64 encoded client key data required to communicate with your cluster. Add this to the client-key-data section of the kubeconfig file for your cluster.
	ClientKey     pulumi.StringPtrInput
	ClusterCaCert pulumi.StringPtrInput
	// cluster local domain
	ClusterDomain pulumi.StringPtrInput
	ClusterSpec   pulumi.StringPtrInput
	// Map of kubernetes cluster connection information.
	Connections               ManagedKubernetesConnectionsPtrInput
	ControlPlaneLogComponents pulumi.StringArrayInput
	ControlPlaneLogProject    pulumi.StringPtrInput
	ControlPlaneLogTtl        pulumi.StringPtrInput
	CustomSan                 pulumi.StringPtrInput
	DeleteOptions             ManagedKubernetesDeleteOptionArrayInput
	DeletionProtection        pulumi.BoolPtrInput
	EnableRrsa                pulumi.BoolPtrInput
	// disk encryption key, only in ack-pro
	EncryptionProviderKey     pulumi.StringPtrInput
	IsEnterpriseSecurityGroup pulumi.BoolPtrInput
	LoadBalancerSpec          pulumi.StringPtrInput
	MaintenanceWindow         ManagedKubernetesMaintenanceWindowPtrInput
	// Node name.
	Name       pulumi.StringPtrInput
	NamePrefix pulumi.StringPtrInput
	// The ID of nat gateway used to launch kubernetes cluster.
	NatGatewayId    pulumi.StringPtrInput
	NewNatGateway   pulumi.BoolPtrInput
	NodeCidrMask    pulumi.IntPtrInput
	PodCidr         pulumi.StringPtrInput
	PodVswitchIds   pulumi.StringArrayInput
	ProxyMode       pulumi.StringPtrInput
	ResourceGroupId pulumi.StringPtrInput
	RetainResources pulumi.StringArrayInput
	// (Optional, Available in v1.185.0+) Nested attribute containing RRSA related data for your cluster.
	RrsaMetadata         ManagedKubernetesRrsaMetadataPtrInput
	SecurityGroupId      pulumi.StringPtrInput
	ServiceAccountIssuer pulumi.StringPtrInput
	ServiceCidr          pulumi.StringPtrInput
	// The ID of APIServer load balancer.
	SlbId pulumi.StringPtrInput
	// The public ip of load balancer.
	SlbInternet        pulumi.StringPtrInput
	SlbInternetEnabled pulumi.BoolPtrInput
	// The ID of private load balancer where the current cluster master node is located.
	SlbIntranet pulumi.StringPtrInput
	Tags        pulumi.MapInput
	Timezone    pulumi.StringPtrInput
	UserCa      pulumi.StringPtrInput
	Version     pulumi.StringPtrInput
	// The ID of VPC where the current cluster is located.
	VpcId pulumi.StringPtrInput
	// The RamRole Name attached to worker node.
	WorkerRamRoleName pulumi.StringPtrInput
	WorkerVswitchIds  pulumi.StringArrayInput
}

func (ManagedKubernetesState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedKubernetesState)(nil)).Elem()
}

type managedKubernetesArgs struct {
	Addons       []ManagedKubernetesAddon `pulumi:"addons"`
	ApiAudiences []string                 `pulumi:"apiAudiences"`
	// The base64 encoded client certificate data required to communicate with your cluster. Add this to the client-certificate-data section of the kubeconfig file for your cluster.
	ClientCert *string `pulumi:"clientCert"`
	// The base64 encoded client key data required to communicate with your cluster. Add this to the client-key-data section of the kubeconfig file for your cluster.
	ClientKey     *string `pulumi:"clientKey"`
	ClusterCaCert *string `pulumi:"clusterCaCert"`
	// cluster local domain
	ClusterDomain             *string                         `pulumi:"clusterDomain"`
	ClusterSpec               *string                         `pulumi:"clusterSpec"`
	ControlPlaneLogComponents []string                        `pulumi:"controlPlaneLogComponents"`
	ControlPlaneLogProject    *string                         `pulumi:"controlPlaneLogProject"`
	ControlPlaneLogTtl        *string                         `pulumi:"controlPlaneLogTtl"`
	CustomSan                 *string                         `pulumi:"customSan"`
	DeleteOptions             []ManagedKubernetesDeleteOption `pulumi:"deleteOptions"`
	DeletionProtection        *bool                           `pulumi:"deletionProtection"`
	EnableRrsa                *bool                           `pulumi:"enableRrsa"`
	// disk encryption key, only in ack-pro
	EncryptionProviderKey     *string                             `pulumi:"encryptionProviderKey"`
	IsEnterpriseSecurityGroup *bool                               `pulumi:"isEnterpriseSecurityGroup"`
	LoadBalancerSpec          *string                             `pulumi:"loadBalancerSpec"`
	MaintenanceWindow         *ManagedKubernetesMaintenanceWindow `pulumi:"maintenanceWindow"`
	// Node name.
	Name                 *string                `pulumi:"name"`
	NamePrefix           *string                `pulumi:"namePrefix"`
	NewNatGateway        *bool                  `pulumi:"newNatGateway"`
	NodeCidrMask         *int                   `pulumi:"nodeCidrMask"`
	PodCidr              *string                `pulumi:"podCidr"`
	PodVswitchIds        []string               `pulumi:"podVswitchIds"`
	ProxyMode            *string                `pulumi:"proxyMode"`
	ResourceGroupId      *string                `pulumi:"resourceGroupId"`
	RetainResources      []string               `pulumi:"retainResources"`
	SecurityGroupId      *string                `pulumi:"securityGroupId"`
	ServiceAccountIssuer *string                `pulumi:"serviceAccountIssuer"`
	ServiceCidr          *string                `pulumi:"serviceCidr"`
	SlbInternetEnabled   *bool                  `pulumi:"slbInternetEnabled"`
	Tags                 map[string]interface{} `pulumi:"tags"`
	Timezone             *string                `pulumi:"timezone"`
	UserCa               *string                `pulumi:"userCa"`
	Version              *string                `pulumi:"version"`
	WorkerVswitchIds     []string               `pulumi:"workerVswitchIds"`
}

// The set of arguments for constructing a ManagedKubernetes resource.
type ManagedKubernetesArgs struct {
	Addons       ManagedKubernetesAddonArrayInput
	ApiAudiences pulumi.StringArrayInput
	// The base64 encoded client certificate data required to communicate with your cluster. Add this to the client-certificate-data section of the kubeconfig file for your cluster.
	ClientCert pulumi.StringPtrInput
	// The base64 encoded client key data required to communicate with your cluster. Add this to the client-key-data section of the kubeconfig file for your cluster.
	ClientKey     pulumi.StringPtrInput
	ClusterCaCert pulumi.StringPtrInput
	// cluster local domain
	ClusterDomain             pulumi.StringPtrInput
	ClusterSpec               pulumi.StringPtrInput
	ControlPlaneLogComponents pulumi.StringArrayInput
	ControlPlaneLogProject    pulumi.StringPtrInput
	ControlPlaneLogTtl        pulumi.StringPtrInput
	CustomSan                 pulumi.StringPtrInput
	DeleteOptions             ManagedKubernetesDeleteOptionArrayInput
	DeletionProtection        pulumi.BoolPtrInput
	EnableRrsa                pulumi.BoolPtrInput
	// disk encryption key, only in ack-pro
	EncryptionProviderKey     pulumi.StringPtrInput
	IsEnterpriseSecurityGroup pulumi.BoolPtrInput
	LoadBalancerSpec          pulumi.StringPtrInput
	MaintenanceWindow         ManagedKubernetesMaintenanceWindowPtrInput
	// Node name.
	Name                 pulumi.StringPtrInput
	NamePrefix           pulumi.StringPtrInput
	NewNatGateway        pulumi.BoolPtrInput
	NodeCidrMask         pulumi.IntPtrInput
	PodCidr              pulumi.StringPtrInput
	PodVswitchIds        pulumi.StringArrayInput
	ProxyMode            pulumi.StringPtrInput
	ResourceGroupId      pulumi.StringPtrInput
	RetainResources      pulumi.StringArrayInput
	SecurityGroupId      pulumi.StringPtrInput
	ServiceAccountIssuer pulumi.StringPtrInput
	ServiceCidr          pulumi.StringPtrInput
	SlbInternetEnabled   pulumi.BoolPtrInput
	Tags                 pulumi.MapInput
	Timezone             pulumi.StringPtrInput
	UserCa               pulumi.StringPtrInput
	Version              pulumi.StringPtrInput
	WorkerVswitchIds     pulumi.StringArrayInput
}

func (ManagedKubernetesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedKubernetesArgs)(nil)).Elem()
}

type ManagedKubernetesInput interface {
	pulumi.Input

	ToManagedKubernetesOutput() ManagedKubernetesOutput
	ToManagedKubernetesOutputWithContext(ctx context.Context) ManagedKubernetesOutput
}

func (*ManagedKubernetes) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedKubernetes)(nil)).Elem()
}

func (i *ManagedKubernetes) ToManagedKubernetesOutput() ManagedKubernetesOutput {
	return i.ToManagedKubernetesOutputWithContext(context.Background())
}

func (i *ManagedKubernetes) ToManagedKubernetesOutputWithContext(ctx context.Context) ManagedKubernetesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedKubernetesOutput)
}

// ManagedKubernetesArrayInput is an input type that accepts ManagedKubernetesArray and ManagedKubernetesArrayOutput values.
// You can construct a concrete instance of `ManagedKubernetesArrayInput` via:
//
//	ManagedKubernetesArray{ ManagedKubernetesArgs{...} }
type ManagedKubernetesArrayInput interface {
	pulumi.Input

	ToManagedKubernetesArrayOutput() ManagedKubernetesArrayOutput
	ToManagedKubernetesArrayOutputWithContext(context.Context) ManagedKubernetesArrayOutput
}

type ManagedKubernetesArray []ManagedKubernetesInput

func (ManagedKubernetesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedKubernetes)(nil)).Elem()
}

func (i ManagedKubernetesArray) ToManagedKubernetesArrayOutput() ManagedKubernetesArrayOutput {
	return i.ToManagedKubernetesArrayOutputWithContext(context.Background())
}

func (i ManagedKubernetesArray) ToManagedKubernetesArrayOutputWithContext(ctx context.Context) ManagedKubernetesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedKubernetesArrayOutput)
}

// ManagedKubernetesMapInput is an input type that accepts ManagedKubernetesMap and ManagedKubernetesMapOutput values.
// You can construct a concrete instance of `ManagedKubernetesMapInput` via:
//
//	ManagedKubernetesMap{ "key": ManagedKubernetesArgs{...} }
type ManagedKubernetesMapInput interface {
	pulumi.Input

	ToManagedKubernetesMapOutput() ManagedKubernetesMapOutput
	ToManagedKubernetesMapOutputWithContext(context.Context) ManagedKubernetesMapOutput
}

type ManagedKubernetesMap map[string]ManagedKubernetesInput

func (ManagedKubernetesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedKubernetes)(nil)).Elem()
}

func (i ManagedKubernetesMap) ToManagedKubernetesMapOutput() ManagedKubernetesMapOutput {
	return i.ToManagedKubernetesMapOutputWithContext(context.Background())
}

func (i ManagedKubernetesMap) ToManagedKubernetesMapOutputWithContext(ctx context.Context) ManagedKubernetesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedKubernetesMapOutput)
}

type ManagedKubernetesOutput struct{ *pulumi.OutputState }

func (ManagedKubernetesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedKubernetes)(nil)).Elem()
}

func (o ManagedKubernetesOutput) ToManagedKubernetesOutput() ManagedKubernetesOutput {
	return o
}

func (o ManagedKubernetesOutput) ToManagedKubernetesOutputWithContext(ctx context.Context) ManagedKubernetesOutput {
	return o
}

func (o ManagedKubernetesOutput) Addons() ManagedKubernetesAddonArrayOutput {
	return o.ApplyT(func(v *ManagedKubernetes) ManagedKubernetesAddonArrayOutput { return v.Addons }).(ManagedKubernetesAddonArrayOutput)
}

func (o ManagedKubernetesOutput) ApiAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringArrayOutput { return v.ApiAudiences }).(pulumi.StringArrayOutput)
}

// (Available in 1.105.0+) Nested attribute containing certificate authority data for your cluster.
func (o ManagedKubernetesOutput) CertificateAuthority() ManagedKubernetesCertificateAuthorityOutput {
	return o.ApplyT(func(v *ManagedKubernetes) ManagedKubernetesCertificateAuthorityOutput { return v.CertificateAuthority }).(ManagedKubernetesCertificateAuthorityOutput)
}

// The base64 encoded client certificate data required to communicate with your cluster. Add this to the client-certificate-data section of the kubeconfig file for your cluster.
func (o ManagedKubernetesOutput) ClientCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringPtrOutput { return v.ClientCert }).(pulumi.StringPtrOutput)
}

// The base64 encoded client key data required to communicate with your cluster. Add this to the client-key-data section of the kubeconfig file for your cluster.
func (o ManagedKubernetesOutput) ClientKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringPtrOutput { return v.ClientKey }).(pulumi.StringPtrOutput)
}

func (o ManagedKubernetesOutput) ClusterCaCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringPtrOutput { return v.ClusterCaCert }).(pulumi.StringPtrOutput)
}

// cluster local domain
func (o ManagedKubernetesOutput) ClusterDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringPtrOutput { return v.ClusterDomain }).(pulumi.StringPtrOutput)
}

func (o ManagedKubernetesOutput) ClusterSpec() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringOutput { return v.ClusterSpec }).(pulumi.StringOutput)
}

// Map of kubernetes cluster connection information.
func (o ManagedKubernetesOutput) Connections() ManagedKubernetesConnectionsOutput {
	return o.ApplyT(func(v *ManagedKubernetes) ManagedKubernetesConnectionsOutput { return v.Connections }).(ManagedKubernetesConnectionsOutput)
}

func (o ManagedKubernetesOutput) ControlPlaneLogComponents() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringArrayOutput { return v.ControlPlaneLogComponents }).(pulumi.StringArrayOutput)
}

func (o ManagedKubernetesOutput) ControlPlaneLogProject() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringOutput { return v.ControlPlaneLogProject }).(pulumi.StringOutput)
}

func (o ManagedKubernetesOutput) ControlPlaneLogTtl() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringOutput { return v.ControlPlaneLogTtl }).(pulumi.StringOutput)
}

func (o ManagedKubernetesOutput) CustomSan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringPtrOutput { return v.CustomSan }).(pulumi.StringPtrOutput)
}

func (o ManagedKubernetesOutput) DeleteOptions() ManagedKubernetesDeleteOptionArrayOutput {
	return o.ApplyT(func(v *ManagedKubernetes) ManagedKubernetesDeleteOptionArrayOutput { return v.DeleteOptions }).(ManagedKubernetesDeleteOptionArrayOutput)
}

func (o ManagedKubernetesOutput) DeletionProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.BoolPtrOutput { return v.DeletionProtection }).(pulumi.BoolPtrOutput)
}

func (o ManagedKubernetesOutput) EnableRrsa() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.BoolPtrOutput { return v.EnableRrsa }).(pulumi.BoolPtrOutput)
}

// disk encryption key, only in ack-pro
func (o ManagedKubernetesOutput) EncryptionProviderKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringPtrOutput { return v.EncryptionProviderKey }).(pulumi.StringPtrOutput)
}

func (o ManagedKubernetesOutput) IsEnterpriseSecurityGroup() pulumi.BoolOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.BoolOutput { return v.IsEnterpriseSecurityGroup }).(pulumi.BoolOutput)
}

func (o ManagedKubernetesOutput) LoadBalancerSpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringPtrOutput { return v.LoadBalancerSpec }).(pulumi.StringPtrOutput)
}

func (o ManagedKubernetesOutput) MaintenanceWindow() ManagedKubernetesMaintenanceWindowOutput {
	return o.ApplyT(func(v *ManagedKubernetes) ManagedKubernetesMaintenanceWindowOutput { return v.MaintenanceWindow }).(ManagedKubernetesMaintenanceWindowOutput)
}

// Node name.
func (o ManagedKubernetesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedKubernetesOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringPtrOutput { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

// The ID of nat gateway used to launch kubernetes cluster.
func (o ManagedKubernetesOutput) NatGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringOutput { return v.NatGatewayId }).(pulumi.StringOutput)
}

func (o ManagedKubernetesOutput) NewNatGateway() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.BoolPtrOutput { return v.NewNatGateway }).(pulumi.BoolPtrOutput)
}

func (o ManagedKubernetesOutput) NodeCidrMask() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.IntPtrOutput { return v.NodeCidrMask }).(pulumi.IntPtrOutput)
}

func (o ManagedKubernetesOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringPtrOutput { return v.PodCidr }).(pulumi.StringPtrOutput)
}

func (o ManagedKubernetesOutput) PodVswitchIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringArrayOutput { return v.PodVswitchIds }).(pulumi.StringArrayOutput)
}

func (o ManagedKubernetesOutput) ProxyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringPtrOutput { return v.ProxyMode }).(pulumi.StringPtrOutput)
}

func (o ManagedKubernetesOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

func (o ManagedKubernetesOutput) RetainResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringArrayOutput { return v.RetainResources }).(pulumi.StringArrayOutput)
}

// (Optional, Available in v1.185.0+) Nested attribute containing RRSA related data for your cluster.
func (o ManagedKubernetesOutput) RrsaMetadata() ManagedKubernetesRrsaMetadataOutput {
	return o.ApplyT(func(v *ManagedKubernetes) ManagedKubernetesRrsaMetadataOutput { return v.RrsaMetadata }).(ManagedKubernetesRrsaMetadataOutput)
}

func (o ManagedKubernetesOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

func (o ManagedKubernetesOutput) ServiceAccountIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringPtrOutput { return v.ServiceAccountIssuer }).(pulumi.StringPtrOutput)
}

func (o ManagedKubernetesOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringPtrOutput { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

// The ID of APIServer load balancer.
func (o ManagedKubernetesOutput) SlbId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringOutput { return v.SlbId }).(pulumi.StringOutput)
}

// The public ip of load balancer.
func (o ManagedKubernetesOutput) SlbInternet() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringOutput { return v.SlbInternet }).(pulumi.StringOutput)
}

func (o ManagedKubernetesOutput) SlbInternetEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.BoolPtrOutput { return v.SlbInternetEnabled }).(pulumi.BoolPtrOutput)
}

// The ID of private load balancer where the current cluster master node is located.
func (o ManagedKubernetesOutput) SlbIntranet() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringOutput { return v.SlbIntranet }).(pulumi.StringOutput)
}

func (o ManagedKubernetesOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

func (o ManagedKubernetesOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

func (o ManagedKubernetesOutput) UserCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringPtrOutput { return v.UserCa }).(pulumi.StringPtrOutput)
}

func (o ManagedKubernetesOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// The ID of VPC where the current cluster is located.
func (o ManagedKubernetesOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The RamRole Name attached to worker node.
func (o ManagedKubernetesOutput) WorkerRamRoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringOutput { return v.WorkerRamRoleName }).(pulumi.StringOutput)
}

func (o ManagedKubernetesOutput) WorkerVswitchIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedKubernetes) pulumi.StringArrayOutput { return v.WorkerVswitchIds }).(pulumi.StringArrayOutput)
}

type ManagedKubernetesArrayOutput struct{ *pulumi.OutputState }

func (ManagedKubernetesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedKubernetes)(nil)).Elem()
}

func (o ManagedKubernetesArrayOutput) ToManagedKubernetesArrayOutput() ManagedKubernetesArrayOutput {
	return o
}

func (o ManagedKubernetesArrayOutput) ToManagedKubernetesArrayOutputWithContext(ctx context.Context) ManagedKubernetesArrayOutput {
	return o
}

func (o ManagedKubernetesArrayOutput) Index(i pulumi.IntInput) ManagedKubernetesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ManagedKubernetes {
		return vs[0].([]*ManagedKubernetes)[vs[1].(int)]
	}).(ManagedKubernetesOutput)
}

type ManagedKubernetesMapOutput struct{ *pulumi.OutputState }

func (ManagedKubernetesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedKubernetes)(nil)).Elem()
}

func (o ManagedKubernetesMapOutput) ToManagedKubernetesMapOutput() ManagedKubernetesMapOutput {
	return o
}

func (o ManagedKubernetesMapOutput) ToManagedKubernetesMapOutputWithContext(ctx context.Context) ManagedKubernetesMapOutput {
	return o
}

func (o ManagedKubernetesMapOutput) MapIndex(k pulumi.StringInput) ManagedKubernetesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ManagedKubernetes {
		return vs[0].(map[string]*ManagedKubernetes)[vs[1].(string)]
	}).(ManagedKubernetesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedKubernetesInput)(nil)).Elem(), &ManagedKubernetes{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedKubernetesArrayInput)(nil)).Elem(), ManagedKubernetesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedKubernetesMapInput)(nil)).Elem(), ManagedKubernetesMap{})
	pulumi.RegisterOutputType(ManagedKubernetesOutput{})
	pulumi.RegisterOutputType(ManagedKubernetesArrayOutput{})
	pulumi.RegisterOutputType(ManagedKubernetesMapOutput{})
}
