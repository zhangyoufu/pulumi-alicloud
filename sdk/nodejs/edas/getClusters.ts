// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of EDAS clusters in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.82.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const clusters = alicloud.edas.getClusters({
 *     logicalRegionId: "cn-shenzhen:xxx",
 *     ids: ["addfs-dfsasd"],
 *     outputFile: "clusters.txt",
 * });
 * export const firstClusterName = data.alicloud_alikafka_consumer_groups.clusters.clusters[0].cluster_name;
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getClusters(args: GetClustersArgs, opts?: pulumi.InvokeOptions): Promise<GetClustersResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:edas/getClusters:getClusters", {
        "ids": args.ids,
        "logicalRegionId": args.logicalRegionId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusters.
 */
export interface GetClustersArgs {
    /**
     * An ids string to filter results by the cluster id.
     */
    ids?: string[];
    /**
     * ID of the namespace in EDAS.
     */
    logicalRegionId: string;
    /**
     * A regex string to filter results by the cluster name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getClusters.
 */
export interface GetClustersResult {
    /**
     * A list of clusters.
     */
    readonly clusters: outputs.edas.GetClustersCluster[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of cluster IDs.
     */
    readonly ids: string[];
    readonly logicalRegionId: string;
    readonly nameRegex?: string;
    /**
     * A list of cluster names.
     */
    readonly names: string[];
    readonly outputFile?: string;
}
/**
 * This data source provides a list of EDAS clusters in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.82.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const clusters = alicloud.edas.getClusters({
 *     logicalRegionId: "cn-shenzhen:xxx",
 *     ids: ["addfs-dfsasd"],
 *     outputFile: "clusters.txt",
 * });
 * export const firstClusterName = data.alicloud_alikafka_consumer_groups.clusters.clusters[0].cluster_name;
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getClustersOutput(args: GetClustersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClustersResult> {
    return pulumi.output(args).apply((a: any) => getClusters(a, opts))
}

/**
 * A collection of arguments for invoking getClusters.
 */
export interface GetClustersOutputArgs {
    /**
     * An ids string to filter results by the cluster id.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the namespace in EDAS.
     */
    logicalRegionId: pulumi.Input<string>;
    /**
     * A regex string to filter results by the cluster name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
