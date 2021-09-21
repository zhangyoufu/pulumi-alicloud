// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Sddp Instances of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.136.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.sddp.getInstances({});
 * export const sddpInstanceId = _default.then(_default => _default.instances[0]);
 * ```
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:sddp/getInstances:getInstances", {
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instances: outputs.sddp.GetInstancesInstance[];
    readonly outputFile?: string;
}
