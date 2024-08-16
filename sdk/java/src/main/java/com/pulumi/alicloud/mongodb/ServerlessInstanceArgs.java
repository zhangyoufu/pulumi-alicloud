// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mongodb;

import com.pulumi.alicloud.mongodb.inputs.ServerlessInstanceSecurityIpGroupArgs;
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


public final class ServerlessInstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServerlessInstanceArgs Empty = new ServerlessInstanceArgs();

    /**
     * The password of the database logon account.
     * * The password length is `8` to `32` bits.
     * * The password consists of at least any three of uppercase letters, lowercase letters, numbers, and special characters. The special character is `!#$%^&amp;*()_+-=`. The MongoDB Serverless instance provides a default database login account. This account cannot be modified. You can only set or modify the password for this account.
     * 
     */
    @Import(name="accountPassword", required=true)
    private Output<String> accountPassword;

    /**
     * @return The password of the database logon account.
     * * The password length is `8` to `32` bits.
     * * The password consists of at least any three of uppercase letters, lowercase letters, numbers, and special characters. The special character is `!#$%^&amp;*()_+-=`. The MongoDB Serverless instance provides a default database login account. This account cannot be modified. You can only set or modify the password for this account.
     * 
     */
    public Output<String> accountPassword() {
        return this.accountPassword;
    }

    /**
     * Set whether the instance is automatically renewed.
     * 
     */
    @Import(name="autoRenew")
    private @Nullable Output<Boolean> autoRenew;

    /**
     * @return Set whether the instance is automatically renewed.
     * 
     */
    public Optional<Output<Boolean>> autoRenew() {
        return Optional.ofNullable(this.autoRenew);
    }

    /**
     * The I/O throughput consumed by the instance. Valid values: `100` to `8000`.
     * 
     */
    @Import(name="capacityUnit", required=true)
    private Output<Integer> capacityUnit;

    /**
     * @return The I/O throughput consumed by the instance. Valid values: `100` to `8000`.
     * 
     */
    public Output<Integer> capacityUnit() {
        return this.capacityUnit;
    }

    /**
     * The db instance description.
     * 
     */
    @Import(name="dbInstanceDescription")
    private @Nullable Output<String> dbInstanceDescription;

    /**
     * @return The db instance description.
     * 
     */
    public Optional<Output<String>> dbInstanceDescription() {
        return Optional.ofNullable(this.dbInstanceDescription);
    }

    /**
     * The db instance storage. Valid values: `1` to `100`.
     * 
     */
    @Import(name="dbInstanceStorage", required=true)
    private Output<Integer> dbInstanceStorage;

    /**
     * @return The db instance storage. Valid values: `1` to `100`.
     * 
     */
    public Output<Integer> dbInstanceStorage() {
        return this.dbInstanceStorage;
    }

    /**
     * The database engine of the instance. Valid values: `MongoDB`.
     * 
     */
    @Import(name="engine")
    private @Nullable Output<String> engine;

    /**
     * @return The database engine of the instance. Valid values: `MongoDB`.
     * 
     */
    public Optional<Output<String>> engine() {
        return Optional.ofNullable(this.engine);
    }

    /**
     * The database version number. Valid values: `4.2`.
     * 
     */
    @Import(name="engineVersion", required=true)
    private Output<String> engineVersion;

    /**
     * @return The database version number. Valid values: `4.2`.
     * 
     */
    public Output<String> engineVersion() {
        return this.engineVersion;
    }

    /**
     * The end time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC. **NOTE:** The difference between the start time and end time must be one hour. For example, if `maintain_start_time` is `01:00Z`, `maintain_end_time` must be `02:00Z`.
     * 
     */
    @Import(name="maintainEndTime")
    private @Nullable Output<String> maintainEndTime;

    /**
     * @return The end time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC. **NOTE:** The difference between the start time and end time must be one hour. For example, if `maintain_start_time` is `01:00Z`, `maintain_end_time` must be `02:00Z`.
     * 
     */
    public Optional<Output<String>> maintainEndTime() {
        return Optional.ofNullable(this.maintainEndTime);
    }

    /**
     * The start time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC.
     * 
     */
    @Import(name="maintainStartTime")
    private @Nullable Output<String> maintainStartTime;

    /**
     * @return The start time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC.
     * 
     */
    public Optional<Output<String>> maintainStartTime() {
        return Optional.ofNullable(this.maintainStartTime);
    }

    /**
     * The purchase duration of the instance, in months. Valid values: `1` to `9`, `12`, `24`, `36`, `60`.
     * 
     */
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return The purchase duration of the instance, in months. Valid values: `1` to `9`, `12`, `24`, `36`, `60`.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * The period price type. Valid values: `Day`, `Month`.
     * 
     */
    @Import(name="periodPriceType")
    private @Nullable Output<String> periodPriceType;

    /**
     * @return The period price type. Valid values: `Day`, `Month`.
     * 
     */
    public Optional<Output<String>> periodPriceType() {
        return Optional.ofNullable(this.periodPriceType);
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
     * An array that consists of the information of IP whitelists.
     * 
     */
    @Import(name="securityIpGroups")
    private @Nullable Output<List<ServerlessInstanceSecurityIpGroupArgs>> securityIpGroups;

    /**
     * @return An array that consists of the information of IP whitelists.
     * 
     */
    public Optional<Output<List<ServerlessInstanceSecurityIpGroupArgs>>> securityIpGroups() {
        return Optional.ofNullable(this.securityIpGroups);
    }

    /**
     * The storage engine used by the instance. Valid values: `WiredTiger`.
     * 
     */
    @Import(name="storageEngine")
    private @Nullable Output<String> storageEngine;

    /**
     * @return The storage engine used by the instance. Valid values: `WiredTiger`.
     * 
     */
    public Optional<Output<String>> storageEngine() {
        return Optional.ofNullable(this.storageEngine);
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
     * The ID of the VPC network.
     * 
     */
    @Import(name="vpcId", required=true)
    private Output<String> vpcId;

    /**
     * @return The ID of the VPC network.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     * The of the vswitch.
     * 
     */
    @Import(name="vswitchId", required=true)
    private Output<String> vswitchId;

    /**
     * @return The of the vswitch.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    /**
     * The ID of the zone. Use this parameter to specify the zone created by the instance.
     * 
     */
    @Import(name="zoneId", required=true)
    private Output<String> zoneId;

    /**
     * @return The ID of the zone. Use this parameter to specify the zone created by the instance.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    private ServerlessInstanceArgs() {}

    private ServerlessInstanceArgs(ServerlessInstanceArgs $) {
        this.accountPassword = $.accountPassword;
        this.autoRenew = $.autoRenew;
        this.capacityUnit = $.capacityUnit;
        this.dbInstanceDescription = $.dbInstanceDescription;
        this.dbInstanceStorage = $.dbInstanceStorage;
        this.engine = $.engine;
        this.engineVersion = $.engineVersion;
        this.maintainEndTime = $.maintainEndTime;
        this.maintainStartTime = $.maintainStartTime;
        this.period = $.period;
        this.periodPriceType = $.periodPriceType;
        this.resourceGroupId = $.resourceGroupId;
        this.securityIpGroups = $.securityIpGroups;
        this.storageEngine = $.storageEngine;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
        this.vswitchId = $.vswitchId;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerlessInstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerlessInstanceArgs $;

        public Builder() {
            $ = new ServerlessInstanceArgs();
        }

        public Builder(ServerlessInstanceArgs defaults) {
            $ = new ServerlessInstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountPassword The password of the database logon account.
         * * The password length is `8` to `32` bits.
         * * The password consists of at least any three of uppercase letters, lowercase letters, numbers, and special characters. The special character is `!#$%^&amp;*()_+-=`. The MongoDB Serverless instance provides a default database login account. This account cannot be modified. You can only set or modify the password for this account.
         * 
         * @return builder
         * 
         */
        public Builder accountPassword(Output<String> accountPassword) {
            $.accountPassword = accountPassword;
            return this;
        }

        /**
         * @param accountPassword The password of the database logon account.
         * * The password length is `8` to `32` bits.
         * * The password consists of at least any three of uppercase letters, lowercase letters, numbers, and special characters. The special character is `!#$%^&amp;*()_+-=`. The MongoDB Serverless instance provides a default database login account. This account cannot be modified. You can only set or modify the password for this account.
         * 
         * @return builder
         * 
         */
        public Builder accountPassword(String accountPassword) {
            return accountPassword(Output.of(accountPassword));
        }

        /**
         * @param autoRenew Set whether the instance is automatically renewed.
         * 
         * @return builder
         * 
         */
        public Builder autoRenew(@Nullable Output<Boolean> autoRenew) {
            $.autoRenew = autoRenew;
            return this;
        }

        /**
         * @param autoRenew Set whether the instance is automatically renewed.
         * 
         * @return builder
         * 
         */
        public Builder autoRenew(Boolean autoRenew) {
            return autoRenew(Output.of(autoRenew));
        }

        /**
         * @param capacityUnit The I/O throughput consumed by the instance. Valid values: `100` to `8000`.
         * 
         * @return builder
         * 
         */
        public Builder capacityUnit(Output<Integer> capacityUnit) {
            $.capacityUnit = capacityUnit;
            return this;
        }

        /**
         * @param capacityUnit The I/O throughput consumed by the instance. Valid values: `100` to `8000`.
         * 
         * @return builder
         * 
         */
        public Builder capacityUnit(Integer capacityUnit) {
            return capacityUnit(Output.of(capacityUnit));
        }

        /**
         * @param dbInstanceDescription The db instance description.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceDescription(@Nullable Output<String> dbInstanceDescription) {
            $.dbInstanceDescription = dbInstanceDescription;
            return this;
        }

        /**
         * @param dbInstanceDescription The db instance description.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceDescription(String dbInstanceDescription) {
            return dbInstanceDescription(Output.of(dbInstanceDescription));
        }

        /**
         * @param dbInstanceStorage The db instance storage. Valid values: `1` to `100`.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceStorage(Output<Integer> dbInstanceStorage) {
            $.dbInstanceStorage = dbInstanceStorage;
            return this;
        }

        /**
         * @param dbInstanceStorage The db instance storage. Valid values: `1` to `100`.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceStorage(Integer dbInstanceStorage) {
            return dbInstanceStorage(Output.of(dbInstanceStorage));
        }

        /**
         * @param engine The database engine of the instance. Valid values: `MongoDB`.
         * 
         * @return builder
         * 
         */
        public Builder engine(@Nullable Output<String> engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param engine The database engine of the instance. Valid values: `MongoDB`.
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            return engine(Output.of(engine));
        }

        /**
         * @param engineVersion The database version number. Valid values: `4.2`.
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(Output<String> engineVersion) {
            $.engineVersion = engineVersion;
            return this;
        }

        /**
         * @param engineVersion The database version number. Valid values: `4.2`.
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(String engineVersion) {
            return engineVersion(Output.of(engineVersion));
        }

        /**
         * @param maintainEndTime The end time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC. **NOTE:** The difference between the start time and end time must be one hour. For example, if `maintain_start_time` is `01:00Z`, `maintain_end_time` must be `02:00Z`.
         * 
         * @return builder
         * 
         */
        public Builder maintainEndTime(@Nullable Output<String> maintainEndTime) {
            $.maintainEndTime = maintainEndTime;
            return this;
        }

        /**
         * @param maintainEndTime The end time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC. **NOTE:** The difference between the start time and end time must be one hour. For example, if `maintain_start_time` is `01:00Z`, `maintain_end_time` must be `02:00Z`.
         * 
         * @return builder
         * 
         */
        public Builder maintainEndTime(String maintainEndTime) {
            return maintainEndTime(Output.of(maintainEndTime));
        }

        /**
         * @param maintainStartTime The start time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC.
         * 
         * @return builder
         * 
         */
        public Builder maintainStartTime(@Nullable Output<String> maintainStartTime) {
            $.maintainStartTime = maintainStartTime;
            return this;
        }

        /**
         * @param maintainStartTime The start time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC.
         * 
         * @return builder
         * 
         */
        public Builder maintainStartTime(String maintainStartTime) {
            return maintainStartTime(Output.of(maintainStartTime));
        }

        /**
         * @param period The purchase duration of the instance, in months. Valid values: `1` to `9`, `12`, `24`, `36`, `60`.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period The purchase duration of the instance, in months. Valid values: `1` to `9`, `12`, `24`, `36`, `60`.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param periodPriceType The period price type. Valid values: `Day`, `Month`.
         * 
         * @return builder
         * 
         */
        public Builder periodPriceType(@Nullable Output<String> periodPriceType) {
            $.periodPriceType = periodPriceType;
            return this;
        }

        /**
         * @param periodPriceType The period price type. Valid values: `Day`, `Month`.
         * 
         * @return builder
         * 
         */
        public Builder periodPriceType(String periodPriceType) {
            return periodPriceType(Output.of(periodPriceType));
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
         * @param securityIpGroups An array that consists of the information of IP whitelists.
         * 
         * @return builder
         * 
         */
        public Builder securityIpGroups(@Nullable Output<List<ServerlessInstanceSecurityIpGroupArgs>> securityIpGroups) {
            $.securityIpGroups = securityIpGroups;
            return this;
        }

        /**
         * @param securityIpGroups An array that consists of the information of IP whitelists.
         * 
         * @return builder
         * 
         */
        public Builder securityIpGroups(List<ServerlessInstanceSecurityIpGroupArgs> securityIpGroups) {
            return securityIpGroups(Output.of(securityIpGroups));
        }

        /**
         * @param securityIpGroups An array that consists of the information of IP whitelists.
         * 
         * @return builder
         * 
         */
        public Builder securityIpGroups(ServerlessInstanceSecurityIpGroupArgs... securityIpGroups) {
            return securityIpGroups(List.of(securityIpGroups));
        }

        /**
         * @param storageEngine The storage engine used by the instance. Valid values: `WiredTiger`.
         * 
         * @return builder
         * 
         */
        public Builder storageEngine(@Nullable Output<String> storageEngine) {
            $.storageEngine = storageEngine;
            return this;
        }

        /**
         * @param storageEngine The storage engine used by the instance. Valid values: `WiredTiger`.
         * 
         * @return builder
         * 
         */
        public Builder storageEngine(String storageEngine) {
            return storageEngine(Output.of(storageEngine));
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
         * @param vpcId The ID of the VPC network.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the VPC network.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param vswitchId The of the vswitch.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId The of the vswitch.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        /**
         * @param zoneId The ID of the zone. Use this parameter to specify the zone created by the instance.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        /**
         * @param zoneId The ID of the zone. Use this parameter to specify the zone created by the instance.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public ServerlessInstanceArgs build() {
            if ($.accountPassword == null) {
                throw new MissingRequiredPropertyException("ServerlessInstanceArgs", "accountPassword");
            }
            if ($.capacityUnit == null) {
                throw new MissingRequiredPropertyException("ServerlessInstanceArgs", "capacityUnit");
            }
            if ($.dbInstanceStorage == null) {
                throw new MissingRequiredPropertyException("ServerlessInstanceArgs", "dbInstanceStorage");
            }
            if ($.engineVersion == null) {
                throw new MissingRequiredPropertyException("ServerlessInstanceArgs", "engineVersion");
            }
            if ($.vpcId == null) {
                throw new MissingRequiredPropertyException("ServerlessInstanceArgs", "vpcId");
            }
            if ($.vswitchId == null) {
                throw new MissingRequiredPropertyException("ServerlessInstanceArgs", "vswitchId");
            }
            if ($.zoneId == null) {
                throw new MissingRequiredPropertyException("ServerlessInstanceArgs", "zoneId");
            }
            return $;
        }
    }

}
