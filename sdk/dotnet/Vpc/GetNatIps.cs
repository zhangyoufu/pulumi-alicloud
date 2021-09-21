// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    public static class GetNatIps
    {
        /// <summary>
        /// This data source provides the Vpc Nat Ips of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.136.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///         var ids = Output.Create(AliCloud.Vpc.GetNatIps.InvokeAsync(new AliCloud.Vpc.GetNatIpsArgs
        ///         {
        ///             NatGatewayId = "example_value",
        ///             Ids = 
        ///             {
        ///                 "example_value-1",
        ///                 "example_value-2",
        ///             },
        ///         }));
        ///         this.VpcNatIpId1 = ids.Apply(ids =&gt; ids.Ips[0].Id);
        ///         var nameRegex = Output.Create(AliCloud.Vpc.GetNatIps.InvokeAsync(new AliCloud.Vpc.GetNatIpsArgs
        ///         {
        ///             NatGatewayId = "example_value",
        ///             NameRegex = "^my-NatIp",
        ///         }));
        ///         this.VpcNatIpId2 = nameRegex.Apply(nameRegex =&gt; nameRegex.Ips[0].Id);
        ///         var natIpCidr = Output.Create(AliCloud.Vpc.GetNatIps.InvokeAsync(new AliCloud.Vpc.GetNatIpsArgs
        ///         {
        ///             NatGatewayId = "example_value",
        ///             NatIpCidr = "example_value",
        ///             NameRegex = "^my-NatIp",
        ///         }));
        ///         this.VpcNatIpId3 = natIpCidr.Apply(natIpCidr =&gt; natIpCidr.Ips[0].Id);
        ///         var natIpName = Output.Create(AliCloud.Vpc.GetNatIps.InvokeAsync(new AliCloud.Vpc.GetNatIpsArgs
        ///         {
        ///             NatGatewayId = "example_value",
        ///             Ids = 
        ///             {
        ///                 "example_value",
        ///             },
        ///             NatIpNames = 
        ///             {
        ///                 "example_value",
        ///             },
        ///         }));
        ///         this.VpcNatIpId4 = natIpName.Apply(natIpName =&gt; natIpName.Ips[0].Id);
        ///         var natIpIds = Output.Create(AliCloud.Vpc.GetNatIps.InvokeAsync(new AliCloud.Vpc.GetNatIpsArgs
        ///         {
        ///             NatGatewayId = "example_value",
        ///             Ids = 
        ///             {
        ///                 "example_value",
        ///             },
        ///             NatIpIds = 
        ///             {
        ///                 "example_value",
        ///             },
        ///         }));
        ///         this.VpcNatIpId5 = natIpIds.Apply(natIpIds =&gt; natIpIds.Ips[0].Id);
        ///         var status = Output.Create(AliCloud.Vpc.GetNatIps.InvokeAsync(new AliCloud.Vpc.GetNatIpsArgs
        ///         {
        ///             NatGatewayId = "example_value",
        ///             Ids = 
        ///             {
        ///                 "example_value",
        ///             },
        ///             Status = "example_value",
        ///         }));
        ///         this.VpcNatIpId6 = status.Apply(status =&gt; status.Ips[0].Id);
        ///     }
        /// 
        ///     [Output("vpcNatIpId1")]
        ///     public Output&lt;string&gt; VpcNatIpId1 { get; set; }
        ///     [Output("vpcNatIpId2")]
        ///     public Output&lt;string&gt; VpcNatIpId2 { get; set; }
        ///     [Output("vpcNatIpId3")]
        ///     public Output&lt;string&gt; VpcNatIpId3 { get; set; }
        ///     [Output("vpcNatIpId4")]
        ///     public Output&lt;string&gt; VpcNatIpId4 { get; set; }
        ///     [Output("vpcNatIpId5")]
        ///     public Output&lt;string&gt; VpcNatIpId5 { get; set; }
        ///     [Output("vpcNatIpId6")]
        ///     public Output&lt;string&gt; VpcNatIpId6 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNatIpsResult> InvokeAsync(GetNatIpsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNatIpsResult>("alicloud:vpc/getNatIps:getNatIps", args ?? new GetNatIpsArgs(), options.WithVersion());
    }


    public sealed class GetNatIpsArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Nat Ip IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Nat Ip name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// The ID of the Virtual Private Cloud (VPC) NAT gateway to which the NAT IP address belongs.
        /// </summary>
        [Input("natGatewayId", required: true)]
        public string NatGatewayId { get; set; } = null!;

        /// <summary>
        /// The CIDR block to which the NAT IP address belongs.
        /// </summary>
        [Input("natIpCidr")]
        public string? NatIpCidr { get; set; }

        [Input("natIpIds")]
        private List<string>? _natIpIds;
        public List<string> NatIpIds
        {
            get => _natIpIds ?? (_natIpIds = new List<string>());
            set => _natIpIds = value;
        }

        [Input("natIpNames")]
        private List<string>? _natIpNames;

        /// <summary>
        /// The name of the NAT IP address.
        /// </summary>
        public List<string> NatIpNames
        {
            get => _natIpNames ?? (_natIpNames = new List<string>());
            set => _natIpNames = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the NAT IP address. Valid values: `Available`, `Deleting` and `Creating`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetNatIpsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNatIpsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<Outputs.GetNatIpsIpResult> Ips;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string NatGatewayId;
        public readonly string? NatIpCidr;
        public readonly ImmutableArray<string> NatIpIds;
        public readonly ImmutableArray<string> NatIpNames;
        public readonly string? OutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetNatIpsResult(
            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetNatIpsIpResult> ips,

            string? nameRegex,

            ImmutableArray<string> names,

            string natGatewayId,

            string? natIpCidr,

            ImmutableArray<string> natIpIds,

            ImmutableArray<string> natIpNames,

            string? outputFile,

            string? status)
        {
            Id = id;
            Ids = ids;
            Ips = ips;
            NameRegex = nameRegex;
            Names = names;
            NatGatewayId = natGatewayId;
            NatIpCidr = natIpCidr;
            NatIpIds = natIpIds;
            NatIpNames = natIpNames;
            OutputFile = outputFile;
            Status = status;
        }
    }
}
