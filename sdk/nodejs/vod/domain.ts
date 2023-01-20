// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a VOD Domain resource.
 *
 * For information about VOD Domain and how to use it, see [What is Domain](https://www.alibabacloud.com/help/product/29932.html).
 *
 * > **NOTE:** Available in v1.136.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.vod.Domain("default", {
 *     domainName: "your_domain_name",
 *     scope: "domestic",
 *     sources: [{
 *         sourceContent: "your_source_content",
 *         sourcePort: "80",
 *         sourceType: "domain",
 *     }],
 *     tags: {
 *         key1: "value1",
 *         key2: "value2",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * VOD Domain can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vod/domain:Domain example <domain_name>
 * ```
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainState, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vod/domain:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    /**
     * The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
     */
    public /*out*/ readonly certName!: pulumi.Output<string>;
    /**
     * The URL that is used for health checks.
     */
    public readonly checkUrl!: pulumi.Output<string | undefined>;
    /**
     * The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
     */
    public /*out*/ readonly cname!: pulumi.Output<string>;
    /**
     * The description of the domain name for CDN.
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     */
    public /*out*/ readonly gmtCreated!: pulumi.Output<string>;
    /**
     * The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     */
    public /*out*/ readonly gmtModified!: pulumi.Output<string>;
    /**
     * This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
     */
    public readonly sources!: pulumi.Output<outputs.vod.DomainSource[]>;
    /**
     * Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
     */
    public /*out*/ readonly sslProtocol!: pulumi.Output<string>;
    /**
     * The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
     */
    public /*out*/ readonly sslPub!: pulumi.Output<string>;
    /**
     * The status of the domain name for CDN. Valid values:
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The top-level domain name.
     */
    public readonly topLevelDomain!: pulumi.Output<string | undefined>;
    /**
     * The weight of the origin server.
     */
    public /*out*/ readonly weight!: pulumi.Output<string>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainArgs | DomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainState | undefined;
            resourceInputs["certName"] = state ? state.certName : undefined;
            resourceInputs["checkUrl"] = state ? state.checkUrl : undefined;
            resourceInputs["cname"] = state ? state.cname : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["gmtCreated"] = state ? state.gmtCreated : undefined;
            resourceInputs["gmtModified"] = state ? state.gmtModified : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["sources"] = state ? state.sources : undefined;
            resourceInputs["sslProtocol"] = state ? state.sslProtocol : undefined;
            resourceInputs["sslPub"] = state ? state.sslPub : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["topLevelDomain"] = state ? state.topLevelDomain : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.sources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sources'");
            }
            resourceInputs["checkUrl"] = args ? args.checkUrl : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["sources"] = args ? args.sources : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["topLevelDomain"] = args ? args.topLevelDomain : undefined;
            resourceInputs["certName"] = undefined /*out*/;
            resourceInputs["cname"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["gmtCreated"] = undefined /*out*/;
            resourceInputs["gmtModified"] = undefined /*out*/;
            resourceInputs["sslProtocol"] = undefined /*out*/;
            resourceInputs["sslPub"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["weight"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Domain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Domain resources.
 */
export interface DomainState {
    /**
     * The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
     */
    certName?: pulumi.Input<string>;
    /**
     * The URL that is used for health checks.
     */
    checkUrl?: pulumi.Input<string>;
    /**
     * The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
     */
    cname?: pulumi.Input<string>;
    /**
     * The description of the domain name for CDN.
     */
    description?: pulumi.Input<string>;
    /**
     * The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     */
    gmtCreated?: pulumi.Input<string>;
    /**
     * The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     */
    gmtModified?: pulumi.Input<string>;
    /**
     * This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
     */
    scope?: pulumi.Input<string>;
    /**
     * The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
     */
    sources?: pulumi.Input<pulumi.Input<inputs.vod.DomainSource>[]>;
    /**
     * Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
     */
    sslProtocol?: pulumi.Input<string>;
    /**
     * The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
     */
    sslPub?: pulumi.Input<string>;
    /**
     * The status of the domain name for CDN. Valid values:
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The top-level domain name.
     */
    topLevelDomain?: pulumi.Input<string>;
    /**
     * The weight of the origin server.
     */
    weight?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * The URL that is used for health checks.
     */
    checkUrl?: pulumi.Input<string>;
    /**
     * The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
     */
    domainName: pulumi.Input<string>;
    /**
     * This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
     */
    scope?: pulumi.Input<string>;
    /**
     * The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
     */
    sources: pulumi.Input<pulumi.Input<inputs.vod.DomainSource>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The top-level domain name.
     */
    topLevelDomain?: pulumi.Input<string>;
}
