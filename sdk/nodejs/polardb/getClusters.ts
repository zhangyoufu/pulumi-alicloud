// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The `alicloud.polardb.getClusters` data source provides a collection of PolarDB clusters available in Alibaba Cloud account.
 * Filters support regular expression for the cluster description, searches by tags, and other filters which are listed below.
 *
 * > **NOTE:** Available since v1.66.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const this = alicloud.polardb.getNodeClasses({
 *     dbType: "MySQL",
 *     dbVersion: "8.0",
 *     payType: "PostPaid",
 *     category: "Normal",
 * });
 * const _default = new alicloud.vpc.Network("default", {
 *     vpcName: "terraform-example",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     vpcId: _default.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: _this.then(_this => _this.classes?.[0]?.zoneId),
 *     vswitchName: "terraform-example",
 * });
 * const cluster = new alicloud.polardb.Cluster("cluster", {
 *     dbType: "MySQL",
 *     dbVersion: "8.0",
 *     payType: "PostPaid",
 *     dbNodeCount: 2,
 *     dbNodeClass: _this.then(_this => _this.classes?.[0]?.supportedEngines?.[0]?.availableResources?.[0]?.dbNodeClass),
 *     vswitchId: defaultSwitch.id,
 * });
 * const polardbClustersDs = alicloud.polardb.getClustersOutput({
 *     descriptionRegex: cluster.id,
 *     status: "Running",
 * });
 * export const firstPolardbClusterId = polardbClustersDs.apply(polardbClustersDs => polardbClustersDs.clusters?.[0]?.id);
 * ```
 */
export function getClusters(args?: GetClustersArgs, opts?: pulumi.InvokeOptions): Promise<GetClustersResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:polardb/getClusters:getClusters", {
        "dbType": args.dbType,
        "descriptionRegex": args.descriptionRegex,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "status": args.status,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusters.
 */
export interface GetClustersArgs {
    /**
     * Database type. Options are `MySQL`, `Oracle` and `PostgreSQL`. If no value is specified, all types are returned.
     */
    dbType?: string;
    /**
     * A regex string to filter results by cluster description.
     */
    descriptionRegex?: string;
    /**
     * A list of PolarDB cluster IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * status of the cluster.
     */
    status?: string;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getClusters.
 */
export interface GetClustersResult {
    /**
     * A list of PolarDB clusters. Each element contains the following attributes:
     */
    readonly clusters: outputs.polardb.GetClustersCluster[];
    /**
     * `Primary` for primary cluster, `ReadOnly` for read-only cluster, `Guard` for disaster recovery cluster, and `Temp` for temporary cluster.
     */
    readonly dbType?: string;
    readonly descriptionRegex?: string;
    /**
     * A list of RDS cluster descriptions.
     */
    readonly descriptions: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of RDS cluster IDs.
     */
    readonly ids: string[];
    readonly outputFile?: string;
    /**
     * Status of the cluster.
     */
    readonly status?: string;
    readonly tags?: {[key: string]: string};
}
/**
 * The `alicloud.polardb.getClusters` data source provides a collection of PolarDB clusters available in Alibaba Cloud account.
 * Filters support regular expression for the cluster description, searches by tags, and other filters which are listed below.
 *
 * > **NOTE:** Available since v1.66.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const this = alicloud.polardb.getNodeClasses({
 *     dbType: "MySQL",
 *     dbVersion: "8.0",
 *     payType: "PostPaid",
 *     category: "Normal",
 * });
 * const _default = new alicloud.vpc.Network("default", {
 *     vpcName: "terraform-example",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     vpcId: _default.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: _this.then(_this => _this.classes?.[0]?.zoneId),
 *     vswitchName: "terraform-example",
 * });
 * const cluster = new alicloud.polardb.Cluster("cluster", {
 *     dbType: "MySQL",
 *     dbVersion: "8.0",
 *     payType: "PostPaid",
 *     dbNodeCount: 2,
 *     dbNodeClass: _this.then(_this => _this.classes?.[0]?.supportedEngines?.[0]?.availableResources?.[0]?.dbNodeClass),
 *     vswitchId: defaultSwitch.id,
 * });
 * const polardbClustersDs = alicloud.polardb.getClustersOutput({
 *     descriptionRegex: cluster.id,
 *     status: "Running",
 * });
 * export const firstPolardbClusterId = polardbClustersDs.apply(polardbClustersDs => polardbClustersDs.clusters?.[0]?.id);
 * ```
 */
export function getClustersOutput(args?: GetClustersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClustersResult> {
    return pulumi.output(args).apply((a: any) => getClusters(a, opts))
}

/**
 * A collection of arguments for invoking getClusters.
 */
export interface GetClustersOutputArgs {
    /**
     * Database type. Options are `MySQL`, `Oracle` and `PostgreSQL`. If no value is specified, all types are returned.
     */
    dbType?: pulumi.Input<string>;
    /**
     * A regex string to filter results by cluster description.
     */
    descriptionRegex?: pulumi.Input<string>;
    /**
     * A list of PolarDB cluster IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * status of the cluster.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
