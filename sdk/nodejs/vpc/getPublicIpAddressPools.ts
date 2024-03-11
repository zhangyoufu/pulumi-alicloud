// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Vpc Public Ip Address Pools of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.186.0+.
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
 * const ids = alicloud.vpc.getPublicIpAddressPools({
 *     ids: ["example_id"],
 * });
 * export const vpcPublicIpAddressPoolId1 = ids.then(ids => ids.pools?.[0]?.id);
 * const nameRegex = alicloud.vpc.getPublicIpAddressPools({
 *     nameRegex: "example_name",
 * });
 * export const vpcPublicIpAddressPoolId2 = nameRegex.then(nameRegex => nameRegex.pools?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPublicIpAddressPools(args?: GetPublicIpAddressPoolsArgs, opts?: pulumi.InvokeOptions): Promise<GetPublicIpAddressPoolsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:vpc/getPublicIpAddressPools:getPublicIpAddressPools", {
        "ids": args.ids,
        "isp": args.isp,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "publicIpAddressPoolIds": args.publicIpAddressPoolIds,
        "publicIpAddressPoolName": args.publicIpAddressPoolName,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getPublicIpAddressPools.
 */
export interface GetPublicIpAddressPoolsArgs {
    /**
     * A list of Vpc Public Ip Address Pool IDs.
     */
    ids?: string[];
    /**
     * The Internet service provider.
     */
    isp?: string;
    /**
     * A regex string to filter results by Vpc Public Ip Address Pool name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The IDs of the Vpc Public IP address pools.
     */
    publicIpAddressPoolIds?: string[];
    /**
     * The name of the Vpc Public Ip Address Pool.
     */
    publicIpAddressPoolName?: string;
    /**
     * The status of the Vpc Public Ip Address Pool.
     */
    status?: string;
}

/**
 * A collection of values returned by getPublicIpAddressPools.
 */
export interface GetPublicIpAddressPoolsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly isp?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly pools: outputs.vpc.GetPublicIpAddressPoolsPool[];
    readonly publicIpAddressPoolIds?: string[];
    readonly publicIpAddressPoolName?: string;
    readonly status?: string;
}
/**
 * This data source provides the Vpc Public Ip Address Pools of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.186.0+.
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
 * const ids = alicloud.vpc.getPublicIpAddressPools({
 *     ids: ["example_id"],
 * });
 * export const vpcPublicIpAddressPoolId1 = ids.then(ids => ids.pools?.[0]?.id);
 * const nameRegex = alicloud.vpc.getPublicIpAddressPools({
 *     nameRegex: "example_name",
 * });
 * export const vpcPublicIpAddressPoolId2 = nameRegex.then(nameRegex => nameRegex.pools?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPublicIpAddressPoolsOutput(args?: GetPublicIpAddressPoolsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPublicIpAddressPoolsResult> {
    return pulumi.output(args).apply((a: any) => getPublicIpAddressPools(a, opts))
}

/**
 * A collection of arguments for invoking getPublicIpAddressPools.
 */
export interface GetPublicIpAddressPoolsOutputArgs {
    /**
     * A list of Vpc Public Ip Address Pool IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Internet service provider.
     */
    isp?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Vpc Public Ip Address Pool name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The IDs of the Vpc Public IP address pools.
     */
    publicIpAddressPoolIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Vpc Public Ip Address Pool.
     */
    publicIpAddressPoolName?: pulumi.Input<string>;
    /**
     * The status of the Vpc Public Ip Address Pool.
     */
    status?: pulumi.Input<string>;
}
