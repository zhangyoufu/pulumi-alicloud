// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Global Accelerator (GA) Endpoint Group resource.
 *
 * For information about Global Accelerator (GA) Endpoint Group and how to use it, see [What is Endpoint Group](https://www.alibabacloud.com/help/en/doc-detail/153259.htm).
 *
 * > **NOTE:** Available in v1.113.0+.
 *
 * > **NOTE:** Listeners that use different protocols support different types of endpoint groups:
 * * For a TCP or UDP listener, you can create only one default endpoint group.
 * * For an HTTP or HTTPS listener, you can create one default endpoint group and one virtual endpoint group. By default, you can create only one virtual endpoint group.
 *   * A default endpoint group refers to the endpoint group that you configure when you create an HTTP or HTTPS listener.
 *   * A virtual endpoint group refers to the endpoint group that you can create on the Endpoint Group page after you create a listener.
 * * After you create a virtual endpoint group for an HTTP or HTTPS listener, you can create a forwarding rule and associate the forwarding rule with the virtual endpoint group. Then, the HTTP or HTTPS listener forwards requests with different destination domain names or paths to the default or virtual endpoint group based on the forwarding rule. This way, you can use one Global Accelerator (GA) instance to accelerate access to multiple domain names or paths. For more information about how to create a forwarding rule, see [Manage forwarding rules](https://www.alibabacloud.com/help/en/doc-detail/204224.htm).
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
 * const deBandwidthPackage = new alicloud.ga.BandwidthPackage("deBandwidthPackage", {
 *     bandwidth: 100,
 *     type: "Basic",
 *     bandwidthType: "Basic",
 *     paymentType: "PayAsYouGo",
 *     billingType: "PayBy95",
 *     ratio: 30,
 * });
 * const deBandwidthPackageAttachment = new alicloud.ga.BandwidthPackageAttachment("deBandwidthPackageAttachment", {
 *     acceleratorId: exampleAccelerator.id,
 *     bandwidthPackageId: deBandwidthPackage.id,
 * });
 * const exampleListener = new alicloud.ga.Listener("exampleListener", {
 *     acceleratorId: exampleAccelerator.id,
 *     portRanges: [{
 *         fromPort: 60,
 *         toPort: 70,
 *     }],
 * }, {
 *     dependsOn: [deBandwidthPackageAttachment],
 * });
 * const exampleEipAddress = new alicloud.ecs.EipAddress("exampleEipAddress", {
 *     bandwidth: "10",
 *     internetChargeType: "PayByBandwidth",
 * });
 * const exampleEndpointGroup = new alicloud.ga.EndpointGroup("exampleEndpointGroup", {
 *     acceleratorId: exampleAccelerator.id,
 *     endpointConfigurations: [{
 *         endpoint: exampleEipAddress.ipAddress,
 *         type: "PublicIp",
 *         weight: 20,
 *     }],
 *     endpointGroupRegion: "cn-hangzhou",
 *     listenerId: exampleListener.id,
 * });
 * ```
 *
 * ## Import
 *
 * Ga Endpoint Group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ga/endpointGroup:EndpointGroup example <id>
 * ```
 */
export class EndpointGroup extends pulumi.CustomResource {
    /**
     * Get an existing EndpointGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointGroupState, opts?: pulumi.CustomResourceOptions): EndpointGroup {
        return new EndpointGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ga/endpointGroup:EndpointGroup';

    /**
     * Returns true if the given object is an instance of EndpointGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointGroup.__pulumiType;
    }

    /**
     * The ID of the Global Accelerator instance to which the endpoint group will be added.
     */
    public readonly acceleratorId!: pulumi.Output<string>;
    /**
     * The description of the endpoint group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The endpointConfigurations of the endpoint group. See the following `Block endpointConfigurations`.
     */
    public readonly endpointConfigurations!: pulumi.Output<outputs.ga.EndpointGroupEndpointConfiguration[]>;
    /**
     * The ID of the region where the endpoint group is deployed.
     */
    public readonly endpointGroupRegion!: pulumi.Output<string>;
    /**
     * The endpoint group type. Valid values: `default`, `virtual`. Default value is `default`.
     */
    public readonly endpointGroupType!: pulumi.Output<string>;
    /**
     * The endpoint request protocol. Valid value: `HTTP`, `HTTPS`.
     */
    public readonly endpointRequestProtocol!: pulumi.Output<string | undefined>;
    /**
     * The interval between two consecutive health checks. Unit: seconds.
     */
    public readonly healthCheckIntervalSeconds!: pulumi.Output<number | undefined>;
    /**
     * The path specified as the destination of the targets for health checks.
     */
    public readonly healthCheckPath!: pulumi.Output<string | undefined>;
    /**
     * The port that is used for health checks.
     */
    public readonly healthCheckPort!: pulumi.Output<number | undefined>;
    /**
     * The protocol that is used to connect to the targets for health checks. Valid values: `http`, `https`, `tcp`.
     */
    public readonly healthCheckProtocol!: pulumi.Output<string | undefined>;
    /**
     * The ID of the listener that is associated with the endpoint group.
     */
    public readonly listenerId!: pulumi.Output<string>;
    /**
     * The name of the endpoint group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Mapping between listening port and forwarding port of boarding point. See the following `Block portOverrides`.
     */
    public readonly portOverrides!: pulumi.Output<outputs.ga.EndpointGroupPortOverrides | undefined>;
    /**
     * The status of the endpoint group.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value is `3`.
     */
    public readonly thresholdCount!: pulumi.Output<number | undefined>;
    /**
     * The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
     */
    public readonly trafficPercentage!: pulumi.Output<number | undefined>;

    /**
     * Create a EndpointGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointGroupArgs | EndpointGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointGroupState | undefined;
            resourceInputs["acceleratorId"] = state ? state.acceleratorId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["endpointConfigurations"] = state ? state.endpointConfigurations : undefined;
            resourceInputs["endpointGroupRegion"] = state ? state.endpointGroupRegion : undefined;
            resourceInputs["endpointGroupType"] = state ? state.endpointGroupType : undefined;
            resourceInputs["endpointRequestProtocol"] = state ? state.endpointRequestProtocol : undefined;
            resourceInputs["healthCheckIntervalSeconds"] = state ? state.healthCheckIntervalSeconds : undefined;
            resourceInputs["healthCheckPath"] = state ? state.healthCheckPath : undefined;
            resourceInputs["healthCheckPort"] = state ? state.healthCheckPort : undefined;
            resourceInputs["healthCheckProtocol"] = state ? state.healthCheckProtocol : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["portOverrides"] = state ? state.portOverrides : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["thresholdCount"] = state ? state.thresholdCount : undefined;
            resourceInputs["trafficPercentage"] = state ? state.trafficPercentage : undefined;
        } else {
            const args = argsOrState as EndpointGroupArgs | undefined;
            if ((!args || args.acceleratorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceleratorId'");
            }
            if ((!args || args.endpointConfigurations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointConfigurations'");
            }
            if ((!args || args.endpointGroupRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointGroupRegion'");
            }
            if ((!args || args.listenerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerId'");
            }
            resourceInputs["acceleratorId"] = args ? args.acceleratorId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endpointConfigurations"] = args ? args.endpointConfigurations : undefined;
            resourceInputs["endpointGroupRegion"] = args ? args.endpointGroupRegion : undefined;
            resourceInputs["endpointGroupType"] = args ? args.endpointGroupType : undefined;
            resourceInputs["endpointRequestProtocol"] = args ? args.endpointRequestProtocol : undefined;
            resourceInputs["healthCheckIntervalSeconds"] = args ? args.healthCheckIntervalSeconds : undefined;
            resourceInputs["healthCheckPath"] = args ? args.healthCheckPath : undefined;
            resourceInputs["healthCheckPort"] = args ? args.healthCheckPort : undefined;
            resourceInputs["healthCheckProtocol"] = args ? args.healthCheckProtocol : undefined;
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["portOverrides"] = args ? args.portOverrides : undefined;
            resourceInputs["thresholdCount"] = args ? args.thresholdCount : undefined;
            resourceInputs["trafficPercentage"] = args ? args.trafficPercentage : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EndpointGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndpointGroup resources.
 */
export interface EndpointGroupState {
    /**
     * The ID of the Global Accelerator instance to which the endpoint group will be added.
     */
    acceleratorId?: pulumi.Input<string>;
    /**
     * The description of the endpoint group.
     */
    description?: pulumi.Input<string>;
    /**
     * The endpointConfigurations of the endpoint group. See the following `Block endpointConfigurations`.
     */
    endpointConfigurations?: pulumi.Input<pulumi.Input<inputs.ga.EndpointGroupEndpointConfiguration>[]>;
    /**
     * The ID of the region where the endpoint group is deployed.
     */
    endpointGroupRegion?: pulumi.Input<string>;
    /**
     * The endpoint group type. Valid values: `default`, `virtual`. Default value is `default`.
     */
    endpointGroupType?: pulumi.Input<string>;
    /**
     * The endpoint request protocol. Valid value: `HTTP`, `HTTPS`.
     */
    endpointRequestProtocol?: pulumi.Input<string>;
    /**
     * The interval between two consecutive health checks. Unit: seconds.
     */
    healthCheckIntervalSeconds?: pulumi.Input<number>;
    /**
     * The path specified as the destination of the targets for health checks.
     */
    healthCheckPath?: pulumi.Input<string>;
    /**
     * The port that is used for health checks.
     */
    healthCheckPort?: pulumi.Input<number>;
    /**
     * The protocol that is used to connect to the targets for health checks. Valid values: `http`, `https`, `tcp`.
     */
    healthCheckProtocol?: pulumi.Input<string>;
    /**
     * The ID of the listener that is associated with the endpoint group.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The name of the endpoint group.
     */
    name?: pulumi.Input<string>;
    /**
     * Mapping between listening port and forwarding port of boarding point. See the following `Block portOverrides`.
     */
    portOverrides?: pulumi.Input<inputs.ga.EndpointGroupPortOverrides>;
    /**
     * The status of the endpoint group.
     */
    status?: pulumi.Input<string>;
    /**
     * The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value is `3`.
     */
    thresholdCount?: pulumi.Input<number>;
    /**
     * The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
     */
    trafficPercentage?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a EndpointGroup resource.
 */
export interface EndpointGroupArgs {
    /**
     * The ID of the Global Accelerator instance to which the endpoint group will be added.
     */
    acceleratorId: pulumi.Input<string>;
    /**
     * The description of the endpoint group.
     */
    description?: pulumi.Input<string>;
    /**
     * The endpointConfigurations of the endpoint group. See the following `Block endpointConfigurations`.
     */
    endpointConfigurations: pulumi.Input<pulumi.Input<inputs.ga.EndpointGroupEndpointConfiguration>[]>;
    /**
     * The ID of the region where the endpoint group is deployed.
     */
    endpointGroupRegion: pulumi.Input<string>;
    /**
     * The endpoint group type. Valid values: `default`, `virtual`. Default value is `default`.
     */
    endpointGroupType?: pulumi.Input<string>;
    /**
     * The endpoint request protocol. Valid value: `HTTP`, `HTTPS`.
     */
    endpointRequestProtocol?: pulumi.Input<string>;
    /**
     * The interval between two consecutive health checks. Unit: seconds.
     */
    healthCheckIntervalSeconds?: pulumi.Input<number>;
    /**
     * The path specified as the destination of the targets for health checks.
     */
    healthCheckPath?: pulumi.Input<string>;
    /**
     * The port that is used for health checks.
     */
    healthCheckPort?: pulumi.Input<number>;
    /**
     * The protocol that is used to connect to the targets for health checks. Valid values: `http`, `https`, `tcp`.
     */
    healthCheckProtocol?: pulumi.Input<string>;
    /**
     * The ID of the listener that is associated with the endpoint group.
     */
    listenerId: pulumi.Input<string>;
    /**
     * The name of the endpoint group.
     */
    name?: pulumi.Input<string>;
    /**
     * Mapping between listening port and forwarding port of boarding point. See the following `Block portOverrides`.
     */
    portOverrides?: pulumi.Input<inputs.ga.EndpointGroupPortOverrides>;
    /**
     * The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value is `3`.
     */
    thresholdCount?: pulumi.Input<number>;
    /**
     * The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
     */
    trafficPercentage?: pulumi.Input<number>;
}
