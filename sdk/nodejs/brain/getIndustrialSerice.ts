// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Using this data source can open Brain Industrial service automatically. If the service has been opened, it will return opened.
 *
 * > **NOTE:** Available in v1.115.0+
 *
 * > **NOTE:** The Brain Industrial service is not support in the international site.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const open = alicloud.brain.getIndustrialSerice({
 *     enable: "On",
 * });
 * ```
 */
export function getIndustrialSerice(args?: GetIndustrialSericeArgs, opts?: pulumi.InvokeOptions): Promise<GetIndustrialSericeResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:brain/getIndustrialSerice:getIndustrialSerice", {
        "enable": args.enable,
    }, opts);
}

/**
 * A collection of arguments for invoking getIndustrialSerice.
 */
export interface GetIndustrialSericeArgs {
    /**
     * Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
     *
     * > **NOTE:** Setting `enable = "On"` to open the Brain Industrial service. The service can not closed once it is opened.
     */
    enable?: string;
}

/**
 * A collection of values returned by getIndustrialSerice.
 */
export interface GetIndustrialSericeResult {
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
 * Using this data source can open Brain Industrial service automatically. If the service has been opened, it will return opened.
 *
 * > **NOTE:** Available in v1.115.0+
 *
 * > **NOTE:** The Brain Industrial service is not support in the international site.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const open = alicloud.brain.getIndustrialSerice({
 *     enable: "On",
 * });
 * ```
 */
export function getIndustrialSericeOutput(args?: GetIndustrialSericeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIndustrialSericeResult> {
    return pulumi.output(args).apply((a: any) => getIndustrialSerice(a, opts))
}

/**
 * A collection of arguments for invoking getIndustrialSerice.
 */
export interface GetIndustrialSericeOutputArgs {
    /**
     * Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
     *
     * > **NOTE:** Setting `enable = "On"` to open the Brain Industrial service. The service can not closed once it is opened.
     */
    enable?: pulumi.Input<string>;
}
