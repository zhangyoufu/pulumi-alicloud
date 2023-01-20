// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ECS Snapshot resource.
 *
 * For information about ECS Snapshot and how to use it, see [What is Snapshot](https://www.alibabacloud.com/help/en/doc-detail/25524.htm).
 *
 * > **NOTE:** Available in v1.120.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.ecs.EcsSnapshot("default", {
 *     category: "standard",
 *     description: "Test For Terraform",
 *     diskId: "d-gw8csgxxxxxxxxx",
 *     retentionDays: 20,
 *     snapshotName: "tf-test",
 *     tags: {
 *         Created: "TF",
 *         For: "Acceptance-test",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ECS Snapshot can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ecs/ecsSnapshot:EcsSnapshot example <id>
 * ```
 */
export class EcsSnapshot extends pulumi.CustomResource {
    /**
     * Get an existing EcsSnapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EcsSnapshotState, opts?: pulumi.CustomResourceOptions): EcsSnapshot {
        return new EcsSnapshot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/ecsSnapshot:EcsSnapshot';

    /**
     * Returns true if the given object is an instance of EcsSnapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EcsSnapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EcsSnapshot.__pulumiType;
    }

    /**
     * The category of the snapshot. Valid Values: `standard` and `flash`.
     */
    public readonly category!: pulumi.Output<string | undefined>;
    /**
     * The description of the snapshot.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the disk.
     */
    public readonly diskId!: pulumi.Output<string>;
    /**
     * Specifies whether to forcibly delete the snapshot that has been used to create disks.
     */
    public readonly force!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether to enable the instant access feature.
     */
    public readonly instantAccess!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
     */
    public readonly instantAccessRetentionDays!: pulumi.Output<number | undefined>;
    /**
     * Field `name` has been deprecated from provider version 1.120.0. New field `snapshotName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The resource group id.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * The retention period of the snapshot.
     */
    public readonly retentionDays!: pulumi.Output<number | undefined>;
    /**
     * The name of the snapshot.
     */
    public readonly snapshotName!: pulumi.Output<string>;
    /**
     * The status of snapshot.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the snapshot.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a EcsSnapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EcsSnapshotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EcsSnapshotArgs | EcsSnapshotState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EcsSnapshotState | undefined;
            resourceInputs["category"] = state ? state.category : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["diskId"] = state ? state.diskId : undefined;
            resourceInputs["force"] = state ? state.force : undefined;
            resourceInputs["instantAccess"] = state ? state.instantAccess : undefined;
            resourceInputs["instantAccessRetentionDays"] = state ? state.instantAccessRetentionDays : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["retentionDays"] = state ? state.retentionDays : undefined;
            resourceInputs["snapshotName"] = state ? state.snapshotName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EcsSnapshotArgs | undefined;
            if ((!args || args.diskId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskId'");
            }
            resourceInputs["category"] = args ? args.category : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["diskId"] = args ? args.diskId : undefined;
            resourceInputs["force"] = args ? args.force : undefined;
            resourceInputs["instantAccess"] = args ? args.instantAccess : undefined;
            resourceInputs["instantAccessRetentionDays"] = args ? args.instantAccessRetentionDays : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["retentionDays"] = args ? args.retentionDays : undefined;
            resourceInputs["snapshotName"] = args ? args.snapshotName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EcsSnapshot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EcsSnapshot resources.
 */
export interface EcsSnapshotState {
    /**
     * The category of the snapshot. Valid Values: `standard` and `flash`.
     */
    category?: pulumi.Input<string>;
    /**
     * The description of the snapshot.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the disk.
     */
    diskId?: pulumi.Input<string>;
    /**
     * Specifies whether to forcibly delete the snapshot that has been used to create disks.
     */
    force?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable the instant access feature.
     */
    instantAccess?: pulumi.Input<boolean>;
    /**
     * Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
     */
    instantAccessRetentionDays?: pulumi.Input<number>;
    /**
     * Field `name` has been deprecated from provider version 1.120.0. New field `snapshotName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The resource group id.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The retention period of the snapshot.
     */
    retentionDays?: pulumi.Input<number>;
    /**
     * The name of the snapshot.
     */
    snapshotName?: pulumi.Input<string>;
    /**
     * The status of snapshot.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the snapshot.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a EcsSnapshot resource.
 */
export interface EcsSnapshotArgs {
    /**
     * The category of the snapshot. Valid Values: `standard` and `flash`.
     */
    category?: pulumi.Input<string>;
    /**
     * The description of the snapshot.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the disk.
     */
    diskId: pulumi.Input<string>;
    /**
     * Specifies whether to forcibly delete the snapshot that has been used to create disks.
     */
    force?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable the instant access feature.
     */
    instantAccess?: pulumi.Input<boolean>;
    /**
     * Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
     */
    instantAccessRetentionDays?: pulumi.Input<number>;
    /**
     * Field `name` has been deprecated from provider version 1.120.0. New field `snapshotName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The resource group id.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The retention period of the snapshot.
     */
    retentionDays?: pulumi.Input<number>;
    /**
     * The name of the snapshot.
     */
    snapshotName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the snapshot.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
