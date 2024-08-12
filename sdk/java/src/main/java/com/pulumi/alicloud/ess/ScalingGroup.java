// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ess.ScalingGroupArgs;
import com.pulumi.alicloud.ess.inputs.ScalingGroupState;
import com.pulumi.alicloud.ess.outputs.ScalingGroupAlbServerGroup;
import com.pulumi.alicloud.ess.outputs.ScalingGroupLaunchTemplateOverride;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ESS scaling group resource which is a collection of ECS instances with the same application scenarios.
 * 
 * It defines the maximum and minimum numbers of ECS instances in the group, and their associated Server Load Balancer instances, RDS instances, and other attributes.
 * 
 * &gt; **NOTE:** You can launch an ESS scaling group for a VPC network via specifying parameter `vswitch_ids`.
 * 
 * For information about ess scaling rule, see [CreateScalingGroup](https://www.alibabacloud.com/help/en/auto-scaling/latest/createscalinggroup).
 * 
 * &gt; **NOTE:** Available since v1.39.0.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecs.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.ecs.inputs.GetImagesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroupRule;
 * import com.pulumi.alicloud.ecs.SecurityGroupRuleArgs;
 * import com.pulumi.alicloud.ess.ScalingGroup;
 * import com.pulumi.alicloud.ess.ScalingGroupArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var name = config.get("name").orElse("terraform-example");
 *         var defaultInteger = new Integer("defaultInteger", IntegerArgs.builder()
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         final var myName = String.format("%s-%s", name,defaultInteger.result());
 * 
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableDiskCategory("cloud_efficiency")
 *             .availableResourceCreation("VSwitch")
 *             .build());
 * 
 *         final var defaultGetInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(default_.zones()[0].id())
 *             .cpuCoreCount(2)
 *             .memorySize(4)
 *             .build());
 * 
 *         final var defaultGetImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex("^ubuntu_18.*64")
 *             .mostRecent(true)
 *             .owners("system")
 *             .build());
 * 
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()
 *             .vpcName(myName)
 *             .cidrBlock("172.16.0.0/16")
 *             .build());
 * 
 *         var defaultSwitch = new Switch("defaultSwitch", SwitchArgs.builder()
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock("172.16.0.0/24")
 *             .zoneId(default_.zones()[0].id())
 *             .vswitchName(myName)
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup("defaultSecurityGroup", SecurityGroupArgs.builder()
 *             .name(myName)
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultSecurityGroupRule = new SecurityGroupRule("defaultSecurityGroupRule", SecurityGroupRuleArgs.builder()
 *             .type("ingress")
 *             .ipProtocol("tcp")
 *             .nicType("intranet")
 *             .policy("accept")
 *             .portRange("22/22")
 *             .priority(1)
 *             .securityGroupId(defaultSecurityGroup.id())
 *             .cidrIp("172.16.0.0/24")
 *             .build());
 * 
 *         var default2 = new Switch("default2", SwitchArgs.builder()
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock("172.16.1.0/24")
 *             .zoneId(default_.zones()[0].id())
 *             .vswitchName(String.format("%s-bar", name))
 *             .build());
 * 
 *         var defaultScalingGroup = new ScalingGroup("defaultScalingGroup", ScalingGroupArgs.builder()
 *             .minSize(1)
 *             .maxSize(1)
 *             .scalingGroupName(myName)
 *             .defaultCooldown(20)
 *             .vswitchIds(            
 *                 defaultSwitch.id(),
 *                 default2.id())
 *             .removalPolicies(            
 *                 "OldestInstance",
 *                 "NewestInstance")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Module Support
 * 
 * You can use to the existing autoscaling module
 * to create a scaling group, configuration and lifecycle hook one-click.
 * 
 * ## Import
 * 
 * ESS scaling group can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ess/scalingGroup:ScalingGroup example asg-abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:ess/scalingGroup:ScalingGroup")
public class ScalingGroup extends com.pulumi.resources.CustomResource {
    /**
     * If a Serve ALB instance is specified in the scaling group, the scaling group automatically attaches its ECS instances to the Server ALB instance.  See `alb_server_group` below for details.
     * 
     */
    @Export(name="albServerGroups", refs={List.class,ScalingGroupAlbServerGroup.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ScalingGroupAlbServerGroup>> albServerGroups;

    /**
     * @return If a Serve ALB instance is specified in the scaling group, the scaling group automatically attaches its ECS instances to the Server ALB instance.  See `alb_server_group` below for details.
     * 
     */
    public Output<Optional<List<ScalingGroupAlbServerGroup>>> albServerGroups() {
        return Codegen.optional(this.albServerGroups);
    }
    /**
     * The allocation policy of instances. Auto Scaling selects instance types based on the allocation policy to create instances. The policy can be applied to pay-as-you-go instances and preemptible instances. This parameter takes effect only if you set MultiAZPolicy to COMPOSABLE.
     * 
     */
    @Export(name="allocationStrategy", refs={String.class}, tree="[0]")
    private Output<String> allocationStrategy;

    /**
     * @return The allocation policy of instances. Auto Scaling selects instance types based on the allocation policy to create instances. The policy can be applied to pay-as-you-go instances and preemptible instances. This parameter takes effect only if you set MultiAZPolicy to COMPOSABLE.
     * 
     */
    public Output<String> allocationStrategy() {
        return this.allocationStrategy;
    }
    /**
     * Specifies whether to evenly distribute instances in the scaling group across multiple zones. This parameter takes effect only if you set MultiAZPolicy to COMPOSABLE.
     * 
     */
    @Export(name="azBalance", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> azBalance;

    /**
     * @return Specifies whether to evenly distribute instances in the scaling group across multiple zones. This parameter takes effect only if you set MultiAZPolicy to COMPOSABLE.
     * 
     */
    public Output<Optional<Boolean>> azBalance() {
        return Codegen.optional(this.azBalance);
    }
    /**
     * If an RDS instance is specified in the scaling group, the scaling group automatically attaches the Intranet IP addresses of its ECS instances to the RDS access whitelist.
     * - The specified RDS instance must be in running status.
     * - The specified RDS instance’s whitelist must have room for more IP addresses.
     * 
     */
    @Export(name="dbInstanceIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> dbInstanceIds;

    /**
     * @return If an RDS instance is specified in the scaling group, the scaling group automatically attaches the Intranet IP addresses of its ECS instances to the RDS access whitelist.
     * - The specified RDS instance must be in running status.
     * - The specified RDS instance’s whitelist must have room for more IP addresses.
     * 
     */
    public Output<Optional<List<String>>> dbInstanceIds() {
        return Codegen.optional(this.dbInstanceIds);
    }
    /**
     * Default cool-down time (in seconds) of the scaling group. Value range: [0, 86400]. The default value is 300s.
     * 
     */
    @Export(name="defaultCooldown", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> defaultCooldown;

    /**
     * @return Default cool-down time (in seconds) of the scaling group. Value range: [0, 86400]. The default value is 300s.
     * 
     */
    public Output<Optional<Integer>> defaultCooldown() {
        return Codegen.optional(this.defaultCooldown);
    }
    /**
     * Expected number of ECS instances in the scaling group. Value range: [min_size, max_size].
     * 
     */
    @Export(name="desiredCapacity", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> desiredCapacity;

    /**
     * @return Expected number of ECS instances in the scaling group. Value range: [min_size, max_size].
     * 
     */
    public Output<Optional<Integer>> desiredCapacity() {
        return Codegen.optional(this.desiredCapacity);
    }
    /**
     * Specifies whether the scaling group deletion protection is enabled. `true` or `false`, Default value: `false`.
     * 
     */
    @Export(name="groupDeletionProtection", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> groupDeletionProtection;

    /**
     * @return Specifies whether the scaling group deletion protection is enabled. `true` or `false`, Default value: `false`.
     * 
     */
    public Output<Optional<Boolean>> groupDeletionProtection() {
        return Codegen.optional(this.groupDeletionProtection);
    }
    /**
     * Resource type within scaling group. Optional values: ECS, ECI. Default to ECS.
     * 
     */
    @Export(name="groupType", refs={String.class}, tree="[0]")
    private Output<String> groupType;

    /**
     * @return Resource type within scaling group. Optional values: ECS, ECI. Default to ECS.
     * 
     */
    public Output<String> groupType() {
        return this.groupType;
    }
    /**
     * Resource type within scaling group. Optional values: ECS, NONE, LOAD_BALANCER. Default to ECS.
     * 
     */
    @Export(name="healthCheckType", refs={String.class}, tree="[0]")
    private Output<String> healthCheckType;

    /**
     * @return Resource type within scaling group. Optional values: ECS, NONE, LOAD_BALANCER. Default to ECS.
     * 
     */
    public Output<String> healthCheckType() {
        return this.healthCheckType;
    }
    /**
     * Instance launch template ID, scaling group obtains launch configuration from instance launch template, see [Launch Template](https://www.alibabacloud.com/help/doc-detail/73916.html). Creating scaling group from launch template enable group automatically.
     * 
     */
    @Export(name="launchTemplateId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> launchTemplateId;

    /**
     * @return Instance launch template ID, scaling group obtains launch configuration from instance launch template, see [Launch Template](https://www.alibabacloud.com/help/doc-detail/73916.html). Creating scaling group from launch template enable group automatically.
     * 
     */
    public Output<Optional<String>> launchTemplateId() {
        return Codegen.optional(this.launchTemplateId);
    }
    /**
     * The details of the instance types that are specified by using the Extend Instance Type of Launch Template feature.  See `launch_template_override` below for details.
     * 
     */
    @Export(name="launchTemplateOverrides", refs={List.class,ScalingGroupLaunchTemplateOverride.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ScalingGroupLaunchTemplateOverride>> launchTemplateOverrides;

    /**
     * @return The details of the instance types that are specified by using the Extend Instance Type of Launch Template feature.  See `launch_template_override` below for details.
     * 
     */
    public Output<Optional<List<ScalingGroupLaunchTemplateOverride>>> launchTemplateOverrides() {
        return Codegen.optional(this.launchTemplateOverrides);
    }
    /**
     * The version number of the launch template. Valid values are the version number, `Latest`, or `Default`, Default value: `Default`.
     * 
     */
    @Export(name="launchTemplateVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> launchTemplateVersion;

    /**
     * @return The version number of the launch template. Valid values are the version number, `Latest`, or `Default`, Default value: `Default`.
     * 
     */
    public Output<Optional<String>> launchTemplateVersion() {
        return Codegen.optional(this.launchTemplateVersion);
    }
    /**
     * If a Server Load Balancer instance is specified in the scaling group, the scaling group automatically attaches its ECS instances to the Server Load Balancer instance.
     * - The Server Load Balancer instance must be enabled.
     * - At least one listener must be configured for each Server Load Balancer and it HealthCheck must be on. Otherwise, creation will fail (it may be useful to add a `depends_on` argument
     *   targeting your `alicloud.slb.Listener` in order to make sure the listener with its HealthCheck configuration is ready before creating your scaling group).
     * - The Server Load Balancer instance attached with VPC-type ECS instances cannot be attached to the scaling group.
     * - The default weight of an ECS instance attached to the Server Load Balancer instance is 50.
     * 
     */
    @Export(name="loadbalancerIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> loadbalancerIds;

    /**
     * @return If a Server Load Balancer instance is specified in the scaling group, the scaling group automatically attaches its ECS instances to the Server Load Balancer instance.
     * - The Server Load Balancer instance must be enabled.
     * - At least one listener must be configured for each Server Load Balancer and it HealthCheck must be on. Otherwise, creation will fail (it may be useful to add a `depends_on` argument
     *   targeting your `alicloud.slb.Listener` in order to make sure the listener with its HealthCheck configuration is ready before creating your scaling group).
     * - The Server Load Balancer instance attached with VPC-type ECS instances cannot be attached to the scaling group.
     * - The default weight of an ECS instance attached to the Server Load Balancer instance is 50.
     * 
     */
    public Output<Optional<List<String>>> loadbalancerIds() {
        return Codegen.optional(this.loadbalancerIds);
    }
    /**
     * The maximum life span of an instance in the scaling group. Unit: seconds.
     * 
     */
    @Export(name="maxInstanceLifetime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxInstanceLifetime;

    /**
     * @return The maximum life span of an instance in the scaling group. Unit: seconds.
     * 
     */
    public Output<Optional<Integer>> maxInstanceLifetime() {
        return Codegen.optional(this.maxInstanceLifetime);
    }
    /**
     * Maximum number of ECS instances in the scaling group. Value range: [0, 2000].
     * **NOTE:** From version 1.204.1, `max_size` can be set to `2000`.
     * 
     */
    @Export(name="maxSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxSize;

    /**
     * @return Maximum number of ECS instances in the scaling group. Value range: [0, 2000].
     * **NOTE:** From version 1.204.1, `max_size` can be set to `2000`.
     * 
     */
    public Output<Integer> maxSize() {
        return this.maxSize;
    }
    /**
     * Minimum number of ECS instances in the scaling group. Value range: [0, 2000].
     * **NOTE:** From version 1.204.1, `min_size` can be set to `2000`.
     * 
     */
    @Export(name="minSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> minSize;

    /**
     * @return Minimum number of ECS instances in the scaling group. Value range: [0, 2000].
     * **NOTE:** From version 1.204.1, `min_size` can be set to `2000`.
     * 
     */
    public Output<Integer> minSize() {
        return this.minSize;
    }
    /**
     * Multi-AZ scaling group ECS instance expansion and contraction strategy. PRIORITY, COMPOSABLE, BALANCE or COST_OPTIMIZED(Available since v1.54.0).
     * 
     */
    @Export(name="multiAzPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> multiAzPolicy;

    /**
     * @return Multi-AZ scaling group ECS instance expansion and contraction strategy. PRIORITY, COMPOSABLE, BALANCE or COST_OPTIMIZED(Available since v1.54.0).
     * 
     */
    public Output<Optional<String>> multiAzPolicy() {
        return Codegen.optional(this.multiAzPolicy);
    }
    /**
     * The minimum amount of the Auto Scaling group&#39;s capacity that must be fulfilled by On-Demand Instances. This base portion is provisioned first as your group scales.
     * 
     */
    @Export(name="onDemandBaseCapacity", refs={Integer.class}, tree="[0]")
    private Output<Integer> onDemandBaseCapacity;

    /**
     * @return The minimum amount of the Auto Scaling group&#39;s capacity that must be fulfilled by On-Demand Instances. This base portion is provisioned first as your group scales.
     * 
     */
    public Output<Integer> onDemandBaseCapacity() {
        return this.onDemandBaseCapacity;
    }
    /**
     * Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond OnDemandBaseCapacity.
     * 
     */
    @Export(name="onDemandPercentageAboveBaseCapacity", refs={Integer.class}, tree="[0]")
    private Output<Integer> onDemandPercentageAboveBaseCapacity;

    /**
     * @return Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond OnDemandBaseCapacity.
     * 
     */
    public Output<Integer> onDemandPercentageAboveBaseCapacity() {
        return this.onDemandPercentageAboveBaseCapacity;
    }
    /**
     * Set or unset instances within group into protected status.
     * 
     */
    @Export(name="protectedInstances", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> protectedInstances;

    /**
     * @return Set or unset instances within group into protected status.
     * 
     */
    public Output<Optional<List<String>>> protectedInstances() {
        return Codegen.optional(this.protectedInstances);
    }
    /**
     * RemovalPolicy is used to select the ECS instances you want to remove from the scaling group when multiple candidates for removal exist. Optional values:
     * - OldestInstance: removes the ECS instance that is added to the scaling group at the earliest point in time.
     * - NewestInstance: removes the ECS instance that is added to the scaling group at the latest point in time.
     * - OldestScalingConfiguration: removes the ECS instance that is created based on the earliest scaling configuration.
     * - Default values: Default value of RemovalPolicy.1: OldestScalingConfiguration. Default value of RemovalPolicy.2: OldestInstance.
     * 
     */
    @Export(name="removalPolicies", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> removalPolicies;

    /**
     * @return RemovalPolicy is used to select the ECS instances you want to remove from the scaling group when multiple candidates for removal exist. Optional values:
     * - OldestInstance: removes the ECS instance that is added to the scaling group at the earliest point in time.
     * - NewestInstance: removes the ECS instance that is added to the scaling group at the latest point in time.
     * - OldestScalingConfiguration: removes the ECS instance that is created based on the earliest scaling configuration.
     * - Default values: Default value of RemovalPolicy.1: OldestScalingConfiguration. Default value of RemovalPolicy.2: OldestInstance.
     * 
     */
    public Output<List<String>> removalPolicies() {
        return this.removalPolicies;
    }
    /**
     * The ID of the resource group to which you want to add the scaling group.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group to which you want to add the scaling group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * Name shown for the scaling group, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain numbers, underscores `_`, hyphens `-`, and decimal points `.`. If this parameter is not specified, the default value is ScalingGroupId.
     * 
     */
    @Export(name="scalingGroupName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> scalingGroupName;

    /**
     * @return Name shown for the scaling group, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain numbers, underscores `_`, hyphens `-`, and decimal points `.`. If this parameter is not specified, the default value is ScalingGroupId.
     * 
     */
    public Output<Optional<String>> scalingGroupName() {
        return Codegen.optional(this.scalingGroupName);
    }
    /**
     * The reclaim mode of the scaling group. Optional values: recycle, release, forceRecycle, forceRelease.
     * 
     */
    @Export(name="scalingPolicy", refs={String.class}, tree="[0]")
    private Output<String> scalingPolicy;

    /**
     * @return The reclaim mode of the scaling group. Optional values: recycle, release, forceRecycle, forceRelease.
     * 
     */
    public Output<String> scalingPolicy() {
        return this.scalingPolicy;
    }
    /**
     * The allocation policy of preemptible instances. You can use this parameter to individually specify the allocation policy for preemptible instances. This parameter takes effect only if you set MultiAZPolicy to COMPOSABLE.
     * 
     */
    @Export(name="spotAllocationStrategy", refs={String.class}, tree="[0]")
    private Output<String> spotAllocationStrategy;

    /**
     * @return The allocation policy of preemptible instances. You can use this parameter to individually specify the allocation policy for preemptible instances. This parameter takes effect only if you set MultiAZPolicy to COMPOSABLE.
     * 
     */
    public Output<String> spotAllocationStrategy() {
        return this.spotAllocationStrategy;
    }
    /**
     * The number of Spot pools to use to allocate your Spot capacity. The Spot pools is composed of instance types of lowest price.
     * 
     */
    @Export(name="spotInstancePools", refs={Integer.class}, tree="[0]")
    private Output<Integer> spotInstancePools;

    /**
     * @return The number of Spot pools to use to allocate your Spot capacity. The Spot pools is composed of instance types of lowest price.
     * 
     */
    public Output<Integer> spotInstancePools() {
        return this.spotInstancePools;
    }
    /**
     * Whether to replace spot instances with newly created spot/onDemand instance when receive a spot recycling message.
     * 
     */
    @Export(name="spotInstanceRemedy", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> spotInstanceRemedy;

    /**
     * @return Whether to replace spot instances with newly created spot/onDemand instance when receive a spot recycling message.
     * 
     */
    public Output<Boolean> spotInstanceRemedy() {
        return this.spotInstanceRemedy;
    }
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * It has been deprecated from version 1.7.1 and new field &#39;vswitch_ids&#39; replaces it.
     * 
     * @deprecated
     * Field &#39;vswitch_id&#39; has been deprecated from provider version 1.7.1, and new field &#39;vswitch_ids&#39; can replace it.
     * 
     */
    @Deprecated /* Field 'vswitch_id' has been deprecated from provider version 1.7.1, and new field 'vswitch_ids' can replace it. */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vswitchId;

    /**
     * @return It has been deprecated from version 1.7.1 and new field &#39;vswitch_ids&#39; replaces it.
     * 
     */
    public Output<Optional<String>> vswitchId() {
        return Codegen.optional(this.vswitchId);
    }
    /**
     * List of virtual switch IDs in which the ecs instances to be launched.
     * 
     */
    @Export(name="vswitchIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> vswitchIds;

    /**
     * @return List of virtual switch IDs in which the ecs instances to be launched.
     * 
     */
    public Output<Optional<List<String>>> vswitchIds() {
        return Codegen.optional(this.vswitchIds);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ScalingGroup(java.lang.String name) {
        this(name, ScalingGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ScalingGroup(java.lang.String name, ScalingGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ScalingGroup(java.lang.String name, ScalingGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ess/scalingGroup:ScalingGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ScalingGroup(java.lang.String name, Output<java.lang.String> id, @Nullable ScalingGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ess/scalingGroup:ScalingGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static ScalingGroupArgs makeArgs(ScalingGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ScalingGroupArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ScalingGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable ScalingGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ScalingGroup(name, id, state, options);
    }
}
