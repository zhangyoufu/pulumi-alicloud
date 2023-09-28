// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbase.Inputs
{

    public sealed class InstanceSlbConnAddrArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Phoenix address.
        /// </summary>
        [Input("connAddr")]
        public Input<string>? ConnAddr { get; set; }

        /// <summary>
        /// The number of the port over which Phoenix connects to the instance.
        /// </summary>
        [Input("connAddrPort")]
        public Input<string>? ConnAddrPort { get; set; }

        /// <summary>
        /// The type of the network. Valid values:
        /// </summary>
        [Input("netType")]
        public Input<string>? NetType { get; set; }

        public InstanceSlbConnAddrArgs()
        {
        }
        public static new InstanceSlbConnAddrArgs Empty => new InstanceSlbConnAddrArgs();
    }
}
