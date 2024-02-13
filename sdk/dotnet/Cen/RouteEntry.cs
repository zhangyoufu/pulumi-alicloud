// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    /// <summary>
    /// Provides a CEN route entry resource. Cloud Enterprise Network (CEN) supports publishing and withdrawing route entries of attached networks. You can publish a route entry of an attached VPC or VBR to a CEN instance, then other attached networks can learn the route if there is no route conflict. You can withdraw a published route entry when CEN does not need it any more.
    /// 
    /// For information about CEN route entries publishment and how to use it, see [Manage network routes](https://www.alibabacloud.com/help/doc-detail/86980.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.20.0.
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
    ///     var @default = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var exampleZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "Instance",
    ///     });
    /// 
    ///     var exampleInstanceTypes = AliCloud.Ecs.GetInstanceTypes.Invoke(new()
    ///     {
    ///         AvailabilityZone = exampleZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         CpuCoreCount = 1,
    ///         MemorySize = 2,
    ///     });
    /// 
    ///     var exampleImages = AliCloud.Ecs.GetImages.Invoke(new()
    ///     {
    ///         NameRegex = "^ubuntu_[0-9]+_[0-9]+_x64*",
    ///         Owners = "system",
    ///     });
    /// 
    ///     var exampleNetwork = new AliCloud.Vpc.Network("exampleNetwork", new()
    ///     {
    ///         VpcName = "terraform-example",
    ///         CidrBlock = "172.17.3.0/24",
    ///     });
    /// 
    ///     var exampleSwitch = new AliCloud.Vpc.Switch("exampleSwitch", new()
    ///     {
    ///         VswitchName = "terraform-example",
    ///         CidrBlock = "172.17.3.0/24",
    ///         VpcId = exampleNetwork.Id,
    ///         ZoneId = exampleZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var exampleSecurityGroup = new AliCloud.Ecs.SecurityGroup("exampleSecurityGroup", new()
    ///     {
    ///         VpcId = exampleNetwork.Id,
    ///     });
    /// 
    ///     var exampleInstance = new AliCloud.Ecs.Instance("exampleInstance", new()
    ///     {
    ///         AvailabilityZone = exampleZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         InstanceName = "terraform-example",
    ///         ImageId = exampleImages.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.Id),
    ///         InstanceType = exampleInstanceTypes.Apply(getInstanceTypesResult =&gt; getInstanceTypesResult.InstanceTypes[0]?.Id),
    ///         SecurityGroups = new[]
    ///         {
    ///             exampleSecurityGroup.Id,
    ///         },
    ///         VswitchId = exampleSwitch.Id,
    ///         InternetMaxBandwidthOut = 5,
    ///     });
    /// 
    ///     var exampleCen_instanceInstance = new AliCloud.Cen.Instance("exampleCen/instanceInstance", new()
    ///     {
    ///         CenInstanceName = "tf_example",
    ///         Description = "an example for cen",
    ///     });
    /// 
    ///     var exampleInstanceAttachment = new AliCloud.Cen.InstanceAttachment("exampleInstanceAttachment", new()
    ///     {
    ///         InstanceId = exampleCen / instanceInstance.Id,
    ///         ChildInstanceId = exampleNetwork.Id,
    ///         ChildInstanceType = "VPC",
    ///         ChildInstanceRegionId = @default.Apply(@default =&gt; @default.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)),
    ///     });
    /// 
    ///     var exampleRouteEntry = new AliCloud.Vpc.RouteEntry("exampleRouteEntry", new()
    ///     {
    ///         RouteTableId = exampleNetwork.RouteTableId,
    ///         DestinationCidrblock = "11.0.0.0/16",
    ///         NexthopType = "Instance",
    ///         NexthopId = exampleInstance.Id,
    ///     });
    /// 
    ///     var exampleCen_routeEntryRouteEntry = new AliCloud.Cen.RouteEntry("exampleCen/routeEntryRouteEntry", new()
    ///     {
    ///         InstanceId = exampleInstanceAttachment.InstanceId,
    ///         RouteTableId = exampleNetwork.RouteTableId,
    ///         CidrBlock = exampleRouteEntry.DestinationCidrblock,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CEN instance can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cen/routeEntry:RouteEntry example cen-abc123456:vtb-abc123:192.168.0.0/24
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/routeEntry:RouteEntry")]
    public partial class RouteEntry : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The destination CIDR block of the route entry to publish.
        /// 
        /// -&gt;**NOTE:** The "alicloud_cen_instance_route_entries" resource depends on the related "alicloud.cen.InstanceAttachment" resource.
        /// 
        /// -&gt;**NOTE:** The "alicloud.cen.InstanceAttachment" resource should depend on the related "alicloud.vpc.Switch" resource.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The route table of the attached VBR or VPC.
        /// </summary>
        [Output("routeTableId")]
        public Output<string> RouteTableId { get; private set; } = null!;


        /// <summary>
        /// Create a RouteEntry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteEntry(string name, RouteEntryArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/routeEntry:RouteEntry", name, args ?? new RouteEntryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteEntry(string name, Input<string> id, RouteEntryState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/routeEntry:RouteEntry", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RouteEntry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteEntry Get(string name, Input<string> id, RouteEntryState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteEntry(name, id, state, options);
        }
    }

    public sealed class RouteEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination CIDR block of the route entry to publish.
        /// 
        /// -&gt;**NOTE:** The "alicloud_cen_instance_route_entries" resource depends on the related "alicloud.cen.InstanceAttachment" resource.
        /// 
        /// -&gt;**NOTE:** The "alicloud.cen.InstanceAttachment" resource should depend on the related "alicloud.vpc.Switch" resource.
        /// </summary>
        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The route table of the attached VBR or VPC.
        /// </summary>
        [Input("routeTableId", required: true)]
        public Input<string> RouteTableId { get; set; } = null!;

        public RouteEntryArgs()
        {
        }
        public static new RouteEntryArgs Empty => new RouteEntryArgs();
    }

    public sealed class RouteEntryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination CIDR block of the route entry to publish.
        /// 
        /// -&gt;**NOTE:** The "alicloud_cen_instance_route_entries" resource depends on the related "alicloud.cen.InstanceAttachment" resource.
        /// 
        /// -&gt;**NOTE:** The "alicloud.cen.InstanceAttachment" resource should depend on the related "alicloud.vpc.Switch" resource.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The route table of the attached VBR or VPC.
        /// </summary>
        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        public RouteEntryState()
        {
        }
        public static new RouteEntryState Empty => new RouteEntryState();
    }
}
