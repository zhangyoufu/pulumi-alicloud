// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Market product item details of Alibaba Cloud.
 * 
 * > **NOTE:** Available in 1.69.0+
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const defaultProduct = alicloud.marketplace.getProduct({
 *     productCode: "cmapi022206",
 * });
 * 
 * export const productName = defaultProduct.products[0].name;
 * export const firstProductSkuCode = defaultProduct.products[0].skuses[0].skuCode;
 * export const firstProductPackageVersion = defaultProduct.products[0].skuses[0].packageVersions[0].packageVersion;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/market_product.html.markdown.
 */
export function getProduct(args: GetProductArgs, opts?: pulumi.InvokeOptions): Promise<GetProductResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:marketplace/getProduct:getProduct", {
        "availableRegion": args.availableRegion,
        "productCode": args.productCode,
    }, opts);
}

/**
 * A collection of arguments for invoking getProduct.
 */
export interface GetProductArgs {
    /**
     * A available region id used to filter market place Ecs images.
     */
    readonly availableRegion?: string;
    /**
     * The product code of the market product.
     */
    readonly productCode: string;
}

/**
 * A collection of values returned by getProduct.
 */
export interface GetProductResult {
    readonly availableRegion?: string;
    /**
     * A product. It contains the following attributes:
     */
    readonly products: outputs.marketplace.GetProductProduct[];
    readonly productCode: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
