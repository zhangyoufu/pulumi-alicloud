// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Inputs
{

    public sealed class RuleRuleConditionCookieConfigValueGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key of the header field. The key must be 1 to 40 characters in length, and can contain letters, digits, hyphens (-) and underscores (_). The key does not support Cookie or Host.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The value must be 1 to 128 characters in length, and can contain lowercase letters, printable characters, asterisks (*), and question marks (?). The value cannot contain spaces or the following special characters: # [ ] { } \ | &lt; &gt; &amp;.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public RuleRuleConditionCookieConfigValueGetArgs()
        {
        }
        public static new RuleRuleConditionCookieConfigValueGetArgs Empty => new RuleRuleConditionCookieConfigValueGetArgs();
    }
}
