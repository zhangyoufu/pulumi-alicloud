// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides Dcdn Waf Rule available to the user.[What is Waf Rule](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-batchcreatedcdnwafrules)
 *
 * > **NOTE:** Available since v1.201.0.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.dcdn.getWafRules({
 *     ids: [defaultAlicloudDcdnWafRule.id],
 * });
 * export const alicloudDcdnWafRuleExampleId = _default.then(_default => _default.wafRules?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getWafRules(args?: GetWafRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetWafRulesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:dcdn/getWafRules:getWafRules", {
        "ids": args.ids,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
        "queryArgs": args.queryArgs,
    }, opts);
}

/**
 * A collection of arguments for invoking getWafRules.
 */
export interface GetWafRulesArgs {
    /**
     * A list of Waf Rule IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
    /**
     * The query conditions. The value is a string in the JSON format.
     */
    queryArgs?: string;
}

/**
 * A collection of values returned by getWafRules.
 */
export interface GetWafRulesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
    readonly queryArgs?: string;
    /**
     * A list of Waf Rule Entries. Each element contains the following attributes:
     */
    readonly wafRules: outputs.dcdn.GetWafRulesWafRule[];
}
/**
 * This data source provides Dcdn Waf Rule available to the user.[What is Waf Rule](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-batchcreatedcdnwafrules)
 *
 * > **NOTE:** Available since v1.201.0.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.dcdn.getWafRules({
 *     ids: [defaultAlicloudDcdnWafRule.id],
 * });
 * export const alicloudDcdnWafRuleExampleId = _default.then(_default => _default.wafRules?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getWafRulesOutput(args?: GetWafRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWafRulesResult> {
    return pulumi.output(args).apply((a: any) => getWafRules(a, opts))
}

/**
 * A collection of arguments for invoking getWafRules.
 */
export interface GetWafRulesOutputArgs {
    /**
     * A list of Waf Rule IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
    /**
     * The query conditions. The value is a string in the JSON format.
     */
    queryArgs?: pulumi.Input<string>;
}
