// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Inputs
{

    public sealed class RuleRuleActionTrafficMirrorConfigMirrorGroupConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("serverGroupTuples")]
        private InputList<Inputs.RuleRuleActionTrafficMirrorConfigMirrorGroupConfigServerGroupTupleArgs>? _serverGroupTuples;

        /// <summary>
        /// The destination server group to which requests are forwarded. See `server_group_tuples` below.
        /// </summary>
        public InputList<Inputs.RuleRuleActionTrafficMirrorConfigMirrorGroupConfigServerGroupTupleArgs> ServerGroupTuples
        {
            get => _serverGroupTuples ?? (_serverGroupTuples = new InputList<Inputs.RuleRuleActionTrafficMirrorConfigMirrorGroupConfigServerGroupTupleArgs>());
            set => _serverGroupTuples = value;
        }

        public RuleRuleActionTrafficMirrorConfigMirrorGroupConfigArgs()
        {
        }
        public static new RuleRuleActionTrafficMirrorConfigMirrorGroupConfigArgs Empty => new RuleRuleActionTrafficMirrorConfigMirrorGroupConfigArgs();
    }
}
