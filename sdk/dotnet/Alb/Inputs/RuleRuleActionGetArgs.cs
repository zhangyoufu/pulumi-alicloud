// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Inputs
{

    public sealed class RuleRuleActionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Request forwarding based on CORS. See `cors_config` below.
        /// </summary>
        [Input("corsConfig")]
        public Input<Inputs.RuleRuleActionCorsConfigGetArgs>? CorsConfig { get; set; }

        /// <summary>
        /// The configuration of the fixed response. See `fixed_response_config` below.
        /// </summary>
        [Input("fixedResponseConfig")]
        public Input<Inputs.RuleRuleActionFixedResponseConfigGetArgs>? FixedResponseConfig { get; set; }

        /// <summary>
        /// The forward response action within ALB. See `forward_group_config` below.
        /// </summary>
        [Input("forwardGroupConfig")]
        public Input<Inputs.RuleRuleActionForwardGroupConfigGetArgs>? ForwardGroupConfig { get; set; }

        /// <summary>
        /// The configuration of the inserted header field. See `insert_header_config` below.
        /// </summary>
        [Input("insertHeaderConfig")]
        public Input<Inputs.RuleRuleActionInsertHeaderConfigGetArgs>? InsertHeaderConfig { get; set; }

        /// <summary>
        /// The order of the forwarding rule actions. Valid values: `1` to `50000`. The actions are performed in ascending order. You cannot leave this parameter empty. Each value must be unique.
        /// </summary>
        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        /// <summary>
        /// The configuration of the external redirect action. See `redirect_config` below.
        /// </summary>
        [Input("redirectConfig")]
        public Input<Inputs.RuleRuleActionRedirectConfigGetArgs>? RedirectConfig { get; set; }

        /// <summary>
        /// The redirect action within ALB. See `rewrite_config` below.
        /// </summary>
        [Input("rewriteConfig")]
        public Input<Inputs.RuleRuleActionRewriteConfigGetArgs>? RewriteConfig { get; set; }

        /// <summary>
        /// The Flow speed limit. See `traffic_limit_config` below.
        /// </summary>
        [Input("trafficLimitConfig")]
        public Input<Inputs.RuleRuleActionTrafficLimitConfigGetArgs>? TrafficLimitConfig { get; set; }

        /// <summary>
        /// The Traffic mirroring. See `traffic_mirror_config` below.
        /// </summary>
        [Input("trafficMirrorConfig")]
        public Input<Inputs.RuleRuleActionTrafficMirrorConfigGetArgs>? TrafficMirrorConfig { get; set; }

        /// <summary>
        /// The action type. Valid values: `ForwardGroup`, `Redirect`, `FixedResponse`, `Rewrite`, `InsertHeader`, `TrafficLimit`, `TrafficMirror` and `Cors`.
        /// **Note:** The preceding actions can be classified into two types:  `FinalType`: A forwarding rule can contain only one `FinalType` action, which is executed last. This type of action can contain only one `ForwardGroup`, `Redirect` or `FixedResponse` action. `ExtType`: A forwarding rule can contain one or more `ExtType` actions, which are executed before `FinalType` actions and need to coexist with the `FinalType` actions. This type of action can contain multiple `InsertHeader` actions or one `Rewrite` action.
        /// **NOTE:** The `TrafficLimit` and `TrafficMirror` option is available since 1.162.0.
        /// **NOTE:** From version 1.205.0, `type` can be set to `Cors`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public RuleRuleActionGetArgs()
        {
        }
        public static new RuleRuleActionGetArgs Empty => new RuleRuleActionGetArgs();
    }
}
