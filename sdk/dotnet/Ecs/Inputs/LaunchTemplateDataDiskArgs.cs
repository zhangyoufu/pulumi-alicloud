// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs.Inputs
{

    public sealed class LaunchTemplateDataDiskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The category of the disk:
        /// - cloud: Basic cloud disk.
        /// - cloud_efficiency: Ultra cloud disk.
        /// - cloud_ssd: SSD cloud Disks.
        /// - ephemeral_ssd: local SSD Disks
        /// - cloud_essd: ESSD cloud Disks.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Delete this data disk when the instance is destroyed. It only works on cloud, cloud_efficiency, cloud_ssd and cloud_essd disk. If the category of this data disk was ephemeral_ssd, please don't set this param.
        /// </summary>
        [Input("deleteWithInstance")]
        public Input<bool>? DeleteWithInstance { get; set; }

        /// <summary>
        /// The description of the data disk.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// -(Optional, Bool) Encrypted the data in this disk.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// The name of the data disk.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The size of the data disk.
        /// - cloud：[5, 2000]
        /// - cloud_efficiency：[20, 32768]
        /// - cloud_ssd：[20, 32768]
        /// - cloud_essd：[20, 32768]
        /// - ephemeral_ssd: [5, 800]
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The snapshot ID used to initialize the data disk. If the size specified by snapshot is greater that the size of the disk, use the size specified by snapshot as the size of the data disk.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public LaunchTemplateDataDiskArgs()
        {
        }
    }
}
