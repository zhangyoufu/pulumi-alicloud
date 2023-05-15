// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dcdn
{
    public static class GetIpaDomains
    {
        /// <summary>
        /// This data source provides the Dcdn Ipa Domains of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.158.0+.
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
        ///     var ids = AliCloud.Dcdn.GetIpaDomains.Invoke(new()
        ///     {
        ///         DomainName = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///     });
        /// 
        ///     var status = AliCloud.Dcdn.GetIpaDomains.Invoke(new()
        ///     {
        ///         Status = "online",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["dcdnIpaDomainId1"] = ids.Apply(getIpaDomainsResult =&gt; getIpaDomainsResult.Domains[0]?.Id),
        ///         ["dcdnIpaDomainId2"] = status.Apply(getIpaDomainsResult =&gt; getIpaDomainsResult.Domains[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIpaDomainsResult> InvokeAsync(GetIpaDomainsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpaDomainsResult>("alicloud:dcdn/getIpaDomains:getIpaDomains", args ?? new GetIpaDomainsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Dcdn Ipa Domains of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.158.0+.
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
        ///     var ids = AliCloud.Dcdn.GetIpaDomains.Invoke(new()
        ///     {
        ///         DomainName = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///     });
        /// 
        ///     var status = AliCloud.Dcdn.GetIpaDomains.Invoke(new()
        ///     {
        ///         Status = "online",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["dcdnIpaDomainId1"] = ids.Apply(getIpaDomainsResult =&gt; getIpaDomainsResult.Domains[0]?.Id),
        ///         ["dcdnIpaDomainId2"] = status.Apply(getIpaDomainsResult =&gt; getIpaDomainsResult.Domains[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIpaDomainsResult> Invoke(GetIpaDomainsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpaDomainsResult>("alicloud:dcdn/getIpaDomains:getIpaDomains", args ?? new GetIpaDomainsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpaDomainsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The accelerated domain names.
        /// </summary>
        [Input("domainName")]
        public string? DomainName { get; set; }

        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Ipa Domain IDs.
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
        /// The status of the accelerated domain name.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetIpaDomainsArgs()
        {
        }
        public static new GetIpaDomainsArgs Empty => new GetIpaDomainsArgs();
    }

    public sealed class GetIpaDomainsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The accelerated domain names.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Ipa Domain IDs.
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
        /// The status of the accelerated domain name.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetIpaDomainsInvokeArgs()
        {
        }
        public static new GetIpaDomainsInvokeArgs Empty => new GetIpaDomainsInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpaDomainsResult
    {
        public readonly string? DomainName;
        public readonly ImmutableArray<Outputs.GetIpaDomainsDomainResult> Domains;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetIpaDomainsResult(
            string? domainName,

            ImmutableArray<Outputs.GetIpaDomainsDomainResult> domains,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<string> names,

            string? outputFile,

            string? status)
        {
            DomainName = domainName;
            Domains = domains;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            Names = names;
            OutputFile = outputFile;
            Status = status;
        }
    }
}
