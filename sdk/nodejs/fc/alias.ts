// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a Function Compute service alias. Creates an alias that points to the specified Function Compute service version.
 *  For the detailed information, please refer to the [developer guide](https://www.alibabacloud.com/help/doc-detail/171635.htm).
 *
 * > **NOTE:** Available in 1.104.0+
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.fc.Alias("example", {
 *     aliasName: "my_alias",
 *     description: "a sample description",
 *     routingConfig: {
 *         additionalVersionWeights: {
 *             "2": 0.5,
 *         },
 *     },
 *     serviceName: "my_service_name",
 *     serviceVersion: "1",
 * });
 * ```
 *
 * ## Import
 *
 * Function Compute alias can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:fc/alias:Alias example my_alias_id
 * ```
 */
export class Alias extends pulumi.CustomResource {
    /**
     * Get an existing Alias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AliasState, opts?: pulumi.CustomResourceOptions): Alias {
        return new Alias(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:fc/alias:Alias';

    /**
     * Returns true if the given object is an instance of Alias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alias.__pulumiType;
    }

    /**
     * Name for the alias you are creating.
     */
    public readonly aliasName!: pulumi.Output<string>;
    /**
     * Description of the alias.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Function Compute alias' route configuration settings. Fields documented below.
     *
     * **routing_config** includes the following arguments:
     */
    public readonly routingConfig!: pulumi.Output<outputs.fc.AliasRoutingConfig | undefined>;
    /**
     * The Function Compute service name.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * The Function Compute service version for which you are creating the alias. Pattern: (LATEST|[0-9]+).
     */
    public readonly serviceVersion!: pulumi.Output<string>;

    /**
     * Create a Alias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AliasArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AliasArgs | AliasState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AliasState | undefined;
            resourceInputs["aliasName"] = state ? state.aliasName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["routingConfig"] = state ? state.routingConfig : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["serviceVersion"] = state ? state.serviceVersion : undefined;
        } else {
            const args = argsOrState as AliasArgs | undefined;
            if ((!args || args.aliasName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aliasName'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.serviceVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceVersion'");
            }
            resourceInputs["aliasName"] = args ? args.aliasName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["routingConfig"] = args ? args.routingConfig : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["serviceVersion"] = args ? args.serviceVersion : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Alias.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Alias resources.
 */
export interface AliasState {
    /**
     * Name for the alias you are creating.
     */
    aliasName?: pulumi.Input<string>;
    /**
     * Description of the alias.
     */
    description?: pulumi.Input<string>;
    /**
     * The Function Compute alias' route configuration settings. Fields documented below.
     *
     * **routing_config** includes the following arguments:
     */
    routingConfig?: pulumi.Input<inputs.fc.AliasRoutingConfig>;
    /**
     * The Function Compute service name.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * The Function Compute service version for which you are creating the alias. Pattern: (LATEST|[0-9]+).
     */
    serviceVersion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Alias resource.
 */
export interface AliasArgs {
    /**
     * Name for the alias you are creating.
     */
    aliasName: pulumi.Input<string>;
    /**
     * Description of the alias.
     */
    description?: pulumi.Input<string>;
    /**
     * The Function Compute alias' route configuration settings. Fields documented below.
     *
     * **routing_config** includes the following arguments:
     */
    routingConfig?: pulumi.Input<inputs.fc.AliasRoutingConfig>;
    /**
     * The Function Compute service name.
     */
    serviceName: pulumi.Input<string>;
    /**
     * The Function Compute service version for which you are creating the alias. Pattern: (LATEST|[0-9]+).
     */
    serviceVersion: pulumi.Input<string>;
}
