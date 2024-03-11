// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of OOS Templates in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in v1.92.0+.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.oos.getTemplates({
 *     hasTrigger: false,
 *     nameRegex: "test",
 *     shareType: "Private",
 *     tags: {
 *         Created: "TF",
 *         For: "template Test",
 *     },
 * });
 * export const firstTemplateName = example.then(example => example.templates?.[0]?.templateName);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getTemplates(args?: GetTemplatesArgs, opts?: pulumi.InvokeOptions): Promise<GetTemplatesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:oos/getTemplates:getTemplates", {
        "category": args.category,
        "createdBy": args.createdBy,
        "createdDate": args.createdDate,
        "createdDateAfter": args.createdDateAfter,
        "hasTrigger": args.hasTrigger,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "shareType": args.shareType,
        "sortField": args.sortField,
        "sortOrder": args.sortOrder,
        "tags": args.tags,
        "templateFormat": args.templateFormat,
        "templateType": args.templateType,
    }, opts);
}

/**
 * A collection of arguments for invoking getTemplates.
 */
export interface GetTemplatesArgs {
    /**
     * The category of template.
     */
    category?: string;
    /**
     * The creator of the template.
     */
    createdBy?: string;
    /**
     * The template whose creation time is less than or equal to the specified time. The format is: YYYY-MM-DDThh:mm::ssZ.
     */
    createdDate?: string;
    /**
     * Create a template whose time is greater than or equal to the specified time. The format is: YYYY-MM-DDThh:mm:ssZ.
     */
    createdDateAfter?: string;
    /**
     * Is it triggered successfully.
     */
    hasTrigger?: boolean;
    /**
     * A list of OOS Template ids. Each element in the list is same as template_name.
     */
    ids?: string[];
    /**
     * A regex string to filter the results by the template_name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The sharing type of the template. Valid values: `Private`, `Public`.
     */
    shareType?: string;
    /**
     * Sort field. Valid values: `TotalExecutionCount`, `Popularity`, `TemplateName` and `CreatedDate`. Default to `TotalExecutionCount`.
     */
    sortField?: string;
    /**
     * Sort order. Valid values: `Ascending`, `Descending`. Default to `Descending`
     */
    sortOrder?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: {[key: string]: any};
    /**
     * The format of the template. Valid values: `JSON`, `YAML`.
     */
    templateFormat?: string;
    /**
     * The type of OOS Template.
     */
    templateType?: string;
}

/**
 * A collection of values returned by getTemplates.
 */
export interface GetTemplatesResult {
    readonly category?: string;
    readonly createdBy?: string;
    readonly createdDate?: string;
    readonly createdDateAfter?: string;
    readonly hasTrigger?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of OOS Template ids. Each element in the list is same as template_name.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * (Available in v1.114.0+) A list of OOS Template names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly shareType?: string;
    readonly sortField?: string;
    readonly sortOrder?: string;
    readonly tags?: {[key: string]: any};
    readonly templateFormat?: string;
    readonly templateType?: string;
    /**
     * A list of OOS Templates. Each element contains the following attributes:
     */
    readonly templates: outputs.oos.GetTemplatesTemplate[];
}
/**
 * This data source provides a list of OOS Templates in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in v1.92.0+.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.oos.getTemplates({
 *     hasTrigger: false,
 *     nameRegex: "test",
 *     shareType: "Private",
 *     tags: {
 *         Created: "TF",
 *         For: "template Test",
 *     },
 * });
 * export const firstTemplateName = example.then(example => example.templates?.[0]?.templateName);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getTemplatesOutput(args?: GetTemplatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTemplatesResult> {
    return pulumi.output(args).apply((a: any) => getTemplates(a, opts))
}

/**
 * A collection of arguments for invoking getTemplates.
 */
export interface GetTemplatesOutputArgs {
    /**
     * The category of template.
     */
    category?: pulumi.Input<string>;
    /**
     * The creator of the template.
     */
    createdBy?: pulumi.Input<string>;
    /**
     * The template whose creation time is less than or equal to the specified time. The format is: YYYY-MM-DDThh:mm::ssZ.
     */
    createdDate?: pulumi.Input<string>;
    /**
     * Create a template whose time is greater than or equal to the specified time. The format is: YYYY-MM-DDThh:mm:ssZ.
     */
    createdDateAfter?: pulumi.Input<string>;
    /**
     * Is it triggered successfully.
     */
    hasTrigger?: pulumi.Input<boolean>;
    /**
     * A list of OOS Template ids. Each element in the list is same as template_name.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter the results by the template_name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The sharing type of the template. Valid values: `Private`, `Public`.
     */
    shareType?: pulumi.Input<string>;
    /**
     * Sort field. Valid values: `TotalExecutionCount`, `Popularity`, `TemplateName` and `CreatedDate`. Default to `TotalExecutionCount`.
     */
    sortField?: pulumi.Input<string>;
    /**
     * Sort order. Valid values: `Ascending`, `Descending`. Default to `Descending`
     */
    sortOrder?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The format of the template. Valid values: `JSON`, `YAML`.
     */
    templateFormat?: pulumi.Input<string>;
    /**
     * The type of OOS Template.
     */
    templateType?: pulumi.Input<string>;
}
