// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This data source provides the CR Endpoint Acl Service of the current Alibaba Cloud user.
 *
 * For information about Event Bridge and how to use it, see [What is CR Endpoint Acl](https://www.alibabacloud.com/help/en/doc-detail/142246.htm).
 *
 * > **NOTE:** Available in v1.139.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.cr.getEndpointAclService({
 *     enable: true,
 *     endpointType: "internet",
 *     instanceId: "example_id",
 *     moduleName: "Registry",
 * });
 * ```
 */
export function getEndpointAclService(args: GetEndpointAclServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetEndpointAclServiceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cr/getEndpointAclService:getEndpointAclService", {
        "enable": args.enable,
        "endpointType": args.endpointType,
        "instanceId": args.instanceId,
        "moduleName": args.moduleName,
    }, opts);
}

/**
 * A collection of arguments for invoking getEndpointAclService.
 */
export interface GetEndpointAclServiceArgs {
    /**
     * Whether to enable Acl Service.  Valid values: `true` and `false`.
     */
    enable: boolean;
    /**
     * The type of endpoint. Valid values: `internet`.
     */
    endpointType: string;
    /**
     * The ID of the CR Instance.
     */
    instanceId: string;
    /**
     * The ModuleName. Valid values: `Registry`.
     */
    moduleName?: string;
}

/**
 * A collection of values returned by getEndpointAclService.
 */
export interface GetEndpointAclServiceResult {
    readonly enable: boolean;
    readonly endpointType: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly moduleName?: string;
    readonly status: string;
}
/**
 * This data source provides the CR Endpoint Acl Service of the current Alibaba Cloud user.
 *
 * For information about Event Bridge and how to use it, see [What is CR Endpoint Acl](https://www.alibabacloud.com/help/en/doc-detail/142246.htm).
 *
 * > **NOTE:** Available in v1.139.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.cr.getEndpointAclService({
 *     enable: true,
 *     endpointType: "internet",
 *     instanceId: "example_id",
 *     moduleName: "Registry",
 * });
 * ```
 */
export function getEndpointAclServiceOutput(args: GetEndpointAclServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEndpointAclServiceResult> {
    return pulumi.output(args).apply((a: any) => getEndpointAclService(a, opts))
}

/**
 * A collection of arguments for invoking getEndpointAclService.
 */
export interface GetEndpointAclServiceOutputArgs {
    /**
     * Whether to enable Acl Service.  Valid values: `true` and `false`.
     */
    enable: pulumi.Input<boolean>;
    /**
     * The type of endpoint. Valid values: `internet`.
     */
    endpointType: pulumi.Input<string>;
    /**
     * The ID of the CR Instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The ModuleName. Valid values: `Registry`.
     */
    moduleName?: pulumi.Input<string>;
}
