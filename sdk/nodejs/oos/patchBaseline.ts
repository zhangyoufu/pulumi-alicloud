// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a OOS Patch Baseline resource.
 *
 * For information about OOS Patch Baseline and how to use it, see [What is Patch Baseline](https://www.alibabacloud.com/help/en/doc-detail/268700.html).
 *
 * > **NOTE:** Available in v1.146.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.oos.PatchBaseline("example", {
 *     approvalRules: "example_value",
 *     operationSystem: "Windows",
 *     patchBaselineName: "my-PatchBaseline",
 * });
 * ```
 *
 * ## Import
 *
 * OOS Patch Baseline can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:oos/patchBaseline:PatchBaseline example <patch_baseline_name>
 * ```
 */
export class PatchBaseline extends pulumi.CustomResource {
    /**
     * Get an existing PatchBaseline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PatchBaselineState, opts?: pulumi.CustomResourceOptions): PatchBaseline {
        return new PatchBaseline(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:oos/patchBaseline:PatchBaseline';

    /**
     * Returns true if the given object is an instance of PatchBaseline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PatchBaseline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PatchBaseline.__pulumiType;
    }

    /**
     * Accept the rules. This value follows the json format. For more details, see the [description of ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/doc-detail/311002.html).
     */
    public readonly approvalRules!: pulumi.Output<string>;
    /**
     * Patches baseline description information.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Operating system type. Valid values: `AliyunLinux`, `Anolis`, `CentOS`, `Debian`, `RedhatEnterpriseLinux`, `Ubuntu`, `Windows`.
     */
    public readonly operationSystem!: pulumi.Output<string>;
    /**
     * The name of the patch baseline.
     */
    public readonly patchBaselineName!: pulumi.Output<string>;

    /**
     * Create a PatchBaseline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PatchBaselineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PatchBaselineArgs | PatchBaselineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PatchBaselineState | undefined;
            resourceInputs["approvalRules"] = state ? state.approvalRules : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["operationSystem"] = state ? state.operationSystem : undefined;
            resourceInputs["patchBaselineName"] = state ? state.patchBaselineName : undefined;
        } else {
            const args = argsOrState as PatchBaselineArgs | undefined;
            if ((!args || args.approvalRules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'approvalRules'");
            }
            if ((!args || args.operationSystem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'operationSystem'");
            }
            if ((!args || args.patchBaselineName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'patchBaselineName'");
            }
            resourceInputs["approvalRules"] = args ? args.approvalRules : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["operationSystem"] = args ? args.operationSystem : undefined;
            resourceInputs["patchBaselineName"] = args ? args.patchBaselineName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PatchBaseline.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PatchBaseline resources.
 */
export interface PatchBaselineState {
    /**
     * Accept the rules. This value follows the json format. For more details, see the [description of ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/doc-detail/311002.html).
     */
    approvalRules?: pulumi.Input<string>;
    /**
     * Patches baseline description information.
     */
    description?: pulumi.Input<string>;
    /**
     * Operating system type. Valid values: `AliyunLinux`, `Anolis`, `CentOS`, `Debian`, `RedhatEnterpriseLinux`, `Ubuntu`, `Windows`.
     */
    operationSystem?: pulumi.Input<string>;
    /**
     * The name of the patch baseline.
     */
    patchBaselineName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PatchBaseline resource.
 */
export interface PatchBaselineArgs {
    /**
     * Accept the rules. This value follows the json format. For more details, see the [description of ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/doc-detail/311002.html).
     */
    approvalRules: pulumi.Input<string>;
    /**
     * Patches baseline description information.
     */
    description?: pulumi.Input<string>;
    /**
     * Operating system type. Valid values: `AliyunLinux`, `Anolis`, `CentOS`, `Debian`, `RedhatEnterpriseLinux`, `Ubuntu`, `Windows`.
     */
    operationSystem: pulumi.Input<string>;
    /**
     * The name of the patch baseline.
     */
    patchBaselineName: pulumi.Input<string>;
}
