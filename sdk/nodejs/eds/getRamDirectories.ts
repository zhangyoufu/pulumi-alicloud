// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Ecd Ram Directories of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.174.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.eds.getRamDirectories({
 *     ids: ["example_id"],
 * });
 * export const ecdRamDirectoryId1 = ids.then(ids => ids.directories?[0]?.id);
 * const nameRegex = alicloud.eds.getRamDirectories({
 *     nameRegex: "^my-RamDirectory",
 * });
 * export const ecdRamDirectoryId2 = nameRegex.then(nameRegex => nameRegex.directories?[0]?.id);
 * ```
 */
export function getRamDirectories(args?: GetRamDirectoriesArgs, opts?: pulumi.InvokeOptions): Promise<GetRamDirectoriesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:eds/getRamDirectories:getRamDirectories", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getRamDirectories.
 */
export interface GetRamDirectoriesArgs {
    /**
     * A list of Ram Directory IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Ram Directory name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The status of directory.
     */
    status?: string;
}

/**
 * A collection of values returned by getRamDirectories.
 */
export interface GetRamDirectoriesResult {
    readonly directories: outputs.eds.GetRamDirectoriesDirectory[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly status?: string;
}

export function getRamDirectoriesOutput(args?: GetRamDirectoriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRamDirectoriesResult> {
    return pulumi.output(args).apply(a => getRamDirectories(a, opts))
}

/**
 * A collection of arguments for invoking getRamDirectories.
 */
export interface GetRamDirectoriesOutputArgs {
    /**
     * A list of Ram Directory IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Ram Directory name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The status of directory.
     */
    status?: pulumi.Input<string>;
}
