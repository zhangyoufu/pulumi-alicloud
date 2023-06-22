// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess.inputs;

import com.pulumi.alicloud.ess.inputs.EciScalingConfigurationVolumeConfigFileVolumeConfigFileToPathArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EciScalingConfigurationVolumeArgs extends com.pulumi.resources.ResourceArgs {

    public static final EciScalingConfigurationVolumeArgs Empty = new EciScalingConfigurationVolumeArgs();

    /**
     * ConfigFileVolumeConfigFileToPaths.
     * See `config_file_volume_config_file_to_paths` below for details.
     * 
     */
    @Import(name="configFileVolumeConfigFileToPaths")
    private @Nullable Output<List<EciScalingConfigurationVolumeConfigFileVolumeConfigFileToPathArgs>> configFileVolumeConfigFileToPaths;

    /**
     * @return ConfigFileVolumeConfigFileToPaths.
     * See `config_file_volume_config_file_to_paths` below for details.
     * 
     */
    public Optional<Output<List<EciScalingConfigurationVolumeConfigFileVolumeConfigFileToPathArgs>>> configFileVolumeConfigFileToPaths() {
        return Optional.ofNullable(this.configFileVolumeConfigFileToPaths);
    }

    /**
     * The ID of DiskVolume.
     * 
     */
    @Import(name="diskVolumeDiskId")
    private @Nullable Output<String> diskVolumeDiskId;

    /**
     * @return The ID of DiskVolume.
     * 
     */
    public Optional<Output<String>> diskVolumeDiskId() {
        return Optional.ofNullable(this.diskVolumeDiskId);
    }

    /**
     * The disk size of DiskVolume.
     * 
     */
    @Import(name="diskVolumeDiskSize")
    private @Nullable Output<Integer> diskVolumeDiskSize;

    /**
     * @return The disk size of DiskVolume.
     * 
     */
    public Optional<Output<Integer>> diskVolumeDiskSize() {
        return Optional.ofNullable(this.diskVolumeDiskSize);
    }

    /**
     * The system type of DiskVolume.
     * 
     */
    @Import(name="diskVolumeFsType")
    private @Nullable Output<String> diskVolumeFsType;

    /**
     * @return The system type of DiskVolume.
     * 
     */
    public Optional<Output<String>> diskVolumeFsType() {
        return Optional.ofNullable(this.diskVolumeFsType);
    }

    /**
     * The name of the FlexVolume driver.
     * 
     */
    @Import(name="flexVolumeDriver")
    private @Nullable Output<String> flexVolumeDriver;

    /**
     * @return The name of the FlexVolume driver.
     * 
     */
    public Optional<Output<String>> flexVolumeDriver() {
        return Optional.ofNullable(this.flexVolumeDriver);
    }

    /**
     * The type of the mounted file system. The default value is determined by the script
     * of FlexVolume.
     * 
     */
    @Import(name="flexVolumeFsType")
    private @Nullable Output<String> flexVolumeFsType;

    /**
     * @return The type of the mounted file system. The default value is determined by the script
     * of FlexVolume.
     * 
     */
    public Optional<Output<String>> flexVolumeFsType() {
        return Optional.ofNullable(this.flexVolumeFsType);
    }

    /**
     * The list of FlexVolume objects. Each object is a key-value pair contained in a JSON
     * string.
     * 
     */
    @Import(name="flexVolumeOptions")
    private @Nullable Output<String> flexVolumeOptions;

    /**
     * @return The list of FlexVolume objects. Each object is a key-value pair contained in a JSON
     * string.
     * 
     */
    public Optional<Output<String>> flexVolumeOptions() {
        return Optional.ofNullable(this.flexVolumeOptions);
    }

    /**
     * The name of the volume.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the volume.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The path to the NFS volume.
     * 
     */
    @Import(name="nfsVolumePath")
    private @Nullable Output<String> nfsVolumePath;

    /**
     * @return The path to the NFS volume.
     * 
     */
    public Optional<Output<String>> nfsVolumePath() {
        return Optional.ofNullable(this.nfsVolumePath);
    }

    /**
     * The nfs volume read only. Default to `false`.
     * 
     */
    @Import(name="nfsVolumeReadOnly")
    private @Nullable Output<Boolean> nfsVolumeReadOnly;

    /**
     * @return The nfs volume read only. Default to `false`.
     * 
     */
    public Optional<Output<Boolean>> nfsVolumeReadOnly() {
        return Optional.ofNullable(this.nfsVolumeReadOnly);
    }

    /**
     * The address of the NFS server.
     * 
     * &gt; **NOTE:** Every volume mounted must have a name and type attributes.
     * 
     */
    @Import(name="nfsVolumeServer")
    private @Nullable Output<String> nfsVolumeServer;

    /**
     * @return The address of the NFS server.
     * 
     * &gt; **NOTE:** Every volume mounted must have a name and type attributes.
     * 
     */
    public Optional<Output<String>> nfsVolumeServer() {
        return Optional.ofNullable(this.nfsVolumeServer);
    }

    /**
     * The type of the volume.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type of the volume.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private EciScalingConfigurationVolumeArgs() {}

    private EciScalingConfigurationVolumeArgs(EciScalingConfigurationVolumeArgs $) {
        this.configFileVolumeConfigFileToPaths = $.configFileVolumeConfigFileToPaths;
        this.diskVolumeDiskId = $.diskVolumeDiskId;
        this.diskVolumeDiskSize = $.diskVolumeDiskSize;
        this.diskVolumeFsType = $.diskVolumeFsType;
        this.flexVolumeDriver = $.flexVolumeDriver;
        this.flexVolumeFsType = $.flexVolumeFsType;
        this.flexVolumeOptions = $.flexVolumeOptions;
        this.name = $.name;
        this.nfsVolumePath = $.nfsVolumePath;
        this.nfsVolumeReadOnly = $.nfsVolumeReadOnly;
        this.nfsVolumeServer = $.nfsVolumeServer;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EciScalingConfigurationVolumeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EciScalingConfigurationVolumeArgs $;

        public Builder() {
            $ = new EciScalingConfigurationVolumeArgs();
        }

        public Builder(EciScalingConfigurationVolumeArgs defaults) {
            $ = new EciScalingConfigurationVolumeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configFileVolumeConfigFileToPaths ConfigFileVolumeConfigFileToPaths.
         * See `config_file_volume_config_file_to_paths` below for details.
         * 
         * @return builder
         * 
         */
        public Builder configFileVolumeConfigFileToPaths(@Nullable Output<List<EciScalingConfigurationVolumeConfigFileVolumeConfigFileToPathArgs>> configFileVolumeConfigFileToPaths) {
            $.configFileVolumeConfigFileToPaths = configFileVolumeConfigFileToPaths;
            return this;
        }

        /**
         * @param configFileVolumeConfigFileToPaths ConfigFileVolumeConfigFileToPaths.
         * See `config_file_volume_config_file_to_paths` below for details.
         * 
         * @return builder
         * 
         */
        public Builder configFileVolumeConfigFileToPaths(List<EciScalingConfigurationVolumeConfigFileVolumeConfigFileToPathArgs> configFileVolumeConfigFileToPaths) {
            return configFileVolumeConfigFileToPaths(Output.of(configFileVolumeConfigFileToPaths));
        }

        /**
         * @param configFileVolumeConfigFileToPaths ConfigFileVolumeConfigFileToPaths.
         * See `config_file_volume_config_file_to_paths` below for details.
         * 
         * @return builder
         * 
         */
        public Builder configFileVolumeConfigFileToPaths(EciScalingConfigurationVolumeConfigFileVolumeConfigFileToPathArgs... configFileVolumeConfigFileToPaths) {
            return configFileVolumeConfigFileToPaths(List.of(configFileVolumeConfigFileToPaths));
        }

        /**
         * @param diskVolumeDiskId The ID of DiskVolume.
         * 
         * @return builder
         * 
         */
        public Builder diskVolumeDiskId(@Nullable Output<String> diskVolumeDiskId) {
            $.diskVolumeDiskId = diskVolumeDiskId;
            return this;
        }

        /**
         * @param diskVolumeDiskId The ID of DiskVolume.
         * 
         * @return builder
         * 
         */
        public Builder diskVolumeDiskId(String diskVolumeDiskId) {
            return diskVolumeDiskId(Output.of(diskVolumeDiskId));
        }

        /**
         * @param diskVolumeDiskSize The disk size of DiskVolume.
         * 
         * @return builder
         * 
         */
        public Builder diskVolumeDiskSize(@Nullable Output<Integer> diskVolumeDiskSize) {
            $.diskVolumeDiskSize = diskVolumeDiskSize;
            return this;
        }

        /**
         * @param diskVolumeDiskSize The disk size of DiskVolume.
         * 
         * @return builder
         * 
         */
        public Builder diskVolumeDiskSize(Integer diskVolumeDiskSize) {
            return diskVolumeDiskSize(Output.of(diskVolumeDiskSize));
        }

        /**
         * @param diskVolumeFsType The system type of DiskVolume.
         * 
         * @return builder
         * 
         */
        public Builder diskVolumeFsType(@Nullable Output<String> diskVolumeFsType) {
            $.diskVolumeFsType = diskVolumeFsType;
            return this;
        }

        /**
         * @param diskVolumeFsType The system type of DiskVolume.
         * 
         * @return builder
         * 
         */
        public Builder diskVolumeFsType(String diskVolumeFsType) {
            return diskVolumeFsType(Output.of(diskVolumeFsType));
        }

        /**
         * @param flexVolumeDriver The name of the FlexVolume driver.
         * 
         * @return builder
         * 
         */
        public Builder flexVolumeDriver(@Nullable Output<String> flexVolumeDriver) {
            $.flexVolumeDriver = flexVolumeDriver;
            return this;
        }

        /**
         * @param flexVolumeDriver The name of the FlexVolume driver.
         * 
         * @return builder
         * 
         */
        public Builder flexVolumeDriver(String flexVolumeDriver) {
            return flexVolumeDriver(Output.of(flexVolumeDriver));
        }

        /**
         * @param flexVolumeFsType The type of the mounted file system. The default value is determined by the script
         * of FlexVolume.
         * 
         * @return builder
         * 
         */
        public Builder flexVolumeFsType(@Nullable Output<String> flexVolumeFsType) {
            $.flexVolumeFsType = flexVolumeFsType;
            return this;
        }

        /**
         * @param flexVolumeFsType The type of the mounted file system. The default value is determined by the script
         * of FlexVolume.
         * 
         * @return builder
         * 
         */
        public Builder flexVolumeFsType(String flexVolumeFsType) {
            return flexVolumeFsType(Output.of(flexVolumeFsType));
        }

        /**
         * @param flexVolumeOptions The list of FlexVolume objects. Each object is a key-value pair contained in a JSON
         * string.
         * 
         * @return builder
         * 
         */
        public Builder flexVolumeOptions(@Nullable Output<String> flexVolumeOptions) {
            $.flexVolumeOptions = flexVolumeOptions;
            return this;
        }

        /**
         * @param flexVolumeOptions The list of FlexVolume objects. Each object is a key-value pair contained in a JSON
         * string.
         * 
         * @return builder
         * 
         */
        public Builder flexVolumeOptions(String flexVolumeOptions) {
            return flexVolumeOptions(Output.of(flexVolumeOptions));
        }

        /**
         * @param name The name of the volume.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the volume.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nfsVolumePath The path to the NFS volume.
         * 
         * @return builder
         * 
         */
        public Builder nfsVolumePath(@Nullable Output<String> nfsVolumePath) {
            $.nfsVolumePath = nfsVolumePath;
            return this;
        }

        /**
         * @param nfsVolumePath The path to the NFS volume.
         * 
         * @return builder
         * 
         */
        public Builder nfsVolumePath(String nfsVolumePath) {
            return nfsVolumePath(Output.of(nfsVolumePath));
        }

        /**
         * @param nfsVolumeReadOnly The nfs volume read only. Default to `false`.
         * 
         * @return builder
         * 
         */
        public Builder nfsVolumeReadOnly(@Nullable Output<Boolean> nfsVolumeReadOnly) {
            $.nfsVolumeReadOnly = nfsVolumeReadOnly;
            return this;
        }

        /**
         * @param nfsVolumeReadOnly The nfs volume read only. Default to `false`.
         * 
         * @return builder
         * 
         */
        public Builder nfsVolumeReadOnly(Boolean nfsVolumeReadOnly) {
            return nfsVolumeReadOnly(Output.of(nfsVolumeReadOnly));
        }

        /**
         * @param nfsVolumeServer The address of the NFS server.
         * 
         * &gt; **NOTE:** Every volume mounted must have a name and type attributes.
         * 
         * @return builder
         * 
         */
        public Builder nfsVolumeServer(@Nullable Output<String> nfsVolumeServer) {
            $.nfsVolumeServer = nfsVolumeServer;
            return this;
        }

        /**
         * @param nfsVolumeServer The address of the NFS server.
         * 
         * &gt; **NOTE:** Every volume mounted must have a name and type attributes.
         * 
         * @return builder
         * 
         */
        public Builder nfsVolumeServer(String nfsVolumeServer) {
            return nfsVolumeServer(Output.of(nfsVolumeServer));
        }

        /**
         * @param type The type of the volume.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of the volume.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public EciScalingConfigurationVolumeArgs build() {
            return $;
        }
    }

}
