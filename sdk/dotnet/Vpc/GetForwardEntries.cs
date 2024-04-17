// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    public static class GetForwardEntries
    {
        /// <summary>
        /// This data source provides a list of Forward Entries owned by an Alibaba Cloud account.
        /// 
        /// &gt; **NOTE:** Available in 1.37.0+.
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
        ///     var name = config.Get("name") ?? "forward-entry-config-example-name";
        ///     var @default = AliCloud.GetZones.Invoke(new()
        ///     {
        ///         AvailableResourceCreation = "VSwitch",
        ///     });
        /// 
        ///     var defaultNetwork = new AliCloud.Vpc.Network("default", new()
        ///     {
        ///         VpcName = name,
        ///         CidrBlock = "172.16.0.0/12",
        ///     });
        /// 
        ///     var defaultSwitch = new AliCloud.Vpc.Switch("default", new()
        ///     {
        ///         VpcId = defaultNetwork.Id,
        ///         CidrBlock = "172.16.0.0/21",
        ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id)),
        ///         VswitchName = name,
        ///     });
        /// 
        ///     var defaultNatGateway = new AliCloud.Vpc.NatGateway("default", new()
        ///     {
        ///         VpcId = defaultNetwork.Id,
        ///         InternetChargeType = "PayByLcu",
        ///         NatGatewayName = name,
        ///         NatType = "Enhanced",
        ///         VswitchId = defaultSwitch.Id,
        ///     });
        /// 
        ///     var defaultEipAddress = new AliCloud.Ecs.EipAddress("default", new()
        ///     {
        ///         AddressName = name,
        ///     });
        /// 
        ///     var defaultEipAssociation = new AliCloud.Ecs.EipAssociation("default", new()
        ///     {
        ///         AllocationId = defaultEipAddress.Id,
        ///         InstanceId = defaultNatGateway.Id,
        ///     });
        /// 
        ///     var defaultForwardEntry = new AliCloud.Vpc.ForwardEntry("default", new()
        ///     {
        ///         ForwardTableId = defaultNatGateway.ForwardTableIds,
        ///         ExternalIp = defaultEipAddress.IpAddress,
        ///         ExternalPort = "80",
        ///         IpProtocol = "tcp",
        ///         InternalIp = "172.16.0.3",
        ///         InternalPort = "8080",
        ///     });
        /// 
        ///     var defaultGetForwardEntries = AliCloud.Vpc.GetForwardEntries.Invoke(new()
        ///     {
        ///         ForwardTableId = defaultForwardEntry.ForwardTableId,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetForwardEntriesResult> InvokeAsync(GetForwardEntriesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetForwardEntriesResult>("alicloud:vpc/getForwardEntries:getForwardEntries", args ?? new GetForwardEntriesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides a list of Forward Entries owned by an Alibaba Cloud account.
        /// 
        /// &gt; **NOTE:** Available in 1.37.0+.
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
        ///     var name = config.Get("name") ?? "forward-entry-config-example-name";
        ///     var @default = AliCloud.GetZones.Invoke(new()
        ///     {
        ///         AvailableResourceCreation = "VSwitch",
        ///     });
        /// 
        ///     var defaultNetwork = new AliCloud.Vpc.Network("default", new()
        ///     {
        ///         VpcName = name,
        ///         CidrBlock = "172.16.0.0/12",
        ///     });
        /// 
        ///     var defaultSwitch = new AliCloud.Vpc.Switch("default", new()
        ///     {
        ///         VpcId = defaultNetwork.Id,
        ///         CidrBlock = "172.16.0.0/21",
        ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id)),
        ///         VswitchName = name,
        ///     });
        /// 
        ///     var defaultNatGateway = new AliCloud.Vpc.NatGateway("default", new()
        ///     {
        ///         VpcId = defaultNetwork.Id,
        ///         InternetChargeType = "PayByLcu",
        ///         NatGatewayName = name,
        ///         NatType = "Enhanced",
        ///         VswitchId = defaultSwitch.Id,
        ///     });
        /// 
        ///     var defaultEipAddress = new AliCloud.Ecs.EipAddress("default", new()
        ///     {
        ///         AddressName = name,
        ///     });
        /// 
        ///     var defaultEipAssociation = new AliCloud.Ecs.EipAssociation("default", new()
        ///     {
        ///         AllocationId = defaultEipAddress.Id,
        ///         InstanceId = defaultNatGateway.Id,
        ///     });
        /// 
        ///     var defaultForwardEntry = new AliCloud.Vpc.ForwardEntry("default", new()
        ///     {
        ///         ForwardTableId = defaultNatGateway.ForwardTableIds,
        ///         ExternalIp = defaultEipAddress.IpAddress,
        ///         ExternalPort = "80",
        ///         IpProtocol = "tcp",
        ///         InternalIp = "172.16.0.3",
        ///         InternalPort = "8080",
        ///     });
        /// 
        ///     var defaultGetForwardEntries = AliCloud.Vpc.GetForwardEntries.Invoke(new()
        ///     {
        ///         ForwardTableId = defaultForwardEntry.ForwardTableId,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetForwardEntriesResult> Invoke(GetForwardEntriesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetForwardEntriesResult>("alicloud:vpc/getForwardEntries:getForwardEntries", args ?? new GetForwardEntriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetForwardEntriesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The public IP address.
        /// </summary>
        [Input("externalIp")]
        public string? ExternalIp { get; set; }

        /// <summary>
        /// The public port.
        /// </summary>
        [Input("externalPort")]
        public string? ExternalPort { get; set; }

        /// <summary>
        /// The name of forward entry.
        /// </summary>
        [Input("forwardEntryName")]
        public string? ForwardEntryName { get; set; }

        /// <summary>
        /// The ID of the Forward table.
        /// </summary>
        [Input("forwardTableId", required: true)]
        public string ForwardTableId { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Forward Entries IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The private IP address.
        /// </summary>
        [Input("internalIp")]
        public string? InternalIp { get; set; }

        /// <summary>
        /// The internal port.
        /// </summary>
        [Input("internalPort")]
        public string? InternalPort { get; set; }

        /// <summary>
        /// The ip protocol. Valid values: `any`,`tcp` and `udp`.
        /// </summary>
        [Input("ipProtocol")]
        public string? IpProtocol { get; set; }

        /// <summary>
        /// A regex string to filter results by forward entry name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of farward entry. Valid value `Available`, `Deleting` and `Pending`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetForwardEntriesArgs()
        {
        }
        public static new GetForwardEntriesArgs Empty => new GetForwardEntriesArgs();
    }

    public sealed class GetForwardEntriesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The public IP address.
        /// </summary>
        [Input("externalIp")]
        public Input<string>? ExternalIp { get; set; }

        /// <summary>
        /// The public port.
        /// </summary>
        [Input("externalPort")]
        public Input<string>? ExternalPort { get; set; }

        /// <summary>
        /// The name of forward entry.
        /// </summary>
        [Input("forwardEntryName")]
        public Input<string>? ForwardEntryName { get; set; }

        /// <summary>
        /// The ID of the Forward table.
        /// </summary>
        [Input("forwardTableId", required: true)]
        public Input<string> ForwardTableId { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Forward Entries IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The private IP address.
        /// </summary>
        [Input("internalIp")]
        public Input<string>? InternalIp { get; set; }

        /// <summary>
        /// The internal port.
        /// </summary>
        [Input("internalPort")]
        public Input<string>? InternalPort { get; set; }

        /// <summary>
        /// The ip protocol. Valid values: `any`,`tcp` and `udp`.
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        /// <summary>
        /// A regex string to filter results by forward entry name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of farward entry. Valid value `Available`, `Deleting` and `Pending`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetForwardEntriesInvokeArgs()
        {
        }
        public static new GetForwardEntriesInvokeArgs Empty => new GetForwardEntriesInvokeArgs();
    }


    [OutputType]
    public sealed class GetForwardEntriesResult
    {
        /// <summary>
        /// A list of Forward Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardEntriesEntryResult> Entries;
        /// <summary>
        /// The public IP address.
        /// </summary>
        public readonly string? ExternalIp;
        /// <summary>
        /// The public port.
        /// </summary>
        public readonly string? ExternalPort;
        /// <summary>
        /// The name of forward entry.
        /// </summary>
        public readonly string? ForwardEntryName;
        public readonly string ForwardTableId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Forward Entries IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The private IP address.
        /// </summary>
        public readonly string? InternalIp;
        /// <summary>
        /// The private port.
        /// </summary>
        public readonly string? InternalPort;
        /// <summary>
        /// The protocol type.
        /// </summary>
        public readonly string? IpProtocol;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of Forward Entries names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// The status of forward entry.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private GetForwardEntriesResult(
            ImmutableArray<Outputs.GetForwardEntriesEntryResult> entries,

            string? externalIp,

            string? externalPort,

            string? forwardEntryName,

            string forwardTableId,

            string id,

            ImmutableArray<string> ids,

            string? internalIp,

            string? internalPort,

            string? ipProtocol,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? status)
        {
            Entries = entries;
            ExternalIp = externalIp;
            ExternalPort = externalPort;
            ForwardEntryName = forwardEntryName;
            ForwardTableId = forwardTableId;
            Id = id;
            Ids = ids;
            InternalIp = internalIp;
            InternalPort = internalPort;
            IpProtocol = ipProtocol;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
        }
    }
}
