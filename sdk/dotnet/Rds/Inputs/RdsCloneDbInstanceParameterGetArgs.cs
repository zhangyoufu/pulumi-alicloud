// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds.Inputs
{

    public sealed class RdsCloneDbInstanceParameterGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The parameters name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The parameters value.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public RdsCloneDbInstanceParameterGetArgs()
        {
        }
        public static new RdsCloneDbInstanceParameterGetArgs Empty => new RdsCloneDbInstanceParameterGetArgs();
    }
}
