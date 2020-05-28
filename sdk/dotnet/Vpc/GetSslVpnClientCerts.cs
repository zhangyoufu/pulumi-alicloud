// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    public static class GetSslVpnClientCerts
    {
        /// <summary>
        /// The SSL-VPN client certificates data source lists lots of SSL-VPN client certificates resource information owned by an Alicloud account.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var foo = Output.Create(AliCloud.Vpc.GetSslVpnClientCerts.InvokeAsync(new AliCloud.Vpc.GetSslVpnClientCertsArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "fake-cert-id",
        ///             },
        ///             NameRegex = "^foo",
        ///             OutputFile = "/tmp/clientcert",
        ///             SslVpnServerId = "fake-server-id",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSslVpnClientCertsResult> InvokeAsync(GetSslVpnClientCertsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSslVpnClientCertsResult>("alicloud:vpc/getSslVpnClientCerts:getSslVpnClientCerts", args ?? new GetSslVpnClientCertsArgs(), options.WithVersion());
    }


    public sealed class GetSslVpnClientCertsArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// IDs of the SSL-VPN client certificates.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string of SSL-VPN client certificate name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// Save the result to the file.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Use the SSL-VPN server ID as the search key.
        /// </summary>
        [Input("sslVpnServerId")]
        public string? SslVpnServerId { get; set; }

        public GetSslVpnClientCertsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSslVpnClientCertsResult
    {
        public readonly ImmutableArray<Outputs.GetSslVpnClientCertsCertResult> Certs;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of SSL-VPN client cert IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of SSL-VPN client cert names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// ID of the SSL-VPN Server.
        /// </summary>
        public readonly string? SslVpnServerId;

        [OutputConstructor]
        private GetSslVpnClientCertsResult(
            ImmutableArray<Outputs.GetSslVpnClientCertsCertResult> certs,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? sslVpnServerId)
        {
            Certs = certs;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            SslVpnServerId = sslVpnServerId;
        }
    }
}
