// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cfg
{
    public static class GetRules
    {
        /// <summary>
        /// This data source provides the Config Rules of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:**  Available in 1.99.0+.
        /// 
        /// &gt; **NOTE:** The Cloud Config region only support `cn-shanghai` and `ap-northeast-1`.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(AliCloud.Cfg.GetRules.InvokeAsync(new AliCloud.Cfg.GetRulesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "cr-ed4bad756057********",
        ///             },
        ///             NameRegex = "tftest",
        ///         }));
        ///         this.FirstConfigRuleId = example.Apply(example =&gt; example.Rules[0].Id);
        ///     }
        /// 
        ///     [Output("firstConfigRuleId")]
        ///     public Output&lt;string&gt; FirstConfigRuleId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRulesResult> InvokeAsync(GetRulesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRulesResult>("alicloud:cfg/getRules:getRules", args ?? new GetRulesArgs(), options.WithVersion());
    }


    public sealed class GetRulesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Field `config_rule_state` has been deprecated from provider version 1.124.1. New field `status` instead.
        /// </summary>
        [Input("configRuleState")]
        public string? ConfigRuleState { get; set; }

        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Config Rule IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The ID of the member account to which the rule to be queried belongs. The default is empty. When `multi_account` is set to true, this parameter is valid.
        /// </summary>
        [Input("memberId")]
        public int? MemberId { get; set; }

        /// <summary>
        /// Whether the enterprise management account queries the rule details of member accounts.
        /// </summary>
        [Input("multiAccount")]
        public bool? MultiAccount { get; set; }

        /// <summary>
        /// A regex string to filter results by rule name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The risk level of Config Rule. Valid values: `1`: Critical ,`2`: Warning , `3`: Info.
        /// </summary>
        [Input("riskLevel")]
        public int? RiskLevel { get; set; }

        /// <summary>
        /// The name of config rule.
        /// </summary>
        [Input("ruleName")]
        public string? RuleName { get; set; }

        /// <summary>
        /// The status of the config rule, valid values: `ACTIVE`, `DELETING`, `EVALUATING` and `INACTIVE`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetRulesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRulesResult
    {
        public readonly string? ConfigRuleState;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Config Rule IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly int? MemberId;
        public readonly bool? MultiAccount;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of Config Rule names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly int? RiskLevel;
        public readonly string? RuleName;
        /// <summary>
        /// A list of Config Rules. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRulesRuleResult> Rules;
        /// <summary>
        /// (Available in 1.124.1+) The status of config rule.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private GetRulesResult(
            string? configRuleState,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            int? memberId,

            bool? multiAccount,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            int? riskLevel,

            string? ruleName,

            ImmutableArray<Outputs.GetRulesRuleResult> rules,

            string? status)
        {
            ConfigRuleState = configRuleState;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            MemberId = memberId;
            MultiAccount = multiAccount;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            RiskLevel = riskLevel;
            RuleName = ruleName;
            Rules = rules;
            Status = status;
        }
    }
}
