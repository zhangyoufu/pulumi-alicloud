// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a KMS Instance resource.
 *
 * For information about KMS Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/zh/key-management-service/latest/kms-instance-management).
 *
 * > **NOTE:** Available since v1.210.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const default = alicloud.vpc.getNetworks({
 *     nameRegex: "^default-NODELETING$",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultGetSwitches = _default.then(_default => alicloud.vpc.getSwitches({
 *     vpcId: _default.ids?.[0],
 *     zoneId: "cn-hangzhou-h",
 * }));
 * const defaultInstance = new alicloud.kms.Instance("default", {
 *     productVersion: "3",
 *     vpcId: _default.then(_default => _default.ids?.[0]),
 *     zoneIds: [
 *         "cn-hangzhou-h",
 *         "cn-hangzhou-g",
 *     ],
 *     vswitchIds: [defaultGetSwitches.then(defaultGetSwitches => defaultGetSwitches.ids?.[0])],
 *     vpcNum: 1,
 *     keyNum: 1000,
 *     secretNum: 0,
 *     spec: 1000,
 * });
 * ```
 *
 * ## Import
 *
 * KMS Instance can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:kms/instance:Instance example <id>
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
    public static readonly __pulumiType = 'alicloud:kms/instance:Instance';

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
     * Aucillary VPCs used to access this KMS instance. See `bindVpcs` below.
     */
    public readonly bindVpcs!: pulumi.Output<outputs.kms.InstanceBindVpc[] | undefined>;
    /**
     * KMS instance certificate chain in PEM format.
     */
    public /*out*/ readonly caCertificateChainPem!: pulumi.Output<string>;
    /**
     * The creation time of the resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Whether to force deletion even without backup.
     */
    public readonly forceDeleteWithoutBackup!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly instanceName!: pulumi.Output<string>;
    /**
     * Maximum number of stored keys. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    public readonly keyNum!: pulumi.Output<number | undefined>;
    /**
     * Instance Audit Log Switch. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    public readonly log!: pulumi.Output<string>;
    /**
     * Instance log capacity. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    public readonly logStorage!: pulumi.Output<number>;
    /**
     * Payment type, valid values:  `Subscription`: Prepaid. `PayAsYouGo`: Postpaid, available since v1.223.2.
     */
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * Purchase cycle, in months. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * KMS Instance commodity type (software/hardware).
     */
    public readonly productVersion!: pulumi.Output<string | undefined>;
    /**
     * Automatic renewal period, in months. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    public readonly renewPeriod!: pulumi.Output<number | undefined>;
    /**
     * Renewal options. Valid values: `AutoRenewal`, `ManualRenewal`. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    public readonly renewStatus!: pulumi.Output<string | undefined>;
    /**
     * Maximum number of Secrets. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    public readonly secretNum!: pulumi.Output<number | undefined>;
    /**
     * The computation performance level of the KMS instance. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    public readonly spec!: pulumi.Output<number | undefined>;
    /**
     * Instance status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Instance VPC id.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The number of managed accesses. The maximum number of VPCs that can access this KMS instance. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    public readonly vpcNum!: pulumi.Output<number | undefined>;
    /**
     * Instance bind vswitches.
     */
    public readonly vswitchIds!: pulumi.Output<string[]>;
    /**
     * zone id.
     */
    public readonly zoneIds!: pulumi.Output<string[]>;

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
            resourceInputs["bindVpcs"] = state ? state.bindVpcs : undefined;
            resourceInputs["caCertificateChainPem"] = state ? state.caCertificateChainPem : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["forceDeleteWithoutBackup"] = state ? state.forceDeleteWithoutBackup : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["keyNum"] = state ? state.keyNum : undefined;
            resourceInputs["log"] = state ? state.log : undefined;
            resourceInputs["logStorage"] = state ? state.logStorage : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["productVersion"] = state ? state.productVersion : undefined;
            resourceInputs["renewPeriod"] = state ? state.renewPeriod : undefined;
            resourceInputs["renewStatus"] = state ? state.renewStatus : undefined;
            resourceInputs["secretNum"] = state ? state.secretNum : undefined;
            resourceInputs["spec"] = state ? state.spec : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vpcNum"] = state ? state.vpcNum : undefined;
            resourceInputs["vswitchIds"] = state ? state.vswitchIds : undefined;
            resourceInputs["zoneIds"] = state ? state.zoneIds : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            if ((!args || args.vswitchIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchIds'");
            }
            if ((!args || args.zoneIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneIds'");
            }
            resourceInputs["bindVpcs"] = args ? args.bindVpcs : undefined;
            resourceInputs["forceDeleteWithoutBackup"] = args ? args.forceDeleteWithoutBackup : undefined;
            resourceInputs["keyNum"] = args ? args.keyNum : undefined;
            resourceInputs["log"] = args ? args.log : undefined;
            resourceInputs["logStorage"] = args ? args.logStorage : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["productVersion"] = args ? args.productVersion : undefined;
            resourceInputs["renewPeriod"] = args ? args.renewPeriod : undefined;
            resourceInputs["renewStatus"] = args ? args.renewStatus : undefined;
            resourceInputs["secretNum"] = args ? args.secretNum : undefined;
            resourceInputs["spec"] = args ? args.spec : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vpcNum"] = args ? args.vpcNum : undefined;
            resourceInputs["vswitchIds"] = args ? args.vswitchIds : undefined;
            resourceInputs["zoneIds"] = args ? args.zoneIds : undefined;
            resourceInputs["caCertificateChainPem"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["instanceName"] = undefined /*out*/;
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
     * Aucillary VPCs used to access this KMS instance. See `bindVpcs` below.
     */
    bindVpcs?: pulumi.Input<pulumi.Input<inputs.kms.InstanceBindVpc>[]>;
    /**
     * KMS instance certificate chain in PEM format.
     */
    caCertificateChainPem?: pulumi.Input<string>;
    /**
     * The creation time of the resource.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Whether to force deletion even without backup.
     */
    forceDeleteWithoutBackup?: pulumi.Input<string>;
    /**
     * The name of the resource.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Maximum number of stored keys. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    keyNum?: pulumi.Input<number>;
    /**
     * Instance Audit Log Switch. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    log?: pulumi.Input<string>;
    /**
     * Instance log capacity. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    logStorage?: pulumi.Input<number>;
    /**
     * Payment type, valid values:  `Subscription`: Prepaid. `PayAsYouGo`: Postpaid, available since v1.223.2.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * Purchase cycle, in months. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    period?: pulumi.Input<number>;
    /**
     * KMS Instance commodity type (software/hardware).
     */
    productVersion?: pulumi.Input<string>;
    /**
     * Automatic renewal period, in months. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    renewPeriod?: pulumi.Input<number>;
    /**
     * Renewal options. Valid values: `AutoRenewal`, `ManualRenewal`. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    renewStatus?: pulumi.Input<string>;
    /**
     * Maximum number of Secrets. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    secretNum?: pulumi.Input<number>;
    /**
     * The computation performance level of the KMS instance. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    spec?: pulumi.Input<number>;
    /**
     * Instance status.
     */
    status?: pulumi.Input<string>;
    /**
     * Instance VPC id.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The number of managed accesses. The maximum number of VPCs that can access this KMS instance. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    vpcNum?: pulumi.Input<number>;
    /**
     * Instance bind vswitches.
     */
    vswitchIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * zone id.
     */
    zoneIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Aucillary VPCs used to access this KMS instance. See `bindVpcs` below.
     */
    bindVpcs?: pulumi.Input<pulumi.Input<inputs.kms.InstanceBindVpc>[]>;
    /**
     * Whether to force deletion even without backup.
     */
    forceDeleteWithoutBackup?: pulumi.Input<string>;
    /**
     * Maximum number of stored keys. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    keyNum?: pulumi.Input<number>;
    /**
     * Instance Audit Log Switch. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    log?: pulumi.Input<string>;
    /**
     * Instance log capacity. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    logStorage?: pulumi.Input<number>;
    /**
     * Payment type, valid values:  `Subscription`: Prepaid. `PayAsYouGo`: Postpaid, available since v1.223.2.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * Purchase cycle, in months. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    period?: pulumi.Input<number>;
    /**
     * KMS Instance commodity type (software/hardware).
     */
    productVersion?: pulumi.Input<string>;
    /**
     * Automatic renewal period, in months. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    renewPeriod?: pulumi.Input<number>;
    /**
     * Renewal options. Valid values: `AutoRenewal`, `ManualRenewal`. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    renewStatus?: pulumi.Input<string>;
    /**
     * Maximum number of Secrets. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    secretNum?: pulumi.Input<number>;
    /**
     * The computation performance level of the KMS instance. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    spec?: pulumi.Input<number>;
    /**
     * Instance VPC id.
     */
    vpcId: pulumi.Input<string>;
    /**
     * The number of managed accesses. The maximum number of VPCs that can access this KMS instance. The attribute is valid when the attribute `paymentType` is `Subscription`.
     */
    vpcNum?: pulumi.Input<number>;
    /**
     * Instance bind vswitches.
     */
    vswitchIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * zone id.
     */
    zoneIds: pulumi.Input<pulumi.Input<string>[]>;
}
