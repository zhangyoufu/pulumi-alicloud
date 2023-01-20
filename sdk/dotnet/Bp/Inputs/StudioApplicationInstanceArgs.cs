// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Bp.Inputs
{

    public sealed class StudioApplicationInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the instance.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the instance.
        /// </summary>
        [Input("nodeName")]
        public Input<string>? NodeName { get; set; }

        /// <summary>
        /// The type of the instance.
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        public StudioApplicationInstanceArgs()
        {
        }
        public static new StudioApplicationInstanceArgs Empty => new StudioApplicationInstanceArgs();
    }
}
