// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.rds.BackupPolicyArgs;
import com.pulumi.alicloud.rds.inputs.BackupPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an RDS instance backup policy resource and used to configure instance backup policy, see [What is DB Backup Policy](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/api-rds-2014-08-15-modifybackuppolicy).
 * 
 * &gt; **NOTE:** Each DB instance has a backup policy and it will be set default values when destroying the resource.
 * 
 * &gt; **NOTE:** Available since v1.5.0.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.rds.RdsFunctions;
 * import com.pulumi.alicloud.rds.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.rds.Instance;
 * import com.pulumi.alicloud.rds.InstanceArgs;
 * import com.pulumi.alicloud.rds.BackupPolicy;
 * import com.pulumi.alicloud.rds.BackupPolicyArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-example&#34;);
 *         final var defaultZones = RdsFunctions.getZones(GetZonesArgs.builder()
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;5.6&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .vswitchName(name)
 *             .build());
 * 
 *         var instance = new Instance(&#34;instance&#34;, InstanceArgs.builder()        
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;5.6&#34;)
 *             .instanceType(&#34;rds.mysql.s1.small&#34;)
 *             .instanceStorage(&#34;10&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .instanceName(name)
 *             .build());
 * 
 *         var policy = new BackupPolicy(&#34;policy&#34;, BackupPolicyArgs.builder()        
 *             .instanceId(instance.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * RDS backup policy can be imported using the id or instance id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:rds/backupPolicy:BackupPolicy example &#34;rm-12345678&#34;
 * ```
 * 
 */
@ResourceType(type="alicloud:rds/backupPolicy:BackupPolicy")
public class BackupPolicy extends com.pulumi.resources.CustomResource {
    /**
     * Instance archive backup keep count. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. When `archive_backup_keep_policy` is `ByMonth` Valid values: [1-31]. When `archive_backup_keep_policy` is `ByWeek` Valid values: [1-7].
     * 
     */
    @Export(name="archiveBackupKeepCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> archiveBackupKeepCount;

    /**
     * @return Instance archive backup keep count. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. When `archive_backup_keep_policy` is `ByMonth` Valid values: [1-31]. When `archive_backup_keep_policy` is `ByWeek` Valid values: [1-7].
     * 
     */
    public Output<Integer> archiveBackupKeepCount() {
        return this.archiveBackupKeepCount;
    }
    /**
     * Instance archive backup keep policy. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. Valid values are `ByMonth`, `ByWeek`, `KeepAll`.
     * 
     */
    @Export(name="archiveBackupKeepPolicy", refs={String.class}, tree="[0]")
    private Output<String> archiveBackupKeepPolicy;

    /**
     * @return Instance archive backup keep policy. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. Valid values are `ByMonth`, `ByWeek`, `KeepAll`.
     * 
     */
    public Output<String> archiveBackupKeepPolicy() {
        return this.archiveBackupKeepPolicy;
    }
    /**
     * Instance archive backup retention days. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. Valid values: [30-1095], and `archive_backup_retention_period` must larger than `backup_retention_period` 730.
     * 
     */
    @Export(name="archiveBackupRetentionPeriod", refs={Integer.class}, tree="[0]")
    private Output<Integer> archiveBackupRetentionPeriod;

    /**
     * @return Instance archive backup retention days. Valid when the `enable_backup_log` is `true` and instance is mysql local disk. Valid values: [30-1095], and `archive_backup_retention_period` must larger than `backup_retention_period` 730.
     * 
     */
    public Output<Integer> archiveBackupRetentionPeriod() {
        return this.archiveBackupRetentionPeriod;
    }
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
     * &gt; **NOTE:** Currently, the SQLServer instance does not support to modify `log_backup_retention_period`.
     * 
     */
    @Export(name="backupInterval", refs={String.class}, tree="[0]")
    private Output<String> backupInterval;

    /**
     * @return The frequency at which you want to perform a snapshot backup on the instance. Valid values:
     * - -1: No backup frequencies are specified.
     * - 30: A snapshot backup is performed once every 30 minutes.
     * - 60: A snapshot backup is performed once every 60 minutes.
     * - 120: A snapshot backup is performed once every 120 minutes.
     * - 240: A snapshot backup is performed once every 240 minutes.
     * - 360: A snapshot backup is performed once every 360 minutes.
     * - 480: A snapshot backup is performed once every 480 minutes.
     * - 720: A snapshot backup is performed once every 720 minutes.
     * 
     * &gt; **NOTE:** Currently, the SQLServer instance does not support to modify `log_backup_retention_period`.
     * 
     */
    public Output<String> backupInterval() {
        return this.backupInterval;
    }
    /**
     * It has been deprecated from version 1.69.0, and use field &#39;preferred_backup_period&#39; instead.
     * 
     * @deprecated
     * Attribute &#39;backup_period&#39; has been deprecated from version 1.69.0. Use `preferred_backup_period` instead
     * 
     */
    @Deprecated /* Attribute 'backup_period' has been deprecated from version 1.69.0. Use `preferred_backup_period` instead */
    @Export(name="backupPeriods", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> backupPeriods;

    /**
     * @return It has been deprecated from version 1.69.0, and use field &#39;preferred_backup_period&#39; instead.
     * 
     */
    public Output<List<String>> backupPeriods() {
        return this.backupPeriods;
    }
    /**
     * Instance backup retention days. Valid values: [7-730]. Default to 7. But mysql local disk is unlimited.
     * 
     */
    @Export(name="backupRetentionPeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> backupRetentionPeriod;

    /**
     * @return Instance backup retention days. Valid values: [7-730]. Default to 7. But mysql local disk is unlimited.
     * 
     */
    public Output<Optional<Integer>> backupRetentionPeriod() {
        return Codegen.optional(this.backupRetentionPeriod);
    }
    /**
     * It has been deprecated from version 1.69.0, and use field &#39;preferred_backup_time&#39; instead.
     * 
     * @deprecated
     * Attribute &#39;backup_time&#39; has been deprecated from version 1.69.0. Use `preferred_backup_time` instead
     * 
     */
    @Deprecated /* Attribute 'backup_time' has been deprecated from version 1.69.0. Use `preferred_backup_time` instead */
    @Export(name="backupTime", refs={String.class}, tree="[0]")
    private Output<String> backupTime;

    /**
     * @return It has been deprecated from version 1.69.0, and use field &#39;preferred_backup_time&#39; instead.
     * 
     */
    public Output<String> backupTime() {
        return this.backupTime;
    }
    /**
     * Whether to enable second level backup.Valid values are `Flash`, `Standard`, Note:It only takes effect when the BackupPolicyMode parameter is DataBackupPolicy.
     * &gt; **NOTE:** You can configure a backup policy by using this parameter and the PreferredBackupPeriod parameter. For example, if you set the PreferredBackupPeriod parameter to Saturday,Sunday and the BackupInterval parameter to -1, a snapshot backup is performed on every Saturday and Sunday.If the instance runs PostgreSQL, the BackupInterval parameter is supported only when the instance is equipped with standard SSDs or enhanced SSDs (ESSDs).This parameter takes effect only when you set the BackupPolicyMode parameter to DataBackupPolicy.
     * 
     */
    @Export(name="category", refs={String.class}, tree="[0]")
    private Output<String> category;

    /**
     * @return Whether to enable second level backup.Valid values are `Flash`, `Standard`, Note:It only takes effect when the BackupPolicyMode parameter is DataBackupPolicy.
     * &gt; **NOTE:** You can configure a backup policy by using this parameter and the PreferredBackupPeriod parameter. For example, if you set the PreferredBackupPeriod parameter to Saturday,Sunday and the BackupInterval parameter to -1, a snapshot backup is performed on every Saturday and Sunday.If the instance runs PostgreSQL, the BackupInterval parameter is supported only when the instance is equipped with standard SSDs or enhanced SSDs (ESSDs).This parameter takes effect only when you set the BackupPolicyMode parameter to DataBackupPolicy.
     * 
     */
    public Output<String> category() {
        return this.category;
    }
    /**
     * The compress type of instance policy. Valid values are `1`, `4`, `8`.
     * 
     */
    @Export(name="compressType", refs={String.class}, tree="[0]")
    private Output<String> compressType;

    /**
     * @return The compress type of instance policy. Valid values are `1`, `4`, `8`.
     * 
     */
    public Output<String> compressType() {
        return this.compressType;
    }
    /**
     * Whether to backup instance log. Valid values are `true`, `false`, Default to `true`. Note: The &#39;Basic Edition&#39; category Rds instance does not support setting log backup. [What is Basic Edition](https://www.alibabacloud.com/help/doc-detail/48980.htm).
     * 
     */
    @Export(name="enableBackupLog", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enableBackupLog;

    /**
     * @return Whether to backup instance log. Valid values are `true`, `false`, Default to `true`. Note: The &#39;Basic Edition&#39; category Rds instance does not support setting log backup. [What is Basic Edition](https://www.alibabacloud.com/help/doc-detail/48980.htm).
     * 
     */
    public Output<Boolean> enableBackupLog() {
        return this.enableBackupLog;
    }
    /**
     * Instance high space usage protection policy. Valid when the `enable_backup_log` is `true`. Valid values are `Enable`, `Disable`.
     * 
     */
    @Export(name="highSpaceUsageProtection", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> highSpaceUsageProtection;

    /**
     * @return Instance high space usage protection policy. Valid when the `enable_backup_log` is `true`. Valid values are `Enable`, `Disable`.
     * 
     */
    public Output<Optional<String>> highSpaceUsageProtection() {
        return Codegen.optional(this.highSpaceUsageProtection);
    }
    /**
     * The Id of instance that can run database.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The Id of instance that can run database.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Instance log backup local retention hours. Valid when the `enable_backup_log` is `true`. Valid values: [0-7*24].
     * 
     */
    @Export(name="localLogRetentionHours", refs={Integer.class}, tree="[0]")
    private Output<Integer> localLogRetentionHours;

    /**
     * @return Instance log backup local retention hours. Valid when the `enable_backup_log` is `true`. Valid values: [0-7*24].
     * 
     */
    public Output<Integer> localLogRetentionHours() {
        return this.localLogRetentionHours;
    }
    /**
     * Instance log backup local retention space. Valid when the `enable_backup_log` is `true`. Valid values: [0-50].
     * 
     */
    @Export(name="localLogRetentionSpace", refs={Integer.class}, tree="[0]")
    private Output<Integer> localLogRetentionSpace;

    /**
     * @return Instance log backup local retention space. Valid when the `enable_backup_log` is `true`. Valid values: [0-50].
     * 
     */
    public Output<Integer> localLogRetentionSpace() {
        return this.localLogRetentionSpace;
    }
    /**
     * It has been deprecated from version 1.68.0, and use field &#39;enable_backup_log&#39; instead.
     * 
     * @deprecated
     * Attribute &#39;log_backup&#39; has been deprecated from version 1.68.0. Use `enable_backup_log` instead
     * 
     */
    @Deprecated /* Attribute 'log_backup' has been deprecated from version 1.68.0. Use `enable_backup_log` instead */
    @Export(name="logBackup", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> logBackup;

    /**
     * @return It has been deprecated from version 1.68.0, and use field &#39;enable_backup_log&#39; instead.
     * 
     */
    public Output<Boolean> logBackup() {
        return this.logBackup;
    }
    /**
     * Instance log backup frequency. Valid when the instance engine is `SQLServer`. Valid values are `LogInterval`.
     * 
     */
    @Export(name="logBackupFrequency", refs={String.class}, tree="[0]")
    private Output<String> logBackupFrequency;

    /**
     * @return Instance log backup frequency. Valid when the instance engine is `SQLServer`. Valid values are `LogInterval`.
     * 
     */
    public Output<String> logBackupFrequency() {
        return this.logBackupFrequency;
    }
    /**
     * Instance log backup retention days. Valid when the `enable_backup_log` is `1`. Valid values: [7-730]. Default to 7. It cannot be larger than `backup_retention_period`.
     * 
     */
    @Export(name="logBackupRetentionPeriod", refs={Integer.class}, tree="[0]")
    private Output<Integer> logBackupRetentionPeriod;

    /**
     * @return Instance log backup retention days. Valid when the `enable_backup_log` is `1`. Valid values: [7-730]. Default to 7. It cannot be larger than `backup_retention_period`.
     * 
     */
    public Output<Integer> logBackupRetentionPeriod() {
        return this.logBackupRetentionPeriod;
    }
    /**
     * It has been deprecated from version 1.69.0, and use field &#39;log_backup_retention_period&#39; instead.
     * 
     * @deprecated
     * Attribute &#39;log_retention_period&#39; has been deprecated from version 1.69.0. Use `log_backup_retention_period` instead
     * 
     */
    @Deprecated /* Attribute 'log_retention_period' has been deprecated from version 1.69.0. Use `log_backup_retention_period` instead */
    @Export(name="logRetentionPeriod", refs={Integer.class}, tree="[0]")
    private Output<Integer> logRetentionPeriod;

    /**
     * @return It has been deprecated from version 1.69.0, and use field &#39;log_backup_retention_period&#39; instead.
     * 
     */
    public Output<Integer> logRetentionPeriod() {
        return this.logRetentionPeriod;
    }
    /**
     * DB Instance backup period. Please set at least two days to ensure backing up at least twice a week. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
     * 
     */
    @Export(name="preferredBackupPeriods", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> preferredBackupPeriods;

    /**
     * @return DB Instance backup period. Please set at least two days to ensure backing up at least twice a week. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
     * 
     */
    public Output<List<String>> preferredBackupPeriods() {
        return this.preferredBackupPeriods;
    }
    /**
     * DB instance backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to &#34;02:00Z-03:00Z&#34;. China time is 8 hours behind it.
     * 
     */
    @Export(name="preferredBackupTime", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> preferredBackupTime;

    /**
     * @return DB instance backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to &#34;02:00Z-03:00Z&#34;. China time is 8 hours behind it.
     * 
     */
    public Output<Optional<String>> preferredBackupTime() {
        return Codegen.optional(this.preferredBackupTime);
    }
    /**
     * The policy based on which ApsaraDB RDS retains archived backup files if the instance is released. Default value: None. Valid values:
     * * **None**: No archived backup files are retained.
     * * **Lastest**: Only the most recent archived backup file is retained.
     * * **All**: All archived backup files are retained.
     * 
     */
    @Export(name="releasedKeepPolicy", refs={String.class}, tree="[0]")
    private Output<String> releasedKeepPolicy;

    /**
     * @return The policy based on which ApsaraDB RDS retains archived backup files if the instance is released. Default value: None. Valid values:
     * * **None**: No archived backup files are retained.
     * * **Lastest**: Only the most recent archived backup file is retained.
     * * **All**: All archived backup files are retained.
     * 
     */
    public Output<String> releasedKeepPolicy() {
        return this.releasedKeepPolicy;
    }
    /**
     * It has been deprecated from version 1.69.0, and use field &#39;backup_retention_period&#39; instead.
     * 
     * @deprecated
     * Attribute &#39;retention_period&#39; has been deprecated from version 1.69.0. Use `backup_retention_period` instead
     * 
     */
    @Deprecated /* Attribute 'retention_period' has been deprecated from version 1.69.0. Use `backup_retention_period` instead */
    @Export(name="retentionPeriod", refs={Integer.class}, tree="[0]")
    private Output<Integer> retentionPeriod;

    /**
     * @return It has been deprecated from version 1.69.0, and use field &#39;backup_retention_period&#39; instead.
     * 
     */
    public Output<Integer> retentionPeriod() {
        return this.retentionPeriod;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BackupPolicy(String name) {
        this(name, BackupPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BackupPolicy(String name, BackupPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BackupPolicy(String name, BackupPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/backupPolicy:BackupPolicy", name, args == null ? BackupPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BackupPolicy(String name, Output<String> id, @Nullable BackupPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/backupPolicy:BackupPolicy", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static BackupPolicy get(String name, Output<String> id, @Nullable BackupPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BackupPolicy(name, id, state, options);
    }
}
