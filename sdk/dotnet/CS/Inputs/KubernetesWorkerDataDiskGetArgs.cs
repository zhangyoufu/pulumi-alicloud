// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class KubernetesWorkerDataDiskGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional, Available in 1.120.0+) Worker node data disk auto snapshot policy.
        /// </summary>
        [Input("autoSnapshotPolicyId")]
        public Input<string>? AutoSnapshotPolicyId { get; set; }

        /// <summary>
        /// The type of the data disks. Valid values: `cloud`, `cloud_efficiency`, `cloud_ssd` and `cloud_essd`. Default to `cloud_efficiency`.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Specifies whether to encrypt data disks. Valid values: true and false.
        /// </summary>
        [Input("encrypted")]
        public Input<string>? Encrypted { get; set; }

        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// The kubernetes cluster's name. It is unique in one Alicloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Optional, Available in 1.120.0+) Worker node data disk performance level, when `category` values `cloud_essd`, the optional values are `PL0`, `PL1`, `PL2` or `PL3`, but the specific performance level is related to the disk capacity. For more information, see [Enhanced SSDs](https://www.alibabacloud.com/help/doc-detail/122389.htm). Default is `PL1`.
        /// </summary>
        [Input("performanceLevel")]
        public Input<string>? PerformanceLevel { get; set; }

        /// <summary>
        /// The size of a data disk, Its valid value range [40~32768] in GB. Unit: GiB.
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public KubernetesWorkerDataDiskGetArgs()
        {
        }
        public static new KubernetesWorkerDataDiskGetArgs Empty => new KubernetesWorkerDataDiskGetArgs();
    }
}
