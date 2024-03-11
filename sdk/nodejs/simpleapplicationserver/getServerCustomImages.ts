// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Simple Application Server Custom Images of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.143.0+.
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
 * const ids = alicloud.simpleapplicationserver.getServerCustomImages({
 *     ids: ["example_id"],
 * });
 * export const simpleApplicationServerCustomImageId1 = ids.then(ids => ids.images?.[0]?.id);
 * const nameRegex = alicloud.simpleapplicationserver.getServerCustomImages({
 *     nameRegex: "^my-CustomImage",
 * });
 * export const simpleApplicationServerCustomImageId2 = nameRegex.then(nameRegex => nameRegex.images?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getServerCustomImages(args?: GetServerCustomImagesArgs, opts?: pulumi.InvokeOptions): Promise<GetServerCustomImagesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:simpleapplicationserver/getServerCustomImages:getServerCustomImages", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getServerCustomImages.
 */
export interface GetServerCustomImagesArgs {
    /**
     * A list of Custom Image IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Custom Image name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getServerCustomImages.
 */
export interface GetServerCustomImagesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly images: outputs.simpleapplicationserver.GetServerCustomImagesImage[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
}
/**
 * This data source provides the Simple Application Server Custom Images of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.143.0+.
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
 * const ids = alicloud.simpleapplicationserver.getServerCustomImages({
 *     ids: ["example_id"],
 * });
 * export const simpleApplicationServerCustomImageId1 = ids.then(ids => ids.images?.[0]?.id);
 * const nameRegex = alicloud.simpleapplicationserver.getServerCustomImages({
 *     nameRegex: "^my-CustomImage",
 * });
 * export const simpleApplicationServerCustomImageId2 = nameRegex.then(nameRegex => nameRegex.images?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getServerCustomImagesOutput(args?: GetServerCustomImagesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServerCustomImagesResult> {
    return pulumi.output(args).apply((a: any) => getServerCustomImages(a, opts))
}

/**
 * A collection of arguments for invoking getServerCustomImages.
 */
export interface GetServerCustomImagesOutputArgs {
    /**
     * A list of Custom Image IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Custom Image name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
