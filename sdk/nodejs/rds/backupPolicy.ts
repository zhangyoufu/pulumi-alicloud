// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an RDS instance backup policy resource and used to configure instance backup policy, see [What is DB Backup Policy](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/api-rds-2014-08-15-modifybackuppolicy).
 *
 * > **NOTE:** Each DB instance has a backup policy and it will be set default values when destroying the resource.
 *
 * > **NOTE:** Available since v1.5.0.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const default = alicloud.rds.getZones({
 *     engine: "MySQL",
 *     engineVersion: "5.6",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     vpcName: name,
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: _default.then(_default => _default.zones?.[0]?.id),
 *     vswitchName: name,
 * });
 * const instance = new alicloud.rds.Instance("instance", {
 *     engine: "MySQL",
 *     engineVersion: "5.6",
 *     instanceType: "rds.mysql.s1.small",
 *     instanceStorage: 10,
 *     vswitchId: defaultSwitch.id,
 *     instanceName: name,
 * });
 * const policy = new alicloud.rds.BackupPolicy("policy", {instanceId: instance.id});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * RDS backup policy can be imported using the id or instance id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:rds/backupPolicy:BackupPolicy example "rm-12345678"
 * ```
 */
export class BackupPolicy extends pulumi.CustomResource {
    /**
     * Get an existing BackupPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackupPolicyState, opts?: pulumi.CustomResourceOptions): BackupPolicy {
        return new BackupPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:rds/backupPolicy:BackupPolicy';

    /**
     * Returns true if the given object is an instance of BackupPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackupPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackupPolicy.__pulumiType;
    }

    /**
     * Instance archive backup keep count. Valid when the `enableBackupLog` is `true` and instance is mysql local disk. When `archiveBackupKeepPolicy` is `ByMonth` Valid values: [1-31]. When `archiveBackupKeepPolicy` is `ByWeek` Valid values: [1-7].
     */
    public readonly archiveBackupKeepCount!: pulumi.Output<number>;
    /**
     * Instance archive backup keep policy. Valid when the `enableBackupLog` is `true` and instance is mysql local disk. Valid values are `ByMonth`, `ByWeek`, `KeepAll`.
     */
    public readonly archiveBackupKeepPolicy!: pulumi.Output<string>;
    /**
     * Instance archive backup retention days. Valid when the `enableBackupLog` is `true` and instance is mysql local disk. Valid values: [30-1095], and `archiveBackupRetentionPeriod` must larger than `backupRetentionPeriod` 730.
     */
    public readonly archiveBackupRetentionPeriod!: pulumi.Output<number>;
    /**
     * The frequency at which you want to perform a snapshot backup on the instance. Valid values:
     * - -1: No backup frequencies are specified.
     * - 30: A snapshot backup is performed once every 30 minutes.
     * - 60: A snapshot backup is performed once every 60 minutes.
     * - 120: A snapshot backup is performed once every 120 minutes.
     * - 240: A snapshot backup is performed once every 240 minutes.
     * - 360: A snapshot backup is performed once every 360 minutes.
     * - 480: A snapshot backup is performed once every 480 minutes.
     * - 720: A snapshot backup is performed once every 720 minutes.
     *
     * > **NOTE:** Currently, the SQLServer instance does not support to modify `logBackupRetentionPeriod`.
     */
    public readonly backupInterval!: pulumi.Output<string>;
    /**
     * It has been deprecated from version 1.69.0, and use field 'preferred_backup_period' instead.
     *
     * @deprecated Attribute 'backup_period' has been deprecated from version 1.69.0. Use `preferredBackupPeriod` instead
     */
    public readonly backupPeriods!: pulumi.Output<string[]>;
    /**
     * Instance backup retention days. Valid values: [7-730]. Default to 7. But mysql local disk is unlimited.
     */
    public readonly backupRetentionPeriod!: pulumi.Output<number | undefined>;
    /**
     * It has been deprecated from version 1.69.0, and use field 'preferred_backup_time' instead.
     *
     * @deprecated Attribute 'backup_time' has been deprecated from version 1.69.0. Use `preferredBackupTime` instead
     */
    public readonly backupTime!: pulumi.Output<string>;
    /**
     * Whether to enable second level backup.Valid values are `Flash`, `Standard`, Note:It only takes effect when the BackupPolicyMode parameter is DataBackupPolicy. 
     * > **NOTE:** You can configure a backup policy by using this parameter and the PreferredBackupPeriod parameter. For example, if you set the PreferredBackupPeriod parameter to Saturday,Sunday and the BackupInterval parameter to -1, a snapshot backup is performed on every Saturday and Sunday.If the instance runs PostgreSQL, the BackupInterval parameter is supported only when the instance is equipped with standard SSDs or enhanced SSDs (ESSDs).This parameter takes effect only when you set the BackupPolicyMode parameter to DataBackupPolicy.
     */
    public readonly category!: pulumi.Output<string>;
    /**
     * The compress type of instance policy. Valid values are `1`, `4`, `8`.
     */
    public readonly compressType!: pulumi.Output<string>;
    /**
     * Whether to backup instance log. Valid values are `true`, `false`, Default to `true`. Note: The 'Basic Edition' category Rds instance does not support setting log backup. [What is Basic Edition](https://www.alibabacloud.com/help/doc-detail/48980.htm).
     */
    public readonly enableBackupLog!: pulumi.Output<boolean>;
    /**
     * Instance high space usage protection policy. Valid when the `enableBackupLog` is `true`. Valid values are `Enable`, `Disable`.
     */
    public readonly highSpaceUsageProtection!: pulumi.Output<string | undefined>;
    /**
     * The Id of instance that can run database.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Instance log backup local retention hours. Valid when the `enableBackupLog` is `true`. Valid values: [0-7*24].
     */
    public readonly localLogRetentionHours!: pulumi.Output<number>;
    /**
     * Instance log backup local retention space. Valid when the `enableBackupLog` is `true`. Valid values: [0-50].
     */
    public readonly localLogRetentionSpace!: pulumi.Output<number>;
    /**
     * It has been deprecated from version 1.68.0, and use field 'enable_backup_log' instead.
     *
     * @deprecated Attribute 'log_backup' has been deprecated from version 1.68.0. Use `enableBackupLog` instead
     */
    public readonly logBackup!: pulumi.Output<boolean>;
    /**
     * Instance log backup frequency. Valid when the instance engine is `SQLServer`. Valid values are `LogInterval`.
     */
    public readonly logBackupFrequency!: pulumi.Output<string>;
    /**
     * Instance log backup retention days. Valid when the `enableBackupLog` is `1`. Valid values: [7-730]. Default to 7. It cannot be larger than `backupRetentionPeriod`.
     */
    public readonly logBackupRetentionPeriod!: pulumi.Output<number>;
    /**
     * It has been deprecated from version 1.69.0, and use field 'log_backup_retention_period' instead.
     *
     * @deprecated Attribute 'log_retention_period' has been deprecated from version 1.69.0. Use `logBackupRetentionPeriod` instead
     */
    public readonly logRetentionPeriod!: pulumi.Output<number>;
    /**
     * DB Instance backup period. Please set at least two days to ensure backing up at least twice a week. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
     */
    public readonly preferredBackupPeriods!: pulumi.Output<string[]>;
    /**
     * DB instance backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to "02:00Z-03:00Z". China time is 8 hours behind it.
     */
    public readonly preferredBackupTime!: pulumi.Output<string | undefined>;
    /**
     * The policy based on which ApsaraDB RDS retains archived backup files if the instance is released. Default value: None. Valid values:
     * * **None**: No archived backup files are retained.
     * * **Lastest**: Only the most recent archived backup file is retained.
     * * **All**: All archived backup files are retained.
     */
    public readonly releasedKeepPolicy!: pulumi.Output<string>;
    /**
     * It has been deprecated from version 1.69.0, and use field 'backup_retention_period' instead.
     *
     * @deprecated Attribute 'retention_period' has been deprecated from version 1.69.0. Use `backupRetentionPeriod` instead
     */
    public readonly retentionPeriod!: pulumi.Output<number>;

    /**
     * Create a BackupPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackupPolicyArgs | BackupPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackupPolicyState | undefined;
            resourceInputs["archiveBackupKeepCount"] = state ? state.archiveBackupKeepCount : undefined;
            resourceInputs["archiveBackupKeepPolicy"] = state ? state.archiveBackupKeepPolicy : undefined;
            resourceInputs["archiveBackupRetentionPeriod"] = state ? state.archiveBackupRetentionPeriod : undefined;
            resourceInputs["backupInterval"] = state ? state.backupInterval : undefined;
            resourceInputs["backupPeriods"] = state ? state.backupPeriods : undefined;
            resourceInputs["backupRetentionPeriod"] = state ? state.backupRetentionPeriod : undefined;
            resourceInputs["backupTime"] = state ? state.backupTime : undefined;
            resourceInputs["category"] = state ? state.category : undefined;
            resourceInputs["compressType"] = state ? state.compressType : undefined;
            resourceInputs["enableBackupLog"] = state ? state.enableBackupLog : undefined;
            resourceInputs["highSpaceUsageProtection"] = state ? state.highSpaceUsageProtection : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["localLogRetentionHours"] = state ? state.localLogRetentionHours : undefined;
            resourceInputs["localLogRetentionSpace"] = state ? state.localLogRetentionSpace : undefined;
            resourceInputs["logBackup"] = state ? state.logBackup : undefined;
            resourceInputs["logBackupFrequency"] = state ? state.logBackupFrequency : undefined;
            resourceInputs["logBackupRetentionPeriod"] = state ? state.logBackupRetentionPeriod : undefined;
            resourceInputs["logRetentionPeriod"] = state ? state.logRetentionPeriod : undefined;
            resourceInputs["preferredBackupPeriods"] = state ? state.preferredBackupPeriods : undefined;
            resourceInputs["preferredBackupTime"] = state ? state.preferredBackupTime : undefined;
            resourceInputs["releasedKeepPolicy"] = state ? state.releasedKeepPolicy : undefined;
            resourceInputs["retentionPeriod"] = state ? state.retentionPeriod : undefined;
        } else {
            const args = argsOrState as BackupPolicyArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["archiveBackupKeepCount"] = args ? args.archiveBackupKeepCount : undefined;
            resourceInputs["archiveBackupKeepPolicy"] = args ? args.archiveBackupKeepPolicy : undefined;
            resourceInputs["archiveBackupRetentionPeriod"] = args ? args.archiveBackupRetentionPeriod : undefined;
            resourceInputs["backupInterval"] = args ? args.backupInterval : undefined;
            resourceInputs["backupPeriods"] = args ? args.backupPeriods : undefined;
            resourceInputs["backupRetentionPeriod"] = args ? args.backupRetentionPeriod : undefined;
            resourceInputs["backupTime"] = args ? args.backupTime : undefined;
            resourceInputs["category"] = args ? args.category : undefined;
            resourceInputs["compressType"] = args ? args.compressType : undefined;
            resourceInputs["enableBackupLog"] = args ? args.enableBackupLog : undefined;
            resourceInputs["highSpaceUsageProtection"] = args ? args.highSpaceUsageProtection : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["localLogRetentionHours"] = args ? args.localLogRetentionHours : undefined;
            resourceInputs["localLogRetentionSpace"] = args ? args.localLogRetentionSpace : undefined;
            resourceInputs["logBackup"] = args ? args.logBackup : undefined;
            resourceInputs["logBackupFrequency"] = args ? args.logBackupFrequency : undefined;
            resourceInputs["logBackupRetentionPeriod"] = args ? args.logBackupRetentionPeriod : undefined;
            resourceInputs["logRetentionPeriod"] = args ? args.logRetentionPeriod : undefined;
            resourceInputs["preferredBackupPeriods"] = args ? args.preferredBackupPeriods : undefined;
            resourceInputs["preferredBackupTime"] = args ? args.preferredBackupTime : undefined;
            resourceInputs["releasedKeepPolicy"] = args ? args.releasedKeepPolicy : undefined;
            resourceInputs["retentionPeriod"] = args ? args.retentionPeriod : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BackupPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackupPolicy resources.
 */
export interface BackupPolicyState {
    /**
     * Instance archive backup keep count. Valid when the `enableBackupLog` is `true` and instance is mysql local disk. When `archiveBackupKeepPolicy` is `ByMonth` Valid values: [1-31]. When `archiveBackupKeepPolicy` is `ByWeek` Valid values: [1-7].
     */
    archiveBackupKeepCount?: pulumi.Input<number>;
    /**
     * Instance archive backup keep policy. Valid when the `enableBackupLog` is `true` and instance is mysql local disk. Valid values are `ByMonth`, `ByWeek`, `KeepAll`.
     */
    archiveBackupKeepPolicy?: pulumi.Input<string>;
    /**
     * Instance archive backup retention days. Valid when the `enableBackupLog` is `true` and instance is mysql local disk. Valid values: [30-1095], and `archiveBackupRetentionPeriod` must larger than `backupRetentionPeriod` 730.
     */
    archiveBackupRetentionPeriod?: pulumi.Input<number>;
    /**
     * The frequency at which you want to perform a snapshot backup on the instance. Valid values:
     * - -1: No backup frequencies are specified.
     * - 30: A snapshot backup is performed once every 30 minutes.
     * - 60: A snapshot backup is performed once every 60 minutes.
     * - 120: A snapshot backup is performed once every 120 minutes.
     * - 240: A snapshot backup is performed once every 240 minutes.
     * - 360: A snapshot backup is performed once every 360 minutes.
     * - 480: A snapshot backup is performed once every 480 minutes.
     * - 720: A snapshot backup is performed once every 720 minutes.
     *
     * > **NOTE:** Currently, the SQLServer instance does not support to modify `logBackupRetentionPeriod`.
     */
    backupInterval?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.69.0, and use field 'preferred_backup_period' instead.
     *
     * @deprecated Attribute 'backup_period' has been deprecated from version 1.69.0. Use `preferredBackupPeriod` instead
     */
    backupPeriods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Instance backup retention days. Valid values: [7-730]. Default to 7. But mysql local disk is unlimited.
     */
    backupRetentionPeriod?: pulumi.Input<number>;
    /**
     * It has been deprecated from version 1.69.0, and use field 'preferred_backup_time' instead.
     *
     * @deprecated Attribute 'backup_time' has been deprecated from version 1.69.0. Use `preferredBackupTime` instead
     */
    backupTime?: pulumi.Input<string>;
    /**
     * Whether to enable second level backup.Valid values are `Flash`, `Standard`, Note:It only takes effect when the BackupPolicyMode parameter is DataBackupPolicy. 
     * > **NOTE:** You can configure a backup policy by using this parameter and the PreferredBackupPeriod parameter. For example, if you set the PreferredBackupPeriod parameter to Saturday,Sunday and the BackupInterval parameter to -1, a snapshot backup is performed on every Saturday and Sunday.If the instance runs PostgreSQL, the BackupInterval parameter is supported only when the instance is equipped with standard SSDs or enhanced SSDs (ESSDs).This parameter takes effect only when you set the BackupPolicyMode parameter to DataBackupPolicy.
     */
    category?: pulumi.Input<string>;
    /**
     * The compress type of instance policy. Valid values are `1`, `4`, `8`.
     */
    compressType?: pulumi.Input<string>;
    /**
     * Whether to backup instance log. Valid values are `true`, `false`, Default to `true`. Note: The 'Basic Edition' category Rds instance does not support setting log backup. [What is Basic Edition](https://www.alibabacloud.com/help/doc-detail/48980.htm).
     */
    enableBackupLog?: pulumi.Input<boolean>;
    /**
     * Instance high space usage protection policy. Valid when the `enableBackupLog` is `true`. Valid values are `Enable`, `Disable`.
     */
    highSpaceUsageProtection?: pulumi.Input<string>;
    /**
     * The Id of instance that can run database.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Instance log backup local retention hours. Valid when the `enableBackupLog` is `true`. Valid values: [0-7*24].
     */
    localLogRetentionHours?: pulumi.Input<number>;
    /**
     * Instance log backup local retention space. Valid when the `enableBackupLog` is `true`. Valid values: [0-50].
     */
    localLogRetentionSpace?: pulumi.Input<number>;
    /**
     * It has been deprecated from version 1.68.0, and use field 'enable_backup_log' instead.
     *
     * @deprecated Attribute 'log_backup' has been deprecated from version 1.68.0. Use `enableBackupLog` instead
     */
    logBackup?: pulumi.Input<boolean>;
    /**
     * Instance log backup frequency. Valid when the instance engine is `SQLServer`. Valid values are `LogInterval`.
     */
    logBackupFrequency?: pulumi.Input<string>;
    /**
     * Instance log backup retention days. Valid when the `enableBackupLog` is `1`. Valid values: [7-730]. Default to 7. It cannot be larger than `backupRetentionPeriod`.
     */
    logBackupRetentionPeriod?: pulumi.Input<number>;
    /**
     * It has been deprecated from version 1.69.0, and use field 'log_backup_retention_period' instead.
     *
     * @deprecated Attribute 'log_retention_period' has been deprecated from version 1.69.0. Use `logBackupRetentionPeriod` instead
     */
    logRetentionPeriod?: pulumi.Input<number>;
    /**
     * DB Instance backup period. Please set at least two days to ensure backing up at least twice a week. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
     */
    preferredBackupPeriods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * DB instance backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to "02:00Z-03:00Z". China time is 8 hours behind it.
     */
    preferredBackupTime?: pulumi.Input<string>;
    /**
     * The policy based on which ApsaraDB RDS retains archived backup files if the instance is released. Default value: None. Valid values:
     * * **None**: No archived backup files are retained.
     * * **Lastest**: Only the most recent archived backup file is retained.
     * * **All**: All archived backup files are retained.
     */
    releasedKeepPolicy?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.69.0, and use field 'backup_retention_period' instead.
     *
     * @deprecated Attribute 'retention_period' has been deprecated from version 1.69.0. Use `backupRetentionPeriod` instead
     */
    retentionPeriod?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a BackupPolicy resource.
 */
export interface BackupPolicyArgs {
    /**
     * Instance archive backup keep count. Valid when the `enableBackupLog` is `true` and instance is mysql local disk. When `archiveBackupKeepPolicy` is `ByMonth` Valid values: [1-31]. When `archiveBackupKeepPolicy` is `ByWeek` Valid values: [1-7].
     */
    archiveBackupKeepCount?: pulumi.Input<number>;
    /**
     * Instance archive backup keep policy. Valid when the `enableBackupLog` is `true` and instance is mysql local disk. Valid values are `ByMonth`, `ByWeek`, `KeepAll`.
     */
    archiveBackupKeepPolicy?: pulumi.Input<string>;
    /**
     * Instance archive backup retention days. Valid when the `enableBackupLog` is `true` and instance is mysql local disk. Valid values: [30-1095], and `archiveBackupRetentionPeriod` must larger than `backupRetentionPeriod` 730.
     */
    archiveBackupRetentionPeriod?: pulumi.Input<number>;
    /**
     * The frequency at which you want to perform a snapshot backup on the instance. Valid values:
     * - -1: No backup frequencies are specified.
     * - 30: A snapshot backup is performed once every 30 minutes.
     * - 60: A snapshot backup is performed once every 60 minutes.
     * - 120: A snapshot backup is performed once every 120 minutes.
     * - 240: A snapshot backup is performed once every 240 minutes.
     * - 360: A snapshot backup is performed once every 360 minutes.
     * - 480: A snapshot backup is performed once every 480 minutes.
     * - 720: A snapshot backup is performed once every 720 minutes.
     *
     * > **NOTE:** Currently, the SQLServer instance does not support to modify `logBackupRetentionPeriod`.
     */
    backupInterval?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.69.0, and use field 'preferred_backup_period' instead.
     *
     * @deprecated Attribute 'backup_period' has been deprecated from version 1.69.0. Use `preferredBackupPeriod` instead
     */
    backupPeriods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Instance backup retention days. Valid values: [7-730]. Default to 7. But mysql local disk is unlimited.
     */
    backupRetentionPeriod?: pulumi.Input<number>;
    /**
     * It has been deprecated from version 1.69.0, and use field 'preferred_backup_time' instead.
     *
     * @deprecated Attribute 'backup_time' has been deprecated from version 1.69.0. Use `preferredBackupTime` instead
     */
    backupTime?: pulumi.Input<string>;
    /**
     * Whether to enable second level backup.Valid values are `Flash`, `Standard`, Note:It only takes effect when the BackupPolicyMode parameter is DataBackupPolicy. 
     * > **NOTE:** You can configure a backup policy by using this parameter and the PreferredBackupPeriod parameter. For example, if you set the PreferredBackupPeriod parameter to Saturday,Sunday and the BackupInterval parameter to -1, a snapshot backup is performed on every Saturday and Sunday.If the instance runs PostgreSQL, the BackupInterval parameter is supported only when the instance is equipped with standard SSDs or enhanced SSDs (ESSDs).This parameter takes effect only when you set the BackupPolicyMode parameter to DataBackupPolicy.
     */
    category?: pulumi.Input<string>;
    /**
     * The compress type of instance policy. Valid values are `1`, `4`, `8`.
     */
    compressType?: pulumi.Input<string>;
    /**
     * Whether to backup instance log. Valid values are `true`, `false`, Default to `true`. Note: The 'Basic Edition' category Rds instance does not support setting log backup. [What is Basic Edition](https://www.alibabacloud.com/help/doc-detail/48980.htm).
     */
    enableBackupLog?: pulumi.Input<boolean>;
    /**
     * Instance high space usage protection policy. Valid when the `enableBackupLog` is `true`. Valid values are `Enable`, `Disable`.
     */
    highSpaceUsageProtection?: pulumi.Input<string>;
    /**
     * The Id of instance that can run database.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Instance log backup local retention hours. Valid when the `enableBackupLog` is `true`. Valid values: [0-7*24].
     */
    localLogRetentionHours?: pulumi.Input<number>;
    /**
     * Instance log backup local retention space. Valid when the `enableBackupLog` is `true`. Valid values: [0-50].
     */
    localLogRetentionSpace?: pulumi.Input<number>;
    /**
     * It has been deprecated from version 1.68.0, and use field 'enable_backup_log' instead.
     *
     * @deprecated Attribute 'log_backup' has been deprecated from version 1.68.0. Use `enableBackupLog` instead
     */
    logBackup?: pulumi.Input<boolean>;
    /**
     * Instance log backup frequency. Valid when the instance engine is `SQLServer`. Valid values are `LogInterval`.
     */
    logBackupFrequency?: pulumi.Input<string>;
    /**
     * Instance log backup retention days. Valid when the `enableBackupLog` is `1`. Valid values: [7-730]. Default to 7. It cannot be larger than `backupRetentionPeriod`.
     */
    logBackupRetentionPeriod?: pulumi.Input<number>;
    /**
     * It has been deprecated from version 1.69.0, and use field 'log_backup_retention_period' instead.
     *
     * @deprecated Attribute 'log_retention_period' has been deprecated from version 1.69.0. Use `logBackupRetentionPeriod` instead
     */
    logRetentionPeriod?: pulumi.Input<number>;
    /**
     * DB Instance backup period. Please set at least two days to ensure backing up at least twice a week. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
     */
    preferredBackupPeriods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * DB instance backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to "02:00Z-03:00Z". China time is 8 hours behind it.
     */
    preferredBackupTime?: pulumi.Input<string>;
    /**
     * The policy based on which ApsaraDB RDS retains archived backup files if the instance is released. Default value: None. Valid values:
     * * **None**: No archived backup files are retained.
     * * **Lastest**: Only the most recent archived backup file is retained.
     * * **All**: All archived backup files are retained.
     */
    releasedKeepPolicy?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.69.0, and use field 'backup_retention_period' instead.
     *
     * @deprecated Attribute 'retention_period' has been deprecated from version 1.69.0. Use `backupRetentionPeriod` instead
     */
    retentionPeriod?: pulumi.Input<number>;
}
