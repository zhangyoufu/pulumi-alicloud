// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds
{
    /// <summary>
    /// Provides an RDS instance backup policy resource and used to configure instance backup policy.
    /// 
    /// &gt; **NOTE:** Each DB instance has a backup policy and it will be set default values when destroying the resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var creation = config.Get("creation") ?? "Rds";
    ///     var name = config.Get("name") ?? "dbbackuppolicybasic";
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = creation,
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         VswitchName = name,
    ///     });
    /// 
    ///     var instance = new AliCloud.Rds.Instance("instance", new()
    ///     {
    ///         Engine = "MySQL",
    ///         EngineVersion = "5.6",
    ///         InstanceType = "rds.mysql.s1.small",
    ///         InstanceStorage = 10,
    ///         VswitchId = defaultSwitch.Id,
    ///         InstanceName = name,
    ///     });
    /// 
    ///     var policy = new AliCloud.Rds.BackupPolicy("policy", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RDS backup policy can be imported using the id or instance id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:rds/backupPolicy:BackupPolicy example "rm-12345678"
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:rds/backupPolicy:BackupPolicy")]
    public partial class BackupPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Instance archive backup keep count. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. When `archive_backup_keep_policy` is `ByMonth` Valid values: [1-31]. When `archive_backup_keep_policy` is `ByWeek` Valid values: [1-7].
        /// </summary>
        [Output("archiveBackupKeepCount")]
        public Output<int> ArchiveBackupKeepCount { get; private set; } = null!;

        /// <summary>
        /// Instance archive backup keep policy. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. Valid values are `ByMonth`, `ByWeek`, `KeepAll`.
        /// </summary>
        [Output("archiveBackupKeepPolicy")]
        public Output<string> ArchiveBackupKeepPolicy { get; private set; } = null!;

        /// <summary>
        /// Instance archive backup retention days. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. Valid values: [30-1095], and `archive_backup_retention_period` must larger than `backup_retention_period` 730.
        /// </summary>
        [Output("archiveBackupRetentionPeriod")]
        public Output<int> ArchiveBackupRetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// The frequency at which you want to perform a snapshot backup on the instance. Valid values:
        /// - -1: No backup frequencies are specified.
        /// - 30: A snapshot backup is performed once every 30 minutes.
        /// - 60: A snapshot backup is performed once every 60 minutes.
        /// - 120: A snapshot backup is performed once every 120 minutes.
        /// - 240: A snapshot backup is performed once every 240 minutes.
        /// - 360: A snapshot backup is performed once every 360 minutes.
        /// - 480: A snapshot backup is performed once every 480 minutes.
        /// - 720: A snapshot backup is performed once every 720 minutes.
        /// </summary>
        [Output("backupInterval")]
        public Output<string> BackupInterval { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.69.0, and use field 'preferred_backup_period' instead.
        /// </summary>
        [Output("backupPeriods")]
        public Output<ImmutableArray<string>> BackupPeriods { get; private set; } = null!;

        /// <summary>
        /// Instance backup retention days. Valid values: [7-730]. Default to 7. But mysql local disk is unlimited.
        /// </summary>
        [Output("backupRetentionPeriod")]
        public Output<int?> BackupRetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.69.0, and use field 'preferred_backup_time' instead.
        /// </summary>
        [Output("backupTime")]
        public Output<string> BackupTime { get; private set; } = null!;

        /// <summary>
        /// Whether to enable second level backup.Valid values are `Flash`, `Standard`, Note:It only takes effect when the BackupPolicyMode parameter is DataBackupPolicy. 
        /// &gt; **NOTE:** You can configure a backup policy by using this parameter and the PreferredBackupPeriod parameter. For example, if you set the PreferredBackupPeriod parameter to Saturday,Sunday and the BackupInterval parameter to -1, a snapshot backup is performed on every Saturday and Sunday.If the instance runs PostgreSQL, the BackupInterval parameter is supported only when the instance is equipped with standard SSDs or enhanced SSDs (ESSDs).This parameter takes effect only when you set the BackupPolicyMode parameter to DataBackupPolicy.
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// The compress type of instance policy. Valid values are `1`, `4`, `8`.
        /// </summary>
        [Output("compressType")]
        public Output<string> CompressType { get; private set; } = null!;

        /// <summary>
        /// Whether to backup instance log. Valid values are `true`, `false`, Default to `true`. Note: The 'Basic Edition' category Rds instance does not support setting log backup. [What is Basic Edition](https://www.alibabacloud.com/help/doc-detail/48980.htm).
        /// </summary>
        [Output("enableBackupLog")]
        public Output<bool> EnableBackupLog { get; private set; } = null!;

        /// <summary>
        /// Instance high space usage protection policy. Valid when the `enable_backup_log` is `true`. Valid values are `Enable`, `Disable`.
        /// </summary>
        [Output("highSpaceUsageProtection")]
        public Output<string?> HighSpaceUsageProtection { get; private set; } = null!;

        /// <summary>
        /// The Id of instance that can run database.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Instance log backup local retention hours. Valid when the `enable_backup_log` is `true`. Valid values: [0-7*24].
        /// </summary>
        [Output("localLogRetentionHours")]
        public Output<int> LocalLogRetentionHours { get; private set; } = null!;

        /// <summary>
        /// Instance log backup local retention space. Valid when the `enable_backup_log` is `true`. Valid values: [0-50].
        /// </summary>
        [Output("localLogRetentionSpace")]
        public Output<int> LocalLogRetentionSpace { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.68.0, and use field 'enable_backup_log' instead.
        /// </summary>
        [Output("logBackup")]
        public Output<bool> LogBackup { get; private set; } = null!;

        /// <summary>
        /// Instance log backup frequency. Valid when the instance engine is `SQLServer`. Valid values are `LogInterval`.
        /// </summary>
        [Output("logBackupFrequency")]
        public Output<string> LogBackupFrequency { get; private set; } = null!;

        /// <summary>
        /// Instance log backup retention days. Valid when the `enable_backup_log` is `1`. Valid values: [7-730]. Default to 7. It cannot be larger than `backup_retention_period`.
        /// </summary>
        [Output("logBackupRetentionPeriod")]
        public Output<int> LogBackupRetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.69.0, and use field 'log_backup_retention_period' instead.
        /// </summary>
        [Output("logRetentionPeriod")]
        public Output<int> LogRetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// DB Instance backup period. Please set at least two days to ensure backing up at least twice a week. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
        /// </summary>
        [Output("preferredBackupPeriods")]
        public Output<ImmutableArray<string>> PreferredBackupPeriods { get; private set; } = null!;

        /// <summary>
        /// DB instance backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to "02:00Z-03:00Z". China time is 8 hours behind it.
        /// </summary>
        [Output("preferredBackupTime")]
        public Output<string?> PreferredBackupTime { get; private set; } = null!;

        /// <summary>
        /// The policy based on which ApsaraDB RDS retains archived backup files if the instance is released. Default value: None. Valid values:
        /// * **None**: No archived backup files are retained.
        /// * **Lastest**: Only the most recent archived backup file is retained.
        /// * **All**: All archived backup files are retained.
        /// </summary>
        [Output("releasedKeepPolicy")]
        public Output<string> ReleasedKeepPolicy { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.69.0, and use field 'backup_retention_period' instead.
        /// </summary>
        [Output("retentionPeriod")]
        public Output<int> RetentionPeriod { get; private set; } = null!;


        /// <summary>
        /// Create a BackupPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupPolicy(string name, BackupPolicyArgs args, CustomResourceOptions? options = null)
            : base("alicloud:rds/backupPolicy:BackupPolicy", name, args ?? new BackupPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupPolicy(string name, Input<string> id, BackupPolicyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:rds/backupPolicy:BackupPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Instance archive backup keep count. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. When `archive_backup_keep_policy` is `ByMonth` Valid values: [1-31]. When `archive_backup_keep_policy` is `ByWeek` Valid values: [1-7].
        /// </summary>
        [Input("archiveBackupKeepCount")]
        public Input<int>? ArchiveBackupKeepCount { get; set; }

        /// <summary>
        /// Instance archive backup keep policy. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. Valid values are `ByMonth`, `ByWeek`, `KeepAll`.
        /// </summary>
        [Input("archiveBackupKeepPolicy")]
        public Input<string>? ArchiveBackupKeepPolicy { get; set; }

        /// <summary>
        /// Instance archive backup retention days. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. Valid values: [30-1095], and `archive_backup_retention_period` must larger than `backup_retention_period` 730.
        /// </summary>
        [Input("archiveBackupRetentionPeriod")]
        public Input<int>? ArchiveBackupRetentionPeriod { get; set; }

        /// <summary>
        /// The frequency at which you want to perform a snapshot backup on the instance. Valid values:
        /// - -1: No backup frequencies are specified.
        /// - 30: A snapshot backup is performed once every 30 minutes.
        /// - 60: A snapshot backup is performed once every 60 minutes.
        /// - 120: A snapshot backup is performed once every 120 minutes.
        /// - 240: A snapshot backup is performed once every 240 minutes.
        /// - 360: A snapshot backup is performed once every 360 minutes.
        /// - 480: A snapshot backup is performed once every 480 minutes.
        /// - 720: A snapshot backup is performed once every 720 minutes.
        /// </summary>
        [Input("backupInterval")]
        public Input<string>? BackupInterval { get; set; }

        [Input("backupPeriods")]
        private InputList<string>? _backupPeriods;

        /// <summary>
        /// It has been deprecated from version 1.69.0, and use field 'preferred_backup_period' instead.
        /// </summary>
        [Obsolete(@"Attribute 'backup_period' has been deprecated from version 1.69.0. Use `preferred_backup_period` instead")]
        public InputList<string> BackupPeriods
        {
            get => _backupPeriods ?? (_backupPeriods = new InputList<string>());
            set => _backupPeriods = value;
        }

        /// <summary>
        /// Instance backup retention days. Valid values: [7-730]. Default to 7. But mysql local disk is unlimited.
        /// </summary>
        [Input("backupRetentionPeriod")]
        public Input<int>? BackupRetentionPeriod { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.69.0, and use field 'preferred_backup_time' instead.
        /// </summary>
        [Input("backupTime")]
        public Input<string>? BackupTime { get; set; }

        /// <summary>
        /// Whether to enable second level backup.Valid values are `Flash`, `Standard`, Note:It only takes effect when the BackupPolicyMode parameter is DataBackupPolicy. 
        /// &gt; **NOTE:** You can configure a backup policy by using this parameter and the PreferredBackupPeriod parameter. For example, if you set the PreferredBackupPeriod parameter to Saturday,Sunday and the BackupInterval parameter to -1, a snapshot backup is performed on every Saturday and Sunday.If the instance runs PostgreSQL, the BackupInterval parameter is supported only when the instance is equipped with standard SSDs or enhanced SSDs (ESSDs).This parameter takes effect only when you set the BackupPolicyMode parameter to DataBackupPolicy.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// The compress type of instance policy. Valid values are `1`, `4`, `8`.
        /// </summary>
        [Input("compressType")]
        public Input<string>? CompressType { get; set; }

        /// <summary>
        /// Whether to backup instance log. Valid values are `true`, `false`, Default to `true`. Note: The 'Basic Edition' category Rds instance does not support setting log backup. [What is Basic Edition](https://www.alibabacloud.com/help/doc-detail/48980.htm).
        /// </summary>
        [Input("enableBackupLog")]
        public Input<bool>? EnableBackupLog { get; set; }

        /// <summary>
        /// Instance high space usage protection policy. Valid when the `enable_backup_log` is `true`. Valid values are `Enable`, `Disable`.
        /// </summary>
        [Input("highSpaceUsageProtection")]
        public Input<string>? HighSpaceUsageProtection { get; set; }

        /// <summary>
        /// The Id of instance that can run database.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Instance log backup local retention hours. Valid when the `enable_backup_log` is `true`. Valid values: [0-7*24].
        /// </summary>
        [Input("localLogRetentionHours")]
        public Input<int>? LocalLogRetentionHours { get; set; }

        /// <summary>
        /// Instance log backup local retention space. Valid when the `enable_backup_log` is `true`. Valid values: [0-50].
        /// </summary>
        [Input("localLogRetentionSpace")]
        public Input<int>? LocalLogRetentionSpace { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.68.0, and use field 'enable_backup_log' instead.
        /// </summary>
        [Input("logBackup")]
        public Input<bool>? LogBackup { get; set; }

        /// <summary>
        /// Instance log backup frequency. Valid when the instance engine is `SQLServer`. Valid values are `LogInterval`.
        /// </summary>
        [Input("logBackupFrequency")]
        public Input<string>? LogBackupFrequency { get; set; }

        /// <summary>
        /// Instance log backup retention days. Valid when the `enable_backup_log` is `1`. Valid values: [7-730]. Default to 7. It cannot be larger than `backup_retention_period`.
        /// </summary>
        [Input("logBackupRetentionPeriod")]
        public Input<int>? LogBackupRetentionPeriod { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.69.0, and use field 'log_backup_retention_period' instead.
        /// </summary>
        [Input("logRetentionPeriod")]
        public Input<int>? LogRetentionPeriod { get; set; }

        [Input("preferredBackupPeriods")]
        private InputList<string>? _preferredBackupPeriods;

        /// <summary>
        /// DB Instance backup period. Please set at least two days to ensure backing up at least twice a week. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
        /// </summary>
        public InputList<string> PreferredBackupPeriods
        {
            get => _preferredBackupPeriods ?? (_preferredBackupPeriods = new InputList<string>());
            set => _preferredBackupPeriods = value;
        }

        /// <summary>
        /// DB instance backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to "02:00Z-03:00Z". China time is 8 hours behind it.
        /// </summary>
        [Input("preferredBackupTime")]
        public Input<string>? PreferredBackupTime { get; set; }

        /// <summary>
        /// The policy based on which ApsaraDB RDS retains archived backup files if the instance is released. Default value: None. Valid values:
        /// * **None**: No archived backup files are retained.
        /// * **Lastest**: Only the most recent archived backup file is retained.
        /// * **All**: All archived backup files are retained.
        /// </summary>
        [Input("releasedKeepPolicy")]
        public Input<string>? ReleasedKeepPolicy { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.69.0, and use field 'backup_retention_period' instead.
        /// </summary>
        [Input("retentionPeriod")]
        public Input<int>? RetentionPeriod { get; set; }

        public BackupPolicyArgs()
        {
        }
        public static new BackupPolicyArgs Empty => new BackupPolicyArgs();
    }

    public sealed class BackupPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance archive backup keep count. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. When `archive_backup_keep_policy` is `ByMonth` Valid values: [1-31]. When `archive_backup_keep_policy` is `ByWeek` Valid values: [1-7].
        /// </summary>
        [Input("archiveBackupKeepCount")]
        public Input<int>? ArchiveBackupKeepCount { get; set; }

        /// <summary>
        /// Instance archive backup keep policy. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. Valid values are `ByMonth`, `ByWeek`, `KeepAll`.
        /// </summary>
        [Input("archiveBackupKeepPolicy")]
        public Input<string>? ArchiveBackupKeepPolicy { get; set; }

        /// <summary>
        /// Instance archive backup retention days. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. Valid values: [30-1095], and `archive_backup_retention_period` must larger than `backup_retention_period` 730.
        /// </summary>
        [Input("archiveBackupRetentionPeriod")]
        public Input<int>? ArchiveBackupRetentionPeriod { get; set; }

        /// <summary>
        /// The frequency at which you want to perform a snapshot backup on the instance. Valid values:
        /// - -1: No backup frequencies are specified.
        /// - 30: A snapshot backup is performed once every 30 minutes.
        /// - 60: A snapshot backup is performed once every 60 minutes.
        /// - 120: A snapshot backup is performed once every 120 minutes.
        /// - 240: A snapshot backup is performed once every 240 minutes.
        /// - 360: A snapshot backup is performed once every 360 minutes.
        /// - 480: A snapshot backup is performed once every 480 minutes.
        /// - 720: A snapshot backup is performed once every 720 minutes.
        /// </summary>
        [Input("backupInterval")]
        public Input<string>? BackupInterval { get; set; }

        [Input("backupPeriods")]
        private InputList<string>? _backupPeriods;

        /// <summary>
        /// It has been deprecated from version 1.69.0, and use field 'preferred_backup_period' instead.
        /// </summary>
        [Obsolete(@"Attribute 'backup_period' has been deprecated from version 1.69.0. Use `preferred_backup_period` instead")]
        public InputList<string> BackupPeriods
        {
            get => _backupPeriods ?? (_backupPeriods = new InputList<string>());
            set => _backupPeriods = value;
        }

        /// <summary>
        /// Instance backup retention days. Valid values: [7-730]. Default to 7. But mysql local disk is unlimited.
        /// </summary>
        [Input("backupRetentionPeriod")]
        public Input<int>? BackupRetentionPeriod { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.69.0, and use field 'preferred_backup_time' instead.
        /// </summary>
        [Input("backupTime")]
        public Input<string>? BackupTime { get; set; }

        /// <summary>
        /// Whether to enable second level backup.Valid values are `Flash`, `Standard`, Note:It only takes effect when the BackupPolicyMode parameter is DataBackupPolicy. 
        /// &gt; **NOTE:** You can configure a backup policy by using this parameter and the PreferredBackupPeriod parameter. For example, if you set the PreferredBackupPeriod parameter to Saturday,Sunday and the BackupInterval parameter to -1, a snapshot backup is performed on every Saturday and Sunday.If the instance runs PostgreSQL, the BackupInterval parameter is supported only when the instance is equipped with standard SSDs or enhanced SSDs (ESSDs).This parameter takes effect only when you set the BackupPolicyMode parameter to DataBackupPolicy.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// The compress type of instance policy. Valid values are `1`, `4`, `8`.
        /// </summary>
        [Input("compressType")]
        public Input<string>? CompressType { get; set; }

        /// <summary>
        /// Whether to backup instance log. Valid values are `true`, `false`, Default to `true`. Note: The 'Basic Edition' category Rds instance does not support setting log backup. [What is Basic Edition](https://www.alibabacloud.com/help/doc-detail/48980.htm).
        /// </summary>
        [Input("enableBackupLog")]
        public Input<bool>? EnableBackupLog { get; set; }

        /// <summary>
        /// Instance high space usage protection policy. Valid when the `enable_backup_log` is `true`. Valid values are `Enable`, `Disable`.
        /// </summary>
        [Input("highSpaceUsageProtection")]
        public Input<string>? HighSpaceUsageProtection { get; set; }

        /// <summary>
        /// The Id of instance that can run database.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Instance log backup local retention hours. Valid when the `enable_backup_log` is `true`. Valid values: [0-7*24].
        /// </summary>
        [Input("localLogRetentionHours")]
        public Input<int>? LocalLogRetentionHours { get; set; }

        /// <summary>
        /// Instance log backup local retention space. Valid when the `enable_backup_log` is `true`. Valid values: [0-50].
        /// </summary>
        [Input("localLogRetentionSpace")]
        public Input<int>? LocalLogRetentionSpace { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.68.0, and use field 'enable_backup_log' instead.
        /// </summary>
        [Input("logBackup")]
        public Input<bool>? LogBackup { get; set; }

        /// <summary>
        /// Instance log backup frequency. Valid when the instance engine is `SQLServer`. Valid values are `LogInterval`.
        /// </summary>
        [Input("logBackupFrequency")]
        public Input<string>? LogBackupFrequency { get; set; }

        /// <summary>
        /// Instance log backup retention days. Valid when the `enable_backup_log` is `1`. Valid values: [7-730]. Default to 7. It cannot be larger than `backup_retention_period`.
        /// </summary>
        [Input("logBackupRetentionPeriod")]
        public Input<int>? LogBackupRetentionPeriod { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.69.0, and use field 'log_backup_retention_period' instead.
        /// </summary>
        [Input("logRetentionPeriod")]
        public Input<int>? LogRetentionPeriod { get; set; }

        [Input("preferredBackupPeriods")]
        private InputList<string>? _preferredBackupPeriods;

        /// <summary>
        /// DB Instance backup period. Please set at least two days to ensure backing up at least twice a week. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
        /// </summary>
        public InputList<string> PreferredBackupPeriods
        {
            get => _preferredBackupPeriods ?? (_preferredBackupPeriods = new InputList<string>());
            set => _preferredBackupPeriods = value;
        }

        /// <summary>
        /// DB instance backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to "02:00Z-03:00Z". China time is 8 hours behind it.
        /// </summary>
        [Input("preferredBackupTime")]
        public Input<string>? PreferredBackupTime { get; set; }

        /// <summary>
        /// The policy based on which ApsaraDB RDS retains archived backup files if the instance is released. Default value: None. Valid values:
        /// * **None**: No archived backup files are retained.
        /// * **Lastest**: Only the most recent archived backup file is retained.
        /// * **All**: All archived backup files are retained.
        /// </summary>
        [Input("releasedKeepPolicy")]
        public Input<string>? ReleasedKeepPolicy { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.69.0, and use field 'backup_retention_period' instead.
        /// </summary>
        [Input("retentionPeriod")]
        public Input<int>? RetentionPeriod { get; set; }

        public BackupPolicyState()
        {
        }
        public static new BackupPolicyState Empty => new BackupPolicyState();
    }
}
