// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This resource will help you to manager a Serverless Kubernetes Cluster, see [What is serverless kubernetes](https://www.alibabacloud.com/help/en/ack/ack-managed-and-ack-dedicated/developer-reference/create-a-dedicated-kubernetes-cluster-that-supports-sandboxed-containers). The cluster is same as container service created by web console.
 *
 * > **NOTE:** Available since v1.58.0.
 *
 * > **NOTE:** Serverless Kubernetes cluster only supports VPC network and it can access internet while creating kubernetes cluster.
 * A Nat Gateway and configuring a SNAT for it can ensure one VPC network access internet. If there is no nat gateway in the
 * VPC, you can set `newNatGateway` to "true" to create one automatically.
 *
 * > **NOTE:** Creating serverless kubernetes cluster need to install several packages and it will cost about 5 minutes. Please be patient.
 *
 * > **NOTE:** The provider supports to download kube config, client certificate, client key and cluster ca certificate
 * after creating cluster successfully, and you can put them into the specified location, like '~/.kube/config'.
 *
 * > **NOTE:** If you want to manage serverless Kubernetes, you can use Kubernetes Provider.
 *
 * > **NOTE:** You need to activate several other products and confirm Authorization Policy used by Container Service before using this resource.
 * Please refer to the `Authorization management` and `Cluster management` sections in the [Document Center](https://www.alibabacloud.com/help/doc-detail/86488.htm).
 *
 * > **NOTE:** From version 1.162.0, support for creating professional serverless cluster.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "ask-example";
 * const default = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     vpcName: name,
 *     cidrBlock: "10.1.0.0/21",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     vswitchName: name,
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "10.1.1.0/24",
 *     zoneId: _default.then(_default => _default.zones?.[0]?.id),
 * });
 * const serverless = new alicloud.cs.ServerlessKubernetes("serverless", {
 *     namePrefix: name,
 *     clusterSpec: "ack.pro.small",
 *     vpcId: defaultNetwork.id,
 *     vswitchIds: [defaultSwitch.id],
 *     newNatGateway: true,
 *     endpointPublicAccessEnabled: true,
 *     deletionProtection: false,
 *     loadBalancerSpec: "slb.s2.small",
 *     timeZone: "Asia/Shanghai",
 *     serviceCidr: "172.21.0.0/20",
 *     serviceDiscoveryTypes: ["PrivateZone"],
 *     loggingType: "SLS",
 *     tags: {
 *         "k-aa": "v-aa",
 *         "k-bb": "v-aa",
 *     },
 *     addons: [
 *         {
 *             name: "alb-ingress-controller",
 *         },
 *         {
 *             name: "metrics-server",
 *         },
 *         {
 *             name: "knative",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Serverless Kubernetes cluster can be imported using the id, e.g. Then complete the main.tf accords to the result of `pulumi preview`.
 *
 * ```sh
 * $ pulumi import alicloud:cs/serverlessKubernetes:ServerlessKubernetes main ce4273f9156874b46bb
 * ```
 */
export class ServerlessKubernetes extends pulumi.CustomResource {
    /**
     * Get an existing ServerlessKubernetes resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerlessKubernetesState, opts?: pulumi.CustomResourceOptions): ServerlessKubernetes {
        return new ServerlessKubernetes(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cs/serverlessKubernetes:ServerlessKubernetes';

    /**
     * Returns true if the given object is an instance of ServerlessKubernetes.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerlessKubernetes {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerlessKubernetes.__pulumiType;
    }

    /**
     * You can specific network plugin,log component,ingress component and so on. See `addons` below.
     */
    public readonly addons!: pulumi.Output<outputs.cs.ServerlessKubernetesAddon[]>;
    /**
     * The path of client certificate, like `~/.kube/client-cert.pem`.
     */
    public readonly clientCert!: pulumi.Output<string | undefined>;
    /**
     * The path of client key, like `~/.kube/client-key.pem`.
     */
    public readonly clientKey!: pulumi.Output<string | undefined>;
    /**
     * The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
     */
    public readonly clusterCaCert!: pulumi.Output<string | undefined>;
    /**
     * The cluster specifications of serverless kubernetes cluster, which can be empty. Valid values:
     * - ack.standard: Standard serverless clusters.
     * - ack.pro.small: Professional serverless clusters.
     */
    public readonly clusterSpec!: pulumi.Output<string>;
    /**
     * whether to create a v2 version cluster.
     *
     * *Removed params*
     */
    public readonly createV2Cluster!: pulumi.Output<boolean>;
    /**
     * Whether enable the deletion protection or not.
     * - true: Enable deletion protection.
     * - false: Disable deletion protection.
     */
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to enable cluster to support RRSA for version 1.22.3+. Default to `false`. Once the RRSA function is turned on, it is not allowed to turn off. If your cluster has enabled this function, please manually modify your tf file and add the rrsa configuration to the file, learn more [RAM Roles for Service Accounts](https://www.alibabacloud.com/help/zh/container-service-for-kubernetes/latest/use-rrsa-to-enforce-access-control).
     */
    public readonly enableRrsa!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to create internet  eip for API Server. Default to false.
     */
    public readonly endpointPublicAccessEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Default false, when you want to change `vpcId` and `vswitchId`, you have to set this field to true, then the cluster will be recreated.
     */
    public readonly forceUpdate!: pulumi.Output<boolean | undefined>;
    /**
     * The path of kube config, like `~/.kube/config`.
     *
     * @deprecated Field 'kube_config' has been deprecated from provider version 1.187.0. New DataSource 'alicloud_cs_cluster_credential' manage your cluster's kube config.
     */
    public readonly kubeConfig!: pulumi.Output<string | undefined>;
    /**
     * The cluster api server load balance instance specification, default `slb.s2.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
     */
    public readonly loadBalancerSpec!: pulumi.Output<string>;
    /**
     * Enable log service, Valid value `SLS`.
     */
    public readonly loggingType!: pulumi.Output<string | undefined>;
    /**
     * The kubernetes cluster's name. It is the only in one Alicloud account.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * Whether to create a new nat gateway while creating kubernetes cluster. SNAT must be configured when a new VPC is automatically created. Default is `true`.
     */
    public readonly newNatGateway!: pulumi.Output<boolean | undefined>;
    /**
     * Has been deprecated from provider version 1.123.1. `PrivateZone` is used as the enumeration value of `serviceDiscoveryTypes`.
     *
     * @deprecated Field 'private_zone' has been deprecated from provider version 1.123.1. New field 'service_discovery_types' replace it.
     */
    public readonly privateZone!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    public readonly retainResources!: pulumi.Output<string[] | undefined>;
    /**
     * Nested attribute containing RRSA related data for your cluster. See `rrsaMetadata` below.
     */
    public readonly rrsaMetadata!: pulumi.Output<outputs.cs.ServerlessKubernetesRrsaMetadata>;
    /**
     * The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * CIDR block of the service network. The specified CIDR block cannot overlap with that of the VPC or those of the ACK clusters that are deployed in the VPC. The CIDR block cannot be modified after the cluster is created.
     */
    public readonly serviceCidr!: pulumi.Output<string | undefined>;
    /**
     * Service discovery type. If the value is empty, it means that service discovery is not enabled. Valid values are `CoreDNS` and `PrivateZone`.
     */
    public readonly serviceDiscoveryTypes!: pulumi.Output<string[] | undefined>;
    /**
     * If you use an existing SLS project, you must specify `slsProjectName`.
     */
    public readonly slsProjectName!: pulumi.Output<string>;
    /**
     * Default nil, A map of tags assigned to the kubernetes cluster and work nodes.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The time zone of the cluster.
     */
    public readonly timeZone!: pulumi.Output<string>;
    /**
     * Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used.
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * The vpc where new kubernetes cluster will be located. Specify one vpc's id, if it is not specified, a new VPC  will be built.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availabilityZone` specified.
     *
     * @deprecated Field 'vswitch_id' has been deprecated from provider version 1.91.0. New field 'vswitch_ids' replace it.
     */
    public readonly vswitchId!: pulumi.Output<string>;
    /**
     * The vswitches where new kubernetes cluster will be located.
     */
    public readonly vswitchIds!: pulumi.Output<string[]>;
    /**
     * When creating a cluster using automatic VPC creation, you need to specify the zone where the VPC is located.
     */
    public readonly zoneId!: pulumi.Output<string | undefined>;

    /**
     * Create a ServerlessKubernetes resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerlessKubernetesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerlessKubernetesArgs | ServerlessKubernetesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerlessKubernetesState | undefined;
            resourceInputs["addons"] = state ? state.addons : undefined;
            resourceInputs["clientCert"] = state ? state.clientCert : undefined;
            resourceInputs["clientKey"] = state ? state.clientKey : undefined;
            resourceInputs["clusterCaCert"] = state ? state.clusterCaCert : undefined;
            resourceInputs["clusterSpec"] = state ? state.clusterSpec : undefined;
            resourceInputs["createV2Cluster"] = state ? state.createV2Cluster : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["enableRrsa"] = state ? state.enableRrsa : undefined;
            resourceInputs["endpointPublicAccessEnabled"] = state ? state.endpointPublicAccessEnabled : undefined;
            resourceInputs["forceUpdate"] = state ? state.forceUpdate : undefined;
            resourceInputs["kubeConfig"] = state ? state.kubeConfig : undefined;
            resourceInputs["loadBalancerSpec"] = state ? state.loadBalancerSpec : undefined;
            resourceInputs["loggingType"] = state ? state.loggingType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["newNatGateway"] = state ? state.newNatGateway : undefined;
            resourceInputs["privateZone"] = state ? state.privateZone : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["retainResources"] = state ? state.retainResources : undefined;
            resourceInputs["rrsaMetadata"] = state ? state.rrsaMetadata : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["serviceCidr"] = state ? state.serviceCidr : undefined;
            resourceInputs["serviceDiscoveryTypes"] = state ? state.serviceDiscoveryTypes : undefined;
            resourceInputs["slsProjectName"] = state ? state.slsProjectName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["timeZone"] = state ? state.timeZone : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
            resourceInputs["vswitchIds"] = state ? state.vswitchIds : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ServerlessKubernetesArgs | undefined;
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["addons"] = args ? args.addons : undefined;
            resourceInputs["clientCert"] = args ? args.clientCert : undefined;
            resourceInputs["clientKey"] = args ? args.clientKey : undefined;
            resourceInputs["clusterCaCert"] = args ? args.clusterCaCert : undefined;
            resourceInputs["clusterSpec"] = args ? args.clusterSpec : undefined;
            resourceInputs["createV2Cluster"] = args ? args.createV2Cluster : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["enableRrsa"] = args ? args.enableRrsa : undefined;
            resourceInputs["endpointPublicAccessEnabled"] = args ? args.endpointPublicAccessEnabled : undefined;
            resourceInputs["forceUpdate"] = args ? args.forceUpdate : undefined;
            resourceInputs["kubeConfig"] = args ? args.kubeConfig : undefined;
            resourceInputs["loadBalancerSpec"] = args ? args.loadBalancerSpec : undefined;
            resourceInputs["loggingType"] = args ? args.loggingType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["newNatGateway"] = args ? args.newNatGateway : undefined;
            resourceInputs["privateZone"] = args ? args.privateZone : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["retainResources"] = args ? args.retainResources : undefined;
            resourceInputs["rrsaMetadata"] = args ? args.rrsaMetadata : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["serviceCidr"] = args ? args.serviceCidr : undefined;
            resourceInputs["serviceDiscoveryTypes"] = args ? args.serviceDiscoveryTypes : undefined;
            resourceInputs["slsProjectName"] = args ? args.slsProjectName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["timeZone"] = args ? args.timeZone : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["vswitchIds"] = args ? args.vswitchIds : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerlessKubernetes.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerlessKubernetes resources.
 */
export interface ServerlessKubernetesState {
    /**
     * You can specific network plugin,log component,ingress component and so on. See `addons` below.
     */
    addons?: pulumi.Input<pulumi.Input<inputs.cs.ServerlessKubernetesAddon>[]>;
    /**
     * The path of client certificate, like `~/.kube/client-cert.pem`.
     */
    clientCert?: pulumi.Input<string>;
    /**
     * The path of client key, like `~/.kube/client-key.pem`.
     */
    clientKey?: pulumi.Input<string>;
    /**
     * The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
     */
    clusterCaCert?: pulumi.Input<string>;
    /**
     * The cluster specifications of serverless kubernetes cluster, which can be empty. Valid values:
     * - ack.standard: Standard serverless clusters.
     * - ack.pro.small: Professional serverless clusters.
     */
    clusterSpec?: pulumi.Input<string>;
    /**
     * whether to create a v2 version cluster.
     *
     * *Removed params*
     */
    createV2Cluster?: pulumi.Input<boolean>;
    /**
     * Whether enable the deletion protection or not.
     * - true: Enable deletion protection.
     * - false: Disable deletion protection.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * Whether to enable cluster to support RRSA for version 1.22.3+. Default to `false`. Once the RRSA function is turned on, it is not allowed to turn off. If your cluster has enabled this function, please manually modify your tf file and add the rrsa configuration to the file, learn more [RAM Roles for Service Accounts](https://www.alibabacloud.com/help/zh/container-service-for-kubernetes/latest/use-rrsa-to-enforce-access-control).
     */
    enableRrsa?: pulumi.Input<boolean>;
    /**
     * Whether to create internet  eip for API Server. Default to false.
     */
    endpointPublicAccessEnabled?: pulumi.Input<boolean>;
    /**
     * Default false, when you want to change `vpcId` and `vswitchId`, you have to set this field to true, then the cluster will be recreated.
     */
    forceUpdate?: pulumi.Input<boolean>;
    /**
     * The path of kube config, like `~/.kube/config`.
     *
     * @deprecated Field 'kube_config' has been deprecated from provider version 1.187.0. New DataSource 'alicloud_cs_cluster_credential' manage your cluster's kube config.
     */
    kubeConfig?: pulumi.Input<string>;
    /**
     * The cluster api server load balance instance specification, default `slb.s2.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
     */
    loadBalancerSpec?: pulumi.Input<string>;
    /**
     * Enable log service, Valid value `SLS`.
     */
    loggingType?: pulumi.Input<string>;
    /**
     * The kubernetes cluster's name. It is the only in one Alicloud account.
     */
    name?: pulumi.Input<string>;
    namePrefix?: pulumi.Input<string>;
    /**
     * Whether to create a new nat gateway while creating kubernetes cluster. SNAT must be configured when a new VPC is automatically created. Default is `true`.
     */
    newNatGateway?: pulumi.Input<boolean>;
    /**
     * Has been deprecated from provider version 1.123.1. `PrivateZone` is used as the enumeration value of `serviceDiscoveryTypes`.
     *
     * @deprecated Field 'private_zone' has been deprecated from provider version 1.123.1. New field 'service_discovery_types' replace it.
     */
    privateZone?: pulumi.Input<boolean>;
    /**
     * The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    retainResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Nested attribute containing RRSA related data for your cluster. See `rrsaMetadata` below.
     */
    rrsaMetadata?: pulumi.Input<inputs.cs.ServerlessKubernetesRrsaMetadata>;
    /**
     * The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * CIDR block of the service network. The specified CIDR block cannot overlap with that of the VPC or those of the ACK clusters that are deployed in the VPC. The CIDR block cannot be modified after the cluster is created.
     */
    serviceCidr?: pulumi.Input<string>;
    /**
     * Service discovery type. If the value is empty, it means that service discovery is not enabled. Valid values are `CoreDNS` and `PrivateZone`.
     */
    serviceDiscoveryTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If you use an existing SLS project, you must specify `slsProjectName`.
     */
    slsProjectName?: pulumi.Input<string>;
    /**
     * Default nil, A map of tags assigned to the kubernetes cluster and work nodes.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The time zone of the cluster.
     */
    timeZone?: pulumi.Input<string>;
    /**
     * Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used.
     */
    version?: pulumi.Input<string>;
    /**
     * The vpc where new kubernetes cluster will be located. Specify one vpc's id, if it is not specified, a new VPC  will be built.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availabilityZone` specified.
     *
     * @deprecated Field 'vswitch_id' has been deprecated from provider version 1.91.0. New field 'vswitch_ids' replace it.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The vswitches where new kubernetes cluster will be located.
     */
    vswitchIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When creating a cluster using automatic VPC creation, you need to specify the zone where the VPC is located.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServerlessKubernetes resource.
 */
export interface ServerlessKubernetesArgs {
    /**
     * You can specific network plugin,log component,ingress component and so on. See `addons` below.
     */
    addons?: pulumi.Input<pulumi.Input<inputs.cs.ServerlessKubernetesAddon>[]>;
    /**
     * The path of client certificate, like `~/.kube/client-cert.pem`.
     */
    clientCert?: pulumi.Input<string>;
    /**
     * The path of client key, like `~/.kube/client-key.pem`.
     */
    clientKey?: pulumi.Input<string>;
    /**
     * The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
     */
    clusterCaCert?: pulumi.Input<string>;
    /**
     * The cluster specifications of serverless kubernetes cluster, which can be empty. Valid values:
     * - ack.standard: Standard serverless clusters.
     * - ack.pro.small: Professional serverless clusters.
     */
    clusterSpec?: pulumi.Input<string>;
    /**
     * whether to create a v2 version cluster.
     *
     * *Removed params*
     */
    createV2Cluster?: pulumi.Input<boolean>;
    /**
     * Whether enable the deletion protection or not.
     * - true: Enable deletion protection.
     * - false: Disable deletion protection.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * Whether to enable cluster to support RRSA for version 1.22.3+. Default to `false`. Once the RRSA function is turned on, it is not allowed to turn off. If your cluster has enabled this function, please manually modify your tf file and add the rrsa configuration to the file, learn more [RAM Roles for Service Accounts](https://www.alibabacloud.com/help/zh/container-service-for-kubernetes/latest/use-rrsa-to-enforce-access-control).
     */
    enableRrsa?: pulumi.Input<boolean>;
    /**
     * Whether to create internet  eip for API Server. Default to false.
     */
    endpointPublicAccessEnabled?: pulumi.Input<boolean>;
    /**
     * Default false, when you want to change `vpcId` and `vswitchId`, you have to set this field to true, then the cluster will be recreated.
     */
    forceUpdate?: pulumi.Input<boolean>;
    /**
     * The path of kube config, like `~/.kube/config`.
     *
     * @deprecated Field 'kube_config' has been deprecated from provider version 1.187.0. New DataSource 'alicloud_cs_cluster_credential' manage your cluster's kube config.
     */
    kubeConfig?: pulumi.Input<string>;
    /**
     * The cluster api server load balance instance specification, default `slb.s2.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
     */
    loadBalancerSpec?: pulumi.Input<string>;
    /**
     * Enable log service, Valid value `SLS`.
     */
    loggingType?: pulumi.Input<string>;
    /**
     * The kubernetes cluster's name. It is the only in one Alicloud account.
     */
    name?: pulumi.Input<string>;
    namePrefix?: pulumi.Input<string>;
    /**
     * Whether to create a new nat gateway while creating kubernetes cluster. SNAT must be configured when a new VPC is automatically created. Default is `true`.
     */
    newNatGateway?: pulumi.Input<boolean>;
    /**
     * Has been deprecated from provider version 1.123.1. `PrivateZone` is used as the enumeration value of `serviceDiscoveryTypes`.
     *
     * @deprecated Field 'private_zone' has been deprecated from provider version 1.123.1. New field 'service_discovery_types' replace it.
     */
    privateZone?: pulumi.Input<boolean>;
    /**
     * The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    retainResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Nested attribute containing RRSA related data for your cluster. See `rrsaMetadata` below.
     */
    rrsaMetadata?: pulumi.Input<inputs.cs.ServerlessKubernetesRrsaMetadata>;
    /**
     * The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * CIDR block of the service network. The specified CIDR block cannot overlap with that of the VPC or those of the ACK clusters that are deployed in the VPC. The CIDR block cannot be modified after the cluster is created.
     */
    serviceCidr?: pulumi.Input<string>;
    /**
     * Service discovery type. If the value is empty, it means that service discovery is not enabled. Valid values are `CoreDNS` and `PrivateZone`.
     */
    serviceDiscoveryTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If you use an existing SLS project, you must specify `slsProjectName`.
     */
    slsProjectName?: pulumi.Input<string>;
    /**
     * Default nil, A map of tags assigned to the kubernetes cluster and work nodes.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The time zone of the cluster.
     */
    timeZone?: pulumi.Input<string>;
    /**
     * Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used.
     */
    version?: pulumi.Input<string>;
    /**
     * The vpc where new kubernetes cluster will be located. Specify one vpc's id, if it is not specified, a new VPC  will be built.
     */
    vpcId: pulumi.Input<string>;
    /**
     * The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availabilityZone` specified.
     *
     * @deprecated Field 'vswitch_id' has been deprecated from provider version 1.91.0. New field 'vswitch_ids' replace it.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The vswitches where new kubernetes cluster will be located.
     */
    vswitchIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When creating a cluster using automatic VPC creation, you need to specify the zone where the VPC is located.
     */
    zoneId?: pulumi.Input<string>;
}
