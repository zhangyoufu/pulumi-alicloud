// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Inputs
{

    public sealed class RuleRuleActionTrafficLimitConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Number of requests per second. Valid values: `1` to `100000`.
        /// </summary>
        [Input("qps")]
        public Input<int>? Qps { get; set; }

        public RuleRuleActionTrafficLimitConfigArgs()
        {
        }
        public static new RuleRuleActionTrafficLimitConfigArgs Empty => new RuleRuleActionTrafficLimitConfigArgs();
    }
}
