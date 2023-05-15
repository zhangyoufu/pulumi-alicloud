// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Apsara File Storage for HDFS Access Groups of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.133.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.dfs.getAccessGroups({
 *     ids: ["example_id"],
 * });
 * export const dfsAccessGroupId1 = ids.then(ids => ids.groups?.[0]?.id);
 * const nameRegex = alicloud.dfs.getAccessGroups({
 *     nameRegex: "^my-AccessGroup",
 * });
 * export const dfsAccessGroupId2 = nameRegex.then(nameRegex => nameRegex.groups?.[0]?.id);
 * ```
 */
export function getAccessGroups(args?: GetAccessGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessGroupsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:dfs/getAccessGroups:getAccessGroups", {
        "ids": args.ids,
        "limit": args.limit,
        "nameRegex": args.nameRegex,
        "orderBy": args.orderBy,
        "orderType": args.orderType,
        "outputFile": args.outputFile,
        "startOffset": args.startOffset,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessGroups.
 */
export interface GetAccessGroupsArgs {
    /**
     * A list of Access Group IDs.
     */
    ids?: string[];
    limit?: number;
    /**
     * A regex string to filter results by Access Group name.
     */
    nameRegex?: string;
    orderBy?: string;
    orderType?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    startOffset?: number;
}

/**
 * A collection of values returned by getAccessGroups.
 */
export interface GetAccessGroupsResult {
    readonly groups: outputs.dfs.GetAccessGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly limit?: number;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly orderBy?: string;
    readonly orderType?: string;
    readonly outputFile?: string;
    readonly startOffset?: number;
}
/**
 * This data source provides the Apsara File Storage for HDFS Access Groups of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.133.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.dfs.getAccessGroups({
 *     ids: ["example_id"],
 * });
 * export const dfsAccessGroupId1 = ids.then(ids => ids.groups?.[0]?.id);
 * const nameRegex = alicloud.dfs.getAccessGroups({
 *     nameRegex: "^my-AccessGroup",
 * });
 * export const dfsAccessGroupId2 = nameRegex.then(nameRegex => nameRegex.groups?.[0]?.id);
 * ```
 */
export function getAccessGroupsOutput(args?: GetAccessGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccessGroupsResult> {
    return pulumi.output(args).apply((a: any) => getAccessGroups(a, opts))
}

/**
 * A collection of arguments for invoking getAccessGroups.
 */
export interface GetAccessGroupsOutputArgs {
    /**
     * A list of Access Group IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    limit?: pulumi.Input<number>;
    /**
     * A regex string to filter results by Access Group name.
     */
    nameRegex?: pulumi.Input<string>;
    orderBy?: pulumi.Input<string>;
    orderType?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    startOffset?: pulumi.Input<number>;
}
