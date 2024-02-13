// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * EIP Address can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ecs/eipAddress:EipAddress example <id>
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
     * Special activity ID. This parameter is not required.
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
     * The time when the EIP was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Whether the delete protection function is turned on.
     * - **true**: enabled.
     * - **false**: not enabled.
     */
    public readonly deletionProtection!: pulumi.Output<boolean>;
    /**
     * The description of the EIP.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the second-level monitoring is enabled for the EIP.
     * - **OFF**: not enabled.
     * - **ON**: enabled.
     */
    public readonly highDefinitionMonitorLogStatus!: pulumi.Output<string>;
    /**
     * . Field 'instance_charge_type' has been deprecated from provider version 1.126.0. New field 'payment_type' instead.
     *
     * @deprecated Field 'instance_charge_type' has been deprecated since provider version 1.126.0. New field 'payment_type' instead.
     */
    public readonly instanceChargeType!: pulumi.Output<string>;
    /**
     * Renewal Payment type.
     * - **PayByBandwidth**: billed by fixed bandwidth.
     * - **PayByTraffic**: Billing by traffic.
     */
    public readonly internetChargeType!: pulumi.Output<string>;
    /**
     * The IP address of the EIP.
     */
    public readonly ipAddress!: pulumi.Output<string>;
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
     * . Field 'name' has been deprecated from provider version 1.126.0. New field 'address_name' instead.
     *
     * @deprecated Field 'name' has been deprecated since provider version 1.126.0. New field 'address_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of the network. Valid value is `public` (Internet).
     */
    public readonly netmode!: pulumi.Output<string>;
    /**
     * The billing method of the EIP. Valid values:  `Subscription`, `PayAsYouGo`.
     */
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * When the PricingCycle is set to Month, the Period value ranges from 1 to 9.  When the PricingCycle is set to Year, the Period range is 1 to 5.  If the value of the InstanceChargeType parameter is PrePaid, this parameter is required. If the value of the InstanceChargeType parameter is PostPaid, this parameter is not filled in.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Value: Month (default): Pay monthly. Year: Pay per Year. This parameter is required when the value of the InstanceChargeType parameter is Subscription(PrePaid). This parameter is optional when the value of the InstanceChargeType parameter is PayAsYouGo(PostPaid).
     */
    public readonly pricingCycle!: pulumi.Output<string | undefined>;
    /**
     * The ID of the IP address pool to which the EIP belongs.
     */
    public readonly publicIpAddressPoolId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * Security protection level.
     * - When the return is empty, the basic DDoS protection is specified.
     * - When **antidos_enhanced** is returned, it indicates DDoS protection (enhanced version).
     */
    public readonly securityProtectionTypes!: pulumi.Output<string[] | undefined>;
    /**
     * The status of the EIP.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The tag of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The zone of the EIP.  This parameter is returned only for whitelist users that are visible to the zone.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    public readonly zone!: pulumi.Output<string>;

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
            resourceInputs["createTime"] = state ? state.createTime : undefined;
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
            resourceInputs["pricingCycle"] = state ? state.pricingCycle : undefined;
            resourceInputs["publicIpAddressPoolId"] = state ? state.publicIpAddressPoolId : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["securityProtectionTypes"] = state ? state.securityProtectionTypes : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
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
            resourceInputs["ipAddress"] = args ? args.ipAddress : undefined;
            resourceInputs["isp"] = args ? args.isp : undefined;
            resourceInputs["logProject"] = args ? args.logProject : undefined;
            resourceInputs["logStore"] = args ? args.logStore : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["netmode"] = args ? args.netmode : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["pricingCycle"] = args ? args.pricingCycle : undefined;
            resourceInputs["publicIpAddressPoolId"] = args ? args.publicIpAddressPoolId : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["securityProtectionTypes"] = args ? args.securityProtectionTypes : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
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
     * Special activity ID. This parameter is not required.
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
     * The time when the EIP was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Whether the delete protection function is turned on.
     * - **true**: enabled.
     * - **false**: not enabled.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The description of the EIP.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the second-level monitoring is enabled for the EIP.
     * - **OFF**: not enabled.
     * - **ON**: enabled.
     */
    highDefinitionMonitorLogStatus?: pulumi.Input<string>;
    /**
     * . Field 'instance_charge_type' has been deprecated from provider version 1.126.0. New field 'payment_type' instead.
     *
     * @deprecated Field 'instance_charge_type' has been deprecated since provider version 1.126.0. New field 'payment_type' instead.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * Renewal Payment type.
     * - **PayByBandwidth**: billed by fixed bandwidth.
     * - **PayByTraffic**: Billing by traffic.
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * The IP address of the EIP.
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
     * . Field 'name' has been deprecated from provider version 1.126.0. New field 'address_name' instead.
     *
     * @deprecated Field 'name' has been deprecated since provider version 1.126.0. New field 'address_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of the network. Valid value is `public` (Internet).
     */
    netmode?: pulumi.Input<string>;
    /**
     * The billing method of the EIP. Valid values:  `Subscription`, `PayAsYouGo`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * When the PricingCycle is set to Month, the Period value ranges from 1 to 9.  When the PricingCycle is set to Year, the Period range is 1 to 5.  If the value of the InstanceChargeType parameter is PrePaid, this parameter is required. If the value of the InstanceChargeType parameter is PostPaid, this parameter is not filled in.
     */
    period?: pulumi.Input<number>;
    /**
     * Value: Month (default): Pay monthly. Year: Pay per Year. This parameter is required when the value of the InstanceChargeType parameter is Subscription(PrePaid). This parameter is optional when the value of the InstanceChargeType parameter is PayAsYouGo(PostPaid).
     */
    pricingCycle?: pulumi.Input<string>;
    /**
     * The ID of the IP address pool to which the EIP belongs.
     */
    publicIpAddressPoolId?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Security protection level.
     * - When the return is empty, the basic DDoS protection is specified.
     * - When **antidos_enhanced** is returned, it indicates DDoS protection (enhanced version).
     */
    securityProtectionTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of the EIP.
     */
    status?: pulumi.Input<string>;
    /**
     * The tag of the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The zone of the EIP.  This parameter is returned only for whitelist users that are visible to the zone.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EipAddress resource.
 */
