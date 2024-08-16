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
    public sealed class GetVirtualNodesNodeResult
    {
        /// <summary>
        /// The Number of CPU.
        /// </summary>
        public readonly int Cpu;
        /// <summary>
        /// The creation time of the virtual node.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The ENI instance ID.
        /// </summary>
        public readonly string EniInstanceId;
        /// <summary>
        /// The event list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualNodesNodeEventResult> Events;
        /// <summary>
        /// The ID of the Virtual Node.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The IP address of a public network.
        /// </summary>
        public readonly string InternetIp;
        /// <summary>
        /// The private IP address of the RDS instance.
        /// </summary>
        public readonly string IntranetIp;
        /// <summary>
        /// The memory size.
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// The ram role.
        /// </summary>
        public readonly string RamRoleName;
        /// <summary>
        /// The resource group ID.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// The security group ID.
        /// </summary>
        public readonly string SecurityGroupId;
        /// <summary>
        /// The Status of the virtual node.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Of the virtual node number.
        /// </summary>
        public readonly string VirtualNodeId;
        /// <summary>
        /// The name of the virtual node.
        /// </summary>
        public readonly string VirtualNodeName;
        public readonly string VpcId;
        /// <summary>
        /// The vswitch id.
        /// </summary>
        public readonly string VswitchId;
        /// <summary>
        /// The Zone.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private GetVirtualNodesNodeResult(
            int cpu,

            string createTime,

            string eniInstanceId,

            ImmutableArray<Outputs.GetVirtualNodesNodeEventResult> events,

            string id,

            string internetIp,

            string intranetIp,

            int memory,

            string ramRoleName,

            string resourceGroupId,

            string securityGroupId,

            string status,

            ImmutableDictionary<string, string> tags,

            string virtualNodeId,

            string virtualNodeName,

            string vpcId,

            string vswitchId,

            string zoneId)
        {
            Cpu = cpu;
            CreateTime = createTime;
            EniInstanceId = eniInstanceId;
            Events = events;
            Id = id;
            InternetIp = internetIp;
            IntranetIp = intranetIp;
            Memory = memory;
            RamRoleName = ramRoleName;
            ResourceGroupId = resourceGroupId;
            SecurityGroupId = securityGroupId;
            Status = status;
            Tags = tags;
            VirtualNodeId = virtualNodeId;
            VirtualNodeName = virtualNodeName;
            VpcId = vpcId;
            VswitchId = vswitchId;
            ZoneId = zoneId;
        }
    }
}
