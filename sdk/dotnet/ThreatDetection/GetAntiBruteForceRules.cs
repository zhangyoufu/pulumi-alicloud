// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ThreatDetection
{
    public static class GetAntiBruteForceRules
    {
        /// <summary>
        /// This data source provides Threat Detection Anti Brute Force Rule available to the user.[What is Anti Brute Force Rule](https://www.alibabacloud.com/help/en/security-center/latest/api-sas-2018-12-03-createantibruteforcerule)
        /// 
        /// &gt; **NOTE:** Available since v1.195.0.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var config = new Config();
        ///     var name = config.Get("name") ?? "example_value";
        ///     var defaultAntiBruteForceRule = new AliCloud.ThreatDetection.AntiBruteForceRule("default", new()
        ///     {
        ///         AntiBruteForceRuleName = name,
        ///         ForbiddenTime = 360,
        ///         UuidLists = new[]
        ///         {
        ///             "7567806c-4ec5-4597-9543-7c9543381a13",
        ///         },
        ///         FailCount = 80,
        ///         Span = 10,
        ///     });
        /// 
        ///     var @default = AliCloud.ThreatDetection.GetAntiBruteForceRules.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             defaultAntiBruteForceRule.Id,
        ///         },
        ///         NameRegex = defaultAntiBruteForceRule.Name,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudThreatDetectionAntiBruteForceRuleExampleId"] = @default.Apply(@default =&gt; @default.Apply(getAntiBruteForceRulesResult =&gt; getAntiBruteForceRulesResult.Rules[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetAntiBruteForceRulesResult> InvokeAsync(GetAntiBruteForceRulesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAntiBruteForceRulesResult>("alicloud:threatdetection/getAntiBruteForceRules:getAntiBruteForceRules", args ?? new GetAntiBruteForceRulesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Threat Detection Anti Brute Force Rule available to the user.[What is Anti Brute Force Rule](https://www.alibabacloud.com/help/en/security-center/latest/api-sas-2018-12-03-createantibruteforcerule)
        /// 
        /// &gt; **NOTE:** Available since v1.195.0.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var config = new Config();
        ///     var name = config.Get("name") ?? "example_value";
        ///     var defaultAntiBruteForceRule = new AliCloud.ThreatDetection.AntiBruteForceRule("default", new()
        ///     {
        ///         AntiBruteForceRuleName = name,
        ///         ForbiddenTime = 360,
        ///         UuidLists = new[]
        ///         {
        ///             "7567806c-4ec5-4597-9543-7c9543381a13",
        ///         },
        ///         FailCount = 80,
        ///         Span = 10,
        ///     });
        /// 
        ///     var @default = AliCloud.ThreatDetection.GetAntiBruteForceRules.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             defaultAntiBruteForceRule.Id,
        ///         },
        ///         NameRegex = defaultAntiBruteForceRule.Name,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudThreatDetectionAntiBruteForceRuleExampleId"] = @default.Apply(@default =&gt; @default.Apply(getAntiBruteForceRulesResult =&gt; getAntiBruteForceRulesResult.Rules[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetAntiBruteForceRulesResult> Invoke(GetAntiBruteForceRulesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAntiBruteForceRulesResult>("alicloud:threatdetection/getAntiBruteForceRules:getAntiBruteForceRules", args ?? new GetAntiBruteForceRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAntiBruteForceRulesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Anti-Brute Force Rule IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by the name of the defense rule.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetAntiBruteForceRulesArgs()
        {
        }
        public static new GetAntiBruteForceRulesArgs Empty => new GetAntiBruteForceRulesArgs();
    }

    public sealed class GetAntiBruteForceRulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Anti-Brute Force Rule IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by the name of the defense rule.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetAntiBruteForceRulesInvokeArgs()
        {
        }
        public static new GetAntiBruteForceRulesInvokeArgs Empty => new GetAntiBruteForceRulesInvokeArgs();
    }


    [OutputType]
    public sealed class GetAntiBruteForceRulesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Anti Brute Force Rule IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of name of Anti Brute Force Rules.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of Anti Brute Force Rule Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAntiBruteForceRulesRuleResult> Rules;

        [OutputConstructor]
        private GetAntiBruteForceRulesResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetAntiBruteForceRulesRuleResult> rules)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Rules = rules;
        }
    }
}
