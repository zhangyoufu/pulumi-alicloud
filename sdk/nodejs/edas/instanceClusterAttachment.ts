// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an EDAS instance cluster attachment resource, see [What is EDAS Instance Cluster Attachment](https://www.alibabacloud.com/help/en/edas/developer-reference/api-edas-2017-08-01-installagent).
 *
 * > **NOTE:** Available since v1.82.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const defaultRegions = alicloud.getRegions({
 *     current: true,
 * });
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_[0-9]+_[0-9]+_x64*",
 *     owners: "system",
 * });
 * const defaultInstanceTypes = defaultZones.then(defaultZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones?.[0]?.id,
 *     cpuCoreCount: 1,
 *     memorySize: 2,
 * }));
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "10.4.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vswitchName: name,
 *     cidrBlock: "10.4.0.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {vpcId: defaultNetwork.id});
 * const defaultInstance = new alicloud.ecs.Instance("defaultInstance", {
 *     availabilityZone: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     instanceName: name,
 *     imageId: defaultImages.then(defaultImages => defaultImages.images?.[0]?.id),
 *     instanceType: defaultInstanceTypes.then(defaultInstanceTypes => defaultInstanceTypes.instanceTypes?.[0]?.id),
 *     securityGroups: [defaultSecurityGroup.id],
 *     vswitchId: defaultSwitch.id,
 * });
 * const defaultCluster = new alicloud.edas.Cluster("defaultCluster", {
 *     clusterName: name,
 *     clusterType: 2,
 *     networkMode: 2,
 *     logicalRegionId: defaultRegions.then(defaultRegions => defaultRegions.regions?.[0]?.id),
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultInstanceClusterAttachment = new alicloud.edas.InstanceClusterAttachment("defaultInstanceClusterAttachment", {
 *     clusterId: defaultCluster.id,
 *     instanceIds: [defaultInstance.id],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class InstanceClusterAttachment extends pulumi.CustomResource {
    /**
     * Get an existing InstanceClusterAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceClusterAttachmentState, opts?: pulumi.CustomResourceOptions): InstanceClusterAttachment {
        return new InstanceClusterAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:edas/instanceClusterAttachment:InstanceClusterAttachment';

    /**
     * Returns true if the given object is an instance of InstanceClusterAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceClusterAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceClusterAttachment.__pulumiType;
    }

    /**
     * The ID of the cluster that you want to create the application.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The cluster members map of the resource supplied above. The key is instanceId and the value is cluster_member_id.
     */
    public /*out*/ readonly clusterMemberIds!: pulumi.Output<{[key: string]: string}>;
    /**
     * The ecu map of the resource supplied above. The key is instanceId and the value is ecu_id.
     */
    public /*out*/ readonly ecuMap!: pulumi.Output<{[key: string]: string}>;
    /**
     * The ID of instance. Type: list.
     */
    public readonly instanceIds!: pulumi.Output<string[]>;
    /**
     * The status map of the resource supplied above. The key is instanceId and the values are 1(running) 0(converting) -1(failed) and -2(offline).
     */
    public /*out*/ readonly statusMap!: pulumi.Output<{[key: string]: number}>;

    /**
     * Create a InstanceClusterAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceClusterAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceClusterAttachmentArgs | InstanceClusterAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceClusterAttachmentState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["clusterMemberIds"] = state ? state.clusterMemberIds : undefined;
            resourceInputs["ecuMap"] = state ? state.ecuMap : undefined;
            resourceInputs["instanceIds"] = state ? state.instanceIds : undefined;
            resourceInputs["statusMap"] = state ? state.statusMap : undefined;
        } else {
            const args = argsOrState as InstanceClusterAttachmentArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.instanceIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceIds'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["instanceIds"] = args ? args.instanceIds : undefined;
            resourceInputs["clusterMemberIds"] = undefined /*out*/;
            resourceInputs["ecuMap"] = undefined /*out*/;
            resourceInputs["statusMap"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceClusterAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceClusterAttachment resources.
 */
export interface InstanceClusterAttachmentState {
    /**
     * The ID of the cluster that you want to create the application.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The cluster members map of the resource supplied above. The key is instanceId and the value is cluster_member_id.
     */
    clusterMemberIds?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ecu map of the resource supplied above. The key is instanceId and the value is ecu_id.
     */
    ecuMap?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of instance. Type: list.
     */
    instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status map of the resource supplied above. The key is instanceId and the values are 1(running) 0(converting) -1(failed) and -2(offline).
     */
    statusMap?: pulumi.Input<{[key: string]: pulumi.Input<number>}>;
}

/**
 * The set of arguments for constructing a InstanceClusterAttachment resource.
 */
export interface InstanceClusterAttachmentArgs {
    /**
     * The ID of the cluster that you want to create the application.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The ID of instance. Type: list.
     */
    instanceIds: pulumi.Input<pulumi.Input<string>[]>;
}
