// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Event Bridge Event Buses of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.129.0+.
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
 * const ids = alicloud.eventbridge.getEventBuses({});
 * export const eventBridgeEventBusId1 = ids.then(ids => ids.buses?.[0]?.id);
 * const nameRegex = alicloud.eventbridge.getEventBuses({
 *     nameRegex: "^my-EventBus",
 * });
 * export const eventBridgeEventBusId2 = nameRegex.then(nameRegex => nameRegex.buses?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getEventBuses(args?: GetEventBusesArgs, opts?: pulumi.InvokeOptions): Promise<GetEventBusesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:eventbridge/getEventBuses:getEventBuses", {
        "eventBusType": args.eventBusType,
        "ids": args.ids,
        "namePrefix": args.namePrefix,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getEventBuses.
 */
export interface GetEventBusesArgs {
    /**
     * The event bus type.
     */
    eventBusType?: string;
    /**
     * A list of Event Bus IDs. Its element value is same as Event Bus Name.
     */
    ids?: string[];
    /**
     * The name prefix.
     */
    namePrefix?: string;
    /**
     * A regex string to filter results by Event Bus name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getEventBuses.
 */
export interface GetEventBusesResult {
    readonly buses: outputs.eventbridge.GetEventBusesBus[];
    readonly eventBusType?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly namePrefix?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
}
/**
 * This data source provides the Event Bridge Event Buses of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.129.0+.
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
 * const ids = alicloud.eventbridge.getEventBuses({});
 * export const eventBridgeEventBusId1 = ids.then(ids => ids.buses?.[0]?.id);
 * const nameRegex = alicloud.eventbridge.getEventBuses({
 *     nameRegex: "^my-EventBus",
 * });
 * export const eventBridgeEventBusId2 = nameRegex.then(nameRegex => nameRegex.buses?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getEventBusesOutput(args?: GetEventBusesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEventBusesResult> {
    return pulumi.output(args).apply((a: any) => getEventBuses(a, opts))
}

/**
 * A collection of arguments for invoking getEventBuses.
 */
export interface GetEventBusesOutputArgs {
    /**
     * The event bus type.
     */
    eventBusType?: pulumi.Input<string>;
    /**
     * A list of Event Bus IDs. Its element value is same as Event Bus Name.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name prefix.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Event Bus name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
