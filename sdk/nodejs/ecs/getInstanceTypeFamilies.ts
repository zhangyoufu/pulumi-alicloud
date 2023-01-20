// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the ECS instance type families of Alibaba Cloud.
 *
 * > **NOTE:** Available in 1.54.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.ecs.getInstanceTypeFamilies({
 *     instanceChargeType: "PrePaid",
 * });
 * export const firstInstanceTypeFamilyId = _default.then(_default => _default.families?.[0]?.id);
 * export const instanceIds = _default.then(_default => _default.ids);
 * ```
 */
export function getInstanceTypeFamilies(args?: GetInstanceTypeFamiliesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceTypeFamiliesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ecs/getInstanceTypeFamilies:getInstanceTypeFamilies", {
        "generation": args.generation,
        "instanceChargeType": args.instanceChargeType,
        "outputFile": args.outputFile,
        "spotStrategy": args.spotStrategy,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceTypeFamilies.
 */
export interface GetInstanceTypeFamiliesArgs {
    /**
     * The generation of the instance type family, Valid values: `ecs-1`, `ecs-2`, `ecs-3`, `ecs-4`, `ecs-5`, `ecs-6`. For more information, see [Instance type families](https://www.alibabacloud.com/help/doc-detail/25378.htm).
     */
    generation?: string;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
     */
    instanceChargeType?: string;
    outputFile?: string;
    /**
     * Filter the results by ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
     */
    spotStrategy?: string;
    /**
     * The Zone to launch the instance.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by getInstanceTypeFamilies.
 */
export interface GetInstanceTypeFamiliesResult {
    readonly families: outputs.ecs.GetInstanceTypeFamiliesFamily[];
    /**
     * The generation of the instance type family.
     */
    readonly generation?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of instance type family IDs.
     */
    readonly ids: string[];
    readonly instanceChargeType?: string;
    readonly outputFile?: string;
    readonly spotStrategy?: string;
    readonly zoneId?: string;
}
/**
 * This data source provides the ECS instance type families of Alibaba Cloud.
 *
 * > **NOTE:** Available in 1.54.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.ecs.getInstanceTypeFamilies({
 *     instanceChargeType: "PrePaid",
 * });
 * export const firstInstanceTypeFamilyId = _default.then(_default => _default.families?.[0]?.id);
 * export const instanceIds = _default.then(_default => _default.ids);
 * ```
 */
export function getInstanceTypeFamiliesOutput(args?: GetInstanceTypeFamiliesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceTypeFamiliesResult> {
    return pulumi.output(args).apply((a: any) => getInstanceTypeFamilies(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceTypeFamilies.
 */
export interface GetInstanceTypeFamiliesOutputArgs {
    /**
     * The generation of the instance type family, Valid values: `ecs-1`, `ecs-2`, `ecs-3`, `ecs-4`, `ecs-5`, `ecs-6`. For more information, see [Instance type families](https://www.alibabacloud.com/help/doc-detail/25378.htm).
     */
    generation?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
     */
    instanceChargeType?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * Filter the results by ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
     */
    spotStrategy?: pulumi.Input<string>;
    /**
     * The Zone to launch the instance.
     */
    zoneId?: pulumi.Input<string>;
}
