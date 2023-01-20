// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds.Outputs
{

    [OutputType]
    public sealed class GetInstanceClassInfosInfoResult
    {
        /// <summary>
        /// The code of the instance type.
        /// </summary>
        public readonly string? ClassCode;
        /// <summary>
        /// The instance family of the instance.
        /// </summary>
        public readonly string? ClassGroup;
        /// <summary>
        /// The number of cores that are supported by the instance type. Unit: cores.
        /// </summary>
        public readonly string? Cpu;
        /// <summary>
        /// The architecture of the instance type.
        /// </summary>
        public readonly string? InstructionSetArch;
        /// <summary>
        /// The maximum number of connections that are supported by the instance type. Unit: connections.
        /// </summary>
        public readonly string? MaxConnections;
        /// <summary>
        /// The maximum I/O bandwidth that is supported by the instance type. Unit: Mbit/s.
        /// </summary>
        public readonly string? MaxIombps;
        /// <summary>
        /// The maximum input/output operations per second (IOPS) that is supported by the instance type. Unit: operations per second.
        /// </summary>
        public readonly string? MaxIops;
        /// <summary>
        /// The memory capacity that is supported by the instance type. Unit: GB.
        /// </summary>
        public readonly string? MemoryClass;
        /// <summary>
        /// The fee that you must pay for the instance type. Unit: cent (USD).
        /// </summary>
        public readonly string? ReferencePrice;

        [OutputConstructor]
        private GetInstanceClassInfosInfoResult(
            string? classCode,

            string? classGroup,

            string? cpu,

            string? instructionSetArch,

            string? maxConnections,

            string? maxIombps,

            string? maxIops,

            string? memoryClass,

            string? referencePrice)
        {
            ClassCode = classCode;
            ClassGroup = classGroup;
            Cpu = cpu;
            InstructionSetArch = instructionSetArch;
            MaxConnections = maxConnections;
            MaxIombps = maxIombps;
            MaxIops = maxIops;
            MemoryClass = memoryClass;
            ReferencePrice = referencePrice;
        }
    }
}
