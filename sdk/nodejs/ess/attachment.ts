// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Attaches several ECS instances to a specified scaling group or remove them from it.
 *
 * > **NOTE:** ECS instances can be attached or remove only when the scaling group is active, and it has no scaling activity in progress.
 *
 * > **NOTE:** There are two types ECS instances in a scaling group: "AutoCreated" and "Attached". The total number of them can not larger than the scaling group "MaxSize".
 *
 * > **NOTE:** Available since v1.6.0.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const default = alicloud.getZones({
 *     availableDiskCategory: "cloud_efficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultGetInstanceTypes = _default.then(_default => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: _default.zones?.[0]?.id,
 *     cpuCoreCount: 2,
 *     memorySize: 4,
 * }));
 * const defaultGetImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_18.*64",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     vpcName: name,
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: _default.then(_default => _default.zones?.[0]?.id),
 *     vswitchName: name,
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("default", {
 *     name: name,
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultSecurityGroupRule = new alicloud.ecs.SecurityGroupRule("default", {
 *     type: "ingress",
 *     ipProtocol: "tcp",
 *     nicType: "intranet",
 *     policy: "accept",
 *     portRange: "22/22",
 *     priority: 1,
 *     securityGroupId: defaultSecurityGroup.id,
 *     cidrIp: "172.16.0.0/24",
 * });
 * const defaultScalingGroup = new alicloud.ess.ScalingGroup("default", {
 *     minSize: 0,
 *     maxSize: 2,
 *     scalingGroupName: name,
 *     removalPolicies: [
 *         "OldestInstance",
 *         "NewestInstance",
 *     ],
 *     vswitchIds: [defaultSwitch.id],
 * });
 * const defaultScalingConfiguration = new alicloud.ess.ScalingConfiguration("default", {
 *     scalingGroupId: defaultScalingGroup.id,
 *     imageId: defaultGetImages.then(defaultGetImages => defaultGetImages.images?.[0]?.id),
 *     instanceType: defaultGetInstanceTypes.then(defaultGetInstanceTypes => defaultGetInstanceTypes.instanceTypes?.[0]?.id),
 *     securityGroupId: defaultSecurityGroup.id,
 *     forceDelete: true,
 *     active: true,
 *     enable: true,
 * });
 * const defaultInstance: alicloud.ecs.Instance[] = [];
 * for (const range = {value: 0}; range.value < 2; range.value++) {
 *     defaultInstance.push(new alicloud.ecs.Instance(`default-${range.value}`, {
 *         imageId: defaultGetImages.then(defaultGetImages => defaultGetImages.images?.[0]?.id),
 *         instanceType: defaultGetInstanceTypes.then(defaultGetInstanceTypes => defaultGetInstanceTypes.instanceTypes?.[0]?.id),
 *         securityGroups: [defaultSecurityGroup.id],
 *         internetChargeType: "PayByTraffic",
 *         internetMaxBandwidthOut: 10,
 *         instanceChargeType: "PostPaid",
 *         systemDiskCategory: "cloud_efficiency",
 *         vswitchId: defaultSwitch.id,
 *         instanceName: name,
 *     }));
 * }
 * const defaultAttachment = new alicloud.ess.Attachment("default", {
 *     scalingGroupId: defaultScalingGroup.id,
 *     instanceIds: [
 *         defaultInstance[0].id,
 *         defaultInstance[1].id,
 *     ],
 *     force: true,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ESS attachment can be imported using the id or scaling group id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ess/attachment:Attachment example asg-abc123456
 * ```
 */
export class Attachment extends pulumi.CustomResource {
    /**
     * Get an existing Attachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AttachmentState, opts?: pulumi.CustomResourceOptions): Attachment {
        return new Attachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ess/attachment:Attachment';

    /**
     * Returns true if the given object is an instance of Attachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Attachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Attachment.__pulumiType;
    }

    /**
     * Specifies whether the scaling group manages the lifecycles of the instances that are manually added to the scaling group.
     */
    public readonly entrusted!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to remove forcibly "AutoCreated" ECS instances in order to release scaling group capacity "MaxSize" for attaching ECS instances. Default to false.
     */
    public readonly force!: pulumi.Output<boolean | undefined>;
    /**
     * ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.
     */
    public readonly instanceIds!: pulumi.Output<string[]>;
    /**
     * Specifies whether to trigger a lifecycle hook for the scaling group to which instances are being added.
     */
    public readonly lifecycleHook!: pulumi.Output<boolean | undefined>;
    /**
     * The weight of ECS instance N or elastic container instance N as a backend server of the associated Server Load Balancer (SLB) instance. Valid values of N: 1 to 20. Valid values of this parameter: 1 to 100.
     *
     * > **NOTE:** "AutoCreated" ECS instance will be deleted after it is removed from scaling group, but "Attached" will be not.
     *
     * > **NOTE:** Restrictions on attaching ECS instances:
     *
     * - The attached ECS instances and the scaling group must have the same region and network type(`Classic` or `VPC`).
     * - The attached ECS instances and the instance with active scaling configurations must have the same instance type.
     * - The attached ECS instances must in the running state.
     * - The attached ECS instances has not been attached to other scaling groups.
     * - The attached ECS instances supports Subscription and Pay-As-You-Go payment methods.
     */
    public readonly loadBalancerWeights!: pulumi.Output<number[]>;
    /**
     * ID of the scaling group of a scaling configuration.
     */
    public readonly scalingGroupId!: pulumi.Output<string>;

    /**
     * Create a Attachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AttachmentArgs | AttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AttachmentState | undefined;
            resourceInputs["entrusted"] = state ? state.entrusted : undefined;
            resourceInputs["force"] = state ? state.force : undefined;
            resourceInputs["instanceIds"] = state ? state.instanceIds : undefined;
            resourceInputs["lifecycleHook"] = state ? state.lifecycleHook : undefined;
            resourceInputs["loadBalancerWeights"] = state ? state.loadBalancerWeights : undefined;
            resourceInputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
        } else {
            const args = argsOrState as AttachmentArgs | undefined;
            if ((!args || args.instanceIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceIds'");
            }
            if ((!args || args.scalingGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            resourceInputs["entrusted"] = args ? args.entrusted : undefined;
            resourceInputs["force"] = args ? args.force : undefined;
            resourceInputs["instanceIds"] = args ? args.instanceIds : undefined;
            resourceInputs["lifecycleHook"] = args ? args.lifecycleHook : undefined;
            resourceInputs["loadBalancerWeights"] = args ? args.loadBalancerWeights : undefined;
            resourceInputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Attachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Attachment resources.
 */
export interface AttachmentState {
    /**
     * Specifies whether the scaling group manages the lifecycles of the instances that are manually added to the scaling group.
     */
    entrusted?: pulumi.Input<boolean>;
    /**
     * Whether to remove forcibly "AutoCreated" ECS instances in order to release scaling group capacity "MaxSize" for attaching ECS instances. Default to false.
     */
    force?: pulumi.Input<boolean>;
    /**
     * ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.
     */
    instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies whether to trigger a lifecycle hook for the scaling group to which instances are being added.
     */
    lifecycleHook?: pulumi.Input<boolean>;
    /**
     * The weight of ECS instance N or elastic container instance N as a backend server of the associated Server Load Balancer (SLB) instance. Valid values of N: 1 to 20. Valid values of this parameter: 1 to 100.
     *
     * > **NOTE:** "AutoCreated" ECS instance will be deleted after it is removed from scaling group, but "Attached" will be not.
     *
     * > **NOTE:** Restrictions on attaching ECS instances:
     *
     * - The attached ECS instances and the scaling group must have the same region and network type(`Classic` or `VPC`).
     * - The attached ECS instances and the instance with active scaling configurations must have the same instance type.
     * - The attached ECS instances must in the running state.
     * - The attached ECS instances has not been attached to other scaling groups.
     * - The attached ECS instances supports Subscription and Pay-As-You-Go payment methods.
     */
    loadBalancerWeights?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * ID of the scaling group of a scaling configuration.
     */
    scalingGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Attachment resource.
 */
export interface AttachmentArgs {
    /**
     * Specifies whether the scaling group manages the lifecycles of the instances that are manually added to the scaling group.
     */
    entrusted?: pulumi.Input<boolean>;
    /**
     * Whether to remove forcibly "AutoCreated" ECS instances in order to release scaling group capacity "MaxSize" for attaching ECS instances. Default to false.
     */
    force?: pulumi.Input<boolean>;
    /**
     * ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.
     */
    instanceIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies whether to trigger a lifecycle hook for the scaling group to which instances are being added.
     */
    lifecycleHook?: pulumi.Input<boolean>;
    /**
     * The weight of ECS instance N or elastic container instance N as a backend server of the associated Server Load Balancer (SLB) instance. Valid values of N: 1 to 20. Valid values of this parameter: 1 to 100.
     *
     * > **NOTE:** "AutoCreated" ECS instance will be deleted after it is removed from scaling group, but "Attached" will be not.
     *
     * > **NOTE:** Restrictions on attaching ECS instances:
     *
     * - The attached ECS instances and the scaling group must have the same region and network type(`Classic` or `VPC`).
     * - The attached ECS instances and the instance with active scaling configurations must have the same instance type.
     * - The attached ECS instances must in the running state.
     * - The attached ECS instances has not been attached to other scaling groups.
     * - The attached ECS instances supports Subscription and Pay-As-You-Go payment methods.
     */
    loadBalancerWeights?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * ID of the scaling group of a scaling configuration.
     */
    scalingGroupId: pulumi.Input<string>;
}
