// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Nas Access Rule resource.
 *
 * When NAS is activated, the Default VPC Permission Group is automatically generated. It allows all IP addresses in a VPC to access the mount point with full permissions. Full permissions include Read/Write permission with no restriction on root users.
 *
 * > **NOTE:** Available in v1.34.0+.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const fooAccessGroup = new alicloud.nas.AccessGroup("foo", {
 *     description: "tf-testAccNasConfig",
 *     type: "Vpc",
 * });
 * const fooAccessRule = new alicloud.nas.AccessRule("foo", {
 *     accessGroupName: fooAccessGroup.id,
 *     priority: 2,
 *     rwAccessType: "RDWR",
 *     sourceCidrIp: "168.1.1.0/16",
 *     userAccessType: "noSquash",
 * });
 * ```
 */
export class AccessRule extends pulumi.CustomResource {
    /**
     * Get an existing AccessRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessRuleState, opts?: pulumi.CustomResourceOptions): AccessRule {
        return new AccessRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:nas/accessRule:AccessRule';

    /**
     * Returns true if the given object is an instance of AccessRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessRule.__pulumiType;
    }

    /**
     * Permission group name.
     */
    public readonly accessGroupName!: pulumi.Output<string>;
    /**
     * The nas access rule ID.
     */
    public /*out*/ readonly accessRuleId!: pulumi.Output<string>;
    /**
     * Priority level. Range: 1-100. Default value: 1.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * Read-write permission type: RDWR (default), RDONLY.
     */
    public readonly rwAccessType!: pulumi.Output<string | undefined>;
    /**
     * Address or address segment.
     */
    public readonly sourceCidrIp!: pulumi.Output<string>;
    /**
     * User permission type: noSquash (default), root_squash, all_squash.
     */
    public readonly userAccessType!: pulumi.Output<string | undefined>;

    /**
     * Create a AccessRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessRuleArgs | AccessRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AccessRuleState | undefined;
            inputs["accessGroupName"] = state ? state.accessGroupName : undefined;
            inputs["accessRuleId"] = state ? state.accessRuleId : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["rwAccessType"] = state ? state.rwAccessType : undefined;
            inputs["sourceCidrIp"] = state ? state.sourceCidrIp : undefined;
            inputs["userAccessType"] = state ? state.userAccessType : undefined;
        } else {
            const args = argsOrState as AccessRuleArgs | undefined;
            if (!args || args.accessGroupName === undefined) {
                throw new Error("Missing required property 'accessGroupName'");
            }
            if (!args || args.sourceCidrIp === undefined) {
                throw new Error("Missing required property 'sourceCidrIp'");
            }
            inputs["accessGroupName"] = args ? args.accessGroupName : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["rwAccessType"] = args ? args.rwAccessType : undefined;
            inputs["sourceCidrIp"] = args ? args.sourceCidrIp : undefined;
            inputs["userAccessType"] = args ? args.userAccessType : undefined;
            inputs["accessRuleId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AccessRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessRule resources.
 */
export interface AccessRuleState {
    /**
     * Permission group name.
     */
    readonly accessGroupName?: pulumi.Input<string>;
    /**
     * The nas access rule ID.
     */
    readonly accessRuleId?: pulumi.Input<string>;
    /**
     * Priority level. Range: 1-100. Default value: 1.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * Read-write permission type: RDWR (default), RDONLY.
     */
    readonly rwAccessType?: pulumi.Input<string>;
    /**
     * Address or address segment.
     */
    readonly sourceCidrIp?: pulumi.Input<string>;
    /**
     * User permission type: noSquash (default), root_squash, all_squash.
     */
    readonly userAccessType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessRule resource.
 */
export interface AccessRuleArgs {
    /**
     * Permission group name.
     */
    readonly accessGroupName: pulumi.Input<string>;
    /**
     * Priority level. Range: 1-100. Default value: 1.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * Read-write permission type: RDWR (default), RDONLY.
     */
    readonly rwAccessType?: pulumi.Input<string>;
    /**
     * Address or address segment.
     */
    readonly sourceCidrIp: pulumi.Input<string>;
    /**
     * User permission type: noSquash (default), root_squash, all_squash.
     */
    readonly userAccessType?: pulumi.Input<string>;
}
