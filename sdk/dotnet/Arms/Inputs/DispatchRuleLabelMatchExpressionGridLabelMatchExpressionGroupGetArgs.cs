// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Arms.Inputs
{

    public sealed class DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupGetArgs : Pulumi.ResourceArgs
    {
        [Input("labelMatchExpressions", required: true)]
        private InputList<Inputs.DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpressionGetArgs>? _labelMatchExpressions;

        /// <summary>
        /// Sets the dispatch rule. See the following `Block label_match_expressions`.
        /// </summary>
        public InputList<Inputs.DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpressionGetArgs> LabelMatchExpressions
        {
            get => _labelMatchExpressions ?? (_labelMatchExpressions = new InputList<Inputs.DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpressionGetArgs>());
            set => _labelMatchExpressions = value;
        }

        public DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupGetArgs()
        {
        }
    }
}
