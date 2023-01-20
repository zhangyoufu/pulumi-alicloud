// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Hbr OssBackupPlans of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.131.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.hbr.getOssBackupPlans({
 *     nameRegex: "^my-OssBackupPlan",
 * });
 * export const hbrOssBackupPlanId = ids.then(ids => ids.plans?.[0]?.id);
 * ```
 */
export function getOssBackupPlans(args?: GetOssBackupPlansArgs, opts?: pulumi.InvokeOptions): Promise<GetOssBackupPlansResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:hbr/getOssBackupPlans:getOssBackupPlans", {
        "bucket": args.bucket,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "vaultId": args.vaultId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOssBackupPlans.
 */
export interface GetOssBackupPlansArgs {
    /**
     * The name of OSS bucket.
     */
    bucket?: string;
    /**
     * A list of OssBackupPlan IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by OssBackupPlan name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The ID of backup vault.
     */
    vaultId?: string;
}

/**
 * A collection of values returned by getOssBackupPlans.
 */
export interface GetOssBackupPlansResult {
    readonly bucket?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly plans: outputs.hbr.GetOssBackupPlansPlan[];
    readonly vaultId?: string;
}
/**
 * This data source provides the Hbr OssBackupPlans of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.131.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.hbr.getOssBackupPlans({
 *     nameRegex: "^my-OssBackupPlan",
 * });
 * export const hbrOssBackupPlanId = ids.then(ids => ids.plans?.[0]?.id);
 * ```
 */
export function getOssBackupPlansOutput(args?: GetOssBackupPlansOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOssBackupPlansResult> {
    return pulumi.output(args).apply((a: any) => getOssBackupPlans(a, opts))
}

/**
 * A collection of arguments for invoking getOssBackupPlans.
 */
export interface GetOssBackupPlansOutputArgs {
    /**
     * The name of OSS bucket.
     */
    bucket?: pulumi.Input<string>;
    /**
     * A list of OssBackupPlan IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by OssBackupPlan name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of backup vault.
     */
    vaultId?: pulumi.Input<string>;
}
