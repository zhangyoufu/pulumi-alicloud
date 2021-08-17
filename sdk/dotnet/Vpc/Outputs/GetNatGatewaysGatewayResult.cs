// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc.Outputs
{

    [OutputType]
    public sealed class GetNatGatewaysGatewayResult
    {
        /// <summary>
        /// The state of the NAT gateway.
        /// </summary>
        public readonly string BusinessStatus;
        /// <summary>
        /// Indicates whether deletion protection is enabled.
        /// </summary>
        public readonly bool DeletionProtection;
        /// <summary>
        /// The description of the NAT gateway.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Indicates whether the traffic monitoring feature is enabled.
        /// </summary>
        public readonly bool EcsMetricEnabled;
        /// <summary>
        /// The time when the NAT gateway expires.
        /// </summary>
        public readonly string ExpiredTime;
        /// <summary>
        /// The ID of the DNAT table.
        /// </summary>
        public readonly ImmutableArray<string> ForwardTableIds;
        /// <summary>
        /// The ID of the NAT gateway.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The metering method of the NAT gateway.
        /// </summary>
        public readonly string InternetChargeType;
        /// <summary>
        /// The ip address of the bind eip.
        /// </summary>
        public readonly ImmutableArray<string> IpLists;
        /// <summary>
        /// Name of the NAT gateway.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of the NAT gateway.
        /// </summary>
        public readonly string NatGatewayId;
        /// <summary>
        /// The name of NAT gateway.
        /// </summary>
        public readonly string NatGatewayName;
        /// <summary>
        /// The nat type of NAT gateway. Valid values `Enhanced` and `Normal`.
        /// </summary>
        public readonly string NatType;
        /// <summary>
        /// The payment type of NAT gateway. Valid values `PayAsYouGo` and `Subscription`.
        /// </summary>
        public readonly string PaymentType;
        /// <summary>
        /// The resource group id of NAT gateway.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// The ID of the SNAT table that is associated with the NAT gateway.
        /// </summary>
        public readonly ImmutableArray<string> SnatTableIds;
        /// <summary>
        /// The specification of the NAT gateway.
        /// </summary>
        public readonly string Spec;
        /// <summary>
        /// The specification of NAT gateway. Valid values `Middle`, `Large`, `Small` and `XLarge.1`. Default value is `Small`.
        /// </summary>
        public readonly string Specification;
        /// <summary>
        /// The status of NAT gateway. Valid values `Available`, `Converting`, `Creating`, `Deleting` and `Modifying`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The tags of NAT gateway.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// The ID of the VPC.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The ID of the vSwitch to which the NAT gateway belongs.
        /// </summary>
        public readonly string VswitchId;

        [OutputConstructor]
        private GetNatGatewaysGatewayResult(
            string businessStatus,

            bool deletionProtection,

            string description,

            bool ecsMetricEnabled,

            string expiredTime,

            ImmutableArray<string> forwardTableIds,

            string id,

            string internetChargeType,

            ImmutableArray<string> ipLists,

            string name,

            string natGatewayId,

            string natGatewayName,

            string natType,

            string paymentType,

            string resourceGroupId,

            ImmutableArray<string> snatTableIds,

            string spec,

            string specification,

            string status,

            ImmutableDictionary<string, object> tags,

            string vpcId,

            string vswitchId)
        {
            BusinessStatus = businessStatus;
            DeletionProtection = deletionProtection;
            Description = description;
            EcsMetricEnabled = ecsMetricEnabled;
            ExpiredTime = expiredTime;
            ForwardTableIds = forwardTableIds;
            Id = id;
            InternetChargeType = internetChargeType;
            IpLists = ipLists;
            Name = name;
            NatGatewayId = natGatewayId;
            NatGatewayName = natGatewayName;
            NatType = natType;
            PaymentType = paymentType;
            ResourceGroupId = resourceGroupId;
            SnatTableIds = snatTableIds;
            Spec = spec;
            Specification = specification;
            Status = status;
            Tags = tags;
            VpcId = vpcId;
            VswitchId = vswitchId;
        }
    }
}
