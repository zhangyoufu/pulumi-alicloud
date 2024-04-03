// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServerlessKubernetesAddon {
    /**
     * @return The ACK add-on configurations. For more config information, see cs_kubernetes_addon_metadata.
     * 
     */
    private @Nullable String config;
    /**
     * @return Disables the automatic installation of a component. Default is `false`.
     * 
     * The following example is the definition of addons block, The type of this field is list:
     * 
     */
    private @Nullable Boolean disabled;
    /**
     * @return Name of the ACK add-on. The name must match one of the names returned by [DescribeAddons](https://help.aliyun.com/document_detail/171524.html).
     * 
     */
    private @Nullable String name;

    private ServerlessKubernetesAddon() {}
    /**
     * @return The ACK add-on configurations. For more config information, see cs_kubernetes_addon_metadata.
     * 
     */
    public Optional<String> config() {
        return Optional.ofNullable(this.config);
    }
    /**
     * @return Disables the automatic installation of a component. Default is `false`.
     * 
     * The following example is the definition of addons block, The type of this field is list:
     * 
     */
    public Optional<Boolean> disabled() {
        return Optional.ofNullable(this.disabled);
    }
    /**
     * @return Name of the ACK add-on. The name must match one of the names returned by [DescribeAddons](https://help.aliyun.com/document_detail/171524.html).
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServerlessKubernetesAddon defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String config;
        private @Nullable Boolean disabled;
        private @Nullable String name;
        public Builder() {}
        public Builder(ServerlessKubernetesAddon defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.config = defaults.config;
    	      this.disabled = defaults.disabled;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder config(@Nullable String config) {

            this.config = config;
            return this;
        }
        @CustomType.Setter
        public Builder disabled(@Nullable Boolean disabled) {

            this.disabled = disabled;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        public ServerlessKubernetesAddon build() {
            final var _resultValue = new ServerlessKubernetesAddon();
            _resultValue.config = config;
            _resultValue.disabled = disabled;
            _resultValue.name = name;
            return _resultValue;
        }
    }
}
