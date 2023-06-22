// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc.Inputs
{

    public sealed class DhcpOptionsSetAssociateVpcGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The status of the VPC associated with the DHCP option set.
        /// </summary>
        [Input("associateStatus")]
        public Input<string>? AssociateStatus { get; set; }

        /// <summary>
        /// The ID of the VPC network that is associated with the DHCP options set.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public DhcpOptionsSetAssociateVpcGetArgs()
        {
        }
        public static new DhcpOptionsSetAssociateVpcGetArgs Empty => new DhcpOptionsSetAssociateVpcGetArgs();
    }
}
