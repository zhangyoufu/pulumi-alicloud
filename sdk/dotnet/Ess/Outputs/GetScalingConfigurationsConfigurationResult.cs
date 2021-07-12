// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ess.Outputs
{

    [OutputType]
    public sealed class GetScalingConfigurationsConfigurationResult
    {
        /// <summary>
        /// Creation time of the scaling configuration.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// Performance mode of the t5 burstable instance.
        /// </summary>
        public readonly string CreditSpecification;
        /// <summary>
        /// Data disks of the scaling configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetScalingConfigurationsConfigurationDataDiskResult> DataDisks;
        /// <summary>
        /// ID of the scaling rule.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Image ID of the scaling configuration.
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// Instance type of the scaling configuration.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// Internet charge type of the scaling configuration.
        /// </summary>
        public readonly string InternetChargeType;
        /// <summary>
        /// Internet max bandwidth in of the scaling configuration.
        /// </summary>
        public readonly int InternetMaxBandwidthIn;
        /// <summary>
        /// Internet max bandwidth of the scaling configuration.
        /// </summary>
        public readonly int InternetMaxBandwidthOut;
        /// <summary>
        /// Lifecycle state of the scaling configuration.
        /// </summary>
        public readonly string LifecycleState;
        /// <summary>
        /// Name of the scaling configuration.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Scaling group id the scaling configurations belong to.
        /// </summary>
        public readonly string ScalingGroupId;
        /// <summary>
        /// Security group ID of the scaling configuration.
        /// </summary>
        public readonly string SecurityGroupId;
        /// <summary>
        /// System disk category of the scaling configuration.
        /// </summary>
        public readonly string SystemDiskCategory;
        /// <summary>
        /// The performance level of the ESSD used as the system disk.
        /// </summary>
        public readonly string SystemDiskPerformanceLevel;
        /// <summary>
        /// System disk size of the scaling configuration.
        /// </summary>
        public readonly int SystemDiskSize;

        [OutputConstructor]
        private GetScalingConfigurationsConfigurationResult(
            string creationTime,

            string creditSpecification,

            ImmutableArray<Outputs.GetScalingConfigurationsConfigurationDataDiskResult> dataDisks,

            string id,

            string imageId,

            string instanceType,

            string internetChargeType,

            int internetMaxBandwidthIn,

            int internetMaxBandwidthOut,

            string lifecycleState,

            string name,

            string scalingGroupId,

            string securityGroupId,

            string systemDiskCategory,

            string systemDiskPerformanceLevel,

            int systemDiskSize)
        {
            CreationTime = creationTime;
            CreditSpecification = creditSpecification;
            DataDisks = dataDisks;
            Id = id;
            ImageId = imageId;
            InstanceType = instanceType;
            InternetChargeType = internetChargeType;
            InternetMaxBandwidthIn = internetMaxBandwidthIn;
            InternetMaxBandwidthOut = internetMaxBandwidthOut;
            LifecycleState = lifecycleState;
            Name = name;
            ScalingGroupId = scalingGroupId;
            SecurityGroupId = securityGroupId;
            SystemDiskCategory = systemDiskCategory;
            SystemDiskPerformanceLevel = systemDiskPerformanceLevel;
            SystemDiskSize = systemDiskSize;
        }
    }
}
