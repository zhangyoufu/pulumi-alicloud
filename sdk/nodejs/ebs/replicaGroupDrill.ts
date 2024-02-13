// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a EBS Replica Group Drill resource.
 *
 * For information about EBS Replica Group Drill and how to use it, see [What is Replica Group Drill](https://www.alibabacloud.com/help/en/).
 *
 * > **NOTE:** Available since v1.215.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const _default = new alicloud.ebs.ReplicaGroupDrill("default", {groupId: "pg-m1H9aaOUIGsDUwgZ"});
 * ```
 *
 * ## Import
 *
 * EBS Replica Group Drill can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ebs/replicaGroupDrill:ReplicaGroupDrill example <group_id>:<replica_group_drill_id>
 * ```
 */
export class ReplicaGroupDrill extends pulumi.CustomResource {
    /**
     * Get an existing ReplicaGroupDrill resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReplicaGroupDrillState, opts?: pulumi.CustomResourceOptions): ReplicaGroupDrill {
        return new ReplicaGroupDrill(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ebs/replicaGroupDrill:ReplicaGroupDrill';

    /**
     * Returns true if the given object is an instance of ReplicaGroupDrill.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicaGroupDrill {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicaGroupDrill.__pulumiType;
    }

    /**
     * The ID of the replication group. You can use the describediskreplicaggroups interface to query the asynchronous replication group list to obtain the value of the replication group ID input parameter.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * The first ID of the resource.
     */
    public /*out*/ readonly replicaGroupDrillId!: pulumi.Output<string>;
    /**
     * Walkthrough status. _failed: Execution failed._failed: Cleanup failed.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a ReplicaGroupDrill resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicaGroupDrillArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplicaGroupDrillArgs | ReplicaGroupDrillState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReplicaGroupDrillState | undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["replicaGroupDrillId"] = state ? state.replicaGroupDrillId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ReplicaGroupDrillArgs | undefined;
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["replicaGroupDrillId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReplicaGroupDrill.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReplicaGroupDrill resources.
 */
export interface ReplicaGroupDrillState {
    /**
     * The ID of the replication group. You can use the describediskreplicaggroups interface to query the asynchronous replication group list to obtain the value of the replication group ID input parameter.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The first ID of the resource.
     */
    replicaGroupDrillId?: pulumi.Input<string>;
    /**
     * Walkthrough status. _failed: Execution failed._failed: Cleanup failed.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReplicaGroupDrill resource.
 */
export interface ReplicaGroupDrillArgs {
    /**
     * The ID of the replication group. You can use the describediskreplicaggroups interface to query the asynchronous replication group list to obtain the value of the replication group ID input parameter.
     */
    groupId: pulumi.Input<string>;
}
