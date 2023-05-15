// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Cms Namespaces of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.171.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.cms.getNamespaces({
 *     ids: ["example_id"],
 * });
 * export const cmsNamespaceId1 = ids.then(ids => ids.namespaces?.[0]?.id);
 * ```
 */
export function getNamespaces(args?: GetNamespacesArgs, opts?: pulumi.InvokeOptions): Promise<GetNamespacesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cms/getNamespaces:getNamespaces", {
        "ids": args.ids,
        "keyword": args.keyword,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
    }, opts);
}

/**
 * A collection of arguments for invoking getNamespaces.
 */
export interface GetNamespacesArgs {
    /**
     * A list of Namespace IDs.
     */
    ids?: string[];
    /**
     * The keywords of the `namespace` or `description` of the namespace.
     */
    keyword?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
}

/**
 * A collection of values returned by getNamespaces.
 */
export interface GetNamespacesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly keyword?: string;
    readonly namespaces: outputs.cms.GetNamespacesNamespace[];
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
}
/**
 * This data source provides the Cms Namespaces of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.171.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.cms.getNamespaces({
 *     ids: ["example_id"],
 * });
 * export const cmsNamespaceId1 = ids.then(ids => ids.namespaces?.[0]?.id);
 * ```
 */
export function getNamespacesOutput(args?: GetNamespacesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNamespacesResult> {
    return pulumi.output(args).apply((a: any) => getNamespaces(a, opts))
}

/**
 * A collection of arguments for invoking getNamespaces.
 */
export interface GetNamespacesOutputArgs {
    /**
     * A list of Namespace IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The keywords of the `namespace` or `description` of the namespace.
     */
    keyword?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
}
