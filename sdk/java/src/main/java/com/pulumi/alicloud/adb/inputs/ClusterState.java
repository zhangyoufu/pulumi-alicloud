// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.adb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterState extends com.pulumi.resources.ResourceArgs {

    public static final ClusterState Empty = new ClusterState();

    /**
     * Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
     * 
     */
    @Import(name="autoRenewPeriod")
    private @Nullable Output<Integer> autoRenewPeriod;

    /**
     * @return Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
     * 
     */
    public Optional<Output<Integer>> autoRenewPeriod() {
        return Optional.ofNullable(this.autoRenewPeriod);
    }

    @Import(name="computeResource")
    private @Nullable Output<String> computeResource;

    public Optional<Output<String>> computeResource() {
        return Optional.ofNullable(this.computeResource);
    }

    /**
     * (Available in 1.93.0+) The connection string of the ADB cluster.
     * 
     */
    @Import(name="connectionString")
    private @Nullable Output<String> connectionString;

    /**
     * @return (Available in 1.93.0+) The connection string of the ADB cluster.
     * 
     */
    public Optional<Output<String>> connectionString() {
        return Optional.ofNullable(this.connectionString);
    }

    /**
     * Cluster category. Value options: `Basic`, `Cluster`.
     * 
     */
    @Import(name="dbClusterCategory")
    private @Nullable Output<String> dbClusterCategory;

    /**
     * @return Cluster category. Value options: `Basic`, `Cluster`.
     * 
     */
    public Optional<Output<String>> dbClusterCategory() {
        return Optional.ofNullable(this.dbClusterCategory);
    }

    /**
     * @deprecated
     * It duplicates with attribute db_node_class and is deprecated from 1.121.2.
     * 
     */
    @Deprecated /* It duplicates with attribute db_node_class and is deprecated from 1.121.2. */
    @Import(name="dbClusterClass")
    private @Nullable Output<String> dbClusterClass;

    /**
     * @deprecated
     * It duplicates with attribute db_node_class and is deprecated from 1.121.2.
     * 
     */
    @Deprecated /* It duplicates with attribute db_node_class and is deprecated from 1.121.2. */
    public Optional<Output<String>> dbClusterClass() {
        return Optional.ofNullable(this.dbClusterClass);
    }

    /**
     * Cluster version. Value options: `3.0`, Default to `3.0`.
     * 
     */
    @Import(name="dbClusterVersion")
    private @Nullable Output<String> dbClusterVersion;

    /**
     * @return Cluster version. Value options: `3.0`, Default to `3.0`.
     * 
     */
    public Optional<Output<String>> dbClusterVersion() {
        return Optional.ofNullable(this.dbClusterVersion);
    }

    /**
     * The db_node_class of cluster node.
     * 
     */
    @Import(name="dbNodeClass")
    private @Nullable Output<String> dbNodeClass;

    /**
     * @return The db_node_class of cluster node.
     * 
     */
    public Optional<Output<String>> dbNodeClass() {
        return Optional.ofNullable(this.dbNodeClass);
    }

    /**
     * The db_node_count of cluster node.
     * 
     */
    @Import(name="dbNodeCount")
    private @Nullable Output<Integer> dbNodeCount;

    /**
     * @return The db_node_count of cluster node.
     * 
     */
    public Optional<Output<Integer>> dbNodeCount() {
        return Optional.ofNullable(this.dbNodeCount);
    }

    /**
     * The db_node_storage of cluster node.
     * 
     */
    @Import(name="dbNodeStorage")
    private @Nullable Output<Integer> dbNodeStorage;

    /**
     * @return The db_node_storage of cluster node.
     * 
     */
    public Optional<Output<Integer>> dbNodeStorage() {
        return Optional.ofNullable(this.dbNodeStorage);
    }

    /**
     * The description of cluster.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of cluster.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="diskPerformanceLevel")
    private @Nullable Output<String> diskPerformanceLevel;

    public Optional<Output<String>> diskPerformanceLevel() {
        return Optional.ofNullable(this.diskPerformanceLevel);
    }

    @Import(name="elasticIoResource")
    private @Nullable Output<Integer> elasticIoResource;

    public Optional<Output<Integer>> elasticIoResource() {
        return Optional.ofNullable(this.elasticIoResource);
    }

    @Import(name="elasticIoResourceSize")
    private @Nullable Output<String> elasticIoResourceSize;

    public Optional<Output<String>> elasticIoResourceSize() {
        return Optional.ofNullable(this.elasticIoResourceSize);
    }

    /**
     * Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
     * 
     */
    @Import(name="maintainTime")
    private @Nullable Output<String> maintainTime;

    /**
     * @return Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
     * 
     */
    public Optional<Output<String>> maintainTime() {
        return Optional.ofNullable(this.maintainTime);
    }

    @Import(name="mode")
    private @Nullable Output<String> mode;

    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    @Import(name="modifyType")
    private @Nullable Output<String> modifyType;

    public Optional<Output<String>> modifyType() {
        return Optional.ofNullable(this.modifyType);
    }

    /**
     * Field `pay_type` has been deprecated. New field `payment_type` instead.
     * 
     * @deprecated
     * Attribute &#39;pay_type&#39; has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute &#39;payment_type&#39; instead.
     * 
     */
    @Deprecated /* Attribute 'pay_type' has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead. */
    @Import(name="payType")
    private @Nullable Output<String> payType;

    /**
     * @return Field `pay_type` has been deprecated. New field `payment_type` instead.
     * 
     * @deprecated
     * Attribute &#39;pay_type&#39; has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute &#39;payment_type&#39; instead.
     * 
     */
    @Deprecated /* Attribute 'pay_type' has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead. */
    public Optional<Output<String>> payType() {
        return Optional.ofNullable(this.payType);
    }

    /**
     * The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `payment_type` supports updating from v1.166.0+.
     * 
     */
    @Import(name="paymentType")
    private @Nullable Output<String> paymentType;

    /**
     * @return The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `payment_type` supports updating from v1.166.0+.
     * 
     */
    public Optional<Output<String>> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }

    /**
     * The duration that you will buy DB cluster (in month). It is valid when pay_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
     * 
     */
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return The duration that you will buy DB cluster (in month). It is valid when pay_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * (Available in 1.196.0+) The connection port of the ADB cluster.
     * 
     */
    @Import(name="port")
    private @Nullable Output<String> port;

    /**
     * @return (Available in 1.196.0+) The connection port of the ADB cluster.
     * 
     */
    public Optional<Output<String>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
     * 
     */
    @Import(name="renewalStatus")
    private @Nullable Output<String> renewalStatus;

    /**
     * @return Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
     * 
     */
    public Optional<Output<String>> renewalStatus() {
        return Optional.ofNullable(this.renewalStatus);
    }

    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     * 
     */
    @Import(name="securityIps")
    private @Nullable Output<List<String>> securityIps;

    /**
     * @return List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     * 
     */
    public Optional<Output<List<String>>> securityIps() {
        return Optional.ofNullable(this.securityIps);
    }

    @Import(name="status")
    private @Nullable Output<String> status;

    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     * &gt; **NOTE:** Because of data backup and migration, change DB cluster type and storage would cost 15~30 minutes. Please make full preparation before changing them.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     * &gt; **NOTE:** Because of data backup and migration, change DB cluster type and storage would cost 15~30 minutes. Please make full preparation before changing them.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * The virtual switch ID to launch DB instances in one VPC.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable Output<String> vswitchId;

    /**
     * @return The virtual switch ID to launch DB instances in one VPC.
     * 
     */
    public Optional<Output<String>> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    /**
     * The Zone to launch the DB cluster.
     * 
     */
    @Import(name="zoneId")
    private @Nullable Output<String> zoneId;

    /**
     * @return The Zone to launch the DB cluster.
     * 
     */
    public Optional<Output<String>> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    private ClusterState() {}

    private ClusterState(ClusterState $) {
        this.autoRenewPeriod = $.autoRenewPeriod;
        this.computeResource = $.computeResource;
        this.connectionString = $.connectionString;
        this.dbClusterCategory = $.dbClusterCategory;
        this.dbClusterClass = $.dbClusterClass;
        this.dbClusterVersion = $.dbClusterVersion;
        this.dbNodeClass = $.dbNodeClass;
        this.dbNodeCount = $.dbNodeCount;
        this.dbNodeStorage = $.dbNodeStorage;
        this.description = $.description;
        this.diskPerformanceLevel = $.diskPerformanceLevel;
        this.elasticIoResource = $.elasticIoResource;
        this.elasticIoResourceSize = $.elasticIoResourceSize;
        this.maintainTime = $.maintainTime;
        this.mode = $.mode;
        this.modifyType = $.modifyType;
        this.payType = $.payType;
        this.paymentType = $.paymentType;
        this.period = $.period;
        this.port = $.port;
        this.renewalStatus = $.renewalStatus;
        this.resourceGroupId = $.resourceGroupId;
        this.securityIps = $.securityIps;
        this.status = $.status;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
        this.vswitchId = $.vswitchId;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterState $;

        public Builder() {
            $ = new ClusterState();
        }

        public Builder(ClusterState defaults) {
            $ = new ClusterState(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoRenewPeriod Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
         * 
         * @return builder
         * 
         */
        public Builder autoRenewPeriod(@Nullable Output<Integer> autoRenewPeriod) {
            $.autoRenewPeriod = autoRenewPeriod;
            return this;
        }

        /**
         * @param autoRenewPeriod Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
         * 
         * @return builder
         * 
         */
        public Builder autoRenewPeriod(Integer autoRenewPeriod) {
            return autoRenewPeriod(Output.of(autoRenewPeriod));
        }

        public Builder computeResource(@Nullable Output<String> computeResource) {
            $.computeResource = computeResource;
            return this;
        }

        public Builder computeResource(String computeResource) {
            return computeResource(Output.of(computeResource));
        }

        /**
         * @param connectionString (Available in 1.93.0+) The connection string of the ADB cluster.
         * 
         * @return builder
         * 
         */
        public Builder connectionString(@Nullable Output<String> connectionString) {
            $.connectionString = connectionString;
            return this;
        }

        /**
         * @param connectionString (Available in 1.93.0+) The connection string of the ADB cluster.
         * 
         * @return builder
         * 
         */
        public Builder connectionString(String connectionString) {
            return connectionString(Output.of(connectionString));
        }

        /**
         * @param dbClusterCategory Cluster category. Value options: `Basic`, `Cluster`.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterCategory(@Nullable Output<String> dbClusterCategory) {
            $.dbClusterCategory = dbClusterCategory;
            return this;
        }

        /**
         * @param dbClusterCategory Cluster category. Value options: `Basic`, `Cluster`.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterCategory(String dbClusterCategory) {
            return dbClusterCategory(Output.of(dbClusterCategory));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * It duplicates with attribute db_node_class and is deprecated from 1.121.2.
         * 
         */
        @Deprecated /* It duplicates with attribute db_node_class and is deprecated from 1.121.2. */
        public Builder dbClusterClass(@Nullable Output<String> dbClusterClass) {
            $.dbClusterClass = dbClusterClass;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * It duplicates with attribute db_node_class and is deprecated from 1.121.2.
         * 
         */
        @Deprecated /* It duplicates with attribute db_node_class and is deprecated from 1.121.2. */
        public Builder dbClusterClass(String dbClusterClass) {
            return dbClusterClass(Output.of(dbClusterClass));
        }

        /**
         * @param dbClusterVersion Cluster version. Value options: `3.0`, Default to `3.0`.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterVersion(@Nullable Output<String> dbClusterVersion) {
            $.dbClusterVersion = dbClusterVersion;
            return this;
        }

        /**
         * @param dbClusterVersion Cluster version. Value options: `3.0`, Default to `3.0`.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterVersion(String dbClusterVersion) {
            return dbClusterVersion(Output.of(dbClusterVersion));
        }

        /**
         * @param dbNodeClass The db_node_class of cluster node.
         * 
         * @return builder
         * 
         */
        public Builder dbNodeClass(@Nullable Output<String> dbNodeClass) {
            $.dbNodeClass = dbNodeClass;
            return this;
        }

        /**
         * @param dbNodeClass The db_node_class of cluster node.
         * 
         * @return builder
         * 
         */
        public Builder dbNodeClass(String dbNodeClass) {
            return dbNodeClass(Output.of(dbNodeClass));
        }

        /**
         * @param dbNodeCount The db_node_count of cluster node.
         * 
         * @return builder
         * 
         */
        public Builder dbNodeCount(@Nullable Output<Integer> dbNodeCount) {
            $.dbNodeCount = dbNodeCount;
            return this;
        }

        /**
         * @param dbNodeCount The db_node_count of cluster node.
         * 
         * @return builder
         * 
         */
        public Builder dbNodeCount(Integer dbNodeCount) {
            return dbNodeCount(Output.of(dbNodeCount));
        }

        /**
         * @param dbNodeStorage The db_node_storage of cluster node.
         * 
         * @return builder
         * 
         */
        public Builder dbNodeStorage(@Nullable Output<Integer> dbNodeStorage) {
            $.dbNodeStorage = dbNodeStorage;
            return this;
        }

        /**
         * @param dbNodeStorage The db_node_storage of cluster node.
         * 
         * @return builder
         * 
         */
        public Builder dbNodeStorage(Integer dbNodeStorage) {
            return dbNodeStorage(Output.of(dbNodeStorage));
        }

        /**
         * @param description The description of cluster.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of cluster.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder diskPerformanceLevel(@Nullable Output<String> diskPerformanceLevel) {
            $.diskPerformanceLevel = diskPerformanceLevel;
            return this;
        }

        public Builder diskPerformanceLevel(String diskPerformanceLevel) {
            return diskPerformanceLevel(Output.of(diskPerformanceLevel));
        }

        public Builder elasticIoResource(@Nullable Output<Integer> elasticIoResource) {
            $.elasticIoResource = elasticIoResource;
            return this;
        }

        public Builder elasticIoResource(Integer elasticIoResource) {
            return elasticIoResource(Output.of(elasticIoResource));
        }

        public Builder elasticIoResourceSize(@Nullable Output<String> elasticIoResourceSize) {
            $.elasticIoResourceSize = elasticIoResourceSize;
            return this;
        }

        public Builder elasticIoResourceSize(String elasticIoResourceSize) {
            return elasticIoResourceSize(Output.of(elasticIoResourceSize));
        }

        /**
         * @param maintainTime Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
         * 
         * @return builder
         * 
         */
        public Builder maintainTime(@Nullable Output<String> maintainTime) {
            $.maintainTime = maintainTime;
            return this;
        }

        /**
         * @param maintainTime Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
         * 
         * @return builder
         * 
         */
        public Builder maintainTime(String maintainTime) {
            return maintainTime(Output.of(maintainTime));
        }

        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        public Builder modifyType(@Nullable Output<String> modifyType) {
            $.modifyType = modifyType;
            return this;
        }

        public Builder modifyType(String modifyType) {
            return modifyType(Output.of(modifyType));
        }

        /**
         * @param payType Field `pay_type` has been deprecated. New field `payment_type` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Attribute &#39;pay_type&#39; has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute &#39;payment_type&#39; instead.
         * 
         */
        @Deprecated /* Attribute 'pay_type' has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead. */
        public Builder payType(@Nullable Output<String> payType) {
            $.payType = payType;
            return this;
        }

        /**
         * @param payType Field `pay_type` has been deprecated. New field `payment_type` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Attribute &#39;pay_type&#39; has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute &#39;payment_type&#39; instead.
         * 
         */
        @Deprecated /* Attribute 'pay_type' has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead. */
        public Builder payType(String payType) {
            return payType(Output.of(payType));
        }

        /**
         * @param paymentType The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `payment_type` supports updating from v1.166.0+.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(@Nullable Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param paymentType The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `payment_type` supports updating from v1.166.0+.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(String paymentType) {
            return paymentType(Output.of(paymentType));
        }

        /**
         * @param period The duration that you will buy DB cluster (in month). It is valid when pay_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period The duration that you will buy DB cluster (in month). It is valid when pay_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param port (Available in 1.196.0+) The connection port of the ADB cluster.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<String> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port (Available in 1.196.0+) The connection port of the ADB cluster.
         * 
         * @return builder
         * 
         */
        public Builder port(String port) {
            return port(Output.of(port));
        }

        /**
         * @param renewalStatus Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(@Nullable Output<String> renewalStatus) {
            $.renewalStatus = renewalStatus;
            return this;
        }

        /**
         * @param renewalStatus Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(String renewalStatus) {
            return renewalStatus(Output.of(renewalStatus));
        }

        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param securityIps List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
         * 
         * @return builder
         * 
         */
        public Builder securityIps(@Nullable Output<List<String>> securityIps) {
            $.securityIps = securityIps;
            return this;
        }

        /**
         * @param securityIps List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
         * 
         * @return builder
         * 
         */
        public Builder securityIps(List<String> securityIps) {
            return securityIps(Output.of(securityIps));
        }

        /**
         * @param securityIps List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
         * 
         * @return builder
         * 
         */
        public Builder securityIps(String... securityIps) {
            return securityIps(List.of(securityIps));
        }

        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
         * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
         * 
         * &gt; **NOTE:** Because of data backup and migration, change DB cluster type and storage would cost 15~30 minutes. Please make full preparation before changing them.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
         * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
         * 
         * &gt; **NOTE:** Because of data backup and migration, change DB cluster type and storage would cost 15~30 minutes. Please make full preparation before changing them.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param vswitchId The virtual switch ID to launch DB instances in one VPC.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId The virtual switch ID to launch DB instances in one VPC.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        /**
         * @param zoneId The Zone to launch the DB cluster.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(@Nullable Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        /**
         * @param zoneId The Zone to launch the DB cluster.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public ClusterState build() {
            return $;
        }
    }

}
