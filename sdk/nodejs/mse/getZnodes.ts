// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Mse Znodes of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.162.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.mse.getZnodes({
 *     clusterId: "example_value",
 *     path: "/",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 * });
 * export const mseZnodeId1 = ids.then(ids => ids.znodes?.[0]?.id);
 * const nameRegex = alicloud.mse.getZnodes({
 *     path: "/",
 *     clusterId: "example_value",
 *     nameRegex: "^my-Znode",
 * });
 * export const mseZnodeId2 = nameRegex.then(nameRegex => nameRegex.znodes?.[0]?.id);
 * ```
 */
export function getZnodes(args: GetZnodesArgs, opts?: pulumi.InvokeOptions): Promise<GetZnodesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:mse/getZnodes:getZnodes", {
        "acceptLanguage": args.acceptLanguage,
        "clusterId": args.clusterId,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "path": args.path,
    }, opts);
}

/**
 * A collection of arguments for invoking getZnodes.
 */
export interface GetZnodesArgs {
    /**
     * The language type of the returned information. Valid values: `zh` or `en`.
     */
    acceptLanguage?: string;
    /**
     * The ID of the Cluster.
     */
    clusterId: string;
    /**
     * A list of Znode IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Znode name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The Node path.
     */
    path: string;
}

/**
 * A collection of values returned by getZnodes.
 */
export interface GetZnodesResult {
    readonly acceptLanguage?: string;
    readonly clusterId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly path: string;
    readonly znodes: outputs.mse.GetZnodesZnode[];
}
/**
 * This data source provides the Mse Znodes of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.162.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.mse.getZnodes({
 *     clusterId: "example_value",
 *     path: "/",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 * });
 * export const mseZnodeId1 = ids.then(ids => ids.znodes?.[0]?.id);
 * const nameRegex = alicloud.mse.getZnodes({
 *     path: "/",
 *     clusterId: "example_value",
 *     nameRegex: "^my-Znode",
 * });
 * export const mseZnodeId2 = nameRegex.then(nameRegex => nameRegex.znodes?.[0]?.id);
 * ```
 */
export function getZnodesOutput(args: GetZnodesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZnodesResult> {
    return pulumi.output(args).apply((a: any) => getZnodes(a, opts))
}

/**
 * A collection of arguments for invoking getZnodes.
 */
export interface GetZnodesOutputArgs {
    /**
     * The language type of the returned information. Valid values: `zh` or `en`.
     */
    acceptLanguage?: pulumi.Input<string>;
    /**
     * The ID of the Cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * A list of Znode IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Znode name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The Node path.
     */
    path: pulumi.Input<string>;
}
