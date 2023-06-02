// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Serverless App Engine (SAE) Namespace resource.
 *
 * For information about SAE Namespace and how to use it, see [What is Namespace](https://www.alibabacloud.com/help/en/sae/latest/createnamespace).
 *
 * > **NOTE:** Available in v1.129.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.sae.Namespace("example", {
 *     namespaceDescription: "your_description",
 *     namespaceId: "cn-hangzhou:yourname",
 *     namespaceName: "example_value",
 * });
 * ```
 *
 * ## Import
 *
 * Serverless App Engine (SAE) Namespace can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:sae/namespace:Namespace example <namespace_id>
 * ```
 */
export class Namespace extends pulumi.CustomResource {
    /**
     * Get an existing Namespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NamespaceState, opts?: pulumi.CustomResourceOptions): Namespace {
        return new Namespace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:sae/namespace:Namespace';

    /**
     * Returns true if the given object is an instance of Namespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Namespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Namespace.__pulumiType;
    }

    /**
     * Specifies whether to enable the SAE built-in registry. If you do not use the built-in registry, you can set `enableMicroRegistration` to `false` to accelerate the creation of the namespace. Default value: `true`. Valid values:
     */
    public readonly enableMicroRegistration!: pulumi.Output<boolean>;
    /**
     * The Description of Namespace.
     */
    public readonly namespaceDescription!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Namespace. It can contain 2 to 32 lowercase characters. The value is in format `{RegionId}:{namespace}`.
     */
    public readonly namespaceId!: pulumi.Output<string>;
    /**
     * The Name of Namespace.
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * The short ID of the Namespace. You do not need to specify a region ID. The value of `namespaceShortId` can be up to 20 characters in length and can contain only lowercase letters and digits.
     */
    public readonly namespaceShortId!: pulumi.Output<string>;

    /**
     * Create a Namespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NamespaceArgs | NamespaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NamespaceState | undefined;
            resourceInputs["enableMicroRegistration"] = state ? state.enableMicroRegistration : undefined;
            resourceInputs["namespaceDescription"] = state ? state.namespaceDescription : undefined;
            resourceInputs["namespaceId"] = state ? state.namespaceId : undefined;
            resourceInputs["namespaceName"] = state ? state.namespaceName : undefined;
            resourceInputs["namespaceShortId"] = state ? state.namespaceShortId : undefined;
        } else {
            const args = argsOrState as NamespaceArgs | undefined;
            if ((!args || args.namespaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceName'");
            }
            resourceInputs["enableMicroRegistration"] = args ? args.enableMicroRegistration : undefined;
            resourceInputs["namespaceDescription"] = args ? args.namespaceDescription : undefined;
            resourceInputs["namespaceId"] = args ? args.namespaceId : undefined;
            resourceInputs["namespaceName"] = args ? args.namespaceName : undefined;
            resourceInputs["namespaceShortId"] = args ? args.namespaceShortId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Namespace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Namespace resources.
 */
export interface NamespaceState {
    /**
     * Specifies whether to enable the SAE built-in registry. If you do not use the built-in registry, you can set `enableMicroRegistration` to `false` to accelerate the creation of the namespace. Default value: `true`. Valid values:
     */
    enableMicroRegistration?: pulumi.Input<boolean>;
    /**
     * The Description of Namespace.
     */
    namespaceDescription?: pulumi.Input<string>;
    /**
     * The ID of the Namespace. It can contain 2 to 32 lowercase characters. The value is in format `{RegionId}:{namespace}`.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * The Name of Namespace.
     */
    namespaceName?: pulumi.Input<string>;
    /**
     * The short ID of the Namespace. You do not need to specify a region ID. The value of `namespaceShortId` can be up to 20 characters in length and can contain only lowercase letters and digits.
     */
    namespaceShortId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Namespace resource.
 */
export interface NamespaceArgs {
    /**
     * Specifies whether to enable the SAE built-in registry. If you do not use the built-in registry, you can set `enableMicroRegistration` to `false` to accelerate the creation of the namespace. Default value: `true`. Valid values:
     */
    enableMicroRegistration?: pulumi.Input<boolean>;
    /**
     * The Description of Namespace.
     */
    namespaceDescription?: pulumi.Input<string>;
    /**
     * The ID of the Namespace. It can contain 2 to 32 lowercase characters. The value is in format `{RegionId}:{namespace}`.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * The Name of Namespace.
     */
    namespaceName: pulumi.Input<string>;
    /**
     * The short ID of the Namespace. You do not need to specify a region ID. The value of `namespaceShortId` can be up to 20 characters in length and can contain only lowercase letters and digits.
     */
    namespaceShortId?: pulumi.Input<string>;
}
