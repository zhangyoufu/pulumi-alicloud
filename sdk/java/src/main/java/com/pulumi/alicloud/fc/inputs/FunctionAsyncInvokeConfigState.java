// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.fc.inputs;

import com.pulumi.alicloud.fc.inputs.FunctionAsyncInvokeConfigDestinationConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FunctionAsyncInvokeConfigState extends com.pulumi.resources.ResourceArgs {

    public static final FunctionAsyncInvokeConfigState Empty = new FunctionAsyncInvokeConfigState();

    /**
     * The date this resource was created.
     * 
     */
    @Import(name="createdTime")
    private @Nullable Output<String> createdTime;

    /**
     * @return The date this resource was created.
     * 
     */
    public Optional<Output<String>> createdTime() {
        return Optional.ofNullable(this.createdTime);
    }

    /**
     * Configuration block with destination configuration. See `destination_config` below.
     * 
     */
    @Import(name="destinationConfig")
    private @Nullable Output<FunctionAsyncInvokeConfigDestinationConfigArgs> destinationConfig;

    /**
     * @return Configuration block with destination configuration. See `destination_config` below.
     * 
     */
    public Optional<Output<FunctionAsyncInvokeConfigDestinationConfigArgs>> destinationConfig() {
        return Optional.ofNullable(this.destinationConfig);
    }

    /**
     * Name of the Function Compute Function.
     * 
     */
    @Import(name="functionName")
    private @Nullable Output<String> functionName;

    /**
     * @return Name of the Function Compute Function.
     * 
     */
    public Optional<Output<String>> functionName() {
        return Optional.ofNullable(this.functionName);
    }

    /**
     * The date this resource was last modified.
     * 
     */
    @Import(name="lastModifiedTime")
    private @Nullable Output<String> lastModifiedTime;

    /**
     * @return The date this resource was last modified.
     * 
     */
    public Optional<Output<String>> lastModifiedTime() {
        return Optional.ofNullable(this.lastModifiedTime);
    }

    /**
     * Maximum age of a request that Function Compute sends to a function for processing in seconds. Valid values between 1 and 2592000 (between 60 and 21600 before v1.167.0).
     * 
     */
    @Import(name="maximumEventAgeInSeconds")
    private @Nullable Output<Integer> maximumEventAgeInSeconds;

    /**
     * @return Maximum age of a request that Function Compute sends to a function for processing in seconds. Valid values between 1 and 2592000 (between 60 and 21600 before v1.167.0).
     * 
     */
    public Optional<Output<Integer>> maximumEventAgeInSeconds() {
        return Optional.ofNullable(this.maximumEventAgeInSeconds);
    }

    /**
     * Maximum number of times to retry when the function returns an error. Valid values between 0 and 8 (between 0 and 2 before v1.167.0). Defaults to 2.
     * 
     */
    @Import(name="maximumRetryAttempts")
    private @Nullable Output<Integer> maximumRetryAttempts;

    /**
     * @return Maximum number of times to retry when the function returns an error. Valid values between 0 and 8 (between 0 and 2 before v1.167.0). Defaults to 2.
     * 
     */
    public Optional<Output<Integer>> maximumRetryAttempts() {
        return Optional.ofNullable(this.maximumRetryAttempts);
    }

    /**
     * Function Compute Function published version, `LATEST`, or Function Compute Alias name. The default value is `LATEST`.
     * 
     */
    @Import(name="qualifier")
    private @Nullable Output<String> qualifier;

    /**
     * @return Function Compute Function published version, `LATEST`, or Function Compute Alias name. The default value is `LATEST`.
     * 
     */
    public Optional<Output<String>> qualifier() {
        return Optional.ofNullable(this.qualifier);
    }

    /**
     * Name of the Function Compute Function, omitting any version or alias qualifier.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return Name of the Function Compute Function, omitting any version or alias qualifier.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Function Compute async job configuration(also known as Task Mode). valid values true or false, default `false`
     * 
     */
    @Import(name="statefulInvocation")
    private @Nullable Output<Boolean> statefulInvocation;

    /**
     * @return Function Compute async job configuration(also known as Task Mode). valid values true or false, default `false`
     * 
     */
    public Optional<Output<Boolean>> statefulInvocation() {
        return Optional.ofNullable(this.statefulInvocation);
    }

    private FunctionAsyncInvokeConfigState() {}

    private FunctionAsyncInvokeConfigState(FunctionAsyncInvokeConfigState $) {
        this.createdTime = $.createdTime;
        this.destinationConfig = $.destinationConfig;
        this.functionName = $.functionName;
        this.lastModifiedTime = $.lastModifiedTime;
        this.maximumEventAgeInSeconds = $.maximumEventAgeInSeconds;
        this.maximumRetryAttempts = $.maximumRetryAttempts;
        this.qualifier = $.qualifier;
        this.serviceName = $.serviceName;
        this.statefulInvocation = $.statefulInvocation;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FunctionAsyncInvokeConfigState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FunctionAsyncInvokeConfigState $;

        public Builder() {
            $ = new FunctionAsyncInvokeConfigState();
        }

        public Builder(FunctionAsyncInvokeConfigState defaults) {
            $ = new FunctionAsyncInvokeConfigState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createdTime The date this resource was created.
         * 
         * @return builder
         * 
         */
        public Builder createdTime(@Nullable Output<String> createdTime) {
            $.createdTime = createdTime;
            return this;
        }

        /**
         * @param createdTime The date this resource was created.
         * 
         * @return builder
         * 
         */
        public Builder createdTime(String createdTime) {
            return createdTime(Output.of(createdTime));
        }

        /**
         * @param destinationConfig Configuration block with destination configuration. See `destination_config` below.
         * 
         * @return builder
         * 
         */
        public Builder destinationConfig(@Nullable Output<FunctionAsyncInvokeConfigDestinationConfigArgs> destinationConfig) {
            $.destinationConfig = destinationConfig;
            return this;
        }

        /**
         * @param destinationConfig Configuration block with destination configuration. See `destination_config` below.
         * 
         * @return builder
         * 
         */
        public Builder destinationConfig(FunctionAsyncInvokeConfigDestinationConfigArgs destinationConfig) {
            return destinationConfig(Output.of(destinationConfig));
        }

        /**
         * @param functionName Name of the Function Compute Function.
         * 
         * @return builder
         * 
         */
        public Builder functionName(@Nullable Output<String> functionName) {
            $.functionName = functionName;
            return this;
        }

        /**
         * @param functionName Name of the Function Compute Function.
         * 
         * @return builder
         * 
         */
        public Builder functionName(String functionName) {
            return functionName(Output.of(functionName));
        }

        /**
         * @param lastModifiedTime The date this resource was last modified.
         * 
         * @return builder
         * 
         */
        public Builder lastModifiedTime(@Nullable Output<String> lastModifiedTime) {
            $.lastModifiedTime = lastModifiedTime;
            return this;
        }

        /**
         * @param lastModifiedTime The date this resource was last modified.
         * 
         * @return builder
         * 
         */
        public Builder lastModifiedTime(String lastModifiedTime) {
            return lastModifiedTime(Output.of(lastModifiedTime));
        }

        /**
         * @param maximumEventAgeInSeconds Maximum age of a request that Function Compute sends to a function for processing in seconds. Valid values between 1 and 2592000 (between 60 and 21600 before v1.167.0).
         * 
         * @return builder
         * 
         */
        public Builder maximumEventAgeInSeconds(@Nullable Output<Integer> maximumEventAgeInSeconds) {
            $.maximumEventAgeInSeconds = maximumEventAgeInSeconds;
            return this;
        }

        /**
         * @param maximumEventAgeInSeconds Maximum age of a request that Function Compute sends to a function for processing in seconds. Valid values between 1 and 2592000 (between 60 and 21600 before v1.167.0).
         * 
         * @return builder
         * 
         */
        public Builder maximumEventAgeInSeconds(Integer maximumEventAgeInSeconds) {
            return maximumEventAgeInSeconds(Output.of(maximumEventAgeInSeconds));
        }

        /**
         * @param maximumRetryAttempts Maximum number of times to retry when the function returns an error. Valid values between 0 and 8 (between 0 and 2 before v1.167.0). Defaults to 2.
         * 
         * @return builder
         * 
         */
        public Builder maximumRetryAttempts(@Nullable Output<Integer> maximumRetryAttempts) {
            $.maximumRetryAttempts = maximumRetryAttempts;
            return this;
        }

        /**
         * @param maximumRetryAttempts Maximum number of times to retry when the function returns an error. Valid values between 0 and 8 (between 0 and 2 before v1.167.0). Defaults to 2.
         * 
         * @return builder
         * 
         */
        public Builder maximumRetryAttempts(Integer maximumRetryAttempts) {
            return maximumRetryAttempts(Output.of(maximumRetryAttempts));
        }

        /**
         * @param qualifier Function Compute Function published version, `LATEST`, or Function Compute Alias name. The default value is `LATEST`.
         * 
         * @return builder
         * 
         */
        public Builder qualifier(@Nullable Output<String> qualifier) {
            $.qualifier = qualifier;
            return this;
        }

        /**
         * @param qualifier Function Compute Function published version, `LATEST`, or Function Compute Alias name. The default value is `LATEST`.
         * 
         * @return builder
         * 
         */
        public Builder qualifier(String qualifier) {
            return qualifier(Output.of(qualifier));
        }

        /**
         * @param serviceName Name of the Function Compute Function, omitting any version or alias qualifier.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName Name of the Function Compute Function, omitting any version or alias qualifier.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param statefulInvocation Function Compute async job configuration(also known as Task Mode). valid values true or false, default `false`
         * 
         * @return builder
         * 
         */
        public Builder statefulInvocation(@Nullable Output<Boolean> statefulInvocation) {
            $.statefulInvocation = statefulInvocation;
            return this;
        }

        /**
         * @param statefulInvocation Function Compute async job configuration(also known as Task Mode). valid values true or false, default `false`
         * 
         * @return builder
         * 
         */
        public Builder statefulInvocation(Boolean statefulInvocation) {
            return statefulInvocation(Output.of(statefulInvocation));
        }

        public FunctionAsyncInvokeConfigState build() {
            return $;
        }
    }

}
