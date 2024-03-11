// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the HBR Replication Vault Regions of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.152.0+.
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
 * const default = alicloud.hbr.getReplicationVaultRegions({});
 * export const hbrReplicationVaultRegionRegionId1 = _default.then(_default => _default.regions?.[0]?.replicationRegionId);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getReplicationVaultRegions(args?: GetReplicationVaultRegionsArgs, opts?: pulumi.InvokeOptions): Promise<GetReplicationVaultRegionsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:hbr/getReplicationVaultRegions:getReplicationVaultRegions", {
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getReplicationVaultRegions.
 */
export interface GetReplicationVaultRegionsArgs {
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getReplicationVaultRegions.
 */
export interface GetReplicationVaultRegionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    readonly regions: outputs.hbr.GetReplicationVaultRegionsRegion[];
}
/**
 * This data source provides the HBR Replication Vault Regions of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.152.0+.
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
 * const default = alicloud.hbr.getReplicationVaultRegions({});
 * export const hbrReplicationVaultRegionRegionId1 = _default.then(_default => _default.regions?.[0]?.replicationRegionId);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getReplicationVaultRegionsOutput(args?: GetReplicationVaultRegionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReplicationVaultRegionsResult> {
    return pulumi.output(args).apply((a: any) => getReplicationVaultRegions(a, opts))
}

/**
 * A collection of arguments for invoking getReplicationVaultRegions.
 */
export interface GetReplicationVaultRegionsOutputArgs {
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
