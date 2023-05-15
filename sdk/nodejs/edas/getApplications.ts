// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of EDAS application in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.82.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const applications = alicloud.edas.getApplications({
 *     ids: ["xxx"],
 *     outputFile: "application.txt",
 * });
 * export const firstApplicationName = applications.then(applications => applications.applications?.[0]?.appName);
 * ```
 */
export function getApplications(args?: GetApplicationsArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:edas/getApplications:getApplications", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplications.
 */
export interface GetApplicationsArgs {
    /**
     * An ids string to filter results by the application id.
     */
    ids?: string[];
    /**
     * A regex string to filter results by the application name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getApplications.
 */
export interface GetApplicationsResult {
    /**
     * A list of applications.
     */
    readonly applications: outputs.edas.GetApplicationsApplication[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of application IDs.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of applications names.
     */
    readonly names: string[];
    readonly outputFile?: string;
}
/**
 * This data source provides a list of EDAS application in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.82.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const applications = alicloud.edas.getApplications({
 *     ids: ["xxx"],
 *     outputFile: "application.txt",
 * });
 * export const firstApplicationName = applications.then(applications => applications.applications?.[0]?.appName);
 * ```
 */
export function getApplicationsOutput(args?: GetApplicationsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationsResult> {
    return pulumi.output(args).apply((a: any) => getApplications(a, opts))
}

/**
 * A collection of arguments for invoking getApplications.
 */
export interface GetApplicationsOutputArgs {
    /**
     * An ids string to filter results by the application id.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by the application name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
