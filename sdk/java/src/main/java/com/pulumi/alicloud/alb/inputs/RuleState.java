// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.inputs;

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


public final class RuleState extends com.pulumi.resources.ResourceArgs {

    public static final RuleState Empty = new RuleState();

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
    @Import(name="listenerId")
    private @Nullable Output<String> listenerId;

    /**
     * @return The ID of the listener to which the forwarding rule belongs.
     * 
     */
    public Optional<Output<String>> listenerId() {
        return Optional.ofNullable(this.listenerId);
    }

    /**
     * The priority of the rule. Valid values: 1 to 10000. A smaller value indicates a higher priority. **Note*:* The priority of each rule within the same listener must be unique.
     * 
     */
    @Import(name="priority")
    private @Nullable Output<Integer> priority;

    /**
     * @return The priority of the rule. Valid values: 1 to 10000. A smaller value indicates a higher priority. **Note*:* The priority of each rule within the same listener must be unique.
     * 
     */
    public Optional<Output<Integer>> priority() {
        return Optional.ofNullable(this.priority);
    }

    /**
     * The actions of the forwarding rules. See the following `Block rule_actions`.
     * 
     */
    @Import(name="ruleActions")
    private @Nullable Output<List<RuleRuleActionArgs>> ruleActions;

    /**
     * @return The actions of the forwarding rules. See the following `Block rule_actions`.
     * 
     */
    public Optional<Output<List<RuleRuleActionArgs>>> ruleActions() {
        return Optional.ofNullable(this.ruleActions);
    }

    /**
     * The conditions of the forwarding rule. See the following `Block rule_conditions`.
     * 
     */
    @Import(name="ruleConditions")
    private @Nullable Output<List<RuleRuleConditionArgs>> ruleConditions;

    /**
     * @return The conditions of the forwarding rule. See the following `Block rule_conditions`.
     * 
     */
    public Optional<Output<List<RuleRuleConditionArgs>>> ruleConditions() {
        return Optional.ofNullable(this.ruleConditions);
    }

    /**
     * The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    @Import(name="ruleName")
    private @Nullable Output<String> ruleName;

    /**
     * @return The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    public Optional<Output<String>> ruleName() {
        return Optional.ofNullable(this.ruleName);
    }

    /**
     * The status of the resource.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private RuleState() {}

    private RuleState(RuleState $) {
        this.direction = $.direction;
        this.dryRun = $.dryRun;
        this.listenerId = $.listenerId;
        this.priority = $.priority;
        this.ruleActions = $.ruleActions;
        this.ruleConditions = $.ruleConditions;
        this.ruleName = $.ruleName;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuleState $;

        public Builder() {
            $ = new RuleState();
        }

        public Builder(RuleState defaults) {
            $ = new RuleState(Objects.requireNonNull(defaults));
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
        public Builder listenerId(@Nullable Output<String> listenerId) {
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
        public Builder priority(@Nullable Output<Integer> priority) {
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
        public Builder ruleActions(@Nullable Output<List<RuleRuleActionArgs>> ruleActions) {
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
        public Builder ruleConditions(@Nullable Output<List<RuleRuleConditionArgs>> ruleConditions) {
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
        public Builder ruleName(@Nullable Output<String> ruleName) {
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

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public RuleState build() {
            return $;
        }
    }

}
