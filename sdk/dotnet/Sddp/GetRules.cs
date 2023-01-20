// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sddp
{
    public static class GetRules
    {
        /// <summary>
        /// This data source provides the Sddp Rules of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.132.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var defaultRule = new AliCloud.Sddp.Rule("defaultRule", new()
        ///     {
        ///         Category = 0,
        ///         Content = "content",
        ///         RuleName = "rule_name",
        ///         RiskLevelId = "4",
        ///         ProductCode = "ODPS",
        ///     });
        /// 
        ///     var defaultRules = AliCloud.Sddp.GetRules.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             defaultRule.Id,
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["sddpRuleId"] = defaultRules.Apply(getRulesResult =&gt; getRulesResult.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRulesResult> InvokeAsync(GetRulesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRulesResult>("alicloud:sddp/getRules:getRules", args ?? new GetRulesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Sddp Rules of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.132.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var defaultRule = new AliCloud.Sddp.Rule("defaultRule", new()
        ///     {
        ///         Category = 0,
        ///         Content = "content",
        ///         RuleName = "rule_name",
        ///         RiskLevelId = "4",
        ///         ProductCode = "ODPS",
        ///     });
        /// 
        ///     var defaultRules = AliCloud.Sddp.GetRules.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             defaultRule.Id,
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["sddpRuleId"] = defaultRules.Apply(getRulesResult =&gt; getRulesResult.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRulesResult> Invoke(GetRulesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRulesResult>("alicloud:sddp/getRules:getRules", args ?? new GetRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRulesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Sensitive Data Identification Rules for the Type of.
        /// </summary>
        [Input("category")]
        public int? Category { get; set; }

        /// <summary>
        /// The Content Classification.
        /// </summary>
        [Input("contentCategory")]
        public string? ContentCategory { get; set; }

        /// <summary>
        /// Sensitive Data Identification Rules of Type. 0: the Built-in 1: The User-Defined.
        /// </summary>
        [Input("customType")]
        public int? CustomType { get; set; }

        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Rule IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of rule.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// A regex string to filter results by Rule name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Product ID.
        /// </summary>
        [Input("productId")]
        public string? ProductId { get; set; }

        /// <summary>
        /// Sensitive Data Identification Rules of Risk Level ID. Valid values:1:S1, Weak Risk Level. 2:S2, Medium Risk Level. 3:S3 High Risk Level. 4:S4, the Highest Risk Level.
        /// </summary>
        [Input("riskLevelId")]
        public string? RiskLevelId { get; set; }

        /// <summary>
        /// Rule Type.
        /// </summary>
        [Input("ruleType")]
        public int? RuleType { get; set; }

        /// <summary>
        /// Sensitive Data Identification Rules Detection State of.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The Level of Risk.
        /// </summary>
        [Input("warnLevel")]
        public int? WarnLevel { get; set; }

        public GetRulesArgs()
        {
        }
        public static new GetRulesArgs Empty => new GetRulesArgs();
    }

    public sealed class GetRulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Sensitive Data Identification Rules for the Type of.
        /// </summary>
        [Input("category")]
        public Input<int>? Category { get; set; }

        /// <summary>
        /// The Content Classification.
        /// </summary>
        [Input("contentCategory")]
        public Input<string>? ContentCategory { get; set; }

        /// <summary>
        /// Sensitive Data Identification Rules of Type. 0: the Built-in 1: The User-Defined.
        /// </summary>
        [Input("customType")]
        public Input<int>? CustomType { get; set; }

        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Rule IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A regex string to filter results by Rule name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Product ID.
        /// </summary>
        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        /// <summary>
        /// Sensitive Data Identification Rules of Risk Level ID. Valid values:1:S1, Weak Risk Level. 2:S2, Medium Risk Level. 3:S3 High Risk Level. 4:S4, the Highest Risk Level.
        /// </summary>
        [Input("riskLevelId")]
        public Input<string>? RiskLevelId { get; set; }

        /// <summary>
        /// Rule Type.
        /// </summary>
        [Input("ruleType")]
        public Input<int>? RuleType { get; set; }

        /// <summary>
        /// Sensitive Data Identification Rules Detection State of.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The Level of Risk.
        /// </summary>
        [Input("warnLevel")]
        public Input<int>? WarnLevel { get; set; }

        public GetRulesInvokeArgs()
        {
        }
        public static new GetRulesInvokeArgs Empty => new GetRulesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRulesResult
    {
        public readonly int? Category;
        public readonly string? ContentCategory;
        public readonly int? CustomType;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? Name;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? ProductId;
        public readonly string? RiskLevelId;
        public readonly int? RuleType;
        public readonly ImmutableArray<Outputs.GetRulesRuleResult> Rules;
        public readonly string? Status;
        public readonly int? WarnLevel;

        [OutputConstructor]
        private GetRulesResult(
            int? category,

            string? contentCategory,

            int? customType,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? name,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? productId,

            string? riskLevelId,

            int? ruleType,

            ImmutableArray<Outputs.GetRulesRuleResult> rules,

            string? status,

            int? warnLevel)
        {
            Category = category;
            ContentCategory = contentCategory;
            CustomType = customType;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            Name = name;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ProductId = productId;
            RiskLevelId = riskLevelId;
            RuleType = ruleType;
            Rules = rules;
            Status = status;
            WarnLevel = warnLevel;
        }
    }
}
