// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.threatdetection.BaselineStrategyArgs;
import com.pulumi.alicloud.threatdetection.inputs.BaselineStrategyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Threat Detection Baseline Strategy resource.
 * 
 * For information about Threat Detection Baseline Strategy and how to use it, see [What is Baseline Strategy](https://www.alibabacloud.com/help/en/security-center/latest/api-sas-2018-12-03-modifystrategy).
 * 
 * &gt; **NOTE:** Available since v1.195.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.threatdetection.BaselineStrategy;
 * import com.pulumi.alicloud.threatdetection.BaselineStrategyArgs;
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
 *         var default_ = new BaselineStrategy(&#34;default&#34;, BaselineStrategyArgs.builder()        
 *             .customType(&#34;custom&#34;)
 *             .endTime(&#34;08:00:00&#34;)
 *             .baselineStrategyName(&#34;apispec&#34;)
 *             .cycleDays(3)
 *             .targetType(&#34;groupId&#34;)
 *             .startTime(&#34;05:00:00&#34;)
 *             .riskSubTypeName(&#34;hc_exploit_redis&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Threat Detection Baseline Strategy can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:threatdetection/baselineStrategy:BaselineStrategy example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:threatdetection/baselineStrategy:BaselineStrategy")
public class BaselineStrategy extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the baseline check policy.
     * 
     */
    @Export(name="baselineStrategyId", refs={String.class}, tree="[0]")
    private Output<String> baselineStrategyId;

    /**
     * @return The ID of the baseline check policy.
     * 
     */
    public Output<String> baselineStrategyId() {
        return this.baselineStrategyId;
    }
    /**
     * Policy name.
     * 
     */
    @Export(name="baselineStrategyName", refs={String.class}, tree="[0]")
    private Output<String> baselineStrategyName;

    /**
     * @return Policy name.
     * 
     */
    public Output<String> baselineStrategyName() {
        return this.baselineStrategyName;
    }
    /**
     * The type of policy. Value:
     * * **common**: standard policy
     * * **custom**: custom policy
     * 
     */
    @Export(name="customType", refs={String.class}, tree="[0]")
    private Output<String> customType;

    /**
     * @return The type of policy. Value:
     * * **common**: standard policy
     * * **custom**: custom policy
     * 
     */
    public Output<String> customType() {
        return this.customType;
    }
    /**
     * The detection period of the policy.
     * 
     */
    @Export(name="cycleDays", refs={Integer.class}, tree="[0]")
    private Output<Integer> cycleDays;

    /**
     * @return The detection period of the policy.
     * 
     */
    public Output<Integer> cycleDays() {
        return this.cycleDays;
    }
    /**
     * The detection period of the policy. Value:
     * * **0**: 0:00~06:00
     * * **6**: 6:00~12:00
     * * **12**: 12:00~18:00
     * * **18**: 18:00~24:00
     * 
     */
    @Export(name="cycleStartTime", refs={Integer.class}, tree="[0]")
    private Output<Integer> cycleStartTime;

    /**
     * @return The detection period of the policy. Value:
     * * **0**: 0:00~06:00
     * * **6**: 6:00~12:00
     * * **12**: 12:00~18:00
     * * **18**: 18:00~24:00
     * 
     */
    public Output<Integer> cycleStartTime() {
        return this.cycleStartTime;
    }
    /**
     * The baseline check policy execution end time.
     * 
     */
    @Export(name="endTime", refs={String.class}, tree="[0]")
    private Output<String> endTime;

    /**
     * @return The baseline check policy execution end time.
     * 
     */
    public Output<String> endTime() {
        return this.endTime;
    }
    /**
     * Detection item subtype.
     * 
     */
    @Export(name="riskSubTypeName", refs={String.class}, tree="[0]")
    private Output<String> riskSubTypeName;

    /**
     * @return Detection item subtype.
     * 
     */
    public Output<String> riskSubTypeName() {
        return this.riskSubTypeName;
    }
    /**
     * The baseline check policy start time.
     * 
     */
    @Export(name="startTime", refs={String.class}, tree="[0]")
    private Output<String> startTime;

    /**
     * @return The baseline check policy start time.
     * 
     */
    public Output<String> startTime() {
        return this.startTime;
    }
    /**
     * The method of adding assets that take effect from the policy. Value:
     * * **groupId**: Added by asset group.
     * * **uuid**: Add by single asset.
     * 
     */
    @Export(name="targetType", refs={String.class}, tree="[0]")
    private Output<String> targetType;

    /**
     * @return The method of adding assets that take effect from the policy. Value:
     * * **groupId**: Added by asset group.
     * * **uuid**: Add by single asset.
     * 
     */
    public Output<String> targetType() {
        return this.targetType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BaselineStrategy(String name) {
        this(name, BaselineStrategyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BaselineStrategy(String name, BaselineStrategyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BaselineStrategy(String name, BaselineStrategyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/baselineStrategy:BaselineStrategy", name, args == null ? BaselineStrategyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BaselineStrategy(String name, Output<String> id, @Nullable BaselineStrategyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/baselineStrategy:BaselineStrategy", name, state, makeResourceOptions(options, id));
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
    public static BaselineStrategy get(String name, Output<String> id, @Nullable BaselineStrategyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BaselineStrategy(name, id, state, options);
    }
}
