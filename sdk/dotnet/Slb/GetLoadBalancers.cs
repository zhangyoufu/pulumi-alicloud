// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb
{
    public static class GetLoadBalancers
    {
        /// <summary>
        /// &gt; **DEPRECATED:** This datasource has been renamed to alicloud.slb.getApplicationLoadBalancers from version 1.123.1.
        /// 
        /// This data source provides the server load balancers of the current Alibaba Cloud user.
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
        ///     var @default = new AliCloud.Slb.LoadBalancer("default");
        /// 
        ///     var slbsDs = AliCloud.Slb.GetLoadBalancers.Invoke(new()
        ///     {
        ///         NameRegex = "sample_slb",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstSlbId"] = slbsDs.Apply(getLoadBalancersResult =&gt; getLoadBalancersResult.Slbs[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetLoadBalancersResult> InvokeAsync(GetLoadBalancersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLoadBalancersResult>("alicloud:slb/getLoadBalancers:getLoadBalancers", args ?? new GetLoadBalancersArgs(), options.WithDefaults());

        /// <summary>
        /// &gt; **DEPRECATED:** This datasource has been renamed to alicloud.slb.getApplicationLoadBalancers from version 1.123.1.
        /// 
        /// This data source provides the server load balancers of the current Alibaba Cloud user.
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
        ///     var @default = new AliCloud.Slb.LoadBalancer("default");
        /// 
        ///     var slbsDs = AliCloud.Slb.GetLoadBalancers.Invoke(new()
        ///     {
        ///         NameRegex = "sample_slb",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstSlbId"] = slbsDs.Apply(getLoadBalancersResult =&gt; getLoadBalancersResult.Slbs[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetLoadBalancersResult> Invoke(GetLoadBalancersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoadBalancersResult>("alicloud:slb/getLoadBalancers:getLoadBalancers", args ?? new GetLoadBalancersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoadBalancersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Service address of the SLBs.
        /// </summary>
        [Input("address")]
        public string? Address { get; set; }

        [Input("addressIpVersion")]
        public string? AddressIpVersion { get; set; }

        [Input("addressType")]
        public string? AddressType { get; set; }

        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of SLBs IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("internetChargeType")]
        public string? InternetChargeType { get; set; }

        [Input("loadBalancerName")]
        public string? LoadBalancerName { get; set; }

        [Input("masterZoneId")]
        public string? MasterZoneId { get; set; }

        /// <summary>
        /// A regex string to filter results by SLB name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// Network type of the SLBs. Valid values: `vpc` and `classic`.
        /// </summary>
        [Input("networkType")]
        public string? NetworkType { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        [Input("paymentType")]
        public string? PaymentType { get; set; }

        /// <summary>
        /// The Id of resource group which SLB belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("serverId")]
        public string? ServerId { get; set; }

        [Input("serverIntranetAddress")]
        public string? ServerIntranetAddress { get; set; }

        [Input("slaveZoneId")]
        public string? SlaveZoneId { get; set; }

        /// <summary>
        /// SLB current status. Possible values: `inactive`, `active` and `locked`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A map of tags assigned to the SLB instances. The `tags` can have a maximum of 5 tag. It must be in the format:
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var taggedInstances = AliCloud.Slb.GetLoadBalancers.Invoke(new()
        ///     {
        ///         Tags = 
        ///         {
        ///             { "tagKey1", "tagValue1" },
        ///             { "tagKey2", "tagValue2" },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of the VPC linked to the SLBs.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        /// <summary>
        /// ID of the VSwitch linked to the SLBs.
        /// </summary>
        [Input("vswitchId")]
        public string? VswitchId { get; set; }

        public GetLoadBalancersArgs()
        {
        }
        public static new GetLoadBalancersArgs Empty => new GetLoadBalancersArgs();
    }

    public sealed class GetLoadBalancersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Service address of the SLBs.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        [Input("addressIpVersion")]
        public Input<string>? AddressIpVersion { get; set; }

        [Input("addressType")]
        public Input<string>? AddressType { get; set; }

        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of SLBs IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("internetChargeType")]
        public Input<string>? InternetChargeType { get; set; }

        [Input("loadBalancerName")]
        public Input<string>? LoadBalancerName { get; set; }

        [Input("masterZoneId")]
        public Input<string>? MasterZoneId { get; set; }

        /// <summary>
        /// A regex string to filter results by SLB name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// Network type of the SLBs. Valid values: `vpc` and `classic`.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The Id of resource group which SLB belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        [Input("serverIntranetAddress")]
        public Input<string>? ServerIntranetAddress { get; set; }

        [Input("slaveZoneId")]
        public Input<string>? SlaveZoneId { get; set; }

        /// <summary>
        /// SLB current status. Possible values: `inactive`, `active` and `locked`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags assigned to the SLB instances. The `tags` can have a maximum of 5 tag. It must be in the format:
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var taggedInstances = AliCloud.Slb.GetLoadBalancers.Invoke(new()
        ///     {
        ///         Tags = 
        ///         {
        ///             { "tagKey1", "tagValue1" },
        ///             { "tagKey2", "tagValue2" },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of the VPC linked to the SLBs.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// ID of the VSwitch linked to the SLBs.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public GetLoadBalancersInvokeArgs()
        {
        }
        public static new GetLoadBalancersInvokeArgs Empty => new GetLoadBalancersInvokeArgs();
    }


    [OutputType]
    public sealed class GetLoadBalancersResult
    {
        /// <summary>
        /// Service address of the SLB.
        /// </summary>
        public readonly string? Address;
        public readonly string? AddressIpVersion;
        public readonly string? AddressType;
        public readonly ImmutableArray<Outputs.GetLoadBalancersBalancerResult> Balancers;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of slb IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? InternetChargeType;
        public readonly string? LoadBalancerName;
        public readonly string? MasterZoneId;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of slb names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        /// <summary>
        /// Network type of the SLB. Possible values: `vpc` and `classic`.
        /// </summary>
        public readonly string? NetworkType;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        public readonly string? PaymentType;
        public readonly string? ResourceGroupId;
        public readonly string? ServerId;
        public readonly string? ServerIntranetAddress;
        public readonly string? SlaveZoneId;
        /// <summary>
        /// A list of SLBs. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLoadBalancersSlbResult> Slbs;
        /// <summary>
        /// SLB current status. Possible values: `inactive`, `active` and `locked`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// A map of tags assigned to the SLB instance.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;
        public readonly int TotalCount;
        /// <summary>
        /// ID of the VPC the SLB belongs to.
        /// </summary>
        public readonly string? VpcId;
        /// <summary>
        /// ID of the VSwitch the SLB belongs to.
        /// </summary>
        public readonly string? VswitchId;

        [OutputConstructor]
        private GetLoadBalancersResult(
            string? address,

            string? addressIpVersion,

            string? addressType,

            ImmutableArray<Outputs.GetLoadBalancersBalancerResult> balancers,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? internetChargeType,

            string? loadBalancerName,

            string? masterZoneId,

            string? nameRegex,

            ImmutableArray<string> names,

            string? networkType,

            string? outputFile,

            int? pageNumber,

            int? pageSize,

            string? paymentType,

            string? resourceGroupId,

            string? serverId,

            string? serverIntranetAddress,

            string? slaveZoneId,

            ImmutableArray<Outputs.GetLoadBalancersSlbResult> slbs,

            string? status,

            ImmutableDictionary<string, object>? tags,

            int totalCount,

            string? vpcId,

            string? vswitchId)
        {
            Address = address;
            AddressIpVersion = addressIpVersion;
            AddressType = addressType;
            Balancers = balancers;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            InternetChargeType = internetChargeType;
            LoadBalancerName = loadBalancerName;
            MasterZoneId = masterZoneId;
            NameRegex = nameRegex;
            Names = names;
            NetworkType = networkType;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
            PaymentType = paymentType;
            ResourceGroupId = resourceGroupId;
            ServerId = serverId;
            ServerIntranetAddress = serverIntranetAddress;
            SlaveZoneId = slaveZoneId;
            Slbs = slbs;
            Status = status;
            Tags = tags;
            TotalCount = totalCount;
            VpcId = vpcId;
            VswitchId = vswitchId;
        }
    }
}
