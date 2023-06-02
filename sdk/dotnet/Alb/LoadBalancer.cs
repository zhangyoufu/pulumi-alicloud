// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb
{
    /// <summary>
    /// Provides a ALB Load Balancer resource.
    /// 
    /// For information about ALB Load Balancer and how to use it, see [What is Load Balancer](https://www.alibabacloud.com/help/doc-detail/197341.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.132.0+.
    /// 
    /// ## Import
    /// 
    /// ALB Load Balancer can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:alb/loadBalancer:LoadBalancer example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:alb/loadBalancer:LoadBalancer")]
    public partial class LoadBalancer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Access Logging Configuration Structure. See the following `Block access_log_config`.
        /// </summary>
        [Output("accessLogConfig")]
        public Output<Outputs.LoadBalancerAccessLogConfig?> AccessLogConfig { get; private set; } = null!;

        /// <summary>
        /// The method in which IP addresses are assigned. Valid values: `Fixed` and `Dynamic`. Default value: `Dynamic`.
        /// </summary>
        [Output("addressAllocatedMode")]
        public Output<string?> AddressAllocatedMode { get; private set; } = null!;

        /// <summary>
        /// The IP version. Valid values: `Ipv4`, `DualStack`.
        /// </summary>
        [Output("addressIpVersion")]
        public Output<string> AddressIpVersion { get; private set; } = null!;

        /// <summary>
        /// The type of IP address that the ALB instance uses to provide services. Valid values: `Intranet`, `Internet`. **NOTE:** From version 1.193.1, `address_type` can be modified.
        /// </summary>
        [Output("addressType")]
        public Output<string> AddressType { get; private set; } = null!;

        /// <summary>
        /// The deletion protection enabled. Valid values: `true` and `false`. Default value: `false`.
        /// </summary>
        [Output("deletionProtectionEnabled")]
        public Output<bool?> DeletionProtectionEnabled { get; private set; } = null!;

        /// <summary>
        /// The domain name of the ALB instance. **NOTE:** Available in v1.158.0+.
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to precheck the API request. Valid values: `true` and `false`.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The configuration of the billing method. See the following `Block load_balancer_billing_config`.
        /// </summary>
        [Output("loadBalancerBillingConfig")]
        public Output<Outputs.LoadBalancerLoadBalancerBillingConfig> LoadBalancerBillingConfig { get; private set; } = null!;

        /// <summary>
        /// The edition of the ALB instance. Different editions have different limits and billing methods. Valid values: `Basic`, `Standard` and `StandardWithWaf`(Available in v1.193.1+).
        /// </summary>
        [Output("loadBalancerEdition")]
        public Output<string> LoadBalancerEdition { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("loadBalancerName")]
        public Output<string> LoadBalancerName { get; private set; } = null!;

        /// <summary>
        /// Modify the Protection Configuration. See the following `Block modification_protection_config`.
        /// </summary>
        [Output("modificationProtectionConfig")]
        public Output<Outputs.LoadBalancerModificationProtectionConfig> ModificationProtectionConfig { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable the configuration read-only mode for the ALB instance. Valid values: `NonProtection` and `ConsoleProtection`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource. **NOTE:** The Key of `tags` cannot begin with "aliyun", "acs:", "http://", "https://", "ack" or "ingress".
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the virtual private cloud (VPC) where the ALB instance is deployed.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The zones and vSwitches. You must specify at least two zones. See the following `Block zone_mappings`.
        /// </summary>
        [Output("zoneMappings")]
        public Output<ImmutableArray<Outputs.LoadBalancerZoneMapping>> ZoneMappings { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancer(string name, LoadBalancerArgs args, CustomResourceOptions? options = null)
            : base("alicloud:alb/loadBalancer:LoadBalancer", name, args ?? new LoadBalancerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancer(string name, Input<string> id, LoadBalancerState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:alb/loadBalancer:LoadBalancer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancer Get(string name, Input<string> id, LoadBalancerState? state = null, CustomResourceOptions? options = null)
        {
            return new LoadBalancer(name, id, state, options);
        }
    }

    public sealed class LoadBalancerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Access Logging Configuration Structure. See the following `Block access_log_config`.
        /// </summary>
        [Input("accessLogConfig")]
        public Input<Inputs.LoadBalancerAccessLogConfigArgs>? AccessLogConfig { get; set; }

        /// <summary>
        /// The method in which IP addresses are assigned. Valid values: `Fixed` and `Dynamic`. Default value: `Dynamic`.
        /// </summary>
        [Input("addressAllocatedMode")]
        public Input<string>? AddressAllocatedMode { get; set; }

        /// <summary>
        /// The IP version. Valid values: `Ipv4`, `DualStack`.
        /// </summary>
        [Input("addressIpVersion")]
        public Input<string>? AddressIpVersion { get; set; }

        /// <summary>
        /// The type of IP address that the ALB instance uses to provide services. Valid values: `Intranet`, `Internet`. **NOTE:** From version 1.193.1, `address_type` can be modified.
        /// </summary>
        [Input("addressType", required: true)]
        public Input<string> AddressType { get; set; } = null!;

        /// <summary>
        /// The deletion protection enabled. Valid values: `true` and `false`. Default value: `false`.
        /// </summary>
        [Input("deletionProtectionEnabled")]
        public Input<bool>? DeletionProtectionEnabled { get; set; }

        /// <summary>
        /// Specifies whether to precheck the API request. Valid values: `true` and `false`.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The configuration of the billing method. See the following `Block load_balancer_billing_config`.
        /// </summary>
        [Input("loadBalancerBillingConfig", required: true)]
        public Input<Inputs.LoadBalancerLoadBalancerBillingConfigArgs> LoadBalancerBillingConfig { get; set; } = null!;

        /// <summary>
        /// The edition of the ALB instance. Different editions have different limits and billing methods. Valid values: `Basic`, `Standard` and `StandardWithWaf`(Available in v1.193.1+).
        /// </summary>
        [Input("loadBalancerEdition", required: true)]
        public Input<string> LoadBalancerEdition { get; set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("loadBalancerName", required: true)]
        public Input<string> LoadBalancerName { get; set; } = null!;

        /// <summary>
        /// Modify the Protection Configuration. See the following `Block modification_protection_config`.
        /// </summary>
        [Input("modificationProtectionConfig")]
        public Input<Inputs.LoadBalancerModificationProtectionConfigArgs>? ModificationProtectionConfig { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource. **NOTE:** The Key of `tags` cannot begin with "aliyun", "acs:", "http://", "https://", "ack" or "ingress".
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the virtual private cloud (VPC) where the ALB instance is deployed.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        [Input("zoneMappings", required: true)]
        private InputList<Inputs.LoadBalancerZoneMappingArgs>? _zoneMappings;

        /// <summary>
        /// The zones and vSwitches. You must specify at least two zones. See the following `Block zone_mappings`.
        /// </summary>
        public InputList<Inputs.LoadBalancerZoneMappingArgs> ZoneMappings
        {
            get => _zoneMappings ?? (_zoneMappings = new InputList<Inputs.LoadBalancerZoneMappingArgs>());
            set => _zoneMappings = value;
        }

        public LoadBalancerArgs()
        {
        }
        public static new LoadBalancerArgs Empty => new LoadBalancerArgs();
    }

    public sealed class LoadBalancerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Access Logging Configuration Structure. See the following `Block access_log_config`.
        /// </summary>
        [Input("accessLogConfig")]
        public Input<Inputs.LoadBalancerAccessLogConfigGetArgs>? AccessLogConfig { get; set; }

        /// <summary>
        /// The method in which IP addresses are assigned. Valid values: `Fixed` and `Dynamic`. Default value: `Dynamic`.
        /// </summary>
        [Input("addressAllocatedMode")]
        public Input<string>? AddressAllocatedMode { get; set; }

        /// <summary>
        /// The IP version. Valid values: `Ipv4`, `DualStack`.
        /// </summary>
        [Input("addressIpVersion")]
        public Input<string>? AddressIpVersion { get; set; }

        /// <summary>
        /// The type of IP address that the ALB instance uses to provide services. Valid values: `Intranet`, `Internet`. **NOTE:** From version 1.193.1, `address_type` can be modified.
        /// </summary>
        [Input("addressType")]
        public Input<string>? AddressType { get; set; }

        /// <summary>
        /// The deletion protection enabled. Valid values: `true` and `false`. Default value: `false`.
        /// </summary>
        [Input("deletionProtectionEnabled")]
        public Input<bool>? DeletionProtectionEnabled { get; set; }

        /// <summary>
        /// The domain name of the ALB instance. **NOTE:** Available in v1.158.0+.
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// Specifies whether to precheck the API request. Valid values: `true` and `false`.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The configuration of the billing method. See the following `Block load_balancer_billing_config`.
        /// </summary>
        [Input("loadBalancerBillingConfig")]
        public Input<Inputs.LoadBalancerLoadBalancerBillingConfigGetArgs>? LoadBalancerBillingConfig { get; set; }

        /// <summary>
        /// The edition of the ALB instance. Different editions have different limits and billing methods. Valid values: `Basic`, `Standard` and `StandardWithWaf`(Available in v1.193.1+).
        /// </summary>
        [Input("loadBalancerEdition")]
        public Input<string>? LoadBalancerEdition { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("loadBalancerName")]
        public Input<string>? LoadBalancerName { get; set; }

        /// <summary>
        /// Modify the Protection Configuration. See the following `Block modification_protection_config`.
        /// </summary>
        [Input("modificationProtectionConfig")]
        public Input<Inputs.LoadBalancerModificationProtectionConfigGetArgs>? ModificationProtectionConfig { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Specifies whether to enable the configuration read-only mode for the ALB instance. Valid values: `NonProtection` and `ConsoleProtection`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource. **NOTE:** The Key of `tags` cannot begin with "aliyun", "acs:", "http://", "https://", "ack" or "ingress".
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the virtual private cloud (VPC) where the ALB instance is deployed.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        [Input("zoneMappings")]
        private InputList<Inputs.LoadBalancerZoneMappingGetArgs>? _zoneMappings;

        /// <summary>
        /// The zones and vSwitches. You must specify at least two zones. See the following `Block zone_mappings`.
        /// </summary>
        public InputList<Inputs.LoadBalancerZoneMappingGetArgs> ZoneMappings
        {
            get => _zoneMappings ?? (_zoneMappings = new InputList<Inputs.LoadBalancerZoneMappingGetArgs>());
            set => _zoneMappings = value;
        }

        public LoadBalancerState()
        {
        }
        public static new LoadBalancerState Empty => new LoadBalancerState();
    }
}
