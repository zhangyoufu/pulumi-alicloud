// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides available lifecycle hook resources. 
 *
 * > **NOTE:** Available in 1.72.0+
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ds = pulumi.output(alicloud.ess.getLifecycleHooks({
 *     nameRegex: "lifecyclehookName",
 *     scalingGroupId: "scalingGroupId",
 * }, { async: true }));
 *
 * export const firstLifecycleHook = ds.hooks[0].id;
 * ```
 */
export function getLifecycleHooks(args?: GetLifecycleHooksArgs, opts?: pulumi.InvokeOptions): Promise<GetLifecycleHooksResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:ess/getLifecycleHooks:getLifecycleHooks", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "scalingGroupId": args.scalingGroupId,
    }, opts);
}

/**
 * A collection of arguments for invoking getLifecycleHooks.
 */
export interface GetLifecycleHooksArgs {
    /**
     * A list of lifecycle hook IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter resulting lifecycle hook by name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * Scaling group id the lifecycle hooks belong to.
     */
    readonly scalingGroupId?: string;
}

/**
 * A collection of values returned by getLifecycleHooks.
 */
export interface GetLifecycleHooksResult {
    /**
     * A list of lifecycle hooks. Each element contains the following attributes:
     */
    readonly hooks: outputs.ess.GetLifecycleHooksHook[];
    /**
     * A list of lifecycle hook ids.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of lifecycle hook names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * ID of the scaling group.
     */
    readonly scalingGroupId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
