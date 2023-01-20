// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetProductArgs, GetProductResult, GetProductOutputArgs } from "./getProduct";
export const getProduct: typeof import("./getProduct").getProduct = null as any;
export const getProductOutput: typeof import("./getProduct").getProductOutput = null as any;
utilities.lazyLoad(exports, ["getProduct","getProductOutput"], () => require("./getProduct"));

export { GetProductsArgs, GetProductsResult, GetProductsOutputArgs } from "./getProducts";
export const getProducts: typeof import("./getProducts").getProducts = null as any;
export const getProductsOutput: typeof import("./getProducts").getProductsOutput = null as any;
utilities.lazyLoad(exports, ["getProducts","getProductsOutput"], () => require("./getProducts"));

export { OrderArgs, OrderState } from "./order";
export type Order = import("./order").Order;
export const Order: typeof import("./order").Order = null as any;
utilities.lazyLoad(exports, ["Order"], () => require("./order"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:marketplace/order:Order":
                return new Order(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "marketplace/order", _module)
