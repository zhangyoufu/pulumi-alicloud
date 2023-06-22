// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ForwardingRuleRuleActionForwardGroupConfigServerGroupTuple {
    /**
     * @return The ID of the endpoint group.
     * 
     */
    private String endpointGroupId;

    private ForwardingRuleRuleActionForwardGroupConfigServerGroupTuple() {}
    /**
     * @return The ID of the endpoint group.
     * 
     */
    public String endpointGroupId() {
        return this.endpointGroupId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ForwardingRuleRuleActionForwardGroupConfigServerGroupTuple defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String endpointGroupId;
        public Builder() {}
        public Builder(ForwardingRuleRuleActionForwardGroupConfigServerGroupTuple defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.endpointGroupId = defaults.endpointGroupId;
        }

        @CustomType.Setter
        public Builder endpointGroupId(String endpointGroupId) {
            this.endpointGroupId = Objects.requireNonNull(endpointGroupId);
            return this;
        }
        public ForwardingRuleRuleActionForwardGroupConfigServerGroupTuple build() {
            final var o = new ForwardingRuleRuleActionForwardGroupConfigServerGroupTuple();
            o.endpointGroupId = endpointGroupId;
            return o;
        }
    }
}
