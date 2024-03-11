// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Monitor Service Dynamic Tag Group resource.
 *
 * For information about Cloud Monitor Service Dynamic Tag Group and how to use it, see [What is Dynamic Tag Group](https://www.alibabacloud.com/help/en/cloudmonitor/latest/createdynamictaggroup).
 *
 * > **NOTE:** Available since v1.142.0.
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
 * const defaultAlarmContactGroup = new alicloud.cms.AlarmContactGroup("defaultAlarmContactGroup", {
 *     alarmContactGroupName: "example_value",
 *     describe: "example_value",
 *     enableSubscribed: true,
 * });
 * const defaultDynamicTagGroup = new alicloud.cms.DynamicTagGroup("defaultDynamicTagGroup", {
 *     contactGroupLists: [defaultAlarmContactGroup.id],
 *     tagKey: "your_tag_key",
 *     matchExpresses: [{
 *         tagValue: "your_tag_value",
 *         tagValueMatchFunction: "all",
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Cloud Monitor Service Dynamic Tag Group can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cms/dynamicTagGroup:DynamicTagGroup example <id>
 * ```
 */
export class DynamicTagGroup extends pulumi.CustomResource {
    /**
     * Get an existing DynamicTagGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DynamicTagGroupState, opts?: pulumi.CustomResourceOptions): DynamicTagGroup {
        return new DynamicTagGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cms/dynamicTagGroup:DynamicTagGroup';

    /**
     * Returns true if the given object is an instance of DynamicTagGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DynamicTagGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DynamicTagGroup.__pulumiType;
    }

    /**
     * Alarm contact group. The value range of N is 1~100. The alarm notification of the application group is sent to the alarm contact in the alarm contact group.
     */
    public readonly contactGroupLists!: pulumi.Output<string[]>;
    /**
     * The relationship between conditional expressions. Valid values: `and`, `or`.
     */
    public readonly matchExpressFilterRelation!: pulumi.Output<string>;
    /**
     * The label generates a matching expression that applies the grouping. See `matchExpress` below.
     */
    public readonly matchExpresses!: pulumi.Output<outputs.cms.DynamicTagGroupMatchExpress[]>;
    /**
     * The status of the resource. Valid values: `RUNNING`, `FINISH`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The tag key of the tag.
     */
    public readonly tagKey!: pulumi.Output<string>;
    /**
     * Alarm template ID list.
     */
    public readonly templateIdLists!: pulumi.Output<string[] | undefined>;

    /**
     * Create a DynamicTagGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DynamicTagGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DynamicTagGroupArgs | DynamicTagGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DynamicTagGroupState | undefined;
            resourceInputs["contactGroupLists"] = state ? state.contactGroupLists : undefined;
            resourceInputs["matchExpressFilterRelation"] = state ? state.matchExpressFilterRelation : undefined;
            resourceInputs["matchExpresses"] = state ? state.matchExpresses : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tagKey"] = state ? state.tagKey : undefined;
            resourceInputs["templateIdLists"] = state ? state.templateIdLists : undefined;
        } else {
            const args = argsOrState as DynamicTagGroupArgs | undefined;
            if ((!args || args.contactGroupLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contactGroupLists'");
            }
            if ((!args || args.matchExpresses === undefined) && !opts.urn) {
                throw new Error("Missing required property 'matchExpresses'");
            }
            if ((!args || args.tagKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tagKey'");
            }
            resourceInputs["contactGroupLists"] = args ? args.contactGroupLists : undefined;
            resourceInputs["matchExpressFilterRelation"] = args ? args.matchExpressFilterRelation : undefined;
            resourceInputs["matchExpresses"] = args ? args.matchExpresses : undefined;
            resourceInputs["tagKey"] = args ? args.tagKey : undefined;
            resourceInputs["templateIdLists"] = args ? args.templateIdLists : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DynamicTagGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DynamicTagGroup resources.
 */
export interface DynamicTagGroupState {
    /**
     * Alarm contact group. The value range of N is 1~100. The alarm notification of the application group is sent to the alarm contact in the alarm contact group.
     */
    contactGroupLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The relationship between conditional expressions. Valid values: `and`, `or`.
     */
    matchExpressFilterRelation?: pulumi.Input<string>;
    /**
     * The label generates a matching expression that applies the grouping. See `matchExpress` below.
     */
    matchExpresses?: pulumi.Input<pulumi.Input<inputs.cms.DynamicTagGroupMatchExpress>[]>;
    /**
     * The status of the resource. Valid values: `RUNNING`, `FINISH`.
     */
    status?: pulumi.Input<string>;
    /**
     * The tag key of the tag.
     */
    tagKey?: pulumi.Input<string>;
    /**
     * Alarm template ID list.
     */
    templateIdLists?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a DynamicTagGroup resource.
 */
export interface DynamicTagGroupArgs {
    /**
     * Alarm contact group. The value range of N is 1~100. The alarm notification of the application group is sent to the alarm contact in the alarm contact group.
     */
    contactGroupLists: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The relationship between conditional expressions. Valid values: `and`, `or`.
     */
    matchExpressFilterRelation?: pulumi.Input<string>;
    /**
     * The label generates a matching expression that applies the grouping. See `matchExpress` below.
     */
    matchExpresses: pulumi.Input<pulumi.Input<inputs.cms.DynamicTagGroupMatchExpress>[]>;
    /**
     * The tag key of the tag.
     */
    tagKey: pulumi.Input<string>;
    /**
     * Alarm template ID list.
     */
    templateIdLists?: pulumi.Input<pulumi.Input<string>[]>;
}
