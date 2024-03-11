// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cfg;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cfg.CompliancePackArgs;
import com.pulumi.alicloud.cfg.inputs.CompliancePackState;
import com.pulumi.alicloud.cfg.outputs.CompliancePackConfigRule;
import com.pulumi.alicloud.cfg.outputs.CompliancePackConfigRuleId;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Config Compliance Pack resource.
 * 
 * For information about Cloud Config Compliance Pack and how to use it, see [What is Compliance Pack](https://www.alibabacloud.com/help/en/cloud-config/latest/api-config-2020-09-07-createcompliancepack).
 * 
 * &gt; **NOTE:** Available since v1.124.0.
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetRegionsArgs;
 * import com.pulumi.alicloud.cfg.Rule;
 * import com.pulumi.alicloud.cfg.RuleArgs;
 * import com.pulumi.alicloud.cfg.CompliancePack;
 * import com.pulumi.alicloud.cfg.CompliancePackArgs;
 * import com.pulumi.alicloud.cfg.inputs.CompliancePackConfigRuleIdArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-example-config-name&#34;);
 *         final var defaultRegions = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         var rule1 = new Rule(&#34;rule1&#34;, RuleArgs.builder()        
 *             .description(name)
 *             .sourceOwner(&#34;ALIYUN&#34;)
 *             .sourceIdentifier(&#34;ram-user-ak-create-date-expired-check&#34;)
 *             .riskLevel(1)
 *             .maximumExecutionFrequency(&#34;TwentyFour_Hours&#34;)
 *             .regionIdsScope(defaultRegions.applyValue(getRegionsResult -&gt; getRegionsResult.regions()[0].id()))
 *             .configRuleTriggerTypes(&#34;ScheduledNotification&#34;)
 *             .resourceTypesScopes(&#34;ACS::RAM::User&#34;)
 *             .ruleName(&#34;ciscompliancecheck_ram-user-ak-create-date-expired-check&#34;)
 *             .inputParameters(Map.of(&#34;days&#34;, &#34;90&#34;))
 *             .build());
 * 
 *         var rule2 = new Rule(&#34;rule2&#34;, RuleArgs.builder()        
 *             .description(name)
 *             .sourceOwner(&#34;ALIYUN&#34;)
 *             .sourceIdentifier(&#34;adb-cluster-maintain-time-check&#34;)
 *             .riskLevel(2)
 *             .regionIdsScope(defaultRegions.applyValue(getRegionsResult -&gt; getRegionsResult.regions()[0].id()))
 *             .configRuleTriggerTypes(&#34;ScheduledNotification&#34;)
 *             .resourceTypesScopes(&#34;ACS::ADB::DBCluster&#34;)
 *             .ruleName(&#34;governance-evaluation-adb-cluster-maintain-time-check&#34;)
 *             .inputParameters(Map.of(&#34;maintainTimes&#34;, &#34;02:00-04:00,06:00-08:00,12:00-13:00&#34;))
 *             .build());
 * 
 *         var defaultCompliancePack = new CompliancePack(&#34;defaultCompliancePack&#34;, CompliancePackArgs.builder()        
 *             .compliancePackName(name)
 *             .description(&#34;CloudGovernanceCenter evaluation&#34;)
 *             .riskLevel(&#34;2&#34;)
 *             .configRuleIds(            
 *                 CompliancePackConfigRuleIdArgs.builder()
 *                     .configRuleId(rule1.id())
 *                     .build(),
 *                 CompliancePackConfigRuleIdArgs.builder()
 *                     .configRuleId(rule2.id())
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Cloud Config Compliance Pack can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cfg/compliancePack:CompliancePack example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cfg/compliancePack:CompliancePack")
public class CompliancePack extends com.pulumi.resources.CustomResource {
    /**
     * The Compliance Package Name. **NOTE:** From version 1.146.0, `compliance_pack_name` can be modified.
     * 
     */
    @Export(name="compliancePackName", refs={String.class}, tree="[0]")
    private Output<String> compliancePackName;

    /**
     * @return The Compliance Package Name. **NOTE:** From version 1.146.0, `compliance_pack_name` can be modified.
     * 
     */
    public Output<String> compliancePackName() {
        return this.compliancePackName;
    }
    /**
     * Compliance Package Template Id.
     * 
     */
    @Export(name="compliancePackTemplateId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> compliancePackTemplateId;

    /**
     * @return Compliance Package Template Id.
     * 
     */
    public Output<Optional<String>> compliancePackTemplateId() {
        return Codegen.optional(this.compliancePackTemplateId);
    }
    /**
     * A list of Config Rule IDs. See `config_rule_ids` below.
     * 
     */
    @Export(name="configRuleIds", refs={List.class,CompliancePackConfigRuleId.class}, tree="[0,1]")
    private Output</* @Nullable */ List<CompliancePackConfigRuleId>> configRuleIds;

    /**
     * @return A list of Config Rule IDs. See `config_rule_ids` below.
     * 
     */
    public Output<Optional<List<CompliancePackConfigRuleId>>> configRuleIds() {
        return Codegen.optional(this.configRuleIds);
    }
    /**
     * A list of Config Rules. See `config_rules` below. **NOTE:** Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead.
     * 
     * @deprecated
     * Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead.
     * 
     */
    @Deprecated /* Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead. */
    @Export(name="configRules", refs={List.class,CompliancePackConfigRule.class}, tree="[0,1]")
    private Output</* @Nullable */ List<CompliancePackConfigRule>> configRules;

    /**
     * @return A list of Config Rules. See `config_rules` below. **NOTE:** Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead.
     * 
     */
    public Output<Optional<List<CompliancePackConfigRule>>> configRules() {
        return Codegen.optional(this.configRules);
    }
    /**
     * The Description of compliance pack.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The Description of compliance pack.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The Risk Level. Valid values:
     * 
     */
    @Export(name="riskLevel", refs={Integer.class}, tree="[0]")
    private Output<Integer> riskLevel;

    /**
     * @return The Risk Level. Valid values:
     * 
     */
    public Output<Integer> riskLevel() {
        return this.riskLevel;
    }
    /**
     * The status of the Compliance Pack.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the Compliance Pack.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CompliancePack(String name) {
        this(name, CompliancePackArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CompliancePack(String name, CompliancePackArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CompliancePack(String name, CompliancePackArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cfg/compliancePack:CompliancePack", name, args == null ? CompliancePackArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CompliancePack(String name, Output<String> id, @Nullable CompliancePackState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cfg/compliancePack:CompliancePack", name, state, makeResourceOptions(options, id));
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
    public static CompliancePack get(String name, Output<String> id, @Nullable CompliancePackState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CompliancePack(name, id, state, options);
    }
}
