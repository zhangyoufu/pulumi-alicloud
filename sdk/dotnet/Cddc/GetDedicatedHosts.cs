// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cddc
{
    public static class GetDedicatedHosts
    {
        /// <summary>
        /// This data source provides the Cddc Dedicated Hosts of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.147.0+.
        /// 
        /// ## Example Usage
        /// 
        /// Basic Usage
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
        ///     var ids = AliCloud.Cddc.GetDedicatedHosts.Invoke(new()
        ///     {
        ///         DedicatedHostGroupId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///     });
        /// 
        ///     var status = AliCloud.Cddc.GetDedicatedHosts.Invoke(new()
        ///     {
        ///         DedicatedHostGroupId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///         Status = "1",
        ///     });
        /// 
        ///     var zoneId = AliCloud.Cddc.GetDedicatedHosts.Invoke(new()
        ///     {
        ///         DedicatedHostGroupId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///         ZoneId = "example_value",
        ///     });
        /// 
        ///     var allocationStatus = AliCloud.Cddc.GetDedicatedHosts.Invoke(new()
        ///     {
        ///         DedicatedHostGroupId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///         AllocationStatus = "Allocatable",
        ///     });
        /// 
        ///     var hostType = AliCloud.Cddc.GetDedicatedHosts.Invoke(new()
        ///     {
        ///         DedicatedHostGroupId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///         HostType = "dhg_cloud_ssd",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cddcDedicatedHostId1"] = ids.Apply(getDedicatedHostsResult =&gt; getDedicatedHostsResult.Hosts[0]?.Id),
        ///         ["cddcDedicatedHostId2"] = status.Apply(getDedicatedHostsResult =&gt; getDedicatedHostsResult.Hosts[0]?.Id),
        ///         ["cddcDedicatedHostId3"] = zoneId.Apply(getDedicatedHostsResult =&gt; getDedicatedHostsResult.Hosts[0]?.Id),
        ///         ["cddcDedicatedHostId4"] = allocationStatus.Apply(getDedicatedHostsResult =&gt; getDedicatedHostsResult.Hosts[0]?.Id),
        ///         ["cddcDedicatedHostId5"] = hostType.Apply(getDedicatedHostsResult =&gt; getDedicatedHostsResult.Hosts[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetDedicatedHostsResult> InvokeAsync(GetDedicatedHostsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDedicatedHostsResult>("alicloud:cddc/getDedicatedHosts:getDedicatedHosts", args ?? new GetDedicatedHostsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Cddc Dedicated Hosts of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.147.0+.
        /// 
        /// ## Example Usage
        /// 
        /// Basic Usage
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
        ///     var ids = AliCloud.Cddc.GetDedicatedHosts.Invoke(new()
        ///     {
        ///         DedicatedHostGroupId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///     });
        /// 
        ///     var status = AliCloud.Cddc.GetDedicatedHosts.Invoke(new()
        ///     {
        ///         DedicatedHostGroupId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///         Status = "1",
        ///     });
        /// 
        ///     var zoneId = AliCloud.Cddc.GetDedicatedHosts.Invoke(new()
        ///     {
        ///         DedicatedHostGroupId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///         ZoneId = "example_value",
        ///     });
        /// 
        ///     var allocationStatus = AliCloud.Cddc.GetDedicatedHosts.Invoke(new()
        ///     {
        ///         DedicatedHostGroupId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///         AllocationStatus = "Allocatable",
        ///     });
        /// 
        ///     var hostType = AliCloud.Cddc.GetDedicatedHosts.Invoke(new()
        ///     {
        ///         DedicatedHostGroupId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///         HostType = "dhg_cloud_ssd",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cddcDedicatedHostId1"] = ids.Apply(getDedicatedHostsResult =&gt; getDedicatedHostsResult.Hosts[0]?.Id),
        ///         ["cddcDedicatedHostId2"] = status.Apply(getDedicatedHostsResult =&gt; getDedicatedHostsResult.Hosts[0]?.Id),
        ///         ["cddcDedicatedHostId3"] = zoneId.Apply(getDedicatedHostsResult =&gt; getDedicatedHostsResult.Hosts[0]?.Id),
        ///         ["cddcDedicatedHostId4"] = allocationStatus.Apply(getDedicatedHostsResult =&gt; getDedicatedHostsResult.Hosts[0]?.Id),
        ///         ["cddcDedicatedHostId5"] = hostType.Apply(getDedicatedHostsResult =&gt; getDedicatedHostsResult.Hosts[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetDedicatedHostsResult> Invoke(GetDedicatedHostsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDedicatedHostsResult>("alicloud:cddc/getDedicatedHosts:getDedicatedHosts", args ?? new GetDedicatedHostsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDedicatedHostsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies whether instances can be created on the host. Valid values: `1` or `0`. `1`: Instances can be created on the host. `0`: Instances cannot be created on the host.
        /// </summary>
        [Input("allocationStatus")]
        public string? AllocationStatus { get; set; }

        /// <summary>
        /// The ID of the dedicated cluster in which the host is created.
        /// </summary>
        [Input("dedicatedHostGroupId", required: true)]
        public string DedicatedHostGroupId { get; set; } = null!;

        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        /// <summary>
        /// The storage type of the host.
        /// </summary>
        [Input("hostType")]
        public string? HostType { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Dedicated Host IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The ID of the order.
        /// </summary>
        [Input("orderId")]
        public string? OrderId { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The state of the host.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone ID of the host.
        /// </summary>
        [Input("zoneId")]
        public string? ZoneId { get; set; }

        public GetDedicatedHostsArgs()
        {
        }
        public static new GetDedicatedHostsArgs Empty => new GetDedicatedHostsArgs();
    }

    public sealed class GetDedicatedHostsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies whether instances can be created on the host. Valid values: `1` or `0`. `1`: Instances can be created on the host. `0`: Instances cannot be created on the host.
        /// </summary>
        [Input("allocationStatus")]
        public Input<string>? AllocationStatus { get; set; }

        /// <summary>
        /// The ID of the dedicated cluster in which the host is created.
        /// </summary>
        [Input("dedicatedHostGroupId", required: true)]
        public Input<string> DedicatedHostGroupId { get; set; } = null!;

        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        /// <summary>
        /// The storage type of the host.
        /// </summary>
        [Input("hostType")]
        public Input<string>? HostType { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Dedicated Host IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The ID of the order.
        /// </summary>
        [Input("orderId")]
        public Input<string>? OrderId { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The state of the host.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone ID of the host.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public GetDedicatedHostsInvokeArgs()
        {
        }
        public static new GetDedicatedHostsInvokeArgs Empty => new GetDedicatedHostsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDedicatedHostsResult
    {
        public readonly string? AllocationStatus;
        public readonly string DedicatedHostGroupId;
        public readonly bool? EnableDetails;
        public readonly string? HostType;
        public readonly ImmutableArray<Outputs.GetDedicatedHostsHostResult> Hosts;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OrderId;
        public readonly string? OutputFile;
        public readonly string? Status;
        public readonly ImmutableDictionary<string, object>? Tags;
        public readonly string? ZoneId;

        [OutputConstructor]
        private GetDedicatedHostsResult(
            string? allocationStatus,

            string dedicatedHostGroupId,

            bool? enableDetails,

            string? hostType,

            ImmutableArray<Outputs.GetDedicatedHostsHostResult> hosts,

            string id,

            ImmutableArray<string> ids,

            string? orderId,

            string? outputFile,

            string? status,

            ImmutableDictionary<string, object>? tags,

            string? zoneId)
        {
            AllocationStatus = allocationStatus;
            DedicatedHostGroupId = dedicatedHostGroupId;
            EnableDetails = enableDetails;
            HostType = hostType;
            Hosts = hosts;
            Id = id;
            Ids = ids;
            OrderId = orderId;
            OutputFile = outputFile;
            Status = status;
            Tags = tags;
            ZoneId = zoneId;
        }
    }
}
