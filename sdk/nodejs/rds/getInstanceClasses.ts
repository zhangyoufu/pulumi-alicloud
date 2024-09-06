// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the RDS instance classes resource available info of Alibaba Cloud.
 *
 * > **NOTE:** Available in v1.46.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const resources = alicloud.rds.getInstanceClasses({
 *     instanceChargeType: "PostPaid",
 *     engine: "MySQL",
 *     engineVersion: "5.6",
 *     outputFile: "./classes.txt",
 * });
 * export const firstDbInstanceClass = resources.then(resources => resources.instanceClasses?.[0]?.instanceClass);
 * ```
 */
export function getInstanceClasses(args?: GetInstanceClassesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceClassesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:rds/getInstanceClasses:getInstanceClasses", {
        "category": args.category,
        "commodityCode": args.commodityCode,
        "dbInstanceClass": args.dbInstanceClass,
        "dbInstanceId": args.dbInstanceId,
        "dbInstanceStorageType": args.dbInstanceStorageType,
        "engine": args.engine,
        "engineVersion": args.engineVersion,
        "instanceChargeType": args.instanceChargeType,
        "multiZone": args.multiZone,
        "outputFile": args.outputFile,
        "sortedBy": args.sortedBy,
        "storageType": args.storageType,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceClasses.
 */
export interface GetInstanceClassesArgs {
    /**
     * DB Instance category. the value like [`Basic`, `HighAvailability`, `Finance`, `AlwaysOn`, `serverlessBasic`, `serverlessStandard`, `serverlessHa`, `cluster`], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).
     */
    category?: string;
    /**
     * The commodity code of the instance. Valid values:
     * * **bards**: The instance is a pay-as-you-go primary instance. This value is available on the China site (aliyun.com).
     * * **rds**: The instance is a subscription primary instance. This value is available on the China site (aliyun.com).
     * * **rords**: The instance is a pay-as-you-go read-only instance. This value is available on the China site (aliyun.com).
     * * **rds_rordspre_public_cn**: The instance is a subscription read-only instance. This value is available on the China site (aliyun.com).
     * * **bards_intl**: The instance is a pay-as-you-go primary instance. This value is available on the International site (alibabacloud.com).
     * * **rds_intl**: The instance is a subscription primary instance. This value is available on the International site (alibabacloud.com).
     * * **rords_intl**: The instance is a pay-as-you-go read-only instance. This value is available on the International site (alibabacloud.com).
     * * **rds_rordspre_public_intl**: The instance is a subscription read-only instance. This value is available on the International site (alibabacloud.com).
     * * **rds_serverless_public_cn**: The instance is a subscription serverless instance. This value is available on the China site (aliyun.com).
     * * **rds_serverless_public_intl**: The instance is a subscription serverless instance. This value is available on the International site (alibabacloud.com).
     *
     * > **NOTE**: The field `dbInstanceId` will be ignored when `commodityCode` is not a read-only type.
     */
    commodityCode?: string;
    /**
     * The DB instance class type by the user.
     */
    dbInstanceClass?: string;
    /**
     * The ID of the instance.
     *
     * > **NOTE**: The field `dbInstanceId` will be ignored when `commodityCode` is not a read-only type.
     */
    dbInstanceId?: string;
    /**
     * The DB instance storage space required by the user. Valid values: "cloudSsd", "localSsd", "cloudEssd", "cloudEssd2", "cloudEssd3".
     */
    dbInstanceStorageType?: string;
    /**
     * Database type. Valid values:"MySQL", "SQLServer", "PostgreSQL", "MariaDB". If not set, it will match all of engines.
     */
    engine?: string;
    /**
     * Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
     */
    engineVersion?: string;
    /**
     * Filter the results by charge type. Valid values: `PrePaid` and `PostPaid` and `Serverless`. Default to `PostPaid`.
     */
    instanceChargeType?: string;
    /**
     * Whether to show multi available zone. Default false to not show multi availability zone.
     */
    multiZone?: boolean;
    /**
     * File name where to save data source results (after running `pulumi up`).
     */
    outputFile?: string;
    sortedBy?: string;
    /**
     * It has been deprecated from version 1.134.0+ and using `dbInstanceStorageType` instead.
     */
    storageType?: string;
    /**
     * The Zone to launch the DB instance.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by getInstanceClasses.
 */
export interface GetInstanceClassesResult {
    readonly category?: string;
    readonly commodityCode?: string;
    readonly dbInstanceClass?: string;
    readonly dbInstanceId?: string;
    readonly dbInstanceStorageType?: string;
    readonly engine?: string;
    readonly engineVersion?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Available in 1.60.0+) A list of Rds instance class codes.
     */
    readonly ids: string[];
    readonly instanceChargeType?: string;
    /**
     * A list of Rds available resource. Each element contains the following attributes:
     */
    readonly instanceClasses: outputs.rds.GetInstanceClassesInstanceClass[];
    readonly multiZone?: boolean;
    readonly outputFile?: string;
    readonly sortedBy?: string;
    readonly storageType?: string;
    readonly zoneId?: string;
}
/**
 * This data source provides the RDS instance classes resource available info of Alibaba Cloud.
 *
 * > **NOTE:** Available in v1.46.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const resources = alicloud.rds.getInstanceClasses({
 *     instanceChargeType: "PostPaid",
 *     engine: "MySQL",
 *     engineVersion: "5.6",
 *     outputFile: "./classes.txt",
 * });
 * export const firstDbInstanceClass = resources.then(resources => resources.instanceClasses?.[0]?.instanceClass);
 * ```
 */
