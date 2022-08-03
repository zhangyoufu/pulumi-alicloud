// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HybridMonitorFcTaskArgs extends com.pulumi.resources.ResourceArgs {

    public static final HybridMonitorFcTaskArgs Empty = new HybridMonitorFcTaskArgs();

    /**
     * The index warehouse where the host belongs.
     * 
     */
    @Import(name="namespace", required=true)
    private Output<String> namespace;

    /**
     * @return The index warehouse where the host belongs.
     * 
     */
    public Output<String> namespace() {
        return this.namespace;
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
     * - `namespace`: the namespace of the Alibaba Cloud service.
     * - `metric_list`: the metrics of the Alibaba Cloud service.
     * 
     */
    @Import(name="yarmConfig", required=true)
    private Output<String> yarmConfig;

    /**
     * @return The configuration file of the Alibaba Cloud service that you want to monitor by using Hybrid Cloud Monitoring.
     * - `namespace`: the namespace of the Alibaba Cloud service.
     * - `metric_list`: the metrics of the Alibaba Cloud service.
     * 
     */
    public Output<String> yarmConfig() {
        return this.yarmConfig;
    }

    private HybridMonitorFcTaskArgs() {}

    private HybridMonitorFcTaskArgs(HybridMonitorFcTaskArgs $) {
        this.namespace = $.namespace;
        this.targetUserId = $.targetUserId;
        this.yarmConfig = $.yarmConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HybridMonitorFcTaskArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HybridMonitorFcTaskArgs $;

        public Builder() {
            $ = new HybridMonitorFcTaskArgs();
        }

        public Builder(HybridMonitorFcTaskArgs defaults) {
            $ = new HybridMonitorFcTaskArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param namespace The index warehouse where the host belongs.
         * 
         * @return builder
         * 
         */
        public Builder namespace(Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace The index warehouse where the host belongs.
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
         * - `namespace`: the namespace of the Alibaba Cloud service.
         * - `metric_list`: the metrics of the Alibaba Cloud service.
         * 
         * @return builder
         * 
         */
        public Builder yarmConfig(Output<String> yarmConfig) {
            $.yarmConfig = yarmConfig;
            return this;
        }

        /**
         * @param yarmConfig The configuration file of the Alibaba Cloud service that you want to monitor by using Hybrid Cloud Monitoring.
         * - `namespace`: the namespace of the Alibaba Cloud service.
         * - `metric_list`: the metrics of the Alibaba Cloud service.
         * 
         * @return builder
         * 
         */
        public Builder yarmConfig(String yarmConfig) {
            return yarmConfig(Output.of(yarmConfig));
        }

        public HybridMonitorFcTaskArgs build() {
            $.namespace = Objects.requireNonNull($.namespace, "expected parameter 'namespace' to be non-null");
            $.yarmConfig = Objects.requireNonNull($.yarmConfig, "expected parameter 'yarmConfig' to be non-null");
            return $;
        }
    }

}
