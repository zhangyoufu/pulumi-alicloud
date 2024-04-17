// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a DFS Access Rule resource.
 *
 * For information about DFS Access Rule and how to use it, see [What is Access Rule](https://www.alibabacloud.com/help/en/aibaba-cloud-storage-services/latest/apsara-file-storage-for-hdfs).
 *
 * > **NOTE:** Available since v1.140.0.
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
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const _default = new alicloud.dfs.AccessGroup("default", {
 *     description: "example",
 *     networkType: "VPC",
 *     accessGroupName: name,
 * });
 * const defaultAccessRule = new alicloud.dfs.AccessRule("default", {
 *     description: "example",
 *     rwAccessType: "RDWR",
 *     priority: 1,
 *     networkSegment: "192.168.81.1",
 *     accessGroupId: _default.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * DFS Access Rule can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:dfs/accessRule:AccessRule example <access_group_id>:<access_rule_id>
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
    public static readonly __pulumiType = 'alicloud:dfs/accessRule:AccessRule';

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
     * Permission group resource ID. You must specify the permission group ID when creating a permission rule.
     */
    public readonly accessGroupId!: pulumi.Output<string>;
    /**
     * The unique identity of the permission rule, which is used to retrieve the permission rule for a specific day in the permission group.
     */
    public /*out*/ readonly accessRuleId!: pulumi.Output<string>;
    /**
     * Permission rule resource creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Permission rule description.  No more than 32 characters in length.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The IP address or network segment of the authorized object.
     */
    public readonly networkSegment!: pulumi.Output<string>;
    /**
     * Permission rule priority. When the same authorization object matches multiple rules, the high-priority rule takes effect. Value range: 1~100,1 is the highest priority.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * The read and write permissions of the authorized object on the file system. Value: RDWR: readable and writable RDONLY: Read only.
     */
    public readonly rwAccessType!: pulumi.Output<string>;

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
            resourceInputs["accessGroupId"] = state ? state.accessGroupId : undefined;
            resourceInputs["accessRuleId"] = state ? state.accessRuleId : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["networkSegment"] = state ? state.networkSegment : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["rwAccessType"] = state ? state.rwAccessType : undefined;
        } else {
            const args = argsOrState as AccessRuleArgs | undefined;
            if ((!args || args.accessGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessGroupId'");
            }
            if ((!args || args.networkSegment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkSegment'");
            }
            if ((!args || args.priority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'priority'");
            }
            if ((!args || args.rwAccessType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rwAccessType'");
            }
            resourceInputs["accessGroupId"] = args ? args.accessGroupId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["networkSegment"] = args ? args.networkSegment : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["rwAccessType"] = args ? args.rwAccessType : undefined;
            resourceInputs["accessRuleId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
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
     * Permission group resource ID. You must specify the permission group ID when creating a permission rule.
     */
    accessGroupId?: pulumi.Input<string>;
    /**
     * The unique identity of the permission rule, which is used to retrieve the permission rule for a specific day in the permission group.
     */
    accessRuleId?: pulumi.Input<string>;
    /**
     * Permission rule resource creation time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Permission rule description.  No more than 32 characters in length.
     */
    description?: pulumi.Input<string>;
    /**
     * The IP address or network segment of the authorized object.
     */
    networkSegment?: pulumi.Input<string>;
    /**
     * Permission rule priority. When the same authorization object matches multiple rules, the high-priority rule takes effect. Value range: 1~100,1 is the highest priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * The read and write permissions of the authorized object on the file system. Value: RDWR: readable and writable RDONLY: Read only.
     */
    rwAccessType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessRule resource.
 */
export interface AccessRuleArgs {
    /**
     * Permission group resource ID. You must specify the permission group ID when creating a permission rule.
     */
    accessGroupId: pulumi.Input<string>;
    /**
     * Permission rule description.  No more than 32 characters in length.
     */
    description?: pulumi.Input<string>;
    /**
     * The IP address or network segment of the authorized object.
     */
    networkSegment: pulumi.Input<string>;
    /**
     * Permission rule priority. When the same authorization object matches multiple rules, the high-priority rule takes effect. Value range: 1~100,1 is the highest priority.
     */
    priority: pulumi.Input<number>;
    /**
     * The read and write permissions of the authorized object on the file system. Value: RDWR: readable and writable RDONLY: Read only.
     */
    rwAccessType: pulumi.Input<string>;
}
