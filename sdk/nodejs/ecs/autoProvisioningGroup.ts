// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a ECS auto provisioning group resource which is a solution that uses preemptive instances and payAsYouGo instances to rapidly deploy clusters.
 *
 * > **NOTE:** Available in 1.79.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "auto_provisioning_group";
 * const defaultZones = alicloud.getZones({
 *     availableDiskCategory: "cloud_efficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?[0]?.id),
 *     vswitchName: name,
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {vpcId: defaultNetwork.id});
 * const defaultImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_18.*64",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const template = new alicloud.ecs.EcsLaunchTemplate("template", {
 *     imageId: defaultImages.then(defaultImages => defaultImages.images?[0]?.id),
 *     instanceType: "ecs.n1.tiny",
 *     securityGroupId: defaultSecurityGroup.id,
 * });
 * const defaultAutoProvisioningGroup = new alicloud.ecs.AutoProvisioningGroup("defaultAutoProvisioningGroup", {
 *     launchTemplateId: template.id,
 *     totalTargetCapacity: "4",
 *     payAsYouGoTargetCapacity: "1",
 *     spotTargetCapacity: "2",
 *     launchTemplateConfigs: [{
 *         instanceType: "ecs.n1.small",
 *         vswitchId: defaultSwitch.id,
 *         weightedCapacity: "2",
 *         maxPrice: "2",
 *     }],
 * });
 * ```
 * ## Block config
 *
 * The config mapping supports the following:
 * * `instanceType` - (Optional) The instance type of the Nth extended configurations of the launch template.
 * * `maxPrice` - (Required) The maximum price of the instance type specified in the Nth extended configurations of the launch template.
 * * `vswitchId` - (Required) The ID of the VSwitch in the Nth extended configurations of the launch template.
 * * `weightedCapacity` - (Required) The weight of the instance type specified in the Nth extended configurations of the launch template.
 * * `priority` - (Optional) The priority of the instance type specified in the Nth extended configurations of the launch template. A value of 0 indicates the highest priority.
 *
 * ## Import
 *
 * ECS auto provisioning group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ecs/autoProvisioningGroup:AutoProvisioningGroup example asg-abc123456
 * ```
 */
