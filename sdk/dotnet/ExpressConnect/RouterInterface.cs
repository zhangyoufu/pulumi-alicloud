// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ExpressConnect
{
    /// <summary>
    /// Provides a Express Connect Router Interface resource.
    /// 
    /// For information about Express Connect Router Interface and how to use it, see What is Router Interface.
    /// 
    /// &gt; **NOTE:** Available since v1.199.0.
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
    ///     var name = config.Get("name") ?? "tf_example";
    ///     var defaultNetwork = new AliCloud.Vpc.Network("default", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "172.16.0.0/12",
    ///     });
    /// 
    ///     var @default = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var defaultRouterInterface = new AliCloud.ExpressConnect.RouterInterface("default", new()
    ///     {
    ///         Description = name,
    ///         OppositeRegionId = @default.Apply(@default =&gt; @default.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)),
    ///         RouterId = defaultNetwork.RouterId,
    ///         Role = "InitiatingSide",
    ///         RouterType = "VRouter",
    ///         PaymentType = "PayAsYouGo",
    ///         RouterInterfaceName = name,
    ///         Spec = "Mini.2",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Express Connect Router Interface can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:expressconnect/routerInterface:RouterInterface example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:expressconnect/routerInterface:RouterInterface")]
    public partial class RouterInterface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access point ID to which the VBR belongs.
        /// </summary>
        [Output("accessPointId")]
        public Output<string?> AccessPointId { get; private set; } = null!;

        /// <summary>
        /// Whether to pay automatically, value:-**false** (default): automatic payment is not enabled. After generating an order, you need to complete the payment at the order center.-**true**: Enable automatic payment to automatically pay for orders.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
        /// </summary>
        [Output("autoPay")]
        public Output<bool?> AutoPay { get; private set; } = null!;

        /// <summary>
        /// The bandwidth of the resource.
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// The businessStatus of the resource. Valid Values: `Normal`, `FinancialLocked`, `SecurityLocked`.
        /// </summary>
        [Output("businessStatus")]
        public Output<string> BusinessStatus { get; private set; } = null!;

        /// <summary>
        /// The connected time of the resource.
        /// </summary>
        [Output("connectedTime")]
        public Output<string> ConnectedTime { get; private set; } = null!;

        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The cross border of the resource.
        /// </summary>
        [Output("crossBorder")]
        public Output<bool> CrossBorder { get; private set; } = null!;

        /// <summary>
        /// Whether to delete the health check IP address configured on the router interface. Value:-**true**: deletes the health check IP address.-**false** (default): does not delete the health check IP address.
        /// </summary>
        [Output("deleteHealthCheckIp")]
        public Output<bool?> DeleteHealthCheckIp { get; private set; } = null!;

        /// <summary>
        /// The description of the router interface. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The end time of the resource.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// The has reservation data of the resource.
        /// </summary>
        [Output("hasReservationData")]
        public Output<string> HasReservationData { get; private set; } = null!;

        /// <summary>
        /// The health check rate. Unit: seconds. The recommended value is 2. This indicates the interval between successive probe messages sent during the specified health check.
        /// </summary>
        [Output("hcRate")]
        public Output<int?> HcRate { get; private set; } = null!;

        /// <summary>
        /// The health check thresholds. Unit: pcs. The recommended value is 8. This indicates the number of probe messages to be sent during the specified health check.
        /// </summary>
        [Output("hcThreshold")]
        public Output<string?> HcThreshold { get; private set; } = null!;

        /// <summary>
        /// The health check source IP address, must be an unused IP within the local VPC.
        /// </summary>
        [Output("healthCheckSourceIp")]
        public Output<string?> HealthCheckSourceIp { get; private set; } = null!;

        /// <summary>
        /// The IP address for health screening purposes.
        /// </summary>
        [Output("healthCheckTargetIp")]
        public Output<string?> HealthCheckTargetIp { get; private set; } = null!;

        /// <summary>
        /// The Access point ID to which the other end belongs.
        /// </summary>
        [Output("oppositeAccessPointId")]
        public Output<string?> OppositeAccessPointId { get; private set; } = null!;

        /// <summary>
        /// The opposite bandwidth of the router on the other side.
        /// </summary>
        [Output("oppositeBandwidth")]
        public Output<int> OppositeBandwidth { get; private set; } = null!;

        /// <summary>
        /// The opposite interface business status of the router on the other side. Valid Values: `Normal`, `FinancialLocked`, `SecurityLocked`.
        /// </summary>
        [Output("oppositeInterfaceBusinessStatus")]
        public Output<string> OppositeInterfaceBusinessStatus { get; private set; } = null!;

        /// <summary>
        /// The Interface ID of the router at the other end.
        /// </summary>
        [Output("oppositeInterfaceId")]
        public Output<string?> OppositeInterfaceId { get; private set; } = null!;

        /// <summary>
        /// The AliCloud account ID of the owner of the router interface on the other end.
        /// </summary>
        [Output("oppositeInterfaceOwnerId")]
        public Output<string?> OppositeInterfaceOwnerId { get; private set; } = null!;

        /// <summary>
        /// The opposite interface spec of the router on the other side. Valid Values: `Mini.2`, `Mini.5`, `Mini.5`, `Small.2`, `Small.5`, `Middle.1`, `Middle.2`, `Middle.5`, `Large.1`, `Large.2`, `Large.5`, `XLarge.1`, `Negative`.
        /// </summary>
        [Output("oppositeInterfaceSpec")]
        public Output<string> OppositeInterfaceSpec { get; private set; } = null!;

        /// <summary>
        /// The opposite interface status of the router on the other side. Valid Values: `Idle`, `AcceptingConnecting`, `Connecting`, `Activating`, `Active`, `Modifying`, `Deactivating`, `Inactive`, `Deleting`.
        /// </summary>
        [Output("oppositeInterfaceStatus")]
        public Output<string> OppositeInterfaceStatus { get; private set; } = null!;

        /// <summary>
        /// The geographical ID of the location of the receiving end of the connection.
        /// </summary>
        [Output("oppositeRegionId")]
        public Output<string> OppositeRegionId { get; private set; } = null!;

        /// <summary>
        /// The id of the router at the other end.
        /// </summary>
        [Output("oppositeRouterId")]
        public Output<string?> OppositeRouterId { get; private set; } = null!;

        /// <summary>
        /// The opposite router type of the router on the other side. Valid Values: `VRouter`, `VBR`.
        /// </summary>
        [Output("oppositeRouterType")]
        public Output<string> OppositeRouterType { get; private set; } = null!;

        /// <summary>
        /// The opposite vpc instance id of the router on the other side.
        /// </summary>
        [Output("oppositeVpcInstanceId")]
        public Output<string> OppositeVpcInstanceId { get; private set; } = null!;

        /// <summary>
        /// The payment methods for router interfaces. Valid Values: `PayAsYouGo`, `Subscription`.
        /// </summary>
        [Output("paymentType")]
        public Output<string?> PaymentType { get; private set; } = null!;

        /// <summary>
        /// Purchase duration, value:-When you choose to pay on a monthly basis, the value range is **1 to 9 * *.-When you choose to pay per year, the value range is **1 to 3 * *.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// The billing cycle of the prepaid fee. Valid values:-**Month** (default): monthly payment.-**Year**: Pay per Year.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
        /// </summary>
        [Output("pricingCycle")]
        public Output<string?> PricingCycle { get; private set; } = null!;

        /// <summary>
        /// The reservation active time of the resource.
        /// </summary>
        [Output("reservationActiveTime")]
        public Output<string> ReservationActiveTime { get; private set; } = null!;

        /// <summary>
        /// The reservation bandwidth of the resource.
        /// </summary>
        [Output("reservationBandwidth")]
        public Output<string> ReservationBandwidth { get; private set; } = null!;

        /// <summary>
        /// The reservation internet charge type of the resource.
        /// </summary>
        [Output("reservationInternetChargeType")]
        public Output<string> ReservationInternetChargeType { get; private set; } = null!;

        /// <summary>
        /// The reservation order type of the resource.
        /// </summary>
        [Output("reservationOrderType")]
        public Output<string> ReservationOrderType { get; private set; } = null!;

        /// <summary>
        /// The role of the router interface. Valid Values: `InitiatingSide`, `AcceptingSide`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The router id associated with the router interface.
        /// </summary>
        [Output("routerId")]
        public Output<string> RouterId { get; private set; } = null!;

        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        [Output("routerInterfaceId")]
        public Output<string> RouterInterfaceId { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("routerInterfaceName")]
        public Output<string?> RouterInterfaceName { get; private set; } = null!;

        /// <summary>
        /// The type of router associated with the router interface. Valid Values: `VRouter`, `VBR`.
        /// </summary>
        [Output("routerType")]
        public Output<string> RouterType { get; private set; } = null!;

        /// <summary>
        /// The specification of the router interface. Valid Values: `Mini.2`, `Mini.5`, `Mini.5`, `Small.2`, `Small.5`, `Middle.1`, `Middle.2`, `Middle.5`, `Large.1`, `Large.2`, `Large.5`, `XLarge.1`, `Negative`.
        /// </summary>
        [Output("spec")]
        public Output<string> Spec { get; private set; } = null!;

        /// <summary>
        /// The status of the resource. Valid Values: `Idle`, `AcceptingConnecting`, `Connecting`, `Activating`, `Active`, `Modifying`, `Deactivating`, `Inactive`, `Deleting`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The vpc instance id of the resource.
        /// </summary>
        [Output("vpcInstanceId")]
        public Output<string> VpcInstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a RouterInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouterInterface(string name, RouterInterfaceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:expressconnect/routerInterface:RouterInterface", name, args ?? new RouterInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouterInterface(string name, Input<string> id, RouterInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:expressconnect/routerInterface:RouterInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouterInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouterInterface Get(string name, Input<string> id, RouterInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new RouterInterface(name, id, state, options);
        }
    }

    public sealed class RouterInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access point ID to which the VBR belongs.
        /// </summary>
        [Input("accessPointId")]
        public Input<string>? AccessPointId { get; set; }

        /// <summary>
        /// Whether to pay automatically, value:-**false** (default): automatic payment is not enabled. After generating an order, you need to complete the payment at the order center.-**true**: Enable automatic payment to automatically pay for orders.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
        /// </summary>
        [Input("autoPay")]
        public Input<bool>? AutoPay { get; set; }

        /// <summary>
        /// Whether to delete the health check IP address configured on the router interface. Value:-**true**: deletes the health check IP address.-**false** (default): does not delete the health check IP address.
        /// </summary>
        [Input("deleteHealthCheckIp")]
        public Input<bool>? DeleteHealthCheckIp { get; set; }

        /// <summary>
        /// The description of the router interface. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The health check rate. Unit: seconds. The recommended value is 2. This indicates the interval between successive probe messages sent during the specified health check.
        /// </summary>
        [Input("hcRate")]
        public Input<int>? HcRate { get; set; }

        /// <summary>
        /// The health check thresholds. Unit: pcs. The recommended value is 8. This indicates the number of probe messages to be sent during the specified health check.
        /// </summary>
        [Input("hcThreshold")]
        public Input<string>? HcThreshold { get; set; }

        /// <summary>
        /// The health check source IP address, must be an unused IP within the local VPC.
        /// </summary>
        [Input("healthCheckSourceIp")]
        public Input<string>? HealthCheckSourceIp { get; set; }

        /// <summary>
        /// The IP address for health screening purposes.
        /// </summary>
        [Input("healthCheckTargetIp")]
        public Input<string>? HealthCheckTargetIp { get; set; }

        /// <summary>
        /// The Access point ID to which the other end belongs.
        /// </summary>
        [Input("oppositeAccessPointId")]
        public Input<string>? OppositeAccessPointId { get; set; }

        /// <summary>
        /// The Interface ID of the router at the other end.
        /// </summary>
        [Input("oppositeInterfaceId")]
        public Input<string>? OppositeInterfaceId { get; set; }

        /// <summary>
        /// The AliCloud account ID of the owner of the router interface on the other end.
        /// </summary>
        [Input("oppositeInterfaceOwnerId")]
        public Input<string>? OppositeInterfaceOwnerId { get; set; }

        /// <summary>
        /// The geographical ID of the location of the receiving end of the connection.
        /// </summary>
        [Input("oppositeRegionId", required: true)]
        public Input<string> OppositeRegionId { get; set; } = null!;

        /// <summary>
        /// The id of the router at the other end.
        /// </summary>
        [Input("oppositeRouterId")]
        public Input<string>? OppositeRouterId { get; set; }

        /// <summary>
        /// The opposite router type of the router on the other side. Valid Values: `VRouter`, `VBR`.
        /// </summary>
        [Input("oppositeRouterType")]
        public Input<string>? OppositeRouterType { get; set; }

        /// <summary>
        /// The payment methods for router interfaces. Valid Values: `PayAsYouGo`, `Subscription`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// Purchase duration, value:-When you choose to pay on a monthly basis, the value range is **1 to 9 * *.-When you choose to pay per year, the value range is **1 to 3 * *.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The billing cycle of the prepaid fee. Valid values:-**Month** (default): monthly payment.-**Year**: Pay per Year.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
        /// </summary>
        [Input("pricingCycle")]
        public Input<string>? PricingCycle { get; set; }

        /// <summary>
        /// The role of the router interface. Valid Values: `InitiatingSide`, `AcceptingSide`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The router id associated with the router interface.
        /// </summary>
        [Input("routerId", required: true)]
        public Input<string> RouterId { get; set; } = null!;

        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        [Input("routerInterfaceId")]
        public Input<string>? RouterInterfaceId { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("routerInterfaceName")]
        public Input<string>? RouterInterfaceName { get; set; }

        /// <summary>
        /// The type of router associated with the router interface. Valid Values: `VRouter`, `VBR`.
        /// </summary>
        [Input("routerType", required: true)]
        public Input<string> RouterType { get; set; } = null!;

        /// <summary>
        /// The specification of the router interface. Valid Values: `Mini.2`, `Mini.5`, `Mini.5`, `Small.2`, `Small.5`, `Middle.1`, `Middle.2`, `Middle.5`, `Large.1`, `Large.2`, `Large.5`, `XLarge.1`, `Negative`.
        /// </summary>
        [Input("spec", required: true)]
        public Input<string> Spec { get; set; } = null!;

        /// <summary>
        /// The status of the resource. Valid Values: `Idle`, `AcceptingConnecting`, `Connecting`, `Activating`, `Active`, `Modifying`, `Deactivating`, `Inactive`, `Deleting`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public RouterInterfaceArgs()
        {
        }
        public static new RouterInterfaceArgs Empty => new RouterInterfaceArgs();
    }

    public sealed class RouterInterfaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access point ID to which the VBR belongs.
        /// </summary>
        [Input("accessPointId")]
        public Input<string>? AccessPointId { get; set; }

        /// <summary>
        /// Whether to pay automatically, value:-**false** (default): automatic payment is not enabled. After generating an order, you need to complete the payment at the order center.-**true**: Enable automatic payment to automatically pay for orders.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
        /// </summary>
        [Input("autoPay")]
        public Input<bool>? AutoPay { get; set; }

        /// <summary>
        /// The bandwidth of the resource.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The businessStatus of the resource. Valid Values: `Normal`, `FinancialLocked`, `SecurityLocked`.
        /// </summary>
        [Input("businessStatus")]
        public Input<string>? BusinessStatus { get; set; }

        /// <summary>
        /// The connected time of the resource.
        /// </summary>
        [Input("connectedTime")]
        public Input<string>? ConnectedTime { get; set; }

        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The cross border of the resource.
        /// </summary>
        [Input("crossBorder")]
        public Input<bool>? CrossBorder { get; set; }

        /// <summary>
        /// Whether to delete the health check IP address configured on the router interface. Value:-**true**: deletes the health check IP address.-**false** (default): does not delete the health check IP address.
        /// </summary>
        [Input("deleteHealthCheckIp")]
        public Input<bool>? DeleteHealthCheckIp { get; set; }

        /// <summary>
        /// The description of the router interface. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The end time of the resource.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// The has reservation data of the resource.
        /// </summary>
        [Input("hasReservationData")]
        public Input<string>? HasReservationData { get; set; }

        /// <summary>
        /// The health check rate. Unit: seconds. The recommended value is 2. This indicates the interval between successive probe messages sent during the specified health check.
        /// </summary>
        [Input("hcRate")]
        public Input<int>? HcRate { get; set; }

        /// <summary>
        /// The health check thresholds. Unit: pcs. The recommended value is 8. This indicates the number of probe messages to be sent during the specified health check.
        /// </summary>
        [Input("hcThreshold")]
        public Input<string>? HcThreshold { get; set; }

        /// <summary>
        /// The health check source IP address, must be an unused IP within the local VPC.
        /// </summary>
        [Input("healthCheckSourceIp")]
        public Input<string>? HealthCheckSourceIp { get; set; }

        /// <summary>
        /// The IP address for health screening purposes.
        /// </summary>
        [Input("healthCheckTargetIp")]
        public Input<string>? HealthCheckTargetIp { get; set; }

        /// <summary>
        /// The Access point ID to which the other end belongs.
        /// </summary>
        [Input("oppositeAccessPointId")]
        public Input<string>? OppositeAccessPointId { get; set; }

        /// <summary>
        /// The opposite bandwidth of the router on the other side.
        /// </summary>
        [Input("oppositeBandwidth")]
        public Input<int>? OppositeBandwidth { get; set; }

        /// <summary>
        /// The opposite interface business status of the router on the other side. Valid Values: `Normal`, `FinancialLocked`, `SecurityLocked`.
        /// </summary>
        [Input("oppositeInterfaceBusinessStatus")]
        public Input<string>? OppositeInterfaceBusinessStatus { get; set; }

        /// <summary>
        /// The Interface ID of the router at the other end.
        /// </summary>
        [Input("oppositeInterfaceId")]
        public Input<string>? OppositeInterfaceId { get; set; }

        /// <summary>
        /// The AliCloud account ID of the owner of the router interface on the other end.
        /// </summary>
        [Input("oppositeInterfaceOwnerId")]
        public Input<string>? OppositeInterfaceOwnerId { get; set; }

        /// <summary>
        /// The opposite interface spec of the router on the other side. Valid Values: `Mini.2`, `Mini.5`, `Mini.5`, `Small.2`, `Small.5`, `Middle.1`, `Middle.2`, `Middle.5`, `Large.1`, `Large.2`, `Large.5`, `XLarge.1`, `Negative`.
        /// </summary>
        [Input("oppositeInterfaceSpec")]
        public Input<string>? OppositeInterfaceSpec { get; set; }

        /// <summary>
        /// The opposite interface status of the router on the other side. Valid Values: `Idle`, `AcceptingConnecting`, `Connecting`, `Activating`, `Active`, `Modifying`, `Deactivating`, `Inactive`, `Deleting`.
        /// </summary>
        [Input("oppositeInterfaceStatus")]
        public Input<string>? OppositeInterfaceStatus { get; set; }

        /// <summary>
        /// The geographical ID of the location of the receiving end of the connection.
        /// </summary>
        [Input("oppositeRegionId")]
        public Input<string>? OppositeRegionId { get; set; }

        /// <summary>
        /// The id of the router at the other end.
        /// </summary>
        [Input("oppositeRouterId")]
        public Input<string>? OppositeRouterId { get; set; }

        /// <summary>
        /// The opposite router type of the router on the other side. Valid Values: `VRouter`, `VBR`.
        /// </summary>
        [Input("oppositeRouterType")]
        public Input<string>? OppositeRouterType { get; set; }

        /// <summary>
        /// The opposite vpc instance id of the router on the other side.
        /// </summary>
        [Input("oppositeVpcInstanceId")]
        public Input<string>? OppositeVpcInstanceId { get; set; }

        /// <summary>
        /// The payment methods for router interfaces. Valid Values: `PayAsYouGo`, `Subscription`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// Purchase duration, value:-When you choose to pay on a monthly basis, the value range is **1 to 9 * *.-When you choose to pay per year, the value range is **1 to 3 * *.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The billing cycle of the prepaid fee. Valid values:-**Month** (default): monthly payment.-**Year**: Pay per Year.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
        /// </summary>
        [Input("pricingCycle")]
        public Input<string>? PricingCycle { get; set; }

        /// <summary>
        /// The reservation active time of the resource.
        /// </summary>
        [Input("reservationActiveTime")]
        public Input<string>? ReservationActiveTime { get; set; }

        /// <summary>
        /// The reservation bandwidth of the resource.
        /// </summary>
        [Input("reservationBandwidth")]
        public Input<string>? ReservationBandwidth { get; set; }

        /// <summary>
        /// The reservation internet charge type of the resource.
        /// </summary>
        [Input("reservationInternetChargeType")]
        public Input<string>? ReservationInternetChargeType { get; set; }

        /// <summary>
        /// The reservation order type of the resource.
        /// </summary>
        [Input("reservationOrderType")]
        public Input<string>? ReservationOrderType { get; set; }

        /// <summary>
        /// The role of the router interface. Valid Values: `InitiatingSide`, `AcceptingSide`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The router id associated with the router interface.
        /// </summary>
        [Input("routerId")]
        public Input<string>? RouterId { get; set; }

        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        [Input("routerInterfaceId")]
        public Input<string>? RouterInterfaceId { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("routerInterfaceName")]
        public Input<string>? RouterInterfaceName { get; set; }

        /// <summary>
        /// The type of router associated with the router interface. Valid Values: `VRouter`, `VBR`.
        /// </summary>
        [Input("routerType")]
        public Input<string>? RouterType { get; set; }

        /// <summary>
        /// The specification of the router interface. Valid Values: `Mini.2`, `Mini.5`, `Mini.5`, `Small.2`, `Small.5`, `Middle.1`, `Middle.2`, `Middle.5`, `Large.1`, `Large.2`, `Large.5`, `XLarge.1`, `Negative`.
        /// </summary>
        [Input("spec")]
        public Input<string>? Spec { get; set; }

        /// <summary>
        /// The status of the resource. Valid Values: `Idle`, `AcceptingConnecting`, `Connecting`, `Activating`, `Active`, `Modifying`, `Deactivating`, `Inactive`, `Deleting`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The vpc instance id of the resource.
        /// </summary>
        [Input("vpcInstanceId")]
        public Input<string>? VpcInstanceId { get; set; }

        public RouterInterfaceState()
        {
        }
        public static new RouterInterfaceState Empty => new RouterInterfaceState();
    }
}
