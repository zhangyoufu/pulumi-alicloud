// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list Container Registry Enterprise Edition namespaces on Alibaba Cloud.
 *
 * > **NOTE:** Available in v1.86.0+
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Declare the data source
 * const myNamespaces = pulumi.output(alicloud.cs.getRegistryEnterpriseNamespaces({
 *     instanceId: "cri-xxx",
 *     nameRegex: "my-namespace",
 *     outputFile: "my-namespace-json",
 * }, { async: true }));
 *
 * export const output = myNamespaces.namespaces;
 * ```
 */
export function getRegistryEnterpriseNamespaces(args: GetRegistryEnterpriseNamespacesArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistryEnterpriseNamespacesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:cs/getRegistryEnterpriseNamespaces:getRegistryEnterpriseNamespaces", {
        "ids": args.ids,
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegistryEnterpriseNamespaces.
 */
export interface GetRegistryEnterpriseNamespacesArgs {
    /**
     * A list of ids to filter results by namespace id.
     */
    readonly ids?: string[];
    /**
     * ID of Container Registry Enterprise Edition instance.
     */
    readonly instanceId: string;
    /**
     * A regex string to filter results by namespace name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getRegistryEnterpriseNamespaces.
 */
export interface GetRegistryEnterpriseNamespacesResult {
    /**
     * A list of matched Container Registry Enterprise Edition namespaces. Its element is a namespace uuid.
     */
    readonly ids: string[];
    /**
     * ID of Container Registry Enterprise Edition instance.
     */
    readonly instanceId: string;
    readonly nameRegex?: string;
    /**
     * A list of namespace names.
     */
    readonly names: string[];
    /**
     * A list of matched Container Registry Enterprise Edition namespaces. Each element contains the following attributes:
     */
    readonly namespaces: outputs.cs.GetRegistryEnterpriseNamespacesNamespace[];
    readonly outputFile?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
