// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Event Bridge Api Destination resource.
 *
 * For information about Event Bridge Api Destination and how to use it, see [What is Api Destination](https://www.alibabacloud.com/help/en/eventbridge/latest/api-eventbridge-2020-04-01-createapidestination).
 *
 * > **NOTE:** Available since v1.211.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const region = config.get("region") || "cn-chengdu";
 * const name = config.get("name") || "terraform-example";
 * const defaultConnection = new alicloud.eventbridge.Connection("defaultConnection", {
 *     connectionName: name,
 *     networkParameters: {
 *         networkType: "PublicNetwork",
 *     },
 * });
 * const defaultApiDestination = new alicloud.eventbridge.ApiDestination("defaultApiDestination", {
 *     connectionName: defaultConnection.connectionName,
 *     apiDestinationName: name,
 *     description: "test-api-destination-connection",
 *     httpApiParameters: {
 *         endpoint: "http://127.0.0.1:8001",
 *         method: "POST",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Event Bridge Api Destination can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:eventbridge/apiDestination:ApiDestination example <id>
 * ```
 */
export class ApiDestination extends pulumi.CustomResource {
    /**
     * Get an existing ApiDestination resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiDestinationState, opts?: pulumi.CustomResourceOptions): ApiDestination {
        return new ApiDestination(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eventbridge/apiDestination:ApiDestination';

    /**
     * Returns true if the given object is an instance of ApiDestination.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiDestination {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiDestination.__pulumiType;
    }

    /**
     * The name of the API destination.
     */
    public readonly apiDestinationName!: pulumi.Output<string>;
    /**
     * The name of the connection.
     */
    public readonly connectionName!: pulumi.Output<string>;
    /**
     * The creation time of the Api Destination.
     */
    public /*out*/ readonly createTime!: pulumi.Output<number>;
    /**
     * The description of the API destination.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The parameters that are configured for the API destination. See `httpApiParameters` below.
     */
    public readonly httpApiParameters!: pulumi.Output<outputs.eventbridge.ApiDestinationHttpApiParameters>;

    /**
     * Create a ApiDestination resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiDestinationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiDestinationArgs | ApiDestinationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiDestinationState | undefined;
            resourceInputs["apiDestinationName"] = state ? state.apiDestinationName : undefined;
            resourceInputs["connectionName"] = state ? state.connectionName : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["httpApiParameters"] = state ? state.httpApiParameters : undefined;
        } else {
            const args = argsOrState as ApiDestinationArgs | undefined;
            if ((!args || args.apiDestinationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiDestinationName'");
            }
            if ((!args || args.connectionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionName'");
            }
            if ((!args || args.httpApiParameters === undefined) && !opts.urn) {
                throw new Error("Missing required property 'httpApiParameters'");
            }
            resourceInputs["apiDestinationName"] = args ? args.apiDestinationName : undefined;
            resourceInputs["connectionName"] = args ? args.connectionName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["httpApiParameters"] = args ? args.httpApiParameters : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApiDestination.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiDestination resources.
 */
export interface ApiDestinationState {
    /**
     * The name of the API destination.
     */
    apiDestinationName?: pulumi.Input<string>;
    /**
     * The name of the connection.
     */
    connectionName?: pulumi.Input<string>;
    /**
     * The creation time of the Api Destination.
     */
    createTime?: pulumi.Input<number>;
    /**
     * The description of the API destination.
     */
    description?: pulumi.Input<string>;
    /**
     * The parameters that are configured for the API destination. See `httpApiParameters` below.
     */
    httpApiParameters?: pulumi.Input<inputs.eventbridge.ApiDestinationHttpApiParameters>;
}

/**
 * The set of arguments for constructing a ApiDestination resource.
 */
export interface ApiDestinationArgs {
    /**
     * The name of the API destination.
     */
    apiDestinationName: pulumi.Input<string>;
    /**
     * The name of the connection.
     */
    connectionName: pulumi.Input<string>;
    /**
     * The description of the API destination.
     */
    description?: pulumi.Input<string>;
    /**
     * The parameters that are configured for the API destination. See `httpApiParameters` below.
     */
    httpApiParameters: pulumi.Input<inputs.eventbridge.ApiDestinationHttpApiParameters>;
}
