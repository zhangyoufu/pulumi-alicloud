// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Quotas
{
    public static class GetQuotas
    {
        /// <summary>
        /// This data source provides the Quotas Quotas of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.115.0+.
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
        ///     var example = AliCloud.Quotas.GetQuotas.Invoke(new()
        ///     {
        ///         ProductCode = "ecs",
        ///         NameRegex = "专有宿主机总数量上限",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstQuotasQuotaId"] = example.Apply(getQuotasResult =&gt; getQuotasResult.Quotas[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetQuotasResult> InvokeAsync(GetQuotasArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetQuotasResult>("alicloud:quotas/getQuotas:getQuotas", args ?? new GetQuotasArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Quotas Quotas of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.115.0+.
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
        ///     var example = AliCloud.Quotas.GetQuotas.Invoke(new()
        ///     {
        ///         ProductCode = "ecs",
        ///         NameRegex = "专有宿主机总数量上限",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstQuotasQuotaId"] = example.Apply(getQuotasResult =&gt; getQuotasResult.Quotas[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetQuotasResult> Invoke(GetQuotasInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetQuotasResult>("alicloud:quotas/getQuotas:getQuotas", args ?? new GetQuotasInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetQuotasArgs : global::Pulumi.InvokeArgs
    {
        [Input("dimensions")]
        private List<Inputs.GetQuotasDimensionArgs>? _dimensions;

        /// <summary>
        /// The dimensions.
        /// </summary>
        public List<Inputs.GetQuotasDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new List<Inputs.GetQuotasDimensionArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// The group code.
        /// </summary>
        [Input("groupCode")]
        public string? GroupCode { get; set; }

        /// <summary>
        /// The key word.
        /// </summary>
        [Input("keyWord")]
        public string? KeyWord { get; set; }

        /// <summary>
        /// A regex string to filter results by Quota name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The product code.
        /// </summary>
        [Input("productCode", required: true)]
        public string ProductCode { get; set; } = null!;

        /// <summary>
        /// The quota action code.
        /// </summary>
        [Input("quotaActionCode")]
        public string? QuotaActionCode { get; set; }

        /// <summary>
        /// The category of quota. Valid Values: `FlowControl` and `CommonQuota`.
        /// </summary>
        [Input("quotaCategory")]
        public string? QuotaCategory { get; set; }

        /// <summary>
        /// Cloud service ECS specification quota supports setting sorting fields. Valid Values: `TIME`, `TOTAL` and `RESERVED`.
        /// </summary>
        [Input("sortField")]
        public string? SortField { get; set; }

        /// <summary>
        /// Ranking of cloud service ECS specification quota support. Valid Values: `Ascending` and `Descending`.
        /// </summary>
        [Input("sortOrder")]
        public string? SortOrder { get; set; }

        public GetQuotasArgs()
        {
        }
        public static new GetQuotasArgs Empty => new GetQuotasArgs();
    }

    public sealed class GetQuotasInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("dimensions")]
        private InputList<Inputs.GetQuotasDimensionInputArgs>? _dimensions;

        /// <summary>
        /// The dimensions.
        /// </summary>
        public InputList<Inputs.GetQuotasDimensionInputArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.GetQuotasDimensionInputArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// The group code.
        /// </summary>
        [Input("groupCode")]
        public Input<string>? GroupCode { get; set; }

        /// <summary>
        /// The key word.
        /// </summary>
        [Input("keyWord")]
        public Input<string>? KeyWord { get; set; }

        /// <summary>
        /// A regex string to filter results by Quota name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The product code.
        /// </summary>
        [Input("productCode", required: true)]
        public Input<string> ProductCode { get; set; } = null!;

        /// <summary>
        /// The quota action code.
        /// </summary>
        [Input("quotaActionCode")]
        public Input<string>? QuotaActionCode { get; set; }

        /// <summary>
        /// The category of quota. Valid Values: `FlowControl` and `CommonQuota`.
        /// </summary>
        [Input("quotaCategory")]
        public Input<string>? QuotaCategory { get; set; }

        /// <summary>
        /// Cloud service ECS specification quota supports setting sorting fields. Valid Values: `TIME`, `TOTAL` and `RESERVED`.
        /// </summary>
        [Input("sortField")]
        public Input<string>? SortField { get; set; }

        /// <summary>
        /// Ranking of cloud service ECS specification quota support. Valid Values: `Ascending` and `Descending`.
        /// </summary>
        [Input("sortOrder")]
        public Input<string>? SortOrder { get; set; }

        public GetQuotasInvokeArgs()
        {
        }
        public static new GetQuotasInvokeArgs Empty => new GetQuotasInvokeArgs();
    }


    [OutputType]
    public sealed class GetQuotasResult
    {
        public readonly ImmutableArray<Outputs.GetQuotasDimensionResult> Dimensions;
        public readonly string? GroupCode;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? KeyWord;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string ProductCode;
        public readonly string? QuotaActionCode;
        public readonly string? QuotaCategory;
        public readonly ImmutableArray<Outputs.GetQuotasQuotaResult> Quotas;
        public readonly string? SortField;
        public readonly string? SortOrder;

        [OutputConstructor]
        private GetQuotasResult(
            ImmutableArray<Outputs.GetQuotasDimensionResult> dimensions,

            string? groupCode,

            string id,

            ImmutableArray<string> ids,

            string? keyWord,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string productCode,

            string? quotaActionCode,

            string? quotaCategory,

            ImmutableArray<Outputs.GetQuotasQuotaResult> quotas,

            string? sortField,

            string? sortOrder)
        {
            Dimensions = dimensions;
            GroupCode = groupCode;
            Id = id;
            Ids = ids;
            KeyWord = keyWord;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ProductCode = productCode;
            QuotaActionCode = quotaActionCode;
            QuotaCategory = quotaCategory;
            Quotas = quotas;
            SortField = sortField;
            SortOrder = sortOrder;
        }
    }
}
