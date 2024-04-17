// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides Eflo Subnet available to the user.[What is Subnet](https://help.aliyun.com/document_detail/604977.html)
 *
 * > **NOTE:** Available in 1.204.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.eflo.getSubnets({
 *     nameRegex: defaultAlicloudEfloSubnet.name,
 *     subnetName: "SubnetTestForTerraform",
 *     vpdId: vpdId,
 *     zoneId: zoneId,
 * });
 * export const alicloudEfloSubnetExampleId = _default.then(_default => _default.subnets?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSubnets(args?: GetSubnetsArgs, opts?: pulumi.InvokeOptions): Promise<GetSubnetsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:eflo/getSubnets:getSubnets", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
        "resourceGroupId": args.resourceGroupId,
        "status": args.status,
        "subnetId": args.subnetId,
        "subnetName": args.subnetName,
        "type": args.type,
        "vpdId": args.vpdId,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubnets.
 */
export interface GetSubnetsArgs {
    ids?: string[];
    /**
     * A regex string to filter results by Group Metric Rule name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
    /**
     * Resource Group ID.
     */
    resourceGroupId?: string;
    /**
     * The status of the resource.
     */
    status?: string;
    /**
     * Primary key ID.
     */
    subnetId?: string;
    /**
     * The Subnet name.
     */
    subnetName?: string;
    /**
     * Eflo subnet usage type, optional value: 
     * - General type is not filled in
     * - OOB:OOB type
     * - LB: LB type
     */
    type?: string;
    /**
     * The Eflo VPD ID.
     */
    vpdId?: string;
    /**
     * The zone ID of the resource.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by getSubnets.
 */
export interface GetSubnetsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of name of Subnets.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
    /**
     * Resource Group ID.
     */
    readonly resourceGroupId?: string;
    /**
     * The status of the resource.
     */
    readonly status?: string;
    /**
     * The Eflo subnet ID.
     */
    readonly subnetId?: string;
    /**
     * The Subnet name.
     */
    readonly subnetName?: string;
    /**
     * A list of Subnet Entries. Each element contains the following attributes:
     */
    readonly subnets: outputs.eflo.GetSubnetsSubnet[];
    /**
     * Eflo subnet usage type.
     */
    readonly type?: string;
    /**
     * Eflo VPD ID.
     */
    readonly vpdId?: string;
    /**
     * The zone ID of the resource.
     */
    readonly zoneId?: string;
}
/**
 * This data source provides Eflo Subnet available to the user.[What is Subnet](https://help.aliyun.com/document_detail/604977.html)
 *
 * > **NOTE:** Available in 1.204.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.eflo.getSubnets({
 *     nameRegex: defaultAlicloudEfloSubnet.name,
 *     subnetName: "SubnetTestForTerraform",
 *     vpdId: vpdId,
 *     zoneId: zoneId,
 * });
 * export const alicloudEfloSubnetExampleId = _default.then(_default => _default.subnets?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSubnetsOutput(args?: GetSubnetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubnetsResult> {
    return pulumi.output(args).apply((a: any) => getSubnets(a, opts))
}

/**
 * A collection of arguments for invoking getSubnets.
 */
export interface GetSubnetsOutputArgs {
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Group Metric Rule name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
    /**
     * Resource Group ID.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
    /**
     * Primary key ID.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The Subnet name.
     */
    subnetName?: pulumi.Input<string>;
    /**
     * Eflo subnet usage type, optional value: 
     * - General type is not filled in
     * - OOB:OOB type
     * - LB: LB type
     */
    type?: pulumi.Input<string>;
    /**
     * The Eflo VPD ID.
     */
    vpdId?: pulumi.Input<string>;
    /**
     * The zone ID of the resource.
     */
    zoneId?: pulumi.Input<string>;
}
