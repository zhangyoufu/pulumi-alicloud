// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nlb
{
    /// <summary>
    /// Provides a NLB Security Policy resource.
    /// 
    /// For information about NLB Security Policy and how to use it, see [What is Security Policy](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createsecuritypolicy-nlb).
    /// 
    /// &gt; **NOTE:** Available since v1.187.0.
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
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var defaultResourceGroups = AliCloud.ResourceManager.GetResourceGroups.Invoke();
    /// 
    ///     var defaultSecurityPolicy = new AliCloud.Nlb.SecurityPolicy("defaultSecurityPolicy", new()
    ///     {
    ///         ResourceGroupId = defaultResourceGroups.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0]),
    ///         SecurityPolicyName = name,
    ///         Ciphers = new[]
    ///         {
    ///             "ECDHE-RSA-AES128-SHA",
    ///             "ECDHE-ECDSA-AES128-SHA",
    ///         },
    ///         TlsVersions = new[]
    ///         {
    ///             "TLSv1.0",
    ///             "TLSv1.1",
    ///             "TLSv1.2",
    ///         },
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///             { "For", "example" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// NLB Security Policy can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:nlb/securityPolicy:SecurityPolicy example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:nlb/securityPolicy:SecurityPolicy")]
    public partial class SecurityPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
        /// - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
        /// - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
        /// - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`
        /// </summary>
        [Output("ciphers")]
        public Output<ImmutableArray<string>> Ciphers { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
        /// </summary>
        [Output("securityPolicyName")]
        public Output<string> SecurityPolicyName { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
        /// </summary>
        [Output("tlsVersions")]
        public Output<ImmutableArray<string>> TlsVersions { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityPolicy(string name, SecurityPolicyArgs args, CustomResourceOptions? options = null)
            : base("alicloud:nlb/securityPolicy:SecurityPolicy", name, args ?? new SecurityPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityPolicy(string name, Input<string> id, SecurityPolicyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:nlb/securityPolicy:SecurityPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityPolicy Get(string name, Input<string> id, SecurityPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityPolicy(name, id, state, options);
        }
    }

    public sealed class SecurityPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("ciphers", required: true)]
        private InputList<string>? _ciphers;

        /// <summary>
        /// The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
        /// - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
        /// - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
        /// - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`
        /// </summary>
        public InputList<string> Ciphers
        {
            get => _ciphers ?? (_ciphers = new InputList<string>());
            set => _ciphers = value;
        }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
        /// </summary>
        [Input("securityPolicyName")]
        public Input<string>? SecurityPolicyName { get; set; }

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

        [Input("tlsVersions", required: true)]
        private InputList<string>? _tlsVersions;

        /// <summary>
        /// The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
        /// </summary>
        public InputList<string> TlsVersions
        {
            get => _tlsVersions ?? (_tlsVersions = new InputList<string>());
            set => _tlsVersions = value;
        }

        public SecurityPolicyArgs()
        {
        }
        public static new SecurityPolicyArgs Empty => new SecurityPolicyArgs();
    }

    public sealed class SecurityPolicyState : global::Pulumi.ResourceArgs
    {
        [Input("ciphers")]
        private InputList<string>? _ciphers;

        /// <summary>
        /// The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
        /// - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
        /// - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
        /// - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`
        /// </summary>
        public InputList<string> Ciphers
        {
            get => _ciphers ?? (_ciphers = new InputList<string>());
            set => _ciphers = value;
        }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
        /// </summary>
        [Input("securityPolicyName")]
        public Input<string>? SecurityPolicyName { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

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

        [Input("tlsVersions")]
        private InputList<string>? _tlsVersions;

        /// <summary>
        /// The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
        /// </summary>
        public InputList<string> TlsVersions
        {
            get => _tlsVersions ?? (_tlsVersions = new InputList<string>());
            set => _tlsVersions = value;
        }

        public SecurityPolicyState()
        {
        }
        public static new SecurityPolicyState Empty => new SecurityPolicyState();
    }
}
