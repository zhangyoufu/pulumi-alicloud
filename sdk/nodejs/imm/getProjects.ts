// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Intelligent Media Management Projects of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.134.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.imm.getProjects({
 *     ids: ["example_id"],
 * });
 * export const immProjectId1 = ids.then(ids => ids.projects?.[0]?.id);
 * ```
 */
export function getProjects(args?: GetProjectsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:imm/getProjects:getProjects", {
        "ids": args.ids,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjects.
 */
export interface GetProjectsArgs {
    /**
     * A list of Project IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getProjects.
 */
export interface GetProjectsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly projects: outputs.imm.GetProjectsProject[];
}
/**
 * This data source provides the Intelligent Media Management Projects of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.134.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.imm.getProjects({
 *     ids: ["example_id"],
 * });
 * export const immProjectId1 = ids.then(ids => ids.projects?.[0]?.id);
 * ```
 */
export function getProjectsOutput(args?: GetProjectsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectsResult> {
    return pulumi.output(args).apply((a: any) => getProjects(a, opts))
}

/**
 * A collection of arguments for invoking getProjects.
 */
export interface GetProjectsOutputArgs {
    /**
     * A list of Project IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
