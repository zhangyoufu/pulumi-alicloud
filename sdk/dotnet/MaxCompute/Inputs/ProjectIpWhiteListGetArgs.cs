// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.MaxCompute.Inputs
{

    public sealed class ProjectIpWhiteListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Classic network IP white list.
        /// </summary>
        [Input("ipList")]
        public Input<string>? IpList { get; set; }

        /// <summary>
        /// VPC network whitelist.
        /// </summary>
        [Input("vpcIpList")]
        public Input<string>? VpcIpList { get; set; }

        public ProjectIpWhiteListGetArgs()
        {
        }
        public static new ProjectIpWhiteListGetArgs Empty => new ProjectIpWhiteListGetArgs();
    }
}
