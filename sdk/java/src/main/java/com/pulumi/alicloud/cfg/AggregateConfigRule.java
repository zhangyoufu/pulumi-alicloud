// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cfg;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cfg.AggregateConfigRuleArgs;
import com.pulumi.alicloud.cfg.inputs.AggregateConfigRuleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Config Aggregate Config Rule resource.
 * 
 * For information about Cloud Config Aggregate Config Rule and how to use it, see [What is Aggregate Config Rule](https://www.alibabacloud.com/help/doc-detail/154216.html).
 * 
 * &gt; **NOTE:** Available in v1.124.0+.
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
 * import com.pulumi.alicloud.cfg.Aggregator;
 * import com.pulumi.alicloud.cfg.AggregatorArgs;
 * import com.pulumi.alicloud.cfg.inputs.AggregatorAggregatorAccountArgs;
 * import com.pulumi.alicloud.cfg.AggregateConfigRule;
 * import com.pulumi.alicloud.cfg.AggregateConfigRuleArgs;
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
 *         var exampleAggregator = new Aggregator(&#34;exampleAggregator&#34;, AggregatorArgs.builder()        
 *             .aggregatorAccounts(AggregatorAggregatorAccountArgs.builder()
 *                 .accountId(&#34;140278452670****&#34;)
 *                 .accountName(&#34;test-2&#34;)
 *                 .accountType(&#34;ResourceDirectory&#34;)
 *                 .build())
 *             .aggregatorName(&#34;tf-testaccaggregator&#34;)
 *             .description(&#34;tf-testaccaggregator&#34;)
 *             .build());
 * 
 *         var exampleAggregateConfigRule = new AggregateConfigRule(&#34;exampleAggregateConfigRule&#34;, AggregateConfigRuleArgs.builder()        
 *             .aggregateConfigRuleName(&#34;tf-testaccconfig1234&#34;)
 *             .aggregatorId(exampleAggregator.id())
 *             .configRuleTriggerTypes(&#34;ConfigurationItemChangeNotification&#34;)
 *             .sourceOwner(&#34;ALIYUN&#34;)
 *             .sourceIdentifier(&#34;ecs-cpu-min-count-limit&#34;)
 *             .riskLevel(1)
 *             .resourceTypesScopes(&#34;ACS::ECS::Instance&#34;)
 *             .inputParameters(Map.of(&#34;cpuCount&#34;, &#34;4&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cloud Config Aggregate Config Rule can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:cfg/aggregateConfigRule:AggregateConfigRule example &lt;aggregator_id&gt;:&lt;config_rule_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cfg/aggregateConfigRule:AggregateConfigRule")
public class AggregateConfigRule extends com.pulumi.resources.CustomResource {
    /**
     * The name of the rule.
     * 
     */
    @Export(name="aggregateConfigRuleName", type=String.class, parameters={})
    private Output<String> aggregateConfigRuleName;

    /**
     * @return The name of the rule.
     * 
     */
    public Output<String> aggregateConfigRuleName() {
        return this.aggregateConfigRuleName;
    }
    /**
     * The Aggregator Id.
     * 
     */
    @Export(name="aggregatorId", type=String.class, parameters={})
    private Output<String> aggregatorId;

    /**
     * @return The Aggregator Id.
     * 
     */
    public Output<String> aggregatorId() {
        return this.aggregatorId;
    }
    /**
     * (Available in 1.141.0+) The rule ID of Aggregate Config Rule.
     * 
     */
    @Export(name="configRuleId", type=String.class, parameters={})
    private Output<String> configRuleId;

    /**
     * @return (Available in 1.141.0+) The rule ID of Aggregate Config Rule.
     * 
     */
    public Output<String> configRuleId() {
        return this.configRuleId;
    }
    /**
     * The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
     * 
     */
    @Export(name="configRuleTriggerTypes", type=String.class, parameters={})
    private Output<String> configRuleTriggerTypes;

    /**
     * @return The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
     * 
     */
    public Output<String> configRuleTriggerTypes() {
        return this.configRuleTriggerTypes;
    }
    /**
     * The description of the rule.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the rule.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
     * 
     */
    @Export(name="excludeResourceIdsScope", type=String.class, parameters={})
    private Output</* @Nullable */ String> excludeResourceIdsScope;

    /**
     * @return The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
     * 
     */
    public Output<Optional<String>> excludeResourceIdsScope() {
        return Codegen.optional(this.excludeResourceIdsScope);
    }
    /**
     * The settings map of the input parameters for the rule.
     * 
     */
    @Export(name="inputParameters", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> inputParameters;

    /**
     * @return The settings map of the input parameters for the rule.
     * 
     */
    public Output<Optional<Map<String,Object>>> inputParameters() {
        return Codegen.optional(this.inputParameters);
    }
    /**
     * The frequency of the compliance evaluations. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`. System default value is `TwentyFour_Hours` and valid when the `config_rule_trigger_types` is `ScheduledNotification`.
     * 
     */
    @Export(name="maximumExecutionFrequency", type=String.class, parameters={})
    private Output<String> maximumExecutionFrequency;

    /**
     * @return The frequency of the compliance evaluations. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`. System default value is `TwentyFour_Hours` and valid when the `config_rule_trigger_types` is `ScheduledNotification`.
     * 
     */
    public Output<String> maximumExecutionFrequency() {
        return this.maximumExecutionFrequency;
    }
    /**
     * The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
     * 
     */
    @Export(name="regionIdsScope", type=String.class, parameters={})
    private Output</* @Nullable */ String> regionIdsScope;

    /**
     * @return The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
     * 
     */
    public Output<Optional<String>> regionIdsScope() {
        return Codegen.optional(this.regionIdsScope);
    }
    /**
     * The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
     * 
     */
    @Export(name="resourceGroupIdsScope", type=String.class, parameters={})
    private Output</* @Nullable */ String> resourceGroupIdsScope;

    /**
     * @return The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
     * 
     */
    public Output<Optional<String>> resourceGroupIdsScope() {
        return Codegen.optional(this.resourceGroupIdsScope);
    }
    /**
     * Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
     * 
     */
    @Export(name="resourceTypesScopes", type=List.class, parameters={String.class})
    private Output<List<String>> resourceTypesScopes;

    /**
     * @return Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
     * 
     */
    public Output<List<String>> resourceTypesScopes() {
        return this.resourceTypesScopes;
    }
    /**
     * The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info.
     * 
     */
    @Export(name="riskLevel", type=Integer.class, parameters={})
    private Output<Integer> riskLevel;

    /**
     * @return The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info.
     * 
     */
    public Output<Integer> riskLevel() {
        return this.riskLevel;
    }
    /**
     * The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
     * 
     */
    @Export(name="sourceIdentifier", type=String.class, parameters={})
    private Output<String> sourceIdentifier;

    /**
     * @return The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
     * 
     */
    public Output<String> sourceIdentifier() {
        return this.sourceIdentifier;
    }
    /**
     * Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
     * 
     */
    @Export(name="sourceOwner", type=String.class, parameters={})
    private Output<String> sourceOwner;

    /**
     * @return Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
     * 
     */
    public Output<String> sourceOwner() {
        return this.sourceOwner;
    }
    /**
     * The rule status. The valid values: `ACTIVE`, `INACTIVE`.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The rule status. The valid values: `ACTIVE`, `INACTIVE`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The rule monitors the tag key, only applies to rules created based on managed rules.
     * 
     */
    @Export(name="tagKeyScope", type=String.class, parameters={})
    private Output</* @Nullable */ String> tagKeyScope;

    /**
     * @return The rule monitors the tag key, only applies to rules created based on managed rules.
     * 
     */
    public Output<Optional<String>> tagKeyScope() {
        return Codegen.optional(this.tagKeyScope);
    }
    /**
     * The rule monitors the tag value, use with the `tag_key_scope` options. only applies to rules created based on managed rules.
     * 
     */
    @Export(name="tagValueScope", type=String.class, parameters={})
    private Output</* @Nullable */ String> tagValueScope;

    /**
     * @return The rule monitors the tag value, use with the `tag_key_scope` options. only applies to rules created based on managed rules.
     * 
     */
    public Output<Optional<String>> tagValueScope() {
        return Codegen.optional(this.tagValueScope);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AggregateConfigRule(String name) {
        this(name, AggregateConfigRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AggregateConfigRule(String name, AggregateConfigRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AggregateConfigRule(String name, AggregateConfigRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cfg/aggregateConfigRule:AggregateConfigRule", name, args == null ? AggregateConfigRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AggregateConfigRule(String name, Output<String> id, @Nullable AggregateConfigRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cfg/aggregateConfigRule:AggregateConfigRule", name, state, makeResourceOptions(options, id));
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
    public static AggregateConfigRule get(String name, Output<String> id, @Nullable AggregateConfigRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AggregateConfigRule(name, id, state, options);
    }
}
