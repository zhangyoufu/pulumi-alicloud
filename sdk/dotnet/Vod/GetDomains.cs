// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vod
{
    public static class GetDomains
    {
        /// <summary>
        /// This data source provides the Vod Domains of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.136.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var defaultDomain = new AliCloud.Vod.Domain("defaultDomain", new AliCloud.Vod.DomainArgs
        ///         {
        ///             DomainName = "your_domain_name",
        ///             Scope = "domestic",
        ///             Sources = 
        ///             {
        ///                 new AliCloud.Vod.Inputs.DomainSourceArgs
        ///                 {
        ///                     SourceType = "domain",
        ///                     SourceContent = "your_source_content",
        ///                     SourcePort = "80",
        ///                 },
        ///             },
        ///             Tags = 
        ///             {
        ///                 { "key1", "value1" },
        ///                 { "key2", "value2" },
        ///             },
        ///         });
        ///         var defaultDomains = defaultDomain.Id.Apply(id =&gt; AliCloud.Vod.GetDomains.InvokeAsync(new AliCloud.Vod.GetDomainsArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 id,
        ///             },
        ///             Tags = 
        ///             {
        ///                 { "key1", "value1" },
        ///                 { "key2", "value2" },
        ///             },
        ///         }));
        ///         this.VodDomain = defaultDomains.Apply(defaultDomains =&gt; defaultDomains.Domains[0]);
        ///     }
        /// 
        ///     [Output("vodDomain")]
        ///     public Output&lt;string&gt; VodDomain { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDomainsResult> InvokeAsync(GetDomainsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainsResult>("alicloud:vod/getDomains:getDomains", args ?? new GetDomainsArgs(), options.WithVersion());
    }


    public sealed class GetDomainsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The search method. Valid values:
        /// * `fuzzy_match`: fuzzy match. This is the default value.
        /// * `pre_match`: prefix match.
        /// * `suf_match`: suffix match.
        /// * `full_match`: exact match
        /// </summary>
        [Input("domainSearchType")]
        public string? DomainSearchType { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Domain IDs. Its element value is same as Domain Name.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Domain name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
        /// * `Value`: It can be up to 128 characters in length. It can be a null string.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetDomainsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDomainsResult
    {
        public readonly string? DomainSearchType;
        public readonly ImmutableArray<Outputs.GetDomainsDomainResult> Domains;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? Status;
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetDomainsResult(
            string? domainSearchType,

            ImmutableArray<Outputs.GetDomainsDomainResult> domains,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? status,

            ImmutableDictionary<string, object>? tags)
        {
            DomainSearchType = domainSearchType;
            Domains = domains;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
            Tags = tags;
        }
    }
}
