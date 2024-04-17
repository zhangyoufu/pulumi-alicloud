// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a DBFS Snapshot resource.
 *
 * For information about DBFS Snapshot and how to use it.
 *
 * > **NOTE:** Available since v1.156.0.
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
 * const zoneId = "cn-hangzhou-i";
 * const example = alicloud.ecs.getInstanceTypes({
 *     availabilityZone: zoneId,
 *     instanceTypeFamily: "ecs.g7se",
 * });
 * const exampleGetImages = Promise.all([example, example.then(example => example.instanceTypes).length]).then(([example, length]) => alicloud.ecs.getImages({
 *     instanceType: example.instanceTypes[length - 1].id,
 *     nameRegex: "^aliyun_2_1903_x64_20G_alibase_20231221.vhd",
 *     owners: "system",
 * }));
 * const default = alicloud.vpc.getNetworks({
 *     nameRegex: "^default-NODELETING$",
 * });
 * const defaultGetSwitches = _default.then(_default => alicloud.vpc.getSwitches({
 *     vpcId: _default.ids?.[0],
 *     zoneId: zoneId,
 * }));
 * const exampleSecurityGroup = new alicloud.ecs.SecurityGroup("example", {
 *     name: name,
 *     vpcId: _default.then(_default => _default.ids?.[0]),
 * });
 * const defaultInstance = new alicloud.ecs.Instance("default", {
 *     availabilityZone: zoneId,
 *     instanceName: name,
 *     imageId: exampleGetImages.then(exampleGetImages => exampleGetImages.images?.[0]?.id),
 *     instanceType: Promise.all([example, example.then(example => example.instanceTypes).length]).then(([example, length]) => example.instanceTypes[length - 1].id),
 *     securityGroups: [exampleSecurityGroup.id],
 *     vswitchId: defaultGetSwitches.then(defaultGetSwitches => defaultGetSwitches.ids?.[0]),
 *     systemDiskCategory: "cloud_essd",
 * });
 * const defaultInstance2 = new alicloud.databasefilesystem.Instance("default", {
 *     category: "enterprise",
 *     zoneId: defaultInstance.availabilityZone,
 *     performanceLevel: "PL1",
 *     fsName: name,
 *     size: 100,
 * });
 * const defaultInstanceAttachment = new alicloud.databasefilesystem.InstanceAttachment("default", {
 *     ecsId: defaultInstance.id,
 *     instanceId: defaultInstance2.id,
 * });
 * const exampleSnapshot = new alicloud.databasefilesystem.Snapshot("example", {
 *     instanceId: defaultInstanceAttachment.instanceId,
 *     snapshotName: name,
 *     description: name,
 *     retentionDays: 30,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * DBFS Snapshot can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:databasefilesystem/snapshot:Snapshot example <id>
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
    public static readonly __pulumiType = 'alicloud:databasefilesystem/snapshot:Snapshot';

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
     * Description of the snapshot. The description must be `2` to `256` characters in length. It must start with a letter, and cannot start with `http://` or `https://`.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether to force deletion of snapshots.
     */
    public readonly force!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the database file system.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The retention time of the snapshot. Unit: days. Snapshots are automatically released after the retention time expires. Valid values: `1` to `65536`.
     */
    public readonly retentionDays!: pulumi.Output<number | undefined>;
    /**
     * The display name of the snapshot. The length is `2` to `128` characters. It must start with a large or small letter or Chinese, and cannot start with `http://` and `https://`. It can contain numbers, colons (:), underscores (_), or hyphens (-). To prevent name conflicts with automatic snapshots, you cannot start with `auto`.
     */
    public readonly snapshotName!: pulumi.Output<string | undefined>;
    /**
     * The status of the Snapshot.
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
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["force"] = state ? state.force : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["retentionDays"] = state ? state.retentionDays : undefined;
            resourceInputs["snapshotName"] = state ? state.snapshotName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as SnapshotArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["force"] = args ? args.force : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["retentionDays"] = args ? args.retentionDays : undefined;
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
     * Description of the snapshot. The description must be `2` to `256` characters in length. It must start with a letter, and cannot start with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether to force deletion of snapshots.
     */
    force?: pulumi.Input<boolean>;
    /**
     * The ID of the database file system.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The retention time of the snapshot. Unit: days. Snapshots are automatically released after the retention time expires. Valid values: `1` to `65536`.
     */
    retentionDays?: pulumi.Input<number>;
    /**
     * The display name of the snapshot. The length is `2` to `128` characters. It must start with a large or small letter or Chinese, and cannot start with `http://` and `https://`. It can contain numbers, colons (:), underscores (_), or hyphens (-). To prevent name conflicts with automatic snapshots, you cannot start with `auto`.
     */
    snapshotName?: pulumi.Input<string>;
    /**
     * The status of the Snapshot.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Snapshot resource.
 */
export interface SnapshotArgs {
    /**
     * Description of the snapshot. The description must be `2` to `256` characters in length. It must start with a letter, and cannot start with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether to force deletion of snapshots.
     */
    force?: pulumi.Input<boolean>;
    /**
     * The ID of the database file system.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The retention time of the snapshot. Unit: days. Snapshots are automatically released after the retention time expires. Valid values: `1` to `65536`.
     */
    retentionDays?: pulumi.Input<number>;
    /**
     * The display name of the snapshot. The length is `2` to `128` characters. It must start with a large or small letter or Chinese, and cannot start with `http://` and `https://`. It can contain numbers, colons (:), underscores (_), or hyphens (-). To prevent name conflicts with automatic snapshots, you cannot start with `auto`.
     */
    snapshotName?: pulumi.Input<string>;
}