export function getInstanceClassesOutput(args?: GetInstanceClassesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceClassesResult> {
    return pulumi.output(args).apply((a: any) => getInstanceClasses(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceClasses.
 */
export interface GetInstanceClassesOutputArgs {
    /**
     * DB Instance category. the value like [`Basic`, `HighAvailability`, `Finance`, `AlwaysOn`, `serverlessBasic`, `serverlessStandard`, `serverlessHa`, `cluster`], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).
     */
    category?: pulumi.Input<string>;
    /**
     * The commodity code of the instance. Valid values:
     * * **bards**: The instance is a pay-as-you-go primary instance. This value is available on the China site (aliyun.com).
     * * **rds**: The instance is a subscription primary instance. This value is available on the China site (aliyun.com).
     * * **rords**: The instance is a pay-as-you-go read-only instance. This value is available on the China site (aliyun.com).
     * * **rds_rordspre_public_cn**: The instance is a subscription read-only instance. This value is available on the China site (aliyun.com).
     * * **bards_intl**: The instance is a pay-as-you-go primary instance. This value is available on the International site (alibabacloud.com).
     * * **rds_intl**: The instance is a subscription primary instance. This value is available on the International site (alibabacloud.com).
     * * **rords_intl**: The instance is a pay-as-you-go read-only instance. This value is available on the International site (alibabacloud.com).
     * * **rds_rordspre_public_intl**: The instance is a subscription read-only instance. This value is available on the International site (alibabacloud.com).
     * * **rds_serverless_public_cn**: The instance is a subscription serverless instance. This value is available on the China site (aliyun.com).
     * * **rds_serverless_public_intl**: The instance is a subscription serverless instance. This value is available on the International site (alibabacloud.com).
     *
     * > **NOTE**: The field `dbInstanceId` will be ignored when `commodityCode` is not a read-only type.
     */
    commodityCode?: pulumi.Input<string>;
    /**
     * The DB instance class type by the user.
     */
    dbInstanceClass?: pulumi.Input<string>;
    /**
     * The ID of the instance.
     *
     * > **NOTE**: The field `dbInstanceId` will be ignored when `commodityCode` is not a read-only type.
     */
    dbInstanceId?: pulumi.Input<string>;
    /**
     * The DB instance storage space required by the user. Valid values: "cloudSsd", "localSsd", "cloudEssd", "cloudEssd2", "cloudEssd3".
     */
    dbInstanceStorageType?: pulumi.Input<string>;
    /**
     * Database type. Valid values:"MySQL", "SQLServer", "PostgreSQL", "MariaDB". If not set, it will match all of engines.
     */
    engine?: pulumi.Input<string>;
    /**
     * Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Filter the results by charge type. Valid values: `PrePaid` and `PostPaid` and `Serverless`. Default to `PostPaid`.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * Whether to show multi available zone. Default false to not show multi availability zone.
     */
    multiZone?: pulumi.Input<boolean>;
    /**
     * File name where to save data source results (after running `pulumi up`).
     */
    outputFile?: pulumi.Input<string>;
    sortedBy?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.134.0+ and using `dbInstanceStorageType` instead.
     */
    storageType?: pulumi.Input<string>;
    /**
     * The Zone to launch the DB instance.
     */
    zoneId?: pulumi.Input<string>;
}
