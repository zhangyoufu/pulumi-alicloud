// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.message;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceQueueArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceQueueArgs Empty = new ServiceQueueArgs();

    /**
     * This means that messages sent to the queue can only be consumed after the delay time set by this parameter, in seconds.
     * 
     */
    @Import(name="delaySeconds")
    private @Nullable Output<Integer> delaySeconds;

    /**
     * @return This means that messages sent to the queue can only be consumed after the delay time set by this parameter, in seconds.
     * 
     */
    public Optional<Output<Integer>> delaySeconds() {
        return Optional.ofNullable(this.delaySeconds);
    }

    /**
     * Represents whether the log management function is enabled.
     * 
     */
    @Import(name="loggingEnabled")
    private @Nullable Output<Boolean> loggingEnabled;

    /**
     * @return Represents whether the log management function is enabled.
     * 
     */
    public Optional<Output<Boolean>> loggingEnabled() {
        return Optional.ofNullable(this.loggingEnabled);
    }

    /**
     * Represents the maximum length of the message body sent to the Queue, in Byte.
     * 
     */
    @Import(name="maximumMessageSize")
    private @Nullable Output<Integer> maximumMessageSize;

    /**
     * @return Represents the maximum length of the message body sent to the Queue, in Byte.
     * 
     */
    public Optional<Output<Integer>> maximumMessageSize() {
        return Optional.ofNullable(this.maximumMessageSize);
    }

    /**
     * Represents the longest life time of the message in the Queue.
     * 
     */
    @Import(name="messageRetentionPeriod")
    private @Nullable Output<Integer> messageRetentionPeriod;

    /**
     * @return Represents the longest life time of the message in the Queue.
     * 
     */
    public Optional<Output<Integer>> messageRetentionPeriod() {
        return Optional.ofNullable(this.messageRetentionPeriod);
    }

    /**
     * The longest waiting time for a Queue request when the number of messages is empty, in seconds.
     * 
     */
    @Import(name="pollingWaitSeconds")
    private @Nullable Output<Integer> pollingWaitSeconds;

    /**
     * @return The longest waiting time for a Queue request when the number of messages is empty, in seconds.
     * 
     */
    public Optional<Output<Integer>> pollingWaitSeconds() {
        return Optional.ofNullable(this.pollingWaitSeconds);
    }

    /**
     * Representative resources.
     * 
     */
    @Import(name="queueName", required=true)
    private Output<String> queueName;

    /**
     * @return Representative resources.
     * 
     */
    public Output<String> queueName() {
        return this.queueName;
    }

    /**
     * Represents the duration after the message is removed from the Queue and changed from the Active state to the Inactive state.
     * 
     */
    @Import(name="visibilityTimeout")
    private @Nullable Output<Integer> visibilityTimeout;

    /**
     * @return Represents the duration after the message is removed from the Queue and changed from the Active state to the Inactive state.
     * 
     */
    public Optional<Output<Integer>> visibilityTimeout() {
        return Optional.ofNullable(this.visibilityTimeout);
    }

    private ServiceQueueArgs() {}

    private ServiceQueueArgs(ServiceQueueArgs $) {
        this.delaySeconds = $.delaySeconds;
        this.loggingEnabled = $.loggingEnabled;
        this.maximumMessageSize = $.maximumMessageSize;
        this.messageRetentionPeriod = $.messageRetentionPeriod;
        this.pollingWaitSeconds = $.pollingWaitSeconds;
        this.queueName = $.queueName;
        this.visibilityTimeout = $.visibilityTimeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceQueueArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceQueueArgs $;

        public Builder() {
            $ = new ServiceQueueArgs();
        }

        public Builder(ServiceQueueArgs defaults) {
            $ = new ServiceQueueArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param delaySeconds This means that messages sent to the queue can only be consumed after the delay time set by this parameter, in seconds.
         * 
         * @return builder
         * 
         */
        public Builder delaySeconds(@Nullable Output<Integer> delaySeconds) {
            $.delaySeconds = delaySeconds;
            return this;
        }

        /**
         * @param delaySeconds This means that messages sent to the queue can only be consumed after the delay time set by this parameter, in seconds.
         * 
         * @return builder
         * 
         */
        public Builder delaySeconds(Integer delaySeconds) {
            return delaySeconds(Output.of(delaySeconds));
        }

        /**
         * @param loggingEnabled Represents whether the log management function is enabled.
         * 
         * @return builder
         * 
         */
        public Builder loggingEnabled(@Nullable Output<Boolean> loggingEnabled) {
            $.loggingEnabled = loggingEnabled;
            return this;
        }

        /**
         * @param loggingEnabled Represents whether the log management function is enabled.
         * 
         * @return builder
         * 
         */
        public Builder loggingEnabled(Boolean loggingEnabled) {
            return loggingEnabled(Output.of(loggingEnabled));
        }

        /**
         * @param maximumMessageSize Represents the maximum length of the message body sent to the Queue, in Byte.
         * 
         * @return builder
         * 
         */
        public Builder maximumMessageSize(@Nullable Output<Integer> maximumMessageSize) {
            $.maximumMessageSize = maximumMessageSize;
            return this;
        }

        /**
         * @param maximumMessageSize Represents the maximum length of the message body sent to the Queue, in Byte.
         * 
         * @return builder
         * 
         */
        public Builder maximumMessageSize(Integer maximumMessageSize) {
            return maximumMessageSize(Output.of(maximumMessageSize));
        }

        /**
         * @param messageRetentionPeriod Represents the longest life time of the message in the Queue.
         * 
         * @return builder
         * 
         */
        public Builder messageRetentionPeriod(@Nullable Output<Integer> messageRetentionPeriod) {
            $.messageRetentionPeriod = messageRetentionPeriod;
            return this;
        }

        /**
         * @param messageRetentionPeriod Represents the longest life time of the message in the Queue.
         * 
         * @return builder
         * 
         */
        public Builder messageRetentionPeriod(Integer messageRetentionPeriod) {
            return messageRetentionPeriod(Output.of(messageRetentionPeriod));
        }

        /**
         * @param pollingWaitSeconds The longest waiting time for a Queue request when the number of messages is empty, in seconds.
         * 
         * @return builder
         * 
         */
        public Builder pollingWaitSeconds(@Nullable Output<Integer> pollingWaitSeconds) {
            $.pollingWaitSeconds = pollingWaitSeconds;
            return this;
        }

        /**
         * @param pollingWaitSeconds The longest waiting time for a Queue request when the number of messages is empty, in seconds.
         * 
         * @return builder
         * 
         */
        public Builder pollingWaitSeconds(Integer pollingWaitSeconds) {
            return pollingWaitSeconds(Output.of(pollingWaitSeconds));
        }

        /**
         * @param queueName Representative resources.
         * 
         * @return builder
         * 
         */
        public Builder queueName(Output<String> queueName) {
            $.queueName = queueName;
            return this;
        }

        /**
         * @param queueName Representative resources.
         * 
         * @return builder
         * 
         */
        public Builder queueName(String queueName) {
            return queueName(Output.of(queueName));
        }

        /**
         * @param visibilityTimeout Represents the duration after the message is removed from the Queue and changed from the Active state to the Inactive state.
         * 
         * @return builder
         * 
         */
        public Builder visibilityTimeout(@Nullable Output<Integer> visibilityTimeout) {
            $.visibilityTimeout = visibilityTimeout;
            return this;
        }

        /**
         * @param visibilityTimeout Represents the duration after the message is removed from the Queue and changed from the Active state to the Inactive state.
         * 
         * @return builder
         * 
         */
        public Builder visibilityTimeout(Integer visibilityTimeout) {
            return visibilityTimeout(Output.of(visibilityTimeout));
        }

        public ServiceQueueArgs build() {
            if ($.queueName == null) {
                throw new MissingRequiredPropertyException("ServiceQueueArgs", "queueName");
            }
            return $;
        }
    }

}
