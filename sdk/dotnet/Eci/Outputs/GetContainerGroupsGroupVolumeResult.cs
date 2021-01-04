// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eci.Outputs
{

    [OutputType]
    public sealed class GetContainerGroupsGroupVolumeResult
    {
        /// <summary>
        /// The list of configuration file paths.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetContainerGroupsGroupVolumeConfigFileVolumeConfigFileToPathResult> ConfigFileVolumeConfigFileToPaths;
        /// <summary>
        /// The ID of DiskVolume.
        /// </summary>
        public readonly string DiskVolumeDiskId;
        /// <summary>
        /// The type of DiskVolume.
        /// </summary>
        public readonly string DiskVolumeFsType;
        /// <summary>
        /// The name of the FlexVolume driver.
        /// </summary>
        public readonly string FlexVolumeDriver;
        /// <summary>
        /// The type of the mounted file system. The default value is determined by the script of FlexVolume.
        /// </summary>
        public readonly string FlexVolumeFsType;
        /// <summary>
        /// The list of FlexVolume objects.
        /// </summary>
        public readonly string FlexVolumeOptions;
        /// <summary>
        /// The name of the volume.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The path to the NFS volume.
        /// </summary>
        public readonly string NfsVolumePath;
        /// <summary>
        /// Default value: `false`.
        /// </summary>
        public readonly bool NfsVolumeReadOnly;
        /// <summary>
        /// The address of the NFS server.
        /// </summary>
        public readonly string NfsVolumeServer;
        /// <summary>
        /// The type of the volume. Currently, the following types of volumes are supported: EmptyDirVolume, NFSVolume, ConfigFileVolume, and FlexVolume.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetContainerGroupsGroupVolumeResult(
            ImmutableArray<Outputs.GetContainerGroupsGroupVolumeConfigFileVolumeConfigFileToPathResult> configFileVolumeConfigFileToPaths,

            string diskVolumeDiskId,

            string diskVolumeFsType,

            string flexVolumeDriver,

            string flexVolumeFsType,

            string flexVolumeOptions,

            string name,

            string nfsVolumePath,

            bool nfsVolumeReadOnly,

            string nfsVolumeServer,

            string type)
        {
            ConfigFileVolumeConfigFileToPaths = configFileVolumeConfigFileToPaths;
            DiskVolumeDiskId = diskVolumeDiskId;
            DiskVolumeFsType = diskVolumeFsType;
            FlexVolumeDriver = flexVolumeDriver;
            FlexVolumeFsType = flexVolumeFsType;
            FlexVolumeOptions = flexVolumeOptions;
            Name = name;
            NfsVolumePath = nfsVolumePath;
            NfsVolumeReadOnly = nfsVolumeReadOnly;
            NfsVolumeServer = nfsVolumeServer;
            Type = type;
        }
    }
}
