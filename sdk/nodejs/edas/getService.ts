// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Using this data source can open EDAS service automatically. If the service has been opened, it will return opened.
 *
 * For information about EDAS and how to use it, see [What is EDAS](https://www.alibabacloud.com/help/product/29500.htm).
 *
 * > **NOTE:** Available in v1.98.0+
 *
 * > **NOTE:** The EDAS service is not support to be open automatically in the international site.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const open = alicloud.edas.getService({
 *     enable: "On",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getService(args?: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:edas/getService:getService", {
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
     * > **NOTE:** Setting `enable = "On"` to open the EDAS service that means you have read and agreed the [EDAS Terms of Service](https://help.aliyun.com/document_detail/44633.html). The service can not closed once it is opened.
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
 * Using this data source can open EDAS service automatically. If the service has been opened, it will return opened.
 *
 * For information about EDAS and how to use it, see [What is EDAS](https://www.alibabacloud.com/help/product/29500.htm).
 *
 * > **NOTE:** Available in v1.98.0+
 *
 * > **NOTE:** The EDAS service is not support to be open automatically in the international site.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const open = alicloud.edas.getService({
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
     * > **NOTE:** Setting `enable = "On"` to open the EDAS service that means you have read and agreed the [EDAS Terms of Service](https://help.aliyun.com/document_detail/44633.html). The service can not closed once it is opened.
     */
    enable?: pulumi.Input<string>;
}
