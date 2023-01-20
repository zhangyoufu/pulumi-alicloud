// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpn.Outputs
{

    [OutputType]
    public sealed class GetCustomerGatewaysGatewayResult
    {
        /// <summary>
        /// The autonomous system number of the local data center gateway device of the VPN customer gateway.
        /// </summary>
        public readonly int Asn;
        /// <summary>
        /// The creation time of the VPN customer gateway.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of the VPN customer gateway.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// ID of the VPN customer gateway .
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ip address of the VPN customer gateway.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The name of the VPN customer gateway.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetCustomerGatewaysGatewayResult(
            int asn,

            string createTime,

            string description,

            string id,

            string ipAddress,

            string name)
        {
            Asn = asn;
            CreateTime = createTime;
            Description = description;
            Id = id;
            IpAddress = ipAddress;
            Name = name;
        }
    }
}
