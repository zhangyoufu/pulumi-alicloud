// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPC Traffic Mirror Session resource. Traffic mirroring session.
 *
 * For information about VPC Traffic Mirror Session and how to use it, see [What is Traffic Mirror Session](https://www.alibabacloud.com/help/en/doc-detail/261364.htm).
 *
 * > **NOTE:** Available since v1.142.0.
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
 * const name = config.get("name") || "tf-example";
 * const defaultInstanceTypes = alicloud.ecs.getInstanceTypes({
 *     instanceTypeFamily: "ecs.g7",
 * });
 * const defaultZones = defaultInstanceTypes.then(defaultInstanceTypes => alicloud.getZones({
 *     availableResourceCreation: "Instance",
 *     availableInstanceType: defaultInstanceTypes.instanceTypes?.[0]?.id,
 * }));
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "10.4.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vswitchName: name,
 *     cidrBlock: "10.4.0.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {
 *     description: name,
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_[0-9]+_[0-9]+_x64*",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const defaultInstance: alicloud.ecs.Instance[] = [];
 * for (const range = {value: 0}; range.value < 2; range.value++) {
 *     defaultInstance.push(new alicloud.ecs.Instance(`defaultInstance-${range.value}`, {
 *         availabilityZone: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *         instanceName: name,
 *         hostName: name,
 *         imageId: defaultImages.then(defaultImages => defaultImages.images?.[0]?.id),
 *         instanceType: defaultInstanceTypes.then(defaultInstanceTypes => defaultInstanceTypes.instanceTypes?.[0]?.id),
 *         securityGroups: [defaultSecurityGroup.id],
 *         vswitchId: defaultSwitch.id,
 *         systemDiskCategory: "cloud_essd",
 *     }));
 * }
 * const defaultEcsNetworkInterface: alicloud.ecs.EcsNetworkInterface[] = [];
 * for (const range = {value: 0}; range.value < 2; range.value++) {
 *     defaultEcsNetworkInterface.push(new alicloud.ecs.EcsNetworkInterface(`defaultEcsNetworkInterface-${range.value}`, {
 *         networkInterfaceName: name,
 *         vswitchId: defaultSwitch.id,
 *         securityGroupIds: [defaultSecurityGroup.id],
 *     }));
 * }
 * const defaultEcsNetworkInterfaceAttachment: alicloud.ecs.EcsNetworkInterfaceAttachment[] = [];
 * for (const range = {value: 0}; range.value < 2; range.value++) {
 *     defaultEcsNetworkInterfaceAttachment.push(new alicloud.ecs.EcsNetworkInterfaceAttachment(`defaultEcsNetworkInterfaceAttachment-${range.value}`, {
 *         instanceId: defaultInstance[range.value].id,
 *         networkInterfaceId: defaultEcsNetworkInterface[range.value].id,
 *     }));
 * }
 * const defaultTrafficMirrorFilter = new alicloud.vpc.TrafficMirrorFilter("defaultTrafficMirrorFilter", {
 *     trafficMirrorFilterName: name,
 *     trafficMirrorFilterDescription: name,
 * });
 * const defaultTrafficMirrorSession = new alicloud.vpc.TrafficMirrorSession("defaultTrafficMirrorSession", {
 *     priority: 1,
 *     virtualNetworkId: 10,
 *     trafficMirrorSessionDescription: name,
 *     trafficMirrorSessionName: name,
 *     trafficMirrorTargetId: defaultEcsNetworkInterfaceAttachment[0].networkInterfaceId,
 *     trafficMirrorSourceIds: [defaultEcsNetworkInterfaceAttachment[1].networkInterfaceId],
 *     trafficMirrorFilterId: defaultTrafficMirrorFilter.id,
 *     trafficMirrorTargetType: "NetworkInterface",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * VPC Traffic Mirror Session can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:vpc/trafficMirrorSession:TrafficMirrorSession example <id>
 * ```
 */
export class TrafficMirrorSession extends pulumi.CustomResource {
    /**
     * Get an existing TrafficMirrorSession resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TrafficMirrorSessionState, opts?: pulumi.CustomResourceOptions): TrafficMirrorSession {
        return new TrafficMirrorSession(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/trafficMirrorSession:TrafficMirrorSession';

    /**
     * Returns true if the given object is an instance of TrafficMirrorSession.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TrafficMirrorSession {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TrafficMirrorSession.__pulumiType;
    }

    /**
     * Whether to PreCheck only this request, value:
     * - **true**: sends a check request and does not create a mirror session. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
     * - **false** (default): Sends a normal request and directly creates a mirror session after checking.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether to enable traffic mirror sessions. default to `false`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Maximum Transmission Unit (MTU).
     */
    public readonly packetLength!: pulumi.Output<number>;
    /**
     * The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The status of the resource.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The tags of this resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The ID of the filter.
     */
    public readonly trafficMirrorFilterId!: pulumi.Output<string>;
    /**
     * The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
     */
    public readonly trafficMirrorSessionDescription!: pulumi.Output<string | undefined>;
    /**
     * The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
     */
    public readonly trafficMirrorSessionName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the image source instance. Currently, the Eni is supported as the image source. The default value of N is 1, that is, only one mirror source can be added to a mirror session.
     */
    public readonly trafficMirrorSourceIds!: pulumi.Output<string[]>;
    /**
     * The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
     */
    public readonly trafficMirrorTargetId!: pulumi.Output<string>;
    /**
     * The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance.
     */
    public readonly trafficMirrorTargetType!: pulumi.Output<string>;
    /**
     * The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
     */
    public readonly virtualNetworkId!: pulumi.Output<number>;

    /**
     * Create a TrafficMirrorSession resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TrafficMirrorSessionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TrafficMirrorSessionArgs | TrafficMirrorSessionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TrafficMirrorSessionState | undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["packetLength"] = state ? state.packetLength : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["trafficMirrorFilterId"] = state ? state.trafficMirrorFilterId : undefined;
            resourceInputs["trafficMirrorSessionDescription"] = state ? state.trafficMirrorSessionDescription : undefined;
            resourceInputs["trafficMirrorSessionName"] = state ? state.trafficMirrorSessionName : undefined;
            resourceInputs["trafficMirrorSourceIds"] = state ? state.trafficMirrorSourceIds : undefined;
            resourceInputs["trafficMirrorTargetId"] = state ? state.trafficMirrorTargetId : undefined;
            resourceInputs["trafficMirrorTargetType"] = state ? state.trafficMirrorTargetType : undefined;
            resourceInputs["virtualNetworkId"] = state ? state.virtualNetworkId : undefined;
        } else {
            const args = argsOrState as TrafficMirrorSessionArgs | undefined;
            if ((!args || args.priority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'priority'");
            }
            if ((!args || args.trafficMirrorFilterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trafficMirrorFilterId'");
            }
            if ((!args || args.trafficMirrorSourceIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trafficMirrorSourceIds'");
            }
            if ((!args || args.trafficMirrorTargetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trafficMirrorTargetId'");
            }
            if ((!args || args.trafficMirrorTargetType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trafficMirrorTargetType'");
            }
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["packetLength"] = args ? args.packetLength : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["trafficMirrorFilterId"] = args ? args.trafficMirrorFilterId : undefined;
            resourceInputs["trafficMirrorSessionDescription"] = args ? args.trafficMirrorSessionDescription : undefined;
            resourceInputs["trafficMirrorSessionName"] = args ? args.trafficMirrorSessionName : undefined;
            resourceInputs["trafficMirrorSourceIds"] = args ? args.trafficMirrorSourceIds : undefined;
            resourceInputs["trafficMirrorTargetId"] = args ? args.trafficMirrorTargetId : undefined;
            resourceInputs["trafficMirrorTargetType"] = args ? args.trafficMirrorTargetType : undefined;
            resourceInputs["virtualNetworkId"] = args ? args.virtualNetworkId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TrafficMirrorSession.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TrafficMirrorSession resources.
 */
export interface TrafficMirrorSessionState {
    /**
     * Whether to PreCheck only this request, value:
     * - **true**: sends a check request and does not create a mirror session. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
     * - **false** (default): Sends a normal request and directly creates a mirror session after checking.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable traffic mirror sessions. default to `false`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Maximum Transmission Unit (MTU).
     */
    packetLength?: pulumi.Input<number>;
    /**
     * The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
     */
    priority?: pulumi.Input<number>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
    /**
     * The tags of this resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the filter.
     */
    trafficMirrorFilterId?: pulumi.Input<string>;
    /**
     * The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
     */
    trafficMirrorSessionDescription?: pulumi.Input<string>;
    /**
     * The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
     */
    trafficMirrorSessionName?: pulumi.Input<string>;
    /**
     * The ID of the image source instance. Currently, the Eni is supported as the image source. The default value of N is 1, that is, only one mirror source can be added to a mirror session.
     */
    trafficMirrorSourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
     */
    trafficMirrorTargetId?: pulumi.Input<string>;
    /**
     * The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance.
     */
    trafficMirrorTargetType?: pulumi.Input<string>;
    /**
     * The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
     */
    virtualNetworkId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a TrafficMirrorSession resource.
 */
export interface TrafficMirrorSessionArgs {
    /**
     * Whether to PreCheck only this request, value:
     * - **true**: sends a check request and does not create a mirror session. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
     * - **false** (default): Sends a normal request and directly creates a mirror session after checking.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable traffic mirror sessions. default to `false`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Maximum Transmission Unit (MTU).
     */
    packetLength?: pulumi.Input<number>;
    /**
     * The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
     */
    priority: pulumi.Input<number>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The tags of this resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the filter.
     */
    trafficMirrorFilterId: pulumi.Input<string>;
    /**
     * The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
     */
    trafficMirrorSessionDescription?: pulumi.Input<string>;
    /**
     * The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
     */
    trafficMirrorSessionName?: pulumi.Input<string>;
    /**
     * The ID of the image source instance. Currently, the Eni is supported as the image source. The default value of N is 1, that is, only one mirror source can be added to a mirror session.
     */
    trafficMirrorSourceIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
     */
    trafficMirrorTargetId: pulumi.Input<string>;
    /**
     * The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance.
     */
    trafficMirrorTargetType: pulumi.Input<string>;
    /**
     * The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
     */
    virtualNetworkId?: pulumi.Input<number>;
}
