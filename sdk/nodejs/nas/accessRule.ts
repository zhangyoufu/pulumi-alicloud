// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a NAS Access Rule resource.
 *
 * For information about NAS Access Rule and how to use it, see [What is Access Rule](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-createaccessrule).
 *
 * > **NOTE:** Available since v1.34.0.
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
 * const foo = new alicloud.nas.AccessGroup("foo", {
 *     accessGroupName: "tf-NasConfigName",
 *     accessGroupType: "Vpc",
 *     description: "tf-testAccNasConfig",
 * });
 * const fooAccessRule = new alicloud.nas.AccessRule("foo", {
 *     accessGroupName: foo.accessGroupName,
 *     sourceCidrIp: "168.1.1.0/16",
 *     rwAccessType: "RDWR",
 *     userAccessType: "no_squash",
 *     priority: 2,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * NAS Access Rule can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:nas/accessRule:AccessRule example <access_group_name>:<file_system_type>:<access_rule_id>
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
     * @param opts Optional settings to control the behavior of the CustomResource.
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
     * AccessGroupName.
     */
    public readonly accessGroupName!: pulumi.Output<string>;
    /**
     * The first ID of the resource.
     */
    public /*out*/ readonly accessRuleId!: pulumi.Output<string>;
    /**
     * filesystem type. include standard, extreme.
     */
    public readonly fileSystemType!: pulumi.Output<string>;
    /**
     * Ipv6SourceCidrIp.
     */
    public readonly ipv6SourceCidrIp!: pulumi.Output<string | undefined>;
    /**
     * Priority.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * RWAccess.
     */
    public readonly rwAccessType!: pulumi.Output<string>;
    /**
     * SourceCidrIp.
     */
    public readonly sourceCidrIp!: pulumi.Output<string | undefined>;
    /**
     * UserAccess.
     */
    public readonly userAccessType!: pulumi.Output<string>;

    /**
     * Create a AccessRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessRuleArgs | AccessRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessRuleState | undefined;
            resourceInputs["accessGroupName"] = state ? state.accessGroupName : undefined;
            resourceInputs["accessRuleId"] = state ? state.accessRuleId : undefined;
            resourceInputs["fileSystemType"] = state ? state.fileSystemType : undefined;
            resourceInputs["ipv6SourceCidrIp"] = state ? state.ipv6SourceCidrIp : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["rwAccessType"] = state ? state.rwAccessType : undefined;
            resourceInputs["sourceCidrIp"] = state ? state.sourceCidrIp : undefined;
            resourceInputs["userAccessType"] = state ? state.userAccessType : undefined;
        } else {
            const args = argsOrState as AccessRuleArgs | undefined;
            if ((!args || args.accessGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessGroupName'");
            }
            resourceInputs["accessGroupName"] = args ? args.accessGroupName : undefined;
            resourceInputs["fileSystemType"] = args ? args.fileSystemType : undefined;
            resourceInputs["ipv6SourceCidrIp"] = args ? args.ipv6SourceCidrIp : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["rwAccessType"] = args ? args.rwAccessType : undefined;
            resourceInputs["sourceCidrIp"] = args ? args.sourceCidrIp : undefined;
            resourceInputs["userAccessType"] = args ? args.userAccessType : undefined;
            resourceInputs["accessRuleId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccessRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessRule resources.
 */
export interface AccessRuleState {
    /**
     * AccessGroupName.
     */
    accessGroupName?: pulumi.Input<string>;
    /**
     * The first ID of the resource.
     */
    accessRuleId?: pulumi.Input<string>;
    /**
     * filesystem type. include standard, extreme.
     */
    fileSystemType?: pulumi.Input<string>;
    /**
     * Ipv6SourceCidrIp.
     */
    ipv6SourceCidrIp?: pulumi.Input<string>;
    /**
     * Priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * RWAccess.
     */
    rwAccessType?: pulumi.Input<string>;
    /**
     * SourceCidrIp.
     */
    sourceCidrIp?: pulumi.Input<string>;
    /**
     * UserAccess.
     */
    userAccessType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessRule resource.
 */
export interface AccessRuleArgs {
    /**
     * AccessGroupName.
     */
    accessGroupName: pulumi.Input<string>;
    /**
     * filesystem type. include standard, extreme.
     */
    fileSystemType?: pulumi.Input<string>;
    /**
     * Ipv6SourceCidrIp.
     */
    ipv6SourceCidrIp?: pulumi.Input<string>;
    /**
     * Priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * RWAccess.
     */
    rwAccessType?: pulumi.Input<string>;
    /**
     * SourceCidrIp.
     */
    sourceCidrIp?: pulumi.Input<string>;
    /**
     * UserAccess.
     */
    userAccessType?: pulumi.Input<string>;
}
