// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides ASM available versions in the specified region.
 *
 * > **NOTE:** Available in v1.161.0+.
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
 * const default = alicloud.servicemesh.getVersions({
 *     edition: "Default",
 * });
 * export const serviceMeshVersion = versions[0].version;
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getVersions(args?: GetVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetVersionsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:servicemesh/getVersions:getVersions", {
        "edition": args.edition,
        "ids": args.ids,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getVersions.
 */
export interface GetVersionsArgs {
    /**
     * The edition of the ASM instance.
     */
    edition?: string;
    /**
     * A list of ASM versions. Its element formats as `<edition>:<version>`.
     */
    ids?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getVersions.
 */
export interface GetVersionsResult {
    readonly edition?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly versions: outputs.servicemesh.GetVersionsVersion[];
}
/**
 * This data source provides ASM available versions in the specified region.
 *
 * > **NOTE:** Available in v1.161.0+.
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
 * const default = alicloud.servicemesh.getVersions({
 *     edition: "Default",
 * });
 * export const serviceMeshVersion = versions[0].version;
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getVersionsOutput(args?: GetVersionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVersionsResult> {
    return pulumi.output(args).apply((a: any) => getVersions(a, opts))
}

/**
 * A collection of arguments for invoking getVersions.
 */
export interface GetVersionsOutputArgs {
    /**
     * The edition of the ASM instance.
     */
    edition?: pulumi.Input<string>;
    /**
     * A list of ASM versions. Its element formats as `<edition>:<version>`.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
