// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list Container Service Managed Kubernetes Clusters on Alibaba Cloud.
 *
 * > **NOTE:** Available in v1.35.0+
 *
 * > **NOTE:** From version 1.177.0+, We supported batch export of clusters' kube config information by `kubeConfigFilePrefix`.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Declare the data source
 * const k8sClusters = alicloud.cs.getManagedKubernetesClusters({
 *     nameRegex: "my-first-k8s",
 *     outputFile: "my-first-k8s-json",
 *     kubeConfigFilePrefix: "~/.kube/managed",
 * });
 * export const output = k8sClusters.then(k8sClusters => k8sClusters.clusters);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getManagedKubernetesClusters(args?: GetManagedKubernetesClustersArgs, opts?: pulumi.InvokeOptions): Promise<GetManagedKubernetesClustersResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cs/getManagedKubernetesClusters:getManagedKubernetesClusters", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "kubeConfigFilePrefix": args.kubeConfigFilePrefix,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getManagedKubernetesClusters.
 */
export interface GetManagedKubernetesClustersArgs {
    enableDetails?: boolean;
    /**
     * Cluster IDs to filter.
     */
    ids?: string[];
    /**
     * The path prefix of kube config. You could store kube config in a specified directory by specifying this field, like `~/.kube/managed`, then it will be named with `~/.kube/managed-clusterID-kubeconfig`. From version 1.187.0+, kubeConfig will not export kubeConfig if this field is not set.
     */
    kubeConfigFilePrefix?: string;
    /**
     * A regex string to filter results by cluster name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getManagedKubernetesClusters.
 */
export interface GetManagedKubernetesClustersResult {
    /**
     * A list of matched Kubernetes clusters. Each element contains the following attributes:
     */
    readonly clusters: outputs.cs.GetManagedKubernetesClustersCluster[];
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of matched Kubernetes clusters' ids.
     */
    readonly ids: string[];
    readonly kubeConfigFilePrefix?: string;
    readonly nameRegex?: string;
    /**
     * A list of matched Kubernetes clusters' names.
     */
    readonly names: string[];
    readonly outputFile?: string;
}
/**
 * This data source provides a list Container Service Managed Kubernetes Clusters on Alibaba Cloud.
 *
 * > **NOTE:** Available in v1.35.0+
 *
 * > **NOTE:** From version 1.177.0+, We supported batch export of clusters' kube config information by `kubeConfigFilePrefix`.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Declare the data source
 * const k8sClusters = alicloud.cs.getManagedKubernetesClusters({
 *     nameRegex: "my-first-k8s",
 *     outputFile: "my-first-k8s-json",
 *     kubeConfigFilePrefix: "~/.kube/managed",
 * });
 * export const output = k8sClusters.then(k8sClusters => k8sClusters.clusters);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getManagedKubernetesClustersOutput(args?: GetManagedKubernetesClustersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetManagedKubernetesClustersResult> {
    return pulumi.output(args).apply((a: any) => getManagedKubernetesClusters(a, opts))
}

/**
 * A collection of arguments for invoking getManagedKubernetesClusters.
 */
export interface GetManagedKubernetesClustersOutputArgs {
    enableDetails?: pulumi.Input<boolean>;
    /**
     * Cluster IDs to filter.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The path prefix of kube config. You could store kube config in a specified directory by specifying this field, like `~/.kube/managed`, then it will be named with `~/.kube/managed-clusterID-kubeconfig`. From version 1.187.0+, kubeConfig will not export kubeConfig if this field is not set.
     */
    kubeConfigFilePrefix?: pulumi.Input<string>;
    /**
     * A regex string to filter results by cluster name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
