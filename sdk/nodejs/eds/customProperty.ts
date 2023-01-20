// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a ECD Custom Property resource.
 *
 * For information about ECD Custom Property and how to use it, see [What is Custom Property](https://help.aliyun.com/document_detail/436381.html).
 *
 * > **NOTE:** Available in v1.176.0+.
 *
 * > **NOTE:** Up to 10 different attributes can be created under an alibaba cloud account. Up to 50 different attribute values can be added under an attribute.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.eds.CustomProperty("example", {
 *     propertyKey: "example_key",
 *     propertyValues: [{
 *         propertyValue: "example_value",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * ECD Custom Property can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:eds/customProperty:CustomProperty example <id>
 * ```
 */
export class CustomProperty extends pulumi.CustomResource {
    /**
     * Get an existing CustomProperty resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomPropertyState, opts?: pulumi.CustomResourceOptions): CustomProperty {
        return new CustomProperty(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eds/customProperty:CustomProperty';

    /**
     * Returns true if the given object is an instance of CustomProperty.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomProperty {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomProperty.__pulumiType;
    }

    /**
     * The Custom attribute key.
     */
    public readonly propertyKey!: pulumi.Output<string>;
    /**
     * Custom attribute sets the value of. See the following `Block propertyValues`.
     */
    public readonly propertyValues!: pulumi.Output<outputs.eds.CustomPropertyPropertyValue[] | undefined>;

    /**
     * Create a CustomProperty resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomPropertyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomPropertyArgs | CustomPropertyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomPropertyState | undefined;
            resourceInputs["propertyKey"] = state ? state.propertyKey : undefined;
            resourceInputs["propertyValues"] = state ? state.propertyValues : undefined;
        } else {
            const args = argsOrState as CustomPropertyArgs | undefined;
            if ((!args || args.propertyKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'propertyKey'");
            }
            resourceInputs["propertyKey"] = args ? args.propertyKey : undefined;
            resourceInputs["propertyValues"] = args ? args.propertyValues : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomProperty.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomProperty resources.
 */
export interface CustomPropertyState {
    /**
     * The Custom attribute key.
     */
    propertyKey?: pulumi.Input<string>;
    /**
     * Custom attribute sets the value of. See the following `Block propertyValues`.
     */
    propertyValues?: pulumi.Input<pulumi.Input<inputs.eds.CustomPropertyPropertyValue>[]>;
}

/**
 * The set of arguments for constructing a CustomProperty resource.
 */
export interface CustomPropertyArgs {
    /**
     * The Custom attribute key.
     */
    propertyKey: pulumi.Input<string>;
    /**
     * Custom attribute sets the value of. See the following `Block propertyValues`.
     */
    propertyValues?: pulumi.Input<pulumi.Input<inputs.eds.CustomPropertyPropertyValue>[]>;
}
