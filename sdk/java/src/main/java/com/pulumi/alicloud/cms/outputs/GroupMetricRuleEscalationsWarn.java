// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GroupMetricRuleEscalationsWarn {
    /**
     * @return The comparison operator of the threshold for critical-level alerts.
     * 
     */
    private @Nullable String comparisonOperator;
    /**
     * @return The statistical aggregation method for critical-level alerts.
     * 
     */
    private @Nullable String statistics;
    /**
     * @return The threshold for critical-level alerts.
     * 
     */
    private @Nullable String threshold;
    /**
     * @return The consecutive number of times for which the metric value is measured before a critical-level alert is triggered.
     * 
     */
    private @Nullable Integer times;

    private GroupMetricRuleEscalationsWarn() {}
    /**
     * @return The comparison operator of the threshold for critical-level alerts.
     * 
     */
    public Optional<String> comparisonOperator() {
        return Optional.ofNullable(this.comparisonOperator);
    }
    /**
     * @return The statistical aggregation method for critical-level alerts.
     * 
     */
    public Optional<String> statistics() {
        return Optional.ofNullable(this.statistics);
    }
    /**
     * @return The threshold for critical-level alerts.
     * 
     */
    public Optional<String> threshold() {
        return Optional.ofNullable(this.threshold);
    }
    /**
     * @return The consecutive number of times for which the metric value is measured before a critical-level alert is triggered.
     * 
     */
    public Optional<Integer> times() {
        return Optional.ofNullable(this.times);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GroupMetricRuleEscalationsWarn defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String comparisonOperator;
        private @Nullable String statistics;
        private @Nullable String threshold;
        private @Nullable Integer times;
        public Builder() {}
        public Builder(GroupMetricRuleEscalationsWarn defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.comparisonOperator = defaults.comparisonOperator;
    	      this.statistics = defaults.statistics;
    	      this.threshold = defaults.threshold;
    	      this.times = defaults.times;
        }

        @CustomType.Setter
        public Builder comparisonOperator(@Nullable String comparisonOperator) {
            this.comparisonOperator = comparisonOperator;
            return this;
        }
        @CustomType.Setter
        public Builder statistics(@Nullable String statistics) {
            this.statistics = statistics;
            return this;
        }
        @CustomType.Setter
        public Builder threshold(@Nullable String threshold) {
            this.threshold = threshold;
            return this;
        }
        @CustomType.Setter
        public Builder times(@Nullable Integer times) {
            this.times = times;
            return this;
        }
        public GroupMetricRuleEscalationsWarn build() {
            final var o = new GroupMetricRuleEscalationsWarn();
            o.comparisonOperator = comparisonOperator;
            o.statistics = statistics;
            o.threshold = threshold;
            o.times = times;
            return o;
        }
    }
}
