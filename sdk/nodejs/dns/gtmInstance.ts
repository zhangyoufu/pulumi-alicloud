// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Alidns Gtm Instance resource.
 *
 * For information about Alidns Gtm Instance and how to use it, see [What is Gtm Instance](https://www.alibabacloud.com/help/en/doc-detail/204852.html).
 *
 * > **NOTE:** Available since v1.151.0.
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
 * const domainName = config.get("domainName") || "alicloud-provider.com";
 * const defaultResourceGroups = alicloud.resourcemanager.getResourceGroups({});
 * const defaultAlarmContactGroup = new alicloud.cms.AlarmContactGroup("defaultAlarmContactGroup", {alarmContactGroupName: "tf_example"});
 * const defaultGtmInstance = new alicloud.dns.GtmInstance("defaultGtmInstance", {
 *     instanceName: "tf_example",
 *     paymentType: "Subscription",
 *     period: 1,
 *     renewalStatus: "ManualRenewal",
 *     packageEdition: "standard",
 *     healthCheckTaskCount: 100,
 *     smsNotificationCount: 1000,
 *     publicCnameMode: "SYSTEM_ASSIGN",
 *     ttl: 60,
 *     cnameType: "PUBLIC",
 *     resourceGroupId: defaultResourceGroups.then(defaultResourceGroups => defaultResourceGroups.groups?.[0]?.id),
 *     alertGroups: [defaultAlarmContactGroup.alarmContactGroupName],
 *     publicUserDomainName: domainName,
 *     alertConfigs: [{
 *         smsNotice: true,
 *         noticeType: "ADDR_ALERT",
 *         emailNotice: true,
 *         dingtalkNotice: true,
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Alidns Gtm Instance can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:dns/gtmInstance:GtmInstance example <id>
 * ```
 */
export class GtmInstance extends pulumi.CustomResource {
    /**
     * Get an existing GtmInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GtmInstanceState, opts?: pulumi.CustomResourceOptions): GtmInstance {
        return new GtmInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dns/gtmInstance:GtmInstance';

    /**
     * Returns true if the given object is an instance of GtmInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GtmInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GtmInstance.__pulumiType;
    }

    /**
     * The alert notification methods. See `alertConfig` below for details.
     */
    public readonly alertConfigs!: pulumi.Output<outputs.dns.GtmInstanceAlertConfig[] | undefined>;
    /**
     * The alert group.
     */
    public readonly alertGroups!: pulumi.Output<string[] | undefined>;
    /**
     * The access type of the CNAME domain name. Valid value: `PUBLIC`.
     */
    public readonly cnameType!: pulumi.Output<string>;
    /**
     * The force update.
     */
    public readonly forceUpdate!: pulumi.Output<boolean | undefined>;
    /**
     * The quota of detection tasks.
     */
    public readonly healthCheckTaskCount!: pulumi.Output<number>;
    /**
     * The name of the instance.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * The lang.
     */
    public readonly lang!: pulumi.Output<string | undefined>;
    /**
     * Paid package version. Valid values: `ultimate`, `standard`.
     */
    public readonly packageEdition!: pulumi.Output<string>;
    /**
     * The Payment Type of the resource. Valid value: `Subscription`.
     */
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
     */
    public readonly period!: pulumi.Output<number>;
    /**
     * The Public Network domain name access method. Valid values: `CUSTOM`, `SYSTEM_ASSIGN`.
     */
    public readonly publicCnameMode!: pulumi.Output<string>;
    /**
     * The CNAME access domain name.
     */
    public readonly publicRr!: pulumi.Output<string>;
    /**
     * The website domain name that the user uses on the Internet.
     */
    public readonly publicUserDomainName!: pulumi.Output<string>;
    /**
     * The domain name that is used to access GTM over the Internet.
     */
    public readonly publicZoneName!: pulumi.Output<string>;
    /**
     * Automatic renewal period, the unit is month. When setting `renewalStatus` to AutoRenewal, it must be set.
     */
    public readonly renewPeriod!: pulumi.Output<number | undefined>;
    /**
     * Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`.
     */
    public readonly renewalStatus!: pulumi.Output<string>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * The quota of SMS notifications.
     */
    public readonly smsNotificationCount!: pulumi.Output<number>;
    /**
     * The type of the access policy. Valid values: `GEO`, `LATENCY`.
     */
    public readonly strategyMode!: pulumi.Output<string>;
    /**
     * The global time to live. Valid values: `60`, `120`, `300`, `600`. Unit: second.
     */
    public readonly ttl!: pulumi.Output<number | undefined>;

    /**
     * Create a GtmInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GtmInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GtmInstanceArgs | GtmInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GtmInstanceState | undefined;
            resourceInputs["alertConfigs"] = state ? state.alertConfigs : undefined;
            resourceInputs["alertGroups"] = state ? state.alertGroups : undefined;
            resourceInputs["cnameType"] = state ? state.cnameType : undefined;
            resourceInputs["forceUpdate"] = state ? state.forceUpdate : undefined;
            resourceInputs["healthCheckTaskCount"] = state ? state.healthCheckTaskCount : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["lang"] = state ? state.lang : undefined;
            resourceInputs["packageEdition"] = state ? state.packageEdition : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["publicCnameMode"] = state ? state.publicCnameMode : undefined;
            resourceInputs["publicRr"] = state ? state.publicRr : undefined;
            resourceInputs["publicUserDomainName"] = state ? state.publicUserDomainName : undefined;
            resourceInputs["publicZoneName"] = state ? state.publicZoneName : undefined;
            resourceInputs["renewPeriod"] = state ? state.renewPeriod : undefined;
            resourceInputs["renewalStatus"] = state ? state.renewalStatus : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["smsNotificationCount"] = state ? state.smsNotificationCount : undefined;
            resourceInputs["strategyMode"] = state ? state.strategyMode : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
        } else {
            const args = argsOrState as GtmInstanceArgs | undefined;
            if ((!args || args.healthCheckTaskCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'healthCheckTaskCount'");
            }
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            if ((!args || args.packageEdition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'packageEdition'");
            }
            if ((!args || args.paymentType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paymentType'");
            }
            if ((!args || args.period === undefined) && !opts.urn) {
                throw new Error("Missing required property 'period'");
            }
            if ((!args || args.smsNotificationCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'smsNotificationCount'");
            }
            resourceInputs["alertConfigs"] = args ? args.alertConfigs : undefined;
            resourceInputs["alertGroups"] = args ? args.alertGroups : undefined;
            resourceInputs["cnameType"] = args ? args.cnameType : undefined;
            resourceInputs["forceUpdate"] = args ? args.forceUpdate : undefined;
            resourceInputs["healthCheckTaskCount"] = args ? args.healthCheckTaskCount : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["lang"] = args ? args.lang : undefined;
            resourceInputs["packageEdition"] = args ? args.packageEdition : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["publicCnameMode"] = args ? args.publicCnameMode : undefined;
            resourceInputs["publicRr"] = args ? args.publicRr : undefined;
            resourceInputs["publicUserDomainName"] = args ? args.publicUserDomainName : undefined;
            resourceInputs["publicZoneName"] = args ? args.publicZoneName : undefined;
            resourceInputs["renewPeriod"] = args ? args.renewPeriod : undefined;
            resourceInputs["renewalStatus"] = args ? args.renewalStatus : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["smsNotificationCount"] = args ? args.smsNotificationCount : undefined;
            resourceInputs["strategyMode"] = args ? args.strategyMode : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GtmInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GtmInstance resources.
 */
export interface GtmInstanceState {
    /**
     * The alert notification methods. See `alertConfig` below for details.
     */
    alertConfigs?: pulumi.Input<pulumi.Input<inputs.dns.GtmInstanceAlertConfig>[]>;
    /**
     * The alert group.
     */
    alertGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The access type of the CNAME domain name. Valid value: `PUBLIC`.
     */
    cnameType?: pulumi.Input<string>;
    /**
     * The force update.
     */
    forceUpdate?: pulumi.Input<boolean>;
    /**
     * The quota of detection tasks.
     */
    healthCheckTaskCount?: pulumi.Input<number>;
    /**
     * The name of the instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The lang.
     */
    lang?: pulumi.Input<string>;
    /**
     * Paid package version. Valid values: `ultimate`, `standard`.
     */
    packageEdition?: pulumi.Input<string>;
    /**
     * The Payment Type of the resource. Valid value: `Subscription`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
     */
    period?: pulumi.Input<number>;
    /**
     * The Public Network domain name access method. Valid values: `CUSTOM`, `SYSTEM_ASSIGN`.
     */
    publicCnameMode?: pulumi.Input<string>;
    /**
     * The CNAME access domain name.
     */
    publicRr?: pulumi.Input<string>;
    /**
     * The website domain name that the user uses on the Internet.
     */
    publicUserDomainName?: pulumi.Input<string>;
    /**
     * The domain name that is used to access GTM over the Internet.
     */
    publicZoneName?: pulumi.Input<string>;
    /**
     * Automatic renewal period, the unit is month. When setting `renewalStatus` to AutoRenewal, it must be set.
     */
    renewPeriod?: pulumi.Input<number>;
    /**
     * Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`.
     */
    renewalStatus?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The quota of SMS notifications.
     */
    smsNotificationCount?: pulumi.Input<number>;
    /**
     * The type of the access policy. Valid values: `GEO`, `LATENCY`.
     */
    strategyMode?: pulumi.Input<string>;
    /**
     * The global time to live. Valid values: `60`, `120`, `300`, `600`. Unit: second.
     */
    ttl?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a GtmInstance resource.
 */
export interface GtmInstanceArgs {
    /**
     * The alert notification methods. See `alertConfig` below for details.
     */
    alertConfigs?: pulumi.Input<pulumi.Input<inputs.dns.GtmInstanceAlertConfig>[]>;
    /**
     * The alert group.
     */
    alertGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The access type of the CNAME domain name. Valid value: `PUBLIC`.
     */
    cnameType?: pulumi.Input<string>;
    /**
     * The force update.
     */
    forceUpdate?: pulumi.Input<boolean>;
    /**
     * The quota of detection tasks.
     */
    healthCheckTaskCount: pulumi.Input<number>;
    /**
     * The name of the instance.
     */
    instanceName: pulumi.Input<string>;
    /**
     * The lang.
     */
    lang?: pulumi.Input<string>;
    /**
     * Paid package version. Valid values: `ultimate`, `standard`.
     */
    packageEdition: pulumi.Input<string>;
    /**
     * The Payment Type of the resource. Valid value: `Subscription`.
     */
    paymentType: pulumi.Input<string>;
    /**
     * Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
     */
    period: pulumi.Input<number>;
    /**
     * The Public Network domain name access method. Valid values: `CUSTOM`, `SYSTEM_ASSIGN`.
     */
    publicCnameMode?: pulumi.Input<string>;
    /**
     * The CNAME access domain name.
     */
    publicRr?: pulumi.Input<string>;
    /**
     * The website domain name that the user uses on the Internet.
     */
    publicUserDomainName?: pulumi.Input<string>;
    /**
     * The domain name that is used to access GTM over the Internet.
     */
    publicZoneName?: pulumi.Input<string>;
    /**
     * Automatic renewal period, the unit is month. When setting `renewalStatus` to AutoRenewal, it must be set.
     */
    renewPeriod?: pulumi.Input<number>;
    /**
     * Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`.
     */
    renewalStatus?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The quota of SMS notifications.
     */
    smsNotificationCount: pulumi.Input<number>;
    /**
     * The type of the access policy. Valid values: `GEO`, `LATENCY`.
     */
    strategyMode?: pulumi.Input<string>;
    /**
     * The global time to live. Valid values: `60`, `120`, `300`, `600`. Unit: second.
     */
    ttl?: pulumi.Input<number>;
}
