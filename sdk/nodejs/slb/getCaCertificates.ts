// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the CA certificate list.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const sampleDs = alicloud.slb.getCaCertificates({});
 * export const firstSlbCaCertificateId = sampleDs.then(sampleDs => sampleDs.certificates?.[0]?.id);
 * ```
 */
export function getCaCertificates(args?: GetCaCertificatesArgs, opts?: pulumi.InvokeOptions): Promise<GetCaCertificatesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:slb/getCaCertificates:getCaCertificates", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getCaCertificates.
 */
export interface GetCaCertificatesArgs {
    /**
     * A list of ca certificates IDs to filter results.
     */
    ids?: string[];
    /**
     * A regex string to filter results by ca certificate name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The Id of resource group which ca certificates belongs.
     */
    resourceGroupId?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getCaCertificates.
 */
export interface GetCaCertificatesResult {
    /**
     * A list of SLB ca certificates. Each element contains the following attributes:
     */
    readonly certificates: outputs.slb.GetCaCertificatesCertificate[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of SLB ca certificates IDs.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of SLB ca certificates names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The resource group Id of CA certificate.
     */
    readonly resourceGroupId?: string;
    /**
     * (Available in v1.66.0+) A mapping of tags to assign to the resource.
     */
    readonly tags?: {[key: string]: string};
}
/**
 * This data source provides the CA certificate list.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const sampleDs = alicloud.slb.getCaCertificates({});
 * export const firstSlbCaCertificateId = sampleDs.then(sampleDs => sampleDs.certificates?.[0]?.id);
 * ```
 */
export function getCaCertificatesOutput(args?: GetCaCertificatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCaCertificatesResult> {
    return pulumi.output(args).apply((a: any) => getCaCertificates(a, opts))
}

/**
 * A collection of arguments for invoking getCaCertificates.
 */
export interface GetCaCertificatesOutputArgs {
    /**
     * A list of ca certificates IDs to filter results.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by ca certificate name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The Id of resource group which ca certificates belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
