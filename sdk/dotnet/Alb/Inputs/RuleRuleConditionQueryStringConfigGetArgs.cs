// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Inputs
{

    public sealed class RuleRuleConditionQueryStringConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("values")]
        private InputList<Inputs.RuleRuleConditionQueryStringConfigValueGetArgs>? _values;

        /// <summary>
        /// The query string.
        /// </summary>
        public InputList<Inputs.RuleRuleConditionQueryStringConfigValueGetArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.RuleRuleConditionQueryStringConfigValueGetArgs>());
            set => _values = value;
        }

        public RuleRuleConditionQueryStringConfigGetArgs()
        {
        }
    }
}
