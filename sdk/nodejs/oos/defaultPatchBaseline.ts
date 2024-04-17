// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Oos Default Patch Baseline resource.
 *
 * For information about Oos Default Patch Baseline and how to use it, see [What is Default Patch Baseline](https://www.alibabacloud.com/help/en/operation-orchestration-service/latest/api-oos-2019-06-01-registerdefaultpatchbaseline).
 *
 * > **NOTE:** Available in v1.203.0+.
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
 * const _default = new alicloud.oos.PatchBaseline("default", {
 *     operationSystem: "Windows",
 *     patchBaselineName: "terraform-example",
 *     description: "terraform-example",
 *     approvalRules: "{\"PatchRules\":[{\"PatchFilterGroup\":[{\"Key\":\"PatchSet\",\"Values\":[\"OS\"]},{\"Key\":\"ProductFamily\",\"Values\":[\"Windows\"]},{\"Key\":\"Product\",\"Values\":[\"Windows 10\",\"Windows 7\"]},{\"Key\":\"Classification\",\"Values\":[\"Security Updates\",\"Updates\",\"Update Rollups\",\"Critical Updates\"]},{\"Key\":\"Severity\",\"Values\":[\"Critical\",\"Important\",\"Moderate\"]}],\"ApproveAfterDays\":7,\"EnableNonSecurity\":true,\"ComplianceLevel\":\"Medium\"}]}",
 * });
 * const defaultDefaultPatchBaseline = new alicloud.oos.DefaultPatchBaseline("default", {patchBaselineName: _default.id});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Oos Default Patch Baseline can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:oos/defaultPatchBaseline:DefaultPatchBaseline example <id>
 * ```
 */
export class DefaultPatchBaseline extends pulumi.CustomResource {
    /**
     * Get an existing DefaultPatchBaseline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefaultPatchBaselineState, opts?: pulumi.CustomResourceOptions): DefaultPatchBaseline {
        return new DefaultPatchBaseline(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:oos/defaultPatchBaseline:DefaultPatchBaseline';

    /**
     * Returns true if the given object is an instance of DefaultPatchBaseline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefaultPatchBaseline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefaultPatchBaseline.__pulumiType;
    }

    /**
     * The ID of the patch baseline.
     */
    public /*out*/ readonly patchBaselineId!: pulumi.Output<string>;
    /**
     * The name of the patch baseline.
     */
    public readonly patchBaselineName!: pulumi.Output<string>;

    /**
     * Create a DefaultPatchBaseline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DefaultPatchBaselineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefaultPatchBaselineArgs | DefaultPatchBaselineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DefaultPatchBaselineState | undefined;
            resourceInputs["patchBaselineId"] = state ? state.patchBaselineId : undefined;
            resourceInputs["patchBaselineName"] = state ? state.patchBaselineName : undefined;
        } else {
            const args = argsOrState as DefaultPatchBaselineArgs | undefined;
            if ((!args || args.patchBaselineName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'patchBaselineName'");
            }
            resourceInputs["patchBaselineName"] = args ? args.patchBaselineName : undefined;
            resourceInputs["patchBaselineId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DefaultPatchBaseline.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefaultPatchBaseline resources.
 */
export interface DefaultPatchBaselineState {
    /**
     * The ID of the patch baseline.
     */
    patchBaselineId?: pulumi.Input<string>;
    /**
     * The name of the patch baseline.
     */
    patchBaselineName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DefaultPatchBaseline resource.
 */
export interface DefaultPatchBaselineArgs {
    /**
     * The name of the patch baseline.
     */
    patchBaselineName: pulumi.Input<string>;
}
