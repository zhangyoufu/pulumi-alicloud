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
    /// Provides a DCDN Domain resource.
    /// 
    /// Full station accelerated domain name.
    /// 
    /// For information about DCDN Domain and how to use it, see [What is Domain](https://www.alibabacloud.com/help/en/doc-detail/130628.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.94.0.
    /// 
    /// &gt; **NOTE:** Field `force_set`, `security_token` has been removed from provider version 1.227.1.
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
    ///     var config = new Config();
    ///     var domainName = config.Get("domainName") ?? "tf-example.com";
    ///     var @default = new Random.Index.Integer("default", new()
    ///     {
    ///         Min = 10000,
    ///         Max = 99999,
    ///     });
    /// 
    ///     var example = new AliCloud.Dcdn.Domain("example", new()
    ///     {
    ///         DomainName = $"{domainName}-{@default.Result}",
    ///         Scope = "overseas",
    ///         Sources = new[]
    ///         {
    ///             new AliCloud.Dcdn.Inputs.DomainSourceArgs
    ///             {
    ///                 Content = "1.1.1.1",
    ///                 Port = 80,
    ///                 Priority = "20",
    ///                 Type = "ipaddr",
    ///                 Weight = "10",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DCDN Domain can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:dcdn/domain:Domain example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dcdn/domain:Domain")]
    public partial class Domain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
        /// </summary>
        [Output("certId")]
        public Output<string> CertId { get; private set; } = null!;

        /// <summary>
        /// The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
        /// </summary>
        [Output("certName")]
        public Output<string> CertName { get; private set; } = null!;

        /// <summary>
        /// The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
        /// </summary>
        [Output("certRegion")]
        public Output<string> CertRegion { get; private set; } = null!;

        /// <summary>
        /// The certificate type.
        /// </summary>
        [Output("certType")]
        public Output<string> CertType { get; private set; } = null!;

        /// <summary>
        /// The URL that is used for health checks.
        /// </summary>
        [Output("checkUrl")]
        public Output<string?> CheckUrl { get; private set; } = null!;

        /// <summary>
        /// The CNAME domain name corresponding to the accelerated domain name.
        /// </summary>
        [Output("cname")]
        public Output<string> Cname { get; private set; } = null!;

        /// <summary>
        /// The time when the accelerated domain name was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
        /// </summary>
        [Output("env")]
        public Output<string?> Env { get; private set; } = null!;

        /// <summary>
        /// Computing service type. Valid values:
        /// </summary>
        [Output("functionType")]
        public Output<string?> FunctionType { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The Acceleration scen. Supported:
        /// </summary>
        [Output("scene")]
        public Output<string?> Scene { get; private set; } = null!;

        /// <summary>
        /// The region where the acceleration service is deployed. Valid values:
        /// </summary>
        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        /// <summary>
        /// Source  See `sources` below.
        /// </summary>
        [Output("sources")]
        public Output<ImmutableArray<Outputs.DomainSource>> Sources { get; private set; } = null!;

        /// <summary>
        /// The private key. Specify the private key only if you want to enable the SSL certificate.
        /// </summary>
        [Output("sslPri")]
        public Output<string?> SslPri { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable the SSL certificate. Valid values:
        /// </summary>
        [Output("sslProtocol")]
        public Output<string?> SslProtocol { get; private set; } = null!;

        /// <summary>
        /// The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
        /// </summary>
        [Output("sslPub")]
        public Output<string> SslPub { get; private set; } = null!;

        /// <summary>
        /// The status of the domain name. Valid values:
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tag of the resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The top-level domain.
        /// </summary>
        [Output("topLevelDomain")]
        public Output<string?> TopLevelDomain { get; private set; } = null!;


        /// <summary>
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dcdn/domain:Domain", name, args ?? new DomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dcdn/domain:Domain", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "sslPri",
                },
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
        /// The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
        /// </summary>
        [Input("certId")]
        public Input<string>? CertId { get; set; }

        /// <summary>
        /// The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
        /// </summary>
        [Input("certName")]
        public Input<string>? CertName { get; set; }

        /// <summary>
        /// The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
        /// </summary>
        [Input("certRegion")]
        public Input<string>? CertRegion { get; set; }

        /// <summary>
        /// The certificate type.
        /// </summary>
        [Input("certType")]
        public Input<string>? CertType { get; set; }

        /// <summary>
        /// The URL that is used for health checks.
        /// </summary>
        [Input("checkUrl")]
        public Input<string>? CheckUrl { get; set; }

        /// <summary>
        /// The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
        /// </summary>
        [Input("env")]
        public Input<string>? Env { get; set; }

        /// <summary>
        /// Computing service type. Valid values:
        /// </summary>
        [Input("functionType")]
        public Input<string>? FunctionType { get; set; }

        /// <summary>
        /// The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The Acceleration scen. Supported:
        /// </summary>
        [Input("scene")]
        public Input<string>? Scene { get; set; }

        /// <summary>
        /// The region where the acceleration service is deployed. Valid values:
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("sources")]
        private InputList<Inputs.DomainSourceArgs>? _sources;

        /// <summary>
        /// Source  See `sources` below.
        /// </summary>
        public InputList<Inputs.DomainSourceArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.DomainSourceArgs>());
            set => _sources = value;
        }

        [Input("sslPri")]
        private Input<string>? _sslPri;

        /// <summary>
        /// The private key. Specify the private key only if you want to enable the SSL certificate.
        /// </summary>
        public Input<string>? SslPri
        {
            get => _sslPri;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _sslPri = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies whether to enable the SSL certificate. Valid values:
        /// </summary>
        [Input("sslProtocol")]
        public Input<string>? SslProtocol { get; set; }

        /// <summary>
        /// The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
        /// </summary>
        [Input("sslPub")]
        public Input<string>? SslPub { get; set; }

        /// <summary>
        /// The status of the domain name. Valid values:
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tag of the resource
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The top-level domain.
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
        /// The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
        /// </summary>
        [Input("certId")]
        public Input<string>? CertId { get; set; }

        /// <summary>
        /// The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
        /// </summary>
        [Input("certName")]
        public Input<string>? CertName { get; set; }

        /// <summary>
        /// The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
        /// </summary>
        [Input("certRegion")]
        public Input<string>? CertRegion { get; set; }

        /// <summary>
        /// The certificate type.
        /// </summary>
        [Input("certType")]
        public Input<string>? CertType { get; set; }

        /// <summary>
        /// The URL that is used for health checks.
        /// </summary>
        [Input("checkUrl")]
        public Input<string>? CheckUrl { get; set; }

        /// <summary>
        /// The CNAME domain name corresponding to the accelerated domain name.
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        /// <summary>
        /// The time when the accelerated domain name was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
        /// </summary>
        [Input("env")]
        public Input<string>? Env { get; set; }

        /// <summary>
        /// Computing service type. Valid values:
        /// </summary>
        [Input("functionType")]
        public Input<string>? FunctionType { get; set; }

        /// <summary>
        /// The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The Acceleration scen. Supported:
        /// </summary>
        [Input("scene")]
        public Input<string>? Scene { get; set; }

        /// <summary>
        /// The region where the acceleration service is deployed. Valid values:
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("sources")]
        private InputList<Inputs.DomainSourceGetArgs>? _sources;

        /// <summary>
        /// Source  See `sources` below.
        /// </summary>
        public InputList<Inputs.DomainSourceGetArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.DomainSourceGetArgs>());
            set => _sources = value;
        }

        [Input("sslPri")]
        private Input<string>? _sslPri;

        /// <summary>
        /// The private key. Specify the private key only if you want to enable the SSL certificate.
        /// </summary>
        public Input<string>? SslPri
        {
            get => _sslPri;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _sslPri = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies whether to enable the SSL certificate. Valid values:
        /// </summary>
        [Input("sslProtocol")]
        public Input<string>? SslProtocol { get; set; }

        /// <summary>
        /// The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
        /// </summary>
        [Input("sslPub")]
        public Input<string>? SslPub { get; set; }

        /// <summary>
        /// The status of the domain name. Valid values:
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tag of the resource
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The top-level domain.
        /// </summary>
        [Input("topLevelDomain")]
        public Input<string>? TopLevelDomain { get; set; }

        public DomainState()
        {
        }
        public static new DomainState Empty => new DomainState();
    }
}
