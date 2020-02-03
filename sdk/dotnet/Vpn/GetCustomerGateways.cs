// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpn
{
    public static partial class Invokes
    {
        /// <summary>
        /// The VPN customers gateways data source lists a number of VPN customer gateways resource information owned by an Alicloud account.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/vpn_customer_gateways.html.markdown.
        /// </summary>
        public static Task<GetCustomerGatewaysResult> GetCustomerGateways(GetCustomerGatewaysArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCustomerGatewaysResult>("alicloud:vpn/getCustomerGateways:getCustomerGateways", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetCustomerGatewaysArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// ID of the VPN customer gateways.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string of VPN customer gateways name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// Save the result to the file.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetCustomerGatewaysArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetCustomerGatewaysResult
    {
        /// <summary>
        /// A list of VPN customer gateways. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCustomerGatewaysGatewaysResult> Gateways;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetCustomerGatewaysResult(
            ImmutableArray<Outputs.GetCustomerGatewaysGatewaysResult> gateways,
            ImmutableArray<string> ids,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            string id)
        {
            Gateways = gateways;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetCustomerGatewaysGatewaysResult
    {
        /// <summary>
        /// The creation time of the VPN customer gateway.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of the VPN customer gateway.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// ID of the VPN customer gateway .
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ip address of the VPN customer gateway.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The name of the VPN customer gateway.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetCustomerGatewaysGatewaysResult(
            string createTime,
            string description,
            string id,
            string ipAddress,
            string name)
        {
            CreateTime = createTime;
            Description = description;
            Id = id;
            IpAddress = ipAddress;
            Name = name;
        }
    }
    }
}
