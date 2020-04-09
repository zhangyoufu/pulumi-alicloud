// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CEN bandwidth package resource. The CEN bandwidth package is an abstracted object that includes an interconnection bandwidth and interconnection areas. To buy a bandwidth package, you must specify the areas to connect. An area consists of one or more Alibaba Cloud regions. The areas in CEN include Mainland China, Asia Pacific, North America, and Europe.
 * 
 * For information about CEN and how to use it, see [Manage bandwidth packages](https://www.alibabacloud.com/help/doc-detail/65982.htm).
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const foo = new alicloud.cen.BandwidthPackage("foo", {
 *     bandwidth: 5,
 *     geographicRegionIds: [
 *         "China",
 *         "Asia-Pacific",
 *     ],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cen_bandwidth_package.html.markdown.
 */
export class BandwidthPackage extends pulumi.CustomResource {
    /**
     * Get an existing BandwidthPackage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BandwidthPackageState, opts?: pulumi.CustomResourceOptions): BandwidthPackage {
        return new BandwidthPackage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cen/bandwidthPackage:BandwidthPackage';

    /**
     * Returns true if the given object is an instance of BandwidthPackage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BandwidthPackage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BandwidthPackage.__pulumiType;
    }

    /**
     * The bandwidth in Mbps of the bandwidth package. Cannot be less than 2Mbps.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * The billing method. Valid value: PostPaid | PrePaid. Default to PostPaid. If set to PrePaid, the bandwidth package can't be deleted before expired time.
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * The description of the bandwidth package. Default to null.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The time of the bandwidth package to expire.
     */
    public /*out*/ readonly expiredTime!: pulumi.Output<string>;
    /**
     * List of the two areas to connect. Valid value: China | North-America | Asia-Pacific | Europe | Middle-East.
     */
    public readonly geographicRegionIds!: pulumi.Output<string[]>;
    /**
     * The name of the bandwidth package. Defaults to null.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The purchase period in month. Valid value: 1, 2, 3, 6, 12. Default to 1.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The status of the bandwidth, including "InUse" and "Idle".
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a BandwidthPackage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BandwidthPackageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BandwidthPackageArgs | BandwidthPackageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BandwidthPackageState | undefined;
            inputs["bandwidth"] = state ? state.bandwidth : undefined;
            inputs["chargeType"] = state ? state.chargeType : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["expiredTime"] = state ? state.expiredTime : undefined;
            inputs["geographicRegionIds"] = state ? state.geographicRegionIds : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as BandwidthPackageArgs | undefined;
            if (!args || args.bandwidth === undefined) {
                throw new Error("Missing required property 'bandwidth'");
            }
            if (!args || args.geographicRegionIds === undefined) {
                throw new Error("Missing required property 'geographicRegionIds'");
            }
            inputs["bandwidth"] = args ? args.bandwidth : undefined;
            inputs["chargeType"] = args ? args.chargeType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["geographicRegionIds"] = args ? args.geographicRegionIds : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["expiredTime"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BandwidthPackage.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BandwidthPackage resources.
 */
export interface BandwidthPackageState {
    /**
     * The bandwidth in Mbps of the bandwidth package. Cannot be less than 2Mbps.
     */
    readonly bandwidth?: pulumi.Input<number>;
    /**
     * The billing method. Valid value: PostPaid | PrePaid. Default to PostPaid. If set to PrePaid, the bandwidth package can't be deleted before expired time.
     */
    readonly chargeType?: pulumi.Input<string>;
    /**
     * The description of the bandwidth package. Default to null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The time of the bandwidth package to expire.
     */
    readonly expiredTime?: pulumi.Input<string>;
    /**
     * List of the two areas to connect. Valid value: China | North-America | Asia-Pacific | Europe | Middle-East.
     */
    readonly geographicRegionIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the bandwidth package. Defaults to null.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The purchase period in month. Valid value: 1, 2, 3, 6, 12. Default to 1.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * The status of the bandwidth, including "InUse" and "Idle".
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BandwidthPackage resource.
 */
export interface BandwidthPackageArgs {
    /**
     * The bandwidth in Mbps of the bandwidth package. Cannot be less than 2Mbps.
     */
    readonly bandwidth: pulumi.Input<number>;
    /**
     * The billing method. Valid value: PostPaid | PrePaid. Default to PostPaid. If set to PrePaid, the bandwidth package can't be deleted before expired time.
     */
    readonly chargeType?: pulumi.Input<string>;
    /**
     * The description of the bandwidth package. Default to null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * List of the two areas to connect. Valid value: China | North-America | Asia-Pacific | Europe | Middle-East.
     */
    readonly geographicRegionIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the bandwidth package. Defaults to null.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The purchase period in month. Valid value: 1, 2, 3, 6, 12. Default to 1.
     */
    readonly period?: pulumi.Input<number>;
}
