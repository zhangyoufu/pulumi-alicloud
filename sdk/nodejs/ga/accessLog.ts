// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Global Accelerator (GA) Access Log resource.
 *
 * For information about Global Accelerator (GA) Access Log and how to use it, see [What is Access Log](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-attachlogstoretoendpointgroup).
 *
 * > **NOTE:** Available since v1.187.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const region = config.get("region") || "cn-hangzhou";
 * const _default = new random.index.Integer("default", {
 *     max: 99999,
 *     min: 10000,
 * });
 * const defaultProject = new alicloud.log.Project("default", {name: `terraform-example-${_default.result}`});
 * const defaultStore = new alicloud.log.Store("default", {
 *     project: defaultProject.name,
 *     name: "terraform-example",
 * });
 * const defaultAccelerator = new alicloud.ga.Accelerator("default", {
 *     duration: 1,
 *     autoUseCoupon: true,
 *     spec: "2",
 * });
 * const defaultBandwidthPackage = new alicloud.ga.BandwidthPackage("default", {
 *     bandwidth: 100,
 *     type: "Basic",
 *     bandwidthType: "Basic",
 *     paymentType: "PayAsYouGo",
 *     billingType: "PayBy95",
 *     ratio: 30,
 * });
 * const defaultBandwidthPackageAttachment = new alicloud.ga.BandwidthPackageAttachment("default", {
 *     acceleratorId: defaultAccelerator.id,
 *     bandwidthPackageId: defaultBandwidthPackage.id,
 * });
 * const defaultListener = new alicloud.ga.Listener("default", {
 *     acceleratorId: defaultBandwidthPackageAttachment.acceleratorId,
 *     clientAffinity: "SOURCE_IP",
 *     protocol: "HTTP",
 *     name: "terraform-example",
 *     portRanges: [{
 *         fromPort: 70,
 *         toPort: 70,
 *     }],
 * });
 * const defaultEipAddress = new alicloud.ecs.EipAddress("default", {
 *     bandwidth: "10",
 *     internetChargeType: "PayByBandwidth",
 *     addressName: "terraform-example",
 * });
 * const defaultEndpointGroup = new alicloud.ga.EndpointGroup("default", {
 *     acceleratorId: defaultListener.acceleratorId,
 *     endpointConfigurations: [{
 *         endpoint: defaultEipAddress.ipAddress,
 *         type: "PublicIp",
 *         weight: 20,
 *     }],
 *     endpointGroupRegion: region,
 *     listenerId: defaultListener.id,
 * });
 * const defaultAccessLog = new alicloud.ga.AccessLog("default", {
 *     acceleratorId: defaultAccelerator.id,
 *     listenerId: defaultListener.id,
 *     endpointGroupId: defaultEndpointGroup.id,
 *     slsProjectName: defaultProject.name,
 *     slsLogStoreName: defaultStore.name,
 *     slsRegionId: region,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Global Accelerator (GA) Access Log can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ga/accessLog:AccessLog example <accelerator_id>:<listener_id>:<endpoint_group_id>
 * ```
 */
export class AccessLog extends pulumi.CustomResource {
    /**
     * Get an existing AccessLog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessLogState, opts?: pulumi.CustomResourceOptions): AccessLog {
        return new AccessLog(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ga/accessLog:AccessLog';

    /**
     * Returns true if the given object is an instance of AccessLog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessLog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessLog.__pulumiType;
    }

    /**
     * The ID of the global acceleration instance.
     */
    public readonly acceleratorId!: pulumi.Output<string>;
    /**
     * The ID of the endpoint group instance.
     */
    public readonly endpointGroupId!: pulumi.Output<string>;
    /**
     * The ID of the listener.
     */
    public readonly listenerId!: pulumi.Output<string>;
    /**
     * The name of the Log Store.
     */
    public readonly slsLogStoreName!: pulumi.Output<string>;
    /**
     * The name of the Log Service project.
     */
    public readonly slsProjectName!: pulumi.Output<string>;
    /**
     * The region ID of the Log Service project.
     */
    public readonly slsRegionId!: pulumi.Output<string>;
    /**
     * Whether access log is enabled.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a AccessLog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessLogArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessLogArgs | AccessLogState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessLogState | undefined;
            resourceInputs["acceleratorId"] = state ? state.acceleratorId : undefined;
            resourceInputs["endpointGroupId"] = state ? state.endpointGroupId : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["slsLogStoreName"] = state ? state.slsLogStoreName : undefined;
            resourceInputs["slsProjectName"] = state ? state.slsProjectName : undefined;
            resourceInputs["slsRegionId"] = state ? state.slsRegionId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as AccessLogArgs | undefined;
            if ((!args || args.acceleratorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceleratorId'");
            }
            if ((!args || args.endpointGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointGroupId'");
            }
            if ((!args || args.listenerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerId'");
            }
            if ((!args || args.slsLogStoreName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slsLogStoreName'");
            }
            if ((!args || args.slsProjectName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slsProjectName'");
            }
            if ((!args || args.slsRegionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slsRegionId'");
            }
            resourceInputs["acceleratorId"] = args ? args.acceleratorId : undefined;
            resourceInputs["endpointGroupId"] = args ? args.endpointGroupId : undefined;
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["slsLogStoreName"] = args ? args.slsLogStoreName : undefined;
            resourceInputs["slsProjectName"] = args ? args.slsProjectName : undefined;
            resourceInputs["slsRegionId"] = args ? args.slsRegionId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccessLog.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessLog resources.
 */
export interface AccessLogState {
    /**
     * The ID of the global acceleration instance.
     */
    acceleratorId?: pulumi.Input<string>;
    /**
     * The ID of the endpoint group instance.
     */
    endpointGroupId?: pulumi.Input<string>;
    /**
     * The ID of the listener.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The name of the Log Store.
     */
    slsLogStoreName?: pulumi.Input<string>;
    /**
     * The name of the Log Service project.
     */
    slsProjectName?: pulumi.Input<string>;
    /**
     * The region ID of the Log Service project.
     */
    slsRegionId?: pulumi.Input<string>;
    /**
     * Whether access log is enabled.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessLog resource.
 */
export interface AccessLogArgs {
    /**
     * The ID of the global acceleration instance.
     */
    acceleratorId: pulumi.Input<string>;
    /**
     * The ID of the endpoint group instance.
     */
    endpointGroupId: pulumi.Input<string>;
    /**
     * The ID of the listener.
     */
    listenerId: pulumi.Input<string>;
    /**
     * The name of the Log Store.
     */
    slsLogStoreName: pulumi.Input<string>;
    /**
     * The name of the Log Service project.
     */
    slsProjectName: pulumi.Input<string>;
    /**
     * The region ID of the Log Service project.
     */
    slsRegionId: pulumi.Input<string>;
}
