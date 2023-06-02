// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NamespaceArgs extends com.pulumi.resources.ResourceArgs {

    public static final NamespaceArgs Empty = new NamespaceArgs();

    /**
     * Specifies whether to enable the SAE built-in registry. If you do not use the built-in registry, you can set `enable_micro_registration` to `false` to accelerate the creation of the namespace. Default value: `true`. Valid values:
     * 
     */
    @Import(name="enableMicroRegistration")
    private @Nullable Output<Boolean> enableMicroRegistration;

    /**
     * @return Specifies whether to enable the SAE built-in registry. If you do not use the built-in registry, you can set `enable_micro_registration` to `false` to accelerate the creation of the namespace. Default value: `true`. Valid values:
     * 
     */
    public Optional<Output<Boolean>> enableMicroRegistration() {
        return Optional.ofNullable(this.enableMicroRegistration);
    }

    /**
     * The Description of Namespace.
     * 
     */
    @Import(name="namespaceDescription")
    private @Nullable Output<String> namespaceDescription;

    /**
     * @return The Description of Namespace.
     * 
     */
    public Optional<Output<String>> namespaceDescription() {
        return Optional.ofNullable(this.namespaceDescription);
    }

    /**
     * The ID of the Namespace. It can contain 2 to 32 lowercase characters. The value is in format `{RegionId}:{namespace}`.
     * 
     */
    @Import(name="namespaceId")
    private @Nullable Output<String> namespaceId;

    /**
     * @return The ID of the Namespace. It can contain 2 to 32 lowercase characters. The value is in format `{RegionId}:{namespace}`.
     * 
     */
    public Optional<Output<String>> namespaceId() {
        return Optional.ofNullable(this.namespaceId);
    }

    /**
     * The Name of Namespace.
     * 
     */
    @Import(name="namespaceName", required=true)
    private Output<String> namespaceName;

    /**
     * @return The Name of Namespace.
     * 
     */
    public Output<String> namespaceName() {
        return this.namespaceName;
    }

    /**
     * The short ID of the Namespace. You do not need to specify a region ID. The value of `namespace_short_id` can be up to 20 characters in length and can contain only lowercase letters and digits.
     * 
     */
    @Import(name="namespaceShortId")
    private @Nullable Output<String> namespaceShortId;

    /**
     * @return The short ID of the Namespace. You do not need to specify a region ID. The value of `namespace_short_id` can be up to 20 characters in length and can contain only lowercase letters and digits.
     * 
     */
    public Optional<Output<String>> namespaceShortId() {
        return Optional.ofNullable(this.namespaceShortId);
    }

    private NamespaceArgs() {}

    private NamespaceArgs(NamespaceArgs $) {
        this.enableMicroRegistration = $.enableMicroRegistration;
        this.namespaceDescription = $.namespaceDescription;
        this.namespaceId = $.namespaceId;
        this.namespaceName = $.namespaceName;
        this.namespaceShortId = $.namespaceShortId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NamespaceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NamespaceArgs $;

        public Builder() {
            $ = new NamespaceArgs();
        }

        public Builder(NamespaceArgs defaults) {
            $ = new NamespaceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enableMicroRegistration Specifies whether to enable the SAE built-in registry. If you do not use the built-in registry, you can set `enable_micro_registration` to `false` to accelerate the creation of the namespace. Default value: `true`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder enableMicroRegistration(@Nullable Output<Boolean> enableMicroRegistration) {
            $.enableMicroRegistration = enableMicroRegistration;
            return this;
        }

        /**
         * @param enableMicroRegistration Specifies whether to enable the SAE built-in registry. If you do not use the built-in registry, you can set `enable_micro_registration` to `false` to accelerate the creation of the namespace. Default value: `true`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder enableMicroRegistration(Boolean enableMicroRegistration) {
            return enableMicroRegistration(Output.of(enableMicroRegistration));
        }

        /**
         * @param namespaceDescription The Description of Namespace.
         * 
         * @return builder
         * 
         */
        public Builder namespaceDescription(@Nullable Output<String> namespaceDescription) {
            $.namespaceDescription = namespaceDescription;
            return this;
        }

        /**
         * @param namespaceDescription The Description of Namespace.
         * 
         * @return builder
         * 
         */
        public Builder namespaceDescription(String namespaceDescription) {
            return namespaceDescription(Output.of(namespaceDescription));
        }

        /**
         * @param namespaceId The ID of the Namespace. It can contain 2 to 32 lowercase characters. The value is in format `{RegionId}:{namespace}`.
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(@Nullable Output<String> namespaceId) {
            $.namespaceId = namespaceId;
            return this;
        }

        /**
         * @param namespaceId The ID of the Namespace. It can contain 2 to 32 lowercase characters. The value is in format `{RegionId}:{namespace}`.
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(String namespaceId) {
            return namespaceId(Output.of(namespaceId));
        }

        /**
         * @param namespaceName The Name of Namespace.
         * 
         * @return builder
         * 
         */
        public Builder namespaceName(Output<String> namespaceName) {
            $.namespaceName = namespaceName;
            return this;
        }

        /**
         * @param namespaceName The Name of Namespace.
         * 
         * @return builder
         * 
         */
        public Builder namespaceName(String namespaceName) {
            return namespaceName(Output.of(namespaceName));
        }

        /**
         * @param namespaceShortId The short ID of the Namespace. You do not need to specify a region ID. The value of `namespace_short_id` can be up to 20 characters in length and can contain only lowercase letters and digits.
         * 
         * @return builder
         * 
         */
        public Builder namespaceShortId(@Nullable Output<String> namespaceShortId) {
            $.namespaceShortId = namespaceShortId;
            return this;
        }

        /**
         * @param namespaceShortId The short ID of the Namespace. You do not need to specify a region ID. The value of `namespace_short_id` can be up to 20 characters in length and can contain only lowercase letters and digits.
         * 
         * @return builder
         * 
         */
        public Builder namespaceShortId(String namespaceShortId) {
            return namespaceShortId(Output.of(namespaceShortId));
        }

        public NamespaceArgs build() {
            $.namespaceName = Objects.requireNonNull($.namespaceName, "expected parameter 'namespaceName' to be non-null");
            return $;
        }
    }

}
