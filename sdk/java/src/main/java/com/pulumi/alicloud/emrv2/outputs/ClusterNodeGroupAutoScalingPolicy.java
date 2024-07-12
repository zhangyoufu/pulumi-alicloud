// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.emrv2.outputs;

import com.pulumi.alicloud.emrv2.outputs.ClusterNodeGroupAutoScalingPolicyConstraints;
import com.pulumi.alicloud.emrv2.outputs.ClusterNodeGroupAutoScalingPolicyScalingRule;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClusterNodeGroupAutoScalingPolicy {
    private @Nullable ClusterNodeGroupAutoScalingPolicyConstraints constraints;
    private @Nullable List<ClusterNodeGroupAutoScalingPolicyScalingRule> scalingRules;

    private ClusterNodeGroupAutoScalingPolicy() {}
    public Optional<ClusterNodeGroupAutoScalingPolicyConstraints> constraints() {
        return Optional.ofNullable(this.constraints);
    }
    public List<ClusterNodeGroupAutoScalingPolicyScalingRule> scalingRules() {
        return this.scalingRules == null ? List.of() : this.scalingRules;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterNodeGroupAutoScalingPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable ClusterNodeGroupAutoScalingPolicyConstraints constraints;
        private @Nullable List<ClusterNodeGroupAutoScalingPolicyScalingRule> scalingRules;
        public Builder() {}
        public Builder(ClusterNodeGroupAutoScalingPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.constraints = defaults.constraints;
    	      this.scalingRules = defaults.scalingRules;
        }

        @CustomType.Setter
        public Builder constraints(@Nullable ClusterNodeGroupAutoScalingPolicyConstraints constraints) {

            this.constraints = constraints;
            return this;
        }
        @CustomType.Setter
        public Builder scalingRules(@Nullable List<ClusterNodeGroupAutoScalingPolicyScalingRule> scalingRules) {

            this.scalingRules = scalingRules;
            return this;
        }
        public Builder scalingRules(ClusterNodeGroupAutoScalingPolicyScalingRule... scalingRules) {
            return scalingRules(List.of(scalingRules));
        }
        public ClusterNodeGroupAutoScalingPolicy build() {
            final var _resultValue = new ClusterNodeGroupAutoScalingPolicy();
            _resultValue.constraints = constraints;
            _resultValue.scalingRules = scalingRules;
            return _resultValue;
        }
    }
}
