// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dcdn
{
    /// <summary>
    /// Provides a DCDN Ipa Domain resource.
    /// 
    /// For information about DCDN Ipa Domain and how to use it, see [What is Ipa Domain](https://www.alibabacloud.com/help/en/doc-detail/130634.html).
    /// 
    /// &gt; **NOTE:** Available since v1.158.0.
    /// 
    /// ## Example Usage
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
    ///     var domainName = config.Get("domainName") ?? "example.com";
    ///     var @default = AliCloud.ResourceManager.GetResourceGroups.Invoke();
    /// 
    ///     var example = new AliCloud.Dcdn.IpaDomain("example", new()
    ///     {
    ///         DomainName = domainName,
    ///         ResourceGroupId = @default.Apply(@default =&gt; @default.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Groups[0]?.Id)),
    ///         Scope = "global",
    ///         Status = "online",
    ///         Sources = new[]
    ///         {
    ///             new AliCloud.Dcdn.Inputs.IpaDomainSourceArgs
    ///             {
    ///                 Content = "www.alicloud-provider.cn",
    ///                 Port = 80,
    ///                 Priority = "20",
    ///                 Type = "domain",
    ///                 Weight = 10,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DCDN Ipa Domain can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:dcdn/ipaDomain:IpaDomain example &lt;domain_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dcdn/ipaDomain:IpaDomain")]
    public partial class IpaDomain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The accelerated region. Valid values: `domestic`, `global`, `overseas`.
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;

        /// <summary>
        /// Sources. See `sources` below.
        /// </summary>
        [Output("sources")]
        public Output<ImmutableArray<Outputs.IpaDomainSource>> Sources { get; private set; } = null!;

        /// <summary>
        /// The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a IpaDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpaDomain(string name, IpaDomainArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dcdn/ipaDomain:IpaDomain", name, args ?? new IpaDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpaDomain(string name, Input<string> id, IpaDomainState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dcdn/ipaDomain:IpaDomain", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IpaDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpaDomain Get(string name, Input<string> id, IpaDomainState? state = null, CustomResourceOptions? options = null)
        {
            return new IpaDomain(name, id, state, options);
        }
    }

    public sealed class IpaDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The accelerated region. Valid values: `domestic`, `global`, `overseas`.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("sources", required: true)]
        private InputList<Inputs.IpaDomainSourceArgs>? _sources;

        /// <summary>
        /// Sources. See `sources` below.
        /// </summary>
        public InputList<Inputs.IpaDomainSourceArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.IpaDomainSourceArgs>());
            set => _sources = value;
        }

        /// <summary>
        /// The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public IpaDomainArgs()
        {
        }
        public static new IpaDomainArgs Empty => new IpaDomainArgs();
    }

    public sealed class IpaDomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The accelerated region. Valid values: `domestic`, `global`, `overseas`.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("sources")]
        private InputList<Inputs.IpaDomainSourceGetArgs>? _sources;

        /// <summary>
        /// Sources. See `sources` below.
        /// </summary>
        public InputList<Inputs.IpaDomainSourceGetArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.IpaDomainSourceGetArgs>());
            set => _sources = value;
        }

        /// <summary>
        /// The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public IpaDomainState()
        {
        }
        public static new IpaDomainState Empty => new IpaDomainState();
    }
}
