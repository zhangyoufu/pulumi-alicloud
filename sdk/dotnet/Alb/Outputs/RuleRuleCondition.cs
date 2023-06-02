// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Outputs
{

    [OutputType]
    public sealed class RuleRuleCondition
    {
        /// <summary>
        /// The configuration of the cookie. See the following `Block cookie_config`.
        /// </summary>
        public readonly Outputs.RuleRuleConditionCookieConfig? CookieConfig;
        /// <summary>
        /// The configuration of the header field. See the following `Block header_config`.
        /// </summary>
        public readonly Outputs.RuleRuleConditionHeaderConfig? HeaderConfig;
        /// <summary>
        /// The configuration of the host field. See the following `Block host_config`.
        /// </summary>
        public readonly Outputs.RuleRuleConditionHostConfig? HostConfig;
        /// <summary>
        /// The configuration of the request method. See the following `Block method_config`.
        /// </summary>
        public readonly Outputs.RuleRuleConditionMethodConfig? MethodConfig;
        /// <summary>
        /// The configuration of the path for the request to be forwarded. See the following `Block path_config`.
        /// </summary>
        public readonly Outputs.RuleRuleConditionPathConfig? PathConfig;
        /// <summary>
        /// The configuration of the query string. See the following `Block query_string_config`.
        /// </summary>
        public readonly Outputs.RuleRuleConditionQueryStringConfig? QueryStringConfig;
        /// <summary>
        /// The Based on source IP traffic matching. Required and valid when Type is SourceIP. See the following `Block source_ip_config`.
        /// </summary>
        public readonly Outputs.RuleRuleConditionSourceIpConfig? SourceIpConfig;
        /// <summary>
        /// The type of the forwarding rule. Valid values: `Header`, `Host`, `Path`,  `Cookie`, `QueryString`, `Method` and `SourceIp`.
        /// **Note:**  The preceding actions can be classified into two types:  `FinalType`: A forwarding rule can contain only one `FinalType` action, which is executed last. This type of action can contain only one `ForwardGroup`, `Redirect` or `FixedResponse` action. `ExtType`: A forwarding rule can contain one or more `ExtType` actions, which are executed before `FinalType` actions and need to coexist with the `FinalType` actions. This type of action can contain multiple `InsertHeader` actions or one `Rewrite` action.
        /// **NOTE:** The `TrafficLimit` and `TrafficMirror` option is available in 1.162.0+.
        /// **NOTE:** From version 1.205.0+, `type` can be set to `Cors`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private RuleRuleCondition(
            Outputs.RuleRuleConditionCookieConfig? cookieConfig,

            Outputs.RuleRuleConditionHeaderConfig? headerConfig,

            Outputs.RuleRuleConditionHostConfig? hostConfig,

            Outputs.RuleRuleConditionMethodConfig? methodConfig,

            Outputs.RuleRuleConditionPathConfig? pathConfig,

            Outputs.RuleRuleConditionQueryStringConfig? queryStringConfig,

            Outputs.RuleRuleConditionSourceIpConfig? sourceIpConfig,

            string type)
        {
            CookieConfig = cookieConfig;
            HeaderConfig = headerConfig;
            HostConfig = hostConfig;
            MethodConfig = methodConfig;
            PathConfig = pathConfig;
            QueryStringConfig = queryStringConfig;
            SourceIpConfig = sourceIpConfig;
            Type = type;
        }
    }
}
