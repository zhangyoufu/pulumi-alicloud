// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.alicloud.ga.outputs.ForwardingRuleRuleActionForwardGroupConfig;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ForwardingRuleRuleAction {
    /**
     * @return Forwarding configuration. See `forward_group_config` below.
     * &gt; **NOTE:** From version 1.207.0, We recommend that you do not use `forward_group_config`, and we recommend that you use the `rule_action_type` and `rule_action_value` to configure forwarding actions.
     * 
     */
    private @Nullable ForwardingRuleRuleActionForwardGroupConfig forwardGroupConfig;
    /**
     * @return Forwarding priority.
     * 
     */
    private Integer order;
    /**
     * @return Forward action type.
     * 
     */
    private String ruleActionType;
    /**
     * @return The value of the forwarding action type. For more information, see [How to use it](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-doc-ga-2019-11-20-api-doc-createforwardingrules).
     * 
     */
    private @Nullable String ruleActionValue;

    private ForwardingRuleRuleAction() {}
    /**
     * @return Forwarding configuration. See `forward_group_config` below.
     * &gt; **NOTE:** From version 1.207.0, We recommend that you do not use `forward_group_config`, and we recommend that you use the `rule_action_type` and `rule_action_value` to configure forwarding actions.
     * 
     */
    public Optional<ForwardingRuleRuleActionForwardGroupConfig> forwardGroupConfig() {
        return Optional.ofNullable(this.forwardGroupConfig);
    }
    /**
     * @return Forwarding priority.
     * 
     */
    public Integer order() {
        return this.order;
    }
    /**
     * @return Forward action type.
     * 
     */
    public String ruleActionType() {
        return this.ruleActionType;
    }
    /**
     * @return The value of the forwarding action type. For more information, see [How to use it](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-doc-ga-2019-11-20-api-doc-createforwardingrules).
     * 
     */
    public Optional<String> ruleActionValue() {
        return Optional.ofNullable(this.ruleActionValue);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ForwardingRuleRuleAction defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable ForwardingRuleRuleActionForwardGroupConfig forwardGroupConfig;
        private Integer order;
        private String ruleActionType;
        private @Nullable String ruleActionValue;
        public Builder() {}
        public Builder(ForwardingRuleRuleAction defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.forwardGroupConfig = defaults.forwardGroupConfig;
    	      this.order = defaults.order;
    	      this.ruleActionType = defaults.ruleActionType;
    	      this.ruleActionValue = defaults.ruleActionValue;
        }

        @CustomType.Setter
        public Builder forwardGroupConfig(@Nullable ForwardingRuleRuleActionForwardGroupConfig forwardGroupConfig) {
            this.forwardGroupConfig = forwardGroupConfig;
            return this;
        }
        @CustomType.Setter
        public Builder order(Integer order) {
            this.order = Objects.requireNonNull(order);
            return this;
        }
        @CustomType.Setter
        public Builder ruleActionType(String ruleActionType) {
            this.ruleActionType = Objects.requireNonNull(ruleActionType);
            return this;
        }
        @CustomType.Setter
        public Builder ruleActionValue(@Nullable String ruleActionValue) {
            this.ruleActionValue = ruleActionValue;
            return this;
        }
        public ForwardingRuleRuleAction build() {
            final var o = new ForwardingRuleRuleAction();
            o.forwardGroupConfig = forwardGroupConfig;
            o.order = order;
            o.ruleActionType = ruleActionType;
            o.ruleActionValue = ruleActionValue;
            return o;
        }
    }
}
