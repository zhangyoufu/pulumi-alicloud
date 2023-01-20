// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms.Inputs
{

    public sealed class HybridMonitorSlsTaskSlsProcessConfigFilterFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The method that is used to filter logs imported from Log Service. Valid values: `&gt;`, `&gt;=`, `=`, `&lt;=`, `&lt;`, `!=`, `contain`, `notContain`.
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        /// <summary>
        /// The name of the key that is used to aggregate logs imported from Log Service.
        /// </summary>
        [Input("slsKeyName")]
        public Input<string>? SlsKeyName { get; set; }

        /// <summary>
        /// The value of the key that is used to filter logs imported from Log Service.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public HybridMonitorSlsTaskSlsProcessConfigFilterFilterArgs()
        {
        }
        public static new HybridMonitorSlsTaskSlsProcessConfigFilterFilterArgs Empty => new HybridMonitorSlsTaskSlsProcessConfigFilterFilterArgs();
    }
}
