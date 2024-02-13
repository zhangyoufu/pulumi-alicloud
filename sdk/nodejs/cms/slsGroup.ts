// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Monitor Service Sls Group resource.
 *
 * For information about Cloud Monitor Service Sls Group and how to use it, see [What is Sls Group](https://www.alibabacloud.com/help/doc-detail/28608.htm).
 *
 * > **NOTE:** Available since v1.171.0.
 *
 * ## Import
 *
 * Cloud Monitor Service Sls Group can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cms/slsGroup:SlsGroup example <sls_group_name>
 * ```
 */
export class SlsGroup extends pulumi.CustomResource {
    /**
     * Get an existing SlsGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SlsGroupState, opts?: pulumi.CustomResourceOptions): SlsGroup {
        return new SlsGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cms/slsGroup:SlsGroup';

    /**
     * Returns true if the given object is an instance of SlsGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SlsGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SlsGroup.__pulumiType;
    }

    /**
     * The Config of the Sls Group. You can specify up to 25 Config. See `slsGroupConfig` below.
     */
    public readonly slsGroupConfigs!: pulumi.Output<outputs.cms.SlsGroupSlsGroupConfig[]>;
    /**
     * The Description of the Sls Group.
     */
    public readonly slsGroupDescription!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource. The name must be `2` to `32` characters in length, and can contain letters, digits and underscores (_). It must start with a letter.
     */
    public readonly slsGroupName!: pulumi.Output<string>;

    /**
     * Create a SlsGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SlsGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SlsGroupArgs | SlsGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SlsGroupState | undefined;
            resourceInputs["slsGroupConfigs"] = state ? state.slsGroupConfigs : undefined;
            resourceInputs["slsGroupDescription"] = state ? state.slsGroupDescription : undefined;
            resourceInputs["slsGroupName"] = state ? state.slsGroupName : undefined;
        } else {
            const args = argsOrState as SlsGroupArgs | undefined;
            if ((!args || args.slsGroupConfigs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slsGroupConfigs'");
            }
            if ((!args || args.slsGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slsGroupName'");
            }
            resourceInputs["slsGroupConfigs"] = args ? args.slsGroupConfigs : undefined;
            resourceInputs["slsGroupDescription"] = args ? args.slsGroupDescription : undefined;
            resourceInputs["slsGroupName"] = args ? args.slsGroupName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SlsGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SlsGroup resources.
 */
export interface SlsGroupState {
    /**
     * The Config of the Sls Group. You can specify up to 25 Config. See `slsGroupConfig` below.
     */
    slsGroupConfigs?: pulumi.Input<pulumi.Input<inputs.cms.SlsGroupSlsGroupConfig>[]>;
    /**
     * The Description of the Sls Group.
     */
    slsGroupDescription?: pulumi.Input<string>;
    /**
     * The name of the resource. The name must be `2` to `32` characters in length, and can contain letters, digits and underscores (_). It must start with a letter.
     */
    slsGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SlsGroup resource.
 */
export interface SlsGroupArgs {
    /**
     * The Config of the Sls Group. You can specify up to 25 Config. See `slsGroupConfig` below.
     */
    slsGroupConfigs: pulumi.Input<pulumi.Input<inputs.cms.SlsGroupSlsGroupConfig>[]>;
    /**
     * The Description of the Sls Group.
     */
    slsGroupDescription?: pulumi.Input<string>;
    /**
     * The name of the resource. The name must be `2` to `32` characters in length, and can contain letters, digits and underscores (_). It must start with a letter.
     */
    slsGroupName: pulumi.Input<string>;
}
