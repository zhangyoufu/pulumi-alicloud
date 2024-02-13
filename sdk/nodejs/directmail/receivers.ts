// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Direct Mail Receivers resource.
 *
 * For information about Direct Mail Receivers and how to use it, see [What is Direct Mail Receivers](https://www.alibabacloud.com/help/en/doc-detail/29414.htm).
 *
 * > **NOTE:** Available since v1.125.0.
 *
 * ## Import
 *
 * Direct Mail Receivers can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:directmail/receivers:Receivers example <id>
 * ```
 */
export class Receivers extends pulumi.CustomResource {
    /**
     * Get an existing Receivers resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReceiversState, opts?: pulumi.CustomResourceOptions): Receivers {
        return new Receivers(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:directmail/receivers:Receivers';

    /**
     * Returns true if the given object is an instance of Receivers.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Receivers {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Receivers.__pulumiType;
    }

    /**
     * The description of receivers and 1-50 characters in length.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The alias of receivers. Must email address and less than 30 characters in length.
     */
    public readonly receiversAlias!: pulumi.Output<string>;
    /**
     * The name of the resource. The length that cannot be repeated is 1-30 characters.
     */
    public readonly receiversName!: pulumi.Output<string>;
    /**
     * The status of the resource. `0` means uploading, `1` means upload completed.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;

    /**
     * Create a Receivers resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReceiversArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReceiversArgs | ReceiversState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReceiversState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["receiversAlias"] = state ? state.receiversAlias : undefined;
            resourceInputs["receiversName"] = state ? state.receiversName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ReceiversArgs | undefined;
            if ((!args || args.receiversAlias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'receiversAlias'");
            }
            if ((!args || args.receiversName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'receiversName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["receiversAlias"] = args ? args.receiversAlias : undefined;
            resourceInputs["receiversName"] = args ? args.receiversName : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Receivers.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Receivers resources.
 */
export interface ReceiversState {
    /**
     * The description of receivers and 1-50 characters in length.
     */
    description?: pulumi.Input<string>;
    /**
     * The alias of receivers. Must email address and less than 30 characters in length.
     */
    receiversAlias?: pulumi.Input<string>;
    /**
     * The name of the resource. The length that cannot be repeated is 1-30 characters.
     */
    receiversName?: pulumi.Input<string>;
    /**
     * The status of the resource. `0` means uploading, `1` means upload completed.
     */
    status?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Receivers resource.
 */
export interface ReceiversArgs {
    /**
     * The description of receivers and 1-50 characters in length.
     */
    description?: pulumi.Input<string>;
    /**
     * The alias of receivers. Must email address and less than 30 characters in length.
     */
    receiversAlias: pulumi.Input<string>;
    /**
     * The name of the resource. The length that cannot be repeated is 1-30 characters.
     */
    receiversName: pulumi.Input<string>;
}
