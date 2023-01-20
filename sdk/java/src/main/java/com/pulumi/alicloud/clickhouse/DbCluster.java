// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.clickhouse;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.clickhouse.DbClusterArgs;
import com.pulumi.alicloud.clickhouse.inputs.DbClusterState;
import com.pulumi.alicloud.clickhouse.outputs.DbClusterDbClusterAccessWhiteList;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Click House DBCluster resource.
 * 
 * For information about Click House DBCluster and how to use it, see [What is DBCluster](https://www.alibabacloud.com/product/clickhouse).
 * 
 * &gt; **NOTE:** Available in v1.134.0+.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.clickhouse.ClickhouseFunctions;
 * import com.pulumi.alicloud.clickhouse.inputs.GetRegionsArgs;
 * import com.pulumi.alicloud.vpc.VpcFunctions;
 * import com.pulumi.alicloud.vpc.inputs.GetNetworksArgs;
 * import com.pulumi.alicloud.vpc.inputs.GetSwitchesArgs;
 * import com.pulumi.alicloud.clickhouse.DbCluster;
 * import com.pulumi.alicloud.clickhouse.DbClusterArgs;
 * import com.pulumi.alicloud.clickhouse.inputs.DbClusterDbClusterAccessWhiteListArgs;
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
 *         final var defaultRegions = ClickhouseFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         final var defaultNetworks = VpcFunctions.getNetworks(GetNetworksArgs.builder()
 *             .nameRegex(&#34;default-NODELETING&#34;)
 *             .build());
 * 
 *         final var defaultSwitches = VpcFunctions.getSwitches(GetSwitchesArgs.builder()
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .zoneId(defaultRegions.applyValue(getRegionsResult -&gt; getRegionsResult.regions()[0].zoneIds()[0].zoneId()))
 *             .build());
 * 
 *         var defaultDbCluster = new DbCluster(&#34;defaultDbCluster&#34;, DbClusterArgs.builder()        
 *             .dbClusterVersion(&#34;20.3.10.75&#34;)
 *             .category(&#34;Basic&#34;)
 *             .dbClusterClass(&#34;S8&#34;)
 *             .dbClusterNetworkType(&#34;vpc&#34;)
 *             .dbNodeGroupCount(&#34;1&#34;)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .dbNodeStorage(&#34;500&#34;)
 *             .storageType(&#34;cloud_essd&#34;)
 *             .vswitchId(defaultSwitches.applyValue(getSwitchesResult -&gt; getSwitchesResult.ids()[0]))
 *             .dbClusterAccessWhiteLists(DbClusterDbClusterAccessWhiteListArgs.builder()
 *                 .dbClusterIpArrayAttribute(&#34;test&#34;)
 *                 .dbClusterIpArrayName(&#34;test&#34;)
 *                 .securityIpList(&#34;192.168.0.1&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Click House DBCluster can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:clickhouse/dbCluster:DbCluster example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:clickhouse/dbCluster:DbCluster")
public class DbCluster extends com.pulumi.resources.CustomResource {
    /**
     * The Category of DBCluster. Valid values: `Basic`,`HighAvailability`.
     * 
     */
    @Export(name="category", type=String.class, parameters={})
    private Output<String> category;

    /**
     * @return The Category of DBCluster. Valid values: `Basic`,`HighAvailability`.
     * 
     */
    public Output<String> category() {
        return this.category;
    }
    /**
     * (Available in 1.196.0+) - The connection string of the cluster.
     * 
     */
    @Export(name="connectionString", type=String.class, parameters={})
    private Output<String> connectionString;

    /**
     * @return (Available in 1.196.0+) - The connection string of the cluster.
     * 
     */
    public Output<String> connectionString() {
        return this.connectionString;
    }
    /**
     * The db cluster access white list.
     * 
     */
    @Export(name="dbClusterAccessWhiteLists", type=List.class, parameters={DbClusterDbClusterAccessWhiteList.class})
    private Output</* @Nullable */ List<DbClusterDbClusterAccessWhiteList>> dbClusterAccessWhiteLists;

    /**
     * @return The db cluster access white list.
     * 
     */
    public Output<Optional<List<DbClusterDbClusterAccessWhiteList>>> dbClusterAccessWhiteLists() {
        return Codegen.optional(this.dbClusterAccessWhiteLists);
    }
    /**
     * The DBCluster class. According to the category, db_cluster_class has two value ranges:
     * * Under the condition that the category is the `Basic`, Valid values: `S4-NEW`, `S8`, `S16`, `S32`, `S64`, `S104`.
     * * Under the condition that the category is the `HighAvailability`, Valid values: `C4-NEW`, `C8`, `C16`, `C32`, `C64`, `C104`.
     * 
     */
    @Export(name="dbClusterClass", type=String.class, parameters={})
    private Output<String> dbClusterClass;

    /**
     * @return The DBCluster class. According to the category, db_cluster_class has two value ranges:
     * * Under the condition that the category is the `Basic`, Valid values: `S4-NEW`, `S8`, `S16`, `S32`, `S64`, `S104`.
     * * Under the condition that the category is the `HighAvailability`, Valid values: `C4-NEW`, `C8`, `C16`, `C32`, `C64`, `C104`.
     * 
     */
    public Output<String> dbClusterClass() {
        return this.dbClusterClass;
    }
    /**
     * The DBCluster description.
     * 
     */
    @Export(name="dbClusterDescription", type=String.class, parameters={})
    private Output<String> dbClusterDescription;

    /**
     * @return The DBCluster description.
     * 
     */
    public Output<String> dbClusterDescription() {
        return this.dbClusterDescription;
    }
    /**
     * The DBCluster network type. Valid values: `vpc`.
     * 
     */
    @Export(name="dbClusterNetworkType", type=String.class, parameters={})
    private Output<String> dbClusterNetworkType;

    /**
     * @return The DBCluster network type. Valid values: `vpc`.
     * 
     */
    public Output<String> dbClusterNetworkType() {
        return this.dbClusterNetworkType;
    }
    /**
     * The DBCluster version. Valid values: `20.3.10.75`, `20.8.7.15`, `21.8.10.19`, `22.8.5.29`. **NOTE:** `19.15.2.2` is no longer supported. From version 1.191.0, `db_cluster_version` can be set to `22.8.5.29`.
     * 
     */
    @Export(name="dbClusterVersion", type=String.class, parameters={})
    private Output<String> dbClusterVersion;

    /**
     * @return The DBCluster version. Valid values: `20.3.10.75`, `20.8.7.15`, `21.8.10.19`, `22.8.5.29`. **NOTE:** `19.15.2.2` is no longer supported. From version 1.191.0, `db_cluster_version` can be set to `22.8.5.29`.
     * 
     */
    public Output<String> dbClusterVersion() {
        return this.dbClusterVersion;
    }
    /**
     * The db node group count. The number should between 1 and 48.
     * 
     */
    @Export(name="dbNodeGroupCount", type=Integer.class, parameters={})
    private Output<Integer> dbNodeGroupCount;

    /**
     * @return The db node group count. The number should between 1 and 48.
     * 
     */
    public Output<Integer> dbNodeGroupCount() {
        return this.dbNodeGroupCount;
    }
    /**
     * The db node storage.
     * 
     */
    @Export(name="dbNodeStorage", type=String.class, parameters={})
    private Output<String> dbNodeStorage;

    /**
     * @return The db node storage.
     * 
     */
    public Output<String> dbNodeStorage() {
        return this.dbNodeStorage;
    }
    /**
     * Key management service KMS key ID.
     * 
     */
    @Export(name="encryptionKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> encryptionKey;

    /**
     * @return Key management service KMS key ID.
     * 
     */
    public Output<Optional<String>> encryptionKey() {
        return Codegen.optional(this.encryptionKey);
    }
    /**
     * Currently only supports ECS disk encryption, with a value of CloudDisk, not encrypted when empty.
     * 
     */
    @Export(name="encryptionType", type=String.class, parameters={})
    private Output</* @Nullable */ String> encryptionType;

    /**
     * @return Currently only supports ECS disk encryption, with a value of CloudDisk, not encrypted when empty.
     * 
     */
    public Output<Optional<String>> encryptionType() {
        return Codegen.optional(this.encryptionType);
    }
    /**
     * The maintenance window of DBCluster. Valid format: `hh:mmZ-hh:mm Z`.
     * 
     */
    @Export(name="maintainTime", type=String.class, parameters={})
    private Output<String> maintainTime;

    /**
     * @return The maintenance window of DBCluster. Valid format: `hh:mmZ-hh:mm Z`.
     * 
     */
    public Output<String> maintainTime() {
        return this.maintainTime;
    }
    /**
     * The payment type of the resource. Valid values: `PayAsYouGo`,`Subscription`.
     * 
     */
    @Export(name="paymentType", type=String.class, parameters={})
    private Output<String> paymentType;

    /**
     * @return The payment type of the resource. Valid values: `PayAsYouGo`,`Subscription`.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * Pre-paid cluster of the pay-as-you-go cycle. Valid values: `Month`, `Year`.
     * 
     */
    @Export(name="period", type=String.class, parameters={})
    private Output</* @Nullable */ String> period;

    /**
     * @return Pre-paid cluster of the pay-as-you-go cycle. Valid values: `Month`, `Year`.
     * 
     */
    public Output<Optional<String>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * (Available in 1.196.0+) The connection port of the cluster.
     * 
     */
    @Export(name="port", type=String.class, parameters={})
    private Output<String> port;

    /**
     * @return (Available in 1.196.0+) The connection port of the cluster.
     * 
     */
    public Output<String> port() {
        return this.port;
    }
    /**
     * The status of the resource. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the resource. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Storage type of DBCluster. Valid values: `cloud_essd`, `cloud_efficiency`, `cloud_essd_pl2`, `cloud_essd_pl3`.
     * 
     */
    @Export(name="storageType", type=String.class, parameters={})
    private Output<String> storageType;

    /**
     * @return Storage type of DBCluster. Valid values: `cloud_essd`, `cloud_efficiency`, `cloud_essd_pl2`, `cloud_essd_pl3`.
     * 
     */
    public Output<String> storageType() {
        return this.storageType;
    }
    /**
     * The used time of DBCluster.
     * 
     */
    @Export(name="usedTime", type=String.class, parameters={})
    private Output</* @Nullable */ String> usedTime;

    /**
     * @return The used time of DBCluster.
     * 
     */
    public Output<Optional<String>> usedTime() {
        return Codegen.optional(this.usedTime);
    }
    /**
     * The id of the VPC.
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output<String> vpcId;

    /**
     * @return The id of the VPC.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The vswitch id of DBCluster.
     * 
     */
    @Export(name="vswitchId", type=String.class, parameters={})
    private Output</* @Nullable */ String> vswitchId;

    /**
     * @return The vswitch id of DBCluster.
     * 
     */
    public Output<Optional<String>> vswitchId() {
        return Codegen.optional(this.vswitchId);
    }
    /**
     * The zone ID of the instance.
     * 
     */
    @Export(name="zoneId", type=String.class, parameters={})
    private Output<String> zoneId;

    /**
     * @return The zone ID of the instance.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DbCluster(String name) {
        this(name, DbClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DbCluster(String name, DbClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DbCluster(String name, DbClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:clickhouse/dbCluster:DbCluster", name, args == null ? DbClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DbCluster(String name, Output<String> id, @Nullable DbClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:clickhouse/dbCluster:DbCluster", name, state, makeResourceOptions(options, id));
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
    public static DbCluster get(String name, Output<String> id, @Nullable DbClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DbCluster(name, id, state, options);
    }
}
