// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Global Accelerator (GA) Custom Routing Endpoint Group Destination resource.
 *
 * For information about Global Accelerator (GA) Custom Routing Endpoint Group Destination and how to use it, see [What is Custom Routing Endpoint Group Destination](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createcustomroutingendpointgroupdestinations).
 *
 * > **NOTE:** Available since v1.197.0.
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
 * const config = new pulumi.Config();
 * const region = config.get("region") || "cn-hangzhou";
 * const defaultAccelerator = new alicloud.ga.Accelerator("defaultAccelerator", {
 *     duration: 1,
 *     autoUseCoupon: true,
 *     spec: "1",
 * });
 * const defaultBandwidthPackage = new alicloud.ga.BandwidthPackage("defaultBandwidthPackage", {
 *     bandwidth: 100,
 *     type: "Basic",
 *     bandwidthType: "Basic",
 *     paymentType: "PayAsYouGo",
 *     billingType: "PayBy95",
 *     ratio: 30,
 * });
 * const defaultBandwidthPackageAttachment = new alicloud.ga.BandwidthPackageAttachment("defaultBandwidthPackageAttachment", {
 *     acceleratorId: defaultAccelerator.id,
 *     bandwidthPackageId: defaultBandwidthPackage.id,
 * });
 * const defaultListener = new alicloud.ga.Listener("defaultListener", {
 *     acceleratorId: defaultBandwidthPackageAttachment.acceleratorId,
 *     listenerType: "CustomRouting",
 *     portRanges: [{
 *         fromPort: 10000,
 *         toPort: 16000,
 *     }],
 * });
 * const defaultCustomRoutingEndpointGroup = new alicloud.ga.CustomRoutingEndpointGroup("defaultCustomRoutingEndpointGroup", {
 *     acceleratorId: defaultListener.acceleratorId,
 *     listenerId: defaultListener.id,
 *     endpointGroupRegion: region,
 *     customRoutingEndpointGroupName: "terraform-example",
 *     description: "terraform-example",
 * });
 * const defaultCustomRoutingEndpointGroupDestination = new alicloud.ga.CustomRoutingEndpointGroupDestination("defaultCustomRoutingEndpointGroupDestination", {
 *     endpointGroupId: defaultCustomRoutingEndpointGroup.id,
 *     protocols: ["TCP"],
 *     fromPort: 1,
 *     toPort: 2,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Global Accelerator (GA) Custom Routing Endpoint Group Destination can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ga/customRoutingEndpointGroupDestination:CustomRoutingEndpointGroupDestination example <endpoint_group_id>:<custom_routing_endpoint_group_destination_id>
 * ```
 */
export class CustomRoutingEndpointGroupDestination extends pulumi.CustomResource {
    /**
     * Get an existing CustomRoutingEndpointGroupDestination resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomRoutingEndpointGroupDestinationState, opts?: pulumi.CustomResourceOptions): CustomRoutingEndpointGroupDestination {
        return new CustomRoutingEndpointGroupDestination(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ga/customRoutingEndpointGroupDestination:CustomRoutingEndpointGroupDestination';

    /**
     * Returns true if the given object is an instance of CustomRoutingEndpointGroupDestination.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomRoutingEndpointGroupDestination {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomRoutingEndpointGroupDestination.__pulumiType;
    }

    /**
     * The ID of the GA instance.
     */
    public /*out*/ readonly acceleratorId!: pulumi.Output<string>;
    /**
     * The ID of the Custom Routing Endpoint Group Destination.
     */
    public /*out*/ readonly customRoutingEndpointGroupDestinationId!: pulumi.Output<string>;
    /**
     * The ID of the endpoint group.
     */
    public readonly endpointGroupId!: pulumi.Output<string>;
    /**
     * The start port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
     */
    public readonly fromPort!: pulumi.Output<number>;
    /**
     * The ID of the listener.
     */
    public /*out*/ readonly listenerId!: pulumi.Output<string>;
    /**
     * The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
     */
    public readonly protocols!: pulumi.Output<string[]>;
    /**
     * The status of the Custom Routing Endpoint Group Destination.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The end port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
     */
    public readonly toPort!: pulumi.Output<number>;

    /**
     * Create a CustomRoutingEndpointGroupDestination resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomRoutingEndpointGroupDestinationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomRoutingEndpointGroupDestinationArgs | CustomRoutingEndpointGroupDestinationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomRoutingEndpointGroupDestinationState | undefined;
            resourceInputs["acceleratorId"] = state ? state.acceleratorId : undefined;
            resourceInputs["customRoutingEndpointGroupDestinationId"] = state ? state.customRoutingEndpointGroupDestinationId : undefined;
            resourceInputs["endpointGroupId"] = state ? state.endpointGroupId : undefined;
            resourceInputs["fromPort"] = state ? state.fromPort : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["protocols"] = state ? state.protocols : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["toPort"] = state ? state.toPort : undefined;
        } else {
            const args = argsOrState as CustomRoutingEndpointGroupDestinationArgs | undefined;
            if ((!args || args.endpointGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointGroupId'");
            }
            if ((!args || args.fromPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fromPort'");
            }
            if ((!args || args.protocols === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocols'");
            }
            if ((!args || args.toPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'toPort'");
            }
            resourceInputs["endpointGroupId"] = args ? args.endpointGroupId : undefined;
            resourceInputs["fromPort"] = args ? args.fromPort : undefined;
            resourceInputs["protocols"] = args ? args.protocols : undefined;
            resourceInputs["toPort"] = args ? args.toPort : undefined;
            resourceInputs["acceleratorId"] = undefined /*out*/;
            resourceInputs["customRoutingEndpointGroupDestinationId"] = undefined /*out*/;
            resourceInputs["listenerId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomRoutingEndpointGroupDestination.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomRoutingEndpointGroupDestination resources.
 */
export interface CustomRoutingEndpointGroupDestinationState {
    /**
     * The ID of the GA instance.
     */
    acceleratorId?: pulumi.Input<string>;
    /**
     * The ID of the Custom Routing Endpoint Group Destination.
     */
    customRoutingEndpointGroupDestinationId?: pulumi.Input<string>;
    /**
     * The ID of the endpoint group.
     */
    endpointGroupId?: pulumi.Input<string>;
    /**
     * The start port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
     */
    fromPort?: pulumi.Input<number>;
    /**
     * The ID of the listener.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
     */
    protocols?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of the Custom Routing Endpoint Group Destination.
     */
    status?: pulumi.Input<string>;
    /**
     * The end port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
     */
    toPort?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CustomRoutingEndpointGroupDestination resource.
 */
export interface CustomRoutingEndpointGroupDestinationArgs {
    /**
     * The ID of the endpoint group.
     */
    endpointGroupId: pulumi.Input<string>;
    /**
     * The start port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
     */
    fromPort: pulumi.Input<number>;
    /**
     * The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
     */
    protocols: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The end port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
     */
    toPort: pulumi.Input<number>;
}
