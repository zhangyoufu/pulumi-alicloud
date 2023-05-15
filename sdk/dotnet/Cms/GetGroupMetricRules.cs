// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms
{
    public static class GetGroupMetricRules
    {
        /// <summary>
        /// This data source provides the Cms Group Metric Rules of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.104.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AliCloud.Cms.GetGroupMetricRules.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "4a9a8978-a9cc-55ca-aa7c-530ccd91ae57",
        ///         },
        ///         NameRegex = "the_resource_name",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstCmsGroupMetricRuleId"] = example.Apply(getGroupMetricRulesResult =&gt; getGroupMetricRulesResult.Rules[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupMetricRulesResult> InvokeAsync(GetGroupMetricRulesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupMetricRulesResult>("alicloud:cms/getGroupMetricRules:getGroupMetricRules", args ?? new GetGroupMetricRulesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Cms Group Metric Rules of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.104.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AliCloud.Cms.GetGroupMetricRules.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "4a9a8978-a9cc-55ca-aa7c-530ccd91ae57",
        ///         },
        ///         NameRegex = "the_resource_name",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstCmsGroupMetricRuleId"] = example.Apply(getGroupMetricRulesResult =&gt; getGroupMetricRulesResult.Rules[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGroupMetricRulesResult> Invoke(GetGroupMetricRulesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupMetricRulesResult>("alicloud:cms/getGroupMetricRules:getGroupMetricRules", args ?? new GetGroupMetricRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupMetricRulesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The dimensions that specify the resources to be associated with the alert rule.
        /// </summary>
        [Input("dimensions")]
        public string? Dimensions { get; set; }

        /// <summary>
        /// Indicates whether the alert rule is enabled.
        /// </summary>
        [Input("enableState")]
        public bool? EnableState { get; set; }

        /// <summary>
        /// The ID of the application group.
        /// </summary>
        [Input("groupId")]
        public string? GroupId { get; set; }

        /// <summary>
        /// The name of the alert rule.
        /// </summary>
        [Input("groupMetricRuleName")]
        public string? GroupMetricRuleName { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Group Metric Rule IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of the metric.
        /// </summary>
        [Input("metricName")]
        public string? MetricName { get; set; }

        /// <summary>
        /// A regex string to filter results by Group Metric Rule name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// The namespace of the service.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of Group Metric Rule..
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetGroupMetricRulesArgs()
        {
        }
        public static new GetGroupMetricRulesArgs Empty => new GetGroupMetricRulesArgs();
    }

    public sealed class GetGroupMetricRulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The dimensions that specify the resources to be associated with the alert rule.
        /// </summary>
        [Input("dimensions")]
        public Input<string>? Dimensions { get; set; }

        /// <summary>
        /// Indicates whether the alert rule is enabled.
        /// </summary>
        [Input("enableState")]
        public Input<bool>? EnableState { get; set; }

        /// <summary>
        /// The ID of the application group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the alert rule.
        /// </summary>
        [Input("groupMetricRuleName")]
        public Input<string>? GroupMetricRuleName { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Group Metric Rule IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of the metric.
        /// </summary>
        [Input("metricName")]
        public Input<string>? MetricName { get; set; }

        /// <summary>
        /// A regex string to filter results by Group Metric Rule name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// The namespace of the service.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of Group Metric Rule..
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetGroupMetricRulesInvokeArgs()
        {
        }
        public static new GetGroupMetricRulesInvokeArgs Empty => new GetGroupMetricRulesInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupMetricRulesResult
    {
        public readonly string? Dimensions;
        public readonly bool? EnableState;
        public readonly string? GroupId;
        public readonly string? GroupMetricRuleName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? MetricName;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? Namespace;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetGroupMetricRulesRuleResult> Rules;
        public readonly string? Status;

        [OutputConstructor]
        private GetGroupMetricRulesResult(
            string? dimensions,

            bool? enableState,

            string? groupId,

            string? groupMetricRuleName,

            string id,

            ImmutableArray<string> ids,

            string? metricName,

            string? nameRegex,

            ImmutableArray<string> names,

            string? @namespace,

            string? outputFile,

            ImmutableArray<Outputs.GetGroupMetricRulesRuleResult> rules,

            string? status)
        {
            Dimensions = dimensions;
            EnableState = enableState;
            GroupId = groupId;
            GroupMetricRuleName = groupMetricRuleName;
            Id = id;
            Ids = ids;
            MetricName = metricName;
            NameRegex = nameRegex;
            Names = names;
            Namespace = @namespace;
            OutputFile = outputFile;
            Rules = rules;
            Status = status;
        }
    }
}
