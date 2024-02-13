// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Nlb Load Balancer Security Group Attachment resource.
 *
 * For information about Nlb Load Balancer Security Group Attachment and how to use it, see [What is Load Balancer Security Group Attachment](https://www.alibabacloud.com/help/en/server-load-balancer/latest/loadbalancerjoinsecuritygroup).
 *
 * > **NOTE:** Available since v1.198.0.
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
 * const defaultLoadBalancerSecurityGroupAttachment = new alicloud.nlb.LoadBalancerSecurityGroupAttachment("defaultLoadBalancerSecurityGroupAttachment", {
 *     securityGroupId: defaultSecurityGroup.id,
 *     loadBalancerId: defaultLoadBalancer.id,
 * });
 * ```
 *
 * ## Import
 *
 * Nlb Load Balancer Security Group Attachment can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:nlb/loadBalancerSecurityGroupAttachment:LoadBalancerSecurityGroupAttachment example <LoadBalancerId>:<SecurityGroupId>
 * ```
 */
export class LoadBalancerSecurityGroupAttachment extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancerSecurityGroupAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerSecurityGroupAttachmentState, opts?: pulumi.CustomResourceOptions): LoadBalancerSecurityGroupAttachment {
        return new LoadBalancerSecurityGroupAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:nlb/loadBalancerSecurityGroupAttachment:LoadBalancerSecurityGroupAttachment';

    /**
     * Returns true if the given object is an instance of LoadBalancerSecurityGroupAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancerSecurityGroupAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancerSecurityGroupAttachment.__pulumiType;
    }

    /**
     * Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind a security group to the instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the network-based server load balancer instance to be bound to the security group.
     */
    public readonly loadBalancerId!: pulumi.Output<string>;
    /**
     * The ID of security groups.
     */
    public readonly securityGroupId!: pulumi.Output<string>;

    /**
     * Create a LoadBalancerSecurityGroupAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerSecurityGroupAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerSecurityGroupAttachmentArgs | LoadBalancerSecurityGroupAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadBalancerSecurityGroupAttachmentState | undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["loadBalancerId"] = state ? state.loadBalancerId : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
        } else {
            const args = argsOrState as LoadBalancerSecurityGroupAttachmentArgs | undefined;
            if ((!args || args.loadBalancerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerId'");
            }
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["loadBalancerId"] = args ? args.loadBalancerId : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadBalancerSecurityGroupAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancerSecurityGroupAttachment resources.
 */
export interface LoadBalancerSecurityGroupAttachmentState {
    /**
     * Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind a security group to the instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the network-based server load balancer instance to be bound to the security group.
     */
    loadBalancerId?: pulumi.Input<string>;
    /**
     * The ID of security groups.
     */
    securityGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadBalancerSecurityGroupAttachment resource.
 */
export interface LoadBalancerSecurityGroupAttachmentArgs {
    /**
     * Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind a security group to the instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the network-based server load balancer instance to be bound to the security group.
     */
    loadBalancerId: pulumi.Input<string>;
    /**
     * The ID of security groups.
     */
    securityGroupId: pulumi.Input<string>;
}
