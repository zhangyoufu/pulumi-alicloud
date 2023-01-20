// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms.Inputs
{

    public sealed class HybridMonitorSlsTaskSlsProcessConfigFilterGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("filters")]
        private InputList<Inputs.HybridMonitorSlsTaskSlsProcessConfigFilterFilterGetArgs>? _filters;

        /// <summary>
        /// The conditions that are used to filter logs imported from Log Service. See the following `Block filters`.
        /// </summary>
        public InputList<Inputs.HybridMonitorSlsTaskSlsProcessConfigFilterFilterGetArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.HybridMonitorSlsTaskSlsProcessConfigFilterFilterGetArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The relationship between multiple filter conditions. Valid values: `and`(default value), `or`.
        /// </summary>
        [Input("relation")]
        public Input<string>? Relation { get; set; }

        public HybridMonitorSlsTaskSlsProcessConfigFilterGetArgs()
        {
        }
        public static new HybridMonitorSlsTaskSlsProcessConfigFilterGetArgs Empty => new HybridMonitorSlsTaskSlsProcessConfigFilterGetArgs();
    }
}
