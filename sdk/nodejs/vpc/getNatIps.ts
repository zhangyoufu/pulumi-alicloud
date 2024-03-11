// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Vpc Nat Ips of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.136.0+.
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
 * const ids = alicloud.vpc.getNatIps({
 *     natGatewayId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 * });
 * export const vpcNatIpId1 = ids.then(ids => ids.ips?.[0]?.id);
 * const nameRegex = alicloud.vpc.getNatIps({
 *     natGatewayId: "example_value",
 *     nameRegex: "^my-NatIp",
 * });
 * export const vpcNatIpId2 = nameRegex.then(nameRegex => nameRegex.ips?.[0]?.id);
 * const natIpCidr = alicloud.vpc.getNatIps({
 *     natGatewayId: "example_value",
 *     natIpCidr: "example_value",
 *     nameRegex: "^my-NatIp",
 * });
 * export const vpcNatIpId3 = natIpCidr.then(natIpCidr => natIpCidr.ips?.[0]?.id);
 * const natIpName = alicloud.vpc.getNatIps({
 *     natGatewayId: "example_value",
 *     ids: ["example_value"],
 *     natIpNames: ["example_value"],
 * });
 * export const vpcNatIpId4 = natIpName.then(natIpName => natIpName.ips?.[0]?.id);
 * const natIpIds = alicloud.vpc.getNatIps({
 *     natGatewayId: "example_value",
 *     ids: ["example_value"],
 *     natIpIds: ["example_value"],
 * });
 * export const vpcNatIpId5 = natIpIds.then(natIpIds => natIpIds.ips?.[0]?.id);
 * const status = alicloud.vpc.getNatIps({
 *     natGatewayId: "example_value",
 *     ids: ["example_value"],
 *     status: "example_value",
 * });
 * export const vpcNatIpId6 = status.then(status => status.ips?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getNatIps(args: GetNatIpsArgs, opts?: pulumi.InvokeOptions): Promise<GetNatIpsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:vpc/getNatIps:getNatIps", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "natGatewayId": args.natGatewayId,
        "natIpCidr": args.natIpCidr,
        "natIpIds": args.natIpIds,
        "natIpNames": args.natIpNames,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getNatIps.
 */
export interface GetNatIpsArgs {
    /**
     * A list of Nat Ip IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Nat Ip name.
     */
    nameRegex?: string;
    /**
     * The ID of the Virtual Private Cloud (VPC) NAT gateway to which the NAT IP address belongs.
     */
    natGatewayId: string;
    /**
     * The CIDR block to which the NAT IP address belongs.
     */
    natIpCidr?: string;
    natIpIds?: string[];
    /**
     * The name of the NAT IP address.
     */
    natIpNames?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The status of the NAT IP address. Valid values: `Available`, `Deleting` and `Creating`.
     */
    status?: string;
}

/**
 * A collection of values returned by getNatIps.
 */
export interface GetNatIpsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly ips: outputs.vpc.GetNatIpsIp[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly natGatewayId: string;
    readonly natIpCidr?: string;
    readonly natIpIds?: string[];
    readonly natIpNames?: string[];
    readonly outputFile?: string;
    readonly status?: string;
}
/**
 * This data source provides the Vpc Nat Ips of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.136.0+.
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
 * const ids = alicloud.vpc.getNatIps({
 *     natGatewayId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 * });
 * export const vpcNatIpId1 = ids.then(ids => ids.ips?.[0]?.id);
 * const nameRegex = alicloud.vpc.getNatIps({
 *     natGatewayId: "example_value",
 *     nameRegex: "^my-NatIp",
 * });
 * export const vpcNatIpId2 = nameRegex.then(nameRegex => nameRegex.ips?.[0]?.id);
 * const natIpCidr = alicloud.vpc.getNatIps({
 *     natGatewayId: "example_value",
 *     natIpCidr: "example_value",
 *     nameRegex: "^my-NatIp",
 * });
 * export const vpcNatIpId3 = natIpCidr.then(natIpCidr => natIpCidr.ips?.[0]?.id);
 * const natIpName = alicloud.vpc.getNatIps({
 *     natGatewayId: "example_value",
 *     ids: ["example_value"],
 *     natIpNames: ["example_value"],
 * });
 * export const vpcNatIpId4 = natIpName.then(natIpName => natIpName.ips?.[0]?.id);
 * const natIpIds = alicloud.vpc.getNatIps({
 *     natGatewayId: "example_value",
 *     ids: ["example_value"],
 *     natIpIds: ["example_value"],
 * });
 * export const vpcNatIpId5 = natIpIds.then(natIpIds => natIpIds.ips?.[0]?.id);
 * const status = alicloud.vpc.getNatIps({
 *     natGatewayId: "example_value",
 *     ids: ["example_value"],
 *     status: "example_value",
 * });
 * export const vpcNatIpId6 = status.then(status => status.ips?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getNatIpsOutput(args: GetNatIpsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNatIpsResult> {
    return pulumi.output(args).apply((a: any) => getNatIps(a, opts))
}

/**
 * A collection of arguments for invoking getNatIps.
 */
export interface GetNatIpsOutputArgs {
    /**
     * A list of Nat Ip IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Nat Ip name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * The ID of the Virtual Private Cloud (VPC) NAT gateway to which the NAT IP address belongs.
     */
    natGatewayId: pulumi.Input<string>;
    /**
     * The CIDR block to which the NAT IP address belongs.
     */
    natIpCidr?: pulumi.Input<string>;
    natIpIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the NAT IP address.
     */
    natIpNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the NAT IP address. Valid values: `Available`, `Deleting` and `Creating`.
     */
    status?: pulumi.Input<string>;
}
