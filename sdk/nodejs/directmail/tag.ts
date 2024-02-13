// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Direct Mail Tag resource.
 *
 * For information about Direct Mail Tag and how to use it, see [What is Tag](https://www.alibabacloud.com/help/en/directmail/latest/createtag).
 *
 * > **NOTE:** Available since v1.144.0.
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
 * const name = config.get("name") || "example";
 * const example = new alicloud.directmail.Tag("example", {tagName: name});
 * ```
 *
 * ## Import
 *
 * Direct Mail Tag can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:directmail/tag:Tag example <id>
 * ```
 */
export class Tag extends pulumi.CustomResource {
    /**
     * Get an existing Tag resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagState, opts?: pulumi.CustomResourceOptions): Tag {
        return new Tag(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:directmail/tag:Tag';

    /**
     * Returns true if the given object is an instance of Tag.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Tag {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Tag.__pulumiType;
    }

    /**
     * The name of the tag. The name must be `1` to `50` characters in length, and can contain letters and digits.
     */
    public readonly tagName!: pulumi.Output<string>;

    /**
     * Create a Tag resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagArgs | TagState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TagState | undefined;
            resourceInputs["tagName"] = state ? state.tagName : undefined;
        } else {
            const args = argsOrState as TagArgs | undefined;
            if ((!args || args.tagName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tagName'");
            }
            resourceInputs["tagName"] = args ? args.tagName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Tag.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Tag resources.
 */
export interface TagState {
    /**
     * The name of the tag. The name must be `1` to `50` characters in length, and can contain letters and digits.
     */
    tagName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Tag resource.
 */
export interface TagArgs {
    /**
     * The name of the tag. The name must be `1` to `50` characters in length, and can contain letters and digits.
     */
    tagName: pulumi.Input<string>;
}
