// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class KubernetesMasterNodeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the node.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The kubernetes cluster's name. It is unique in one Alicloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The private IP address of node.
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        public KubernetesMasterNodeArgs()
        {
        }
        public static new KubernetesMasterNodeArgs Empty => new KubernetesMasterNodeArgs();
    }
}
