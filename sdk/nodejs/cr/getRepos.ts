// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list Container Registry repositories on Alibaba Cloud.
 *
 * > **NOTE:** Available in v1.35.0+
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Declare the data source
 * const myRepos = pulumi.output(alicloud.cr.getRepos({
 *     nameRegex: "my-repos",
 *     outputFile: "my-repo-json",
 * }, { async: true }));
 *
 * export const output = myRepos.repos;
 * ```
 */
export function getRepos(args?: GetReposArgs, opts?: pulumi.InvokeOptions): Promise<GetReposResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:cr/getRepos:getRepos", {
        "enableDetails": args.enableDetails,
        "nameRegex": args.nameRegex,
        "namespace": args.namespace,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepos.
 */
export interface GetReposArgs {
    /**
     * Boolean, false by default, only repository attributes are exported. Set to true if domain list and tags belong to this repository are needed. See `tags` in attributes.
     */
    readonly enableDetails?: boolean;
    /**
     * A regex string to filter results by repository name.
     */
    readonly nameRegex?: string;
    /**
     * Name of container registry namespace where the repositories are located in.
     */
    readonly namespace?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getRepos.
 */
export interface GetReposResult {
    readonly enableDetails?: boolean;
    /**
     * A list of matched Container Registry Repositories. Its element is set to `names`.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of repository names.
     */
    readonly names: string[];
    /**
     * Name of container registry namespace where repo is located.
     */
    readonly namespace?: string;
    readonly outputFile?: string;
    /**
     * A list of matched Container Registry Namespaces. Each element contains the following attributes:
     */
    readonly repos: outputs.cr.GetReposRepo[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
