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
    /// A Load Balancer CA Certificate is used by the listener of the protocol https.
    /// 
    /// For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).
    /// 
    /// For information about CA Certificate and how to use it, see [Configure CA Certificate](https://www.alibabacloud.com/help/doc-detail/85968.htm).
    /// 
    /// ## Example Usage
    /// 
    /// * using CA certificate content
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var foo = new AliCloud.Slb.CaCertificate("foo", new()
    ///     {
    ///         Certificate = @"-----BEGIN CERTIFICATE-----
    /// MIIDRjCCAq+gAwIBAgIJAJn3ox4K13PoMA0GCSqGSIb3DQEBBQUAMHYxCzAJBgNV
    /// BAYTAkNOMQswCQYDVQQIEwJCSjELMAkGA1UEBxMCQkoxDDAKBgNVBAoTA0FMSTEP
    /// MA0GA1UECxMGQUxJWVVOMQ0wCwYDVQQDEwR0ZXN0MR8wHQYJKoZIhvcNAQkBFhB0
    /// ZXN0QGhvdG1haWwuY29tMB4XDTE0MTEyNDA2MDQyNVoXDTI0MTEyMTA2MDQyNVow
    /// djELMAkGA1UEBhMCQ04xCzAJBgNVBAgTAkJKMQswCQYDVQQHEwJCSjEMMAoGA1UE
    /// ChMDQUxJMQ8wDQYDVQQLEwZBTElZVU4xDTALBgNVBAMTBHRlc3QxHzAdBgkqhkiG
    /// 9w0BCQEWEHRlc3RAaG90bWFpbC5jb20wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJ
    /// AoGBAM7SS3e9+Nj0HKAsRuIDNSsS3UK6b+62YQb2uuhKrp1HMrOx61WSDR2qkAnB
    /// coG00Uz38EE+9DLYNUVQBK7aSgLP5M1Ak4wr4GqGyCgjejzzh3DshUzLCCy2rook
    /// KOyRTlPX+Q5l7rE1fcSNzgepcae5i2sE1XXXzLRIDIvQxcspAgMBAAGjgdswgdgw
    /// HQYDVR0OBBYEFBdy+OuMsvbkV7R14f0OyoLoh2z4MIGoBgNVHSMEgaAwgZ2AFBdy
    /// +OuMsvbkV7R14f0OyoLoh2z4oXqkeDB2MQswCQYDVQQGEwJDTjELMAkGA1UECBMC
    /// QkoxCzAJBgNVBAcTAkJKMQwwCgYDVQQKEwNBTEkxDzANBgNVBAsTBkFMSVlVTjEN
    /// MAsGA1UEAxMEdGVzdDEfMB0GCSqGSIb3DQEJARYQdGVzdEBob3RtYWlsLmNvbYIJ
    /// AJn3ox4K13PoMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEFBQADgYEAY7KOsnyT
    /// cQzfhiiG7ASjiPakw5wXoycHt5GCvLG5htp2TKVzgv9QTliA3gtfv6oV4zRZx7X1
    /// Ofi6hVgErtHaXJheuPVeW6eAW8mHBoEfvDAfU3y9waYrtUevSl07643bzKL6v+Qd
    /// DUBTxOAvSYfXTtI90EAxEG/bJJyOm5LqoiA=
    /// -----END CERTIFICATE-----
    /// ",
    ///         CaCertificateName = "tf-testAccSlbCACertificate",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// * using CA certificate file
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var foo_file = new AliCloud.Slb.CaCertificate("foo-file", new()
    ///     {
    ///         CaCertificateName = "tf-testAccSlbCACertificate",
    ///         Certificate = File.ReadAllText($"{path.Module}/ca_certificate.pem"),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Server Load balancer CA Certificate can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:slb/caCertificate:CaCertificate example abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:slb/caCertificate:CaCertificate")]
    public partial class CaCertificate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// the content of the CA certificate.
        /// </summary>
        [Output("caCertificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// Name of the CA Certificate.
        /// </summary>
        [Output("caCertificateName")]
        public Output<string> CaCertificateName { get; private set; } = null!;

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.123.1. New field `ca_certificate_name` instead
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Id of resource group which the slb_ca certificate belongs.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a CaCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CaCertificate(string name, CaCertificateArgs args, CustomResourceOptions? options = null)
            : base("alicloud:slb/caCertificate:CaCertificate", name, args ?? new CaCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CaCertificate(string name, Input<string> id, CaCertificateState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:slb/caCertificate:CaCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CaCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CaCertificate Get(string name, Input<string> id, CaCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new CaCertificate(name, id, state, options);
        }
    }

    public sealed class CaCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the content of the CA certificate.
        /// </summary>
        [Input("caCertificate", required: true)]
        public Input<string> Certificate { get; set; } = null!;

        /// <summary>
        /// Name of the CA Certificate.
        /// </summary>
        [Input("caCertificateName")]
        public Input<string>? CaCertificateName { get; set; }

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.123.1. New field `ca_certificate_name` instead
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Id of resource group which the slb_ca certificate belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

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

        public CaCertificateArgs()
        {
        }
        public static new CaCertificateArgs Empty => new CaCertificateArgs();
    }

    public sealed class CaCertificateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the content of the CA certificate.
        /// </summary>
        [Input("caCertificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// Name of the CA Certificate.
        /// </summary>
        [Input("caCertificateName")]
        public Input<string>? CaCertificateName { get; set; }

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.123.1. New field `ca_certificate_name` instead
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Id of resource group which the slb_ca certificate belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

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

        public CaCertificateState()
        {
        }
        public static new CaCertificateState Empty => new CaCertificateState();
    }
}
