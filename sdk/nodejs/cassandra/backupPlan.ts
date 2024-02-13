// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Cassandra Backup Plan resource.
 *
 * For information about Cassandra Backup Plan and how to use it, see [What is Backup Plan](https://www.alibabacloud.com/help/doc-detail/157522.htm).
 *
 * > **NOTE:** Available in v1.128.0+.
 *
 * ## Import
 *
 * Cassandra Backup Plan can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cassandra/backupPlan:BackupPlan example <cluster_id>:<data_center_id>
 * ```
 */
export class BackupPlan extends pulumi.CustomResource {
    /**
     * Get an existing BackupPlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackupPlanState, opts?: pulumi.CustomResourceOptions): BackupPlan {
        return new BackupPlan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cassandra/backupPlan:BackupPlan';

    /**
     * Returns true if the given object is an instance of BackupPlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackupPlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackupPlan.__pulumiType;
    }

    /**
     * Specifies whether to activate the backup plan. Valid values: `True`, `False`. Default value: `True`.
     */
    public readonly active!: pulumi.Output<boolean>;
    /**
     * The backup cycle. Valid values: `Friday`, `Monday`, `Saturday`, `Sunday`, `Thursday`, `Tuesday`, `Wednesday`.
     */
    public readonly backupPeriod!: pulumi.Output<string>;
    /**
     * The start time of the backup task each day. The time is displayed in UTC and denoted by Z.
     */
    public readonly backupTime!: pulumi.Output<string>;
    /**
     * The ID of the cluster for the backup.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The ID of the data center for the backup in the cluster.
     */
    public readonly dataCenterId!: pulumi.Output<string>;
    /**
     * The duration for which you want to retain the backup. Valid values: 1 to 30. Unit: days. Default value: `30`.
     */
    public readonly retentionPeriod!: pulumi.Output<number>;

    /**
     * Create a BackupPlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupPlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackupPlanArgs | BackupPlanState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackupPlanState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["backupPeriod"] = state ? state.backupPeriod : undefined;
            resourceInputs["backupTime"] = state ? state.backupTime : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["dataCenterId"] = state ? state.dataCenterId : undefined;
            resourceInputs["retentionPeriod"] = state ? state.retentionPeriod : undefined;
        } else {
            const args = argsOrState as BackupPlanArgs | undefined;
            if ((!args || args.backupTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupTime'");
            }
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.dataCenterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataCenterId'");
            }
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["backupPeriod"] = args ? args.backupPeriod : undefined;
            resourceInputs["backupTime"] = args ? args.backupTime : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["dataCenterId"] = args ? args.dataCenterId : undefined;
            resourceInputs["retentionPeriod"] = args ? args.retentionPeriod : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BackupPlan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackupPlan resources.
 */
export interface BackupPlanState {
    /**
     * Specifies whether to activate the backup plan. Valid values: `True`, `False`. Default value: `True`.
     */
    active?: pulumi.Input<boolean>;
    /**
     * The backup cycle. Valid values: `Friday`, `Monday`, `Saturday`, `Sunday`, `Thursday`, `Tuesday`, `Wednesday`.
     */
    backupPeriod?: pulumi.Input<string>;
    /**
     * The start time of the backup task each day. The time is displayed in UTC and denoted by Z.
     */
    backupTime?: pulumi.Input<string>;
    /**
     * The ID of the cluster for the backup.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The ID of the data center for the backup in the cluster.
     */
    dataCenterId?: pulumi.Input<string>;
    /**
     * The duration for which you want to retain the backup. Valid values: 1 to 30. Unit: days. Default value: `30`.
     */
    retentionPeriod?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a BackupPlan resource.
 */
export interface BackupPlanArgs {
    /**
     * Specifies whether to activate the backup plan. Valid values: `True`, `False`. Default value: `True`.
     */
    active?: pulumi.Input<boolean>;
    /**
     * The backup cycle. Valid values: `Friday`, `Monday`, `Saturday`, `Sunday`, `Thursday`, `Tuesday`, `Wednesday`.
     */
    backupPeriod?: pulumi.Input<string>;
    /**
     * The start time of the backup task each day. The time is displayed in UTC and denoted by Z.
     */
    backupTime: pulumi.Input<string>;
    /**
     * The ID of the cluster for the backup.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The ID of the data center for the backup in the cluster.
     */
    dataCenterId: pulumi.Input<string>;
    /**
     * The duration for which you want to retain the backup. Valid values: 1 to 30. Unit: days. Default value: `30`.
     */
    retentionPeriod?: pulumi.Input<number>;
}
