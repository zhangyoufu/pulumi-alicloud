// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.PolarDB.Inputs
{

    public sealed class ParameterGroupParameterGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of a parameter in the parameter template.
        /// </summary>
        [Input("paramName", required: true)]
        public Input<string> ParamName { get; set; } = null!;

        /// <summary>
        /// The value of a parameter in the parameter template.
        /// </summary>
        [Input("paramValue", required: true)]
        public Input<string> ParamValue { get; set; } = null!;

        public ParameterGroupParameterGetArgs()
        {
        }
        public static new ParameterGroupParameterGetArgs Empty => new ParameterGroupParameterGetArgs();
    }
}
