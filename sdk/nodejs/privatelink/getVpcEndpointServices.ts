// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Privatelink Vpc Endpoint Services of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.109.0+.
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
 * const example = alicloud.privatelink.getVpcEndpointServices({
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstPrivatelinkVpcEndpointServiceId = example.then(example => example.services?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getVpcEndpointServices(args?: GetVpcEndpointServicesArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcEndpointServicesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:privatelink/getVpcEndpointServices:getVpcEndpointServices", {
        "autoAcceptConnection": args.autoAcceptConnection,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "serviceBusinessStatus": args.serviceBusinessStatus,
        "status": args.status,
        "vpcEndpointServiceName": args.vpcEndpointServiceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcEndpointServices.
 */
export interface GetVpcEndpointServicesArgs {
    /**
     * Whether to automatically accept terminal node connections..
     */
    autoAcceptConnection?: boolean;
    /**
     * A list of Vpc Endpoint Service IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Vpc Endpoint Service name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The business status of the terminal node service..
     */
    serviceBusinessStatus?: string;
    /**
     * The Status of Vpc Endpoint Service.
     */
    status?: string;
    /**
     * The name of Vpc Endpoint Service.
     */
    vpcEndpointServiceName?: string;
}

/**
 * A collection of values returned by getVpcEndpointServices.
 */
export interface GetVpcEndpointServicesResult {
    readonly autoAcceptConnection?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly serviceBusinessStatus?: string;
    readonly services: outputs.privatelink.GetVpcEndpointServicesService[];
    readonly status?: string;
    readonly vpcEndpointServiceName?: string;
}
/**
 * This data source provides the Privatelink Vpc Endpoint Services of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.109.0+.
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
 * const example = alicloud.privatelink.getVpcEndpointServices({
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstPrivatelinkVpcEndpointServiceId = example.then(example => example.services?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getVpcEndpointServicesOutput(args?: GetVpcEndpointServicesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcEndpointServicesResult> {
    return pulumi.output(args).apply((a: any) => getVpcEndpointServices(a, opts))
}

/**
 * A collection of arguments for invoking getVpcEndpointServices.
 */
export interface GetVpcEndpointServicesOutputArgs {
    /**
     * Whether to automatically accept terminal node connections..
     */
    autoAcceptConnection?: pulumi.Input<boolean>;
    /**
     * A list of Vpc Endpoint Service IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Vpc Endpoint Service name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The business status of the terminal node service..
     */
    serviceBusinessStatus?: pulumi.Input<string>;
    /**
     * The Status of Vpc Endpoint Service.
     */
    status?: pulumi.Input<string>;
    /**
     * The name of Vpc Endpoint Service.
     */
    vpcEndpointServiceName?: pulumi.Input<string>;
}
