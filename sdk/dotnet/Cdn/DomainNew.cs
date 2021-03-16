// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cdn
{
    /// <summary>
    /// Provides a CDN Accelerated Domain resource. This resource is based on CDN's new version OpenAPI.
    /// 
    /// For information about Cdn Domain New and how to use it, see [Add a domain](https://www.alibabacloud.com/help/doc-detail/91176.html).
    /// 
    /// &gt; **NOTE:** Available in v1.34.0+.
    /// 
    /// ## Example Usage
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
    ///         // Create a new Domain.
    ///         var domain = new AliCloud.Cdn.DomainNew("domain", new AliCloud.Cdn.DomainNewArgs
    ///         {
    ///             CdnType = "web",
    ///             DomainName = "terraform.test.com",
    ///             Scope = "overseas",
    ///             Sources = 
    ///             {
    ///                 new AliCloud.Cdn.Inputs.DomainNewSourceArgs
    ///                 {
    ///                     Content = "1.1.1.1",
    ///                     Port = 80,
    ///                     Priority = 20,
    ///                     Type = "ipaddr",
    ///                     Weight = 10,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CDN domain can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cdn/domainNew:DomainNew example xxxx.com
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cdn/domainNew:DomainNew")]
    public partial class DomainNew : Pulumi.CustomResource
    {
        /// <summary>
        /// Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        /// </summary>
        [Output("cdnType")]
        public Output<string> CdnType { get; private set; } = null!;

        /// <summary>
        /// Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        /// </summary>
        [Output("certificateConfig")]
        public Output<Outputs.DomainNewCertificateConfig> CertificateConfig { get; private set; } = null!;

        /// <summary>
        /// (Available in v1.90.0+) The CNAME of the CDN domain.
        /// </summary>
        [Output("cname")]
        public Output<string> Cname { get; private set; } = null!;

        /// <summary>
        /// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// Resource group ID.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;

        /// <summary>
        /// The source address list of the accelerated domain. Defaults to null. See Block Sources.
        /// </summary>
        [Output("sources")]
        public Output<ImmutableArray<Outputs.DomainNewSource>> Sources { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DomainNew resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainNew(string name, DomainNewArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cdn/domainNew:DomainNew", name, args ?? new DomainNewArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainNew(string name, Input<string> id, DomainNewState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cdn/domainNew:DomainNew", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DomainNew resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainNew Get(string name, Input<string> id, DomainNewState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainNew(name, id, state, options);
        }
    }

    public sealed class DomainNewArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        /// </summary>
        [Input("cdnType", required: true)]
        public Input<string> CdnType { get; set; } = null!;

        /// <summary>
        /// Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        /// </summary>
        [Input("certificateConfig")]
        public Input<Inputs.DomainNewCertificateConfigArgs>? CertificateConfig { get; set; }

        /// <summary>
        /// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// Resource group ID.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("sources", required: true)]
        private InputList<Inputs.DomainNewSourceArgs>? _sources;

        /// <summary>
        /// The source address list of the accelerated domain. Defaults to null. See Block Sources.
        /// </summary>
        public InputList<Inputs.DomainNewSourceArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.DomainNewSourceArgs>());
            set => _sources = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public DomainNewArgs()
        {
        }
    }

    public sealed class DomainNewState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        /// </summary>
        [Input("cdnType")]
        public Input<string>? CdnType { get; set; }

        /// <summary>
        /// Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        /// </summary>
        [Input("certificateConfig")]
        public Input<Inputs.DomainNewCertificateConfigGetArgs>? CertificateConfig { get; set; }

        /// <summary>
        /// (Available in v1.90.0+) The CNAME of the CDN domain.
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        /// <summary>
        /// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Resource group ID.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("sources")]
        private InputList<Inputs.DomainNewSourceGetArgs>? _sources;

        /// <summary>
        /// The source address list of the accelerated domain. Defaults to null. See Block Sources.
        /// </summary>
        public InputList<Inputs.DomainNewSourceGetArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.DomainNewSourceGetArgs>());
            set => _sources = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public DomainNewState()
        {
        }
    }
}
