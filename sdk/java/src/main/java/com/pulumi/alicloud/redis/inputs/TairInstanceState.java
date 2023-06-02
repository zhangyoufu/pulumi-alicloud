// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.redis.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TairInstanceState extends com.pulumi.resources.ResourceArgs {

    public static final TairInstanceState Empty = new TairInstanceState();

    /**
     * Specifies whether to enable auto-renewal for the instance. Default value: false. Valid values: true(enables auto-renewal), false(disables auto-renewal).
     * 
     */
    @Import(name="autoRenew")
    private @Nullable Output<String> autoRenew;

    /**
     * @return Specifies whether to enable auto-renewal for the instance. Default value: false. Valid values: true(enables auto-renewal), false(disables auto-renewal).
     * 
     */
    public Optional<Output<String>> autoRenew() {
        return Optional.ofNullable(this.autoRenew);
    }

    /**
     * The subscription duration that is supported by auto-renewal. Unit: months. Valid values: 1, 2, 3, 6, and 12. This parameter is required only if the AutoRenew parameter is set to true.
     * 
     */
    @Import(name="autoRenewPeriod")
    private @Nullable Output<String> autoRenewPeriod;

    /**
     * @return The subscription duration that is supported by auto-renewal. Unit: months. Valid values: 1, 2, 3, 6, and 12. This parameter is required only if the AutoRenew parameter is set to true.
     * 
     */
    public Optional<Output<String>> autoRenewPeriod() {
        return Optional.ofNullable(this.autoRenewPeriod);
    }

    /**
     * The time when the instance was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The time when the instance was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * The time when to change the configurations. Default value: Immediately. Valid values: Immediately (The configurations are immediately changed), MaintainTime (The configurations are changed within the maintenance window).
     * 
     */
    @Import(name="effectiveTime")
    private @Nullable Output<String> effectiveTime;

    /**
     * @return The time when to change the configurations. Default value: Immediately. Valid values: Immediately (The configurations are immediately changed), MaintainTime (The configurations are changed within the maintenance window).
     * 
     */
    public Optional<Output<String>> effectiveTime() {
        return Optional.ofNullable(this.effectiveTime);
    }

    /**
     * The database engine version of the instance. Default value: 1.0. The default version is developed by Alibaba Cloud and compatible with Redis 5.0.
     * 
     */
    @Import(name="engineVersion")
    private @Nullable Output<String> engineVersion;

    /**
     * @return The database engine version of the instance. Default value: 1.0. The default version is developed by Alibaba Cloud and compatible with Redis 5.0.
     * 
     */
    public Optional<Output<String>> engineVersion() {
        return Optional.ofNullable(this.engineVersion);
    }

    /**
     * Specifies whether to forcefully change the configurations of the instance. Default value: true. Valid values: false (The system does not forcefully change the configurations), true (The system forcefully changes the configurations).
     * 
     */
    @Import(name="forceUpgrade")
    private @Nullable Output<Boolean> forceUpgrade;

    /**
     * @return Specifies whether to forcefully change the configurations of the instance. Default value: true. Valid values: false (The system does not forcefully change the configurations), true (The system forcefully changes the configurations).
     * 
     */
    public Optional<Output<Boolean>> forceUpgrade() {
        return Optional.ofNullable(this.forceUpgrade);
    }

    /**
     * The instance type of the instance. For more information, see [Instance types](https://www.alibabacloud.com/help/en/apsaradb-for-redis/latest/instance-types).
     * 
     */
    @Import(name="instanceClass")
    private @Nullable Output<String> instanceClass;

    /**
     * @return The instance type of the instance. For more information, see [Instance types](https://www.alibabacloud.com/help/en/apsaradb-for-redis/latest/instance-types).
     * 
     */
    public Optional<Output<String>> instanceClass() {
        return Optional.ofNullable(this.instanceClass);
    }

    /**
     * The storage medium of the instance. Valid values: tair_rdb, tair_scm, tair_essd.
     * 
     */
    @Import(name="instanceType")
    private @Nullable Output<String> instanceType;

    /**
     * @return The storage medium of the instance. Valid values: tair_rdb, tair_scm, tair_essd.
     * 
     */
    public Optional<Output<String>> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }

    /**
     * The password that is used to connect to the instance. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include ! @ # $ % ^ &amp; * ( ) _ + - =.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The password that is used to connect to the instance. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include ! @ # $ % ^ &amp; * ( ) _ + - =.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * Payment type: Subscription (prepaid), PayAsYouGo (postpaid). Default PayAsYouGo.
     * 
     */
    @Import(name="paymentType")
    private @Nullable Output<String> paymentType;

    /**
     * @return Payment type: Subscription (prepaid), PayAsYouGo (postpaid). Default PayAsYouGo.
     * 
     */
    public Optional<Output<String>> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }

    /**
     * The subscription duration. Unit: months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24,36, and 60. This parameter is required only if you set the PaymentType parameter to Subscription.
     * 
     */
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return The subscription duration. Unit: months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24,36, and 60. This parameter is required only if you set the PaymentType parameter to Subscription.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * The Tair service port. The service port of the instance. Valid values: 1024 to 65535. Default value: 6379.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return The Tair service port. The service port of the instance. Valid values: 1024 to 65535. Default value: 6379.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * The ID of the resource group to which the instance belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group to which the instance belongs.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The ID of the secondary zone.This parameter is returned only if the instance is deployed in two zones.
     * 
     */
    @Import(name="secondaryZoneId")
    private @Nullable Output<String> secondaryZoneId;

    /**
     * @return The ID of the secondary zone.This parameter is returned only if the instance is deployed in two zones.
     * 
     */
    public Optional<Output<String>> secondaryZoneId() {
        return Optional.ofNullable(this.secondaryZoneId);
    }

    /**
     * The number of data nodes in the instance. When 1 is passed, it means that the instance created is a standard architecture with only one data node. You can create an instance in the standard architecture that contains only a single data node. 2 to 32: You can create an instance in the cluster architecture that contains the specified number of data nodes. Only persistent memory-optimized instances can use the cluster architecture. Therefore, you can set this parameter to an integer from 2 to 32 only if you set the InstanceType parameter to tair_scm.
     * 
     */
    @Import(name="shardCount")
    private @Nullable Output<Integer> shardCount;

    /**
     * @return The number of data nodes in the instance. When 1 is passed, it means that the instance created is a standard architecture with only one data node. You can create an instance in the standard architecture that contains only a single data node. 2 to 32: You can create an instance in the cluster architecture that contains the specified number of data nodes. Only persistent memory-optimized instances can use the cluster architecture. Therefore, you can set this parameter to an integer from 2 to 32 only if you set the InstanceType parameter to tair_scm.
     * 
     */
    public Optional<Output<Integer>> shardCount() {
        return Optional.ofNullable(this.shardCount);
    }

    /**
     * The status of the instance. Valid values: Normal, Creating, Changing, Inactive, Flushing, Released, Transforming, Unavailable, Error, Migrating, and upgrading, networkModifying, SSLModifying, and MajorVersionUpgrading.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the instance. Valid values: Normal, Creating, Changing, Inactive, Flushing, Released, Transforming, Unavailable, Error, Migrating, and upgrading, networkModifying, SSLModifying, and MajorVersionUpgrading.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The name of the resource.
     * 
     */
    @Import(name="tairInstanceName")
    private @Nullable Output<String> tairInstanceName;

    /**
     * @return The name of the resource.
     * 
     */
    public Optional<Output<String>> tairInstanceName() {
        return Optional.ofNullable(this.tairInstanceName);
    }

    /**
     * The ID of the virtual private cloud (VPC).
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The ID of the virtual private cloud (VPC).
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * The ID of the vSwitch to which the instance is connected.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable Output<String> vswitchId;

    /**
     * @return The ID of the vSwitch to which the instance is connected.
     * 
     */
    public Optional<Output<String>> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    /**
     * The zone ID of the instance.
     * 
     */
    @Import(name="zoneId")
    private @Nullable Output<String> zoneId;

    /**
     * @return The zone ID of the instance.
     * 
     */
    public Optional<Output<String>> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    private TairInstanceState() {}

    private TairInstanceState(TairInstanceState $) {
        this.autoRenew = $.autoRenew;
        this.autoRenewPeriod = $.autoRenewPeriod;
        this.createTime = $.createTime;
        this.effectiveTime = $.effectiveTime;
        this.engineVersion = $.engineVersion;
        this.forceUpgrade = $.forceUpgrade;
        this.instanceClass = $.instanceClass;
        this.instanceType = $.instanceType;
        this.password = $.password;
        this.paymentType = $.paymentType;
        this.period = $.period;
        this.port = $.port;
        this.resourceGroupId = $.resourceGroupId;
        this.secondaryZoneId = $.secondaryZoneId;
        this.shardCount = $.shardCount;
        this.status = $.status;
        this.tairInstanceName = $.tairInstanceName;
        this.vpcId = $.vpcId;
        this.vswitchId = $.vswitchId;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TairInstanceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TairInstanceState $;

        public Builder() {
            $ = new TairInstanceState();
        }

        public Builder(TairInstanceState defaults) {
            $ = new TairInstanceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoRenew Specifies whether to enable auto-renewal for the instance. Default value: false. Valid values: true(enables auto-renewal), false(disables auto-renewal).
         * 
         * @return builder
         * 
         */
        public Builder autoRenew(@Nullable Output<String> autoRenew) {
            $.autoRenew = autoRenew;
            return this;
        }

        /**
         * @param autoRenew Specifies whether to enable auto-renewal for the instance. Default value: false. Valid values: true(enables auto-renewal), false(disables auto-renewal).
         * 
         * @return builder
         * 
         */
        public Builder autoRenew(String autoRenew) {
            return autoRenew(Output.of(autoRenew));
        }

        /**
         * @param autoRenewPeriod The subscription duration that is supported by auto-renewal. Unit: months. Valid values: 1, 2, 3, 6, and 12. This parameter is required only if the AutoRenew parameter is set to true.
         * 
         * @return builder
         * 
         */
        public Builder autoRenewPeriod(@Nullable Output<String> autoRenewPeriod) {
            $.autoRenewPeriod = autoRenewPeriod;
            return this;
        }

        /**
         * @param autoRenewPeriod The subscription duration that is supported by auto-renewal. Unit: months. Valid values: 1, 2, 3, 6, and 12. This parameter is required only if the AutoRenew parameter is set to true.
         * 
         * @return builder
         * 
         */
        public Builder autoRenewPeriod(String autoRenewPeriod) {
            return autoRenewPeriod(Output.of(autoRenewPeriod));
        }

        /**
         * @param createTime The time when the instance was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The time when the instance was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param effectiveTime The time when to change the configurations. Default value: Immediately. Valid values: Immediately (The configurations are immediately changed), MaintainTime (The configurations are changed within the maintenance window).
         * 
         * @return builder
         * 
         */
        public Builder effectiveTime(@Nullable Output<String> effectiveTime) {
            $.effectiveTime = effectiveTime;
            return this;
        }

        /**
         * @param effectiveTime The time when to change the configurations. Default value: Immediately. Valid values: Immediately (The configurations are immediately changed), MaintainTime (The configurations are changed within the maintenance window).
         * 
         * @return builder
         * 
         */
        public Builder effectiveTime(String effectiveTime) {
            return effectiveTime(Output.of(effectiveTime));
        }

        /**
         * @param engineVersion The database engine version of the instance. Default value: 1.0. The default version is developed by Alibaba Cloud and compatible with Redis 5.0.
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(@Nullable Output<String> engineVersion) {
            $.engineVersion = engineVersion;
            return this;
        }

        /**
         * @param engineVersion The database engine version of the instance. Default value: 1.0. The default version is developed by Alibaba Cloud and compatible with Redis 5.0.
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(String engineVersion) {
            return engineVersion(Output.of(engineVersion));
        }

        /**
         * @param forceUpgrade Specifies whether to forcefully change the configurations of the instance. Default value: true. Valid values: false (The system does not forcefully change the configurations), true (The system forcefully changes the configurations).
         * 
         * @return builder
         * 
         */
        public Builder forceUpgrade(@Nullable Output<Boolean> forceUpgrade) {
            $.forceUpgrade = forceUpgrade;
            return this;
        }

        /**
         * @param forceUpgrade Specifies whether to forcefully change the configurations of the instance. Default value: true. Valid values: false (The system does not forcefully change the configurations), true (The system forcefully changes the configurations).
         * 
         * @return builder
         * 
         */
        public Builder forceUpgrade(Boolean forceUpgrade) {
            return forceUpgrade(Output.of(forceUpgrade));
        }

        /**
         * @param instanceClass The instance type of the instance. For more information, see [Instance types](https://www.alibabacloud.com/help/en/apsaradb-for-redis/latest/instance-types).
         * 
         * @return builder
         * 
         */
        public Builder instanceClass(@Nullable Output<String> instanceClass) {
            $.instanceClass = instanceClass;
            return this;
        }

        /**
         * @param instanceClass The instance type of the instance. For more information, see [Instance types](https://www.alibabacloud.com/help/en/apsaradb-for-redis/latest/instance-types).
         * 
         * @return builder
         * 
         */
        public Builder instanceClass(String instanceClass) {
            return instanceClass(Output.of(instanceClass));
        }

        /**
         * @param instanceType The storage medium of the instance. Valid values: tair_rdb, tair_scm, tair_essd.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(@Nullable Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType The storage medium of the instance. Valid values: tair_rdb, tair_scm, tair_essd.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        /**
         * @param password The password that is used to connect to the instance. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include ! @ # $ % ^ &amp; * ( ) _ + - =.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password that is used to connect to the instance. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include ! @ # $ % ^ &amp; * ( ) _ + - =.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param paymentType Payment type: Subscription (prepaid), PayAsYouGo (postpaid). Default PayAsYouGo.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(@Nullable Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param paymentType Payment type: Subscription (prepaid), PayAsYouGo (postpaid). Default PayAsYouGo.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(String paymentType) {
            return paymentType(Output.of(paymentType));
        }

        /**
         * @param period The subscription duration. Unit: months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24,36, and 60. This parameter is required only if you set the PaymentType parameter to Subscription.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period The subscription duration. Unit: months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24,36, and 60. This parameter is required only if you set the PaymentType parameter to Subscription.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param port The Tair service port. The service port of the instance. Valid values: 1024 to 65535. Default value: 6379.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The Tair service port. The service port of the instance. Valid values: 1024 to 65535. Default value: 6379.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param resourceGroupId The ID of the resource group to which the instance belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group to which the instance belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param secondaryZoneId The ID of the secondary zone.This parameter is returned only if the instance is deployed in two zones.
         * 
         * @return builder
         * 
         */
        public Builder secondaryZoneId(@Nullable Output<String> secondaryZoneId) {
            $.secondaryZoneId = secondaryZoneId;
            return this;
        }

        /**
         * @param secondaryZoneId The ID of the secondary zone.This parameter is returned only if the instance is deployed in two zones.
         * 
         * @return builder
         * 
         */
        public Builder secondaryZoneId(String secondaryZoneId) {
            return secondaryZoneId(Output.of(secondaryZoneId));
        }

        /**
         * @param shardCount The number of data nodes in the instance. When 1 is passed, it means that the instance created is a standard architecture with only one data node. You can create an instance in the standard architecture that contains only a single data node. 2 to 32: You can create an instance in the cluster architecture that contains the specified number of data nodes. Only persistent memory-optimized instances can use the cluster architecture. Therefore, you can set this parameter to an integer from 2 to 32 only if you set the InstanceType parameter to tair_scm.
         * 
         * @return builder
         * 
         */
        public Builder shardCount(@Nullable Output<Integer> shardCount) {
            $.shardCount = shardCount;
            return this;
        }

        /**
         * @param shardCount The number of data nodes in the instance. When 1 is passed, it means that the instance created is a standard architecture with only one data node. You can create an instance in the standard architecture that contains only a single data node. 2 to 32: You can create an instance in the cluster architecture that contains the specified number of data nodes. Only persistent memory-optimized instances can use the cluster architecture. Therefore, you can set this parameter to an integer from 2 to 32 only if you set the InstanceType parameter to tair_scm.
         * 
         * @return builder
         * 
         */
        public Builder shardCount(Integer shardCount) {
            return shardCount(Output.of(shardCount));
        }

        /**
         * @param status The status of the instance. Valid values: Normal, Creating, Changing, Inactive, Flushing, Released, Transforming, Unavailable, Error, Migrating, and upgrading, networkModifying, SSLModifying, and MajorVersionUpgrading.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the instance. Valid values: Normal, Creating, Changing, Inactive, Flushing, Released, Transforming, Unavailable, Error, Migrating, and upgrading, networkModifying, SSLModifying, and MajorVersionUpgrading.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tairInstanceName The name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder tairInstanceName(@Nullable Output<String> tairInstanceName) {
            $.tairInstanceName = tairInstanceName;
            return this;
        }

        /**
         * @param tairInstanceName The name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder tairInstanceName(String tairInstanceName) {
            return tairInstanceName(Output.of(tairInstanceName));
        }

        /**
         * @param vpcId The ID of the virtual private cloud (VPC).
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the virtual private cloud (VPC).
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param vswitchId The ID of the vSwitch to which the instance is connected.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId The ID of the vSwitch to which the instance is connected.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        /**
         * @param zoneId The zone ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(@Nullable Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        /**
         * @param zoneId The zone ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public TairInstanceState build() {
            return $;
        }
    }

}
