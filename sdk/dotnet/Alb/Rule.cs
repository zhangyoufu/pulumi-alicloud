// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb
{
    /// <summary>
    /// Provides a Application Load Balancer (ALB) Rule resource.
    /// 
    /// For information about Application Load Balancer (ALB) Rule and how to use it, see [What is Rule](https://www.alibabacloud.com/help/doc-detail/214375.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.133.0+.
    /// 
    /// &gt; **NOTE:** This version only supports forwarding rules in the request direction.
    /// 
    /// ## Import
    /// 
    /// Application Load Balancer (ALB) Rule can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:alb/rule:Rule example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:alb/rule:Rule")]
    public partial class Rule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether to precheck this request.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The ID of the listener to which the forwarding rule belongs.
        /// </summary>
        [Output("listenerId")]
        public Output<string> ListenerId { get; private set; } = null!;

        /// <summary>
        /// The priority of the rule. Valid values: 1 to 10000. A smaller value indicates a higher priority. **Note*:* The priority of each rule within the same listener must be unique.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// The actions of the forwarding rules. See the following `Block rule_actions`.
        /// </summary>
        [Output("ruleActions")]
        public Output<ImmutableArray<Outputs.RuleRuleAction>> RuleActions { get; private set; } = null!;

        /// <summary>
        /// The conditions of the forwarding rule. See the following `Block rule_conditions`.
        /// </summary>
        [Output("ruleConditions")]
        public Output<ImmutableArray<Outputs.RuleRuleCondition>> RuleConditions { get; private set; } = null!;

        /// <summary>
        /// The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
        /// </summary>
        [Output("ruleName")]
        public Output<string> RuleName { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Rule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Rule(string name, RuleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:alb/rule:Rule", name, args ?? new RuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Rule(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:alb/rule:Rule", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Rule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Rule Get(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
        {
            return new Rule(name, id, state, options);
        }
    }

    public sealed class RuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to precheck this request.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the listener to which the forwarding rule belongs.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        /// <summary>
        /// The priority of the rule. Valid values: 1 to 10000. A smaller value indicates a higher priority. **Note*:* The priority of each rule within the same listener must be unique.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        [Input("ruleActions", required: true)]
        private InputList<Inputs.RuleRuleActionArgs>? _ruleActions;

        /// <summary>
        /// The actions of the forwarding rules. See the following `Block rule_actions`.
        /// </summary>
        public InputList<Inputs.RuleRuleActionArgs> RuleActions
        {
            get => _ruleActions ?? (_ruleActions = new InputList<Inputs.RuleRuleActionArgs>());
            set => _ruleActions = value;
        }

        [Input("ruleConditions", required: true)]
        private InputList<Inputs.RuleRuleConditionArgs>? _ruleConditions;

        /// <summary>
        /// The conditions of the forwarding rule. See the following `Block rule_conditions`.
        /// </summary>
        public InputList<Inputs.RuleRuleConditionArgs> RuleConditions
        {
            get => _ruleConditions ?? (_ruleConditions = new InputList<Inputs.RuleRuleConditionArgs>());
            set => _ruleConditions = value;
        }

        /// <summary>
        /// The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        public RuleArgs()
        {
        }
        public static new RuleArgs Empty => new RuleArgs();
    }

    public sealed class RuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to precheck this request.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the listener to which the forwarding rule belongs.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// The priority of the rule. Valid values: 1 to 10000. A smaller value indicates a higher priority. **Note*:* The priority of each rule within the same listener must be unique.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("ruleActions")]
        private InputList<Inputs.RuleRuleActionGetArgs>? _ruleActions;

        /// <summary>
        /// The actions of the forwarding rules. See the following `Block rule_actions`.
        /// </summary>
        public InputList<Inputs.RuleRuleActionGetArgs> RuleActions
        {
            get => _ruleActions ?? (_ruleActions = new InputList<Inputs.RuleRuleActionGetArgs>());
            set => _ruleActions = value;
        }

        [Input("ruleConditions")]
        private InputList<Inputs.RuleRuleConditionGetArgs>? _ruleConditions;

        /// <summary>
        /// The conditions of the forwarding rule. See the following `Block rule_conditions`.
        /// </summary>
        public InputList<Inputs.RuleRuleConditionGetArgs> RuleConditions
        {
            get => _ruleConditions ?? (_ruleConditions = new InputList<Inputs.RuleRuleConditionGetArgs>());
            set => _ruleConditions = value;
        }

        /// <summary>
        /// The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
        /// </summary>
        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public RuleState()
        {
        }
        public static new RuleState Empty => new RuleState();
    }
}
