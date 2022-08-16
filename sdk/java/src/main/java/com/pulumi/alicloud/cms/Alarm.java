// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cms.AlarmArgs;
import com.pulumi.alicloud.cms.inputs.AlarmState;
import com.pulumi.alicloud.cms.outputs.AlarmEscalationsCritical;
import com.pulumi.alicloud.cms.outputs.AlarmEscalationsInfo;
import com.pulumi.alicloud.cms.outputs.AlarmEscalationsWarn;
import com.pulumi.alicloud.cms.outputs.AlarmPrometheus;
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
 * This resource provides a alarm rule resource and it can be used to monitor several cloud services according different metrics.
 * Details for [alarm rule](https://www.alibabacloud.com/help/doc-detail/28608.htm).
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.cms.Alarm;
 * import com.pulumi.alicloud.cms.AlarmArgs;
 * import com.pulumi.alicloud.cms.inputs.AlarmEscalationsCriticalArgs;
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
 *         var basic = new Alarm(&#34;basic&#34;, AlarmArgs.builder()        
 *             .contactGroups(&#34;test-group&#34;)
 *             .effectiveInterval(&#34;0:00-2:00&#34;)
 *             .escalationsCritical(AlarmEscalationsCriticalArgs.builder()
 *                 .comparisonOperator(&#34;&lt;=&#34;)
 *                 .statistics(&#34;Average&#34;)
 *                 .threshold(35)
 *                 .times(2)
 *                 .build())
 *             .metricDimensions(&#34;[{\&#34;instanceId\&#34;:\&#34;i-bp1247jeep0y53nu3bnk\&#34;,\&#34;device\&#34;:\&#34;/dev/vda1\&#34;},{\&#34;instanceId\&#34;:\&#34;i-bp11gdcik8z6dl5jm84p\&#34;,\&#34;device\&#34;:\&#34;/dev/vdb1\&#34;}]&#34;)
 *             .period(900)
 *             .project(&#34;acs_ecs_dashboard&#34;)
 *             .webhook(String.format(&#34;https://%s.eu-central-1.fc.aliyuncs.com/2016-08-15/proxy/Terraform/AlarmEndpointMock/&#34;, data.alicloud_account().current().id()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Alarm rule can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:cms/alarm:Alarm alarm abc12345
 * ```
 * 
 */
@ResourceType(type="alicloud:cms/alarm:Alarm")
public class Alarm extends com.pulumi.resources.CustomResource {
    /**
     * List contact groups of the alarm rule, which must have been created on the console.
     * 
     */
    @Export(name="contactGroups", type=List.class, parameters={String.class})
    private Output<List<String>> contactGroups;

    /**
     * @return List contact groups of the alarm rule, which must have been created on the console.
     * 
     */
    public Output<List<String>> contactGroups() {
        return this.contactGroups;
    }
    /**
     * Field `dimensions` has been deprecated from version 1.95.0. Use `metric_dimensions` instead.
     * 
     * @deprecated
     * Field &#39;dimensions&#39; has been deprecated from version 1.173.0. Use &#39;metric_dimensions&#39; instead.
     * 
     */
    @Deprecated /* Field 'dimensions' has been deprecated from version 1.173.0. Use 'metric_dimensions' instead. */
    @Export(name="dimensions", type=Map.class, parameters={String.class, Object.class})
    private Output<Map<String,Object>> dimensions;

    /**
     * @return Field `dimensions` has been deprecated from version 1.95.0. Use `metric_dimensions` instead.
     * 
     */
    public Output<Map<String,Object>> dimensions() {
        return this.dimensions;
    }
    /**
     * The interval of effecting alarm rule. It format as &#34;hh:mm-hh:mm&#34;, like &#34;0:00-4:00&#34;. Default to &#34;00:00-23:59&#34;.
     * 
     */
    @Export(name="effectiveInterval", type=String.class, parameters={})
    private Output</* @Nullable */ String> effectiveInterval;

    /**
     * @return The interval of effecting alarm rule. It format as &#34;hh:mm-hh:mm&#34;, like &#34;0:00-4:00&#34;. Default to &#34;00:00-23:59&#34;.
     * 
     */
    public Output<Optional<String>> effectiveInterval() {
        return Codegen.optional(this.effectiveInterval);
    }
    /**
     * Whether to enable alarm rule. Default to true.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Whether to enable alarm rule. Default to true.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * It has been deprecated from provider version 1.50.0 and &#39;effective_interval&#39; instead.
     * 
     * @deprecated
     * Field &#39;end_time&#39; has been deprecated from provider version 1.50.0. New field &#39;effective_interval&#39; instead.
     * 
     */
    @Deprecated /* Field 'end_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead. */
    @Export(name="endTime", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> endTime;

    /**
     * @return It has been deprecated from provider version 1.50.0 and &#39;effective_interval&#39; instead.
     * 
     */
    public Output<Optional<Integer>> endTime() {
        return Codegen.optional(this.endTime);
    }
    /**
     * A configuration of critical alarm (documented below).
     * 
     */
    @Export(name="escalationsCritical", type=AlarmEscalationsCritical.class, parameters={})
    private Output</* @Nullable */ AlarmEscalationsCritical> escalationsCritical;

    /**
     * @return A configuration of critical alarm (documented below).
     * 
     */
    public Output<Optional<AlarmEscalationsCritical>> escalationsCritical() {
        return Codegen.optional(this.escalationsCritical);
    }
    /**
     * A configuration of critical info (documented below).
     * 
     */
    @Export(name="escalationsInfo", type=AlarmEscalationsInfo.class, parameters={})
    private Output</* @Nullable */ AlarmEscalationsInfo> escalationsInfo;

    /**
     * @return A configuration of critical info (documented below).
     * 
     */
    public Output<Optional<AlarmEscalationsInfo>> escalationsInfo() {
        return Codegen.optional(this.escalationsInfo);
    }
    /**
     * A configuration of critical warn (documented below).
     * 
     */
    @Export(name="escalationsWarn", type=AlarmEscalationsWarn.class, parameters={})
    private Output</* @Nullable */ AlarmEscalationsWarn> escalationsWarn;

    /**
     * @return A configuration of critical warn (documented below).
     * 
     */
    public Output<Optional<AlarmEscalationsWarn>> escalationsWarn() {
        return Codegen.optional(this.escalationsWarn);
    }
    /**
     * Name of the monitoring metrics corresponding to a project, such as &#34;CPUUtilization&#34; and &#34;networkin_rate&#34;. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     * 
     */
    @Export(name="metric", type=String.class, parameters={})
    private Output<String> metric;

    /**
     * @return Name of the monitoring metrics corresponding to a project, such as &#34;CPUUtilization&#34; and &#34;networkin_rate&#34;. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     * 
     */
    public Output<String> metric() {
        return this.metric;
    }
    /**
     * Map of the resources associated with the alarm rule, such as &#34;instanceId&#34;, &#34;device&#34; and &#34;port&#34;. Each key&#39;s value is a string, and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     * 
     */
    @Export(name="metricDimensions", type=String.class, parameters={})
    private Output<String> metricDimensions;

    /**
     * @return Map of the resources associated with the alarm rule, such as &#34;instanceId&#34;, &#34;device&#34; and &#34;port&#34;. Each key&#39;s value is a string, and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     * 
     */
    public Output<String> metricDimensions() {
        return this.metricDimensions;
    }
    /**
     * The alarm rule name.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The alarm rule name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * It has been deprecated from provider version 1.94.0 and &#39;escalations_critical.comparison_operator&#39; instead.
     * 
     * @deprecated
     * Field &#39;operator&#39; has been deprecated from provider version 1.94.0. New field &#39;escalations_critical.comparison_operator&#39; instead.
     * 
     */
    @Deprecated /* Field 'operator' has been deprecated from provider version 1.94.0. New field 'escalations_critical.comparison_operator' instead. */
    @Export(name="operator", type=String.class, parameters={})
    private Output<String> operator;

    /**
     * @return It has been deprecated from provider version 1.94.0 and &#39;escalations_critical.comparison_operator&#39; instead.
     * 
     */
    public Output<String> operator() {
        return this.operator;
    }
    /**
     * Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
     * 
     */
    @Export(name="period", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> period;

    /**
     * @return Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * Monitor project name, such as &#34;acs_ecs_dashboard&#34; and &#34;acs_rds_dashboard&#34;. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     * **NOTE:** The `dimensions` and `metric_dimensions` must be empty when `project` is `acs_prometheus`, otherwise, one of them must be set.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return Monitor project name, such as &#34;acs_ecs_dashboard&#34; and &#34;acs_rds_dashboard&#34;. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     * **NOTE:** The `dimensions` and `metric_dimensions` must be empty when `project` is `acs_prometheus`, otherwise, one of them must be set.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The Prometheus alert rule. See the following `Block prometheus`. **Note:** This parameter is required only when you create a Prometheus alert rule for Hybrid Cloud Monitoring.
     * 
     */
    @Export(name="prometheuses", type=List.class, parameters={AlarmPrometheus.class})
    private Output</* @Nullable */ List<AlarmPrometheus>> prometheuses;

    /**
     * @return The Prometheus alert rule. See the following `Block prometheus`. **Note:** This parameter is required only when you create a Prometheus alert rule for Hybrid Cloud Monitoring.
     * 
     */
    public Output<Optional<List<AlarmPrometheus>>> prometheuses() {
        return Codegen.optional(this.prometheuses);
    }
    /**
     * Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
     * 
     */
    @Export(name="silenceTime", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> silenceTime;

    /**
     * @return Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
     * 
     */
    public Output<Optional<Integer>> silenceTime() {
        return Codegen.optional(this.silenceTime);
    }
    /**
     * It has been deprecated from provider version 1.50.0 and &#39;effective_interval&#39; instead.
     * 
     * @deprecated
     * Field &#39;start_time&#39; has been deprecated from provider version 1.50.0. New field &#39;effective_interval&#39; instead.
     * 
     */
    @Deprecated /* Field 'start_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead. */
    @Export(name="startTime", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> startTime;

    /**
     * @return It has been deprecated from provider version 1.50.0 and &#39;effective_interval&#39; instead.
     * 
     */
    public Output<Optional<Integer>> startTime() {
        return Codegen.optional(this.startTime);
    }
    /**
     * Critical level alarm statistics method. It must be consistent with that defined for metrics. Valid values: [&#34;Availability&#34;,&#34;Average&#34;, &#34;Minimum&#34;, &#34;Maximum&#34;, &#34;Value&#34;, &#34;ErrorCodeMaximum&#34;, &#34;Sum&#34;, &#34;Count&#34;]. Default to &#34;Average&#34;.
     * 
     * @deprecated
     * Field &#39;statistics&#39; has been deprecated from provider version 1.94.0. New field &#39;escalations_critical.statistics&#39; instead.
     * 
     */
    @Deprecated /* Field 'statistics' has been deprecated from provider version 1.94.0. New field 'escalations_critical.statistics' instead. */
    @Export(name="statistics", type=String.class, parameters={})
    private Output<String> statistics;

    /**
     * @return Critical level alarm statistics method. It must be consistent with that defined for metrics. Valid values: [&#34;Availability&#34;,&#34;Average&#34;, &#34;Minimum&#34;, &#34;Maximum&#34;, &#34;Value&#34;, &#34;ErrorCodeMaximum&#34;, &#34;Sum&#34;, &#34;Count&#34;]. Default to &#34;Average&#34;.
     * 
     */
    public Output<String> statistics() {
        return this.statistics;
    }
    /**
     * The current alarm rule status.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The current alarm rule status.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Critical level alarm threshold value, which must be a numeric value currently.
     * 
     * @deprecated
     * Field &#39;threshold&#39; has been deprecated from provider version 1.94.0. New field &#39;escalations_critical.threshold&#39; instead.
     * 
     */
    @Deprecated /* Field 'threshold' has been deprecated from provider version 1.94.0. New field 'escalations_critical.threshold' instead. */
    @Export(name="threshold", type=String.class, parameters={})
    private Output<String> threshold;

    /**
     * @return Critical level alarm threshold value, which must be a numeric value currently.
     * 
     */
    public Output<String> threshold() {
        return this.threshold;
    }
    /**
     * It has been deprecated from provider version 1.94.0 and &#39;escalations_critical.times&#39; instead.
     * 
     * @deprecated
     * Field &#39;triggered_count&#39; has been deprecated from provider version 1.94.0. New field &#39;escalations_critical.times&#39; instead.
     * 
     */
    @Deprecated /* Field 'triggered_count' has been deprecated from provider version 1.94.0. New field 'escalations_critical.times' instead. */
    @Export(name="triggeredCount", type=Integer.class, parameters={})
    private Output<Integer> triggeredCount;

    /**
     * @return It has been deprecated from provider version 1.94.0 and &#39;escalations_critical.times&#39; instead.
     * 
     */
    public Output<Integer> triggeredCount() {
        return this.triggeredCount;
    }
    /**
     * The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
     * 
     */
    @Export(name="webhook", type=String.class, parameters={})
    private Output</* @Nullable */ String> webhook;

    /**
     * @return The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
     * 
     */
    public Output<Optional<String>> webhook() {
        return Codegen.optional(this.webhook);
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
        super("alicloud:cms/alarm:Alarm", name, args == null ? AlarmArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Alarm(String name, Output<String> id, @Nullable AlarmState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cms/alarm:Alarm", name, state, makeResourceOptions(options, id));
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
