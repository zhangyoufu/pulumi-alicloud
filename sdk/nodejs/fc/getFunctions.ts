// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Function Compute functions of the current Alibaba Cloud user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const functionsDs = pulumi.output(alicloud.fc.getFunctions({
 *     nameRegex: "sample_fc_function",
 *     serviceName: "sample_service",
 * }, { async: true }));
 *
 * export const firstFcFunctionName = functionsDs.functions[0].name;
 * ```
 */
export function getFunctions(args: GetFunctionsArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:fc/getFunctions:getFunctions", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getFunctions.
 */
export interface GetFunctionsArgs {
    /**
     * - A list of functions ids.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by function name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * Name of the service that contains the functions to find.
     */
    readonly serviceName: string;
}

/**
 * A collection of values returned by getFunctions.
 */
export interface GetFunctionsResult {
    /**
     * A list of functions. Each element contains the following attributes:
     */
    readonly functions: outputs.fc.GetFunctionsFunction[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of functions ids.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of functions names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly serviceName: string;
}
