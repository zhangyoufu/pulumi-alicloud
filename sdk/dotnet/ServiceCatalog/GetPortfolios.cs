// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ServiceCatalog
{
    public static class GetPortfolios
    {
        /// <summary>
        /// This data source provides Service Catalog Portfolio available to the user.[What is Portfolio](https://www.alibabacloud.com/help/en/service-catalog/developer-reference/api-servicecatalog-2021-09-01-createportfolio)
        /// 
        /// &gt; **NOTE:** Available in 1.204.0+
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
        ///     var @default = AliCloud.ServiceCatalog.GetPortfolios.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             defaultAlicloudServiceCatalogPortfolio.Id,
        ///         },
        ///         NameRegex = defaultAlicloudServiceCatalogPortfolio.Name,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudServiceCatalogPortfolioExampleId"] = @default.Apply(@default =&gt; @default.Apply(getPortfoliosResult =&gt; getPortfoliosResult.Portfolios[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetPortfoliosResult> InvokeAsync(GetPortfoliosArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPortfoliosResult>("alicloud:servicecatalog/getPortfolios:getPortfolios", args ?? new GetPortfoliosArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Service Catalog Portfolio available to the user.[What is Portfolio](https://www.alibabacloud.com/help/en/service-catalog/developer-reference/api-servicecatalog-2021-09-01-createportfolio)
        /// 
        /// &gt; **NOTE:** Available in 1.204.0+
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
        ///     var @default = AliCloud.ServiceCatalog.GetPortfolios.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             defaultAlicloudServiceCatalogPortfolio.Id,
        ///         },
        ///         NameRegex = defaultAlicloudServiceCatalogPortfolio.Name,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudServiceCatalogPortfolioExampleId"] = @default.Apply(@default =&gt; @default.Apply(getPortfoliosResult =&gt; getPortfoliosResult.Portfolios[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetPortfoliosResult> Invoke(GetPortfoliosInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPortfoliosResult>("alicloud:servicecatalog/getPortfolios:getPortfolios", args ?? new GetPortfoliosInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPortfoliosArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Portfolio IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Group Metric Rule name.
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
        /// The ID of the product.
        /// </summary>
        [Input("productId")]
        public string? ProductId { get; set; }

        /// <summary>
        /// The query scope. Valid values: `Local`(default), `Import`, `All`.
        /// </summary>
        [Input("scope")]
        public string? Scope { get; set; }

        /// <summary>
        /// The field that is used to sort the queried data. The value is fixed as CreateTime, which specifies the creation time of product portfolios.
        /// </summary>
        [Input("sortBy")]
        public string? SortBy { get; set; }

        /// <summary>
        /// The order in which you want to sort the queried data. Valid values: `Asc`, `Desc`.
        /// </summary>
        [Input("sortOrder")]
        public string? SortOrder { get; set; }

        public GetPortfoliosArgs()
        {
        }
        public static new GetPortfoliosArgs Empty => new GetPortfoliosArgs();
    }

    public sealed class GetPortfoliosInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Portfolio IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Group Metric Rule name.
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
        /// The ID of the product.
        /// </summary>
        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        /// <summary>
        /// The query scope. Valid values: `Local`(default), `Import`, `All`.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The field that is used to sort the queried data. The value is fixed as CreateTime, which specifies the creation time of product portfolios.
        /// </summary>
        [Input("sortBy")]
        public Input<string>? SortBy { get; set; }

        /// <summary>
        /// The order in which you want to sort the queried data. Valid values: `Asc`, `Desc`.
        /// </summary>
        [Input("sortOrder")]
        public Input<string>? SortOrder { get; set; }

        public GetPortfoliosInvokeArgs()
        {
        }
        public static new GetPortfoliosInvokeArgs Empty => new GetPortfoliosInvokeArgs();
    }


    [OutputType]
    public sealed class GetPortfoliosResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Portfolio IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of name of Portfolios.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        /// <summary>
        /// A list of Portfolio Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPortfoliosPortfolioResult> Portfolios;
        public readonly string? ProductId;
        public readonly string? Scope;
        public readonly string? SortBy;
        public readonly string? SortOrder;

        [OutputConstructor]
        private GetPortfoliosResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            int? pageNumber,

            int? pageSize,

            ImmutableArray<Outputs.GetPortfoliosPortfolioResult> portfolios,

            string? productId,

            string? scope,

            string? sortBy,

            string? sortOrder)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
            Portfolios = portfolios;
            ProductId = productId;
            Scope = scope;
            SortBy = sortBy;
            SortOrder = sortOrder;
        }
    }
}
