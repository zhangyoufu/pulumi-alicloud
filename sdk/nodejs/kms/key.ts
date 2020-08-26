// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A kms key can help user to protect data security in the transmission process. For information about Alikms Key and how to use it, see [What is Resource Alikms Key](https://www.alibabacloud.com/help/doc-detail/28947.htm).
 *
 * > **NOTE:** Available in v1.85.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const key = new alicloud.kms.Key("key", {
 *     description: "Hello KMS",
 *     keyState: "Enabled",
 *     pendingWindowInDays: 7,
 * });
 * ```
 */
export class Key extends pulumi.CustomResource {
    /**
     * Get an existing Key resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyState, opts?: pulumi.CustomResourceOptions): Key {
        return new Key(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:kms/key:Key';

    /**
     * Returns true if the given object is an instance of Key.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Key {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Key.__pulumiType;
    }

    /**
     * The Alicloud Resource Name (ARN) of the key.
     * * `creationDate` -The date and time when the CMK was created. The time is displayed in UTC.
     * * `creator` -The creator of the CMK.
     * * `deleteDate` -The scheduled date to delete CMK. The time is displayed in UTC. This value is returned only when the KeyState value is PendingDeletion.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Specifies whether to enable automatic key rotation. Default:"Disabled".
     */
    public readonly automaticRotation!: pulumi.Output<string | undefined>;
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    public /*out*/ readonly creator!: pulumi.Output<string>;
    public /*out*/ readonly deleteDate!: pulumi.Output<string>;
    /**
     * Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
     *
     * @deprecated Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
     */
    public readonly deletionWindowInDays!: pulumi.Output<number | undefined>;
    /**
     * The description of the key as viewed in Alicloud console.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
     *
     * @deprecated Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
     */
    public readonly isEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The type of the CMK.
     */
    public readonly keySpec!: pulumi.Output<string>;
    /**
     * The status of CMK. Defaults to Enabled.
     */
    public readonly keyState!: pulumi.Output<string | undefined>;
    /**
     * Specifies the usage of CMK. Currently, default to 'ENCRYPT/DECRYPT', indicating that CMK is used for encryption and decryption.
     */
    public readonly keyUsage!: pulumi.Output<string | undefined>;
    /**
     * The date and time the last rotation was performed. The time is displayed in UTC.
     */
    public /*out*/ readonly lastRotationDate!: pulumi.Output<string>;
    /**
     * The time and date the key material for the CMK expires. The time is displayed in UTC. If the value is empty, the key material for the CMK does not expire.
     */
    public /*out*/ readonly materialExpireTime!: pulumi.Output<string>;
    /**
     * The time the next rotation is scheduled for execution.
     */
    public /*out*/ readonly nextRotationDate!: pulumi.Output<string>;
    /**
     * The source of the key material for the CMK. Defaults to "Aliyun_KMS".
     */
    public readonly origin!: pulumi.Output<string | undefined>;
    /**
     * Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
     */
    public readonly pendingWindowInDays!: pulumi.Output<number | undefined>;
    /**
     * The ID of the current primary key version of the symmetric CMK.
     */
    public /*out*/ readonly primaryKeyVersion!: pulumi.Output<string>;
    /**
     * The protection level of the CMK. Defaults to "SOFTWARE".
     */
    public readonly protectionLevel!: pulumi.Output<string | undefined>;
    /**
     * The period of automatic key rotation. Unit: seconds.
     */
    public readonly rotationInterval!: pulumi.Output<string | undefined>;

    /**
     * Create a Key resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: KeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyArgs | KeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as KeyState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["automaticRotation"] = state ? state.automaticRotation : undefined;
            inputs["creationDate"] = state ? state.creationDate : undefined;
            inputs["creator"] = state ? state.creator : undefined;
            inputs["deleteDate"] = state ? state.deleteDate : undefined;
            inputs["deletionWindowInDays"] = state ? state.deletionWindowInDays : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["isEnabled"] = state ? state.isEnabled : undefined;
            inputs["keySpec"] = state ? state.keySpec : undefined;
            inputs["keyState"] = state ? state.keyState : undefined;
            inputs["keyUsage"] = state ? state.keyUsage : undefined;
            inputs["lastRotationDate"] = state ? state.lastRotationDate : undefined;
            inputs["materialExpireTime"] = state ? state.materialExpireTime : undefined;
            inputs["nextRotationDate"] = state ? state.nextRotationDate : undefined;
            inputs["origin"] = state ? state.origin : undefined;
            inputs["pendingWindowInDays"] = state ? state.pendingWindowInDays : undefined;
            inputs["primaryKeyVersion"] = state ? state.primaryKeyVersion : undefined;
            inputs["protectionLevel"] = state ? state.protectionLevel : undefined;
            inputs["rotationInterval"] = state ? state.rotationInterval : undefined;
        } else {
            const args = argsOrState as KeyArgs | undefined;
            inputs["automaticRotation"] = args ? args.automaticRotation : undefined;
            inputs["deletionWindowInDays"] = args ? args.deletionWindowInDays : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["isEnabled"] = args ? args.isEnabled : undefined;
            inputs["keySpec"] = args ? args.keySpec : undefined;
            inputs["keyState"] = args ? args.keyState : undefined;
            inputs["keyUsage"] = args ? args.keyUsage : undefined;
            inputs["origin"] = args ? args.origin : undefined;
            inputs["pendingWindowInDays"] = args ? args.pendingWindowInDays : undefined;
            inputs["protectionLevel"] = args ? args.protectionLevel : undefined;
            inputs["rotationInterval"] = args ? args.rotationInterval : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["creationDate"] = undefined /*out*/;
            inputs["creator"] = undefined /*out*/;
            inputs["deleteDate"] = undefined /*out*/;
            inputs["lastRotationDate"] = undefined /*out*/;
            inputs["materialExpireTime"] = undefined /*out*/;
            inputs["nextRotationDate"] = undefined /*out*/;
            inputs["primaryKeyVersion"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Key.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Key resources.
 */
export interface KeyState {
    /**
     * The Alicloud Resource Name (ARN) of the key.
     * * `creationDate` -The date and time when the CMK was created. The time is displayed in UTC.
     * * `creator` -The creator of the CMK.
     * * `deleteDate` -The scheduled date to delete CMK. The time is displayed in UTC. This value is returned only when the KeyState value is PendingDeletion.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Specifies whether to enable automatic key rotation. Default:"Disabled".
     */
    readonly automaticRotation?: pulumi.Input<string>;
    readonly creationDate?: pulumi.Input<string>;
    readonly creator?: pulumi.Input<string>;
    readonly deleteDate?: pulumi.Input<string>;
    /**
     * Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
     *
     * @deprecated Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
     */
    readonly deletionWindowInDays?: pulumi.Input<number>;
    /**
     * The description of the key as viewed in Alicloud console.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
     *
     * @deprecated Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
     */
    readonly isEnabled?: pulumi.Input<boolean>;
    /**
     * The type of the CMK.
     */
    readonly keySpec?: pulumi.Input<string>;
    /**
     * The status of CMK. Defaults to Enabled.
     */
    readonly keyState?: pulumi.Input<string>;
    /**
     * Specifies the usage of CMK. Currently, default to 'ENCRYPT/DECRYPT', indicating that CMK is used for encryption and decryption.
     */
    readonly keyUsage?: pulumi.Input<string>;
    /**
     * The date and time the last rotation was performed. The time is displayed in UTC.
     */
    readonly lastRotationDate?: pulumi.Input<string>;
    /**
     * The time and date the key material for the CMK expires. The time is displayed in UTC. If the value is empty, the key material for the CMK does not expire.
     */
    readonly materialExpireTime?: pulumi.Input<string>;
    /**
     * The time the next rotation is scheduled for execution.
     */
    readonly nextRotationDate?: pulumi.Input<string>;
    /**
     * The source of the key material for the CMK. Defaults to "Aliyun_KMS".
     */
    readonly origin?: pulumi.Input<string>;
    /**
     * Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
     */
    readonly pendingWindowInDays?: pulumi.Input<number>;
    /**
     * The ID of the current primary key version of the symmetric CMK.
     */
    readonly primaryKeyVersion?: pulumi.Input<string>;
    /**
     * The protection level of the CMK. Defaults to "SOFTWARE".
     */
    readonly protectionLevel?: pulumi.Input<string>;
    /**
     * The period of automatic key rotation. Unit: seconds.
     */
    readonly rotationInterval?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Key resource.
 */
export interface KeyArgs {
    /**
     * Specifies whether to enable automatic key rotation. Default:"Disabled".
     */
    readonly automaticRotation?: pulumi.Input<string>;
    /**
     * Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
     *
     * @deprecated Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
     */
    readonly deletionWindowInDays?: pulumi.Input<number>;
    /**
     * The description of the key as viewed in Alicloud console.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
     *
     * @deprecated Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
     */
    readonly isEnabled?: pulumi.Input<boolean>;
    /**
     * The type of the CMK.
     */
    readonly keySpec?: pulumi.Input<string>;
    /**
     * The status of CMK. Defaults to Enabled.
     */
    readonly keyState?: pulumi.Input<string>;
    /**
     * Specifies the usage of CMK. Currently, default to 'ENCRYPT/DECRYPT', indicating that CMK is used for encryption and decryption.
     */
    readonly keyUsage?: pulumi.Input<string>;
    /**
     * The source of the key material for the CMK. Defaults to "Aliyun_KMS".
     */
    readonly origin?: pulumi.Input<string>;
    /**
     * Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
     */
    readonly pendingWindowInDays?: pulumi.Input<number>;
    /**
     * The protection level of the CMK. Defaults to "SOFTWARE".
     */
    readonly protectionLevel?: pulumi.Input<string>;
    /**
     * The period of automatic key rotation. Unit: seconds.
     */
    readonly rotationInterval?: pulumi.Input<string>;
}
