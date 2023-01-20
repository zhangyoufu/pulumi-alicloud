// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Global Accelerator (GA) Bandwidth Package Attachment resource.
 *
 * For information about Global Accelerator (GA) Bandwidth Package Attachment and how to use it, see [What is Bandwidth Package Attachment](https://www.alibabacloud.com/help/en/global-accelerator/latest/bandwidthpackageaddaccelerator).
 *
 * > **NOTE:** Available in v1.113.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const exampleAccelerator = new alicloud.ga.Accelerator("exampleAccelerator", {
 *     duration: 1,
 *     autoUseCoupon: true,
 *     spec: "1",
 * });
 * const exampleBandwidthPackage = new alicloud.ga.BandwidthPackage("exampleBandwidthPackage", {
 *     bandwidth: 20,
 *     type: "Basic",
 *     bandwidthType: "Basic",
 *     duration: "1",
 *     autoPay: true,
 *     ratio: 30,
 * });
 * const exampleBandwidthPackageAttachment = new alicloud.ga.BandwidthPackageAttachment("exampleBandwidthPackageAttachment", {
 *     acceleratorId: exampleAccelerator.id,
 *     bandwidthPackageId: exampleBandwidthPackage.id,
 * });
 * ```
 *
 * ## Import
 *
 * Ga Bandwidth Package Attachment can be imported using the id. Format to `<accelerator_id>:<bandwidth_package_id>`, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ga/bandwidthPackageAttachment:BandwidthPackageAttachment example your_accelerator_id:your_bandwidth_package_id
 * ```
 */
export class BandwidthPackageAttachment extends pulumi.CustomResource {
    /**
     * Get an existing BandwidthPackageAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BandwidthPackageAttachmentState, opts?: pulumi.CustomResourceOptions): BandwidthPackageAttachment {
        return new BandwidthPackageAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ga/bandwidthPackageAttachment:BandwidthPackageAttachment';

    /**
     * Returns true if the given object is an instance of BandwidthPackageAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BandwidthPackageAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BandwidthPackageAttachment.__pulumiType;
    }

    /**
     * The ID of the Global Accelerator instance from which you want to disassociate the bandwidth plan.
     */
    public readonly acceleratorId!: pulumi.Output<string>;
    /**
     * Accelerators bound with current Bandwidth Package.
     */
    public /*out*/ readonly accelerators!: pulumi.Output<string[]>;
    /**
     * The ID of the bandwidth plan to disassociate. **NOTE:** From version 1.192.0, `bandwidthPackageId` can be modified.
     */
    public readonly bandwidthPackageId!: pulumi.Output<string>;
    /**
     * State of Bandwidth Package.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a BandwidthPackageAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BandwidthPackageAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BandwidthPackageAttachmentArgs | BandwidthPackageAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BandwidthPackageAttachmentState | undefined;
            resourceInputs["acceleratorId"] = state ? state.acceleratorId : undefined;
            resourceInputs["accelerators"] = state ? state.accelerators : undefined;
            resourceInputs["bandwidthPackageId"] = state ? state.bandwidthPackageId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as BandwidthPackageAttachmentArgs | undefined;
            if ((!args || args.acceleratorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceleratorId'");
            }
            if ((!args || args.bandwidthPackageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidthPackageId'");
            }
            resourceInputs["acceleratorId"] = args ? args.acceleratorId : undefined;
            resourceInputs["bandwidthPackageId"] = args ? args.bandwidthPackageId : undefined;
            resourceInputs["accelerators"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BandwidthPackageAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BandwidthPackageAttachment resources.
 */
export interface BandwidthPackageAttachmentState {
    /**
     * The ID of the Global Accelerator instance from which you want to disassociate the bandwidth plan.
     */
    acceleratorId?: pulumi.Input<string>;
    /**
     * Accelerators bound with current Bandwidth Package.
     */
    accelerators?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the bandwidth plan to disassociate. **NOTE:** From version 1.192.0, `bandwidthPackageId` can be modified.
     */
    bandwidthPackageId?: pulumi.Input<string>;
    /**
     * State of Bandwidth Package.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BandwidthPackageAttachment resource.
 */
export interface BandwidthPackageAttachmentArgs {
    /**
     * The ID of the Global Accelerator instance from which you want to disassociate the bandwidth plan.
     */
    acceleratorId: pulumi.Input<string>;
    /**
     * The ID of the bandwidth plan to disassociate. **NOTE:** From version 1.192.0, `bandwidthPackageId` can be modified.
     */
    bandwidthPackageId: pulumi.Input<string>;
}
