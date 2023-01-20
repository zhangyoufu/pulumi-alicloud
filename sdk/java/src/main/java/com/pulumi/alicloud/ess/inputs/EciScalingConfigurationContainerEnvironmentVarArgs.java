// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EciScalingConfigurationContainerEnvironmentVarArgs extends com.pulumi.resources.ResourceArgs {

    public static final EciScalingConfigurationContainerEnvironmentVarArgs Empty = new EciScalingConfigurationContainerEnvironmentVarArgs();

    /**
     * The name of the variable. The name can be 1 to 128 characters in length and can contain letters,
     * digits, and underscores (_). It cannot start with a digit.
     * digits, and underscores (_). It cannot start with a digit.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return The name of the variable. The name can be 1 to 128 characters in length and can contain letters,
     * digits, and underscores (_). It cannot start with a digit.
     * digits, and underscores (_). It cannot start with a digit.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * The value of the variable. The value can be 0 to 256 characters in length.
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return The value of the variable. The value can be 0 to 256 characters in length.
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private EciScalingConfigurationContainerEnvironmentVarArgs() {}

    private EciScalingConfigurationContainerEnvironmentVarArgs(EciScalingConfigurationContainerEnvironmentVarArgs $) {
        this.key = $.key;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EciScalingConfigurationContainerEnvironmentVarArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EciScalingConfigurationContainerEnvironmentVarArgs $;

        public Builder() {
            $ = new EciScalingConfigurationContainerEnvironmentVarArgs();
        }

        public Builder(EciScalingConfigurationContainerEnvironmentVarArgs defaults) {
            $ = new EciScalingConfigurationContainerEnvironmentVarArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param key The name of the variable. The name can be 1 to 128 characters in length and can contain letters,
         * digits, and underscores (_). It cannot start with a digit.
         * digits, and underscores (_). It cannot start with a digit.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The name of the variable. The name can be 1 to 128 characters in length and can contain letters,
         * digits, and underscores (_). It cannot start with a digit.
         * digits, and underscores (_). It cannot start with a digit.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param value The value of the variable. The value can be 0 to 256 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The value of the variable. The value can be 0 to 256 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public EciScalingConfigurationContainerEnvironmentVarArgs build() {
            return $;
        }
    }

}
