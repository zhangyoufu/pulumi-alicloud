// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Resource Manager Shared Target resource.
 *
 * For information about Resource Manager Shared Target and how to use it, see [What is Shared Target](https://www.alibabacloud.com/help/en/doc-detail/94475.htm).
 *
 * > **NOTE:** Available in v1.111.0+.
 *
 * ## Import
 *
 * Resource Manager Shared Target can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:resourcemanager/sharedTarget:SharedTarget example <resource_share_id>:<target_id>
 * ```
 */
export class SharedTarget extends pulumi.CustomResource {
    /**
     * Get an existing SharedTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SharedTargetState, opts?: pulumi.CustomResourceOptions): SharedTarget {
        return new SharedTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:resourcemanager/sharedTarget:SharedTarget';

    /**
     * Returns true if the given object is an instance of SharedTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SharedTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SharedTarget.__pulumiType;
    }

    /**
     * The resource share ID of resource manager.
     */
    public readonly resourceShareId!: pulumi.Output<string>;
    /**
     * The status of shared target.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The member account ID in resource directory.
     */
    public readonly targetId!: pulumi.Output<string>;

    /**
     * Create a SharedTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SharedTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SharedTargetArgs | SharedTargetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SharedTargetState | undefined;
            inputs["resourceShareId"] = state ? state.resourceShareId : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["targetId"] = state ? state.targetId : undefined;
        } else {
            const args = argsOrState as SharedTargetArgs | undefined;
            if (!args || args.resourceShareId === undefined) {
                throw new Error("Missing required property 'resourceShareId'");
            }
            if (!args || args.targetId === undefined) {
                throw new Error("Missing required property 'targetId'");
            }
            inputs["resourceShareId"] = args ? args.resourceShareId : undefined;
            inputs["targetId"] = args ? args.targetId : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SharedTarget.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SharedTarget resources.
 */
export interface SharedTargetState {
    /**
     * The resource share ID of resource manager.
     */
    readonly resourceShareId?: pulumi.Input<string>;
    /**
     * The status of shared target.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The member account ID in resource directory.
     */
    readonly targetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SharedTarget resource.
 */
export interface SharedTargetArgs {
    /**
     * The resource share ID of resource manager.
     */
    readonly resourceShareId: pulumi.Input<string>;
    /**
     * The member account ID in resource directory.
     */
    readonly targetId: pulumi.Input<string>;
}
