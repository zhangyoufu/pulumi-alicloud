// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eds
{
    /// <summary>
    /// Provides a ECD Ad Connector Directory resource.
    /// 
    /// For information about ECD Ad Connector Directory and how to use it, see [What is Ad Connector Directory](https://www.alibabacloud.com/help/en/elastic-desktop-service/latest/api-doc-ecd-2020-09-30-api-doc-createadconnectordirectory).
    /// 
    /// &gt; **NOTE:** Available since v1.174.0.
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
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var defaultZones = AliCloud.Eds.GetZones.Invoke();
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Ids[0]),
    ///         VswitchName = name,
    ///     });
    /// 
    ///     var defaultAdConnectorDirectory = new AliCloud.Eds.AdConnectorDirectory("defaultAdConnectorDirectory", new()
    ///     {
    ///         DirectoryName = name,
    ///         DesktopAccessType = "INTERNET",
    ///         DnsAddresses = new[]
    ///         {
    ///             "127.0.0.2",
    ///         },
    ///         DomainName = "corp.example.com",
    ///         DomainPassword = "Example1234",
    ///         DomainUserName = "sAMAccountName",
    ///         EnableAdminAccess = false,
    ///         MfaEnabled = false,
    ///         Specification = 1,
    ///         SubDomainDnsAddresses = new[]
    ///         {
    ///             "127.0.0.3",
    ///         },
    ///         SubDomainName = "child.example.com",
    ///         VswitchIds = new[]
    ///         {
    ///             defaultSwitch.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECD Ad Connector Directory can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:eds/adConnectorDirectory:AdConnectorDirectory example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:eds/adConnectorDirectory:AdConnectorDirectory")]
    public partial class AdConnectorDirectory : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
        /// </summary>
        [Output("desktopAccessType")]
        public Output<string> DesktopAccessType { get; private set; } = null!;

        /// <summary>
        /// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        /// </summary>
        [Output("directoryName")]
        public Output<string> DirectoryName { get; private set; } = null!;

        /// <summary>
        /// The DNS address list.
        /// </summary>
        [Output("dnsAddresses")]
        public Output<ImmutableArray<string>> DnsAddresses { get; private set; } = null!;

        /// <summary>
        /// The name of the domain.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// The user password of the domain administrator. The maximum number of English characters is 64.
        /// </summary>
        [Output("domainPassword")]
        public Output<string> DomainPassword { get; private set; } = null!;

        /// <summary>
        /// The username of the domain administrator. The maximum number of English characters is 64.
        /// </summary>
        [Output("domainUserName")]
        public Output<string> DomainUserName { get; private set; } = null!;

        /// <summary>
        /// Whether to grant local administrator rights to users who use cloud desktops.
        /// </summary>
        [Output("enableAdminAccess")]
        public Output<bool> EnableAdminAccess { get; private set; } = null!;

        /// <summary>
        /// Whether MFA authentication is enabled. After all AD users in this directory log on to the cloud desktop, enter the correct password and then enter the dynamic verification code generated by the MFA device. **NOTE:** The MFA device needs to be bound when logging in for the first time.
        /// </summary>
        [Output("mfaEnabled")]
        public Output<bool> MfaEnabled { get; private set; } = null!;

        /// <summary>
        /// The AD Connector specifications. Valid values: `1`, `2`.
        /// </summary>
        [Output("specification")]
        public Output<int?> Specification { get; private set; } = null!;

        /// <summary>
        /// The status of directory.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The Enterprise already has the DNS address of the AD subdomain. If `sub_domain_name` is set and this parameter is not set, the child Domain DNS is considered consistent with the parent domain.
        /// </summary>
        [Output("subDomainDnsAddresses")]
        public Output<ImmutableArray<string>> SubDomainDnsAddresses { get; private set; } = null!;

        /// <summary>
        /// The Enterprise already has a fully qualified domain name (FQDN) of an AD subdomain, with both a host name and a domain name.
        /// </summary>
        [Output("subDomainName")]
        public Output<string?> SubDomainName { get; private set; } = null!;

        /// <summary>
        /// List of VSwitch IDs in the directory.
        /// </summary>
        [Output("vswitchIds")]
        public Output<ImmutableArray<string>> VswitchIds { get; private set; } = null!;


        /// <summary>
        /// Create a AdConnectorDirectory resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AdConnectorDirectory(string name, AdConnectorDirectoryArgs args, CustomResourceOptions? options = null)
            : base("alicloud:eds/adConnectorDirectory:AdConnectorDirectory", name, args ?? new AdConnectorDirectoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AdConnectorDirectory(string name, Input<string> id, AdConnectorDirectoryState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:eds/adConnectorDirectory:AdConnectorDirectory", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "domainPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AdConnectorDirectory resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AdConnectorDirectory Get(string name, Input<string> id, AdConnectorDirectoryState? state = null, CustomResourceOptions? options = null)
        {
            return new AdConnectorDirectory(name, id, state, options);
        }
    }

    public sealed class AdConnectorDirectoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
        /// </summary>
        [Input("desktopAccessType")]
        public Input<string>? DesktopAccessType { get; set; }

        /// <summary>
        /// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        /// </summary>
        [Input("directoryName", required: true)]
        public Input<string> DirectoryName { get; set; } = null!;

        [Input("dnsAddresses", required: true)]
        private InputList<string>? _dnsAddresses;

        /// <summary>
        /// The DNS address list.
        /// </summary>
        public InputList<string> DnsAddresses
        {
            get => _dnsAddresses ?? (_dnsAddresses = new InputList<string>());
            set => _dnsAddresses = value;
        }

        /// <summary>
        /// The name of the domain.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        [Input("domainPassword", required: true)]
        private Input<string>? _domainPassword;

        /// <summary>
        /// The user password of the domain administrator. The maximum number of English characters is 64.
        /// </summary>
        public Input<string>? DomainPassword
        {
            get => _domainPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _domainPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The username of the domain administrator. The maximum number of English characters is 64.
        /// </summary>
        [Input("domainUserName", required: true)]
        public Input<string> DomainUserName { get; set; } = null!;

        /// <summary>
        /// Whether to grant local administrator rights to users who use cloud desktops.
        /// </summary>
        [Input("enableAdminAccess")]
        public Input<bool>? EnableAdminAccess { get; set; }

        /// <summary>
        /// Whether MFA authentication is enabled. After all AD users in this directory log on to the cloud desktop, enter the correct password and then enter the dynamic verification code generated by the MFA device. **NOTE:** The MFA device needs to be bound when logging in for the first time.
        /// </summary>
        [Input("mfaEnabled")]
        public Input<bool>? MfaEnabled { get; set; }

        /// <summary>
        /// The AD Connector specifications. Valid values: `1`, `2`.
        /// </summary>
        [Input("specification")]
        public Input<int>? Specification { get; set; }

        [Input("subDomainDnsAddresses")]
        private InputList<string>? _subDomainDnsAddresses;

        /// <summary>
        /// The Enterprise already has the DNS address of the AD subdomain. If `sub_domain_name` is set and this parameter is not set, the child Domain DNS is considered consistent with the parent domain.
        /// </summary>
        public InputList<string> SubDomainDnsAddresses
        {
            get => _subDomainDnsAddresses ?? (_subDomainDnsAddresses = new InputList<string>());
            set => _subDomainDnsAddresses = value;
        }

        /// <summary>
        /// The Enterprise already has a fully qualified domain name (FQDN) of an AD subdomain, with both a host name and a domain name.
        /// </summary>
        [Input("subDomainName")]
        public Input<string>? SubDomainName { get; set; }

        [Input("vswitchIds", required: true)]
        private InputList<string>? _vswitchIds;

        /// <summary>
        /// List of VSwitch IDs in the directory.
        /// </summary>
        public InputList<string> VswitchIds
        {
            get => _vswitchIds ?? (_vswitchIds = new InputList<string>());
            set => _vswitchIds = value;
        }

        public AdConnectorDirectoryArgs()
        {
        }
        public static new AdConnectorDirectoryArgs Empty => new AdConnectorDirectoryArgs();
    }

    public sealed class AdConnectorDirectoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
        /// </summary>
        [Input("desktopAccessType")]
        public Input<string>? DesktopAccessType { get; set; }

        /// <summary>
        /// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        /// </summary>
        [Input("directoryName")]
        public Input<string>? DirectoryName { get; set; }

        [Input("dnsAddresses")]
        private InputList<string>? _dnsAddresses;

        /// <summary>
        /// The DNS address list.
        /// </summary>
        public InputList<string> DnsAddresses
        {
            get => _dnsAddresses ?? (_dnsAddresses = new InputList<string>());
            set => _dnsAddresses = value;
        }

        /// <summary>
        /// The name of the domain.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("domainPassword")]
        private Input<string>? _domainPassword;

        /// <summary>
        /// The user password of the domain administrator. The maximum number of English characters is 64.
        /// </summary>
        public Input<string>? DomainPassword
        {
            get => _domainPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _domainPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The username of the domain administrator. The maximum number of English characters is 64.
        /// </summary>
        [Input("domainUserName")]
        public Input<string>? DomainUserName { get; set; }

        /// <summary>
        /// Whether to grant local administrator rights to users who use cloud desktops.
        /// </summary>
        [Input("enableAdminAccess")]
        public Input<bool>? EnableAdminAccess { get; set; }

        /// <summary>
        /// Whether MFA authentication is enabled. After all AD users in this directory log on to the cloud desktop, enter the correct password and then enter the dynamic verification code generated by the MFA device. **NOTE:** The MFA device needs to be bound when logging in for the first time.
        /// </summary>
        [Input("mfaEnabled")]
        public Input<bool>? MfaEnabled { get; set; }

        /// <summary>
        /// The AD Connector specifications. Valid values: `1`, `2`.
        /// </summary>
        [Input("specification")]
        public Input<int>? Specification { get; set; }

        /// <summary>
        /// The status of directory.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("subDomainDnsAddresses")]
        private InputList<string>? _subDomainDnsAddresses;

        /// <summary>
        /// The Enterprise already has the DNS address of the AD subdomain. If `sub_domain_name` is set and this parameter is not set, the child Domain DNS is considered consistent with the parent domain.
        /// </summary>
        public InputList<string> SubDomainDnsAddresses
        {
            get => _subDomainDnsAddresses ?? (_subDomainDnsAddresses = new InputList<string>());
            set => _subDomainDnsAddresses = value;
        }

        /// <summary>
        /// The Enterprise already has a fully qualified domain name (FQDN) of an AD subdomain, with both a host name and a domain name.
        /// </summary>
        [Input("subDomainName")]
        public Input<string>? SubDomainName { get; set; }

        [Input("vswitchIds")]
        private InputList<string>? _vswitchIds;

        /// <summary>
        /// List of VSwitch IDs in the directory.
        /// </summary>
        public InputList<string> VswitchIds
        {
            get => _vswitchIds ?? (_vswitchIds = new InputList<string>());
            set => _vswitchIds = value;
        }

        public AdConnectorDirectoryState()
        {
        }
        public static new AdConnectorDirectoryState Empty => new AdConnectorDirectoryState();
    }
}