export interface EipAddressArgs {
    /**
     * Special activity ID. This parameter is not required.
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
     * Whether the delete protection function is turned on.
     * - **true**: enabled.
     * - **false**: not enabled.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The description of the EIP.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the second-level monitoring is enabled for the EIP.
     * - **OFF**: not enabled.
     * - **ON**: enabled.
     */
    highDefinitionMonitorLogStatus?: pulumi.Input<string>;
    /**
     * . Field 'instance_charge_type' has been deprecated from provider version 1.126.0. New field 'payment_type' instead.
     *
     * @deprecated Field 'instance_charge_type' has been deprecated since provider version 1.126.0. New field 'payment_type' instead.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * Renewal Payment type.
     * - **PayByBandwidth**: billed by fixed bandwidth.
     * - **PayByTraffic**: Billing by traffic.
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * The IP address of the EIP.
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
     * . Field 'name' has been deprecated from provider version 1.126.0. New field 'address_name' instead.
     *
     * @deprecated Field 'name' has been deprecated since provider version 1.126.0. New field 'address_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of the network. Valid value is `public` (Internet).
     */
    netmode?: pulumi.Input<string>;
    /**
     * The billing method of the EIP. Valid values:  `Subscription`, `PayAsYouGo`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * When the PricingCycle is set to Month, the Period value ranges from 1 to 9.  When the PricingCycle is set to Year, the Period range is 1 to 5.  If the value of the InstanceChargeType parameter is PrePaid, this parameter is required. If the value of the InstanceChargeType parameter is PostPaid, this parameter is not filled in.
     */
    period?: pulumi.Input<number>;
    /**
     * Value: Month (default): Pay monthly. Year: Pay per Year. This parameter is required when the value of the InstanceChargeType parameter is Subscription(PrePaid). This parameter is optional when the value of the InstanceChargeType parameter is PayAsYouGo(PostPaid).
     */
    pricingCycle?: pulumi.Input<string>;
    /**
     * The ID of the IP address pool to which the EIP belongs.
     */
    publicIpAddressPoolId?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Security protection level.
     * - When the return is empty, the basic DDoS protection is specified.
     * - When **antidos_enhanced** is returned, it indicates DDoS protection (enhanced version).
     */
    securityProtectionTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tag of the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The zone of the EIP.  This parameter is returned only for whitelist users that are visible to the zone.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    zone?: pulumi.Input<string>;
}
