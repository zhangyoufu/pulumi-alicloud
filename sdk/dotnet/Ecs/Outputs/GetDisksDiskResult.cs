// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs.Outputs
{

    [OutputType]
    public sealed class GetDisksDiskResult
    {
        /// <summary>
        /// Disk attachment time.
        /// </summary>
        public readonly string AttachedTime;
        public readonly string AutoSnapshotPolicyId;
        /// <summary>
        /// Availability zone of the disk.
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// Disk category. Possible values: `cloud` (basic cloud disk), `cloud_efficiency` (ultra cloud disk), `ephemeral_ssd` (local SSD cloud disk), `cloud_ssd` (SSD cloud disk), and `cloud_essd` (ESSD cloud disk), `cloud_essd_entry`.
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// Disk creation time.
        /// </summary>
        public readonly string CreationTime;
        public readonly bool DeleteAutoSnapshot;
        public readonly bool DeleteWithInstance;
        /// <summary>
        /// Disk description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Disk detachment time.
        /// </summary>
        public readonly string DetachedTime;
        public readonly string Device;
        public readonly string DiskId;
        public readonly string DiskName;
        public readonly string DiskType;
        public readonly bool EnableAutoSnapshot;
        public readonly bool EnableAutomatedSnapshotPolicy;
        /// <summary>
        /// Indicate whether the disk is encrypted or not. Possible values: `on` and `off`.
        /// </summary>
        public readonly string Encrypted;
        public readonly string ExpiredTime;
        /// <summary>
        /// ID of the disk.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID of the image from which the disk is created. It is null unless the disk is created using an image.
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// Filter the results by the specified ECS instance ID.
        /// </summary>
        public readonly string InstanceId;
        public readonly int Iops;
        public readonly int IopsRead;
        public readonly int IopsWrite;
        public readonly string KmsKeyId;
        public readonly int MountInstanceNum;
        public readonly ImmutableArray<Outputs.GetDisksDiskMountInstanceResult> MountInstances;
        /// <summary>
        /// Disk name.
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetDisksDiskOperationLockResult> OperationLocks;
        public readonly string PaymentType;
        public readonly string PerformanceLevel;
        public readonly bool Portable;
        public readonly string ProductCode;
        /// <summary>
        /// Region ID the disk belongs to.
        /// </summary>
        public readonly string RegionId;
        /// <summary>
        /// The Id of resource group which the disk belongs.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// Disk size in GiB.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// Snapshot used to create the disk. It is null if no snapshot is used to create the disk.
        /// </summary>
        public readonly string SnapshotId;
        /// <summary>
        /// Current status. Possible values: `In_use`, `Available`, `Attaching`, `Detaching`, `Creating` and `ReIniting`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A map of tags assigned to the disks. It must be in the format:
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var disksDs = AliCloud.Ecs.GetDisks.Invoke(new()
        ///     {
        ///         Tags = 
        ///         {
        ///             { "tagKey1", "tagValue1" },
        ///             { "tagKey2", "tagValue2" },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// Disk type. Possible values: `system` and `data`.
        /// </summary>
        public readonly string Type;
        public readonly string ZoneId;

        [OutputConstructor]
        private GetDisksDiskResult(
            string attachedTime,

            string autoSnapshotPolicyId,

            string availabilityZone,

            string category,

            string creationTime,

            bool deleteAutoSnapshot,

            bool deleteWithInstance,

            string description,

            string detachedTime,

            string device,

            string diskId,

            string diskName,

            string diskType,

            bool enableAutoSnapshot,

            bool enableAutomatedSnapshotPolicy,

            string encrypted,

            string expiredTime,

            string id,

            string imageId,

            string instanceId,

            int iops,

            int iopsRead,

            int iopsWrite,

            string kmsKeyId,

            int mountInstanceNum,

            ImmutableArray<Outputs.GetDisksDiskMountInstanceResult> mountInstances,

            string name,

            ImmutableArray<Outputs.GetDisksDiskOperationLockResult> operationLocks,

            string paymentType,

            string performanceLevel,

            bool portable,

            string productCode,

            string regionId,

            string resourceGroupId,

            int size,

            string snapshotId,

            string status,

            ImmutableDictionary<string, object> tags,

            string type,

            string zoneId)
        {
            AttachedTime = attachedTime;
            AutoSnapshotPolicyId = autoSnapshotPolicyId;
            AvailabilityZone = availabilityZone;
            Category = category;
            CreationTime = creationTime;
            DeleteAutoSnapshot = deleteAutoSnapshot;
            DeleteWithInstance = deleteWithInstance;
            Description = description;
            DetachedTime = detachedTime;
            Device = device;
            DiskId = diskId;
            DiskName = diskName;
            DiskType = diskType;
            EnableAutoSnapshot = enableAutoSnapshot;
            EnableAutomatedSnapshotPolicy = enableAutomatedSnapshotPolicy;
            Encrypted = encrypted;
            ExpiredTime = expiredTime;
            Id = id;
            ImageId = imageId;
            InstanceId = instanceId;
            Iops = iops;
            IopsRead = iopsRead;
            IopsWrite = iopsWrite;
            KmsKeyId = kmsKeyId;
            MountInstanceNum = mountInstanceNum;
            MountInstances = mountInstances;
            Name = name;
            OperationLocks = operationLocks;
            PaymentType = paymentType;
            PerformanceLevel = performanceLevel;
            Portable = portable;
            ProductCode = productCode;
            RegionId = regionId;
            ResourceGroupId = resourceGroupId;
            Size = size;
            SnapshotId = snapshotId;
            Status = status;
            Tags = tags;
            Type = type;
            ZoneId = zoneId;
        }
    }
}
