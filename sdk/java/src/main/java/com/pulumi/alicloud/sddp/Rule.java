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
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         final var name = config.get("name").orElse("tf-example-name");
 *         var default_ = new Rule("default", RuleArgs.builder()        
 *             .ruleName(name)
 *             .category("2")
 *             .content("""
 *   [
 *     {
 *       "rule": [
 *         {
 *           "operator": "contains",
 *           "target": "content",
 *           "value": "tf-testACCContent"
 *         }
 *       ],
 *       "ruleRelation": "AND"
 *     }
 *   ]
 *             """)
 *             .riskLevelId("4")
 *             .productCode("OSS")
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
     * The content type of the sensitive data detection rule. Valid values:
     * 
     */
    @Export(name="category", refs={Integer.class}, tree="[0]")
    private Output<Integer> category;

    /**
     * @return The content type of the sensitive data detection rule. Valid values:
     * 
     */
    public Output<Integer> category() {
        return this.category;
    }
    /**
     * The content of the sensitive data detection rule. **NOTE:** From version 1.222.0, `content` can be modified.
     * 
     */
    @Export(name="content", refs={String.class}, tree="[0]")
    private Output<String> content;

    /**
     * @return The content of the sensitive data detection rule. **NOTE:** From version 1.222.0, `content` can be modified.
     * 
     */
    public Output<String> content() {
        return this.content;
    }
    /**
     * The type of the content in the sensitive data detection rule. **NOTE:** From version 1.222.0, `content_category` cannot be modified.
     * 
     */
    @Export(name="contentCategory", refs={String.class}, tree="[0]")
    private Output<String> contentCategory;

    /**
     * @return The type of the content in the sensitive data detection rule. **NOTE:** From version 1.222.0, `content_category` cannot be modified.
     * 
     */
    public Output<String> contentCategory() {
        return this.contentCategory;
    }
    /**
     * The type of the sensitive data detection rule. **NOTE:** From version 1.222.0, `custom_type` cannot be specified when create Rule.
     * 
     */
    @Export(name="customType", refs={Integer.class}, tree="[0]")
    private Output<Integer> customType;

    /**
     * @return The type of the sensitive data detection rule. **NOTE:** From version 1.222.0, `custom_type` cannot be specified when create Rule.
     * 
     */
    public Output<Integer> customType() {
        return this.customType;
    }
    /**
     * The description of the rule. **NOTE:** From version 1.222.0, `description` cannot be modified.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the rule. **NOTE:** From version 1.222.0, `description` cannot be modified.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The language of the content within the request and response. Default value: `zh`. Valid values:
     * 
     */
    @Export(name="lang", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lang;

    /**
     * @return The language of the content within the request and response. Default value: `zh`. Valid values:
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }
    /**
     * The name of the service to which data in the column of the table belongs. Valid values: `OSS`, `RDS`, `ODPS`(MaxCompute).
     * 
     */
    @Export(name="productCode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> productCode;

    /**
     * @return The name of the service to which data in the column of the table belongs. Valid values: `OSS`, `RDS`, `ODPS`(MaxCompute).
     * 
     */
    public Output<Optional<String>> productCode() {
        return Codegen.optional(this.productCode);
    }
    /**
     * The ID of the service to which the data asset belongs. Valid values:
     * 
     */
    @Export(name="productId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> productId;

    /**
     * @return The ID of the service to which the data asset belongs. Valid values:
     * 
     */
    public Output<Optional<String>> productId() {
        return Codegen.optional(this.productId);
    }
    /**
     * The sensitivity level of the sensitive data that hits the sensitive data detection rule. Valid values:
     * 
     */
    @Export(name="riskLevelId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> riskLevelId;

    /**
     * @return The sensitivity level of the sensitive data that hits the sensitive data detection rule. Valid values:
     * 
     */
    public Output<Optional<String>> riskLevelId() {
        return Codegen.optional(this.riskLevelId);
    }
    /**
     * The name of the sensitive data detection rule. **NOTE:** From version 1.222.0, `rule_name` can be modified.
     * 
     */
    @Export(name="ruleName", refs={String.class}, tree="[0]")
    private Output<String> ruleName;

    /**
     * @return The name of the sensitive data detection rule. **NOTE:** From version 1.222.0, `rule_name` can be modified.
     * 
     */
    public Output<String> ruleName() {
        return this.ruleName;
    }
    /**
     * The type of the sensitive data detection rule. Valid values:
     * 
     */
    @Export(name="ruleType", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> ruleType;

    /**
     * @return The type of the sensitive data detection rule. Valid values:
     * 
     */
    public Output<Optional<Integer>> ruleType() {
        return Codegen.optional(this.ruleType);
    }
    /**
     * The statistical expression. **NOTE:** From version 1.222.0, `stat_express` cannot be modified.
     * 
     */
    @Export(name="statExpress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> statExpress;

    /**
     * @return The statistical expression. **NOTE:** From version 1.222.0, `stat_express` cannot be modified.
     * 
     */
    public Output<Optional<String>> statExpress() {
        return Codegen.optional(this.statExpress);
    }
    /**
     * Sensitive Specifies whether to enable the sensitive data detection rule. Valid values:
     * 
     */
    @Export(name="status", refs={Integer.class}, tree="[0]")
    private Output<Integer> status;

    /**
     * @return Sensitive Specifies whether to enable the sensitive data detection rule. Valid values:
     * 
     */
    public Output<Integer> status() {
        return this.status;
    }
    /**
     * The code of the service to which the sensitive data detection rule is applied. **NOTE:** From version 1.222.0, `target` cannot be modified.
     * 
     */
    @Export(name="target", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> target;

    /**
     * @return The code of the service to which the sensitive data detection rule is applied. **NOTE:** From version 1.222.0, `target` cannot be modified.
     * 
     */
    public Output<Optional<String>> target() {
        return Codegen.optional(this.target);
    }
    /**
     * The risk level of the alert that is triggered. Valid values:
     * 
     */
    @Export(name="warnLevel", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> warnLevel;

    /**
     * @return The risk level of the alert that is triggered. Valid values:
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
