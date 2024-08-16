// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga;

import com.pulumi.alicloud.ga.inputs.EndpointGroupEndpointConfigurationArgs;
import com.pulumi.alicloud.ga.inputs.EndpointGroupPortOverridesArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EndpointGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final EndpointGroupArgs Empty = new EndpointGroupArgs();

    /**
     * The ID of the Global Accelerator instance to which the endpoint group will be added.
     * 
     */
    @Import(name="acceleratorId", required=true)
    private Output<String> acceleratorId;

    /**
     * @return The ID of the Global Accelerator instance to which the endpoint group will be added.
     * 
     */
    public Output<String> acceleratorId() {
        return this.acceleratorId;
    }

    /**
     * The description of the endpoint group.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the endpoint group.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The endpointConfigurations of the endpoint group. See `endpoint_configurations` below.
     * 
     */
    @Import(name="endpointConfigurations", required=true)
    private Output<List<EndpointGroupEndpointConfigurationArgs>> endpointConfigurations;

    /**
     * @return The endpointConfigurations of the endpoint group. See `endpoint_configurations` below.
     * 
     */
    public Output<List<EndpointGroupEndpointConfigurationArgs>> endpointConfigurations() {
        return this.endpointConfigurations;
    }

    /**
     * The ID of the region where the endpoint group is deployed.
     * 
     */
    @Import(name="endpointGroupRegion", required=true)
    private Output<String> endpointGroupRegion;

    /**
     * @return The ID of the region where the endpoint group is deployed.
     * 
     */
    public Output<String> endpointGroupRegion() {
        return this.endpointGroupRegion;
    }

    /**
     * The endpoint group type. Default value: `default`. Valid values: `default`, `virtual`.
     * &gt; **NOTE:** Currently, only `HTTP` or `HTTPS` protocol listener can directly create a `virtual` Endpoint Group. If it is `TCP` protocol listener, and you want to create a `virtual` Endpoint Group, please ensure that the `default` Endpoint Group has been created.
     * 
     */
    @Import(name="endpointGroupType")
    private @Nullable Output<String> endpointGroupType;

    /**
     * @return The endpoint group type. Default value: `default`. Valid values: `default`, `virtual`.
     * &gt; **NOTE:** Currently, only `HTTP` or `HTTPS` protocol listener can directly create a `virtual` Endpoint Group. If it is `TCP` protocol listener, and you want to create a `virtual` Endpoint Group, please ensure that the `default` Endpoint Group has been created.
     * 
     */
    public Optional<Output<String>> endpointGroupType() {
        return Optional.ofNullable(this.endpointGroupType);
    }

    /**
     * The endpoint request protocol. Valid values: `HTTP`, `HTTPS`.
     * &gt; **NOTE:** This item is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. For the listening instance of HTTP protocol, the back-end service protocol supports and only supports HTTP.
     * 
     */
    @Import(name="endpointRequestProtocol")
    private @Nullable Output<String> endpointRequestProtocol;

    /**
     * @return The endpoint request protocol. Valid values: `HTTP`, `HTTPS`.
     * &gt; **NOTE:** This item is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. For the listening instance of HTTP protocol, the back-end service protocol supports and only supports HTTP.
     * 
     */
    public Optional<Output<String>> endpointRequestProtocol() {
        return Optional.ofNullable(this.endpointRequestProtocol);
    }

    /**
     * Specifies whether to enable the health check feature. Valid values:
     * 
     */
    @Import(name="healthCheckEnabled")
    private @Nullable Output<Boolean> healthCheckEnabled;

    /**
     * @return Specifies whether to enable the health check feature. Valid values:
     * 
     */
    public Optional<Output<Boolean>> healthCheckEnabled() {
        return Optional.ofNullable(this.healthCheckEnabled);
    }

    /**
     * The interval between two consecutive health checks. Unit: seconds.
     * 
     */
    @Import(name="healthCheckIntervalSeconds")
    private @Nullable Output<Integer> healthCheckIntervalSeconds;

    /**
     * @return The interval between two consecutive health checks. Unit: seconds.
     * 
     */
    public Optional<Output<Integer>> healthCheckIntervalSeconds() {
        return Optional.ofNullable(this.healthCheckIntervalSeconds);
    }

    /**
     * The path specified as the destination of the targets for health checks.
     * 
     */
    @Import(name="healthCheckPath")
    private @Nullable Output<String> healthCheckPath;

    /**
     * @return The path specified as the destination of the targets for health checks.
     * 
     */
    public Optional<Output<String>> healthCheckPath() {
        return Optional.ofNullable(this.healthCheckPath);
    }

    /**
     * The port that is used for health checks.
     * 
     */
    @Import(name="healthCheckPort")
    private @Nullable Output<Integer> healthCheckPort;

    /**
     * @return The port that is used for health checks.
     * 
     */
    public Optional<Output<Integer>> healthCheckPort() {
        return Optional.ofNullable(this.healthCheckPort);
    }

    /**
     * The protocol that is used to connect to the targets for health checks. Valid values:
     * - `TCP` or `tcp`: TCP protocol.
     * - `HTTP` or `http`: HTTP protocol.
     * - `HTTPS` or `https`: HTTPS protocol.
     * &gt; **NOTE:** From version 1.223.0, `health_check_protocol` can be set to `TCP`, `HTTP`, `HTTPS`.
     * 
     */
    @Import(name="healthCheckProtocol")
    private @Nullable Output<String> healthCheckProtocol;

    /**
     * @return The protocol that is used to connect to the targets for health checks. Valid values:
     * - `TCP` or `tcp`: TCP protocol.
     * - `HTTP` or `http`: HTTP protocol.
     * - `HTTPS` or `https`: HTTPS protocol.
     * &gt; **NOTE:** From version 1.223.0, `health_check_protocol` can be set to `TCP`, `HTTP`, `HTTPS`.
     * 
     */
    public Optional<Output<String>> healthCheckProtocol() {
        return Optional.ofNullable(this.healthCheckProtocol);
    }

    /**
     * The ID of the listener that is associated with the endpoint group.
     * 
     */
    @Import(name="listenerId", required=true)
    private Output<String> listenerId;

    /**
     * @return The ID of the listener that is associated with the endpoint group.
     * 
     */
    public Output<String> listenerId() {
        return this.listenerId;
    }

    /**
     * The name of the endpoint group.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the endpoint group.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Mapping between listening port and forwarding port of boarding point. See `port_overrides` below.
     * &gt; **NOTE:** Port mapping is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. The listening port in the port map must be consistent with the listening port of the current listening instance.
     * 
     */
    @Import(name="portOverrides")
    private @Nullable Output<EndpointGroupPortOverridesArgs> portOverrides;

    /**
     * @return Mapping between listening port and forwarding port of boarding point. See `port_overrides` below.
     * &gt; **NOTE:** Port mapping is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. The listening port in the port map must be consistent with the listening port of the current listening instance.
     * 
     */
    public Optional<Output<EndpointGroupPortOverridesArgs>> portOverrides() {
        return Optional.ofNullable(this.portOverrides);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value: `3`.
     * 
     */
    @Import(name="thresholdCount")
    private @Nullable Output<Integer> thresholdCount;

    /**
     * @return The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value: `3`.
     * 
     */
    public Optional<Output<Integer>> thresholdCount() {
        return Optional.ofNullable(this.thresholdCount);
    }

    /**
     * The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
     * 
     */
    @Import(name="trafficPercentage")
    private @Nullable Output<Integer> trafficPercentage;

    /**
     * @return The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
     * 
     */
    public Optional<Output<Integer>> trafficPercentage() {
        return Optional.ofNullable(this.trafficPercentage);
    }

    private EndpointGroupArgs() {}

    private EndpointGroupArgs(EndpointGroupArgs $) {
        this.acceleratorId = $.acceleratorId;
        this.description = $.description;
        this.endpointConfigurations = $.endpointConfigurations;
        this.endpointGroupRegion = $.endpointGroupRegion;
        this.endpointGroupType = $.endpointGroupType;
        this.endpointRequestProtocol = $.endpointRequestProtocol;
        this.healthCheckEnabled = $.healthCheckEnabled;
        this.healthCheckIntervalSeconds = $.healthCheckIntervalSeconds;
        this.healthCheckPath = $.healthCheckPath;
        this.healthCheckPort = $.healthCheckPort;
        this.healthCheckProtocol = $.healthCheckProtocol;
        this.listenerId = $.listenerId;
        this.name = $.name;
        this.portOverrides = $.portOverrides;
        this.tags = $.tags;
        this.thresholdCount = $.thresholdCount;
        this.trafficPercentage = $.trafficPercentage;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EndpointGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EndpointGroupArgs $;

        public Builder() {
            $ = new EndpointGroupArgs();
        }

        public Builder(EndpointGroupArgs defaults) {
            $ = new EndpointGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acceleratorId The ID of the Global Accelerator instance to which the endpoint group will be added.
         * 
         * @return builder
         * 
         */
        public Builder acceleratorId(Output<String> acceleratorId) {
            $.acceleratorId = acceleratorId;
            return this;
        }

        /**
         * @param acceleratorId The ID of the Global Accelerator instance to which the endpoint group will be added.
         * 
         * @return builder
         * 
         */
        public Builder acceleratorId(String acceleratorId) {
            return acceleratorId(Output.of(acceleratorId));
        }

        /**
         * @param description The description of the endpoint group.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the endpoint group.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param endpointConfigurations The endpointConfigurations of the endpoint group. See `endpoint_configurations` below.
         * 
         * @return builder
         * 
         */
        public Builder endpointConfigurations(Output<List<EndpointGroupEndpointConfigurationArgs>> endpointConfigurations) {
            $.endpointConfigurations = endpointConfigurations;
            return this;
        }

        /**
         * @param endpointConfigurations The endpointConfigurations of the endpoint group. See `endpoint_configurations` below.
         * 
         * @return builder
         * 
         */
        public Builder endpointConfigurations(List<EndpointGroupEndpointConfigurationArgs> endpointConfigurations) {
            return endpointConfigurations(Output.of(endpointConfigurations));
        }

        /**
         * @param endpointConfigurations The endpointConfigurations of the endpoint group. See `endpoint_configurations` below.
         * 
         * @return builder
         * 
         */
        public Builder endpointConfigurations(EndpointGroupEndpointConfigurationArgs... endpointConfigurations) {
            return endpointConfigurations(List.of(endpointConfigurations));
        }

        /**
         * @param endpointGroupRegion The ID of the region where the endpoint group is deployed.
         * 
         * @return builder
         * 
         */
        public Builder endpointGroupRegion(Output<String> endpointGroupRegion) {
            $.endpointGroupRegion = endpointGroupRegion;
            return this;
        }

        /**
         * @param endpointGroupRegion The ID of the region where the endpoint group is deployed.
         * 
         * @return builder
         * 
         */
        public Builder endpointGroupRegion(String endpointGroupRegion) {
            return endpointGroupRegion(Output.of(endpointGroupRegion));
        }

        /**
         * @param endpointGroupType The endpoint group type. Default value: `default`. Valid values: `default`, `virtual`.
         * &gt; **NOTE:** Currently, only `HTTP` or `HTTPS` protocol listener can directly create a `virtual` Endpoint Group. If it is `TCP` protocol listener, and you want to create a `virtual` Endpoint Group, please ensure that the `default` Endpoint Group has been created.
         * 
         * @return builder
         * 
         */
        public Builder endpointGroupType(@Nullable Output<String> endpointGroupType) {
            $.endpointGroupType = endpointGroupType;
            return this;
        }

        /**
         * @param endpointGroupType The endpoint group type. Default value: `default`. Valid values: `default`, `virtual`.
         * &gt; **NOTE:** Currently, only `HTTP` or `HTTPS` protocol listener can directly create a `virtual` Endpoint Group. If it is `TCP` protocol listener, and you want to create a `virtual` Endpoint Group, please ensure that the `default` Endpoint Group has been created.
         * 
         * @return builder
         * 
         */
        public Builder endpointGroupType(String endpointGroupType) {
            return endpointGroupType(Output.of(endpointGroupType));
        }

        /**
         * @param endpointRequestProtocol The endpoint request protocol. Valid values: `HTTP`, `HTTPS`.
         * &gt; **NOTE:** This item is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. For the listening instance of HTTP protocol, the back-end service protocol supports and only supports HTTP.
         * 
         * @return builder
         * 
         */
        public Builder endpointRequestProtocol(@Nullable Output<String> endpointRequestProtocol) {
            $.endpointRequestProtocol = endpointRequestProtocol;
            return this;
        }

        /**
         * @param endpointRequestProtocol The endpoint request protocol. Valid values: `HTTP`, `HTTPS`.
         * &gt; **NOTE:** This item is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. For the listening instance of HTTP protocol, the back-end service protocol supports and only supports HTTP.
         * 
         * @return builder
         * 
         */
        public Builder endpointRequestProtocol(String endpointRequestProtocol) {
            return endpointRequestProtocol(Output.of(endpointRequestProtocol));
        }

        /**
         * @param healthCheckEnabled Specifies whether to enable the health check feature. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder healthCheckEnabled(@Nullable Output<Boolean> healthCheckEnabled) {
            $.healthCheckEnabled = healthCheckEnabled;
            return this;
        }

        /**
         * @param healthCheckEnabled Specifies whether to enable the health check feature. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder healthCheckEnabled(Boolean healthCheckEnabled) {
            return healthCheckEnabled(Output.of(healthCheckEnabled));
        }

        /**
         * @param healthCheckIntervalSeconds The interval between two consecutive health checks. Unit: seconds.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckIntervalSeconds(@Nullable Output<Integer> healthCheckIntervalSeconds) {
            $.healthCheckIntervalSeconds = healthCheckIntervalSeconds;
            return this;
        }

        /**
         * @param healthCheckIntervalSeconds The interval between two consecutive health checks. Unit: seconds.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckIntervalSeconds(Integer healthCheckIntervalSeconds) {
            return healthCheckIntervalSeconds(Output.of(healthCheckIntervalSeconds));
        }

        /**
         * @param healthCheckPath The path specified as the destination of the targets for health checks.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckPath(@Nullable Output<String> healthCheckPath) {
            $.healthCheckPath = healthCheckPath;
            return this;
        }

        /**
         * @param healthCheckPath The path specified as the destination of the targets for health checks.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckPath(String healthCheckPath) {
            return healthCheckPath(Output.of(healthCheckPath));
        }

        /**
         * @param healthCheckPort The port that is used for health checks.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckPort(@Nullable Output<Integer> healthCheckPort) {
            $.healthCheckPort = healthCheckPort;
            return this;
        }

        /**
         * @param healthCheckPort The port that is used for health checks.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckPort(Integer healthCheckPort) {
            return healthCheckPort(Output.of(healthCheckPort));
        }

        /**
         * @param healthCheckProtocol The protocol that is used to connect to the targets for health checks. Valid values:
         * - `TCP` or `tcp`: TCP protocol.
         * - `HTTP` or `http`: HTTP protocol.
         * - `HTTPS` or `https`: HTTPS protocol.
         * &gt; **NOTE:** From version 1.223.0, `health_check_protocol` can be set to `TCP`, `HTTP`, `HTTPS`.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckProtocol(@Nullable Output<String> healthCheckProtocol) {
            $.healthCheckProtocol = healthCheckProtocol;
            return this;
        }

        /**
         * @param healthCheckProtocol The protocol that is used to connect to the targets for health checks. Valid values:
         * - `TCP` or `tcp`: TCP protocol.
         * - `HTTP` or `http`: HTTP protocol.
         * - `HTTPS` or `https`: HTTPS protocol.
         * &gt; **NOTE:** From version 1.223.0, `health_check_protocol` can be set to `TCP`, `HTTP`, `HTTPS`.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckProtocol(String healthCheckProtocol) {
            return healthCheckProtocol(Output.of(healthCheckProtocol));
        }

        /**
         * @param listenerId The ID of the listener that is associated with the endpoint group.
         * 
         * @return builder
         * 
         */
        public Builder listenerId(Output<String> listenerId) {
            $.listenerId = listenerId;
            return this;
        }

        /**
         * @param listenerId The ID of the listener that is associated with the endpoint group.
         * 
         * @return builder
         * 
         */
        public Builder listenerId(String listenerId) {
            return listenerId(Output.of(listenerId));
        }

        /**
         * @param name The name of the endpoint group.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the endpoint group.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param portOverrides Mapping between listening port and forwarding port of boarding point. See `port_overrides` below.
         * &gt; **NOTE:** Port mapping is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. The listening port in the port map must be consistent with the listening port of the current listening instance.
         * 
         * @return builder
         * 
         */
        public Builder portOverrides(@Nullable Output<EndpointGroupPortOverridesArgs> portOverrides) {
            $.portOverrides = portOverrides;
            return this;
        }

        /**
         * @param portOverrides Mapping between listening port and forwarding port of boarding point. See `port_overrides` below.
         * &gt; **NOTE:** Port mapping is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. The listening port in the port map must be consistent with the listening port of the current listening instance.
         * 
         * @return builder
         * 
         */
        public Builder portOverrides(EndpointGroupPortOverridesArgs portOverrides) {
            return portOverrides(Output.of(portOverrides));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param thresholdCount The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value: `3`.
         * 
         * @return builder
         * 
         */
        public Builder thresholdCount(@Nullable Output<Integer> thresholdCount) {
            $.thresholdCount = thresholdCount;
            return this;
        }

        /**
         * @param thresholdCount The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value: `3`.
         * 
         * @return builder
         * 
         */
        public Builder thresholdCount(Integer thresholdCount) {
            return thresholdCount(Output.of(thresholdCount));
        }

        /**
         * @param trafficPercentage The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
         * 
         * @return builder
         * 
         */
        public Builder trafficPercentage(@Nullable Output<Integer> trafficPercentage) {
            $.trafficPercentage = trafficPercentage;
            return this;
        }

        /**
         * @param trafficPercentage The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
         * 
         * @return builder
         * 
         */
        public Builder trafficPercentage(Integer trafficPercentage) {
            return trafficPercentage(Output.of(trafficPercentage));
        }

        public EndpointGroupArgs build() {
            if ($.acceleratorId == null) {
                throw new MissingRequiredPropertyException("EndpointGroupArgs", "acceleratorId");
            }
            if ($.endpointConfigurations == null) {
                throw new MissingRequiredPropertyException("EndpointGroupArgs", "endpointConfigurations");
            }
            if ($.endpointGroupRegion == null) {
                throw new MissingRequiredPropertyException("EndpointGroupArgs", "endpointGroupRegion");
            }
            if ($.listenerId == null) {
                throw new MissingRequiredPropertyException("EndpointGroupArgs", "listenerId");
            }
            return $;
        }
    }

}
