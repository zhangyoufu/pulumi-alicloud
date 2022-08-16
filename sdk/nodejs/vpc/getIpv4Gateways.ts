// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Vpc Ipv4 Gateways of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.181.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.vpc.getIpv4Gateways({});
 * export const vpcIpv4GatewayId1 = ids.then(ids => ids.gateways?[0]?.id);
 * const nameRegex = alicloud.vpc.getIpv4Gateways({
 *     nameRegex: "^my-Ipv4Gateway",
 * });
 * export const vpcIpv4GatewayId2 = nameRegex.then(nameRegex => nameRegex.gateways?[0]?.id);
 * ```
 */
export function getIpv4Gateways(args?: GetIpv4GatewaysArgs, opts?: pulumi.InvokeOptions): Promise<GetIpv4GatewaysResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:vpc/getIpv4Gateways:getIpv4Gateways", {
        "ids": args.ids,
        "ipv4GatewayName": args.ipv4GatewayName,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpv4Gateways.
 */
export interface GetIpv4GatewaysArgs {
    /**
     * A list of Ipv4 Gateway IDs.
     */
    ids?: string[];
    /**
     * The name of the IPv4 gateway.
     */
    ipv4GatewayName?: string;
    /**
     * A regex string to filter results by Ipv4 Gateway name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The status of the resource.
     */
    status?: string;
    /**
     * The ID of the VPC associated with the IPv4 Gateway.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getIpv4Gateways.
 */
export interface GetIpv4GatewaysResult {
    readonly gateways: outputs.vpc.GetIpv4GatewaysGateway[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly ipv4GatewayName?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly status?: string;
    readonly vpcId?: string;
}

export function getIpv4GatewaysOutput(args?: GetIpv4GatewaysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIpv4GatewaysResult> {
    return pulumi.output(args).apply(a => getIpv4Gateways(a, opts))
}

/**
 * A collection of arguments for invoking getIpv4Gateways.
 */
export interface GetIpv4GatewaysOutputArgs {
    /**
     * A list of Ipv4 Gateway IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the IPv4 gateway.
     */
    ipv4GatewayName?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Ipv4 Gateway name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the VPC associated with the IPv4 Gateway.
     */
    vpcId?: pulumi.Input<string>;
}
