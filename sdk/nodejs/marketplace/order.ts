// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/market_order.html.markdown.
 */
export class Order extends pulumi.CustomResource {
    /**
     * Get an existing Order resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrderState, opts?: pulumi.CustomResourceOptions): Order {
        return new Order(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:marketplace/order:Order';

    /**
     * Returns true if the given object is an instance of Order.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Order {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Order.__pulumiType;
    }

    /**
     * Service providers customize additional components.
     */
    public readonly components!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The coupon id of the market product.
     */
    public readonly couponId!: pulumi.Output<string | undefined>;
    /**
     * The number of purchase cycles.
     */
    public readonly duration!: pulumi.Output<number | undefined>;
    /**
     * The package version of the market product.
     */
    public readonly packageVersion!: pulumi.Output<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
     */
    public readonly payType!: pulumi.Output<string | undefined>;
    /**
     * The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
     */
    public readonly pricingCycle!: pulumi.Output<string>;
    /**
     * The productCode of market place product.
     */
    public readonly productCode!: pulumi.Output<string>;
    /**
     * The quantity of the market product will be purchased.
     */
    public readonly quantity!: pulumi.Output<number | undefined>;

    /**
     * Create a Order resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrderArgs | OrderState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as OrderState | undefined;
            inputs["components"] = state ? state.components : undefined;
            inputs["couponId"] = state ? state.couponId : undefined;
            inputs["duration"] = state ? state.duration : undefined;
            inputs["packageVersion"] = state ? state.packageVersion : undefined;
            inputs["payType"] = state ? state.payType : undefined;
            inputs["pricingCycle"] = state ? state.pricingCycle : undefined;
            inputs["productCode"] = state ? state.productCode : undefined;
            inputs["quantity"] = state ? state.quantity : undefined;
        } else {
            const args = argsOrState as OrderArgs | undefined;
            if (!args || args.packageVersion === undefined) {
                throw new Error("Missing required property 'packageVersion'");
            }
            if (!args || args.pricingCycle === undefined) {
                throw new Error("Missing required property 'pricingCycle'");
            }
            if (!args || args.productCode === undefined) {
                throw new Error("Missing required property 'productCode'");
            }
            inputs["components"] = args ? args.components : undefined;
            inputs["couponId"] = args ? args.couponId : undefined;
            inputs["duration"] = args ? args.duration : undefined;
            inputs["packageVersion"] = args ? args.packageVersion : undefined;
            inputs["payType"] = args ? args.payType : undefined;
            inputs["pricingCycle"] = args ? args.pricingCycle : undefined;
            inputs["productCode"] = args ? args.productCode : undefined;
            inputs["quantity"] = args ? args.quantity : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Order.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Order resources.
 */
export interface OrderState {
    /**
     * Service providers customize additional components.
     */
    readonly components?: pulumi.Input<{[key: string]: any}>;
    /**
     * The coupon id of the market product.
     */
    readonly couponId?: pulumi.Input<string>;
    /**
     * The number of purchase cycles.
     */
    readonly duration?: pulumi.Input<number>;
    /**
     * The package version of the market product.
     */
    readonly packageVersion?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
     */
    readonly payType?: pulumi.Input<string>;
    /**
     * The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
     */
    readonly pricingCycle?: pulumi.Input<string>;
    /**
     * The productCode of market place product.
     */
    readonly productCode?: pulumi.Input<string>;
    /**
     * The quantity of the market product will be purchased.
     */
    readonly quantity?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Order resource.
 */
export interface OrderArgs {
    /**
     * Service providers customize additional components.
     */
    readonly components?: pulumi.Input<{[key: string]: any}>;
    /**
     * The coupon id of the market product.
     */
    readonly couponId?: pulumi.Input<string>;
    /**
     * The number of purchase cycles.
     */
    readonly duration?: pulumi.Input<number>;
    /**
     * The package version of the market product.
     */
    readonly packageVersion: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
     */
    readonly payType?: pulumi.Input<string>;
    /**
     * The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
     */
    readonly pricingCycle: pulumi.Input<string>;
    /**
     * The productCode of market place product.
     */
    readonly productCode: pulumi.Input<string>;
    /**
     * The quantity of the market product will be purchased.
     */
    readonly quantity?: pulumi.Input<number>;
}
