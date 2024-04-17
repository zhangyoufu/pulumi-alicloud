// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.gpdb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.gpdb.ElasticInstanceArgs;
import com.pulumi.alicloud.gpdb.inputs.ElasticInstanceState;
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
 * Provides a AnalyticDB for PostgreSQL instance resource which storage type is flexible. Compared to the reserved storage ADB PG instance, you can scale up each disk and smoothly scale out nodes online.\
 * For more detail product introduction, see [here](https://www.alibabacloud.com/help/doc-detail/141368.htm).
 * 
 * &gt; **DEPRECATED:**  This resource  has been deprecated from version `1.147.0`. Please use new resource alicloud_gpdb_instance.
 * 
 * &gt; **NOTE:**  Available in 1.127.0+
 * 
 * ## Example Usage
 * 
 * ### Create a AnalyticDB for PostgreSQL instance
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.gpdb.ElasticInstance;
 * import com.pulumi.alicloud.gpdb.ElasticInstanceArgs;
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
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;Gpdb&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .name(&#34;vpc-123456&#34;)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .zoneId(default_.zones()[0].id())
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .vswitchName(&#34;vpc-123456&#34;)
 *             .build());
 * 
 *         var adbPgInstance = new ElasticInstance(&#34;adbPgInstance&#34;, ElasticInstanceArgs.builder()        
 *             .engine(&#34;gpdb&#34;)
 *             .engineVersion(&#34;6.0&#34;)
 *             .segStorageType(&#34;cloud_essd&#34;)
 *             .segNodeNum(4)
 *             .storageSize(50)
 *             .instanceSpec(&#34;2C16G&#34;)
 *             .dbInstanceDescription(&#34;Created by terraform&#34;)
 *             .instanceNetworkType(&#34;VPC&#34;)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * AnalyticDB for PostgreSQL can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:gpdb/elasticInstance:ElasticInstance adb_pg_instance gp-bpxxxxxxxxxxxxxx
 * ```
 * 
 */
@ResourceType(type="alicloud:gpdb/elasticInstance:ElasticInstance")
public class ElasticInstance extends com.pulumi.resources.CustomResource {
    /**
     * ADB PG instance connection string.
     * 
     */
    @Export(name="connectionString", refs={String.class}, tree="[0]")
    private Output<String> connectionString;

    /**
     * @return ADB PG instance connection string.
     * 
     */
    public Output<String> connectionString() {
        return this.connectionString;
    }
    /**
     * The edition of the instance. Valid values: `Basic`, `HighAvailability`. Default value: `HighAvailability`.
     * 
     */
    @Export(name="dbInstanceCategory", refs={String.class}, tree="[0]")
    private Output<String> dbInstanceCategory;

    /**
     * @return The edition of the instance. Valid values: `Basic`, `HighAvailability`. Default value: `HighAvailability`.
     * 
     */
    public Output<String> dbInstanceCategory() {
        return this.dbInstanceCategory;
    }
    /**
     * The description of ADB PG instance. It is a string of 2 to 256 characters.
     * 
     */
    @Export(name="dbInstanceDescription", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dbInstanceDescription;

    /**
     * @return The description of ADB PG instance. It is a string of 2 to 256 characters.
     * 
     */
    public Output<Optional<String>> dbInstanceDescription() {
        return Codegen.optional(this.dbInstanceDescription);
    }
    /**
     * The ID of the encryption key. **Note:** If the `encryption_type` parameter is set to `CloudDisk`, you must specify this parameter to the encryption key that is in the same region as the disk that is specified by the EncryptionType parameter. Otherwise, leave this parameter empty.
     * 
     */
    @Export(name="encryptionKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> encryptionKey;

    /**
     * @return The ID of the encryption key. **Note:** If the `encryption_type` parameter is set to `CloudDisk`, you must specify this parameter to the encryption key that is in the same region as the disk that is specified by the EncryptionType parameter. Otherwise, leave this parameter empty.
     * 
     */
    public Output<Optional<String>> encryptionKey() {
        return Codegen.optional(this.encryptionKey);
    }
    /**
     * The type of the encryption. Valid values: `CloudDisk`. **Note:** Disk encryption cannot be disabled after it is enabled.
     * 
     */
    @Export(name="encryptionType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> encryptionType;

    /**
     * @return The type of the encryption. Valid values: `CloudDisk`. **Note:** Disk encryption cannot be disabled after it is enabled.
     * 
     */
    public Output<Optional<String>> encryptionType() {
        return Codegen.optional(this.encryptionType);
    }
    /**
     * Database engine: `gpdb`.
     * 
     */
    @Export(name="engine", refs={String.class}, tree="[0]")
    private Output<String> engine;

    /**
     * @return Database engine: `gpdb`.
     * 
     */
    public Output<String> engine() {
        return this.engine;
    }
    /**
     * Database version. Valid value is `6.0`.
     * 
     */
    @Export(name="engineVersion", refs={String.class}, tree="[0]")
    private Output<String> engineVersion;

    /**
     * @return Database version. Valid value is `6.0`.
     * 
     */
    public Output<String> engineVersion() {
        return this.engineVersion;
    }
    /**
     * The network type of ADB PG instance. Only `VPC` supported now.
     * 
     */
    @Export(name="instanceNetworkType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> instanceNetworkType;

    /**
     * @return The network type of ADB PG instance. Only `VPC` supported now.
     * 
     */
    public Output<Optional<String>> instanceNetworkType() {
        return Codegen.optional(this.instanceNetworkType);
    }
    /**
     * The specification of segment nodes.
     * * When `db_instance_category` is `HighAvailability`, Valid values: `2C16G`, `4C32G`, `16C128G`.
     * * When `db_instance_category` is `Basic`, Valid values: `2C8G`, `4C16G`, `8C32G`, `16C64G`.
     * 
     */
    @Export(name="instanceSpec", refs={String.class}, tree="[0]")
    private Output<String> instanceSpec;

    /**
     * @return The specification of segment nodes.
     * * When `db_instance_category` is `HighAvailability`, Valid values: `2C16G`, `4C32G`, `16C128G`.
     * * When `db_instance_category` is `Basic`, Valid values: `2C8G`, `4C16G`, `8C32G`, `16C64G`.
     * 
     */
    public Output<String> instanceSpec() {
        return this.instanceSpec;
    }
    /**
     * The subscription period. Valid values: [1~12]. It is valid when payment_type is `Subscription`.\
     * **NOTE:** Will not take effect after modifying `payment_duration` for now, if you want to renew a PayAsYouGo instance, need to do in on aliyun console.
     * 
     */
    @Export(name="paymentDuration", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> paymentDuration;

    /**
     * @return The subscription period. Valid values: [1~12]. It is valid when payment_type is `Subscription`.\
     * **NOTE:** Will not take effect after modifying `payment_duration` for now, if you want to renew a PayAsYouGo instance, need to do in on aliyun console.
     * 
     */
    public Output<Optional<Integer>> paymentDuration() {
        return Codegen.optional(this.paymentDuration);
    }
    /**
     * The unit of the subscription period. Valid values: `Month`, `Year`. It is valid when payment_type is `Subscription`.\
     * **NOTE:** Will not take effect after modifying `payment_duration_unit` for now, if you want to renew a PayAsYouGo instance, need to do in on aliyun console.
     * 
     */
    @Export(name="paymentDurationUnit", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> paymentDurationUnit;

    /**
     * @return The unit of the subscription period. Valid values: `Month`, `Year`. It is valid when payment_type is `Subscription`.\
     * **NOTE:** Will not take effect after modifying `payment_duration_unit` for now, if you want to renew a PayAsYouGo instance, need to do in on aliyun console.
     * 
     */
    public Output<Optional<String>> paymentDurationUnit() {
        return Codegen.optional(this.paymentDurationUnit);
    }
    /**
     * Valid values are `PayAsYouGo`, `Subscription`. Default to `PayAsYouGo`.
     * 
     */
    @Export(name="paymentType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> paymentType;

    /**
     * @return Valid values are `PayAsYouGo`, `Subscription`. Default to `PayAsYouGo`.
     * 
     */
    public Output<Optional<String>> paymentType() {
        return Codegen.optional(this.paymentType);
    }
    /**
     * (Available in 1.196.0+) The connection port of the instance.
     * 
     */
    @Export(name="port", refs={String.class}, tree="[0]")
    private Output<String> port;

    /**
     * @return (Available in 1.196.0+) The connection port of the instance.
     * 
     */
    public Output<String> port() {
        return this.port;
    }
    /**
     * List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     * 
     */
    @Export(name="securityIpLists", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> securityIpLists;

    /**
     * @return List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     * 
     */
    public Output<List<String>> securityIpLists() {
        return this.securityIpLists;
    }
    /**
     * The number of segment nodes. Minimum is `4`, max is `256`, step is `4`.
     * 
     */
    @Export(name="segNodeNum", refs={Integer.class}, tree="[0]")
    private Output<Integer> segNodeNum;

    /**
     * @return The number of segment nodes. Minimum is `4`, max is `256`, step is `4`.
     * 
     */
    public Output<Integer> segNodeNum() {
        return this.segNodeNum;
    }
    /**
     * The disk type of segment nodes. Valid values: `cloud_essd`, `cloud_efficiency`.
     * 
     */
    @Export(name="segStorageType", refs={String.class}, tree="[0]")
    private Output<String> segStorageType;

    /**
     * @return The disk type of segment nodes. Valid values: `cloud_essd`, `cloud_efficiency`.
     * 
     */
    public Output<String> segStorageType() {
        return this.segStorageType;
    }
    /**
     * Instance status.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Instance status.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The storage capacity of per segment node. Unit: GB. Minimum is `50`, max is `4000`, step is `50`.
     * 
     */
    @Export(name="storageSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> storageSize;

    /**
     * @return The storage capacity of per segment node. Unit: GB. Minimum is `50`, max is `4000`, step is `50`.
     * 
     */
    public Output<Integer> storageSize() {
        return this.storageSize;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The virtual switch ID to launch ADB PG instances in one VPC.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    /**
     * @return The virtual switch ID to launch ADB PG instances in one VPC.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }
    /**
     * The Zone to launch the ADB PG instance. If specified, must be consistent with the zone where the vswitch is located.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The Zone to launch the ADB PG instance. If specified, must be consistent with the zone where the vswitch is located.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ElasticInstance(String name) {
        this(name, ElasticInstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ElasticInstance(String name, ElasticInstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ElasticInstance(String name, ElasticInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:gpdb/elasticInstance:ElasticInstance", name, args == null ? ElasticInstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ElasticInstance(String name, Output<String> id, @Nullable ElasticInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:gpdb/elasticInstance:ElasticInstance", name, state, makeResourceOptions(options, id));
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
    public static ElasticInstance get(String name, Output<String> id, @Nullable ElasticInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ElasticInstance(name, id, state, options);
    }
}
