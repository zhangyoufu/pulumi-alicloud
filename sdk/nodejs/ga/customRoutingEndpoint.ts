// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Global Accelerator (GA) Custom Routing Endpoint resource.
 *
 * For information about Global Accelerator (GA) Custom Routing Endpoint and how to use it, see [What is Custom Routing Endpoint](https://www.alibabacloud.com/help/en/global-accelerator/latest/createcustomroutingendpoints).
 *
 * > **NOTE:** Available since v1.197.0.
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
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: "terraform-example",
 *     cidrBlock: "172.17.3.0/24",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vswitchName: "terraform-example",
 *     cidrBlock: "172.17.3.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * });
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
 * const defaultCustomRoutingEndpoint = new alicloud.ga.CustomRoutingEndpoint("defaultCustomRoutingEndpoint", {
 *     endpointGroupId: defaultCustomRoutingEndpointGroup.id,
 *     endpoint: defaultSwitch.id,
 *     type: "PrivateSubNet",
 *     trafficToEndpointPolicy: "DenyAll",
 * });
 * ```
 *
 * ## Import
 *
 * Global Accelerator (GA) Custom Routing Endpoint can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ga/customRoutingEndpoint:CustomRoutingEndpoint example <endpoint_group_id>:<custom_routing_endpoint_id>
 * ```
 */
export class CustomRoutingEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing CustomRoutingEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomRoutingEndpointState, opts?: pulumi.CustomResourceOptions): CustomRoutingEndpoint {
        return new CustomRoutingEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ga/customRoutingEndpoint:CustomRoutingEndpoint';

    /**
     * Returns true if the given object is an instance of CustomRoutingEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomRoutingEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomRoutingEndpoint.__pulumiType;
    }

    /**
     * The ID of the GA instance with which the endpoint is associated.
     */
    public /*out*/ readonly acceleratorId!: pulumi.Output<string>;
    /**
     * The ID of the Custom Routing Endpoint.
     */
    public /*out*/ readonly customRoutingEndpointId!: pulumi.Output<string>;
    /**
     * The ID of the endpoint (vSwitch).
     */
    public readonly endpoint!: pulumi.Output<string>;
    /**
     * The ID of the endpoint group in which to create endpoints.
     */
    public readonly endpointGroupId!: pulumi.Output<string>;
    /**
     * The ID of the listener with which the endpoint is associated.
     */
    public /*out*/ readonly listenerId!: pulumi.Output<string>;
    /**
     * The status of the Custom Routing Endpoint.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The access policy of traffic to the endpoint. Default value: `DenyAll`. Valid values:
     */
    public readonly trafficToEndpointPolicy!: pulumi.Output<string>;
    /**
     * The backend service type of the endpoint. Valid values: `PrivateSubNet`.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a CustomRoutingEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomRoutingEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomRoutingEndpointArgs | CustomRoutingEndpointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomRoutingEndpointState | undefined;
            resourceInputs["acceleratorId"] = state ? state.acceleratorId : undefined;
            resourceInputs["customRoutingEndpointId"] = state ? state.customRoutingEndpointId : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["endpointGroupId"] = state ? state.endpointGroupId : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["trafficToEndpointPolicy"] = state ? state.trafficToEndpointPolicy : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as CustomRoutingEndpointArgs | undefined;
            if ((!args || args.endpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpoint'");
            }
            if ((!args || args.endpointGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointGroupId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["endpoint"] = args ? args.endpoint : undefined;
            resourceInputs["endpointGroupId"] = args ? args.endpointGroupId : undefined;
            resourceInputs["trafficToEndpointPolicy"] = args ? args.trafficToEndpointPolicy : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["acceleratorId"] = undefined /*out*/;
            resourceInputs["customRoutingEndpointId"] = undefined /*out*/;
            resourceInputs["listenerId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomRoutingEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomRoutingEndpoint resources.
 */
export interface CustomRoutingEndpointState {
    /**
     * The ID of the GA instance with which the endpoint is associated.
     */
    acceleratorId?: pulumi.Input<string>;
    /**
     * The ID of the Custom Routing Endpoint.
     */
    customRoutingEndpointId?: pulumi.Input<string>;
    /**
     * The ID of the endpoint (vSwitch).
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The ID of the endpoint group in which to create endpoints.
     */
    endpointGroupId?: pulumi.Input<string>;
    /**
     * The ID of the listener with which the endpoint is associated.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The status of the Custom Routing Endpoint.
     */
    status?: pulumi.Input<string>;
    /**
     * The access policy of traffic to the endpoint. Default value: `DenyAll`. Valid values:
     */
    trafficToEndpointPolicy?: pulumi.Input<string>;
    /**
     * The backend service type of the endpoint. Valid values: `PrivateSubNet`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomRoutingEndpoint resource.
 */
export interface CustomRoutingEndpointArgs {
    /**
     * The ID of the endpoint (vSwitch).
     */
    endpoint: pulumi.Input<string>;
    /**
     * The ID of the endpoint group in which to create endpoints.
     */
    endpointGroupId: pulumi.Input<string>;
    /**
     * The access policy of traffic to the endpoint. Default value: `DenyAll`. Valid values:
     */
    trafficToEndpointPolicy?: pulumi.Input<string>;
    /**
     * The backend service type of the endpoint. Valid values: `PrivateSubNet`.
     */
    type: pulumi.Input<string>;
}
