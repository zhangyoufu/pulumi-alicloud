// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Wafv3
{
    public static class GetDomains
    {
        /// <summary>
        /// This data source provides the Wafv3 Domains of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available since v1.200.0.
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
        ///     var defaultInstances = AliCloud.Wafv3.GetInstances.Invoke();
        /// 
        ///     var ids = AliCloud.Wafv3.GetDomains.Invoke(new()
        ///     {
        ///         InstanceId = defaultInstances.Apply(getInstancesResult =&gt; getInstancesResult.Ids[0]),
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var defaultDomains = AliCloud.Wafv3.GetDomains.Invoke(new()
        ///     {
        ///         InstanceId = defaultInstances.Apply(getInstancesResult =&gt; getInstancesResult.Ids[0]),
        ///         Domain = "zctest12.wafqax.top",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["wafv3DomainsId1"] = ids.Apply(getDomainsResult =&gt; getDomainsResult.Domains[0]?.Id),
        ///         ["wafv3DomainsId2"] = defaultDomains.Apply(getDomainsResult =&gt; getDomainsResult.Domains[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetDomainsResult> InvokeAsync(GetDomainsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDomainsResult>("alicloud:wafv3/getDomains:getDomains", args ?? new GetDomainsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Wafv3 Domains of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available since v1.200.0.
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
        ///     var defaultInstances = AliCloud.Wafv3.GetInstances.Invoke();
        /// 
        ///     var ids = AliCloud.Wafv3.GetDomains.Invoke(new()
        ///     {
        ///         InstanceId = defaultInstances.Apply(getInstancesResult =&gt; getInstancesResult.Ids[0]),
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var defaultDomains = AliCloud.Wafv3.GetDomains.Invoke(new()
        ///     {
        ///         InstanceId = defaultInstances.Apply(getInstancesResult =&gt; getInstancesResult.Ids[0]),
        ///         Domain = "zctest12.wafqax.top",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["wafv3DomainsId1"] = ids.Apply(getDomainsResult =&gt; getDomainsResult.Domains[0]?.Id),
        ///         ["wafv3DomainsId2"] = defaultDomains.Apply(getDomainsResult =&gt; getDomainsResult.Domains[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetDomainsResult> Invoke(GetDomainsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainsResult>("alicloud:wafv3/getDomains:getDomains", args ?? new GetDomainsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The address type of the origin server. The address can be an IP address or a domain name. You can specify only one type of address.
        /// </summary>
        [Input("backend")]
        public string? Backend { get; set; }

        /// <summary>
        /// The name of the domain name to query.
        /// </summary>
        [Input("domain")]
        public string? Domain { get; set; }

        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of domain IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The WAF instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        public GetDomainsArgs()
        {
        }
        public static new GetDomainsArgs Empty => new GetDomainsArgs();
    }

    public sealed class GetDomainsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The address type of the origin server. The address can be an IP address or a domain name. You can specify only one type of address.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The name of the domain name to query.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of domain IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The WAF instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        public GetDomainsInvokeArgs()
        {
        }
        public static new GetDomainsInvokeArgs Empty => new GetDomainsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainsResult
    {
        public readonly string? Backend;
        /// <summary>
        /// The name of the domain.
        /// </summary>
        public readonly string? Domain;
        /// <summary>
        /// A list of Domain Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainsDomainResult> Domains;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string InstanceId;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;

        [OutputConstructor]
        private GetDomainsResult(
            string? backend,

            string? domain,

            ImmutableArray<Outputs.GetDomainsDomainResult> domains,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string instanceId,

            string? outputFile,

            int? pageNumber,

            int? pageSize)
        {
            Backend = backend;
            Domain = domain;
            Domains = domains;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            InstanceId = instanceId;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
        }
    }
}
