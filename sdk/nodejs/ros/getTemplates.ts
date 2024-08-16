// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Ros Templates of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.108.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.ros.getTemplates({
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstRosTemplateId = example.then(example => example.templates?.[0]?.id);
 * ```
 */
export function getTemplates(args?: GetTemplatesArgs, opts?: pulumi.InvokeOptions): Promise<GetTemplatesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ros/getTemplates:getTemplates", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "shareType": args.shareType,
        "tags": args.tags,
        "templateName": args.templateName,
    }, opts);
}

/**
 * A collection of arguments for invoking getTemplates.
 */
export interface GetTemplatesArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: boolean;
    /**
     * A list of Template IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Template name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * Share Type. Valid Values: `Private`, `Shared`
     */
    shareType?: string;
    /**
     * Query the resource bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{"key1":"value1"}`.
     */
    tags?: {[key: string]: string};
    /**
     * The name of the template.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
     */
    templateName?: string;
}

/**
 * A collection of values returned by getTemplates.
 */
export interface GetTemplatesResult {
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly shareType?: string;
    readonly tags?: {[key: string]: string};
    readonly templateName?: string;
    readonly templates: outputs.ros.GetTemplatesTemplate[];
}
/**
 * This data source provides the Ros Templates of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.108.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.ros.getTemplates({
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstRosTemplateId = example.then(example => example.templates?.[0]?.id);
 * ```
 */
export function getTemplatesOutput(args?: GetTemplatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTemplatesResult> {
    return pulumi.output(args).apply((a: any) => getTemplates(a, opts))
}

/**
 * A collection of arguments for invoking getTemplates.
 */
export interface GetTemplatesOutputArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of Template IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Template name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Share Type. Valid Values: `Private`, `Shared`
     */
    shareType?: pulumi.Input<string>;
    /**
     * Query the resource bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{"key1":"value1"}`.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the template.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
     */
    templateName?: pulumi.Input<string>;
}
