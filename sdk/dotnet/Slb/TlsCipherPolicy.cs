// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb
{
    /// <summary>
    /// Provides a SLB Tls Cipher Policy resource.
    /// 
    /// For information about SLB Tls Cipher Policy and how to use it, see [What is Tls Cipher Policy](https://www.alibabacloud.com/help/doc-detail/196714.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.135.0+.
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
    ///         var example = new AliCloud.Slb.TlsCipherPolicy("example", new AliCloud.Slb.TlsCipherPolicyArgs
    ///         {
    ///             Ciphers = 
    ///             {
    ///                 "AES256-SHA256",
    ///                 "AES128-GCM-SHA256",
    ///             },
    ///             TlsCipherPolicyName = "Test-example_value",
    ///             TlsVersions = 
    ///             {
    ///                 "TLSv1.2",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// SLB Tls Cipher Policy can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:slb/tlsCipherPolicy:TlsCipherPolicy example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:slb/tlsCipherPolicy:TlsCipherPolicy")]
    public partial class TlsCipherPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The encryption algorithms supported. It depends on the value of `tls_versions`.
        /// </summary>
        [Output("ciphers")]
        public Output<ImmutableArray<string>> Ciphers { get; private set; } = null!;

        /// <summary>
        /// TLS policy instance state.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// TLS policy name. Length is from 2 to 128, or in both the English and Chinese characters must be with an uppercase/lowercase letter or a Chinese character and the beginning, may contain numbers, in dot `.`, underscore `_` or dash `-`.
        /// </summary>
        [Output("tlsCipherPolicyName")]
        public Output<string> TlsCipherPolicyName { get; private set; } = null!;

        /// <summary>
        /// The version of TLS protocol. You can find the corresponding value description in the document center [What is Tls Cipher Policy](https://www.alibabacloud.com/help/doc-detail/196714.htm).
        /// </summary>
        [Output("tlsVersions")]
        public Output<ImmutableArray<string>> TlsVersions { get; private set; } = null!;


        /// <summary>
        /// Create a TlsCipherPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TlsCipherPolicy(string name, TlsCipherPolicyArgs args, CustomResourceOptions? options = null)
            : base("alicloud:slb/tlsCipherPolicy:TlsCipherPolicy", name, args ?? new TlsCipherPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TlsCipherPolicy(string name, Input<string> id, TlsCipherPolicyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:slb/tlsCipherPolicy:TlsCipherPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TlsCipherPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TlsCipherPolicy Get(string name, Input<string> id, TlsCipherPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new TlsCipherPolicy(name, id, state, options);
        }
    }

    public sealed class TlsCipherPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("ciphers", required: true)]
        private InputList<string>? _ciphers;

        /// <summary>
        /// The encryption algorithms supported. It depends on the value of `tls_versions`.
        /// </summary>
        public InputList<string> Ciphers
        {
            get => _ciphers ?? (_ciphers = new InputList<string>());
            set => _ciphers = value;
        }

        /// <summary>
        /// TLS policy name. Length is from 2 to 128, or in both the English and Chinese characters must be with an uppercase/lowercase letter or a Chinese character and the beginning, may contain numbers, in dot `.`, underscore `_` or dash `-`.
        /// </summary>
        [Input("tlsCipherPolicyName", required: true)]
        public Input<string> TlsCipherPolicyName { get; set; } = null!;

        [Input("tlsVersions", required: true)]
        private InputList<string>? _tlsVersions;

        /// <summary>
        /// The version of TLS protocol. You can find the corresponding value description in the document center [What is Tls Cipher Policy](https://www.alibabacloud.com/help/doc-detail/196714.htm).
        /// </summary>
        public InputList<string> TlsVersions
        {
            get => _tlsVersions ?? (_tlsVersions = new InputList<string>());
            set => _tlsVersions = value;
        }

        public TlsCipherPolicyArgs()
        {
        }
    }

    public sealed class TlsCipherPolicyState : Pulumi.ResourceArgs
    {
        [Input("ciphers")]
        private InputList<string>? _ciphers;

        /// <summary>
        /// The encryption algorithms supported. It depends on the value of `tls_versions`.
        /// </summary>
        public InputList<string> Ciphers
        {
            get => _ciphers ?? (_ciphers = new InputList<string>());
            set => _ciphers = value;
        }

        /// <summary>
        /// TLS policy instance state.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// TLS policy name. Length is from 2 to 128, or in both the English and Chinese characters must be with an uppercase/lowercase letter or a Chinese character and the beginning, may contain numbers, in dot `.`, underscore `_` or dash `-`.
        /// </summary>
        [Input("tlsCipherPolicyName")]
        public Input<string>? TlsCipherPolicyName { get; set; }

        [Input("tlsVersions")]
        private InputList<string>? _tlsVersions;

        /// <summary>
        /// The version of TLS protocol. You can find the corresponding value description in the document center [What is Tls Cipher Policy](https://www.alibabacloud.com/help/doc-detail/196714.htm).
        /// </summary>
        public InputList<string> TlsVersions
        {
            get => _tlsVersions ?? (_tlsVersions = new InputList<string>());
            set => _tlsVersions = value;
        }

        public TlsCipherPolicyState()
        {
        }
    }
}
