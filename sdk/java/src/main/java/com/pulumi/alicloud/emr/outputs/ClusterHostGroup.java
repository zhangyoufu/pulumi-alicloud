// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.emr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClusterHostGroup {
    /**
     * @return Auto renew for prepaid, ’true’ or ‘false’ . Default value: false.
     * 
     */
    private @Nullable Boolean autoRenew;
    /**
     * @return Charge Type for this cluster. Supported value: PostPaid or PrePaid. Default value: PostPaid.
     * 
     */
    private @Nullable String chargeType;
    /**
     * @return Graceful decommission timeout, unit: seconds.
     * 
     */
    private @Nullable Integer decommissionTimeout;
    /**
     * @return Data disk capacity.
     * 
     */
    private @Nullable String diskCapacity;
    /**
     * @return Data disk count.
     * 
     */
    private @Nullable String diskCount;
    /**
     * @return Data disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,local_disk,cloud_essd.
     * 
     */
    private @Nullable String diskType;
    /**
     * @return Enable hadoop cluster of task node graceful decommission, ’true’ or ‘false’ . Default value: false.
     * 
     */
    private @Nullable Boolean enableGracefulDecommission;
    private @Nullable String gpuDriver;
    /**
     * @return host group name.
     * 
     */
    private @Nullable String hostGroupName;
    /**
     * @return host group type, supported value: MASTER, CORE or TASK, supported &#39;GATEWAY&#39; available in 1.61.0+.
     * 
     */
    private @Nullable String hostGroupType;
    /**
     * @return Instance list for cluster scale down. This value follows the json format, e.g. [&#34;instance_id1&#34;,&#34;instance_id2&#34;]. escape character for &#34; is \&#34;.
     * 
     */
    private @Nullable String instanceList;
    /**
     * @return Host Ecs instance type.
     * 
     */
    private @Nullable String instanceType;
    /**
     * @return Host number in this group.
     * 
     */
    private @Nullable String nodeCount;
    /**
     * @return If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
     * 
     */
    private @Nullable Integer period;
    /**
     * @return System disk capacity.
     * 
     */
    private @Nullable String sysDiskCapacity;
    /**
     * @return System disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,cloud_essd.
     * 
     */
    private @Nullable String sysDiskType;

    private ClusterHostGroup() {}
    /**
     * @return Auto renew for prepaid, ’true’ or ‘false’ . Default value: false.
     * 
     */
    public Optional<Boolean> autoRenew() {
        return Optional.ofNullable(this.autoRenew);
    }
    /**
     * @return Charge Type for this cluster. Supported value: PostPaid or PrePaid. Default value: PostPaid.
     * 
     */
    public Optional<String> chargeType() {
        return Optional.ofNullable(this.chargeType);
    }
    /**
     * @return Graceful decommission timeout, unit: seconds.
     * 
     */
    public Optional<Integer> decommissionTimeout() {
        return Optional.ofNullable(this.decommissionTimeout);
    }
    /**
     * @return Data disk capacity.
     * 
     */
    public Optional<String> diskCapacity() {
        return Optional.ofNullable(this.diskCapacity);
    }
    /**
     * @return Data disk count.
     * 
     */
    public Optional<String> diskCount() {
        return Optional.ofNullable(this.diskCount);
    }
    /**
     * @return Data disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,local_disk,cloud_essd.
     * 
     */
    public Optional<String> diskType() {
        return Optional.ofNullable(this.diskType);
    }
    /**
     * @return Enable hadoop cluster of task node graceful decommission, ’true’ or ‘false’ . Default value: false.
     * 
     */
    public Optional<Boolean> enableGracefulDecommission() {
        return Optional.ofNullable(this.enableGracefulDecommission);
    }
    public Optional<String> gpuDriver() {
        return Optional.ofNullable(this.gpuDriver);
    }
    /**
     * @return host group name.
     * 
     */
    public Optional<String> hostGroupName() {
        return Optional.ofNullable(this.hostGroupName);
    }
    /**
     * @return host group type, supported value: MASTER, CORE or TASK, supported &#39;GATEWAY&#39; available in 1.61.0+.
     * 
     */
    public Optional<String> hostGroupType() {
        return Optional.ofNullable(this.hostGroupType);
    }
    /**
     * @return Instance list for cluster scale down. This value follows the json format, e.g. [&#34;instance_id1&#34;,&#34;instance_id2&#34;]. escape character for &#34; is \&#34;.
     * 
     */
    public Optional<String> instanceList() {
        return Optional.ofNullable(this.instanceList);
    }
    /**
     * @return Host Ecs instance type.
     * 
     */
    public Optional<String> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }
    /**
     * @return Host number in this group.
     * 
     */
    public Optional<String> nodeCount() {
        return Optional.ofNullable(this.nodeCount);
    }
    /**
     * @return If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
     * 
     */
    public Optional<Integer> period() {
        return Optional.ofNullable(this.period);
    }
    /**
     * @return System disk capacity.
     * 
     */
    public Optional<String> sysDiskCapacity() {
        return Optional.ofNullable(this.sysDiskCapacity);
    }
    /**
     * @return System disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,cloud_essd.
     * 
     */
    public Optional<String> sysDiskType() {
        return Optional.ofNullable(this.sysDiskType);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterHostGroup defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean autoRenew;
        private @Nullable String chargeType;
        private @Nullable Integer decommissionTimeout;
        private @Nullable String diskCapacity;
        private @Nullable String diskCount;
        private @Nullable String diskType;
        private @Nullable Boolean enableGracefulDecommission;
        private @Nullable String gpuDriver;
        private @Nullable String hostGroupName;
        private @Nullable String hostGroupType;
        private @Nullable String instanceList;
        private @Nullable String instanceType;
        private @Nullable String nodeCount;
        private @Nullable Integer period;
        private @Nullable String sysDiskCapacity;
        private @Nullable String sysDiskType;
        public Builder() {}
        public Builder(ClusterHostGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.autoRenew = defaults.autoRenew;
    	      this.chargeType = defaults.chargeType;
    	      this.decommissionTimeout = defaults.decommissionTimeout;
    	      this.diskCapacity = defaults.diskCapacity;
    	      this.diskCount = defaults.diskCount;
    	      this.diskType = defaults.diskType;
    	      this.enableGracefulDecommission = defaults.enableGracefulDecommission;
    	      this.gpuDriver = defaults.gpuDriver;
    	      this.hostGroupName = defaults.hostGroupName;
    	      this.hostGroupType = defaults.hostGroupType;
    	      this.instanceList = defaults.instanceList;
    	      this.instanceType = defaults.instanceType;
    	      this.nodeCount = defaults.nodeCount;
    	      this.period = defaults.period;
    	      this.sysDiskCapacity = defaults.sysDiskCapacity;
    	      this.sysDiskType = defaults.sysDiskType;
        }

        @CustomType.Setter
        public Builder autoRenew(@Nullable Boolean autoRenew) {
            this.autoRenew = autoRenew;
            return this;
        }
        @CustomType.Setter
        public Builder chargeType(@Nullable String chargeType) {
            this.chargeType = chargeType;
            return this;
        }
        @CustomType.Setter
        public Builder decommissionTimeout(@Nullable Integer decommissionTimeout) {
            this.decommissionTimeout = decommissionTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder diskCapacity(@Nullable String diskCapacity) {
            this.diskCapacity = diskCapacity;
            return this;
        }
        @CustomType.Setter
        public Builder diskCount(@Nullable String diskCount) {
            this.diskCount = diskCount;
            return this;
        }
        @CustomType.Setter
        public Builder diskType(@Nullable String diskType) {
            this.diskType = diskType;
            return this;
        }
        @CustomType.Setter
        public Builder enableGracefulDecommission(@Nullable Boolean enableGracefulDecommission) {
            this.enableGracefulDecommission = enableGracefulDecommission;
            return this;
        }
        @CustomType.Setter
        public Builder gpuDriver(@Nullable String gpuDriver) {
            this.gpuDriver = gpuDriver;
            return this;
        }
        @CustomType.Setter
        public Builder hostGroupName(@Nullable String hostGroupName) {
            this.hostGroupName = hostGroupName;
            return this;
        }
        @CustomType.Setter
        public Builder hostGroupType(@Nullable String hostGroupType) {
            this.hostGroupType = hostGroupType;
            return this;
        }
        @CustomType.Setter
        public Builder instanceList(@Nullable String instanceList) {
            this.instanceList = instanceList;
            return this;
        }
        @CustomType.Setter
        public Builder instanceType(@Nullable String instanceType) {
            this.instanceType = instanceType;
            return this;
        }
        @CustomType.Setter
        public Builder nodeCount(@Nullable String nodeCount) {
            this.nodeCount = nodeCount;
            return this;
        }
        @CustomType.Setter
        public Builder period(@Nullable Integer period) {
            this.period = period;
            return this;
        }
        @CustomType.Setter
        public Builder sysDiskCapacity(@Nullable String sysDiskCapacity) {
            this.sysDiskCapacity = sysDiskCapacity;
            return this;
        }
        @CustomType.Setter
        public Builder sysDiskType(@Nullable String sysDiskType) {
            this.sysDiskType = sysDiskType;
            return this;
        }
        public ClusterHostGroup build() {
            final var o = new ClusterHostGroup();
            o.autoRenew = autoRenew;
            o.chargeType = chargeType;
            o.decommissionTimeout = decommissionTimeout;
            o.diskCapacity = diskCapacity;
            o.diskCount = diskCount;
            o.diskType = diskType;
            o.enableGracefulDecommission = enableGracefulDecommission;
            o.gpuDriver = gpuDriver;
            o.hostGroupName = hostGroupName;
            o.hostGroupType = hostGroupType;
            o.instanceList = instanceList;
            o.instanceType = instanceType;
            o.nodeCount = nodeCount;
            o.period = period;
            o.sysDiskCapacity = sysDiskCapacity;
            o.sysDiskType = sysDiskType;
            return o;
        }
    }
}
