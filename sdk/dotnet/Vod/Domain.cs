// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vod
{
    /// <summary>
    /// Provides a VOD Domain resource.
    /// 
    /// For information about VOD Domain and how to use it, see [What is Domain](https://www.alibabacloud.com/help/product/29932.html).
    /// 
    /// &gt; **NOTE:** Available since v1.136.0+.
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
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Random.Index.Integer("default", new()
    ///     {
    ///         Min = 10000,
    ///         Max = 99999,
    ///     });
    /// 
    ///     var defaultDomain = new AliCloud.Vod.Domain("default", new()
    ///     {
    ///         DomainName = $"example-{@default.Result}.com",
    ///         Scope = "domestic",
    ///         Sources = new[]
    ///         {
    ///             new AliCloud.Vod.Inputs.DomainSourceArgs
    ///             {
    ///                 SourceType = "domain",
    ///                 SourceContent = "outin-c7405446108111ec9a7100163e0eb78b.oss-cn-beijing.aliyuncs.com",
    ///                 SourcePort = "443",
    ///             },
    ///         },
    ///         Tags = 
    ///         {
    ///             { "Created", "terraform" },
    ///             { "For", "example" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VOD Domain can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vod/domain:Domain example &lt;domain_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vod/domain:Domain")]
    public partial class Domain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
        /// </summary>
        [Output("certName")]
        public Output<string> CertName { get; private set; } = null!;

        /// <summary>
        /// The URL that is used for health checks.
        /// </summary>
        [Output("checkUrl")]
        public Output<string?> CheckUrl { get; private set; } = null!;

        /// <summary>
        /// The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
        /// </summary>
        [Output("cname")]
        public Output<string> Cname { get; private set; } = null!;

        /// <summary>
        /// The description of the domain name for CDN.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
        /// </summary>
        [Output("gmtCreated")]
        public Output<string> GmtCreated { get; private set; } = null!;

        /// <summary>
        /// The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
        /// </summary>
        [Output("gmtModified")]
        public Output<string> GmtModified { get; private set; } = null!;

        /// <summary>
        /// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
        /// </summary>
        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        /// <summary>
        /// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
        /// </summary>
        [Output("sources")]
        public Output<ImmutableArray<Outputs.DomainSource>> Sources { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
        /// </summary>
        [Output("sslProtocol")]
        public Output<string> SslProtocol { get; private set; } = null!;

        /// <summary>
        /// The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
        /// </summary>
        [Output("sslPub")]
        public Output<string> SslPub { get; private set; } = null!;

        /// <summary>
        /// The status of the domain name for CDN. Valid values:
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
        /// * `Value`: It can be up to 128 characters in length. It can be a null string.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The top-level domain name.
        /// </summary>
        [Output("topLevelDomain")]
        public Output<string?> TopLevelDomain { get; private set; } = null!;

        /// <summary>
        /// The weight of the origin server.
        /// </summary>
        [Output("weight")]
        public Output<string> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vod/domain:Domain", name, args ?? new DomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vod/domain:Domain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Domain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domain Get(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
        {
            return new Domain(name, id, state, options);
        }
    }

    public sealed class DomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL that is used for health checks.
        /// </summary>
        [Input("checkUrl")]
        public Input<string>? CheckUrl { get; set; }

        /// <summary>
        /// The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("sources", required: true)]
        private InputList<Inputs.DomainSourceArgs>? _sources;

        /// <summary>
        /// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
        /// </summary>
        public InputList<Inputs.DomainSourceArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.DomainSourceArgs>());
            set => _sources = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
        /// * `Value`: It can be up to 128 characters in length. It can be a null string.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The top-level domain name.
        /// </summary>
        [Input("topLevelDomain")]
        public Input<string>? TopLevelDomain { get; set; }

        public DomainArgs()
        {
        }
        public static new DomainArgs Empty => new DomainArgs();
    }

    public sealed class DomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
        /// </summary>
        [Input("certName")]
        public Input<string>? CertName { get; set; }

        /// <summary>
        /// The URL that is used for health checks.
        /// </summary>
        [Input("checkUrl")]
        public Input<string>? CheckUrl { get; set; }

        /// <summary>
        /// The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        /// <summary>
        /// The description of the domain name for CDN.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
        /// </summary>
        [Input("gmtCreated")]
        public Input<string>? GmtCreated { get; set; }

        /// <summary>
        /// The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
        /// </summary>
        [Input("gmtModified")]
        public Input<string>? GmtModified { get; set; }

        /// <summary>
        /// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("sources")]
        private InputList<Inputs.DomainSourceGetArgs>? _sources;

        /// <summary>
        /// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
        /// </summary>
        public InputList<Inputs.DomainSourceGetArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.DomainSourceGetArgs>());
            set => _sources = value;
        }

        /// <summary>
        /// Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
        /// </summary>
        [Input("sslProtocol")]
        public Input<string>? SslProtocol { get; set; }

        /// <summary>
        /// The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
        /// </summary>
        [Input("sslPub")]
        public Input<string>? SslPub { get; set; }

        /// <summary>
        /// The status of the domain name for CDN. Valid values:
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
        /// * `Value`: It can be up to 128 characters in length. It can be a null string.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The top-level domain name.
        /// </summary>
        [Input("topLevelDomain")]
        public Input<string>? TopLevelDomain { get; set; }

        /// <summary>
        /// The weight of the origin server.
        /// </summary>
        [Input("weight")]
        public Input<string>? Weight { get; set; }

        public DomainState()
        {
        }
        public static new DomainState Empty => new DomainState();
    }
}
