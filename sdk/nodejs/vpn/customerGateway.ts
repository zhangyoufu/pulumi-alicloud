// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
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
 * const _default = new alicloud.vpn.CustomerGateway("default", {
 *     description: name,
 *     ipAddress: "4.3.2.10",
 *     asn: "1219002",
 *     customerGatewayName: name,
 * });
 * ```
 *
 * ## Import
 *
 * VPN customer gateway can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:vpn/customerGateway:CustomerGateway example <id>
 * ```
 */
export class CustomerGateway extends pulumi.CustomResource {
    /**
     * Get an existing CustomerGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomerGatewayState, opts?: pulumi.CustomResourceOptions): CustomerGateway {
        return new CustomerGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpn/customerGateway:CustomerGateway';

    /**
     * Returns true if the given object is an instance of CustomerGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomerGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomerGateway.__pulumiType;
    }

    /**
     * Asn.
     */
    public readonly asn!: pulumi.Output<string | undefined>;
    /**
     * The time when the customer gateway was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<number>;
    /**
     * The name of the customer gateway.
     */
    public readonly customerGatewayName!: pulumi.Output<string>;
    /**
     * The description of the customer gateway.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The IP address of the customer gateway.
     */
    public readonly ipAddress!: pulumi.Output<string>;
    /**
     * . Field 'name' has been deprecated from provider version 1.216.0. New field 'customer_gateway_name' instead.
     *
     * @deprecated Field 'name' has been deprecated since provider version 1.210.0. New field 'customer_gateway_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * tag.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a CustomerGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomerGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomerGatewayArgs | CustomerGatewayState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomerGatewayState | undefined;
            resourceInputs["asn"] = state ? state.asn : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["customerGatewayName"] = state ? state.customerGatewayName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as CustomerGatewayArgs | undefined;
            if ((!args || args.ipAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipAddress'");
            }
            resourceInputs["asn"] = args ? args.asn : undefined;
            resourceInputs["customerGatewayName"] = args ? args.customerGatewayName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ipAddress"] = args ? args.ipAddress : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomerGateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomerGateway resources.
 */
export interface CustomerGatewayState {
    /**
     * Asn.
     */
    asn?: pulumi.Input<string>;
    /**
     * The time when the customer gateway was created.
     */
    createTime?: pulumi.Input<number>;
    /**
     * The name of the customer gateway.
     */
    customerGatewayName?: pulumi.Input<string>;
    /**
     * The description of the customer gateway.
     */
    description?: pulumi.Input<string>;
    /**
     * The IP address of the customer gateway.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * . Field 'name' has been deprecated from provider version 1.216.0. New field 'customer_gateway_name' instead.
     *
     * @deprecated Field 'name' has been deprecated since provider version 1.210.0. New field 'customer_gateway_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * tag.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a CustomerGateway resource.
 */
export interface CustomerGatewayArgs {
    /**
     * Asn.
     */
    asn?: pulumi.Input<string>;
    /**
     * The name of the customer gateway.
     */
    customerGatewayName?: pulumi.Input<string>;
    /**
     * The description of the customer gateway.
     */
    description?: pulumi.Input<string>;
    /**
     * The IP address of the customer gateway.
     */
    ipAddress: pulumi.Input<string>;
    /**
     * . Field 'name' has been deprecated from provider version 1.216.0. New field 'customer_gateway_name' instead.
     *
     * @deprecated Field 'name' has been deprecated since provider version 1.210.0. New field 'customer_gateway_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * tag.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
