// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Private Link Vpc Endpoint Connection resource. vpc endpoint connection.
 *
 * For information about Private Link Vpc Endpoint Connection and how to use it, see [What is Vpc Endpoint Connection](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-enablevpcendpointzoneconnection).
 *
 * > **NOTE:** Available since v1.110.0.
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
 * const name = config.get("name") || "tf_example";
 * const example = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const exampleVpcEndpointService = new alicloud.privatelink.VpcEndpointService("example", {
 *     serviceDescription: name,
 *     connectBandwidth: 103,
 *     autoAcceptConnection: false,
 * });
 * const exampleNetwork = new alicloud.vpc.Network("example", {
 *     vpcName: name,
 *     cidrBlock: "10.0.0.0/8",
 * });
 * const exampleSwitch = new alicloud.vpc.Switch("example", {
 *     vswitchName: name,
 *     cidrBlock: "10.1.0.0/16",
 *     vpcId: exampleNetwork.id,
 *     zoneId: example.then(example => example.zones?.[0]?.id),
 * });
 * const exampleSecurityGroup = new alicloud.ecs.SecurityGroup("example", {
 *     name: name,
 *     vpcId: exampleNetwork.id,
 * });
 * const exampleApplicationLoadBalancer = new alicloud.slb.ApplicationLoadBalancer("example", {
 *     loadBalancerName: name,
 *     vswitchId: exampleSwitch.id,
 *     loadBalancerSpec: "slb.s2.small",
 *     addressType: "intranet",
 * });
 * const exampleVpcEndpointServiceResource = new alicloud.privatelink.VpcEndpointServiceResource("example", {
 *     serviceId: exampleVpcEndpointService.id,
 *     resourceId: exampleApplicationLoadBalancer.id,
 *     resourceType: "slb",
 * });
 * const exampleVpcEndpoint = new alicloud.privatelink.VpcEndpoint("example", {
 *     serviceId: exampleVpcEndpointServiceResource.serviceId,
 *     securityGroupIds: [exampleSecurityGroup.id],
 *     vpcId: exampleNetwork.id,
 *     vpcEndpointName: name,
 * });
 * const exampleVpcEndpointServiceConnection = new alicloud.privatelink.VpcEndpointServiceConnection("example", {
 *     endpointId: exampleVpcEndpoint.id,
 *     serviceId: exampleVpcEndpoint.serviceId,
 *     bandwidth: 1024,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Private Link Vpc Endpoint Connection can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection example <service_id>:<endpoint_id>
 * ```
 */
export class VpcEndpointServiceConnection extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpointServiceConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEndpointServiceConnectionState, opts?: pulumi.CustomResourceOptions): VpcEndpointServiceConnection {
        return new VpcEndpointServiceConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection';

    /**
     * Returns true if the given object is an instance of VpcEndpointServiceConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcEndpointServiceConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcEndpointServiceConnection.__pulumiType;
    }

    /**
     * The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The endpoint ID.
     */
    public readonly endpointId!: pulumi.Output<string>;
    /**
     * The endpoint service ID.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * The state of the endpoint connection.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a VpcEndpointServiceConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcEndpointServiceConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcEndpointServiceConnectionArgs | VpcEndpointServiceConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcEndpointServiceConnectionState | undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["endpointId"] = state ? state.endpointId : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as VpcEndpointServiceConnectionArgs | undefined;
            if ((!args || args.endpointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointId'");
            }
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["endpointId"] = args ? args.endpointId : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcEndpointServiceConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEndpointServiceConnection resources.
 */
export interface VpcEndpointServiceConnectionState {
    /**
     * The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The endpoint ID.
     */
    endpointId?: pulumi.Input<string>;
    /**
     * The endpoint service ID.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * The state of the endpoint connection.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcEndpointServiceConnection resource.
 */
export interface VpcEndpointServiceConnectionArgs {
    /**
     * The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The endpoint ID.
     */
    endpointId: pulumi.Input<string>;
    /**
     * The endpoint service ID.
     */
    serviceId: pulumi.Input<string>;
}
