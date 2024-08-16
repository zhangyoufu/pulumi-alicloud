// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eventbridge;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EventSourceArgs extends com.pulumi.resources.ResourceArgs {

    public static final EventSourceArgs Empty = new EventSourceArgs();

    /**
     * The detail describe of event source.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The detail describe of event source.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of event bus.
     * 
     */
    @Import(name="eventBusName", required=true)
    private Output<String> eventBusName;

    /**
     * @return The name of event bus.
     * 
     */
    public Output<String> eventBusName() {
        return this.eventBusName;
    }

    /**
     * The code name of event source.
     * 
     */
    @Import(name="eventSourceName", required=true)
    private Output<String> eventSourceName;

    /**
     * @return The code name of event source.
     * 
     */
    public Output<String> eventSourceName() {
        return this.eventSourceName;
    }

    /**
     * The config of external source.
     * When `external_source_type` is `RabbitMQ`, The following attributes are supported:
     * `RegionId` - The region ID of RabbitMQ.
     * `InstanceId` - The instance ID of RabbitMQ.
     * `VirtualHostName` - The virtual host name of RabbitMQ.
     * `QueueName` - The queue name of RabbitMQ.
     * When `external_source_type` is `RabbitMQ`, The following attributes are supported:
     * `RegionId` - The region ID of RabbitMQ.
     * `InstanceId` - The instance ID of RabbitMQ.
     * `Topic` - The topic of RabbitMQ.
     * `Offset` -  The offset of RabbitMQ, valid values: `CONSUME_FROM_FIRST_OFFSET`, `CONSUME_FROM_LAST_OFFSET` and `CONSUME_FROM_TIMESTAMP`.
     * `GroupID` - The group ID of consumer.
     * When `external_source_type` is `MNS`, The following attributes are supported:
     * `QueueName` - The queue name of MNS.
     * 
     */
    @Import(name="externalSourceConfig")
    private @Nullable Output<Map<String,String>> externalSourceConfig;

    /**
     * @return The config of external source.
     * When `external_source_type` is `RabbitMQ`, The following attributes are supported:
     * `RegionId` - The region ID of RabbitMQ.
     * `InstanceId` - The instance ID of RabbitMQ.
     * `VirtualHostName` - The virtual host name of RabbitMQ.
     * `QueueName` - The queue name of RabbitMQ.
     * When `external_source_type` is `RabbitMQ`, The following attributes are supported:
     * `RegionId` - The region ID of RabbitMQ.
     * `InstanceId` - The instance ID of RabbitMQ.
     * `Topic` - The topic of RabbitMQ.
     * `Offset` -  The offset of RabbitMQ, valid values: `CONSUME_FROM_FIRST_OFFSET`, `CONSUME_FROM_LAST_OFFSET` and `CONSUME_FROM_TIMESTAMP`.
     * `GroupID` - The group ID of consumer.
     * When `external_source_type` is `MNS`, The following attributes are supported:
     * `QueueName` - The queue name of MNS.
     * 
     */
    public Optional<Output<Map<String,String>>> externalSourceConfig() {
        return Optional.ofNullable(this.externalSourceConfig);
    }

    /**
     * The type of external data source. Valid value : `RabbitMQ`, `RocketMQ` and `MNS`. **NOTE:** Only When `linked_external_source` is `true`, This field is valid.
     * 
     */
    @Import(name="externalSourceType")
    private @Nullable Output<String> externalSourceType;

    /**
     * @return The type of external data source. Valid value : `RabbitMQ`, `RocketMQ` and `MNS`. **NOTE:** Only When `linked_external_source` is `true`, This field is valid.
     * 
     */
    public Optional<Output<String>> externalSourceType() {
        return Optional.ofNullable(this.externalSourceType);
    }

    /**
     * Whether to connect to an external data source. Default value: `false`
     * 
     */
    @Import(name="linkedExternalSource")
    private @Nullable Output<Boolean> linkedExternalSource;

    /**
     * @return Whether to connect to an external data source. Default value: `false`
     * 
     */
    public Optional<Output<Boolean>> linkedExternalSource() {
        return Optional.ofNullable(this.linkedExternalSource);
    }

    private EventSourceArgs() {}

    private EventSourceArgs(EventSourceArgs $) {
        this.description = $.description;
        this.eventBusName = $.eventBusName;
        this.eventSourceName = $.eventSourceName;
        this.externalSourceConfig = $.externalSourceConfig;
        this.externalSourceType = $.externalSourceType;
        this.linkedExternalSource = $.linkedExternalSource;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EventSourceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EventSourceArgs $;

        public Builder() {
            $ = new EventSourceArgs();
        }

        public Builder(EventSourceArgs defaults) {
            $ = new EventSourceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The detail describe of event source.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The detail describe of event source.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param eventBusName The name of event bus.
         * 
         * @return builder
         * 
         */
        public Builder eventBusName(Output<String> eventBusName) {
            $.eventBusName = eventBusName;
            return this;
        }

        /**
         * @param eventBusName The name of event bus.
         * 
         * @return builder
         * 
         */
        public Builder eventBusName(String eventBusName) {
            return eventBusName(Output.of(eventBusName));
        }

        /**
         * @param eventSourceName The code name of event source.
         * 
         * @return builder
         * 
         */
        public Builder eventSourceName(Output<String> eventSourceName) {
            $.eventSourceName = eventSourceName;
            return this;
        }

        /**
         * @param eventSourceName The code name of event source.
         * 
         * @return builder
         * 
         */
        public Builder eventSourceName(String eventSourceName) {
            return eventSourceName(Output.of(eventSourceName));
        }

        /**
         * @param externalSourceConfig The config of external source.
         * When `external_source_type` is `RabbitMQ`, The following attributes are supported:
         * `RegionId` - The region ID of RabbitMQ.
         * `InstanceId` - The instance ID of RabbitMQ.
         * `VirtualHostName` - The virtual host name of RabbitMQ.
         * `QueueName` - The queue name of RabbitMQ.
         * When `external_source_type` is `RabbitMQ`, The following attributes are supported:
         * `RegionId` - The region ID of RabbitMQ.
         * `InstanceId` - The instance ID of RabbitMQ.
         * `Topic` - The topic of RabbitMQ.
         * `Offset` -  The offset of RabbitMQ, valid values: `CONSUME_FROM_FIRST_OFFSET`, `CONSUME_FROM_LAST_OFFSET` and `CONSUME_FROM_TIMESTAMP`.
         * `GroupID` - The group ID of consumer.
         * When `external_source_type` is `MNS`, The following attributes are supported:
         * `QueueName` - The queue name of MNS.
         * 
         * @return builder
         * 
         */
        public Builder externalSourceConfig(@Nullable Output<Map<String,String>> externalSourceConfig) {
            $.externalSourceConfig = externalSourceConfig;
            return this;
        }

        /**
         * @param externalSourceConfig The config of external source.
         * When `external_source_type` is `RabbitMQ`, The following attributes are supported:
         * `RegionId` - The region ID of RabbitMQ.
         * `InstanceId` - The instance ID of RabbitMQ.
         * `VirtualHostName` - The virtual host name of RabbitMQ.
         * `QueueName` - The queue name of RabbitMQ.
         * When `external_source_type` is `RabbitMQ`, The following attributes are supported:
         * `RegionId` - The region ID of RabbitMQ.
         * `InstanceId` - The instance ID of RabbitMQ.
         * `Topic` - The topic of RabbitMQ.
         * `Offset` -  The offset of RabbitMQ, valid values: `CONSUME_FROM_FIRST_OFFSET`, `CONSUME_FROM_LAST_OFFSET` and `CONSUME_FROM_TIMESTAMP`.
         * `GroupID` - The group ID of consumer.
         * When `external_source_type` is `MNS`, The following attributes are supported:
         * `QueueName` - The queue name of MNS.
         * 
         * @return builder
         * 
         */
        public Builder externalSourceConfig(Map<String,String> externalSourceConfig) {
            return externalSourceConfig(Output.of(externalSourceConfig));
        }

        /**
         * @param externalSourceType The type of external data source. Valid value : `RabbitMQ`, `RocketMQ` and `MNS`. **NOTE:** Only When `linked_external_source` is `true`, This field is valid.
         * 
         * @return builder
         * 
         */
        public Builder externalSourceType(@Nullable Output<String> externalSourceType) {
            $.externalSourceType = externalSourceType;
            return this;
        }

        /**
         * @param externalSourceType The type of external data source. Valid value : `RabbitMQ`, `RocketMQ` and `MNS`. **NOTE:** Only When `linked_external_source` is `true`, This field is valid.
         * 
         * @return builder
         * 
         */
        public Builder externalSourceType(String externalSourceType) {
            return externalSourceType(Output.of(externalSourceType));
        }

        /**
         * @param linkedExternalSource Whether to connect to an external data source. Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder linkedExternalSource(@Nullable Output<Boolean> linkedExternalSource) {
            $.linkedExternalSource = linkedExternalSource;
            return this;
        }

        /**
         * @param linkedExternalSource Whether to connect to an external data source. Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder linkedExternalSource(Boolean linkedExternalSource) {
            return linkedExternalSource(Output.of(linkedExternalSource));
        }

        public EventSourceArgs build() {
            if ($.eventBusName == null) {
                throw new MissingRequiredPropertyException("EventSourceArgs", "eventBusName");
            }
            if ($.eventSourceName == null) {
                throw new MissingRequiredPropertyException("EventSourceArgs", "eventSourceName");
            }
            return $;
        }
    }

}
