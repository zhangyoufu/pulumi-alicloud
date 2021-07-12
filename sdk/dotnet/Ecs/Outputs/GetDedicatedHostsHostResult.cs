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
    public sealed class GetDedicatedHostsHostResult
    {
        /// <summary>
        /// The policy used to migrate the instances from the dedicated host when the dedicated host fails or needs to be repaired online.
        /// </summary>
        public readonly string ActionOnMaintenance;
        /// <summary>
        /// Specifies whether to add the dedicated host to the resource pool for automatic deployment.
        /// </summary>
        public readonly string AutoPlacement;
        /// <summary>
        /// The automatic release time of the dedicated host.
        /// </summary>
        public readonly string AutoReleaseTime;
        /// <summary>
        /// (Available in 1.123.1+) A collection of proprietary host performance indicators.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDedicatedHostsHostCapacityResult> Capacities;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly int Cores;
        /// <summary>
        /// (Available in 1.123.1+) CPU oversold ratio.
        /// </summary>
        public readonly double CpuOverCommitRatio;
        /// <summary>
        /// The ID of ECS Dedicated Host.
        /// </summary>
        public readonly string DedicatedHostId;
        /// <summary>
        /// The name of ECS Dedicated Host.
        /// </summary>
        public readonly string DedicatedHostName;
        /// <summary>
        /// The type of the dedicated host.
        /// </summary>
        public readonly string DedicatedHostType;
        /// <summary>
        /// The description of the dedicated host.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The expiration time of the subscription dedicated host.
        /// </summary>
        public readonly string ExpiredTime;
        /// <summary>
        /// The GPU model.
        /// </summary>
        public readonly string GpuSpec;
        /// <summary>
        /// ID of the ECS Dedicated Host.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The machine code of the dedicated host.
        /// </summary>
        public readonly string MachineId;
        /// <summary>
        /// dedicated host network parameters. contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDedicatedHostsHostNetworkAttributeResult> NetworkAttributes;
        /// <summary>
        /// The reason why the dedicated host resource is locked.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDedicatedHostsHostOperationLockResult> OperationLocks;
        /// <summary>
        /// The billing method of the dedicated host.
        /// </summary>
        public readonly string PaymentType;
        /// <summary>
        /// The number of physical GPUs.
        /// </summary>
        public readonly int PhysicalGpus;
        /// <summary>
        /// The ID of the resource group to which the ECS Dedicated Host belongs.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// The unit of the subscription billing method.
        /// </summary>
        public readonly string SaleCycle;
        /// <summary>
        /// The number of physical CPUs.
        /// </summary>
        public readonly int Sockets;
        /// <summary>
        /// The status of the ECS Dedicated Host. validate value: `Available`, `Creating`, `PermanentFailure`, `Released`, `UnderAssessment`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// (Available in 1.123.1+) A custom instance type family supported by a dedicated host.
        /// </summary>
        public readonly ImmutableArray<string> SupportedCustomInstanceTypeFamilies;
        /// <summary>
        /// (Available in 1.123.1+) ECS instance type family supported by the dedicated host.
        /// </summary>
        public readonly ImmutableArray<string> SupportedInstanceTypeFamilies;
        /// <summary>
        /// The list of ECS instance
        /// </summary>
        public readonly ImmutableArray<string> SupportedInstanceTypesLists;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// The zone ID of the ECS Dedicated Host.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private GetDedicatedHostsHostResult(
            string actionOnMaintenance,

            string autoPlacement,

            string autoReleaseTime,

            ImmutableArray<Outputs.GetDedicatedHostsHostCapacityResult> capacities,

            int cores,

            double cpuOverCommitRatio,

            string dedicatedHostId,

            string dedicatedHostName,

            string dedicatedHostType,

            string description,

            string expiredTime,

            string gpuSpec,

            string id,

            string machineId,

            ImmutableArray<Outputs.GetDedicatedHostsHostNetworkAttributeResult> networkAttributes,

            ImmutableArray<Outputs.GetDedicatedHostsHostOperationLockResult> operationLocks,

            string paymentType,

            int physicalGpus,

            string resourceGroupId,

            string saleCycle,

            int sockets,

            string status,

            ImmutableArray<string> supportedCustomInstanceTypeFamilies,

            ImmutableArray<string> supportedInstanceTypeFamilies,

            ImmutableArray<string> supportedInstanceTypesLists,

            ImmutableDictionary<string, object> tags,

            string zoneId)
        {
            ActionOnMaintenance = actionOnMaintenance;
            AutoPlacement = autoPlacement;
            AutoReleaseTime = autoReleaseTime;
            Capacities = capacities;
            Cores = cores;
            CpuOverCommitRatio = cpuOverCommitRatio;
            DedicatedHostId = dedicatedHostId;
            DedicatedHostName = dedicatedHostName;
            DedicatedHostType = dedicatedHostType;
            Description = description;
            ExpiredTime = expiredTime;
            GpuSpec = gpuSpec;
            Id = id;
            MachineId = machineId;
            NetworkAttributes = networkAttributes;
            OperationLocks = operationLocks;
            PaymentType = paymentType;
            PhysicalGpus = physicalGpus;
            ResourceGroupId = resourceGroupId;
            SaleCycle = saleCycle;
            Sockets = sockets;
            Status = status;
            SupportedCustomInstanceTypeFamilies = supportedCustomInstanceTypeFamilies;
            SupportedInstanceTypeFamilies = supportedInstanceTypeFamilies;
            SupportedInstanceTypesLists = supportedInstanceTypesLists;
            Tags = tags;
            ZoneId = zoneId;
        }
    }
}
