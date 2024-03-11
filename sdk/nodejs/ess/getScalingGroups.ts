// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides available scaling group resources.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const scalinggroupsDs = alicloud.ess.getScalingGroups({
 *     ids: [
 *         "scaling_group_id1",
 *         "scaling_group_id2",
 *     ],
 *     nameRegex: "scaling_group_name",
 * });
 * export const firstScalingGroup = scalinggroupsDs.then(scalinggroupsDs => scalinggroupsDs.groups?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getScalingGroups(args?: GetScalingGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetScalingGroupsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ess/getScalingGroups:getScalingGroups", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getScalingGroups.
 */
export interface GetScalingGroupsArgs {
    /**
     * A list of scaling group IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter resulting scaling groups by name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getScalingGroups.
 */
export interface GetScalingGroupsResult {
    /**
     * A list of scaling groups. Each element contains the following attributes:
     */
    readonly groups: outputs.ess.GetScalingGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of scaling group ids.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of scaling group names.
     */
    readonly names: string[];
    readonly outputFile?: string;
}
/**
 * This data source provides available scaling group resources.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const scalinggroupsDs = alicloud.ess.getScalingGroups({
 *     ids: [
 *         "scaling_group_id1",
 *         "scaling_group_id2",
 *     ],
 *     nameRegex: "scaling_group_name",
 * });
 * export const firstScalingGroup = scalinggroupsDs.then(scalinggroupsDs => scalinggroupsDs.groups?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getScalingGroupsOutput(args?: GetScalingGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetScalingGroupsResult> {
    return pulumi.output(args).apply((a: any) => getScalingGroups(a, opts))
}

/**
 * A collection of arguments for invoking getScalingGroups.
 */
export interface GetScalingGroupsOutputArgs {
    /**
     * A list of scaling group IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter resulting scaling groups by name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
