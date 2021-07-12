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
    public sealed class GetNetworkInterfacesInterfaceResult
    {
        /// <summary>
        /// Creation time of the ENI.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// Description of the ENI.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// ID of the ENI.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID of the instance that the ENI is attached to.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// MAC address of the ENI.
        /// </summary>
        public readonly string Mac;
        /// <summary>
        /// Name of the ENI.
        /// </summary>
        public readonly string Name;
        public readonly string NetworkInterfaceId;
        public readonly string NetworkInterfaceName;
        public readonly string PrimaryIpAddress;
        /// <summary>
        /// Primary private IP of the ENI.
        /// </summary>
        public readonly string PrivateIp;
        public readonly ImmutableArray<string> PrivateIpAddresses;
        /// <summary>
        /// A list of secondary private IP address that is assigned to the ENI.
        /// </summary>
        public readonly ImmutableArray<string> PrivateIps;
        public readonly int QueueNumber;
        /// <summary>
        /// The Id of resource group.
        /// </summary>
        public readonly string ResourceGroupId;
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// A list of security group that the ENI belongs to.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;
        public readonly int ServiceId;
        public readonly bool ServiceManaged;
        /// <summary>
        /// Current status of the ENI.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A map of tags assigned to the ENI.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        public readonly string Type;
        /// <summary>
        /// ID of the VPC that the ENI belongs to.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// ID of the VSwitch that the ENI is linked to.
        /// </summary>
        public readonly string VswitchId;
        /// <summary>
        /// ID of the availability zone that the ENI belongs to.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private GetNetworkInterfacesInterfaceResult(
            string creationTime,

            string description,

            string id,

            string instanceId,

            string mac,

            string name,

            string networkInterfaceId,

            string networkInterfaceName,

            string primaryIpAddress,

            string privateIp,

            ImmutableArray<string> privateIpAddresses,

            ImmutableArray<string> privateIps,

            int queueNumber,

            string resourceGroupId,

            ImmutableArray<string> securityGroupIds,

            ImmutableArray<string> securityGroups,

            int serviceId,

            bool serviceManaged,

            string status,

            ImmutableDictionary<string, object> tags,

            string type,

            string vpcId,

            string vswitchId,

            string zoneId)
        {
            CreationTime = creationTime;
            Description = description;
            Id = id;
            InstanceId = instanceId;
            Mac = mac;
            Name = name;
            NetworkInterfaceId = networkInterfaceId;
            NetworkInterfaceName = networkInterfaceName;
            PrimaryIpAddress = primaryIpAddress;
            PrivateIp = privateIp;
            PrivateIpAddresses = privateIpAddresses;
            PrivateIps = privateIps;
            QueueNumber = queueNumber;
            ResourceGroupId = resourceGroupId;
            SecurityGroupIds = securityGroupIds;
            SecurityGroups = securityGroups;
            ServiceId = serviceId;
            ServiceManaged = serviceManaged;
            Status = status;
            Tags = tags;
            Type = type;
            VpcId = vpcId;
            VswitchId = vswitchId;
            ZoneId = zoneId;
        }
    }
}
