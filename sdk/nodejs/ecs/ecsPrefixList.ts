// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a ECS Prefix List resource.
 *
 * For information about ECS Prefix List and how to use it, see [What is Prefix List.](https://www.alibabacloud.com/help/en/doc-detail/207969.html).
 *
 * > **NOTE:** Available in v1.152.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.ecs.EcsPrefixList("default", {
 *     addressFamily: "IPv4",
 *     description: "description",
 *     entries: [{
 *         cidr: "192.168.0.0/24",
 *         description: "description",
 *     }],
 *     maxEntries: 2,
 *     prefixListName: "tftest",
 * });
 * ```
 *
 * ## Import
 *
 * ECS Prefix List can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ecs/ecsPrefixList:EcsPrefixList example <id>
 * ```
 */
export class EcsPrefixList extends pulumi.CustomResource {
    /**
     * Get an existing EcsPrefixList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EcsPrefixListState, opts?: pulumi.CustomResourceOptions): EcsPrefixList {
        return new EcsPrefixList(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/ecsPrefixList:EcsPrefixList';

    /**
     * Returns true if the given object is an instance of EcsPrefixList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EcsPrefixList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EcsPrefixList.__pulumiType;
    }

    /**
     * The IP address family. Valid values: `IPv4`,`IPv6`.
     */
    public readonly addressFamily!: pulumi.Output<string>;
    /**
     * The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Entry. The details see Block `entry`.
     */
    public readonly entries!: pulumi.Output<outputs.ecs.EcsPrefixListEntry[]>;
    /**
     * The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
     */
    public readonly maxEntries!: pulumi.Output<number>;
    /**
     * The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
     */
    public readonly prefixListName!: pulumi.Output<string>;

    /**
     * Create a EcsPrefixList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EcsPrefixListArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EcsPrefixListArgs | EcsPrefixListState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EcsPrefixListState | undefined;
            resourceInputs["addressFamily"] = state ? state.addressFamily : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["entries"] = state ? state.entries : undefined;
            resourceInputs["maxEntries"] = state ? state.maxEntries : undefined;
            resourceInputs["prefixListName"] = state ? state.prefixListName : undefined;
        } else {
            const args = argsOrState as EcsPrefixListArgs | undefined;
            if ((!args || args.addressFamily === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if ((!args || args.entries === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entries'");
            }
            if ((!args || args.maxEntries === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxEntries'");
            }
            if ((!args || args.prefixListName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'prefixListName'");
            }
            resourceInputs["addressFamily"] = args ? args.addressFamily : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["entries"] = args ? args.entries : undefined;
            resourceInputs["maxEntries"] = args ? args.maxEntries : undefined;
            resourceInputs["prefixListName"] = args ? args.prefixListName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EcsPrefixList.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EcsPrefixList resources.
 */
export interface EcsPrefixListState {
    /**
     * The IP address family. Valid values: `IPv4`,`IPv6`.
     */
    addressFamily?: pulumi.Input<string>;
    /**
     * The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * The Entry. The details see Block `entry`.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.ecs.EcsPrefixListEntry>[]>;
    /**
     * The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
     */
    maxEntries?: pulumi.Input<number>;
    /**
     * The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
     */
    prefixListName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EcsPrefixList resource.
 */
export interface EcsPrefixListArgs {
    /**
     * The IP address family. Valid values: `IPv4`,`IPv6`.
     */
    addressFamily: pulumi.Input<string>;
    /**
     * The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * The Entry. The details see Block `entry`.
     */
    entries: pulumi.Input<pulumi.Input<inputs.ecs.EcsPrefixListEntry>[]>;
    /**
     * The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
     */
    maxEntries: pulumi.Input<number>;
    /**
     * The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
     */
    prefixListName: pulumi.Input<string>;
}
