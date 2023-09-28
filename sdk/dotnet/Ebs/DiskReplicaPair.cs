// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ebs
{
    /// <summary>
    /// Provides a Ebs Disk Replica Pair resource.
    /// 
    /// For information about Ebs Disk Replica Pair and how to use it, see [What is Disk Replica Pair](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/CreateDiskReplicaPair).
    /// 
    /// &gt; **NOTE:** Available since v1.196.0.
    /// 
    /// ## Import
    /// 
    /// Ebs Disk Replica Pair can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ebs/diskReplicaPair:DiskReplicaPair example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ebs/diskReplicaPair:DiskReplicaPair")]
    public partial class DiskReplicaPair : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The bandwidth for asynchronous data replication between cloud disks. The unit is Kbps. Value range:-10240 Kbps: equal to 10 Mbps.-20480 Kbps: equal to 20 Mbps.-51200 Kbps: equal to 50 Mbps.-102400 Kbps: equal to 100 Mbps.Default value: 10240.This parameter cannot be specified when the ChargeType value is POSTPAY. The system value is 0, which indicates that the disk is dynamically allocated according to data write changes during asynchronous replication.
        /// </summary>
        [Output("bandwidth")]
        public Output<string> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// The creation time of the resource
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The description of the asynchronous replication relationship. 2 to 256 English or Chinese characters in length and cannot start with' http:// 'or' https.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the standby disk.
        /// </summary>
        [Output("destinationDiskId")]
        public Output<string> DestinationDiskId { get; private set; } = null!;

        /// <summary>
        /// The ID of the region to which the disaster recovery site belongs.
        /// </summary>
        [Output("destinationRegionId")]
        public Output<string> DestinationRegionId { get; private set; } = null!;

        /// <summary>
        /// The ID of the zone to which the disaster recovery site belongs.
        /// </summary>
        [Output("destinationZoneId")]
        public Output<string> DestinationZoneId { get; private set; } = null!;

        /// <summary>
        /// The ID of the primary disk.
        /// </summary>
        [Output("diskId")]
        public Output<string> DiskId { get; private set; } = null!;

        /// <summary>
        /// The name of the asynchronous replication relationship. The length must be 2 to 128 characters in length and must start with a letter or Chinese name. It cannot start with http:// or https. It can contain Chinese, English, numbers, half-width colons (:), underscores (_), half-width periods (.), or dashes (-).
        /// </summary>
        [Output("pairName")]
        public Output<string?> PairName { get; private set; } = null!;

        /// <summary>
        /// The payment type of the resource
        /// </summary>
        [Output("paymentType")]
        public Output<string?> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The length of the purchase for the asynchronous replication relationship. When ChargeType=PrePay, this parameter is mandatory. The unit of duration is specified by PeriodUnit and takes on a range of values. When PeriodUnit=Week, this parameter takes values in the range `1`, `2`, `3` and `4`. When PeriodUnit=Month, the parameter takes on the values `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `12`, `24`, `36`, `48`, `60`.
        /// </summary>
        [Output("period")]
        public Output<string?> Period { get; private set; } = null!;

        /// <summary>
        /// The units of asynchronous replication relationship purchase length. Valid values: `Week` and `Month`. Default value: `Month`.
        /// </summary>
        [Output("periodUnit")]
        public Output<string?> PeriodUnit { get; private set; } = null!;

        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        [Output("replicaPairId")]
        public Output<string> ReplicaPairId { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The RPO value set by the consistency group in seconds. Currently only 900 seconds are supported.
        /// </summary>
        [Output("rpo")]
        public Output<string> Rpo { get; private set; } = null!;

        /// <summary>
        /// The ID of the zone to which the production site belongs.
        /// </summary>
        [Output("sourceZoneId")]
        public Output<string> SourceZoneId { get; private set; } = null!;

        /// <summary>
        /// The status of the resource
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a DiskReplicaPair resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DiskReplicaPair(string name, DiskReplicaPairArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ebs/diskReplicaPair:DiskReplicaPair", name, args ?? new DiskReplicaPairArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DiskReplicaPair(string name, Input<string> id, DiskReplicaPairState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ebs/diskReplicaPair:DiskReplicaPair", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DiskReplicaPair resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DiskReplicaPair Get(string name, Input<string> id, DiskReplicaPairState? state = null, CustomResourceOptions? options = null)
        {
            return new DiskReplicaPair(name, id, state, options);
        }
    }

    public sealed class DiskReplicaPairArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth for asynchronous data replication between cloud disks. The unit is Kbps. Value range:-10240 Kbps: equal to 10 Mbps.-20480 Kbps: equal to 20 Mbps.-51200 Kbps: equal to 50 Mbps.-102400 Kbps: equal to 100 Mbps.Default value: 10240.This parameter cannot be specified when the ChargeType value is POSTPAY. The system value is 0, which indicates that the disk is dynamically allocated according to data write changes during asynchronous replication.
        /// </summary>
        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        /// <summary>
        /// The description of the asynchronous replication relationship. 2 to 256 English or Chinese characters in length and cannot start with' http:// 'or' https.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the standby disk.
        /// </summary>
        [Input("destinationDiskId", required: true)]
        public Input<string> DestinationDiskId { get; set; } = null!;

        /// <summary>
        /// The ID of the region to which the disaster recovery site belongs.
        /// </summary>
        [Input("destinationRegionId", required: true)]
        public Input<string> DestinationRegionId { get; set; } = null!;

        /// <summary>
        /// The ID of the zone to which the disaster recovery site belongs.
        /// </summary>
        [Input("destinationZoneId", required: true)]
        public Input<string> DestinationZoneId { get; set; } = null!;

        /// <summary>
        /// The ID of the primary disk.
        /// </summary>
        [Input("diskId", required: true)]
        public Input<string> DiskId { get; set; } = null!;

        /// <summary>
        /// The name of the asynchronous replication relationship. The length must be 2 to 128 characters in length and must start with a letter or Chinese name. It cannot start with http:// or https. It can contain Chinese, English, numbers, half-width colons (:), underscores (_), half-width periods (.), or dashes (-).
        /// </summary>
        [Input("pairName")]
        public Input<string>? PairName { get; set; }

        /// <summary>
        /// The payment type of the resource
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The length of the purchase for the asynchronous replication relationship. When ChargeType=PrePay, this parameter is mandatory. The unit of duration is specified by PeriodUnit and takes on a range of values. When PeriodUnit=Week, this parameter takes values in the range `1`, `2`, `3` and `4`. When PeriodUnit=Month, the parameter takes on the values `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `12`, `24`, `36`, `48`, `60`.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        /// <summary>
        /// The units of asynchronous replication relationship purchase length. Valid values: `Week` and `Month`. Default value: `Month`.
        /// </summary>
        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        [Input("replicaPairId")]
        public Input<string>? ReplicaPairId { get; set; }

        /// <summary>
        /// The RPO value set by the consistency group in seconds. Currently only 900 seconds are supported.
        /// </summary>
        [Input("rpo")]
        public Input<string>? Rpo { get; set; }

        /// <summary>
        /// The ID of the zone to which the production site belongs.
        /// </summary>
        [Input("sourceZoneId", required: true)]
        public Input<string> SourceZoneId { get; set; } = null!;

        public DiskReplicaPairArgs()
        {
        }
        public static new DiskReplicaPairArgs Empty => new DiskReplicaPairArgs();
    }

    public sealed class DiskReplicaPairState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth for asynchronous data replication between cloud disks. The unit is Kbps. Value range:-10240 Kbps: equal to 10 Mbps.-20480 Kbps: equal to 20 Mbps.-51200 Kbps: equal to 50 Mbps.-102400 Kbps: equal to 100 Mbps.Default value: 10240.This parameter cannot be specified when the ChargeType value is POSTPAY. The system value is 0, which indicates that the disk is dynamically allocated according to data write changes during asynchronous replication.
        /// </summary>
        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        /// <summary>
        /// The creation time of the resource
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The description of the asynchronous replication relationship. 2 to 256 English or Chinese characters in length and cannot start with' http:// 'or' https.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the standby disk.
        /// </summary>
        [Input("destinationDiskId")]
        public Input<string>? DestinationDiskId { get; set; }

        /// <summary>
        /// The ID of the region to which the disaster recovery site belongs.
        /// </summary>
        [Input("destinationRegionId")]
        public Input<string>? DestinationRegionId { get; set; }

        /// <summary>
        /// The ID of the zone to which the disaster recovery site belongs.
        /// </summary>
        [Input("destinationZoneId")]
        public Input<string>? DestinationZoneId { get; set; }

        /// <summary>
        /// The ID of the primary disk.
        /// </summary>
        [Input("diskId")]
        public Input<string>? DiskId { get; set; }

        /// <summary>
        /// The name of the asynchronous replication relationship. The length must be 2 to 128 characters in length and must start with a letter or Chinese name. It cannot start with http:// or https. It can contain Chinese, English, numbers, half-width colons (:), underscores (_), half-width periods (.), or dashes (-).
        /// </summary>
        [Input("pairName")]
        public Input<string>? PairName { get; set; }

        /// <summary>
        /// The payment type of the resource
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The length of the purchase for the asynchronous replication relationship. When ChargeType=PrePay, this parameter is mandatory. The unit of duration is specified by PeriodUnit and takes on a range of values. When PeriodUnit=Week, this parameter takes values in the range `1`, `2`, `3` and `4`. When PeriodUnit=Month, the parameter takes on the values `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `12`, `24`, `36`, `48`, `60`.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        /// <summary>
        /// The units of asynchronous replication relationship purchase length. Valid values: `Week` and `Month`. Default value: `Month`.
        /// </summary>
        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        [Input("replicaPairId")]
        public Input<string>? ReplicaPairId { get; set; }

        /// <summary>
        /// The ID of the resource group
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The RPO value set by the consistency group in seconds. Currently only 900 seconds are supported.
        /// </summary>
        [Input("rpo")]
        public Input<string>? Rpo { get; set; }

        /// <summary>
        /// The ID of the zone to which the production site belongs.
        /// </summary>
        [Input("sourceZoneId")]
        public Input<string>? SourceZoneId { get; set; }

        /// <summary>
        /// The status of the resource
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public DiskReplicaPairState()
        {
        }
        public static new DiskReplicaPairState Empty => new DiskReplicaPairState();
    }
}
