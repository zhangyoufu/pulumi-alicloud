// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Gpdb
{
    /// <summary>
    /// ## Import
    /// 
    /// GPDB Backup Policy can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:gpdb/backupPolicy:BackupPolicy example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:gpdb/backupPolicy:BackupPolicy")]
    public partial class BackupPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Data backup retention days.
        /// </summary>
        [Output("backupRetentionPeriod")]
        public Output<int> BackupRetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// The instance ID.
        /// &gt; **NOTE:**  You can call the DescribeDBInstances operation to view the details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
        /// </summary>
        [Output("dbInstanceId")]
        public Output<string> DbInstanceId { get; private set; } = null!;

        /// <summary>
        /// Whether to enable automatic recovery points. Value Description:
        /// - **true**: enabled.
        /// - **false**: Closed.
        /// </summary>
        [Output("enableRecoveryPoint")]
        public Output<bool> EnableRecoveryPoint { get; private set; } = null!;

        /// <summary>
        /// Data backup cycle. Separate multiple values with commas (,). Value Description:
        /// - **Monday**: Monday.
        /// - **Tuesday**: Tuesday.
        /// - **Wednesday**: Wednesday.
        /// - **Thursday**: Thursday.
        /// - **Friday**: Friday.
        /// - **Saturday**: Saturday.
        /// - **Sunday**: Sunday.
        /// </summary>
        [Output("preferredBackupPeriod")]
        public Output<string> PreferredBackupPeriod { get; private set; } = null!;

        /// <summary>
        /// Data backup time. Format: HH:mmZ-HH:mmZ(UTC time).
        /// </summary>
        [Output("preferredBackupTime")]
        public Output<string> PreferredBackupTime { get; private set; } = null!;

        /// <summary>
        /// Recovery point frequency. Value Description:
        /// - **1**: Hourly.
        /// - **2**: Every two hours.
        /// - **4**: Every four hours.
        /// - **8**: Every eight hours.
        /// </summary>
        [Output("recoveryPointPeriod")]
        public Output<string> RecoveryPointPeriod { get; private set; } = null!;


        /// <summary>
        /// Create a BackupPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupPolicy(string name, BackupPolicyArgs args, CustomResourceOptions? options = null)
            : base("alicloud:gpdb/backupPolicy:BackupPolicy", name, args ?? new BackupPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupPolicy(string name, Input<string> id, BackupPolicyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:gpdb/backupPolicy:BackupPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackupPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackupPolicy Get(string name, Input<string> id, BackupPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new BackupPolicy(name, id, state, options);
        }
    }

    public sealed class BackupPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Data backup retention days.
        /// </summary>
        [Input("backupRetentionPeriod")]
        public Input<int>? BackupRetentionPeriod { get; set; }

        /// <summary>
        /// The instance ID.
        /// &gt; **NOTE:**  You can call the DescribeDBInstances operation to view the details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        /// <summary>
        /// Whether to enable automatic recovery points. Value Description:
        /// - **true**: enabled.
        /// - **false**: Closed.
        /// </summary>
        [Input("enableRecoveryPoint")]
        public Input<bool>? EnableRecoveryPoint { get; set; }

        /// <summary>
        /// Data backup cycle. Separate multiple values with commas (,). Value Description:
        /// - **Monday**: Monday.
        /// - **Tuesday**: Tuesday.
        /// - **Wednesday**: Wednesday.
        /// - **Thursday**: Thursday.
        /// - **Friday**: Friday.
        /// - **Saturday**: Saturday.
        /// - **Sunday**: Sunday.
        /// </summary>
        [Input("preferredBackupPeriod", required: true)]
        public Input<string> PreferredBackupPeriod { get; set; } = null!;

        /// <summary>
        /// Data backup time. Format: HH:mmZ-HH:mmZ(UTC time).
        /// </summary>
        [Input("preferredBackupTime", required: true)]
        public Input<string> PreferredBackupTime { get; set; } = null!;

        /// <summary>
        /// Recovery point frequency. Value Description:
        /// - **1**: Hourly.
        /// - **2**: Every two hours.
        /// - **4**: Every four hours.
        /// - **8**: Every eight hours.
        /// </summary>
        [Input("recoveryPointPeriod")]
        public Input<string>? RecoveryPointPeriod { get; set; }

        public BackupPolicyArgs()
        {
        }
        public static new BackupPolicyArgs Empty => new BackupPolicyArgs();
    }

    public sealed class BackupPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Data backup retention days.
        /// </summary>
        [Input("backupRetentionPeriod")]
        public Input<int>? BackupRetentionPeriod { get; set; }

        /// <summary>
        /// The instance ID.
        /// &gt; **NOTE:**  You can call the DescribeDBInstances operation to view the details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
        /// </summary>
        [Input("dbInstanceId")]
        public Input<string>? DbInstanceId { get; set; }

        /// <summary>
        /// Whether to enable automatic recovery points. Value Description:
        /// - **true**: enabled.
        /// - **false**: Closed.
        /// </summary>
        [Input("enableRecoveryPoint")]
        public Input<bool>? EnableRecoveryPoint { get; set; }

        /// <summary>
        /// Data backup cycle. Separate multiple values with commas (,). Value Description:
        /// - **Monday**: Monday.
        /// - **Tuesday**: Tuesday.
        /// - **Wednesday**: Wednesday.
        /// - **Thursday**: Thursday.
        /// - **Friday**: Friday.
        /// - **Saturday**: Saturday.
        /// - **Sunday**: Sunday.
        /// </summary>
        [Input("preferredBackupPeriod")]
        public Input<string>? PreferredBackupPeriod { get; set; }

        /// <summary>
        /// Data backup time. Format: HH:mmZ-HH:mmZ(UTC time).
        /// </summary>
        [Input("preferredBackupTime")]
        public Input<string>? PreferredBackupTime { get; set; }

        /// <summary>
        /// Recovery point frequency. Value Description:
        /// - **1**: Hourly.
        /// - **2**: Every two hours.
        /// - **4**: Every four hours.
        /// - **8**: Every eight hours.
        /// </summary>
        [Input("recoveryPointPeriod")]
        public Input<string>? RecoveryPointPeriod { get; set; }

        public BackupPolicyState()
        {
        }
        public static new BackupPolicyState Empty => new BackupPolicyState();
    }
}
