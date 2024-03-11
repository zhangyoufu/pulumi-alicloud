// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Monitor Service Namespace resource.
 *
 * For information about Cloud Monitor Service Namespace and how to use it, see [What is Namespace](https://www.alibabacloud.com/help/en/cloudmonitor/latest/createhybridmonitornamespace).
 *
 * > **NOTE:** Available since v1.171.0.
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
 * const example = new alicloud.cms.Namespace("example", {
 *     description: "tf-example",
 *     namespace: "tf-example",
 *     specification: "cms.s1.large",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Cloud Monitor Service Namespace can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cms/namespace:Namespace example <id>
 * ```
 */
export class Namespace extends pulumi.CustomResource {
    /**
     * Get an existing Namespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NamespaceState, opts?: pulumi.CustomResourceOptions): Namespace {
        return new Namespace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cms/namespace:Namespace';

    /**
     * Returns true if the given object is an instance of Namespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Namespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Namespace.__pulumiType;
    }

    /**
     * The description of the namespace.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the namespace. The name can contain lowercase letters, digits, and hyphens (-).
     */
    public readonly namespace!: pulumi.Output<string>;
    /**
     * The data retention period. Default value: `cms.s1.3xlarge`. Valid values:
     * - `cms.s1.large`: Data storage duration is 15 days.
     * - `cms.s1.xlarge`: Data storage duration is 32 days.
     * - `cms.s1.2xlarge`: Data storage duration 63 days.
     * - `cms.s1.3xlarge`: Data storage duration 93 days.
     * - `cms.s1.6xlarge`: Data storage duration 185 days.
     * - `cms.s1.12xlarge`: Data storage duration 376 days.
     */
    public readonly specification!: pulumi.Output<string>;

    /**
     * Create a Namespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NamespaceArgs | NamespaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NamespaceState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["specification"] = state ? state.specification : undefined;
        } else {
            const args = argsOrState as NamespaceArgs | undefined;
            if ((!args || args.namespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespace'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["specification"] = args ? args.specification : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Namespace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Namespace resources.
 */
export interface NamespaceState {
    /**
     * The description of the namespace.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the namespace. The name can contain lowercase letters, digits, and hyphens (-).
     */
    namespace?: pulumi.Input<string>;
    /**
     * The data retention period. Default value: `cms.s1.3xlarge`. Valid values:
     * - `cms.s1.large`: Data storage duration is 15 days.
     * - `cms.s1.xlarge`: Data storage duration is 32 days.
     * - `cms.s1.2xlarge`: Data storage duration 63 days.
     * - `cms.s1.3xlarge`: Data storage duration 93 days.
     * - `cms.s1.6xlarge`: Data storage duration 185 days.
     * - `cms.s1.12xlarge`: Data storage duration 376 days.
     */
    specification?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Namespace resource.
 */
export interface NamespaceArgs {
    /**
     * The description of the namespace.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the namespace. The name can contain lowercase letters, digits, and hyphens (-).
     */
    namespace: pulumi.Input<string>;
    /**
     * The data retention period. Default value: `cms.s1.3xlarge`. Valid values:
     * - `cms.s1.large`: Data storage duration is 15 days.
     * - `cms.s1.xlarge`: Data storage duration is 32 days.
     * - `cms.s1.2xlarge`: Data storage duration 63 days.
     * - `cms.s1.3xlarge`: Data storage duration 93 days.
     * - `cms.s1.6xlarge`: Data storage duration 185 days.
     * - `cms.s1.12xlarge`: Data storage duration 376 days.
     */
    specification?: pulumi.Input<string>;
}
