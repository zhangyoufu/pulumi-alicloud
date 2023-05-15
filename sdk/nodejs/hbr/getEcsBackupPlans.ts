// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Hbr EcsBackupPlans of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.132.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.hbr.getEcsBackupPlans({
 *     nameRegex: "plan-name",
 * });
 * export const hbrEcsBackupPlanId = ids.then(ids => ids.plans?.[0]?.id);
 * ```
 */
export function getEcsBackupPlans(args?: GetEcsBackupPlansArgs, opts?: pulumi.InvokeOptions): Promise<GetEcsBackupPlansResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:hbr/getEcsBackupPlans:getEcsBackupPlans", {
        "ids": args.ids,
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "vaultId": args.vaultId,
    }, opts);
}

/**
 * A collection of arguments for invoking getEcsBackupPlans.
 */
export interface GetEcsBackupPlansArgs {
    /**
     * A list of EcsBackupPlan IDs.
     */
    ids?: string[];
    /**
     * The ID of ECS instance.
     */
    instanceId?: string;
    /**
     * A regex string to filter results by EcsBackupPlan name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The ID of Backup vault.
     */
    vaultId?: string;
}

/**
 * A collection of values returned by getEcsBackupPlans.
 */
export interface GetEcsBackupPlansResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly instanceId?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly plans: outputs.hbr.GetEcsBackupPlansPlan[];
    readonly vaultId?: string;
}
/**
 * This data source provides the Hbr EcsBackupPlans of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.132.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.hbr.getEcsBackupPlans({
 *     nameRegex: "plan-name",
 * });
 * export const hbrEcsBackupPlanId = ids.then(ids => ids.plans?.[0]?.id);
 * ```
 */
export function getEcsBackupPlansOutput(args?: GetEcsBackupPlansOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEcsBackupPlansResult> {
    return pulumi.output(args).apply((a: any) => getEcsBackupPlans(a, opts))
}

/**
 * A collection of arguments for invoking getEcsBackupPlans.
 */
export interface GetEcsBackupPlansOutputArgs {
    /**
     * A list of EcsBackupPlan IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of ECS instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * A regex string to filter results by EcsBackupPlan name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of Backup vault.
     */
    vaultId?: pulumi.Input<string>;
}
