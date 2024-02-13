// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    /// <summary>
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var defaultCommonBandwithPackage = new AliCloud.Vpc.CommonBandwithPackage("defaultCommonBandwithPackage", new()
    ///     {
    ///         Bandwidth = "3",
    ///         InternetChargeType = "PayByBandwidth",
    ///     });
    /// 
    ///     var defaultEipAddress = new AliCloud.Ecs.EipAddress("defaultEipAddress", new()
    ///     {
    ///         Bandwidth = "3",
    ///         InternetChargeType = "PayByTraffic",
    ///     });
    /// 
    ///     var defaultCommonBandwithPackageAttachment = new AliCloud.Vpc.CommonBandwithPackageAttachment("defaultCommonBandwithPackageAttachment", new()
    ///     {
    ///         BandwidthPackageId = defaultCommonBandwithPackage.Id,
    ///         InstanceId = defaultEipAddress.Id,
    ///         BandwidthPackageBandwidth = "2",
    ///         IpType = "EIP",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// cbwp Common Bandwidth Package Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpc/commonBandwithPackageAttachment:CommonBandwithPackageAttachment example &lt;bandwidth_package_id&gt;:&lt;instance_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/commonBandwithPackageAttachment:CommonBandwithPackageAttachment")]
    public partial class CommonBandwithPackageAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
        /// </summary>
        [Output("bandwidthPackageBandwidth")]
        public Output<string> BandwidthPackageBandwidth { get; private set; } = null!;

        /// <summary>
        /// The bandwidth_package_id of the common bandwidth package attachment, the field can't be changed.
        /// </summary>
        [Output("bandwidthPackageId")]
        public Output<string> BandwidthPackageId { get; private set; } = null!;

        /// <summary>
        /// Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
        /// </summary>
        [Output("cancelCommonBandwidthPackageIpBandwidth")]
        public Output<bool?> CancelCommonBandwidthPackageIpBandwidth { get; private set; } = null!;

        /// <summary>
        /// The instance_id of the common bandwidth package attachment, the field can't be changed.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
        /// </summary>
        [Output("ipType")]
        public Output<string?> IpType { get; private set; } = null!;

        /// <summary>
        /// The status of the Internet Shared Bandwidth instance.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a CommonBandwithPackageAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CommonBandwithPackageAttachment(string name, CommonBandwithPackageAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/commonBandwithPackageAttachment:CommonBandwithPackageAttachment", name, args ?? new CommonBandwithPackageAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CommonBandwithPackageAttachment(string name, Input<string> id, CommonBandwithPackageAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/commonBandwithPackageAttachment:CommonBandwithPackageAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CommonBandwithPackageAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CommonBandwithPackageAttachment Get(string name, Input<string> id, CommonBandwithPackageAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new CommonBandwithPackageAttachment(name, id, state, options);
        }
    }

    public sealed class CommonBandwithPackageAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
        /// </summary>
        [Input("bandwidthPackageBandwidth")]
        public Input<string>? BandwidthPackageBandwidth { get; set; }

        /// <summary>
        /// The bandwidth_package_id of the common bandwidth package attachment, the field can't be changed.
        /// </summary>
        [Input("bandwidthPackageId", required: true)]
        public Input<string> BandwidthPackageId { get; set; } = null!;

        /// <summary>
        /// Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
        /// </summary>
        [Input("cancelCommonBandwidthPackageIpBandwidth")]
        public Input<bool>? CancelCommonBandwidthPackageIpBandwidth { get; set; }

        /// <summary>
        /// The instance_id of the common bandwidth package attachment, the field can't be changed.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
        /// </summary>
        [Input("ipType")]
        public Input<string>? IpType { get; set; }

        public CommonBandwithPackageAttachmentArgs()
        {
        }
        public static new CommonBandwithPackageAttachmentArgs Empty => new CommonBandwithPackageAttachmentArgs();
    }

    public sealed class CommonBandwithPackageAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
        /// </summary>
        [Input("bandwidthPackageBandwidth")]
        public Input<string>? BandwidthPackageBandwidth { get; set; }

        /// <summary>
        /// The bandwidth_package_id of the common bandwidth package attachment, the field can't be changed.
        /// </summary>
        [Input("bandwidthPackageId")]
        public Input<string>? BandwidthPackageId { get; set; }

        /// <summary>
        /// Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
        /// </summary>
        [Input("cancelCommonBandwidthPackageIpBandwidth")]
        public Input<bool>? CancelCommonBandwidthPackageIpBandwidth { get; set; }

        /// <summary>
        /// The instance_id of the common bandwidth package attachment, the field can't be changed.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
        /// </summary>
        [Input("ipType")]
        public Input<string>? IpType { get; set; }

        /// <summary>
        /// The status of the Internet Shared Bandwidth instance.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public CommonBandwithPackageAttachmentState()
        {
        }
        public static new CommonBandwithPackageAttachmentState Empty => new CommonBandwithPackageAttachmentState();
    }
}
