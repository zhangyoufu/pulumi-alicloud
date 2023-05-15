// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms
{
    public static class GetEventRules
    {
        /// <summary>
        /// This data source provides the Cms Event Rules of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.182.0+.
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
        ///     var ids = AliCloud.Cms.GetEventRules.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Cms.GetEventRules.Invoke(new()
        ///     {
        ///         NameRegex = "^my-EventRule",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cmsEventRuleId1"] = ids.Apply(getEventRulesResult =&gt; getEventRulesResult.Rules[0]?.Id),
        ///         ["cmsEventRuleId2"] = nameRegex.Apply(getEventRulesResult =&gt; getEventRulesResult.Rules[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEventRulesResult> InvokeAsync(GetEventRulesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEventRulesResult>("alicloud:cms/getEventRules:getEventRules", args ?? new GetEventRulesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Cms Event Rules of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.182.0+.
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
        ///     var ids = AliCloud.Cms.GetEventRules.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Cms.GetEventRules.Invoke(new()
        ///     {
        ///         NameRegex = "^my-EventRule",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cmsEventRuleId1"] = ids.Apply(getEventRulesResult =&gt; getEventRulesResult.Rules[0]?.Id),
        ///         ["cmsEventRuleId2"] = nameRegex.Apply(getEventRulesResult =&gt; getEventRulesResult.Rules[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEventRulesResult> Invoke(GetEventRulesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEventRulesResult>("alicloud:cms/getEventRules:getEventRules", args ?? new GetEventRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEventRulesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Event Rule IDs. Its element value is same as Event Rule Name.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name prefix.
        /// </summary>
        [Input("namePrefix")]
        public string? NamePrefix { get; set; }

        /// <summary>
        /// A regex string to filter results by Event Rule name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetEventRulesArgs()
        {
        }
        public static new GetEventRulesArgs Empty => new GetEventRulesArgs();
    }

    public sealed class GetEventRulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Event Rule IDs. Its element value is same as Event Rule Name.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name prefix.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// A regex string to filter results by Event Rule name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetEventRulesInvokeArgs()
        {
        }
        public static new GetEventRulesInvokeArgs Empty => new GetEventRulesInvokeArgs();
    }


    [OutputType]
    public sealed class GetEventRulesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NamePrefix;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        public readonly ImmutableArray<Outputs.GetEventRulesRuleResult> Rules;
        public readonly string? Status;

        [OutputConstructor]
        private GetEventRulesResult(
            string id,

            ImmutableArray<string> ids,

            string? namePrefix,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            int? pageNumber,

            int? pageSize,

            ImmutableArray<Outputs.GetEventRulesRuleResult> rules,

            string? status)
        {
            Id = id;
            Ids = ids;
            NamePrefix = namePrefix;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
            Rules = rules;
            Status = status;
        }
    }
}
