// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.quotas;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.quotas.QuotaAlarmArgs;
import com.pulumi.alicloud.quotas.inputs.QuotaAlarmState;
import com.pulumi.alicloud.quotas.outputs.QuotaAlarmQuotaDimension;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Quotas Quota Alarm resource.
 * 
 * For information about Quotas Quota Alarm and how to use it, see [What is Quota Alarm](https://www.alibabacloud.com/help/en/quota-center/latest/api-doc-quotas-2020-05-10-api-doc-createquotaalarm).
 * 
 * &gt; **NOTE:** Available since v1.116.0.
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
 * import com.pulumi.alicloud.quotas.QuotaAlarm;
 * import com.pulumi.alicloud.quotas.QuotaAlarmArgs;
 * import com.pulumi.alicloud.quotas.inputs.QuotaAlarmQuotaDimensionArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         var default_ = new QuotaAlarm(&#34;default&#34;, QuotaAlarmArgs.builder()        
 *             .quotaActionCode(&#34;q_desktop-count&#34;)
 *             .quotaDimensions(QuotaAlarmQuotaDimensionArgs.builder()
 *                 .key(&#34;regionId&#34;)
 *                 .value(&#34;cn-hangzhou&#34;)
 *                 .build())
 *             .thresholdPercent(80)
 *             .productCode(&#34;gws&#34;)
 *             .quotaAlarmName(name)
 *             .thresholdType(&#34;used&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Quotas Quota Alarm can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:quotas/quotaAlarm:QuotaAlarm example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:quotas/quotaAlarm:QuotaAlarm")
public class QuotaAlarm extends com.pulumi.resources.CustomResource {
    /**
     * The creation time of the resource.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return The creation time of the resource.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The Product Code.
     * 
     */
    @Export(name="productCode", type=String.class, parameters={})
    private Output<String> productCode;

    /**
     * @return The Product Code.
     * 
     */
    public Output<String> productCode() {
        return this.productCode;
    }
    /**
     * The Quota Action Code.
     * 
     */
    @Export(name="quotaActionCode", type=String.class, parameters={})
    private Output<String> quotaActionCode;

    /**
     * @return The Quota Action Code.
     * 
     */
    public Output<String> quotaActionCode() {
        return this.quotaActionCode;
    }
    /**
     * The name of Quota Alarm.
     * 
     */
    @Export(name="quotaAlarmName", type=String.class, parameters={})
    private Output<String> quotaAlarmName;

    /**
     * @return The name of Quota Alarm.
     * 
     */
    public Output<String> quotaAlarmName() {
        return this.quotaAlarmName;
    }
    /**
     * The Quota Dimensions. See `quota_dimensions` below.
     * 
     */
    @Export(name="quotaDimensions", type=List.class, parameters={QuotaAlarmQuotaDimension.class})
    private Output</* @Nullable */ List<QuotaAlarmQuotaDimension>> quotaDimensions;

    /**
     * @return The Quota Dimensions. See `quota_dimensions` below.
     * 
     */
    public Output<Optional<List<QuotaAlarmQuotaDimension>>> quotaDimensions() {
        return Codegen.optional(this.quotaDimensions);
    }
    /**
     * The threshold of Quota Alarm.
     * 
     */
    @Export(name="threshold", type=Double.class, parameters={})
    private Output</* @Nullable */ Double> threshold;

    /**
     * @return The threshold of Quota Alarm.
     * 
     */
    public Output<Optional<Double>> threshold() {
        return Codegen.optional(this.threshold);
    }
    /**
     * The threshold percent of Quota Alarm.
     * 
     */
    @Export(name="thresholdPercent", type=Double.class, parameters={})
    private Output</* @Nullable */ Double> thresholdPercent;

    /**
     * @return The threshold percent of Quota Alarm.
     * 
     */
    public Output<Optional<Double>> thresholdPercent() {
        return Codegen.optional(this.thresholdPercent);
    }
    /**
     * Quota alarm type. Value:
     * - used: Quota used alarm.
     * - usable: alarm for the remaining available quota.
     * 
     */
    @Export(name="thresholdType", type=String.class, parameters={})
    private Output<String> thresholdType;

    /**
     * @return Quota alarm type. Value:
     * - used: Quota used alarm.
     * - usable: alarm for the remaining available quota.
     * 
     */
    public Output<String> thresholdType() {
        return this.thresholdType;
    }
    /**
     * The WebHook of Quota Alarm.
     * 
     */
    @Export(name="webHook", type=String.class, parameters={})
    private Output</* @Nullable */ String> webHook;

    /**
     * @return The WebHook of Quota Alarm.
     * 
     */
    public Output<Optional<String>> webHook() {
        return Codegen.optional(this.webHook);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public QuotaAlarm(String name) {
        this(name, QuotaAlarmArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QuotaAlarm(String name, QuotaAlarmArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QuotaAlarm(String name, QuotaAlarmArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:quotas/quotaAlarm:QuotaAlarm", name, args == null ? QuotaAlarmArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private QuotaAlarm(String name, Output<String> id, @Nullable QuotaAlarmState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:quotas/quotaAlarm:QuotaAlarm", name, state, makeResourceOptions(options, id));
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
    public static QuotaAlarm get(String name, Output<String> id, @Nullable QuotaAlarmState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QuotaAlarm(name, id, state, options);
    }
}
