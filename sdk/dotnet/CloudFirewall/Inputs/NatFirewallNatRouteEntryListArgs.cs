// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CloudFirewall.Inputs
{

    public sealed class NatFirewallNatRouteEntryListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination network segment of the default route.
        /// </summary>
        [Input("destinationCidr", required: true)]
        public Input<string> DestinationCidr { get; set; } = null!;

        /// <summary>
        /// The next hop address of the original NAT gateway.
        /// </summary>
        [Input("nexthopId", required: true)]
        public Input<string> NexthopId { get; set; } = null!;

        /// <summary>
        /// The network type of the next hop. Value: NatGateway : NAT Gateway.
        /// </summary>
        [Input("nexthopType", required: true)]
        public Input<string> NexthopType { get; set; } = null!;

        /// <summary>
        /// The route table where the default route of the NAT gateway is located.
        /// </summary>
        [Input("routeTableId", required: true)]
        public Input<string> RouteTableId { get; set; } = null!;

        public NatFirewallNatRouteEntryListArgs()
        {
        }
        public static new NatFirewallNatRouteEntryListArgs Empty => new NatFirewallNatRouteEntryListArgs();
    }
}
