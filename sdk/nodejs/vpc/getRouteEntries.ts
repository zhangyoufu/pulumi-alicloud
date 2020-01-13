// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of Route Entries owned by an Alibaba Cloud account.
 * 
 * > **NOTE:** Available in 1.37.0+.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-testAccRouteEntryConfig";
 * 
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultInstanceTypes = alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones[0].id,
 *     cpuCoreCount: 1,
 *     memorySize: 2,
 * });
 * const defaultImages = alicloud.ecs.getImages({
 *     mostRecent: true,
 *     nameRegex: "^ubuntu_18.*64",
 *     owners: "system",
 * });
 * const fooNetwork = new alicloud.vpc.Network("foo", {
 *     cidrBlock: "10.1.0.0/21",
 * });
 * const fooSwitch = new alicloud.vpc.Switch("foo", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "10.1.1.0/24",
 *     vpcId: fooNetwork.id,
 * });
 * const tfTestFoo = new alicloud.ecs.SecurityGroup("tfTestFoo", {
 *     description: "foo",
 *     vpcId: fooNetwork.id,
 * });
 * const fooInstance = new alicloud.ecs.Instance("foo", {
 *     allocatePublicIp: true,
 *     imageId: defaultImages.images[0].id,
 *     // series III
 *     instanceChargeType: "PostPaid",
 *     instanceName: name,
 *     instanceType: defaultInstanceTypes.instanceTypes[0].id,
 *     internetChargeType: "PayByTraffic",
 *     internetMaxBandwidthOut: 5,
 *     // cn-beijing
 *     securityGroups: [tfTestFoo.id],
 *     systemDiskCategory: "cloudEfficiency",
 *     vswitchId: fooSwitch.id,
 * });
 * const fooRouteEntry = new alicloud.vpc.RouteEntry("foo", {
 *     destinationCidrblock: "172.11.1.1/32",
 *     nexthopId: fooInstance.id,
 *     nexthopType: "Instance",
 *     routeTableId: fooNetwork.routeTableId,
 * });
 * const ingress = new alicloud.ecs.SecurityGroupRule("ingress", {
 *     cidrIp: "0.0.0.0/0",
 *     ipProtocol: "tcp",
 *     nicType: "intranet",
 *     policy: "accept",
 *     portRange: "22/22",
 *     priority: 1,
 *     securityGroupId: tfTestFoo.id,
 *     type: "ingress",
 * });
 * const fooRouteEntries = fooRouteEntry.routeTableId.apply(routeTableId => alicloud.vpc.getRouteEntries({
 *     routeTableId: routeTableId,
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/route_entries.html.markdown.
 */
export function getRouteEntries(args: GetRouteEntriesArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteEntriesResult> & GetRouteEntriesResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetRouteEntriesResult> = pulumi.runtime.invoke("alicloud:vpc/getRouteEntries:getRouteEntries", {
        "cidrBlock": args.cidrBlock,
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
        "routeTableId": args.routeTableId,
        "type": args.type,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getRouteEntries.
 */
export interface GetRouteEntriesArgs {
    /**
     * The destination CIDR block of the route entry.
     */
    readonly cidrBlock?: string;
    /**
     * The instance ID of the next hop.
     */
    readonly instanceId?: string;
    readonly outputFile?: string;
    /**
     * The ID of the router table to which the route entry belongs.
     */
    readonly routeTableId: string;
    /**
     * The type of the route entry.
     */
    readonly type?: string;
}

/**
 * A collection of values returned by getRouteEntries.
 */
export interface GetRouteEntriesResult {
    /**
     * The destination CIDR block of the route entry.
     */
    readonly cidrBlock?: string;
    /**
     * A list of Route Entries. Each element contains the following attributes:
     */
    readonly entries: outputs.vpc.GetRouteEntriesEntry[];
    /**
     * The instance ID of the next hop.
     */
    readonly instanceId?: string;
    readonly outputFile?: string;
    /**
     * The ID of the router table to which the route entry belongs.
     */
    readonly routeTableId: string;
    /**
     * The type of the route entry.
     */
    readonly type?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
