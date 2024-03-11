// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Elastic Cloud Phone (ECP) Key Pair resource.
 *
 * For information about Elastic Cloud Phone (ECP) Key Pair and how to use it, see [What is Key Pair](https://help.aliyun.com/document_detail/257197.html).
 *
 * > **NOTE:** Available in v1.130.0+.
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
 * const example = new alicloud.ecp.KeyPair("example", {
 *     keyPairName: "my-KeyPair",
 *     publicKeyBody: "ssh-rsa AAAAxxxxxxxxxxtyuudsfsg",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Elastic Cloud Phone (ECP) Key Pair can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ecp/keyPair:KeyPair example <key_pair_name>
 * ```
 */
export class KeyPair extends pulumi.CustomResource {
    /**
     * Get an existing KeyPair resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyPairState, opts?: pulumi.CustomResourceOptions): KeyPair {
        return new KeyPair(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecp/keyPair:KeyPair';

    /**
     * Returns true if the given object is an instance of KeyPair.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeyPair {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeyPair.__pulumiType;
    }

    /**
     * The Key Name.
     */
    public readonly keyPairName!: pulumi.Output<string>;
    /**
     * The public key body.
     */
    public readonly publicKeyBody!: pulumi.Output<string>;

    /**
     * Create a KeyPair resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyPairArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyPairArgs | KeyPairState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeyPairState | undefined;
            resourceInputs["keyPairName"] = state ? state.keyPairName : undefined;
            resourceInputs["publicKeyBody"] = state ? state.publicKeyBody : undefined;
        } else {
            const args = argsOrState as KeyPairArgs | undefined;
            if ((!args || args.keyPairName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyPairName'");
            }
            if ((!args || args.publicKeyBody === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicKeyBody'");
            }
            resourceInputs["keyPairName"] = args ? args.keyPairName : undefined;
            resourceInputs["publicKeyBody"] = args ? args.publicKeyBody : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KeyPair.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeyPair resources.
 */
export interface KeyPairState {
    /**
     * The Key Name.
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * The public key body.
     */
    publicKeyBody?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KeyPair resource.
 */
export interface KeyPairArgs {
    /**
     * The Key Name.
     */
    keyPairName: pulumi.Input<string>;
    /**
     * The public key body.
     */
    publicKeyBody: pulumi.Input<string>;
}
