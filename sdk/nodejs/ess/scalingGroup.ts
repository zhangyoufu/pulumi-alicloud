// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ScalingGroup extends pulumi.CustomResource {
    /**
     * Get an existing ScalingGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScalingGroupState, opts?: pulumi.CustomResourceOptions): ScalingGroup {
        return new ScalingGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ess/scalingGroup:ScalingGroup';

    /**
     * Returns true if the given object is an instance of ScalingGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScalingGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScalingGroup.__pulumiType;
    }

    /**
     * If an RDS instance is specified in the scaling group, the scaling group automatically attaches the Intranet IP addresses of its ECS instances to the RDS access whitelist.
     * - The specified RDS instance must be in running status.
     * - The specified RDS instance’s whitelist must have room for more IP addresses.
     */
    public readonly dbInstanceIds!: pulumi.Output<string[] | undefined>;
    /**
     * Default cool-down time (in seconds) of the scaling group. Value range: [0, 86400]. The default value is 300s.
     */
    public readonly defaultCooldown!: pulumi.Output<number | undefined>;
    /**
     * Expected number of ECS instances in the scaling group. Value range: [min_size, maxSize].
     */
    public readonly desiredCapacity!: pulumi.Output<number | undefined>;
    /**
     * If a Server Load Balancer instance is specified in the scaling group, the scaling group automatically attaches its ECS instances to the Server Load Balancer instance.
     * - The Server Load Balancer instance must be enabled.
     * - At least one listener must be configured for each Server Load Balancer and it HealthCheck must be on. Otherwise, creation will fail (it may be useful to add a `dependsOn` argument
     * targeting your `alicloud.slb.Listener` in order to make sure the listener with its HealthCheck configuration is ready before creating your scaling group).
     * - The Server Load Balancer instance attached with VPC-type ECS instances cannot be attached to the scaling group.
     * - The default weight of an ECS instance attached to the Server Load Balancer instance is 50.
     */
    public readonly loadbalancerIds!: pulumi.Output<string[] | undefined>;
    /**
     * Maximum number of ECS instances in the scaling group. Value range: [0, 1000].
     */
    public readonly maxSize!: pulumi.Output<number>;
    /**
     * Minimum number of ECS instances in the scaling group. Value range: [0, 1000].
     */
    public readonly minSize!: pulumi.Output<number>;
    /**
     * Multi-AZ scaling group ECS instance expansion and contraction strategy. PRIORITY, BALANCE or COST_OPTIMIZED(Available in 1.54.0+).
     */
    public readonly multiAzPolicy!: pulumi.Output<string | undefined>;
    /**
     * The minimum amount of the Auto Scaling group's capacity that must be fulfilled by On-Demand Instances. This base portion is provisioned first as your group scales.
     */
    public readonly onDemandBaseCapacity!: pulumi.Output<number>;
    /**
     * Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond OnDemandBaseCapacity.  
     */
    public readonly onDemandPercentageAboveBaseCapacity!: pulumi.Output<number>;
    /**
     * RemovalPolicy is used to select the ECS instances you want to remove from the scaling group when multiple candidates for removal exist. Optional values:
     * - OldestInstance: removes the first ECS instance attached to the scaling group.
     * - NewestInstance: removes the first ECS instance attached to the scaling group.
     * - OldestScalingConfiguration: removes the ECS instance with the oldest scaling configuration.
     * - Default values: OldestScalingConfiguration and OldestInstance. You can enter up to two removal policies.
     */
    public readonly removalPolicies!: pulumi.Output<string[]>;
    /**
     * Name shown for the scaling group, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain numbers, underscores `_`, hyphens `-`, and decimal points `.`. If this parameter is not specified, the default value is ScalingGroupId.
     */
    public readonly scalingGroupName!: pulumi.Output<string | undefined>;
    /**
     * The number of Spot pools to use to allocate your Spot capacity. The Spot pools is composed of instance types of lowest price.
     */
    public readonly spotInstancePools!: pulumi.Output<number>;
    /**
     * Whether to replace spot instances with newly created spot/onDemand instance when receive a spot recycling message.            
     */
    public readonly spotInstanceRemedy!: pulumi.Output<boolean>;
    /**
     * It has been deprecated from version 1.7.1 and new field 'vswitch_ids' replaces it.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;
    /**
     * List of virtual switch IDs in which the ecs instances to be launched.
     */
    public readonly vswitchIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a ScalingGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScalingGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScalingGroupArgs | ScalingGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ScalingGroupState | undefined;
            inputs["dbInstanceIds"] = state ? state.dbInstanceIds : undefined;
            inputs["defaultCooldown"] = state ? state.defaultCooldown : undefined;
            inputs["desiredCapacity"] = state ? state.desiredCapacity : undefined;
            inputs["loadbalancerIds"] = state ? state.loadbalancerIds : undefined;
            inputs["maxSize"] = state ? state.maxSize : undefined;
            inputs["minSize"] = state ? state.minSize : undefined;
            inputs["multiAzPolicy"] = state ? state.multiAzPolicy : undefined;
            inputs["onDemandBaseCapacity"] = state ? state.onDemandBaseCapacity : undefined;
            inputs["onDemandPercentageAboveBaseCapacity"] = state ? state.onDemandPercentageAboveBaseCapacity : undefined;
            inputs["removalPolicies"] = state ? state.removalPolicies : undefined;
            inputs["scalingGroupName"] = state ? state.scalingGroupName : undefined;
            inputs["spotInstancePools"] = state ? state.spotInstancePools : undefined;
            inputs["spotInstanceRemedy"] = state ? state.spotInstanceRemedy : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
            inputs["vswitchIds"] = state ? state.vswitchIds : undefined;
        } else {
            const args = argsOrState as ScalingGroupArgs | undefined;
            if (!args || args.maxSize === undefined) {
                throw new Error("Missing required property 'maxSize'");
            }
            if (!args || args.minSize === undefined) {
                throw new Error("Missing required property 'minSize'");
            }
            inputs["dbInstanceIds"] = args ? args.dbInstanceIds : undefined;
            inputs["defaultCooldown"] = args ? args.defaultCooldown : undefined;
            inputs["desiredCapacity"] = args ? args.desiredCapacity : undefined;
            inputs["loadbalancerIds"] = args ? args.loadbalancerIds : undefined;
            inputs["maxSize"] = args ? args.maxSize : undefined;
            inputs["minSize"] = args ? args.minSize : undefined;
            inputs["multiAzPolicy"] = args ? args.multiAzPolicy : undefined;
            inputs["onDemandBaseCapacity"] = args ? args.onDemandBaseCapacity : undefined;
            inputs["onDemandPercentageAboveBaseCapacity"] = args ? args.onDemandPercentageAboveBaseCapacity : undefined;
            inputs["removalPolicies"] = args ? args.removalPolicies : undefined;
            inputs["scalingGroupName"] = args ? args.scalingGroupName : undefined;
            inputs["spotInstancePools"] = args ? args.spotInstancePools : undefined;
            inputs["spotInstanceRemedy"] = args ? args.spotInstanceRemedy : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["vswitchIds"] = args ? args.vswitchIds : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ScalingGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScalingGroup resources.
 */
export interface ScalingGroupState {
    /**
     * If an RDS instance is specified in the scaling group, the scaling group automatically attaches the Intranet IP addresses of its ECS instances to the RDS access whitelist.
     * - The specified RDS instance must be in running status.
     * - The specified RDS instance’s whitelist must have room for more IP addresses.
     */
    readonly dbInstanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default cool-down time (in seconds) of the scaling group. Value range: [0, 86400]. The default value is 300s.
     */
    readonly defaultCooldown?: pulumi.Input<number>;
    /**
     * Expected number of ECS instances in the scaling group. Value range: [min_size, maxSize].
     */
    readonly desiredCapacity?: pulumi.Input<number>;
    /**
     * If a Server Load Balancer instance is specified in the scaling group, the scaling group automatically attaches its ECS instances to the Server Load Balancer instance.
     * - The Server Load Balancer instance must be enabled.
     * - At least one listener must be configured for each Server Load Balancer and it HealthCheck must be on. Otherwise, creation will fail (it may be useful to add a `dependsOn` argument
     * targeting your `alicloud.slb.Listener` in order to make sure the listener with its HealthCheck configuration is ready before creating your scaling group).
     * - The Server Load Balancer instance attached with VPC-type ECS instances cannot be attached to the scaling group.
     * - The default weight of an ECS instance attached to the Server Load Balancer instance is 50.
     */
    readonly loadbalancerIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Maximum number of ECS instances in the scaling group. Value range: [0, 1000].
     */
    readonly maxSize?: pulumi.Input<number>;
    /**
     * Minimum number of ECS instances in the scaling group. Value range: [0, 1000].
     */
    readonly minSize?: pulumi.Input<number>;
    /**
     * Multi-AZ scaling group ECS instance expansion and contraction strategy. PRIORITY, BALANCE or COST_OPTIMIZED(Available in 1.54.0+).
     */
    readonly multiAzPolicy?: pulumi.Input<string>;
    /**
     * The minimum amount of the Auto Scaling group's capacity that must be fulfilled by On-Demand Instances. This base portion is provisioned first as your group scales.
     */
    readonly onDemandBaseCapacity?: pulumi.Input<number>;
    /**
     * Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond OnDemandBaseCapacity.  
     */
    readonly onDemandPercentageAboveBaseCapacity?: pulumi.Input<number>;
    /**
     * RemovalPolicy is used to select the ECS instances you want to remove from the scaling group when multiple candidates for removal exist. Optional values:
     * - OldestInstance: removes the first ECS instance attached to the scaling group.
     * - NewestInstance: removes the first ECS instance attached to the scaling group.
     * - OldestScalingConfiguration: removes the ECS instance with the oldest scaling configuration.
     * - Default values: OldestScalingConfiguration and OldestInstance. You can enter up to two removal policies.
     */
    readonly removalPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name shown for the scaling group, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain numbers, underscores `_`, hyphens `-`, and decimal points `.`. If this parameter is not specified, the default value is ScalingGroupId.
     */
    readonly scalingGroupName?: pulumi.Input<string>;
    /**
     * The number of Spot pools to use to allocate your Spot capacity. The Spot pools is composed of instance types of lowest price.
     */
    readonly spotInstancePools?: pulumi.Input<number>;
    /**
     * Whether to replace spot instances with newly created spot/onDemand instance when receive a spot recycling message.            
     */
    readonly spotInstanceRemedy?: pulumi.Input<boolean>;
    /**
     * It has been deprecated from version 1.7.1 and new field 'vswitch_ids' replaces it.
     * 
     * @deprecated Field 'vswitch_id' has been deprecated from provider version 1.7.1, and new field 'vswitch_ids' can replace it.
     */
    readonly vswitchId?: pulumi.Input<string>;
    /**
     * List of virtual switch IDs in which the ecs instances to be launched.
     */
    readonly vswitchIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ScalingGroup resource.
 */
export interface ScalingGroupArgs {
    /**
     * If an RDS instance is specified in the scaling group, the scaling group automatically attaches the Intranet IP addresses of its ECS instances to the RDS access whitelist.
     * - The specified RDS instance must be in running status.
     * - The specified RDS instance’s whitelist must have room for more IP addresses.
     */
    readonly dbInstanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default cool-down time (in seconds) of the scaling group. Value range: [0, 86400]. The default value is 300s.
     */
    readonly defaultCooldown?: pulumi.Input<number>;
    /**
     * Expected number of ECS instances in the scaling group. Value range: [min_size, maxSize].
     */
    readonly desiredCapacity?: pulumi.Input<number>;
    /**
     * If a Server Load Balancer instance is specified in the scaling group, the scaling group automatically attaches its ECS instances to the Server Load Balancer instance.
     * - The Server Load Balancer instance must be enabled.
     * - At least one listener must be configured for each Server Load Balancer and it HealthCheck must be on. Otherwise, creation will fail (it may be useful to add a `dependsOn` argument
     * targeting your `alicloud.slb.Listener` in order to make sure the listener with its HealthCheck configuration is ready before creating your scaling group).
     * - The Server Load Balancer instance attached with VPC-type ECS instances cannot be attached to the scaling group.
     * - The default weight of an ECS instance attached to the Server Load Balancer instance is 50.
     */
    readonly loadbalancerIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Maximum number of ECS instances in the scaling group. Value range: [0, 1000].
     */
    readonly maxSize: pulumi.Input<number>;
    /**
     * Minimum number of ECS instances in the scaling group. Value range: [0, 1000].
     */
    readonly minSize: pulumi.Input<number>;
    /**
     * Multi-AZ scaling group ECS instance expansion and contraction strategy. PRIORITY, BALANCE or COST_OPTIMIZED(Available in 1.54.0+).
     */
    readonly multiAzPolicy?: pulumi.Input<string>;
    /**
     * The minimum amount of the Auto Scaling group's capacity that must be fulfilled by On-Demand Instances. This base portion is provisioned first as your group scales.
     */
    readonly onDemandBaseCapacity?: pulumi.Input<number>;
    /**
     * Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond OnDemandBaseCapacity.  
     */
    readonly onDemandPercentageAboveBaseCapacity?: pulumi.Input<number>;
    /**
     * RemovalPolicy is used to select the ECS instances you want to remove from the scaling group when multiple candidates for removal exist. Optional values:
     * - OldestInstance: removes the first ECS instance attached to the scaling group.
     * - NewestInstance: removes the first ECS instance attached to the scaling group.
     * - OldestScalingConfiguration: removes the ECS instance with the oldest scaling configuration.
     * - Default values: OldestScalingConfiguration and OldestInstance. You can enter up to two removal policies.
     */
    readonly removalPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name shown for the scaling group, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain numbers, underscores `_`, hyphens `-`, and decimal points `.`. If this parameter is not specified, the default value is ScalingGroupId.
     */
    readonly scalingGroupName?: pulumi.Input<string>;
    /**
     * The number of Spot pools to use to allocate your Spot capacity. The Spot pools is composed of instance types of lowest price.
     */
    readonly spotInstancePools?: pulumi.Input<number>;
    /**
     * Whether to replace spot instances with newly created spot/onDemand instance when receive a spot recycling message.            
     */
    readonly spotInstanceRemedy?: pulumi.Input<boolean>;
    /**
     * It has been deprecated from version 1.7.1 and new field 'vswitch_ids' replaces it.
     * 
     * @deprecated Field 'vswitch_id' has been deprecated from provider version 1.7.1, and new field 'vswitch_ids' can replace it.
     */
    readonly vswitchId?: pulumi.Input<string>;
    /**
     * List of virtual switch IDs in which the ecs instances to be launched.
     */
    readonly vswitchIds?: pulumi.Input<pulumi.Input<string>[]>;
}
