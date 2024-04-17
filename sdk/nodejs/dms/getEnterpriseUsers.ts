// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of DMS Enterprise Users in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.90.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Declare the data source
 * const dmsEnterpriseUsersDs = alicloud.dms.getEnterpriseUsers({
 *     ids: ["uid"],
 *     role: "USER",
 *     status: "NORMAL",
 * });
 * export const firstUserId = dmsEnterpriseUsersDs.then(dmsEnterpriseUsersDs => dmsEnterpriseUsersDs.users?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getEnterpriseUsers(args?: GetEnterpriseUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetEnterpriseUsersResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:dms/getEnterpriseUsers:getEnterpriseUsers", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "role": args.role,
        "searchKey": args.searchKey,
        "status": args.status,
        "tid": args.tid,
    }, opts);
}

/**
 * A collection of arguments for invoking getEnterpriseUsers.
 */
export interface GetEnterpriseUsersArgs {
    /**
     * A list of DMS Enterprise User IDs (UID).
     */
    ids?: string[];
    /**
     * A regex string to filter the results by the DMS Enterprise User nick_name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The role of the user to query.
     */
    role?: string;
    /**
     * The keyword used to query users.
     */
    searchKey?: string;
    /**
     * The status of the user.
     */
    status?: string;
    /**
     * The ID of the tenant in DMS Enterprise.
     */
    tid?: number;
}

/**
 * A collection of values returned by getEnterpriseUsers.
 */
export interface GetEnterpriseUsersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of DMS Enterprise User IDs (UID).
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of DMS Enterprise User names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly role?: string;
    readonly searchKey?: string;
    /**
     * The status of the user.
     */
    readonly status?: string;
    readonly tid?: number;
    /**
     * A list of DMS Enterprise Users. Each element contains the following attributes:
     */
    readonly users: outputs.dms.GetEnterpriseUsersUser[];
}
/**
 * This data source provides a list of DMS Enterprise Users in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.90.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Declare the data source
 * const dmsEnterpriseUsersDs = alicloud.dms.getEnterpriseUsers({
 *     ids: ["uid"],
 *     role: "USER",
 *     status: "NORMAL",
 * });
 * export const firstUserId = dmsEnterpriseUsersDs.then(dmsEnterpriseUsersDs => dmsEnterpriseUsersDs.users?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getEnterpriseUsersOutput(args?: GetEnterpriseUsersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEnterpriseUsersResult> {
    return pulumi.output(args).apply((a: any) => getEnterpriseUsers(a, opts))
}

/**
 * A collection of arguments for invoking getEnterpriseUsers.
 */
export interface GetEnterpriseUsersOutputArgs {
    /**
     * A list of DMS Enterprise User IDs (UID).
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter the results by the DMS Enterprise User nick_name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The role of the user to query.
     */
    role?: pulumi.Input<string>;
    /**
     * The keyword used to query users.
     */
    searchKey?: pulumi.Input<string>;
    /**
     * The status of the user.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the tenant in DMS Enterprise.
     */
    tid?: pulumi.Input<number>;
}
