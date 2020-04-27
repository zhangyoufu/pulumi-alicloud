// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source compute file crc64 checksum.
 * 
 * > **NOTE:** Available in 1.59.0+.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const defaultFileCrc64Checksum = pulumi.output(alicloud.getFileCrc64Checksum({
 *     filename: "exampleFileName",
 * }, { async: true }));
 * 
 * export const fileCrc64Checksum = alicloud_file_crc64_checksum_defualt.checksum;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/file_crc64_checksum.html.markdown.
 */
export function getFileCrc64Checksum(args: GetFileCrc64ChecksumArgs, opts?: pulumi.InvokeOptions): Promise<GetFileCrc64ChecksumResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:index/getFileCrc64Checksum:getFileCrc64Checksum", {
        "filename": args.filename,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getFileCrc64Checksum.
 */
export interface GetFileCrc64ChecksumArgs {
    /**
     * The name of the file to be computed crc64 checksum.
     */
    readonly filename: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getFileCrc64Checksum.
 */
export interface GetFileCrc64ChecksumResult {
    /**
     * the file checksum of crc64.
     */
    readonly checksum: string;
    readonly filename: string;
    readonly outputFile?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
