// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a key pair resource.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const basic = new alicloud.ecs.KeyPair("basic", {
 *     keyName: "terraform-test-key-pair",
 * });
 * // Using name prefix to build key pair
 * const prefix = new alicloud.ecs.KeyPair("prefix", {
 *     keyNamePrefix: "terraform-test-key-pair-prefix",
 * });
 * // Import an existing public key to build a alicloud key pair
 * const publickey = new alicloud.ecs.KeyPair("publickey", {
 *     keyName: "my_public_key",
 *     publicKey: "ssh-rsa AAAAB3Nza12345678qwertyuudsfsg",
 * });
 * ```
 *
 * ## Import
 *
 * Key pair can be imported using the name, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ecs/keyPair:KeyPair example my_public_key
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
    public static readonly __pulumiType = 'alicloud:ecs/keyPair:KeyPair';

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

    public /*out*/ readonly fingerPrint!: pulumi.Output<string>;
    /**
     * The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
     */
    public readonly keyFile!: pulumi.Output<string | undefined>;
    /**
     * The key pair's name. It is the only in one Alicloud account.
     */
    public readonly keyName!: pulumi.Output<string>;
    public readonly keyNamePrefix!: pulumi.Output<string | undefined>;
    /**
     * You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resourceGroupId` is the key pair belongs.
     */
    public readonly publicKey!: pulumi.Output<string | undefined>;
    /**
     * The Id of resource group which the key pair belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a KeyPair resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: KeyPairArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyPairArgs | KeyPairState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeyPairState | undefined;
            inputs["fingerPrint"] = state ? state.fingerPrint : undefined;
            inputs["keyFile"] = state ? state.keyFile : undefined;
            inputs["keyName"] = state ? state.keyName : undefined;
            inputs["keyNamePrefix"] = state ? state.keyNamePrefix : undefined;
            inputs["publicKey"] = state ? state.publicKey : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as KeyPairArgs | undefined;
            inputs["keyFile"] = args ? args.keyFile : undefined;
            inputs["keyName"] = args ? args.keyName : undefined;
            inputs["keyNamePrefix"] = args ? args.keyNamePrefix : undefined;
            inputs["publicKey"] = args ? args.publicKey : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["fingerPrint"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(KeyPair.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeyPair resources.
 */
export interface KeyPairState {
    readonly fingerPrint?: pulumi.Input<string>;
    /**
     * The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
     */
    readonly keyFile?: pulumi.Input<string>;
    /**
     * The key pair's name. It is the only in one Alicloud account.
     */
    readonly keyName?: pulumi.Input<string>;
    readonly keyNamePrefix?: pulumi.Input<string>;
    /**
     * You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resourceGroupId` is the key pair belongs.
     */
    readonly publicKey?: pulumi.Input<string>;
    /**
     * The Id of resource group which the key pair belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a KeyPair resource.
 */
export interface KeyPairArgs {
    /**
     * The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
     */
    readonly keyFile?: pulumi.Input<string>;
    /**
     * The key pair's name. It is the only in one Alicloud account.
     */
    readonly keyName?: pulumi.Input<string>;
    readonly keyNamePrefix?: pulumi.Input<string>;
    /**
     * You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resourceGroupId` is the key pair belongs.
     */
    readonly publicKey?: pulumi.Input<string>;
    /**
     * The Id of resource group which the key pair belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
