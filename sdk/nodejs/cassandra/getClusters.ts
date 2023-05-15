// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The `alicloud.cassandra.getClusters` data source provides a collection of Cassandra clusters available in Alicloud account.
 * Filters support regular expression for the cluster name, ids or tags.
 *
 * > **NOTE:**  Available in 1.88.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const cassandra = alicloud.cassandra.getClusters({
 *     nameRegex: "tf_testAccCassandra",
 * });
 * ```
 */
export function getClusters(args?: GetClustersArgs, opts?: pulumi.InvokeOptions): Promise<GetClustersResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cassandra/getClusters:getClusters", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusters.
 */
export interface GetClustersArgs {
    /**
     * The list of Cassandra cluster ids.
     */
    ids?: string[];
    /**
     * A regex string to apply to the cluster name.
     */
    nameRegex?: string;
    /**
     * The name of file that can save the collection of clusters after running `pulumi preview`.
     */
    outputFile?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getClusters.
 */
export interface GetClustersResult {
    /**
     * A list of Cassandra clusters. Its every element contains the following attributes:
     */
    readonly clusters: outputs.cassandra.GetClustersCluster[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of Cassandra cluster ids.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * The name list of Cassandra clusters.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: {[key: string]: any};
}
/**
 * The `alicloud.cassandra.getClusters` data source provides a collection of Cassandra clusters available in Alicloud account.
 * Filters support regular expression for the cluster name, ids or tags.
 *
 * > **NOTE:**  Available in 1.88.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const cassandra = alicloud.cassandra.getClusters({
 *     nameRegex: "tf_testAccCassandra",
 * });
 * ```
 */
export function getClustersOutput(args?: GetClustersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClustersResult> {
    return pulumi.output(args).apply((a: any) => getClusters(a, opts))
}

/**
 * A collection of arguments for invoking getClusters.
 */
export interface GetClustersOutputArgs {
    /**
     * The list of Cassandra cluster ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to apply to the cluster name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * The name of file that can save the collection of clusters after running `pulumi preview`.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
