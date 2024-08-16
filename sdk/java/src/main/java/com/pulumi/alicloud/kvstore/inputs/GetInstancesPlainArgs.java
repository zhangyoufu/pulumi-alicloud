// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.kvstore.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstancesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstancesPlainArgs Empty = new GetInstancesPlainArgs();

    /**
     * The type of the architecture. Valid values: `cluster`, `standard` and `SplitRW`.
     * 
     */
    @Import(name="architectureType")
    private @Nullable String architectureType;

    /**
     * @return The type of the architecture. Valid values: `cluster`, `standard` and `SplitRW`.
     * 
     */
    public Optional<String> architectureType() {
        return Optional.ofNullable(this.architectureType);
    }

    /**
     * Used to retrieve instances belong to specified `vswitch` resources.  Valid values: `Enterprise`, `Community`.
     * 
     */
    @Import(name="editionType")
    private @Nullable String editionType;

    /**
     * @return Used to retrieve instances belong to specified `vswitch` resources.  Valid values: `Enterprise`, `Community`.
     * 
     */
    public Optional<String> editionType() {
        return Optional.ofNullable(this.editionType);
    }

    /**
     * Default to `false`. Set it to true can output more details.
     * 
     */
    @Import(name="enableDetails")
    private @Nullable Boolean enableDetails;

    /**
     * @return Default to `false`. Set it to true can output more details.
     * 
     */
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * The engine version. Valid values: `2.8`, `4.0`, `5.0`, `6.0`, `7.0`.
     * 
     */
    @Import(name="engineVersion")
    private @Nullable String engineVersion;

    /**
     * @return The engine version. Valid values: `2.8`, `4.0`, `5.0`, `6.0`, `7.0`.
     * 
     */
    public Optional<String> engineVersion() {
        return Optional.ofNullable(this.engineVersion);
    }

    /**
     * The expiration status of the instance.
     * 
     */
    @Import(name="expired")
    private @Nullable String expired;

    /**
     * @return The expiration status of the instance.
     * 
     */
    public Optional<String> expired() {
        return Optional.ofNullable(this.expired);
    }

    /**
     * Whether to create a distributed cache.
     * 
     */
    @Import(name="globalInstance")
    private @Nullable Boolean globalInstance;

    /**
     * @return Whether to create a distributed cache.
     * 
     */
    public Optional<Boolean> globalInstance() {
        return Optional.ofNullable(this.globalInstance);
    }

    /**
     * A list of KVStore DBInstance IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of KVStore DBInstance IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * Type of the applied ApsaraDB for Redis instance. For more information, see [Instance type table](https://help.aliyun.com/zh/redis/developer-reference/instance-types).
     * 
     */
    @Import(name="instanceClass")
    private @Nullable String instanceClass;

    /**
     * @return Type of the applied ApsaraDB for Redis instance. For more information, see [Instance type table](https://help.aliyun.com/zh/redis/developer-reference/instance-types).
     * 
     */
    public Optional<String> instanceClass() {
        return Optional.ofNullable(this.instanceClass);
    }

    /**
     * The engine type of the KVStore DBInstance. Options are `Memcache`, and `Redis`. If no value is specified, all types are returned.
     * 
     */
    @Import(name="instanceType")
    private @Nullable String instanceType;

    /**
     * @return The engine type of the KVStore DBInstance. Options are `Memcache`, and `Redis`. If no value is specified, all types are returned.
     * 
     */
    public Optional<String> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }

    /**
     * A regex string to apply to the instance name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to apply to the instance name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * The type of the network. Valid values: `CLASSIC`, `VPC`.
     * 
     */
    @Import(name="networkType")
    private @Nullable String networkType;

    /**
     * @return The type of the network. Valid values: `CLASSIC`, `VPC`.
     * 
     */
    public Optional<String> networkType() {
        return Optional.ofNullable(this.networkType);
    }

    /**
     * The name of file that can save the collection of instances after running `pulumi preview`.
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return The name of file that can save the collection of instances after running `pulumi preview`.
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The payment type. Valid values: `PostPaid`, `PrePaid`.
     * 
     */
    @Import(name="paymentType")
    private @Nullable String paymentType;

    /**
     * @return The payment type. Valid values: `PostPaid`, `PrePaid`.
     * 
     */
    public Optional<String> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }

    /**
     * The ID of the resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable String resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The name of the instance.
     * 
     */
    @Import(name="searchKey")
    private @Nullable String searchKey;

    /**
     * @return The name of the instance.
     * 
     */
    public Optional<String> searchKey() {
        return Optional.ofNullable(this.searchKey);
    }

    /**
     * The status of the KVStore DBInstance. Valid values: `Changing`, `CleaningUpExpiredData`, `Creating`, `Flushing`, `HASwitching`, `Inactive`, `MajorVersionUpgrading`, `Migrating`, `NetworkModifying`, `Normal`, `Rebooting`, `SSLModifying`, `Transforming`, `ZoneMigrating`.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the KVStore DBInstance. Valid values: `Changing`, `CleaningUpExpiredData`, `Creating`, `Flushing`, `HASwitching`, `Inactive`, `MajorVersionUpgrading`, `Migrating`, `NetworkModifying`, `Normal`, `Rebooting`, `SSLModifying`, `Transforming`, `ZoneMigrating`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{&#34;key1&#34;:&#34;value1&#34;}`.
     * 
     */
    @Import(name="tags")
    private @Nullable Map<String,String> tags;

    /**
     * @return Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{&#34;key1&#34;:&#34;value1&#34;}`.
     * 
     */
    public Optional<Map<String,String>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Used to retrieve instances belong to specified VPC.
     * 
     */
    @Import(name="vpcId")
    private @Nullable String vpcId;

    /**
     * @return Used to retrieve instances belong to specified VPC.
     * 
     */
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * Used to retrieve instances belong to specified `vswitch` resources.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable String vswitchId;

    /**
     * @return Used to retrieve instances belong to specified `vswitch` resources.
     * 
     */
    public Optional<String> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    /**
     * The ID of the zone.
     * 
     */
    @Import(name="zoneId")
    private @Nullable String zoneId;

    /**
     * @return The ID of the zone.
     * 
     */
    public Optional<String> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    private GetInstancesPlainArgs() {}

    private GetInstancesPlainArgs(GetInstancesPlainArgs $) {
        this.architectureType = $.architectureType;
        this.editionType = $.editionType;
        this.enableDetails = $.enableDetails;
        this.engineVersion = $.engineVersion;
        this.expired = $.expired;
        this.globalInstance = $.globalInstance;
        this.ids = $.ids;
        this.instanceClass = $.instanceClass;
        this.instanceType = $.instanceType;
        this.nameRegex = $.nameRegex;
        this.networkType = $.networkType;
        this.outputFile = $.outputFile;
        this.paymentType = $.paymentType;
        this.resourceGroupId = $.resourceGroupId;
        this.searchKey = $.searchKey;
        this.status = $.status;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
        this.vswitchId = $.vswitchId;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstancesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstancesPlainArgs $;

        public Builder() {
            $ = new GetInstancesPlainArgs();
        }

        public Builder(GetInstancesPlainArgs defaults) {
            $ = new GetInstancesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param architectureType The type of the architecture. Valid values: `cluster`, `standard` and `SplitRW`.
         * 
         * @return builder
         * 
         */
        public Builder architectureType(@Nullable String architectureType) {
            $.architectureType = architectureType;
            return this;
        }

        /**
         * @param editionType Used to retrieve instances belong to specified `vswitch` resources.  Valid values: `Enterprise`, `Community`.
         * 
         * @return builder
         * 
         */
        public Builder editionType(@Nullable String editionType) {
            $.editionType = editionType;
            return this;
        }

        /**
         * @param enableDetails Default to `false`. Set it to true can output more details.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param engineVersion The engine version. Valid values: `2.8`, `4.0`, `5.0`, `6.0`, `7.0`.
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(@Nullable String engineVersion) {
            $.engineVersion = engineVersion;
            return this;
        }

        /**
         * @param expired The expiration status of the instance.
         * 
         * @return builder
         * 
         */
        public Builder expired(@Nullable String expired) {
            $.expired = expired;
            return this;
        }

        /**
         * @param globalInstance Whether to create a distributed cache.
         * 
         * @return builder
         * 
         */
        public Builder globalInstance(@Nullable Boolean globalInstance) {
            $.globalInstance = globalInstance;
            return this;
        }

        /**
         * @param ids A list of KVStore DBInstance IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of KVStore DBInstance IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceClass Type of the applied ApsaraDB for Redis instance. For more information, see [Instance type table](https://help.aliyun.com/zh/redis/developer-reference/instance-types).
         * 
         * @return builder
         * 
         */
        public Builder instanceClass(@Nullable String instanceClass) {
            $.instanceClass = instanceClass;
            return this;
        }

        /**
         * @param instanceType The engine type of the KVStore DBInstance. Options are `Memcache`, and `Redis`. If no value is specified, all types are returned.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(@Nullable String instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param nameRegex A regex string to apply to the instance name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param networkType The type of the network. Valid values: `CLASSIC`, `VPC`.
         * 
         * @return builder
         * 
         */
        public Builder networkType(@Nullable String networkType) {
            $.networkType = networkType;
            return this;
        }

        /**
         * @param outputFile The name of file that can save the collection of instances after running `pulumi preview`.
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param paymentType The payment type. Valid values: `PostPaid`, `PrePaid`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(@Nullable String paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param searchKey The name of the instance.
         * 
         * @return builder
         * 
         */
        public Builder searchKey(@Nullable String searchKey) {
            $.searchKey = searchKey;
            return this;
        }

        /**
         * @param status The status of the KVStore DBInstance. Valid values: `Changing`, `CleaningUpExpiredData`, `Creating`, `Flushing`, `HASwitching`, `Inactive`, `MajorVersionUpgrading`, `Migrating`, `NetworkModifying`, `Normal`, `Rebooting`, `SSLModifying`, `Transforming`, `ZoneMigrating`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param tags Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{&#34;key1&#34;:&#34;value1&#34;}`.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Map<String,String> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param vpcId Used to retrieve instances belong to specified VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable String vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vswitchId Used to retrieve instances belong to specified `vswitch` resources.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable String vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param zoneId The ID of the zone.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(@Nullable String zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        public GetInstancesPlainArgs build() {
            return $;
        }
    }

}
