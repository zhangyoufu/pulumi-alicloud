// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Application Load Balancer resource.
 *
 * > **NOTE:** Available in 1.123.1+
 *
 * > **NOTE:** At present, to avoid some unnecessary regulation confusion, SLB can not support alicloud international account to create `PayByBandwidth` instance.
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
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {vpcName: name});
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/21",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones[0].id),
 *     vswitchName: name,
 * });
 * const defaultApplicationLoadBalancer = new alicloud.slb.ApplicationLoadBalancer("defaultApplicationLoadBalancer", {
 *     loadBalancerName: name,
 *     addressType: "intranet",
 *     loadBalancerSpec: "slb.s2.small",
 *     vswitchId: defaultSwitch.id,
 *     tags: {
 *         info: "create for internet",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Load balancer can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:slb/applicationLoadBalancer:ApplicationLoadBalancer example lb-abc123456
 * ```
 */
export class ApplicationLoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationLoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationLoadBalancerState, opts?: pulumi.CustomResourceOptions): ApplicationLoadBalancer {
        return new ApplicationLoadBalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:slb/applicationLoadBalancer:ApplicationLoadBalancer';

    /**
     * Returns true if the given object is an instance of ApplicationLoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationLoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationLoadBalancer.__pulumiType;
    }

    /**
     * Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * The IP version of the SLB instance to be created, which can be set to `ipv4` or `ipv6` . Default to `ipv4`. Now, only internet instance support `ipv6` address.
     */
    public readonly addressIpVersion!: pulumi.Output<string | undefined>;
    /**
     * The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be `intranet`.
     * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
     * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
     */
    public readonly addressType!: pulumi.Output<string>;
    /**
     * Valid value is between 1 and 1000, If argument `internetChargeType` is `PayByTraffic`, then this value will be ignore.
     */
    public readonly bandwidth!: pulumi.Output<number | undefined>;
    /**
     * Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
     */
    public readonly deleteProtection!: pulumi.Output<string | undefined>;
    /**
     * @deprecated Field 'instance_charge_type' has been deprecated from provider version 1.124. Use 'payment_type' replaces it.
     */
    public readonly instanceChargeType!: pulumi.Output<string>;
    /**
     * Valid values are `PayByBandwidth`, `PayByTraffic`. If this value is `PayByBandwidth`, then argument `addressType` must be `internet`. Default is `PayByTraffic`. If load balancer launched in VPC, this value must be `PayByTraffic`. Before version 1.10.1, the valid values are `paybybandwidth` and `paybytraffic`.
     */
    public readonly internetChargeType!: pulumi.Output<string | undefined>;
    public readonly loadBalancerName!: pulumi.Output<string>;
    /**
     * The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
     * Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: `slb.s1.small`, `slb.s2.small`, `slb.s2.medium`,
     * `slb.s3.small`, `slb.s3.medium`, `slb.s3.large` and `slb.s4.large`.
     */
    public readonly loadBalancerSpec!: pulumi.Output<string>;
    /**
     * The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the [DescribeZone](https://help.aliyun.com/document_detail/27585.htm) API.
     */
    public readonly masterZoneId!: pulumi.Output<string>;
    /**
     * The resource of modification protection. It's effective when modification protection is `ConsoleProtection`.
     */
    public readonly modificationProtectionReason!: pulumi.Output<string | undefined>;
    /**
     * The status of modification protection. Valid values: `ConsoleProtection` and `NonProtection`. Default value is `NonProtection`.
     */
    public readonly modificationProtectionStatus!: pulumi.Output<string>;
    /**
     * @deprecated Field 'name' has been deprecated from provider version 1.123.1. New field 'load_balancer_name' instead
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The billing method of the load balancer. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
     */
    public readonly paymentType!: pulumi.Output<string>;
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
     * @deprecated Field 'specification' has been deprecated from provider version 1.123.1. New field 'load_balancer_spec' instead
     */
    public readonly specification!: pulumi.Output<string>;
    /**
     * The status of slb load balancer. Valid values: `actice` and `inactice`. The system default value is `active`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The VSwitch ID to launch in. If `addressType` is internet, it will be ignore.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;

    /**
     * Create a ApplicationLoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ApplicationLoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationLoadBalancerArgs | ApplicationLoadBalancerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationLoadBalancerState | undefined;
            inputs["address"] = state ? state.address : undefined;
            inputs["addressIpVersion"] = state ? state.addressIpVersion : undefined;
            inputs["addressType"] = state ? state.addressType : undefined;
            inputs["bandwidth"] = state ? state.bandwidth : undefined;
            inputs["deleteProtection"] = state ? state.deleteProtection : undefined;
            inputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            inputs["internetChargeType"] = state ? state.internetChargeType : undefined;
            inputs["loadBalancerName"] = state ? state.loadBalancerName : undefined;
            inputs["loadBalancerSpec"] = state ? state.loadBalancerSpec : undefined;
            inputs["masterZoneId"] = state ? state.masterZoneId : undefined;
            inputs["modificationProtectionReason"] = state ? state.modificationProtectionReason : undefined;
            inputs["modificationProtectionStatus"] = state ? state.modificationProtectionStatus : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["paymentType"] = state ? state.paymentType : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["slaveZoneId"] = state ? state.slaveZoneId : undefined;
            inputs["specification"] = state ? state.specification : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as ApplicationLoadBalancerArgs | undefined;
            inputs["address"] = args ? args.address : undefined;
            inputs["addressIpVersion"] = args ? args.addressIpVersion : undefined;
            inputs["addressType"] = args ? args.addressType : undefined;
            inputs["bandwidth"] = args ? args.bandwidth : undefined;
            inputs["deleteProtection"] = args ? args.deleteProtection : undefined;
            inputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            inputs["internetChargeType"] = args ? args.internetChargeType : undefined;
            inputs["loadBalancerName"] = args ? args.loadBalancerName : undefined;
            inputs["loadBalancerSpec"] = args ? args.loadBalancerSpec : undefined;
            inputs["masterZoneId"] = args ? args.masterZoneId : undefined;
            inputs["modificationProtectionReason"] = args ? args.modificationProtectionReason : undefined;
            inputs["modificationProtectionStatus"] = args ? args.modificationProtectionStatus : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["paymentType"] = args ? args.paymentType : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["slaveZoneId"] = args ? args.slaveZoneId : undefined;
            inputs["specification"] = args ? args.specification : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ApplicationLoadBalancer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationLoadBalancer resources.
 */
export interface ApplicationLoadBalancerState {
    /**
     * Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
     */
    readonly address?: pulumi.Input<string>;
    /**
     * The IP version of the SLB instance to be created, which can be set to `ipv4` or `ipv6` . Default to `ipv4`. Now, only internet instance support `ipv6` address.
     */
    readonly addressIpVersion?: pulumi.Input<string>;
    /**
     * The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be `intranet`.
     * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
     * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
     */
    readonly addressType?: pulumi.Input<string>;
    /**
     * Valid value is between 1 and 1000, If argument `internetChargeType` is `PayByTraffic`, then this value will be ignore.
     */
    readonly bandwidth?: pulumi.Input<number>;
    /**
     * Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
     */
    readonly deleteProtection?: pulumi.Input<string>;
    /**
     * @deprecated Field 'instance_charge_type' has been deprecated from provider version 1.124. Use 'payment_type' replaces it.
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * Valid values are `PayByBandwidth`, `PayByTraffic`. If this value is `PayByBandwidth`, then argument `addressType` must be `internet`. Default is `PayByTraffic`. If load balancer launched in VPC, this value must be `PayByTraffic`. Before version 1.10.1, the valid values are `paybybandwidth` and `paybytraffic`.
     */
    readonly internetChargeType?: pulumi.Input<string>;
    readonly loadBalancerName?: pulumi.Input<string>;
    /**
     * The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
     * Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: `slb.s1.small`, `slb.s2.small`, `slb.s2.medium`,
     * `slb.s3.small`, `slb.s3.medium`, `slb.s3.large` and `slb.s4.large`.
     */
    readonly loadBalancerSpec?: pulumi.Input<string>;
    /**
     * The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the [DescribeZone](https://help.aliyun.com/document_detail/27585.htm) API.
     */
    readonly masterZoneId?: pulumi.Input<string>;
    /**
     * The resource of modification protection. It's effective when modification protection is `ConsoleProtection`.
     */
    readonly modificationProtectionReason?: pulumi.Input<string>;
    /**
     * The status of modification protection. Valid values: `ConsoleProtection` and `NonProtection`. Default value is `NonProtection`.
     */
    readonly modificationProtectionStatus?: pulumi.Input<string>;
    /**
     * @deprecated Field 'name' has been deprecated from provider version 1.123.1. New field 'load_balancer_name' instead
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The billing method of the load balancer. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
     */
    readonly paymentType?: pulumi.Input<string>;
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
     * @deprecated Field 'specification' has been deprecated from provider version 1.123.1. New field 'load_balancer_spec' instead
     */
    readonly specification?: pulumi.Input<string>;
    /**
     * The status of slb load balancer. Valid values: `actice` and `inactice`. The system default value is `active`.
     */
    readonly status?: pulumi.Input<string>;
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
 * The set of arguments for constructing a ApplicationLoadBalancer resource.
 */
export interface ApplicationLoadBalancerArgs {
    /**
     * Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
     */
    readonly address?: pulumi.Input<string>;
    /**
     * The IP version of the SLB instance to be created, which can be set to `ipv4` or `ipv6` . Default to `ipv4`. Now, only internet instance support `ipv6` address.
     */
    readonly addressIpVersion?: pulumi.Input<string>;
    /**
     * The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be `intranet`.
     * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
     * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
     */
    readonly addressType?: pulumi.Input<string>;
    /**
     * Valid value is between 1 and 1000, If argument `internetChargeType` is `PayByTraffic`, then this value will be ignore.
     */
    readonly bandwidth?: pulumi.Input<number>;
    /**
     * Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
     */
    readonly deleteProtection?: pulumi.Input<string>;
    /**
     * @deprecated Field 'instance_charge_type' has been deprecated from provider version 1.124. Use 'payment_type' replaces it.
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * Valid values are `PayByBandwidth`, `PayByTraffic`. If this value is `PayByBandwidth`, then argument `addressType` must be `internet`. Default is `PayByTraffic`. If load balancer launched in VPC, this value must be `PayByTraffic`. Before version 1.10.1, the valid values are `paybybandwidth` and `paybytraffic`.
     */
    readonly internetChargeType?: pulumi.Input<string>;
    readonly loadBalancerName?: pulumi.Input<string>;
    /**
     * The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
     * Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: `slb.s1.small`, `slb.s2.small`, `slb.s2.medium`,
     * `slb.s3.small`, `slb.s3.medium`, `slb.s3.large` and `slb.s4.large`.
     */
    readonly loadBalancerSpec?: pulumi.Input<string>;
    /**
     * The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the [DescribeZone](https://help.aliyun.com/document_detail/27585.htm) API.
     */
    readonly masterZoneId?: pulumi.Input<string>;
    /**
     * The resource of modification protection. It's effective when modification protection is `ConsoleProtection`.
     */
    readonly modificationProtectionReason?: pulumi.Input<string>;
    /**
     * The status of modification protection. Valid values: `ConsoleProtection` and `NonProtection`. Default value is `NonProtection`.
     */
    readonly modificationProtectionStatus?: pulumi.Input<string>;
    /**
     * @deprecated Field 'name' has been deprecated from provider version 1.123.1. New field 'load_balancer_name' instead
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The billing method of the load balancer. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
     */
    readonly paymentType?: pulumi.Input<string>;
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
     * @deprecated Field 'specification' has been deprecated from provider version 1.123.1. New field 'load_balancer_spec' instead
     */
    readonly specification?: pulumi.Input<string>;
    /**
     * The status of slb load balancer. Valid values: `actice` and `inactice`. The system default value is `active`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VSwitch ID to launch in. If `addressType` is internet, it will be ignore.
     */
    readonly vswitchId?: pulumi.Input<string>;
}
