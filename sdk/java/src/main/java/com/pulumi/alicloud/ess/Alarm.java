// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ess.AlarmArgs;
import com.pulumi.alicloud.ess.inputs.AlarmState;
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
 * Provides a ESS alarm task resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecs.inputs.GetImagesArgs;
 * import com.pulumi.alicloud.ecs.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ess.ScalingGroup;
 * import com.pulumi.alicloud.ess.ScalingGroupArgs;
 * import com.pulumi.alicloud.ess.ScalingRule;
 * import com.pulumi.alicloud.ess.ScalingRuleArgs;
 * import com.pulumi.alicloud.ess.Alarm;
 * import com.pulumi.alicloud.ess.AlarmArgs;
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
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableDiskCategory(&#34;cloud_efficiency&#34;)
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         final var ecsImage = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .mostRecent(true)
 *             .nameRegex(&#34;^centos_6\\w{1,5}[64].*&#34;)
 *             .build());
 * 
 *         final var defaultInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .cpuCoreCount(1)
 *             .memorySize(2)
 *             .build());
 * 
 *         var fooNetwork = new Network(&#34;fooNetwork&#34;, NetworkArgs.builder()        
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var fooSwitch = new Switch(&#34;fooSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(&#34;tf-testAccEssAlarm_basic_foo&#34;)
 *             .vpcId(fooNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var bar = new Switch(&#34;bar&#34;, SwitchArgs.builder()        
 *             .vswitchName(&#34;tf-testAccEssAlarm_basic_bar&#34;)
 *             .vpcId(fooNetwork.id())
 *             .cidrBlock(&#34;172.16.1.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var fooScalingGroup = new ScalingGroup(&#34;fooScalingGroup&#34;, ScalingGroupArgs.builder()        
 *             .minSize(1)
 *             .maxSize(1)
 *             .scalingGroupName(&#34;tf-testAccEssAlarm_basic&#34;)
 *             .removalPolicies(            
 *                 &#34;OldestInstance&#34;,
 *                 &#34;NewestInstance&#34;)
 *             .vswitchIds(            
 *                 fooSwitch.id(),
 *                 bar.id())
 *             .build());
 * 
 *         var fooScalingRule = new ScalingRule(&#34;fooScalingRule&#34;, ScalingRuleArgs.builder()        
 *             .scalingRuleName(&#34;tf-testAccEssAlarm_basic&#34;)
 *             .scalingGroupId(fooScalingGroup.id())
 *             .adjustmentType(&#34;TotalCapacity&#34;)
 *             .adjustmentValue(2)
 *             .cooldown(60)
 *             .build());
 * 
 *         var fooAlarm = new Alarm(&#34;fooAlarm&#34;, AlarmArgs.builder()        
 *             .description(&#34;Acc alarm test&#34;)
 *             .alarmActions(fooScalingRule.ari())
 *             .scalingGroupId(fooScalingGroup.id())
 *             .metricType(&#34;system&#34;)
 *             .metricName(&#34;CpuUtilization&#34;)
 *             .period(300)
 *             .statistics(&#34;Average&#34;)
 *             .threshold(200.3)
 *             .comparisonOperator(&#34;&gt;=&#34;)
 *             .evaluationCount(2)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Module Support
 * 
 * You can use to the existing autoscaling-rule module
 * to create alarm task, different type rules and scheduled task one-click.
 * 
 * ## Block metricNames_and_dimensions
 * 
 * Supported metric names and dimensions :
 * 
 * | MetricName         | Dimensions                   |
 * | ------------------ | ---------------------------- |
 * | CpuUtilization     | user_id,scaling_group        |
 * | ClassicInternetRx  | user_id,scaling_group        |
 * | ClassicInternetTx  | user_id,scaling_group        |
 * | VpcInternetRx      | user_id,scaling_group        |
 * | VpcInternetTx      | user_id,scaling_group        |
 * | IntranetRx         | user_id,scaling_group        |
 * | IntranetTx         | user_id,scaling_group        |
 * | LoadAverage        | user_id,scaling_group        |
 * | MemoryUtilization  | user_id,scaling_group        |
 * | SystemDiskReadBps  | user_id,scaling_group        |
 * | SystemDiskWriteBps | user_id,scaling_group        |
 * | SystemDiskReadOps  | user_id,scaling_group        |
 * | SystemDiskWriteOps | user_id,scaling_group        |
 * | PackagesNetIn      | user_id,scaling_group,device |
 * | PackagesNetOut     | user_id,scaling_group,device |
 * | TcpConnection      | user_id,scaling_group,state  |
 * 
 * &gt; **NOTE:** Dimension `user_id` and `scaling_group` is automatically filled, which means you only need to care about dimension `device` and `state` when needed.
 * 
 * ## Import
 * 
 * Ess alarm can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ess/alarm:Alarm example asg-2ze500_045efffe-4d05
 * ```
 * 
 */
@ResourceType(type="alicloud:ess/alarm:Alarm")
public class Alarm extends com.pulumi.resources.CustomResource {
    /**
     * The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.
     * 
     */
    @Export(name="alarmActions", type=List.class, parameters={String.class})
    private Output<List<String>> alarmActions;

    /**
     * @return The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.
     * 
     */
    public Output<List<String>> alarmActions() {
        return this.alarmActions;
    }
    /**
     * Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.
     * 
     */
    @Export(name="cloudMonitorGroupId", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> cloudMonitorGroupId;

    /**
     * @return Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.
     * 
     */
    public Output<Optional<Integer>> cloudMonitorGroupId() {
        return Codegen.optional(this.cloudMonitorGroupId);
    }
    /**
     * The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: &gt;=, &lt;=, &gt;, &lt;. Defaults to &gt;=.
     * 
     */
    @Export(name="comparisonOperator", type=String.class, parameters={})
    private Output</* @Nullable */ String> comparisonOperator;

    /**
     * @return The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: &gt;=, &lt;=, &gt;, &lt;. Defaults to &gt;=.
     * 
     */
    public Output<Optional<String>> comparisonOperator() {
        return Codegen.optional(this.comparisonOperator);
    }
    /**
     * The description for the alarm.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description for the alarm.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The dimension map for the alarm&#39;s associated metric (documented below). For all metrics, you can not set the dimension key as &#34;scaling_group&#34; or &#34;userId&#34;, which is set by default, the second dimension for metric, such as &#34;device&#34; for &#34;PackagesNetIn&#34;, need to be set by users.
     * 
     */
    @Export(name="dimensions", type=Map.class, parameters={String.class, Object.class})
    private Output<Map<String,Object>> dimensions;

    /**
     * @return The dimension map for the alarm&#39;s associated metric (documented below). For all metrics, you can not set the dimension key as &#34;scaling_group&#34; or &#34;userId&#34;, which is set by default, the second dimension for metric, such as &#34;device&#34; for &#34;PackagesNetIn&#34;, need to be set by users.
     * 
     */
    public Output<Map<String,Object>> dimensions() {
        return this.dimensions;
    }
    /**
     * Whether to enable specific ess alarm. Default to true.
     * 
     */
    @Export(name="enable", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enable;

    /**
     * @return Whether to enable specific ess alarm. Default to true.
     * 
     */
    public Output<Optional<Boolean>> enable() {
        return Codegen.optional(this.enable);
    }
    /**
     * The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.
     * 
     */
    @Export(name="evaluationCount", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> evaluationCount;

    /**
     * @return The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.
     * 
     */
    public Output<Optional<Integer>> evaluationCount() {
        return Codegen.optional(this.evaluationCount);
    }
    /**
     * The name for the alarm&#39;s associated metric. See Block_metricNames_and_dimensions below for details.
     * 
     */
    @Export(name="metricName", type=String.class, parameters={})
    private Output<String> metricName;

    /**
     * @return The name for the alarm&#39;s associated metric. See Block_metricNames_and_dimensions below for details.
     * 
     */
    public Output<String> metricName() {
        return this.metricName;
    }
    /**
     * The type for the alarm&#39;s associated metric. Supported value: system, custom. &#34;system&#34; means the metric data is collected by Aliyun Cloud Monitor Service(CMS), &#34;custom&#34; means the metric data is upload to CMS by users. Defaults to system.
     * 
     */
    @Export(name="metricType", type=String.class, parameters={})
    private Output</* @Nullable */ String> metricType;

    /**
     * @return The type for the alarm&#39;s associated metric. Supported value: system, custom. &#34;system&#34; means the metric data is collected by Aliyun Cloud Monitor Service(CMS), &#34;custom&#34; means the metric data is upload to CMS by users. Defaults to system.
     * 
     */
    public Output<Optional<String>> metricType() {
        return Codegen.optional(this.metricType);
    }
    /**
     * The name for ess alarm.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name for ess alarm.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.
     * 
     */
    @Export(name="period", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> period;

    /**
     * @return The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * The scaling group associated with this alarm, the &#39;ForceNew&#39; attribute is available in 1.56.0+.
     * 
     */
    @Export(name="scalingGroupId", type=String.class, parameters={})
    private Output<String> scalingGroupId;

    /**
     * @return The scaling group associated with this alarm, the &#39;ForceNew&#39; attribute is available in 1.56.0+.
     * 
     */
    public Output<String> scalingGroupId() {
        return this.scalingGroupId;
    }
    /**
     * The state of specified alarm.
     * 
     */
    @Export(name="state", type=String.class, parameters={})
    private Output<String> state;

    /**
     * @return The state of specified alarm.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The statistic to apply to the alarm&#39;s associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.
     * 
     */
    @Export(name="statistics", type=String.class, parameters={})
    private Output</* @Nullable */ String> statistics;

    /**
     * @return The statistic to apply to the alarm&#39;s associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.
     * 
     */
    public Output<Optional<String>> statistics() {
        return Codegen.optional(this.statistics);
    }
    /**
     * The value against which the specified statistics is compared.
     * 
     */
    @Export(name="threshold", type=String.class, parameters={})
    private Output<String> threshold;

    /**
     * @return The value against which the specified statistics is compared.
     * 
     */
    public Output<String> threshold() {
        return this.threshold;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Alarm(String name) {
        this(name, AlarmArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Alarm(String name, AlarmArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Alarm(String name, AlarmArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ess/alarm:Alarm", name, args == null ? AlarmArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Alarm(String name, Output<String> id, @Nullable AlarmState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ess/alarm:Alarm", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Alarm get(String name, Output<String> id, @Nullable AlarmState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Alarm(name, id, state, options);
    }
}
