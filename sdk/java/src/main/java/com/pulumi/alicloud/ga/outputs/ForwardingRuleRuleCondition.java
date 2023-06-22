// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.alicloud.ga.outputs.ForwardingRuleRuleConditionHostConfig;
import com.pulumi.alicloud.ga.outputs.ForwardingRuleRuleConditionPathConfig;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ForwardingRuleRuleCondition {
    /**
     * @return The configuration of the domain name. See `host_config` below.
     * 
     */
    private @Nullable List<ForwardingRuleRuleConditionHostConfig> hostConfigs;
    /**
     * @return The configuration of the path. See `path_config` below.
     * 
     */
    private @Nullable ForwardingRuleRuleConditionPathConfig pathConfig;
    /**
     * @return The type of the forwarding conditions. Valid values: `Host`, `Path`.
     * 
     */
    private String ruleConditionType;

    private ForwardingRuleRuleCondition() {}
    /**
     * @return The configuration of the domain name. See `host_config` below.
     * 
     */
    public List<ForwardingRuleRuleConditionHostConfig> hostConfigs() {
        return this.hostConfigs == null ? List.of() : this.hostConfigs;
    }
    /**
     * @return The configuration of the path. See `path_config` below.
     * 
     */
    public Optional<ForwardingRuleRuleConditionPathConfig> pathConfig() {
        return Optional.ofNullable(this.pathConfig);
    }
    /**
     * @return The type of the forwarding conditions. Valid values: `Host`, `Path`.
     * 
     */
    public String ruleConditionType() {
        return this.ruleConditionType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ForwardingRuleRuleCondition defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<ForwardingRuleRuleConditionHostConfig> hostConfigs;
        private @Nullable ForwardingRuleRuleConditionPathConfig pathConfig;
        private String ruleConditionType;
        public Builder() {}
        public Builder(ForwardingRuleRuleCondition defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.hostConfigs = defaults.hostConfigs;
    	      this.pathConfig = defaults.pathConfig;
    	      this.ruleConditionType = defaults.ruleConditionType;
        }

        @CustomType.Setter
        public Builder hostConfigs(@Nullable List<ForwardingRuleRuleConditionHostConfig> hostConfigs) {
            this.hostConfigs = hostConfigs;
            return this;
        }
        public Builder hostConfigs(ForwardingRuleRuleConditionHostConfig... hostConfigs) {
            return hostConfigs(List.of(hostConfigs));
        }
        @CustomType.Setter
        public Builder pathConfig(@Nullable ForwardingRuleRuleConditionPathConfig pathConfig) {
            this.pathConfig = pathConfig;
            return this;
        }
        @CustomType.Setter
        public Builder ruleConditionType(String ruleConditionType) {
            this.ruleConditionType = Objects.requireNonNull(ruleConditionType);
            return this;
        }
        public ForwardingRuleRuleCondition build() {
            final var o = new ForwardingRuleRuleCondition();
            o.hostConfigs = hostConfigs;
            o.pathConfig = pathConfig;
            o.ruleConditionType = ruleConditionType;
            return o;
        }
    }
}
