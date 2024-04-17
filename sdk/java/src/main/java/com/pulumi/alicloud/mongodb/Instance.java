// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mongodb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.mongodb.InstanceArgs;
import com.pulumi.alicloud.mongodb.inputs.InstanceState;
import com.pulumi.alicloud.mongodb.outputs.InstanceParameter;
import com.pulumi.alicloud.mongodb.outputs.InstanceReplicaSet;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a MongoDB instance resource supports replica set instances only. the MongoDB provides stable, reliable, and automatic scalable database services.
 * It offers a full range of database solutions, such as disaster recovery, backup, recovery, monitoring, and alarms.
 * You can see detail product introduction [here](https://www.alibabacloud.com/help/doc-detail/26558.htm)
 * 
 * &gt; **NOTE:** Available since v1.37.0.
 * 
 * &gt; **NOTE:**  The following regions don&#39;t support create Classic network MongoDB instance.
 * [`cn-zhangjiakou`,`cn-huhehaote`,`ap-southeast-3`,`ap-southeast-5`,`me-east-1`,`ap-northeast-1`,`eu-west-1`]
 * 
 * &gt; **NOTE:**  Create MongoDB instance or change instance type and storage would cost 5~10 minutes. Please make full preparation
 * 
 * ## Example Usage
 * 
 * ### Create a Mongodb instance
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
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
 * import com.pulumi.alicloud.mongodb.Instance;
 * import com.pulumi.alicloud.mongodb.InstanceArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         final var default = MongodbFunctions.getZones();
 * 
 *         final var index = default_.zones().length() - 1;
 * 
 *         final var zoneId = default_.zones()[index].id();
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(zoneId)
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .engineVersion(&#34;4.2&#34;)
 *             .dbInstanceClass(&#34;dds.mongo.mid&#34;)
 *             .dbInstanceStorage(10)
 *             .vswitchId(defaultSwitch.id())
 *             .securityIpLists(            
 *                 &#34;10.168.1.12&#34;,
 *                 &#34;100.69.7.112&#34;)
 *             .name(name)
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;Created&#34;, &#34;TF&#34;),
 *                 Map.entry(&#34;For&#34;, &#34;example&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Module Support
 * 
 * You can use to the existing mongodb module
 * to create a MongoDB instance resource one-click.
 * 
 * ## Import
 * 
 * MongoDB instance can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:mongodb/instance:Instance example dds-bp1291daeda44194
 * ```
 * 
 */
@ResourceType(type="alicloud:mongodb/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
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
     * &gt; **NOTE:** The start time to the end time must be 1 hour. For example, the MaintainStartTime is 01:00Z, then the MaintainEndTime must be 02:00Z.
     * 
     */
    @Export(name="autoRenew", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoRenew;

    /**
     * @return Auto renew for prepaid. Default value: `false`. Valid values: `true`, `false`.
     * &gt; **NOTE:** The start time to the end time must be 1 hour. For example, the MaintainStartTime is 01:00Z, then the MaintainEndTime must be 02:00Z.
     * 
     */
    public Output<Optional<Boolean>> autoRenew() {
        return Codegen.optional(this.autoRenew);
    }
    /**
     * The frequency at which high-frequency backups are created. Valid values: `-1`, `15`, `30`, `60`, `120`, `180`, `240`, `360`, `480`, `720`.
     * 
     */
    @Export(name="backupInterval", refs={String.class}, tree="[0]")
    private Output<String> backupInterval;

    /**
     * @return The frequency at which high-frequency backups are created. Valid values: `-1`, `15`, `30`, `60`, `120`, `180`, `240`, `360`, `480`, `720`.
     * 
     */
    public Output<String> backupInterval() {
        return this.backupInterval;
    }
    /**
     * MongoDB Instance backup period. It is required when `backup_time` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
     * 
     */
    @Export(name="backupPeriods", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> backupPeriods;

    /**
     * @return MongoDB Instance backup period. It is required when `backup_time` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
     * 
     */
    public Output<List<String>> backupPeriods() {
        return this.backupPeriods;
    }
    /**
     * The retention period of full backups.
     * 
     */
    @Export(name="backupRetentionPeriod", refs={Integer.class}, tree="[0]")
    private Output<Integer> backupRetentionPeriod;

    /**
     * @return The retention period of full backups.
     * 
     */
    public Output<Integer> backupRetentionPeriod() {
        return this.backupRetentionPeriod;
    }
    /**
     * MongoDB instance backup time. It is required when `backup_period` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like &#34;23:00Z-24:00Z&#34;.
     * 
     */
    @Export(name="backupTime", refs={String.class}, tree="[0]")
    private Output<String> backupTime;

    /**
     * @return MongoDB instance backup time. It is required when `backup_period` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like &#34;23:00Z-24:00Z&#34;.
     * 
     */
    public Output<String> backupTime() {
        return this.backupTime;
    }
    /**
     * The ID of the encryption key.
     * 
     */
    @Export(name="cloudDiskEncryptionKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cloudDiskEncryptionKey;

    /**
     * @return The ID of the encryption key.
     * 
     */
    public Output<Optional<String>> cloudDiskEncryptionKey() {
        return Codegen.optional(this.cloudDiskEncryptionKey);
    }
    /**
     * Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
     * 
     */
    @Export(name="dbInstanceClass", refs={String.class}, tree="[0]")
    private Output<String> dbInstanceClass;

    /**
     * @return Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
     * 
     */
    public Output<String> dbInstanceClass() {
        return this.dbInstanceClass;
    }
    /**
     * User-defined DB instance storage space.Unit: GB. Value range:
     * - Custom storage space.
     * - 10-GB increments.
     * 
     */
    @Export(name="dbInstanceStorage", refs={Integer.class}, tree="[0]")
    private Output<Integer> dbInstanceStorage;

    /**
     * @return User-defined DB instance storage space.Unit: GB. Value range:
     * - Custom storage space.
     * - 10-GB increments.
     * 
     */
    public Output<Integer> dbInstanceStorage() {
        return this.dbInstanceStorage;
    }
    /**
     * The time when the changed configurations take effect. Valid values: `Immediately`, `MaintainTime`.
     * 
     */
    @Export(name="effectiveTime", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> effectiveTime;

    /**
     * @return The time when the changed configurations take effect. Valid values: `Immediately`, `MaintainTime`.
     * 
     */
    public Output<Optional<String>> effectiveTime() {
        return Codegen.optional(this.effectiveTime);
    }
    /**
     * Whether to enable cloud disk encryption. Default value: `false`. Valid values: `true`, `false`.
     * 
     */
    @Export(name="encrypted", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> encrypted;

    /**
     * @return Whether to enable cloud disk encryption. Default value: `false`. Valid values: `true`, `false`.
     * 
     */
    public Output<Optional<Boolean>> encrypted() {
        return Codegen.optional(this.encrypted);
    }
    /**
     * The ID of the custom key.
     * 
     */
    @Export(name="encryptionKey", refs={String.class}, tree="[0]")
    private Output<String> encryptionKey;

    /**
     * @return The ID of the custom key.
     * 
     */
    public Output<String> encryptionKey() {
        return this.encryptionKey;
    }
    /**
     * The encryption method. **NOTE:** `encryptor_name` is valid only when `tde_status` is set to `enabled`.
     * 
     */
    @Export(name="encryptorName", refs={String.class}, tree="[0]")
    private Output<String> encryptorName;

    /**
     * @return The encryption method. **NOTE:** `encryptor_name` is valid only when `tde_status` is set to `enabled`.
     * 
     */
    public Output<String> encryptorName() {
        return this.encryptorName;
    }
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/61763.htm) `EngineVersion`.
     * 
     */
    @Export(name="engineVersion", refs={String.class}, tree="[0]")
    private Output<String> engineVersion;

    /**
     * @return Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/61763.htm) `EngineVersion`.
     * 
     */
    public Output<String> engineVersion() {
        return this.engineVersion;
    }
    /**
     * Configure the zone where the hidden node is located to deploy multiple zones. **NOTE:** This parameter value cannot be the same as `zone_id` and `secondary_zone_id` parameter values.
     * 
     */
    @Export(name="hiddenZoneId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> hiddenZoneId;

    /**
     * @return Configure the zone where the hidden node is located to deploy multiple zones. **NOTE:** This parameter value cannot be the same as `zone_id` and `secondary_zone_id` parameter values.
     * 
     */
    public Output<Optional<String>> hiddenZoneId() {
        return Codegen.optional(this.hiddenZoneId);
    }
    /**
     * The billing method of the instance. Default value: `PostPaid`. Valid values: `PrePaid`, `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version 1.63.0.
     * 
     */
    @Export(name="instanceChargeType", refs={String.class}, tree="[0]")
    private Output<String> instanceChargeType;

    /**
     * @return The billing method of the instance. Default value: `PostPaid`. Valid values: `PrePaid`, `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version 1.63.0.
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
    @Export(name="kmsEncryptionContext", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> kmsEncryptionContext;

    /**
     * @return An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    public Output<Optional<Map<String,Object>>> kmsEncryptionContext() {
        return Codegen.optional(this.kmsEncryptionContext);
    }
    /**
     * The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     * 
     */
    @Export(name="maintainEndTime", refs={String.class}, tree="[0]")
    private Output<String> maintainEndTime;

    /**
     * @return The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     * 
     */
    public Output<String> maintainEndTime() {
        return this.maintainEndTime;
    }
    /**
     * The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     * 
     */
    @Export(name="maintainStartTime", refs={String.class}, tree="[0]")
    private Output<String> maintainStartTime;

    /**
     * @return The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     * 
     */
    public Output<String> maintainStartTime() {
        return this.maintainStartTime;
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
     * The network type of the instance. Valid values:`Classic`, `VPC`.
     * 
     */
    @Export(name="networkType", refs={String.class}, tree="[0]")
    private Output<String> networkType;

    /**
     * @return The network type of the instance. Valid values:`Classic`, `VPC`.
     * 
     */
    public Output<String> networkType() {
        return this.networkType;
    }
    /**
     * The type of configuration changes performed. Default value: `DOWNGRADE`. Valid values:
     * 
     */
    @Export(name="orderType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> orderType;

    /**
     * @return The type of configuration changes performed. Default value: `DOWNGRADE`. Valid values:
     * 
     */
    public Output<Optional<String>> orderType() {
        return Codegen.optional(this.orderType);
    }
    /**
     * Set of parameters needs to be set after mongodb instance was launched. See `parameters` below.
     * 
     */
    @Export(name="parameters", refs={List.class,InstanceParameter.class}, tree="[0,1]")
    private Output<List<InstanceParameter>> parameters;

    /**
     * @return Set of parameters needs to be set after mongodb instance was launched. See `parameters` below.
     * 
     */
    public Output<List<InstanceParameter>> parameters() {
        return this.parameters;
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
     * The number of read-only nodes in the replica set instance. Default value: 0. Valid values: 0 to 5.
     * 
     */
    @Export(name="readonlyReplicas", refs={Integer.class}, tree="[0]")
    private Output<Integer> readonlyReplicas;

    /**
     * @return The number of read-only nodes in the replica set instance. Default value: 0. Valid values: 0 to 5.
     * 
     */
    public Output<Integer> readonlyReplicas() {
        return this.readonlyReplicas;
    }
    /**
     * The name of the mongo replica set.
     * 
     */
    @Export(name="replicaSetName", refs={String.class}, tree="[0]")
    private Output<String> replicaSetName;

    /**
     * @return The name of the mongo replica set.
     * 
     */
    public Output<String> replicaSetName() {
        return this.replicaSetName;
    }
    /**
     * Replica set instance information.
     * 
     */
    @Export(name="replicaSets", refs={List.class,InstanceReplicaSet.class}, tree="[0,1]")
    private Output<List<InstanceReplicaSet>> replicaSets;

    /**
     * @return Replica set instance information.
     * 
     */
    public Output<List<InstanceReplicaSet>> replicaSets() {
        return this.replicaSets;
    }
    /**
     * Number of replica set nodes. Valid values: `1`, `3`, `5`, `7`.
     * 
     */
    @Export(name="replicationFactor", refs={Integer.class}, tree="[0]")
    private Output<Integer> replicationFactor;

    /**
     * @return Number of replica set nodes. Valid values: `1`, `3`, `5`, `7`.
     * 
     */
    public Output<Integer> replicationFactor() {
        return this.replicationFactor;
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
     * Instance data backup retention days. Available since v1.42.0.
     * 
     */
    @Export(name="retentionPeriod", refs={Integer.class}, tree="[0]")
    private Output<Integer> retentionPeriod;

    /**
     * @return Instance data backup retention days. Available since v1.42.0.
     * 
     */
    public Output<Integer> retentionPeriod() {
        return this.retentionPeriod;
    }
    /**
     * The Alibaba Cloud Resource Name (ARN) of the specified Resource Access Management (RAM) role.
     * 
     */
    @Export(name="roleArn", refs={String.class}, tree="[0]")
    private Output<String> roleArn;

    /**
     * @return The Alibaba Cloud Resource Name (ARN) of the specified Resource Access Management (RAM) role.
     * 
     */
    public Output<String> roleArn() {
        return this.roleArn;
    }
    /**
     * Configure the available area where the slave node (Secondary node) is located to realize multi-available area deployment. **NOTE:** This parameter value cannot be the same as `zone_id` and `hidden_zone_id` parameter values.
     * 
     */
    @Export(name="secondaryZoneId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> secondaryZoneId;

    /**
     * @return Configure the available area where the slave node (Secondary node) is located to realize multi-available area deployment. **NOTE:** This parameter value cannot be the same as `zone_id` and `hidden_zone_id` parameter values.
     * 
     */
    public Output<Optional<String>> secondaryZoneId() {
        return Codegen.optional(this.secondaryZoneId);
    }
    /**
     * The Security Group ID of ECS.
     * 
     */
    @Export(name="securityGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> securityGroupId;

    /**
     * @return The Security Group ID of ECS.
     * 
     */
    public Output<Optional<String>> securityGroupId() {
        return Codegen.optional(this.securityGroupId);
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
     * The snapshot backup type. Default value: `Standard`. Valid values:
     * 
     */
    @Export(name="snapshotBackupType", refs={String.class}, tree="[0]")
    private Output<String> snapshotBackupType;

    /**
     * @return The snapshot backup type. Default value: `Standard`. Valid values:
     * 
     */
    public Output<String> snapshotBackupType() {
        return this.snapshotBackupType;
    }
    /**
     * Actions performed on SSL functions. Valid values:
     * 
     */
    @Export(name="sslAction", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sslAction;

    /**
     * @return Actions performed on SSL functions. Valid values:
     * 
     */
    public Output<Optional<String>> sslAction() {
        return Codegen.optional(this.sslAction);
    }
    /**
     * Status of the SSL feature.
     * 
     */
    @Export(name="sslStatus", refs={String.class}, tree="[0]")
    private Output<String> sslStatus;

    /**
     * @return Status of the SSL feature.
     * 
     */
    public Output<String> sslStatus() {
        return this.sslStatus;
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
     * The TDE(Transparent Data Encryption) status. Valid values: `enabled`.
     * 
     */
    @Export(name="tdeStatus", refs={String.class}, tree="[0]")
    private Output<String> tdeStatus;

    /**
     * @return The TDE(Transparent Data Encryption) status. Valid values: `enabled`.
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
     * The Zone to launch the DB instance. it supports multiple zone.
     * If it is a multi-zone and `vswitch_id` is specified, the vswitch must in one of them.
     * The multiple zone ID can be retrieved by setting `multi` to &#34;true&#34; in the data source `alicloud.getZones`.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The Zone to launch the DB instance. it supports multiple zone.
     * If it is a multi-zone and `vswitch_id` is specified, the vswitch must in one of them.
     * The multiple zone ID can be retrieved by setting `multi` to &#34;true&#34; in the data source `alicloud.getZones`.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mongodb/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mongodb/instance:Instance", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Instance get(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
