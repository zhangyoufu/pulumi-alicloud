// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an RDS read write splitting connection resource to allocate an Intranet connection string for RDS instance.
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
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: creation,
 * });
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
 * const defaultReadWriteSplittingConnection = new alicloud.rds.ReadWriteSplittingConnection("default", {
 *     connectionPrefix: "t-con-123",
 *     distributionType: "Standard",
 *     instanceId: defaultInstance.id,
 * }, {dependsOn: [defaultReadOnlyInstance]});
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/db_read_write_splitting_connection.html.markdown.
 */
export class ReadWriteSplittingConnection extends pulumi.CustomResource {
    /**
     * Get an existing ReadWriteSplittingConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReadWriteSplittingConnectionState, opts?: pulumi.CustomResourceOptions): ReadWriteSplittingConnection {
        return new ReadWriteSplittingConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:rds/readWriteSplittingConnection:ReadWriteSplittingConnection';

    /**
     * Returns true if the given object is an instance of ReadWriteSplittingConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReadWriteSplittingConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReadWriteSplittingConnection.__pulumiType;
    }

    /**
     * Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
     */
    public readonly connectionPrefix!: pulumi.Output<string | undefined>;
    /**
     * Connection instance string.
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution. 
     */
    public readonly distributionType!: pulumi.Output<string>;
    /**
     * The Id of instance that can run database.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.  
     */
    public readonly maxDelayTime!: pulumi.Output<number>;
    /**
     * Intranet connection port. Valid value: [3001-3999]. Default to 3306.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom. 
     */
    public readonly weight!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a ReadWriteSplittingConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReadWriteSplittingConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReadWriteSplittingConnectionArgs | ReadWriteSplittingConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ReadWriteSplittingConnectionState | undefined;
            inputs["connectionPrefix"] = state ? state.connectionPrefix : undefined;
            inputs["connectionString"] = state ? state.connectionString : undefined;
            inputs["distributionType"] = state ? state.distributionType : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["maxDelayTime"] = state ? state.maxDelayTime : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as ReadWriteSplittingConnectionArgs | undefined;
            if (!args || args.distributionType === undefined) {
                throw new Error("Missing required property 'distributionType'");
            }
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            inputs["connectionPrefix"] = args ? args.connectionPrefix : undefined;
            inputs["distributionType"] = args ? args.distributionType : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["maxDelayTime"] = args ? args.maxDelayTime : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["weight"] = args ? args.weight : undefined;
            inputs["connectionString"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ReadWriteSplittingConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReadWriteSplittingConnection resources.
 */
export interface ReadWriteSplittingConnectionState {
    /**
     * Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
     */
    readonly connectionPrefix?: pulumi.Input<string>;
    /**
     * Connection instance string.
     */
    readonly connectionString?: pulumi.Input<string>;
    /**
     * Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution. 
     */
    readonly distributionType?: pulumi.Input<string>;
    /**
     * The Id of instance that can run database.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.  
     */
    readonly maxDelayTime?: pulumi.Input<number>;
    /**
     * Intranet connection port. Valid value: [3001-3999]. Default to 3306.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom. 
     */
    readonly weight?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a ReadWriteSplittingConnection resource.
 */
export interface ReadWriteSplittingConnectionArgs {
    /**
     * Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
     */
    readonly connectionPrefix?: pulumi.Input<string>;
    /**
     * Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution. 
     */
    readonly distributionType: pulumi.Input<string>;
    /**
     * The Id of instance that can run database.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.  
     */
    readonly maxDelayTime?: pulumi.Input<number>;
    /**
     * Intranet connection port. Valid value: [3001-3999]. Default to 3306.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom. 
     */
    readonly weight?: pulumi.Input<{[key: string]: any}>;
}
