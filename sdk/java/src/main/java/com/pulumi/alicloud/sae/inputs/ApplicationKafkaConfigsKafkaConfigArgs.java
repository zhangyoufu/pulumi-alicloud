// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ApplicationKafkaConfigsKafkaConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ApplicationKafkaConfigsKafkaConfigArgs Empty = new ApplicationKafkaConfigsKafkaConfigArgs();

    /**
     * The topic of the Kafka.
     * 
     */
    @Import(name="kafkaTopic")
    private @Nullable Output<String> kafkaTopic;

    /**
     * @return The topic of the Kafka.
     * 
     */
    public Optional<Output<String>> kafkaTopic() {
        return Optional.ofNullable(this.kafkaTopic);
    }

    /**
     * The path in which logs are stored.
     * 
     */
    @Import(name="logDir")
    private @Nullable Output<String> logDir;

    /**
     * @return The path in which logs are stored.
     * 
     */
    public Optional<Output<String>> logDir() {
        return Optional.ofNullable(this.logDir);
    }

    /**
     * The type of the log.
     * 
     */
    @Import(name="logType")
    private @Nullable Output<String> logType;

    /**
     * @return The type of the log.
     * 
     */
    public Optional<Output<String>> logType() {
        return Optional.ofNullable(this.logType);
    }

    private ApplicationKafkaConfigsKafkaConfigArgs() {}

    private ApplicationKafkaConfigsKafkaConfigArgs(ApplicationKafkaConfigsKafkaConfigArgs $) {
        this.kafkaTopic = $.kafkaTopic;
        this.logDir = $.logDir;
        this.logType = $.logType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApplicationKafkaConfigsKafkaConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApplicationKafkaConfigsKafkaConfigArgs $;

        public Builder() {
            $ = new ApplicationKafkaConfigsKafkaConfigArgs();
        }

        public Builder(ApplicationKafkaConfigsKafkaConfigArgs defaults) {
            $ = new ApplicationKafkaConfigsKafkaConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param kafkaTopic The topic of the Kafka.
         * 
         * @return builder
         * 
         */
        public Builder kafkaTopic(@Nullable Output<String> kafkaTopic) {
            $.kafkaTopic = kafkaTopic;
            return this;
        }

        /**
         * @param kafkaTopic The topic of the Kafka.
         * 
         * @return builder
         * 
         */
        public Builder kafkaTopic(String kafkaTopic) {
            return kafkaTopic(Output.of(kafkaTopic));
        }

        /**
         * @param logDir The path in which logs are stored.
         * 
         * @return builder
         * 
         */
        public Builder logDir(@Nullable Output<String> logDir) {
            $.logDir = logDir;
            return this;
        }

        /**
         * @param logDir The path in which logs are stored.
         * 
         * @return builder
         * 
         */
        public Builder logDir(String logDir) {
            return logDir(Output.of(logDir));
        }

        /**
         * @param logType The type of the log.
         * 
         * @return builder
         * 
         */
        public Builder logType(@Nullable Output<String> logType) {
            $.logType = logType;
            return this;
        }

        /**
         * @param logType The type of the log.
         * 
         * @return builder
         * 
         */
        public Builder logType(String logType) {
            return logType(Output.of(logType));
        }

        public ApplicationKafkaConfigsKafkaConfigArgs build() {
            return $;
        }
    }

}
