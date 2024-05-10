// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.arms.DispatchRuleArgs;
import com.pulumi.alicloud.arms.inputs.DispatchRuleState;
import com.pulumi.alicloud.arms.outputs.DispatchRuleGroupRule;
import com.pulumi.alicloud.arms.outputs.DispatchRuleLabelMatchExpressionGrid;
import com.pulumi.alicloud.arms.outputs.DispatchRuleNotifyRule;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Application Real-Time Monitoring Service (ARMS) Alert Dispatch Rule resource.
 * 
 * For information about Application Real-Time Monitoring Service (ARMS) Alert Dispatch Rule and how to use it, see [What is Alert Dispatch_Rule](https://next.api.alibabacloud.com/document/ARMS/2019-08-08/CreateDispatchRule).
 * 
 * &gt; **NOTE:** Available since v1.136.0.
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
 * import com.pulumi.alicloud.arms.AlertContact;
 * import com.pulumi.alicloud.arms.AlertContactArgs;
 * import com.pulumi.alicloud.arms.AlertContactGroup;
 * import com.pulumi.alicloud.arms.AlertContactGroupArgs;
 * import com.pulumi.alicloud.arms.DispatchRule;
 * import com.pulumi.alicloud.arms.DispatchRuleArgs;
 * import com.pulumi.alicloud.arms.inputs.DispatchRuleGroupRuleArgs;
 * import com.pulumi.alicloud.arms.inputs.DispatchRuleLabelMatchExpressionGridArgs;
 * import com.pulumi.alicloud.arms.inputs.DispatchRuleNotifyRuleArgs;
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
 *         var default_ = new AlertContact("default", AlertContactArgs.builder()        
 *             .alertContactName("example_value")
 *             .email("example_value{@literal @}aaa.com")
 *             .build());
 * 
 *         var defaultAlertContactGroup = new AlertContactGroup("defaultAlertContactGroup", AlertContactGroupArgs.builder()        
 *             .alertContactGroupName("example_value")
 *             .contactIds(default_.id())
 *             .build());
 * 
 *         var defaultDispatchRule = new DispatchRule("defaultDispatchRule", DispatchRuleArgs.builder()        
 *             .dispatchRuleName("example_value")
 *             .dispatchType("CREATE_ALERT")
 *             .groupRules(DispatchRuleGroupRuleArgs.builder()
 *                 .groupWaitTime(5)
 *                 .groupInterval(15)
 *                 .repeatInterval(100)
 *                 .groupingFields("alertname")
 *                 .build())
 *             .labelMatchExpressionGrids(DispatchRuleLabelMatchExpressionGridArgs.builder()
 *                 .labelMatchExpressionGroups(DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupArgs.builder()
 *                     .labelMatchExpressions(DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpressionArgs.builder()
 *                         .key("_aliyun_arms_involvedObject_kind")
 *                         .value("app")
 *                         .operator("eq")
 *                         .build())
 *                     .build())
 *                 .build())
 *             .notifyRules(DispatchRuleNotifyRuleArgs.builder()
 *                 .notifyObjects(                
 *                     DispatchRuleNotifyRuleNotifyObjectArgs.builder()
 *                         .notifyObjectId(default_.id())
 *                         .notifyType("ARMS_CONTACT")
 *                         .name("example_value")
 *                         .build(),
 *                     DispatchRuleNotifyRuleNotifyObjectArgs.builder()
 *                         .notifyObjectId(defaultAlertContactGroup.id())
 *                         .notifyType("ARMS_CONTACT_GROUP")
 *                         .name("example_value")
 *                         .build())
 *                 .notifyChannels(                
 *                     "dingTalk",
 *                     "wechat")
 *                 .build())
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
 * Application Real-Time Monitoring Service (ARMS) Alert Contact can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:arms/dispatchRule:DispatchRule example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:arms/dispatchRule:DispatchRule")
public class DispatchRule extends com.pulumi.resources.CustomResource {
    /**
     * The name of the dispatch policy.
     * 
     */
    @Export(name="dispatchRuleName", refs={String.class}, tree="[0]")
    private Output<String> dispatchRuleName;

    /**
     * @return The name of the dispatch policy.
     * 
     */
    public Output<String> dispatchRuleName() {
        return this.dispatchRuleName;
    }
    /**
     * The alert handling method. Valid values: CREATE_ALERT: generates an alert. DISCARD_ALERT: discards the alert event and generates no alert.
     * 
     */
    @Export(name="dispatchType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dispatchType;

    /**
     * @return The alert handling method. Valid values: CREATE_ALERT: generates an alert. DISCARD_ALERT: discards the alert event and generates no alert.
     * 
     */
    public Output<Optional<String>> dispatchType() {
        return Codegen.optional(this.dispatchType);
    }
    /**
     * Sets the event group. See `group_rules` below. It will be ignored  when `dispatch_type = &#34;DISCARD_ALERT&#34;`.
     * 
     */
    @Export(name="groupRules", refs={List.class,DispatchRuleGroupRule.class}, tree="[0,1]")
    private Output<List<DispatchRuleGroupRule>> groupRules;

    /**
     * @return Sets the event group. See `group_rules` below. It will be ignored  when `dispatch_type = &#34;DISCARD_ALERT&#34;`.
     * 
     */
    public Output<List<DispatchRuleGroupRule>> groupRules() {
        return this.groupRules;
    }
    /**
     * Specifies whether to send the restored alert. Valid values: true: sends the alert. false: does not send the alert.
     * 
     */
    @Export(name="isRecover", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isRecover;

    /**
     * @return Specifies whether to send the restored alert. Valid values: true: sends the alert. false: does not send the alert.
     * 
     */
    public Output<Optional<Boolean>> isRecover() {
        return Codegen.optional(this.isRecover);
    }
    /**
     * Sets the dispatch rule. See `label_match_expression_grid` below.
     * 
     */
    @Export(name="labelMatchExpressionGrids", refs={List.class,DispatchRuleLabelMatchExpressionGrid.class}, tree="[0,1]")
    private Output<List<DispatchRuleLabelMatchExpressionGrid>> labelMatchExpressionGrids;

    /**
     * @return Sets the dispatch rule. See `label_match_expression_grid` below.
     * 
     */
    public Output<List<DispatchRuleLabelMatchExpressionGrid>> labelMatchExpressionGrids() {
        return this.labelMatchExpressionGrids;
    }
    /**
     * Sets the notification rule. See `notify_rules` below. It will be ignored  when `dispatch_type = &#34;DISCARD_ALERT&#34;`.
     * 
     */
    @Export(name="notifyRules", refs={List.class,DispatchRuleNotifyRule.class}, tree="[0,1]")
    private Output<List<DispatchRuleNotifyRule>> notifyRules;

    /**
     * @return Sets the notification rule. See `notify_rules` below. It will be ignored  when `dispatch_type = &#34;DISCARD_ALERT&#34;`.
     * 
     */
    public Output<List<DispatchRuleNotifyRule>> notifyRules() {
        return this.notifyRules;
    }
    /**
     * The resource status of Alert Dispatch Rule.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The resource status of Alert Dispatch Rule.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DispatchRule(String name) {
        this(name, DispatchRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DispatchRule(String name, DispatchRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DispatchRule(String name, DispatchRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:arms/dispatchRule:DispatchRule", name, args == null ? DispatchRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DispatchRule(String name, Output<String> id, @Nullable DispatchRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:arms/dispatchRule:DispatchRule", name, state, makeResourceOptions(options, id));
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
    public static DispatchRule get(String name, Output<String> id, @Nullable DispatchRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DispatchRule(name, id, state, options);
    }
}
