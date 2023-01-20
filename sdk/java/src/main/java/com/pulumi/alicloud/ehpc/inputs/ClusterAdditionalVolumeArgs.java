// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ehpc.inputs;

import com.pulumi.alicloud.ehpc.inputs.ClusterAdditionalVolumeRoleArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterAdditionalVolumeArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterAdditionalVolumeArgs Empty = new ClusterAdditionalVolumeArgs();

    /**
     * The queue to which the compute nodes are added.
     * 
     */
    @Import(name="jobQueue")
    private @Nullable Output<String> jobQueue;

    /**
     * @return The queue to which the compute nodes are added.
     * 
     */
    public Optional<Output<String>> jobQueue() {
        return Optional.ofNullable(this.jobQueue);
    }

    /**
     * The local directory on which the additional file system is mounted.
     * 
     */
    @Import(name="localDirectory")
    private @Nullable Output<String> localDirectory;

    /**
     * @return The local directory on which the additional file system is mounted.
     * 
     */
    public Optional<Output<String>> localDirectory() {
        return Optional.ofNullable(this.localDirectory);
    }

    /**
     * The type of the cluster. Valid value: `PublicCloud`.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The type of the cluster. Valid value: `PublicCloud`.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The remote directory to which the file system is mounted.
     * 
     */
    @Import(name="remoteDirectory")
    private @Nullable Output<String> remoteDirectory;

    /**
     * @return The remote directory to which the file system is mounted.
     * 
     */
    public Optional<Output<String>> remoteDirectory() {
        return Optional.ofNullable(this.remoteDirectory);
    }

    /**
     * The roles. See the following `Block roles`.
     * 
     */
    @Import(name="roles")
    private @Nullable Output<List<ClusterAdditionalVolumeRoleArgs>> roles;

    /**
     * @return The roles. See the following `Block roles`.
     * 
     */
    public Optional<Output<List<ClusterAdditionalVolumeRoleArgs>>> roles() {
        return Optional.ofNullable(this.roles);
    }

    /**
     * The ID of the file system. If you leave the parameter empty, a Performance NAS file system is created by default.
     * 
     */
    @Import(name="volumeId")
    private @Nullable Output<String> volumeId;

    /**
     * @return The ID of the file system. If you leave the parameter empty, a Performance NAS file system is created by default.
     * 
     */
    public Optional<Output<String>> volumeId() {
        return Optional.ofNullable(this.volumeId);
    }

    /**
     * The mount options of the file system.
     * 
     */
    @Import(name="volumeMountOption")
    private @Nullable Output<String> volumeMountOption;

    /**
     * @return The mount options of the file system.
     * 
     */
    public Optional<Output<String>> volumeMountOption() {
        return Optional.ofNullable(this.volumeMountOption);
    }

    /**
     * The mount target of the file system. Take note of the following information:
     * - If you do not specify the VolumeId parameter, you can leave the VolumeMountpoint parameter empty. A mount target is created by default.
     * - If you specify the VolumeId parameter, the VolumeMountpoint parameter is required.
     * 
     */
    @Import(name="volumeMountpoint")
    private @Nullable Output<String> volumeMountpoint;

    /**
     * @return The mount target of the file system. Take note of the following information:
     * - If you do not specify the VolumeId parameter, you can leave the VolumeMountpoint parameter empty. A mount target is created by default.
     * - If you specify the VolumeId parameter, the VolumeMountpoint parameter is required.
     * 
     */
    public Optional<Output<String>> volumeMountpoint() {
        return Optional.ofNullable(this.volumeMountpoint);
    }

    /**
     * The type of the protocol that is used by the file system. Valid values: `NFS`, `SMB`. Default value: `NFS`.
     * 
     */
    @Import(name="volumeProtocol")
    private @Nullable Output<String> volumeProtocol;

    /**
     * @return The type of the protocol that is used by the file system. Valid values: `NFS`, `SMB`. Default value: `NFS`.
     * 
     */
    public Optional<Output<String>> volumeProtocol() {
        return Optional.ofNullable(this.volumeProtocol);
    }

    /**
     * The type of the shared storage. Only Apsara File Storage NAS file systems are supported.
     * 
     */
    @Import(name="volumeType")
    private @Nullable Output<String> volumeType;

    /**
     * @return The type of the shared storage. Only Apsara File Storage NAS file systems are supported.
     * 
     */
    public Optional<Output<String>> volumeType() {
        return Optional.ofNullable(this.volumeType);
    }

    private ClusterAdditionalVolumeArgs() {}

    private ClusterAdditionalVolumeArgs(ClusterAdditionalVolumeArgs $) {
        this.jobQueue = $.jobQueue;
        this.localDirectory = $.localDirectory;
        this.location = $.location;
        this.remoteDirectory = $.remoteDirectory;
        this.roles = $.roles;
        this.volumeId = $.volumeId;
        this.volumeMountOption = $.volumeMountOption;
        this.volumeMountpoint = $.volumeMountpoint;
        this.volumeProtocol = $.volumeProtocol;
        this.volumeType = $.volumeType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterAdditionalVolumeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterAdditionalVolumeArgs $;

        public Builder() {
            $ = new ClusterAdditionalVolumeArgs();
        }

        public Builder(ClusterAdditionalVolumeArgs defaults) {
            $ = new ClusterAdditionalVolumeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param jobQueue The queue to which the compute nodes are added.
         * 
         * @return builder
         * 
         */
        public Builder jobQueue(@Nullable Output<String> jobQueue) {
            $.jobQueue = jobQueue;
            return this;
        }

        /**
         * @param jobQueue The queue to which the compute nodes are added.
         * 
         * @return builder
         * 
         */
        public Builder jobQueue(String jobQueue) {
            return jobQueue(Output.of(jobQueue));
        }

        /**
         * @param localDirectory The local directory on which the additional file system is mounted.
         * 
         * @return builder
         * 
         */
        public Builder localDirectory(@Nullable Output<String> localDirectory) {
            $.localDirectory = localDirectory;
            return this;
        }

        /**
         * @param localDirectory The local directory on which the additional file system is mounted.
         * 
         * @return builder
         * 
         */
        public Builder localDirectory(String localDirectory) {
            return localDirectory(Output.of(localDirectory));
        }

        /**
         * @param location The type of the cluster. Valid value: `PublicCloud`.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The type of the cluster. Valid value: `PublicCloud`.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param remoteDirectory The remote directory to which the file system is mounted.
         * 
         * @return builder
         * 
         */
        public Builder remoteDirectory(@Nullable Output<String> remoteDirectory) {
            $.remoteDirectory = remoteDirectory;
            return this;
        }

        /**
         * @param remoteDirectory The remote directory to which the file system is mounted.
         * 
         * @return builder
         * 
         */
        public Builder remoteDirectory(String remoteDirectory) {
            return remoteDirectory(Output.of(remoteDirectory));
        }

        /**
         * @param roles The roles. See the following `Block roles`.
         * 
         * @return builder
         * 
         */
        public Builder roles(@Nullable Output<List<ClusterAdditionalVolumeRoleArgs>> roles) {
            $.roles = roles;
            return this;
        }

        /**
         * @param roles The roles. See the following `Block roles`.
         * 
         * @return builder
         * 
         */
        public Builder roles(List<ClusterAdditionalVolumeRoleArgs> roles) {
            return roles(Output.of(roles));
        }

        /**
         * @param roles The roles. See the following `Block roles`.
         * 
         * @return builder
         * 
         */
        public Builder roles(ClusterAdditionalVolumeRoleArgs... roles) {
            return roles(List.of(roles));
        }

        /**
         * @param volumeId The ID of the file system. If you leave the parameter empty, a Performance NAS file system is created by default.
         * 
         * @return builder
         * 
         */
        public Builder volumeId(@Nullable Output<String> volumeId) {
            $.volumeId = volumeId;
            return this;
        }

        /**
         * @param volumeId The ID of the file system. If you leave the parameter empty, a Performance NAS file system is created by default.
         * 
         * @return builder
         * 
         */
        public Builder volumeId(String volumeId) {
            return volumeId(Output.of(volumeId));
        }

        /**
         * @param volumeMountOption The mount options of the file system.
         * 
         * @return builder
         * 
         */
        public Builder volumeMountOption(@Nullable Output<String> volumeMountOption) {
            $.volumeMountOption = volumeMountOption;
            return this;
        }

        /**
         * @param volumeMountOption The mount options of the file system.
         * 
         * @return builder
         * 
         */
        public Builder volumeMountOption(String volumeMountOption) {
            return volumeMountOption(Output.of(volumeMountOption));
        }

        /**
         * @param volumeMountpoint The mount target of the file system. Take note of the following information:
         * - If you do not specify the VolumeId parameter, you can leave the VolumeMountpoint parameter empty. A mount target is created by default.
         * - If you specify the VolumeId parameter, the VolumeMountpoint parameter is required.
         * 
         * @return builder
         * 
         */
        public Builder volumeMountpoint(@Nullable Output<String> volumeMountpoint) {
            $.volumeMountpoint = volumeMountpoint;
            return this;
        }

        /**
         * @param volumeMountpoint The mount target of the file system. Take note of the following information:
         * - If you do not specify the VolumeId parameter, you can leave the VolumeMountpoint parameter empty. A mount target is created by default.
         * - If you specify the VolumeId parameter, the VolumeMountpoint parameter is required.
         * 
         * @return builder
         * 
         */
        public Builder volumeMountpoint(String volumeMountpoint) {
            return volumeMountpoint(Output.of(volumeMountpoint));
        }

        /**
         * @param volumeProtocol The type of the protocol that is used by the file system. Valid values: `NFS`, `SMB`. Default value: `NFS`.
         * 
         * @return builder
         * 
         */
        public Builder volumeProtocol(@Nullable Output<String> volumeProtocol) {
            $.volumeProtocol = volumeProtocol;
            return this;
        }

        /**
         * @param volumeProtocol The type of the protocol that is used by the file system. Valid values: `NFS`, `SMB`. Default value: `NFS`.
         * 
         * @return builder
         * 
         */
        public Builder volumeProtocol(String volumeProtocol) {
            return volumeProtocol(Output.of(volumeProtocol));
        }

        /**
         * @param volumeType The type of the shared storage. Only Apsara File Storage NAS file systems are supported.
         * 
         * @return builder
         * 
         */
        public Builder volumeType(@Nullable Output<String> volumeType) {
            $.volumeType = volumeType;
            return this;
        }

        /**
         * @param volumeType The type of the shared storage. Only Apsara File Storage NAS file systems are supported.
         * 
         * @return builder
         * 
         */
        public Builder volumeType(String volumeType) {
            return volumeType(Output.of(volumeType));
        }

        public ClusterAdditionalVolumeArgs build() {
            return $;
        }
    }

}
