// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ECS Hpc Cluster resource.
 *
 * For information about ECS Hpc Cluster and how to use it, see [What is Hpc Cluster](https://www.alibabacloud.com/help/en/doc-detail/109138.htm).
 *
 * > **NOTE:** Available in v1.116.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.ecs.HpcCluster("example", {description: "For Terraform Test"});
 * ```
 *
 * ## Import
 *
 * ECS Hpc Cluster can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ecs/hpcCluster:HpcCluster example <id>
 * ```
 */
export class HpcCluster extends pulumi.CustomResource {
    /**
     * Get an existing HpcCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HpcClusterState, opts?: pulumi.CustomResourceOptions): HpcCluster {
        return new HpcCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/hpcCluster:HpcCluster';

    /**
     * Returns true if the given object is an instance of HpcCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HpcCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HpcCluster.__pulumiType;
    }

    /**
     * The description of ECS Hpc Cluster.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of ECS Hpc Cluster.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a HpcCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: HpcClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HpcClusterArgs | HpcClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HpcClusterState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as HpcClusterArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HpcCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HpcCluster resources.
 */
export interface HpcClusterState {
    /**
     * The description of ECS Hpc Cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of ECS Hpc Cluster.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HpcCluster resource.
 */
export interface HpcClusterArgs {
    /**
     * The description of ECS Hpc Cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of ECS Hpc Cluster.
     */
    name?: pulumi.Input<string>;
}
