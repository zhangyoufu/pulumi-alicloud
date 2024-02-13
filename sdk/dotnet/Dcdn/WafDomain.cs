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
    /// Provides a DCDN Waf Domain resource.
    /// 
    /// For information about DCDN Waf Domain and how to use it, see [What is Waf Domain](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-batchsetdcdnwafdomainconfigs).
    /// 
    /// &gt; **NOTE:** Available since v1.185.0.
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
    ///     var @default = new Random.RandomInteger("default", new()
    ///     {
    ///         Min = 10000,
    ///         Max = 99999,
    ///     });
    /// 
    ///     var exampleDomain = new AliCloud.Dcdn.Domain("exampleDomain", new()
    ///     {
    ///         DomainName = @default.Result.Apply(result =&gt; $"{domainName}-{result}"),
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
    ///     var exampleWafDomain = new AliCloud.Dcdn.WafDomain("exampleWafDomain", new()
    ///     {
    ///         DomainName = exampleDomain.DomainName,
    ///         ClientIpTag = "X-Forwarded-For",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DCDN Waf Domain can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:dcdn/wafDomain:WafDomain example &lt;domain_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dcdn/wafDomain:WafDomain")]
    public partial class WafDomain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The client ip tag.
        /// </summary>
        [Output("clientIpTag")]
        public Output<string?> ClientIpTag { get; private set; } = null!;

        /// <summary>
        /// The accelerated domain name.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;


        /// <summary>
        /// Create a WafDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WafDomain(string name, WafDomainArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dcdn/wafDomain:WafDomain", name, args ?? new WafDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WafDomain(string name, Input<string> id, WafDomainState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dcdn/wafDomain:WafDomain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WafDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WafDomain Get(string name, Input<string> id, WafDomainState? state = null, CustomResourceOptions? options = null)
        {
            return new WafDomain(name, id, state, options);
        }
    }

    public sealed class WafDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client ip tag.
        /// </summary>
        [Input("clientIpTag")]
        public Input<string>? ClientIpTag { get; set; }

        /// <summary>
        /// The accelerated domain name.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        public WafDomainArgs()
        {
        }
        public static new WafDomainArgs Empty => new WafDomainArgs();
    }

    public sealed class WafDomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client ip tag.
        /// </summary>
        [Input("clientIpTag")]
        public Input<string>? ClientIpTag { get; set; }

        /// <summary>
        /// The accelerated domain name.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        public WafDomainState()
        {
        }
        public static new WafDomainState Empty => new WafDomainState();
    }
}
