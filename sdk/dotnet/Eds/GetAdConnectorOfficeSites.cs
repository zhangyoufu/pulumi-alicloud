// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eds
{
    public static class GetAdConnectorOfficeSites
    {
        /// <summary>
        /// This data source provides the Ecd Ad Connector Office Sites of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.176.0+.
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
        ///     var ids = AliCloud.Eds.GetAdConnectorOfficeSites.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Eds.GetAdConnectorOfficeSites.Invoke(new()
        ///     {
        ///         NameRegex = "^my-AdConnectorOfficeSite",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ecdAdConnectorOfficeSiteId1"] = ids.Apply(getAdConnectorOfficeSitesResult =&gt; getAdConnectorOfficeSitesResult.Sites[0]?.Id),
        ///         ["ecdAdConnectorOfficeSiteId2"] = nameRegex.Apply(getAdConnectorOfficeSitesResult =&gt; getAdConnectorOfficeSitesResult.Sites[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetAdConnectorOfficeSitesResult> InvokeAsync(GetAdConnectorOfficeSitesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAdConnectorOfficeSitesResult>("alicloud:eds/getAdConnectorOfficeSites:getAdConnectorOfficeSites", args ?? new GetAdConnectorOfficeSitesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Ecd Ad Connector Office Sites of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.176.0+.
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
        ///     var ids = AliCloud.Eds.GetAdConnectorOfficeSites.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Eds.GetAdConnectorOfficeSites.Invoke(new()
        ///     {
        ///         NameRegex = "^my-AdConnectorOfficeSite",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ecdAdConnectorOfficeSiteId1"] = ids.Apply(getAdConnectorOfficeSitesResult =&gt; getAdConnectorOfficeSitesResult.Sites[0]?.Id),
        ///         ["ecdAdConnectorOfficeSiteId2"] = nameRegex.Apply(getAdConnectorOfficeSitesResult =&gt; getAdConnectorOfficeSitesResult.Sites[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetAdConnectorOfficeSitesResult> Invoke(GetAdConnectorOfficeSitesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAdConnectorOfficeSitesResult>("alicloud:eds/getAdConnectorOfficeSites:getAdConnectorOfficeSites", args ?? new GetAdConnectorOfficeSitesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAdConnectorOfficeSitesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Ad Connector Office Site IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Ad Connector Office Site name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The workspace status.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetAdConnectorOfficeSitesArgs()
        {
        }
        public static new GetAdConnectorOfficeSitesArgs Empty => new GetAdConnectorOfficeSitesArgs();
    }

    public sealed class GetAdConnectorOfficeSitesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Ad Connector Office Site IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Ad Connector Office Site name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The workspace status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetAdConnectorOfficeSitesInvokeArgs()
        {
        }
        public static new GetAdConnectorOfficeSitesInvokeArgs Empty => new GetAdConnectorOfficeSitesInvokeArgs();
    }


    [OutputType]
    public sealed class GetAdConnectorOfficeSitesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetAdConnectorOfficeSitesSiteResult> Sites;
        public readonly string? Status;

        [OutputConstructor]
        private GetAdConnectorOfficeSitesResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetAdConnectorOfficeSitesSiteResult> sites,

            string? status)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Sites = sites;
            Status = status;
        }
    }
}
