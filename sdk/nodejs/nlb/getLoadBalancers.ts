// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Nlb Load Balancers of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.191.0+.
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
 * const ids = alicloud.nlb.getLoadBalancers({
 *     ids: ["example_id"],
 * });
 * export const nlbLoadBalancerId1 = ids.then(ids => ids.balancers?.[0]?.id);
 * const nameRegex = alicloud.nlb.getLoadBalancers({
 *     nameRegex: "^my-LoadBalancer",
 * });
 * export const nlbLoadBalancerId2 = nameRegex.then(nameRegex => nameRegex.balancers?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getLoadBalancers(args?: GetLoadBalancersArgs, opts?: pulumi.InvokeOptions): Promise<GetLoadBalancersResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:nlb/getLoadBalancers:getLoadBalancers", {
        "addressIpVersion": args.addressIpVersion,
        "addressType": args.addressType,
        "dnsName": args.dnsName,
        "ids": args.ids,
        "ipv6AddressType": args.ipv6AddressType,
        "loadBalancerBusinessStatus": args.loadBalancerBusinessStatus,
        "loadBalancerNames": args.loadBalancerNames,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
        "status": args.status,
        "tags": args.tags,
        "vpcIds": args.vpcIds,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getLoadBalancers.
 */
export interface GetLoadBalancersArgs {
    /**
     * The IP version.
     */
    addressIpVersion?: string;
    /**
     * The type of IPv4 address used by the NLB instance.
     */
    addressType?: string;
    /**
     * The domain name of the NLB instance.
     */
    dnsName?: string;
    /**
     * A list of Load Balancer IDs.
     */
    ids?: string[];
    /**
     * The type of IPv6 address used by the NLB instance.
     */
    ipv6AddressType?: string;
    /**
     * The business status of the NLB instance.
     */
    loadBalancerBusinessStatus?: string;
    /**
     * The name of the NLB instance. You can specify at most 10 names.
     */
    loadBalancerNames?: string[];
    /**
     * A regex string to filter results by Load Balancer name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: string;
    /**
     * The status of the NLB instance.
     */
    status?: string;
    /**
     * The tag of the resource.
     */
    tags?: {[key: string]: any};
    /**
     * The ID of the virtual private cloud (VPC) where the NLB instance is deployed. You can specify at most 10 IDs.
     */
    vpcIds?: string[];
    /**
     * The name of the zone.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by getLoadBalancers.
 */
export interface GetLoadBalancersResult {
    readonly addressIpVersion?: string;
    readonly addressType?: string;
    readonly balancers: outputs.nlb.GetLoadBalancersBalancer[];
    readonly dnsName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly ipv6AddressType?: string;
    readonly loadBalancerBusinessStatus?: string;
    readonly loadBalancerNames?: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly resourceGroupId?: string;
    readonly status?: string;
    readonly tags?: {[key: string]: any};
    readonly vpcIds?: string[];
    readonly zoneId?: string;
}
/**
 * This data source provides the Nlb Load Balancers of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.191.0+.
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
 * const ids = alicloud.nlb.getLoadBalancers({
 *     ids: ["example_id"],
 * });
 * export const nlbLoadBalancerId1 = ids.then(ids => ids.balancers?.[0]?.id);
 * const nameRegex = alicloud.nlb.getLoadBalancers({
 *     nameRegex: "^my-LoadBalancer",
 * });
 * export const nlbLoadBalancerId2 = nameRegex.then(nameRegex => nameRegex.balancers?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getLoadBalancersOutput(args?: GetLoadBalancersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLoadBalancersResult> {
    return pulumi.output(args).apply((a: any) => getLoadBalancers(a, opts))
}

/**
 * A collection of arguments for invoking getLoadBalancers.
 */
export interface GetLoadBalancersOutputArgs {
    /**
     * The IP version.
     */
    addressIpVersion?: pulumi.Input<string>;
    /**
     * The type of IPv4 address used by the NLB instance.
     */
    addressType?: pulumi.Input<string>;
    /**
     * The domain name of the NLB instance.
     */
    dnsName?: pulumi.Input<string>;
    /**
     * A list of Load Balancer IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of IPv6 address used by the NLB instance.
     */
    ipv6AddressType?: pulumi.Input<string>;
    /**
     * The business status of the NLB instance.
     */
    loadBalancerBusinessStatus?: pulumi.Input<string>;
    /**
     * The name of the NLB instance. You can specify at most 10 names.
     */
    loadBalancerNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Load Balancer name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The status of the NLB instance.
     */
    status?: pulumi.Input<string>;
    /**
     * The tag of the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the virtual private cloud (VPC) where the NLB instance is deployed. You can specify at most 10 IDs.
     */
    vpcIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the zone.
     */
    zoneId?: pulumi.Input<string>;
}
