// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Using this data source can open Private Zone service automatically. If the service has been opened, it will return opened.
 *
 * For information about Priavte Zone and how to use it, see [What is Private Zone](https://www.alibabacloud.com/help/en/product/64583.htm).
 *
 * > **NOTE:** Available in v1.114.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const open = alicloud.pvtz.getService({
 *     enable: "On",
 * });
 * ```
 */
export function getService(args?: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:pvtz/getService:getService", {
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
     * > **NOTE:** Setting `enable = "On"` to open the Private Zone service that means you have read and agreed the [Private Zone Terms of Service](https://help.aliyun.com/document_detail/65657.html). The service can not closed once it is opened.
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
 * Using this data source can open Private Zone service automatically. If the service has been opened, it will return opened.
 *
 * For information about Priavte Zone and how to use it, see [What is Private Zone](https://www.alibabacloud.com/help/en/product/64583.htm).
 *
 * > **NOTE:** Available in v1.114.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const open = alicloud.pvtz.getService({
 *     enable: "On",
 * });
 * ```
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
     * > **NOTE:** Setting `enable = "On"` to open the Private Zone service that means you have read and agreed the [Private Zone Terms of Service](https://help.aliyun.com/document_detail/65657.html). The service can not closed once it is opened.
     */
    enable?: pulumi.Input<string>;
}
