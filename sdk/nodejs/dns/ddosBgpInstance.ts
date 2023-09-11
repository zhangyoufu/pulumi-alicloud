// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Anti-DDoS Advanced instance resource. "Ddosbgp" is the short term of this product.
 *
 * > **NOTE:** The endpoint of bssopenapi used only support "business.aliyuncs.com" at present.
 *
 * > **NOTE:** Available since v1.183.0.
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
 * const name = config.get("name") || "tf-example";
 * const instance = new alicloud.ddos.DdosBgpInstance("instance", {
 *     baseBandwidth: 20,
 *     bandwidth: -1,
 *     ipCount: 100,
 *     ipType: "IPv4",
 *     normalBandwidth: 100,
 *     type: "Enterprise",
 * });
 * ```
 *
 * ## Import
 *
 * Ddosbgp instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:dns/ddosBgpInstance:DdosBgpInstance example ddosbgp-abc123456
 * ```
 *
 * @deprecated alicloud.dns.DdosBgpInstance has been deprecated in favor of alicloud.ddos.DdosBgpInstance
 */
export class DdosBgpInstance extends pulumi.CustomResource {
    /**
     * Get an existing DdosBgpInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DdosBgpInstanceState, opts?: pulumi.CustomResourceOptions): DdosBgpInstance {
        pulumi.log.warn("DdosBgpInstance is deprecated: alicloud.dns.DdosBgpInstance has been deprecated in favor of alicloud.ddos.DdosBgpInstance")
        return new DdosBgpInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dns/ddosBgpInstance:DdosBgpInstance';

    /**
     * Returns true if the given object is an instance of DdosBgpInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DdosBgpInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DdosBgpInstance.__pulumiType;
    }

    /**
     * Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
     */
    public readonly baseBandwidth!: pulumi.Output<number | undefined>;
    /**
     * IP count of the instance. Valid values: 100.
     */
    public readonly ipCount!: pulumi.Output<number>;
    /**
     * IP version of the instance. Valid values: IPv4,IPv6.
     */
    public readonly ipType!: pulumi.Output<string>;
    /**
     * Name of the instance. This name can have a string of 1 to 63 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Normal defend bandwidth of the instance. The unit is Gbps.
     */
    public readonly normalBandwidth!: pulumi.Output<number>;
    /**
     * The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Type of the instance. Valid values: `Enterprise`, `Professional`. Default to `Enterprise`
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a DdosBgpInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated alicloud.dns.DdosBgpInstance has been deprecated in favor of alicloud.ddos.DdosBgpInstance */
    constructor(name: string, args: DdosBgpInstanceArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated alicloud.dns.DdosBgpInstance has been deprecated in favor of alicloud.ddos.DdosBgpInstance */
    constructor(name: string, argsOrState?: DdosBgpInstanceArgs | DdosBgpInstanceState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DdosBgpInstance is deprecated: alicloud.dns.DdosBgpInstance has been deprecated in favor of alicloud.ddos.DdosBgpInstance")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DdosBgpInstanceState | undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["baseBandwidth"] = state ? state.baseBandwidth : undefined;
            resourceInputs["ipCount"] = state ? state.ipCount : undefined;
            resourceInputs["ipType"] = state ? state.ipType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["normalBandwidth"] = state ? state.normalBandwidth : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as DdosBgpInstanceArgs | undefined;
            if ((!args || args.bandwidth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidth'");
            }
            if ((!args || args.ipCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipCount'");
            }
            if ((!args || args.ipType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipType'");
            }
            if ((!args || args.normalBandwidth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'normalBandwidth'");
            }
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["baseBandwidth"] = args ? args.baseBandwidth : undefined;
            resourceInputs["ipCount"] = args ? args.ipCount : undefined;
            resourceInputs["ipType"] = args ? args.ipType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["normalBandwidth"] = args ? args.normalBandwidth : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DdosBgpInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DdosBgpInstance resources.
 */
export interface DdosBgpInstanceState {
    /**
     * Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
     */
    baseBandwidth?: pulumi.Input<number>;
    /**
     * IP count of the instance. Valid values: 100.
     */
    ipCount?: pulumi.Input<number>;
    /**
     * IP version of the instance. Valid values: IPv4,IPv6.
     */
    ipType?: pulumi.Input<string>;
    /**
     * Name of the instance. This name can have a string of 1 to 63 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Normal defend bandwidth of the instance. The unit is Gbps.
     */
    normalBandwidth?: pulumi.Input<number>;
    /**
     * The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
     */
    period?: pulumi.Input<number>;
    /**
     * Type of the instance. Valid values: `Enterprise`, `Professional`. Default to `Enterprise`
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DdosBgpInstance resource.
 */
export interface DdosBgpInstanceArgs {
    /**
     * Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
     */
    bandwidth: pulumi.Input<number>;
    /**
     * Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
     */
    baseBandwidth?: pulumi.Input<number>;
    /**
     * IP count of the instance. Valid values: 100.
     */
    ipCount: pulumi.Input<number>;
    /**
     * IP version of the instance. Valid values: IPv4,IPv6.
     */
    ipType: pulumi.Input<string>;
    /**
     * Name of the instance. This name can have a string of 1 to 63 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Normal defend bandwidth of the instance. The unit is Gbps.
     */
    normalBandwidth: pulumi.Input<number>;
    /**
     * The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
     */
    period?: pulumi.Input<number>;
    /**
     * Type of the instance. Valid values: `Enterprise`, `Professional`. Default to `Enterprise`
     */
    type?: pulumi.Input<string>;
}