export class AutoProvisioningGroup extends pulumi.CustomResource {
    /**
     * Get an existing AutoProvisioningGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoProvisioningGroupState, opts?: pulumi.CustomResourceOptions): AutoProvisioningGroup {
        return new AutoProvisioningGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/autoProvisioningGroup:AutoProvisioningGroup';

    /**
     * Returns true if the given object is an instance of AutoProvisioningGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutoProvisioningGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutoProvisioningGroup.__pulumiType;
    }

    /**
     * The name of the auto provisioning group to be created. It must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-)
     */
    public readonly autoProvisioningGroupName!: pulumi.Output<string>;
    /**
     * The type of the auto provisioning group. Valid values:`request` and `maintain`,Default value: `maintain`.
     */
    public readonly autoProvisioningGroupType!: pulumi.Output<string | undefined>;
    /**
     * The type of supplemental instances. When the total value of `PayAsYouGoTargetCapacity` and `SpotTargetCapacity` is smaller than the value of TotalTargetCapacity, the auto provisioning group will create instances of the specified type to meet the capacity requirements. Valid values:`PayAsYouGo`: Pay-as-you-go instances; `Spot`: Preemptible instances, Default value: `Spot`.
     */
    public readonly defaultTargetCapacityType!: pulumi.Output<string | undefined>;
    /**
     * The description of the auto provisioning group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The shutdown policy for excess preemptible instances followed when the capacity of the auto provisioning group exceeds the target capacity. Valid values: `no-termination` and `termination`,Default value: `no-termination`.
     */
    public readonly excessCapacityTerminationPolicy!: pulumi.Output<string | undefined>;
    /**
     * DataDisk mappings to attach to ecs instance. See Block config below for details.
     */
    public readonly launchTemplateConfigs!: pulumi.Output<outputs.ecs.AutoProvisioningGroupLaunchTemplateConfig[]>;
    /**
     * The ID of the instance launch template associated with the auto provisioning group.
     */
    public readonly launchTemplateId!: pulumi.Output<string>;
    /**
     * The version of the instance launch template associated with the auto provisioning group.
     */
    public readonly launchTemplateVersion!: pulumi.Output<string>;
    /**
     * The global maximum price for preemptible instances in the auto provisioning group. If both the `MaxSpotPrice` and `LaunchTemplateConfig.N.MaxPrice` parameters are specified, the maximum price is the lower value of the two.
     */
    public readonly maxSpotPrice!: pulumi.Output<number>;
    /**
     * The scale-out policy for pay-as-you-go instances. Valid values: `lowest-price` and `prioritized`,Default value: `lowest-price`.
     */
    public readonly payAsYouGoAllocationStrategy!: pulumi.Output<string | undefined>;
    /**
     * The target capacity of pay-as-you-go instances in the auto provisioning group.
     */
    public readonly payAsYouGoTargetCapacity!: pulumi.Output<string | undefined>;
    /**
     * The scale-out policy for preemptible instances. Valid values:`lowest-price` and `diversified`,Default value: `lowest-price`.
     */
    public readonly spotAllocationStrategy!: pulumi.Output<string | undefined>;
    /**
     * The default behavior after preemptible instances are shut down. Valid values: `stop` and `terminate`,Default value: `stop`.
     */
    public readonly spotInstanceInterruptionBehavior!: pulumi.Output<string | undefined>;
    /**
     * This parameter takes effect when the `SpotAllocationStrategy` parameter is set to `lowest-price`. The auto provisioning group selects instance types of the lowest cost to create instances.
     */
    public readonly spotInstancePoolsToUseCount!: pulumi.Output<number>;
    /**
     * The target capacity of preemptible instances in the auto provisioning group.
     */
    public readonly spotTargetCapacity!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to release instances of the auto provisioning group. Valid values:`false` and `true`, default value: `false`.
     */
    public readonly terminateInstances!: pulumi.Output<boolean | undefined>;
    /**
     * The shutdown policy for preemptible instances when the auto provisioning group expires. Valid values: `false` and `true`, default value: `false`.
     */
    public readonly terminateInstancesWithExpiration!: pulumi.Output<boolean | undefined>;
    /**
     * The total target capacity of the auto provisioning group. The target capacity consists of the following three parts:PayAsYouGoTargetCapacity,SpotTargetCapacity and the supplemental capacity besides PayAsYouGoTargetCapacity and SpotTargetCapacity.
     */
    public readonly totalTargetCapacity!: pulumi.Output<string>;
    /**
     * The time when the auto provisioning group is started. The period of time between this point in time and the point in time specified by the `validUntil` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group is immediately started after creation.
     */
    public readonly validFrom!: pulumi.Output<string>;
    /**
     * The time when the auto provisioning group expires. The period of time between this point in time and the point in time specified by the `validFrom` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group never expires.
     */
    public readonly validUntil!: pulumi.Output<string>;

