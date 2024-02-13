// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbr
{
    /// <summary>
    /// Provides a Hybrid Backup Recovery (HBR) Replication Vault resource.
    /// 
    /// For information about Hybrid Backup Recovery (HBR) Replication Vault and how to use it, see [What is Replication Vault](https://www.alibabacloud.com/help/en/doc-detail/345603.html).
    /// 
    /// &gt; **NOTE:** Available in v1.152.0+.
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
    ///     var sourceRegion = config.Get("sourceRegion") ?? "cn-hangzhou";
    ///     var source = new AliCloud.Provider("source", new()
    ///     {
    ///         Region = sourceRegion,
    ///     });
    /// 
    ///     var defaultReplicationVaultRegions = AliCloud.Hbr.GetReplicationVaultRegions.Invoke();
    /// 
    ///     var replication = new AliCloud.Provider("replication", new()
    ///     {
    ///         Region = defaultReplicationVaultRegions.Apply(getReplicationVaultRegionsResult =&gt; getReplicationVaultRegionsResult.Regions[0]?.ReplicationRegionId),
    ///     });
    /// 
    ///     var defaultVault = new AliCloud.Hbr.Vault("defaultVault", new()
    ///     {
    ///         VaultName = "terraform-example",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Source,
    ///     });
    /// 
    ///     var defaultReplicationVault = new AliCloud.Hbr.ReplicationVault("defaultReplicationVault", new()
    ///     {
    ///         ReplicationSourceRegionId = sourceRegion,
    ///         ReplicationSourceVaultId = defaultVault.Id,
    ///         VaultName = "terraform-example",
    ///         VaultStorageClass = "STANDARD",
    ///         Description = "terraform-example",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Replication,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Hybrid Backup Recovery (HBR) Replication Vault can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:hbr/replicationVault:ReplicationVault example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:hbr/replicationVault:ReplicationVault")]
    public partial class ReplicationVault : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the backup vault. The description must be 0 to 255 characters in length.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the region where the source vault resides.
        /// </summary>
        [Output("replicationSourceRegionId")]
        public Output<string> ReplicationSourceRegionId { get; private set; } = null!;

        /// <summary>
        /// The ID of the source vault.
        /// </summary>
        [Output("replicationSourceVaultId")]
        public Output<string> ReplicationSourceVaultId { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The name of the backup vault. The name must be 1 to 64 characters in length.
        /// </summary>
        [Output("vaultName")]
        public Output<string> VaultName { get; private set; } = null!;

        /// <summary>
        /// The storage type of the backup vault. Valid values: `STANDARD`.
        /// </summary>
        [Output("vaultStorageClass")]
        public Output<string> VaultStorageClass { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicationVault resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicationVault(string name, ReplicationVaultArgs args, CustomResourceOptions? options = null)
            : base("alicloud:hbr/replicationVault:ReplicationVault", name, args ?? new ReplicationVaultArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReplicationVault(string name, Input<string> id, ReplicationVaultState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:hbr/replicationVault:ReplicationVault", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReplicationVault resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicationVault Get(string name, Input<string> id, ReplicationVaultState? state = null, CustomResourceOptions? options = null)
        {
            return new ReplicationVault(name, id, state, options);
        }
    }

    public sealed class ReplicationVaultArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the backup vault. The description must be 0 to 255 characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the region where the source vault resides.
        /// </summary>
        [Input("replicationSourceRegionId", required: true)]
        public Input<string> ReplicationSourceRegionId { get; set; } = null!;

        /// <summary>
        /// The ID of the source vault.
        /// </summary>
        [Input("replicationSourceVaultId", required: true)]
        public Input<string> ReplicationSourceVaultId { get; set; } = null!;

        /// <summary>
        /// The name of the backup vault. The name must be 1 to 64 characters in length.
        /// </summary>
        [Input("vaultName", required: true)]
        public Input<string> VaultName { get; set; } = null!;

        /// <summary>
        /// The storage type of the backup vault. Valid values: `STANDARD`.
        /// </summary>
        [Input("vaultStorageClass")]
        public Input<string>? VaultStorageClass { get; set; }

        public ReplicationVaultArgs()
        {
        }
        public static new ReplicationVaultArgs Empty => new ReplicationVaultArgs();
    }

    public sealed class ReplicationVaultState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the backup vault. The description must be 0 to 255 characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the region where the source vault resides.
        /// </summary>
        [Input("replicationSourceRegionId")]
        public Input<string>? ReplicationSourceRegionId { get; set; }

        /// <summary>
        /// The ID of the source vault.
        /// </summary>
        [Input("replicationSourceVaultId")]
        public Input<string>? ReplicationSourceVaultId { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The name of the backup vault. The name must be 1 to 64 characters in length.
        /// </summary>
        [Input("vaultName")]
        public Input<string>? VaultName { get; set; }

        /// <summary>
        /// The storage type of the backup vault. Valid values: `STANDARD`.
        /// </summary>
        [Input("vaultStorageClass")]
        public Input<string>? VaultStorageClass { get; set; }

        public ReplicationVaultState()
        {
        }
        public static new ReplicationVaultState Empty => new ReplicationVaultState();
    }
}
