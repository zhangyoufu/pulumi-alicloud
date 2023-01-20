// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Enterprise Network (CEN) Transit Router Multicast Domain resource.
 *
 * For information about Cloud Enterprise Network (CEN) Transit Router Multicast Domain and how to use it, see [What is Transit Router Multicast Domain](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-createtransitroutermulticastdomain).
 *
 * > **NOTE:** Available in v1.195.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultInstance = new alicloud.cen.Instance("defaultInstance", {cenInstanceName: "tf-example"});
 * const defaultTransitRouter = new alicloud.cen.TransitRouter("defaultTransitRouter", {
 *     cenId: defaultInstance.id,
 *     supportMulticast: true,
 * });
 * const defaultTransitRouterMulticastDomain = new alicloud.cen.TransitRouterMulticastDomain("defaultTransitRouterMulticastDomain", {
 *     transitRouterId: defaultTransitRouter.transitRouterId,
 *     transitRouterMulticastDomainName: "tf-example-name",
 *     transitRouterMulticastDomainDescription: "tf-example-description",
 * });
 * ```
 *
 * ## Import
 *
 * Cloud Enterprise Network (CEN) Transit Router Multicast Domain can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cen/transitRouterMulticastDomain:TransitRouterMulticastDomain example <id>
 * ```
 */
export class TransitRouterMulticastDomain extends pulumi.CustomResource {
    /**
     * Get an existing TransitRouterMulticastDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TransitRouterMulticastDomainState, opts?: pulumi.CustomResourceOptions): TransitRouterMulticastDomain {
        return new TransitRouterMulticastDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cen/transitRouterMulticastDomain:TransitRouterMulticastDomain';

    /**
     * Returns true if the given object is an instance of TransitRouterMulticastDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitRouterMulticastDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitRouterMulticastDomain.__pulumiType;
    }

    /**
     * The status of the Transit Router Multicast Domain.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The ID of the transit router.
     */
    public readonly transitRouterId!: pulumi.Output<string>;
    /**
     * The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
     */
    public readonly transitRouterMulticastDomainDescription!: pulumi.Output<string | undefined>;
    /**
     * The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
     */
    public readonly transitRouterMulticastDomainName!: pulumi.Output<string | undefined>;

    /**
     * Create a TransitRouterMulticastDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransitRouterMulticastDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TransitRouterMulticastDomainArgs | TransitRouterMulticastDomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TransitRouterMulticastDomainState | undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["transitRouterId"] = state ? state.transitRouterId : undefined;
            resourceInputs["transitRouterMulticastDomainDescription"] = state ? state.transitRouterMulticastDomainDescription : undefined;
            resourceInputs["transitRouterMulticastDomainName"] = state ? state.transitRouterMulticastDomainName : undefined;
        } else {
            const args = argsOrState as TransitRouterMulticastDomainArgs | undefined;
            if ((!args || args.transitRouterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitRouterId'");
            }
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["transitRouterId"] = args ? args.transitRouterId : undefined;
            resourceInputs["transitRouterMulticastDomainDescription"] = args ? args.transitRouterMulticastDomainDescription : undefined;
            resourceInputs["transitRouterMulticastDomainName"] = args ? args.transitRouterMulticastDomainName : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TransitRouterMulticastDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TransitRouterMulticastDomain resources.
 */
export interface TransitRouterMulticastDomainState {
    /**
     * The status of the Transit Router Multicast Domain.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the transit router.
     */
    transitRouterId?: pulumi.Input<string>;
    /**
     * The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
     */
    transitRouterMulticastDomainDescription?: pulumi.Input<string>;
    /**
     * The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
     */
    transitRouterMulticastDomainName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TransitRouterMulticastDomain resource.
 */
export interface TransitRouterMulticastDomainArgs {
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the transit router.
     */
    transitRouterId: pulumi.Input<string>;
    /**
     * The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
     */
    transitRouterMulticastDomainDescription?: pulumi.Input<string>;
    /**
     * The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
     */
    transitRouterMulticastDomainName?: pulumi.Input<string>;
}
