// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DiskState extends com.pulumi.resources.ResourceArgs {

    public static final DiskState Empty = new DiskState();

    @Import(name="advancedFeatures")
    private @Nullable Output<String> advancedFeatures;

    public Optional<Output<String>> advancedFeatures() {
        return Optional.ofNullable(this.advancedFeatures);
    }

    /**
     * The Zone to create the disk in.
     * 
     * @deprecated
     * Field &#39;availability_zone&#39; has been deprecated from provider version 1.122.0. New field &#39;zone_id&#39; instead
     * 
     */
    @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead */
    @Import(name="availabilityZone")
    private @Nullable Output<String> availabilityZone;

    /**
     * @return The Zone to create the disk in.
     * 
     * @deprecated
     * Field &#39;availability_zone&#39; has been deprecated from provider version 1.122.0. New field &#39;zone_id&#39; instead
     * 
     */
    @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead */
    public Optional<Output<String>> availabilityZone() {
        return Optional.ofNullable(this.availabilityZone);
    }

    /**
     * Category of the disk. Valid values are `cloud`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd`, `cloud_essd_entry`. Default is `cloud_efficiency`.
     * 
     */
    @Import(name="category")
    private @Nullable Output<String> category;

    /**
     * @return Category of the disk. Valid values are `cloud`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd`, `cloud_essd_entry`. Default is `cloud_efficiency`.
     * 
     */
    public Optional<Output<String>> category() {
        return Optional.ofNullable(this.category);
    }

    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
     * 
     */
    @Import(name="deleteAutoSnapshot")
    private @Nullable Output<Boolean> deleteAutoSnapshot;

    /**
     * @return Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
     * 
     */
    public Optional<Output<Boolean>> deleteAutoSnapshot() {
        return Optional.ofNullable(this.deleteAutoSnapshot);
    }

    /**
     * Indicates whether the disk is released together with the instance: Default value: false.
     * 
     */
    @Import(name="deleteWithInstance")
    private @Nullable Output<Boolean> deleteWithInstance;

    /**
     * @return Indicates whether the disk is released together with the instance: Default value: false.
     * 
     */
    public Optional<Output<Boolean>> deleteWithInstance() {
        return Optional.ofNullable(this.deleteWithInstance);
    }

    /**
     * Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="diskName")
    private @Nullable Output<String> diskName;

    public Optional<Output<String>> diskName() {
        return Optional.ofNullable(this.diskName);
    }

    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
     * 
     */
    @Import(name="enableAutoSnapshot")
    private @Nullable Output<Boolean> enableAutoSnapshot;

    /**
     * @return Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
     * 
     */
    public Optional<Output<Boolean>> enableAutoSnapshot() {
        return Optional.ofNullable(this.enableAutoSnapshot);
    }

    @Import(name="encryptAlgorithm")
    private @Nullable Output<String> encryptAlgorithm;

    public Optional<Output<String>> encryptAlgorithm() {
        return Optional.ofNullable(this.encryptAlgorithm);
    }

    /**
     * If true, the disk will be encrypted, conflict with `snapshot_id`.
     * 
     */
    @Import(name="encrypted")
    private @Nullable Output<Boolean> encrypted;

    /**
     * @return If true, the disk will be encrypted, conflict with `snapshot_id`.
     * 
     */
    public Optional<Output<Boolean>> encrypted() {
        return Optional.ofNullable(this.encrypted);
    }

    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
     * 
     */
    @Import(name="kmsKeyId")
    private @Nullable Output<String> kmsKeyId;

    /**
     * @return The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
     * 
     */
    public Optional<Output<String>> kmsKeyId() {
        return Optional.ofNullable(this.kmsKeyId);
    }

    /**
     * Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.122.0. New field &#39;disk_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead. */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.122.0. New field &#39;disk_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead. */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="paymentType")
    private @Nullable Output<String> paymentType;

    public Optional<Output<String>> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }

    /**
     * Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:
     * * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
     * * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
     * * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
     * 
     */
    @Import(name="performanceLevel")
    private @Nullable Output<String> performanceLevel;

    /**
     * @return Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:
     * * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
     * * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
     * * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
     * 
     */
    public Optional<Output<String>> performanceLevel() {
        return Optional.ofNullable(this.performanceLevel);
    }

    /**
     * The Id of resource group which the disk belongs.
     * &gt; **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloud_efficiency` and `cloud_ssd` disk.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The Id of resource group which the disk belongs.
     * &gt; **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloud_efficiency` and `cloud_ssd` disk.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     * 
     */
    @Import(name="size")
    private @Nullable Output<Integer> size;

    /**
     * @return The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     * 
     */
    public Optional<Output<Integer>> size() {
        return Optional.ofNullable(this.size);
    }

    /**
     * A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     * 
     */
    @Import(name="snapshotId")
    private @Nullable Output<String> snapshotId;

    /**
     * @return A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     * 
     */
    public Optional<Output<String>> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
    }

    /**
     * The disk status.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The disk status.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    @Import(name="storageSetId")
    private @Nullable Output<String> storageSetId;

    public Optional<Output<String>> storageSetId() {
        return Optional.ofNullable(this.storageSetId);
    }

    @Import(name="storageSetPartitionNumber")
    private @Nullable Output<Integer> storageSetPartitionNumber;

    public Optional<Output<Integer>> storageSetPartitionNumber() {
        return Optional.ofNullable(this.storageSetPartitionNumber);
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

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    @Import(name="zoneId")
    private @Nullable Output<String> zoneId;

    public Optional<Output<String>> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    private DiskState() {}

    private DiskState(DiskState $) {
        this.advancedFeatures = $.advancedFeatures;
        this.availabilityZone = $.availabilityZone;
        this.category = $.category;
        this.deleteAutoSnapshot = $.deleteAutoSnapshot;
        this.deleteWithInstance = $.deleteWithInstance;
        this.description = $.description;
        this.diskName = $.diskName;
        this.dryRun = $.dryRun;
        this.enableAutoSnapshot = $.enableAutoSnapshot;
        this.encryptAlgorithm = $.encryptAlgorithm;
        this.encrypted = $.encrypted;
        this.instanceId = $.instanceId;
        this.kmsKeyId = $.kmsKeyId;
        this.name = $.name;
        this.paymentType = $.paymentType;
        this.performanceLevel = $.performanceLevel;
        this.resourceGroupId = $.resourceGroupId;
        this.size = $.size;
        this.snapshotId = $.snapshotId;
        this.status = $.status;
        this.storageSetId = $.storageSetId;
        this.storageSetPartitionNumber = $.storageSetPartitionNumber;
        this.tags = $.tags;
        this.type = $.type;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DiskState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DiskState $;

        public Builder() {
            $ = new DiskState();
        }

        public Builder(DiskState defaults) {
            $ = new DiskState(Objects.requireNonNull(defaults));
        }

        public Builder advancedFeatures(@Nullable Output<String> advancedFeatures) {
            $.advancedFeatures = advancedFeatures;
            return this;
        }

        public Builder advancedFeatures(String advancedFeatures) {
            return advancedFeatures(Output.of(advancedFeatures));
        }

        /**
         * @param availabilityZone The Zone to create the disk in.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;availability_zone&#39; has been deprecated from provider version 1.122.0. New field &#39;zone_id&#39; instead
         * 
         */
        @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead */
        public Builder availabilityZone(@Nullable Output<String> availabilityZone) {
            $.availabilityZone = availabilityZone;
            return this;
        }

        /**
         * @param availabilityZone The Zone to create the disk in.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;availability_zone&#39; has been deprecated from provider version 1.122.0. New field &#39;zone_id&#39; instead
         * 
         */
        @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead */
        public Builder availabilityZone(String availabilityZone) {
            return availabilityZone(Output.of(availabilityZone));
        }

        /**
         * @param category Category of the disk. Valid values are `cloud`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd`, `cloud_essd_entry`. Default is `cloud_efficiency`.
         * 
         * @return builder
         * 
         */
        public Builder category(@Nullable Output<String> category) {
            $.category = category;
            return this;
        }

        /**
         * @param category Category of the disk. Valid values are `cloud`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd`, `cloud_essd_entry`. Default is `cloud_efficiency`.
         * 
         * @return builder
         * 
         */
        public Builder category(String category) {
            return category(Output.of(category));
        }

        /**
         * @param deleteAutoSnapshot Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
         * 
         * @return builder
         * 
         */
        public Builder deleteAutoSnapshot(@Nullable Output<Boolean> deleteAutoSnapshot) {
            $.deleteAutoSnapshot = deleteAutoSnapshot;
            return this;
        }

        /**
         * @param deleteAutoSnapshot Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
         * 
         * @return builder
         * 
         */
        public Builder deleteAutoSnapshot(Boolean deleteAutoSnapshot) {
            return deleteAutoSnapshot(Output.of(deleteAutoSnapshot));
        }

        /**
         * @param deleteWithInstance Indicates whether the disk is released together with the instance: Default value: false.
         * 
         * @return builder
         * 
         */
        public Builder deleteWithInstance(@Nullable Output<Boolean> deleteWithInstance) {
            $.deleteWithInstance = deleteWithInstance;
            return this;
        }

        /**
         * @param deleteWithInstance Indicates whether the disk is released together with the instance: Default value: false.
         * 
         * @return builder
         * 
         */
        public Builder deleteWithInstance(Boolean deleteWithInstance) {
            return deleteWithInstance(Output.of(deleteWithInstance));
        }

        /**
         * @param description Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder diskName(@Nullable Output<String> diskName) {
            $.diskName = diskName;
            return this;
        }

        public Builder diskName(String diskName) {
            return diskName(Output.of(diskName));
        }

        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param enableAutoSnapshot Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
         * 
         * @return builder
         * 
         */
        public Builder enableAutoSnapshot(@Nullable Output<Boolean> enableAutoSnapshot) {
            $.enableAutoSnapshot = enableAutoSnapshot;
            return this;
        }

        /**
         * @param enableAutoSnapshot Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
         * 
         * @return builder
         * 
         */
        public Builder enableAutoSnapshot(Boolean enableAutoSnapshot) {
            return enableAutoSnapshot(Output.of(enableAutoSnapshot));
        }

        public Builder encryptAlgorithm(@Nullable Output<String> encryptAlgorithm) {
            $.encryptAlgorithm = encryptAlgorithm;
            return this;
        }

        public Builder encryptAlgorithm(String encryptAlgorithm) {
            return encryptAlgorithm(Output.of(encryptAlgorithm));
        }

        /**
         * @param encrypted If true, the disk will be encrypted, conflict with `snapshot_id`.
         * 
         * @return builder
         * 
         */
        public Builder encrypted(@Nullable Output<Boolean> encrypted) {
            $.encrypted = encrypted;
            return this;
        }

        /**
         * @param encrypted If true, the disk will be encrypted, conflict with `snapshot_id`.
         * 
         * @return builder
         * 
         */
        public Builder encrypted(Boolean encrypted) {
            return encrypted(Output.of(encrypted));
        }

        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param kmsKeyId The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(@Nullable Output<String> kmsKeyId) {
            $.kmsKeyId = kmsKeyId;
            return this;
        }

        /**
         * @param kmsKeyId The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(String kmsKeyId) {
            return kmsKeyId(Output.of(kmsKeyId));
        }

        /**
         * @param name Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.122.0. New field &#39;disk_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead. */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.122.0. New field &#39;disk_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead. */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder paymentType(@Nullable Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        public Builder paymentType(String paymentType) {
            return paymentType(Output.of(paymentType));
        }

        /**
         * @param performanceLevel Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:
         * * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
         * * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
         * * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
         * 
         * @return builder
         * 
         */
        public Builder performanceLevel(@Nullable Output<String> performanceLevel) {
            $.performanceLevel = performanceLevel;
            return this;
        }

        /**
         * @param performanceLevel Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:
         * * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
         * * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
         * * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
         * 
         * @return builder
         * 
         */
        public Builder performanceLevel(String performanceLevel) {
            return performanceLevel(Output.of(performanceLevel));
        }

        /**
         * @param resourceGroupId The Id of resource group which the disk belongs.
         * &gt; **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloud_efficiency` and `cloud_ssd` disk.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The Id of resource group which the disk belongs.
         * &gt; **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloud_efficiency` and `cloud_ssd` disk.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param size The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
         * 
         * @return builder
         * 
         */
        public Builder size(@Nullable Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param snapshotId A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(@Nullable Output<String> snapshotId) {
            $.snapshotId = snapshotId;
            return this;
        }

        /**
         * @param snapshotId A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(String snapshotId) {
            return snapshotId(Output.of(snapshotId));
        }

        /**
         * @param status The disk status.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The disk status.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public Builder storageSetId(@Nullable Output<String> storageSetId) {
            $.storageSetId = storageSetId;
            return this;
        }

        public Builder storageSetId(String storageSetId) {
            return storageSetId(Output.of(storageSetId));
        }

        public Builder storageSetPartitionNumber(@Nullable Output<Integer> storageSetPartitionNumber) {
            $.storageSetPartitionNumber = storageSetPartitionNumber;
            return this;
        }

        public Builder storageSetPartitionNumber(Integer storageSetPartitionNumber) {
            return storageSetPartitionNumber(Output.of(storageSetPartitionNumber));
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

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public Builder zoneId(@Nullable Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public DiskState build() {
            return $;
        }
    }

}
