// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Resource Manager Shared Targets of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.111.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.resourcemanager.getSharedTargets({
 *     ids: ["15681091********"],
 * });
 * export const firstResourceManagerSharedTargetId = example.then(example => example.targets[0].id);
 * ```
 */
export function getSharedTargets(args?: GetSharedTargetsArgs, opts?: pulumi.InvokeOptions): Promise<GetSharedTargetsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
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
    readonly ids?: string[];
    readonly outputFile?: string;
    /**
     * The resource shared ID of resource manager.
     */
    readonly resourceShareId?: string;
    /**
     * The status of shared target.
     */
    readonly status?: string;
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
    readonly resourceShareId?: string;
    readonly status?: string;
    readonly targets: outputs.resourcemanager.GetSharedTargetsTarget[];
}
