// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a WAFV3 Defense Template resource.
 *
 * For information about WAFV3 Defense Template and how to use it, see [What is Defense Template](https://www.alibabacloud.com/help/en/web-application-firewall/latest/api-waf-openapi-2021-10-01-createdefensetemplate).
 *
 * > **NOTE:** Available since v1.218.0.
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
 * const name = config.get("name") || "terraform-example";
 * const defaultInstances = alicloud.wafv3.getInstances({});
 * const defaultDefenseTemplate = new alicloud.wafv3.DefenseTemplate("defaultDefenseTemplate", {
 *     status: "1",
 *     instanceId: defaultInstances.then(defaultInstances => defaultInstances.ids?.[0]),
 *     defenseTemplateName: name,
 *     templateType: "user_custom",
 *     templateOrigin: "custom",
 *     defenseScene: "antiscan",
 *     resourceManagerResourceGroupId: "example",
 *     description: name,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * WAFV3 Defense Template can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:wafv3/defenseTemplate:DefenseTemplate example <instance_id>:<defense_template_id>
 * ```
 */
export class DefenseTemplate extends pulumi.CustomResource {
    /**
     * Get an existing DefenseTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefenseTemplateState, opts?: pulumi.CustomResourceOptions): DefenseTemplate {
        return new DefenseTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:wafv3/defenseTemplate:DefenseTemplate';

    /**
     * Returns true if the given object is an instance of DefenseTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefenseTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefenseTemplate.__pulumiType;
    }

    /**
     * The module to which the protection rule that you want to create belongs. Value:
     * - **waf_group**: the basic protection rule module.
     * - **antiscan**: the scan protection module.
     * - **ip_blacklist**: the IP address blacklist module.
     * - **custom_acl**: the custom rule module.
     * - **whitelist**: the whitelist module.
     * - **region_block**: the region blacklist module.
     * - **custom_response**: the custom response module.
     * - **cc**: the HTTP flood protection module.
     * - **tamperproof**: the website tamper-proofing module.
     * - **dlp**: the data leakage prevention module.
     */
    public readonly defenseScene!: pulumi.Output<string>;
    /**
     * Template ID.
     */
    public /*out*/ readonly defenseTemplateId!: pulumi.Output<number>;
    /**
     * The name of the protection rule template.
     */
    public readonly defenseTemplateName!: pulumi.Output<string>;
    /**
     * The description of the protection rule template. .
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Web Application Firewall (WAF) instance. .
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The ID of the Alibaba Cloud resource group. .
     */
    public readonly resourceManagerResourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * The status of the protection rule template. Valid values:
     * - **0**: disabled.
     * - **1**: enabled.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The origin of the protection rule template that you want to create. Set the value to **custom**. The value specifies that the protection rule template is a custom template. .
     */
    public readonly templateOrigin!: pulumi.Output<string>;
    /**
     * The type of the protection rule template. Valid values:
     * - **user_default:** default template.
     * - **user_custom:** custom template.
     */
    public readonly templateType!: pulumi.Output<string>;

    /**
     * Create a DefenseTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DefenseTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefenseTemplateArgs | DefenseTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DefenseTemplateState | undefined;
            resourceInputs["defenseScene"] = state ? state.defenseScene : undefined;
            resourceInputs["defenseTemplateId"] = state ? state.defenseTemplateId : undefined;
            resourceInputs["defenseTemplateName"] = state ? state.defenseTemplateName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["resourceManagerResourceGroupId"] = state ? state.resourceManagerResourceGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["templateOrigin"] = state ? state.templateOrigin : undefined;
            resourceInputs["templateType"] = state ? state.templateType : undefined;
        } else {
            const args = argsOrState as DefenseTemplateArgs | undefined;
            if ((!args || args.defenseScene === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defenseScene'");
            }
            if ((!args || args.defenseTemplateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defenseTemplateName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            if ((!args || args.templateOrigin === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateOrigin'");
            }
            if ((!args || args.templateType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateType'");
            }
            resourceInputs["defenseScene"] = args ? args.defenseScene : undefined;
            resourceInputs["defenseTemplateName"] = args ? args.defenseTemplateName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["resourceManagerResourceGroupId"] = args ? args.resourceManagerResourceGroupId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["templateOrigin"] = args ? args.templateOrigin : undefined;
            resourceInputs["templateType"] = args ? args.templateType : undefined;
            resourceInputs["defenseTemplateId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DefenseTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefenseTemplate resources.
 */
export interface DefenseTemplateState {
    /**
     * The module to which the protection rule that you want to create belongs. Value:
     * - **waf_group**: the basic protection rule module.
     * - **antiscan**: the scan protection module.
     * - **ip_blacklist**: the IP address blacklist module.
     * - **custom_acl**: the custom rule module.
     * - **whitelist**: the whitelist module.
     * - **region_block**: the region blacklist module.
     * - **custom_response**: the custom response module.
     * - **cc**: the HTTP flood protection module.
     * - **tamperproof**: the website tamper-proofing module.
     * - **dlp**: the data leakage prevention module.
     */
    defenseScene?: pulumi.Input<string>;
    /**
     * Template ID.
     */
    defenseTemplateId?: pulumi.Input<number>;
    /**
     * The name of the protection rule template.
     */
    defenseTemplateName?: pulumi.Input<string>;
    /**
     * The description of the protection rule template. .
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the Web Application Firewall (WAF) instance. .
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The ID of the Alibaba Cloud resource group. .
     */
    resourceManagerResourceGroupId?: pulumi.Input<string>;
    /**
     * The status of the protection rule template. Valid values:
     * - **0**: disabled.
     * - **1**: enabled.
     */
    status?: pulumi.Input<string>;
    /**
     * The origin of the protection rule template that you want to create. Set the value to **custom**. The value specifies that the protection rule template is a custom template. .
     */
    templateOrigin?: pulumi.Input<string>;
    /**
     * The type of the protection rule template. Valid values:
     * - **user_default:** default template.
     * - **user_custom:** custom template.
     */
    templateType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DefenseTemplate resource.
 */
export interface DefenseTemplateArgs {
    /**
     * The module to which the protection rule that you want to create belongs. Value:
     * - **waf_group**: the basic protection rule module.
     * - **antiscan**: the scan protection module.
     * - **ip_blacklist**: the IP address blacklist module.
     * - **custom_acl**: the custom rule module.
     * - **whitelist**: the whitelist module.
     * - **region_block**: the region blacklist module.
     * - **custom_response**: the custom response module.
     * - **cc**: the HTTP flood protection module.
     * - **tamperproof**: the website tamper-proofing module.
     * - **dlp**: the data leakage prevention module.
     */
    defenseScene: pulumi.Input<string>;
    /**
     * The name of the protection rule template.
     */
    defenseTemplateName: pulumi.Input<string>;
    /**
     * The description of the protection rule template. .
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the Web Application Firewall (WAF) instance. .
     */
    instanceId: pulumi.Input<string>;
    /**
     * The ID of the Alibaba Cloud resource group. .
     */
    resourceManagerResourceGroupId?: pulumi.Input<string>;
    /**
     * The status of the protection rule template. Valid values:
     * - **0**: disabled.
     * - **1**: enabled.
     */
    status: pulumi.Input<string>;
    /**
     * The origin of the protection rule template that you want to create. Set the value to **custom**. The value specifies that the protection rule template is a custom template. .
     */
    templateOrigin: pulumi.Input<string>;
    /**
     * The type of the protection rule template. Valid values:
     * - **user_default:** default template.
     * - **user_custom:** custom template.
     */
    templateType: pulumi.Input<string>;
}
