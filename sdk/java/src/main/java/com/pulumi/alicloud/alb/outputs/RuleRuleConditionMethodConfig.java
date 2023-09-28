// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class RuleRuleConditionMethodConfig {
    /**
     * @return The values of the cookie. See `values` below.
     * 
     */
    private @Nullable List<String> values;

    private RuleRuleConditionMethodConfig() {}
    /**
     * @return The values of the cookie. See `values` below.
     * 
     */
    public List<String> values() {
        return this.values == null ? List.of() : this.values;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RuleRuleConditionMethodConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> values;
        public Builder() {}
        public Builder(RuleRuleConditionMethodConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.values = defaults.values;
        }

        @CustomType.Setter
        public Builder values(@Nullable List<String> values) {
            this.values = values;
            return this;
        }
        public Builder values(String... values) {
            return values(List.of(values));
        }
        public RuleRuleConditionMethodConfig build() {
            final var o = new RuleRuleConditionMethodConfig();
            o.values = values;
            return o;
        }
    }
}
