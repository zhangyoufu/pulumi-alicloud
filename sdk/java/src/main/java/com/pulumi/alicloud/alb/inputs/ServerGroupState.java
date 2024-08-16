// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.inputs;

import com.pulumi.alicloud.alb.inputs.ServerGroupHealthCheckConfigArgs;
import com.pulumi.alicloud.alb.inputs.ServerGroupServerArgs;
import com.pulumi.alicloud.alb.inputs.ServerGroupStickySessionConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServerGroupState extends com.pulumi.resources.ResourceArgs {

    public static final ServerGroupState Empty = new ServerGroupState();

    /**
     * The dry run.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return The dry run.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * The configuration of health checks. See `health_check_config` below.
     * 
     */
    @Import(name="healthCheckConfig")
    private @Nullable Output<ServerGroupHealthCheckConfigArgs> healthCheckConfig;

    /**
     * @return The configuration of health checks. See `health_check_config` below.
     * 
     */
    public Optional<Output<ServerGroupHealthCheckConfigArgs>> healthCheckConfig() {
        return Optional.ofNullable(this.healthCheckConfig);
    }

    /**
     * The server protocol. Valid values: `  HTTP `, `HTTPS`, `gRPC`. While `server_group_type` is `Fc` this parameter will not take effect. From version 1.215.0, `protocol` can be set to `gRPC`.
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return The server protocol. Valid values: `  HTTP `, `HTTPS`, `gRPC`. While `server_group_type` is `Fc` this parameter will not take effect. From version 1.215.0, `protocol` can be set to `gRPC`.
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * The ID of the resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`. **NOTE:** This parameter takes effect when the `server_group_type` parameter is set to `Instance` or `Ip`.
     * 
     */
    @Import(name="scheduler")
    private @Nullable Output<String> scheduler;

    /**
     * @return The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`. **NOTE:** This parameter takes effect when the `server_group_type` parameter is set to `Instance` or `Ip`.
     * 
     */
    public Optional<Output<String>> scheduler() {
        return Optional.ofNullable(this.scheduler);
    }

    /**
     * The name of the server group.
     * 
     */
    @Import(name="serverGroupName")
    private @Nullable Output<String> serverGroupName;

    /**
     * @return The name of the server group.
     * 
     */
    public Optional<Output<String>> serverGroupName() {
        return Optional.ofNullable(this.serverGroupName);
    }

    /**
     * The type of the server group. Default value: `Instance`. Valid values:
     * - `Instance`: allows you add servers by specifying Ecs, Ens, or Eci.
     * - `Ip`: allows you to add servers by specifying IP addresses.
     * - `Fc`: allows you to add servers by specifying functions of Function Compute.
     * 
     */
    @Import(name="serverGroupType")
    private @Nullable Output<String> serverGroupType;

    /**
     * @return The type of the server group. Default value: `Instance`. Valid values:
     * - `Instance`: allows you add servers by specifying Ecs, Ens, or Eci.
     * - `Ip`: allows you to add servers by specifying IP addresses.
     * - `Fc`: allows you to add servers by specifying functions of Function Compute.
     * 
     */
    public Optional<Output<String>> serverGroupType() {
        return Optional.ofNullable(this.serverGroupType);
    }

    /**
     * The backend servers. See `servers` below.
     * 
     */
    @Import(name="servers")
    private @Nullable Output<List<ServerGroupServerArgs>> servers;

    /**
     * @return The backend servers. See `servers` below.
     * 
     */
    public Optional<Output<List<ServerGroupServerArgs>>> servers() {
        return Optional.ofNullable(this.servers);
    }

    /**
     * The status of the backend server.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the backend server.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The configuration of session persistence. See `sticky_session_config` below.
     * 
     */
    @Import(name="stickySessionConfig")
    private @Nullable Output<ServerGroupStickySessionConfigArgs> stickySessionConfig;

    /**
     * @return The configuration of session persistence. See `sticky_session_config` below.
     * 
     */
    public Optional<Output<ServerGroupStickySessionConfigArgs>> stickySessionConfig() {
        return Optional.ofNullable(this.stickySessionConfig);
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
     * The ID of the VPC that you want to access. **NOTE:** This parameter takes effect when the `server_group_type` parameter is set to `Instance` or `Ip`.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The ID of the VPC that you want to access. **NOTE:** This parameter takes effect when the `server_group_type` parameter is set to `Instance` or `Ip`.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    private ServerGroupState() {}

    private ServerGroupState(ServerGroupState $) {
        this.dryRun = $.dryRun;
        this.healthCheckConfig = $.healthCheckConfig;
        this.protocol = $.protocol;
        this.resourceGroupId = $.resourceGroupId;
        this.scheduler = $.scheduler;
        this.serverGroupName = $.serverGroupName;
        this.serverGroupType = $.serverGroupType;
        this.servers = $.servers;
        this.status = $.status;
        this.stickySessionConfig = $.stickySessionConfig;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerGroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerGroupState $;

        public Builder() {
            $ = new ServerGroupState();
        }

        public Builder(ServerGroupState defaults) {
            $ = new ServerGroupState(Objects.requireNonNull(defaults));
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param healthCheckConfig The configuration of health checks. See `health_check_config` below.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckConfig(@Nullable Output<ServerGroupHealthCheckConfigArgs> healthCheckConfig) {
            $.healthCheckConfig = healthCheckConfig;
            return this;
        }

        /**
         * @param healthCheckConfig The configuration of health checks. See `health_check_config` below.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckConfig(ServerGroupHealthCheckConfigArgs healthCheckConfig) {
            return healthCheckConfig(Output.of(healthCheckConfig));
        }

        /**
         * @param protocol The server protocol. Valid values: `  HTTP `, `HTTPS`, `gRPC`. While `server_group_type` is `Fc` this parameter will not take effect. From version 1.215.0, `protocol` can be set to `gRPC`.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol The server protocol. Valid values: `  HTTP `, `HTTPS`, `gRPC`. While `server_group_type` is `Fc` this parameter will not take effect. From version 1.215.0, `protocol` can be set to `gRPC`.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param scheduler The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`. **NOTE:** This parameter takes effect when the `server_group_type` parameter is set to `Instance` or `Ip`.
         * 
         * @return builder
         * 
         */
        public Builder scheduler(@Nullable Output<String> scheduler) {
            $.scheduler = scheduler;
            return this;
        }

        /**
         * @param scheduler The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`. **NOTE:** This parameter takes effect when the `server_group_type` parameter is set to `Instance` or `Ip`.
         * 
         * @return builder
         * 
         */
        public Builder scheduler(String scheduler) {
            return scheduler(Output.of(scheduler));
        }

        /**
         * @param serverGroupName The name of the server group.
         * 
         * @return builder
         * 
         */
        public Builder serverGroupName(@Nullable Output<String> serverGroupName) {
            $.serverGroupName = serverGroupName;
            return this;
        }

        /**
         * @param serverGroupName The name of the server group.
         * 
         * @return builder
         * 
         */
        public Builder serverGroupName(String serverGroupName) {
            return serverGroupName(Output.of(serverGroupName));
        }

        /**
         * @param serverGroupType The type of the server group. Default value: `Instance`. Valid values:
         * - `Instance`: allows you add servers by specifying Ecs, Ens, or Eci.
         * - `Ip`: allows you to add servers by specifying IP addresses.
         * - `Fc`: allows you to add servers by specifying functions of Function Compute.
         * 
         * @return builder
         * 
         */
        public Builder serverGroupType(@Nullable Output<String> serverGroupType) {
            $.serverGroupType = serverGroupType;
            return this;
        }

        /**
         * @param serverGroupType The type of the server group. Default value: `Instance`. Valid values:
         * - `Instance`: allows you add servers by specifying Ecs, Ens, or Eci.
         * - `Ip`: allows you to add servers by specifying IP addresses.
         * - `Fc`: allows you to add servers by specifying functions of Function Compute.
         * 
         * @return builder
         * 
         */
        public Builder serverGroupType(String serverGroupType) {
            return serverGroupType(Output.of(serverGroupType));
        }

        /**
         * @param servers The backend servers. See `servers` below.
         * 
         * @return builder
         * 
         */
        public Builder servers(@Nullable Output<List<ServerGroupServerArgs>> servers) {
            $.servers = servers;
            return this;
        }

        /**
         * @param servers The backend servers. See `servers` below.
         * 
         * @return builder
         * 
         */
        public Builder servers(List<ServerGroupServerArgs> servers) {
            return servers(Output.of(servers));
        }

        /**
         * @param servers The backend servers. See `servers` below.
         * 
         * @return builder
         * 
         */
        public Builder servers(ServerGroupServerArgs... servers) {
            return servers(List.of(servers));
        }

        /**
         * @param status The status of the backend server.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the backend server.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param stickySessionConfig The configuration of session persistence. See `sticky_session_config` below.
         * 
         * @return builder
         * 
         */
        public Builder stickySessionConfig(@Nullable Output<ServerGroupStickySessionConfigArgs> stickySessionConfig) {
            $.stickySessionConfig = stickySessionConfig;
            return this;
        }

        /**
         * @param stickySessionConfig The configuration of session persistence. See `sticky_session_config` below.
         * 
         * @return builder
         * 
         */
        public Builder stickySessionConfig(ServerGroupStickySessionConfigArgs stickySessionConfig) {
            return stickySessionConfig(Output.of(stickySessionConfig));
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
         * @param vpcId The ID of the VPC that you want to access. **NOTE:** This parameter takes effect when the `server_group_type` parameter is set to `Instance` or `Ip`.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the VPC that you want to access. **NOTE:** This parameter takes effect when the `server_group_type` parameter is set to `Instance` or `Ip`.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        public ServerGroupState build() {
            return $;
        }
    }

}
