// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Mse.Inputs
{

    public sealed class GatewaySlbListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The associate id.
        /// </summary>
        [Input("associateId")]
        public Input<string>? AssociateId { get; set; }

        /// <summary>
        /// The Mode of the gateway slb.
        /// </summary>
        [Input("gatewaySlbMode")]
        public Input<string>? GatewaySlbMode { get; set; }

        /// <summary>
        /// The Status of the gateway slb.
        /// </summary>
        [Input("gatewaySlbStatus")]
        public Input<string>? GatewaySlbStatus { get; set; }

        /// <summary>
        /// The creation time of the gateway slb.
        /// </summary>
        [Input("gmtCreate")]
        public Input<string>? GmtCreate { get; set; }

        /// <summary>
        /// The ID of the gateway slb.
        /// </summary>
        [Input("slbId")]
        public Input<string>? SlbId { get; set; }

        /// <summary>
        /// The ip of the gateway slb.
        /// </summary>
        [Input("slbIp")]
        public Input<string>? SlbIp { get; set; }

        /// <summary>
        /// The port of the gateway slb.
        /// </summary>
        [Input("slbPort")]
        public Input<string>? SlbPort { get; set; }

        /// <summary>
        /// The type of the gateway slb.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GatewaySlbListGetArgs()
        {
        }
        public static new GatewaySlbListGetArgs Empty => new GatewaySlbListGetArgs();
    }
}
