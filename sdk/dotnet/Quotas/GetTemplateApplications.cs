// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Quotas
{
    public static class GetTemplateApplications
    {
        /// <summary>
        /// This data source provides Quotas Template Applications available to the user.[What is Template Applications](https://www.alibabacloud.com/help/en/quota-center/developer-reference/api-quotas-2020-05-10-createquotaapplicationsfortemplate)
        /// 
        /// &gt; **NOTE:** Available since v1.214.0.
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
        ///     var @default = AliCloud.ResourceManager.GetAccounts.Invoke(new()
        ///     {
        ///         Status = "CreateSuccess",
        ///     });
        /// 
        ///     var defaultTemplateApplications = new AliCloud.Quotas.TemplateApplications("default", new()
        ///     {
        ///         QuotaActionCode = "vpc_whitelist/ha_vip_whitelist",
        ///         ProductCode = "vpc",
        ///         QuotaCategory = "FlowControl",
        ///         AliyunUids = new[]
        ///         {
        ///             @default.Apply(@default =&gt; @default.Apply(getAccountsResult =&gt; getAccountsResult.Ids[0])),
        ///         },
        ///         DesireValue = 6,
        ///         NoticeType = 0,
        ///         EnvLanguage = "zh",
        ///         Reason = "example",
        ///         Dimensions = new[]
        ///         {
        ///             new AliCloud.Quotas.Inputs.TemplateApplicationsDimensionArgs
        ///             {
        ///                 Key = "apiName",
        ///                 Value = "GetProductQuotaDimension",
        ///             },
        ///             new AliCloud.Quotas.Inputs.TemplateApplicationsDimensionArgs
        ///             {
        ///                 Key = "apiVersion",
        ///                 Value = "2020-05-10",
        ///             },
        ///             new AliCloud.Quotas.Inputs.TemplateApplicationsDimensionArgs
        ///             {
        ///                 Key = "regionId",
        ///                 Value = "cn-hangzhou",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var defaultGetTemplateApplications = AliCloud.Quotas.GetTemplateApplications.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             defaultTemplateApplications.Id,
        ///         },
        ///         ProductCode = "vpc",
        ///         QuotaActionCode = "vpc_whitelist/ha_vip_whitelist",
        ///         QuotaCategory = "FlowControl",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudQuotasTemplateApplicationsExampleId"] = defaultGetTemplateApplications.Apply(getTemplateApplicationsResult =&gt; getTemplateApplicationsResult.Applications[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetTemplateApplicationsResult> InvokeAsync(GetTemplateApplicationsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTemplateApplicationsResult>("alicloud:quotas/getTemplateApplications:getTemplateApplications", args ?? new GetTemplateApplicationsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Quotas Template Applications available to the user.[What is Template Applications](https://www.alibabacloud.com/help/en/quota-center/developer-reference/api-quotas-2020-05-10-createquotaapplicationsfortemplate)
        /// 
        /// &gt; **NOTE:** Available since v1.214.0.
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
        ///     var @default = AliCloud.ResourceManager.GetAccounts.Invoke(new()
        ///     {
        ///         Status = "CreateSuccess",
        ///     });
        /// 
        ///     var defaultTemplateApplications = new AliCloud.Quotas.TemplateApplications("default", new()
        ///     {
        ///         QuotaActionCode = "vpc_whitelist/ha_vip_whitelist",
        ///         ProductCode = "vpc",
        ///         QuotaCategory = "FlowControl",
        ///         AliyunUids = new[]
        ///         {
        ///             @default.Apply(@default =&gt; @default.Apply(getAccountsResult =&gt; getAccountsResult.Ids[0])),
        ///         },
        ///         DesireValue = 6,
        ///         NoticeType = 0,
        ///         EnvLanguage = "zh",
        ///         Reason = "example",
        ///         Dimensions = new[]
        ///         {
        ///             new AliCloud.Quotas.Inputs.TemplateApplicationsDimensionArgs
        ///             {
        ///                 Key = "apiName",
        ///                 Value = "GetProductQuotaDimension",
        ///             },
        ///             new AliCloud.Quotas.Inputs.TemplateApplicationsDimensionArgs
        ///             {
        ///                 Key = "apiVersion",
        ///                 Value = "2020-05-10",
        ///             },
        ///             new AliCloud.Quotas.Inputs.TemplateApplicationsDimensionArgs
        ///             {
        ///                 Key = "regionId",
        ///                 Value = "cn-hangzhou",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var defaultGetTemplateApplications = AliCloud.Quotas.GetTemplateApplications.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             defaultTemplateApplications.Id,
        ///         },
        ///         ProductCode = "vpc",
        ///         QuotaActionCode = "vpc_whitelist/ha_vip_whitelist",
        ///         QuotaCategory = "FlowControl",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudQuotasTemplateApplicationsExampleId"] = defaultGetTemplateApplications.Apply(getTemplateApplicationsResult =&gt; getTemplateApplicationsResult.Applications[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetTemplateApplicationsResult> Invoke(GetTemplateApplicationsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTemplateApplicationsResult>("alicloud:quotas/getTemplateApplications:getTemplateApplications", args ?? new GetTemplateApplicationsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTemplateApplicationsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the quota application batch.
        /// </summary>
        [Input("batchQuotaApplicationId")]
        public string? BatchQuotaApplicationId { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Template Applications IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Cloud service name abbreviation.&gt; For more information about cloud services that support quota centers, see Cloud services that support quota centers.
        /// </summary>
        [Input("productCode")]
        public string? ProductCode { get; set; }

        /// <summary>
        /// The quota ID.
        /// </summary>
        [Input("quotaActionCode")]
        public string? QuotaActionCode { get; set; }

        /// <summary>
        /// The quota type. Value: `CommonQuota`, `FlowControl` and `WhiteListLabel`.
        /// </summary>
        [Input("quotaCategory")]
        public string? QuotaCategory { get; set; }

        public GetTemplateApplicationsArgs()
        {
        }
        public static new GetTemplateApplicationsArgs Empty => new GetTemplateApplicationsArgs();
    }

    public sealed class GetTemplateApplicationsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the quota application batch.
        /// </summary>
        [Input("batchQuotaApplicationId")]
        public Input<string>? BatchQuotaApplicationId { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Template Applications IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Cloud service name abbreviation.&gt; For more information about cloud services that support quota centers, see Cloud services that support quota centers.
        /// </summary>
        [Input("productCode")]
        public Input<string>? ProductCode { get; set; }

        /// <summary>
        /// The quota ID.
        /// </summary>
        [Input("quotaActionCode")]
        public Input<string>? QuotaActionCode { get; set; }

        /// <summary>
        /// The quota type. Value: `CommonQuota`, `FlowControl` and `WhiteListLabel`.
        /// </summary>
        [Input("quotaCategory")]
        public Input<string>? QuotaCategory { get; set; }

        public GetTemplateApplicationsInvokeArgs()
        {
        }
        public static new GetTemplateApplicationsInvokeArgs Empty => new GetTemplateApplicationsInvokeArgs();
    }


    [OutputType]
    public sealed class GetTemplateApplicationsResult
    {
        /// <summary>
        /// A list of Template Applications Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTemplateApplicationsApplicationResult> Applications;
        /// <summary>
        /// The ID of the quota application batch.
        /// </summary>
        public readonly string? BatchQuotaApplicationId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Template Applications IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        /// <summary>
        /// Cloud service name abbreviation.&gt; For more information about cloud services that support quota centers, see Cloud services that support quota centers.
        /// </summary>
        public readonly string? ProductCode;
        /// <summary>
        /// The quota ID.
        /// </summary>
        public readonly string? QuotaActionCode;
        /// <summary>
        /// The quota type. Value:-CommonQuota (default): Generic quota.-FlowControl:API rate quota.-WhiteListLabel: Equity quota.
        /// </summary>
        public readonly string? QuotaCategory;

        [OutputConstructor]
        private GetTemplateApplicationsResult(
            ImmutableArray<Outputs.GetTemplateApplicationsApplicationResult> applications,

            string? batchQuotaApplicationId,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            string? productCode,

            string? quotaActionCode,

            string? quotaCategory)
        {
            Applications = applications;
            BatchQuotaApplicationId = batchQuotaApplicationId;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            ProductCode = productCode;
            QuotaActionCode = quotaActionCode;
            QuotaCategory = quotaCategory;
        }
    }
}
