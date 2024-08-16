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
    /// A Load Balancer Server Certificate is an ssl Certificate used by the listener of the protocol https.
    /// 
    /// For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).
    /// 
    /// For information about Server Certificate and how to use it, see [Configure Server Certificate](https://www.alibabacloud.com/help/doc-detail/85968.htm).
    /// 
    /// ## Example Usage
    /// 
    /// * using server_certificate/private content as string example
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // create a server certificate
    ///     var foo = new AliCloud.Slb.ServerCertificate("foo", new()
    ///     {
    ///         Name = "slbservercertificate",
    ///         Certificate = @"-----BEGIN CERTIFICATE-----
    /// MIICWDCCAcGgAwIBAgIJAP7vOtjPtQIjMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV
    /// BAYTAkNOMRMwEQYDVQQIDApjbi1iZWlqaW5nMSEwHwYDVQQKDBhJbnRlcm5ldCBX
    /// aWRnaXRzIFB0eSBMdGQwHhcNMjAxMDIwMDYxOTUxWhcNMjAxMTE5MDYxOTUxWjBF
    /// MQswCQYDVQQGEwJDTjETMBEGA1UECAwKY24tYmVpamluZzEhMB8GA1UECgwYSW50
    /// ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKB
    /// gQDEdoyaJ0kdtjtbLRx5X9qwI7FblhJPRcScvhQSE8P5y/b/T8J9BVuFIBoU8nrP
    /// Y9ABz4JFklZ6SznxLbFBqtXoJTmzV6ixyjjH+AGEw6hCiA8Pqy2CNIzxr9DjCzN5
    /// tWruiHqO60O3Bve6cHipH0VyLAhrB85mflvOZSH4xGsJkwIDAQABo1AwTjAdBgNV
    /// HQ4EFgQUYDwuuqC2a2UPrfm1v31vE7+GRM4wHwYDVR0jBBgwFoAUYDwuuqC2a2UP
    /// rfm1v31vE7+GRM4wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOBgQAovSB0
    /// 5JRKrg7lYR/KlTuKHmozfyL9UER0/dpTSoqsCyt8yc1BbtAKUJWh09BujBE1H22f
    /// lKvCAjhPmnNdfd/l9GrmAWNDWEDPLdUTkGSkKAScMpdS+mLmOBuYWgdnOtq3eQGf
    /// t07tlBL+dtzrrohHpfLeuNyYb40g8VQdp3RRRQ==
    /// -----END CERTIFICATE-----",
    ///         PrivateKey = @"-----BEGIN RSA PRIVATE KEY-----
    /// MIICXAIBAAKBgQDEdoyaJ0kdtjtbLRx5X9qwI7FblhJPRcScvhQSE8P5y/b/T8J9
    /// BVuFIBoU8nrPY9ABz4JFklZ6SznxLbFBqtXoJTmzV6ixyjjH+AGEw6hCiA8Pqy2C
    /// NIzxr9DjCzN5tWruiHqO60O3Bve6cHipH0VyLAhrB85mflvOZSH4xGsJkwIDAQAB
    /// AoGARe2oaCo5lTDK+c4Zx3392hoqQ94r0DmWHPBvNmwAooYd+YxLPrLMe5sMjY4t
    /// dmohnLNevCK1Uzw5eIX6BNSo5CORBcIDRmiAgwiYiS3WOv2+qi9g5uIdMiDr+EED
    /// K8wZJjB5E2WyfxL507vtW4T5L36yfr8SkmqH3GvzpI2jCqECQQDsy0AmBzyfK0tG
    /// Nw1+iF9SReJWgb1f5iHvz+6Dt5ueVQngrl/5++Gp5bNoaQMkLEDsy0iHIj9j43ji
    /// 0DON05uDAkEA1GXgGn8MXXKyuzYuoyYXCBH7aF579d7KEGET/jjnXx9DHcfRJZBY
    /// B9ghMnnonSOGboF04Zsdd3xwYF/3OHYssQJAekd/SeQEzyE5TvoQ8t2Tc9X4yrlW
    /// xNX/gmp6/fPr3biGUEtb7qi+4NBodCt+XsingmB7hKUP3RJTk7T2WnAC5wJAMqHi
    /// jY5x3SkFkHl3Hq9q2CKpQxUbCd7FXqg1wum/xj5GmqfSpNjHE3+jUkwbdrJMTrWP
    /// rmRy3tQMWf0mixAo0QJBAN4IcZChanq8cZyNqqoNbxGm4hkxUmE0W4hxHmLC2CYZ
    /// V4JpNm8dpi4CiMWLasF6TYlVMgX+aPxYRUWc/qqf1/Q=
    /// -----END RSA PRIVATE KEY-----",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// * using server_certificate/private file example
    /// 
    /// ## Import
    /// 
    /// Server Load balancer Server Certificate can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:slb/serverCertificate:ServerCertificate example abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:slb/serverCertificate:ServerCertificate")]
    public partial class ServerCertificate : global::Pulumi.CustomResource
    {
        [Output("alicloudCertifacteId")]
        public Output<string?> AlicloudCertifacteId { get; private set; } = null!;

        [Output("alicloudCertifacteName")]
        public Output<string?> AlicloudCertifacteName { get; private set; } = null!;

        /// <summary>
        /// an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
        /// </summary>
        [Output("alicloudCertificateId")]
        public Output<string?> AlicloudCertificateId { get; private set; } = null!;

        /// <summary>
        /// the name of the certificate specified by `alicloud_certificate_id`.but it is not supported on the international site of alibaba cloud now.
        /// </summary>
        [Output("alicloudCertificateName")]
        public Output<string?> AlicloudCertificateName { get; private set; } = null!;

        /// <summary>
        /// the region of the certificate specified by `alicloud_certificate_id`. but it is not supported on the international site of alibaba cloud now.
        /// </summary>
        [Output("alicloudCertificateRegionId")]
        public Output<string?> AlicloudCertificateRegionId { get; private set; } = null!;

        /// <summary>
        /// Name of the Server Certificate.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// the content of privat key of the ssl certificate specified by `server_certificate`. where `alicloud_certificate_id` is null, it is required, otherwise it is ignored.
        /// </summary>
        [Output("privateKey")]
        public Output<string?> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// The Id of resource group which the slb server certificate belongs.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// the content of the ssl certificate. where `alicloud_certificate_id` is null, it is required, otherwise it is ignored.
        /// </summary>
        [Output("serverCertificate")]
        public Output<string?> Certificate { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ServerCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerCertificate(string name, ServerCertificateArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:slb/serverCertificate:ServerCertificate", name, args ?? new ServerCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerCertificate(string name, Input<string> id, ServerCertificateState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:slb/serverCertificate:ServerCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerCertificate Get(string name, Input<string> id, ServerCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerCertificate(name, id, state, options);
        }
    }

    public sealed class ServerCertificateArgs : global::Pulumi.ResourceArgs
    {
        [Input("alicloudCertifacteId")]
        public Input<string>? AlicloudCertifacteId { get; set; }

        [Input("alicloudCertifacteName")]
        public Input<string>? AlicloudCertifacteName { get; set; }

        /// <summary>
        /// an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
        /// </summary>
        [Input("alicloudCertificateId")]
        public Input<string>? AlicloudCertificateId { get; set; }

        /// <summary>
        /// the name of the certificate specified by `alicloud_certificate_id`.but it is not supported on the international site of alibaba cloud now.
        /// </summary>
        [Input("alicloudCertificateName")]
        public Input<string>? AlicloudCertificateName { get; set; }

        /// <summary>
        /// the region of the certificate specified by `alicloud_certificate_id`. but it is not supported on the international site of alibaba cloud now.
        /// </summary>
        [Input("alicloudCertificateRegionId")]
        public Input<string>? AlicloudCertificateRegionId { get; set; }

        /// <summary>
        /// Name of the Server Certificate.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// the content of privat key of the ssl certificate specified by `server_certificate`. where `alicloud_certificate_id` is null, it is required, otherwise it is ignored.
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        /// <summary>
        /// The Id of resource group which the slb server certificate belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// the content of the ssl certificate. where `alicloud_certificate_id` is null, it is required, otherwise it is ignored.
        /// </summary>
        [Input("serverCertificate")]
        public Input<string>? Certificate { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServerCertificateArgs()
        {
        }
        public static new ServerCertificateArgs Empty => new ServerCertificateArgs();
    }

    public sealed class ServerCertificateState : global::Pulumi.ResourceArgs
    {
        [Input("alicloudCertifacteId")]
        public Input<string>? AlicloudCertifacteId { get; set; }

        [Input("alicloudCertifacteName")]
        public Input<string>? AlicloudCertifacteName { get; set; }

        /// <summary>
        /// an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
        /// </summary>
        [Input("alicloudCertificateId")]
        public Input<string>? AlicloudCertificateId { get; set; }

        /// <summary>
        /// the name of the certificate specified by `alicloud_certificate_id`.but it is not supported on the international site of alibaba cloud now.
        /// </summary>
        [Input("alicloudCertificateName")]
        public Input<string>? AlicloudCertificateName { get; set; }

        /// <summary>
        /// the region of the certificate specified by `alicloud_certificate_id`. but it is not supported on the international site of alibaba cloud now.
        /// </summary>
        [Input("alicloudCertificateRegionId")]
        public Input<string>? AlicloudCertificateRegionId { get; set; }

        /// <summary>
        /// Name of the Server Certificate.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// the content of privat key of the ssl certificate specified by `server_certificate`. where `alicloud_certificate_id` is null, it is required, otherwise it is ignored.
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        /// <summary>
        /// The Id of resource group which the slb server certificate belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// the content of the ssl certificate. where `alicloud_certificate_id` is null, it is required, otherwise it is ignored.
        /// </summary>
        [Input("serverCertificate")]
        public Input<string>? Certificate { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServerCertificateState()
        {
        }
        public static new ServerCertificateState Empty => new ServerCertificateState();
    }
}
