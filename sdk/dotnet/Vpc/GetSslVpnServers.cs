// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    public static class GetSslVpnServers
    {
        /// <summary>
        /// The SSL-VPN servers data source lists lots of SSL-VPN servers resource information owned by an Alicloud account.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = AliCloud.Vpc.GetSslVpnServers.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "fake-server-id",
        ///         },
        ///         NameRegex = "^foo",
        ///         OutputFile = "/tmp/sslserver",
        ///         VpnGatewayId = "fake-vpn-id",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSslVpnServersResult> InvokeAsync(GetSslVpnServersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSslVpnServersResult>("alicloud:vpc/getSslVpnServers:getSslVpnServers", args ?? new GetSslVpnServersArgs(), options.WithDefaults());

        /// <summary>
        /// The SSL-VPN servers data source lists lots of SSL-VPN servers resource information owned by an Alicloud account.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = AliCloud.Vpc.GetSslVpnServers.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "fake-server-id",
        ///         },
        ///         NameRegex = "^foo",
        ///         OutputFile = "/tmp/sslserver",
        ///         VpnGatewayId = "fake-vpn-id",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSslVpnServersResult> Invoke(GetSslVpnServersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSslVpnServersResult>("alicloud:vpc/getSslVpnServers:getSslVpnServers", args ?? new GetSslVpnServersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSslVpnServersArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// IDs of the SSL-VPN servers.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string of SSL-VPN server name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// Save the result to the file.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Use the VPN gateway ID as the search key.
        /// </summary>
        [Input("vpnGatewayId")]
        public string? VpnGatewayId { get; set; }

        public GetSslVpnServersArgs()
        {
        }
        public static new GetSslVpnServersArgs Empty => new GetSslVpnServersArgs();
    }

    public sealed class GetSslVpnServersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// IDs of the SSL-VPN servers.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string of SSL-VPN server name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// Save the result to the file.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Use the VPN gateway ID as the search key.
        /// </summary>
        [Input("vpnGatewayId")]
        public Input<string>? VpnGatewayId { get; set; }

        public GetSslVpnServersInvokeArgs()
        {
        }
        public static new GetSslVpnServersInvokeArgs Empty => new GetSslVpnServersInvokeArgs();
    }


    [OutputType]
    public sealed class GetSslVpnServersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of SSL-VPN server IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of SSL-VPN server names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of SSL-VPN servers. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSslVpnServersServerResult> Servers;
        /// <summary>
        /// The ID of the VPN gateway instance.
        /// </summary>
        public readonly string? VpnGatewayId;

        [OutputConstructor]
        private GetSslVpnServersResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetSslVpnServersServerResult> servers,

            string? vpnGatewayId)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Servers = servers;
            VpnGatewayId = vpnGatewayId;
        }
    }
}
