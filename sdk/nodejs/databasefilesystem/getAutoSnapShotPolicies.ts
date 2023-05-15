// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides Dbfs Auto Snap Shot Policy available to the user.[What is Auto Snap Shot Policy](https://help.aliyun.com/document_detail/469597.html)
 *
 * > **NOTE:** Available in 1.202.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.databasefilesystem.getAutoSnapShotPolicies({
 *     ids: [alicloud_dbfs_auto_snap_shot_policy["default"].id],
 * });
 * export const alicloudDbfsAutoSnapShotPolicyExampleId = _default.then(_default => _default.autoSnapShotPolicies?.[0]?.id);
 * ```
 */
export function getAutoSnapShotPolicies(args?: GetAutoSnapShotPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetAutoSnapShotPoliciesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:databasefilesystem/getAutoSnapShotPolicies:getAutoSnapShotPolicies", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
    }, opts);
}

/**
 * A collection of arguments for invoking getAutoSnapShotPolicies.
 */
export interface GetAutoSnapShotPoliciesArgs {
    /**
     * A list of Auto Snap Shot Policy IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Auto Snap Shot Policy name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
}

/**
 * A collection of values returned by getAutoSnapShotPolicies.
 */
export interface GetAutoSnapShotPoliciesResult {
    /**
     * A list of Auto Snap Shot Policy Entries. Each element contains the following attributes:
     */
    readonly autoSnapShotPolicies: outputs.databasefilesystem.GetAutoSnapShotPoliciesAutoSnapShotPolicy[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of Auto Snap Shot Policy IDs.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of Auto Snap Shot Policy names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
}
/**
 * This data source provides Dbfs Auto Snap Shot Policy available to the user.[What is Auto Snap Shot Policy](https://help.aliyun.com/document_detail/469597.html)
 *
 * > **NOTE:** Available in 1.202.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.databasefilesystem.getAutoSnapShotPolicies({
 *     ids: [alicloud_dbfs_auto_snap_shot_policy["default"].id],
 * });
 * export const alicloudDbfsAutoSnapShotPolicyExampleId = _default.then(_default => _default.autoSnapShotPolicies?.[0]?.id);
 * ```
 */
export function getAutoSnapShotPoliciesOutput(args?: GetAutoSnapShotPoliciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAutoSnapShotPoliciesResult> {
    return pulumi.output(args).apply((a: any) => getAutoSnapShotPolicies(a, opts))
}

/**
 * A collection of arguments for invoking getAutoSnapShotPolicies.
 */
export interface GetAutoSnapShotPoliciesOutputArgs {
    /**
     * A list of Auto Snap Shot Policy IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Auto Snap Shot Policy name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
}
