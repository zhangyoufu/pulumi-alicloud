// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the master slave server groups related to a server load balancer.
 *
 * > **NOTE:** Available in 1.54.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultZones = alicloud.getZones({
 *     availableDiskCategory: "cloud_efficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultInstanceTypes = defaultZones.then(defaultZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones[0].id,
 *     eniAmount: 2,
 * }));
 * const image = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_18.*64",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-testAccSlbMasterSlaveServerGroupVpc";
 * const number = config.get("number") || "1";
 * const mainNetwork = new alicloud.vpc.Network("mainNetwork", {cidrBlock: "172.16.0.0/16"});
 * const mainSwitch = new alicloud.vpc.Switch("mainSwitch", {
 *     vpcId: mainNetwork.id,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones[0].id),
 *     vswitchName: name,
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const groupSecurityGroup = new alicloud.ecs.SecurityGroup("groupSecurityGroup", {vpcId: mainNetwork.id});
 * const instanceInstance: alicloud.ecs.Instance[];
 * for (const range = {value: 0}; range.value < "2"; range.value++) {
 *     instanceInstance.push(new alicloud.ecs.Instance(`instanceInstance-${range.value}`, {
 *         imageId: image.then(image => image.images[0].id),
 *         instanceType: defaultInstanceTypes.then(defaultInstanceTypes => defaultInstanceTypes.instanceTypes[0].id),
 *         instanceName: name,
 *         securityGroups: [groupSecurityGroup.id],
 *         internetChargeType: "PayByTraffic",
 *         internetMaxBandwidthOut: "10",
 *         availabilityZone: defaultZones.then(defaultZones => defaultZones.zones[0].id),
 *         instanceChargeType: "PostPaid",
 *         systemDiskCategory: "cloud_efficiency",
 *         vswitchId: mainSwitch.id,
 *     }));
 * }
 * const instanceApplicationLoadBalancer = new alicloud.slb.ApplicationLoadBalancer("instanceApplicationLoadBalancer", {
 *     loadBalancerName: name,
 *     vswitchId: mainSwitch.id,
 *     loadBalancerSpec: "slb.s2.small",
 * });
 * const groupMasterSlaveServerGroup = new alicloud.slb.MasterSlaveServerGroup("groupMasterSlaveServerGroup", {
 *     loadBalancerId: instanceApplicationLoadBalancer.id,
 *     servers: [
 *         {
 *             serverId: instanceInstance[0].id,
 *             port: 100,
 *             weight: 100,
 *             serverType: "Master",
 *         },
 *         {
 *             serverId: instanceInstance[1].id,
 *             port: 100,
 *             weight: 100,
 *             serverType: "Slave",
 *         },
 *     ],
 * });
 * const sampleDs = instanceApplicationLoadBalancer.id.apply(id => alicloud.slb.getMasterSlaveServerGroups({
 *     loadBalancerId: id,
 * }));
 * export const firstSlbServerGroupId = sampleDs.groups[0].id;
 * ```
 */
export function getMasterSlaveServerGroups(args: GetMasterSlaveServerGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetMasterSlaveServerGroupsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:slb/getMasterSlaveServerGroups:getMasterSlaveServerGroups", {
        "ids": args.ids,
        "loadBalancerId": args.loadBalancerId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getMasterSlaveServerGroups.
 */
export interface GetMasterSlaveServerGroupsArgs {
    /**
     * A list of master slave server group IDs to filter results.
     */
    readonly ids?: string[];
    /**
     * ID of the SLB.
     */
    readonly loadBalancerId: string;
    /**
     * A regex string to filter results by master slave server group name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getMasterSlaveServerGroups.
 */
export interface GetMasterSlaveServerGroupsResult {
    /**
     * A list of SLB master slave server groups. Each element contains the following attributes:
     */
    readonly groups: outputs.slb.GetMasterSlaveServerGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of SLB master slave server groups IDs.
     */
    readonly ids: string[];
    readonly loadBalancerId: string;
    readonly nameRegex?: string;
    /**
     * A list of SLB master slave server groups names.
     */
    readonly names: string[];
    readonly outputFile?: string;
}
