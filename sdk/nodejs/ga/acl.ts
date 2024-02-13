// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Global Accelerator (GA) Acl resource.
 *
 * For information about Global Accelerator (GA) Acl and how to use it, see [What is Acl](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createacl).
 *
 * > **NOTE:** Available since v1.150.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultAcl = new alicloud.ga.Acl("defaultAcl", {
 *     aclName: "terraform-example",
 *     addressIpVersion: "IPv4",
 * });
 * const defaultAclEntryAttachment = new alicloud.ga.AclEntryAttachment("defaultAclEntryAttachment", {
 *     aclId: defaultAcl.id,
 *     entry: "192.168.1.1/32",
 *     entryDescription: "terraform-example",
 * });
 * ```
 *
 * ## Import
 *
 * Global Accelerator (GA) Acl can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ga/acl:Acl example <id>
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
    public static readonly __pulumiType = 'alicloud:ga/acl:Acl';

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
     * The entries of the Acl. See `aclEntries` below. **NOTE:** "Field `aclEntries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`."
     *
     * @deprecated Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud_ga_acl_entry_attachment`.
     */
    public readonly aclEntries!: pulumi.Output<outputs.ga.AclAclEntry[]>;
    /**
     * The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
     */
    public readonly aclName!: pulumi.Output<string | undefined>;
    /**
     * The IP version. Valid values: `IPv4` and `IPv6`.
     */
    public readonly addressIpVersion!: pulumi.Output<string>;
    /**
     * The dry run.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The status of the resource.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

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
            resourceInputs["addressIpVersion"] = state ? state.addressIpVersion : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as AclArgs | undefined;
            if ((!args || args.addressIpVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addressIpVersion'");
            }
            resourceInputs["aclEntries"] = args ? args.aclEntries : undefined;
            resourceInputs["aclName"] = args ? args.aclName : undefined;
            resourceInputs["addressIpVersion"] = args ? args.addressIpVersion : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
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
     * The entries of the Acl. See `aclEntries` below. **NOTE:** "Field `aclEntries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`."
     *
     * @deprecated Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud_ga_acl_entry_attachment`.
     */
    aclEntries?: pulumi.Input<pulumi.Input<inputs.ga.AclAclEntry>[]>;
    /**
     * The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
     */
    aclName?: pulumi.Input<string>;
    /**
     * The IP version. Valid values: `IPv4` and `IPv6`.
     */
    addressIpVersion?: pulumi.Input<string>;
    /**
     * The dry run.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Acl resource.
 */
export interface AclArgs {
    /**
     * The entries of the Acl. See `aclEntries` below. **NOTE:** "Field `aclEntries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`."
     *
     * @deprecated Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud_ga_acl_entry_attachment`.
     */
    aclEntries?: pulumi.Input<pulumi.Input<inputs.ga.AclAclEntry>[]>;
    /**
     * The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
     */
    aclName?: pulumi.Input<string>;
    /**
     * The IP version. Valid values: `IPv4` and `IPv6`.
     */
    addressIpVersion: pulumi.Input<string>;
    /**
     * The dry run.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
