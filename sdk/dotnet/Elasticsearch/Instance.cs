// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ElasticSearch
{
    /// <summary>
    /// Provides a Elasticsearch instance resource. It contains data nodes, dedicated master node(optional) and etc. It can be associated with private IP whitelists and kibana IP whitelist.
    /// 
    /// &gt; **NOTE:** Only one operation is supported in a request. So if `data_node_spec` and `data_node_disk_size` are both changed, system will respond error.
    /// 
    /// &gt; **NOTE:** At present, `version` can not be modified once instance has been created.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/elasticsearch_instance.html.markdown.
    /// </summary>
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// The Elasticsearch cluster's data node quantity, between 2 and 50.
        /// </summary>
        [Output("dataNodeAmount")]
        public Output<int> DataNodeAmount { get; private set; } = null!;

        /// <summary>
        /// The single data node storage space.
        /// - `cloud_ssd`: An SSD disk, supports a maximum of 2048 GiB (2 TB).
        /// - `cloud_efficiency` An ultra disk, supports a maximum of 5120 GiB (5 TB). If the data to be stored is larger than 2048 GiB, an ultra disk can only support the following data sizes (GiB): [`2560`, `3072`, `3584`, `4096`, `4608`, `5120`].
        /// </summary>
        [Output("dataNodeDiskSize")]
        public Output<int> DataNodeDiskSize { get; private set; } = null!;

        /// <summary>
        /// The data node disk type. Supported values: cloud_ssd, cloud_efficiency.
        /// </summary>
        [Output("dataNodeDiskType")]
        public Output<string> DataNodeDiskType { get; private set; } = null!;

        /// <summary>
        /// The data node specifications of the Elasticsearch instance.
        /// </summary>
        [Output("dataNodeSpec")]
        public Output<string> DataNodeSpec { get; private set; } = null!;

        /// <summary>
        /// The description of instance. It a string of 0 to 30 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Instance connection domain (only VPC network access supported).
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. From version 1.69.0, the Elasticsearch cluster allows you to update your instance_charge_ype from `PostPaid` to `PrePaid`, the following attributes are required: `period`. But, updating from `PostPaid` to `PrePaid` is not supported.
        /// </summary>
        [Output("instanceChargeType")]
        public Output<string?> InstanceChargeType { get; private set; } = null!;

        /// <summary>
        /// Kibana console domain (Internet access supported).
        /// </summary>
        [Output("kibanaDomain")]
        public Output<string> KibanaDomain { get; private set; } = null!;

        /// <summary>
        /// Kibana console port.
        /// </summary>
        [Output("kibanaPort")]
        public Output<int> KibanaPort { get; private set; } = null!;

        /// <summary>
        /// Set the Kibana's IP whitelist in internet network.
        /// </summary>
        [Output("kibanaWhitelists")]
        public Output<ImmutableArray<string>> KibanaWhitelists { get; private set; } = null!;

        /// <summary>
        /// An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored, but you have to specify one of `password` and `kms_encrypted_password` fields.
        /// </summary>
        [Output("kmsEncryptedPassword")]
        public Output<string?> KmsEncryptedPassword { get; private set; } = null!;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        [Output("kmsEncryptionContext")]
        public Output<ImmutableDictionary<string, object>?> KmsEncryptionContext { get; private set; } = null!;

        /// <summary>
        /// The dedicated master node spec. If specified, dedicated master node will be created.
        /// </summary>
        [Output("masterNodeSpec")]
        public Output<string?> MasterNodeSpec { get; private set; } = null!;

        /// <summary>
        /// The password of the instance. The password can be 8 to 32 characters in length and must contain three of the following conditions: uppercase letters, lowercase letters, numbers, and special characters (!@#$%^&amp;*()_+-=).
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The duration that you will buy Elasticsearch instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// Instance connection port.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Set the instance's IP whitelist in VPC network.
        /// </summary>
        [Output("privateWhitelists")]
        public Output<ImmutableArray<string>> PrivateWhitelists { get; private set; } = null!;

        [Output("publicWhitelists")]
        public Output<ImmutableArray<string>> PublicWhitelists { get; private set; } = null!;

        /// <summary>
        /// The Elasticsearch instance status. Includes `active`, `activating`, `inactive`. Some operations are denied when status is not `active`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Elasticsearch version. Supported values: `5.5.3_with_X-Pack`, `6.3_with_X-Pack` and `6.7_with_X-Pack`.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The ID of VSwitch.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The Multi-AZ supported for Elasticsearch, between 1 and 3. The `data_node_amount` value must be an integral multiple of the `zone_count` value.
        /// </summary>
        [Output("zoneCount")]
        public Output<int?> ZoneCount { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:elasticsearch/instance:Instance", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:elasticsearch/instance:Instance", name, state, MakeResourceOptions(options, id))
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

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Elasticsearch cluster's data node quantity, between 2 and 50.
        /// </summary>
        [Input("dataNodeAmount", required: true)]
        public Input<int> DataNodeAmount { get; set; } = null!;

        /// <summary>
        /// The single data node storage space.
        /// - `cloud_ssd`: An SSD disk, supports a maximum of 2048 GiB (2 TB).
        /// - `cloud_efficiency` An ultra disk, supports a maximum of 5120 GiB (5 TB). If the data to be stored is larger than 2048 GiB, an ultra disk can only support the following data sizes (GiB): [`2560`, `3072`, `3584`, `4096`, `4608`, `5120`].
        /// </summary>
        [Input("dataNodeDiskSize", required: true)]
        public Input<int> DataNodeDiskSize { get; set; } = null!;

        /// <summary>
        /// The data node disk type. Supported values: cloud_ssd, cloud_efficiency.
        /// </summary>
        [Input("dataNodeDiskType", required: true)]
        public Input<string> DataNodeDiskType { get; set; } = null!;

        /// <summary>
        /// The data node specifications of the Elasticsearch instance.
        /// </summary>
        [Input("dataNodeSpec", required: true)]
        public Input<string> DataNodeSpec { get; set; } = null!;

        /// <summary>
        /// The description of instance. It a string of 0 to 30 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. From version 1.69.0, the Elasticsearch cluster allows you to update your instance_charge_ype from `PostPaid` to `PrePaid`, the following attributes are required: `period`. But, updating from `PostPaid` to `PrePaid` is not supported.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        [Input("kibanaWhitelists")]
        private InputList<string>? _kibanaWhitelists;

        /// <summary>
        /// Set the Kibana's IP whitelist in internet network.
        /// </summary>
        public InputList<string> KibanaWhitelists
        {
            get => _kibanaWhitelists ?? (_kibanaWhitelists = new InputList<string>());
            set => _kibanaWhitelists = value;
        }

        /// <summary>
        /// An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored, but you have to specify one of `password` and `kms_encrypted_password` fields.
        /// </summary>
        [Input("kmsEncryptedPassword")]
        public Input<string>? KmsEncryptedPassword { get; set; }

        [Input("kmsEncryptionContext")]
        private InputMap<object>? _kmsEncryptionContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        public InputMap<object> KmsEncryptionContext
        {
            get => _kmsEncryptionContext ?? (_kmsEncryptionContext = new InputMap<object>());
            set => _kmsEncryptionContext = value;
        }

        /// <summary>
        /// The dedicated master node spec. If specified, dedicated master node will be created.
        /// </summary>
        [Input("masterNodeSpec")]
        public Input<string>? MasterNodeSpec { get; set; }

        /// <summary>
        /// The password of the instance. The password can be 8 to 32 characters in length and must contain three of the following conditions: uppercase letters, lowercase letters, numbers, and special characters (!@#$%^&amp;*()_+-=).
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The duration that you will buy Elasticsearch instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("privateWhitelists")]
        private InputList<string>? _privateWhitelists;

        /// <summary>
        /// Set the instance's IP whitelist in VPC network.
        /// </summary>
        public InputList<string> PrivateWhitelists
        {
            get => _privateWhitelists ?? (_privateWhitelists = new InputList<string>());
            set => _privateWhitelists = value;
        }

        [Input("publicWhitelists")]
        private InputList<string>? _publicWhitelists;
        public InputList<string> PublicWhitelists
        {
            get => _publicWhitelists ?? (_publicWhitelists = new InputList<string>());
            set => _publicWhitelists = value;
        }

        /// <summary>
        /// Elasticsearch version. Supported values: `5.5.3_with_X-Pack`, `6.3_with_X-Pack` and `6.7_with_X-Pack`.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        /// <summary>
        /// The ID of VSwitch.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        /// <summary>
        /// The Multi-AZ supported for Elasticsearch, between 1 and 3. The `data_node_amount` value must be an integral multiple of the `zone_count` value.
        /// </summary>
        [Input("zoneCount")]
        public Input<int>? ZoneCount { get; set; }

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Elasticsearch cluster's data node quantity, between 2 and 50.
        /// </summary>
        [Input("dataNodeAmount")]
        public Input<int>? DataNodeAmount { get; set; }

        /// <summary>
        /// The single data node storage space.
        /// - `cloud_ssd`: An SSD disk, supports a maximum of 2048 GiB (2 TB).
        /// - `cloud_efficiency` An ultra disk, supports a maximum of 5120 GiB (5 TB). If the data to be stored is larger than 2048 GiB, an ultra disk can only support the following data sizes (GiB): [`2560`, `3072`, `3584`, `4096`, `4608`, `5120`].
        /// </summary>
        [Input("dataNodeDiskSize")]
        public Input<int>? DataNodeDiskSize { get; set; }

        /// <summary>
        /// The data node disk type. Supported values: cloud_ssd, cloud_efficiency.
        /// </summary>
        [Input("dataNodeDiskType")]
        public Input<string>? DataNodeDiskType { get; set; }

        /// <summary>
        /// The data node specifications of the Elasticsearch instance.
        /// </summary>
        [Input("dataNodeSpec")]
        public Input<string>? DataNodeSpec { get; set; }

        /// <summary>
        /// The description of instance. It a string of 0 to 30 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Instance connection domain (only VPC network access supported).
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. From version 1.69.0, the Elasticsearch cluster allows you to update your instance_charge_ype from `PostPaid` to `PrePaid`, the following attributes are required: `period`. But, updating from `PostPaid` to `PrePaid` is not supported.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// Kibana console domain (Internet access supported).
        /// </summary>
        [Input("kibanaDomain")]
        public Input<string>? KibanaDomain { get; set; }

        /// <summary>
        /// Kibana console port.
        /// </summary>
        [Input("kibanaPort")]
        public Input<int>? KibanaPort { get; set; }

        [Input("kibanaWhitelists")]
        private InputList<string>? _kibanaWhitelists;

        /// <summary>
        /// Set the Kibana's IP whitelist in internet network.
        /// </summary>
        public InputList<string> KibanaWhitelists
        {
            get => _kibanaWhitelists ?? (_kibanaWhitelists = new InputList<string>());
            set => _kibanaWhitelists = value;
        }

        /// <summary>
        /// An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored, but you have to specify one of `password` and `kms_encrypted_password` fields.
        /// </summary>
        [Input("kmsEncryptedPassword")]
        public Input<string>? KmsEncryptedPassword { get; set; }

        [Input("kmsEncryptionContext")]
        private InputMap<object>? _kmsEncryptionContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        public InputMap<object> KmsEncryptionContext
        {
            get => _kmsEncryptionContext ?? (_kmsEncryptionContext = new InputMap<object>());
            set => _kmsEncryptionContext = value;
        }

        /// <summary>
        /// The dedicated master node spec. If specified, dedicated master node will be created.
        /// </summary>
        [Input("masterNodeSpec")]
        public Input<string>? MasterNodeSpec { get; set; }

        /// <summary>
        /// The password of the instance. The password can be 8 to 32 characters in length and must contain three of the following conditions: uppercase letters, lowercase letters, numbers, and special characters (!@#$%^&amp;*()_+-=).
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The duration that you will buy Elasticsearch instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Instance connection port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("privateWhitelists")]
        private InputList<string>? _privateWhitelists;

        /// <summary>
        /// Set the instance's IP whitelist in VPC network.
        /// </summary>
        public InputList<string> PrivateWhitelists
        {
            get => _privateWhitelists ?? (_privateWhitelists = new InputList<string>());
            set => _privateWhitelists = value;
        }

        [Input("publicWhitelists")]
        private InputList<string>? _publicWhitelists;
        public InputList<string> PublicWhitelists
        {
            get => _publicWhitelists ?? (_publicWhitelists = new InputList<string>());
            set => _publicWhitelists = value;
        }

        /// <summary>
        /// The Elasticsearch instance status. Includes `active`, `activating`, `inactive`. Some operations are denied when status is not `active`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Elasticsearch version. Supported values: `5.5.3_with_X-Pack`, `6.3_with_X-Pack` and `6.7_with_X-Pack`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The ID of VSwitch.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Multi-AZ supported for Elasticsearch, between 1 and 3. The `data_node_amount` value must be an integral multiple of the `zone_count` value.
        /// </summary>
        [Input("zoneCount")]
        public Input<int>? ZoneCount { get; set; }

        public InstanceState()
        {
        }
    }
}
