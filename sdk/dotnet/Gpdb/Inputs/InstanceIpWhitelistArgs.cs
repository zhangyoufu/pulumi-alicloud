// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Gpdb.Inputs
{

    public sealed class InstanceIpWhitelistArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value of this parameter is empty by default. The attribute of the whitelist group. 
        /// If the value contains `hidden`, this white list item will not output.
        /// </summary>
        [Input("ipGroupAttribute")]
        public Input<string>? IpGroupAttribute { get; set; }

        /// <summary>
        /// IP whitelist group name
        /// </summary>
        [Input("ipGroupName")]
        public Input<string>? IpGroupName { get; set; }

        /// <summary>
        /// Field `security_ip_list` has been deprecated from provider version 1.187.0. New field `ip_whitelist` instead.
        /// </summary>
        [Input("securityIpList")]
        public Input<string>? SecurityIpList { get; set; }

        public InstanceIpWhitelistArgs()
        {
        }
        public static new InstanceIpWhitelistArgs Empty => new InstanceIpWhitelistArgs();
    }
}
