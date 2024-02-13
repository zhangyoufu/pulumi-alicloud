// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a NLB Loadbalancer Common Bandwidth Package Attachment resource. Bandwidth Package Operation.
 *
 * For information about NLB Loadbalancer Common Bandwidth Package Attachment and how to use it, see [What is Loadbalancer Common Bandwidth Package Attachment](https://www.alibabacloud.com/help/en/server-load-balancer/latest/nlb-instances-change).
 *
 * > **NOTE:** Available since v1.209.0.
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
 * const name = config.get("name") || "tf-example";
 * const defaultResourceGroups = alicloud.resourcemanager.getResourceGroups({});
 * const defaultZones = alicloud.nlb.getZones({});
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
 * const default1 = new alicloud.vpc.Switch("default1", {
 *     vswitchName: name,
 *     cidrBlock: "10.4.1.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[1]?.id),
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {vpcId: defaultNetwork.id});
 * const defaultLoadBalancer = new alicloud.nlb.LoadBalancer("defaultLoadBalancer", {
 *     loadBalancerName: name,
 *     resourceGroupId: defaultResourceGroups.then(defaultResourceGroups => defaultResourceGroups.ids?.[0]),
 *     loadBalancerType: "Network",
 *     addressType: "Internet",
 *     addressIpVersion: "Ipv4",
 *     vpcId: defaultNetwork.id,
 *     tags: {
 *         Created: "TF",
 *         For: "example",
 *     },
 *     zoneMappings: [
 *         {
 *             vswitchId: defaultSwitch.id,
 *             zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *         },
 *         {
 *             vswitchId: default1.id,
 *             zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[1]?.id),
 *         },
 *     ],
 * });
 * const defaultCommonBandwithPackage = new alicloud.vpc.CommonBandwithPackage("defaultCommonBandwithPackage", {
 *     bandwidth: "2",
 *     internetChargeType: "PayByBandwidth",
 *     bandwidthPackageName: name,
 *     description: name,
 * });
 * const defaultLoadbalancerCommonBandwidthPackageAttachment = new alicloud.nlb.LoadbalancerCommonBandwidthPackageAttachment("defaultLoadbalancerCommonBandwidthPackageAttachment", {
 *     bandwidthPackageId: defaultCommonBandwithPackage.id,
 *     loadBalancerId: defaultLoadBalancer.id,
 * });
 * ```
 *
 * ## Import
 *
 * NLB Loadbalancer Common Bandwidth Package Attachment can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:nlb/loadbalancerCommonBandwidthPackageAttachment:LoadbalancerCommonBandwidthPackageAttachment example <load_balancer_id>:<bandwidth_package_id>
 * ```
 */
export class LoadbalancerCommonBandwidthPackageAttachment extends pulumi.CustomResource {
    /**
     * Get an existing LoadbalancerCommonBandwidthPackageAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadbalancerCommonBandwidthPackageAttachmentState, opts?: pulumi.CustomResourceOptions): LoadbalancerCommonBandwidthPackageAttachment {
        return new LoadbalancerCommonBandwidthPackageAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:nlb/loadbalancerCommonBandwidthPackageAttachment:LoadbalancerCommonBandwidthPackageAttachment';

    /**
     * Returns true if the given object is an instance of LoadbalancerCommonBandwidthPackageAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadbalancerCommonBandwidthPackageAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadbalancerCommonBandwidthPackageAttachment.__pulumiType;
    }

    /**
     * The ID of the bound shared bandwidth package.
     */
    public readonly bandwidthPackageId!: pulumi.Output<string>;
    /**
     * The ID of the network-based server load balancer instance.
     */
    public readonly loadBalancerId!: pulumi.Output<string>;
    /**
     * Network-based load balancing instance status. Value:, indicating that the instance listener will no longer forward traffic.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a LoadbalancerCommonBandwidthPackageAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadbalancerCommonBandwidthPackageAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadbalancerCommonBandwidthPackageAttachmentArgs | LoadbalancerCommonBandwidthPackageAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadbalancerCommonBandwidthPackageAttachmentState | undefined;
            resourceInputs["bandwidthPackageId"] = state ? state.bandwidthPackageId : undefined;
            resourceInputs["loadBalancerId"] = state ? state.loadBalancerId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as LoadbalancerCommonBandwidthPackageAttachmentArgs | undefined;
            if ((!args || args.bandwidthPackageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidthPackageId'");
            }
            if ((!args || args.loadBalancerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerId'");
            }
            resourceInputs["bandwidthPackageId"] = args ? args.bandwidthPackageId : undefined;
            resourceInputs["loadBalancerId"] = args ? args.loadBalancerId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadbalancerCommonBandwidthPackageAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadbalancerCommonBandwidthPackageAttachment resources.
 */
export interface LoadbalancerCommonBandwidthPackageAttachmentState {
    /**
     * The ID of the bound shared bandwidth package.
     */
    bandwidthPackageId?: pulumi.Input<string>;
    /**
     * The ID of the network-based server load balancer instance.
     */
    loadBalancerId?: pulumi.Input<string>;
    /**
     * Network-based load balancing instance status. Value:, indicating that the instance listener will no longer forward traffic.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadbalancerCommonBandwidthPackageAttachment resource.
 */
export interface LoadbalancerCommonBandwidthPackageAttachmentArgs {
    /**
     * The ID of the bound shared bandwidth package.
     */
    bandwidthPackageId: pulumi.Input<string>;
    /**
     * The ID of the network-based server load balancer instance.
     */
    loadBalancerId: pulumi.Input<string>;
}
