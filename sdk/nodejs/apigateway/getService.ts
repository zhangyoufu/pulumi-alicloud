// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Using this data source can enable API gateway service automatically. If the service has been enabled, it will return `Opened`.
 *
 * For information about API Gateway and how to use it, see [What is API Gateway](https://www.alibabacloud.com/help/product/29462.htm).
 *
 * > **NOTE:** Available in v1.96.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const open = alicloud.apigateway.getService({
 *     enable: "On",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getService(args?: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:apigateway/getService:getService", {
        "enable": args.enable,
    }, opts);
}

/**
 * A collection of arguments for invoking getService.
 */
export interface GetServiceArgs {
    /**
     * Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: "On" or "Off". Default to "Off".
     *
     * > **NOTE:** Setting `enable = "On"` to open the API gateway service that means you have read and agreed the [API Gateway Terms of Service](https://help.aliyun.com/document_detail/35391.html). The service can not closed once it is opened.
     */
    enable?: string;
}

/**
 * A collection of values returned by getService.
 */
export interface GetServiceResult {
    readonly enable?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The current service enable status.
     */
    readonly status: string;
}
/**
 * Using this data source can enable API gateway service automatically. If the service has been enabled, it will return `Opened`.
 *
 * For information about API Gateway and how to use it, see [What is API Gateway](https://www.alibabacloud.com/help/product/29462.htm).
 *
 * > **NOTE:** Available in v1.96.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const open = alicloud.apigateway.getService({
 *     enable: "On",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getServiceOutput(args?: GetServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceResult> {
    return pulumi.output(args).apply((a: any) => getService(a, opts))
}

/**
 * A collection of arguments for invoking getService.
 */
export interface GetServiceOutputArgs {
    /**
     * Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: "On" or "Off". Default to "Off".
     *
     * > **NOTE:** Setting `enable = "On"` to open the API gateway service that means you have read and agreed the [API Gateway Terms of Service](https://help.aliyun.com/document_detail/35391.html). The service can not closed once it is opened.
     */
    enable?: pulumi.Input<string>;
}
