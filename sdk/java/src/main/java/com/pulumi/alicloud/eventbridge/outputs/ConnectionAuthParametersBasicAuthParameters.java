// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eventbridge.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ConnectionAuthParametersBasicAuthParameters {
    /**
     * @return The password for basic authentication.
     * 
     */
    private @Nullable String password;
    /**
     * @return The username for basic authentication.
     * 
     */
    private @Nullable String username;

    private ConnectionAuthParametersBasicAuthParameters() {}
    /**
     * @return The password for basic authentication.
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    /**
     * @return The username for basic authentication.
     * 
     */
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ConnectionAuthParametersBasicAuthParameters defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String password;
        private @Nullable String username;
        public Builder() {}
        public Builder(ConnectionAuthParametersBasicAuthParameters defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.password = defaults.password;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder password(@Nullable String password) {
            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder username(@Nullable String username) {
            this.username = username;
            return this;
        }
        public ConnectionAuthParametersBasicAuthParameters build() {
            final var o = new ConnectionAuthParametersBasicAuthParameters();
            o.password = password;
            o.username = username;
            return o;
        }
    }
}
