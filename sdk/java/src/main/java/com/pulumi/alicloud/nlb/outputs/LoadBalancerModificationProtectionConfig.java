// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nlb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class LoadBalancerModificationProtectionConfig {
    /**
     * @return Opening time.
     * 
     */
    private @Nullable String enabledTime;
    /**
     * @return Reason for opening.
     * 
     */
    private @Nullable String reason;
    /**
     * @return ON.
     * 
     */
    private @Nullable String status;

    private LoadBalancerModificationProtectionConfig() {}
    /**
     * @return Opening time.
     * 
     */
    public Optional<String> enabledTime() {
        return Optional.ofNullable(this.enabledTime);
    }
    /**
     * @return Reason for opening.
     * 
     */
    public Optional<String> reason() {
        return Optional.ofNullable(this.reason);
    }
    /**
     * @return ON.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(LoadBalancerModificationProtectionConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String enabledTime;
        private @Nullable String reason;
        private @Nullable String status;
        public Builder() {}
        public Builder(LoadBalancerModificationProtectionConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabledTime = defaults.enabledTime;
    	      this.reason = defaults.reason;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder enabledTime(@Nullable String enabledTime) {

            this.enabledTime = enabledTime;
            return this;
        }
        @CustomType.Setter
        public Builder reason(@Nullable String reason) {

            this.reason = reason;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {

            this.status = status;
            return this;
        }
        public LoadBalancerModificationProtectionConfig build() {
            final var _resultValue = new LoadBalancerModificationProtectionConfig();
            _resultValue.enabledTime = enabledTime;
            _resultValue.reason = reason;
            _resultValue.status = status;
            return _resultValue;
        }
    }
}
