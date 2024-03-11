// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Vpc Nat Ip Cidrs of the current Alibaba Cloud user.
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
 * const ids = alicloud.vpc.getNatIpCidrs({
 *     natGatewayId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 * });
 * export const vpcNatIpCidrId1 = ids.then(ids => ids.cidrs?.[0]?.id);
 * const nameRegex = alicloud.vpc.getNatIpCidrs({
 *     natGatewayId: "example_value",
 *     nameRegex: "^my-NatIpCidr",
 * });
 * export const vpcNatIpCidrId2 = nameRegex.then(nameRegex => nameRegex.cidrs?.[0]?.id);
 * const status = alicloud.vpc.getNatIpCidrs({
 *     natGatewayId: "example_value",
 *     ids: ["example_value-1"],
 *     status: "Available",
 * });
 * export const vpcNatIpCidrId3 = status.then(status => status.cidrs?.[0]?.id);
 * const natIpCidr = alicloud.vpc.getNatIpCidrs({
 *     natGatewayId: "example_value",
 *     natIpCidrs: ["example_value-1"],
 * });
 * export const vpcNatIpCidrId4 = natIpCidr.then(natIpCidr => natIpCidr.cidrs?.[0]?.id);
 * const atIpCidrName = alicloud.vpc.getNatIpCidrs({
 *     natGatewayId: "example_value",
 *     natIpCidrNames: ["example_value-1"],
 * });
 * export const vpcNatIpCidrId5 = atIpCidrName.then(atIpCidrName => atIpCidrName.cidrs?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getNatIpCidrs(args: GetNatIpCidrsArgs, opts?: pulumi.InvokeOptions): Promise<GetNatIpCidrsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:vpc/getNatIpCidrs:getNatIpCidrs", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "natGatewayId": args.natGatewayId,
        "natIpCidrNames": args.natIpCidrNames,
        "natIpCidrs": args.natIpCidrs,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getNatIpCidrs.
 */
export interface GetNatIpCidrsArgs {
    /**
     * A list of Nat Ip Cidr IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Nat Ip Cidr name.
     */
    nameRegex?: string;
    /**
     * The ID of the VPC NAT gateway.
     */
    natGatewayId: string;
    /**
     * NAT IP ADDRESS the name of the root directory. Length is from `2` to `128` characters, must start with a letter or the Chinese at the beginning can contain numbers, half a period (.), underscore (_) and dash (-). But do not start with `http://` or `https://` at the beginning.
     */
    natIpCidrNames?: string[];
    /**
     * The NAT CIDR block to be created. Support up to `20`. The CIDR block must meet the following conditions: It must be `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`, or one of their subnets. The subnet mask must be `16` to `32` bits in lengths. To use a public CIDR block as the NAT CIDR block, the VPC to which the VPC NAT gateway belongs must be authorized to use public CIDR blocks. For more information, see [Create a VPC NAT gateway](https://www.alibabacloud.com/help/doc-detail/268230.htm).
     */
    natIpCidrs?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The status of the CIDR block of the NAT gateway. If the value is `Available`, the CIDR block is available.
     */
    status?: string;
}

/**
 * A collection of values returned by getNatIpCidrs.
 */
export interface GetNatIpCidrsResult {
    readonly cidrs: outputs.vpc.GetNatIpCidrsCidr[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly natGatewayId: string;
    readonly natIpCidrNames?: string[];
    readonly natIpCidrs?: string[];
    readonly outputFile?: string;
    readonly status?: string;
}
/**
 * This data source provides the Vpc Nat Ip Cidrs of the current Alibaba Cloud user.
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
 * const ids = alicloud.vpc.getNatIpCidrs({
 *     natGatewayId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 * });
 * export const vpcNatIpCidrId1 = ids.then(ids => ids.cidrs?.[0]?.id);
 * const nameRegex = alicloud.vpc.getNatIpCidrs({
 *     natGatewayId: "example_value",
 *     nameRegex: "^my-NatIpCidr",
 * });
 * export const vpcNatIpCidrId2 = nameRegex.then(nameRegex => nameRegex.cidrs?.[0]?.id);
 * const status = alicloud.vpc.getNatIpCidrs({
 *     natGatewayId: "example_value",
 *     ids: ["example_value-1"],
 *     status: "Available",
 * });
 * export const vpcNatIpCidrId3 = status.then(status => status.cidrs?.[0]?.id);
 * const natIpCidr = alicloud.vpc.getNatIpCidrs({
 *     natGatewayId: "example_value",
 *     natIpCidrs: ["example_value-1"],
 * });
 * export const vpcNatIpCidrId4 = natIpCidr.then(natIpCidr => natIpCidr.cidrs?.[0]?.id);
 * const atIpCidrName = alicloud.vpc.getNatIpCidrs({
 *     natGatewayId: "example_value",
 *     natIpCidrNames: ["example_value-1"],
 * });
 * export const vpcNatIpCidrId5 = atIpCidrName.then(atIpCidrName => atIpCidrName.cidrs?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getNatIpCidrsOutput(args: GetNatIpCidrsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNatIpCidrsResult> {
    return pulumi.output(args).apply((a: any) => getNatIpCidrs(a, opts))
}

/**
 * A collection of arguments for invoking getNatIpCidrs.
 */
export interface GetNatIpCidrsOutputArgs {
    /**
     * A list of Nat Ip Cidr IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Nat Ip Cidr name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * The ID of the VPC NAT gateway.
     */
    natGatewayId: pulumi.Input<string>;
    /**
     * NAT IP ADDRESS the name of the root directory. Length is from `2` to `128` characters, must start with a letter or the Chinese at the beginning can contain numbers, half a period (.), underscore (_) and dash (-). But do not start with `http://` or `https://` at the beginning.
     */
    natIpCidrNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The NAT CIDR block to be created. Support up to `20`. The CIDR block must meet the following conditions: It must be `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`, or one of their subnets. The subnet mask must be `16` to `32` bits in lengths. To use a public CIDR block as the NAT CIDR block, the VPC to which the VPC NAT gateway belongs must be authorized to use public CIDR blocks. For more information, see [Create a VPC NAT gateway](https://www.alibabacloud.com/help/doc-detail/268230.htm).
     */
    natIpCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the CIDR block of the NAT gateway. If the value is `Available`, the CIDR block is available.
     */
    status?: pulumi.Input<string>;
}
