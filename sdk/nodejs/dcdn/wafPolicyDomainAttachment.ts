// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a DCDN Waf Policy Domain Attachment resource.
 *
 * For information about DCDN Waf Policy Domain Attachment and how to use it, see [What is Waf Policy Domain Attachment](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-modifydcdnwafpolicydomains).
 *
 * > **NOTE:** Available since v1.186.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const domainName = config.get("domainName") || "tf-example.com";
 * const name = config.get("name") || "tf_example";
 * const _default = new random.RandomInteger("default", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const exampleDomain = new alicloud.dcdn.Domain("exampleDomain", {
 *     domainName: pulumi.interpolate`${domainName}-${_default.result}`,
 *     scope: "overseas",
 *     sources: [{
 *         content: "1.1.1.1",
 *         port: 80,
 *         priority: "20",
 *         type: "ipaddr",
 *         weight: "10",
 *     }],
 * });
 * const exampleWafDomain = new alicloud.dcdn.WafDomain("exampleWafDomain", {
 *     domainName: exampleDomain.domainName,
 *     clientIpTag: "X-Forwarded-For",
 * });
 * const exampleWafPolicy = new alicloud.dcdn.WafPolicy("exampleWafPolicy", {
 *     defenseScene: "waf_group",
 *     policyName: pulumi.interpolate`${name}_${_default.result}`,
 *     policyType: "custom",
 *     status: "on",
 * });
 * const exampleWafPolicyDomainAttachment = new alicloud.dcdn.WafPolicyDomainAttachment("exampleWafPolicyDomainAttachment", {
 *     domainName: exampleWafDomain.domainName,
 *     policyId: exampleWafPolicy.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * DCDN Waf Policy Domain Attachment can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:dcdn/wafPolicyDomainAttachment:WafPolicyDomainAttachment example policy_id:domain_name
 * ```
 */
export class WafPolicyDomainAttachment extends pulumi.CustomResource {
    /**
     * Get an existing WafPolicyDomainAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WafPolicyDomainAttachmentState, opts?: pulumi.CustomResourceOptions): WafPolicyDomainAttachment {
        return new WafPolicyDomainAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dcdn/wafPolicyDomainAttachment:WafPolicyDomainAttachment';

    /**
     * Returns true if the given object is an instance of WafPolicyDomainAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WafPolicyDomainAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WafPolicyDomainAttachment.__pulumiType;
    }

    /**
     * Access the accelerated domain name of the specified protection policy.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The protection policy ID. Only one input is supported.
     */
    public readonly policyId!: pulumi.Output<string>;

    /**
     * Create a WafPolicyDomainAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WafPolicyDomainAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WafPolicyDomainAttachmentArgs | WafPolicyDomainAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WafPolicyDomainAttachmentState | undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["policyId"] = state ? state.policyId : undefined;
        } else {
            const args = argsOrState as WafPolicyDomainAttachmentArgs | undefined;
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.policyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyId'");
            }
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["policyId"] = args ? args.policyId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WafPolicyDomainAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WafPolicyDomainAttachment resources.
 */
export interface WafPolicyDomainAttachmentState {
    /**
     * Access the accelerated domain name of the specified protection policy.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The protection policy ID. Only one input is supported.
     */
    policyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WafPolicyDomainAttachment resource.
 */
export interface WafPolicyDomainAttachmentArgs {
    /**
     * Access the accelerated domain name of the specified protection policy.
     */
    domainName: pulumi.Input<string>;
    /**
     * The protection policy ID. Only one input is supported.
     */
    policyId: pulumi.Input<string>;
}
