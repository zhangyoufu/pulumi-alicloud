// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CR Chart Repository resource.
 *
 * For information about CR Chart Repository and how to use it, see [What is Chart Repository](https://www.alibabacloud.com/help/en/acr/developer-reference/api-cr-2018-12-01-createchartrepository).
 *
 * > **NOTE:** Available since v1.149.0.
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
 * const name = config.get("name") || "tf-example";
 * const example = new alicloud.cr.RegistryEnterpriseInstance("example", {
 *     paymentType: "Subscription",
 *     period: 1,
 *     renewPeriod: 0,
 *     renewalStatus: "ManualRenewal",
 *     instanceType: "Advanced",
 *     instanceName: name,
 * });
 * const exampleChartNamespace = new alicloud.cr.ChartNamespace("example", {
 *     instanceId: example.id,
 *     namespaceName: name,
 * });
 * const exampleChartRepository = new alicloud.cr.ChartRepository("example", {
 *     repoNamespaceName: exampleChartNamespace.namespaceName,
 *     instanceId: exampleChartNamespace.instanceId,
 *     repoName: name,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * CR Chart Repository can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cr/chartRepository:ChartRepository example <instance_id>:<repo_namespace_name>:<repo_name>
 * ```
 */
export class ChartRepository extends pulumi.CustomResource {
    /**
     * Get an existing ChartRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ChartRepositoryState, opts?: pulumi.CustomResourceOptions): ChartRepository {
        return new ChartRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cr/chartRepository:ChartRepository';

    /**
     * Returns true if the given object is an instance of ChartRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ChartRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ChartRepository.__pulumiType;
    }

    /**
     * The ID of the Container Registry instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The name of the repository that you want to create.
     */
    public readonly repoName!: pulumi.Output<string>;
    /**
     * The namespace to which the repository belongs.
     */
    public readonly repoNamespaceName!: pulumi.Output<string>;
    /**
     * The default repository type. Valid values: `PUBLIC`,`PRIVATE`.
     */
    public readonly repoType!: pulumi.Output<string>;
    /**
     * The summary about the repository.
     */
    public readonly summary!: pulumi.Output<string | undefined>;

    /**
     * Create a ChartRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ChartRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ChartRepositoryArgs | ChartRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ChartRepositoryState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["repoName"] = state ? state.repoName : undefined;
            resourceInputs["repoNamespaceName"] = state ? state.repoNamespaceName : undefined;
            resourceInputs["repoType"] = state ? state.repoType : undefined;
            resourceInputs["summary"] = state ? state.summary : undefined;
        } else {
            const args = argsOrState as ChartRepositoryArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.repoName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repoName'");
            }
            if ((!args || args.repoNamespaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repoNamespaceName'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["repoName"] = args ? args.repoName : undefined;
            resourceInputs["repoNamespaceName"] = args ? args.repoNamespaceName : undefined;
            resourceInputs["repoType"] = args ? args.repoType : undefined;
            resourceInputs["summary"] = args ? args.summary : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ChartRepository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ChartRepository resources.
 */
export interface ChartRepositoryState {
    /**
     * The ID of the Container Registry instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The name of the repository that you want to create.
     */
    repoName?: pulumi.Input<string>;
    /**
     * The namespace to which the repository belongs.
     */
    repoNamespaceName?: pulumi.Input<string>;
    /**
     * The default repository type. Valid values: `PUBLIC`,`PRIVATE`.
     */
    repoType?: pulumi.Input<string>;
    /**
     * The summary about the repository.
     */
    summary?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ChartRepository resource.
 */
export interface ChartRepositoryArgs {
    /**
     * The ID of the Container Registry instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The name of the repository that you want to create.
     */
    repoName: pulumi.Input<string>;
    /**
     * The namespace to which the repository belongs.
     */
    repoNamespaceName: pulumi.Input<string>;
    /**
     * The default repository type. Valid values: `PUBLIC`,`PRIVATE`.
     */
    repoType?: pulumi.Input<string>;
    /**
     * The summary about the repository.
     */
    summary?: pulumi.Input<string>;
}
