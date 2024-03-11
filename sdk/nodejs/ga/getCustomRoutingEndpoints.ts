// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Global Accelerator (GA) Custom Routing Endpoints of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in 1.197.0+
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
 * const ids = alicloud.ga.getCustomRoutingEndpoints({
 *     ids: ["example_id"],
 *     acceleratorId: "your_accelerator_id",
 * });
 * export const gaCustomRoutingEndpointsId1 = ids.then(ids => ids.customRoutingEndpoints?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCustomRoutingEndpoints(args: GetCustomRoutingEndpointsArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomRoutingEndpointsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ga/getCustomRoutingEndpoints:getCustomRoutingEndpoints", {
        "acceleratorId": args.acceleratorId,
        "endpointGroupId": args.endpointGroupId,
        "ids": args.ids,
        "listenerId": args.listenerId,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
    }, opts);
}

/**
 * A collection of arguments for invoking getCustomRoutingEndpoints.
 */
export interface GetCustomRoutingEndpointsArgs {
    /**
     * The ID of the GA instance.
     */
    acceleratorId: string;
    /**
     * The ID of the endpoint group.
     */
    endpointGroupId?: string;
    /**
     * A list of Custom Routing Endpoint IDs.
     */
    ids?: string[];
    /**
     * The ID of the custom routing listener.
     */
    listenerId?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
}

/**
 * A collection of values returned by getCustomRoutingEndpoints.
 */
export interface GetCustomRoutingEndpointsResult {
    /**
     * The ID of the GA instance with which the endpoint is associated.
     */
    readonly acceleratorId: string;
    /**
     * A list of Custom Routing Endpoints. Each element contains the following attributes:
     */
    readonly customRoutingEndpoints: outputs.ga.GetCustomRoutingEndpointsCustomRoutingEndpoint[];
    /**
     * The ID of the Custom Routing Endpoint Group.
     */
    readonly endpointGroupId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    /**
     * The ID of the listener with which the endpoint is associated.
     */
    readonly listenerId?: string;
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
}
/**
 * This data source provides the Global Accelerator (GA) Custom Routing Endpoints of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in 1.197.0+
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
 * const ids = alicloud.ga.getCustomRoutingEndpoints({
 *     ids: ["example_id"],
 *     acceleratorId: "your_accelerator_id",
 * });
 * export const gaCustomRoutingEndpointsId1 = ids.then(ids => ids.customRoutingEndpoints?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCustomRoutingEndpointsOutput(args: GetCustomRoutingEndpointsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCustomRoutingEndpointsResult> {
    return pulumi.output(args).apply((a: any) => getCustomRoutingEndpoints(a, opts))
}

/**
 * A collection of arguments for invoking getCustomRoutingEndpoints.
 */
export interface GetCustomRoutingEndpointsOutputArgs {
    /**
     * The ID of the GA instance.
     */
    acceleratorId: pulumi.Input<string>;
    /**
     * The ID of the endpoint group.
     */
    endpointGroupId?: pulumi.Input<string>;
    /**
     * A list of Custom Routing Endpoint IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the custom routing listener.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
}
