// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.inputs;

import com.pulumi.alicloud.ecs.inputs.GetEcsDisksOperationLock;
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


public final class GetEcsDisksPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEcsDisksPlainArgs Empty = new GetEcsDisksPlainArgs();

    /**
     * Other attribute values. Currently, only the incoming value of IOPS is supported, which means to query the IOPS upper limit of the current disk.
     * 
     */
    @Import(name="additionalAttributes")
    private @Nullable List<String> additionalAttributes;

    /**
     * @return Other attribute values. Currently, only the incoming value of IOPS is supported, which means to query the IOPS upper limit of the current disk.
     * 
     */
    public Optional<List<String>> additionalAttributes() {
        return Optional.ofNullable(this.additionalAttributes);
    }

    /**
     * Query cloud disks based on the automatic snapshot policy ID.
     * 
     */
    @Import(name="autoSnapshotPolicyId")
    private @Nullable String autoSnapshotPolicyId;

    /**
     * @return Query cloud disks based on the automatic snapshot policy ID.
     * 
     */
    public Optional<String> autoSnapshotPolicyId() {
        return Optional.ofNullable(this.autoSnapshotPolicyId);
    }

    /**
     * Availability zone of the disk.
     * 
     * @deprecated
     * Field &#39;availability_zone&#39; has been deprecated from provider version 1.122.0. New field &#39;zone_id&#39; instead
     * 
     */
    @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead */
    @Import(name="availabilityZone")
    private @Nullable String availabilityZone;

    /**
     * @return Availability zone of the disk.
     * 
     * @deprecated
     * Field &#39;availability_zone&#39; has been deprecated from provider version 1.122.0. New field &#39;zone_id&#39; instead
     * 
     */
    @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead */
    public Optional<String> availabilityZone() {
        return Optional.ofNullable(this.availabilityZone);
    }

    /**
     * Disk category.
     * 
     */
    @Import(name="category")
    private @Nullable String category;

    /**
     * @return Disk category.
     * 
     */
    public Optional<String> category() {
        return Optional.ofNullable(this.category);
    }

    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released.
     * 
     */
    @Import(name="deleteAutoSnapshot")
    private @Nullable Boolean deleteAutoSnapshot;

    /**
     * @return Indicates whether the automatic snapshot is deleted when the disk is released.
     * 
     */
    public Optional<Boolean> deleteAutoSnapshot() {
        return Optional.ofNullable(this.deleteAutoSnapshot);
    }

    /**
     * Indicates whether the disk is released together with the instance.
     * 
     */
    @Import(name="deleteWithInstance")
    private @Nullable Boolean deleteWithInstance;

    /**
     * @return Indicates whether the disk is released together with the instance.
     * 
     */
    public Optional<Boolean> deleteWithInstance() {
        return Optional.ofNullable(this.deleteWithInstance);
    }

    /**
     * The disk name.
     * 
     */
    @Import(name="diskName")
    private @Nullable String diskName;

    /**
     * @return The disk name.
     * 
     */
    public Optional<String> diskName() {
        return Optional.ofNullable(this.diskName);
    }

    /**
     * The disk type.
     * 
     */
    @Import(name="diskType")
    private @Nullable String diskType;

    /**
     * @return The disk type.
     * 
     */
    public Optional<String> diskType() {
        return Optional.ofNullable(this.diskType);
    }

    /**
     * Specifies whether to check the validity of the request without actually making the request.request Default value: false. Valid values:
     * 
     */
    @Import(name="dryRun")
    private @Nullable Boolean dryRun;

    /**
     * @return Specifies whether to check the validity of the request without actually making the request.request Default value: false. Valid values:
     * 
     */
    public Optional<Boolean> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * Whether the disk implements an automatic snapshot policy.
     * 
     */
    @Import(name="enableAutoSnapshot")
    private @Nullable Boolean enableAutoSnapshot;

    /**
     * @return Whether the disk implements an automatic snapshot policy.
     * 
     */
    public Optional<Boolean> enableAutoSnapshot() {
        return Optional.ofNullable(this.enableAutoSnapshot);
    }

    /**
     * Whether the disk implements an automatic snapshot policy.
     * 
     */
    @Import(name="enableAutomatedSnapshotPolicy")
    private @Nullable Boolean enableAutomatedSnapshotPolicy;

    /**
     * @return Whether the disk implements an automatic snapshot policy.
     * 
     */
    public Optional<Boolean> enableAutomatedSnapshotPolicy() {
        return Optional.ofNullable(this.enableAutomatedSnapshotPolicy);
    }

    /**
     * Whether it is shared block storage.
     * 
     */
    @Import(name="enableShared")
    private @Nullable Boolean enableShared;

    /**
     * @return Whether it is shared block storage.
     * 
     */
    public Optional<Boolean> enableShared() {
        return Optional.ofNullable(this.enableShared);
    }

    /**
     * Indicate whether the disk is encrypted or not.
     * 
     */
    @Import(name="encrypted")
    private @Nullable String encrypted;

    /**
     * @return Indicate whether the disk is encrypted or not.
     * 
     */
    public Optional<String> encrypted() {
        return Optional.ofNullable(this.encrypted);
    }

    /**
     * A list of Disk IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Disk IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The instance ID of the disk mount.
     * 
     */
    @Import(name="instanceId")
    private @Nullable String instanceId;

    /**
     * @return The instance ID of the disk mount.
     * 
     */
    public Optional<String> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * The kms key id.
     * 
     */
    @Import(name="kmsKeyId")
    private @Nullable String kmsKeyId;

    /**
     * @return The kms key id.
     * 
     */
    public Optional<String> kmsKeyId() {
        return Optional.ofNullable(this.kmsKeyId);
    }

    /**
     * A regex string to filter results by Disk name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Disk name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    @Import(name="operationLocks")
    private @Nullable List<GetEcsDisksOperationLock> operationLocks;

    public Optional<List<GetEcsDisksOperationLock>> operationLocks() {
        return Optional.ofNullable(this.operationLocks);
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    @Import(name="pageNumber")
    private @Nullable Integer pageNumber;

    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }

    @Import(name="pageSize")
    private @Nullable Integer pageSize;

    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }

    /**
     * Payment method for disk.
     * 
     */
    @Import(name="paymentType")
    private @Nullable String paymentType;

    /**
     * @return Payment method for disk.
     * 
     */
    public Optional<String> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }

    /**
     * Whether the disk is unmountable.
     * 
     */
    @Import(name="portable")
    private @Nullable Boolean portable;

    /**
     * @return Whether the disk is unmountable.
     * 
     */
    public Optional<Boolean> portable() {
        return Optional.ofNullable(this.portable);
    }

    /**
     * The Id of resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable String resourceGroupId;

    /**
     * @return The Id of resource group.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * Snapshot used to create the disk. It is null if no snapshot is used to create the disk.
     * 
     */
    @Import(name="snapshotId")
    private @Nullable String snapshotId;

    /**
     * @return Snapshot used to create the disk. It is null if no snapshot is used to create the disk.
     * 
     */
    public Optional<String> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
    }

    /**
     * Current status.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return Current status.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * A map of tags assigned to the disk.
     * 
     */
    @Import(name="tags")
    private @Nullable Map<String,Object> tags;

    /**
     * @return A map of tags assigned to the disk.
     * 
     */
    public Optional<Map<String,Object>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Disk type.
     * 
     * @deprecated
     * Field &#39;type&#39; has been deprecated from provider version 1.122.0. New field &#39;disk_type&#39; instead.
     * 
     */
    @Deprecated /* Field 'type' has been deprecated from provider version 1.122.0. New field 'disk_type' instead. */
    @Import(name="type")
    private @Nullable String type;

    /**
     * @return Disk type.
     * 
     * @deprecated
     * Field &#39;type&#39; has been deprecated from provider version 1.122.0. New field &#39;disk_type&#39; instead.
     * 
     */
    @Deprecated /* Field 'type' has been deprecated from provider version 1.122.0. New field 'disk_type' instead. */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * The zone id.
     * 
     */
    @Import(name="zoneId")
    private @Nullable String zoneId;

    /**
     * @return The zone id.
     * 
     */
    public Optional<String> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    private GetEcsDisksPlainArgs() {}

    private GetEcsDisksPlainArgs(GetEcsDisksPlainArgs $) {
        this.additionalAttributes = $.additionalAttributes;
        this.autoSnapshotPolicyId = $.autoSnapshotPolicyId;
        this.availabilityZone = $.availabilityZone;
        this.category = $.category;
        this.deleteAutoSnapshot = $.deleteAutoSnapshot;
        this.deleteWithInstance = $.deleteWithInstance;
        this.diskName = $.diskName;
        this.diskType = $.diskType;
        this.dryRun = $.dryRun;
        this.enableAutoSnapshot = $.enableAutoSnapshot;
        this.enableAutomatedSnapshotPolicy = $.enableAutomatedSnapshotPolicy;
        this.enableShared = $.enableShared;
        this.encrypted = $.encrypted;
        this.ids = $.ids;
        this.instanceId = $.instanceId;
        this.kmsKeyId = $.kmsKeyId;
        this.nameRegex = $.nameRegex;
        this.operationLocks = $.operationLocks;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
        this.paymentType = $.paymentType;
        this.portable = $.portable;
        this.resourceGroupId = $.resourceGroupId;
        this.snapshotId = $.snapshotId;
        this.status = $.status;
        this.tags = $.tags;
        this.type = $.type;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEcsDisksPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEcsDisksPlainArgs $;

        public Builder() {
            $ = new GetEcsDisksPlainArgs();
        }

        public Builder(GetEcsDisksPlainArgs defaults) {
            $ = new GetEcsDisksPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param additionalAttributes Other attribute values. Currently, only the incoming value of IOPS is supported, which means to query the IOPS upper limit of the current disk.
         * 
         * @return builder
         * 
         */
        public Builder additionalAttributes(@Nullable List<String> additionalAttributes) {
            $.additionalAttributes = additionalAttributes;
            return this;
        }

        /**
         * @param additionalAttributes Other attribute values. Currently, only the incoming value of IOPS is supported, which means to query the IOPS upper limit of the current disk.
         * 
         * @return builder
         * 
         */
        public Builder additionalAttributes(String... additionalAttributes) {
            return additionalAttributes(List.of(additionalAttributes));
        }

        /**
         * @param autoSnapshotPolicyId Query cloud disks based on the automatic snapshot policy ID.
         * 
         * @return builder
         * 
         */
        public Builder autoSnapshotPolicyId(@Nullable String autoSnapshotPolicyId) {
            $.autoSnapshotPolicyId = autoSnapshotPolicyId;
            return this;
        }

        /**
         * @param availabilityZone Availability zone of the disk.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;availability_zone&#39; has been deprecated from provider version 1.122.0. New field &#39;zone_id&#39; instead
         * 
         */
        @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead */
        public Builder availabilityZone(@Nullable String availabilityZone) {
            $.availabilityZone = availabilityZone;
            return this;
        }

        /**
         * @param category Disk category.
         * 
         * @return builder
         * 
         */
        public Builder category(@Nullable String category) {
            $.category = category;
            return this;
        }

        /**
         * @param deleteAutoSnapshot Indicates whether the automatic snapshot is deleted when the disk is released.
         * 
         * @return builder
         * 
         */
        public Builder deleteAutoSnapshot(@Nullable Boolean deleteAutoSnapshot) {
            $.deleteAutoSnapshot = deleteAutoSnapshot;
            return this;
        }

        /**
         * @param deleteWithInstance Indicates whether the disk is released together with the instance.
         * 
         * @return builder
         * 
         */
        public Builder deleteWithInstance(@Nullable Boolean deleteWithInstance) {
            $.deleteWithInstance = deleteWithInstance;
            return this;
        }

        /**
         * @param diskName The disk name.
         * 
         * @return builder
         * 
         */
        public Builder diskName(@Nullable String diskName) {
            $.diskName = diskName;
            return this;
        }

        /**
         * @param diskType The disk type.
         * 
         * @return builder
         * 
         */
        public Builder diskType(@Nullable String diskType) {
            $.diskType = diskType;
            return this;
        }

        /**
         * @param dryRun Specifies whether to check the validity of the request without actually making the request.request Default value: false. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Boolean dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param enableAutoSnapshot Whether the disk implements an automatic snapshot policy.
         * 
         * @return builder
         * 
         */
        public Builder enableAutoSnapshot(@Nullable Boolean enableAutoSnapshot) {
            $.enableAutoSnapshot = enableAutoSnapshot;
            return this;
        }

        /**
         * @param enableAutomatedSnapshotPolicy Whether the disk implements an automatic snapshot policy.
         * 
         * @return builder
         * 
         */
        public Builder enableAutomatedSnapshotPolicy(@Nullable Boolean enableAutomatedSnapshotPolicy) {
            $.enableAutomatedSnapshotPolicy = enableAutomatedSnapshotPolicy;
            return this;
        }

        /**
         * @param enableShared Whether it is shared block storage.
         * 
         * @return builder
         * 
         */
        public Builder enableShared(@Nullable Boolean enableShared) {
            $.enableShared = enableShared;
            return this;
        }

        /**
         * @param encrypted Indicate whether the disk is encrypted or not.
         * 
         * @return builder
         * 
         */
        public Builder encrypted(@Nullable String encrypted) {
            $.encrypted = encrypted;
            return this;
        }

        /**
         * @param ids A list of Disk IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Disk IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceId The instance ID of the disk mount.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable String instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param kmsKeyId The kms key id.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(@Nullable String kmsKeyId) {
            $.kmsKeyId = kmsKeyId;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Disk name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        public Builder operationLocks(@Nullable List<GetEcsDisksOperationLock> operationLocks) {
            $.operationLocks = operationLocks;
            return this;
        }

        public Builder operationLocks(GetEcsDisksOperationLock... operationLocks) {
            return operationLocks(List.of(operationLocks));
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder pageNumber(@Nullable Integer pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageSize(@Nullable Integer pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        /**
         * @param paymentType Payment method for disk.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(@Nullable String paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param portable Whether the disk is unmountable.
         * 
         * @return builder
         * 
         */
        public Builder portable(@Nullable Boolean portable) {
            $.portable = portable;
            return this;
        }

        /**
         * @param resourceGroupId The Id of resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param snapshotId Snapshot used to create the disk. It is null if no snapshot is used to create the disk.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(@Nullable String snapshotId) {
            $.snapshotId = snapshotId;
            return this;
        }

        /**
         * @param status Current status.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param tags A map of tags assigned to the disk.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Map<String,Object> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param type Disk type.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;type&#39; has been deprecated from provider version 1.122.0. New field &#39;disk_type&#39; instead.
         * 
         */
        @Deprecated /* Field 'type' has been deprecated from provider version 1.122.0. New field 'disk_type' instead. */
        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        /**
         * @param zoneId The zone id.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(@Nullable String zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        public GetEcsDisksPlainArgs build() {
            return $;
        }
    }

}
