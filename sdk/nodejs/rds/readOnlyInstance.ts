// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an RDS readonly instance resource. 
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
 * const creation = config.get("creation") || "Rds";
 * const name = config.get("name") || "dbInstancevpc";
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
 * const defaultInstance = new alicloud.rds.Instance("default", {
 *     engine: "MySQL",
 *     engineVersion: "5.6",
 *     instanceChargeType: "Postpaid",
 *     instanceName: name,
 *     instanceStorage: 20,
 *     instanceType: "rds.mysql.t1.small",
 *     securityIps: [
 *         "10.168.1.12",
 *         "100.69.7.112",
 *     ],
 *     vswitchId: defaultSwitch.id,
 * });
 * const defaultReadOnlyInstance = new alicloud.rds.ReadOnlyInstance("default", {
 *     engineVersion: defaultInstance.engineVersion,
 *     instanceName: `${name}ro`,
 *     instanceStorage: 30,
 *     instanceType: defaultInstance.instanceType,
 *     masterDbInstanceId: defaultInstance.id,
 *     vswitchId: defaultSwitch.id,
 *     zoneId: defaultInstance.zoneId,
 * });
 * ```
 */
export class ReadOnlyInstance extends pulumi.CustomResource {
    /**
     * Get an existing ReadOnlyInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReadOnlyInstanceState, opts?: pulumi.CustomResourceOptions): ReadOnlyInstance {
        return new ReadOnlyInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:rds/readOnlyInstance:ReadOnlyInstance';

    /**
     * Returns true if the given object is an instance of ReadOnlyInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReadOnlyInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReadOnlyInstance.__pulumiType;
    }

    /**
     * RDS database connection string.
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * Database type.
     */
    public /*out*/ readonly engine!: pulumi.Output<string>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * User-defined DB instance storage space. Value range: [5, 2000] for MySQL/SQL Server HA dual node edition. Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    public readonly instanceStorage!: pulumi.Output<number>;
    /**
     * DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * ID of the master instance.
     */
    public readonly masterDbInstanceId!: pulumi.Output<string>;
    /**
     * Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm).
     */
    public readonly parameters!: pulumi.Output<outputs.rds.ReadOnlyInstanceParameter[]>;
    /**
     * RDS database connection port.
     */
    public /*out*/ readonly port!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;
    /**
     * The Zone to launch the DB instance.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a ReadOnlyInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReadOnlyInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReadOnlyInstanceArgs | ReadOnlyInstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ReadOnlyInstanceState | undefined;
            inputs["connectionString"] = state ? state.connectionString : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["instanceName"] = state ? state.instanceName : undefined;
            inputs["instanceStorage"] = state ? state.instanceStorage : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["masterDbInstanceId"] = state ? state.masterDbInstanceId : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ReadOnlyInstanceArgs | undefined;
            if (!args || args.engineVersion === undefined) {
                throw new Error("Missing required property 'engineVersion'");
            }
            if (!args || args.instanceStorage === undefined) {
                throw new Error("Missing required property 'instanceStorage'");
            }
            if (!args || args.instanceType === undefined) {
                throw new Error("Missing required property 'instanceType'");
            }
            if (!args || args.masterDbInstanceId === undefined) {
                throw new Error("Missing required property 'masterDbInstanceId'");
            }
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["instanceName"] = args ? args.instanceName : undefined;
            inputs["instanceStorage"] = args ? args.instanceStorage : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["masterDbInstanceId"] = args ? args.masterDbInstanceId : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
            inputs["connectionString"] = undefined /*out*/;
            inputs["engine"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ReadOnlyInstance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReadOnlyInstance resources.
 */
export interface ReadOnlyInstanceState {
    /**
     * RDS database connection string.
     */
    readonly connectionString?: pulumi.Input<string>;
    /**
     * Database type.
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    readonly instanceName?: pulumi.Input<string>;
    /**
     * User-defined DB instance storage space. Value range: [5, 2000] for MySQL/SQL Server HA dual node edition. Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    readonly instanceStorage?: pulumi.Input<number>;
    /**
     * DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * ID of the master instance.
     */
    readonly masterDbInstanceId?: pulumi.Input<string>;
    /**
     * Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm).
     */
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.rds.ReadOnlyInstanceParameter>[]>;
    /**
     * RDS database connection port.
     */
    readonly port?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    readonly vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DB instance.
     */
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReadOnlyInstance resource.
 */
export interface ReadOnlyInstanceArgs {
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
     */
    readonly engineVersion: pulumi.Input<string>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    readonly instanceName?: pulumi.Input<string>;
    /**
     * User-defined DB instance storage space. Value range: [5, 2000] for MySQL/SQL Server HA dual node edition. Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    readonly instanceStorage: pulumi.Input<number>;
    /**
     * DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    readonly instanceType: pulumi.Input<string>;
    /**
     * ID of the master instance.
     */
    readonly masterDbInstanceId: pulumi.Input<string>;
    /**
     * Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm).
     */
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.rds.ReadOnlyInstanceParameter>[]>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    readonly vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DB instance.
     */
    readonly zoneId?: pulumi.Input<string>;
}
