// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Mhub
{
    public static class GetProducts
    {
        /// <summary>
        /// This data source provides the Mhub Products of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.138.0+.
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
        ///     var config = new Config();
        ///     var name = config.Get("name") ?? "example_value";
        ///     var @default = new AliCloud.Mhub.Product("default", new()
        ///     {
        ///         ProductName = name,
        ///     });
        /// 
        ///     var ids = AliCloud.Mhub.GetProducts.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Mhub.GetProducts.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Product",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["mhubProductId1"] = ids.Apply(getProductsResult =&gt; getProductsResult.Products[0]?.Id),
        ///         ["mhubProductId2"] = nameRegex.Apply(getProductsResult =&gt; getProductsResult.Products[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProductsResult> InvokeAsync(GetProductsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProductsResult>("alicloud:mhub/getProducts:getProducts", args ?? new GetProductsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Mhub Products of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.138.0+.
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
        ///     var config = new Config();
        ///     var name = config.Get("name") ?? "example_value";
        ///     var @default = new AliCloud.Mhub.Product("default", new()
        ///     {
        ///         ProductName = name,
        ///     });
        /// 
        ///     var ids = AliCloud.Mhub.GetProducts.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Mhub.GetProducts.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Product",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["mhubProductId1"] = ids.Apply(getProductsResult =&gt; getProductsResult.Products[0]?.Id),
        ///         ["mhubProductId2"] = nameRegex.Apply(getProductsResult =&gt; getProductsResult.Products[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProductsResult> Invoke(GetProductsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProductsResult>("alicloud:mhub/getProducts:getProducts", args ?? new GetProductsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProductsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Product IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Product name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetProductsArgs()
        {
        }
        public static new GetProductsArgs Empty => new GetProductsArgs();
    }

    public sealed class GetProductsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Product IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Product name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetProductsInvokeArgs()
        {
        }
        public static new GetProductsInvokeArgs Empty => new GetProductsInvokeArgs();
    }


    [OutputType]
    public sealed class GetProductsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetProductsProductResult> Products;

        [OutputConstructor]
        private GetProductsResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetProductsProductResult> products)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Products = products;
        }
    }
}
