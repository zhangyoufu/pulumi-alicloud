// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hologram
{
    /// <summary>
    /// ## Import
    /// 
    /// Hologram Instance can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:hologram/instance:Instance example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:hologram/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to pay automatically. The default value is true. Value:
        /// - true: automatic payment
        /// - false: only generate orders, not pay
        /// &gt; **NOTE:**  The default value is true. If the balance of your payment method is insufficient, you can set the parameter AutoPay to false, and an unpaid order will be generated. You can log in to the user Center to pay by yourself.
        /// </summary>
        [Output("autoPay")]
        public Output<bool?> AutoPay { get; private set; } = null!;

        /// <summary>
        /// Instance low-frequency storage space. Unit: GB.
        /// &gt; **NOTE:**  PayAsYouGo (PostPaid) instances ignore this parameter.
        /// </summary>
        [Output("coldStorageSize")]
        public Output<int?> ColdStorageSize { get; private set; } = null!;

        /// <summary>
        /// Instance specifications. Value:
        /// - 8 cores 32 GB (number of compute nodes: 1)
        /// - 16 cores 64 GB (number of compute nodes: 1)
        /// - 32 core 128 GB (number of compute nodes: 2)
        /// - 64 core 256 GB (number of compute nodes: 4)
        /// - 96 core 384 GB (number of computing nodes: 6)
        /// - 128 core 512 GB (number of compute nodes: 8)
        /// &gt; **NOTE:** Just fill in the audit number. Please submit a work order application for purchasing 1024 or above specifications. Shared instance types do not need to specify specifications. The specification of - 8 core 32GB (number of computing nodes: 1) is only for experience use and cannot be used for production.
        /// </summary>
        [Output("cpu")]
        public Output<int> Cpu { get; private set; } = null!;

        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The buying cycle. Buy for 2 months. If the Payment type is PayAsYouGo (PostPaid), you do not need to specify it.
        /// </summary>
        [Output("duration")]
        public Output<int?> Duration { get; private set; } = null!;

        /// <summary>
        /// List of domain names. See `endpoints` below.
        /// </summary>
        [Output("endpoints")]
        public Output<ImmutableArray<Outputs.InstanceEndpoint>> Endpoints { get; private set; } = null!;

        /// <summary>
        /// Number of gateway nodes.
        /// </summary>
        [Output("gatewayCount")]
        public Output<int?> GatewayCount { get; private set; } = null!;

        /// <summary>
        /// Initialize the database and split multiple database names ",".
        /// </summary>
        [Output("initialDatabases")]
        public Output<string?> InitialDatabases { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// The instance type. Value:
        /// - Standard: Universal.
        /// - Follower: Read-only slave instance.
        /// - Warehouse: calculation group type.
        /// - Shared: Shared.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The ID of the primary instance.
        /// </summary>
        [Output("leaderInstanceId")]
        public Output<string?> LeaderInstanceId { get; private set; } = null!;

        /// <summary>
        /// The payment type of the resource.
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// Billing cycle. Value:
        /// - Month: monthly billing
        /// - Hour: hourly billing
        /// &gt; **NOTE:**  Subscription instances (PrePaid) only supports Month. PayAsYouGo instances (PostPaid) only supports Hour. The Shared type is automatically set to Hour without specifying it.
        /// </summary>
        [Output("pricingCycle")]
        public Output<string?> PricingCycle { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// Change matching type. Value:
        /// - UPGRADE: UPGRADE
        /// - DOWNGRADE: Downgrading
        /// &gt; **NOTE:** The upgrade specification cannot be less than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is larger than the original specification. The downgrading specification cannot be greater than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is smaller than the original specification.
        /// </summary>
        [Output("scaleType")]
        public Output<string?> ScaleType { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The standard storage space of the instance. Unit: GB.
        /// &gt; **NOTE:**  PayAsYouGo instances (PostPaid) ignore this parameter.
        /// </summary>
        [Output("storageSize")]
        public Output<int?> StorageSize { get; private set; } = null!;

        /// <summary>
        /// Instance tag.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The zone Id. Refer to "Instructions for Use".
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:hologram/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:hologram/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to pay automatically. The default value is true. Value:
        /// - true: automatic payment
        /// - false: only generate orders, not pay
        /// &gt; **NOTE:**  The default value is true. If the balance of your payment method is insufficient, you can set the parameter AutoPay to false, and an unpaid order will be generated. You can log in to the user Center to pay by yourself.
        /// </summary>
        [Input("autoPay")]
        public Input<bool>? AutoPay { get; set; }

        /// <summary>
        /// Instance low-frequency storage space. Unit: GB.
        /// &gt; **NOTE:**  PayAsYouGo (PostPaid) instances ignore this parameter.
        /// </summary>
        [Input("coldStorageSize")]
        public Input<int>? ColdStorageSize { get; set; }

        /// <summary>
        /// Instance specifications. Value:
        /// - 8 cores 32 GB (number of compute nodes: 1)
        /// - 16 cores 64 GB (number of compute nodes: 1)
        /// - 32 core 128 GB (number of compute nodes: 2)
        /// - 64 core 256 GB (number of compute nodes: 4)
        /// - 96 core 384 GB (number of computing nodes: 6)
        /// - 128 core 512 GB (number of compute nodes: 8)
        /// &gt; **NOTE:** Just fill in the audit number. Please submit a work order application for purchasing 1024 or above specifications. Shared instance types do not need to specify specifications. The specification of - 8 core 32GB (number of computing nodes: 1) is only for experience use and cannot be used for production.
        /// </summary>
        [Input("cpu")]
        public Input<int>? Cpu { get; set; }

        /// <summary>
        /// The buying cycle. Buy for 2 months. If the Payment type is PayAsYouGo (PostPaid), you do not need to specify it.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        [Input("endpoints")]
        private InputList<Inputs.InstanceEndpointArgs>? _endpoints;

        /// <summary>
        /// List of domain names. See `endpoints` below.
        /// </summary>
        public InputList<Inputs.InstanceEndpointArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.InstanceEndpointArgs>());
            set => _endpoints = value;
        }

        /// <summary>
        /// Number of gateway nodes.
        /// </summary>
        [Input("gatewayCount")]
        public Input<int>? GatewayCount { get; set; }

        /// <summary>
        /// Initialize the database and split multiple database names ",".
        /// </summary>
        [Input("initialDatabases")]
        public Input<string>? InitialDatabases { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// The instance type. Value:
        /// - Standard: Universal.
        /// - Follower: Read-only slave instance.
        /// - Warehouse: calculation group type.
        /// - Shared: Shared.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The ID of the primary instance.
        /// </summary>
        [Input("leaderInstanceId")]
        public Input<string>? LeaderInstanceId { get; set; }

        /// <summary>
        /// The payment type of the resource.
        /// </summary>
        [Input("paymentType", required: true)]
        public Input<string> PaymentType { get; set; } = null!;

        /// <summary>
        /// Billing cycle. Value:
        /// - Month: monthly billing
        /// - Hour: hourly billing
        /// &gt; **NOTE:**  Subscription instances (PrePaid) only supports Month. PayAsYouGo instances (PostPaid) only supports Hour. The Shared type is automatically set to Hour without specifying it.
        /// </summary>
        [Input("pricingCycle")]
        public Input<string>? PricingCycle { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Change matching type. Value:
        /// - UPGRADE: UPGRADE
        /// - DOWNGRADE: Downgrading
        /// &gt; **NOTE:** The upgrade specification cannot be less than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is larger than the original specification. The downgrading specification cannot be greater than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is smaller than the original specification.
        /// </summary>
        [Input("scaleType")]
        public Input<string>? ScaleType { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The standard storage space of the instance. Unit: GB.
        /// &gt; **NOTE:**  PayAsYouGo instances (PostPaid) ignore this parameter.
        /// </summary>
        [Input("storageSize")]
        public Input<int>? StorageSize { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Instance tag.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone Id. Refer to "Instructions for Use".
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to pay automatically. The default value is true. Value:
        /// - true: automatic payment
        /// - false: only generate orders, not pay
        /// &gt; **NOTE:**  The default value is true. If the balance of your payment method is insufficient, you can set the parameter AutoPay to false, and an unpaid order will be generated. You can log in to the user Center to pay by yourself.
        /// </summary>
        [Input("autoPay")]
        public Input<bool>? AutoPay { get; set; }

        /// <summary>
        /// Instance low-frequency storage space. Unit: GB.
        /// &gt; **NOTE:**  PayAsYouGo (PostPaid) instances ignore this parameter.
        /// </summary>
        [Input("coldStorageSize")]
        public Input<int>? ColdStorageSize { get; set; }

        /// <summary>
        /// Instance specifications. Value:
        /// - 8 cores 32 GB (number of compute nodes: 1)
        /// - 16 cores 64 GB (number of compute nodes: 1)
        /// - 32 core 128 GB (number of compute nodes: 2)
        /// - 64 core 256 GB (number of compute nodes: 4)
        /// - 96 core 384 GB (number of computing nodes: 6)
        /// - 128 core 512 GB (number of compute nodes: 8)
        /// &gt; **NOTE:** Just fill in the audit number. Please submit a work order application for purchasing 1024 or above specifications. Shared instance types do not need to specify specifications. The specification of - 8 core 32GB (number of computing nodes: 1) is only for experience use and cannot be used for production.
        /// </summary>
        [Input("cpu")]
        public Input<int>? Cpu { get; set; }

        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The buying cycle. Buy for 2 months. If the Payment type is PayAsYouGo (PostPaid), you do not need to specify it.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        [Input("endpoints")]
        private InputList<Inputs.InstanceEndpointGetArgs>? _endpoints;

        /// <summary>
        /// List of domain names. See `endpoints` below.
        /// </summary>
        public InputList<Inputs.InstanceEndpointGetArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.InstanceEndpointGetArgs>());
            set => _endpoints = value;
        }

        /// <summary>
        /// Number of gateway nodes.
        /// </summary>
        [Input("gatewayCount")]
        public Input<int>? GatewayCount { get; set; }

        /// <summary>
        /// Initialize the database and split multiple database names ",".
        /// </summary>
        [Input("initialDatabases")]
        public Input<string>? InitialDatabases { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The instance type. Value:
        /// - Standard: Universal.
        /// - Follower: Read-only slave instance.
        /// - Warehouse: calculation group type.
        /// - Shared: Shared.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The ID of the primary instance.
        /// </summary>
        [Input("leaderInstanceId")]
        public Input<string>? LeaderInstanceId { get; set; }

        /// <summary>
        /// The payment type of the resource.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// Billing cycle. Value:
        /// - Month: monthly billing
        /// - Hour: hourly billing
        /// &gt; **NOTE:**  Subscription instances (PrePaid) only supports Month. PayAsYouGo instances (PostPaid) only supports Hour. The Shared type is automatically set to Hour without specifying it.
        /// </summary>
        [Input("pricingCycle")]
        public Input<string>? PricingCycle { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Change matching type. Value:
        /// - UPGRADE: UPGRADE
        /// - DOWNGRADE: Downgrading
        /// &gt; **NOTE:** The upgrade specification cannot be less than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is larger than the original specification. The downgrading specification cannot be greater than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is smaller than the original specification.
        /// </summary>
        [Input("scaleType")]
        public Input<string>? ScaleType { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The standard storage space of the instance. Unit: GB.
        /// &gt; **NOTE:**  PayAsYouGo instances (PostPaid) ignore this parameter.
        /// </summary>
        [Input("storageSize")]
        public Input<int>? StorageSize { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Instance tag.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone Id. Refer to "Instructions for Use".
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
