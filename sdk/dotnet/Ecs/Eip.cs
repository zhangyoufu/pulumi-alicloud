// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    /// <summary>
    /// Provides an elastic IP resource.
    /// 
    /// &gt; **DEPRECATED:**  This resource  has been deprecated from version `1.126.0`. Please use new resource alicloud_eip_address.
    /// 
    /// &gt; **NOTE:** The resource only supports to create `PostPaid PayByTraffic`  or `PrePaid PayByBandwidth` elastic IP for international account. Otherwise, you will happened error `COMMODITY.INVALID_COMPONENT`.
    /// Your account is international if you can use it to login in [International Web Console](https://account.alibabacloud.com/login/login.htm).
    /// 
    /// &gt; **NOTE:** From version 1.10.1, this resource supports creating "PrePaid" EIP. In addition, it supports setting EIP name and description.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create a new EIP.
    ///     var example = new AliCloud.Ecs.Eip("example", new()
    ///     {
    ///         Bandwidth = "10",
    ///         InternetChargeType = "PayByBandwidth",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Module Support
    /// 
    /// You can use the existing eip module
    /// to create several EIP instances and associate them with other resources one-click, like ECS instances, SLB, Nat Gateway and so on.
    /// 
    /// ## Import
    /// 
    /// Elastic IP address can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ecs/eip:Eip example eip-abc12345678
    /// ```
    /// </summary>
    [Obsolete(@"This resource has been deprecated in favour of the EipAddress resource")]
    [AliCloudResourceType("alicloud:ecs/eip:Eip")]
    public partial class Eip : global::Pulumi.CustomResource
    {
        [Output("activityId")]
        public Output<string?> ActivityId { get; private set; } = null!;

        /// <summary>
        /// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        /// </summary>
        [Output("addressName")]
        public Output<string> AddressName { get; private set; } = null!;

        [Output("allocationId")]
        public Output<string?> AllocationId { get; private set; } = null!;

        [Output("autoPay")]
        public Output<bool?> AutoPay { get; private set; } = null!;

        /// <summary>
        /// Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
        /// </summary>
        [Output("bandwidth")]
        public Output<string> Bandwidth { get; private set; } = null!;

        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Whether enable the deletion protection or not. Default value: `false`.
        /// - true: Enable deletion protection.
        /// - false: Disable deletion protection.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("highDefinitionMonitorLogStatus")]
        public Output<string> HighDefinitionMonitorLogStatus { get; private set; } = null!;

        /// <summary>
        /// (It has been deprecated from version 1.126.0 and using new attribute `payment_type` instead) Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
        /// </summary>
        [Output("instanceChargeType")]
        public Output<string> InstanceChargeType { get; private set; } = null!;

        /// <summary>
        /// Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only "PayByBandwidth" when `instance_charge_type` is PrePaid.
        /// </summary>
        [Output("internetChargeType")]
        public Output<string> InternetChargeType { get; private set; } = null!;

        /// <summary>
        /// The elastic ip address
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
        /// </summary>
        [Output("isp")]
        public Output<string> Isp { get; private set; } = null!;

        [Output("logProject")]
        public Output<string?> LogProject { get; private set; } = null!;

        [Output("logStore")]
        public Output<string?> LogStore { get; private set; } = null!;

        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.126.0 and using new attribute `address_name` instead.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("netmode")]
        public Output<string> Netmode { get; private set; } = null!;

        /// <summary>
        /// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The duration that you will buy the resource, in month. It is valid when `instance_charge_type` is `PrePaid`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
        /// **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        [Output("pricingCycle")]
        public Output<string?> PricingCycle { get; private set; } = null!;

        [Output("publicIpAddressPoolId")]
        public Output<string?> PublicIpAddressPoolId { get; private set; } = null!;

        /// <summary>
        /// The Id of resource group which the eip belongs.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        [Output("securityProtectionTypes")]
        public Output<ImmutableArray<string>> SecurityProtectionTypes { get; private set; } = null!;

        /// <summary>
        /// The EIP current status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Eip resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Eip(string name, EipArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/eip:Eip", name, args ?? new EipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Eip(string name, Input<string> id, EipState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/eip:Eip", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Eip resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Eip Get(string name, Input<string> id, EipState? state = null, CustomResourceOptions? options = null)
        {
            return new Eip(name, id, state, options);
        }
    }

    public sealed class EipArgs : global::Pulumi.ResourceArgs
    {
        [Input("activityId")]
        public Input<string>? ActivityId { get; set; }

        /// <summary>
        /// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        /// </summary>
        [Input("addressName")]
        public Input<string>? AddressName { get; set; }

        [Input("allocationId")]
        public Input<string>? AllocationId { get; set; }

        [Input("autoPay")]
        public Input<bool>? AutoPay { get; set; }

        /// <summary>
        /// Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
        /// </summary>
        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        /// <summary>
        /// Whether enable the deletion protection or not. Default value: `false`.
        /// - true: Enable deletion protection.
        /// - false: Disable deletion protection.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("highDefinitionMonitorLogStatus")]
        public Input<string>? HighDefinitionMonitorLogStatus { get; set; }

        /// <summary>
        /// (It has been deprecated from version 1.126.0 and using new attribute `payment_type` instead) Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only "PayByBandwidth" when `instance_charge_type` is PrePaid.
        /// </summary>
        [Input("internetChargeType")]
        public Input<string>? InternetChargeType { get; set; }

        /// <summary>
        /// The elastic ip address
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
        /// </summary>
        [Input("isp")]
        public Input<string>? Isp { get; set; }

        [Input("logProject")]
        public Input<string>? LogProject { get; set; }

        [Input("logStore")]
        public Input<string>? LogStore { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.126.0 and using new attribute `address_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("netmode")]
        public Input<string>? Netmode { get; set; }

        /// <summary>
        /// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The duration that you will buy the resource, in month. It is valid when `instance_charge_type` is `PrePaid`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
        /// **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("pricingCycle")]
        public Input<string>? PricingCycle { get; set; }

        [Input("publicIpAddressPoolId")]
        public Input<string>? PublicIpAddressPoolId { get; set; }

        /// <summary>
        /// The Id of resource group which the eip belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityProtectionTypes")]
        private InputList<string>? _securityProtectionTypes;
        public InputList<string> SecurityProtectionTypes
        {
            get => _securityProtectionTypes ?? (_securityProtectionTypes = new InputList<string>());
            set => _securityProtectionTypes = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public EipArgs()
        {
        }
        public static new EipArgs Empty => new EipArgs();
    }

    public sealed class EipState : global::Pulumi.ResourceArgs
    {
        [Input("activityId")]
        public Input<string>? ActivityId { get; set; }

        /// <summary>
        /// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        /// </summary>
        [Input("addressName")]
        public Input<string>? AddressName { get; set; }

        [Input("allocationId")]
        public Input<string>? AllocationId { get; set; }

        [Input("autoPay")]
        public Input<bool>? AutoPay { get; set; }

        /// <summary>
        /// Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
        /// </summary>
        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Whether enable the deletion protection or not. Default value: `false`.
        /// - true: Enable deletion protection.
        /// - false: Disable deletion protection.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("highDefinitionMonitorLogStatus")]
        public Input<string>? HighDefinitionMonitorLogStatus { get; set; }

        /// <summary>
        /// (It has been deprecated from version 1.126.0 and using new attribute `payment_type` instead) Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only "PayByBandwidth" when `instance_charge_type` is PrePaid.
        /// </summary>
        [Input("internetChargeType")]
        public Input<string>? InternetChargeType { get; set; }

        /// <summary>
        /// The elastic ip address
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
        /// </summary>
        [Input("isp")]
        public Input<string>? Isp { get; set; }

        [Input("logProject")]
        public Input<string>? LogProject { get; set; }

        [Input("logStore")]
        public Input<string>? LogStore { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.126.0 and using new attribute `address_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("netmode")]
        public Input<string>? Netmode { get; set; }

        /// <summary>
        /// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The duration that you will buy the resource, in month. It is valid when `instance_charge_type` is `PrePaid`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
        /// **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("pricingCycle")]
        public Input<string>? PricingCycle { get; set; }

        [Input("publicIpAddressPoolId")]
        public Input<string>? PublicIpAddressPoolId { get; set; }

        /// <summary>
        /// The Id of resource group which the eip belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityProtectionTypes")]
        private InputList<string>? _securityProtectionTypes;
        public InputList<string> SecurityProtectionTypes
        {
            get => _securityProtectionTypes ?? (_securityProtectionTypes = new InputList<string>());
            set => _securityProtectionTypes = value;
        }

        /// <summary>
        /// The EIP current status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public EipState()
        {
        }
        public static new EipState Empty => new EipState();
    }
}
