// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cassandra;

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


public final class ClusterArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterArgs Empty = new ClusterArgs();

    /**
     * Auto renew of dataCenter-1,`true` or `false`. System default to `false`, valid when pay_type = PrePaid.
     * 
     */
    @Import(name="autoRenew")
    private @Nullable Output<Boolean> autoRenew;

    /**
     * @return Auto renew of dataCenter-1,`true` or `false`. System default to `false`, valid when pay_type = PrePaid.
     * 
     */
    public Optional<Output<Boolean>> autoRenew() {
        return Optional.ofNullable(this.autoRenew);
    }

    /**
     * Period of dataCenter-1 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when pay_type = Subscription. Unit: month.
     * 
     */
    @Import(name="autoRenewPeriod")
    private @Nullable Output<Integer> autoRenewPeriod;

    /**
     * @return Period of dataCenter-1 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when pay_type = Subscription. Unit: month.
     * 
     */
    public Optional<Output<Integer>> autoRenewPeriod() {
        return Optional.ofNullable(this.autoRenewPeriod);
    }

    /**
     * Cassandra cluster name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
     * 
     */
    @Import(name="clusterName")
    private @Nullable Output<String> clusterName;

    /**
     * @return Cassandra cluster name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
     * 
     */
    public Optional<Output<String>> clusterName() {
        return Optional.ofNullable(this.clusterName);
    }

    /**
     * Cassandra dataCenter-1 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
     * 
     */
    @Import(name="dataCenterName")
    private @Nullable Output<String> dataCenterName;

    /**
     * @return Cassandra dataCenter-1 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
     * 
     */
    public Optional<Output<String>> dataCenterName() {
        return Optional.ofNullable(this.dataCenterName);
    }

    /**
     * User-defined Cassandra dataCenter-1 one node&#39;s storage space.Unit: GB. Value range:
     * - Custom storage space; value range: [160, 2000].
     * - 80-GB increments.
     * 
     */
    @Import(name="diskSize")
    private @Nullable Output<Integer> diskSize;

    /**
     * @return User-defined Cassandra dataCenter-1 one node&#39;s storage space.Unit: GB. Value range:
     * - Custom storage space; value range: [160, 2000].
     * - 80-GB increments.
     * 
     */
    public Optional<Output<Integer>> diskSize() {
        return Optional.ofNullable(this.diskSize);
    }

    /**
     * The disk type of Cassandra dataCenter-1. Valid values are `cloud_ssd`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`, local_disk size is fixed.
     * 
     */
    @Import(name="diskType")
    private @Nullable Output<String> diskType;

    /**
     * @return The disk type of Cassandra dataCenter-1. Valid values are `cloud_ssd`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`, local_disk size is fixed.
     * 
     */
    public Optional<Output<String>> diskType() {
        return Optional.ofNullable(this.diskType);
    }

    @Import(name="enablePublic")
    private @Nullable Output<Boolean> enablePublic;

    public Optional<Output<Boolean>> enablePublic() {
        return Optional.ofNullable(this.enablePublic);
    }

    /**
     * Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.
     * 
     */
    @Import(name="instanceType", required=true)
    private Output<String> instanceType;

    /**
     * @return Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }

    /**
     * Set the instance&#39;s IP whitelist in VPC network.
     * 
     */
    @Import(name="ipWhite")
    private @Nullable Output<String> ipWhite;

    /**
     * @return Set the instance&#39;s IP whitelist in VPC network.
     * 
     */
    public Optional<Output<String>> ipWhite() {
        return Optional.ofNullable(this.ipWhite);
    }

    /**
     * The end time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).
     * 
     */
    @Import(name="maintainEndTime")
    private @Nullable Output<String> maintainEndTime;

    /**
     * @return The end time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).
     * 
     */
    public Optional<Output<String>> maintainEndTime() {
        return Optional.ofNullable(this.maintainEndTime);
    }

    /**
     * The start time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).
     * 
     */
    @Import(name="maintainStartTime")
    private @Nullable Output<String> maintainStartTime;

    /**
     * @return The start time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).
     * 
     */
    public Optional<Output<String>> maintainStartTime() {
        return Optional.ofNullable(this.maintainStartTime);
    }

    /**
     * Cassandra major version. Now only support version `3.11`.
     * 
     */
    @Import(name="majorVersion", required=true)
    private Output<String> majorVersion;

    /**
     * @return Cassandra major version. Now only support version `3.11`.
     * 
     */
    public Output<String> majorVersion() {
        return this.majorVersion;
    }

    /**
     * The node count of Cassandra dataCenter-1 default to 2.
     * 
     */
    @Import(name="nodeCount", required=true)
    private Output<Integer> nodeCount;

    /**
     * @return The node count of Cassandra dataCenter-1 default to 2.
     * 
     */
    public Output<Integer> nodeCount() {
        return this.nodeCount;
    }

    @Import(name="password")
    private @Nullable Output<String> password;

    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * The pay type of Cassandra dataCenter-1. Valid values are `Subscription`, `PayAsYouGo`,System default to `PayAsYouGo`.
     * 
     */
    @Import(name="payType", required=true)
    private Output<String> payType;

    /**
     * @return The pay type of Cassandra dataCenter-1. Valid values are `Subscription`, `PayAsYouGo`,System default to `PayAsYouGo`.
     * 
     */
    public Output<String> payType() {
        return this.payType;
    }

    @Import(name="period")
    private @Nullable Output<Integer> period;

    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    @Import(name="periodUnit")
    private @Nullable Output<String> periodUnit;

    public Optional<Output<String>> periodUnit() {
        return Optional.ofNullable(this.periodUnit);
    }

    /**
     * A list of security group ids to associate with.
     * 
     * &gt; **NOTE:** Now cluster_name,data_center_name,instance_type,node_count,disk_type,disk_size,maintain_start_time,maintain_end_time,tags,ip_white,security_groups can be change. The others(auto_renew, auto_renew_period and so on) will be supported in the furture.
     * 
     */
    @Import(name="securityGroups")
    private @Nullable Output<List<String>> securityGroups;

    /**
     * @return A list of security group ids to associate with.
     * 
     * &gt; **NOTE:** Now cluster_name,data_center_name,instance_type,node_count,disk_type,disk_size,maintain_start_time,maintain_end_time,tags,ip_white,security_groups can be change. The others(auto_renew, auto_renew_period and so on) will be supported in the furture.
     * 
     */
    public Optional<Output<List<String>>> securityGroups() {
        return Optional.ofNullable(this.securityGroups);
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
     * The vswitch_id of dataCenter-1, can not empty.
     * 
     */
    @Import(name="vswitchId", required=true)
    private Output<String> vswitchId;

    /**
     * @return The vswitch_id of dataCenter-1, can not empty.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    /**
     * The Zone to launch the Cassandra cluster. If vswitch_id is not empty, this zone_id can be &#34;&#34; or consistent.
     * 
     */
    @Import(name="zoneId")
    private @Nullable Output<String> zoneId;

    /**
     * @return The Zone to launch the Cassandra cluster. If vswitch_id is not empty, this zone_id can be &#34;&#34; or consistent.
     * 
     */
    public Optional<Output<String>> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    private ClusterArgs() {}

    private ClusterArgs(ClusterArgs $) {
        this.autoRenew = $.autoRenew;
        this.autoRenewPeriod = $.autoRenewPeriod;
        this.clusterName = $.clusterName;
        this.dataCenterName = $.dataCenterName;
        this.diskSize = $.diskSize;
        this.diskType = $.diskType;
        this.enablePublic = $.enablePublic;
        this.instanceType = $.instanceType;
        this.ipWhite = $.ipWhite;
        this.maintainEndTime = $.maintainEndTime;
        this.maintainStartTime = $.maintainStartTime;
        this.majorVersion = $.majorVersion;
        this.nodeCount = $.nodeCount;
        this.password = $.password;
        this.payType = $.payType;
        this.period = $.period;
        this.periodUnit = $.periodUnit;
        this.securityGroups = $.securityGroups;
        this.tags = $.tags;
        this.vswitchId = $.vswitchId;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterArgs $;

        public Builder() {
            $ = new ClusterArgs();
        }

        public Builder(ClusterArgs defaults) {
            $ = new ClusterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoRenew Auto renew of dataCenter-1,`true` or `false`. System default to `false`, valid when pay_type = PrePaid.
         * 
         * @return builder
         * 
         */
        public Builder autoRenew(@Nullable Output<Boolean> autoRenew) {
            $.autoRenew = autoRenew;
            return this;
        }

        /**
         * @param autoRenew Auto renew of dataCenter-1,`true` or `false`. System default to `false`, valid when pay_type = PrePaid.
         * 
         * @return builder
         * 
         */
        public Builder autoRenew(Boolean autoRenew) {
            return autoRenew(Output.of(autoRenew));
        }

        /**
         * @param autoRenewPeriod Period of dataCenter-1 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when pay_type = Subscription. Unit: month.
         * 
         * @return builder
         * 
         */
        public Builder autoRenewPeriod(@Nullable Output<Integer> autoRenewPeriod) {
            $.autoRenewPeriod = autoRenewPeriod;
            return this;
        }

        /**
         * @param autoRenewPeriod Period of dataCenter-1 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when pay_type = Subscription. Unit: month.
         * 
         * @return builder
         * 
         */
        public Builder autoRenewPeriod(Integer autoRenewPeriod) {
            return autoRenewPeriod(Output.of(autoRenewPeriod));
        }

        /**
         * @param clusterName Cassandra cluster name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(@Nullable Output<String> clusterName) {
            $.clusterName = clusterName;
            return this;
        }

        /**
         * @param clusterName Cassandra cluster name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(String clusterName) {
            return clusterName(Output.of(clusterName));
        }

        /**
         * @param dataCenterName Cassandra dataCenter-1 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
         * 
         * @return builder
         * 
         */
        public Builder dataCenterName(@Nullable Output<String> dataCenterName) {
            $.dataCenterName = dataCenterName;
            return this;
        }

        /**
         * @param dataCenterName Cassandra dataCenter-1 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
         * 
         * @return builder
         * 
         */
        public Builder dataCenterName(String dataCenterName) {
            return dataCenterName(Output.of(dataCenterName));
        }

        /**
         * @param diskSize User-defined Cassandra dataCenter-1 one node&#39;s storage space.Unit: GB. Value range:
         * - Custom storage space; value range: [160, 2000].
         * - 80-GB increments.
         * 
         * @return builder
         * 
         */
        public Builder diskSize(@Nullable Output<Integer> diskSize) {
            $.diskSize = diskSize;
            return this;
        }

        /**
         * @param diskSize User-defined Cassandra dataCenter-1 one node&#39;s storage space.Unit: GB. Value range:
         * - Custom storage space; value range: [160, 2000].
         * - 80-GB increments.
         * 
         * @return builder
         * 
         */
        public Builder diskSize(Integer diskSize) {
            return diskSize(Output.of(diskSize));
        }

        /**
         * @param diskType The disk type of Cassandra dataCenter-1. Valid values are `cloud_ssd`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`, local_disk size is fixed.
         * 
         * @return builder
         * 
         */
        public Builder diskType(@Nullable Output<String> diskType) {
            $.diskType = diskType;
            return this;
        }

        /**
         * @param diskType The disk type of Cassandra dataCenter-1. Valid values are `cloud_ssd`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`, local_disk size is fixed.
         * 
         * @return builder
         * 
         */
        public Builder diskType(String diskType) {
            return diskType(Output.of(diskType));
        }

        public Builder enablePublic(@Nullable Output<Boolean> enablePublic) {
            $.enablePublic = enablePublic;
            return this;
        }

        public Builder enablePublic(Boolean enablePublic) {
            return enablePublic(Output.of(enablePublic));
        }

        /**
         * @param instanceType Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        /**
         * @param ipWhite Set the instance&#39;s IP whitelist in VPC network.
         * 
         * @return builder
         * 
         */
        public Builder ipWhite(@Nullable Output<String> ipWhite) {
            $.ipWhite = ipWhite;
            return this;
        }

        /**
         * @param ipWhite Set the instance&#39;s IP whitelist in VPC network.
         * 
         * @return builder
         * 
         */
        public Builder ipWhite(String ipWhite) {
            return ipWhite(Output.of(ipWhite));
        }

        /**
         * @param maintainEndTime The end time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).
         * 
         * @return builder
         * 
         */
        public Builder maintainEndTime(@Nullable Output<String> maintainEndTime) {
            $.maintainEndTime = maintainEndTime;
            return this;
        }

        /**
         * @param maintainEndTime The end time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).
         * 
         * @return builder
         * 
         */
        public Builder maintainEndTime(String maintainEndTime) {
            return maintainEndTime(Output.of(maintainEndTime));
        }

        /**
         * @param maintainStartTime The start time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).
         * 
         * @return builder
         * 
         */
        public Builder maintainStartTime(@Nullable Output<String> maintainStartTime) {
            $.maintainStartTime = maintainStartTime;
            return this;
        }

        /**
         * @param maintainStartTime The start time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).
         * 
         * @return builder
         * 
         */
        public Builder maintainStartTime(String maintainStartTime) {
            return maintainStartTime(Output.of(maintainStartTime));
        }

        /**
         * @param majorVersion Cassandra major version. Now only support version `3.11`.
         * 
         * @return builder
         * 
         */
        public Builder majorVersion(Output<String> majorVersion) {
            $.majorVersion = majorVersion;
            return this;
        }

        /**
         * @param majorVersion Cassandra major version. Now only support version `3.11`.
         * 
         * @return builder
         * 
         */
        public Builder majorVersion(String majorVersion) {
            return majorVersion(Output.of(majorVersion));
        }

        /**
         * @param nodeCount The node count of Cassandra dataCenter-1 default to 2.
         * 
         * @return builder
         * 
         */
        public Builder nodeCount(Output<Integer> nodeCount) {
            $.nodeCount = nodeCount;
            return this;
        }

        /**
         * @param nodeCount The node count of Cassandra dataCenter-1 default to 2.
         * 
         * @return builder
         * 
         */
        public Builder nodeCount(Integer nodeCount) {
            return nodeCount(Output.of(nodeCount));
        }

        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param payType The pay type of Cassandra dataCenter-1. Valid values are `Subscription`, `PayAsYouGo`,System default to `PayAsYouGo`.
         * 
         * @return builder
         * 
         */
        public Builder payType(Output<String> payType) {
            $.payType = payType;
            return this;
        }

        /**
         * @param payType The pay type of Cassandra dataCenter-1. Valid values are `Subscription`, `PayAsYouGo`,System default to `PayAsYouGo`.
         * 
         * @return builder
         * 
         */
        public Builder payType(String payType) {
            return payType(Output.of(payType));
        }

        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        public Builder periodUnit(@Nullable Output<String> periodUnit) {
            $.periodUnit = periodUnit;
            return this;
        }

        public Builder periodUnit(String periodUnit) {
            return periodUnit(Output.of(periodUnit));
        }

        /**
         * @param securityGroups A list of security group ids to associate with.
         * 
         * &gt; **NOTE:** Now cluster_name,data_center_name,instance_type,node_count,disk_type,disk_size,maintain_start_time,maintain_end_time,tags,ip_white,security_groups can be change. The others(auto_renew, auto_renew_period and so on) will be supported in the furture.
         * 
         * @return builder
         * 
         */
        public Builder securityGroups(@Nullable Output<List<String>> securityGroups) {
            $.securityGroups = securityGroups;
            return this;
        }

        /**
         * @param securityGroups A list of security group ids to associate with.
         * 
         * &gt; **NOTE:** Now cluster_name,data_center_name,instance_type,node_count,disk_type,disk_size,maintain_start_time,maintain_end_time,tags,ip_white,security_groups can be change. The others(auto_renew, auto_renew_period and so on) will be supported in the furture.
         * 
         * @return builder
         * 
         */
        public Builder securityGroups(List<String> securityGroups) {
            return securityGroups(Output.of(securityGroups));
        }

        /**
         * @param securityGroups A list of security group ids to associate with.
         * 
         * &gt; **NOTE:** Now cluster_name,data_center_name,instance_type,node_count,disk_type,disk_size,maintain_start_time,maintain_end_time,tags,ip_white,security_groups can be change. The others(auto_renew, auto_renew_period and so on) will be supported in the furture.
         * 
         * @return builder
         * 
         */
        public Builder securityGroups(String... securityGroups) {
            return securityGroups(List.of(securityGroups));
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
         * @param vswitchId The vswitch_id of dataCenter-1, can not empty.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId The vswitch_id of dataCenter-1, can not empty.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        /**
         * @param zoneId The Zone to launch the Cassandra cluster. If vswitch_id is not empty, this zone_id can be &#34;&#34; or consistent.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(@Nullable Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        /**
         * @param zoneId The Zone to launch the Cassandra cluster. If vswitch_id is not empty, this zone_id can be &#34;&#34; or consistent.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public ClusterArgs build() {
            if ($.instanceType == null) {
                throw new MissingRequiredPropertyException("ClusterArgs", "instanceType");
            }
            if ($.majorVersion == null) {
                throw new MissingRequiredPropertyException("ClusterArgs", "majorVersion");
            }
            if ($.nodeCount == null) {
                throw new MissingRequiredPropertyException("ClusterArgs", "nodeCount");
            }
            if ($.payType == null) {
                throw new MissingRequiredPropertyException("ClusterArgs", "payType");
            }
            if ($.vswitchId == null) {
                throw new MissingRequiredPropertyException("ClusterArgs", "vswitchId");
            }
            return $;
        }
    }

}
