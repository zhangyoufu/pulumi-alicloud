// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Express Connect Access Points of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.132.0+.
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
 * const ids = alicloud.expressconnect.getAccessPoints({
 *     ids: ["ap-cn-hangzhou-yh-C"],
 * });
 * export const expressConnectAccessPointId1 = ids.then(ids => ids.points?.[0]?.id);
 * const nameRegex = alicloud.expressconnect.getAccessPoints({
 *     nameRegex: "^杭州-",
 * });
 * export const expressConnectAccessPointId2 = nameRegex.then(nameRegex => nameRegex.points?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getAccessPoints(args?: GetAccessPointsArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessPointsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:expressconnect/getAccessPoints:getAccessPoints", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessPoints.
 */
export interface GetAccessPointsArgs {
    /**
     * A list of Access Point IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Access Point name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The Physical Connection to Which the Access Point State.
     */
    status?: string;
}

/**
 * A collection of values returned by getAccessPoints.
 */
export interface GetAccessPointsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly points: outputs.expressconnect.GetAccessPointsPoint[];
    readonly status?: string;
}
/**
 * This data source provides the Express Connect Access Points of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.132.0+.
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
 * const ids = alicloud.expressconnect.getAccessPoints({
 *     ids: ["ap-cn-hangzhou-yh-C"],
 * });
 * export const expressConnectAccessPointId1 = ids.then(ids => ids.points?.[0]?.id);
 * const nameRegex = alicloud.expressconnect.getAccessPoints({
 *     nameRegex: "^杭州-",
 * });
 * export const expressConnectAccessPointId2 = nameRegex.then(nameRegex => nameRegex.points?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getAccessPointsOutput(args?: GetAccessPointsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccessPointsResult> {
    return pulumi.output(args).apply((a: any) => getAccessPoints(a, opts))
}

/**
 * A collection of arguments for invoking getAccessPoints.
 */
export interface GetAccessPointsOutputArgs {
    /**
     * A list of Access Point IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Access Point name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The Physical Connection to Which the Access Point State.
     */
    status?: pulumi.Input<string>;
}
