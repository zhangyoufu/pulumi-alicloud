// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides Cloud Monitor Service Metric Rule Black List available to the user.[What is Metric Rule Black List](https://www.alibabacloud.com/help/en/cloudmonitor/latest/describemetricruleblacklist)
 *
 * > **NOTE:** Available in 1.194.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.cms.getMetricRuleBlackLists({
 *     ids: [alicloud_cms_metric_rule_black_lists["default"].id],
 *     category: "ecs",
 *     namespace: "acs_ecs_dashboard",
 * });
 * export const alicloudCmsRuleBlackListExampleId = data.alicloud_cms_metric_rule_black_lists.lists[0].id;
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getMetricRuleBlackLists(args?: GetMetricRuleBlackListsArgs, opts?: pulumi.InvokeOptions): Promise<GetMetricRuleBlackListsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cms/getMetricRuleBlackLists:getMetricRuleBlackLists", {
        "category": args.category,
        "ids": args.ids,
        "metricRuleBlackListId": args.metricRuleBlackListId,
        "nameRegex": args.nameRegex,
        "namespace": args.namespace,
        "order": args.order,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
    }, opts);
}

/**
 * A collection of arguments for invoking getMetricRuleBlackLists.
 */
export interface GetMetricRuleBlackListsArgs {
    /**
     * Cloud service classification. For example, Redis includes kvstore_standard, kvstore_sharding, and kvstore_splitrw.
     */
    category?: string;
    /**
     * A list of Metric Rule Black List IDs.
     */
    ids?: string[];
    /**
     * The first ID of the resource
     */
    metricRuleBlackListId?: string;
    /**
     * A regex string to filter results by Group Metric Rule name.
     */
    nameRegex?: string;
    /**
     * The data namespace of the cloud service.
     */
    namespace?: string;
    order?: number;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
}

/**
 * A collection of values returned by getMetricRuleBlackLists.
 */
export interface GetMetricRuleBlackListsResult {
    /**
     * Cloud service classification. For example, Redis includes kvstore_standard, kvstore_sharding, and kvstore_splitrw.
     */
    readonly category?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of Metric Rule Black List IDs.
     */
    readonly ids: string[];
    /**
     * A list of Metric Rule Black List Entries. Each element contains the following attributes:
     */
    readonly lists: outputs.cms.GetMetricRuleBlackListsList[];
    /**
     * The first ID of the resource
     */
    readonly metricRuleBlackListId?: string;
    readonly nameRegex?: string;
    /**
     * A list of name of Metric Rule Black Lists.
     */
    readonly names: string[];
    /**
     * The data namespace of the cloud service.
     */
    readonly namespace?: string;
    readonly order?: number;
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
}
/**
 * This data source provides Cloud Monitor Service Metric Rule Black List available to the user.[What is Metric Rule Black List](https://www.alibabacloud.com/help/en/cloudmonitor/latest/describemetricruleblacklist)
 *
 * > **NOTE:** Available in 1.194.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.cms.getMetricRuleBlackLists({
 *     ids: [alicloud_cms_metric_rule_black_lists["default"].id],
 *     category: "ecs",
 *     namespace: "acs_ecs_dashboard",
 * });
 * export const alicloudCmsRuleBlackListExampleId = data.alicloud_cms_metric_rule_black_lists.lists[0].id;
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getMetricRuleBlackListsOutput(args?: GetMetricRuleBlackListsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMetricRuleBlackListsResult> {
    return pulumi.output(args).apply((a: any) => getMetricRuleBlackLists(a, opts))
}

/**
 * A collection of arguments for invoking getMetricRuleBlackLists.
 */
export interface GetMetricRuleBlackListsOutputArgs {
    /**
     * Cloud service classification. For example, Redis includes kvstore_standard, kvstore_sharding, and kvstore_splitrw.
     */
    category?: pulumi.Input<string>;
    /**
     * A list of Metric Rule Black List IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The first ID of the resource
     */
    metricRuleBlackListId?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Group Metric Rule name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * The data namespace of the cloud service.
     */
    namespace?: pulumi.Input<string>;
    order?: pulumi.Input<number>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
}
