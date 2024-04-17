// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The `alicloud.emr.getMainVersions` data source provides a collection of emr
 * main versions available in Alibaba Cloud account when create a emr cluster.
 *
 * > **NOTE:** Available in 1.59.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.emr.getMainVersions({
 *     emrVersion: "EMR-3.22.0",
 *     clusterTypes: [
 *         "HADOOP",
 *         "ZOOKEEPER",
 *     ],
 * });
 * export const firstMainVersion = _default.then(_default => _default.mainVersions?.[0]?.emrVersion);
 * export const thisClusterTypes = _default.then(_default => _default.mainVersions?.[0]?.clusterTypes);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getMainVersions(args?: GetMainVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetMainVersionsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:emr/getMainVersions:getMainVersions", {
        "clusterTypes": args.clusterTypes,
        "emrVersion": args.emrVersion,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getMainVersions.
 */
export interface GetMainVersionsArgs {
    /**
     * The supported clusterType of this emr version.
     * Possible values may be any one or combination of these: ["HADOOP", "DRUID", "KAFKA", "ZOOKEEPER", "FLINK", "CLICKHOUSE"]
     */
    clusterTypes?: string[];
    /**
     * The version of the emr cluster instance. Possible values: `EMR-4.0.0`, `EMR-3.23.0`, `EMR-3.22.0`.
     */
    emrVersion?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getMainVersions.
 */
export interface GetMainVersionsResult {
    readonly clusterTypes?: string[];
    /**
     * The version of the emr cluster instance.
     */
    readonly emrVersion?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of emr instance types IDs.
     */
    readonly ids: string[];
    /**
     * A list of versions of the emr cluster instance. Each element contains the following attributes:
     */
    readonly mainVersions: outputs.emr.GetMainVersionsMainVersion[];
    readonly outputFile?: string;
}
/**
 * The `alicloud.emr.getMainVersions` data source provides a collection of emr
 * main versions available in Alibaba Cloud account when create a emr cluster.
 *
 * > **NOTE:** Available in 1.59.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.emr.getMainVersions({
 *     emrVersion: "EMR-3.22.0",
 *     clusterTypes: [
 *         "HADOOP",
 *         "ZOOKEEPER",
 *     ],
 * });
 * export const firstMainVersion = _default.then(_default => _default.mainVersions?.[0]?.emrVersion);
 * export const thisClusterTypes = _default.then(_default => _default.mainVersions?.[0]?.clusterTypes);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getMainVersionsOutput(args?: GetMainVersionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMainVersionsResult> {
    return pulumi.output(args).apply((a: any) => getMainVersions(a, opts))
}

/**
 * A collection of arguments for invoking getMainVersions.
 */
export interface GetMainVersionsOutputArgs {
    /**
     * The supported clusterType of this emr version.
     * Possible values may be any one or combination of these: ["HADOOP", "DRUID", "KAFKA", "ZOOKEEPER", "FLINK", "CLICKHOUSE"]
     */
    clusterTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The version of the emr cluster instance. Possible values: `EMR-4.0.0`, `EMR-3.23.0`, `EMR-3.22.0`.
     */
    emrVersion?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
