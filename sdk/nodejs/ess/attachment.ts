// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Attaches several ECS instances to a specified scaling group or remove them from it.
 *
 * > **NOTE:** ECS instances can be attached or remove only when the scaling group is active and it has no scaling activity in progress.
 *
 * > **NOTE:** There are two types ECS instances in a scaling group: "AutoCreated" and "Attached". The total number of them can not larger than the scaling group "MaxSize".
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "essattachmentconfig";
 * const defaultZones = alicloud.getZones({
 *     availableDiskCategory: "cloud_efficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultInstanceTypes = defaultZones.then(defaultZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones[0].id,
 *     cpuCoreCount: 2,
 *     memorySize: 4,
 * }));
 * const defaultImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_18.*64",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {cidrBlock: "172.16.0.0/16"});
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     availabilityZone: defaultZones.then(defaultZones => defaultZones.zones[0].id),
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {vpcId: defaultNetwork.id});
 * const defaultSecurityGroupRule = new alicloud.ecs.SecurityGroupRule("defaultSecurityGroupRule", {
 *     type: "ingress",
 *     ipProtocol: "tcp",
 *     nicType: "intranet",
 *     policy: "accept",
 *     portRange: "22/22",
 *     priority: 1,
 *     securityGroupId: defaultSecurityGroup.id,
 *     cidrIp: "172.16.0.0/24",
 * });
 * const defaultScalingGroup = new alicloud.ess.ScalingGroup("defaultScalingGroup", {
 *     minSize: 0,
 *     maxSize: 2,
 *     scalingGroupName: name,
 *     removalPolicies: [
 *         "OldestInstance",
 *         "NewestInstance",
 *     ],
 *     vswitchIds: [defaultSwitch.id],
 * });
 * const defaultScalingConfiguration = new alicloud.ess.ScalingConfiguration("defaultScalingConfiguration", {
 *     scalingGroupId: defaultScalingGroup.id,
 *     imageId: defaultImages.then(defaultImages => defaultImages.images[0].id),
 *     instanceType: defaultInstanceTypes.then(defaultInstanceTypes => defaultInstanceTypes.instanceTypes[0].id),
 *     securityGroupId: defaultSecurityGroup.id,
 *     forceDelete: true,
 *     active: true,
 *     enable: true,
 * });
 * const defaultInstance: alicloud.ecs.Instance[];
 * for (const range = {value: 0}; range.value < 2; range.value++) {
 *     defaultInstance.push(new alicloud.ecs.Instance(`defaultInstance-${range.value}`, {
 *         imageId: defaultImages.then(defaultImages => defaultImages.images[0].id),
 *         instanceType: defaultInstanceTypes.then(defaultInstanceTypes => defaultInstanceTypes.instanceTypes[0].id),
 *         securityGroups: [defaultSecurityGroup.id],
 *         internetChargeType: "PayByTraffic",
 *         internetMaxBandwidthOut: "10",
 *         instanceChargeType: "PostPaid",
 *         systemDiskCategory: "cloud_efficiency",
 *         vswitchId: defaultSwitch.id,
 *         instanceName: name,
 *     }));
 * }
 * const defaultAttachment = new alicloud.ess.Attachment("defaultAttachment", {
 *     scalingGroupId: defaultScalingGroup.id,
 *     instanceIds: [
 *         defaultInstance[0].id,
 *         defaultInstance[1].id,
 *     ],
 *     force: true,
 * });
 * ```
 *
 * ## Import
 *
 * ESS attachment can be imported using the id or scaling group id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ess/attachment:Attachment example asg-abc123456
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
     * Whether to remove forcibly "AutoCreated" ECS instances in order to release scaling group capacity "MaxSize" for attaching ECS instances. Default to false.
     */
    public readonly force!: pulumi.Output<boolean | undefined>;
    /**
     * ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.
     */
    public readonly instanceIds!: pulumi.Output<string[]>;
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AttachmentState | undefined;
            inputs["force"] = state ? state.force : undefined;
            inputs["instanceIds"] = state ? state.instanceIds : undefined;
            inputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
        } else {
            const args = argsOrState as AttachmentArgs | undefined;
            if ((!args || args.instanceIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceIds'");
            }
            if ((!args || args.scalingGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            inputs["force"] = args ? args.force : undefined;
            inputs["instanceIds"] = args ? args.instanceIds : undefined;
            inputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Attachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Attachment resources.
 */
export interface AttachmentState {
    /**
     * Whether to remove forcibly "AutoCreated" ECS instances in order to release scaling group capacity "MaxSize" for attaching ECS instances. Default to false.
     */
    readonly force?: pulumi.Input<boolean>;
    /**
     * ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.
     */
    readonly instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the scaling group of a scaling configuration.
     */
    readonly scalingGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Attachment resource.
 */
export interface AttachmentArgs {
    /**
     * Whether to remove forcibly "AutoCreated" ECS instances in order to release scaling group capacity "MaxSize" for attaching ECS instances. Default to false.
     */
    readonly force?: pulumi.Input<boolean>;
    /**
     * ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.
     */
    readonly instanceIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the scaling group of a scaling configuration.
     */
    readonly scalingGroupId: pulumi.Input<string>;
}
