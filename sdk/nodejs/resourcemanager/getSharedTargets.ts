// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Resource Manager Shared Targets of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available since v1.111.0.
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
 * const name = config.get("name") || "tf-example";
 * const default = alicloud.resourcemanager.getAccounts({});
 * const defaultResourceShare = new alicloud.resourcemanager.ResourceShare("default", {resourceShareName: name});
 * const defaultSharedTarget = new alicloud.resourcemanager.SharedTarget("default", {
 *     resourceShareId: defaultResourceShare.id,
 *     targetId: _default.then(_default => _default.ids?.[0]),
 * });
 * const ids = alicloud.resourcemanager.getSharedTargetsOutput({
 *     ids: [defaultSharedTarget.targetId],
 * });
 * export const firstResourceManagerSharedTargetId = ids.apply(ids => ids.targets?.[0]?.id);
 * const resourceShareId = alicloud.resourcemanager.getSharedTargetsOutput({
 *     resourceShareId: defaultSharedTarget.resourceShareId,
 * });
 * export const secondResourceManagerSharedTargetId = resourceShareId.apply(resourceShareId => resourceShareId.targets?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSharedTargets(args?: GetSharedTargetsArgs, opts?: pulumi.InvokeOptions): Promise<GetSharedTargetsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:resourcemanager/getSharedTargets:getSharedTargets", {
        "ids": args.ids,
        "outputFile": args.outputFile,
        "resourceShareId": args.resourceShareId,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getSharedTargets.
 */
export interface GetSharedTargetsArgs {
    /**
     * A list of Shared Target IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The resource share ID of resource manager.
     */
    resourceShareId?: string;
    /**
     * The status of share resource. Valid values: `Associated`, `Associating`, `Disassociated`, `Disassociating` and `Failed`.
     */
    status?: string;
}

/**
 * A collection of values returned by getSharedTargets.
 */
export interface GetSharedTargetsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    /**
     * The resource shared ID of resource manager.
     */
    readonly resourceShareId?: string;
    /**
     * The status of shared target.
     */
    readonly status?: string;
    /**
     * A list of Resource Manager Shared Targets. Each element contains the following attributes:
     */
    readonly targets: outputs.resourcemanager.GetSharedTargetsTarget[];
}
/**
 * This data source provides the Resource Manager Shared Targets of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available since v1.111.0.
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
 * const name = config.get("name") || "tf-example";
 * const default = alicloud.resourcemanager.getAccounts({});
 * const defaultResourceShare = new alicloud.resourcemanager.ResourceShare("default", {resourceShareName: name});
 * const defaultSharedTarget = new alicloud.resourcemanager.SharedTarget("default", {
 *     resourceShareId: defaultResourceShare.id,
 *     targetId: _default.then(_default => _default.ids?.[0]),
 * });
 * const ids = alicloud.resourcemanager.getSharedTargetsOutput({
 *     ids: [defaultSharedTarget.targetId],
 * });
 * export const firstResourceManagerSharedTargetId = ids.apply(ids => ids.targets?.[0]?.id);
 * const resourceShareId = alicloud.resourcemanager.getSharedTargetsOutput({
 *     resourceShareId: defaultSharedTarget.resourceShareId,
 * });
 * export const secondResourceManagerSharedTargetId = resourceShareId.apply(resourceShareId => resourceShareId.targets?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSharedTargetsOutput(args?: GetSharedTargetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSharedTargetsResult> {
    return pulumi.output(args).apply((a: any) => getSharedTargets(a, opts))
}

/**
 * A collection of arguments for invoking getSharedTargets.
 */
export interface GetSharedTargetsOutputArgs {
    /**
     * A list of Shared Target IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The resource share ID of resource manager.
     */
    resourceShareId?: pulumi.Input<string>;
    /**
     * The status of share resource. Valid values: `Associated`, `Associating`, `Disassociated`, `Disassociating` and `Failed`.
     */
    status?: pulumi.Input<string>;
}
