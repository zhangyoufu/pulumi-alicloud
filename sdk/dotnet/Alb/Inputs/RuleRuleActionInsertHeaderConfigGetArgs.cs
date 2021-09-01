// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Inputs
{

    public sealed class RuleRuleActionInsertHeaderConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the inserted header field. The name must be 1 to 40 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). You cannot use the same name in InsertHeader.  Note You cannot use Cookie or Host in the name.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The content of the inserted header field:  If the ValueType parameter is set to SystemDefined, the following values are used:  ClientSrcPort: the port of the client ClientSrcIp: the IP address of the client Protocol: the protocol used by client requests (HTTP or HTTPS) SLBId: the ID of the ALB instance SLBPort: the listener port of the ALB instance If the ValueType parameter is set to UserDefined: The header value must be 1 to 128 characters in length, and can contain lowercase letters, printable characters whose ASCII value is ch &gt;= 32 &amp;&amp; ch &lt; 127, and wildcards such as asterisks (*) and question marks (?). The header value cannot start or end with a space.  If the ValueType parameter is set to ReferenceHeader: The header value must be 1 to 128 characters in length, and can contain lowercase letters, digits, underscores (_), and hyphens (-). Valid values: `ClientSrcPort`, `ClientSrcIp`, `Protocol`, `SLBId`, `SLBPort`, `UserDefined`.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// Valid values:  UserDefined: a custom value ReferenceHeader: uses a field of the user request header. SystemDefined: a system value.
        /// </summary>
        [Input("valueType")]
        public Input<string>? ValueType { get; set; }

        public RuleRuleActionInsertHeaderConfigGetArgs()
        {
        }
    }
}
