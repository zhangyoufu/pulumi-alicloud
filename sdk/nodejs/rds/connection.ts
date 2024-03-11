// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an RDS connection resource to allocate an Internet connection string for RDS instance, see [What is DB Connection](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/api-rds-2014-08-15-allocateinstancepublicconnection).
 *
 * > **NOTE:** Each RDS instance will allocate a intranet connnection string automatically and its prifix is RDS instance ID.
 *  To avoid unnecessary conflict, please specified a internet connection prefix before applying the resource.
 *
 * > **NOTE:** Available since v1.5.0.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf_example";
 * const defaultZones = alicloud.rds.getZones({
 *     engine: "MySQL",
 *     engineVersion: "5.6",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     vswitchName: name,
 * });
 * const defaultInstance = new alicloud.rds.Instance("defaultInstance", {
 *     engine: "MySQL",
 *     engineVersion: "5.6",
 *     instanceType: "rds.mysql.t1.small",
 *     instanceStorage: 10,
 *     vswitchId: defaultSwitch.id,
 *     instanceName: name,
 * });
 * const defaultConnection = new alicloud.rds.Connection("defaultConnection", {
 *     instanceId: defaultInstance.id,
 *     connectionPrefix: "testabc",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * RDS connection can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:rds/connection:Connection example abc12345678
 * ```
 */
export class Connection extends pulumi.CustomResource {
    /**
     * Get an existing Connection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectionState, opts?: pulumi.CustomResourceOptions): Connection {
        return new Connection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:rds/connection:Connection';

    /**
     * Returns true if the given object is an instance of Connection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connection.__pulumiType;
    }

    /**
     * The Tabular Data Stream (TDS) port of the instance for which Babelfish is enabled.
     *
     * > **NOTE:** This parameter applies only to ApsaraDB RDS for PostgreSQL instances. For more information about Babelfish for ApsaraDB RDS for PostgreSQL, see [Introduction to Babelfish](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/babelfish-for-pg).
     */
    public readonly babelfishPort!: pulumi.Output<string>;
    /**
     * Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 40 characters. Default to <instance_id> + 'tf'.
     */
    public readonly connectionPrefix!: pulumi.Output<string>;
    /**
     * Connection instance string.
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * The Id of instance that can run database.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The ip address of connection string.
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * Internet connection port. Valid value: [1000-5999]. Default to 3306.
     */
    public readonly port!: pulumi.Output<string | undefined>;

    /**
     * Create a Connection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionArgs | ConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConnectionState | undefined;
            resourceInputs["babelfishPort"] = state ? state.babelfishPort : undefined;
            resourceInputs["connectionPrefix"] = state ? state.connectionPrefix : undefined;
            resourceInputs["connectionString"] = state ? state.connectionString : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
        } else {
            const args = argsOrState as ConnectionArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["babelfishPort"] = args ? args.babelfishPort : undefined;
            resourceInputs["connectionPrefix"] = args ? args.connectionPrefix : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["connectionString"] = undefined /*out*/;
            resourceInputs["ipAddress"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Connection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Connection resources.
 */
export interface ConnectionState {
    /**
     * The Tabular Data Stream (TDS) port of the instance for which Babelfish is enabled.
     *
     * > **NOTE:** This parameter applies only to ApsaraDB RDS for PostgreSQL instances. For more information about Babelfish for ApsaraDB RDS for PostgreSQL, see [Introduction to Babelfish](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/babelfish-for-pg).
     */
    babelfishPort?: pulumi.Input<string>;
    /**
     * Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 40 characters. Default to <instance_id> + 'tf'.
     */
    connectionPrefix?: pulumi.Input<string>;
    /**
     * Connection instance string.
     */
    connectionString?: pulumi.Input<string>;
    /**
     * The Id of instance that can run database.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The ip address of connection string.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * Internet connection port. Valid value: [1000-5999]. Default to 3306.
     */
    port?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Connection resource.
 */
export interface ConnectionArgs {
    /**
     * The Tabular Data Stream (TDS) port of the instance for which Babelfish is enabled.
     *
     * > **NOTE:** This parameter applies only to ApsaraDB RDS for PostgreSQL instances. For more information about Babelfish for ApsaraDB RDS for PostgreSQL, see [Introduction to Babelfish](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/babelfish-for-pg).
     */
    babelfishPort?: pulumi.Input<string>;
    /**
     * Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 40 characters. Default to <instance_id> + 'tf'.
     */
    connectionPrefix?: pulumi.Input<string>;
    /**
     * The Id of instance that can run database.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Internet connection port. Valid value: [1000-5999]. Default to 3306.
     */
    port?: pulumi.Input<string>;
}
