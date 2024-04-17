// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Simple Application Server Snapshot resource.
 *
 * For information about Simple Application Server Snapshot and how to use it, see [What is Snapshot](https://www.alibabacloud.com/help/doc-detail/190452.htm).
 *
 * > **NOTE:** Available since v1.143.0.
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
 * const name = config.get("name") || "tf_example";
 * const default = alicloud.simpleapplicationserver.getImages({});
 * const defaultGetServerPlans = alicloud.simpleapplicationserver.getServerPlans({});
 * const defaultInstance = new alicloud.simpleapplicationserver.Instance("default", {
 *     paymentType: "Subscription",
 *     planId: defaultGetServerPlans.then(defaultGetServerPlans => defaultGetServerPlans.plans?.[0]?.id),
 *     instanceName: name,
 *     imageId: _default.then(_default => _default.images?.[0]?.id),
 *     period: 1,
 *     dataDiskSize: 100,
 * });
 * const defaultGetServerDisks = alicloud.simpleapplicationserver.getServerDisksOutput({
 *     instanceId: defaultInstance.id,
 * });
 * const defaultSnapshot = new alicloud.simpleapplicationserver.Snapshot("default", {
 *     diskId: defaultGetServerDisks.apply(defaultGetServerDisks => defaultGetServerDisks.ids?.[0]),
 *     snapshotName: name,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Simple Application Server Snapshot can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:simpleapplicationserver/snapshot:Snapshot example <id>
 * ```
 */
export class Snapshot extends pulumi.CustomResource {
    /**
     * Get an existing Snapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnapshotState, opts?: pulumi.CustomResourceOptions): Snapshot {
        return new Snapshot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:simpleapplicationserver/snapshot:Snapshot';

    /**
     * Returns true if the given object is an instance of Snapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Snapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Snapshot.__pulumiType;
    }

    /**
     * The ID of the disk.
     */
    public readonly diskId!: pulumi.Output<string>;
    /**
     * The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
     */
    public readonly snapshotName!: pulumi.Output<string>;
    /**
     * The status of the snapshot. Valid values: `Progressing`, `Accomplished` and `Failed`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Snapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnapshotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnapshotArgs | SnapshotState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnapshotState | undefined;
            resourceInputs["diskId"] = state ? state.diskId : undefined;
            resourceInputs["snapshotName"] = state ? state.snapshotName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as SnapshotArgs | undefined;
            if ((!args || args.diskId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskId'");
            }
            if ((!args || args.snapshotName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'snapshotName'");
            }
            resourceInputs["diskId"] = args ? args.diskId : undefined;
            resourceInputs["snapshotName"] = args ? args.snapshotName : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Snapshot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Snapshot resources.
 */
export interface SnapshotState {
    /**
     * The ID of the disk.
     */
    diskId?: pulumi.Input<string>;
    /**
     * The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
     */
    snapshotName?: pulumi.Input<string>;
    /**
     * The status of the snapshot. Valid values: `Progressing`, `Accomplished` and `Failed`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Snapshot resource.
 */
export interface SnapshotArgs {
    /**
     * The ID of the disk.
     */
    diskId: pulumi.Input<string>;
    /**
     * The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
     */
    snapshotName: pulumi.Input<string>;
}
