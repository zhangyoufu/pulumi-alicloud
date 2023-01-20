// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides available image resources. It contains user's private images, system images provided by Alibaba Cloud,
 * other public images and the ones available on the image market.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const imagesDs = alicloud.ecs.getImages({
 *     nameRegex: "^centos_6",
 *     owners: "system",
 * });
 * export const firstImageId = imagesDs.then(imagesDs => imagesDs.images?.[0]?.id);
 * ```
 */
export function getImages(args?: GetImagesArgs, opts?: pulumi.InvokeOptions): Promise<GetImagesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ecs/getImages:getImages", {
        "actionType": args.actionType,
        "architecture": args.architecture,
        "dryRun": args.dryRun,
        "imageFamily": args.imageFamily,
        "imageId": args.imageId,
        "imageName": args.imageName,
        "imageOwnerId": args.imageOwnerId,
        "instanceType": args.instanceType,
        "isSupportCloudInit": args.isSupportCloudInit,
        "isSupportIoOptimized": args.isSupportIoOptimized,
        "mostRecent": args.mostRecent,
        "nameRegex": args.nameRegex,
        "osType": args.osType,
        "outputFile": args.outputFile,
        "owners": args.owners,
        "resourceGroupId": args.resourceGroupId,
        "snapshotId": args.snapshotId,
        "status": args.status,
        "tags": args.tags,
        "usage": args.usage,
    }, opts);
}

/**
 * A collection of arguments for invoking getImages.
 */
export interface GetImagesArgs {
    /**
     * The scenario in which the image will be used. Default value: `CreateEcs`. Valid values:
     */
    actionType?: string;
    /**
     * The image architecture. Valid values: `i386` and `x8664`.
     */
    architecture?: string;
    /**
     * Specifies whether the image is running on an ECS instance. Default value: `false`. Valid values:
     */
    dryRun?: boolean;
    /**
     * The name of the image family. You can set this parameter to query images of the specified image family. This parameter is empty by default.
     */
    imageFamily?: string;
    /**
     * The ID of the image.
     */
    imageId?: string;
    /**
     * The name of the image.
     */
    imageName?: string;
    /**
     * The ID of the Alibaba Cloud account to which the image belongs. This parameter takes effect only when you query shared images or community images.
     */
    imageOwnerId?: string;
    /**
     * The instance type for which the image can be used.
     */
    instanceType?: string;
    /**
     * Specifies whether the image supports cloud-init.
     */
    isSupportCloudInit?: boolean;
    /**
     * Specifies whether the image can be used on I/O optimized instances.
     */
    isSupportIoOptimized?: boolean;
    /**
     * If more than one result are returned, select the most recent one.
     */
    mostRecent?: boolean;
    /**
     * A regex string to filter resulting images by name.
     */
    nameRegex?: string;
    /**
     * The operating system type of the image. Valid values: `windows` and `linux`.
     */
    osType?: string;
    outputFile?: string;
    /**
     * Filter results by a specific image owner. Valid items are `system`, `self`, `others`, `marketplace`.
     */
    owners?: string;
    /**
     * The ID of the resource group to which the custom image belongs.
     */
    resourceGroupId?: string;
    /**
     * The ID of the snapshot used to create the custom image.
     */
    snapshotId?: string;
    /**
     * The status of the image. The following values are available, Separate multiple parameter values by using commas (,). Default value: `Available`. Valid values:
     */
    status?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: {[key: string]: any};
    /**
     * Specifies whether to check the validity of the request without actually making the request. Valid values:
     */
    usage?: string;
}

/**
 * A collection of values returned by getImages.
 */
export interface GetImagesResult {
    readonly actionType?: string;
    /**
     * Platform type of the image system: i386 or x86_64.
     */
    readonly architecture?: string;
    readonly dryRun?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of image IDs.
     */
    readonly ids: string[];
    readonly imageFamily?: string;
    readonly imageId?: string;
    readonly imageName?: string;
    readonly imageOwnerId?: string;
    /**
     * A list of images. Each element contains the following attributes:
     */
    readonly images: outputs.ecs.GetImagesImage[];
    readonly instanceType?: string;
    readonly isSupportCloudInit?: boolean;
    readonly isSupportIoOptimized?: boolean;
    readonly mostRecent?: boolean;
    readonly nameRegex?: string;
    readonly osType?: string;
    readonly outputFile?: string;
    readonly owners?: string;
    readonly resourceGroupId?: string;
    /**
     * Snapshot ID.
     */
    readonly snapshotId?: string;
    /**
     * Status of the image. Possible values: `UnAvailable`, `Available`, `Creating` and `CreateFailed`.
     */
    readonly status?: string;
    readonly tags?: {[key: string]: any};
    readonly usage?: string;
}
/**
 * This data source provides available image resources. It contains user's private images, system images provided by Alibaba Cloud,
 * other public images and the ones available on the image market.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const imagesDs = alicloud.ecs.getImages({
 *     nameRegex: "^centos_6",
 *     owners: "system",
 * });
 * export const firstImageId = imagesDs.then(imagesDs => imagesDs.images?.[0]?.id);
 * ```
 */
export function getImagesOutput(args?: GetImagesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImagesResult> {
    return pulumi.output(args).apply((a: any) => getImages(a, opts))
}

/**
 * A collection of arguments for invoking getImages.
 */
export interface GetImagesOutputArgs {
    /**
     * The scenario in which the image will be used. Default value: `CreateEcs`. Valid values:
     */
    actionType?: pulumi.Input<string>;
    /**
     * The image architecture. Valid values: `i386` and `x8664`.
     */
    architecture?: pulumi.Input<string>;
    /**
     * Specifies whether the image is running on an ECS instance. Default value: `false`. Valid values:
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The name of the image family. You can set this parameter to query images of the specified image family. This parameter is empty by default.
     */
    imageFamily?: pulumi.Input<string>;
    /**
     * The ID of the image.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The name of the image.
     */
    imageName?: pulumi.Input<string>;
    /**
     * The ID of the Alibaba Cloud account to which the image belongs. This parameter takes effect only when you query shared images or community images.
     */
    imageOwnerId?: pulumi.Input<string>;
    /**
     * The instance type for which the image can be used.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Specifies whether the image supports cloud-init.
     */
    isSupportCloudInit?: pulumi.Input<boolean>;
    /**
     * Specifies whether the image can be used on I/O optimized instances.
     */
    isSupportIoOptimized?: pulumi.Input<boolean>;
    /**
     * If more than one result are returned, select the most recent one.
     */
    mostRecent?: pulumi.Input<boolean>;
    /**
     * A regex string to filter resulting images by name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * The operating system type of the image. Valid values: `windows` and `linux`.
     */
    osType?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * Filter results by a specific image owner. Valid items are `system`, `self`, `others`, `marketplace`.
     */
    owners?: pulumi.Input<string>;
    /**
     * The ID of the resource group to which the custom image belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The ID of the snapshot used to create the custom image.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * The status of the image. The following values are available, Separate multiple parameter values by using commas (,). Default value: `Available`. Valid values:
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies whether to check the validity of the request without actually making the request. Valid values:
     */
    usage?: pulumi.Input<string>;
}
