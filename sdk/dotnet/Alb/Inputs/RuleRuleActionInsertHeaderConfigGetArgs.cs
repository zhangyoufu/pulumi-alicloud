// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Inputs
{

    public sealed class RuleRuleActionInsertHeaderConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the inserted header field. The name must be 1 to 40 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). You cannot use the same name in InsertHeader. Note You cannot use Cookie or Host in the name.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The content of the inserted header field. Valid values:
        /// * If the `value_type` is set to `SystemDefined`, the following values are used:
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// The value type of the inserted header field. Valid values:
        /// </summary>
        [Input("valueType")]
        public Input<string>? ValueType { get; set; }

        public RuleRuleActionInsertHeaderConfigGetArgs()
        {
        }
        public static new RuleRuleActionInsertHeaderConfigGetArgs Empty => new RuleRuleActionInsertHeaderConfigGetArgs();
    }
}
