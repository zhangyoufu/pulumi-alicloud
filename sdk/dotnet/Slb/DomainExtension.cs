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
    /// HTTPS listeners of guaranteed-performance SLB support configuring multiple certificates, allowing you to forward requests with different domain names to different backend servers.
    /// Please refer to the [documentation](https://www.alibabacloud.com/help/doc-detail/85956.htm?spm=a2c63.p38356.b99.40.1c881563Co8p6w) for details.
    /// 
    /// &gt; **NOTE:** Available in 1.60.0+
    /// 
    /// &gt; **NOTE:** The instance with shared loadBalancerSpec doesn't support domainExtension.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var slbDomainExtensionName = config.Get("slbDomainExtensionName") ?? "forDomainExtension";
    ///     var domainExtension = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var domainExtensionNetwork = new AliCloud.Vpc.Network("domain_extension", new()
    ///     {
    ///         VpcName = slbDomainExtensionName,
    ///     });
    /// 
    ///     var domainExtensionSwitch = new AliCloud.Vpc.Switch("domain_extension", new()
    ///     {
    ///         VpcId = domainExtensionNetwork.Id,
    ///         CidrBlock = "172.16.0.0/21",
    ///         ZoneId = domainExtension.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         VswitchName = slbDomainExtensionName,
    ///     });
    /// 
    ///     var instance = new AliCloud.Slb.ApplicationLoadBalancer("instance", new()
    ///     {
    ///         LoadBalancerName = slbDomainExtensionName,
    ///         AddressType = "intranet",
    ///         LoadBalancerSpec = "slb.s2.small",
    ///         VswitchId = domainExtensionSwitch.Id,
    ///     });
    /// 
    ///     var domainExtensionServerCertificate = new AliCloud.Slb.ServerCertificate("domain_extension", new()
    ///     {
    ///         Name = "tf-testAccSlbServerCertificate",
    ///         Certificate = @"-----BEGIN CERTIFICATE-----
    /// MIIDdjCCAl4CCQCcm+erkcKN7DANBgkqhkiG9w0BAQsFADB9MQswCQYDVQQGEwJj
    /// bjELMAkGA1UECAwCYmoxEDAOBgNVBAcMB2JlaWppbmcxDzANBgNVBAoMBmFsaXl1
    /// bjELMAkGA1UECwwCc2MxFTATBgNVBAMMDHd3dy50ZXN0LmNvbTEaMBgGCSqGSIb3
    /// DQEJARYLMTIzQDEyMy5jb20wHhcNMTkwNDI2MDM0ODAxWhcNMjQwNDI1MDM0ODAx
    /// WjB9MQswCQYDVQQGEwJjbjELMAkGA1UECAwCYmoxEDAOBgNVBAcMB2JlaWppbmcx
    /// DzANBgNVBAoMBmFsaXl1bjELMAkGA1UECwwCc2MxFTATBgNVBAMMDHd3dy50ZXN0
    /// LmNvbTEaMBgGCSqGSIb3DQEJARYLMTIzQDEyMy5jb20wggEiMA0GCSqGSIb3DQEB
    /// AQUAA4IBDwAwggEKAoIBAQDKMKF5qmN/uoMjdH3D8aPRcUOA0s8rZpYhG8zbkF1j
    /// 8gHYoB/FDvM7G7dfVsyjbMwLOxKvAhWvHHSpEz/t7gB+QdwrAMiMJwGmtCnXrh2E
    /// WiXgalMe1y4a/T5R7q+m4T1zFATf+kbnHWfkSGF4W7b6UBoaH+9StQ95CnqzNf/2
    /// p/Of7+S0XzCxFXw8GIVzZk0xFe6lHJzaq06f3mvzrD+4rpO56tTUvrgTY/n61gsF
    /// ZP7f0CJ2JQh6eNRFOEUSfxKu/Dy/+IsQxorCJY2Q59ZAf3rXrqDN104jw9PlwnLl
    /// qfZz3RMODN6BWjxE8rvRtT0qMfuAfv1gjBdWZN0hUYBRAgMBAAEwDQYJKoZIhvcN
    /// AQELBQADggEBAABzo82TxGp5poVkd5pLWj5ACgcBv8Cs6oH9D+4Jz9BmyuBUsQXh
    /// 2aG0hQAe1mU61C9konsl/GTW8umJQ4M4lYEztXXwMf5PlBMGwebM0ZbSGg6jKtZg
    /// WCgJ3eP/FMmyXGL5Jji5+e09eObhUDVle4tdi0On97zBoz85W02rgWFAqZJwiEAP
    /// t+c7jX7uOSBq2/38iGStlrX5yB1at/gJXXiA5CL5OtlR3Okvb0/QH37efO1Nu39m
    /// lFi0ODPAVyXjVypAiLguDxPn6AtDTdk9Iw9B19OD4NrzNRWgSSX5vuxo/VcRcgWk
    /// 3gEe9Ca0ZKN20q9XgthAiFFjl1S9ZgdA6Zc=
    /// -----END CERTIFICATE-----",
    ///         PrivateKey = @"-----BEGIN RSA PRIVATE KEY-----
    /// MIIEowIBAAKCAQEAyjCheapjf7qDI3R9w/Gj0XFDgNLPK2aWIRvM25BdY/IB2KAf
    /// xQ7zOxu3X1bMo2zMCzsSrwIVrxx0qRM/7e4AfkHcKwDIjCcBprQp164dhFol4GpT
    /// HtcuGv0+Ue6vpuE9cxQE3/pG5x1n5EhheFu2+lAaGh/vUrUPeQp6szX/9qfzn+/k
    /// tF8wsRV8PBiFc2ZNMRXupRyc2qtOn95r86w/uK6TuerU1L64E2P5+tYLBWT+39Ai
    /// diUIenjURThFEn8Srvw8v/iLEMaKwiWNkOfWQH96166gzddOI8PT5cJy5an2c90T
    /// DgzegVo8RPK70bU9KjH7gH79YIwXVmTdIVGAUQIDAQABAoIBAE1J4a/8biR5S3/W
    /// G+03BYQeY8tuyjqw8FqfoeOcf9agwAvqybouSNQjeCk9qOQfxq/UWQQFK/zQR9gJ
    /// v7pX7GBXFK5rkj3g+0SaQhRsPmRFgY0Tl8qGPt2aSKRRNVv5ZeADmwlzRn86QmiF
    /// Mp0rkfqFfDTYWEepZszCML0ouzuxsW/9tq7rvtSjsgATNt31B3vFa3D3JBi31jUh
    /// 5nfR9A3bATze7mQw3byEDiVl5ASRDgYyur403P1fDnMy9DBHZ8NaPOsFF6OBpJal
    /// BJsG5z00hll5PFN2jfmBQKlvAeU7wfwqdaSnGHOfqf2DeTTaFjIQ4gUhRn/m6pLo
    /// 6kXttLECgYEA9sng0Qz/TcPFfM4tQ1gyvB1cKnnGIwg1FP8sfUjbbEgjaHhA224S
    /// k3BxtX2Kq6fhTXuwusAFc6OVMAZ76FgrQ5K4Ci7+DTsrF28z4b8td+p+lO/DxgP9
    /// lTgN+ddsiTOV4fUef9Z3yY0Zr0CnBUMbQYRaV2UIbCdiB0G4V/bt9TsCgYEA0bya
    /// Oo9wGI0RJV0bYP7qwO74Ra1/i1viWbRlS7jU37Q+AZstrlKcQ5CTPzOjKFKMiUzl
    /// 4miWacZ0/q2n+Mvd7NbXGXTLijahnyOYKaHJYyh4oBymfkgAifRstE0Ki9gdvArb
    /// /I+emC0GvLSyfGN8UUeDJs4NmqdEXGqjo2JOV+MCgYALFv1MR5o9Y1u/hQBRs2fs
    /// PiGDIx+9OUQxYloccyaxEfjNXAIGGkcpavchIbgWiJ++PJ2vdquIC8TLeK8evL+M
    /// 9M3iX0Q5UfxYvD2HmnCvn9D6Xl/cyRcfGnq+TGjrLW9BzSMGuZt+aiHKV0xqFx7l
    /// bc4leTvMqGRmURS4lzcQOwKBgQCDzA/i4sYfN25h21tcHXSpnsG3D2rJyQi5NCo/
    /// ZjunA92/JqOTGuiFcLGHEszhhtY3ZXJET1LNz18vtzKJnpqrvOnYXlOVW/U+SqDQ
    /// 8JDb1c/PVZGuY1KrXkR9HLiW3kz5IJ3S3PFdUVYdeTN8BQxXCyg4V12nJJtJs912
    /// y0zN3wKBgGDS6YttCN6aI4EOABYE8fI1EYQ7vhfiYsaWGWSR1l6bQey7KR6M1ACz
    /// ZzMASNyytVt12yXE4/Emv6/pYqigbDLfL1zQJSLJ3EHJYTh2RxjR+AaGDudYFG/T
    /// liQ9YXhV5Iu2x1pNwrtFnssDdaaGpfA7l3xC00BL7Z+SAJyI4QKA
    /// -----END RSA PRIVATE KEY-----",
    ///     });
    /// 
    ///     var https = new AliCloud.Slb.Listener("https", new()
    ///     {
    ///         LoadBalancerId = instance.Id,
    ///         BackendPort = 80,
    ///         FrontendPort = 443,
    ///         Protocol = "https",
    ///         StickySession = "on",
    ///         StickySessionType = "insert",
    ///         Cookie = "testslblistenercookie",
    ///         CookieTimeout = 86400,
    ///         HealthCheck = "on",
    ///         HealthCheckUri = "/cons",
    ///         HealthCheckConnectPort = 20,
    ///         HealthyThreshold = 8,
    ///         UnhealthyThreshold = 8,
    ///         HealthCheckTimeout = 8,
    ///         HealthCheckInterval = 5,
    ///         HealthCheckHttpCode = "http_2xx,http_3xx",
    ///         Bandwidth = 10,
    ///         ServerCertificateId = domainExtensionServerCertificate.Id,
    ///     });
    /// 
    ///     var example1 = new AliCloud.Slb.DomainExtension("example1", new()
    ///     {
    ///         LoadBalancerId = instance.Id,
    ///         FrontendPort = https.FrontendPort,
    ///         Domain = "www.test.com",
    ///         ServerCertificateId = domainExtensionServerCertificate.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Load balancer domain_extension can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:slb/domainExtension:DomainExtension example de-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:slb/domainExtension:DomainExtension")]
    public partial class DomainExtension : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Output("deleteProtectionValidation")]
        public Output<bool?> DeleteProtectionValidation { get; private set; } = null!;

        /// <summary>
        /// The domain name.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
        /// </summary>
        [Output("frontendPort")]
        public Output<int> FrontendPort { get; private set; } = null!;

        /// <summary>
        /// The ID of the SLB instance.
        /// </summary>
        [Output("loadBalancerId")]
        public Output<string> LoadBalancerId { get; private set; } = null!;

        /// <summary>
        /// The ID of the certificate used by the domain name.
        /// </summary>
        [Output("serverCertificateId")]
        public Output<string> ServerCertificateId { get; private set; } = null!;


        /// <summary>
        /// Create a DomainExtension resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainExtension(string name, DomainExtensionArgs args, CustomResourceOptions? options = null)
            : base("alicloud:slb/domainExtension:DomainExtension", name, args ?? new DomainExtensionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainExtension(string name, Input<string> id, DomainExtensionState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:slb/domainExtension:DomainExtension", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DomainExtension resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainExtension Get(string name, Input<string> id, DomainExtensionState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainExtension(name, id, state, options);
        }
    }

    public sealed class DomainExtensionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Input("deleteProtectionValidation")]
        public Input<bool>? DeleteProtectionValidation { get; set; }

        /// <summary>
        /// The domain name.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
        /// </summary>
        [Input("frontendPort", required: true)]
        public Input<int> FrontendPort { get; set; } = null!;

        /// <summary>
        /// The ID of the SLB instance.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public Input<string> LoadBalancerId { get; set; } = null!;

        /// <summary>
        /// The ID of the certificate used by the domain name.
        /// </summary>
        [Input("serverCertificateId", required: true)]
        public Input<string> ServerCertificateId { get; set; } = null!;

        public DomainExtensionArgs()
        {
        }
        public static new DomainExtensionArgs Empty => new DomainExtensionArgs();
    }

    public sealed class DomainExtensionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Input("deleteProtectionValidation")]
        public Input<bool>? DeleteProtectionValidation { get; set; }

        /// <summary>
        /// The domain name.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
        /// </summary>
        [Input("frontendPort")]
        public Input<int>? FrontendPort { get; set; }

        /// <summary>
        /// The ID of the SLB instance.
        /// </summary>
        [Input("loadBalancerId")]
        public Input<string>? LoadBalancerId { get; set; }

        /// <summary>
        /// The ID of the certificate used by the domain name.
        /// </summary>
        [Input("serverCertificateId")]
        public Input<string>? ServerCertificateId { get; set; }

        public DomainExtensionState()
        {
        }
        public static new DomainExtensionState Empty => new DomainExtensionState();
    }
}
