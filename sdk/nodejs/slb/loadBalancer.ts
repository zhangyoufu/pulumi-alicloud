// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Application Load Balancer resource.
 *
 * > **NOTE:** At present, to avoid some unnecessary regulation confusion, SLB can not support alicloud international account to create "paybybandwidth" instance.
 *
 * > **NOTE:** The supported specifications vary by region. Currently not all regions support guaranteed-performance instances.
 * For more details about guaranteed-performance instance, see [Guaranteed-performance instances](https://www.alibabacloud.com/help/doc-detail/27657.htm).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraformtestslbconfig";
 *
 * const defaultZones = pulumi.output(alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * }, { async: true }));
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     cidrBlock: "172.16.0.0/12",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "172.16.0.0/21",
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultLoadBalancer = new alicloud.slb.LoadBalancer("default", {
 *     specification: "slb.s2.small",
 *     tags: {
 *         tag_a: 1,
 *         tag_b: 2,
 *         tag_c: 3,
 *         tag_d: 4,
 *         tag_e: 5,
 *         tag_f: 6,
 *         tag_g: 7,
 *         tag_h: 8,
 *         tag_i: 9,
 *         tag_j: 10,
 *     },
 *     vswitchId: defaultSwitch.id,
 * });
 * ```
 */
export class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerState, opts?: pulumi.CustomResourceOptions): LoadBalancer {
        return new LoadBalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:slb/loadBalancer:LoadBalancer';

    /**
     * Returns true if the given object is an instance of LoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancer.__pulumiType;
    }

    /**
     * Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to "ipv4". Now, only internet instance support ipv6 address.
     */
    public readonly addressIpVersion!: pulumi.Output<string | undefined>;
    /**
     * The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be "intranet".
     * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
     * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
     */
    public readonly addressType!: pulumi.Output<string>;
    /**
     * Valid
     * value is between 1 and 1000, If argument "internetChargeType" is "paybytraffic", then this value will be ignore.
     */
    public readonly bandwidth!: pulumi.Output<number | undefined>;
    /**
     * Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
     */
    public readonly deleteProtection!: pulumi.Output<string | undefined>;
    /**
     * The billing method of the load balancer. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
     */
    public readonly instanceChargeType!: pulumi.Output<string | undefined>;
    /**
     * Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.
     *
     * @deprecated Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.
     */
    public readonly internet!: pulumi.Output<boolean>;
    /**
     * Valid
     * values are `PayByBandwidth`, `PayByTraffic`. If this value is "PayByBandwidth", then argument "internet" must be "true". Default is "PayByTraffic". If load balancer launched in VPC, this value must be "PayByTraffic".
     * Before version 1.10.1, the valid values are "paybybandwidth" and "paybytraffic".
     */
    public readonly internetChargeType!: pulumi.Output<string | undefined>;
    /**
     * The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     */
    public readonly masterZoneId!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36].
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The Id of resource group which the SLB belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     */
    public readonly slaveZoneId!: pulumi.Output<string>;
    /**
     * The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
     * Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: "slb.s1.small", "slb.s2.small", "slb.s2.medium",
     * "slb.s3.small", "slb.s3.medium", "slb.s3.large" and "slb.s4.large".
     */
    public readonly specification!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The VSwitch ID to launch in. If `addressType` is internet, it will be ignore.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerArgs | LoadBalancerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LoadBalancerState | undefined;
            inputs["address"] = state ? state.address : undefined;
            inputs["addressIpVersion"] = state ? state.addressIpVersion : undefined;
            inputs["addressType"] = state ? state.addressType : undefined;
            inputs["bandwidth"] = state ? state.bandwidth : undefined;
            inputs["deleteProtection"] = state ? state.deleteProtection : undefined;
            inputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            inputs["internet"] = state ? state.internet : undefined;
            inputs["internetChargeType"] = state ? state.internetChargeType : undefined;
            inputs["masterZoneId"] = state ? state.masterZoneId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["slaveZoneId"] = state ? state.slaveZoneId : undefined;
            inputs["specification"] = state ? state.specification : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as LoadBalancerArgs | undefined;
            inputs["address"] = args ? args.address : undefined;
            inputs["addressIpVersion"] = args ? args.addressIpVersion : undefined;
            inputs["addressType"] = args ? args.addressType : undefined;
            inputs["bandwidth"] = args ? args.bandwidth : undefined;
            inputs["deleteProtection"] = args ? args.deleteProtection : undefined;
            inputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            inputs["internet"] = args ? args.internet : undefined;
            inputs["internetChargeType"] = args ? args.internetChargeType : undefined;
            inputs["masterZoneId"] = args ? args.masterZoneId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["slaveZoneId"] = args ? args.slaveZoneId : undefined;
            inputs["specification"] = args ? args.specification : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LoadBalancer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancer resources.
 */
export interface LoadBalancerState {
    /**
     * Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
     */
    readonly address?: pulumi.Input<string>;
    /**
     * The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to "ipv4". Now, only internet instance support ipv6 address.
     */
    readonly addressIpVersion?: pulumi.Input<string>;
    /**
     * The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be "intranet".
     * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
     * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
     */
    readonly addressType?: pulumi.Input<string>;
    /**
     * Valid
     * value is between 1 and 1000, If argument "internetChargeType" is "paybytraffic", then this value will be ignore.
     */
    readonly bandwidth?: pulumi.Input<number>;
    /**
     * Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
     */
    readonly deleteProtection?: pulumi.Input<string>;
    /**
     * The billing method of the load balancer. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.
     *
     * @deprecated Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.
     */
    readonly internet?: pulumi.Input<boolean>;
    /**
     * Valid
     * values are `PayByBandwidth`, `PayByTraffic`. If this value is "PayByBandwidth", then argument "internet" must be "true". Default is "PayByTraffic". If load balancer launched in VPC, this value must be "PayByTraffic".
     * Before version 1.10.1, the valid values are "paybybandwidth" and "paybytraffic".
     */
    readonly internetChargeType?: pulumi.Input<string>;
    /**
     * The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     */
    readonly masterZoneId?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36].
     */
    readonly period?: pulumi.Input<number>;
    /**
     * The Id of resource group which the SLB belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     */
    readonly slaveZoneId?: pulumi.Input<string>;
    /**
     * The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
     * Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: "slb.s1.small", "slb.s2.small", "slb.s2.medium",
     * "slb.s3.small", "slb.s3.medium", "slb.s3.large" and "slb.s4.large".
     */
    readonly specification?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VSwitch ID to launch in. If `addressType` is internet, it will be ignore.
     */
    readonly vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    /**
     * Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
     */
    readonly address?: pulumi.Input<string>;
    /**
     * The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to "ipv4". Now, only internet instance support ipv6 address.
     */
    readonly addressIpVersion?: pulumi.Input<string>;
    /**
     * The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be "intranet".
     * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
     * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
     */
    readonly addressType?: pulumi.Input<string>;
    /**
     * Valid
     * value is between 1 and 1000, If argument "internetChargeType" is "paybytraffic", then this value will be ignore.
     */
    readonly bandwidth?: pulumi.Input<number>;
    /**
     * Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
     */
    readonly deleteProtection?: pulumi.Input<string>;
    /**
     * The billing method of the load balancer. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.
     *
     * @deprecated Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.
     */
    readonly internet?: pulumi.Input<boolean>;
    /**
     * Valid
     * values are `PayByBandwidth`, `PayByTraffic`. If this value is "PayByBandwidth", then argument "internet" must be "true". Default is "PayByTraffic". If load balancer launched in VPC, this value must be "PayByTraffic".
     * Before version 1.10.1, the valid values are "paybybandwidth" and "paybytraffic".
     */
    readonly internetChargeType?: pulumi.Input<string>;
    /**
     * The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     */
    readonly masterZoneId?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36].
     */
    readonly period?: pulumi.Input<number>;
    /**
     * The Id of resource group which the SLB belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     */
    readonly slaveZoneId?: pulumi.Input<string>;
    /**
     * The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
     * Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: "slb.s1.small", "slb.s2.small", "slb.s2.medium",
     * "slb.s3.small", "slb.s3.medium", "slb.s3.large" and "slb.s4.large".
     */
    readonly specification?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VSwitch ID to launch in. If `addressType` is internet, it will be ignore.
     */
    readonly vswitchId?: pulumi.Input<string>;
}
