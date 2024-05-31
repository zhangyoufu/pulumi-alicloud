// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Redis Tair Instance can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:redis/tairInstance:TairInstance example <id>
 * ```
 */
export class TairInstance extends pulumi.CustomResource {
    /**
     * Get an existing TairInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TairInstanceState, opts?: pulumi.CustomResourceOptions): TairInstance {
        return new TairInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:redis/tairInstance:TairInstance';

    /**
     * Returns true if the given object is an instance of TairInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TairInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TairInstance.__pulumiType;
    }

    /**
     * Specifies whether to enable auto-renewal for the instance. Default value: false. Valid values: true(enables auto-renewal), false(disables auto-renewal).
     */
    public readonly autoRenew!: pulumi.Output<string | undefined>;
    /**
     * The subscription duration that is supported by auto-renewal. Unit: months. Valid values: 1, 2, 3, 6, and 12. This parameter is required only if the AutoRenew parameter is set to true.
     */
    public readonly autoRenewPeriod!: pulumi.Output<string | undefined>;
    /**
     * The ID of the backup set of the cluster.  .
     */
    public readonly clusterBackupId!: pulumi.Output<string | undefined>;
    /**
     * The time when the instance was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The time when to change the configurations. Default value: Immediately. Valid values: Immediately (The configurations are immediately changed), MaintainTime (The configurations are changed within the maintenance window).
     */
    public readonly effectiveTime!: pulumi.Output<string | undefined>;
    /**
     * Database version. Default value: 1.0.  Rules for transferring parameters of different tair product types:  tair_rdb:  Compatible with the Redis5.0 and Redis6.0 protocols, and is transmitted to 5.0 or 6.0. tair_scm: The Tair persistent memory is compatible with the Redis6.0 protocol and is passed 1.0. tair_essd: The disk (ESSD/SSD) is compatible with the Redis4.0 and Redis6.0 protocols, and is transmitted to 1.0 and 2.0 respectively.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * Specifies whether to forcefully change the configurations of the instance. Default value: true. Valid values: false (The system does not forcefully change the configurations), true (The system forcefully changes the configurations).
     */
    public readonly forceUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * The instance type of the instance. For more information, see [Instance types](https://www.alibabacloud.com/help/en/apsaradb-for-redis/latest/instance-types).
     */
    public readonly instanceClass!: pulumi.Output<string>;
    /**
     * The storage medium of the instance. Valid values: tair_rdb, tair_scm, tair_essd.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * Node type, value:  MASTER_SLAVE: high availability (dual copy) STAND_ALONE: single copy double: double copy single: single copy Note For Cloud Native instances, select MASTER_SLAVE or STAND_ALONE. For Classic instances, select double or single.
     */
    public readonly nodeType!: pulumi.Output<string>;
    /**
     * The password that is used to connect to the instance. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include ! @ # $ % ^ & * ( ) _ + - =.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The billing method of the instance. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
     */
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * The subscription duration. Unit: months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24,36, and 60. This parameter is required only if you set the PaymentType parameter to Subscription.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The Tair service port. The service port of the instance. Valid values: 1024 to 65535. Default value: 6379.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Number of read-only nodes in the primary zone. Valid values: 0 to 5. This parameter is only applicable to the following conditions:  If the instance is in the cloud disk version standard architecture, you can set this parameter to a value greater than 0 to enable the read/write splitting architecture. If the instance is a cloud disk version read/write splitting architecture instance, you can use this parameter to customize the number of read-only nodes, or set this parameter to 0 to disable the read/write splitting architecture and switch the instance to the standard architecture.
     */
    public readonly readOnlyCount!: pulumi.Output<number | undefined>;
    /**
     * The ID of the resource group to which the instance belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The ID of the secondary zone.This parameter is returned only if the instance is deployed in two zones.
     */
    public readonly secondaryZoneId!: pulumi.Output<string | undefined>;
    /**
     * The number of data nodes in the instance. When 1 is passed, it means that the instance created is a standard architecture with only one data node. You can create an instance in the standard architecture that contains only a single data node. 2 to 32: You can create an instance in the cluster architecture that contains the specified number of data nodes. Only persistent memory-optimized instances can use the cluster architecture. Therefore, you can set this parameter to an integer from 2 to 32 only if you set the InstanceType parameter to tair_scm. It is not allowed to modify the number of shards by modifying this parameter after creating a master-slave architecture instance with or without passing 1.
     */
    public readonly shardCount!: pulumi.Output<number>;
    /**
     * Specifies the number of read-only nodes in the secondary zone when creating a multi-zone read/write splitting instance. Note: To create a multi-zone read/write splitting instance, slaveadonlycount and SecondaryZoneId must be specified at the same time.
     */
    public readonly slaveReadOnlyCount!: pulumi.Output<number | undefined>;
    /**
     * The status of the resource.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The storage type. The value range is [PL1, PL2, and PL3]. The default value is PL1. When the value of instanceType is "tairEssd", this attribute takes effect and is required.
     */
    public readonly storagePerformanceLevel!: pulumi.Output<string | undefined>;
    /**
     * The value range of different specifications is different, see [ESSD-based instances](https://www.alibabacloud.com/help/en/tair/product-overview/essd-based-instances). When the value of instanceType is "tairEssd", this attribute takes effect and is required.
     */
    public readonly storageSizeGb!: pulumi.Output<number>;
    /**
     * The tag of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The name of the resource.
     */
    public readonly tairInstanceName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the virtual private cloud (VPC).
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The ID of the vSwitch to which the instance is connected.
     */
    public readonly vswitchId!: pulumi.Output<string>;
    /**
     * Zone ID.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a TairInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TairInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TairInstanceArgs | TairInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TairInstanceState | undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["autoRenewPeriod"] = state ? state.autoRenewPeriod : undefined;
            resourceInputs["clusterBackupId"] = state ? state.clusterBackupId : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["effectiveTime"] = state ? state.effectiveTime : undefined;
            resourceInputs["engineVersion"] = state ? state.engineVersion : undefined;
            resourceInputs["forceUpgrade"] = state ? state.forceUpgrade : undefined;
            resourceInputs["instanceClass"] = state ? state.instanceClass : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["nodeType"] = state ? state.nodeType : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["readOnlyCount"] = state ? state.readOnlyCount : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["secondaryZoneId"] = state ? state.secondaryZoneId : undefined;
            resourceInputs["shardCount"] = state ? state.shardCount : undefined;
            resourceInputs["slaveReadOnlyCount"] = state ? state.slaveReadOnlyCount : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["storagePerformanceLevel"] = state ? state.storagePerformanceLevel : undefined;
            resourceInputs["storageSizeGb"] = state ? state.storageSizeGb : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tairInstanceName"] = state ? state.tairInstanceName : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as TairInstanceArgs | undefined;
            if ((!args || args.instanceClass === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceClass'");
            }
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["autoRenewPeriod"] = args ? args.autoRenewPeriod : undefined;
            resourceInputs["clusterBackupId"] = args ? args.clusterBackupId : undefined;
            resourceInputs["effectiveTime"] = args ? args.effectiveTime : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["forceUpgrade"] = args ? args.forceUpgrade : undefined;
            resourceInputs["instanceClass"] = args ? args.instanceClass : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["nodeType"] = args ? args.nodeType : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["readOnlyCount"] = args ? args.readOnlyCount : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["secondaryZoneId"] = args ? args.secondaryZoneId : undefined;
            resourceInputs["shardCount"] = args ? args.shardCount : undefined;
            resourceInputs["slaveReadOnlyCount"] = args ? args.slaveReadOnlyCount : undefined;
            resourceInputs["storagePerformanceLevel"] = args ? args.storagePerformanceLevel : undefined;
            resourceInputs["storageSizeGb"] = args ? args.storageSizeGb : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tairInstanceName"] = args ? args.tairInstanceName : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(TairInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TairInstance resources.
 */
export interface TairInstanceState {
    /**
     * Specifies whether to enable auto-renewal for the instance. Default value: false. Valid values: true(enables auto-renewal), false(disables auto-renewal).
     */
    autoRenew?: pulumi.Input<string>;
    /**
     * The subscription duration that is supported by auto-renewal. Unit: months. Valid values: 1, 2, 3, 6, and 12. This parameter is required only if the AutoRenew parameter is set to true.
     */
    autoRenewPeriod?: pulumi.Input<string>;
    /**
     * The ID of the backup set of the cluster.  .
     */
    clusterBackupId?: pulumi.Input<string>;
    /**
     * The time when the instance was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The time when to change the configurations. Default value: Immediately. Valid values: Immediately (The configurations are immediately changed), MaintainTime (The configurations are changed within the maintenance window).
     */
    effectiveTime?: pulumi.Input<string>;
    /**
     * Database version. Default value: 1.0.  Rules for transferring parameters of different tair product types:  tair_rdb:  Compatible with the Redis5.0 and Redis6.0 protocols, and is transmitted to 5.0 or 6.0. tair_scm: The Tair persistent memory is compatible with the Redis6.0 protocol and is passed 1.0. tair_essd: The disk (ESSD/SSD) is compatible with the Redis4.0 and Redis6.0 protocols, and is transmitted to 1.0 and 2.0 respectively.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Specifies whether to forcefully change the configurations of the instance. Default value: true. Valid values: false (The system does not forcefully change the configurations), true (The system forcefully changes the configurations).
     */
    forceUpgrade?: pulumi.Input<boolean>;
    /**
     * The instance type of the instance. For more information, see [Instance types](https://www.alibabacloud.com/help/en/apsaradb-for-redis/latest/instance-types).
     */
    instanceClass?: pulumi.Input<string>;
    /**
     * The storage medium of the instance. Valid values: tair_rdb, tair_scm, tair_essd.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Node type, value:  MASTER_SLAVE: high availability (dual copy) STAND_ALONE: single copy double: double copy single: single copy Note For Cloud Native instances, select MASTER_SLAVE or STAND_ALONE. For Classic instances, select double or single.
     */
    nodeType?: pulumi.Input<string>;
    /**
     * The password that is used to connect to the instance. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include ! @ # $ % ^ & * ( ) _ + - =.
     */
    password?: pulumi.Input<string>;
    /**
     * The billing method of the instance. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * The subscription duration. Unit: months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24,36, and 60. This parameter is required only if you set the PaymentType parameter to Subscription.
     */
    period?: pulumi.Input<number>;
    /**
     * The Tair service port. The service port of the instance. Valid values: 1024 to 65535. Default value: 6379.
     */
    port?: pulumi.Input<number>;
    /**
     * Number of read-only nodes in the primary zone. Valid values: 0 to 5. This parameter is only applicable to the following conditions:  If the instance is in the cloud disk version standard architecture, you can set this parameter to a value greater than 0 to enable the read/write splitting architecture. If the instance is a cloud disk version read/write splitting architecture instance, you can use this parameter to customize the number of read-only nodes, or set this parameter to 0 to disable the read/write splitting architecture and switch the instance to the standard architecture.
     */
    readOnlyCount?: pulumi.Input<number>;
    /**
     * The ID of the resource group to which the instance belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The ID of the secondary zone.This parameter is returned only if the instance is deployed in two zones.
     */
    secondaryZoneId?: pulumi.Input<string>;
    /**
     * The number of data nodes in the instance. When 1 is passed, it means that the instance created is a standard architecture with only one data node. You can create an instance in the standard architecture that contains only a single data node. 2 to 32: You can create an instance in the cluster architecture that contains the specified number of data nodes. Only persistent memory-optimized instances can use the cluster architecture. Therefore, you can set this parameter to an integer from 2 to 32 only if you set the InstanceType parameter to tair_scm. It is not allowed to modify the number of shards by modifying this parameter after creating a master-slave architecture instance with or without passing 1.
     */
    shardCount?: pulumi.Input<number>;
    /**
     * Specifies the number of read-only nodes in the secondary zone when creating a multi-zone read/write splitting instance. Note: To create a multi-zone read/write splitting instance, slaveadonlycount and SecondaryZoneId must be specified at the same time.
     */
    slaveReadOnlyCount?: pulumi.Input<number>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
    /**
     * The storage type. The value range is [PL1, PL2, and PL3]. The default value is PL1. When the value of instanceType is "tairEssd", this attribute takes effect and is required.
     */
    storagePerformanceLevel?: pulumi.Input<string>;
    /**
     * The value range of different specifications is different, see [ESSD-based instances](https://www.alibabacloud.com/help/en/tair/product-overview/essd-based-instances). When the value of instanceType is "tairEssd", this attribute takes effect and is required.
     */
    storageSizeGb?: pulumi.Input<number>;
    /**
     * The tag of the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the resource.
     */
    tairInstanceName?: pulumi.Input<string>;
    /**
     * The ID of the virtual private cloud (VPC).
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The ID of the vSwitch to which the instance is connected.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * Zone ID.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TairInstance resource.
 */
export interface TairInstanceArgs {
    /**
     * Specifies whether to enable auto-renewal for the instance. Default value: false. Valid values: true(enables auto-renewal), false(disables auto-renewal).
     */
    autoRenew?: pulumi.Input<string>;
    /**
     * The subscription duration that is supported by auto-renewal. Unit: months. Valid values: 1, 2, 3, 6, and 12. This parameter is required only if the AutoRenew parameter is set to true.
     */
    autoRenewPeriod?: pulumi.Input<string>;
    /**
     * The ID of the backup set of the cluster.  .
     */
    clusterBackupId?: pulumi.Input<string>;
    /**
     * The time when to change the configurations. Default value: Immediately. Valid values: Immediately (The configurations are immediately changed), MaintainTime (The configurations are changed within the maintenance window).
     */
    effectiveTime?: pulumi.Input<string>;
    /**
     * Database version. Default value: 1.0.  Rules for transferring parameters of different tair product types:  tair_rdb:  Compatible with the Redis5.0 and Redis6.0 protocols, and is transmitted to 5.0 or 6.0. tair_scm: The Tair persistent memory is compatible with the Redis6.0 protocol and is passed 1.0. tair_essd: The disk (ESSD/SSD) is compatible with the Redis4.0 and Redis6.0 protocols, and is transmitted to 1.0 and 2.0 respectively.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Specifies whether to forcefully change the configurations of the instance. Default value: true. Valid values: false (The system does not forcefully change the configurations), true (The system forcefully changes the configurations).
     */
    forceUpgrade?: pulumi.Input<boolean>;
    /**
     * The instance type of the instance. For more information, see [Instance types](https://www.alibabacloud.com/help/en/apsaradb-for-redis/latest/instance-types).
     */
    instanceClass: pulumi.Input<string>;
    /**
     * The storage medium of the instance. Valid values: tair_rdb, tair_scm, tair_essd.
     */
    instanceType: pulumi.Input<string>;
    /**
     * Node type, value:  MASTER_SLAVE: high availability (dual copy) STAND_ALONE: single copy double: double copy single: single copy Note For Cloud Native instances, select MASTER_SLAVE or STAND_ALONE. For Classic instances, select double or single.
     */
    nodeType?: pulumi.Input<string>;
    /**
     * The password that is used to connect to the instance. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include ! @ # $ % ^ & * ( ) _ + - =.
     */
    password?: pulumi.Input<string>;
    /**
     * The billing method of the instance. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * The subscription duration. Unit: months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24,36, and 60. This parameter is required only if you set the PaymentType parameter to Subscription.
     */
    period?: pulumi.Input<number>;
    /**
     * The Tair service port. The service port of the instance. Valid values: 1024 to 65535. Default value: 6379.
     */
    port?: pulumi.Input<number>;
    /**
     * Number of read-only nodes in the primary zone. Valid values: 0 to 5. This parameter is only applicable to the following conditions:  If the instance is in the cloud disk version standard architecture, you can set this parameter to a value greater than 0 to enable the read/write splitting architecture. If the instance is a cloud disk version read/write splitting architecture instance, you can use this parameter to customize the number of read-only nodes, or set this parameter to 0 to disable the read/write splitting architecture and switch the instance to the standard architecture.
     */
    readOnlyCount?: pulumi.Input<number>;
    /**
     * The ID of the resource group to which the instance belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The ID of the secondary zone.This parameter is returned only if the instance is deployed in two zones.
     */
    secondaryZoneId?: pulumi.Input<string>;
    /**
     * The number of data nodes in the instance. When 1 is passed, it means that the instance created is a standard architecture with only one data node. You can create an instance in the standard architecture that contains only a single data node. 2 to 32: You can create an instance in the cluster architecture that contains the specified number of data nodes. Only persistent memory-optimized instances can use the cluster architecture. Therefore, you can set this parameter to an integer from 2 to 32 only if you set the InstanceType parameter to tair_scm. It is not allowed to modify the number of shards by modifying this parameter after creating a master-slave architecture instance with or without passing 1.
     */
    shardCount?: pulumi.Input<number>;
    /**
     * Specifies the number of read-only nodes in the secondary zone when creating a multi-zone read/write splitting instance. Note: To create a multi-zone read/write splitting instance, slaveadonlycount and SecondaryZoneId must be specified at the same time.
     */
    slaveReadOnlyCount?: pulumi.Input<number>;
    /**
     * The storage type. The value range is [PL1, PL2, and PL3]. The default value is PL1. When the value of instanceType is "tairEssd", this attribute takes effect and is required.
     */
    storagePerformanceLevel?: pulumi.Input<string>;
    /**
     * The value range of different specifications is different, see [ESSD-based instances](https://www.alibabacloud.com/help/en/tair/product-overview/essd-based-instances). When the value of instanceType is "tairEssd", this attribute takes effect and is required.
     */
    storageSizeGb?: pulumi.Input<number>;
    /**
     * The tag of the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the resource.
     */
    tairInstanceName?: pulumi.Input<string>;
    /**
     * The ID of the virtual private cloud (VPC).
     */
    vpcId: pulumi.Input<string>;
    /**
     * The ID of the vSwitch to which the instance is connected.
     */
    vswitchId: pulumi.Input<string>;
    /**
     * Zone ID.
     */
    zoneId: pulumi.Input<string>;
}
