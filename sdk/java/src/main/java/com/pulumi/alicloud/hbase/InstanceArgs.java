// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbase;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceArgs Empty = new InstanceArgs();

    /**
     * The account of the cluster web ui. Size [0-128].
     * 
     */
    @Import(name="account")
    private @Nullable Output<String> account;

    /**
     * @return The account of the cluster web ui. Size [0-128].
     * 
     */
    public Optional<Output<String>> account() {
        return Optional.ofNullable(this.account);
    }

    /**
     * Valid values are `true`, `false`, system default to `false`, valid when pay_type = PrePaid.
     * 
     */
    @Import(name="autoRenew")
    private @Nullable Output<Boolean> autoRenew;

    /**
     * @return Valid values are `true`, `false`, system default to `false`, valid when pay_type = PrePaid.
     * 
     */
    public Optional<Output<Boolean>> autoRenew() {
        return Optional.ofNullable(this.autoRenew);
    }

    /**
     * 0 or [800, 100000000], step:10-GB increments. 0 means is_cold_storage = false. [800, 100000000] means is_cold_storage = true.
     * 
     */
    @Import(name="coldStorageSize")
    private @Nullable Output<Integer> coldStorageSize;

    /**
     * @return 0 or [800, 100000000], step:10-GB increments. 0 means is_cold_storage = false. [800, 100000000] means is_cold_storage = true.
     * 
     */
    public Optional<Output<Integer>> coldStorageSize() {
        return Optional.ofNullable(this.coldStorageSize);
    }

    /**
     * User-defined HBase instance one core node&#39;s storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range:
     * - Custom storage space, value range: [20, 64000].
     * - Cluster [400, 64000], step:40-GB increments.
     * - Single [20-500GB], step:1-GB increments.
     * 
     */
    @Import(name="coreDiskSize")
    private @Nullable Output<Integer> coreDiskSize;

    /**
     * @return User-defined HBase instance one core node&#39;s storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range:
     * - Custom storage space, value range: [20, 64000].
     * - Cluster [400, 64000], step:40-GB increments.
     * - Single [20-500GB], step:1-GB increments.
     * 
     */
    public Optional<Output<Integer>> coreDiskSize() {
        return Optional.ofNullable(this.coreDiskSize);
    }

    /**
     * Valid values are `cloud_ssd`, `cloud_essd_pl1`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`，``, local_disk size is fixed. When engine=bds, no need to set disk type(or empty string).
     * 
     */
    @Import(name="coreDiskType")
    private @Nullable Output<String> coreDiskType;

    /**
     * @return Valid values are `cloud_ssd`, `cloud_essd_pl1`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`，``, local_disk size is fixed. When engine=bds, no need to set disk type(or empty string).
     * 
     */
    public Optional<Output<String>> coreDiskType() {
        return Optional.ofNullable(this.coreDiskType);
    }

    /**
     * Default=2, [1-200]. If core_instance_quantity &gt; 1, this is cluster&#39;s instance. If core_instance_quantity = 1, this is a single instance.
     * 
     */
    @Import(name="coreInstanceQuantity")
    private @Nullable Output<Integer> coreInstanceQuantity;

    /**
     * @return Default=2, [1-200]. If core_instance_quantity &gt; 1, this is cluster&#39;s instance. If core_instance_quantity = 1, this is a single instance.
     * 
     */
    public Optional<Output<Integer>> coreInstanceQuantity() {
        return Optional.ofNullable(this.coreInstanceQuantity);
    }

    /**
     * Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
     * 
     */
    @Import(name="coreInstanceType", required=true)
    private Output<String> coreInstanceType;

    /**
     * @return Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
     * 
     */
    public Output<String> coreInstanceType() {
        return this.coreInstanceType;
    }

    /**
     * The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.
     * 
     */
    @Import(name="deletionProtection")
    private @Nullable Output<Boolean> deletionProtection;

    /**
     * @return The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.
     * 
     */
    public Optional<Output<Boolean>> deletionProtection() {
        return Optional.ofNullable(this.deletionProtection);
    }

    /**
     * 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when pay_type = PrePaid,  unit: month. 12, 24, 36 mean 1, 2, 3 years.
     * 
     */
    @Import(name="duration")
    private @Nullable Output<Integer> duration;

    /**
     * @return 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when pay_type = PrePaid,  unit: month. 12, 24, 36 mean 1, 2, 3 years.
     * 
     */
    public Optional<Output<Integer>> duration() {
        return Optional.ofNullable(this.duration);
    }

    /**
     * Valid values are &#34;hbase/hbaseue/bds&#34;. The following types are supported after v1.73.0: `hbaseue` and `bds`. Single hbase instance need to set engine=hbase, core_instance_quantity=1.
     * 
     */
    @Import(name="engine")
    private @Nullable Output<String> engine;

    /**
     * @return Valid values are &#34;hbase/hbaseue/bds&#34;. The following types are supported after v1.73.0: `hbaseue` and `bds`. Single hbase instance need to set engine=hbase, core_instance_quantity=1.
     * 
     */
    public Optional<Output<String>> engine() {
        return Optional.ofNullable(this.engine);
    }

    /**
     * HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://help.aliyun.com/document_detail/144607.html).
     * 
     */
    @Import(name="engineVersion", required=true)
    private Output<String> engineVersion;

    /**
     * @return HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://help.aliyun.com/document_detail/144607.html).
     * 
     */
    public Output<String> engineVersion() {
        return this.engineVersion;
    }

    /**
     * The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.
     * 
     */
    @Import(name="immediateDeleteFlag")
    private @Nullable Output<Boolean> immediateDeleteFlag;

    /**
     * @return The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.
     * 
     */
    public Optional<Output<Boolean>> immediateDeleteFlag() {
        return Optional.ofNullable(this.immediateDeleteFlag);
    }

    /**
     * The white ip list of the cluster.
     * 
     */
    @Import(name="ipWhite")
    private @Nullable Output<String> ipWhite;

    /**
     * @return The white ip list of the cluster.
     * 
     */
    public Optional<Output<String>> ipWhite() {
        return Optional.ofNullable(this.ipWhite);
    }

    /**
     * The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.
     * 
     */
    @Import(name="maintainEndTime")
    private @Nullable Output<String> maintainEndTime;

    /**
     * @return The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.
     * 
     */
    public Optional<Output<String>> maintainEndTime() {
        return Optional.ofNullable(this.maintainEndTime);
    }

    /**
     * The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.
     * 
     */
    @Import(name="maintainStartTime")
    private @Nullable Output<String> maintainStartTime;

    /**
     * @return The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.
     * 
     */
    public Optional<Output<String>> maintainStartTime() {
        return Optional.ofNullable(this.maintainStartTime);
    }

    /**
     * Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
     * 
     */
    @Import(name="masterInstanceType", required=true)
    private Output<String> masterInstanceType;

    /**
     * @return Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
     * 
     */
    public Output<String> masterInstanceType() {
        return this.masterInstanceType;
    }

    /**
     * HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The password of the cluster web ui account. Size [0-128].
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The password of the cluster web ui account. Size [0-128].
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. You can also convert PostPaid to PrePaid. And support convert PrePaid to PostPaid from 1.115.0+.
     * 
     */
    @Import(name="payType")
    private @Nullable Output<String> payType;

    /**
     * @return Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. You can also convert PostPaid to PrePaid. And support convert PrePaid to PostPaid from 1.115.0+.
     * 
     */
    public Optional<Output<String>> payType() {
        return Optional.ofNullable(this.payType);
    }

    /**
     * The security group resource of the cluster.
     * 
     */
    @Import(name="securityGroups")
    private @Nullable Output<List<String>> securityGroups;

    /**
     * @return The security group resource of the cluster.
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
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The id of the VPC.
     * 
     * &gt; **NOTE:** Now only instance name can be change. The others(instance_type, disk_size, core_instance_quantity and so on) will be supported in the furture.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The id of the VPC.
     * 
     * &gt; **NOTE:** Now only instance name can be change. The others(instance_type, disk_size, core_instance_quantity and so on) will be supported in the furture.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * If vswitch_id is not empty, that mean net_type = vpc and has a same region. If vswitch_id is empty, net_type=classic. Intl site not support classic network.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable Output<String> vswitchId;

    /**
     * @return If vswitch_id is not empty, that mean net_type = vpc and has a same region. If vswitch_id is empty, net_type=classic. Intl site not support classic network.
     * 
     */
    public Optional<Output<String>> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    /**
     * The Zone to launch the HBase instance. If vswitch_id is not empty, this zone_id can be &#34;&#34; or consistent.
     * 
     */
    @Import(name="zoneId")
    private @Nullable Output<String> zoneId;

    /**
     * @return The Zone to launch the HBase instance. If vswitch_id is not empty, this zone_id can be &#34;&#34; or consistent.
     * 
     */
    public Optional<Output<String>> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    private InstanceArgs() {}

    private InstanceArgs(InstanceArgs $) {
        this.account = $.account;
        this.autoRenew = $.autoRenew;
        this.coldStorageSize = $.coldStorageSize;
        this.coreDiskSize = $.coreDiskSize;
        this.coreDiskType = $.coreDiskType;
        this.coreInstanceQuantity = $.coreInstanceQuantity;
        this.coreInstanceType = $.coreInstanceType;
        this.deletionProtection = $.deletionProtection;
        this.duration = $.duration;
        this.engine = $.engine;
        this.engineVersion = $.engineVersion;
        this.immediateDeleteFlag = $.immediateDeleteFlag;
        this.ipWhite = $.ipWhite;
        this.maintainEndTime = $.maintainEndTime;
        this.maintainStartTime = $.maintainStartTime;
        this.masterInstanceType = $.masterInstanceType;
        this.name = $.name;
        this.password = $.password;
        this.payType = $.payType;
        this.securityGroups = $.securityGroups;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
        this.vswitchId = $.vswitchId;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceArgs $;

        public Builder() {
            $ = new InstanceArgs();
        }

        public Builder(InstanceArgs defaults) {
            $ = new InstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param account The account of the cluster web ui. Size [0-128].
         * 
         * @return builder
         * 
         */
        public Builder account(@Nullable Output<String> account) {
            $.account = account;
            return this;
        }

        /**
         * @param account The account of the cluster web ui. Size [0-128].
         * 
         * @return builder
         * 
         */
        public Builder account(String account) {
            return account(Output.of(account));
        }

        /**
         * @param autoRenew Valid values are `true`, `false`, system default to `false`, valid when pay_type = PrePaid.
         * 
         * @return builder
         * 
         */
        public Builder autoRenew(@Nullable Output<Boolean> autoRenew) {
            $.autoRenew = autoRenew;
            return this;
        }

        /**
         * @param autoRenew Valid values are `true`, `false`, system default to `false`, valid when pay_type = PrePaid.
         * 
         * @return builder
         * 
         */
        public Builder autoRenew(Boolean autoRenew) {
            return autoRenew(Output.of(autoRenew));
        }

        /**
         * @param coldStorageSize 0 or [800, 100000000], step:10-GB increments. 0 means is_cold_storage = false. [800, 100000000] means is_cold_storage = true.
         * 
         * @return builder
         * 
         */
        public Builder coldStorageSize(@Nullable Output<Integer> coldStorageSize) {
            $.coldStorageSize = coldStorageSize;
            return this;
        }

        /**
         * @param coldStorageSize 0 or [800, 100000000], step:10-GB increments. 0 means is_cold_storage = false. [800, 100000000] means is_cold_storage = true.
         * 
         * @return builder
         * 
         */
        public Builder coldStorageSize(Integer coldStorageSize) {
            return coldStorageSize(Output.of(coldStorageSize));
        }

        /**
         * @param coreDiskSize User-defined HBase instance one core node&#39;s storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range:
         * - Custom storage space, value range: [20, 64000].
         * - Cluster [400, 64000], step:40-GB increments.
         * - Single [20-500GB], step:1-GB increments.
         * 
         * @return builder
         * 
         */
        public Builder coreDiskSize(@Nullable Output<Integer> coreDiskSize) {
            $.coreDiskSize = coreDiskSize;
            return this;
        }

        /**
         * @param coreDiskSize User-defined HBase instance one core node&#39;s storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range:
         * - Custom storage space, value range: [20, 64000].
         * - Cluster [400, 64000], step:40-GB increments.
         * - Single [20-500GB], step:1-GB increments.
         * 
         * @return builder
         * 
         */
        public Builder coreDiskSize(Integer coreDiskSize) {
            return coreDiskSize(Output.of(coreDiskSize));
        }

        /**
         * @param coreDiskType Valid values are `cloud_ssd`, `cloud_essd_pl1`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`，``, local_disk size is fixed. When engine=bds, no need to set disk type(or empty string).
         * 
         * @return builder
         * 
         */
        public Builder coreDiskType(@Nullable Output<String> coreDiskType) {
            $.coreDiskType = coreDiskType;
            return this;
        }

        /**
         * @param coreDiskType Valid values are `cloud_ssd`, `cloud_essd_pl1`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`，``, local_disk size is fixed. When engine=bds, no need to set disk type(or empty string).
         * 
         * @return builder
         * 
         */
        public Builder coreDiskType(String coreDiskType) {
            return coreDiskType(Output.of(coreDiskType));
        }

        /**
         * @param coreInstanceQuantity Default=2, [1-200]. If core_instance_quantity &gt; 1, this is cluster&#39;s instance. If core_instance_quantity = 1, this is a single instance.
         * 
         * @return builder
         * 
         */
        public Builder coreInstanceQuantity(@Nullable Output<Integer> coreInstanceQuantity) {
            $.coreInstanceQuantity = coreInstanceQuantity;
            return this;
        }

        /**
         * @param coreInstanceQuantity Default=2, [1-200]. If core_instance_quantity &gt; 1, this is cluster&#39;s instance. If core_instance_quantity = 1, this is a single instance.
         * 
         * @return builder
         * 
         */
        public Builder coreInstanceQuantity(Integer coreInstanceQuantity) {
            return coreInstanceQuantity(Output.of(coreInstanceQuantity));
        }

        /**
         * @param coreInstanceType Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
         * 
         * @return builder
         * 
         */
        public Builder coreInstanceType(Output<String> coreInstanceType) {
            $.coreInstanceType = coreInstanceType;
            return this;
        }

        /**
         * @param coreInstanceType Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
         * 
         * @return builder
         * 
         */
        public Builder coreInstanceType(String coreInstanceType) {
            return coreInstanceType(Output.of(coreInstanceType));
        }

        /**
         * @param deletionProtection The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(@Nullable Output<Boolean> deletionProtection) {
            $.deletionProtection = deletionProtection;
            return this;
        }

        /**
         * @param deletionProtection The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(Boolean deletionProtection) {
            return deletionProtection(Output.of(deletionProtection));
        }

        /**
         * @param duration 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when pay_type = PrePaid,  unit: month. 12, 24, 36 mean 1, 2, 3 years.
         * 
         * @return builder
         * 
         */
        public Builder duration(@Nullable Output<Integer> duration) {
            $.duration = duration;
            return this;
        }

        /**
         * @param duration 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when pay_type = PrePaid,  unit: month. 12, 24, 36 mean 1, 2, 3 years.
         * 
         * @return builder
         * 
         */
        public Builder duration(Integer duration) {
            return duration(Output.of(duration));
        }

        /**
         * @param engine Valid values are &#34;hbase/hbaseue/bds&#34;. The following types are supported after v1.73.0: `hbaseue` and `bds`. Single hbase instance need to set engine=hbase, core_instance_quantity=1.
         * 
         * @return builder
         * 
         */
        public Builder engine(@Nullable Output<String> engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param engine Valid values are &#34;hbase/hbaseue/bds&#34;. The following types are supported after v1.73.0: `hbaseue` and `bds`. Single hbase instance need to set engine=hbase, core_instance_quantity=1.
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            return engine(Output.of(engine));
        }

        /**
         * @param engineVersion HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://help.aliyun.com/document_detail/144607.html).
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(Output<String> engineVersion) {
            $.engineVersion = engineVersion;
            return this;
        }

        /**
         * @param engineVersion HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://help.aliyun.com/document_detail/144607.html).
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(String engineVersion) {
            return engineVersion(Output.of(engineVersion));
        }

        /**
         * @param immediateDeleteFlag The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.
         * 
         * @return builder
         * 
         */
        public Builder immediateDeleteFlag(@Nullable Output<Boolean> immediateDeleteFlag) {
            $.immediateDeleteFlag = immediateDeleteFlag;
            return this;
        }

        /**
         * @param immediateDeleteFlag The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.
         * 
         * @return builder
         * 
         */
        public Builder immediateDeleteFlag(Boolean immediateDeleteFlag) {
            return immediateDeleteFlag(Output.of(immediateDeleteFlag));
        }

        /**
         * @param ipWhite The white ip list of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder ipWhite(@Nullable Output<String> ipWhite) {
            $.ipWhite = ipWhite;
            return this;
        }

        /**
         * @param ipWhite The white ip list of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder ipWhite(String ipWhite) {
            return ipWhite(Output.of(ipWhite));
        }

        /**
         * @param maintainEndTime The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.
         * 
         * @return builder
         * 
         */
        public Builder maintainEndTime(@Nullable Output<String> maintainEndTime) {
            $.maintainEndTime = maintainEndTime;
            return this;
        }

        /**
         * @param maintainEndTime The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.
         * 
         * @return builder
         * 
         */
        public Builder maintainEndTime(String maintainEndTime) {
            return maintainEndTime(Output.of(maintainEndTime));
        }

        /**
         * @param maintainStartTime The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.
         * 
         * @return builder
         * 
         */
        public Builder maintainStartTime(@Nullable Output<String> maintainStartTime) {
            $.maintainStartTime = maintainStartTime;
            return this;
        }

        /**
         * @param maintainStartTime The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.
         * 
         * @return builder
         * 
         */
        public Builder maintainStartTime(String maintainStartTime) {
            return maintainStartTime(Output.of(maintainStartTime));
        }

        /**
         * @param masterInstanceType Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
         * 
         * @return builder
         * 
         */
        public Builder masterInstanceType(Output<String> masterInstanceType) {
            $.masterInstanceType = masterInstanceType;
            return this;
        }

        /**
         * @param masterInstanceType Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
         * 
         * @return builder
         * 
         */
        public Builder masterInstanceType(String masterInstanceType) {
            return masterInstanceType(Output.of(masterInstanceType));
        }

        /**
         * @param name HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password The password of the cluster web ui account. Size [0-128].
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password of the cluster web ui account. Size [0-128].
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param payType Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. You can also convert PostPaid to PrePaid. And support convert PrePaid to PostPaid from 1.115.0+.
         * 
         * @return builder
         * 
         */
        public Builder payType(@Nullable Output<String> payType) {
            $.payType = payType;
            return this;
        }

        /**
         * @param payType Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. You can also convert PostPaid to PrePaid. And support convert PrePaid to PostPaid from 1.115.0+.
         * 
         * @return builder
         * 
         */
        public Builder payType(String payType) {
            return payType(Output.of(payType));
        }

        /**
         * @param securityGroups The security group resource of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder securityGroups(@Nullable Output<List<String>> securityGroups) {
            $.securityGroups = securityGroups;
            return this;
        }

        /**
         * @param securityGroups The security group resource of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder securityGroups(List<String> securityGroups) {
            return securityGroups(Output.of(securityGroups));
        }

        /**
         * @param securityGroups The security group resource of the cluster.
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
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vpcId The id of the VPC.
         * 
         * &gt; **NOTE:** Now only instance name can be change. The others(instance_type, disk_size, core_instance_quantity and so on) will be supported in the furture.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The id of the VPC.
         * 
         * &gt; **NOTE:** Now only instance name can be change. The others(instance_type, disk_size, core_instance_quantity and so on) will be supported in the furture.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param vswitchId If vswitch_id is not empty, that mean net_type = vpc and has a same region. If vswitch_id is empty, net_type=classic. Intl site not support classic network.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId If vswitch_id is not empty, that mean net_type = vpc and has a same region. If vswitch_id is empty, net_type=classic. Intl site not support classic network.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        /**
         * @param zoneId The Zone to launch the HBase instance. If vswitch_id is not empty, this zone_id can be &#34;&#34; or consistent.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(@Nullable Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        /**
         * @param zoneId The Zone to launch the HBase instance. If vswitch_id is not empty, this zone_id can be &#34;&#34; or consistent.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public InstanceArgs build() {
            $.coreInstanceType = Objects.requireNonNull($.coreInstanceType, "expected parameter 'coreInstanceType' to be non-null");
            $.engineVersion = Objects.requireNonNull($.engineVersion, "expected parameter 'engineVersion' to be non-null");
            $.masterInstanceType = Objects.requireNonNull($.masterInstanceType, "expected parameter 'masterInstanceType' to be non-null");
            return $;
        }
    }

}
