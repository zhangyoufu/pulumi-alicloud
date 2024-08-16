// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource will help you to manager a [Table Store](https://www.alibabacloud.com/help/doc-detail/27280.htm) Instance.
 * It is foundation of creating data table.
 *
 * > **NOTE:** Available since v1.10.0.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const _default = new random.index.Integer("default", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const defaultInstance = new alicloud.ots.Instance("default", {
 *     name: `${name}-${_default.result}`,
 *     description: name,
 *     accessedBy: "Vpc",
 *     tags: {
 *         Created: "TF",
 *         For: "Building table",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * OTS instance can be imported using instance id or name, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ots/instance:Instance foo "my-ots-instance"
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ots/instance:Instance';

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
     * The network limitation of accessing instance. Valid values:
     * * `Any` - Allow all network to access the instance.
     * * `Vpc` - Only can the attached VPC allow to access the instance.
     * * `ConsoleOrVpc` - Allow web console or the attached VPC to access the instance.
     *
     * Default to "Any".
     */
    public readonly accessedBy!: pulumi.Output<string>;
    /**
     * The description of the instance. Currently, it does not support modifying.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The type of instance. Valid values are "Capacity" and "HighPerformance". Default to "HighPerformance".
     */
    public readonly instanceType!: pulumi.Output<string | undefined>;
    /**
     * The name of the instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The set of request sources that are allowed access. Valid optional values:
     * * `TRUST_PROXY` - Trusted proxy, usually the Alibaba Cloud console.
     *
     * Default to ["TRUST_PROXY"].
     */
    public readonly networkSourceAcls!: pulumi.Output<string[]>;
    /**
     * The set of network types that are allowed access. Valid optional values:
     * * `CLASSIC` - Classic network.
     * * `VPC` - VPC network.
     * * `INTERNET` - Public internet.
     *
     * Default to ["VPC", "CLASSIC", "INTERNET"].
     */
    public readonly networkTypeAcls!: pulumi.Output<string[]>;
    /**
     * The resource group the instance belongs to.
     * Default to Alibaba Cloud default resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the instance.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["accessedBy"] = state ? state.accessedBy : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkSourceAcls"] = state ? state.networkSourceAcls : undefined;
            resourceInputs["networkTypeAcls"] = state ? state.networkTypeAcls : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            resourceInputs["accessedBy"] = args ? args.accessedBy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkSourceAcls"] = args ? args.networkSourceAcls : undefined;
            resourceInputs["networkTypeAcls"] = args ? args.networkTypeAcls : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The network limitation of accessing instance. Valid values:
     * * `Any` - Allow all network to access the instance.
     * * `Vpc` - Only can the attached VPC allow to access the instance.
     * * `ConsoleOrVpc` - Allow web console or the attached VPC to access the instance.
     *
     * Default to "Any".
     */
    accessedBy?: pulumi.Input<string>;
    /**
     * The description of the instance. Currently, it does not support modifying.
     */
    description?: pulumi.Input<string>;
    /**
     * The type of instance. Valid values are "Capacity" and "HighPerformance". Default to "HighPerformance".
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The name of the instance.
     */
    name?: pulumi.Input<string>;
    /**
     * The set of request sources that are allowed access. Valid optional values:
     * * `TRUST_PROXY` - Trusted proxy, usually the Alibaba Cloud console.
     *
     * Default to ["TRUST_PROXY"].
     */
    networkSourceAcls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The set of network types that are allowed access. Valid optional values:
     * * `CLASSIC` - Classic network.
     * * `VPC` - VPC network.
     * * `INTERNET` - Public internet.
     *
     * Default to ["VPC", "CLASSIC", "INTERNET"].
     */
    networkTypeAcls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The resource group the instance belongs to.
     * Default to Alibaba Cloud default resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the instance.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The network limitation of accessing instance. Valid values:
     * * `Any` - Allow all network to access the instance.
     * * `Vpc` - Only can the attached VPC allow to access the instance.
     * * `ConsoleOrVpc` - Allow web console or the attached VPC to access the instance.
     *
     * Default to "Any".
     */
    accessedBy?: pulumi.Input<string>;
    /**
     * The description of the instance. Currently, it does not support modifying.
     */
    description?: pulumi.Input<string>;
    /**
     * The type of instance. Valid values are "Capacity" and "HighPerformance". Default to "HighPerformance".
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The name of the instance.
     */
    name?: pulumi.Input<string>;
    /**
     * The set of request sources that are allowed access. Valid optional values:
     * * `TRUST_PROXY` - Trusted proxy, usually the Alibaba Cloud console.
     *
     * Default to ["TRUST_PROXY"].
     */
    networkSourceAcls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The set of network types that are allowed access. Valid optional values:
     * * `CLASSIC` - Classic network.
     * * `VPC` - VPC network.
     * * `INTERNET` - Public internet.
     *
     * Default to ["VPC", "CLASSIC", "INTERNET"].
     */
    networkTypeAcls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The resource group the instance belongs to.
     * Default to Alibaba Cloud default resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the instance.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
