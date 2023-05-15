// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a EIP Address resource.
 *
 * For information about EIP Address and how to use it, see [What is EIP Address](https://www.alibabacloud.com/help/en/doc-detail/36016.htm).
 *
 * > **NOTE:** Available in v1.126.0+.
 *
 * > **NOTE:** BGP (Multi-ISP) lines are supported in all regions. BGP (Multi-ISP) Pro lines are supported only in the China (Hong Kong) region.
 *
 * > **NOTE:** The resource only supports to create `PayAsYouGo PayByTraffic`  or `Subscription PayByBandwidth` elastic IP for international account. Otherwise, you will happened error `COMMODITY.INVALID_COMPONENT`.
 * Your account is international if you can use it to login in [International Web Console](https://account.alibabacloud.com/login/login.htm).
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.ecs.EipAddress("example", {
 *     addressName: "tf-testAcc1234",
 *     internetChargeType: "PayByBandwidth",
 *     isp: "BGP",
 *     paymentType: "PayAsYouGo",
 * });
 * ```
 *
 * ## Import
 *
 * EIP Address can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ecs/eipAddress:EipAddress example <id>
 * ```
 */
export class EipAddress extends pulumi.CustomResource {
    /**
     * Get an existing EipAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EipAddressState, opts?: pulumi.CustomResourceOptions): EipAddress {
        return new EipAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/eipAddress:EipAddress';

    /**
     * Returns true if the given object is an instance of EipAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EipAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EipAddress.__pulumiType;
    }

    /**
     * The activity id.
     */
    public readonly activityId!: pulumi.Output<string | undefined>;
    /**
     * The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
     */
    public readonly addressName!: pulumi.Output<string>;
    /**
     * Whether to pay automatically. Valid values: `true` and `false`. Default value: `true`. When `autoPay` is `true`, The order will be automatically paid. When `autoPay` is `false`, The order needs to go to the order center to complete the payment. **NOTE:** When `paymentType` is `Subscription`, this parameter is valid.
     */
    public readonly autoPay!: pulumi.Output<boolean | undefined>;
    /**
     * The maximum bandwidth of the EIP. Valid values: `1` to `200`. Unit: Mbit/s. Default value: `5`.
     */
    public readonly bandwidth!: pulumi.Output<string>;
    /**
     * Whether enable the deletion protection or not. Default value: `false`.
     */
    public readonly deletionProtection!: pulumi.Output<boolean>;
    /**
     * The description of the EIP.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The status of the EIP. configuring high precision second-by-second monitoring for EIP. Valid values: `ON` and `OFF`.
     */
    public readonly highDefinitionMonitorLogStatus!: pulumi.Output<string>;
    /**
     * Field `instanceChargeType` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `paymentType` instead.
     *
     * @deprecated Field 'instance_charge_type' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
     */
    public readonly instanceChargeType!: pulumi.Output<string>;
    /**
     * The metering method of the EIP. 
     * Valid values: `PayByDominantTraffic`, `PayByBandwidth` and `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** It must be set to "PayByBandwidth" when `paymentType` is "Subscription".
     */
    public readonly internetChargeType!: pulumi.Output<string>;
    /**
     * The address of the EIP.
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * The line type. You can set this parameter only when you create a `PayAsYouGo` EIP. Valid values:
     */
    public readonly isp!: pulumi.Output<string>;
    /**
     * The Name of the logging service LogProject. Current parameter is required when configuring high precision second-by-second monitoring for EIP.
     */
    public readonly logProject!: pulumi.Output<string | undefined>;
    /**
     * The Name of the logging service LogStore. Current parameter is required when configuring high precision second-by-second monitoring for EIP.
     */
    public readonly logStore!: pulumi.Output<string | undefined>;
    /**
     * Field `name` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `addressName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'address_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of the network. Valid value is `public` (Internet).
     */
    public readonly netmode!: pulumi.Output<string | undefined>;
    /**
     * The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
     */
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The ID of the IP address pool. The EIP is allocated from the IP address pool. **NOTE:** The feature is available only to users whose accounts are included in the whitelist. If you want to use the feature,[submit a ticket](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/429100).
     */
    public readonly publicIpAddressPoolId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS is basic level. If you set the value to `AntiDDoS_Enhanced`, High capacity Anti-DDoS Origin is enabled.
     */
    public readonly securityProtectionTypes!: pulumi.Output<string[] | undefined>;
    /**
     * The status of the EIP. Valid values:  `Associating`: The EIP is being associated. `Unassociating`: The EIP is being disassociated. `InUse`: The EIP is allocated. `Available`:The EIP is available.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a EipAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EipAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EipAddressArgs | EipAddressState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EipAddressState | undefined;
            resourceInputs["activityId"] = state ? state.activityId : undefined;
            resourceInputs["addressName"] = state ? state.addressName : undefined;
            resourceInputs["autoPay"] = state ? state.autoPay : undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["highDefinitionMonitorLogStatus"] = state ? state.highDefinitionMonitorLogStatus : undefined;
            resourceInputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            resourceInputs["internetChargeType"] = state ? state.internetChargeType : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["isp"] = state ? state.isp : undefined;
            resourceInputs["logProject"] = state ? state.logProject : undefined;
            resourceInputs["logStore"] = state ? state.logStore : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["netmode"] = state ? state.netmode : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["publicIpAddressPoolId"] = state ? state.publicIpAddressPoolId : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["securityProtectionTypes"] = state ? state.securityProtectionTypes : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EipAddressArgs | undefined;
            resourceInputs["activityId"] = args ? args.activityId : undefined;
            resourceInputs["addressName"] = args ? args.addressName : undefined;
            resourceInputs["autoPay"] = args ? args.autoPay : undefined;
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["highDefinitionMonitorLogStatus"] = args ? args.highDefinitionMonitorLogStatus : undefined;
            resourceInputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            resourceInputs["internetChargeType"] = args ? args.internetChargeType : undefined;
            resourceInputs["isp"] = args ? args.isp : undefined;
            resourceInputs["logProject"] = args ? args.logProject : undefined;
            resourceInputs["logStore"] = args ? args.logStore : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["netmode"] = args ? args.netmode : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["publicIpAddressPoolId"] = args ? args.publicIpAddressPoolId : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["securityProtectionTypes"] = args ? args.securityProtectionTypes : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["ipAddress"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EipAddress.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EipAddress resources.
 */
export interface EipAddressState {
    /**
     * The activity id.
     */
    activityId?: pulumi.Input<string>;
    /**
     * The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
     */
    addressName?: pulumi.Input<string>;
    /**
     * Whether to pay automatically. Valid values: `true` and `false`. Default value: `true`. When `autoPay` is `true`, The order will be automatically paid. When `autoPay` is `false`, The order needs to go to the order center to complete the payment. **NOTE:** When `paymentType` is `Subscription`, this parameter is valid.
     */
    autoPay?: pulumi.Input<boolean>;
    /**
     * The maximum bandwidth of the EIP. Valid values: `1` to `200`. Unit: Mbit/s. Default value: `5`.
     */
    bandwidth?: pulumi.Input<string>;
    /**
     * Whether enable the deletion protection or not. Default value: `false`.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The description of the EIP.
     */
    description?: pulumi.Input<string>;
    /**
     * The status of the EIP. configuring high precision second-by-second monitoring for EIP. Valid values: `ON` and `OFF`.
     */
    highDefinitionMonitorLogStatus?: pulumi.Input<string>;
    /**
     * Field `instanceChargeType` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `paymentType` instead.
     *
     * @deprecated Field 'instance_charge_type' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The metering method of the EIP. 
     * Valid values: `PayByDominantTraffic`, `PayByBandwidth` and `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** It must be set to "PayByBandwidth" when `paymentType` is "Subscription".
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * The address of the EIP.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * The line type. You can set this parameter only when you create a `PayAsYouGo` EIP. Valid values:
     */
    isp?: pulumi.Input<string>;
    /**
     * The Name of the logging service LogProject. Current parameter is required when configuring high precision second-by-second monitoring for EIP.
     */
    logProject?: pulumi.Input<string>;
    /**
     * The Name of the logging service LogStore. Current parameter is required when configuring high precision second-by-second monitoring for EIP.
     */
    logStore?: pulumi.Input<string>;
    /**
     * Field `name` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `addressName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'address_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of the network. Valid value is `public` (Internet).
     */
    netmode?: pulumi.Input<string>;
    /**
     * The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
     */
    period?: pulumi.Input<number>;
    /**
     * The ID of the IP address pool. The EIP is allocated from the IP address pool. **NOTE:** The feature is available only to users whose accounts are included in the whitelist. If you want to use the feature,[submit a ticket](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/429100).
     */
    publicIpAddressPoolId?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS is basic level. If you set the value to `AntiDDoS_Enhanced`, High capacity Anti-DDoS Origin is enabled.
     */
    securityProtectionTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of the EIP. Valid values:  `Associating`: The EIP is being associated. `Unassociating`: The EIP is being disassociated. `InUse`: The EIP is allocated. `Available`:The EIP is available.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a EipAddress resource.
 */
export interface EipAddressArgs {
    /**
     * The activity id.
     */
    activityId?: pulumi.Input<string>;
    /**
     * The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
     */
    addressName?: pulumi.Input<string>;
    /**
     * Whether to pay automatically. Valid values: `true` and `false`. Default value: `true`. When `autoPay` is `true`, The order will be automatically paid. When `autoPay` is `false`, The order needs to go to the order center to complete the payment. **NOTE:** When `paymentType` is `Subscription`, this parameter is valid.
     */
    autoPay?: pulumi.Input<boolean>;
    /**
     * The maximum bandwidth of the EIP. Valid values: `1` to `200`. Unit: Mbit/s. Default value: `5`.
     */
    bandwidth?: pulumi.Input<string>;
    /**
     * Whether enable the deletion protection or not. Default value: `false`.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The description of the EIP.
     */
    description?: pulumi.Input<string>;
    /**
     * The status of the EIP. configuring high precision second-by-second monitoring for EIP. Valid values: `ON` and `OFF`.
     */
    highDefinitionMonitorLogStatus?: pulumi.Input<string>;
    /**
     * Field `instanceChargeType` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `paymentType` instead.
     *
     * @deprecated Field 'instance_charge_type' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The metering method of the EIP. 
     * Valid values: `PayByDominantTraffic`, `PayByBandwidth` and `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** It must be set to "PayByBandwidth" when `paymentType` is "Subscription".
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * The line type. You can set this parameter only when you create a `PayAsYouGo` EIP. Valid values:
     */
    isp?: pulumi.Input<string>;
    /**
     * The Name of the logging service LogProject. Current parameter is required when configuring high precision second-by-second monitoring for EIP.
     */
    logProject?: pulumi.Input<string>;
    /**
     * The Name of the logging service LogStore. Current parameter is required when configuring high precision second-by-second monitoring for EIP.
     */
    logStore?: pulumi.Input<string>;
    /**
     * Field `name` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `addressName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'address_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of the network. Valid value is `public` (Internet).
     */
    netmode?: pulumi.Input<string>;
    /**
     * The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
     */
    period?: pulumi.Input<number>;
    /**
     * The ID of the IP address pool. The EIP is allocated from the IP address pool. **NOTE:** The feature is available only to users whose accounts are included in the whitelist. If you want to use the feature,[submit a ticket](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/429100).
     */
    publicIpAddressPoolId?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS is basic level. If you set the value to `AntiDDoS_Enhanced`, High capacity Anti-DDoS Origin is enabled.
     */
    securityProtectionTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
