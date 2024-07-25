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
    public sealed class NetworkIpv6CidrBlock
    {
        /// <summary>
        /// The IPv6 CIDR block of the default VPC.
        /// 
        /// &gt; **NOTE:**  When `EnableIpv6` is set to `true`, this parameter is required.
        /// </summary>
        public readonly string? Ipv6CidrBlock;
        /// <summary>
        /// The IPv6 address segment type of the VPC. Value:
        /// - `BGP` (default): Alibaba Cloud BGP IPv6.
        /// - `ChinaMobile`: China Mobile (single line).
        /// - `ChinaUnicom`: China Unicom (single line).
        /// - `ChinaTelecom`: China Telecom (single line).
        /// 
        /// &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to `ChinaTelecom` (China Telecom), `ChinaUnicom` (China Unicom), or `ChinaMobile` (China Mobile).
        /// </summary>
        public readonly string? Ipv6Isp;

        [OutputConstructor]
        private NetworkIpv6CidrBlock(
            string? ipv6CidrBlock,

            string? ipv6Isp)
        {
            Ipv6CidrBlock = ipv6CidrBlock;
            Ipv6Isp = ipv6Isp;
        }
    }
}
