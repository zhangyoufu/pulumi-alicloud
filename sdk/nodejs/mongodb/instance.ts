// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a MongoDB instance resource supports replica set instances only. the MongoDB provides stable, reliable, and automatic scalable database services.
 * It offers a full range of database solutions, such as disaster recovery, backup, recovery, monitoring, and alarms.
 * You can see detail product introduction [here](https://www.alibabacloud.com/help/doc-detail/26558.htm)
 *
 * > **NOTE:** Available since v1.37.0.
 *
 * > **NOTE:**  The following regions don't support create Classic network MongoDB instance.
 * [`cn-zhangjiakou`,`cn-huhehaote`,`ap-southeast-2`,`ap-southeast-3`,`ap-southeast-5`,`ap-south-1`,`me-east-1`,`ap-northeast-1`,`eu-west-1`]
 *
 * > **NOTE:**  Create MongoDB instance or change instance type and storage would cost 5~10 minutes. Please make full preparation
 *
 * ## Example Usage
 * ### Create a Mongodb instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const defaultZones = alicloud.mongodb.getZones({});
 * const index = defaultZones.then(defaultZones => defaultZones.zones).length.then(length => length - 1);
 * const zoneId = defaultZones.then(defaultZones => defaultZones.zones[index].id);
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "172.17.3.0/24",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vswitchName: name,
 *     cidrBlock: "172.17.3.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: zoneId,
 * });
 * const defaultInstance = new alicloud.mongodb.Instance("defaultInstance", {
 *     engineVersion: "4.2",
 *     dbInstanceClass: "dds.mongo.mid",
 *     dbInstanceStorage: 10,
 *     vswitchId: defaultSwitch.id,
 *     securityIpLists: [
 *         "10.168.1.12",
 *         "100.69.7.112",
 *     ],
 *     tags: {
 *         Created: "TF",
 *         For: "example",
 *     },
 * });
 * ```
 * ## Module Support
 *
 * You can use to the existing mongodb module
 * to create a MongoDB instance resource one-click.
 *
 * ## Import
 *
 * MongoDB can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:mongodb/instance:Instance example dds-bp1291daeda44194
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:mongodb/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
     */
    public readonly accountPassword!: pulumi.Output<string | undefined>;
    /**
     * Auto renew for prepaid, true of false. Default is false.
     * > **NOTE:** The start time to the end time must be 1 hour. For example, the MaintainStartTime is 01:00Z, then the MaintainEndTime must be 02:00Z.
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
     */
    public readonly backupPeriods!: pulumi.Output<string[]>;
    /**
     * MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
     */
    public readonly backupTime!: pulumi.Output<string>;
    /**
     * Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
     */
    public readonly dbInstanceClass!: pulumi.Output<string>;
    /**
     * User-defined DB instance storage space.Unit: GB. Value range:
     * - Custom storage space.
     * - 10-GB increments.
     */
    public readonly dbInstanceStorage!: pulumi.Output<number>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/61763.htm) `EngineVersion`.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * Configure the zone where the hidden node is located to deploy multiple zones. **NOTE:** This parameter value cannot be the same as `zoneId` and `secondaryZoneId` parameter values.
     */
    public readonly hiddenZoneId!: pulumi.Output<string | undefined>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version 1.63.0.
     */
    public readonly instanceChargeType!: pulumi.Output<string | undefined>;
    /**
     * An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
     */
    public readonly kmsEncryptedPassword!: pulumi.Output<string | undefined>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    public readonly kmsEncryptionContext!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    public readonly maintainEndTime!: pulumi.Output<string>;
    /**
     * The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    public readonly maintainStartTime!: pulumi.Output<string>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
     */
    public readonly networkType!: pulumi.Output<string>;
    /**
     * The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
     * * UPGRADE: The specifications are upgraded.
     * * DOWNGRADE: The specifications are downgraded.
     * **NOTE:** This parameter is only applicable to instances when `instanceChargeType` is PrePaid.
     */
    public readonly orderType!: pulumi.Output<string | undefined>;
    /**
     * Set of parameters needs to be set after mongodb instance was launched. See `parameters` below.
     */
    public readonly parameters!: pulumi.Output<outputs.mongodb.InstanceParameter[]>;
    /**
     * The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
     */
    public readonly period!: pulumi.Output<number>;
    /**
     * The number of read-only nodes in the replica set instance. Default value: 0. Valid values: 0 to 5.
     */
    public readonly readonlyReplicas!: pulumi.Output<number>;
    /**
     * The name of the mongo replica set
     */
    public /*out*/ readonly replicaSetName!: pulumi.Output<string>;
    /**
     * Replica set instance information. The details see Block replica_sets. **NOTE:** Available since v1.140. See `replicaSets` below.
     */
    public /*out*/ readonly replicaSets!: pulumi.Output<outputs.mongodb.InstanceReplicaSet[]>;
    /**
     * Number of replica set nodes. Valid values: [1, 3, 5, 7]
     */
    public readonly replicationFactor!: pulumi.Output<number>;
    /**
     * The ID of the Resource Group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * Instance log backup retention days. Available in 1.42.0+.
     */
    public /*out*/ readonly retentionPeriod!: pulumi.Output<number>;
    /**
     * Configure the available area where the slave node (Secondary node) is located to realize multi-available area deployment. **NOTE:** This parameter value cannot be the same as `zoneId` and `hiddenZoneId` parameter values.
     */
    public readonly secondaryZoneId!: pulumi.Output<string | undefined>;
    /**
     * The Security Group ID of ECS.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     */
    public readonly securityIpLists!: pulumi.Output<string[]>;
    /**
     * Actions performed on SSL functions, Valid values: `Open`: turn on SSL encryption; `Close`: turn off SSL encryption; `Update`: update SSL certificate.
     */
    public readonly sslAction!: pulumi.Output<string>;
    /**
     * Status of the SSL feature. `Open`: SSL is turned on; `Closed`: SSL is turned off.
     */
    public /*out*/ readonly sslStatus!: pulumi.Output<string>;
    /**
     * Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
     */
    public readonly storageEngine!: pulumi.Output<string>;
    /**
     * The storage type of the instance. Valid values: `cloudEssd1`, `cloudEssd2`, `cloudEssd3`, `localSsd`.
     */
    public readonly storageType!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The TDE(Transparent Data Encryption) status.
     */
    public readonly tdeStatus!: pulumi.Output<string | undefined>;
    /**
     * The ID of the VPC. > **NOTE:** This parameter is valid only when NetworkType is set to VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    public readonly vswitchId!: pulumi.Output<string>;
    /**
     * The Zone to launch the DB instance. it supports multiple zone.
     * If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
     * The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `alicloud.getZones`.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["accountPassword"] = state ? state.accountPassword : undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["backupPeriods"] = state ? state.backupPeriods : undefined;
            resourceInputs["backupTime"] = state ? state.backupTime : undefined;
            resourceInputs["dbInstanceClass"] = state ? state.dbInstanceClass : undefined;
            resourceInputs["dbInstanceStorage"] = state ? state.dbInstanceStorage : undefined;
            resourceInputs["engineVersion"] = state ? state.engineVersion : undefined;
            resourceInputs["hiddenZoneId"] = state ? state.hiddenZoneId : undefined;
            resourceInputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            resourceInputs["kmsEncryptedPassword"] = state ? state.kmsEncryptedPassword : undefined;
            resourceInputs["kmsEncryptionContext"] = state ? state.kmsEncryptionContext : undefined;
            resourceInputs["maintainEndTime"] = state ? state.maintainEndTime : undefined;
            resourceInputs["maintainStartTime"] = state ? state.maintainStartTime : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["orderType"] = state ? state.orderType : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["readonlyReplicas"] = state ? state.readonlyReplicas : undefined;
            resourceInputs["replicaSetName"] = state ? state.replicaSetName : undefined;
            resourceInputs["replicaSets"] = state ? state.replicaSets : undefined;
            resourceInputs["replicationFactor"] = state ? state.replicationFactor : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["retentionPeriod"] = state ? state.retentionPeriod : undefined;
            resourceInputs["secondaryZoneId"] = state ? state.secondaryZoneId : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["securityIpLists"] = state ? state.securityIpLists : undefined;
            resourceInputs["sslAction"] = state ? state.sslAction : undefined;
            resourceInputs["sslStatus"] = state ? state.sslStatus : undefined;
            resourceInputs["storageEngine"] = state ? state.storageEngine : undefined;
            resourceInputs["storageType"] = state ? state.storageType : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tdeStatus"] = state ? state.tdeStatus : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.dbInstanceClass === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbInstanceClass'");
            }
            if ((!args || args.dbInstanceStorage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbInstanceStorage'");
            }
            if ((!args || args.engineVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineVersion'");
            }
            resourceInputs["accountPassword"] = args?.accountPassword ? pulumi.secret(args.accountPassword) : undefined;
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["backupPeriods"] = args ? args.backupPeriods : undefined;
            resourceInputs["backupTime"] = args ? args.backupTime : undefined;
            resourceInputs["dbInstanceClass"] = args ? args.dbInstanceClass : undefined;
            resourceInputs["dbInstanceStorage"] = args ? args.dbInstanceStorage : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["hiddenZoneId"] = args ? args.hiddenZoneId : undefined;
            resourceInputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            resourceInputs["kmsEncryptedPassword"] = args ? args.kmsEncryptedPassword : undefined;
            resourceInputs["kmsEncryptionContext"] = args ? args.kmsEncryptionContext : undefined;
            resourceInputs["maintainEndTime"] = args ? args.maintainEndTime : undefined;
            resourceInputs["maintainStartTime"] = args ? args.maintainStartTime : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["orderType"] = args ? args.orderType : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["readonlyReplicas"] = args ? args.readonlyReplicas : undefined;
            resourceInputs["replicationFactor"] = args ? args.replicationFactor : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["secondaryZoneId"] = args ? args.secondaryZoneId : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["securityIpLists"] = args ? args.securityIpLists : undefined;
            resourceInputs["sslAction"] = args ? args.sslAction : undefined;
            resourceInputs["storageEngine"] = args ? args.storageEngine : undefined;
            resourceInputs["storageType"] = args ? args.storageType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tdeStatus"] = args ? args.tdeStatus : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["replicaSetName"] = undefined /*out*/;
            resourceInputs["replicaSets"] = undefined /*out*/;
            resourceInputs["retentionPeriod"] = undefined /*out*/;
            resourceInputs["sslStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accountPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * Auto renew for prepaid, true of false. Default is false.
     * > **NOTE:** The start time to the end time must be 1 hour. For example, the MaintainStartTime is 01:00Z, then the MaintainEndTime must be 02:00Z.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
     */
    backupPeriods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
     */
    backupTime?: pulumi.Input<string>;
    /**
     * Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
     */
    dbInstanceClass?: pulumi.Input<string>;
    /**
     * User-defined DB instance storage space.Unit: GB. Value range:
     * - Custom storage space.
     * - 10-GB increments.
     */
    dbInstanceStorage?: pulumi.Input<number>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/61763.htm) `EngineVersion`.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Configure the zone where the hidden node is located to deploy multiple zones. **NOTE:** This parameter value cannot be the same as `zoneId` and `secondaryZoneId` parameter values.
     */
    hiddenZoneId?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version 1.63.0.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
     */
    kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    maintainEndTime?: pulumi.Input<string>;
    /**
     * The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    maintainStartTime?: pulumi.Input<string>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
     */
    networkType?: pulumi.Input<string>;
    /**
     * The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
     * * UPGRADE: The specifications are upgraded.
     * * DOWNGRADE: The specifications are downgraded.
     * **NOTE:** This parameter is only applicable to instances when `instanceChargeType` is PrePaid.
     */
    orderType?: pulumi.Input<string>;
    /**
     * Set of parameters needs to be set after mongodb instance was launched. See `parameters` below.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.mongodb.InstanceParameter>[]>;
    /**
     * The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
     */
    period?: pulumi.Input<number>;
    /**
     * The number of read-only nodes in the replica set instance. Default value: 0. Valid values: 0 to 5.
     */
    readonlyReplicas?: pulumi.Input<number>;
    /**
     * The name of the mongo replica set
     */
    replicaSetName?: pulumi.Input<string>;
    /**
     * Replica set instance information. The details see Block replica_sets. **NOTE:** Available since v1.140. See `replicaSets` below.
     */
    replicaSets?: pulumi.Input<pulumi.Input<inputs.mongodb.InstanceReplicaSet>[]>;
    /**
     * Number of replica set nodes. Valid values: [1, 3, 5, 7]
     */
    replicationFactor?: pulumi.Input<number>;
    /**
     * The ID of the Resource Group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Instance log backup retention days. Available in 1.42.0+.
     */
    retentionPeriod?: pulumi.Input<number>;
    /**
     * Configure the available area where the slave node (Secondary node) is located to realize multi-available area deployment. **NOTE:** This parameter value cannot be the same as `zoneId` and `hiddenZoneId` parameter values.
     */
    secondaryZoneId?: pulumi.Input<string>;
    /**
     * The Security Group ID of ECS.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     */
    securityIpLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Actions performed on SSL functions, Valid values: `Open`: turn on SSL encryption; `Close`: turn off SSL encryption; `Update`: update SSL certificate.
     */
    sslAction?: pulumi.Input<string>;
    /**
     * Status of the SSL feature. `Open`: SSL is turned on; `Closed`: SSL is turned off.
     */
    sslStatus?: pulumi.Input<string>;
    /**
     * Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
     */
    storageEngine?: pulumi.Input<string>;
    /**
     * The storage type of the instance. Valid values: `cloudEssd1`, `cloudEssd2`, `cloudEssd3`, `localSsd`.
     */
    storageType?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The TDE(Transparent Data Encryption) status.
     */
    tdeStatus?: pulumi.Input<string>;
    /**
     * The ID of the VPC. > **NOTE:** This parameter is valid only when NetworkType is set to VPC.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DB instance. it supports multiple zone.
     * If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
     * The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `alicloud.getZones`.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * Auto renew for prepaid, true of false. Default is false.
     * > **NOTE:** The start time to the end time must be 1 hour. For example, the MaintainStartTime is 01:00Z, then the MaintainEndTime must be 02:00Z.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
     */
    backupPeriods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
     */
    backupTime?: pulumi.Input<string>;
    /**
     * Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
     */
    dbInstanceClass: pulumi.Input<string>;
    /**
     * User-defined DB instance storage space.Unit: GB. Value range:
     * - Custom storage space.
     * - 10-GB increments.
     */
    dbInstanceStorage: pulumi.Input<number>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/61763.htm) `EngineVersion`.
     */
    engineVersion: pulumi.Input<string>;
    /**
     * Configure the zone where the hidden node is located to deploy multiple zones. **NOTE:** This parameter value cannot be the same as `zoneId` and `secondaryZoneId` parameter values.
     */
    hiddenZoneId?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version 1.63.0.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
     */
    kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    maintainEndTime?: pulumi.Input<string>;
    /**
     * The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    maintainStartTime?: pulumi.Input<string>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
     */
    networkType?: pulumi.Input<string>;
    /**
     * The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
     * * UPGRADE: The specifications are upgraded.
     * * DOWNGRADE: The specifications are downgraded.
     * **NOTE:** This parameter is only applicable to instances when `instanceChargeType` is PrePaid.
     */
    orderType?: pulumi.Input<string>;
    /**
     * Set of parameters needs to be set after mongodb instance was launched. See `parameters` below.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.mongodb.InstanceParameter>[]>;
    /**
     * The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
     */
    period?: pulumi.Input<number>;
    /**
     * The number of read-only nodes in the replica set instance. Default value: 0. Valid values: 0 to 5.
     */
    readonlyReplicas?: pulumi.Input<number>;
    /**
     * Number of replica set nodes. Valid values: [1, 3, 5, 7]
     */
    replicationFactor?: pulumi.Input<number>;
    /**
     * The ID of the Resource Group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Configure the available area where the slave node (Secondary node) is located to realize multi-available area deployment. **NOTE:** This parameter value cannot be the same as `zoneId` and `hiddenZoneId` parameter values.
     */
    secondaryZoneId?: pulumi.Input<string>;
    /**
     * The Security Group ID of ECS.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     */
    securityIpLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Actions performed on SSL functions, Valid values: `Open`: turn on SSL encryption; `Close`: turn off SSL encryption; `Update`: update SSL certificate.
     */
    sslAction?: pulumi.Input<string>;
    /**
     * Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
     */
    storageEngine?: pulumi.Input<string>;
    /**
     * The storage type of the instance. Valid values: `cloudEssd1`, `cloudEssd2`, `cloudEssd3`, `localSsd`.
     */
    storageType?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The TDE(Transparent Data Encryption) status.
     */
    tdeStatus?: pulumi.Input<string>;
    /**
     * The ID of the VPC. > **NOTE:** This parameter is valid only when NetworkType is set to VPC.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DB instance. it supports multiple zone.
     * If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
     * The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `alicloud.getZones`.
     */
    zoneId?: pulumi.Input<string>;
}
