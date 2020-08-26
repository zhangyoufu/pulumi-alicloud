// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This data source provides information about the current account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const current = pulumi.output(alicloud.getAccount({ async: true }));
 *
 * export const currentAccountId = current.id;
 * ```
 */
export function getAccount(opts?: pulumi.InvokeOptions): Promise<GetAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:index/getAccount:getAccount", {
    }, opts);
}

/**
 * A collection of values returned by getAccount.
 */
export interface GetAccountResult {
    /**
     * Account ID (e.g. "1239306421830812"). It can be used to construct an ARN.
     */
    readonly id: string;
}
