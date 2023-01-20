// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Resource Manager Accounts of the current Alibaba Cloud user.
 *
 * > **NOTE:**  Available in 1.86.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.resourcemanager.getAccounts({});
 * export const firstAccountId = _default.then(_default => _default.accounts?.[0]?.id);
 * ```
 */
export function getAccounts(args?: GetAccountsArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:resourcemanager/getAccounts:getAccounts", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccounts.
 */
export interface GetAccountsArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: boolean;
    /**
     * A list of account IDs.
     */
    ids?: string[];
    outputFile?: string;
    /**
     * The status of account, valid values: `CreateCancelled`, `CreateExpired`, `CreateFailed`, `CreateSuccess`, `CreateVerifying`, `InviteSuccess`, `PromoteCancelled`, `PromoteExpired`, `PromoteFailed`, `PromoteSuccess`, and `PromoteVerifying`.
     */
    status?: string;
}

/**
 * A collection of values returned by getAccounts.
 */
export interface GetAccountsResult {
    /**
     * A list of accounts. Each element contains the following attributes:
     */
    readonly accounts: outputs.resourcemanager.GetAccountsAccount[];
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of account IDs.
     */
    readonly ids: string[];
    readonly outputFile?: string;
    /**
     * The status of the member account.
     */
    readonly status?: string;
}
/**
 * This data source provides the Resource Manager Accounts of the current Alibaba Cloud user.
 *
 * > **NOTE:**  Available in 1.86.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.resourcemanager.getAccounts({});
 * export const firstAccountId = _default.then(_default => _default.accounts?.[0]?.id);
 * ```
 */
export function getAccountsOutput(args?: GetAccountsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountsResult> {
    return pulumi.output(args).apply((a: any) => getAccounts(a, opts))
}

/**
 * A collection of arguments for invoking getAccounts.
 */
export interface GetAccountsOutputArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of account IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    outputFile?: pulumi.Input<string>;
    /**
     * The status of account, valid values: `CreateCancelled`, `CreateExpired`, `CreateFailed`, `CreateSuccess`, `CreateVerifying`, `InviteSuccess`, `PromoteCancelled`, `PromoteExpired`, `PromoteFailed`, `PromoteSuccess`, and `PromoteVerifying`.
     */
    status?: pulumi.Input<string>;
}
