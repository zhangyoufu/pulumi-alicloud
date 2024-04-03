// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EdgeKubernetesAddonArgs extends com.pulumi.resources.ResourceArgs {

    public static final EdgeKubernetesAddonArgs Empty = new EdgeKubernetesAddonArgs();

    /**
     * The ACK add-on configurations. For more config information, see cs_kubernetes_addon_metadata.
     * 
     */
    @Import(name="config")
    private @Nullable Output<String> config;

    /**
     * @return The ACK add-on configurations. For more config information, see cs_kubernetes_addon_metadata.
     * 
     */
    public Optional<Output<String>> config() {
        return Optional.ofNullable(this.config);
    }

    /**
     * Disables the automatic installation of a component. Default is `false`.
     * 
     * The following example is the definition of addons block, The type of this field is list:
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return Disables the automatic installation of a component. Default is `false`.
     * 
     * The following example is the definition of addons block, The type of this field is list:
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * Name of the ACK add-on. The name must match one of the names returned by [DescribeAddons](https://help.aliyun.com/document_detail/171524.html).
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the ACK add-on. The name must match one of the names returned by [DescribeAddons](https://help.aliyun.com/document_detail/171524.html).
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * It specifies the version of the component.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return It specifies the version of the component.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private EdgeKubernetesAddonArgs() {}

    private EdgeKubernetesAddonArgs(EdgeKubernetesAddonArgs $) {
        this.config = $.config;
        this.disabled = $.disabled;
        this.name = $.name;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EdgeKubernetesAddonArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EdgeKubernetesAddonArgs $;

        public Builder() {
            $ = new EdgeKubernetesAddonArgs();
        }

        public Builder(EdgeKubernetesAddonArgs defaults) {
            $ = new EdgeKubernetesAddonArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param config The ACK add-on configurations. For more config information, see cs_kubernetes_addon_metadata.
         * 
         * @return builder
         * 
         */
        public Builder config(@Nullable Output<String> config) {
            $.config = config;
            return this;
        }

        /**
         * @param config The ACK add-on configurations. For more config information, see cs_kubernetes_addon_metadata.
         * 
         * @return builder
         * 
         */
        public Builder config(String config) {
            return config(Output.of(config));
        }

        /**
         * @param disabled Disables the automatic installation of a component. Default is `false`.
         * 
         * The following example is the definition of addons block, The type of this field is list:
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled Disables the automatic installation of a component. Default is `false`.
         * 
         * The following example is the definition of addons block, The type of this field is list:
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param name Name of the ACK add-on. The name must match one of the names returned by [DescribeAddons](https://help.aliyun.com/document_detail/171524.html).
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the ACK add-on. The name must match one of the names returned by [DescribeAddons](https://help.aliyun.com/document_detail/171524.html).
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param version It specifies the version of the component.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version It specifies the version of the component.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public EdgeKubernetesAddonArgs build() {
            return $;
        }
    }

}
