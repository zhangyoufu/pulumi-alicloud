// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an ApsaraDB Redis / Memcache instance resource. A DB instance is an isolated database environment in the cloud. It can be associated with IP whitelists and backup configuration which are separate resource providers.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const creation = config.get("creation") || "KVStore";
 * const name = config.get("name") || "kvstoreinstancevpc";
 *
 * const defaultZones = pulumi.output(alicloud.getZones({
 *     availableResourceCreation: creation,
 * }, { async: true }));
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "172.16.0.0/24",
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultInstance = new alicloud.kvstore.Instance("default", {
 *     engineVersion: "4.0",
 *     instanceClass: "redis.master.small.default",
 *     instanceName: name,
 *     instanceType: "Redis",
 *     privateIp: "172.16.0.10",
 *     securityIps: ["10.0.0.1"],
 *     vswitchId: defaultSwitch.id,
 * });
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
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:kvstore/instance:Instance';

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
     * Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
     */
    public readonly autoRenewPeriod!: pulumi.Output<number | undefined>;
    /**
     * The Zone to launch the DB instance.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * If an instance created based on a backup set generated by another instance is valid, this parameter indicates the ID of the generated backup set.
     */
    public readonly backupId!: pulumi.Output<string | undefined>;
    /**
     * Instance connection domain (only Intranet access supported).
     */
    public /*out*/ readonly connectionDomain!: pulumi.Output<string>;
    /**
     * Engine version. Supported values: 2.8, 4.0 and 5.0. Default value: 2.8. Only 2.8 can be supported for Memcache Instance.
     */
    public readonly engineVersion!: pulumi.Output<string | undefined>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
     */
    public readonly instanceChargeType!: pulumi.Output<string | undefined>;
    public readonly instanceClass!: pulumi.Output<string>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    public readonly instanceName!: pulumi.Output<string | undefined>;
    /**
     * The engine to use: `Redis` or `Memcache`. Defaults to `Redis`.
     */
    public readonly instanceType!: pulumi.Output<string | undefined>;
    /**
     * An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored.
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
     * Set of parameters needs to be set after instance was launched. Available parameters can refer to the latest docs [Instance configurations table](https://www.alibabacloud.com/help/doc-detail/61209.htm) .
     */
    public readonly parameters!: pulumi.Output<outputs.kvstore.InstanceParameter[]>;
    /**
     * The password of the DB instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Set the instance's private IP.
     */
    public readonly privateIp!: pulumi.Output<string>;
    /**
     * The Security Group ID of ECS.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * Set the instance's IP whitelist of the default security group.
     */
    public readonly securityIps!: pulumi.Output<string[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Only meaningful if instanceType is `Redis` and network type is VPC. Valid values are `Close`, `Open`. Defaults to `Open`.  `Close` means the redis instance can be accessed without authentication. `Open` means authentication is required.
     */
    public readonly vpcAuthMode!: pulumi.Output<string>;
    /**
     * The ID of VSwitch.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceState | undefined;
            inputs["autoRenew"] = state ? state.autoRenew : undefined;
            inputs["autoRenewPeriod"] = state ? state.autoRenewPeriod : undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["backupId"] = state ? state.backupId : undefined;
            inputs["connectionDomain"] = state ? state.connectionDomain : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            inputs["instanceClass"] = state ? state.instanceClass : undefined;
            inputs["instanceName"] = state ? state.instanceName : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["kmsEncryptedPassword"] = state ? state.kmsEncryptedPassword : undefined;
            inputs["kmsEncryptionContext"] = state ? state.kmsEncryptionContext : undefined;
            inputs["maintainEndTime"] = state ? state.maintainEndTime : undefined;
            inputs["maintainStartTime"] = state ? state.maintainStartTime : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["privateIp"] = state ? state.privateIp : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["securityIps"] = state ? state.securityIps : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcAuthMode"] = state ? state.vpcAuthMode : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if (!args || args.instanceClass === undefined) {
                throw new Error("Missing required property 'instanceClass'");
            }
            inputs["autoRenew"] = args ? args.autoRenew : undefined;
            inputs["autoRenewPeriod"] = args ? args.autoRenewPeriod : undefined;
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["backupId"] = args ? args.backupId : undefined;
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            inputs["instanceClass"] = args ? args.instanceClass : undefined;
            inputs["instanceName"] = args ? args.instanceName : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["kmsEncryptedPassword"] = args ? args.kmsEncryptedPassword : undefined;
            inputs["kmsEncryptionContext"] = args ? args.kmsEncryptionContext : undefined;
            inputs["maintainEndTime"] = args ? args.maintainEndTime : undefined;
            inputs["maintainStartTime"] = args ? args.maintainStartTime : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["privateIp"] = args ? args.privateIp : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            inputs["securityIps"] = args ? args.securityIps : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcAuthMode"] = args ? args.vpcAuthMode : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["connectionDomain"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
     */
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
     */
    readonly autoRenewPeriod?: pulumi.Input<number>;
    /**
     * The Zone to launch the DB instance.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * If an instance created based on a backup set generated by another instance is valid, this parameter indicates the ID of the generated backup set.
     */
    readonly backupId?: pulumi.Input<string>;
    /**
     * Instance connection domain (only Intranet access supported).
     */
    readonly connectionDomain?: pulumi.Input<string>;
    /**
     * Engine version. Supported values: 2.8, 4.0 and 5.0. Default value: 2.8. Only 2.8 can be supported for Memcache Instance.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    readonly instanceClass?: pulumi.Input<string>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    readonly instanceName?: pulumi.Input<string>;
    /**
     * The engine to use: `Redis` or `Memcache`. Defaults to `Redis`.
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored.
     */
    readonly kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    readonly kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    readonly maintainEndTime?: pulumi.Input<string>;
    /**
     * The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    readonly maintainStartTime?: pulumi.Input<string>;
    /**
     * Set of parameters needs to be set after instance was launched. Available parameters can refer to the latest docs [Instance configurations table](https://www.alibabacloud.com/help/doc-detail/61209.htm) .
     */
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.kvstore.InstanceParameter>[]>;
    /**
     * The password of the DB instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * Set the instance's private IP.
     */
    readonly privateIp?: pulumi.Input<string>;
    /**
     * The Security Group ID of ECS.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * Set the instance's IP whitelist of the default security group.
     */
    readonly securityIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Only meaningful if instanceType is `Redis` and network type is VPC. Valid values are `Close`, `Open`. Defaults to `Open`.  `Close` means the redis instance can be accessed without authentication. `Open` means authentication is required.
     */
    readonly vpcAuthMode?: pulumi.Input<string>;
    /**
     * The ID of VSwitch.
     */
    readonly vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
     */
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
     */
    readonly autoRenewPeriod?: pulumi.Input<number>;
    /**
     * The Zone to launch the DB instance.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * If an instance created based on a backup set generated by another instance is valid, this parameter indicates the ID of the generated backup set.
     */
    readonly backupId?: pulumi.Input<string>;
    /**
     * Engine version. Supported values: 2.8, 4.0 and 5.0. Default value: 2.8. Only 2.8 can be supported for Memcache Instance.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    readonly instanceClass: pulumi.Input<string>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    readonly instanceName?: pulumi.Input<string>;
    /**
     * The engine to use: `Redis` or `Memcache`. Defaults to `Redis`.
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored.
     */
    readonly kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    readonly kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    readonly maintainEndTime?: pulumi.Input<string>;
    /**
     * The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    readonly maintainStartTime?: pulumi.Input<string>;
    /**
     * Set of parameters needs to be set after instance was launched. Available parameters can refer to the latest docs [Instance configurations table](https://www.alibabacloud.com/help/doc-detail/61209.htm) .
     */
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.kvstore.InstanceParameter>[]>;
    /**
     * The password of the DB instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * Set the instance's private IP.
     */
    readonly privateIp?: pulumi.Input<string>;
    /**
     * The Security Group ID of ECS.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * Set the instance's IP whitelist of the default security group.
     */
    readonly securityIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Only meaningful if instanceType is `Redis` and network type is VPC. Valid values are `Close`, `Open`. Defaults to `Open`.  `Close` means the redis instance can be accessed without authentication. `Open` means authentication is required.
     */
    readonly vpcAuthMode?: pulumi.Input<string>;
    /**
     * The ID of VSwitch.
     */
    readonly vswitchId?: pulumi.Input<string>;
}
