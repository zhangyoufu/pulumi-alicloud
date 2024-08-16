// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Application Load Balancer (ALB) Acl resource.
 *
 * For information about ALB Acl and how to use it, see [What is Acl](https://www.alibabacloud.com/help/en/slb/application-load-balancer/developer-reference/api-alb-2020-06-16-createacl).
 *
 * > **NOTE:** Available since v1.133.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.resourcemanager.getResourceGroups({});
 * const defaultAcl = new alicloud.alb.Acl("default", {
 *     aclName: "tf_example",
 *     resourceGroupId: _default.then(_default => _default.groups?.[0]?.id),
 * });
 * ```
 *
 * ## Import
 *
 * ALB Acl can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:alb/acl:Acl example <id>
 * ```
 */
export class Acl extends pulumi.CustomResource {
    /**
     * Get an existing Acl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AclState, opts?: pulumi.CustomResourceOptions): Acl {
        return new Acl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:alb/acl:Acl';

    /**
     * Returns true if the given object is an instance of Acl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Acl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Acl.__pulumiType;
    }

    /**
     * The list of the ACL entries. You can add up to `20` entries in each call.  See `aclEntries` below for details.
     * **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.166.0 and it will be removed in the future version. Please use the new resource 'alicloud_alb_acl_entry_attachment'.",
     *
     * @deprecated Field 'acl_entries' has been deprecated from provider version 1.166.0 and it will be removed in the future version. Please use the new resource 'alicloud_alb_acl_entry_attachment'.
     */
    public readonly aclEntries!: pulumi.Output<outputs.alb.AclAclEntry[]>;
    /**
     * The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, hyphens (-) and underscores (_). It must start with a letter.
     */
    public readonly aclName!: pulumi.Output<string>;
    /**
     * Specifies whether to precheck the API request.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The state of the ACL. Valid values:`Provisioning`, `Available` and `Configuring`. `Provisioning`: The ACL is being created. `Available`: The ACL is available. `Configuring`: The ACL is being configured.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Acl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AclArgs | AclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AclState | undefined;
            resourceInputs["aclEntries"] = state ? state.aclEntries : undefined;
            resourceInputs["aclName"] = state ? state.aclName : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as AclArgs | undefined;
            if ((!args || args.aclName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aclName'");
            }
            resourceInputs["aclEntries"] = args ? args.aclEntries : undefined;
            resourceInputs["aclName"] = args ? args.aclName : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Acl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Acl resources.
 */
export interface AclState {
    /**
     * The list of the ACL entries. You can add up to `20` entries in each call.  See `aclEntries` below for details.
     * **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.166.0 and it will be removed in the future version. Please use the new resource 'alicloud_alb_acl_entry_attachment'.",
     *
     * @deprecated Field 'acl_entries' has been deprecated from provider version 1.166.0 and it will be removed in the future version. Please use the new resource 'alicloud_alb_acl_entry_attachment'.
     */
    aclEntries?: pulumi.Input<pulumi.Input<inputs.alb.AclAclEntry>[]>;
    /**
     * The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, hyphens (-) and underscores (_). It must start with a letter.
     */
    aclName?: pulumi.Input<string>;
    /**
     * Specifies whether to precheck the API request.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The state of the ACL. Valid values:`Provisioning`, `Available` and `Configuring`. `Provisioning`: The ACL is being created. `Available`: The ACL is available. `Configuring`: The ACL is being configured.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Acl resource.
 */
export interface AclArgs {
    /**
     * The list of the ACL entries. You can add up to `20` entries in each call.  See `aclEntries` below for details.
     * **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.166.0 and it will be removed in the future version. Please use the new resource 'alicloud_alb_acl_entry_attachment'.",
     *
     * @deprecated Field 'acl_entries' has been deprecated from provider version 1.166.0 and it will be removed in the future version. Please use the new resource 'alicloud_alb_acl_entry_attachment'.
     */
    aclEntries?: pulumi.Input<pulumi.Input<inputs.alb.AclAclEntry>[]>;
    /**
     * The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, hyphens (-) and underscores (_). It must start with a letter.
     */
    aclName: pulumi.Input<string>;
    /**
     * Specifies whether to precheck the API request.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
