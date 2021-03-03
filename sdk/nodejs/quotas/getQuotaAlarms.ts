// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Quotas Quota Alarms of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.116.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.quotas.getQuotaAlarms({
 *     ids: ["5VR90-421F886-81E9-xxx"],
 *     nameRegex: "tf-testAcc",
 * });
 * export const firstQuotasQuotaAlarmId = example.then(example => example.alarms[0].id);
 * ```
 */
export function getQuotaAlarms(args?: GetQuotaAlarmsArgs, opts?: pulumi.InvokeOptions): Promise<GetQuotaAlarmsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:quotas/getQuotaAlarms:getQuotaAlarms", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "productCode": args.productCode,
        "quotaActionCode": args.quotaActionCode,
        "quotaAlarmName": args.quotaAlarmName,
        "quotaDimensions": args.quotaDimensions,
    }, opts);
}

/**
 * A collection of arguments for invoking getQuotaAlarms.
 */
export interface GetQuotaAlarmsArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    readonly enableDetails?: boolean;
    /**
     * A list of Quota Alarm IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by Quota Alarm name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The Product Code.
     */
    readonly productCode?: string;
    /**
     * The Quota Action Code.
     */
    readonly quotaActionCode?: string;
    /**
     * The name of Quota Alarm.
     */
    readonly quotaAlarmName?: string;
    /**
     * The Quota Dimensions.
     */
    readonly quotaDimensions?: inputs.quotas.GetQuotaAlarmsQuotaDimension[];
}

/**
 * A collection of values returned by getQuotaAlarms.
 */
export interface GetQuotaAlarmsResult {
    readonly alarms: outputs.quotas.GetQuotaAlarmsAlarm[];
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly productCode?: string;
    readonly quotaActionCode?: string;
    readonly quotaAlarmName?: string;
    readonly quotaDimensions?: outputs.quotas.GetQuotaAlarmsQuotaDimension[];
}
