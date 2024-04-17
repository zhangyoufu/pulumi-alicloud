// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    public static class GetSnatEntries
    {
        /// <summary>
        /// This data source provides a list of Snat Entries owned by an Alibaba Cloud account.
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
        ///     var name = config.Get("name") ?? "snat-entry-example-name";
        ///     var @default = AliCloud.GetZones.Invoke(new()
        ///     {
        ///         AvailableResourceCreation = "VSwitch",
        ///     });
        /// 
        ///     var fooNetwork = new AliCloud.Vpc.Network("foo", new()
        ///     {
        ///         Name = name,
        ///         CidrBlock = "172.16.0.0/12",
        ///     });
        /// 
        ///     var fooSwitch = new AliCloud.Vpc.Switch("foo", new()
        ///     {
        ///         VpcId = fooNetwork.Id,
        ///         CidrBlock = "172.16.0.0/21",
        ///         AvailabilityZone = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id)),
        ///         VswitchName = name,
        ///     });
        /// 
        ///     var fooNatGateway = new AliCloud.Vpc.NatGateway("foo", new()
        ///     {
        ///         VpcId = fooNetwork.Id,
        ///         Specification = "Small",
        ///         Name = name,
        ///     });
        /// 
        ///     var fooEipAddress = new AliCloud.Ecs.EipAddress("foo", new()
        ///     {
        ///         AddressName = name,
        ///     });
        /// 
        ///     var fooEipAssociation = new AliCloud.Ecs.EipAssociation("foo", new()
        ///     {
        ///         AllocationId = fooEipAddress.Id,
        ///         InstanceId = fooNatGateway.Id,
        ///     });
        /// 
        ///     var fooSnatEntry = new AliCloud.Vpc.SnatEntry("foo", new()
        ///     {
        ///         SnatTableId = fooNatGateway.SnatTableIds,
        ///         SourceVswitchId = fooSwitch.Id,
        ///         SnatIp = fooEipAddress.IpAddress,
        ///     });
        /// 
        ///     var foo = AliCloud.Vpc.GetSnatEntries.Invoke(new()
        ///     {
        ///         SnatTableId = fooSnatEntry.SnatTableId,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetSnatEntriesResult> InvokeAsync(GetSnatEntriesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSnatEntriesResult>("alicloud:vpc/getSnatEntries:getSnatEntries", args ?? new GetSnatEntriesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides a list of Snat Entries owned by an Alibaba Cloud account.
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
        ///     var name = config.Get("name") ?? "snat-entry-example-name";
        ///     var @default = AliCloud.GetZones.Invoke(new()
        ///     {
        ///         AvailableResourceCreation = "VSwitch",
        ///     });
        /// 
        ///     var fooNetwork = new AliCloud.Vpc.Network("foo", new()
        ///     {
        ///         Name = name,
        ///         CidrBlock = "172.16.0.0/12",
        ///     });
        /// 
        ///     var fooSwitch = new AliCloud.Vpc.Switch("foo", new()
        ///     {
        ///         VpcId = fooNetwork.Id,
        ///         CidrBlock = "172.16.0.0/21",
        ///         AvailabilityZone = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id)),
        ///         VswitchName = name,
        ///     });
        /// 
        ///     var fooNatGateway = new AliCloud.Vpc.NatGateway("foo", new()
        ///     {
        ///         VpcId = fooNetwork.Id,
        ///         Specification = "Small",
        ///         Name = name,
        ///     });
        /// 
        ///     var fooEipAddress = new AliCloud.Ecs.EipAddress("foo", new()
        ///     {
        ///         AddressName = name,
        ///     });
        /// 
        ///     var fooEipAssociation = new AliCloud.Ecs.EipAssociation("foo", new()
        ///     {
        ///         AllocationId = fooEipAddress.Id,
        ///         InstanceId = fooNatGateway.Id,
        ///     });
        /// 
        ///     var fooSnatEntry = new AliCloud.Vpc.SnatEntry("foo", new()
        ///     {
        ///         SnatTableId = fooNatGateway.SnatTableIds,
        ///         SourceVswitchId = fooSwitch.Id,
        ///         SnatIp = fooEipAddress.IpAddress,
        ///     });
        /// 
        ///     var foo = AliCloud.Vpc.GetSnatEntries.Invoke(new()
        ///     {
        ///         SnatTableId = fooSnatEntry.SnatTableId,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetSnatEntriesResult> Invoke(GetSnatEntriesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnatEntriesResult>("alicloud:vpc/getSnatEntries:getSnatEntries", args ?? new GetSnatEntriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnatEntriesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Snat Entries IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by the resource name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The name of snat entry.
        /// </summary>
        [Input("snatEntryName")]
        public string? SnatEntryName { get; set; }

        /// <summary>
        /// The public IP of the Snat Entry.
        /// </summary>
        [Input("snatIp")]
        public string? SnatIp { get; set; }

        /// <summary>
        /// The ID of the Snat table.
        /// </summary>
        [Input("snatTableId", required: true)]
        public string SnatTableId { get; set; } = null!;

        /// <summary>
        /// The source CIDR block of the Snat Entry.
        /// </summary>
        [Input("sourceCidr")]
        public string? SourceCidr { get; set; }

        /// <summary>
        /// The source vswitch ID.
        /// </summary>
        [Input("sourceVswitchId")]
        public string? SourceVswitchId { get; set; }

        /// <summary>
        /// The status of the Snat Entry. Valid values: `Available`, `Deleting` and `Pending`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetSnatEntriesArgs()
        {
        }
        public static new GetSnatEntriesArgs Empty => new GetSnatEntriesArgs();
    }

    public sealed class GetSnatEntriesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Snat Entries IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by the resource name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The name of snat entry.
        /// </summary>
        [Input("snatEntryName")]
        public Input<string>? SnatEntryName { get; set; }

        /// <summary>
        /// The public IP of the Snat Entry.
        /// </summary>
        [Input("snatIp")]
        public Input<string>? SnatIp { get; set; }

        /// <summary>
        /// The ID of the Snat table.
        /// </summary>
        [Input("snatTableId", required: true)]
        public Input<string> SnatTableId { get; set; } = null!;

        /// <summary>
        /// The source CIDR block of the Snat Entry.
        /// </summary>
        [Input("sourceCidr")]
        public Input<string>? SourceCidr { get; set; }

        /// <summary>
        /// The source vswitch ID.
        /// </summary>
        [Input("sourceVswitchId")]
        public Input<string>? SourceVswitchId { get; set; }

        /// <summary>
        /// The status of the Snat Entry. Valid values: `Available`, `Deleting` and `Pending`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetSnatEntriesInvokeArgs()
        {
        }
        public static new GetSnatEntriesInvokeArgs Empty => new GetSnatEntriesInvokeArgs();
    }


    [OutputType]
    public sealed class GetSnatEntriesResult
    {
        /// <summary>
        /// A list of Snat Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSnatEntriesEntryResult> Entries;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Optional) A list of Snat Entries IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// The name of snat entry.
        /// </summary>
        public readonly string? SnatEntryName;
        /// <summary>
        /// The public IP of the Snat Entry.
        /// </summary>
        public readonly string? SnatIp;
        public readonly string SnatTableId;
        /// <summary>
        /// The source CIDR block of the Snat Entry.
        /// </summary>
        public readonly string? SourceCidr;
        /// <summary>
        /// The source vswitch ID.
        /// </summary>
        public readonly string? SourceVswitchId;
        /// <summary>
        /// The status of the Snat Entry.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private GetSnatEntriesResult(
            ImmutableArray<Outputs.GetSnatEntriesEntryResult> entries,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? snatEntryName,

            string? snatIp,

            string snatTableId,

            string? sourceCidr,

            string? sourceVswitchId,

            string? status)
        {
            Entries = entries;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            SnatEntryName = snatEntryName;
            SnatIp = snatIp;
            SnatTableId = snatTableId;
            SourceCidr = sourceCidr;
            SourceVswitchId = sourceVswitchId;
            Status = status;
        }
    }
}
