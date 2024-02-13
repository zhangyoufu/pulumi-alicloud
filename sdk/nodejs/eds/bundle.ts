// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ECD Bundle resource.
 *
 * For information about ECD Bundle and how to use it, see [What is Bundle](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createbundle).
 *
 * > **NOTE:** Available since v1.170.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const defaultImages = alicloud.eds.getImages({
 *     imageType: "SYSTEM",
 *     osType: "Windows",
 *     desktopInstanceType: "eds.hf.4c8g",
 * });
 * const defaultDesktopTypes = alicloud.eds.getDesktopTypes({
 *     instanceTypeFamily: "eds.hf",
 *     cpuCount: 4,
 *     memorySize: 8192,
 * });
 * const defaultBundle = new alicloud.eds.Bundle("defaultBundle", {
 *     description: name,
 *     desktopType: defaultDesktopTypes.then(defaultDesktopTypes => defaultDesktopTypes.ids?.[0]),
 *     bundleName: name,
 *     imageId: defaultImages.then(defaultImages => defaultImages.ids?.[0]),
 *     userDiskSizeGibs: [70],
 *     rootDiskSizeGib: 80,
 *     rootDiskPerformanceLevel: "PL1",
 *     userDiskPerformanceLevel: "PL1",
 * });
 * ```
 *
 * ## Import
 *
 * ECD Bundle can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:eds/bundle:Bundle example <id>
 * ```
 */
export class Bundle extends pulumi.CustomResource {
    /**
     * Get an existing Bundle resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BundleState, opts?: pulumi.CustomResourceOptions): Bundle {
        return new Bundle(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eds/bundle:Bundle';

    /**
     * Returns true if the given object is an instance of Bundle.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Bundle {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Bundle.__pulumiType;
    }

    /**
     * The name of the bundle.
     */
    public readonly bundleName!: pulumi.Output<string | undefined>;
    /**
     * The description of the bundle.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The desktop type. You can call `alicloud.eds.getDesktopTypes` to query desktop type.
     */
    public readonly desktopType!: pulumi.Output<string>;
    /**
     * The ID of the image.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * The language. Valid values: `zh-CN`, `zh-HK`, `en-US`, `ja-JP`.
     */
    public readonly language!: pulumi.Output<string | undefined>;
    /**
     * The root disk performance level. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
     */
    public readonly rootDiskPerformanceLevel!: pulumi.Output<string>;
    /**
     * The root disk size gib.
     */
    public readonly rootDiskSizeGib!: pulumi.Output<number>;
    /**
     * The user disk performance level. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
     */
    public readonly userDiskPerformanceLevel!: pulumi.Output<string>;
    /**
     * The size of the data disk. Currently, only one data disk can be set. Unit: GiB.
     * - The size of the data disk that supports the setting corresponds to the specification. For more information, see [Overview of Desktop Specifications](https://help.aliyun.com/document_detail/188609.htm?spm=a2c4g.11186623.0.0.6406297bE0U5DG).
     * - The data disk size (user_disk_size_gib) set in the template must be greater than the data disk size (data_disk_size) in the mirror.
     */
    public readonly userDiskSizeGibs!: pulumi.Output<number[]>;

