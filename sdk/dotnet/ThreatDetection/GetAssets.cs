// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ThreatDetection
{
    public static class GetAssets
    {
        /// <summary>
        /// This data source provides Threat Detection Asset available to the user.[What is Asset](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describecloudcenterinstances)
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
        ///     var @default = AliCloud.ThreatDetection.GetAssets.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudThreatDetectionAssetExampleId"] = @default.Apply(@default =&gt; @default.Apply(getAssetsResult =&gt; getAssetsResult.Assets[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetAssetsResult> InvokeAsync(GetAssetsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAssetsResult>("alicloud:threatdetection/getAssets:getAssets", args ?? new GetAssetsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Threat Detection Asset available to the user.[What is Asset](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describecloudcenterinstances)
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
        ///     var @default = AliCloud.ThreatDetection.GetAssets.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudThreatDetectionAssetExampleId"] = @default.Apply(@default =&gt; @default.Apply(getAssetsResult =&gt; getAssetsResult.Assets[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetAssetsResult> Invoke(GetAssetsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAssetsResult>("alicloud:threatdetection/getAssets:getAssets", args ?? new GetAssetsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAssetsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Set the conditions for searching assets. This parameter is in JSON format. Note the case when you enter the parameter. **NOTE:** You can search for assets by using conditions such as the instance ID, instance name, VPC ID, region, and public IP address of the asset.
        /// </summary>
        [Input("criteria")]
        public string? Criteria { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Asset IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Set asset importance. Value:
        /// - **2**: Significant assets
        /// - **1**: General assets
        /// - **0**: Test asset
        /// </summary>
        [Input("importance")]
        public int? Importance { get; set; }

        /// <summary>
        /// Set the logical relationship between multiple search conditions. The default value is **OR**. Valid values:
        /// - **OR**: indicates that the relationship between multiple search conditions is **OR**.
        /// - **AND**: indicates that the relationship between multiple search conditions is **AND**.
        /// </summary>
        [Input("logicalExp")]
        public string? LogicalExp { get; set; }

        /// <summary>
        /// The type of asset to query. Value:
        /// - **ecs**: server.
        /// - **cloud_product**: Cloud product.
        /// </summary>
        [Input("machineTypes")]
        public string? MachineTypes { get; set; }

        /// <summary>
        /// Specifies whether to internationalize the name of the default group. Default value: false
        /// </summary>
        [Input("noGroupTrace")]
        public bool? NoGroupTrace { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        public GetAssetsArgs()
        {
        }
        public static new GetAssetsArgs Empty => new GetAssetsArgs();
    }

    public sealed class GetAssetsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Set the conditions for searching assets. This parameter is in JSON format. Note the case when you enter the parameter. **NOTE:** You can search for assets by using conditions such as the instance ID, instance name, VPC ID, region, and public IP address of the asset.
        /// </summary>
        [Input("criteria")]
        public Input<string>? Criteria { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Asset IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Set asset importance. Value:
        /// - **2**: Significant assets
        /// - **1**: General assets
        /// - **0**: Test asset
        /// </summary>
        [Input("importance")]
        public Input<int>? Importance { get; set; }

        /// <summary>
        /// Set the logical relationship between multiple search conditions. The default value is **OR**. Valid values:
        /// - **OR**: indicates that the relationship between multiple search conditions is **OR**.
        /// - **AND**: indicates that the relationship between multiple search conditions is **AND**.
        /// </summary>
        [Input("logicalExp")]
        public Input<string>? LogicalExp { get; set; }

        /// <summary>
        /// The type of asset to query. Value:
        /// - **ecs**: server.
        /// - **cloud_product**: Cloud product.
        /// </summary>
        [Input("machineTypes")]
        public Input<string>? MachineTypes { get; set; }

        /// <summary>
        /// Specifies whether to internationalize the name of the default group. Default value: false
        /// </summary>
        [Input("noGroupTrace")]
        public Input<bool>? NoGroupTrace { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        public GetAssetsInvokeArgs()
        {
        }
        public static new GetAssetsInvokeArgs Empty => new GetAssetsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAssetsResult
    {
        /// <summary>
        /// A list of Asset Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAssetsAssetResult> Assets;
        public readonly string? Criteria;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Asset IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly int? Importance;
        public readonly string? LogicalExp;
        public readonly string? MachineTypes;
        public readonly bool? NoGroupTrace;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;

        [OutputConstructor]
        private GetAssetsResult(
            ImmutableArray<Outputs.GetAssetsAssetResult> assets,

            string? criteria,

            string id,

            ImmutableArray<string> ids,

            int? importance,

            string? logicalExp,

            string? machineTypes,

            bool? noGroupTrace,

            string? outputFile,

            int? pageNumber,

            int? pageSize)
        {
            Assets = assets;
            Criteria = criteria;
            Id = id;
            Ids = ids;
            Importance = importance;
            LogicalExp = logicalExp;
            MachineTypes = machineTypes;
            NoGroupTrace = noGroupTrace;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
        }
    }
}
