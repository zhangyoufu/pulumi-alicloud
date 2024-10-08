// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ens.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DiskInstanceAttachmentState extends com.pulumi.resources.ResourceArgs {

    public static final DiskInstanceAttachmentState Empty = new DiskInstanceAttachmentState();

    /**
     * Whether the cloud disk to be mounted is released with the instance  Value: true: When the instance is released, the cloud disk is released together with the instance. false: When the instance is released, the cloud disk is retained and is not released together with the instance. Empty means false by default.
     * 
     */
    @Import(name="deleteWithInstance")
    private @Nullable Output<String> deleteWithInstance;

    /**
     * @return Whether the cloud disk to be mounted is released with the instance  Value: true: When the instance is released, the cloud disk is released together with the instance. false: When the instance is released, the cloud disk is retained and is not released together with the instance. Empty means false by default.
     * 
     */
    public Optional<Output<String>> deleteWithInstance() {
        return Optional.ofNullable(this.deleteWithInstance);
    }

    /**
     * The ID of the cloud disk to be mounted. The Cloud Disk (DiskId) and the instance (InstanceId) must be on the same node.
     * 
     */
    @Import(name="diskId")
    private @Nullable Output<String> diskId;

    /**
     * @return The ID of the cloud disk to be mounted. The Cloud Disk (DiskId) and the instance (InstanceId) must be on the same node.
     * 
     */
    public Optional<Output<String>> diskId() {
        return Optional.ofNullable(this.diskId);
    }

    /**
     * Instance ID.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return Instance ID.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    private DiskInstanceAttachmentState() {}

    private DiskInstanceAttachmentState(DiskInstanceAttachmentState $) {
        this.deleteWithInstance = $.deleteWithInstance;
        this.diskId = $.diskId;
        this.instanceId = $.instanceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DiskInstanceAttachmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DiskInstanceAttachmentState $;

        public Builder() {
            $ = new DiskInstanceAttachmentState();
        }

        public Builder(DiskInstanceAttachmentState defaults) {
            $ = new DiskInstanceAttachmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param deleteWithInstance Whether the cloud disk to be mounted is released with the instance  Value: true: When the instance is released, the cloud disk is released together with the instance. false: When the instance is released, the cloud disk is retained and is not released together with the instance. Empty means false by default.
         * 
         * @return builder
         * 
         */
        public Builder deleteWithInstance(@Nullable Output<String> deleteWithInstance) {
            $.deleteWithInstance = deleteWithInstance;
            return this;
        }

        /**
         * @param deleteWithInstance Whether the cloud disk to be mounted is released with the instance  Value: true: When the instance is released, the cloud disk is released together with the instance. false: When the instance is released, the cloud disk is retained and is not released together with the instance. Empty means false by default.
         * 
         * @return builder
         * 
         */
        public Builder deleteWithInstance(String deleteWithInstance) {
            return deleteWithInstance(Output.of(deleteWithInstance));
        }

        /**
         * @param diskId The ID of the cloud disk to be mounted. The Cloud Disk (DiskId) and the instance (InstanceId) must be on the same node.
         * 
         * @return builder
         * 
         */
        public Builder diskId(@Nullable Output<String> diskId) {
            $.diskId = diskId;
            return this;
        }

        /**
         * @param diskId The ID of the cloud disk to be mounted. The Cloud Disk (DiskId) and the instance (InstanceId) must be on the same node.
         * 
         * @return builder
         * 
         */
        public Builder diskId(String diskId) {
            return diskId(Output.of(diskId));
        }

        /**
         * @param instanceId Instance ID.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId Instance ID.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        public DiskInstanceAttachmentState build() {
            return $;
        }
    }

}
