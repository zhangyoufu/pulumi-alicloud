// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Vpn Ipsec Servers of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.161.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.vpc.getIpsecServers({
 *     ids: ["example_id"],
 * });
 * export const vpnIpsecServerId1 = ids.then(ids => ids.servers?.[0]?.id);
 * const nameRegex = alicloud.vpc.getIpsecServers({
 *     nameRegex: "^my-IpsecServer",
 * });
 * export const vpnIpsecServerId2 = nameRegex.then(nameRegex => nameRegex.servers?.[0]?.id);
 * ```
 */
export function getIpsecServers(args?: GetIpsecServersArgs, opts?: pulumi.InvokeOptions): Promise<GetIpsecServersResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:vpc/getIpsecServers:getIpsecServers", {
        "ids": args.ids,
        "ipsecServerName": args.ipsecServerName,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "vpnGatewayId": args.vpnGatewayId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpsecServers.
 */
export interface GetIpsecServersArgs {
    /**
     * A list of Ipsec Server IDs.
     */
    ids?: string[];
    /**
     * The name of the IPsec server.
     */
    ipsecServerName?: string;
    /**
     * A regex string to filter results by Ipsec Server name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The ID of the VPN gateway.
     */
    vpnGatewayId?: string;
}

/**
 * A collection of values returned by getIpsecServers.
 */
export interface GetIpsecServersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly ipsecServerName?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly servers: outputs.vpc.GetIpsecServersServer[];
    readonly vpnGatewayId?: string;
}
/**
 * This data source provides the Vpn Ipsec Servers of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.161.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.vpc.getIpsecServers({
 *     ids: ["example_id"],
 * });
 * export const vpnIpsecServerId1 = ids.then(ids => ids.servers?.[0]?.id);
 * const nameRegex = alicloud.vpc.getIpsecServers({
 *     nameRegex: "^my-IpsecServer",
 * });
 * export const vpnIpsecServerId2 = nameRegex.then(nameRegex => nameRegex.servers?.[0]?.id);
 * ```
 */
export function getIpsecServersOutput(args?: GetIpsecServersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIpsecServersResult> {
    return pulumi.output(args).apply((a: any) => getIpsecServers(a, opts))
}

/**
 * A collection of arguments for invoking getIpsecServers.
 */
export interface GetIpsecServersOutputArgs {
    /**
     * A list of Ipsec Server IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the IPsec server.
     */
    ipsecServerName?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Ipsec Server name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of the VPN gateway.
     */
    vpnGatewayId?: pulumi.Input<string>;
}