    /**
     * Create a AutoProvisioningGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutoProvisioningGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutoProvisioningGroupArgs | AutoProvisioningGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutoProvisioningGroupState | undefined;
            resourceInputs["autoProvisioningGroupName"] = state ? state.autoProvisioningGroupName : undefined;
            resourceInputs["autoProvisioningGroupType"] = state ? state.autoProvisioningGroupType : undefined;
            resourceInputs["defaultTargetCapacityType"] = state ? state.defaultTargetCapacityType : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["excessCapacityTerminationPolicy"] = state ? state.excessCapacityTerminationPolicy : undefined;
            resourceInputs["launchTemplateConfigs"] = state ? state.launchTemplateConfigs : undefined;
            resourceInputs["launchTemplateId"] = state ? state.launchTemplateId : undefined;
            resourceInputs["launchTemplateVersion"] = state ? state.launchTemplateVersion : undefined;
            resourceInputs["maxSpotPrice"] = state ? state.maxSpotPrice : undefined;
            resourceInputs["payAsYouGoAllocationStrategy"] = state ? state.payAsYouGoAllocationStrategy : undefined;
            resourceInputs["payAsYouGoTargetCapacity"] = state ? state.payAsYouGoTargetCapacity : undefined;
            resourceInputs["spotAllocationStrategy"] = state ? state.spotAllocationStrategy : undefined;
            resourceInputs["spotInstanceInterruptionBehavior"] = state ? state.spotInstanceInterruptionBehavior : undefined;
            resourceInputs["spotInstancePoolsToUseCount"] = state ? state.spotInstancePoolsToUseCount : undefined;
            resourceInputs["spotTargetCapacity"] = state ? state.spotTargetCapacity : undefined;
            resourceInputs["terminateInstances"] = state ? state.terminateInstances : undefined;
            resourceInputs["terminateInstancesWithExpiration"] = state ? state.terminateInstancesWithExpiration : undefined;
            resourceInputs["totalTargetCapacity"] = state ? state.totalTargetCapacity : undefined;
            resourceInputs["validFrom"] = state ? state.validFrom : undefined;
            resourceInputs["validUntil"] = state ? state.validUntil : undefined;
        } else {
            const args = argsOrState as AutoProvisioningGroupArgs | undefined;
            if ((!args || args.launchTemplateConfigs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'launchTemplateConfigs'");
            }
            if ((!args || args.launchTemplateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'launchTemplateId'");
            }
            if ((!args || args.totalTargetCapacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'totalTargetCapacity'");
            }
            resourceInputs["autoProvisioningGroupName"] = args ? args.autoProvisioningGroupName : undefined;
            resourceInputs["autoProvisioningGroupType"] = args ? args.autoProvisioningGroupType : undefined;
            resourceInputs["defaultTargetCapacityType"] = args ? args.defaultTargetCapacityType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["excessCapacityTerminationPolicy"] = args ? args.excessCapacityTerminationPolicy : undefined;
            resourceInputs["launchTemplateConfigs"] = args ? args.launchTemplateConfigs : undefined;
            resourceInputs["launchTemplateId"] = args ? args.launchTemplateId : undefined;
            resourceInputs["launchTemplateVersion"] = args ? args.launchTemplateVersion : undefined;
            resourceInputs["maxSpotPrice"] = args ? args.maxSpotPrice : undefined;
            resourceInputs["payAsYouGoAllocationStrategy"] = args ? args.payAsYouGoAllocationStrategy : undefined;
            resourceInputs["payAsYouGoTargetCapacity"] = args ? args.payAsYouGoTargetCapacity : undefined;
            resourceInputs["spotAllocationStrategy"] = args ? args.spotAllocationStrategy : undefined;
            resourceInputs["spotInstanceInterruptionBehavior"] = args ? args.spotInstanceInterruptionBehavior : undefined;
            resourceInputs["spotInstancePoolsToUseCount"] = args ? args.spotInstancePoolsToUseCount : undefined;
            resourceInputs["spotTargetCapacity"] = args ? args.spotTargetCapacity : undefined;
            resourceInputs["terminateInstances"] = args ? args.terminateInstances : undefined;
            resourceInputs["terminateInstancesWithExpiration"] = args ? args.terminateInstancesWithExpiration : undefined;
            resourceInputs["totalTargetCapacity"] = args ? args.totalTargetCapacity : undefined;
            resourceInputs["validFrom"] = args ? args.validFrom : undefined;
            resourceInputs["validUntil"] = args ? args.validUntil : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AutoProvisioningGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AutoProvisioningGroup resources.
 */
export interface AutoProvisioningGroupState {
    /**
     * The name of the auto provisioning group to be created. It must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-)
     */
    autoProvisioningGroupName?: pulumi.Input<string>;
    /**
     * The type of the auto provisioning group. Valid values:`request` and `maintain`,Default value: `maintain`.
     */
    autoProvisioningGroupType?: pulumi.Input<string>;
    /**
     * The type of supplemental instances. When the total value of `PayAsYouGoTargetCapacity` and `SpotTargetCapacity` is smaller than the value of TotalTargetCapacity, the auto provisioning group will create instances of the specified type to meet the capacity requirements. Valid values:`PayAsYouGo`: Pay-as-you-go instances; `Spot`: Preemptible instances, Default value: `Spot`.
     */
    defaultTargetCapacityType?: pulumi.Input<string>;
    /**
     * The description of the auto provisioning group.
     */
    description?: pulumi.Input<string>;
    /**
     * The shutdown policy for excess preemptible instances followed when the capacity of the auto provisioning group exceeds the target capacity. Valid values: `no-termination` and `termination`,Default value: `no-termination`.
     */
    excessCapacityTerminationPolicy?: pulumi.Input<string>;
    /**
     * DataDisk mappings to attach to ecs instance. See Block config below for details.
     */
    launchTemplateConfigs?: pulumi.Input<pulumi.Input<inputs.ecs.AutoProvisioningGroupLaunchTemplateConfig>[]>;
    /**
     * The ID of the instance launch template associated with the auto provisioning group.
     */
    launchTemplateId?: pulumi.Input<string>;
    /**
     * The version of the instance launch template associated with the auto provisioning group.
     */
    launchTemplateVersion?: pulumi.Input<string>;
    /**
     * The global maximum price for preemptible instances in the auto provisioning group. If both the `MaxSpotPrice` and `LaunchTemplateConfig.N.MaxPrice` parameters are specified, the maximum price is the lower value of the two.
     */
    maxSpotPrice?: pulumi.Input<number>;
    /**
     * The scale-out policy for pay-as-you-go instances. Valid values: `lowest-price` and `prioritized`,Default value: `lowest-price`.
     */
    payAsYouGoAllocationStrategy?: pulumi.Input<string>;
    /**
     * The target capacity of pay-as-you-go instances in the auto provisioning group.
     */
    payAsYouGoTargetCapacity?: pulumi.Input<string>;
    /**
     * The scale-out policy for preemptible instances. Valid values:`lowest-price` and `diversified`,Default value: `lowest-price`.
     */
    spotAllocationStrategy?: pulumi.Input<string>;
    /**
     * The default behavior after preemptible instances are shut down. Valid values: `stop` and `terminate`,Default value: `stop`.
     */
    spotInstanceInterruptionBehavior?: pulumi.Input<string>;
    /**
     * This parameter takes effect when the `SpotAllocationStrategy` parameter is set to `lowest-price`. The auto provisioning group selects instance types of the lowest cost to create instances.
     */
    spotInstancePoolsToUseCount?: pulumi.Input<number>;
    /**
     * The target capacity of preemptible instances in the auto provisioning group.
     */
    spotTargetCapacity?: pulumi.Input<string>;
    /**
     * Specifies whether to release instances of the auto provisioning group. Valid values:`false` and `true`, default value: `false`.
     */
    terminateInstances?: pulumi.Input<boolean>;
    /**
     * The shutdown policy for preemptible instances when the auto provisioning group expires. Valid values: `false` and `true`, default value: `false`.
     */
    terminateInstancesWithExpiration?: pulumi.Input<boolean>;
    /**
     * The total target capacity of the auto provisioning group. The target capacity consists of the following three parts:PayAsYouGoTargetCapacity,SpotTargetCapacity and the supplemental capacity besides PayAsYouGoTargetCapacity and SpotTargetCapacity.
     */
    totalTargetCapacity?: pulumi.Input<string>;
    /**
     * The time when the auto provisioning group is started. The period of time between this point in time and the point in time specified by the `validUntil` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group is immediately started after creation.
     */
    validFrom?: pulumi.Input<string>;
    /**
     * The time when the auto provisioning group expires. The period of time between this point in time and the point in time specified by the `validFrom` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group never expires.
     */
    validUntil?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AutoProvisioningGroup resource.
 */
export interface AutoProvisioningGroupArgs {
    /**
     * The name of the auto provisioning group to be created. It must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-)
     */
    autoProvisioningGroupName?: pulumi.Input<string>;
    /**
     * The type of the auto provisioning group. Valid values:`request` and `maintain`,Default value: `maintain`.
     */
    autoProvisioningGroupType?: pulumi.Input<string>;
    /**
     * The type of supplemental instances. When the total value of `PayAsYouGoTargetCapacity` and `SpotTargetCapacity` is smaller than the value of TotalTargetCapacity, the auto provisioning group will create instances of the specified type to meet the capacity requirements. Valid values:`PayAsYouGo`: Pay-as-you-go instances; `Spot`: Preemptible instances, Default value: `Spot`.
     */
    defaultTargetCapacityType?: pulumi.Input<string>;
    /**
     * The description of the auto provisioning group.
     */
    description?: pulumi.Input<string>;
    /**
     * The shutdown policy for excess preemptible instances followed when the capacity of the auto provisioning group exceeds the target capacity. Valid values: `no-termination` and `termination`,Default value: `no-termination`.
     */
    excessCapacityTerminationPolicy?: pulumi.Input<string>;
    /**
     * DataDisk mappings to attach to ecs instance. See Block config below for details.
     */
    launchTemplateConfigs: pulumi.Input<pulumi.Input<inputs.ecs.AutoProvisioningGroupLaunchTemplateConfig>[]>;
    /**
     * The ID of the instance launch template associated with the auto provisioning group.
     */
    launchTemplateId: pulumi.Input<string>;
    /**
     * The version of the instance launch template associated with the auto provisioning group.
     */
    launchTemplateVersion?: pulumi.Input<string>;
    /**
     * The global maximum price for preemptible instances in the auto provisioning group. If both the `MaxSpotPrice` and `LaunchTemplateConfig.N.MaxPrice` parameters are specified, the maximum price is the lower value of the two.
     */
    maxSpotPrice?: pulumi.Input<number>;
    /**
     * The scale-out policy for pay-as-you-go instances. Valid values: `lowest-price` and `prioritized`,Default value: `lowest-price`.
     */
    payAsYouGoAllocationStrategy?: pulumi.Input<string>;
    /**
     * The target capacity of pay-as-you-go instances in the auto provisioning group.
     */
    payAsYouGoTargetCapacity?: pulumi.Input<string>;
    /**
     * The scale-out policy for preemptible instances. Valid values:`lowest-price` and `diversified`,Default value: `lowest-price`.
     */
    spotAllocationStrategy?: pulumi.Input<string>;
    /**
     * The default behavior after preemptible instances are shut down. Valid values: `stop` and `terminate`,Default value: `stop`.
     */
    spotInstanceInterruptionBehavior?: pulumi.Input<string>;
    /**
     * This parameter takes effect when the `SpotAllocationStrategy` parameter is set to `lowest-price`. The auto provisioning group selects instance types of the lowest cost to create instances.
     */
    spotInstancePoolsToUseCount?: pulumi.Input<number>;
    /**
     * The target capacity of preemptible instances in the auto provisioning group.
     */
    spotTargetCapacity?: pulumi.Input<string>;
    /**
     * Specifies whether to release instances of the auto provisioning group. Valid values:`false` and `true`, default value: `false`.
     */
    terminateInstances?: pulumi.Input<boolean>;
    /**
     * The shutdown policy for preemptible instances when the auto provisioning group expires. Valid values: `false` and `true`, default value: `false`.
     */
    terminateInstancesWithExpiration?: pulumi.Input<boolean>;
    /**
     * The total target capacity of the auto provisioning group. The target capacity consists of the following three parts:PayAsYouGoTargetCapacity,SpotTargetCapacity and the supplemental capacity besides PayAsYouGoTargetCapacity and SpotTargetCapacity.
     */
    totalTargetCapacity: pulumi.Input<string>;
    /**
     * The time when the auto provisioning group is started. The period of time between this point in time and the point in time specified by the `validUntil` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group is immediately started after creation.
     */
    validFrom?: pulumi.Input<string>;
    /**
     * The time when the auto provisioning group expires. The period of time between this point in time and the point in time specified by the `validFrom` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group never expires.
     */
    validUntil?: pulumi.Input<string>;
}
