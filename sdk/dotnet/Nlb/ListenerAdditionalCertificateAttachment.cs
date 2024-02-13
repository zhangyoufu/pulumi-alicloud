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
    /// Provides a NLB Listener Additional Certificate Attachment resource.
    /// 
    /// For information about NLB Listener Additional Certificate Attachment and how to use it, see [What is Listener Additional Certificate Attachment](https://www.alibabacloud.com/help/en/server-load-balancer/latest/nlb-instances-change).
    /// 
    /// &gt; **NOTE:** Available since v1.209.0.
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
    ///     var defaultZones = AliCloud.Nlb.GetZones.Invoke();
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.4.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "10.4.0.0/24",
    ///         VpcId = defaultNetwork.Id,
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var default1 = new AliCloud.Vpc.Switch("default1", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "10.4.1.0/24",
    ///         VpcId = defaultNetwork.Id,
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[1]?.Id),
    ///     });
    /// 
    ///     var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("defaultSecurityGroup", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///     });
    /// 
    ///     var defaultLoadBalancer = new AliCloud.Nlb.LoadBalancer("defaultLoadBalancer", new()
    ///     {
    ///         LoadBalancerName = name,
    ///         ResourceGroupId = defaultResourceGroups.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0]),
    ///         LoadBalancerType = "Network",
    ///         AddressType = "Internet",
    ///         AddressIpVersion = "Ipv4",
    ///         VpcId = defaultNetwork.Id,
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///             { "For", "example" },
    ///         },
    ///         ZoneMappings = new[]
    ///         {
    ///             new AliCloud.Nlb.Inputs.LoadBalancerZoneMappingArgs
    ///             {
    ///                 VswitchId = defaultSwitch.Id,
    ///                 ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///             },
    ///             new AliCloud.Nlb.Inputs.LoadBalancerZoneMappingArgs
    ///             {
    ///                 VswitchId = default1.Id,
    ///                 ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[1]?.Id),
    ///             },
    ///         },
    ///     });
    /// 
    ///     var defaultServerGroup = new AliCloud.Nlb.ServerGroup("defaultServerGroup", new()
    ///     {
    ///         ResourceGroupId = defaultResourceGroups.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0]),
    ///         ServerGroupName = name,
    ///         ServerGroupType = "Instance",
    ///         VpcId = defaultNetwork.Id,
    ///         Scheduler = "Wrr",
    ///         Protocol = "TCPSSL",
    ///         ConnectionDrain = true,
    ///         ConnectionDrainTimeout = 60,
    ///         AddressIpVersion = "Ipv4",
    ///         HealthCheck = new AliCloud.Nlb.Inputs.ServerGroupHealthCheckArgs
    ///         {
    ///             HealthCheckUrl = "/adc/index.html",
    ///             HealthCheckDomain = "tf-iac.com",
    ///             HealthCheckEnabled = true,
    ///             HealthCheckType = "TCP",
    ///             HealthCheckConnectPort = 0,
    ///             HealthyThreshold = 2,
    ///             UnhealthyThreshold = 2,
    ///             HealthCheckConnectTimeout = 5,
    ///             HealthCheckInterval = 10,
    ///             HttpCheckMethod = "GET",
    ///             HealthCheckHttpCodes = new[]
    ///             {
    ///                 "http_2xx",
    ///                 "http_3xx",
    ///                 "http_4xx",
    ///             },
    ///         },
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///             { "For", "example" },
    ///         },
    ///     });
    /// 
    ///     var defaultListener = new AliCloud.Nlb.Listener("defaultListener", new()
    ///     {
    ///         ListenerProtocol = "TCPSSL",
    ///         ListenerPort = 1883,
    ///         SecurityPolicyId = "tls_cipher_policy_1_0",
    ///         ListenerDescription = name,
    ///         LoadBalancerId = defaultLoadBalancer.Id,
    ///         ServerGroupId = defaultServerGroup.Id,
    ///         IdleTimeout = 900,
    ///         ProxyProtocolEnabled = true,
    ///         SecSensorEnabled = true,
    ///         AlpnEnabled = true,
    ///         AlpnPolicy = "HTTP2Optional",
    ///         Cps = 10000,
    ///         Mss = 0,
    ///     });
    /// 
    ///     var defaultServiceCertificate = new AliCloud.Cas.ServiceCertificate("defaultServiceCertificate", new()
    ///     {
    ///         CertificateName = name,
    ///         Cert = @"-----BEGIN CERTIFICATE-----
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
    ///         Key = @"-----BEGIN RSA PRIVATE KEY-----
    /// MIICXAIBAAKBgQDO0kt3vfjY9BygLEbiAzUrEt1Cum/utmEG9rroSq6dRzKzsetV
    /// kg0dqpAJwXKBtNFM9/BBPvQy2DVFUASu2koCz+TNQJOMK+BqhsgoI3o884dw7IVM
    /// ywgstq6KJCjskU5T1/kOZe6xNX3Ejc4HqXGnuYtrBNV118y0SAyL0MXLKQIDAQAB
    /// AoGAfe3NxbsGKhN42o4bGsKZPQDfeCHMxayGp5bTd10BtQIE/ST4BcJH+ihAS7Bd
    /// 6FwQlKzivNd4GP1MckemklCXfsVckdL94e8ZbJl23GdWul3v8V+KndJHqv5zVJmP
    /// hwWoKimwIBTb2s0ctVryr2f18N4hhyFw1yGp0VxclGHkjgECQQD9CvllsnOwHpP4
    /// MdrDHbdb29QrobKyKW8pPcDd+sth+kP6Y8MnCVuAKXCKj5FeIsgVtfluPOsZjPzz
    /// 71QQWS1dAkEA0T0KXO8gaBQwJhIoo/w6hy5JGZnrNSpOPp5xvJuMAafs2eyvmhJm
    /// Ev9SN/Pf2VYa1z6FEnBaLOVD6hf6YQIsPQJAX/CZPoW6dzwgvimo1/GcY6eleiWE
    /// qygqjWhsh71e/3bz7yuEAnj5yE3t7Zshcp+dXR3xxGo0eSuLfLFxHgGxwQJAAxf8
    /// 9DzQ5NkPkTCJi0sqbl8/03IUKTgT6hcbpWdDXa7m8J3wRr3o5nUB+TPQ5nzAbthM
    /// zWX931YQeACcwhxvHQJBAN5mTzzJD4w4Ma6YTaNHyXakdYfyAWrOkPIWZxfhMfXe
    /// DrlNdiysTI4Dd1dLeErVpjsckAaOW/JDG5PCSwkaMxk=
    /// -----END RSA PRIVATE KEY-----
    /// ",
    ///     });
    /// 
    ///     var defaultListenerAdditionalCertificateAttachment = new AliCloud.Nlb.ListenerAdditionalCertificateAttachment("defaultListenerAdditionalCertificateAttachment", new()
    ///     {
    ///         CertificateId = defaultServiceCertificate.Id,
    ///         ListenerId = defaultListener.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// NLB Listener Additional Certificate Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:nlb/listenerAdditionalCertificateAttachment:ListenerAdditionalCertificateAttachment example &lt;listener_id&gt;:&lt;certificate_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:nlb/listenerAdditionalCertificateAttachment:ListenerAdditionalCertificateAttachment")]
    public partial class ListenerAdditionalCertificateAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Certificate ID. Currently, only server certificates are supported.
        /// </summary>
        [Output("certificateId")]
        public Output<string> CertificateId { get; private set; } = null!;

        /// <summary>
        /// Whether to PreCheck only this request, value: - **true**: sends a check request and does not create a resource. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '. - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The ID of the tcpssl listener.
        /// </summary>
        [Output("listenerId")]
        public Output<string> ListenerId { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a ListenerAdditionalCertificateAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ListenerAdditionalCertificateAttachment(string name, ListenerAdditionalCertificateAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:nlb/listenerAdditionalCertificateAttachment:ListenerAdditionalCertificateAttachment", name, args ?? new ListenerAdditionalCertificateAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ListenerAdditionalCertificateAttachment(string name, Input<string> id, ListenerAdditionalCertificateAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:nlb/listenerAdditionalCertificateAttachment:ListenerAdditionalCertificateAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ListenerAdditionalCertificateAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ListenerAdditionalCertificateAttachment Get(string name, Input<string> id, ListenerAdditionalCertificateAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ListenerAdditionalCertificateAttachment(name, id, state, options);
        }
    }

    public sealed class ListenerAdditionalCertificateAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate ID. Currently, only server certificates are supported.
        /// </summary>
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        /// <summary>
        /// Whether to PreCheck only this request, value: - **true**: sends a check request and does not create a resource. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '. - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the tcpssl listener.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        public ListenerAdditionalCertificateAttachmentArgs()
        {
        }
        public static new ListenerAdditionalCertificateAttachmentArgs Empty => new ListenerAdditionalCertificateAttachmentArgs();
    }

    public sealed class ListenerAdditionalCertificateAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate ID. Currently, only server certificates are supported.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// Whether to PreCheck only this request, value: - **true**: sends a check request and does not create a resource. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '. - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the tcpssl listener.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ListenerAdditionalCertificateAttachmentState()
        {
        }
        public static new ListenerAdditionalCertificateAttachmentState Empty => new ListenerAdditionalCertificateAttachmentState();
    }
}
