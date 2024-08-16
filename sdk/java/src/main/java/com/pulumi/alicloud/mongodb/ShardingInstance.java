// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mongodb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.mongodb.ShardingInstanceArgs;
import com.pulumi.alicloud.mongodb.inputs.ShardingInstanceState;
import com.pulumi.alicloud.mongodb.outputs.ShardingInstanceConfigServerList;
import com.pulumi.alicloud.mongodb.outputs.ShardingInstanceMongoList;
import com.pulumi.alicloud.mongodb.outputs.ShardingInstanceShardList;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a MongoDB Sharding Instance resource supports replica set instances only. the MongoDB provides stable, reliable, and automatic scalable database services.
 * It offers a full range of database solutions, such as disaster recovery, backup, recovery, monitoring, and alarms.
 * You can see detail product introduction [here](https://www.alibabacloud.com/help/doc-detail/26558.htm)
 * 
 * &gt; **NOTE:** Available since v1.40.0.
 * 
 * &gt; **NOTE:**  The following regions don&#39;t support create Classic network MongoDB Sharding Instance.
 * [`cn-zhangjiakou`,`cn-huhehaote`,`ap-southeast-3`,`ap-southeast-5`,`me-east-1`,`ap-northeast-1`,`eu-west-1`]
 * 
 * &gt; **NOTE:**  Create MongoDB Sharding instance or change instance type and storage would cost 10~20 minutes. Please make full preparation.
 * 
 * ## Example Usage
 * 
 * ### Create a Mongodb Sharding instance
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.mongodb.MongodbFunctions;
 * import com.pulumi.alicloud.mongodb.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.mongodb.ShardingInstance;
 * import com.pulumi.alicloud.mongodb.ShardingInstanceArgs;
 * import com.pulumi.alicloud.mongodb.inputs.ShardingInstanceMongoListArgs;
 * import com.pulumi.alicloud.mongodb.inputs.ShardingInstanceShardListArgs;
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
 *         final var name = config.get("name").orElse("terraform-example");
 *         final var default = MongodbFunctions.getZones();
 * 
 *         final var index = default_.zones().length() - 1;
 * 
 *         final var zoneId = default_.zones()[index].id();
 * 
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()
 *             .vpcName(name)
 *             .cidrBlock("172.17.3.0/24")
 *             .build());
 * 
 *         var defaultSwitch = new Switch("defaultSwitch", SwitchArgs.builder()
 *             .vswitchName(name)
 *             .cidrBlock("172.17.3.0/24")
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(zoneId)
 *             .build());
 * 
 *         var defaultShardingInstance = new ShardingInstance("defaultShardingInstance", ShardingInstanceArgs.builder()
 *             .engineVersion("4.2")
 *             .vswitchId(defaultSwitch.id())
 *             .zoneId(zoneId)
 *             .name(name)
 *             .mongoLists(            
 *                 ShardingInstanceMongoListArgs.builder()
 *                     .nodeClass("dds.mongos.mid")
 *                     .build(),
 *                 ShardingInstanceMongoListArgs.builder()
 *                     .nodeClass("dds.mongos.mid")
 *                     .build())
 *             .shardLists(            
 *                 ShardingInstanceShardListArgs.builder()
 *                     .nodeClass("dds.shard.mid")
 *                     .nodeStorage("10")
 *                     .build(),
 *                 ShardingInstanceShardListArgs.builder()
 *                     .nodeClass("dds.shard.standard")
 *                     .nodeStorage("20")
 *                     .readonlyReplicas("1")
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Module Support
 * 
 * You can use to the existing mongodb-sharding module
 * to create a MongoDB Sharding Instance resource one-click.
 * 
 * ## Import
 * 
 * MongoDB Sharding Instance can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:mongodb/shardingInstance:ShardingInstance example dds-bp1291daeda44195
 * ```
 * 
 */
@ResourceType(type="alicloud:mongodb/shardingInstance:ShardingInstance")
public class ShardingInstance extends com.pulumi.resources.CustomResource {
    /**
     * Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
     * 
     */
    @Export(name="accountPassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accountPassword;

    /**
     * @return Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
     * 
     */
    public Output<Optional<String>> accountPassword() {
        return Codegen.optional(this.accountPassword);
    }
    /**
     * Auto renew for prepaid. Default value: `false`. Valid values: `true`, `false`.
     * 
     */
    @Export(name="autoRenew", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoRenew;

    /**
     * @return Auto renew for prepaid. Default value: `false`. Valid values: `true`, `false`.
     * 
     */
    public Output<Optional<Boolean>> autoRenew() {
        return Codegen.optional(this.autoRenew);
    }
    /**
     * MongoDB Instance backup period. It is required when `backup_time` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
     * 
     */
    @Export(name="backupPeriods", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> backupPeriods;

    /**
     * @return MongoDB Instance backup period. It is required when `backup_time` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
     * 
     */
    public Output<List<String>> backupPeriods() {
        return this.backupPeriods;
    }
    /**
     * Sharding Instance backup time. It is required when `backup_period` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like &#34;23:00Z-24:00Z&#34;.
     * 
     */
    @Export(name="backupTime", refs={String.class}, tree="[0]")
    private Output<String> backupTime;

    /**
     * @return Sharding Instance backup time. It is required when `backup_period` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like &#34;23:00Z-24:00Z&#34;.
     * 
     */
    public Output<String> backupTime() {
        return this.backupTime;
    }
    /**
     * The ConfigServer nodes of the instance. See `config_server_list` below.
     * 
     */
    @Export(name="configServerLists", refs={List.class,ShardingInstanceConfigServerList.class}, tree="[0,1]")
    private Output<List<ShardingInstanceConfigServerList>> configServerLists;

    /**
     * @return The ConfigServer nodes of the instance. See `config_server_list` below.
     * 
     */
    public Output<List<ShardingInstanceConfigServerList>> configServerLists() {
        return this.configServerLists;
    }
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/en/doc-detail/61884.htm) `EngineVersion`. **NOTE:** From version 1.225.1, `engine_version` can be modified.
     * 
     */
    @Export(name="engineVersion", refs={String.class}, tree="[0]")
    private Output<String> engineVersion;

    /**
     * @return Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/en/doc-detail/61884.htm) `EngineVersion`. **NOTE:** From version 1.225.1, `engine_version` can be modified.
     * 
     */
    public Output<String> engineVersion() {
        return this.engineVersion;
    }
    /**
     * The billing method of the instance. Default value: `PostPaid`. Valid values: `PrePaid`, `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version v1.141.0.
     * 
     */
    @Export(name="instanceChargeType", refs={String.class}, tree="[0]")
    private Output<String> instanceChargeType;

    /**
     * @return The billing method of the instance. Default value: `PostPaid`. Valid values: `PrePaid`, `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version v1.141.0.
     * 
     */
    public Output<String> instanceChargeType() {
        return this.instanceChargeType;
    }
    /**
     * An KMS encrypts password used to a instance. If the `account_password` is filled in, this field will be ignored.
     * 
     */
    @Export(name="kmsEncryptedPassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsEncryptedPassword;

    /**
     * @return An KMS encrypts password used to a instance. If the `account_password` is filled in, this field will be ignored.
     * 
     */
    public Output<Optional<String>> kmsEncryptedPassword() {
        return Codegen.optional(this.kmsEncryptedPassword);
    }
    /**
     * An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    @Export(name="kmsEncryptionContext", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> kmsEncryptionContext;

    /**
     * @return An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    public Output<Optional<Map<String,String>>> kmsEncryptionContext() {
        return Codegen.optional(this.kmsEncryptionContext);
    }
    /**
     * The Mongo nodes of the instance. The mongo-node count can be purchased is in range of [2, 32]. See `mongo_list` below.
     * 
     */
    @Export(name="mongoLists", refs={List.class,ShardingInstanceMongoList.class}, tree="[0,1]")
    private Output<List<ShardingInstanceMongoList>> mongoLists;

    /**
     * @return The Mongo nodes of the instance. The mongo-node count can be purchased is in range of [2, 32]. See `mongo_list` below.
     * 
     */
    public Output<List<ShardingInstanceMongoList>> mongoLists() {
        return this.mongoLists;
    }
    /**
     * The name of DB instance. It must be 2 to 256 characters in length.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of DB instance. It must be 2 to 256 characters in length.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The network type of the instance. Valid values:`Classic` or `VPC`.
     * 
     */
    @Export(name="networkType", refs={String.class}, tree="[0]")
    private Output<String> networkType;

    /**
     * @return The network type of the instance. Valid values:`Classic` or `VPC`.
     * 
     */
    public Output<String> networkType() {
        return this.networkType;
    }
    /**
     * The type of configuration changes performed. Default value: `DOWNGRADE`. Valid values:
     * - `UPGRADE`: The specifications are upgraded.
     * - `DOWNGRADE`: The specifications are downgraded.
     *   **NOTE:** `order_type` is only applicable to instances when `instance_charge_type` is `PrePaid`.
     * 
     */
    @Export(name="orderType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> orderType;

    /**
     * @return The type of configuration changes performed. Default value: `DOWNGRADE`. Valid values:
     * - `UPGRADE`: The specifications are upgraded.
     * - `DOWNGRADE`: The specifications are downgraded.
     *   **NOTE:** `order_type` is only applicable to instances when `instance_charge_type` is `PrePaid`.
     * 
     */
    public Output<Optional<String>> orderType() {
        return Codegen.optional(this.orderType);
    }
    /**
     * The duration that you will buy DB instance (in month). It is valid when `instance_charge_type` is `PrePaid`. Default value: `1`. Valid values: [1~9], 12, 24, 36.
     * 
     */
    @Export(name="period", refs={Integer.class}, tree="[0]")
    private Output<Integer> period;

    /**
     * @return The duration that you will buy DB instance (in month). It is valid when `instance_charge_type` is `PrePaid`. Default value: `1`. Valid values: [1~9], 12, 24, 36.
     * 
     */
    public Output<Integer> period() {
        return this.period;
    }
    /**
     * The type of the access protocol. Valid values: `mongodb` or `dynamodb`.
     * 
     */
    @Export(name="protocolType", refs={String.class}, tree="[0]")
    private Output<String> protocolType;

    /**
     * @return The type of the access protocol. Valid values: `mongodb` or `dynamodb`.
     * 
     */
    public Output<String> protocolType() {
        return this.protocolType;
    }
    /**
     * The ID of the Resource Group.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the Resource Group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * (Available since v1.42.0) Instance data backup retention days.
     * 
     */
    @Export(name="retentionPeriod", refs={Integer.class}, tree="[0]")
    private Output<Integer> retentionPeriod;

    /**
     * @return (Available since v1.42.0) Instance data backup retention days.
     * 
     */
    public Output<Integer> retentionPeriod() {
        return this.retentionPeriod;
    }
    /**
     * The Security Group ID of ECS.
     * 
     */
    @Export(name="securityGroupId", refs={String.class}, tree="[0]")
    private Output<String> securityGroupId;

    /**
     * @return The Security Group ID of ECS.
     * 
     */
    public Output<String> securityGroupId() {
        return this.securityGroupId;
    }
    /**
     * List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `[&#34;127.0.0.1&#34;]`.
     * 
     */
    @Export(name="securityIpLists", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> securityIpLists;

    /**
     * @return List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `[&#34;127.0.0.1&#34;]`.
     * 
     */
    public Output<List<String>> securityIpLists() {
        return this.securityIpLists;
    }
    /**
     * The Shard nodes of the instance. The shard-node count can be purchased is in range of [2, 32]. See `shard_list` below.
     * 
     */
    @Export(name="shardLists", refs={List.class,ShardingInstanceShardList.class}, tree="[0,1]")
    private Output<List<ShardingInstanceShardList>> shardLists;

    /**
     * @return The Shard nodes of the instance. The shard-node count can be purchased is in range of [2, 32]. See `shard_list` below.
     * 
     */
    public Output<List<ShardingInstanceShardList>> shardLists() {
        return this.shardLists;
    }
    /**
     * The storage engine of the instance. Default value: `WiredTiger`. Valid values: `WiredTiger`, `RocksDB`.
     * 
     */
    @Export(name="storageEngine", refs={String.class}, tree="[0]")
    private Output<String> storageEngine;

    /**
     * @return The storage engine of the instance. Default value: `WiredTiger`. Valid values: `WiredTiger`, `RocksDB`.
     * 
     */
    public Output<String> storageEngine() {
        return this.storageEngine;
    }
    /**
     * The storage type of the instance. Valid values: `cloud_essd1`, `cloud_essd2`, `cloud_essd3`, `local_ssd`.
     * 
     */
    @Export(name="storageType", refs={String.class}, tree="[0]")
    private Output<String> storageType;

    /**
     * @return The storage type of the instance. Valid values: `cloud_essd1`, `cloud_essd2`, `cloud_essd3`, `local_ssd`.
     * 
     */
    public Output<String> storageType() {
        return this.storageType;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The TDE(Transparent Data Encryption) status. It can be updated from version 1.160.0.
     * 
     */
    @Export(name="tdeStatus", refs={String.class}, tree="[0]")
    private Output<String> tdeStatus;

    /**
     * @return The TDE(Transparent Data Encryption) status. It can be updated from version 1.160.0.
     * 
     */
    public Output<String> tdeStatus() {
        return this.tdeStatus;
    }
    /**
     * The ID of the VPC. &gt; **NOTE:** `vpc_id` is valid only when `network_type` is set to `VPC`.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of the VPC. &gt; **NOTE:** `vpc_id` is valid only when `network_type` is set to `VPC`.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    /**
     * @return The virtual switch ID to launch DB instances in one VPC.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }
    /**
     * The Zone to launch the DB instance. MongoDB Sharding Instance does not support multiple-zone.
     * If it is a multi-zone and `vswitch_id` is specified, the vswitch must in one of them.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The Zone to launch the DB instance. MongoDB Sharding Instance does not support multiple-zone.
     * If it is a multi-zone and `vswitch_id` is specified, the vswitch must in one of them.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ShardingInstance(java.lang.String name) {
        this(name, ShardingInstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ShardingInstance(java.lang.String name, ShardingInstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ShardingInstance(java.lang.String name, ShardingInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mongodb/shardingInstance:ShardingInstance", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ShardingInstance(java.lang.String name, Output<java.lang.String> id, @Nullable ShardingInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mongodb/shardingInstance:ShardingInstance", name, state, makeResourceOptions(options, id), false);
    }

    private static ShardingInstanceArgs makeArgs(ShardingInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ShardingInstanceArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accountPassword"
            ))
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
    public static ShardingInstance get(java.lang.String name, Output<java.lang.String> id, @Nullable ShardingInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ShardingInstance(name, id, state, options);
    }
}
