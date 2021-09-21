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
    public sealed class GetNatIpsIpResult
    {
        /// <summary>
        /// The ID of the Nat Ip.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Indicates whether the BGP Group is the default NAT IP ADDRESS. Valid values: `true`: is the default NAT IP ADDRESS. `false`: it is not the default NAT IP ADDRESS.
        /// </summary>
        public readonly bool IsDefault;
        /// <summary>
        /// The ID of the Virtual Private Cloud (VPC) NAT gateway to which the NAT IP address belongs.
        /// </summary>
        public readonly string NatGatewayId;
        /// <summary>
        /// The NAT IP address that is queried.
        /// </summary>
        public readonly string NatIp;
        /// <summary>
        /// The CIDR block to which the NAT IP address belongs.
        /// </summary>
        public readonly string NatIpCidr;
        /// <summary>
        /// The description of the NAT IP address.
        /// </summary>
        public readonly string NatIpDescription;
        /// <summary>
        /// The ID of the NAT IP address.
        /// </summary>
        public readonly string NatIpId;
        /// <summary>
        /// The name of the NAT IP address.
        /// </summary>
        public readonly string NatIpName;
        /// <summary>
        /// The status of the NAT IP address. Valid values: `Available`, `Deleting` and `Creating`.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetNatIpsIpResult(
            string id,

            bool isDefault,

            string natGatewayId,

            string natIp,

            string natIpCidr,

            string natIpDescription,

            string natIpId,

            string natIpName,

            string status)
        {
            Id = id;
            IsDefault = isDefault;
            NatGatewayId = natGatewayId;
            NatIp = natIp;
            NatIpCidr = natIpCidr;
            NatIpDescription = natIpDescription;
            NatIpId = natIpId;
            NatIpName = natIpName;
            Status = status;
        }
    }
}
