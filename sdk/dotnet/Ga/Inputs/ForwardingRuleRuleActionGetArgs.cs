// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga.Inputs
{

    public sealed class ForwardingRuleRuleActionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Forwarding configuration.
        /// </summary>
        [Input("forwardGroupConfig", required: true)]
        public Input<Inputs.ForwardingRuleRuleActionForwardGroupConfigGetArgs> ForwardGroupConfig { get; set; } = null!;

        /// <summary>
        /// Forwarding priority.
        /// </summary>
        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        /// <summary>
        /// Forward action type. Default: forwardgroup.
        /// </summary>
        [Input("ruleActionType", required: true)]
        public Input<string> RuleActionType { get; set; } = null!;

        public ForwardingRuleRuleActionGetArgs()
        {
        }
        public static new ForwardingRuleRuleActionGetArgs Empty => new ForwardingRuleRuleActionGetArgs();
    }
}
