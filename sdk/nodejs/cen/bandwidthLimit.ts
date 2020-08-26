// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CEN cross-regional interconnection bandwidth resource. To connect networks in different regions, you must set cross-region interconnection bandwidth after buying a bandwidth package. The total bandwidth set for all the interconnected regions of a bandwidth package cannot exceed the bandwidth of the bandwidth package. By default, 1 Kbps bandwidth is provided for connectivity test. To run normal business, you must buy a bandwidth package and set a proper interconnection bandwidth.
 *
 * For example, a CEN instance is bound to a bandwidth package of 20 Mbps and  the interconnection areas are Mainland China and North America. You can set the cross-region interconnection bandwidth between US West 1 and China East 1, China East 2, China South 1, and so on. However, the total bandwidth set for all the interconnected regions cannot exceed 20  Mbps.
 *
 * For information about CEN and how to use it, see [Cross-region interconnection bandwidth](https://www.alibabacloud.com/help/doc-detail/65983.htm)
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
 * const name = config.get("name") || "tf-testAccCenBandwidthLimitConfig";
 *
 * const fra = new alicloud.Provider("fra", {
 *     region: "eu-central-1",
 * });
 * const sh = new alicloud.Provider("sh", {
 *     region: "cn-shanghai",
 * });
 * const vpc1 = new alicloud.vpc.Network("vpc1", {
 *     cidrBlock: "192.168.0.0/16",
 * }, { provider: fra });
 * const vpc2 = new alicloud.vpc.Network("vpc2", {
 *     cidrBlock: "172.16.0.0/12",
 * }, { provider: sh });
 * const cen = new alicloud.cen.Instance("cen", {
 *     description: "tf-testAccCenBandwidthLimitConfigDescription",
 * });
 * const bwp = new alicloud.cen.BandwidthPackage("bwp", {
 *     bandwidth: 5,
 *     geographicRegionIds: [
 *         "Europe",
 *         "China",
 *     ],
 * });
 * const bwpAttach = new alicloud.cen.BandwidthPackageAttachment("bwp_attach", {
 *     bandwidthPackageId: bwp.id,
 *     instanceId: cen.id,
 * });
 * const vpcAttach1 = new alicloud.cen.InstanceAttachment("vpc_attach_1", {
 *     childInstanceId: vpc1.id,
 *     childInstanceRegionId: "eu-central-1",
 *     instanceId: cen.id,
 * });
 * const vpcAttach2 = new alicloud.cen.InstanceAttachment("vpc_attach_2", {
 *     childInstanceId: vpc2.id,
 *     childInstanceRegionId: "cn-shanghai",
 *     instanceId: cen.id,
 * });
 * const foo = new alicloud.cen.BandwidthLimit("foo", {
 *     bandwidthLimit: 4,
 *     instanceId: cen.id,
 *     regionIds: [
 *         "eu-central-1",
 *         "cn-shanghai",
 *     ],
 * }, { dependsOn: [bwpAttach, vpcAttach1, vpcAttach2] });
 * ```
 */
export class BandwidthLimit extends pulumi.CustomResource {
    /**
     * Get an existing BandwidthLimit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BandwidthLimitState, opts?: pulumi.CustomResourceOptions): BandwidthLimit {
        return new BandwidthLimit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cen/bandwidthLimit:BandwidthLimit';

    /**
     * Returns true if the given object is an instance of BandwidthLimit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BandwidthLimit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BandwidthLimit.__pulumiType;
    }

    /**
     * The bandwidth configured for the interconnected regions communication.
     */
    public readonly bandwidthLimit!: pulumi.Output<number>;
    /**
     * The ID of the CEN.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * List of the two regions to interconnect. Must be two different regions.
     */
    public readonly regionIds!: pulumi.Output<string[]>;

    /**
     * Create a BandwidthLimit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BandwidthLimitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BandwidthLimitArgs | BandwidthLimitState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BandwidthLimitState | undefined;
            inputs["bandwidthLimit"] = state ? state.bandwidthLimit : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["regionIds"] = state ? state.regionIds : undefined;
        } else {
            const args = argsOrState as BandwidthLimitArgs | undefined;
            if (!args || args.bandwidthLimit === undefined) {
                throw new Error("Missing required property 'bandwidthLimit'");
            }
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            if (!args || args.regionIds === undefined) {
                throw new Error("Missing required property 'regionIds'");
            }
            inputs["bandwidthLimit"] = args ? args.bandwidthLimit : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["regionIds"] = args ? args.regionIds : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BandwidthLimit.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BandwidthLimit resources.
 */
export interface BandwidthLimitState {
    /**
     * The bandwidth configured for the interconnected regions communication.
     */
    readonly bandwidthLimit?: pulumi.Input<number>;
    /**
     * The ID of the CEN.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * List of the two regions to interconnect. Must be two different regions.
     */
    readonly regionIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a BandwidthLimit resource.
 */
export interface BandwidthLimitArgs {
    /**
     * The bandwidth configured for the interconnected regions communication.
     */
    readonly bandwidthLimit: pulumi.Input<number>;
    /**
     * The ID of the CEN.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * List of the two regions to interconnect. Must be two different regions.
     */
    readonly regionIds: pulumi.Input<pulumi.Input<string>[]>;
}
