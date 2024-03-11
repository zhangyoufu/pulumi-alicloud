// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Config Remediation resource.
 *
 * For information about Config Remediation and how to use it, see [What is Remediation](https://www.alibabacloud.com/help/en/cloud-config/latest/api-config-2020-09-07-createremediation).
 *
 * > **NOTE:** Available since v1.204.0.
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
 * const name = config.get("name") || "tf-example-oss";
 * const defaultRegions = alicloud.getRegions({
 *     current: true,
 * });
 * const defaultBucket = new alicloud.oss.Bucket("defaultBucket", {
 *     bucket: name,
 *     acl: "public-read",
 *     tags: {
 *         For: "example",
 *     },
 * });
 * const defaultRule = new alicloud.cfg.Rule("defaultRule", {
 *     description: "If the ACL policy of the OSS bucket denies read access from the Internet, the configuration is considered compliant.",
 *     sourceOwner: "ALIYUN",
 *     sourceIdentifier: "oss-bucket-public-read-prohibited",
 *     riskLevel: 1,
 *     tagKeyScope: "For",
 *     tagValueScope: "example",
 *     regionIdsScope: defaultRegions.then(defaultRegions => defaultRegions.regions?.[0]?.id),
 *     configRuleTriggerTypes: "ConfigurationItemChangeNotification",
 *     resourceTypesScopes: ["ACS::OSS::Bucket"],
 *     ruleName: "oss-bucket-public-read-prohibited",
 * });
 * const defaultRemediation = new alicloud.cfg.Remediation("defaultRemediation", {
 *     configRuleId: defaultRule.configRuleId,
 *     remediationTemplateId: "ACS-OSS-PutBucketAcl",
 *     remediationSourceType: "ALIYUN",
 *     invokeType: "MANUAL_EXECUTION",
 *     params: pulumi.all([defaultBucket.bucket, defaultRegions]).apply(([bucket, defaultRegions]) => `{"bucketName": "${bucket}", "regionId": "${defaultRegions.regions?.[0]?.id}", "permissionName": "private"}`),
 *     remediationType: "OOS",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Config Remediation can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cfg/remediation:Remediation example <id>
 * ```
 */
export class Remediation extends pulumi.CustomResource {
    /**
     * Get an existing Remediation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RemediationState, opts?: pulumi.CustomResourceOptions): Remediation {
        return new Remediation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cfg/remediation:Remediation';

    /**
     * Returns true if the given object is an instance of Remediation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Remediation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Remediation.__pulumiType;
    }

    /**
     * Rule ID.
     */
    public readonly configRuleId!: pulumi.Output<string>;
    /**
     * Execution type, valid values: `Manual`, `Automatic`.
     */
    public readonly invokeType!: pulumi.Output<string>;
    /**
     * Remediation parameter.
     */
    public readonly params!: pulumi.Output<string>;
    /**
     * Remediation ID.
     */
    public /*out*/ readonly remediationId!: pulumi.Output<string>;
    /**
     * Remediation resource type, valid values: `ALIYUN` , `CUSTOMER`.
     */
    public readonly remediationSourceType!: pulumi.Output<string>;
    /**
     * Remediation template ID.
     */
    public readonly remediationTemplateId!: pulumi.Output<string>;
    /**
     * Remediation type, valid values: `OOS`, `FC`.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    public readonly remediationType!: pulumi.Output<string>;

    /**
     * Create a Remediation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RemediationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RemediationArgs | RemediationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RemediationState | undefined;
            resourceInputs["configRuleId"] = state ? state.configRuleId : undefined;
            resourceInputs["invokeType"] = state ? state.invokeType : undefined;
            resourceInputs["params"] = state ? state.params : undefined;
            resourceInputs["remediationId"] = state ? state.remediationId : undefined;
            resourceInputs["remediationSourceType"] = state ? state.remediationSourceType : undefined;
            resourceInputs["remediationTemplateId"] = state ? state.remediationTemplateId : undefined;
            resourceInputs["remediationType"] = state ? state.remediationType : undefined;
        } else {
            const args = argsOrState as RemediationArgs | undefined;
            if ((!args || args.configRuleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configRuleId'");
            }
            if ((!args || args.invokeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'invokeType'");
            }
            if ((!args || args.params === undefined) && !opts.urn) {
                throw new Error("Missing required property 'params'");
            }
            if ((!args || args.remediationTemplateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remediationTemplateId'");
            }
            if ((!args || args.remediationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remediationType'");
            }
            resourceInputs["configRuleId"] = args ? args.configRuleId : undefined;
            resourceInputs["invokeType"] = args ? args.invokeType : undefined;
            resourceInputs["params"] = args ? args.params : undefined;
            resourceInputs["remediationSourceType"] = args ? args.remediationSourceType : undefined;
            resourceInputs["remediationTemplateId"] = args ? args.remediationTemplateId : undefined;
            resourceInputs["remediationType"] = args ? args.remediationType : undefined;
            resourceInputs["remediationId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Remediation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Remediation resources.
 */
export interface RemediationState {
    /**
     * Rule ID.
     */
    configRuleId?: pulumi.Input<string>;
    /**
     * Execution type, valid values: `Manual`, `Automatic`.
     */
    invokeType?: pulumi.Input<string>;
    /**
     * Remediation parameter.
     */
    params?: pulumi.Input<string>;
    /**
     * Remediation ID.
     */
    remediationId?: pulumi.Input<string>;
    /**
     * Remediation resource type, valid values: `ALIYUN` , `CUSTOMER`.
     */
    remediationSourceType?: pulumi.Input<string>;
    /**
     * Remediation template ID.
     */
    remediationTemplateId?: pulumi.Input<string>;
    /**
     * Remediation type, valid values: `OOS`, `FC`.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    remediationType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Remediation resource.
 */
export interface RemediationArgs {
    /**
     * Rule ID.
     */
    configRuleId: pulumi.Input<string>;
    /**
     * Execution type, valid values: `Manual`, `Automatic`.
     */
    invokeType: pulumi.Input<string>;
    /**
     * Remediation parameter.
     */
    params: pulumi.Input<string>;
    /**
     * Remediation resource type, valid values: `ALIYUN` , `CUSTOMER`.
     */
    remediationSourceType?: pulumi.Input<string>;
    /**
     * Remediation template ID.
     */
    remediationTemplateId: pulumi.Input<string>;
    /**
     * Remediation type, valid values: `OOS`, `FC`.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    remediationType: pulumi.Input<string>;
}
