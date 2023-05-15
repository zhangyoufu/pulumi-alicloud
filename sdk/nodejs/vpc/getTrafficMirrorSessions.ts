// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Vpc Traffic Mirror Sessions of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.142.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 * });
 * export const vpcTrafficMirrorSessionId1 = ids.then(ids => ids.sessions?.[0]?.id);
 * const nameRegex = alicloud.vpc.getTrafficMirrorSessions({
 *     nameRegex: "^my-TrafficMirrorSession",
 * });
 * export const vpcTrafficMirrorSessionId2 = nameRegex.then(nameRegex => nameRegex.sessions?.[0]?.id);
 * const enabled = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 *     enabled: false,
 * });
 * export const vpcTrafficMirrorSessionId3 = enabled.then(enabled => enabled.sessions?.[0]?.id);
 * const priority = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 *     priority: 1,
 * });
 * export const vpcTrafficMirrorSessionId4 = priority.then(priority => priority.sessions?.[0]?.id);
 * const filterId = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 *     trafficMirrorFilterId: "example_value",
 * });
 * export const vpcTrafficMirrorSessionId5 = filterId.then(filterId => filterId.sessions?.[0]?.id);
 * const sessionName = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 *     trafficMirrorSessionName: "example_value",
 * });
 * export const vpcTrafficMirrorSessionId6 = sessionName.then(sessionName => sessionName.sessions?.[0]?.id);
 * const sourceId = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 *     trafficMirrorSourceId: "example_value",
 * });
 * export const vpcTrafficMirrorSessionId7 = sourceId.then(sourceId => sourceId.sessions?.[0]?.id);
 * const targetId = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 *     trafficMirrorTargetId: "example_value",
 * });
 * export const vpcTrafficMirrorSessionId8 = targetId.then(targetId => targetId.sessions?.[0]?.id);
 * const status = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 *     status: "Created",
 * });
 * export const vpcTrafficMirrorSessionId9 = status.then(status => status.sessions?.[0]?.id);
 * ```
 */
export function getTrafficMirrorSessions(args?: GetTrafficMirrorSessionsArgs, opts?: pulumi.InvokeOptions): Promise<GetTrafficMirrorSessionsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:vpc/getTrafficMirrorSessions:getTrafficMirrorSessions", {
        "enabled": args.enabled,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "priority": args.priority,
        "status": args.status,
        "trafficMirrorFilterId": args.trafficMirrorFilterId,
        "trafficMirrorSessionName": args.trafficMirrorSessionName,
        "trafficMirrorSourceId": args.trafficMirrorSourceId,
        "trafficMirrorTargetId": args.trafficMirrorTargetId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTrafficMirrorSessions.
 */
export interface GetTrafficMirrorSessionsArgs {
    /**
     * Indicates whether traffic mirror sessions are enabled. default to `false`.
     */
    enabled?: boolean;
    /**
     * A list of Traffic Mirror Session IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Traffic Mirror Session name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The priority of the traffic mirror session. A smaller value indicates a higher priority.
     */
    priority?: number;
    /**
     * The state of the traffic mirror session. Valid values: `Creating`, `Created`, `Modifying` and `Deleting`.
     */
    status?: string;
    /**
     * The ID of the filter.
     */
    trafficMirrorFilterId?: string;
    /**
     * The name of the traffic mirror session.
     */
    trafficMirrorSessionName?: string;
    /**
     * The ID of the mirror source. You can specify only an elastic network interface (ENI) as the mirror source.
     */
    trafficMirrorSourceId?: string;
    /**
     * The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
     */
    trafficMirrorTargetId?: string;
}

/**
 * A collection of values returned by getTrafficMirrorSessions.
 */
export interface GetTrafficMirrorSessionsResult {
    readonly enabled?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly priority?: number;
    readonly sessions: outputs.vpc.GetTrafficMirrorSessionsSession[];
    readonly status?: string;
    readonly trafficMirrorFilterId?: string;
    readonly trafficMirrorSessionName?: string;
    readonly trafficMirrorSourceId?: string;
    readonly trafficMirrorTargetId?: string;
}
/**
 * This data source provides the Vpc Traffic Mirror Sessions of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.142.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 * });
 * export const vpcTrafficMirrorSessionId1 = ids.then(ids => ids.sessions?.[0]?.id);
 * const nameRegex = alicloud.vpc.getTrafficMirrorSessions({
 *     nameRegex: "^my-TrafficMirrorSession",
 * });
 * export const vpcTrafficMirrorSessionId2 = nameRegex.then(nameRegex => nameRegex.sessions?.[0]?.id);
 * const enabled = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 *     enabled: false,
 * });
 * export const vpcTrafficMirrorSessionId3 = enabled.then(enabled => enabled.sessions?.[0]?.id);
 * const priority = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 *     priority: 1,
 * });
 * export const vpcTrafficMirrorSessionId4 = priority.then(priority => priority.sessions?.[0]?.id);
 * const filterId = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 *     trafficMirrorFilterId: "example_value",
 * });
 * export const vpcTrafficMirrorSessionId5 = filterId.then(filterId => filterId.sessions?.[0]?.id);
 * const sessionName = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 *     trafficMirrorSessionName: "example_value",
 * });
 * export const vpcTrafficMirrorSessionId6 = sessionName.then(sessionName => sessionName.sessions?.[0]?.id);
 * const sourceId = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 *     trafficMirrorSourceId: "example_value",
 * });
 * export const vpcTrafficMirrorSessionId7 = sourceId.then(sourceId => sourceId.sessions?.[0]?.id);
 * const targetId = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 *     trafficMirrorTargetId: "example_value",
 * });
 * export const vpcTrafficMirrorSessionId8 = targetId.then(targetId => targetId.sessions?.[0]?.id);
 * const status = alicloud.vpc.getTrafficMirrorSessions({
 *     ids: ["example_id"],
 *     status: "Created",
 * });
 * export const vpcTrafficMirrorSessionId9 = status.then(status => status.sessions?.[0]?.id);
 * ```
 */
export function getTrafficMirrorSessionsOutput(args?: GetTrafficMirrorSessionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTrafficMirrorSessionsResult> {
    return pulumi.output(args).apply((a: any) => getTrafficMirrorSessions(a, opts))
}

/**
 * A collection of arguments for invoking getTrafficMirrorSessions.
 */
export interface GetTrafficMirrorSessionsOutputArgs {
    /**
     * Indicates whether traffic mirror sessions are enabled. default to `false`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * A list of Traffic Mirror Session IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Traffic Mirror Session name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The priority of the traffic mirror session. A smaller value indicates a higher priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * The state of the traffic mirror session. Valid values: `Creating`, `Created`, `Modifying` and `Deleting`.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the filter.
     */
    trafficMirrorFilterId?: pulumi.Input<string>;
    /**
     * The name of the traffic mirror session.
     */
    trafficMirrorSessionName?: pulumi.Input<string>;
    /**
     * The ID of the mirror source. You can specify only an elastic network interface (ENI) as the mirror source.
     */
    trafficMirrorSourceId?: pulumi.Input<string>;
    /**
     * The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
     */
    trafficMirrorTargetId?: pulumi.Input<string>;
}
