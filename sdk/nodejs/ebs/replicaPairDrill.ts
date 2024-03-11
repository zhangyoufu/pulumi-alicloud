// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a EBS Replica Pair Drill resource.
 *
 * For information about EBS Replica Pair Drill and how to use it, see [What is Replica Pair Drill](https://www.alibabacloud.com/help/en/).
 *
 * > **NOTE:** Available since v1.215.0.
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
 * const _default = new alicloud.ebs.ReplicaPairDrill("default", {pairId: "pair-cn-wwo3kjfq5001"});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * EBS Replica Pair Drill can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ebs/replicaPairDrill:ReplicaPairDrill example <pair_id>:<replica_pair_drill_id>
 * ```
 */
export class ReplicaPairDrill extends pulumi.CustomResource {
    /**
     * Get an existing ReplicaPairDrill resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReplicaPairDrillState, opts?: pulumi.CustomResourceOptions): ReplicaPairDrill {
        return new ReplicaPairDrill(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ebs/replicaPairDrill:ReplicaPairDrill';

    /**
     * Returns true if the given object is an instance of ReplicaPairDrill.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicaPairDrill {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicaPairDrill.__pulumiType;
    }

    /**
     * Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
     */
    public readonly pairId!: pulumi.Output<string>;
    /**
     * The first ID of the resource.
     */
    public /*out*/ readonly replicaPairDrillId!: pulumi.Output<string>;
    /**
     * Walkthrough status. _failed: Execution failed._failed: Cleanup failed.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a ReplicaPairDrill resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicaPairDrillArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplicaPairDrillArgs | ReplicaPairDrillState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReplicaPairDrillState | undefined;
            resourceInputs["pairId"] = state ? state.pairId : undefined;
            resourceInputs["replicaPairDrillId"] = state ? state.replicaPairDrillId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ReplicaPairDrillArgs | undefined;
            if ((!args || args.pairId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pairId'");
            }
            resourceInputs["pairId"] = args ? args.pairId : undefined;
            resourceInputs["replicaPairDrillId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReplicaPairDrill.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReplicaPairDrill resources.
 */
export interface ReplicaPairDrillState {
    /**
     * Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
     */
    pairId?: pulumi.Input<string>;
    /**
     * The first ID of the resource.
     */
    replicaPairDrillId?: pulumi.Input<string>;
    /**
     * Walkthrough status. _failed: Execution failed._failed: Cleanup failed.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReplicaPairDrill resource.
 */
export interface ReplicaPairDrillArgs {
    /**
     * Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
     */
    pairId: pulumi.Input<string>;
}
