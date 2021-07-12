// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the server load balancers of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in 1.123.1+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.slb.getApplicationLoadBalancers({
 *     nameRegex: "sample_slb",
 *     tags: {
 *         tagKey1: "tagValue1",
 *         tagKey2: "tagValue2",
 *     },
 * });
 * export const firstSlbId = example.then(example => example.balancers[0].id);
 * ```
 */
export function getApplicationLoadBalancers(args?: GetApplicationLoadBalancersArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationLoadBalancersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:slb/getApplicationLoadBalancers:getApplicationLoadBalancers", {
        "address": args.address,
        "addressIpVersion": args.addressIpVersion,
        "addressType": args.addressType,
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "internetChargeType": args.internetChargeType,
        "loadBalancerName": args.loadBalancerName,
        "masterZoneId": args.masterZoneId,
        "nameRegex": args.nameRegex,
        "networkType": args.networkType,
        "outputFile": args.outputFile,
        "paymentType": args.paymentType,
        "resourceGroupId": args.resourceGroupId,
        "serverId": args.serverId,
        "serverIntranetAddress": args.serverIntranetAddress,
        "slaveZoneId": args.slaveZoneId,
        "status": args.status,
        "tags": args.tags,
        "vpcId": args.vpcId,
        "vswitchId": args.vswitchId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationLoadBalancers.
 */
export interface GetApplicationLoadBalancersArgs {
    /**
     * Service address of the SLBs.
     */
    readonly address?: string;
    /**
     * The address ip version. Valid values `ipv4` and `ipv6`.
     */
    readonly addressIpVersion?: string;
    /**
     * The address type of the SLB. Valid values `internet` and `intranet`.
     */
    readonly addressType?: string;
    readonly enableDetails?: boolean;
    /**
     * A list of SLBs IDs.
     */
    readonly ids?: string[];
    /**
     * The internet charge type. Valid values `PayByBandwidth` and `PayByTraffic`.
     */
    readonly internetChargeType?: string;
    /**
     * The name of the SLB.
     */
    readonly loadBalancerName?: string;
    /**
     * The master zone id of the SLB.
     */
    readonly masterZoneId?: string;
    /**
     * A regex string to filter results by SLB name.
     */
    readonly nameRegex?: string;
    /**
     * Network type of the SLBs. Valid values: `vpc` and `classic`.
     */
    readonly networkType?: string;
    readonly outputFile?: string;
    /**
     * The payment type of SLB. Valid values `PayAsYouGo` and `Subscription`.
     */
    readonly paymentType?: string;
    /**
     * The Id of resource group which SLB belongs.
     */
    readonly resourceGroupId?: string;
    /**
     * The server ID.
     */
    readonly serverId?: string;
    /**
     * The server intranet address.
     */
    readonly serverIntranetAddress?: string;
    /**
     * The slave zone id of the SLB.
     */
    readonly slaveZoneId?: string;
    /**
     * SLB current status. Possible values: `inactive`, `active` and `locked`.
     */
    readonly status?: string;
    /**
     * A map of tags assigned to the SLB instances. The `tags` can have a maximum of 5 tag. It must be in the format:
     */
    readonly tags?: {[key: string]: any};
    /**
     * ID of the VPC linked to the SLBs.
     */
    readonly vpcId?: string;
    /**
     * ID of the VSwitch linked to the SLBs.
     */
    readonly vswitchId?: string;
}

/**
 * A collection of values returned by getApplicationLoadBalancers.
 */
export interface GetApplicationLoadBalancersResult {
    /**
     * The IP address that the SLB instance uses to provide services.
     */
    readonly address?: string;
    /**
     * The address ip version.
     */
    readonly addressIpVersion?: string;
    /**
     * The address type.
     */
    readonly addressType?: string;
    /**
     * A list of SLBs. Each element contains the following attributes:
     */
    readonly balancers: outputs.slb.GetApplicationLoadBalancersBalancer[];
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of slb IDs.
     */
    readonly ids: string[];
    /**
     * The billing method of the Internet-facing SLB instance.
     */
    readonly internetChargeType?: string;
    /**
     * The name of the SLB.
     */
    readonly loadBalancerName?: string;
    /**
     * Master availability zone of the SLBs.
     */
    readonly masterZoneId?: string;
    readonly nameRegex?: string;
    /**
     * A list of slb names.
     */
    readonly names: string[];
    /**
     * Network type of the SLB. Possible values: `vpc` and `classic`.
     */
    readonly networkType?: string;
    readonly outputFile?: string;
    readonly paymentType?: string;
    /**
     * The ID of the resource group.
     */
    readonly resourceGroupId?: string;
    /**
     * The ID of the Elastic Compute Service (ECS) instance that is specified as a backend server of the CLB instance.
     */
    readonly serverId?: string;
    readonly serverIntranetAddress?: string;
    /**
     * Slave availability zone of the SLBs.
     */
    readonly slaveZoneId?: string;
    /**
     * @deprecated Field 'slbs' has deprecated from v1.123.1 and replace by 'balancers'.
     */
    readonly slbs: outputs.slb.GetApplicationLoadBalancersSlb[];
    /**
     * SLB current status. Possible values: `inactive`, `active` and `locked`.
     */
    readonly status?: string;
    /**
     * The tags of the SLB.
     */
    readonly tags?: {[key: string]: any};
    /**
     * ID of the VPC the SLB belongs to.
     */
    readonly vpcId?: string;
    /**
     * ID of the VSwitch the SLB belongs to.
     */
    readonly vswitchId?: string;
}
