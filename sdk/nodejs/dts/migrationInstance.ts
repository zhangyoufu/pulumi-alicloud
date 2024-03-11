// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a DTS Migration Instance resource.
 *
 * For information about DTS Migration Instance and how to use it, see [What is Synchronization Instance](https://www.alibabacloud.com/help/en/doc-detail/208270.html).
 *
 * > **NOTE:** Available since v1.157.0.
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
 * const defaultRegions = alicloud.getRegions({
 *     current: true,
 * });
 * const defaultMigrationInstance = new alicloud.dts.MigrationInstance("defaultMigrationInstance", {
 *     paymentType: "PayAsYouGo",
 *     sourceEndpointEngineName: "MySQL",
 *     sourceEndpointRegion: defaultRegions.then(defaultRegions => defaultRegions.regions?.[0]?.id),
 *     destinationEndpointEngineName: "MySQL",
 *     destinationEndpointRegion: defaultRegions.then(defaultRegions => defaultRegions.regions?.[0]?.id),
 *     instanceClass: "small",
 *     syncArchitecture: "oneway",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * DTS Migration Instance can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:dts/migrationInstance:MigrationInstance example <id>
 * ```
 */
export class MigrationInstance extends pulumi.CustomResource {
    /**
     * Get an existing MigrationInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MigrationInstanceState, opts?: pulumi.CustomResourceOptions): MigrationInstance {
        return new MigrationInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dts/migrationInstance:MigrationInstance';

    /**
     * Returns true if the given object is an instance of MigrationInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MigrationInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MigrationInstance.__pulumiType;
    }

    /**
     * [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.
     */
    public readonly computeUnit!: pulumi.Output<number | undefined>;
    /**
     * The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when `sourceEndpointEngineName` equals `drds`.
     */
    public readonly databaseCount!: pulumi.Output<number | undefined>;
    /**
     * The type of destination engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
     */
    public readonly destinationEndpointEngineName!: pulumi.Output<string>;
    /**
     * The region of destination instance. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
     */
    public readonly destinationEndpointRegion!: pulumi.Output<string>;
    /**
     * The ID of the Migration Instance.
     */
    public /*out*/ readonly dtsInstanceId!: pulumi.Output<string>;
    /**
     * The instance class. Valid values: `large`, `medium`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
     */
    public readonly instanceClass!: pulumi.Output<string>;
    /**
     * The payment type of the resource. Valid values: `PayAsYouGo`.
     */
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * The type of source endpoint engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
     */
    public readonly sourceEndpointEngineName!: pulumi.Output<string>;
    /**
     * The region of source instance.
     */
    public readonly sourceEndpointRegion!: pulumi.Output<string>;
    /**
     * The status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The sync architecture. Valid values: `oneway`.
     */
    public readonly syncArchitecture!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a MigrationInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MigrationInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MigrationInstanceArgs | MigrationInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MigrationInstanceState | undefined;
            resourceInputs["computeUnit"] = state ? state.computeUnit : undefined;
            resourceInputs["databaseCount"] = state ? state.databaseCount : undefined;
            resourceInputs["destinationEndpointEngineName"] = state ? state.destinationEndpointEngineName : undefined;
            resourceInputs["destinationEndpointRegion"] = state ? state.destinationEndpointRegion : undefined;
            resourceInputs["dtsInstanceId"] = state ? state.dtsInstanceId : undefined;
            resourceInputs["instanceClass"] = state ? state.instanceClass : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["sourceEndpointEngineName"] = state ? state.sourceEndpointEngineName : undefined;
            resourceInputs["sourceEndpointRegion"] = state ? state.sourceEndpointRegion : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["syncArchitecture"] = state ? state.syncArchitecture : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as MigrationInstanceArgs | undefined;
            if ((!args || args.destinationEndpointEngineName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationEndpointEngineName'");
            }
            if ((!args || args.destinationEndpointRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationEndpointRegion'");
            }
            if ((!args || args.paymentType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paymentType'");
            }
            if ((!args || args.sourceEndpointEngineName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceEndpointEngineName'");
            }
            if ((!args || args.sourceEndpointRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceEndpointRegion'");
            }
            resourceInputs["computeUnit"] = args ? args.computeUnit : undefined;
            resourceInputs["databaseCount"] = args ? args.databaseCount : undefined;
            resourceInputs["destinationEndpointEngineName"] = args ? args.destinationEndpointEngineName : undefined;
            resourceInputs["destinationEndpointRegion"] = args ? args.destinationEndpointRegion : undefined;
            resourceInputs["instanceClass"] = args ? args.instanceClass : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["sourceEndpointEngineName"] = args ? args.sourceEndpointEngineName : undefined;
            resourceInputs["sourceEndpointRegion"] = args ? args.sourceEndpointRegion : undefined;
            resourceInputs["syncArchitecture"] = args ? args.syncArchitecture : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["dtsInstanceId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MigrationInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MigrationInstance resources.
 */
export interface MigrationInstanceState {
    /**
     * [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.
     */
    computeUnit?: pulumi.Input<number>;
    /**
     * The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when `sourceEndpointEngineName` equals `drds`.
     */
    databaseCount?: pulumi.Input<number>;
    /**
     * The type of destination engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
     */
    destinationEndpointEngineName?: pulumi.Input<string>;
    /**
     * The region of destination instance. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
     */
    destinationEndpointRegion?: pulumi.Input<string>;
    /**
     * The ID of the Migration Instance.
     */
    dtsInstanceId?: pulumi.Input<string>;
    /**
     * The instance class. Valid values: `large`, `medium`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
     */
    instanceClass?: pulumi.Input<string>;
    /**
     * The payment type of the resource. Valid values: `PayAsYouGo`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * The type of source endpoint engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
     */
    sourceEndpointEngineName?: pulumi.Input<string>;
    /**
     * The region of source instance.
     */
    sourceEndpointRegion?: pulumi.Input<string>;
    /**
     * The status.
     */
    status?: pulumi.Input<string>;
    /**
     * The sync architecture. Valid values: `oneway`.
     */
    syncArchitecture?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a MigrationInstance resource.
 */
export interface MigrationInstanceArgs {
    /**
     * [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.
     */
    computeUnit?: pulumi.Input<number>;
    /**
     * The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when `sourceEndpointEngineName` equals `drds`.
     */
    databaseCount?: pulumi.Input<number>;
    /**
     * The type of destination engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
     */
    destinationEndpointEngineName: pulumi.Input<string>;
    /**
     * The region of destination instance. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
     */
    destinationEndpointRegion: pulumi.Input<string>;
    /**
     * The instance class. Valid values: `large`, `medium`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
     */
    instanceClass?: pulumi.Input<string>;
    /**
     * The payment type of the resource. Valid values: `PayAsYouGo`.
     */
    paymentType: pulumi.Input<string>;
    /**
     * The type of source endpoint engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
     */
    sourceEndpointEngineName: pulumi.Input<string>;
    /**
     * The region of source instance.
     */
    sourceEndpointRegion: pulumi.Input<string>;
    /**
     * The sync architecture. Valid values: `oneway`.
     */
    syncArchitecture?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
