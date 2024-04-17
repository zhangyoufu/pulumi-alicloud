// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga
{
    /// <summary>
    /// Provides a Global Accelerator (GA) Basic Endpoint Group resource.
    /// 
    /// For information about Global Accelerator (GA) Basic Endpoint Group and how to use it, see [What is Basic Endpoint Group](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createbasicendpointgroup).
    /// 
    /// &gt; **NOTE:** Available since v1.194.0.
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
    ///     var config = new Config();
    ///     var region = config.Get("region") ?? "cn-hangzhou";
    ///     var endpointGroupRegion = config.Get("endpointGroupRegion") ?? "cn-beijing";
    ///     var @default = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("default", new()
    ///     {
    ///         VpcName = "terraform-example",
    ///         CidrBlock = "172.17.3.0/24",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("default", new()
    ///     {
    ///         VswitchName = "terraform-example",
    ///         CidrBlock = "172.17.3.0/24",
    ///         VpcId = defaultNetwork.Id,
    ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id)),
    ///     });
    /// 
    ///     var defaultApplicationLoadBalancer = new AliCloud.Slb.ApplicationLoadBalancer("default", new()
    ///     {
    ///         LoadBalancerName = "terraform-example",
    ///         VswitchId = defaultSwitch.Id,
    ///         LoadBalancerSpec = "slb.s2.small",
    ///         AddressType = "intranet",
    ///     });
    /// 
    ///     var defaultBasicAccelerator = new AliCloud.Ga.BasicAccelerator("default", new()
    ///     {
    ///         Duration = 1,
    ///         BasicAcceleratorName = "terraform-example",
    ///         Description = "terraform-example",
    ///         BandwidthBillingType = "CDT",
    ///         AutoUseCoupon = "true",
    ///         AutoPay = true,
    ///     });
    /// 
    ///     var defaultBasicEndpointGroup = new AliCloud.Ga.BasicEndpointGroup("default", new()
    ///     {
    ///         AcceleratorId = defaultBasicAccelerator.Id,
    ///         EndpointGroupRegion = endpointGroupRegion,
    ///         EndpointType = "SLB",
    ///         EndpointAddress = defaultApplicationLoadBalancer.Id,
    ///         EndpointSubAddress = "192.168.0.1",
    ///         BasicEndpointGroupName = "terraform-example",
    ///         Description = "terraform-example",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Global Accelerator (GA) Basic Endpoint Group can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ga/basicEndpointGroup:BasicEndpointGroup example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ga/basicEndpointGroup:BasicEndpointGroup")]
    public partial class BasicEndpointGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the basic GA instance.
        /// </summary>
        [Output("acceleratorId")]
        public Output<string> AcceleratorId { get; private set; } = null!;

        /// <summary>
        /// The name of the endpoint group. The `basic_endpoint_group_name` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
        /// </summary>
        [Output("basicEndpointGroupName")]
        public Output<string?> BasicEndpointGroupName { get; private set; } = null!;

        /// <summary>
        /// The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The address of the endpoint.
        /// </summary>
        [Output("endpointAddress")]
        public Output<string> EndpointAddress { get; private set; } = null!;

        /// <summary>
        /// The ID of the region where you want to create the endpoint group.
        /// </summary>
        [Output("endpointGroupRegion")]
        public Output<string> EndpointGroupRegion { get; private set; } = null!;

        /// <summary>
        /// The sub address of the endpoint.
        /// </summary>
        [Output("endpointSubAddress")]
        public Output<string> EndpointSubAddress { get; private set; } = null!;

        /// <summary>
        /// The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
        /// </summary>
        [Output("endpointType")]
        public Output<string> EndpointType { get; private set; } = null!;

        /// <summary>
        /// The status of the Basic Endpoint Group.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a BasicEndpointGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BasicEndpointGroup(string name, BasicEndpointGroupArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ga/basicEndpointGroup:BasicEndpointGroup", name, args ?? new BasicEndpointGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BasicEndpointGroup(string name, Input<string> id, BasicEndpointGroupState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ga/basicEndpointGroup:BasicEndpointGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BasicEndpointGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BasicEndpointGroup Get(string name, Input<string> id, BasicEndpointGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new BasicEndpointGroup(name, id, state, options);
        }
    }

    public sealed class BasicEndpointGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the basic GA instance.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public Input<string> AcceleratorId { get; set; } = null!;

        /// <summary>
        /// The name of the endpoint group. The `basic_endpoint_group_name` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
        /// </summary>
        [Input("basicEndpointGroupName")]
        public Input<string>? BasicEndpointGroupName { get; set; }

        /// <summary>
        /// The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The address of the endpoint.
        /// </summary>
        [Input("endpointAddress")]
        public Input<string>? EndpointAddress { get; set; }

        /// <summary>
        /// The ID of the region where you want to create the endpoint group.
        /// </summary>
        [Input("endpointGroupRegion", required: true)]
        public Input<string> EndpointGroupRegion { get; set; } = null!;

        /// <summary>
        /// The sub address of the endpoint.
        /// </summary>
        [Input("endpointSubAddress")]
        public Input<string>? EndpointSubAddress { get; set; }

        /// <summary>
        /// The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        public BasicEndpointGroupArgs()
        {
        }
        public static new BasicEndpointGroupArgs Empty => new BasicEndpointGroupArgs();
    }

    public sealed class BasicEndpointGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the basic GA instance.
        /// </summary>
        [Input("acceleratorId")]
        public Input<string>? AcceleratorId { get; set; }

        /// <summary>
        /// The name of the endpoint group. The `basic_endpoint_group_name` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
        /// </summary>
        [Input("basicEndpointGroupName")]
        public Input<string>? BasicEndpointGroupName { get; set; }

        /// <summary>
        /// The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The address of the endpoint.
        /// </summary>
        [Input("endpointAddress")]
        public Input<string>? EndpointAddress { get; set; }

        /// <summary>
        /// The ID of the region where you want to create the endpoint group.
        /// </summary>
        [Input("endpointGroupRegion")]
        public Input<string>? EndpointGroupRegion { get; set; }

        /// <summary>
        /// The sub address of the endpoint.
        /// </summary>
        [Input("endpointSubAddress")]
        public Input<string>? EndpointSubAddress { get; set; }

        /// <summary>
        /// The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// The status of the Basic Endpoint Group.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public BasicEndpointGroupState()
        {
        }
        public static new BasicEndpointGroupState Empty => new BasicEndpointGroupState();
    }
}
