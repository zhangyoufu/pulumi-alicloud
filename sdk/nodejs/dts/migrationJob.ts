// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a DTS Migration Job resource.
 *
 * For information about DTS Migration Job and how to use it, see [What is Migration Job](https://www.alibabacloud.com/help/en/doc-detail/208399.html).
 *
 * > **NOTE:** Available since v1.157.0.
 *
 * ## Notice
 *
 * 1. The expiration time cannot be changed after the work of the annual and monthly subscription suspended;
 * 2. After the pay-as-you-go type job suspended, your job configuration fee will still be charged;
 * 3. If the task suspended for more than 6 hours, the task will not start successfully.
 * 4. Suspending the task will only stop writing to the target library, but will still continue to obtain the incremental log of the source, so that the task can be quickly resumed after the suspension is canceled. Therefore, some resources of the source library, such as bandwidth resources, will continue to be occupied during the period.
 * 5. Charges will continue during the task suspension period. If you need to stop charging, please release the instance
 * 6. When a DTS instance suspended for more than 7 days, the instance cannot be resumed, and the status will change from suspended to failed.
 *
 * ## Import
 *
 * DTS Migration Job can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:dts/migrationJob:MigrationJob example <id>
 * ```
 */
export class MigrationJob extends pulumi.CustomResource {
    /**
     * Get an existing MigrationJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MigrationJobState, opts?: pulumi.CustomResourceOptions): MigrationJob {
        return new MigrationJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dts/migrationJob:MigrationJob';

    /**
     * Returns true if the given object is an instance of MigrationJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MigrationJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MigrationJob.__pulumiType;
    }

    /**
     * Start time in Unix timestamp format.
     */
    public readonly checkpoint!: pulumi.Output<string>;
    /**
     * Whether to execute DTS supports schema migration.
     */
    public readonly dataInitialization!: pulumi.Output<boolean>;
    /**
     * Whether to perform incremental data migration.
     */
    public readonly dataSynchronization!: pulumi.Output<boolean>;
    /**
     * Migration object, in the format of JSON strings. For detailed definition instructions, please refer to [the description of migration, migration or subscription objects](https://help.aliyun.com/document_detail/209545.html).
     */
    public readonly dbList!: pulumi.Output<string>;
    /**
     * The name of migrate the database.
     */
    public readonly destinationEndpointDatabaseName!: pulumi.Output<string | undefined>;
    /**
     * The type of destination database. Valid values: `ADS`, `ADB30`, `AS400`, `DATAHUB`, `DB2`, `GREENPLUM`, `KAFKA`, `MONGODB`, `MSSQL`, `MySQL`, `ORACLE`, `PolarDB`, `POLARDBX20`, `POLARDB_O`, `PostgreSQL`.
     */
    public readonly destinationEndpointEngineName!: pulumi.Output<string>;
    /**
     * The ID of destination instance.
     */
    public readonly destinationEndpointInstanceId!: pulumi.Output<string | undefined>;
    /**
     * The type of destination instance. Valid values: `ADS`, `CEN`, `DATAHUB`, `DG`, `ECS`, `EXPRESS`, `GREENPLUM`, `MONGODB`, `OTHER`, `PolarDB`, `POLARDBX20`, `RDS`.
     */
    public readonly destinationEndpointInstanceType!: pulumi.Output<string>;
    /**
     * The ip of source endpoint.
     */
    public readonly destinationEndpointIp!: pulumi.Output<string | undefined>;
    /**
     * The SID of Oracle database.
     */
    public readonly destinationEndpointOracleSid!: pulumi.Output<string | undefined>;
    /**
     * The password of database account.
     */
    public readonly destinationEndpointPassword!: pulumi.Output<string | undefined>;
    /**
     * The port of source endpoint.
     */
    public readonly destinationEndpointPort!: pulumi.Output<string | undefined>;
    /**
     * The region of destination instance.
     */
    public readonly destinationEndpointRegion!: pulumi.Output<string | undefined>;
    /**
     * The username of database account.
     */
    public readonly destinationEndpointUserName!: pulumi.Output<string | undefined>;
    /**
     * The Migration instance ID. The ID of `alicloud.dts.MigrationInstance`.
     */
    public readonly dtsInstanceId!: pulumi.Output<string>;
    /**
     * The name of migration job.
     */
    public readonly dtsJobName!: pulumi.Output<string>;
    /**
     * The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`, `xxlarge`.
     */
    public readonly instanceClass!: pulumi.Output<string>;
    /**
     * The name of migrate the database.
     */
    public readonly sourceEndpointDatabaseName!: pulumi.Output<string | undefined>;
    /**
     * The type of source database. Valid values: `AS400`, `DB2`, `DMSPOLARDB`, `HBASE`, `MONGODB`, `MSSQL`, `MySQL`, `ORACLE`, `PolarDB`, `POLARDBX20`, `POLARDB_O`, `POSTGRESQL`, `TERADATA`.
     */
    public readonly sourceEndpointEngineName!: pulumi.Output<string>;
    /**
     * The ID of source instance.
     */
    public readonly sourceEndpointInstanceId!: pulumi.Output<string | undefined>;
    /**
     * The type of source instance. Valid values: `CEN`, `DG`, `DISTRIBUTED_DMSLOGICDB`, `ECS`, `EXPRESS`, `MONGODB`, `OTHER`, `PolarDB`, `POLARDBX20`, `RDS`.
     */
    public readonly sourceEndpointInstanceType!: pulumi.Output<string>;
    /**
     * The ip of source endpoint.
     */
    public readonly sourceEndpointIp!: pulumi.Output<string | undefined>;
    /**
     * The SID of Oracle database.
     */
    public readonly sourceEndpointOracleSid!: pulumi.Output<string | undefined>;
    /**
     * The Alibaba Cloud account ID to which the source instance belongs.
     */
    public readonly sourceEndpointOwnerId!: pulumi.Output<string | undefined>;
    /**
     * The password of database account.
     */
    public readonly sourceEndpointPassword!: pulumi.Output<string | undefined>;
    /**
     * The port of source endpoint.
     */
    public readonly sourceEndpointPort!: pulumi.Output<string | undefined>;
    /**
     * The region of source instance.
     */
    public readonly sourceEndpointRegion!: pulumi.Output<string | undefined>;
    /**
     * The name of the role configured for the cloud account to which the source instance belongs.
     */
    public readonly sourceEndpointRole!: pulumi.Output<string | undefined>;
    /**
     * The username of database account.
     */
    public readonly sourceEndpointUserName!: pulumi.Output<string | undefined>;
    /**
     * The status of the resource. Valid values: `Migrating`, `Suspending`. You can suspend the task by specifying `Suspending` and start the task by specifying `Migrating`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Whether to perform a database table structure to migrate.
     */
    public readonly structureInitialization!: pulumi.Output<boolean>;

    /**
     * Create a MigrationJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MigrationJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MigrationJobArgs | MigrationJobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MigrationJobState | undefined;
            resourceInputs["checkpoint"] = state ? state.checkpoint : undefined;
            resourceInputs["dataInitialization"] = state ? state.dataInitialization : undefined;
            resourceInputs["dataSynchronization"] = state ? state.dataSynchronization : undefined;
            resourceInputs["dbList"] = state ? state.dbList : undefined;
            resourceInputs["destinationEndpointDatabaseName"] = state ? state.destinationEndpointDatabaseName : undefined;
            resourceInputs["destinationEndpointEngineName"] = state ? state.destinationEndpointEngineName : undefined;
            resourceInputs["destinationEndpointInstanceId"] = state ? state.destinationEndpointInstanceId : undefined;
            resourceInputs["destinationEndpointInstanceType"] = state ? state.destinationEndpointInstanceType : undefined;
            resourceInputs["destinationEndpointIp"] = state ? state.destinationEndpointIp : undefined;
            resourceInputs["destinationEndpointOracleSid"] = state ? state.destinationEndpointOracleSid : undefined;
            resourceInputs["destinationEndpointPassword"] = state ? state.destinationEndpointPassword : undefined;
            resourceInputs["destinationEndpointPort"] = state ? state.destinationEndpointPort : undefined;
            resourceInputs["destinationEndpointRegion"] = state ? state.destinationEndpointRegion : undefined;
            resourceInputs["destinationEndpointUserName"] = state ? state.destinationEndpointUserName : undefined;
            resourceInputs["dtsInstanceId"] = state ? state.dtsInstanceId : undefined;
            resourceInputs["dtsJobName"] = state ? state.dtsJobName : undefined;
            resourceInputs["instanceClass"] = state ? state.instanceClass : undefined;
            resourceInputs["sourceEndpointDatabaseName"] = state ? state.sourceEndpointDatabaseName : undefined;
            resourceInputs["sourceEndpointEngineName"] = state ? state.sourceEndpointEngineName : undefined;
            resourceInputs["sourceEndpointInstanceId"] = state ? state.sourceEndpointInstanceId : undefined;
            resourceInputs["sourceEndpointInstanceType"] = state ? state.sourceEndpointInstanceType : undefined;
            resourceInputs["sourceEndpointIp"] = state ? state.sourceEndpointIp : undefined;
            resourceInputs["sourceEndpointOracleSid"] = state ? state.sourceEndpointOracleSid : undefined;
            resourceInputs["sourceEndpointOwnerId"] = state ? state.sourceEndpointOwnerId : undefined;
            resourceInputs["sourceEndpointPassword"] = state ? state.sourceEndpointPassword : undefined;
            resourceInputs["sourceEndpointPort"] = state ? state.sourceEndpointPort : undefined;
            resourceInputs["sourceEndpointRegion"] = state ? state.sourceEndpointRegion : undefined;
            resourceInputs["sourceEndpointRole"] = state ? state.sourceEndpointRole : undefined;
            resourceInputs["sourceEndpointUserName"] = state ? state.sourceEndpointUserName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["structureInitialization"] = state ? state.structureInitialization : undefined;
        } else {
            const args = argsOrState as MigrationJobArgs | undefined;
            if ((!args || args.dataInitialization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataInitialization'");
            }
            if ((!args || args.dataSynchronization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataSynchronization'");
            }
            if ((!args || args.dbList === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbList'");
            }
            if ((!args || args.destinationEndpointEngineName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationEndpointEngineName'");
            }
            if ((!args || args.destinationEndpointInstanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationEndpointInstanceType'");
            }
            if ((!args || args.dtsInstanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dtsInstanceId'");
            }
            if ((!args || args.sourceEndpointEngineName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceEndpointEngineName'");
            }
            if ((!args || args.sourceEndpointInstanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceEndpointInstanceType'");
            }
            if ((!args || args.structureInitialization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'structureInitialization'");
            }
            resourceInputs["checkpoint"] = args ? args.checkpoint : undefined;
            resourceInputs["dataInitialization"] = args ? args.dataInitialization : undefined;
            resourceInputs["dataSynchronization"] = args ? args.dataSynchronization : undefined;
            resourceInputs["dbList"] = args ? args.dbList : undefined;
            resourceInputs["destinationEndpointDatabaseName"] = args ? args.destinationEndpointDatabaseName : undefined;
            resourceInputs["destinationEndpointEngineName"] = args ? args.destinationEndpointEngineName : undefined;
            resourceInputs["destinationEndpointInstanceId"] = args ? args.destinationEndpointInstanceId : undefined;
            resourceInputs["destinationEndpointInstanceType"] = args ? args.destinationEndpointInstanceType : undefined;
            resourceInputs["destinationEndpointIp"] = args ? args.destinationEndpointIp : undefined;
            resourceInputs["destinationEndpointOracleSid"] = args ? args.destinationEndpointOracleSid : undefined;
            resourceInputs["destinationEndpointPassword"] = args ? args.destinationEndpointPassword : undefined;
            resourceInputs["destinationEndpointPort"] = args ? args.destinationEndpointPort : undefined;
            resourceInputs["destinationEndpointRegion"] = args ? args.destinationEndpointRegion : undefined;
            resourceInputs["destinationEndpointUserName"] = args ? args.destinationEndpointUserName : undefined;
            resourceInputs["dtsInstanceId"] = args ? args.dtsInstanceId : undefined;
            resourceInputs["dtsJobName"] = args ? args.dtsJobName : undefined;
            resourceInputs["instanceClass"] = args ? args.instanceClass : undefined;
            resourceInputs["sourceEndpointDatabaseName"] = args ? args.sourceEndpointDatabaseName : undefined;
            resourceInputs["sourceEndpointEngineName"] = args ? args.sourceEndpointEngineName : undefined;
            resourceInputs["sourceEndpointInstanceId"] = args ? args.sourceEndpointInstanceId : undefined;
            resourceInputs["sourceEndpointInstanceType"] = args ? args.sourceEndpointInstanceType : undefined;
            resourceInputs["sourceEndpointIp"] = args ? args.sourceEndpointIp : undefined;
            resourceInputs["sourceEndpointOracleSid"] = args ? args.sourceEndpointOracleSid : undefined;
            resourceInputs["sourceEndpointOwnerId"] = args ? args.sourceEndpointOwnerId : undefined;
            resourceInputs["sourceEndpointPassword"] = args ? args.sourceEndpointPassword : undefined;
            resourceInputs["sourceEndpointPort"] = args ? args.sourceEndpointPort : undefined;
            resourceInputs["sourceEndpointRegion"] = args ? args.sourceEndpointRegion : undefined;
            resourceInputs["sourceEndpointRole"] = args ? args.sourceEndpointRole : undefined;
            resourceInputs["sourceEndpointUserName"] = args ? args.sourceEndpointUserName : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["structureInitialization"] = args ? args.structureInitialization : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MigrationJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MigrationJob resources.
 */
export interface MigrationJobState {
    /**
     * Start time in Unix timestamp format.
     */
    checkpoint?: pulumi.Input<string>;
    /**
     * Whether to execute DTS supports schema migration.
     */
    dataInitialization?: pulumi.Input<boolean>;
    /**
     * Whether to perform incremental data migration.
     */
    dataSynchronization?: pulumi.Input<boolean>;
    /**
     * Migration object, in the format of JSON strings. For detailed definition instructions, please refer to [the description of migration, migration or subscription objects](https://help.aliyun.com/document_detail/209545.html).
     */
    dbList?: pulumi.Input<string>;
    /**
     * The name of migrate the database.
     */
    destinationEndpointDatabaseName?: pulumi.Input<string>;
    /**
     * The type of destination database. Valid values: `ADS`, `ADB30`, `AS400`, `DATAHUB`, `DB2`, `GREENPLUM`, `KAFKA`, `MONGODB`, `MSSQL`, `MySQL`, `ORACLE`, `PolarDB`, `POLARDBX20`, `POLARDB_O`, `PostgreSQL`.
     */
    destinationEndpointEngineName?: pulumi.Input<string>;
    /**
     * The ID of destination instance.
     */
    destinationEndpointInstanceId?: pulumi.Input<string>;
    /**
     * The type of destination instance. Valid values: `ADS`, `CEN`, `DATAHUB`, `DG`, `ECS`, `EXPRESS`, `GREENPLUM`, `MONGODB`, `OTHER`, `PolarDB`, `POLARDBX20`, `RDS`.
     */
    destinationEndpointInstanceType?: pulumi.Input<string>;
    /**
     * The ip of source endpoint.
     */
    destinationEndpointIp?: pulumi.Input<string>;
    /**
     * The SID of Oracle database.
     */
    destinationEndpointOracleSid?: pulumi.Input<string>;
    /**
     * The password of database account.
     */
    destinationEndpointPassword?: pulumi.Input<string>;
    /**
     * The port of source endpoint.
     */
    destinationEndpointPort?: pulumi.Input<string>;
    /**
     * The region of destination instance.
     */
    destinationEndpointRegion?: pulumi.Input<string>;
    /**
     * The username of database account.
     */
    destinationEndpointUserName?: pulumi.Input<string>;
    /**
     * The Migration instance ID. The ID of `alicloud.dts.MigrationInstance`.
     */
    dtsInstanceId?: pulumi.Input<string>;
    /**
     * The name of migration job.
     */
    dtsJobName?: pulumi.Input<string>;
    /**
     * The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`, `xxlarge`.
     */
    instanceClass?: pulumi.Input<string>;
    /**
     * The name of migrate the database.
     */
    sourceEndpointDatabaseName?: pulumi.Input<string>;
    /**
     * The type of source database. Valid values: `AS400`, `DB2`, `DMSPOLARDB`, `HBASE`, `MONGODB`, `MSSQL`, `MySQL`, `ORACLE`, `PolarDB`, `POLARDBX20`, `POLARDB_O`, `POSTGRESQL`, `TERADATA`.
     */
    sourceEndpointEngineName?: pulumi.Input<string>;
    /**
     * The ID of source instance.
     */
    sourceEndpointInstanceId?: pulumi.Input<string>;
    /**
     * The type of source instance. Valid values: `CEN`, `DG`, `DISTRIBUTED_DMSLOGICDB`, `ECS`, `EXPRESS`, `MONGODB`, `OTHER`, `PolarDB`, `POLARDBX20`, `RDS`.
     */
    sourceEndpointInstanceType?: pulumi.Input<string>;
    /**
     * The ip of source endpoint.
     */
    sourceEndpointIp?: pulumi.Input<string>;
    /**
     * The SID of Oracle database.
     */
    sourceEndpointOracleSid?: pulumi.Input<string>;
    /**
     * The Alibaba Cloud account ID to which the source instance belongs.
     */
    sourceEndpointOwnerId?: pulumi.Input<string>;
    /**
     * The password of database account.
     */
    sourceEndpointPassword?: pulumi.Input<string>;
    /**
     * The port of source endpoint.
     */
    sourceEndpointPort?: pulumi.Input<string>;
    /**
     * The region of source instance.
     */
    sourceEndpointRegion?: pulumi.Input<string>;
    /**
     * The name of the role configured for the cloud account to which the source instance belongs.
     */
    sourceEndpointRole?: pulumi.Input<string>;
    /**
     * The username of database account.
     */
    sourceEndpointUserName?: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values: `Migrating`, `Suspending`. You can suspend the task by specifying `Suspending` and start the task by specifying `Migrating`.
     */
    status?: pulumi.Input<string>;
    /**
     * Whether to perform a database table structure to migrate.
     */
    structureInitialization?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a MigrationJob resource.
 */
export interface MigrationJobArgs {
    /**
     * Start time in Unix timestamp format.
     */
    checkpoint?: pulumi.Input<string>;
    /**
     * Whether to execute DTS supports schema migration.
     */
    dataInitialization: pulumi.Input<boolean>;
    /**
     * Whether to perform incremental data migration.
     */
    dataSynchronization: pulumi.Input<boolean>;
    /**
     * Migration object, in the format of JSON strings. For detailed definition instructions, please refer to [the description of migration, migration or subscription objects](https://help.aliyun.com/document_detail/209545.html).
     */
    dbList: pulumi.Input<string>;
    /**
     * The name of migrate the database.
     */
    destinationEndpointDatabaseName?: pulumi.Input<string>;
    /**
     * The type of destination database. Valid values: `ADS`, `ADB30`, `AS400`, `DATAHUB`, `DB2`, `GREENPLUM`, `KAFKA`, `MONGODB`, `MSSQL`, `MySQL`, `ORACLE`, `PolarDB`, `POLARDBX20`, `POLARDB_O`, `PostgreSQL`.
     */
    destinationEndpointEngineName: pulumi.Input<string>;
    /**
     * The ID of destination instance.
     */
    destinationEndpointInstanceId?: pulumi.Input<string>;
    /**
     * The type of destination instance. Valid values: `ADS`, `CEN`, `DATAHUB`, `DG`, `ECS`, `EXPRESS`, `GREENPLUM`, `MONGODB`, `OTHER`, `PolarDB`, `POLARDBX20`, `RDS`.
     */
    destinationEndpointInstanceType: pulumi.Input<string>;
    /**
     * The ip of source endpoint.
     */
    destinationEndpointIp?: pulumi.Input<string>;
    /**
     * The SID of Oracle database.
     */
    destinationEndpointOracleSid?: pulumi.Input<string>;
    /**
     * The password of database account.
     */
    destinationEndpointPassword?: pulumi.Input<string>;
    /**
     * The port of source endpoint.
     */
    destinationEndpointPort?: pulumi.Input<string>;
    /**
     * The region of destination instance.
     */
    destinationEndpointRegion?: pulumi.Input<string>;
    /**
     * The username of database account.
     */
    destinationEndpointUserName?: pulumi.Input<string>;
    /**
     * The Migration instance ID. The ID of `alicloud.dts.MigrationInstance`.
     */
    dtsInstanceId: pulumi.Input<string>;
    /**
     * The name of migration job.
     */
    dtsJobName?: pulumi.Input<string>;
    /**
     * The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`, `xxlarge`.
     */
    instanceClass?: pulumi.Input<string>;
    /**
     * The name of migrate the database.
     */
    sourceEndpointDatabaseName?: pulumi.Input<string>;
    /**
     * The type of source database. Valid values: `AS400`, `DB2`, `DMSPOLARDB`, `HBASE`, `MONGODB`, `MSSQL`, `MySQL`, `ORACLE`, `PolarDB`, `POLARDBX20`, `POLARDB_O`, `POSTGRESQL`, `TERADATA`.
     */
    sourceEndpointEngineName: pulumi.Input<string>;
    /**
     * The ID of source instance.
     */
    sourceEndpointInstanceId?: pulumi.Input<string>;
    /**
     * The type of source instance. Valid values: `CEN`, `DG`, `DISTRIBUTED_DMSLOGICDB`, `ECS`, `EXPRESS`, `MONGODB`, `OTHER`, `PolarDB`, `POLARDBX20`, `RDS`.
     */
    sourceEndpointInstanceType: pulumi.Input<string>;
    /**
     * The ip of source endpoint.
     */
    sourceEndpointIp?: pulumi.Input<string>;
    /**
     * The SID of Oracle database.
     */
    sourceEndpointOracleSid?: pulumi.Input<string>;
    /**
     * The Alibaba Cloud account ID to which the source instance belongs.
     */
    sourceEndpointOwnerId?: pulumi.Input<string>;
    /**
     * The password of database account.
     */
    sourceEndpointPassword?: pulumi.Input<string>;
    /**
     * The port of source endpoint.
     */
    sourceEndpointPort?: pulumi.Input<string>;
    /**
     * The region of source instance.
     */
    sourceEndpointRegion?: pulumi.Input<string>;
    /**
     * The name of the role configured for the cloud account to which the source instance belongs.
     */
    sourceEndpointRole?: pulumi.Input<string>;
    /**
     * The username of database account.
     */
    sourceEndpointUserName?: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values: `Migrating`, `Suspending`. You can suspend the task by specifying `Suspending` and start the task by specifying `Migrating`.
     */
    status?: pulumi.Input<string>;
    /**
     * Whether to perform a database table structure to migrate.
     */
    structureInitialization: pulumi.Input<boolean>;
}
