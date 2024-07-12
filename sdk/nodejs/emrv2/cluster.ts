// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a EMR cluster resource. This resource is based on EMR's new version OpenAPI.
 *
 * For information about EMR New and how to use it, see [Add a domain](https://www.alibabacloud.com/help/doc-detail/28068.htm).
 *
 * > **NOTE:** Available since v1.199.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const default = alicloud.resourcemanager.getResourceGroups({
 *     status: "OK",
 * });
 * const defaultGetKeys = alicloud.kms.getKeys({
 *     status: "Enabled",
 * });
 * const defaultGetZones = alicloud.getZones({
 *     availableInstanceType: "ecs.g7.xlarge",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     vpcName: name,
 *     cidrBlock: "172.16.0.0/12",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/21",
 *     zoneId: defaultGetZones.then(defaultGetZones => defaultGetZones.zones?.[0]?.id),
 *     vswitchName: name,
 * });
 * const defaultInteger = new random.index.Integer("default", {
 *     max: 99999,
 *     min: 10000,
 * });
 * const defaultEcsKeyPair = new alicloud.ecs.EcsKeyPair("default", {keyPairName: `${name}-${defaultInteger.result}`});
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("default", {
 *     name: name,
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultRole = new alicloud.ram.Role("default", {
 *     name: name,
 *     document: `    {
 *         "Statement": [
 *         {
 *             "Action": "sts:AssumeRole",
 *             "Effect": "Allow",
 *             "Principal": {
 *             "Service": [
 *                 "emr.aliyuncs.com",
 *                 "ecs.aliyuncs.com"
 *             ]
 *             }
 *         }
 *         ],
 *         "Version": "1"
 *     }
 * `,
 *     description: "this is a role example.",
 *     force: true,
 * });
 * const defaultCluster = new alicloud.emrv2.Cluster("default", {
 *     nodeGroups: [
 *         {
 *             vswitchIds: [defaultSwitch.id],
 *             instanceTypes: ["ecs.g7.xlarge"],
 *             nodeCount: 1,
 *             spotInstanceRemedy: false,
 *             dataDisks: [{
 *                 count: 3,
 *                 category: "cloud_essd",
 *                 size: 80,
 *                 performanceLevel: "PL0",
 *             }],
 *             nodeGroupName: "emr-master",
 *             paymentType: "PayAsYouGo",
 *             withPublicIp: false,
 *             gracefulShutdown: false,
 *             systemDisk: {
 *                 category: "cloud_essd",
 *                 size: 80,
 *                 performanceLevel: "PL0",
 *                 count: 1,
 *             },
 *             nodeGroupType: "MASTER",
 *         },
 *         {
 *             spotInstanceRemedy: false,
 *             nodeGroupType: "CORE",
 *             vswitchIds: [defaultSwitch.id],
 *             nodeCount: 2,
 *             gracefulShutdown: false,
 *             systemDisk: {
 *                 performanceLevel: "PL0",
 *                 count: 1,
 *                 category: "cloud_essd",
 *                 size: 80,
 *             },
 *             dataDisks: [{
 *                 count: 3,
 *                 performanceLevel: "PL0",
 *                 category: "cloud_essd",
 *                 size: 80,
 *             }],
 *             nodeGroupName: "emr-core",
 *             paymentType: "PayAsYouGo",
 *             instanceTypes: ["ecs.g7.xlarge"],
 *             withPublicIp: false,
 *         },
 *     ],
 *     deployMode: "NORMAL",
 *     tags: {
 *         Created: "TF",
 *         For: "example",
 *     },
 *     releaseVersion: "EMR-5.10.0",
 *     applications: [
 *         "HADOOP-COMMON",
 *         "HDFS",
 *         "YARN",
 *     ],
 *     nodeAttributes: [{
 *         zoneId: defaultGetZones.then(defaultGetZones => defaultGetZones.zones?.[0]?.id),
 *         keyPairName: defaultEcsKeyPair.id,
 *         dataDiskEncrypted: true,
 *         dataDiskKmsKeyId: defaultGetKeys.then(defaultGetKeys => defaultGetKeys.ids?.[0]),
 *         vpcId: defaultNetwork.id,
 *         ramRole: defaultRole.name,
 *         securityGroupId: defaultSecurityGroup.id,
 *     }],
 *     resourceGroupId: _default.then(_default => _default.ids?.[0]),
 *     clusterName: name,
 *     paymentType: "PayAsYouGo",
 *     clusterType: "DATAFLOW",
 * });
 * ```
 *
 * ## Import
 *
 * Aliclioud E-MapReduce cluster can be imported using the id e.g.
 *
 * ```sh
 * $ pulumi import alicloud:emrv2/cluster:Cluster default <id>
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:emrv2/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * The application configurations of EMR cluster. See `applicationConfigs` below.
     */
    public readonly applicationConfigs!: pulumi.Output<outputs.emrv2.ClusterApplicationConfig[] | undefined>;
    /**
     * The applications of EMR cluster to be installed, e.g. HADOOP-COMMON, HDFS, YARN, HIVE, SPARK2, SPARK3, ZOOKEEPER etc. You can find all valid applications in emr web console.
     */
    public readonly applications!: pulumi.Output<string[]>;
    /**
     * The bootstrap scripts to be effected when creating emr-cluster or resize emr-cluster, if priority is not specified, the scripts will execute in the declared order. See `bootstrapScripts` below.
     */
    public readonly bootstrapScripts!: pulumi.Output<outputs.emrv2.ClusterBootstrapScript[] | undefined>;
    /**
     * The name of emr cluster. The name length must be less than 64. Supported characters: chinese character, english character, number, "-", "_".
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * EMR Cluster Type, e.g. DATALAKE, OLAP, DATAFLOW, DATASERVING, CUSTOM etc. You can find all valid EMR cluster type in emr web console.
     */
    public readonly clusterType!: pulumi.Output<string>;
    /**
     * The deploy mode of EMR cluster. Supported value: NORMAL or HA.
     */
    public readonly deployMode!: pulumi.Output<string>;
    /**
     * The log collect strategy of EMR cluster.
     */
    public readonly logCollectStrategy!: pulumi.Output<string>;
    /**
     * The node attributes of ecs instances which the emr-cluster belongs. See `nodeAttributes` below.
     */
    public readonly nodeAttributes!: pulumi.Output<outputs.emrv2.ClusterNodeAttribute[]>;
    /**
     * Groups of node, You can specify MASTER as a group, CORE as a group (just like the above example). See `nodeGroups` below. **NOTE:** Since version 1.227.0, the type of `nodeGroups` changed from Set to List.
     */
    public readonly nodeGroups!: pulumi.Output<outputs.emrv2.ClusterNodeGroup[]>;
    /**
     * Payment Type for this cluster. Supported value: PayAsYouGo or Subscription. **NOTE:** From version 1.227.0, `paymentType` can be modified.
     */
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * EMR Version, e.g. EMR-5.10.0. You can find the all valid EMR Version in emr web console.
     */
    public readonly releaseVersion!: pulumi.Output<string>;
    /**
     * The Id of resource group which the emr-cluster belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The security mode of EMR cluster. Supported value: NORMAL or KERBEROS.
     */
    public readonly securityMode!: pulumi.Output<string>;
    /**
     * The detail configuration of subscription payment type. See `subscriptionConfig` below.
     */
    public readonly subscriptionConfig!: pulumi.Output<outputs.emrv2.ClusterSubscriptionConfig | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            resourceInputs["applicationConfigs"] = state ? state.applicationConfigs : undefined;
            resourceInputs["applications"] = state ? state.applications : undefined;
            resourceInputs["bootstrapScripts"] = state ? state.bootstrapScripts : undefined;
            resourceInputs["clusterName"] = state ? state.clusterName : undefined;
            resourceInputs["clusterType"] = state ? state.clusterType : undefined;
            resourceInputs["deployMode"] = state ? state.deployMode : undefined;
            resourceInputs["logCollectStrategy"] = state ? state.logCollectStrategy : undefined;
            resourceInputs["nodeAttributes"] = state ? state.nodeAttributes : undefined;
            resourceInputs["nodeGroups"] = state ? state.nodeGroups : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["releaseVersion"] = state ? state.releaseVersion : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["securityMode"] = state ? state.securityMode : undefined;
            resourceInputs["subscriptionConfig"] = state ? state.subscriptionConfig : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.applications === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applications'");
            }
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.clusterType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterType'");
            }
            if ((!args || args.nodeAttributes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeAttributes'");
            }
            if ((!args || args.nodeGroups === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeGroups'");
            }
            if ((!args || args.releaseVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'releaseVersion'");
            }
            resourceInputs["applicationConfigs"] = args ? args.applicationConfigs : undefined;
            resourceInputs["applications"] = args ? args.applications : undefined;
            resourceInputs["bootstrapScripts"] = args ? args.bootstrapScripts : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["clusterType"] = args ? args.clusterType : undefined;
            resourceInputs["deployMode"] = args ? args.deployMode : undefined;
            resourceInputs["logCollectStrategy"] = args ? args.logCollectStrategy : undefined;
            resourceInputs["nodeAttributes"] = args ? args.nodeAttributes : undefined;
            resourceInputs["nodeGroups"] = args ? args.nodeGroups : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["releaseVersion"] = args ? args.releaseVersion : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["securityMode"] = args ? args.securityMode : undefined;
            resourceInputs["subscriptionConfig"] = args ? args.subscriptionConfig : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * The application configurations of EMR cluster. See `applicationConfigs` below.
     */
    applicationConfigs?: pulumi.Input<pulumi.Input<inputs.emrv2.ClusterApplicationConfig>[]>;
    /**
     * The applications of EMR cluster to be installed, e.g. HADOOP-COMMON, HDFS, YARN, HIVE, SPARK2, SPARK3, ZOOKEEPER etc. You can find all valid applications in emr web console.
     */
    applications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The bootstrap scripts to be effected when creating emr-cluster or resize emr-cluster, if priority is not specified, the scripts will execute in the declared order. See `bootstrapScripts` below.
     */
    bootstrapScripts?: pulumi.Input<pulumi.Input<inputs.emrv2.ClusterBootstrapScript>[]>;
    /**
     * The name of emr cluster. The name length must be less than 64. Supported characters: chinese character, english character, number, "-", "_".
     */
    clusterName?: pulumi.Input<string>;
    /**
     * EMR Cluster Type, e.g. DATALAKE, OLAP, DATAFLOW, DATASERVING, CUSTOM etc. You can find all valid EMR cluster type in emr web console.
     */
    clusterType?: pulumi.Input<string>;
    /**
     * The deploy mode of EMR cluster. Supported value: NORMAL or HA.
     */
    deployMode?: pulumi.Input<string>;
    /**
     * The log collect strategy of EMR cluster.
     */
    logCollectStrategy?: pulumi.Input<string>;
    /**
     * The node attributes of ecs instances which the emr-cluster belongs. See `nodeAttributes` below.
     */
    nodeAttributes?: pulumi.Input<pulumi.Input<inputs.emrv2.ClusterNodeAttribute>[]>;
    /**
     * Groups of node, You can specify MASTER as a group, CORE as a group (just like the above example). See `nodeGroups` below. **NOTE:** Since version 1.227.0, the type of `nodeGroups` changed from Set to List.
     */
    nodeGroups?: pulumi.Input<pulumi.Input<inputs.emrv2.ClusterNodeGroup>[]>;
    /**
     * Payment Type for this cluster. Supported value: PayAsYouGo or Subscription. **NOTE:** From version 1.227.0, `paymentType` can be modified.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * EMR Version, e.g. EMR-5.10.0. You can find the all valid EMR Version in emr web console.
     */
    releaseVersion?: pulumi.Input<string>;
    /**
     * The Id of resource group which the emr-cluster belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The security mode of EMR cluster. Supported value: NORMAL or KERBEROS.
     */
    securityMode?: pulumi.Input<string>;
    /**
     * The detail configuration of subscription payment type. See `subscriptionConfig` below.
     */
    subscriptionConfig?: pulumi.Input<inputs.emrv2.ClusterSubscriptionConfig>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The application configurations of EMR cluster. See `applicationConfigs` below.
     */
    applicationConfigs?: pulumi.Input<pulumi.Input<inputs.emrv2.ClusterApplicationConfig>[]>;
    /**
     * The applications of EMR cluster to be installed, e.g. HADOOP-COMMON, HDFS, YARN, HIVE, SPARK2, SPARK3, ZOOKEEPER etc. You can find all valid applications in emr web console.
     */
    applications: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The bootstrap scripts to be effected when creating emr-cluster or resize emr-cluster, if priority is not specified, the scripts will execute in the declared order. See `bootstrapScripts` below.
     */
    bootstrapScripts?: pulumi.Input<pulumi.Input<inputs.emrv2.ClusterBootstrapScript>[]>;
    /**
     * The name of emr cluster. The name length must be less than 64. Supported characters: chinese character, english character, number, "-", "_".
     */
    clusterName: pulumi.Input<string>;
    /**
     * EMR Cluster Type, e.g. DATALAKE, OLAP, DATAFLOW, DATASERVING, CUSTOM etc. You can find all valid EMR cluster type in emr web console.
     */
    clusterType: pulumi.Input<string>;
    /**
     * The deploy mode of EMR cluster. Supported value: NORMAL or HA.
     */
    deployMode?: pulumi.Input<string>;
    /**
     * The log collect strategy of EMR cluster.
     */
    logCollectStrategy?: pulumi.Input<string>;
    /**
     * The node attributes of ecs instances which the emr-cluster belongs. See `nodeAttributes` below.
     */
    nodeAttributes: pulumi.Input<pulumi.Input<inputs.emrv2.ClusterNodeAttribute>[]>;
    /**
     * Groups of node, You can specify MASTER as a group, CORE as a group (just like the above example). See `nodeGroups` below. **NOTE:** Since version 1.227.0, the type of `nodeGroups` changed from Set to List.
     */
    nodeGroups: pulumi.Input<pulumi.Input<inputs.emrv2.ClusterNodeGroup>[]>;
    /**
     * Payment Type for this cluster. Supported value: PayAsYouGo or Subscription. **NOTE:** From version 1.227.0, `paymentType` can be modified.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * EMR Version, e.g. EMR-5.10.0. You can find the all valid EMR Version in emr web console.
     */
    releaseVersion: pulumi.Input<string>;
    /**
     * The Id of resource group which the emr-cluster belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The security mode of EMR cluster. Supported value: NORMAL or KERBEROS.
     */
    securityMode?: pulumi.Input<string>;
    /**
     * The detail configuration of subscription payment type. See `subscriptionConfig` below.
     */
    subscriptionConfig?: pulumi.Input<inputs.emrv2.ClusterSubscriptionConfig>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
