// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Global Accelerator (GA) Endpoint Group Ip Address Cidr Blocks of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available since v1.213.0.
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
 * const default = alicloud.ga.getEndpointGroupIpAddressCidrBlocks({
 *     endpointGroupRegion: "cn-hangzhou",
 * });
 * export const gaEndpointGroupIpAddressCidrBlocksEndpointGroupRegion = _default.then(_default => _default.endpointGroupIpAddressCidrBlocks?.[0]?.endpointGroupRegion);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getEndpointGroupIpAddressCidrBlocks(args: GetEndpointGroupIpAddressCidrBlocksArgs, opts?: pulumi.InvokeOptions): Promise<GetEndpointGroupIpAddressCidrBlocksResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ga/getEndpointGroupIpAddressCidrBlocks:getEndpointGroupIpAddressCidrBlocks", {
        "endpointGroupRegion": args.endpointGroupRegion,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getEndpointGroupIpAddressCidrBlocks.
 */
export interface GetEndpointGroupIpAddressCidrBlocksArgs {
    /**
     * The region ID of the endpoint group.
     */
    endpointGroupRegion: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getEndpointGroupIpAddressCidrBlocks.
 */
export interface GetEndpointGroupIpAddressCidrBlocksResult {
    /**
     * A list of Endpoint Group Ip Address Cidr Blocks. Each element contains the following attributes:
     */
    readonly endpointGroupIpAddressCidrBlocks: outputs.ga.GetEndpointGroupIpAddressCidrBlocksEndpointGroupIpAddressCidrBlock[];
    /**
     * The region ID of the endpoint group.
     */
    readonly endpointGroupRegion: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
}
/**
 * This data source provides the Global Accelerator (GA) Endpoint Group Ip Address Cidr Blocks of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available since v1.213.0.
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
 * const default = alicloud.ga.getEndpointGroupIpAddressCidrBlocks({
 *     endpointGroupRegion: "cn-hangzhou",
 * });
 * export const gaEndpointGroupIpAddressCidrBlocksEndpointGroupRegion = _default.then(_default => _default.endpointGroupIpAddressCidrBlocks?.[0]?.endpointGroupRegion);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getEndpointGroupIpAddressCidrBlocksOutput(args: GetEndpointGroupIpAddressCidrBlocksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEndpointGroupIpAddressCidrBlocksResult> {
    return pulumi.output(args).apply((a: any) => getEndpointGroupIpAddressCidrBlocks(a, opts))
}

/**
 * A collection of arguments for invoking getEndpointGroupIpAddressCidrBlocks.
 */
export interface GetEndpointGroupIpAddressCidrBlocksOutputArgs {
    /**
     * The region ID of the endpoint group.
     */
    endpointGroupRegion: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
