// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Pvtz
{
    public static class GetZones
    {
        /// <summary>
        /// This data source lists a number of Private Zones resource information owned by an Alibaba Cloud account.
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
        ///     var pvtzZonesDs = AliCloud.Pvtz.GetZones.Invoke(new()
        ///     {
        ///         Keyword = basic.ZoneName,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstZoneId"] = pvtzZonesDs.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetZonesResult> InvokeAsync(GetZonesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetZonesResult>("alicloud:pvtz/getZones:getZones", args ?? new GetZonesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source lists a number of Private Zones resource information owned by an Alibaba Cloud account.
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
        ///     var pvtzZonesDs = AliCloud.Pvtz.GetZones.Invoke(new()
        ///     {
        ///         Keyword = basic.ZoneName,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstZoneId"] = pvtzZonesDs.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetZonesResult> Invoke(GetZonesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetZonesResult>("alicloud:pvtz/getZones:getZones", args ?? new GetZonesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetZonesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to true can output more details.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of zone IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// keyword for zone name.
        /// </summary>
        [Input("keyword")]
        public string? Keyword { get; set; }

        /// <summary>
        /// User language.
        /// </summary>
        [Input("lang")]
        public string? Lang { get; set; }

        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// query_region_id for zone regionId.
        /// </summary>
        [Input("queryRegionId")]
        public string? QueryRegionId { get; set; }

        /// <summary>
        /// query_vpc_id for zone vpcId.
        /// </summary>
        [Input("queryVpcId")]
        public string? QueryVpcId { get; set; }

        /// <summary>
        /// resource_group_id for zone resourceGroupId.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        /// <summary>
        /// Search mode. Value: 
        /// - LIKE: fuzzy search.
        /// - EXACT: precise search. It is not filled in by default.
        /// </summary>
        [Input("searchMode")]
        public string? SearchMode { get; set; }

        public GetZonesArgs()
        {
        }
        public static new GetZonesArgs Empty => new GetZonesArgs();
    }

    public sealed class GetZonesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to true can output more details.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of zone IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// keyword for zone name.
        /// </summary>
        [Input("keyword")]
        public Input<string>? Keyword { get; set; }

        /// <summary>
        /// User language.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// query_region_id for zone regionId.
        /// </summary>
        [Input("queryRegionId")]
        public Input<string>? QueryRegionId { get; set; }

        /// <summary>
        /// query_vpc_id for zone vpcId.
        /// </summary>
        [Input("queryVpcId")]
        public Input<string>? QueryVpcId { get; set; }

        /// <summary>
        /// resource_group_id for zone resourceGroupId.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Search mode. Value: 
        /// - LIKE: fuzzy search.
        /// - EXACT: precise search. It is not filled in by default.
        /// </summary>
        [Input("searchMode")]
        public Input<string>? SearchMode { get; set; }

        public GetZonesInvokeArgs()
        {
        }
        public static new GetZonesInvokeArgs Empty => new GetZonesInvokeArgs();
    }


    [OutputType]
    public sealed class GetZonesResult
    {
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of zone IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? Keyword;
        public readonly string? Lang;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of zone names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? QueryRegionId;
        public readonly string? QueryVpcId;
        /// <summary>
        /// The Id of resource group which the Private Zone belongs.
        /// </summary>
        public readonly string? ResourceGroupId;
        public readonly string? SearchMode;
        /// <summary>
        /// A list of zones. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZonesZoneResult> Zones;

        [OutputConstructor]
        private GetZonesResult(
            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? keyword,

            string? lang,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? queryRegionId,

            string? queryVpcId,

            string? resourceGroupId,

            string? searchMode,

            ImmutableArray<Outputs.GetZonesZoneResult> zones)
        {
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            Keyword = keyword;
            Lang = lang;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            QueryRegionId = queryRegionId;
            QueryVpcId = queryVpcId;
            ResourceGroupId = resourceGroupId;
            SearchMode = searchMode;
            Zones = zones;
        }
    }
}
