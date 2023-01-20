// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a DCDN Ipa Domain resource.
 *
 * For information about DCDN Ipa Domain and how to use it, see [What is Ipa Domain](https://www.alibabacloud.com/help/en/doc-detail/130634.html).
 *
 * > **NOTE:** Available in v1.158.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.resourcemanager.getResourceGroups({
 *     nameRegex: "default",
 * });
 * const example = new alicloud.dcdn.IpaDomain("example", {
 *     domainName: "example.com",
 *     resourceGroupId: _default.then(_default => _default.groups?.[0]?.id),
 *     sources: [{
 *         content: "1.1.1.1",
 *         port: 80,
 *         priority: "20",
 *         type: "ipaddr",
 *         weight: 10,
 *     }],
 *     scope: "overseas",
 *     status: "online",
 * });
 * ```
 *
 * ## Import
 *
 * DCDN Ipa Domain can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:dcdn/ipaDomain:IpaDomain example <domain_name>
 * ```
 */
export class IpaDomain extends pulumi.CustomResource {
    /**
     * Get an existing IpaDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpaDomainState, opts?: pulumi.CustomResourceOptions): IpaDomain {
        return new IpaDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dcdn/ipaDomain:IpaDomain';

    /**
     * Returns true if the given object is an instance of IpaDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpaDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpaDomain.__pulumiType;
    }

    /**
     * The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The accelerated region. Valid values: `domestic`, `global`, `overseas`.
     */
    public readonly scope!: pulumi.Output<string>;
    /**
     * Sources. See the following `Block sources`.
     */
    public readonly sources!: pulumi.Output<outputs.dcdn.IpaDomainSource[]>;
    /**
     * The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
     */
    public readonly status!: pulumi.Output<string>;

    /**
     * Create a IpaDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpaDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpaDomainArgs | IpaDomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpaDomainState | undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["sources"] = state ? state.sources : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as IpaDomainArgs | undefined;
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.sources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sources'");
            }
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["sources"] = args ? args.sources : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpaDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpaDomain resources.
 */
export interface IpaDomainState {
    /**
     * The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
     */
    domainName?: pulumi.Input<string>;
    /**
     * The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The accelerated region. Valid values: `domestic`, `global`, `overseas`.
     */
    scope?: pulumi.Input<string>;
    /**
     * Sources. See the following `Block sources`.
     */
    sources?: pulumi.Input<pulumi.Input<inputs.dcdn.IpaDomainSource>[]>;
    /**
     * The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpaDomain resource.
 */
export interface IpaDomainArgs {
    /**
     * The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
     */
    domainName: pulumi.Input<string>;
    /**
     * The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The accelerated region. Valid values: `domestic`, `global`, `overseas`.
     */
    scope?: pulumi.Input<string>;
    /**
     * Sources. See the following `Block sources`.
     */
    sources: pulumi.Input<pulumi.Input<inputs.dcdn.IpaDomainSource>[]>;
    /**
     * The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
     */
    status?: pulumi.Input<string>;
}
