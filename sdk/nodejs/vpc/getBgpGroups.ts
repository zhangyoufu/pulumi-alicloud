// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Vpc Bgp Groups of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.152.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.vpc.getBgpGroups({
 *     ids: ["example_value"],
 * });
 * export const vpcBgpGroupId1 = ids.then(ids => ids.groups?.[0]?.id);
 * const nameRegex = alicloud.vpc.getBgpGroups({
 *     nameRegex: "^my-BgpGroup",
 * });
 * export const vpcBgpGroupId2 = nameRegex.then(nameRegex => nameRegex.groups?.[0]?.id);
 * ```
 */
export function getBgpGroups(args?: GetBgpGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetBgpGroupsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:vpc/getBgpGroups:getBgpGroups", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "routerId": args.routerId,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getBgpGroups.
 */
export interface GetBgpGroupsArgs {
    /**
     * A list of Bgp Group IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Bgp Group name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The ID of the VBR.
     */
    routerId?: string;
    /**
     * The status of the resource.
     */
    status?: string;
}

/**
 * A collection of values returned by getBgpGroups.
 */
export interface GetBgpGroupsResult {
    readonly groups: outputs.vpc.GetBgpGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly routerId?: string;
    readonly status?: string;
}
/**
 * This data source provides the Vpc Bgp Groups of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.152.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.vpc.getBgpGroups({
 *     ids: ["example_value"],
 * });
 * export const vpcBgpGroupId1 = ids.then(ids => ids.groups?.[0]?.id);
 * const nameRegex = alicloud.vpc.getBgpGroups({
 *     nameRegex: "^my-BgpGroup",
 * });
 * export const vpcBgpGroupId2 = nameRegex.then(nameRegex => nameRegex.groups?.[0]?.id);
 * ```
 */
export function getBgpGroupsOutput(args?: GetBgpGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBgpGroupsResult> {
    return pulumi.output(args).apply((a: any) => getBgpGroups(a, opts))
}

/**
 * A collection of arguments for invoking getBgpGroups.
 */
export interface GetBgpGroupsOutputArgs {
    /**
     * A list of Bgp Group IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Bgp Group name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of the VBR.
     */
    routerId?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
}
