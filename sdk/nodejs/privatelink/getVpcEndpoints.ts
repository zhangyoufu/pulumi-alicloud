// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Privatelink Vpc Endpoints of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.109.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.privatelink.getVpcEndpoints({
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstPrivatelinkVpcEndpointId = example.then(example => example.endpoints?.[0]?.id);
 * ```
 */
export function getVpcEndpoints(args?: GetVpcEndpointsArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcEndpointsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:privatelink/getVpcEndpoints:getVpcEndpoints", {
        "connectionStatus": args.connectionStatus,
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "serviceName": args.serviceName,
        "status": args.status,
        "vpcEndpointName": args.vpcEndpointName,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcEndpoints.
 */
export interface GetVpcEndpointsArgs {
    /**
     * The status of Connection.
     */
    connectionStatus?: string;
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: boolean;
    /**
     * A list of Vpc Endpoint IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Vpc Endpoint name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The name of the terminal node service associated with the terminal node.
     */
    serviceName?: string;
    /**
     * The status of Vpc Endpoint.
     */
    status?: string;
    /**
     * The name of Vpc Endpoint.
     */
    vpcEndpointName?: string;
    /**
     * The private network to which the terminal node belongs.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getVpcEndpoints.
 */
export interface GetVpcEndpointsResult {
    readonly connectionStatus?: string;
    readonly enableDetails?: boolean;
    readonly endpoints: outputs.privatelink.GetVpcEndpointsEndpoint[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly serviceName?: string;
    readonly status?: string;
    readonly vpcEndpointName?: string;
    readonly vpcId?: string;
}
/**
 * This data source provides the Privatelink Vpc Endpoints of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.109.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.privatelink.getVpcEndpoints({
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstPrivatelinkVpcEndpointId = example.then(example => example.endpoints?.[0]?.id);
 * ```
 */
export function getVpcEndpointsOutput(args?: GetVpcEndpointsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcEndpointsResult> {
    return pulumi.output(args).apply((a: any) => getVpcEndpoints(a, opts))
}

/**
 * A collection of arguments for invoking getVpcEndpoints.
 */
export interface GetVpcEndpointsOutputArgs {
    /**
     * The status of Connection.
     */
    connectionStatus?: pulumi.Input<string>;
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of Vpc Endpoint IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Vpc Endpoint name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The name of the terminal node service associated with the terminal node.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * The status of Vpc Endpoint.
     */
    status?: pulumi.Input<string>;
    /**
     * The name of Vpc Endpoint.
     */
    vpcEndpointName?: pulumi.Input<string>;
    /**
     * The private network to which the terminal node belongs.
     */
    vpcId?: pulumi.Input<string>;
}
