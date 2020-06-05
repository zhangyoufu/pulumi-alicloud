// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of Anti-DDoS Advanced instances in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.57.0+ .
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const instanceDdosBgpInstances = pulumi.output(alicloud.ddos.getDdosBgpInstances({
 *     nameRegex: "^ddosbgp",
 * }, { async: true }));
 *
 * export const instance = alicloud_ddosbgp_instances_instance.map(v => v.id);
 * ```
 */
export function getDdosBgpInstances(args?: GetDdosBgpInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetDdosBgpInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:ddos/getDdosBgpInstances:getDdosBgpInstances", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDdosBgpInstances.
 */
export interface GetDdosBgpInstancesArgs {
    /**
     * A list of instance IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by the instance name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getDdosBgpInstances.
 */
export interface GetDdosBgpInstancesResult {
    /**
     * A list of instance IDs.
     */
    readonly ids: string[];
    /**
     * A list of apis. Each element contains the following attributes:
     */
    readonly instances: outputs.ddos.GetDdosBgpInstancesInstance[];
    readonly nameRegex?: string;
    /**
     * A list of instance names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
