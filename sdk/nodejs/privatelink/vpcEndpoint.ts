// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Private Link Vpc Endpoint resource.
 *
 * For information about Private Link Vpc Endpoint and how to use it, see [What is Vpc Endpoint](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-createvpcendpoint).
 *
 * > **NOTE:** Available since v1.109.0.
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
 * const exampleVpcEndpointService = new alicloud.privatelink.VpcEndpointService("exampleVpcEndpointService", {
 *     serviceDescription: name,
 *     connectBandwidth: 103,
 *     autoAcceptConnection: false,
 * });
 * const exampleNetwork = new alicloud.vpc.Network("exampleNetwork", {
 *     vpcName: name,
 *     cidrBlock: "10.0.0.0/8",
 * });
 * const exampleSecurityGroup = new alicloud.ecs.SecurityGroup("exampleSecurityGroup", {vpcId: exampleNetwork.id});
 * const exampleVpcEndpoint = new alicloud.privatelink.VpcEndpoint("exampleVpcEndpoint", {
 *     serviceId: exampleVpcEndpointService.id,
 *     securityGroupIds: [exampleSecurityGroup.id],
 *     vpcId: exampleNetwork.id,
 *     vpcEndpointName: name,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Private Link Vpc Endpoint can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:privatelink/vpcEndpoint:VpcEndpoint example <id>
 * ```
 */
export class VpcEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEndpointState, opts?: pulumi.CustomResourceOptions): VpcEndpoint {
        return new VpcEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:privatelink/vpcEndpoint:VpcEndpoint';

    /**
     * Returns true if the given object is an instance of VpcEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcEndpoint.__pulumiType;
    }

    /**
     * The bandwidth of the endpoint connection.  1024 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
     */
    public /*out*/ readonly bandwidth!: pulumi.Output<number>;
    /**
     * The state of the endpoint connection.
     */
    public /*out*/ readonly connectionStatus!: pulumi.Output<string>;
    /**
     * The time when the endpoint was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The service state of the endpoint.
     */
    public /*out*/ readonly endpointBusinessStatus!: pulumi.Output<string>;
    /**
     * The description of the endpoint.
     */
    public readonly endpointDescription!: pulumi.Output<string | undefined>;
    /**
     * The domain name of the endpoint.
     */
    public /*out*/ readonly endpointDomain!: pulumi.Output<string>;
    /**
     * The endpoint type.Only the value: Interface, indicating the Interface endpoint. You can add the service resource types of Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB).
     */
    public readonly endpointType!: pulumi.Output<string>;
    /**
     * Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
     * - **true**: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
     * - **false (default)**: disables user authentication.
     */
    public readonly protectedEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The resource group ID.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The ID of the security group that is associated with the endpoint ENI. The security group can be used to control data transfer between the VPC and the endpoint ENI.The endpoint can be associated with up to 10 security groups.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * The ID of the endpoint service with which the endpoint is associated.
     */
    public readonly serviceId!: pulumi.Output<string | undefined>;
    /**
     * The name of the endpoint service with which the endpoint is associated.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * The state of the endpoint.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The list of tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The name of the endpoint.
     */
    public readonly vpcEndpointName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the VPC to which the endpoint belongs.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only 1 is returned.
     */
    public readonly zonePrivateIpAddressCount!: pulumi.Output<number>;

    /**
     * Create a VpcEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcEndpointArgs | VpcEndpointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcEndpointState | undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["connectionStatus"] = state ? state.connectionStatus : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["endpointBusinessStatus"] = state ? state.endpointBusinessStatus : undefined;
            resourceInputs["endpointDescription"] = state ? state.endpointDescription : undefined;
            resourceInputs["endpointDomain"] = state ? state.endpointDomain : undefined;
            resourceInputs["endpointType"] = state ? state.endpointType : undefined;
            resourceInputs["protectedEnabled"] = state ? state.protectedEnabled : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcEndpointName"] = state ? state.vpcEndpointName : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zonePrivateIpAddressCount"] = state ? state.zonePrivateIpAddressCount : undefined;
        } else {
            const args = argsOrState as VpcEndpointArgs | undefined;
            if ((!args || args.securityGroupIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupIds'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["endpointDescription"] = args ? args.endpointDescription : undefined;
            resourceInputs["endpointType"] = args ? args.endpointType : undefined;
            resourceInputs["protectedEnabled"] = args ? args.protectedEnabled : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcEndpointName"] = args ? args.vpcEndpointName : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["zonePrivateIpAddressCount"] = args ? args.zonePrivateIpAddressCount : undefined;
            resourceInputs["bandwidth"] = undefined /*out*/;
            resourceInputs["connectionStatus"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["endpointBusinessStatus"] = undefined /*out*/;
            resourceInputs["endpointDomain"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEndpoint resources.
 */
export interface VpcEndpointState {
    /**
     * The bandwidth of the endpoint connection.  1024 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * The state of the endpoint connection.
     */
    connectionStatus?: pulumi.Input<string>;
    /**
     * The time when the endpoint was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The service state of the endpoint.
     */
    endpointBusinessStatus?: pulumi.Input<string>;
    /**
     * The description of the endpoint.
     */
    endpointDescription?: pulumi.Input<string>;
    /**
     * The domain name of the endpoint.
     */
    endpointDomain?: pulumi.Input<string>;
    /**
     * The endpoint type.Only the value: Interface, indicating the Interface endpoint. You can add the service resource types of Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB).
     */
    endpointType?: pulumi.Input<string>;
    /**
     * Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
     * - **true**: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
     * - **false (default)**: disables user authentication.
     */
    protectedEnabled?: pulumi.Input<boolean>;
    /**
     * The resource group ID.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The ID of the security group that is associated with the endpoint ENI. The security group can be used to control data transfer between the VPC and the endpoint ENI.The endpoint can be associated with up to 10 security groups.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the endpoint service with which the endpoint is associated.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * The name of the endpoint service with which the endpoint is associated.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * The state of the endpoint.
     */
    status?: pulumi.Input<string>;
    /**
     * The list of tags.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the endpoint.
     */
    vpcEndpointName?: pulumi.Input<string>;
    /**
     * The ID of the VPC to which the endpoint belongs.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only 1 is returned.
     */
    zonePrivateIpAddressCount?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a VpcEndpoint resource.
 */
export interface VpcEndpointArgs {
    /**
     * Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The description of the endpoint.
     */
    endpointDescription?: pulumi.Input<string>;
    /**
     * The endpoint type.Only the value: Interface, indicating the Interface endpoint. You can add the service resource types of Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB).
     */
    endpointType?: pulumi.Input<string>;
    /**
     * Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
     * - **true**: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
     * - **false (default)**: disables user authentication.
     */
    protectedEnabled?: pulumi.Input<boolean>;
    /**
     * The resource group ID.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The ID of the security group that is associated with the endpoint ENI. The security group can be used to control data transfer between the VPC and the endpoint ENI.The endpoint can be associated with up to 10 security groups.
     */
    securityGroupIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the endpoint service with which the endpoint is associated.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * The name of the endpoint service with which the endpoint is associated.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * The list of tags.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the endpoint.
     */
    vpcEndpointName?: pulumi.Input<string>;
    /**
     * The ID of the VPC to which the endpoint belongs.
     */
    vpcId: pulumi.Input<string>;
    /**
     * The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only 1 is returned.
     */
    zonePrivateIpAddressCount?: pulumi.Input<number>;
}
