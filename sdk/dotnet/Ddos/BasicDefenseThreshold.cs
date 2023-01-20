// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ddos
{
    /// <summary>
    /// Provides a Ddos Basic defense threshold resource.
    /// 
    /// For information about Ddos Basic Antiddos and how to use it, see [What is Defense Threshold](https://www.alibabacloud.com/help/en/ddos-protection/latest/modifydefensethreshold).
    /// 
    /// &gt; **NOTE:** Available in v1.168.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new AliCloud.Ecs.EipAddress("default", new()
    ///     {
    ///         AddressName = @var.Name,
    ///         Isp = "BGP",
    ///         InternetChargeType = "PayByBandwidth",
    ///         PaymentType = "PayAsYouGo",
    ///     });
    /// 
    ///     var example = new AliCloud.Ddos.BasicDefenseThreshold("example", new()
    ///     {
    ///         InstanceId = @default.Id,
    ///         DdosType = "defense",
    ///         InstanceType = "eip",
    ///         Bps = 390,
    ///         Pps = 90000,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Ddos Basic Antiddos can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ddos/basicDefenseThreshold:BasicDefenseThreshold example &lt;instance_id&gt;:&lt;instance_type&gt;:&lt;ddos_type&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ddos/basicDefenseThreshold:BasicDefenseThreshold")]
    public partial class BasicDefenseThreshold : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
        /// </summary>
        [Output("bps")]
        public Output<int> Bps { get; private set; } = null!;

        /// <summary>
        /// The type of the threshold to query. Valid values: `defense`,`blackhole`.
        /// </summary>
        [Output("ddosType")]
        public Output<string> DdosType { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The instance type of the public IP address asset. Value: `ecs`,`slb`,`eip`.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The Internet IP address.
        /// </summary>
        [Output("internetIp")]
        public Output<string> InternetIp { get; private set; } = null!;

        /// <summary>
        /// Whether it is the system default threshold. Value:
        /// </summary>
        [Output("isAuto")]
        public Output<bool> IsAuto { get; private set; } = null!;

        /// <summary>
        /// The maximum traffic scrubbing threshold. Unit: Mbit/s.
        /// </summary>
        [Output("maxBps")]
        public Output<int> MaxBps { get; private set; } = null!;

        /// <summary>
        /// The maximum packet scrubbing threshold. Unit: pps.
        /// </summary>
        [Output("maxPps")]
        public Output<int> MaxPps { get; private set; } = null!;

        /// <summary>
        /// The current message number cleaning threshold. Unit: pps.
        /// </summary>
        [Output("pps")]
        public Output<int> Pps { get; private set; } = null!;


        /// <summary>
        /// Create a BasicDefenseThreshold resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BasicDefenseThreshold(string name, BasicDefenseThresholdArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ddos/basicDefenseThreshold:BasicDefenseThreshold", name, args ?? new BasicDefenseThresholdArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BasicDefenseThreshold(string name, Input<string> id, BasicDefenseThresholdState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ddos/basicDefenseThreshold:BasicDefenseThreshold", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BasicDefenseThreshold resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BasicDefenseThreshold Get(string name, Input<string> id, BasicDefenseThresholdState? state = null, CustomResourceOptions? options = null)
        {
            return new BasicDefenseThreshold(name, id, state, options);
        }
    }

    public sealed class BasicDefenseThresholdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
        /// </summary>
        [Input("bps")]
        public Input<int>? Bps { get; set; }

        /// <summary>
        /// The type of the threshold to query. Valid values: `defense`,`blackhole`.
        /// </summary>
        [Input("ddosType", required: true)]
        public Input<string> DdosType { get; set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The instance type of the public IP address asset. Value: `ecs`,`slb`,`eip`.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The Internet IP address.
        /// </summary>
        [Input("internetIp")]
        public Input<string>? InternetIp { get; set; }

        /// <summary>
        /// Whether it is the system default threshold. Value:
        /// </summary>
        [Input("isAuto")]
        public Input<bool>? IsAuto { get; set; }

        /// <summary>
        /// The current message number cleaning threshold. Unit: pps.
        /// </summary>
        [Input("pps")]
        public Input<int>? Pps { get; set; }

        public BasicDefenseThresholdArgs()
        {
        }
        public static new BasicDefenseThresholdArgs Empty => new BasicDefenseThresholdArgs();
    }

    public sealed class BasicDefenseThresholdState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
        /// </summary>
        [Input("bps")]
        public Input<int>? Bps { get; set; }

        /// <summary>
        /// The type of the threshold to query. Valid values: `defense`,`blackhole`.
        /// </summary>
        [Input("ddosType")]
        public Input<string>? DdosType { get; set; }

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The instance type of the public IP address asset. Value: `ecs`,`slb`,`eip`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The Internet IP address.
        /// </summary>
        [Input("internetIp")]
        public Input<string>? InternetIp { get; set; }

        /// <summary>
        /// Whether it is the system default threshold. Value:
        /// </summary>
        [Input("isAuto")]
        public Input<bool>? IsAuto { get; set; }

        /// <summary>
        /// The maximum traffic scrubbing threshold. Unit: Mbit/s.
        /// </summary>
        [Input("maxBps")]
        public Input<int>? MaxBps { get; set; }

        /// <summary>
        /// The maximum packet scrubbing threshold. Unit: pps.
        /// </summary>
        [Input("maxPps")]
        public Input<int>? MaxPps { get; set; }

        /// <summary>
        /// The current message number cleaning threshold. Unit: pps.
        /// </summary>
        [Input("pps")]
        public Input<int>? Pps { get; set; }

        public BasicDefenseThresholdState()
        {
        }
        public static new BasicDefenseThresholdState Empty => new BasicDefenseThresholdState();
    }
}
