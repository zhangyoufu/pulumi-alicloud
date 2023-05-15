// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides Threat Detection Honeypot Image available to the user.[What is Honeypot Image](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-listavailablehoneypot)
 *
 * > **NOTE:** Available in 1.195.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.threatdetection.getHoneypotImages({
 *     ids: ["sha256:02882320c9a55303410127c5dc4ae2dc470150f9d7f2483102d994f5e5f4d9df"],
 *     nameRegex: "^meta",
 * });
 * export const alicloudThreatDetectionHoneypotImageExampleId = _default.then(_default => _default.images?.[0]?.id);
 * ```
 */
export function getHoneypotImages(args?: GetHoneypotImagesArgs, opts?: pulumi.InvokeOptions): Promise<GetHoneypotImagesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:threatdetection/getHoneypotImages:getHoneypotImages", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "nodeId": args.nodeId,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getHoneypotImages.
 */
export interface GetHoneypotImagesArgs {
    /**
     * A list of Honeypot Image IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Honeypot mirror nam.
     */
    nameRegex?: string;
    nodeId?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getHoneypotImages.
 */
export interface GetHoneypotImagesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of Honeypot Image IDs.
     */
    readonly ids: string[];
    /**
     * A list of Honeypot Image Entries. Each element contains the following attributes:
     */
    readonly images: outputs.threatdetection.GetHoneypotImagesImage[];
    readonly nameRegex?: string;
    /**
     * A list of name of Honeypot Images.
     */
    readonly names: string[];
    readonly nodeId?: string;
    readonly outputFile?: string;
}
/**
 * This data source provides Threat Detection Honeypot Image available to the user.[What is Honeypot Image](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-listavailablehoneypot)
 *
 * > **NOTE:** Available in 1.195.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.threatdetection.getHoneypotImages({
 *     ids: ["sha256:02882320c9a55303410127c5dc4ae2dc470150f9d7f2483102d994f5e5f4d9df"],
 *     nameRegex: "^meta",
 * });
 * export const alicloudThreatDetectionHoneypotImageExampleId = _default.then(_default => _default.images?.[0]?.id);
 * ```
 */
export function getHoneypotImagesOutput(args?: GetHoneypotImagesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHoneypotImagesResult> {
    return pulumi.output(args).apply((a: any) => getHoneypotImages(a, opts))
}

/**
 * A collection of arguments for invoking getHoneypotImages.
 */
export interface GetHoneypotImagesOutputArgs {
    /**
     * A list of Honeypot Image IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Honeypot mirror nam.
     */
    nameRegex?: pulumi.Input<string>;
    nodeId?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
