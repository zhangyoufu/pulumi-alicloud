// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga.Inputs
{

    public sealed class ForwardingRuleRuleActionForwardGroupConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("serverGroupTuples", required: true)]
        private InputList<Inputs.ForwardingRuleRuleActionForwardGroupConfigServerGroupTupleArgs>? _serverGroupTuples;

        /// <summary>
        /// Terminal node group configuration.
        /// </summary>
        public InputList<Inputs.ForwardingRuleRuleActionForwardGroupConfigServerGroupTupleArgs> ServerGroupTuples
        {
            get => _serverGroupTuples ?? (_serverGroupTuples = new InputList<Inputs.ForwardingRuleRuleActionForwardGroupConfigServerGroupTupleArgs>());
            set => _serverGroupTuples = value;
        }

        public ForwardingRuleRuleActionForwardGroupConfigArgs()
        {
        }
        public static new ForwardingRuleRuleActionForwardGroupConfigArgs Empty => new ForwardingRuleRuleActionForwardGroupConfigArgs();
    }
}
