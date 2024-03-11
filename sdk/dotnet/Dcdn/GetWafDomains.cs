// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dcdn
{
    public static class GetWafDomains
    {
        /// <summary>
        /// This data source provides the Dcdn Waf Domains of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.185.0+.
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
        ///     var ids = AliCloud.Dcdn.GetWafDomains.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["dcdnWafDomainId1"] = ids.Apply(getWafDomainsResult =&gt; getWafDomainsResult.Domains[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetWafDomainsResult> InvokeAsync(GetWafDomainsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWafDomainsResult>("alicloud:dcdn/getWafDomains:getWafDomains", args ?? new GetWafDomainsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Dcdn Waf Domains of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.185.0+.
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
        ///     var ids = AliCloud.Dcdn.GetWafDomains.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["dcdnWafDomainId1"] = ids.Apply(getWafDomainsResult =&gt; getWafDomainsResult.Domains[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetWafDomainsResult> Invoke(GetWafDomainsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWafDomainsResult>("alicloud:dcdn/getWafDomains:getWafDomains", args ?? new GetWafDomainsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWafDomainsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Waf Domain IDs.
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
        /// The query conditions. You can filter domain names by name. Fuzzy match is supported `QueryArgs={"DomainName":"Accelerated domain name"}`.
        /// </summary>
        [Input("queryArgs")]
        public string? QueryArgs { get; set; }

        public GetWafDomainsArgs()
        {
        }
        public static new GetWafDomainsArgs Empty => new GetWafDomainsArgs();
    }

    public sealed class GetWafDomainsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Waf Domain IDs.
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
        /// The query conditions. You can filter domain names by name. Fuzzy match is supported `QueryArgs={"DomainName":"Accelerated domain name"}`.
        /// </summary>
        [Input("queryArgs")]
        public Input<string>? QueryArgs { get; set; }

        public GetWafDomainsInvokeArgs()
        {
        }
        public static new GetWafDomainsInvokeArgs Empty => new GetWafDomainsInvokeArgs();
    }


    [OutputType]
    public sealed class GetWafDomainsResult
    {
        public readonly ImmutableArray<Outputs.GetWafDomainsDomainResult> Domains;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly string? QueryArgs;

        [OutputConstructor]
        private GetWafDomainsResult(
            ImmutableArray<Outputs.GetWafDomainsDomainResult> domains,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            string? queryArgs)
        {
            Domains = domains;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            QueryArgs = queryArgs;
        }
    }
}
