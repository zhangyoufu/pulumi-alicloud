// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Global Accelerator (GA) Basic Accelerate IP resource.
 *
 * For information about Global Accelerator (GA) Basic Accelerate IP and how to use it, see [What is Basic Accelerate IP](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createbasicaccelerateip).
 *
 * > **NOTE:** Available since v1.194.0.
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
 * const region = config.get("region") || "cn-hangzhou";
 * const defaultBasicAccelerator = new alicloud.ga.BasicAccelerator("defaultBasicAccelerator", {
 *     duration: 1,
 *     basicAcceleratorName: "terraform-example",
 *     description: "terraform-example",
 *     bandwidthBillingType: "CDT",
 *     autoUseCoupon: "true",
 *     autoPay: true,
 * });
 * const defaultBasicIpSet = new alicloud.ga.BasicIpSet("defaultBasicIpSet", {
 *     acceleratorId: defaultBasicAccelerator.id,
 *     accelerateRegionId: region,
 *     ispType: "BGP",
 *     bandwidth: 5,
 * });
 * const defaultBasicAccelerateIp = new alicloud.ga.BasicAccelerateIp("defaultBasicAccelerateIp", {
 *     acceleratorId: defaultBasicAccelerator.id,
 *     ipSetId: defaultBasicIpSet.id,
 * });
 * ```
 *
 * ## Import
 *
 * Global Accelerator (GA) Basic Accelerate IP can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ga/basicAccelerateIp:BasicAccelerateIp example <id>
 * ```
 */
export class BasicAccelerateIp extends pulumi.CustomResource {
    /**
     * Get an existing BasicAccelerateIp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BasicAccelerateIpState, opts?: pulumi.CustomResourceOptions): BasicAccelerateIp {
        return new BasicAccelerateIp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ga/basicAccelerateIp:BasicAccelerateIp';

    /**
     * Returns true if the given object is an instance of BasicAccelerateIp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BasicAccelerateIp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BasicAccelerateIp.__pulumiType;
    }

    /**
     * The address of the Basic Accelerate IP.
     */
    public /*out*/ readonly accelerateIpAddress!: pulumi.Output<string>;
    /**
     * The ID of the Basic GA instance.
     */
    public readonly acceleratorId!: pulumi.Output<string>;
    /**
     * The ID of the Basic Ip Set.
     */
    public readonly ipSetId!: pulumi.Output<string>;
    /**
     * The status of the Basic Accelerate IP instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a BasicAccelerateIp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BasicAccelerateIpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BasicAccelerateIpArgs | BasicAccelerateIpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BasicAccelerateIpState | undefined;
            resourceInputs["accelerateIpAddress"] = state ? state.accelerateIpAddress : undefined;
            resourceInputs["acceleratorId"] = state ? state.acceleratorId : undefined;
            resourceInputs["ipSetId"] = state ? state.ipSetId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as BasicAccelerateIpArgs | undefined;
            if ((!args || args.acceleratorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceleratorId'");
            }
            if ((!args || args.ipSetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipSetId'");
            }
            resourceInputs["acceleratorId"] = args ? args.acceleratorId : undefined;
            resourceInputs["ipSetId"] = args ? args.ipSetId : undefined;
            resourceInputs["accelerateIpAddress"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BasicAccelerateIp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BasicAccelerateIp resources.
 */
export interface BasicAccelerateIpState {
    /**
     * The address of the Basic Accelerate IP.
     */
    accelerateIpAddress?: pulumi.Input<string>;
    /**
     * The ID of the Basic GA instance.
     */
    acceleratorId?: pulumi.Input<string>;
    /**
     * The ID of the Basic Ip Set.
     */
    ipSetId?: pulumi.Input<string>;
    /**
     * The status of the Basic Accelerate IP instance.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BasicAccelerateIp resource.
 */
export interface BasicAccelerateIpArgs {
    /**
     * The ID of the Basic GA instance.
     */
    acceleratorId: pulumi.Input<string>;
    /**
     * The ID of the Basic Ip Set.
     */
    ipSetId: pulumi.Input<string>;
}
