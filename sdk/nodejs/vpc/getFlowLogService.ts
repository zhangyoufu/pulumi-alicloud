// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Using this data source can open Vpc Flow Log service automatically. If the service has been opened, it will return opened.
 *
 * For information about Vpc Flow Log and how to use it, see [What is Vpc Flow Log](https://www.alibabacloud.com/help/en/vpc/developer-reference/api-openflowlog).
 *
 * > **NOTE:** Available since v1.209.0.
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
 * const default = alicloud.vpc.getFlowLogService({
 *     enable: "On",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getFlowLogService(args?: GetFlowLogServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetFlowLogServiceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:vpc/getFlowLogService:getFlowLogService", {
        "enable": args.enable,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlowLogService.
 */
export interface GetFlowLogServiceArgs {
    /**
     * Setting the value to `On` to enable the service. If has been enabled, return the result. Default value: `Off`. Valid values: `On` and `Off`.
     *
     * > **NOTE:** Setting `enable = "On"` to open the Vpc Flow Log service that means you have read and agreed the [Vpc Flow Log Terms of Service](https://help.aliyun.com/zh/vpc/support/vpc-terms-of-service). The service can not closed once it is opened.
     */
    enable?: string;
}

/**
 * A collection of values returned by getFlowLogService.
 */
export interface GetFlowLogServiceResult {
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
 * Using this data source can open Vpc Flow Log service automatically. If the service has been opened, it will return opened.
 *
 * For information about Vpc Flow Log and how to use it, see [What is Vpc Flow Log](https://www.alibabacloud.com/help/en/vpc/developer-reference/api-openflowlog).
 *
 * > **NOTE:** Available since v1.209.0.
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
 * const default = alicloud.vpc.getFlowLogService({
 *     enable: "On",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getFlowLogServiceOutput(args?: GetFlowLogServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlowLogServiceResult> {
    return pulumi.output(args).apply((a: any) => getFlowLogService(a, opts))
}

/**
 * A collection of arguments for invoking getFlowLogService.
 */
export interface GetFlowLogServiceOutputArgs {
    /**
     * Setting the value to `On` to enable the service. If has been enabled, return the result. Default value: `Off`. Valid values: `On` and `Off`.
     *
     * > **NOTE:** Setting `enable = "On"` to open the Vpc Flow Log service that means you have read and agreed the [Vpc Flow Log Terms of Service](https://help.aliyun.com/zh/vpc/support/vpc-terms-of-service). The service can not closed once it is opened.
     */
    enable?: pulumi.Input<string>;
}
