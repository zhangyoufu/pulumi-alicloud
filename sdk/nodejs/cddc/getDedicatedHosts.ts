// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Cddc Dedicated Hosts of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.147.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.cddc.getDedicatedHosts({
 *     dedicatedHostGroupId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 * });
 * export const cddcDedicatedHostId1 = ids.then(ids => ids.hosts?.[0]?.id);
 * const status = alicloud.cddc.getDedicatedHosts({
 *     dedicatedHostGroupId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 *     status: "1",
 * });
 * export const cddcDedicatedHostId2 = status.then(status => status.hosts?.[0]?.id);
 * const zoneId = alicloud.cddc.getDedicatedHosts({
 *     dedicatedHostGroupId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 *     zoneId: "example_value",
 * });
 * export const cddcDedicatedHostId3 = zoneId.then(zoneId => zoneId.hosts?.[0]?.id);
 * const allocationStatus = alicloud.cddc.getDedicatedHosts({
 *     dedicatedHostGroupId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 *     allocationStatus: "Allocatable",
 * });
 * export const cddcDedicatedHostId4 = allocationStatus.then(allocationStatus => allocationStatus.hosts?.[0]?.id);
 * const hostType = alicloud.cddc.getDedicatedHosts({
 *     dedicatedHostGroupId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 *     hostType: "dhg_cloud_ssd",
 * });
 * export const cddcDedicatedHostId5 = hostType.then(hostType => hostType.hosts?.[0]?.id);
 * ```
 */
export function getDedicatedHosts(args: GetDedicatedHostsArgs, opts?: pulumi.InvokeOptions): Promise<GetDedicatedHostsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cddc/getDedicatedHosts:getDedicatedHosts", {
        "allocationStatus": args.allocationStatus,
        "dedicatedHostGroupId": args.dedicatedHostGroupId,
        "enableDetails": args.enableDetails,
        "hostType": args.hostType,
        "ids": args.ids,
        "orderId": args.orderId,
        "outputFile": args.outputFile,
        "status": args.status,
        "tags": args.tags,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDedicatedHosts.
 */
export interface GetDedicatedHostsArgs {
    /**
     * Specifies whether instances can be created on the host. Valid values: `Allocatable` or `Suspended`. `Allocatable`: Instances can be created on the host. `Suspended`: Instances cannot be created on the host.
     */
    allocationStatus?: string;
    /**
     * The ID of the dedicated cluster.
     */
    dedicatedHostGroupId: string;
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: boolean;
    /**
     * The storage type of the host. Valid values: `dhgLocalSsd` or `dhgCloudSsd`. `dhgLocalSsd`: specifies that the host uses local SSDs. `dhgCloudSsd`: specifies that the host uses enhanced SSDs (ESSDs).
     */
    hostType?: string;
    /**
     * A list of Dedicated Host IDs.
     */
    ids?: string[];
    /**
     * The ID of the order.
     */
    orderId?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The state of the host. Valid values: 
     * * `0:` The host is being created.
     */
    status?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: {[key: string]: string};
    /**
     * The ID of the zone.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by getDedicatedHosts.
 */
export interface GetDedicatedHostsResult {
    readonly allocationStatus?: string;
    readonly dedicatedHostGroupId: string;
    readonly enableDetails?: boolean;
    readonly hostType?: string;
    readonly hosts: outputs.cddc.GetDedicatedHostsHost[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly orderId?: string;
    readonly outputFile?: string;
    readonly status?: string;
    readonly tags?: {[key: string]: string};
    readonly zoneId?: string;
}
/**
 * This data source provides the Cddc Dedicated Hosts of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.147.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.cddc.getDedicatedHosts({
 *     dedicatedHostGroupId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 * });
 * export const cddcDedicatedHostId1 = ids.then(ids => ids.hosts?.[0]?.id);
 * const status = alicloud.cddc.getDedicatedHosts({
 *     dedicatedHostGroupId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 *     status: "1",
 * });
 * export const cddcDedicatedHostId2 = status.then(status => status.hosts?.[0]?.id);
 * const zoneId = alicloud.cddc.getDedicatedHosts({
 *     dedicatedHostGroupId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 *     zoneId: "example_value",
 * });
 * export const cddcDedicatedHostId3 = zoneId.then(zoneId => zoneId.hosts?.[0]?.id);
 * const allocationStatus = alicloud.cddc.getDedicatedHosts({
 *     dedicatedHostGroupId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 *     allocationStatus: "Allocatable",
 * });
 * export const cddcDedicatedHostId4 = allocationStatus.then(allocationStatus => allocationStatus.hosts?.[0]?.id);
 * const hostType = alicloud.cddc.getDedicatedHosts({
 *     dedicatedHostGroupId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 *     hostType: "dhg_cloud_ssd",
 * });
 * export const cddcDedicatedHostId5 = hostType.then(hostType => hostType.hosts?.[0]?.id);
 * ```
 */
export function getDedicatedHostsOutput(args: GetDedicatedHostsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDedicatedHostsResult> {
    return pulumi.output(args).apply((a: any) => getDedicatedHosts(a, opts))
}

/**
 * A collection of arguments for invoking getDedicatedHosts.
 */
export interface GetDedicatedHostsOutputArgs {
    /**
     * Specifies whether instances can be created on the host. Valid values: `Allocatable` or `Suspended`. `Allocatable`: Instances can be created on the host. `Suspended`: Instances cannot be created on the host.
     */
    allocationStatus?: pulumi.Input<string>;
    /**
     * The ID of the dedicated cluster.
     */
    dedicatedHostGroupId: pulumi.Input<string>;
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * The storage type of the host. Valid values: `dhgLocalSsd` or `dhgCloudSsd`. `dhgLocalSsd`: specifies that the host uses local SSDs. `dhgCloudSsd`: specifies that the host uses enhanced SSDs (ESSDs).
     */
    hostType?: pulumi.Input<string>;
    /**
     * A list of Dedicated Host IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the order.
     */
    orderId?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The state of the host. Valid values: 
     * * `0:` The host is being created.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the zone.
     */
    zoneId?: pulumi.Input<string>;
}
