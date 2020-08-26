// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the KVStore instance classes resource available info of Alibaba Cloud.
 *
 * > **NOTE:** Available in v1.49.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const resourcesZones = pulumi.output(alicloud.getZones({
 *     availableResourceCreation: "KVStore",
 * }, { async: true }));
 * const resourcesInstanceClasses = resourcesZones.apply(resourcesZones => alicloud.kvstore.getInstanceClasses({
 *     engine: "Redis",
 *     engineVersion: "5.0",
 *     instanceChargeType: "PrePaid",
 *     outputFile: "./classes.txt",
 *     zoneId: resourcesZones.zones[0].id,
 * }, { async: true }));
 *
 * export const firstKvstoreInstanceClass = resourcesInstanceClasses.instanceClasses;
 * ```
 */
export function getInstanceClasses(args: GetInstanceClassesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceClassesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:kvstore/getInstanceClasses:getInstanceClasses", {
        "architecture": args.architecture,
        "editionType": args.editionType,
        "engine": args.engine,
        "engineVersion": args.engineVersion,
        "instanceChargeType": args.instanceChargeType,
        "nodeType": args.nodeType,
        "outputFile": args.outputFile,
        "packageType": args.packageType,
        "performanceType": args.performanceType,
        "seriesType": args.seriesType,
        "shardNumber": args.shardNumber,
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
     * The KVStore instance system architecture required by the user. Valid values: `standard`, `cluster` and `rwsplit`.
     */
    readonly architecture?: string;
    /**
     * The KVStore instance edition type required by the user. Valid values: `Community` and `Enterprise`.
     */
    readonly editionType?: string;
    /**
     * Database type. Options are `Redis`, `Memcache`. Default to `Redis`.
     */
    readonly engine?: string;
    /**
     * Database version required by the user. Value options of Redis can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/60873.htm) `EngineVersion`. Value of Memcache should be empty.
     */
    readonly engineVersion?: string;
    /**
     * Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PrePaid`.
     */
    readonly instanceChargeType?: string;
    /**
     * The KVStore instance node type required by the user. Valid values: `double`, `single`, `readone`, `readthree` and `readfive`.
     */
    readonly nodeType?: string;
    readonly outputFile?: string;
    /**
     * It has been deprecated from 1.68.0.
     *
     * @deprecated The parameter 'package_type' has been deprecated from 1.68.0.
     */
    readonly packageType?: string;
    /**
     * It has been deprecated from 1.68.0.
     *
     * @deprecated The parameter 'performance_type' has been deprecated from 1.68.0.
     */
    readonly performanceType?: string;
    /**
     * The KVStore instance series type required by the user. Valid values: `enhancedPerformanceType` and `hybridStorage`.
     */
    readonly seriesType?: string;
    /**
     * The number of shard.Valid values: `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`.
     */
    readonly shardNumber?: number;
    readonly sortedBy?: string;
    /**
     * It has been deprecated from 1.68.0.
     *
     * @deprecated The parameter 'storage_type' has been deprecated from 1.68.0.
     */
    readonly storageType?: string;
    /**
     * The Zone to launch the KVStore instance.
     */
    readonly zoneId: string;
}

/**
 * A collection of values returned by getInstanceClasses.
 */
export interface GetInstanceClassesResult {
    readonly architecture?: string;
    /**
     * A list of KVStore available instance classes when the `sortedBy` is "Price". include:
     */
    readonly classes: outputs.kvstore.GetInstanceClassesClass[];
    readonly editionType?: string;
    readonly engine?: string;
    readonly engineVersion?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceChargeType?: string;
    /**
     * A list of KVStore available instance classes.
     */
    readonly instanceClasses: string[];
    readonly nodeType?: string;
    readonly outputFile?: string;
    /**
     * @deprecated The parameter 'package_type' has been deprecated from 1.68.0.
     */
    readonly packageType?: string;
    /**
     * @deprecated The parameter 'performance_type' has been deprecated from 1.68.0.
     */
    readonly performanceType?: string;
    readonly seriesType?: string;
    readonly shardNumber?: number;
    readonly sortedBy?: string;
    /**
     * @deprecated The parameter 'storage_type' has been deprecated from 1.68.0.
     */
    readonly storageType?: string;
    readonly zoneId: string;
}
