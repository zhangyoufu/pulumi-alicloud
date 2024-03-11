// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Api Gateway Backend resource.
 *
 * For information about Api Gateway Backend and how to use it, see [What is Backend](https://www.alibabacloud.com/help/en/api-gateway/developer-reference/api-cloudapi-2016-07-14-createbackend).
 *
 * > **NOTE:** Available since v1.181.0.
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
 * const name = config.get("name") || "tf_example";
 * const _default = new alicloud.apigateway.Backend("default", {
 *     backendName: name,
 *     description: name,
 *     backendType: "HTTP",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Api Gateway Backend can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:apigateway/backend:Backend example <id>
 * ```
 */
export class Backend extends pulumi.CustomResource {
    /**
     * Get an existing Backend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendState, opts?: pulumi.CustomResourceOptions): Backend {
        return new Backend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:apigateway/backend:Backend';

    /**
     * Returns true if the given object is an instance of Backend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Backend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Backend.__pulumiType;
    }

    /**
     * The name of the Backend.
     */
    public readonly backendName!: pulumi.Output<string>;
    /**
     * The type of the Backend. Valid values: `HTTP`, `VPC`, `FC_EVENT`, `FC_HTTP`, `OSS`, `MOCK`.
     */
    public readonly backendType!: pulumi.Output<string>;
    /**
     * Whether to create an Event bus service association role.
     */
    public readonly createEventBridgeServiceLinkedRole!: pulumi.Output<boolean | undefined>;
    /**
     * The description of the Backend.
     */
    public readonly description!: pulumi.Output<string | undefined>;

    /**
     * Create a Backend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackendArgs | BackendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackendState | undefined;
            resourceInputs["backendName"] = state ? state.backendName : undefined;
            resourceInputs["backendType"] = state ? state.backendType : undefined;
            resourceInputs["createEventBridgeServiceLinkedRole"] = state ? state.createEventBridgeServiceLinkedRole : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
        } else {
            const args = argsOrState as BackendArgs | undefined;
            if ((!args || args.backendName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backendName'");
            }
            if ((!args || args.backendType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backendType'");
            }
            resourceInputs["backendName"] = args ? args.backendName : undefined;
            resourceInputs["backendType"] = args ? args.backendType : undefined;
            resourceInputs["createEventBridgeServiceLinkedRole"] = args ? args.createEventBridgeServiceLinkedRole : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Backend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Backend resources.
 */
export interface BackendState {
    /**
     * The name of the Backend.
     */
    backendName?: pulumi.Input<string>;
    /**
     * The type of the Backend. Valid values: `HTTP`, `VPC`, `FC_EVENT`, `FC_HTTP`, `OSS`, `MOCK`.
     */
    backendType?: pulumi.Input<string>;
    /**
     * Whether to create an Event bus service association role.
     */
    createEventBridgeServiceLinkedRole?: pulumi.Input<boolean>;
    /**
     * The description of the Backend.
     */
    description?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Backend resource.
 */
export interface BackendArgs {
    /**
     * The name of the Backend.
     */
    backendName: pulumi.Input<string>;
    /**
     * The type of the Backend. Valid values: `HTTP`, `VPC`, `FC_EVENT`, `FC_HTTP`, `OSS`, `MOCK`.
     */
    backendType: pulumi.Input<string>;
    /**
     * Whether to create an Event bus service association role.
     */
    createEventBridgeServiceLinkedRole?: pulumi.Input<boolean>;
    /**
     * The description of the Backend.
     */
    description?: pulumi.Input<string>;
}
