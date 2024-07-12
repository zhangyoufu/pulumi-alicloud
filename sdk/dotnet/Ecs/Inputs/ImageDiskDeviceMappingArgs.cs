// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs.Inputs
{

    public sealed class ImageDiskDeviceMappingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The device name of disk N in the custom image. Valid values:
        /// - For disks other than basic disks, such as standard SSDs, ultra disks, and enhanced SSDs (ESSDs), the valid values range from /dev/vda to /dev/vdz in alphabetical order.
        /// - For basic disks, the valid values range from /dev/xvda to /dev/xvdz in alphabetical order.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// The type of disk N in the custom image. You can specify this parameter to create the system disk of the custom image from a data disk snapshot. If you do not specify this parameter, the disk type is determined by the corresponding snapshot. Valid values:
        /// - system: system disk. You can specify only one snapshot to use to create the system disk in the custom image.
        /// - data: data disk. You can specify up to 16 snapshots to use to create data disks in the custom image.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// Image format.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// Import the bucket of the OSS to which the image belongs.
        /// </summary>
        [Input("importOssBucket")]
        public Input<string>? ImportOssBucket { get; set; }

        /// <summary>
        /// Import the object of the OSS to which the image file belongs.
        /// </summary>
        [Input("importOssObject")]
        public Input<string>? ImportOssObject { get; set; }

        /// <summary>
        /// Copy the progress of the task.
        /// </summary>
        [Input("progress")]
        public Input<string>? Progress { get; set; }

        /// <summary>
        /// For an image being replicated, return the remaining time of the replication task, in seconds.
        /// </summary>
        [Input("remainTime")]
        public Input<int>? RemainTime { get; set; }

        /// <summary>
        /// The size of disk N in the custom image. Unit: GiB. The valid values and default value of DiskDeviceMapping.N.Size vary based on the value of DiskDeviceMapping.N.SnapshotId.
        /// - If no corresponding snapshot IDs are specified in the value of DiskDeviceMapping.N.SnapshotId, DiskDeviceMapping.N.Size has the following valid values and default values:
        /// *   For basic disks, the valid values range from 5 to 2000, and the default value is 5.
        /// *   For other disks, the valid values range from 20 to 32768, and the default value is 20.
        /// - If a corresponding snapshot ID is specified in the value of DiskDeviceMapping.N.SnapshotId, the value of DiskDeviceMapping.N.Size must be greater than or equal to the size of the specified snapshot. The default value of DiskDeviceMapping.N.Size is the size of the specified snapshot.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The ID of snapshot N to use to create the custom image. .
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public ImageDiskDeviceMappingArgs()
        {
        }
        public static new ImageDiskDeviceMappingArgs Empty => new ImageDiskDeviceMappingArgs();
    }
}
