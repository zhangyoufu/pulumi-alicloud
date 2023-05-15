// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of DMS User Tenants in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.161.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.dms.getUserTenants({
 *     status: "ACTIVE",
 * });
 * export const tid = _default.then(_default => _default.ids?.[0]);
 * ```
 */
export function getUserTenants(args?: GetUserTenantsArgs, opts?: pulumi.InvokeOptions): Promise<GetUserTenantsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:dms/getUserTenants:getUserTenants", {
        "ids": args.ids,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserTenants.
 */
export interface GetUserTenantsArgs {
    /**
     * A list of DMS User Tenant IDs (TID).
     */
    ids?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The status of the user tenant.
     */
    status?: string;
}

/**
 * A collection of values returned by getUserTenants.
 */
export interface GetUserTenantsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of DMS User Tenant IDs (UID).
     */
    readonly ids: string[];
    /**
     * A list of DMS User Tenant names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The status of the user tenant.
     */
    readonly status?: string;
    /**
     * A list of DMS User Tenants. Each element contains the following attributes:
     */
    readonly tenants: outputs.dms.GetUserTenantsTenant[];
}
/**
 * This data source provides a list of DMS User Tenants in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.161.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.dms.getUserTenants({
 *     status: "ACTIVE",
 * });
 * export const tid = _default.then(_default => _default.ids?.[0]);
 * ```
 */
export function getUserTenantsOutput(args?: GetUserTenantsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserTenantsResult> {
    return pulumi.output(args).apply((a: any) => getUserTenants(a, opts))
}

/**
 * A collection of arguments for invoking getUserTenants.
 */
export interface GetUserTenantsOutputArgs {
    /**
     * A list of DMS User Tenant IDs (TID).
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the user tenant.
     */
    status?: pulumi.Input<string>;
}
