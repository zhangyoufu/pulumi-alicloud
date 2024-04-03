// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cassandra;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cassandra.BackupPlanArgs;
import com.pulumi.alicloud.cassandra.inputs.BackupPlanState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Cassandra Backup Plan resource.
 * 
 * For information about Cassandra Backup Plan and how to use it, see [What is Backup Plan](https://www.alibabacloud.com/help/doc-detail/157522.htm).
 * 
 * &gt; **NOTE:** Available in v1.128.0+.
 * 
 * &gt; **DEPRECATED:**  This resource has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
 * 
 * ## Import
 * 
 * Cassandra Backup Plan can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cassandra/backupPlan:BackupPlan example &lt;cluster_id&gt;:&lt;data_center_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cassandra/backupPlan:BackupPlan")
public class BackupPlan extends com.pulumi.resources.CustomResource {
    /**
     * Specifies whether to activate the backup plan. Valid values: `True`, `False`. Default value: `True`.
     * 
     */
    @Export(name="active", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> active;

    /**
     * @return Specifies whether to activate the backup plan. Valid values: `True`, `False`. Default value: `True`.
     * 
     */
    public Output<Boolean> active() {
        return this.active;
    }
    /**
     * The backup cycle. Valid values: `Friday`, `Monday`, `Saturday`, `Sunday`, `Thursday`, `Tuesday`, `Wednesday`.
     * 
     */
    @Export(name="backupPeriod", refs={String.class}, tree="[0]")
    private Output<String> backupPeriod;

    /**
     * @return The backup cycle. Valid values: `Friday`, `Monday`, `Saturday`, `Sunday`, `Thursday`, `Tuesday`, `Wednesday`.
     * 
     */
    public Output<String> backupPeriod() {
        return this.backupPeriod;
    }
    /**
     * The start time of the backup task each day. The time is displayed in UTC and denoted by Z.
     * 
     */
    @Export(name="backupTime", refs={String.class}, tree="[0]")
    private Output<String> backupTime;

    /**
     * @return The start time of the backup task each day. The time is displayed in UTC and denoted by Z.
     * 
     */
    public Output<String> backupTime() {
        return this.backupTime;
    }
    /**
     * The ID of the cluster for the backup.
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output<String> clusterId;

    /**
     * @return The ID of the cluster for the backup.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * The ID of the data center for the backup in the cluster.
     * 
     */
    @Export(name="dataCenterId", refs={String.class}, tree="[0]")
    private Output<String> dataCenterId;

    /**
     * @return The ID of the data center for the backup in the cluster.
     * 
     */
    public Output<String> dataCenterId() {
        return this.dataCenterId;
    }
    /**
     * The duration for which you want to retain the backup. Valid values: 1 to 30. Unit: days. Default value: `30`.
     * 
     */
    @Export(name="retentionPeriod", refs={Integer.class}, tree="[0]")
    private Output<Integer> retentionPeriod;

    /**
     * @return The duration for which you want to retain the backup. Valid values: 1 to 30. Unit: days. Default value: `30`.
     * 
     */
    public Output<Integer> retentionPeriod() {
        return this.retentionPeriod;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BackupPlan(String name) {
        this(name, BackupPlanArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BackupPlan(String name, BackupPlanArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BackupPlan(String name, BackupPlanArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cassandra/backupPlan:BackupPlan", name, args == null ? BackupPlanArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BackupPlan(String name, Output<String> id, @Nullable BackupPlanState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cassandra/backupPlan:BackupPlan", name, state, makeResourceOptions(options, id));
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
    public static BackupPlan get(String name, Output<String> id, @Nullable BackupPlanState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BackupPlan(name, id, state, options);
    }
}
