// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb;

import com.pulumi.alicloud.alb.inputs.RuleRuleActionArgs;
import com.pulumi.alicloud.alb.inputs.RuleRuleConditionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final RuleArgs Empty = new RuleArgs();

    /**
     * The direction to which the forwarding rule is applied. Default value: `Request`. Valid values:
     * 
     */
    @Import(name="direction")
    private @Nullable Output<String> direction;

    /**
     * @return The direction to which the forwarding rule is applied. Default value: `Request`. Valid values:
     * 
     */
    public Optional<Output<String>> direction() {
        return Optional.ofNullable(this.direction);
    }

    /**
     * Specifies whether to precheck this request.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return Specifies whether to precheck this request.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * The ID of the listener to which the forwarding rule belongs.
     * 
     */
    @Import(name="listenerId", required=true)
    private Output<String> listenerId;

    /**
     * @return The ID of the listener to which the forwarding rule belongs.
     * 
     */
    public Output<String> listenerId() {
        return this.listenerId;
    }

    /**
     * The priority of the rule. Valid values: 1 to 10000. A smaller value indicates a higher priority. **Note*:* The priority of each rule within the same listener must be unique.
     * 
     */
    @Import(name="priority", required=true)
    private Output<Integer> priority;

    /**
     * @return The priority of the rule. Valid values: 1 to 10000. A smaller value indicates a higher priority. **Note*:* The priority of each rule within the same listener must be unique.
     * 
     */
    public Output<Integer> priority() {
        return this.priority;
    }

    /**
     * The actions of the forwarding rules. See the following `Block rule_actions`.
     * 
     */
    @Import(name="ruleActions", required=true)
    private Output<List<RuleRuleActionArgs>> ruleActions;

    /**
     * @return The actions of the forwarding rules. See the following `Block rule_actions`.
     * 
     */
    public Output<List<RuleRuleActionArgs>> ruleActions() {
        return this.ruleActions;
    }

    /**
     * The conditions of the forwarding rule. See the following `Block rule_conditions`.
     * 
     */
    @Import(name="ruleConditions", required=true)
    private Output<List<RuleRuleConditionArgs>> ruleConditions;

    /**
     * @return The conditions of the forwarding rule. See the following `Block rule_conditions`.
     * 
     */
    public Output<List<RuleRuleConditionArgs>> ruleConditions() {
        return this.ruleConditions;
    }

    /**
     * The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    @Import(name="ruleName", required=true)
    private Output<String> ruleName;

    /**
     * @return The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    public Output<String> ruleName() {
        return this.ruleName;
    }

    private RuleArgs() {}

    private RuleArgs(RuleArgs $) {
        this.direction = $.direction;
        this.dryRun = $.dryRun;
        this.listenerId = $.listenerId;
        this.priority = $.priority;
        this.ruleActions = $.ruleActions;
        this.ruleConditions = $.ruleConditions;
        this.ruleName = $.ruleName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuleArgs $;

        public Builder() {
            $ = new RuleArgs();
        }

        public Builder(RuleArgs defaults) {
            $ = new RuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param direction The direction to which the forwarding rule is applied. Default value: `Request`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder direction(@Nullable Output<String> direction) {
            $.direction = direction;
            return this;
        }

        /**
         * @param direction The direction to which the forwarding rule is applied. Default value: `Request`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder direction(String direction) {
            return direction(Output.of(direction));
        }

        /**
         * @param dryRun Specifies whether to precheck this request.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun Specifies whether to precheck this request.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param listenerId The ID of the listener to which the forwarding rule belongs.
         * 
         * @return builder
         * 
         */
        public Builder listenerId(Output<String> listenerId) {
            $.listenerId = listenerId;
            return this;
        }

        /**
         * @param listenerId The ID of the listener to which the forwarding rule belongs.
         * 
         * @return builder
         * 
         */
        public Builder listenerId(String listenerId) {
            return listenerId(Output.of(listenerId));
        }

        /**
         * @param priority The priority of the rule. Valid values: 1 to 10000. A smaller value indicates a higher priority. **Note*:* The priority of each rule within the same listener must be unique.
         * 
         * @return builder
         * 
         */
        public Builder priority(Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority The priority of the rule. Valid values: 1 to 10000. A smaller value indicates a higher priority. **Note*:* The priority of each rule within the same listener must be unique.
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
        }

        /**
         * @param ruleActions The actions of the forwarding rules. See the following `Block rule_actions`.
         * 
         * @return builder
         * 
         */
        public Builder ruleActions(Output<List<RuleRuleActionArgs>> ruleActions) {
            $.ruleActions = ruleActions;
            return this;
        }

        /**
         * @param ruleActions The actions of the forwarding rules. See the following `Block rule_actions`.
         * 
         * @return builder
         * 
         */
        public Builder ruleActions(List<RuleRuleActionArgs> ruleActions) {
            return ruleActions(Output.of(ruleActions));
        }

        /**
         * @param ruleActions The actions of the forwarding rules. See the following `Block rule_actions`.
         * 
         * @return builder
         * 
         */
        public Builder ruleActions(RuleRuleActionArgs... ruleActions) {
            return ruleActions(List.of(ruleActions));
        }

        /**
         * @param ruleConditions The conditions of the forwarding rule. See the following `Block rule_conditions`.
         * 
         * @return builder
         * 
         */
        public Builder ruleConditions(Output<List<RuleRuleConditionArgs>> ruleConditions) {
            $.ruleConditions = ruleConditions;
            return this;
        }

        /**
         * @param ruleConditions The conditions of the forwarding rule. See the following `Block rule_conditions`.
         * 
         * @return builder
         * 
         */
        public Builder ruleConditions(List<RuleRuleConditionArgs> ruleConditions) {
            return ruleConditions(Output.of(ruleConditions));
        }

        /**
         * @param ruleConditions The conditions of the forwarding rule. See the following `Block rule_conditions`.
         * 
         * @return builder
         * 
         */
        public Builder ruleConditions(RuleRuleConditionArgs... ruleConditions) {
            return ruleConditions(List.of(ruleConditions));
        }

        /**
         * @param ruleName The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
         * 
         * @return builder
         * 
         */
        public Builder ruleName(Output<String> ruleName) {
            $.ruleName = ruleName;
            return this;
        }

        /**
         * @param ruleName The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
         * 
         * @return builder
         * 
         */
        public Builder ruleName(String ruleName) {
            return ruleName(Output.of(ruleName));
        }

        public RuleArgs build() {
            $.listenerId = Objects.requireNonNull($.listenerId, "expected parameter 'listenerId' to be non-null");
            $.priority = Objects.requireNonNull($.priority, "expected parameter 'priority' to be non-null");
            $.ruleActions = Objects.requireNonNull($.ruleActions, "expected parameter 'ruleActions' to be non-null");
            $.ruleConditions = Objects.requireNonNull($.ruleConditions, "expected parameter 'ruleConditions' to be non-null");
            $.ruleName = Objects.requireNonNull($.ruleName, "expected parameter 'ruleName' to be non-null");
            return $;
        }
    }

}
