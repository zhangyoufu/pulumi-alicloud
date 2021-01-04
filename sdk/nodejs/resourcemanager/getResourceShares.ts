// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Resource Manager Resource Shares of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.111.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.resourcemanager.getResourceShares({
 *     resourceShareOwner: "Self",
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstResourceManagerResourceShareId = example.then(example => example.shares[0].id);
 * ```
 */
export function getResourceShares(args: GetResourceSharesArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceSharesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:resourcemanager/getResourceShares:getResourceShares", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "resourceShareName": args.resourceShareName,
        "resourceShareOwner": args.resourceShareOwner,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourceShares.
 */
export interface GetResourceSharesArgs {
    /**
     * A list of Resource Share IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by Resource Share name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The name of resource share.
     */
    readonly resourceShareName?: string;
    /**
     * The owner of resource share.
     */
    readonly resourceShareOwner: string;
    /**
     * The status of resource share.
     */
    readonly status?: string;
}

/**
 * A collection of values returned by getResourceShares.
 */
export interface GetResourceSharesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly resourceShareName?: string;
    readonly resourceShareOwner: string;
    readonly shares: outputs.resourcemanager.GetResourceSharesShare[];
    readonly status?: string;
}
