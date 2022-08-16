// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.adb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.adb.DBClusterArgs;
import com.pulumi.alicloud.adb.inputs.DBClusterState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a AnalyticDB for MySQL (ADB) DBCluster resource.
 * 
 * For information about AnalyticDB for MySQL (ADB) DBCluster and how to use it, see [What is DBCluster](https://www.alibabacloud.com/help/en/doc-detail/190519.htm).
 * 
 * &gt; **NOTE:** Available in v1.121.0+.
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.adb.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.adb.DBCluster;
 * import com.pulumi.alicloud.adb.DBClusterArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;adbClusterconfig&#34;);
 *         final var creation = config.get(&#34;creation&#34;).orElse(&#34;ADB&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(creation)
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
 *         var this_ = new DBCluster(&#34;this&#34;, DBClusterArgs.builder()        
 *             .dbClusterCategory(&#34;Cluster&#34;)
 *             .dbNodeClass(&#34;C8&#34;)
 *             .dbNodeCount(&#34;4&#34;)
 *             .dbNodeStorage(&#34;400&#34;)
 *             .mode(&#34;reserver&#34;)
 *             .dbClusterVersion(&#34;3.0&#34;)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .description(&#34;Test new adb again.&#34;)
 *             .maintainTime(&#34;23:00Z-00:00Z&#34;)
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;Created&#34;, &#34;TF-update&#34;),
 *                 Map.entry(&#34;For&#34;, &#34;acceptance-test-update&#34;)
 *             ))
 *             .resourceGroupId(&#34;rg-aek2s7ylxx6****&#34;)
 *             .securityIps(            
 *                 &#34;10.168.1.12&#34;,
 *                 &#34;10.168.1.11&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * AnalyticDB for MySQL (ADB) DBCluster can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:adb/dBCluster:DBCluster example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:adb/dBCluster:DBCluster")
public class DBCluster extends com.pulumi.resources.CustomResource {
    /**
     * Auto-renewal period of an cluster, in the unit of the month. It is valid when `payment_type` is `Subscription`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`. Default to `1`.
     * 
     */
    @Export(name="autoRenewPeriod", type=Integer.class, parameters={})
    private Output<Integer> autoRenewPeriod;

    /**
     * @return Auto-renewal period of an cluster, in the unit of the month. It is valid when `payment_type` is `Subscription`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`. Default to `1`.
     * 
     */
    public Output<Integer> autoRenewPeriod() {
        return this.autoRenewPeriod;
    }
    /**
     * The specifications of computing resources in elastic mode. The increase of resources can speed up queries. AnalyticDB for MySQL automatically scales computing resources. For more information, see [ComputeResource](https://www.alibabacloud.com/help/en/doc-detail/144851.htm)
     * 
     */
    @Export(name="computeResource", type=String.class, parameters={})
    private Output</* @Nullable */ String> computeResource;

    /**
     * @return The specifications of computing resources in elastic mode. The increase of resources can speed up queries. AnalyticDB for MySQL automatically scales computing resources. For more information, see [ComputeResource](https://www.alibabacloud.com/help/en/doc-detail/144851.htm)
     * 
     */
    public Output<Optional<String>> computeResource() {
        return Codegen.optional(this.computeResource);
    }
    /**
     * The endpoint of the cluster.
     * 
     */
    @Export(name="connectionString", type=String.class, parameters={})
    private Output<String> connectionString;

    /**
     * @return The endpoint of the cluster.
     * 
     */
    public Output<String> connectionString() {
        return this.connectionString;
    }
    /**
     * The db cluster category. Valid values: `Basic`, `Cluster`, `MixedStorage`.
     * 
     */
    @Export(name="dbClusterCategory", type=String.class, parameters={})
    private Output<String> dbClusterCategory;

    /**
     * @return The db cluster category. Valid values: `Basic`, `Cluster`, `MixedStorage`.
     * 
     */
    public Output<String> dbClusterCategory() {
        return this.dbClusterCategory;
    }
    /**
     * It duplicates with attribute db_node_class and is deprecated from 1.121.2.
     * 
     * @deprecated
     * It duplicates with attribute db_node_class and is deprecated from 1.121.2.
     * 
     */
    @Deprecated /* It duplicates with attribute db_node_class and is deprecated from 1.121.2. */
    @Export(name="dbClusterClass", type=String.class, parameters={})
    private Output</* @Nullable */ String> dbClusterClass;

    /**
     * @return It duplicates with attribute db_node_class and is deprecated from 1.121.2.
     * 
     */
    public Output<Optional<String>> dbClusterClass() {
        return Codegen.optional(this.dbClusterClass);
    }
    /**
     * The db cluster version. Value options: `3.0`, Default to `3.0`.
     * 
     */
    @Export(name="dbClusterVersion", type=String.class, parameters={})
    private Output</* @Nullable */ String> dbClusterVersion;

    /**
     * @return The db cluster version. Value options: `3.0`, Default to `3.0`.
     * 
     */
    public Output<Optional<String>> dbClusterVersion() {
        return Codegen.optional(this.dbClusterVersion);
    }
    /**
     * The db node class. For more information, see [DBClusterClass](https://help.aliyun.com/document_detail/190519.html)
     * 
     */
    @Export(name="dbNodeClass", type=String.class, parameters={})
    private Output<String> dbNodeClass;

    /**
     * @return The db node class. For more information, see [DBClusterClass](https://help.aliyun.com/document_detail/190519.html)
     * 
     */
    public Output<String> dbNodeClass() {
        return this.dbNodeClass;
    }
    /**
     * The db node count.
     * 
     */
    @Export(name="dbNodeCount", type=Integer.class, parameters={})
    private Output<Integer> dbNodeCount;

    /**
     * @return The db node count.
     * 
     */
    public Output<Integer> dbNodeCount() {
        return this.dbNodeCount;
    }
    /**
     * The db node storage.
     * 
     */
    @Export(name="dbNodeStorage", type=Integer.class, parameters={})
    private Output<Integer> dbNodeStorage;

    /**
     * @return The db node storage.
     * 
     */
    public Output<Integer> dbNodeStorage() {
        return this.dbNodeStorage;
    }
    /**
     * The description of DBCluster.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    /**
     * @return The description of DBCluster.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The elastic io resource.
     * 
     */
    @Export(name="elasticIoResource", type=Integer.class, parameters={})
    private Output<Integer> elasticIoResource;

    /**
     * @return The elastic io resource.
     * 
     */
    public Output<Integer> elasticIoResource() {
        return this.elasticIoResource;
    }
    /**
     * The maintenance window of the cluster. Format: hh:mmZ-hh:mmZ.
     * 
     */
    @Export(name="maintainTime", type=String.class, parameters={})
    private Output<String> maintainTime;

    /**
     * @return The maintenance window of the cluster. Format: hh:mmZ-hh:mmZ.
     * 
     */
    public Output<String> maintainTime() {
        return this.maintainTime;
    }
    /**
     * The mode of the cluster. Valid values: `reserver`, `flexible`.
     * 
     */
    @Export(name="mode", type=String.class, parameters={})
    private Output<String> mode;

    /**
     * @return The mode of the cluster. Valid values: `reserver`, `flexible`.
     * 
     */
    public Output<String> mode() {
        return this.mode;
    }
    /**
     * The modify type.
     * 
     */
    @Export(name="modifyType", type=String.class, parameters={})
    private Output</* @Nullable */ String> modifyType;

    /**
     * @return The modify type.
     * 
     */
    public Output<Optional<String>> modifyType() {
        return Codegen.optional(this.modifyType);
    }
    /**
     * Field `pay_type` has been deprecated. New field `payment_type` instead.
     * 
     * @deprecated
     * Attribute &#39;pay_type&#39; has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute &#39;payment_type&#39; instead.
     * 
     */
    @Deprecated /* Attribute 'pay_type' has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead. */
    @Export(name="payType", type=String.class, parameters={})
    private Output<String> payType;

    /**
     * @return Field `pay_type` has been deprecated. New field `payment_type` instead.
     * 
     */
    public Output<String> payType() {
        return this.payType;
    }
    /**
     * The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `payment_type` supports updating from v1.166.0+.
     * 
     */
    @Export(name="paymentType", type=String.class, parameters={})
    private Output<String> paymentType;

    /**
     * @return The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `payment_type` supports updating from v1.166.0+.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    @Export(name="period", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> period;

    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
     * 
     */
    @Export(name="renewalStatus", type=String.class, parameters={})
    private Output<String> renewalStatus;

    /**
     * @return Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
     * 
     */
    public Output<String> renewalStatus() {
        return this.renewalStatus;
    }
    /**
     * The ID of the resource group.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     * 
     */
    @Export(name="securityIps", type=List.class, parameters={String.class})
    private Output<List<String>> securityIps;

    /**
     * @return List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     * 
     */
    public Output<List<String>> securityIps() {
        return this.securityIps;
    }
    /**
     * The status of the resource.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The vpc ID of the resource.
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output<String> vpcId;

    /**
     * @return The vpc ID of the resource.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The vswitch id.
     * 
     */
    @Export(name="vswitchId", type=String.class, parameters={})
    private Output</* @Nullable */ String> vswitchId;

    /**
     * @return The vswitch id.
     * 
     */
    public Output<Optional<String>> vswitchId() {
        return Codegen.optional(this.vswitchId);
    }
    /**
     * The zone ID of the resource.
     * 
     */
    @Export(name="zoneId", type=String.class, parameters={})
    private Output<String> zoneId;

    /**
     * @return The zone ID of the resource.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DBCluster(String name) {
        this(name, DBClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DBCluster(String name, DBClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DBCluster(String name, DBClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:adb/dBCluster:DBCluster", name, args == null ? DBClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DBCluster(String name, Output<String> id, @Nullable DBClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:adb/dBCluster:DBCluster", name, state, makeResourceOptions(options, id));
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
    public static DBCluster get(String name, Output<String> id, @Nullable DBClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DBCluster(name, id, state, options);
    }
}
