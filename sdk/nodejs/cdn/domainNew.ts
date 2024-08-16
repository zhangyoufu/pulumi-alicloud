// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a CDN Domain resource. CDN domain name.
 *
 * For information about CDN Domain and how to use it, see [What is Domain](https://www.alibabacloud.com/help/en/cdn/developer-reference/api-cdn-2018-05-10-addcdndomain).
 *
 * > **NOTE:** Available since v1.34.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const _default = new random.index.Integer("default", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const defaultDomainNew = new alicloud.cdn.DomainNew("default", {
 *     scope: "overseas",
 *     domainName: `mycdndomain-${_default.result}.alicloud-provider.cn`,
 *     cdnType: "web",
 *     sources: [{
 *         type: "ipaddr",
 *         content: "1.1.1.1",
 *         priority: 20,
 *         port: 80,
 *         weight: 15,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * CDN Domain can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cdn/domainNew:DomainNew example <id>
 * ```
 */
export class DomainNew extends pulumi.CustomResource {
    /**
     * Get an existing DomainNew resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainNewState, opts?: pulumi.CustomResourceOptions): DomainNew {
        return new DomainNew(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cdn/domainNew:DomainNew';

    /**
     * Returns true if the given object is an instance of DomainNew.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainNew {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainNew.__pulumiType;
    }

    /**
     * Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
     */
    public readonly cdnType!: pulumi.Output<string>;
    /**
     * Certificate configuration. See `certificateConfig` below.
     */
    public readonly certificateConfig!: pulumi.Output<outputs.cdn.DomainNewCertificateConfig>;
    /**
     * Health test URL.
     */
    public readonly checkUrl!: pulumi.Output<string | undefined>;
    /**
     * The CNAME domain name corresponding to the accelerated domain name.
     */
    public /*out*/ readonly cname!: pulumi.Output<string>;
    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users. Value:
     * - **domestic**: Mainland China only.
     * - **overseas**: Global (excluding Mainland China).
     * - **global**: global.
     * The default value is **domestic**.
     */
    public readonly scope!: pulumi.Output<string>;
    /**
     * The source address list of the accelerated domain. Defaults to null. See `sources` below.
     */
    public readonly sources!: pulumi.Output<outputs.cdn.DomainNewSource[]>;
    /**
     * The status of the resource.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The tag of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a DomainNew resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainNewArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainNewArgs | DomainNewState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainNewState | undefined;
            resourceInputs["cdnType"] = state ? state.cdnType : undefined;
            resourceInputs["certificateConfig"] = state ? state.certificateConfig : undefined;
            resourceInputs["checkUrl"] = state ? state.checkUrl : undefined;
            resourceInputs["cname"] = state ? state.cname : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["sources"] = state ? state.sources : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DomainNewArgs | undefined;
            if ((!args || args.cdnType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cdnType'");
            }
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.sources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sources'");
            }
            resourceInputs["cdnType"] = args ? args.cdnType : undefined;
            resourceInputs["certificateConfig"] = args ? args.certificateConfig : undefined;
            resourceInputs["checkUrl"] = args ? args.checkUrl : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["sources"] = args ? args.sources : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["cname"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DomainNew.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainNew resources.
 */
export interface DomainNewState {
    /**
     * Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
     */
    cdnType?: pulumi.Input<string>;
    /**
     * Certificate configuration. See `certificateConfig` below.
     */
    certificateConfig?: pulumi.Input<inputs.cdn.DomainNewCertificateConfig>;
    /**
     * Health test URL.
     */
    checkUrl?: pulumi.Input<string>;
    /**
     * The CNAME domain name corresponding to the accelerated domain name.
     */
    cname?: pulumi.Input<string>;
    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users. Value:
     * - **domestic**: Mainland China only.
     * - **overseas**: Global (excluding Mainland China).
     * - **global**: global.
     * The default value is **domestic**.
     */
    scope?: pulumi.Input<string>;
    /**
     * The source address list of the accelerated domain. Defaults to null. See `sources` below.
     */
    sources?: pulumi.Input<pulumi.Input<inputs.cdn.DomainNewSource>[]>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
    /**
     * The tag of the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a DomainNew resource.
 */
export interface DomainNewArgs {
    /**
     * Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
     */
    cdnType: pulumi.Input<string>;
    /**
     * Certificate configuration. See `certificateConfig` below.
     */
    certificateConfig?: pulumi.Input<inputs.cdn.DomainNewCertificateConfig>;
    /**
     * Health test URL.
     */
    checkUrl?: pulumi.Input<string>;
    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    domainName: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users. Value:
     * - **domestic**: Mainland China only.
     * - **overseas**: Global (excluding Mainland China).
     * - **global**: global.
     * The default value is **domestic**.
     */
    scope?: pulumi.Input<string>;
    /**
     * The source address list of the accelerated domain. Defaults to null. See `sources` below.
     */
    sources: pulumi.Input<pulumi.Input<inputs.cdn.DomainNewSource>[]>;
    /**
     * The tag of the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
