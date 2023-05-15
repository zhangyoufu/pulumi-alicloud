// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Graph Database Db Instances of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.136.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.graphdatabase.getDbInstances({
 *     ids: ["example_id"],
 * });
 * export const graphDatabaseDbInstanceId1 = ids.then(ids => ids.instances?.[0]?.id);
 * const status = alicloud.graphdatabase.getDbInstances({
 *     ids: ["example_id"],
 *     status: "Running",
 * });
 * export const graphDatabaseDbInstanceId2 = status.then(status => status.instances?.[0]?.id);
 * const description = alicloud.graphdatabase.getDbInstances({
 *     ids: ["example_id"],
 *     dbInstanceDescription: "example_value",
 * });
 * export const graphDatabaseDbInstanceId3 = description.then(description => description.instances?.[0]?.id);
 * ```
 */
export function getDbInstances(args?: GetDbInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetDbInstancesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:graphdatabase/getDbInstances:getDbInstances", {
        "dbInstanceDescription": args.dbInstanceDescription,
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getDbInstances.
 */
export interface GetDbInstancesArgs {
    /**
     * According to the practical example or notes.
     */
    dbInstanceDescription?: string;
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: boolean;
    /**
     * A list of Db Instance IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * Instance status. Value range: `Creating`, `Running`, `Deleting`, `Rebooting`, `DBInstanceClassChanging`, `NetAddressCreating` and `NetAddressDeleting`.
     */
    status?: string;
}

/**
 * A collection of values returned by getDbInstances.
 */
export interface GetDbInstancesResult {
    readonly dbInstanceDescription?: string;
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly instances: outputs.graphdatabase.GetDbInstancesInstance[];
    readonly outputFile?: string;
    readonly status?: string;
}
/**
 * This data source provides the Graph Database Db Instances of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.136.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.graphdatabase.getDbInstances({
 *     ids: ["example_id"],
 * });
 * export const graphDatabaseDbInstanceId1 = ids.then(ids => ids.instances?.[0]?.id);
 * const status = alicloud.graphdatabase.getDbInstances({
 *     ids: ["example_id"],
 *     status: "Running",
 * });
 * export const graphDatabaseDbInstanceId2 = status.then(status => status.instances?.[0]?.id);
 * const description = alicloud.graphdatabase.getDbInstances({
 *     ids: ["example_id"],
 *     dbInstanceDescription: "example_value",
 * });
 * export const graphDatabaseDbInstanceId3 = description.then(description => description.instances?.[0]?.id);
 * ```
 */
export function getDbInstancesOutput(args?: GetDbInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDbInstancesResult> {
    return pulumi.output(args).apply((a: any) => getDbInstances(a, opts))
}

/**
 * A collection of arguments for invoking getDbInstances.
 */
export interface GetDbInstancesOutputArgs {
    /**
     * According to the practical example or notes.
     */
    dbInstanceDescription?: pulumi.Input<string>;
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of Db Instance IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Instance status. Value range: `Creating`, `Running`, `Deleting`, `Rebooting`, `DBInstanceClassChanging`, `NetAddressCreating` and `NetAddressDeleting`.
     */
    status?: pulumi.Input<string>;
}
