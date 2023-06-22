// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ScalingConfigurationDataDisk {
    /**
     * @return The id of auto snapshot policy for data disk.
     * 
     */
    private @Nullable String autoSnapshotPolicyId;
    /**
     * @return Category of data disk. The parameter value options are `ephemeral_ssd`, `cloud_efficiency`, `cloud_ssd` and `cloud`.
     * 
     */
    private @Nullable String category;
    /**
     * @return Whether to delete data disks attached on ecs when release ecs instance. Optional value: `true` or `false`, default to `true`.
     * 
     */
    private @Nullable Boolean deleteWithInstance;
    /**
     * @return The description of data disk N. Valid values of N: 1 to 16. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
     * 
     */
    private @Nullable String description;
    /**
     * @return The mount point of data disk N. Valid values of N: 1 to 16. If this parameter is not specified, the system automatically allocates a mount point to created ECS instances. The name of the mount point ranges from /dev/xvdb to /dev/xvdz in alphabetical order.
     * 
     * @deprecated
     * Attribute device has been deprecated on disk attachment resource. Suggest to remove it from your template.
     * 
     */
    @Deprecated /* Attribute device has been deprecated on disk attachment resource. Suggest to remove it from your template. */
    private @Nullable String device;
    /**
     * @return Specifies whether data disk N is to be encrypted. Valid values of N: 1 to 16. Valid values: `true`: encrypted, `false`: not encrypted. Default value: `false`.
     * 
     */
    private @Nullable Boolean encrypted;
    /**
     * @return The CMK ID for data disk N. Valid values of N: 1 to 16.
     * 
     */
    private @Nullable String kmsKeyId;
    /**
     * @return The name of data disk N. Valid values of N: 1 to 16. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.
     * 
     */
    private @Nullable String name;
    /**
     * @return The performance level of the ESSD used as data disk.
     * 
     */
    private @Nullable String performanceLevel;
    /**
     * @return Size of data disk, in GB. The value ranges [5,2000] for a cloud disk, [5,1024] for an ephemeral disk, [5,800] for an ephemeral_ssd disk, [20,32768] for cloud_efficiency, cloud_ssd, cloud_essd disk.
     * 
     */
    private @Nullable Integer size;
    /**
     * @return Snapshot used for creating the data disk. If this parameter is specified, the size parameter is neglected, and the size of the created disk is the size of the snapshot.
     * 
     */
    private @Nullable String snapshotId;

    private ScalingConfigurationDataDisk() {}
    /**
     * @return The id of auto snapshot policy for data disk.
     * 
     */
    public Optional<String> autoSnapshotPolicyId() {
        return Optional.ofNullable(this.autoSnapshotPolicyId);
    }
    /**
     * @return Category of data disk. The parameter value options are `ephemeral_ssd`, `cloud_efficiency`, `cloud_ssd` and `cloud`.
     * 
     */
    public Optional<String> category() {
        return Optional.ofNullable(this.category);
    }
    /**
     * @return Whether to delete data disks attached on ecs when release ecs instance. Optional value: `true` or `false`, default to `true`.
     * 
     */
    public Optional<Boolean> deleteWithInstance() {
        return Optional.ofNullable(this.deleteWithInstance);
    }
    /**
     * @return The description of data disk N. Valid values of N: 1 to 16. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return The mount point of data disk N. Valid values of N: 1 to 16. If this parameter is not specified, the system automatically allocates a mount point to created ECS instances. The name of the mount point ranges from /dev/xvdb to /dev/xvdz in alphabetical order.
     * 
     * @deprecated
     * Attribute device has been deprecated on disk attachment resource. Suggest to remove it from your template.
     * 
     */
    @Deprecated /* Attribute device has been deprecated on disk attachment resource. Suggest to remove it from your template. */
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    /**
     * @return Specifies whether data disk N is to be encrypted. Valid values of N: 1 to 16. Valid values: `true`: encrypted, `false`: not encrypted. Default value: `false`.
     * 
     */
    public Optional<Boolean> encrypted() {
        return Optional.ofNullable(this.encrypted);
    }
    /**
     * @return The CMK ID for data disk N. Valid values of N: 1 to 16.
     * 
     */
    public Optional<String> kmsKeyId() {
        return Optional.ofNullable(this.kmsKeyId);
    }
    /**
     * @return The name of data disk N. Valid values of N: 1 to 16. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The performance level of the ESSD used as data disk.
     * 
     */
    public Optional<String> performanceLevel() {
        return Optional.ofNullable(this.performanceLevel);
    }
    /**
     * @return Size of data disk, in GB. The value ranges [5,2000] for a cloud disk, [5,1024] for an ephemeral disk, [5,800] for an ephemeral_ssd disk, [20,32768] for cloud_efficiency, cloud_ssd, cloud_essd disk.
     * 
     */
    public Optional<Integer> size() {
        return Optional.ofNullable(this.size);
    }
    /**
     * @return Snapshot used for creating the data disk. If this parameter is specified, the size parameter is neglected, and the size of the created disk is the size of the snapshot.
     * 
     */
    public Optional<String> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ScalingConfigurationDataDisk defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String autoSnapshotPolicyId;
        private @Nullable String category;
        private @Nullable Boolean deleteWithInstance;
        private @Nullable String description;
        private @Nullable String device;
        private @Nullable Boolean encrypted;
        private @Nullable String kmsKeyId;
        private @Nullable String name;
        private @Nullable String performanceLevel;
        private @Nullable Integer size;
        private @Nullable String snapshotId;
        public Builder() {}
        public Builder(ScalingConfigurationDataDisk defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.autoSnapshotPolicyId = defaults.autoSnapshotPolicyId;
    	      this.category = defaults.category;
    	      this.deleteWithInstance = defaults.deleteWithInstance;
    	      this.description = defaults.description;
    	      this.device = defaults.device;
    	      this.encrypted = defaults.encrypted;
    	      this.kmsKeyId = defaults.kmsKeyId;
    	      this.name = defaults.name;
    	      this.performanceLevel = defaults.performanceLevel;
    	      this.size = defaults.size;
    	      this.snapshotId = defaults.snapshotId;
        }

        @CustomType.Setter
        public Builder autoSnapshotPolicyId(@Nullable String autoSnapshotPolicyId) {
            this.autoSnapshotPolicyId = autoSnapshotPolicyId;
            return this;
        }
        @CustomType.Setter
        public Builder category(@Nullable String category) {
            this.category = category;
            return this;
        }
        @CustomType.Setter
        public Builder deleteWithInstance(@Nullable Boolean deleteWithInstance) {
            this.deleteWithInstance = deleteWithInstance;
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        @CustomType.Setter
        public Builder encrypted(@Nullable Boolean encrypted) {
            this.encrypted = encrypted;
            return this;
        }
        @CustomType.Setter
        public Builder kmsKeyId(@Nullable String kmsKeyId) {
            this.kmsKeyId = kmsKeyId;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder performanceLevel(@Nullable String performanceLevel) {
            this.performanceLevel = performanceLevel;
            return this;
        }
        @CustomType.Setter
        public Builder size(@Nullable Integer size) {
            this.size = size;
            return this;
        }
        @CustomType.Setter
        public Builder snapshotId(@Nullable String snapshotId) {
            this.snapshotId = snapshotId;
            return this;
        }
        public ScalingConfigurationDataDisk build() {
            final var o = new ScalingConfigurationDataDisk();
            o.autoSnapshotPolicyId = autoSnapshotPolicyId;
            o.category = category;
            o.deleteWithInstance = deleteWithInstance;
            o.description = description;
            o.device = device;
            o.encrypted = encrypted;
            o.kmsKeyId = kmsKeyId;
            o.name = name;
            o.performanceLevel = performanceLevel;
            o.size = size;
            o.snapshotId = snapshotId;
            return o;
        }
    }
}
