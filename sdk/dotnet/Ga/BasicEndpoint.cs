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
    /// Provides a Global Accelerator (GA) Basic Endpoint resource.
    /// 
    /// For information about Global Accelerator (GA) Basic Endpoint and how to use it, see [What is Basic Endpoint](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createbasicendpoint).
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
    ///     var region = config.Get("region") ?? "cn-shenzhen";
    ///     var endpointRegion = config.Get("endpointRegion") ?? "cn-hangzhou";
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
    ///     var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("default", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///         Name = "terraform-example",
    ///     });
    /// 
    ///     var defaultEcsNetworkInterface = new AliCloud.Ecs.EcsNetworkInterface("default", new()
    ///     {
    ///         VswitchId = defaultSwitch.Id,
    ///         SecurityGroupIds = new[]
    ///         {
    ///             defaultSecurityGroup.Id,
    ///         },
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
    ///         EndpointGroupRegion = region,
    ///         BasicEndpointGroupName = "terraform-example",
    ///         Description = "terraform-example",
    ///     });
    /// 
    ///     var defaultBasicEndpoint = new AliCloud.Ga.BasicEndpoint("default", new()
    ///     {
    ///         AcceleratorId = defaultBasicAccelerator.Id,
    ///         EndpointGroupId = defaultBasicEndpointGroup.Id,
    ///         EndpointType = "ENI",
    ///         EndpointAddress = defaultEcsNetworkInterface.Id,
    ///         EndpointSubAddressType = "secondary",
    ///         EndpointSubAddress = "192.168.0.1",
    ///         BasicEndpointName = "terraform-example",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Global Accelerator (GA) Basic Endpoint can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ga/basicEndpoint:BasicEndpoint example &lt;endpoint_group_id&gt;:&lt;endpoint_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ga/basicEndpoint:BasicEndpoint")]
    public partial class BasicEndpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Basic GA instance.
        /// </summary>
        [Output("acceleratorId")]
        public Output<string> AcceleratorId { get; private set; } = null!;

        /// <summary>
        /// The name of the Basic Endpoint.
        /// </summary>
        [Output("basicEndpointName")]
        public Output<string?> BasicEndpointName { get; private set; } = null!;

        /// <summary>
        /// The address of the Basic Endpoint.
        /// </summary>
        [Output("endpointAddress")]
        public Output<string> EndpointAddress { get; private set; } = null!;

        /// <summary>
        /// The ID of the Basic Endpoint Group.
        /// </summary>
        [Output("endpointGroupId")]
        public Output<string> EndpointGroupId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Basic Endpoint.
        /// </summary>
        [Output("endpointId")]
        public Output<string> EndpointId { get; private set; } = null!;

        /// <summary>
        /// The sub address of the Basic Endpoint.
        /// </summary>
        [Output("endpointSubAddress")]
        public Output<string?> EndpointSubAddress { get; private set; } = null!;

        /// <summary>
        /// The sub address type of the Basic Endpoint. Valid values: `primary`, `secondary`.
        /// </summary>
        [Output("endpointSubAddressType")]
        public Output<string?> EndpointSubAddressType { get; private set; } = null!;

        /// <summary>
        /// The type of the Basic Endpoint. Valid values: `ENI`, `SLB`, `ECS` and `NLB`.
        /// </summary>
        [Output("endpointType")]
        public Output<string> EndpointType { get; private set; } = null!;

        /// <summary>
        /// The zone id of the Basic Endpoint.
        /// </summary>
        [Output("endpointZoneId")]
        public Output<string?> EndpointZoneId { get; private set; } = null!;

        /// <summary>
        /// The status of the Basic Endpoint.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a BasicEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BasicEndpoint(string name, BasicEndpointArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ga/basicEndpoint:BasicEndpoint", name, args ?? new BasicEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BasicEndpoint(string name, Input<string> id, BasicEndpointState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ga/basicEndpoint:BasicEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BasicEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BasicEndpoint Get(string name, Input<string> id, BasicEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new BasicEndpoint(name, id, state, options);
        }
    }

    public sealed class BasicEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Basic GA instance.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public Input<string> AcceleratorId { get; set; } = null!;

        /// <summary>
        /// The name of the Basic Endpoint.
        /// </summary>
        [Input("basicEndpointName")]
        public Input<string>? BasicEndpointName { get; set; }

        /// <summary>
        /// The address of the Basic Endpoint.
        /// </summary>
        [Input("endpointAddress", required: true)]
        public Input<string> EndpointAddress { get; set; } = null!;

        /// <summary>
        /// The ID of the Basic Endpoint Group.
        /// </summary>
        [Input("endpointGroupId", required: true)]
        public Input<string> EndpointGroupId { get; set; } = null!;

        /// <summary>
        /// The sub address of the Basic Endpoint.
        /// </summary>
        [Input("endpointSubAddress")]
        public Input<string>? EndpointSubAddress { get; set; }

        /// <summary>
        /// The sub address type of the Basic Endpoint. Valid values: `primary`, `secondary`.
        /// </summary>
        [Input("endpointSubAddressType")]
        public Input<string>? EndpointSubAddressType { get; set; }

        /// <summary>
        /// The type of the Basic Endpoint. Valid values: `ENI`, `SLB`, `ECS` and `NLB`.
        /// </summary>
        [Input("endpointType", required: true)]
        public Input<string> EndpointType { get; set; } = null!;

        /// <summary>
        /// The zone id of the Basic Endpoint.
        /// </summary>
        [Input("endpointZoneId")]
        public Input<string>? EndpointZoneId { get; set; }

        public BasicEndpointArgs()
        {
        }
        public static new BasicEndpointArgs Empty => new BasicEndpointArgs();
    }

    public sealed class BasicEndpointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Basic GA instance.
        /// </summary>
        [Input("acceleratorId")]
        public Input<string>? AcceleratorId { get; set; }

        /// <summary>
        /// The name of the Basic Endpoint.
        /// </summary>
        [Input("basicEndpointName")]
        public Input<string>? BasicEndpointName { get; set; }

        /// <summary>
        /// The address of the Basic Endpoint.
        /// </summary>
        [Input("endpointAddress")]
        public Input<string>? EndpointAddress { get; set; }

        /// <summary>
        /// The ID of the Basic Endpoint Group.
        /// </summary>
        [Input("endpointGroupId")]
        public Input<string>? EndpointGroupId { get; set; }

        /// <summary>
        /// The ID of the Basic Endpoint.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// The sub address of the Basic Endpoint.
        /// </summary>
        [Input("endpointSubAddress")]
        public Input<string>? EndpointSubAddress { get; set; }

        /// <summary>
        /// The sub address type of the Basic Endpoint. Valid values: `primary`, `secondary`.
        /// </summary>
        [Input("endpointSubAddressType")]
        public Input<string>? EndpointSubAddressType { get; set; }

        /// <summary>
        /// The type of the Basic Endpoint. Valid values: `ENI`, `SLB`, `ECS` and `NLB`.
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// The zone id of the Basic Endpoint.
        /// </summary>
        [Input("endpointZoneId")]
        public Input<string>? EndpointZoneId { get; set; }

        /// <summary>
        /// The status of the Basic Endpoint.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public BasicEndpointState()
        {
        }
        public static new BasicEndpointState Empty => new BasicEndpointState();
    }
}
