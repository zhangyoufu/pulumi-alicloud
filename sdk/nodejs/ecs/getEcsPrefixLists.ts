// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Ecs Prefix Lists of the current Alibaba Cloud user.
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
 * const example = alicloud.ecs.getEcsPrefixLists({
 *     ids: ["E2RY53-xxxx"],
 *     nameRegex: "tf-testAcc",
 * });
 * export const outputId = example.then(example => example.lists?.[0]?.id);
 * ```
 */
export function getEcsPrefixLists(args?: GetEcsPrefixListsArgs, opts?: pulumi.InvokeOptions): Promise<GetEcsPrefixListsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ecs/getEcsPrefixLists:getEcsPrefixLists", {
        "addressFamily": args.addressFamily,
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getEcsPrefixLists.
 */
export interface GetEcsPrefixListsArgs {
    /**
     * The address family of the prefix list. Valid values:`IPv4`,`IPv6`.
     */
    addressFamily?: string;
    enableDetails?: boolean;
    /**
     * A list of Prefix List IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by `prefixListName`.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getEcsPrefixLists.
 */
export interface GetEcsPrefixListsResult {
    readonly addressFamily?: string;
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly lists: outputs.ecs.GetEcsPrefixListsList[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
}
/**
 * This data source provides the Ecs Prefix Lists of the current Alibaba Cloud user.
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
 * const example = alicloud.ecs.getEcsPrefixLists({
 *     ids: ["E2RY53-xxxx"],
 *     nameRegex: "tf-testAcc",
 * });
 * export const outputId = example.then(example => example.lists?.[0]?.id);
 * ```
 */
export function getEcsPrefixListsOutput(args?: GetEcsPrefixListsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEcsPrefixListsResult> {
    return pulumi.output(args).apply((a: any) => getEcsPrefixLists(a, opts))
}

/**
 * A collection of arguments for invoking getEcsPrefixLists.
 */
export interface GetEcsPrefixListsOutputArgs {
    /**
     * The address family of the prefix list. Valid values:`IPv4`,`IPv6`.
     */
    addressFamily?: pulumi.Input<string>;
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of Prefix List IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by `prefixListName`.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}
