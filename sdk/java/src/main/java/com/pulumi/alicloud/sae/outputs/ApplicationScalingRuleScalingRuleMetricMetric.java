// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ApplicationScalingRuleScalingRuleMetricMetric {
    /**
     * @return According to different `metric_type`, set the target value of the corresponding monitoring index.
     * 
     */
    private @Nullable Integer metricTargetAverageUtilization;
    /**
     * @return Monitoring indicator trigger condition. Valid values: `CPU`, `MEMORY`, `tcpActiveConn`, `QPS`, `RT`, `SLB_QPS`, `SLB_RT`, `INTRANET_SLB_QPS` and `INTRANET_SLB_RT`. The values are described as follows:
     * - CPU: CPU usage.
     * - MEMORY: MEMORY usage.
     * - tcpActiveConn: The average number of TCP active connections for a single instance in 30 seconds.
     * - QPS: The average QPS of a single instance within 1 minute of JAVA application.
     * - RT: The average response time of all service interfaces within 1 minute of JAVA application.
     * - SLB_QPS: The average public network SLB QPS of a single instance within 15 seconds.
     * - SLB_RT: The average response time of public network SLB within 15 seconds.
     * - INTRANET_SLB_QPS: The average private network SLB QPS of a single instance within 15 seconds.
     * - INTRANET_SLB_RT: The average response time of private network SLB within 15 seconds.
     *   **NOTE:** From version 1.206.0, `metric_type` can be set to `QPS`, `RT`, `INTRANET_SLB_QPS`, `INTRANET_SLB_RT`.
     * 
     */
    private @Nullable String metricType;
    /**
     * @return SLB ID.
     * 
     */
    private @Nullable String slbId;
    /**
     * @return The log store of the Log Service.
     * 
     */
    private @Nullable String slbLogStore;
    /**
     * @return The project of the Log Service.
     * 
     */
    private @Nullable String slbProject;
    /**
     * @return SLB listening port.
     * 
     */
    private @Nullable String vport;

    private ApplicationScalingRuleScalingRuleMetricMetric() {}
    /**
     * @return According to different `metric_type`, set the target value of the corresponding monitoring index.
     * 
     */
    public Optional<Integer> metricTargetAverageUtilization() {
        return Optional.ofNullable(this.metricTargetAverageUtilization);
    }
    /**
     * @return Monitoring indicator trigger condition. Valid values: `CPU`, `MEMORY`, `tcpActiveConn`, `QPS`, `RT`, `SLB_QPS`, `SLB_RT`, `INTRANET_SLB_QPS` and `INTRANET_SLB_RT`. The values are described as follows:
     * - CPU: CPU usage.
     * - MEMORY: MEMORY usage.
     * - tcpActiveConn: The average number of TCP active connections for a single instance in 30 seconds.
     * - QPS: The average QPS of a single instance within 1 minute of JAVA application.
     * - RT: The average response time of all service interfaces within 1 minute of JAVA application.
     * - SLB_QPS: The average public network SLB QPS of a single instance within 15 seconds.
     * - SLB_RT: The average response time of public network SLB within 15 seconds.
     * - INTRANET_SLB_QPS: The average private network SLB QPS of a single instance within 15 seconds.
     * - INTRANET_SLB_RT: The average response time of private network SLB within 15 seconds.
     *   **NOTE:** From version 1.206.0, `metric_type` can be set to `QPS`, `RT`, `INTRANET_SLB_QPS`, `INTRANET_SLB_RT`.
     * 
     */
    public Optional<String> metricType() {
        return Optional.ofNullable(this.metricType);
    }
    /**
     * @return SLB ID.
     * 
     */
    public Optional<String> slbId() {
        return Optional.ofNullable(this.slbId);
    }
    /**
     * @return The log store of the Log Service.
     * 
     */
    public Optional<String> slbLogStore() {
        return Optional.ofNullable(this.slbLogStore);
    }
    /**
     * @return The project of the Log Service.
     * 
     */
    public Optional<String> slbProject() {
        return Optional.ofNullable(this.slbProject);
    }
    /**
     * @return SLB listening port.
     * 
     */
    public Optional<String> vport() {
        return Optional.ofNullable(this.vport);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApplicationScalingRuleScalingRuleMetricMetric defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer metricTargetAverageUtilization;
        private @Nullable String metricType;
        private @Nullable String slbId;
        private @Nullable String slbLogStore;
        private @Nullable String slbProject;
        private @Nullable String vport;
        public Builder() {}
        public Builder(ApplicationScalingRuleScalingRuleMetricMetric defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.metricTargetAverageUtilization = defaults.metricTargetAverageUtilization;
    	      this.metricType = defaults.metricType;
    	      this.slbId = defaults.slbId;
    	      this.slbLogStore = defaults.slbLogStore;
    	      this.slbProject = defaults.slbProject;
    	      this.vport = defaults.vport;
        }

        @CustomType.Setter
        public Builder metricTargetAverageUtilization(@Nullable Integer metricTargetAverageUtilization) {
            this.metricTargetAverageUtilization = metricTargetAverageUtilization;
            return this;
        }
        @CustomType.Setter
        public Builder metricType(@Nullable String metricType) {
            this.metricType = metricType;
            return this;
        }
        @CustomType.Setter
        public Builder slbId(@Nullable String slbId) {
            this.slbId = slbId;
            return this;
        }
        @CustomType.Setter
        public Builder slbLogStore(@Nullable String slbLogStore) {
            this.slbLogStore = slbLogStore;
            return this;
        }
        @CustomType.Setter
        public Builder slbProject(@Nullable String slbProject) {
            this.slbProject = slbProject;
            return this;
        }
        @CustomType.Setter
        public Builder vport(@Nullable String vport) {
            this.vport = vport;
            return this;
        }
        public ApplicationScalingRuleScalingRuleMetricMetric build() {
            final var o = new ApplicationScalingRuleScalingRuleMetricMetric();
            o.metricTargetAverageUtilization = metricTargetAverageUtilization;
            o.metricType = metricType;
            o.slbId = slbId;
            o.slbLogStore = slbLogStore;
            o.slbProject = slbProject;
            o.vport = vport;
            return o;
        }
    }
}
