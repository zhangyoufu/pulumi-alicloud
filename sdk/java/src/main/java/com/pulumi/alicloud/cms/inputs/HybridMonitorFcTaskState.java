// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HybridMonitorFcTaskState extends com.pulumi.resources.ResourceArgs {

    public static final HybridMonitorFcTaskState Empty = new HybridMonitorFcTaskState();

    /**
     * The ID of the monitoring task.
     * 
     */
    @Import(name="hybridMonitorFcTaskId")
    private @Nullable Output<String> hybridMonitorFcTaskId;

    /**
     * @return The ID of the monitoring task.
     * 
     */
    public Optional<Output<String>> hybridMonitorFcTaskId() {
        return Optional.ofNullable(this.hybridMonitorFcTaskId);
    }

    /**
     * the namespace of the Alibaba Cloud service.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return the namespace of the Alibaba Cloud service.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The ID of the member account. If you call API operations by using a management account, you can connect the Alibaba Cloud services that are activated for a member account in Resource Directory to Hybrid Cloud Monitoring. You can use Resource Directory to monitor Alibaba Cloud services across enterprise accounts.
     * 
     */
    @Import(name="targetUserId")
    private @Nullable Output<String> targetUserId;

    /**
     * @return The ID of the member account. If you call API operations by using a management account, you can connect the Alibaba Cloud services that are activated for a member account in Resource Directory to Hybrid Cloud Monitoring. You can use Resource Directory to monitor Alibaba Cloud services across enterprise accounts.
     * 
     */
    public Optional<Output<String>> targetUserId() {
        return Optional.ofNullable(this.targetUserId);
    }

    /**
     * The configuration file of the Alibaba Cloud service that you want to monitor by using Hybrid Cloud Monitoring.
     * 
     */
    @Import(name="yarmConfig")
    private @Nullable Output<String> yarmConfig;

    /**
     * @return The configuration file of the Alibaba Cloud service that you want to monitor by using Hybrid Cloud Monitoring.
     * 
     */
    public Optional<Output<String>> yarmConfig() {
        return Optional.ofNullable(this.yarmConfig);
    }

    private HybridMonitorFcTaskState() {}

    private HybridMonitorFcTaskState(HybridMonitorFcTaskState $) {
        this.hybridMonitorFcTaskId = $.hybridMonitorFcTaskId;
        this.namespace = $.namespace;
        this.targetUserId = $.targetUserId;
        this.yarmConfig = $.yarmConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HybridMonitorFcTaskState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HybridMonitorFcTaskState $;

        public Builder() {
            $ = new HybridMonitorFcTaskState();
        }

        public Builder(HybridMonitorFcTaskState defaults) {
            $ = new HybridMonitorFcTaskState(Objects.requireNonNull(defaults));
        }

        /**
         * @param hybridMonitorFcTaskId The ID of the monitoring task.
         * 
         * @return builder
         * 
         */
        public Builder hybridMonitorFcTaskId(@Nullable Output<String> hybridMonitorFcTaskId) {
            $.hybridMonitorFcTaskId = hybridMonitorFcTaskId;
            return this;
        }

        /**
         * @param hybridMonitorFcTaskId The ID of the monitoring task.
         * 
         * @return builder
         * 
         */
        public Builder hybridMonitorFcTaskId(String hybridMonitorFcTaskId) {
            return hybridMonitorFcTaskId(Output.of(hybridMonitorFcTaskId));
        }

        /**
         * @param namespace the namespace of the Alibaba Cloud service.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace the namespace of the Alibaba Cloud service.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param targetUserId The ID of the member account. If you call API operations by using a management account, you can connect the Alibaba Cloud services that are activated for a member account in Resource Directory to Hybrid Cloud Monitoring. You can use Resource Directory to monitor Alibaba Cloud services across enterprise accounts.
         * 
         * @return builder
         * 
         */
        public Builder targetUserId(@Nullable Output<String> targetUserId) {
            $.targetUserId = targetUserId;
            return this;
        }

        /**
         * @param targetUserId The ID of the member account. If you call API operations by using a management account, you can connect the Alibaba Cloud services that are activated for a member account in Resource Directory to Hybrid Cloud Monitoring. You can use Resource Directory to monitor Alibaba Cloud services across enterprise accounts.
         * 
         * @return builder
         * 
         */
        public Builder targetUserId(String targetUserId) {
            return targetUserId(Output.of(targetUserId));
        }

        /**
         * @param yarmConfig The configuration file of the Alibaba Cloud service that you want to monitor by using Hybrid Cloud Monitoring.
         * 
         * @return builder
         * 
         */
        public Builder yarmConfig(@Nullable Output<String> yarmConfig) {
            $.yarmConfig = yarmConfig;
            return this;
        }

        /**
         * @param yarmConfig The configuration file of the Alibaba Cloud service that you want to monitor by using Hybrid Cloud Monitoring.
         * 
         * @return builder
         * 
         */
        public Builder yarmConfig(String yarmConfig) {
            return yarmConfig(Output.of(yarmConfig));
        }

        public HybridMonitorFcTaskState build() {
            return $;
        }
    }

}
