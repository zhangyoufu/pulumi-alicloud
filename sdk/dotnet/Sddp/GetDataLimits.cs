// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sddp
{
    public static class GetDataLimits
    {
        /// <summary>
        /// This data source provides the Sddp Data Limits of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.159.0+.
        /// 
        /// ## Example Usage
        /// 
        /// Basic Usage
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
        ///     var ids = AliCloud.Sddp.GetDataLimits.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["sddpDataLimitId1"] = ids.Apply(getDataLimitsResult =&gt; getDataLimitsResult.Limits[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetDataLimitsResult> InvokeAsync(GetDataLimitsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDataLimitsResult>("alicloud:sddp/getDataLimits:getDataLimits", args ?? new GetDataLimitsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Sddp Data Limits of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.159.0+.
        /// 
        /// ## Example Usage
        /// 
        /// Basic Usage
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
        ///     var ids = AliCloud.Sddp.GetDataLimits.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["sddpDataLimitId1"] = ids.Apply(getDataLimitsResult =&gt; getDataLimitsResult.Limits[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetDataLimitsResult> Invoke(GetDataLimitsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataLimitsResult>("alicloud:sddp/getDataLimits:getDataLimits", args ?? new GetDataLimitsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataLimitsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Data Limit IDs.
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
        /// The ID of the data asset.
        /// </summary>
        [Input("parentId")]
        public string? ParentId { get; set; }

        /// <summary>
        /// The type of the service to which the data asset belongs.
        /// </summary>
        [Input("resourceType")]
        public string? ResourceType { get; set; }

        public GetDataLimitsArgs()
        {
        }
        public static new GetDataLimitsArgs Empty => new GetDataLimitsArgs();
    }

    public sealed class GetDataLimitsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Data Limit IDs.
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
        /// The ID of the data asset.
        /// </summary>
        [Input("parentId")]
        public Input<string>? ParentId { get; set; }

        /// <summary>
        /// The type of the service to which the data asset belongs.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        public GetDataLimitsInvokeArgs()
        {
        }
        public static new GetDataLimitsInvokeArgs Empty => new GetDataLimitsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDataLimitsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<Outputs.GetDataLimitsLimitResult> Limits;
        public readonly string? OutputFile;
        public readonly string? ParentId;
        public readonly string? ResourceType;

        [OutputConstructor]
        private GetDataLimitsResult(
            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetDataLimitsLimitResult> limits,

            string? outputFile,

            string? parentId,

            string? resourceType)
        {
            Id = id;
            Ids = ids;
            Limits = limits;
            OutputFile = outputFile;
            ParentId = parentId;
            ResourceType = resourceType;
        }
    }
}
