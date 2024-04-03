// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides Quotas Template Applications available to the user.[What is Template Applications](https://www.alibabacloud.com/help/en/quota-center/developer-reference/api-quotas-2020-05-10-createquotaapplicationsfortemplate)
 *
 * > **NOTE:** Available since v1.214.0.
 */
export function getTemplateApplications(args?: GetTemplateApplicationsArgs, opts?: pulumi.InvokeOptions): Promise<GetTemplateApplicationsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:quotas/getTemplateApplications:getTemplateApplications", {
        "batchQuotaApplicationId": args.batchQuotaApplicationId,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "productCode": args.productCode,
        "quotaActionCode": args.quotaActionCode,
        "quotaCategory": args.quotaCategory,
    }, opts);
}

/**
 * A collection of arguments for invoking getTemplateApplications.
 */
export interface GetTemplateApplicationsArgs {
    /**
     * The ID of the quota application batch.
     */
    batchQuotaApplicationId?: string;
    /**
     * A list of Template Applications IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * Cloud service name abbreviation.> For more information about cloud services that support quota centers, see Cloud services that support quota centers.
     */
    productCode?: string;
    /**
     * The quota ID.
     */
    quotaActionCode?: string;
    /**
     * The quota type. Value: `CommonQuota`, `FlowControl` and `WhiteListLabel`.
     */
    quotaCategory?: string;
}

/**
 * A collection of values returned by getTemplateApplications.
 */
export interface GetTemplateApplicationsResult {
    /**
     * A list of Template Applications Entries. Each element contains the following attributes:
     */
    readonly applications: outputs.quotas.GetTemplateApplicationsApplication[];
    /**
     * The ID of the quota application batch.
     */
    readonly batchQuotaApplicationId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of Template Applications IDs.
     */
    readonly ids: string[];
    readonly outputFile?: string;
    /**
     * Cloud service name abbreviation.> For more information about cloud services that support quota centers, see Cloud services that support quota centers.
     */
    readonly productCode?: string;
    /**
     * The quota ID.
     */
    readonly quotaActionCode?: string;
    /**
     * The quota type. Value:-CommonQuota (default): Generic quota.-FlowControl:API rate quota.-WhiteListLabel: Equity quota.
     */
    readonly quotaCategory?: string;
}
/**
 * This data source provides Quotas Template Applications available to the user.[What is Template Applications](https://www.alibabacloud.com/help/en/quota-center/developer-reference/api-quotas-2020-05-10-createquotaapplicationsfortemplate)
 *
 * > **NOTE:** Available since v1.214.0.
 */
export function getTemplateApplicationsOutput(args?: GetTemplateApplicationsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTemplateApplicationsResult> {
    return pulumi.output(args).apply((a: any) => getTemplateApplications(a, opts))
}

/**
 * A collection of arguments for invoking getTemplateApplications.
 */
export interface GetTemplateApplicationsOutputArgs {
    /**
     * The ID of the quota application batch.
     */
    batchQuotaApplicationId?: pulumi.Input<string>;
    /**
     * A list of Template Applications IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Cloud service name abbreviation.> For more information about cloud services that support quota centers, see Cloud services that support quota centers.
     */
    productCode?: pulumi.Input<string>;
    /**
     * The quota ID.
     */
    quotaActionCode?: pulumi.Input<string>;
    /**
     * The quota type. Value: `CommonQuota`, `FlowControl` and `WhiteListLabel`.
     */
    quotaCategory?: pulumi.Input<string>;
}
