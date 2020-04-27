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
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSslVpnServersResult> InvokeAsync(GetSslVpnServersArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSslVpnServersResult>("alicloud:vpc/getSslVpnServers:getSslVpnServers", args ?? new GetSslVpnServersArgs(), options.WithVersion());
    }


    public sealed class GetSslVpnServersArgs : Pulumi.InvokeArgs
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
