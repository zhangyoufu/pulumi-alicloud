// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ScalingConfigurationDataDiskArgs extends com.pulumi.resources.ResourceArgs {

    public static final ScalingConfigurationDataDiskArgs Empty = new ScalingConfigurationDataDiskArgs();

    /**
     * The id of auto snapshot policy for data disk.
     * 
     */
    @Import(name="autoSnapshotPolicyId")
    private @Nullable Output<String> autoSnapshotPolicyId;

    /**
     * @return The id of auto snapshot policy for data disk.
     * 
     */
    public Optional<Output<String>> autoSnapshotPolicyId() {
        return Optional.ofNullable(this.autoSnapshotPolicyId);
    }

    /**
     * Category of data disk. The parameter value options are `ephemeral_ssd`, `cloud_efficiency`, `cloud_ssd` and `cloud`.
     * 
     */
    @Import(name="category")
    private @Nullable Output<String> category;

    /**
     * @return Category of data disk. The parameter value options are `ephemeral_ssd`, `cloud_efficiency`, `cloud_ssd` and `cloud`.
     * 
     */
    public Optional<Output<String>> category() {
        return Optional.ofNullable(this.category);
    }

    /**
     * Whether to delete data disks attached on ecs when release ecs instance. Optional value: `true` or `false`, default to `true`.
     * 
     */
    @Import(name="deleteWithInstance")
    private @Nullable Output<Boolean> deleteWithInstance;

    /**
     * @return Whether to delete data disks attached on ecs when release ecs instance. Optional value: `true` or `false`, default to `true`.
     * 
     */
    public Optional<Output<Boolean>> deleteWithInstance() {
        return Optional.ofNullable(this.deleteWithInstance);
    }

    /**
     * The description of data disk N. Valid values of N: 1 to 16. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of data disk N. Valid values of N: 1 to 16. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The mount point of data disk N. Valid values of N: 1 to 16. If this parameter is not specified, the system automatically allocates a mount point to created ECS instances. The name of the mount point ranges from /dev/xvdb to /dev/xvdz in alphabetical order.
     * 
     * @deprecated
     * Attribute device has been deprecated on disk attachment resource. Suggest to remove it from your template.
     * 
     */
    @Deprecated /* Attribute device has been deprecated on disk attachment resource. Suggest to remove it from your template. */
    @Import(name="device")
    private @Nullable Output<String> device;

    /**
     * @return The mount point of data disk N. Valid values of N: 1 to 16. If this parameter is not specified, the system automatically allocates a mount point to created ECS instances. The name of the mount point ranges from /dev/xvdb to /dev/xvdz in alphabetical order.
     * 
     * @deprecated
     * Attribute device has been deprecated on disk attachment resource. Suggest to remove it from your template.
     * 
     */
    @Deprecated /* Attribute device has been deprecated on disk attachment resource. Suggest to remove it from your template. */
    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    /**
     * Specifies whether data disk N is to be encrypted. Valid values of N: 1 to 16. Valid values: `true`: encrypted, `false`: not encrypted. Default value: `false`.
     * 
     */
    @Import(name="encrypted")
    private @Nullable Output<Boolean> encrypted;

    /**
     * @return Specifies whether data disk N is to be encrypted. Valid values of N: 1 to 16. Valid values: `true`: encrypted, `false`: not encrypted. Default value: `false`.
     * 
     */
    public Optional<Output<Boolean>> encrypted() {
        return Optional.ofNullable(this.encrypted);
    }

    /**
     * The CMK ID for data disk N. Valid values of N: 1 to 16.
     * 
     */
    @Import(name="kmsKeyId")
    private @Nullable Output<String> kmsKeyId;

    /**
     * @return The CMK ID for data disk N. Valid values of N: 1 to 16.
     * 
     */
    public Optional<Output<String>> kmsKeyId() {
        return Optional.ofNullable(this.kmsKeyId);
    }

    /**
     * The name of data disk N. Valid values of N: 1 to 16. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of data disk N. Valid values of N: 1 to 16. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The performance level of the ESSD used as data disk.
     * 
     */
    @Import(name="performanceLevel")
    private @Nullable Output<String> performanceLevel;

    /**
     * @return The performance level of the ESSD used as data disk.
     * 
     */
    public Optional<Output<String>> performanceLevel() {
        return Optional.ofNullable(this.performanceLevel);
    }

    /**
     * Size of data disk, in GB. The value ranges [5,2000] for a cloud disk, [5,1024] for an ephemeral disk, [5,800] for an ephemeral_ssd disk, [20,32768] for cloud_efficiency, cloud_ssd, cloud_essd disk.
     * 
     */
    @Import(name="size")
    private @Nullable Output<Integer> size;

    /**
     * @return Size of data disk, in GB. The value ranges [5,2000] for a cloud disk, [5,1024] for an ephemeral disk, [5,800] for an ephemeral_ssd disk, [20,32768] for cloud_efficiency, cloud_ssd, cloud_essd disk.
     * 
     */
    public Optional<Output<Integer>> size() {
        return Optional.ofNullable(this.size);
    }

    /**
     * Snapshot used for creating the data disk. If this parameter is specified, the size parameter is neglected, and the size of the created disk is the size of the snapshot.
     * 
     */
    @Import(name="snapshotId")
    private @Nullable Output<String> snapshotId;

    /**
     * @return Snapshot used for creating the data disk. If this parameter is specified, the size parameter is neglected, and the size of the created disk is the size of the snapshot.
     * 
     */
    public Optional<Output<String>> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
    }

    private ScalingConfigurationDataDiskArgs() {}

    private ScalingConfigurationDataDiskArgs(ScalingConfigurationDataDiskArgs $) {
        this.autoSnapshotPolicyId = $.autoSnapshotPolicyId;
        this.category = $.category;
        this.deleteWithInstance = $.deleteWithInstance;
        this.description = $.description;
        this.device = $.device;
        this.encrypted = $.encrypted;
        this.kmsKeyId = $.kmsKeyId;
        this.name = $.name;
        this.performanceLevel = $.performanceLevel;
        this.size = $.size;
        this.snapshotId = $.snapshotId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ScalingConfigurationDataDiskArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ScalingConfigurationDataDiskArgs $;

        public Builder() {
            $ = new ScalingConfigurationDataDiskArgs();
        }

        public Builder(ScalingConfigurationDataDiskArgs defaults) {
            $ = new ScalingConfigurationDataDiskArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoSnapshotPolicyId The id of auto snapshot policy for data disk.
         * 
         * @return builder
         * 
         */
        public Builder autoSnapshotPolicyId(@Nullable Output<String> autoSnapshotPolicyId) {
            $.autoSnapshotPolicyId = autoSnapshotPolicyId;
            return this;
        }

        /**
         * @param autoSnapshotPolicyId The id of auto snapshot policy for data disk.
         * 
         * @return builder
         * 
         */
        public Builder autoSnapshotPolicyId(String autoSnapshotPolicyId) {
            return autoSnapshotPolicyId(Output.of(autoSnapshotPolicyId));
        }

        /**
         * @param category Category of data disk. The parameter value options are `ephemeral_ssd`, `cloud_efficiency`, `cloud_ssd` and `cloud`.
         * 
         * @return builder
         * 
         */
        public Builder category(@Nullable Output<String> category) {
            $.category = category;
            return this;
        }

        /**
         * @param category Category of data disk. The parameter value options are `ephemeral_ssd`, `cloud_efficiency`, `cloud_ssd` and `cloud`.
         * 
         * @return builder
         * 
         */
        public Builder category(String category) {
            return category(Output.of(category));
        }

        /**
         * @param deleteWithInstance Whether to delete data disks attached on ecs when release ecs instance. Optional value: `true` or `false`, default to `true`.
         * 
         * @return builder
         * 
         */
        public Builder deleteWithInstance(@Nullable Output<Boolean> deleteWithInstance) {
            $.deleteWithInstance = deleteWithInstance;
            return this;
        }

        /**
         * @param deleteWithInstance Whether to delete data disks attached on ecs when release ecs instance. Optional value: `true` or `false`, default to `true`.
         * 
         * @return builder
         * 
         */
        public Builder deleteWithInstance(Boolean deleteWithInstance) {
            return deleteWithInstance(Output.of(deleteWithInstance));
        }

        /**
         * @param description The description of data disk N. Valid values of N: 1 to 16. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of data disk N. Valid values of N: 1 to 16. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param device The mount point of data disk N. Valid values of N: 1 to 16. If this parameter is not specified, the system automatically allocates a mount point to created ECS instances. The name of the mount point ranges from /dev/xvdb to /dev/xvdz in alphabetical order.
         * 
         * @return builder
         * 
         * @deprecated
         * Attribute device has been deprecated on disk attachment resource. Suggest to remove it from your template.
         * 
         */
        @Deprecated /* Attribute device has been deprecated on disk attachment resource. Suggest to remove it from your template. */
        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        /**
         * @param device The mount point of data disk N. Valid values of N: 1 to 16. If this parameter is not specified, the system automatically allocates a mount point to created ECS instances. The name of the mount point ranges from /dev/xvdb to /dev/xvdz in alphabetical order.
         * 
         * @return builder
         * 
         * @deprecated
         * Attribute device has been deprecated on disk attachment resource. Suggest to remove it from your template.
         * 
         */
        @Deprecated /* Attribute device has been deprecated on disk attachment resource. Suggest to remove it from your template. */
        public Builder device(String device) {
            return device(Output.of(device));
        }

        /**
         * @param encrypted Specifies whether data disk N is to be encrypted. Valid values of N: 1 to 16. Valid values: `true`: encrypted, `false`: not encrypted. Default value: `false`.
         * 
         * @return builder
         * 
         */
        public Builder encrypted(@Nullable Output<Boolean> encrypted) {
            $.encrypted = encrypted;
            return this;
        }

        /**
         * @param encrypted Specifies whether data disk N is to be encrypted. Valid values of N: 1 to 16. Valid values: `true`: encrypted, `false`: not encrypted. Default value: `false`.
         * 
         * @return builder
         * 
         */
        public Builder encrypted(Boolean encrypted) {
            return encrypted(Output.of(encrypted));
        }

        /**
         * @param kmsKeyId The CMK ID for data disk N. Valid values of N: 1 to 16.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(@Nullable Output<String> kmsKeyId) {
            $.kmsKeyId = kmsKeyId;
            return this;
        }

        /**
         * @param kmsKeyId The CMK ID for data disk N. Valid values of N: 1 to 16.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(String kmsKeyId) {
            return kmsKeyId(Output.of(kmsKeyId));
        }

        /**
         * @param name The name of data disk N. Valid values of N: 1 to 16. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of data disk N. Valid values of N: 1 to 16. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param performanceLevel The performance level of the ESSD used as data disk.
         * 
         * @return builder
         * 
         */
        public Builder performanceLevel(@Nullable Output<String> performanceLevel) {
            $.performanceLevel = performanceLevel;
            return this;
        }

        /**
         * @param performanceLevel The performance level of the ESSD used as data disk.
         * 
         * @return builder
         * 
         */
        public Builder performanceLevel(String performanceLevel) {
            return performanceLevel(Output.of(performanceLevel));
        }

        /**
         * @param size Size of data disk, in GB. The value ranges [5,2000] for a cloud disk, [5,1024] for an ephemeral disk, [5,800] for an ephemeral_ssd disk, [20,32768] for cloud_efficiency, cloud_ssd, cloud_essd disk.
         * 
         * @return builder
         * 
         */
        public Builder size(@Nullable Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size Size of data disk, in GB. The value ranges [5,2000] for a cloud disk, [5,1024] for an ephemeral disk, [5,800] for an ephemeral_ssd disk, [20,32768] for cloud_efficiency, cloud_ssd, cloud_essd disk.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param snapshotId Snapshot used for creating the data disk. If this parameter is specified, the size parameter is neglected, and the size of the created disk is the size of the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(@Nullable Output<String> snapshotId) {
            $.snapshotId = snapshotId;
            return this;
        }

        /**
         * @param snapshotId Snapshot used for creating the data disk. If this parameter is specified, the size parameter is neglected, and the size of the created disk is the size of the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(String snapshotId) {
            return snapshotId(Output.of(snapshotId));
        }

        public ScalingConfigurationDataDiskArgs build() {
            return $;
        }
    }

}