    /**
     * Create a Bundle resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BundleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BundleArgs | BundleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BundleState | undefined;
            resourceInputs["bundleName"] = state ? state.bundleName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["desktopType"] = state ? state.desktopType : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["language"] = state ? state.language : undefined;
            resourceInputs["rootDiskPerformanceLevel"] = state ? state.rootDiskPerformanceLevel : undefined;
            resourceInputs["rootDiskSizeGib"] = state ? state.rootDiskSizeGib : undefined;
            resourceInputs["userDiskPerformanceLevel"] = state ? state.userDiskPerformanceLevel : undefined;
            resourceInputs["userDiskSizeGibs"] = state ? state.userDiskSizeGibs : undefined;
        } else {
            const args = argsOrState as BundleArgs | undefined;
            if ((!args || args.desktopType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'desktopType'");
            }
            if ((!args || args.imageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageId'");
            }
            if ((!args || args.rootDiskSizeGib === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rootDiskSizeGib'");
            }
            if ((!args || args.userDiskSizeGibs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userDiskSizeGibs'");
            }
            resourceInputs["bundleName"] = args ? args.bundleName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["desktopType"] = args ? args.desktopType : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["language"] = args ? args.language : undefined;
            resourceInputs["rootDiskPerformanceLevel"] = args ? args.rootDiskPerformanceLevel : undefined;
            resourceInputs["rootDiskSizeGib"] = args ? args.rootDiskSizeGib : undefined;
            resourceInputs["userDiskPerformanceLevel"] = args ? args.userDiskPerformanceLevel : undefined;
            resourceInputs["userDiskSizeGibs"] = args ? args.userDiskSizeGibs : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Bundle.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Bundle resources.
 */
export interface BundleState {
    /**
     * The name of the bundle.
     */
    bundleName?: pulumi.Input<string>;
    /**
     * The description of the bundle.
     */
    description?: pulumi.Input<string>;
    /**
     * The desktop type. You can call `alicloud.eds.getDesktopTypes` to query desktop type.
     */
    desktopType?: pulumi.Input<string>;
    /**
     * The ID of the image.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The language. Valid values: `zh-CN`, `zh-HK`, `en-US`, `ja-JP`.
     */
    language?: pulumi.Input<string>;
    /**
     * The root disk performance level. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
     */
    rootDiskPerformanceLevel?: pulumi.Input<string>;
    /**
     * The root disk size gib.
     */
    rootDiskSizeGib?: pulumi.Input<number>;
    /**
     * The user disk performance level. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
     */
    userDiskPerformanceLevel?: pulumi.Input<string>;
    /**
     * The size of the data disk. Currently, only one data disk can be set. Unit: GiB.
     * - The size of the data disk that supports the setting corresponds to the specification. For more information, see [Overview of Desktop Specifications](https://help.aliyun.com/document_detail/188609.htm?spm=a2c4g.11186623.0.0.6406297bE0U5DG).
     * - The data disk size (user_disk_size_gib) set in the template must be greater than the data disk size (data_disk_size) in the mirror.
     */
    userDiskSizeGibs?: pulumi.Input<pulumi.Input<number>[]>;
}

/**
 * The set of arguments for constructing a Bundle resource.
 */
export interface BundleArgs {
    /**
     * The name of the bundle.
     */
    bundleName?: pulumi.Input<string>;
    /**
     * The description of the bundle.
     */
    description?: pulumi.Input<string>;
    /**
     * The desktop type. You can call `alicloud.eds.getDesktopTypes` to query desktop type.
     */
    desktopType: pulumi.Input<string>;
    /**
     * The ID of the image.
     */
    imageId: pulumi.Input<string>;
    /**
     * The language. Valid values: `zh-CN`, `zh-HK`, `en-US`, `ja-JP`.
     */
    language?: pulumi.Input<string>;
    /**
     * The root disk performance level. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
     */
    rootDiskPerformanceLevel?: pulumi.Input<string>;
    /**
     * The root disk size gib.
     */
    rootDiskSizeGib: pulumi.Input<number>;
    /**
     * The user disk performance level. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
     */
    userDiskPerformanceLevel?: pulumi.Input<string>;
    /**
     * The size of the data disk. Currently, only one data disk can be set. Unit: GiB.
     * - The size of the data disk that supports the setting corresponds to the specification. For more information, see [Overview of Desktop Specifications](https://help.aliyun.com/document_detail/188609.htm?spm=a2c4g.11186623.0.0.6406297bE0U5DG).
     * - The data disk size (user_disk_size_gib) set in the template must be greater than the data disk size (data_disk_size) in the mirror.
     */
    userDiskSizeGibs: pulumi.Input<pulumi.Input<number>[]>;
}
