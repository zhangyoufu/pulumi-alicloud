// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Nas Access Group resource.
 *
 * In NAS, the permission group acts as a whitelist that allows you to restrict file system access. You can allow specified IP addresses or CIDR blocks to access the file system, and assign different levels of access permission to different IP addresses or CIDR blocks by adding rules to the permission group.
 *
 * > **NOTE:** Available in v1.33.0+.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const foo = new alicloud.nas.AccessGroup("foo", {
 *     description: "test_AccessG",
 *     type: "Classic",
 * });
 * ```
 */
export class AccessGroup extends pulumi.CustomResource {
    /**
     * Get an existing AccessGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessGroupState, opts?: pulumi.CustomResourceOptions): AccessGroup {
        return new AccessGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:nas/accessGroup:AccessGroup';

    /**
     * Returns true if the given object is an instance of AccessGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessGroup.__pulumiType;
    }

    /**
     * The Access Group description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A Name of one Access Group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A Type of one Access Group. Valid values: `Vpc` and `Classic`.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a AccessGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessGroupArgs | AccessGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AccessGroupState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as AccessGroupArgs | undefined;
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AccessGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessGroup resources.
 */
export interface AccessGroupState {
    /**
     * The Access Group description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A Name of one Access Group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A Type of one Access Group. Valid values: `Vpc` and `Classic`.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessGroup resource.
 */
export interface AccessGroupArgs {
    /**
     * The Access Group description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A Name of one Access Group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A Type of one Access Group. Valid values: `Vpc` and `Classic`.
     */
    readonly type: pulumi.Input<string>;
}
