// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.threatdetection.AntiBruteForceRuleArgs;
import com.pulumi.alicloud.threatdetection.inputs.AntiBruteForceRuleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a Threat Detection Anti Brute Force Rule resource.
 * 
 * For information about Threat Detection Anti Brute Force Rule and how to use it, see [What is Anti Brute Force Rule](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createantibruteforcerule).
 * 
 * &gt; **NOTE:** Available since v1.195.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.threatdetection.AntiBruteForceRule;
 * import com.pulumi.alicloud.threatdetection.AntiBruteForceRuleArgs;
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
 *         var default_ = new AntiBruteForceRule("default", AntiBruteForceRuleArgs.builder()
 *             .antiBruteForceRuleName("apispec_example")
 *             .forbiddenTime(360)
 *             .uuidLists("032b618f-b220-4a0d-bd37-fbdc6ef58b6a")
 *             .failCount(80)
 *             .span(10)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Threat Detection Anti Brute Force Rule can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:threatdetection/antiBruteForceRule:AntiBruteForceRule example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:threatdetection/antiBruteForceRule:AntiBruteForceRule")
public class AntiBruteForceRule extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the defense rule.
     * 
     */
    @Export(name="antiBruteForceRuleId", refs={String.class}, tree="[0]")
    private Output<String> antiBruteForceRuleId;

    /**
     * @return The ID of the defense rule.
     * 
     */
    public Output<String> antiBruteForceRuleId() {
        return this.antiBruteForceRuleId;
    }
    /**
     * The name of the defense rule.
     * 
     */
    @Export(name="antiBruteForceRuleName", refs={String.class}, tree="[0]")
    private Output<String> antiBruteForceRuleName;

    /**
     * @return The name of the defense rule.
     * 
     */
    public Output<String> antiBruteForceRuleName() {
        return this.antiBruteForceRuleName;
    }
    /**
     * Specifies whether to set the defense rule as the default rule.
     * 
     */
    @Export(name="defaultRule", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> defaultRule;

    /**
     * @return Specifies whether to set the defense rule as the default rule.
     * 
     */
    public Output<Boolean> defaultRule() {
        return this.defaultRule;
    }
    /**
     * The threshold for the number of failed user logins when the brute-force defense rule takes effect.
     * 
     */
    @Export(name="failCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> failCount;

    /**
     * @return The threshold for the number of failed user logins when the brute-force defense rule takes effect.
     * 
     */
    public Output<Integer> failCount() {
        return this.failCount;
    }
    /**
     * The period of time during which logons from an account are not allowed. Unit: minutes.
     * 
     */
    @Export(name="forbiddenTime", refs={Integer.class}, tree="[0]")
    private Output<Integer> forbiddenTime;

    /**
     * @return The period of time during which logons from an account are not allowed. Unit: minutes.
     * 
     */
    public Output<Integer> forbiddenTime() {
        return this.forbiddenTime;
    }
    /**
     * The period of time during which logon failures from an account are measured. Unit: minutes. If Span is set to 10, the defense rule takes effect when the logon failures measured within 10 minutes reaches the specified threshold. The IP address of attackers cannot be used to log on to the server in the specified period of time.
     * 
     */
    @Export(name="span", refs={Integer.class}, tree="[0]")
    private Output<Integer> span;

    /**
     * @return The period of time during which logon failures from an account are measured. Unit: minutes. If Span is set to 10, the defense rule takes effect when the logon failures measured within 10 minutes reaches the specified threshold. The IP address of attackers cannot be used to log on to the server in the specified period of time.
     * 
     */
    public Output<Integer> span() {
        return this.span;
    }
    /**
     * An array consisting of the UUIDs of servers to which the defense rule is applied.**The binding status must be Enterprise Edition.**
     * 
     */
    @Export(name="uuidLists", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> uuidLists;

    /**
     * @return An array consisting of the UUIDs of servers to which the defense rule is applied.**The binding status must be Enterprise Edition.**
     * 
     */
    public Output<List<String>> uuidLists() {
        return this.uuidLists;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AntiBruteForceRule(java.lang.String name) {
        this(name, AntiBruteForceRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AntiBruteForceRule(java.lang.String name, AntiBruteForceRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AntiBruteForceRule(java.lang.String name, AntiBruteForceRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/antiBruteForceRule:AntiBruteForceRule", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AntiBruteForceRule(java.lang.String name, Output<java.lang.String> id, @Nullable AntiBruteForceRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/antiBruteForceRule:AntiBruteForceRule", name, state, makeResourceOptions(options, id), false);
    }

    private static AntiBruteForceRuleArgs makeArgs(AntiBruteForceRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AntiBruteForceRuleArgs.Empty : args;
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
    public static AntiBruteForceRule get(java.lang.String name, Output<java.lang.String> id, @Nullable AntiBruteForceRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AntiBruteForceRule(name, id, state, options);
    }
}
