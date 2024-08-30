// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Firewall Instance resource.
 *
 * For information about Cloud Firewall Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/product/90174.htm).
 *
 * > **NOTE:** Available since v1.139.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * create a pay-as-you-go instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const payAsYouGo = new alicloud.cloudfirewall.Instance("PayAsYouGo", {paymentType: "PayAsYouGo"});
 * ```
 *
 * create a subscription instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const subscription = new alicloud.cloudfirewall.Instance("Subscription", {
 *     paymentType: "Subscription",
 *     spec: "premium_version",
 *     ipNumber: 20,
 *     bandWidth: 10,
 *     cfwLog: false,
 *     period: 1,
 * });
 * ```
 *
 * ## Import
 *
 * Cloud Firewall Instance can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cloudfirewall/instance:Instance example <id>
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cloudfirewall/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The number of multi account. It will be ignored when `cfwAccount = false`.
     */
    public readonly accountNumber!: pulumi.Output<number | undefined>;
    /**
     * Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
     */
    public readonly bandWidth!: pulumi.Output<number | undefined>;
    /**
     * Whether to use multi-account. Valid values: `true`, `false`.
     */
    public readonly cfwAccount!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to use log audit. Valid values: `true`, `false`.
     */
    public readonly cfwLog!: pulumi.Output<boolean | undefined>;
    /**
     * The log storage capacity. It will be ignored when `cfwLog = false`.
     */
    public readonly cfwLogStorage!: pulumi.Output<number | undefined>;
    /**
     * The creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The end time.
     */
    public /*out*/ readonly endTime!: pulumi.Output<string>;
    /**
     * The number of protected VPCs. It will be ignored when `spec = "premiumVersion"`. Valid values between 2 and 500.
     */
    public readonly fwVpcNumber!: pulumi.Output<number | undefined>;
    /**
     * The number of assets.
     */
    public readonly instanceCount!: pulumi.Output<number | undefined>;
    /**
     * The number of public IPs that can be protected. Valid values: 20 to 4000.
     */
    public readonly ipNumber!: pulumi.Output<number | undefined>;
    /**
     * The logistics.
     */
    public readonly logistics!: pulumi.Output<string | undefined>;
    /**
     * The type of modification. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
     */
    public readonly modifyType!: pulumi.Output<string | undefined>;
    /**
     * The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`. **NOTE:** From version 1.220.0, `paymentType` can be set to `PayAsYouGo`.
     */
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available since 1.204.1. If `paymentType` is set to `Subscription`, `period` is required. Otherwise, it will be ignored.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The release time.
     */
    public /*out*/ readonly releaseTime!: pulumi.Output<string>;
    /**
     * Automatic renewal period. Attribute `renewPeriod` has been deprecated since 1.209.1. Using `renewalDuration` instead.
     *
     * @deprecated Attribute 'renew_period' has been deprecated since 1.209.1. Using 'renewal_duration' instead.
     */
    public readonly renewPeriod!: pulumi.Output<number>;
    /**
     * Auto-Renewal Duration. It is required under the condition that `renewalStatus` is `AutoRenewal`. Valid values: `1`, `2`, `3`, `6`, `12`.
     * **NOTE:** `renewalDuration` takes effect only if `paymentType` is set to `Subscription`, and `renewalStatus` is set to `AutoRenewal`.
     */
    public readonly renewalDuration!: pulumi.Output<number>;
    /**
     * Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
     */
    public readonly renewalDurationUnit!: pulumi.Output<string | undefined>;
    /**
     * Whether to renew an instance automatically or not. Default to "ManualRenewal".
     * - `AutoRenewal`: Auto renewal.
     * - `ManualRenewal`: Manual renewal.
     * - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
     * **NOTE:** `renewalStatus` takes effect only if `paymentType` is set to `Subscription`.
     */
    public readonly renewalStatus!: pulumi.Output<string>;
    /**
     * Current version. Valid values: `premiumVersion`, `enterpriseVersion`,`ultimateVersion`.
     */
    public readonly spec!: pulumi.Output<string | undefined>;
    /**
     * The status of Instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["accountNumber"] = state ? state.accountNumber : undefined;
            resourceInputs["bandWidth"] = state ? state.bandWidth : undefined;
            resourceInputs["cfwAccount"] = state ? state.cfwAccount : undefined;
            resourceInputs["cfwLog"] = state ? state.cfwLog : undefined;
            resourceInputs["cfwLogStorage"] = state ? state.cfwLogStorage : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["fwVpcNumber"] = state ? state.fwVpcNumber : undefined;
            resourceInputs["instanceCount"] = state ? state.instanceCount : undefined;
            resourceInputs["ipNumber"] = state ? state.ipNumber : undefined;
            resourceInputs["logistics"] = state ? state.logistics : undefined;
            resourceInputs["modifyType"] = state ? state.modifyType : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["releaseTime"] = state ? state.releaseTime : undefined;
            resourceInputs["renewPeriod"] = state ? state.renewPeriod : undefined;
            resourceInputs["renewalDuration"] = state ? state.renewalDuration : undefined;
            resourceInputs["renewalDurationUnit"] = state ? state.renewalDurationUnit : undefined;
            resourceInputs["renewalStatus"] = state ? state.renewalStatus : undefined;
            resourceInputs["spec"] = state ? state.spec : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.paymentType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paymentType'");
            }
            resourceInputs["accountNumber"] = args ? args.accountNumber : undefined;
            resourceInputs["bandWidth"] = args ? args.bandWidth : undefined;
            resourceInputs["cfwAccount"] = args ? args.cfwAccount : undefined;
            resourceInputs["cfwLog"] = args ? args.cfwLog : undefined;
            resourceInputs["cfwLogStorage"] = args ? args.cfwLogStorage : undefined;
            resourceInputs["fwVpcNumber"] = args ? args.fwVpcNumber : undefined;
            resourceInputs["instanceCount"] = args ? args.instanceCount : undefined;
            resourceInputs["ipNumber"] = args ? args.ipNumber : undefined;
            resourceInputs["logistics"] = args ? args.logistics : undefined;
            resourceInputs["modifyType"] = args ? args.modifyType : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["renewPeriod"] = args ? args.renewPeriod : undefined;
            resourceInputs["renewalDuration"] = args ? args.renewalDuration : undefined;
            resourceInputs["renewalDurationUnit"] = args ? args.renewalDurationUnit : undefined;
            resourceInputs["renewalStatus"] = args ? args.renewalStatus : undefined;
            resourceInputs["spec"] = args ? args.spec : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["endTime"] = undefined /*out*/;
            resourceInputs["releaseTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The number of multi account. It will be ignored when `cfwAccount = false`.
     */
    accountNumber?: pulumi.Input<number>;
    /**
     * Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
     */
    bandWidth?: pulumi.Input<number>;
    /**
     * Whether to use multi-account. Valid values: `true`, `false`.
     */
    cfwAccount?: pulumi.Input<boolean>;
    /**
     * Whether to use log audit. Valid values: `true`, `false`.
     */
    cfwLog?: pulumi.Input<boolean>;
    /**
     * The log storage capacity. It will be ignored when `cfwLog = false`.
     */
    cfwLogStorage?: pulumi.Input<number>;
    /**
     * The creation time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The end time.
     */
    endTime?: pulumi.Input<string>;
    /**
     * The number of protected VPCs. It will be ignored when `spec = "premiumVersion"`. Valid values between 2 and 500.
     */
    fwVpcNumber?: pulumi.Input<number>;
    /**
     * The number of assets.
     */
    instanceCount?: pulumi.Input<number>;
    /**
     * The number of public IPs that can be protected. Valid values: 20 to 4000.
     */
    ipNumber?: pulumi.Input<number>;
    /**
     * The logistics.
     */
    logistics?: pulumi.Input<string>;
    /**
     * The type of modification. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
     */
    modifyType?: pulumi.Input<string>;
    /**
     * The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`. **NOTE:** From version 1.220.0, `paymentType` can be set to `PayAsYouGo`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available since 1.204.1. If `paymentType` is set to `Subscription`, `period` is required. Otherwise, it will be ignored.
     */
    period?: pulumi.Input<number>;
    /**
     * The release time.
     */
    releaseTime?: pulumi.Input<string>;
    /**
     * Automatic renewal period. Attribute `renewPeriod` has been deprecated since 1.209.1. Using `renewalDuration` instead.
     *
     * @deprecated Attribute 'renew_period' has been deprecated since 1.209.1. Using 'renewal_duration' instead.
     */
    renewPeriod?: pulumi.Input<number>;
    /**
     * Auto-Renewal Duration. It is required under the condition that `renewalStatus` is `AutoRenewal`. Valid values: `1`, `2`, `3`, `6`, `12`.
     * **NOTE:** `renewalDuration` takes effect only if `paymentType` is set to `Subscription`, and `renewalStatus` is set to `AutoRenewal`.
     */
    renewalDuration?: pulumi.Input<number>;
    /**
     * Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
     */
    renewalDurationUnit?: pulumi.Input<string>;
    /**
     * Whether to renew an instance automatically or not. Default to "ManualRenewal".
     * - `AutoRenewal`: Auto renewal.
     * - `ManualRenewal`: Manual renewal.
     * - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
     * **NOTE:** `renewalStatus` takes effect only if `paymentType` is set to `Subscription`.
     */
    renewalStatus?: pulumi.Input<string>;
    /**
     * Current version. Valid values: `premiumVersion`, `enterpriseVersion`,`ultimateVersion`.
     */
    spec?: pulumi.Input<string>;
    /**
     * The status of Instance.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The number of multi account. It will be ignored when `cfwAccount = false`.
     */
    accountNumber?: pulumi.Input<number>;
    /**
     * Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
     */
    bandWidth?: pulumi.Input<number>;
    /**
     * Whether to use multi-account. Valid values: `true`, `false`.
     */
    cfwAccount?: pulumi.Input<boolean>;
    /**
     * Whether to use log audit. Valid values: `true`, `false`.
     */
    cfwLog?: pulumi.Input<boolean>;
    /**
     * The log storage capacity. It will be ignored when `cfwLog = false`.
     */
    cfwLogStorage?: pulumi.Input<number>;
    /**
     * The number of protected VPCs. It will be ignored when `spec = "premiumVersion"`. Valid values between 2 and 500.
     */
    fwVpcNumber?: pulumi.Input<number>;
    /**
     * The number of assets.
     */
    instanceCount?: pulumi.Input<number>;
    /**
     * The number of public IPs that can be protected. Valid values: 20 to 4000.
     */
    ipNumber?: pulumi.Input<number>;
    /**
     * The logistics.
     */
    logistics?: pulumi.Input<string>;
    /**
     * The type of modification. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
     */
    modifyType?: pulumi.Input<string>;
    /**
     * The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`. **NOTE:** From version 1.220.0, `paymentType` can be set to `PayAsYouGo`.
     */
    paymentType: pulumi.Input<string>;
    /**
     * The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available since 1.204.1. If `paymentType` is set to `Subscription`, `period` is required. Otherwise, it will be ignored.
     */
    period?: pulumi.Input<number>;
    /**
     * Automatic renewal period. Attribute `renewPeriod` has been deprecated since 1.209.1. Using `renewalDuration` instead.
     *
     * @deprecated Attribute 'renew_period' has been deprecated since 1.209.1. Using 'renewal_duration' instead.
     */
    renewPeriod?: pulumi.Input<number>;
    /**
     * Auto-Renewal Duration. It is required under the condition that `renewalStatus` is `AutoRenewal`. Valid values: `1`, `2`, `3`, `6`, `12`.
     * **NOTE:** `renewalDuration` takes effect only if `paymentType` is set to `Subscription`, and `renewalStatus` is set to `AutoRenewal`.
     */
    renewalDuration?: pulumi.Input<number>;
    /**
     * Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
     */
    renewalDurationUnit?: pulumi.Input<string>;
    /**
     * Whether to renew an instance automatically or not. Default to "ManualRenewal".
     * - `AutoRenewal`: Auto renewal.
     * - `ManualRenewal`: Manual renewal.
     * - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
     * **NOTE:** `renewalStatus` takes effect only if `paymentType` is set to `Subscription`.
     */
    renewalStatus?: pulumi.Input<string>;
    /**
     * Current version. Valid values: `premiumVersion`, `enterpriseVersion`,`ultimateVersion`.
     */
    spec?: pulumi.Input<string>;
}
