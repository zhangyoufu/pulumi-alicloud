// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CloudStorageGateway
{
    /// <summary>
    /// Provides a Cloud Storage Gateway Gateway Block Volume resource.
    /// 
    /// For information about Cloud Storage Gateway Gateway Block Volume and how to use it, see [What is Gateway Block Volume](https://www.alibabacloud.com/help/en/doc-detail/53972.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.144.0+.
    /// 
    /// ## Import
    /// 
    /// Cloud Storage Gateway Gateway Block Volume can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cloudstoragegateway/gatewayBlockVolume:GatewayBlockVolume example &lt;gateway_id&gt;:&lt;index_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cloudstoragegateway/gatewayBlockVolume:GatewayBlockVolume")]
    public partial class GatewayBlockVolume : Pulumi.CustomResource
    {
        /// <summary>
        /// The Block volume set mode to cache mode. Valid values: `Cache`, `WriteThrough`.
        /// </summary>
        [Output("cacheMode")]
        public Output<string> CacheMode { get; private set; } = null!;

        /// <summary>
        /// Whether to enable iSCSI access of CHAP authentication, which currently supports both CHAP inbound authentication.  Default value: `false`.
        /// </summary>
        [Output("chapEnabled")]
        public Output<bool> ChapEnabled { get; private set; } = null!;

        /// <summary>
        /// The password for inbound authentication when the block volume enables iSCSI access to CHAP authentication. **NOTE:** When the `chap_enabled` is  `true` is,The `chap_in_password` is valid.
        /// </summary>
        [Output("chapInPassword")]
        public Output<string?> ChapInPassword { get; private set; } = null!;

        /// <summary>
        /// The Inbound CHAP user. The `chap_in_user` must be 1 to 32 characters in length, and can contain letters and digits. **NOTE:** When the `chap_enabled` is  `true` is,The `chap_in_password` is valid.
        /// </summary>
        [Output("chapInUser")]
        public Output<string?> ChapInUser { get; private set; } = null!;

        /// <summary>
        /// The Block volume storage allocation unit.  Valid values: `8192`, `16384`, `32768`, `65536`, `131072`. Default value: `32768`. Unit: `Byte`.
        /// </summary>
        [Output("chunkSize")]
        public Output<int> ChunkSize { get; private set; } = null!;

        /// <summary>
        /// The Block volume name. The name must be 1 to 32 characters in length, and can contain lower case letters and digits.
        /// </summary>
        [Output("gatewayBlockVolumeName")]
        public Output<string> GatewayBlockVolumeName { get; private set; } = null!;

        /// <summary>
        /// The Gateway ID.
        /// </summary>
        [Output("gatewayId")]
        public Output<string> GatewayId { get; private set; } = null!;

        /// <summary>
        /// The ID of the index.
        /// </summary>
        [Output("indexId")]
        public Output<string> IndexId { get; private set; } = null!;

        /// <summary>
        /// Whether to delete the source data. Default value `true`. **NOTE:** When `is_source_deletion` is `true`, the data in the OSS Bucket on the cloud is also deleted when deleting the block gateway volume. Please operate with caution.
        /// </summary>
        [Output("isSourceDeletion")]
        public Output<bool?> IsSourceDeletion { get; private set; } = null!;

        /// <summary>
        /// The Cache disk to local path. **NOTE:**  When the `cache_mode` is  `Cache` is,The `chap_in_password` is valid.
        /// </summary>
        [Output("localPath")]
        public Output<string?> LocalPath { get; private set; } = null!;

        /// <summary>
        /// The name of the OSS Bucket.
        /// </summary>
        [Output("ossBucketName")]
        public Output<string> OssBucketName { get; private set; } = null!;

        /// <summary>
        /// Whether to enable SSL access your OSS Buckets. Default value: `true`.
        /// </summary>
        [Output("ossBucketSsl")]
        public Output<bool> OssBucketSsl { get; private set; } = null!;

        /// <summary>
        /// The endpoint of the OSS Bucket.
        /// </summary>
        [Output("ossEndpoint")]
        public Output<string> OssEndpoint { get; private set; } = null!;

        /// <summary>
        /// The Protocol. Valid values: `iSCSI`.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The recovery.
        /// </summary>
        [Output("recovery")]
        public Output<bool?> Recovery { get; private set; } = null!;

        /// <summary>
        /// The Volume size. Valid values: `1` to `262144`. Unit: `Byte`.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// The status of volume. Valid values: 
        /// - `0`: Normal condition.
        /// - `1`: Failed to create volume.
        /// - `2`: Failed to delete volume.
        /// - `3`: Failed to enable target.
        /// - `4`: Failed to disable target.
        /// - `5`: Database error.
        /// - `6`: Failed to enable cache.
        /// - `7`: Failed to disable cache.
        /// - `8`: System error.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayBlockVolume resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayBlockVolume(string name, GatewayBlockVolumeArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cloudstoragegateway/gatewayBlockVolume:GatewayBlockVolume", name, args ?? new GatewayBlockVolumeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayBlockVolume(string name, Input<string> id, GatewayBlockVolumeState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cloudstoragegateway/gatewayBlockVolume:GatewayBlockVolume", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewayBlockVolume resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayBlockVolume Get(string name, Input<string> id, GatewayBlockVolumeState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayBlockVolume(name, id, state, options);
        }
    }

    public sealed class GatewayBlockVolumeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Block volume set mode to cache mode. Valid values: `Cache`, `WriteThrough`.
        /// </summary>
        [Input("cacheMode")]
        public Input<string>? CacheMode { get; set; }

        /// <summary>
        /// Whether to enable iSCSI access of CHAP authentication, which currently supports both CHAP inbound authentication.  Default value: `false`.
        /// </summary>
        [Input("chapEnabled")]
        public Input<bool>? ChapEnabled { get; set; }

        /// <summary>
        /// The password for inbound authentication when the block volume enables iSCSI access to CHAP authentication. **NOTE:** When the `chap_enabled` is  `true` is,The `chap_in_password` is valid.
        /// </summary>
        [Input("chapInPassword")]
        public Input<string>? ChapInPassword { get; set; }

        /// <summary>
        /// The Inbound CHAP user. The `chap_in_user` must be 1 to 32 characters in length, and can contain letters and digits. **NOTE:** When the `chap_enabled` is  `true` is,The `chap_in_password` is valid.
        /// </summary>
        [Input("chapInUser")]
        public Input<string>? ChapInUser { get; set; }

        /// <summary>
        /// The Block volume storage allocation unit.  Valid values: `8192`, `16384`, `32768`, `65536`, `131072`. Default value: `32768`. Unit: `Byte`.
        /// </summary>
        [Input("chunkSize")]
        public Input<int>? ChunkSize { get; set; }

        /// <summary>
        /// The Block volume name. The name must be 1 to 32 characters in length, and can contain lower case letters and digits.
        /// </summary>
        [Input("gatewayBlockVolumeName", required: true)]
        public Input<string> GatewayBlockVolumeName { get; set; } = null!;

        /// <summary>
        /// The Gateway ID.
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        /// <summary>
        /// Whether to delete the source data. Default value `true`. **NOTE:** When `is_source_deletion` is `true`, the data in the OSS Bucket on the cloud is also deleted when deleting the block gateway volume. Please operate with caution.
        /// </summary>
        [Input("isSourceDeletion")]
        public Input<bool>? IsSourceDeletion { get; set; }

        /// <summary>
        /// The Cache disk to local path. **NOTE:**  When the `cache_mode` is  `Cache` is,The `chap_in_password` is valid.
        /// </summary>
        [Input("localPath")]
        public Input<string>? LocalPath { get; set; }

        /// <summary>
        /// The name of the OSS Bucket.
        /// </summary>
        [Input("ossBucketName", required: true)]
        public Input<string> OssBucketName { get; set; } = null!;

        /// <summary>
        /// Whether to enable SSL access your OSS Buckets. Default value: `true`.
        /// </summary>
        [Input("ossBucketSsl")]
        public Input<bool>? OssBucketSsl { get; set; }

        /// <summary>
        /// The endpoint of the OSS Bucket.
        /// </summary>
        [Input("ossEndpoint", required: true)]
        public Input<string> OssEndpoint { get; set; } = null!;

        /// <summary>
        /// The Protocol. Valid values: `iSCSI`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The recovery.
        /// </summary>
        [Input("recovery")]
        public Input<bool>? Recovery { get; set; }

        /// <summary>
        /// The Volume size. Valid values: `1` to `262144`. Unit: `Byte`.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        public GatewayBlockVolumeArgs()
        {
        }
    }

    public sealed class GatewayBlockVolumeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Block volume set mode to cache mode. Valid values: `Cache`, `WriteThrough`.
        /// </summary>
        [Input("cacheMode")]
        public Input<string>? CacheMode { get; set; }

        /// <summary>
        /// Whether to enable iSCSI access of CHAP authentication, which currently supports both CHAP inbound authentication.  Default value: `false`.
        /// </summary>
        [Input("chapEnabled")]
        public Input<bool>? ChapEnabled { get; set; }

        /// <summary>
        /// The password for inbound authentication when the block volume enables iSCSI access to CHAP authentication. **NOTE:** When the `chap_enabled` is  `true` is,The `chap_in_password` is valid.
        /// </summary>
        [Input("chapInPassword")]
        public Input<string>? ChapInPassword { get; set; }

        /// <summary>
        /// The Inbound CHAP user. The `chap_in_user` must be 1 to 32 characters in length, and can contain letters and digits. **NOTE:** When the `chap_enabled` is  `true` is,The `chap_in_password` is valid.
        /// </summary>
        [Input("chapInUser")]
        public Input<string>? ChapInUser { get; set; }

        /// <summary>
        /// The Block volume storage allocation unit.  Valid values: `8192`, `16384`, `32768`, `65536`, `131072`. Default value: `32768`. Unit: `Byte`.
        /// </summary>
        [Input("chunkSize")]
        public Input<int>? ChunkSize { get; set; }

        /// <summary>
        /// The Block volume name. The name must be 1 to 32 characters in length, and can contain lower case letters and digits.
        /// </summary>
        [Input("gatewayBlockVolumeName")]
        public Input<string>? GatewayBlockVolumeName { get; set; }

        /// <summary>
        /// The Gateway ID.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// The ID of the index.
        /// </summary>
        [Input("indexId")]
        public Input<string>? IndexId { get; set; }

        /// <summary>
        /// Whether to delete the source data. Default value `true`. **NOTE:** When `is_source_deletion` is `true`, the data in the OSS Bucket on the cloud is also deleted when deleting the block gateway volume. Please operate with caution.
        /// </summary>
        [Input("isSourceDeletion")]
        public Input<bool>? IsSourceDeletion { get; set; }

        /// <summary>
        /// The Cache disk to local path. **NOTE:**  When the `cache_mode` is  `Cache` is,The `chap_in_password` is valid.
        /// </summary>
        [Input("localPath")]
        public Input<string>? LocalPath { get; set; }

        /// <summary>
        /// The name of the OSS Bucket.
        /// </summary>
        [Input("ossBucketName")]
        public Input<string>? OssBucketName { get; set; }

        /// <summary>
        /// Whether to enable SSL access your OSS Buckets. Default value: `true`.
        /// </summary>
        [Input("ossBucketSsl")]
        public Input<bool>? OssBucketSsl { get; set; }

        /// <summary>
        /// The endpoint of the OSS Bucket.
        /// </summary>
        [Input("ossEndpoint")]
        public Input<string>? OssEndpoint { get; set; }

        /// <summary>
        /// The Protocol. Valid values: `iSCSI`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The recovery.
        /// </summary>
        [Input("recovery")]
        public Input<bool>? Recovery { get; set; }

        /// <summary>
        /// The Volume size. Valid values: `1` to `262144`. Unit: `Byte`.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The status of volume. Valid values: 
        /// - `0`: Normal condition.
        /// - `1`: Failed to create volume.
        /// - `2`: Failed to delete volume.
        /// - `3`: Failed to enable target.
        /// - `4`: Failed to disable target.
        /// - `5`: Database error.
        /// - `6`: Failed to enable cache.
        /// - `7`: Failed to disable cache.
        /// - `8`: System error.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GatewayBlockVolumeState()
        {
        }
    }
}
