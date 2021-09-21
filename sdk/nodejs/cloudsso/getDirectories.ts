// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Cloud Sso Directories of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.135.0+.
 *
 * > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.cloudsso.getDirectories({
 *     ids: ["example_id"],
 * });
 * export const cloudSsoDirectoryId1 = ids.then(ids => ids.directories[0].id);
 * const nameRegex = alicloud.cloudsso.getDirectories({
 *     nameRegex: "^my-Directory",
 * });
 * export const cloudSsoDirectoryId2 = nameRegex.then(nameRegex => nameRegex.directories[0].id);
 * ```
 */
export function getDirectories(args?: GetDirectoriesArgs, opts?: pulumi.InvokeOptions): Promise<GetDirectoriesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:cloudsso/getDirectories:getDirectories", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDirectories.
 */
export interface GetDirectoriesArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    readonly enableDetails?: boolean;
    /**
     * A list of Directory IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by Directory name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getDirectories.
 */
export interface GetDirectoriesResult {
    readonly directories: outputs.cloudsso.GetDirectoriesDirectory[];
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
}
