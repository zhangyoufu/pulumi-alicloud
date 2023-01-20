// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ebs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DedicatedBlockStorageClusterState extends com.pulumi.resources.ResourceArgs {

    public static final DedicatedBlockStorageClusterState Empty = new DedicatedBlockStorageClusterState();

    /**
     * The available capacity of the dedicated block storage cluster. Unit: GiB.
     * 
     */
    @Import(name="availableCapacity")
    private @Nullable Output<String> availableCapacity;

    /**
     * @return The available capacity of the dedicated block storage cluster. Unit: GiB.
     * 
     */
    public Optional<Output<String>> availableCapacity() {
        return Optional.ofNullable(this.availableCapacity);
    }

    /**
     * The type of cloud disk that can be created by a dedicated block storage cluster.
     * 
     */
    @Import(name="category")
    private @Nullable Output<String> category;

    /**
     * @return The type of cloud disk that can be created by a dedicated block storage cluster.
     * 
     */
    public Optional<Output<String>> category() {
        return Optional.ofNullable(this.category);
    }

    /**
     * The creation time of the resource
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The creation time of the resource
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * The first ID of the resource
     * 
     */
    @Import(name="dedicatedBlockStorageClusterId")
    private @Nullable Output<String> dedicatedBlockStorageClusterId;

    /**
     * @return The first ID of the resource
     * 
     */
    public Optional<Output<String>> dedicatedBlockStorageClusterId() {
        return Optional.ofNullable(this.dedicatedBlockStorageClusterId);
    }

    /**
     * The name of the resource
     * 
     */
    @Import(name="dedicatedBlockStorageClusterName")
    private @Nullable Output<String> dedicatedBlockStorageClusterName;

    /**
     * @return The name of the resource
     * 
     */
    public Optional<Output<String>> dedicatedBlockStorageClusterName() {
        return Optional.ofNullable(this.dedicatedBlockStorageClusterName);
    }

    /**
     * Capacity to be delivered in GB.
     * 
     */
    @Import(name="deliveryCapacity")
    private @Nullable Output<String> deliveryCapacity;

    /**
     * @return Capacity to be delivered in GB.
     * 
     */
    public Optional<Output<String>> deliveryCapacity() {
        return Optional.ofNullable(this.deliveryCapacity);
    }

    /**
     * The description of the dedicated block storage cluster.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the dedicated block storage cluster.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The expiration time of the dedicated block storage cluster, in the Unix timestamp format, in seconds.
     * 
     */
    @Import(name="expiredTime")
    private @Nullable Output<String> expiredTime;

    /**
     * @return The expiration time of the dedicated block storage cluster, in the Unix timestamp format, in seconds.
     * 
     */
    public Optional<Output<String>> expiredTime() {
        return Optional.ofNullable(this.expiredTime);
    }

    /**
     * Cloud disk performance level, possible values:-PL0.-PL1.-PL2.-PL3.&gt; Only valid in SupportedCategory = cloud_essd.
     * 
     */
    @Import(name="performanceLevel")
    private @Nullable Output<String> performanceLevel;

    /**
     * @return Cloud disk performance level, possible values:-PL0.-PL1.-PL2.-PL3.&gt; Only valid in SupportedCategory = cloud_essd.
     * 
     */
    public Optional<Output<String>> performanceLevel() {
        return Optional.ofNullable(this.performanceLevel);
    }

    /**
     * The ID of the resource group
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The status of the resource
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the resource
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * This parameter is not supported.
     * 
     */
    @Import(name="supportedCategory")
    private @Nullable Output<String> supportedCategory;

    /**
     * @return This parameter is not supported.
     * 
     */
    public Optional<Output<String>> supportedCategory() {
        return Optional.ofNullable(this.supportedCategory);
    }

    /**
     * The total capacity of the dedicated block storage cluster. Unit: GiB.
     * 
     */
    @Import(name="totalCapacity")
    private @Nullable Output<String> totalCapacity;

    /**
     * @return The total capacity of the dedicated block storage cluster. Unit: GiB.
     * 
     */
    public Optional<Output<String>> totalCapacity() {
        return Optional.ofNullable(this.totalCapacity);
    }

    /**
     * The dedicated block storage cluster performance type. Possible values:-Standard: Basic type. This type of dedicated block storage cluster can create an ESSD PL0 cloud disk.-Premium: performance type. This type of dedicated block storage cluster can create an ESSD PL1 cloud disk.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The dedicated block storage cluster performance type. Possible values:-Standard: Basic type. This type of dedicated block storage cluster can create an ESSD PL0 cloud disk.-Premium: performance type. This type of dedicated block storage cluster can create an ESSD PL1 cloud disk.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * The used (created disk) capacity of the current cluster, in GB
     * 
     */
    @Import(name="usedCapacity")
    private @Nullable Output<String> usedCapacity;

    /**
     * @return The used (created disk) capacity of the current cluster, in GB
     * 
     */
    public Optional<Output<String>> usedCapacity() {
        return Optional.ofNullable(this.usedCapacity);
    }

    /**
     * The zone ID  of the resource
     * 
     */
    @Import(name="zoneId")
    private @Nullable Output<String> zoneId;

    /**
     * @return The zone ID  of the resource
     * 
     */
    public Optional<Output<String>> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    private DedicatedBlockStorageClusterState() {}

    private DedicatedBlockStorageClusterState(DedicatedBlockStorageClusterState $) {
        this.availableCapacity = $.availableCapacity;
        this.category = $.category;
        this.createTime = $.createTime;
        this.dedicatedBlockStorageClusterId = $.dedicatedBlockStorageClusterId;
        this.dedicatedBlockStorageClusterName = $.dedicatedBlockStorageClusterName;
        this.deliveryCapacity = $.deliveryCapacity;
        this.description = $.description;
        this.expiredTime = $.expiredTime;
        this.performanceLevel = $.performanceLevel;
        this.resourceGroupId = $.resourceGroupId;
        this.status = $.status;
        this.supportedCategory = $.supportedCategory;
        this.totalCapacity = $.totalCapacity;
        this.type = $.type;
        this.usedCapacity = $.usedCapacity;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DedicatedBlockStorageClusterState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DedicatedBlockStorageClusterState $;

        public Builder() {
            $ = new DedicatedBlockStorageClusterState();
        }

        public Builder(DedicatedBlockStorageClusterState defaults) {
            $ = new DedicatedBlockStorageClusterState(Objects.requireNonNull(defaults));
        }

        /**
         * @param availableCapacity The available capacity of the dedicated block storage cluster. Unit: GiB.
         * 
         * @return builder
         * 
         */
        public Builder availableCapacity(@Nullable Output<String> availableCapacity) {
            $.availableCapacity = availableCapacity;
            return this;
        }

        /**
         * @param availableCapacity The available capacity of the dedicated block storage cluster. Unit: GiB.
         * 
         * @return builder
         * 
         */
        public Builder availableCapacity(String availableCapacity) {
            return availableCapacity(Output.of(availableCapacity));
        }

        /**
         * @param category The type of cloud disk that can be created by a dedicated block storage cluster.
         * 
         * @return builder
         * 
         */
        public Builder category(@Nullable Output<String> category) {
            $.category = category;
            return this;
        }

        /**
         * @param category The type of cloud disk that can be created by a dedicated block storage cluster.
         * 
         * @return builder
         * 
         */
        public Builder category(String category) {
            return category(Output.of(category));
        }

        /**
         * @param createTime The creation time of the resource
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The creation time of the resource
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param dedicatedBlockStorageClusterId The first ID of the resource
         * 
         * @return builder
         * 
         */
        public Builder dedicatedBlockStorageClusterId(@Nullable Output<String> dedicatedBlockStorageClusterId) {
            $.dedicatedBlockStorageClusterId = dedicatedBlockStorageClusterId;
            return this;
        }

        /**
         * @param dedicatedBlockStorageClusterId The first ID of the resource
         * 
         * @return builder
         * 
         */
        public Builder dedicatedBlockStorageClusterId(String dedicatedBlockStorageClusterId) {
            return dedicatedBlockStorageClusterId(Output.of(dedicatedBlockStorageClusterId));
        }

        /**
         * @param dedicatedBlockStorageClusterName The name of the resource
         * 
         * @return builder
         * 
         */
        public Builder dedicatedBlockStorageClusterName(@Nullable Output<String> dedicatedBlockStorageClusterName) {
            $.dedicatedBlockStorageClusterName = dedicatedBlockStorageClusterName;
            return this;
        }

        /**
         * @param dedicatedBlockStorageClusterName The name of the resource
         * 
         * @return builder
         * 
         */
        public Builder dedicatedBlockStorageClusterName(String dedicatedBlockStorageClusterName) {
            return dedicatedBlockStorageClusterName(Output.of(dedicatedBlockStorageClusterName));
        }

        /**
         * @param deliveryCapacity Capacity to be delivered in GB.
         * 
         * @return builder
         * 
         */
        public Builder deliveryCapacity(@Nullable Output<String> deliveryCapacity) {
            $.deliveryCapacity = deliveryCapacity;
            return this;
        }

        /**
         * @param deliveryCapacity Capacity to be delivered in GB.
         * 
         * @return builder
         * 
         */
        public Builder deliveryCapacity(String deliveryCapacity) {
            return deliveryCapacity(Output.of(deliveryCapacity));
        }

        /**
         * @param description The description of the dedicated block storage cluster.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the dedicated block storage cluster.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param expiredTime The expiration time of the dedicated block storage cluster, in the Unix timestamp format, in seconds.
         * 
         * @return builder
         * 
         */
        public Builder expiredTime(@Nullable Output<String> expiredTime) {
            $.expiredTime = expiredTime;
            return this;
        }

        /**
         * @param expiredTime The expiration time of the dedicated block storage cluster, in the Unix timestamp format, in seconds.
         * 
         * @return builder
         * 
         */
        public Builder expiredTime(String expiredTime) {
            return expiredTime(Output.of(expiredTime));
        }

        /**
         * @param performanceLevel Cloud disk performance level, possible values:-PL0.-PL1.-PL2.-PL3.&gt; Only valid in SupportedCategory = cloud_essd.
         * 
         * @return builder
         * 
         */
        public Builder performanceLevel(@Nullable Output<String> performanceLevel) {
            $.performanceLevel = performanceLevel;
            return this;
        }

        /**
         * @param performanceLevel Cloud disk performance level, possible values:-PL0.-PL1.-PL2.-PL3.&gt; Only valid in SupportedCategory = cloud_essd.
         * 
         * @return builder
         * 
         */
        public Builder performanceLevel(String performanceLevel) {
            return performanceLevel(Output.of(performanceLevel));
        }

        /**
         * @param resourceGroupId The ID of the resource group
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param status The status of the resource
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the resource
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param supportedCategory This parameter is not supported.
         * 
         * @return builder
         * 
         */
        public Builder supportedCategory(@Nullable Output<String> supportedCategory) {
            $.supportedCategory = supportedCategory;
            return this;
        }

        /**
         * @param supportedCategory This parameter is not supported.
         * 
         * @return builder
         * 
         */
        public Builder supportedCategory(String supportedCategory) {
            return supportedCategory(Output.of(supportedCategory));
        }

        /**
         * @param totalCapacity The total capacity of the dedicated block storage cluster. Unit: GiB.
         * 
         * @return builder
         * 
         */
        public Builder totalCapacity(@Nullable Output<String> totalCapacity) {
            $.totalCapacity = totalCapacity;
            return this;
        }

        /**
         * @param totalCapacity The total capacity of the dedicated block storage cluster. Unit: GiB.
         * 
         * @return builder
         * 
         */
        public Builder totalCapacity(String totalCapacity) {
            return totalCapacity(Output.of(totalCapacity));
        }

        /**
         * @param type The dedicated block storage cluster performance type. Possible values:-Standard: Basic type. This type of dedicated block storage cluster can create an ESSD PL0 cloud disk.-Premium: performance type. This type of dedicated block storage cluster can create an ESSD PL1 cloud disk.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The dedicated block storage cluster performance type. Possible values:-Standard: Basic type. This type of dedicated block storage cluster can create an ESSD PL0 cloud disk.-Premium: performance type. This type of dedicated block storage cluster can create an ESSD PL1 cloud disk.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param usedCapacity The used (created disk) capacity of the current cluster, in GB
         * 
         * @return builder
         * 
         */
        public Builder usedCapacity(@Nullable Output<String> usedCapacity) {
            $.usedCapacity = usedCapacity;
            return this;
        }

        /**
         * @param usedCapacity The used (created disk) capacity of the current cluster, in GB
         * 
         * @return builder
         * 
         */
        public Builder usedCapacity(String usedCapacity) {
            return usedCapacity(Output.of(usedCapacity));
        }

        /**
         * @param zoneId The zone ID  of the resource
         * 
         * @return builder
         * 
         */
        public Builder zoneId(@Nullable Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        /**
         * @param zoneId The zone ID  of the resource
         * 
         * @return builder
         * 
         */
        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public DedicatedBlockStorageClusterState build() {
            return $;
        }
    }

}
