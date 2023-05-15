// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Privatelink Vpc Endpoint Service Resources of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.110.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.privatelink.getVpcEndpointServiceResources({
 *     serviceId: "epsrv-gw8ii1xxxx",
 * });
 * export const firstPrivatelinkVpcEndpointServiceResourceId = example.then(example => example.resources?.[0]?.id);
 * ```
 */
export function getVpcEndpointServiceResources(args: GetVpcEndpointServiceResourcesArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcEndpointServiceResourcesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:privatelink/getVpcEndpointServiceResources:getVpcEndpointServiceResources", {
        "outputFile": args.outputFile,
        "serviceId": args.serviceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcEndpointServiceResources.
 */
export interface GetVpcEndpointServiceResourcesArgs {
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The ID of Vpc Endpoint Service.
     */
    serviceId: string;
}

/**
 * A collection of values returned by getVpcEndpointServiceResources.
 */
export interface GetVpcEndpointServiceResourcesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly resources: outputs.privatelink.GetVpcEndpointServiceResourcesResource[];
    readonly serviceId: string;
}
/**
 * This data source provides the Privatelink Vpc Endpoint Service Resources of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.110.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.privatelink.getVpcEndpointServiceResources({
 *     serviceId: "epsrv-gw8ii1xxxx",
 * });
 * export const firstPrivatelinkVpcEndpointServiceResourceId = example.then(example => example.resources?.[0]?.id);
 * ```
 */
export function getVpcEndpointServiceResourcesOutput(args: GetVpcEndpointServiceResourcesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcEndpointServiceResourcesResult> {
    return pulumi.output(args).apply((a: any) => getVpcEndpointServiceResources(a, opts))
}

/**
 * A collection of arguments for invoking getVpcEndpointServiceResources.
 */
export interface GetVpcEndpointServiceResourcesOutputArgs {
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of Vpc Endpoint Service.
     */
    serviceId: pulumi.Input<string>;
}
