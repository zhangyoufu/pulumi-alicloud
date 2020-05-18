// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a WAF Instance resource to create instance in the Web Application Firewall.
 * 
 * For information about WAF and how to use it, see [What is Alibaba Cloud WAF](https://www.alibabacloud.com/help/doc-detail/28517.htm).
 * 
 * > **NOTE:** Available in 1.83.0+ .
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const defaultInstance = new alicloud.waf.Instance("default", {
 *     bigScreen: "0",
 *     exclusiveIpPackage: "1",
 *     extBandwidth: "50",
 *     extDomainPackage: "1",
 *     logStorage: "3",
 *     logTime: "180",
 *     packageCode: "version3",
 *     period: 1,
 *     prefessionalService: "false",
 *     resourceGroupId: "rs-abc12345",
 *     subscriptionType: "Subscription",
 *     wafLog: "false",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/waf_instance.html.markdown.
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:waf/instance:Instance';

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
     * Specify whether big screen is supported. Valid values: ["0", "1"]. "0" for false and "1" for true.
     */
    public readonly bigScreen!: pulumi.Output<string>;
    /**
     * Specify the number of exclusive WAF IP addresses.
     */
    public readonly exclusiveIpPackage!: pulumi.Output<string>;
    /**
     * The extra bandwidth. Unit: Mbit/s.
     */
    public readonly extBandwidth!: pulumi.Output<string>;
    /**
     * The number of extra domains.
     */
    public readonly extDomainPackage!: pulumi.Output<string>;
    /**
     * Log storage size. Unit: T. Valid values: [3, 5, 10, 20, 50].
     */
    public readonly logStorage!: pulumi.Output<string>;
    /**
     * Log storage period. Unit: day. Valid values: [180, 360].
     */
    public readonly logTime!: pulumi.Output<string>;
    /**
     * Type of configuration change. Valid value: Upgrade.
     */
    public readonly modifyType!: pulumi.Output<string | undefined>;
    /**
     * Subscription plan:
     * * China site customers can purchase the following versions of China Mainland region, valid values: ["version3", "version4", "version5"].
     * * China site customers can purchase the following versions of International region, valid values: ["versionProAsia", "versionBusinessAsia", "versionEnterpriseAsia"]
     * * International site customers can purchase the following versions of China Mainland region: ["versionProChina", "versionBusinessChina", "versionEnterpriseChina"]
     * * International site customers can purchase the following versions of International region: ["versionPro", "versionBusiness", "versionEnterprise"].
     */
    public readonly packageCode!: pulumi.Output<string>;
    /**
     * Service time of Web Application Firewall.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Specify whether professional service is supported. Valid values: ["true", "false"]
     */
    public readonly prefessionalService!: pulumi.Output<string>;
    /**
     * Renewal period of WAF service. Unit: month
     */
    public readonly renewPeriod!: pulumi.Output<number | undefined>;
    /**
     * Renewal status of WAF service. Valid values: 
     * * AutoRenewal: The service time of WAF is renewed automatically.
     * * ManualRenewal (default): The service time of WAF is renewed manually.Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: "On" and "Off". Default to "Off".
     */
    public readonly renewalStatus!: pulumi.Output<string | undefined>;
    /**
     * The resource group ID.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * The status of the instance.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * Subscription of WAF service. Valid values: ["Subscription", "PayAsYouGo"].
     */
    public readonly subscriptionType!: pulumi.Output<string>;
    /**
     * Specify whether Log service is supported. Valid values: ["true", "false"]                                           
     */
    public readonly wafLog!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceState | undefined;
            inputs["bigScreen"] = state ? state.bigScreen : undefined;
            inputs["exclusiveIpPackage"] = state ? state.exclusiveIpPackage : undefined;
            inputs["extBandwidth"] = state ? state.extBandwidth : undefined;
            inputs["extDomainPackage"] = state ? state.extDomainPackage : undefined;
            inputs["logStorage"] = state ? state.logStorage : undefined;
            inputs["logTime"] = state ? state.logTime : undefined;
            inputs["modifyType"] = state ? state.modifyType : undefined;
            inputs["packageCode"] = state ? state.packageCode : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["prefessionalService"] = state ? state.prefessionalService : undefined;
            inputs["renewPeriod"] = state ? state.renewPeriod : undefined;
            inputs["renewalStatus"] = state ? state.renewalStatus : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["subscriptionType"] = state ? state.subscriptionType : undefined;
            inputs["wafLog"] = state ? state.wafLog : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if (!args || args.bigScreen === undefined) {
                throw new Error("Missing required property 'bigScreen'");
            }
            if (!args || args.exclusiveIpPackage === undefined) {
                throw new Error("Missing required property 'exclusiveIpPackage'");
            }
            if (!args || args.extBandwidth === undefined) {
                throw new Error("Missing required property 'extBandwidth'");
            }
            if (!args || args.extDomainPackage === undefined) {
                throw new Error("Missing required property 'extDomainPackage'");
            }
            if (!args || args.logStorage === undefined) {
                throw new Error("Missing required property 'logStorage'");
            }
            if (!args || args.logTime === undefined) {
                throw new Error("Missing required property 'logTime'");
            }
            if (!args || args.packageCode === undefined) {
                throw new Error("Missing required property 'packageCode'");
            }
            if (!args || args.prefessionalService === undefined) {
                throw new Error("Missing required property 'prefessionalService'");
            }
            if (!args || args.subscriptionType === undefined) {
                throw new Error("Missing required property 'subscriptionType'");
            }
            if (!args || args.wafLog === undefined) {
                throw new Error("Missing required property 'wafLog'");
            }
            inputs["bigScreen"] = args ? args.bigScreen : undefined;
            inputs["exclusiveIpPackage"] = args ? args.exclusiveIpPackage : undefined;
            inputs["extBandwidth"] = args ? args.extBandwidth : undefined;
            inputs["extDomainPackage"] = args ? args.extDomainPackage : undefined;
            inputs["logStorage"] = args ? args.logStorage : undefined;
            inputs["logTime"] = args ? args.logTime : undefined;
            inputs["modifyType"] = args ? args.modifyType : undefined;
            inputs["packageCode"] = args ? args.packageCode : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["prefessionalService"] = args ? args.prefessionalService : undefined;
            inputs["renewPeriod"] = args ? args.renewPeriod : undefined;
            inputs["renewalStatus"] = args ? args.renewalStatus : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["subscriptionType"] = args ? args.subscriptionType : undefined;
            inputs["wafLog"] = args ? args.wafLog : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Specify whether big screen is supported. Valid values: ["0", "1"]. "0" for false and "1" for true.
     */
    readonly bigScreen?: pulumi.Input<string>;
    /**
     * Specify the number of exclusive WAF IP addresses.
     */
    readonly exclusiveIpPackage?: pulumi.Input<string>;
    /**
     * The extra bandwidth. Unit: Mbit/s.
     */
    readonly extBandwidth?: pulumi.Input<string>;
    /**
     * The number of extra domains.
     */
    readonly extDomainPackage?: pulumi.Input<string>;
    /**
     * Log storage size. Unit: T. Valid values: [3, 5, 10, 20, 50].
     */
    readonly logStorage?: pulumi.Input<string>;
    /**
     * Log storage period. Unit: day. Valid values: [180, 360].
     */
    readonly logTime?: pulumi.Input<string>;
    /**
     * Type of configuration change. Valid value: Upgrade.
     */
    readonly modifyType?: pulumi.Input<string>;
    /**
     * Subscription plan:
     * * China site customers can purchase the following versions of China Mainland region, valid values: ["version3", "version4", "version5"].
     * * China site customers can purchase the following versions of International region, valid values: ["versionProAsia", "versionBusinessAsia", "versionEnterpriseAsia"]
     * * International site customers can purchase the following versions of China Mainland region: ["versionProChina", "versionBusinessChina", "versionEnterpriseChina"]
     * * International site customers can purchase the following versions of International region: ["versionPro", "versionBusiness", "versionEnterprise"].
     */
    readonly packageCode?: pulumi.Input<string>;
    /**
     * Service time of Web Application Firewall.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * Specify whether professional service is supported. Valid values: ["true", "false"]
     */
    readonly prefessionalService?: pulumi.Input<string>;
    /**
     * Renewal period of WAF service. Unit: month
     */
    readonly renewPeriod?: pulumi.Input<number>;
    /**
     * Renewal status of WAF service. Valid values: 
     * * AutoRenewal: The service time of WAF is renewed automatically.
     * * ManualRenewal (default): The service time of WAF is renewed manually.Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: "On" and "Off". Default to "Off".
     */
    readonly renewalStatus?: pulumi.Input<string>;
    /**
     * The resource group ID.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The status of the instance.
     */
    readonly status?: pulumi.Input<number>;
    /**
     * Subscription of WAF service. Valid values: ["Subscription", "PayAsYouGo"].
     */
    readonly subscriptionType?: pulumi.Input<string>;
    /**
     * Specify whether Log service is supported. Valid values: ["true", "false"]                                           
     */
    readonly wafLog?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Specify whether big screen is supported. Valid values: ["0", "1"]. "0" for false and "1" for true.
     */
    readonly bigScreen: pulumi.Input<string>;
    /**
     * Specify the number of exclusive WAF IP addresses.
     */
    readonly exclusiveIpPackage: pulumi.Input<string>;
    /**
     * The extra bandwidth. Unit: Mbit/s.
     */
    readonly extBandwidth: pulumi.Input<string>;
    /**
     * The number of extra domains.
     */
    readonly extDomainPackage: pulumi.Input<string>;
    /**
     * Log storage size. Unit: T. Valid values: [3, 5, 10, 20, 50].
     */
    readonly logStorage: pulumi.Input<string>;
    /**
     * Log storage period. Unit: day. Valid values: [180, 360].
     */
    readonly logTime: pulumi.Input<string>;
    /**
     * Type of configuration change. Valid value: Upgrade.
     */
    readonly modifyType?: pulumi.Input<string>;
    /**
     * Subscription plan:
     * * China site customers can purchase the following versions of China Mainland region, valid values: ["version3", "version4", "version5"].
     * * China site customers can purchase the following versions of International region, valid values: ["versionProAsia", "versionBusinessAsia", "versionEnterpriseAsia"]
     * * International site customers can purchase the following versions of China Mainland region: ["versionProChina", "versionBusinessChina", "versionEnterpriseChina"]
     * * International site customers can purchase the following versions of International region: ["versionPro", "versionBusiness", "versionEnterprise"].
     */
    readonly packageCode: pulumi.Input<string>;
    /**
     * Service time of Web Application Firewall.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * Specify whether professional service is supported. Valid values: ["true", "false"]
     */
    readonly prefessionalService: pulumi.Input<string>;
    /**
     * Renewal period of WAF service. Unit: month
     */
    readonly renewPeriod?: pulumi.Input<number>;
    /**
     * Renewal status of WAF service. Valid values: 
     * * AutoRenewal: The service time of WAF is renewed automatically.
     * * ManualRenewal (default): The service time of WAF is renewed manually.Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: "On" and "Off". Default to "Off".
     */
    readonly renewalStatus?: pulumi.Input<string>;
    /**
     * The resource group ID.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * Subscription of WAF service. Valid values: ["Subscription", "PayAsYouGo"].
     */
    readonly subscriptionType: pulumi.Input<string>;
    /**
     * Specify whether Log service is supported. Valid values: ["true", "false"]                                           
     */
    readonly wafLog: pulumi.Input<string>;
}
