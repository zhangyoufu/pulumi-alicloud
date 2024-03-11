// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides Service Catalog Launch Option available to the user. [What is Launch Option](https://www.alibabacloud.com/help/en/service-catalog/developer-reference/api-servicecatalog-2021-09-01-listlaunchoptions).
 *
 * > **NOTE:** Available since v1.196.0.
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
 * const defaultEndUserProducts = alicloud.servicecatalog.getEndUserProducts({
 *     nameRegex: "ram模板创建",
 * });
 * const defaultLaunchOptions = alicloud.servicecatalog.getLaunchOptions({
 *     productId: "data.alicloud_service_catalog_end_user_products.default.end_user_products.0.id",
 * });
 * export const alicloudServiceCatalogLaunchOptionExampleId = defaultLaunchOptions.then(defaultLaunchOptions => defaultLaunchOptions.launchOptions?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getLaunchOptions(args: GetLaunchOptionsArgs, opts?: pulumi.InvokeOptions): Promise<GetLaunchOptionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:servicecatalog/getLaunchOptions:getLaunchOptions", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "productId": args.productId,
    }, opts);
}

/**
 * A collection of arguments for invoking getLaunchOptions.
 */
export interface GetLaunchOptionsArgs {
    /**
     * A list of Launch Option IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by portfolio name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * Product ID.
     */
    productId: string;
}

/**
 * A collection of values returned by getLaunchOptions.
 */
export interface GetLaunchOptionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    /**
     * (Available since v1.197.0) A list of Launch Option Entries. Each element contains the following attributes:
     */
    readonly launchOptions: outputs.servicecatalog.GetLaunchOptionsLaunchOption[];
    readonly nameRegex?: string;
    /**
     * (Deprecated since v1.197.0) A list of Launch Option Entries. Each element contains the following attributes:
     *
     * @deprecated Field 'options' has been deprecated from provider version 1.197.0.
     */
    readonly options: outputs.servicecatalog.GetLaunchOptionsOption[];
    readonly outputFile?: string;
    readonly productId: string;
}
/**
 * This data source provides Service Catalog Launch Option available to the user. [What is Launch Option](https://www.alibabacloud.com/help/en/service-catalog/developer-reference/api-servicecatalog-2021-09-01-listlaunchoptions).
 *
 * > **NOTE:** Available since v1.196.0.
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
 * const defaultEndUserProducts = alicloud.servicecatalog.getEndUserProducts({
 *     nameRegex: "ram模板创建",
 * });
 * const defaultLaunchOptions = alicloud.servicecatalog.getLaunchOptions({
 *     productId: "data.alicloud_service_catalog_end_user_products.default.end_user_products.0.id",
 * });
 * export const alicloudServiceCatalogLaunchOptionExampleId = defaultLaunchOptions.then(defaultLaunchOptions => defaultLaunchOptions.launchOptions?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getLaunchOptionsOutput(args: GetLaunchOptionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLaunchOptionsResult> {
    return pulumi.output(args).apply((a: any) => getLaunchOptions(a, opts))
}

/**
 * A collection of arguments for invoking getLaunchOptions.
 */
export interface GetLaunchOptionsOutputArgs {
    /**
     * A list of Launch Option IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by portfolio name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Product ID.
     */
    productId: pulumi.Input<string>;
}
