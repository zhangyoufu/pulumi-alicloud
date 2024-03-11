// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Havips of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.120.0+.
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
 * const example = alicloud.vpc.getHavips({
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstHavipId = example.then(example => example.havips?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getHavips(args?: GetHavipsArgs, opts?: pulumi.InvokeOptions): Promise<GetHavipsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:vpc/getHavips:getHavips", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getHavips.
 */
export interface GetHavipsArgs {
    /**
     * A list of Ha Vip IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Ha Vip name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The status.
     */
    status?: string;
}

/**
 * A collection of values returned by getHavips.
 */
export interface GetHavipsResult {
    readonly havips: outputs.vpc.GetHavipsHavip[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly status?: string;
}
/**
 * This data source provides the Havips of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.120.0+.
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
 * const example = alicloud.vpc.getHavips({
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstHavipId = example.then(example => example.havips?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getHavipsOutput(args?: GetHavipsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHavipsResult> {
    return pulumi.output(args).apply((a: any) => getHavips(a, opts))
}

/**
 * A collection of arguments for invoking getHavips.
 */
export interface GetHavipsOutputArgs {
    /**
     * A list of Ha Vip IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Ha Vip name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status.
     */
    status?: pulumi.Input<string>;
}
