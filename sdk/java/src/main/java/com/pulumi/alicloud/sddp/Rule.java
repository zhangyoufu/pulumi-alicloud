// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sddp;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.sddp.RuleArgs;
import com.pulumi.alicloud.sddp.inputs.RuleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Data Security Center Rule resource.
 * 
 * For information about Data Security Center Rule and how to use it, see [What is Rule](https://www.alibabacloud.com/help/en/data-security-center/latest/api-sddp-2019-01-03-createrule).
 * 
 * &gt; **NOTE:** Available since v1.132.0.
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
 * import com.pulumi.alicloud.sddp.Rule;
 * import com.pulumi.alicloud.sddp.RuleArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf_example_name&#34;);
 *         var default_ = new Rule(&#34;default&#34;, RuleArgs.builder()        
 *             .category(&#34;0&#34;)
 *             .content(&#34;content&#34;)
 *             .ruleName(name)
 *             .riskLevelId(&#34;4&#34;)
 *             .productCode(&#34;OSS&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Data Security Center Rule can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:sddp/rule:Rule example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:sddp/rule:Rule")
public class Rule extends com.pulumi.resources.CustomResource {
    /**
     * Sensitive Data Identification Rules for the Type of. Valid values:
     * 
     */
    @Export(name="category", refs={Integer.class}, tree="[0]")
    private Output<Integer> category;

    /**
     * @return Sensitive Data Identification Rules for the Type of. Valid values:
     * 
     */
    public Output<Integer> category() {
        return this.category;
    }
    /**
     * Sensitive Data Identification Rules the Content.
     * 
     */
    @Export(name="content", refs={String.class}, tree="[0]")
    private Output<String> content;

    /**
     * @return Sensitive Data Identification Rules the Content.
     * 
     */
    public Output<String> content() {
        return this.content;
    }
    /**
     * The Content Classification.
     * 
     */
    @Export(name="contentCategory", refs={String.class}, tree="[0]")
    private Output<String> contentCategory;

    /**
     * @return The Content Classification.
     * 
     */
    public Output<String> contentCategory() {
        return this.contentCategory;
    }
    /**
     * Sensitive Data Identification Rules of Type. Valid values:
     * 
     */
    @Export(name="customType", refs={Integer.class}, tree="[0]")
    private Output<Integer> customType;

    /**
     * @return Sensitive Data Identification Rules of Type. Valid values:
     * 
     */
    public Output<Integer> customType() {
        return this.customType;
    }
    /**
     * Sensitive Data Identification a Description of the Rule Information.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Sensitive Data Identification a Description of the Rule Information.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The Request and Receive the Language of the Message Type. Valid values:
     * 
     */
    @Export(name="lang", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lang;

    /**
     * @return The Request and Receive the Language of the Message Type. Valid values:
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }
    /**
     * Product Code. Valid values: `OSS`,`RDS`,`ODPS`(MaxCompute).
     * 
     */
    @Export(name="productCode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> productCode;

    /**
     * @return Product Code. Valid values: `OSS`,`RDS`,`ODPS`(MaxCompute).
     * 
     */
    public Output<Optional<String>> productCode() {
        return Codegen.optional(this.productCode);
    }
    /**
     * Product ID. Valid values:
     * 
     */
    @Export(name="productId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> productId;

    /**
     * @return Product ID. Valid values:
     * 
     */
    public Output<Optional<String>> productId() {
        return Codegen.optional(this.productId);
    }
    /**
     * Sensitive Data Identification Rules of Risk Level ID. Valid values:
     * 
     */
    @Export(name="riskLevelId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> riskLevelId;

    /**
     * @return Sensitive Data Identification Rules of Risk Level ID. Valid values:
     * 
     */
    public Output<Optional<String>> riskLevelId() {
        return Codegen.optional(this.riskLevelId);
    }
    /**
     * Sensitive Data Identification Name of the Rule.
     * 
     */
    @Export(name="ruleName", refs={String.class}, tree="[0]")
    private Output<String> ruleName;

    /**
     * @return Sensitive Data Identification Name of the Rule.
     * 
     */
    public Output<String> ruleName() {
        return this.ruleName;
    }
    /**
     * Rule Type.
     * 
     */
    @Export(name="ruleType", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> ruleType;

    /**
     * @return Rule Type.
     * 
     */
    public Output<Optional<Integer>> ruleType() {
        return Codegen.optional(this.ruleType);
    }
    /**
     * Triggered the Alarm Conditions.
     * 
     */
    @Export(name="statExpress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> statExpress;

    /**
     * @return Triggered the Alarm Conditions.
     * 
     */
    public Output<Optional<String>> statExpress() {
        return Codegen.optional(this.statExpress);
    }
    /**
     * Sensitive Data Identification Rules Detection State of.
     * 
     */
    @Export(name="status", refs={Integer.class}, tree="[0]")
    private Output<Integer> status;

    /**
     * @return Sensitive Data Identification Rules Detection State of.
     * 
     */
    public Output<Integer> status() {
        return this.status;
    }
    /**
     * The Target of rule.
     * 
     */
    @Export(name="target", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> target;

    /**
     * @return The Target of rule.
     * 
     */
    public Output<Optional<String>> target() {
        return Codegen.optional(this.target);
    }
    /**
     * The Level of Risk. Valid values:
     * 
     */
    @Export(name="warnLevel", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> warnLevel;

    /**
     * @return The Level of Risk. Valid values:
     * 
     */
    public Output<Optional<Integer>> warnLevel() {
        return Codegen.optional(this.warnLevel);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Rule(String name) {
        this(name, RuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Rule(String name, RuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Rule(String name, RuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:sddp/rule:Rule", name, args == null ? RuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Rule(String name, Output<String> id, @Nullable RuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:sddp/rule:Rule", name, state, makeResourceOptions(options, id));
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
    public static Rule get(String name, Output<String> id, @Nullable RuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Rule(name, id, state, options);
    }
}
