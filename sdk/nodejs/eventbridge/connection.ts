// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Event Bridge Connection resource.
 *
 * For information about Event Bridge Connection and how to use it, see [What is Connection](https://www.alibabacloud.com/help/en/eventbridge/latest/api-eventbridge-2020-04-01-createconnection).
 *
 * > **NOTE:** Available since v1.210.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const region = config.get("region") || "cn-chengdu";
 * const name = config.get("name") || "terraform-example";
 * const default = alicloud.getZones({});
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     vpcName: name,
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: _default.then(_default => _default.zones?.[0]?.id),
 *     vswitchName: name,
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("default", {
 *     name: name,
 *     vpcId: defaultSwitch.vpcId,
 * });
 * const defaultConnection = new alicloud.eventbridge.Connection("default", {
 *     connectionName: name,
 *     description: "test-connection-basic-pre",
 *     networkParameters: {
 *         networkType: "PublicNetwork",
 *         vpcId: defaultNetwork.id,
 *         vswitcheId: defaultSwitch.id,
 *         securityGroupId: defaultSecurityGroup.id,
 *     },
 *     authParameters: {
 *         authorizationType: "BASIC_AUTH",
 *         apiKeyAuthParameters: {
 *             apiKeyName: "Token",
 *             apiKeyValue: "Token-value",
 *         },
 *         basicAuthParameters: {
 *             username: "admin",
 *             password: "admin",
 *         },
 *         oauthParameters: {
 *             authorizationEndpoint: "http://127.0.0.1:8080",
 *             httpMethod: "POST",
 *             clientParameters: {
 *                 clientId: "ClientId",
 *                 clientSecret: "ClientSecret",
 *             },
 *             oauthHttpParameters: {
 *                 headerParameters: [{
 *                     key: "name",
 *                     value: "name",
 *                     isValueSecret: "true",
 *                 }],
 *                 bodyParameters: [{
 *                     key: "name",
 *                     value: "name",
 *                     isValueSecret: "true",
 *                 }],
 *                 queryStringParameters: [{
 *                     key: "name",
 *                     value: "name",
 *                     isValueSecret: "true",
 *                 }],
 *             },
 *         },
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Event Bridge Connection can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:eventbridge/connection:Connection example <id>
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
    public static readonly __pulumiType = 'alicloud:eventbridge/connection:Connection';

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
     * The parameters that are configured for authentication. See `authParameters` below.
     */
    public readonly authParameters!: pulumi.Output<outputs.eventbridge.ConnectionAuthParameters | undefined>;
    /**
     * The name of the connection.
     */
    public readonly connectionName!: pulumi.Output<string>;
    /**
     * The creation time of the Connection.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The description of the connection.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The parameters that are configured for the network. See `networkParameters` below.
     */
    public readonly networkParameters!: pulumi.Output<outputs.eventbridge.ConnectionNetworkParameters>;

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
            resourceInputs["authParameters"] = state ? state.authParameters : undefined;
            resourceInputs["connectionName"] = state ? state.connectionName : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["networkParameters"] = state ? state.networkParameters : undefined;
        } else {
            const args = argsOrState as ConnectionArgs | undefined;
            if ((!args || args.connectionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionName'");
            }
            if ((!args || args.networkParameters === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkParameters'");
            }
            resourceInputs["authParameters"] = args ? args.authParameters : undefined;
            resourceInputs["connectionName"] = args ? args.connectionName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["networkParameters"] = args ? args.networkParameters : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
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
     * The parameters that are configured for authentication. See `authParameters` below.
     */
    authParameters?: pulumi.Input<inputs.eventbridge.ConnectionAuthParameters>;
    /**
     * The name of the connection.
     */
    connectionName?: pulumi.Input<string>;
    /**
     * The creation time of the Connection.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The description of the connection.
     */
    description?: pulumi.Input<string>;
    /**
     * The parameters that are configured for the network. See `networkParameters` below.
     */
    networkParameters?: pulumi.Input<inputs.eventbridge.ConnectionNetworkParameters>;
}

/**
 * The set of arguments for constructing a Connection resource.
 */
export interface ConnectionArgs {
    /**
     * The parameters that are configured for authentication. See `authParameters` below.
     */
    authParameters?: pulumi.Input<inputs.eventbridge.ConnectionAuthParameters>;
    /**
     * The name of the connection.
     */
    connectionName: pulumi.Input<string>;
    /**
     * The description of the connection.
     */
    description?: pulumi.Input<string>;
    /**
     * The parameters that are configured for the network. See `networkParameters` below.
     */
    networkParameters: pulumi.Input<inputs.eventbridge.ConnectionNetworkParameters>;
}
