// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ens
{
    /// <summary>
    /// Provides a ENS Disk resource. The disk. When you use it for the first time, please contact the product classmates to add a resource whitelist.
    /// 
    /// For information about ENS Disk and how to use it, see [What is Disk](https://www.alibabacloud.com/help/en/ens/developer-reference/api-createdisk).
    /// 
    /// &gt; **NOTE:** Available since v1.213.0.
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var @default = new AliCloud.Ens.Disk("default", new()
    ///     {
    ///         Category = "cloud_ssd",
    ///         EnsRegionId = "cn-chongqing-11",
    ///         PaymentType = "PayAsYouGo",
    ///         Size = 20,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ENS Disk can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ens/disk:Disk example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ens/disk:Disk")]
    public partial class Disk : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Types of disk instancesValues: cloud_efficiency (high-efficiency cloud disk),cloud_ssd (full Flash cloud disk),local_hdd (local HDD),local_ssd (local ssd).
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// Disk instance creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Name of the disk instance.
        /// </summary>
        [Output("diskName")]
        public Output<string?> DiskName { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the cloud disk is Encrypted. If Encrypted = true, the default service key is used when KMSKeyId is not entered. Value range:`true`, `false`(default).
        /// </summary>
        [Output("encrypted")]
        public Output<bool?> Encrypted { get; private set; } = null!;

        /// <summary>
        /// Ens node IDExample value: cn-chengdu-telecom.
        /// </summary>
        [Output("ensRegionId")]
        public Output<string> EnsRegionId { get; private set; } = null!;

        /// <summary>
        /// The ID of the KMS key used by the cloud disk. If Encrypted is set to true, the service default key is used when KMSKeyId is empty.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// Billing type of the disk instanceValue: PayAsYouGo.
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The size of the disk instance. Unit: GiB.
        /// </summary>
        [Output("size")]
        public Output<int?> Size { get; private set; } = null!;

        /// <summary>
        /// The ID of the snapshot used to create the cloud disk.
        /// 
        /// The SnapshotId and Size parameters have the following limitations:
        /// - If the snapshot capacity corresponding to the **SnapshotId** parameter is greater than the specified **Size** parameter, the Size of the cloud disk created is the Size of the specified snapshot.
        /// - If the snapshot capacity corresponding to the **SnapshotId** parameter is less than the set **Size** parameter value, the Size of the cloud disk created is the specified **Size** parameter value.
        /// </summary>
        [Output("snapshotId")]
        public Output<string?> SnapshotId { get; private set; } = null!;

        /// <summary>
        /// Status of the disk instance:Value:In-use: In useAvailable: To be mountedAttaching: AttachingDetaching: uninstallingCreating: CreatingReIniting: Resetting.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Disk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Disk(string name, DiskArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ens/disk:Disk", name, args ?? new DiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Disk(string name, Input<string> id, DiskState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ens/disk:Disk", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Disk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Disk Get(string name, Input<string> id, DiskState? state = null, CustomResourceOptions? options = null)
        {
            return new Disk(name, id, state, options);
        }
    }

    public sealed class DiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Types of disk instancesValues: cloud_efficiency (high-efficiency cloud disk),cloud_ssd (full Flash cloud disk),local_hdd (local HDD),local_ssd (local ssd).
        /// </summary>
        [Input("category", required: true)]
        public Input<string> Category { get; set; } = null!;

        /// <summary>
        /// Name of the disk instance.
        /// </summary>
        [Input("diskName")]
        public Input<string>? DiskName { get; set; }

        /// <summary>
        /// Indicates whether the cloud disk is Encrypted. If Encrypted = true, the default service key is used when KMSKeyId is not entered. Value range:`true`, `false`(default).
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// Ens node IDExample value: cn-chengdu-telecom.
        /// </summary>
        [Input("ensRegionId", required: true)]
        public Input<string> EnsRegionId { get; set; } = null!;

        /// <summary>
        /// The ID of the KMS key used by the cloud disk. If Encrypted is set to true, the service default key is used when KMSKeyId is empty.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Billing type of the disk instanceValue: PayAsYouGo.
        /// </summary>
        [Input("paymentType", required: true)]
        public Input<string> PaymentType { get; set; } = null!;

        /// <summary>
        /// The size of the disk instance. Unit: GiB.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The ID of the snapshot used to create the cloud disk.
        /// 
        /// The SnapshotId and Size parameters have the following limitations:
        /// - If the snapshot capacity corresponding to the **SnapshotId** parameter is greater than the specified **Size** parameter, the Size of the cloud disk created is the Size of the specified snapshot.
        /// - If the snapshot capacity corresponding to the **SnapshotId** parameter is less than the set **Size** parameter value, the Size of the cloud disk created is the specified **Size** parameter value.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public DiskArgs()
        {
        }
        public static new DiskArgs Empty => new DiskArgs();
    }

    public sealed class DiskState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Types of disk instancesValues: cloud_efficiency (high-efficiency cloud disk),cloud_ssd (full Flash cloud disk),local_hdd (local HDD),local_ssd (local ssd).
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Disk instance creation time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Name of the disk instance.
        /// </summary>
        [Input("diskName")]
        public Input<string>? DiskName { get; set; }

        /// <summary>
        /// Indicates whether the cloud disk is Encrypted. If Encrypted = true, the default service key is used when KMSKeyId is not entered. Value range:`true`, `false`(default).
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// Ens node IDExample value: cn-chengdu-telecom.
        /// </summary>
        [Input("ensRegionId")]
        public Input<string>? EnsRegionId { get; set; }

        /// <summary>
        /// The ID of the KMS key used by the cloud disk. If Encrypted is set to true, the service default key is used when KMSKeyId is empty.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Billing type of the disk instanceValue: PayAsYouGo.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The size of the disk instance. Unit: GiB.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The ID of the snapshot used to create the cloud disk.
        /// 
        /// The SnapshotId and Size parameters have the following limitations:
        /// - If the snapshot capacity corresponding to the **SnapshotId** parameter is greater than the specified **Size** parameter, the Size of the cloud disk created is the Size of the specified snapshot.
        /// - If the snapshot capacity corresponding to the **SnapshotId** parameter is less than the set **Size** parameter value, the Size of the cloud disk created is the specified **Size** parameter value.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        /// <summary>
        /// Status of the disk instance:Value:In-use: In useAvailable: To be mountedAttaching: AttachingDetaching: uninstallingCreating: CreatingReIniting: Resetting.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public DiskState()
        {
        }
        public static new DiskState Empty => new DiskState();
    }
}
