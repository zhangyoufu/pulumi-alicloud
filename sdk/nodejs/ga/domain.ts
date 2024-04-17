// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Ga Domain resource.
 *
 * For information about Ga Domain and how to use it, see [What is Domain](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createdomain).
 *
 * > **NOTE:** Available since v1.197.0.
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
 * const _default = new alicloud.ga.Accelerator("default", {
 *     duration: 1,
 *     autoUseCoupon: true,
 *     spec: "1",
 * });
 * const defaultDomain = new alicloud.ga.Domain("default", {
 *     domain: "changes.com.cn",
 *     acceleratorId: _default.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Ga Domain can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ga/domain:Domain example <accelerator_id>:<domain>
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
    public static readonly __pulumiType = 'alicloud:ga/domain:Domain';

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
     * The ID of the global acceleration instance.
     */
    public readonly acceleratorId!: pulumi.Output<string>;
    /**
     * The accelerated domain name to be added. only top-level domain names are supported, such as 'example.com'.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * The status of the resource
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

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
            resourceInputs["acceleratorId"] = state ? state.acceleratorId : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            if ((!args || args.acceleratorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceleratorId'");
            }
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            resourceInputs["acceleratorId"] = args ? args.acceleratorId : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["status"] = undefined /*out*/;
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
     * The ID of the global acceleration instance.
     */
    acceleratorId?: pulumi.Input<string>;
    /**
     * The accelerated domain name to be added. only top-level domain names are supported, such as 'example.com'.
     */
    domain?: pulumi.Input<string>;
    /**
     * The status of the resource
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * The ID of the global acceleration instance.
     */
    acceleratorId: pulumi.Input<string>;
    /**
     * The accelerated domain name to be added. only top-level domain names are supported, such as 'example.com'.
     */
    domain: pulumi.Input<string>;
}
