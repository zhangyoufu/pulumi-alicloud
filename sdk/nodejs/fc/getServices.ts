// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Function Compute services of the current Alibaba Cloud user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const fcServicesDs = pulumi.output(alicloud.fc.getServices({
 *     nameRegex: "sample_fc_service",
 * }, { async: true }));
 *
 * export const firstFcServiceName = fcServicesDs.services[0].name;
 * ```
 */
export function getServices(args?: GetServicesArgs, opts?: pulumi.InvokeOptions): Promise<GetServicesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:fc/getServices:getServices", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getServices.
 */
export interface GetServicesArgs {
    /**
     * - A list of FC services ids.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by FC service name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getServices.
 */
export interface GetServicesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of FC services ids.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of FC services names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of FC services. Each element contains the following attributes:
     */
    readonly services: outputs.fc.GetServicesService[];
}
