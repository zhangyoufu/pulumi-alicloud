// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Rds Accounts of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.120.0+.
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
 * const example = alicloud.rds.getAccounts({
 *     dbInstanceId: "example_value",
 *     nameRegex: "the_resource_name",
 * });
 * export const firstRdsAccountId = example.then(example => example.accounts?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getAccounts(args: GetAccountsArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:rds/getAccounts:getAccounts", {
        "dbInstanceId": args.dbInstanceId,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccounts.
 */
export interface GetAccountsArgs {
    /**
     * The db instance id.
     */
    dbInstanceId: string;
    /**
     * A list of Account IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Account name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The status of the resource.
     */
    status?: string;
}

/**
 * A collection of values returned by getAccounts.
 */
export interface GetAccountsResult {
    readonly accounts: outputs.rds.GetAccountsAccount[];
    readonly dbInstanceId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly status?: string;
}
/**
 * This data source provides the Rds Accounts of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.120.0+.
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
 * const example = alicloud.rds.getAccounts({
 *     dbInstanceId: "example_value",
 *     nameRegex: "the_resource_name",
 * });
 * export const firstRdsAccountId = example.then(example => example.accounts?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getAccountsOutput(args: GetAccountsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountsResult> {
    return pulumi.output(args).apply((a: any) => getAccounts(a, opts))
}

/**
 * A collection of arguments for invoking getAccounts.
 */
export interface GetAccountsOutputArgs {
    /**
     * The db instance id.
     */
    dbInstanceId: pulumi.Input<string>;
    /**
     * A list of Account IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Account name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
}
