// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Adb DBCluster Lake Versions of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.190.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.adb.getDBClusterLakeVersions({
 *     ids: ["example_id"],
 * });
 * export const adbDbClusterLakeVersionId1 = ids.then(ids => ids.versions?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDBClusterLakeVersions(args?: GetDBClusterLakeVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetDBClusterLakeVersionsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:adb/getDBClusterLakeVersions:getDBClusterLakeVersions", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
        "resourceGroupId": args.resourceGroupId,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getDBClusterLakeVersions.
 */
export interface GetDBClusterLakeVersionsArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: boolean;
    /**
     * A list of DBCluster IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: string;
    /**
     * The status of the resource.
     */
    status?: string;
}

/**
 * A collection of values returned by getDBClusterLakeVersions.
 */
export interface GetDBClusterLakeVersionsResult {
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
    readonly resourceGroupId?: string;
    readonly status?: string;
    readonly versions: outputs.adb.GetDBClusterLakeVersionsVersion[];
}
/**
 * This data source provides the Adb DBCluster Lake Versions of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.190.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.adb.getDBClusterLakeVersions({
 *     ids: ["example_id"],
 * });
 * export const adbDbClusterLakeVersionId1 = ids.then(ids => ids.versions?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDBClusterLakeVersionsOutput(args?: GetDBClusterLakeVersionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDBClusterLakeVersionsResult> {
    return pulumi.output(args).apply((a: any) => getDBClusterLakeVersions(a, opts))
}

/**
 * A collection of arguments for invoking getDBClusterLakeVersions.
 */
export interface GetDBClusterLakeVersionsOutputArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of DBCluster IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
}
