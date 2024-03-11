// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Private Link Vpc Endpoint Service Resource resource. Endpoint service resource.
 *
 * For information about Private Link Vpc Endpoint Service Resource and how to use it, see [What is Vpc Endpoint Service Resource](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-attachresourcetovpcendpointservice).
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
 * const exampleZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const exampleVpcEndpointService = new alicloud.privatelink.VpcEndpointService("exampleVpcEndpointService", {
 *     serviceDescription: name,
 *     connectBandwidth: 103,
 *     autoAcceptConnection: false,
 * });
 * const exampleNetwork = new alicloud.vpc.Network("exampleNetwork", {
 *     vpcName: name,
 *     cidrBlock: "10.0.0.0/8",
 * });
 * const exampleSwitch = new alicloud.vpc.Switch("exampleSwitch", {
 *     vswitchName: name,
 *     cidrBlock: "10.1.0.0/16",
 *     vpcId: exampleNetwork.id,
 *     zoneId: exampleZones.then(exampleZones => exampleZones.zones?.[0]?.id),
 * });
 * const exampleSecurityGroup = new alicloud.ecs.SecurityGroup("exampleSecurityGroup", {vpcId: exampleNetwork.id});
 * const exampleApplicationLoadBalancer = new alicloud.slb.ApplicationLoadBalancer("exampleApplicationLoadBalancer", {
 *     loadBalancerName: name,
 *     vswitchId: exampleSwitch.id,
 *     loadBalancerSpec: "slb.s2.small",
 *     addressType: "intranet",
 * });
 * const exampleVpcEndpoint = new alicloud.privatelink.VpcEndpoint("exampleVpcEndpoint", {
 *     serviceId: exampleVpcEndpointService.id,
 *     securityGroupIds: [exampleSecurityGroup.id],
 *     vpcId: exampleNetwork.id,
 *     vpcEndpointName: name,
 * });
 * const exampleVpcEndpointServiceResource = new alicloud.privatelink.VpcEndpointServiceResource("exampleVpcEndpointServiceResource", {
 *     serviceId: exampleVpcEndpointService.id,
 *     resourceId: exampleApplicationLoadBalancer.id,
 *     resourceType: "slb",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Private Link Vpc Endpoint Service Resource can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource example <service_id>:<resource_id>:<zone_id>
 * ```
 */
export class VpcEndpointServiceResource extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpointServiceResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEndpointServiceResourceState, opts?: pulumi.CustomResourceOptions): VpcEndpointServiceResource {
        return new VpcEndpointServiceResource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource';

    /**
     * Returns true if the given object is an instance of VpcEndpointServiceResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcEndpointServiceResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcEndpointServiceResource.__pulumiType;
    }

    /**
     * Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The service resource ID.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * Service resource type, value:
     * - **slb**: indicates that the service resource type is Classic Load Balancer (CLB).
     * - **alb**: indicates that the service resource type is Application Load Balancer (ALB).
     * - **nlb**: indicates that the service resource type is Network Load Balancer (NLB).
     */
    public readonly resourceType!: pulumi.Output<string>;
    /**
     * The endpoint service ID.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * The ID of the zone to which the service resource belongs. (valid when the resource type is nlb/alb).
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a VpcEndpointServiceResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcEndpointServiceResourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcEndpointServiceResourceArgs | VpcEndpointServiceResourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcEndpointServiceResourceState | undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as VpcEndpointServiceResourceArgs | undefined;
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcEndpointServiceResource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEndpointServiceResource resources.
 */
export interface VpcEndpointServiceResourceState {
    /**
     * Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The service resource ID.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * Service resource type, value:
     * - **slb**: indicates that the service resource type is Classic Load Balancer (CLB).
     * - **alb**: indicates that the service resource type is Application Load Balancer (ALB).
     * - **nlb**: indicates that the service resource type is Network Load Balancer (NLB).
     */
    resourceType?: pulumi.Input<string>;
    /**
     * The endpoint service ID.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * The ID of the zone to which the service resource belongs. (valid when the resource type is nlb/alb).
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcEndpointServiceResource resource.
 */
export interface VpcEndpointServiceResourceArgs {
    /**
     * Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The service resource ID.
     */
    resourceId: pulumi.Input<string>;
    /**
     * Service resource type, value:
     * - **slb**: indicates that the service resource type is Classic Load Balancer (CLB).
     * - **alb**: indicates that the service resource type is Application Load Balancer (ALB).
     * - **nlb**: indicates that the service resource type is Network Load Balancer (NLB).
     */
    resourceType: pulumi.Input<string>;
    /**
     * The endpoint service ID.
     */
    serviceId: pulumi.Input<string>;
    /**
     * The ID of the zone to which the service resource belongs. (valid when the resource type is nlb/alb).
     */
    zoneId?: pulumi.Input<string>;
}
