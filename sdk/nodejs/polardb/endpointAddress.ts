// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a PolarDB endpoint address resource to allocate an Internet endpoint address string for PolarDB instance.
 *
 * > **NOTE:** Available in v1.68.0+. Each PolarDB instance will allocate a intranet connection string automatically and its prefix is Cluster ID.
 *  To avoid unnecessary conflict, please specified a internet connection prefix before applying the resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const creation = config.get("creation") || "PolarDB";
 * const name = config.get("name") || "polardbconnectionbasic";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: creation,
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {cidrBlock: "172.16.0.0/16"});
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     availabilityZone: defaultZones.then(defaultZones => defaultZones.zones[0].id),
 * });
 * const defaultCluster = new alicloud.polardb.Cluster("defaultCluster", {
 *     dbType: "MySQL",
 *     dbVersion: "8.0",
 *     payType: "PostPaid",
 *     dbNodeClass: "polar.mysql.x4.large",
 *     vswitchId: defaultSwitch.id,
 *     description: name,
 * });
 * const defaultEndpoints = defaultCluster.id.apply(id => alicloud.polardb.getEndpoints({
 *     dbClusterId: id,
 * }));
 * const endpoint = new alicloud.polardb.EndpointAddress("endpoint", {
 *     dbClusterId: defaultCluster.id,
 *     dbEndpointId: defaultEndpoints.endpoints[0].dbEndpointId,
 *     connectionPrefix: "testpolardbconn",
 *     netType: "Public",
 * });
 * ```
 *
 * ## Import
 *
 * PolarDB endpoint address can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:polardb/endpointAddress:EndpointAddress example pc-abc123456:pe-abc123456
 * ```
 */
export class EndpointAddress extends pulumi.CustomResource {
    /**
     * Get an existing EndpointAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointAddressState, opts?: pulumi.CustomResourceOptions): EndpointAddress {
        return new EndpointAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:polardb/endpointAddress:EndpointAddress';

    /**
     * Returns true if the given object is an instance of EndpointAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointAddress.__pulumiType;
    }

    /**
     * Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
     */
    public readonly connectionPrefix!: pulumi.Output<string>;
    /**
     * Connection cluster or endpoint string.
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * The Id of cluster that can run database.
     */
    public readonly dbClusterId!: pulumi.Output<string>;
    /**
     * The Id of endpoint that can run database.
     */
    public readonly dbEndpointId!: pulumi.Output<string>;
    /**
     * The ip address of connection string.
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
     */
    public readonly netType!: pulumi.Output<string | undefined>;
    /**
     * Connection cluster or endpoint port.
     */
    public /*out*/ readonly port!: pulumi.Output<string>;

    /**
     * Create a EndpointAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointAddressArgs | EndpointAddressState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointAddressState | undefined;
            inputs["connectionPrefix"] = state ? state.connectionPrefix : undefined;
            inputs["connectionString"] = state ? state.connectionString : undefined;
            inputs["dbClusterId"] = state ? state.dbClusterId : undefined;
            inputs["dbEndpointId"] = state ? state.dbEndpointId : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["netType"] = state ? state.netType : undefined;
            inputs["port"] = state ? state.port : undefined;
        } else {
            const args = argsOrState as EndpointAddressArgs | undefined;
            if ((!args || args.dbClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbClusterId'");
            }
            if ((!args || args.dbEndpointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbEndpointId'");
            }
            inputs["connectionPrefix"] = args ? args.connectionPrefix : undefined;
            inputs["dbClusterId"] = args ? args.dbClusterId : undefined;
            inputs["dbEndpointId"] = args ? args.dbEndpointId : undefined;
            inputs["netType"] = args ? args.netType : undefined;
            inputs["connectionString"] = undefined /*out*/;
            inputs["ipAddress"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EndpointAddress.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndpointAddress resources.
 */
export interface EndpointAddressState {
    /**
     * Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
     */
    readonly connectionPrefix?: pulumi.Input<string>;
    /**
     * Connection cluster or endpoint string.
     */
    readonly connectionString?: pulumi.Input<string>;
    /**
     * The Id of cluster that can run database.
     */
    readonly dbClusterId?: pulumi.Input<string>;
    /**
     * The Id of endpoint that can run database.
     */
    readonly dbEndpointId?: pulumi.Input<string>;
    /**
     * The ip address of connection string.
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
     */
    readonly netType?: pulumi.Input<string>;
    /**
     * Connection cluster or endpoint port.
     */
    readonly port?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EndpointAddress resource.
 */
export interface EndpointAddressArgs {
    /**
     * Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
     */
    readonly connectionPrefix?: pulumi.Input<string>;
    /**
     * The Id of cluster that can run database.
     */
    readonly dbClusterId: pulumi.Input<string>;
    /**
     * The Id of endpoint that can run database.
     */
    readonly dbEndpointId: pulumi.Input<string>;
    /**
     * Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
     */
    readonly netType?: pulumi.Input<string>;
}
