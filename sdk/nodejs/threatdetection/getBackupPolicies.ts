// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Threat Detection Backup Policies of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.195.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.threatdetection.getBackupPolicies({
 *     ids: ["example_id"],
 * });
 * export const threatDetectionBackupPoliciesId1 = ids.then(ids => ids.policies?.[0]?.id);
 * const nameRegex = alicloud.threatdetection.getBackupPolicies({
 *     nameRegex: "tf-example",
 * });
 * export const threatDetectionBackupPoliciesId2 = nameRegex.then(nameRegex => nameRegex.policies?.[0]?.id);
 * ```
 */
export function getBackupPolicies(args?: GetBackupPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetBackupPoliciesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:threatdetection/getBackupPolicies:getBackupPolicies", {
        "currentPage": args.currentPage,
        "ids": args.ids,
        "machineRemark": args.machineRemark,
        "name": args.name,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "pageSize": args.pageSize,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackupPolicies.
 */
export interface GetBackupPoliciesArgs {
    currentPage?: number;
    /**
     * A list of Threat Detection Backup Policies IDs.
     */
    ids?: string[];
    /**
     * The information that you want to use to identify the servers protected by the anti-ransomware policy. You can enter the IP address or ID of a server.
     */
    machineRemark?: string;
    /**
     * The name of the anti-ransomware policy that you want to query.
     */
    name?: string;
    /**
     * A regex string to filter results by Threat Detection Backup Policies name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    pageSize?: number;
    /**
     * The status of the anti-ransomware policy. Valid Value: `enabled`, `disabled`, `closed`.
     */
    status?: string;
}

/**
 * A collection of values returned by getBackupPolicies.
 */
export interface GetBackupPoliciesResult {
    readonly currentPage?: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly machineRemark?: string;
    readonly name?: string;
    readonly nameRegex?: string;
    /**
     * A list of Threat Detection Backup Policy names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly pageSize?: number;
    /**
     * A list of Threat Detection Backup policies. Each element contains the following attributes:
     */
    readonly policies: outputs.threatdetection.GetBackupPoliciesPolicy[];
    /**
     * The status of the anti-ransomware policy.
     */
    readonly status?: string;
}
/**
 * This data source provides the Threat Detection Backup Policies of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.195.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.threatdetection.getBackupPolicies({
 *     ids: ["example_id"],
 * });
 * export const threatDetectionBackupPoliciesId1 = ids.then(ids => ids.policies?.[0]?.id);
 * const nameRegex = alicloud.threatdetection.getBackupPolicies({
 *     nameRegex: "tf-example",
 * });
 * export const threatDetectionBackupPoliciesId2 = nameRegex.then(nameRegex => nameRegex.policies?.[0]?.id);
 * ```
 */
export function getBackupPoliciesOutput(args?: GetBackupPoliciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackupPoliciesResult> {
    return pulumi.output(args).apply((a: any) => getBackupPolicies(a, opts))
}

/**
 * A collection of arguments for invoking getBackupPolicies.
 */
export interface GetBackupPoliciesOutputArgs {
    currentPage?: pulumi.Input<number>;
    /**
     * A list of Threat Detection Backup Policies IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The information that you want to use to identify the servers protected by the anti-ransomware policy. You can enter the IP address or ID of a server.
     */
    machineRemark?: pulumi.Input<string>;
    /**
     * The name of the anti-ransomware policy that you want to query.
     */
    name?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Threat Detection Backup Policies name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    pageSize?: pulumi.Input<number>;
    /**
     * The status of the anti-ransomware policy. Valid Value: `enabled`, `disabled`, `closed`.
     */
    status?: pulumi.Input<string>;
}
