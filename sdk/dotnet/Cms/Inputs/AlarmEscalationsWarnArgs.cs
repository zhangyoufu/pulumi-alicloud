// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms.Inputs
{

    public sealed class AlarmEscalationsWarnArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Critical level alarm comparison operator. Valid values: ["&lt;=", "&lt;", "&gt;", "&gt;=", "==", "!="]. Default to "==".
        /// </summary>
        [Input("comparisonOperator")]
        public Input<string>? ComparisonOperator { get; set; }

        /// <summary>
        /// Critical level alarm statistics method. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum", "Value", "ErrorCodeMaximum", "Sum", "Count"]. Default to "Average".
        /// </summary>
        [Input("statistics")]
        public Input<string>? Statistics { get; set; }

        /// <summary>
        /// Critical level alarm threshold value, which must be a numeric value currently.
        /// </summary>
        [Input("threshold")]
        public Input<string>? Threshold { get; set; }

        /// <summary>
        /// Critical level alarm retry times. Default to 3.
        /// </summary>
        [Input("times")]
        public Input<int>? Times { get; set; }

        public AlarmEscalationsWarnArgs()
        {
        }
    }
}
