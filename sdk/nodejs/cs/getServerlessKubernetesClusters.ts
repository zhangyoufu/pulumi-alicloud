// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list Container Service Serverless Kubernetes Clusters on Alibaba Cloud.
 *
 * > **NOTE:** Available in 1.58.0+
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
 * const k8sClusters = pulumi.output(alicloud.cs.getServerlessKubernetesClusters({
 *     nameRegex: "my-first-k8s",
 *     outputFile: "my-first-k8s-json",
 * }, { async: true }));
 *
 * export const output = k8sClusters.clusters;
 * ```
 */
export function getServerlessKubernetesClusters(args?: GetServerlessKubernetesClustersArgs, opts?: pulumi.InvokeOptions): Promise<GetServerlessKubernetesClustersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:cs/getServerlessKubernetesClusters:getServerlessKubernetesClusters", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getServerlessKubernetesClusters.
 */
export interface GetServerlessKubernetesClustersArgs {
    readonly enableDetails?: boolean;
    /**
     * Cluster IDs to filter.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by cluster name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getServerlessKubernetesClusters.
 */
export interface GetServerlessKubernetesClustersResult {
    /**
     * A list of matched Kubernetes clusters. Each element contains the following attributes:
     */
    readonly clusters: outputs.cs.GetServerlessKubernetesClustersCluster[];
    readonly enableDetails?: boolean;
    /**
     * A list of matched Kubernetes clusters' ids.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of matched Kubernetes clusters' names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
