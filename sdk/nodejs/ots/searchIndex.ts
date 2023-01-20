// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an OTS search index resource.
 *
 * For information about OTS search index and how to use it, see [Search index overview](https://www.alibabacloud.com/help/en/tablestore/latest/search-index-overview).
 *
 * > **NOTE:** Available in v1.187.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraformtest";
 * const instance1 = new alicloud.ots.Instance("instance1", {
 *     description: name,
 *     accessedBy: "Any",
 *     tags: {
 *         Created: "TF",
 *         For: "acceptance test",
 *     },
 * });
 * const table1 = new alicloud.ots.Table("table1", {
 *     instanceName: instance1.name,
 *     tableName: name,
 *     primaryKeys: [
 *         {
 *             name: "pk1",
 *             type: "Integer",
 *         },
 *         {
 *             name: "pk2",
 *             type: "String",
 *         },
 *     ],
 *     definedColumns: [
 *         {
 *             name: "col1",
 *             type: "String",
 *         },
 *         {
 *             name: "col2",
 *             type: "Integer",
 *         },
 *     ],
 *     timeToLive: -1,
 *     maxVersion: 1,
 *     deviationCellVersionInSec: "1",
 * });
 * const _default = new alicloud.ots.SearchIndex("default", {
 *     instanceName: instance1.name,
 *     tableName: table1.tableName,
 *     indexName: name,
 *     timeToLive: -1,
 *     schemas: [{
 *         fieldSchemas: [
 *             {
 *                 fieldName: "col1",
 *                 fieldType: "Text",
 *                 isArray: false,
 *                 index: true,
 *                 analyzer: "Split",
 *                 store: true,
 *             },
 *             {
 *                 fieldName: "col2",
 *                 fieldType: "Long",
 *                 enableSortAndAgg: true,
 *             },
 *             {
 *                 fieldName: "pk1",
 *                 fieldType: "Long",
 *             },
 *             {
 *                 fieldName: "pk2",
 *                 fieldType: "Text",
 *             },
 *         ],
 *         indexSettings: [{
 *             routingFields: [
 *                 "pk1",
 *                 "pk2",
 *             ],
 *         }],
 *         indexSorts: [{
 *             sorters: [
 *                 {
 *                     sorterType: "PrimaryKeySort",
 *                     order: "Asc",
 *                 },
 *                 {
 *                     sorterType: "FieldSort",
 *                     order: "Desc",
 *                     fieldName: "col2",
 *                     mode: "Max",
 *                 },
 *             ],
 *         }],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * OTS search index can be imported using id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ots/searchIndex:SearchIndex index1 <instance_name>:<table_name>:<index_name>:<index_type>
 * ```
 */
export class SearchIndex extends pulumi.CustomResource {
    /**
     * Get an existing SearchIndex resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SearchIndexState, opts?: pulumi.CustomResourceOptions): SearchIndex {
        return new SearchIndex(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ots/searchIndex:SearchIndex';

    /**
     * Returns true if the given object is an instance of SearchIndex.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SearchIndex {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SearchIndex.__pulumiType;
    }

    /**
     * The search index create time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<number>;
    /**
     * The timestamp for sync phase.
     */
    public /*out*/ readonly currentSyncTimestamp!: pulumi.Output<number>;
    /**
     * The index id of the search index which could not be changed.
     */
    public /*out*/ readonly indexId!: pulumi.Output<string>;
    /**
     * The index name of the OTS Table. If changed, a new index would be created.
     */
    public readonly indexName!: pulumi.Output<string>;
    /**
     * The name of the OTS instance in which table will located.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * The schema of the search index. If changed, a new index would be created.
     */
    public readonly schemas!: pulumi.Output<outputs.ots.SearchIndexSchema[]>;
    /**
     * The search index sync phase. possible values: `Full`, `Incr`.
     */
    public /*out*/ readonly syncPhase!: pulumi.Output<string>;
    /**
     * The name of the OTS table. If changed, a new table would be created.
     */
    public readonly tableName!: pulumi.Output<string>;
    /**
     * The index type of the OTS Table. Specifies the retention period of data in the search index. Unit: seconds. Default value: -1.
     * If the retention period exceeds the TTL value, OTS automatically deletes expired data.
     */
    public readonly timeToLive!: pulumi.Output<number | undefined>;

    /**
     * Create a SearchIndex resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SearchIndexArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SearchIndexArgs | SearchIndexState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SearchIndexState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["currentSyncTimestamp"] = state ? state.currentSyncTimestamp : undefined;
            resourceInputs["indexId"] = state ? state.indexId : undefined;
            resourceInputs["indexName"] = state ? state.indexName : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["schemas"] = state ? state.schemas : undefined;
            resourceInputs["syncPhase"] = state ? state.syncPhase : undefined;
            resourceInputs["tableName"] = state ? state.tableName : undefined;
            resourceInputs["timeToLive"] = state ? state.timeToLive : undefined;
        } else {
            const args = argsOrState as SearchIndexArgs | undefined;
            if ((!args || args.indexName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'indexName'");
            }
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            if ((!args || args.schemas === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schemas'");
            }
            if ((!args || args.tableName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tableName'");
            }
            resourceInputs["indexName"] = args ? args.indexName : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["schemas"] = args ? args.schemas : undefined;
            resourceInputs["tableName"] = args ? args.tableName : undefined;
            resourceInputs["timeToLive"] = args ? args.timeToLive : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["currentSyncTimestamp"] = undefined /*out*/;
            resourceInputs["indexId"] = undefined /*out*/;
            resourceInputs["syncPhase"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SearchIndex.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SearchIndex resources.
 */
export interface SearchIndexState {
    /**
     * The search index create time.
     */
    createTime?: pulumi.Input<number>;
    /**
     * The timestamp for sync phase.
     */
    currentSyncTimestamp?: pulumi.Input<number>;
    /**
     * The index id of the search index which could not be changed.
     */
    indexId?: pulumi.Input<string>;
    /**
     * The index name of the OTS Table. If changed, a new index would be created.
     */
    indexName?: pulumi.Input<string>;
    /**
     * The name of the OTS instance in which table will located.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The schema of the search index. If changed, a new index would be created.
     */
    schemas?: pulumi.Input<pulumi.Input<inputs.ots.SearchIndexSchema>[]>;
    /**
     * The search index sync phase. possible values: `Full`, `Incr`.
     */
    syncPhase?: pulumi.Input<string>;
    /**
     * The name of the OTS table. If changed, a new table would be created.
     */
    tableName?: pulumi.Input<string>;
    /**
     * The index type of the OTS Table. Specifies the retention period of data in the search index. Unit: seconds. Default value: -1.
     * If the retention period exceeds the TTL value, OTS automatically deletes expired data.
     */
    timeToLive?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SearchIndex resource.
 */
export interface SearchIndexArgs {
    /**
     * The index name of the OTS Table. If changed, a new index would be created.
     */
    indexName: pulumi.Input<string>;
    /**
     * The name of the OTS instance in which table will located.
     */
    instanceName: pulumi.Input<string>;
    /**
     * The schema of the search index. If changed, a new index would be created.
     */
    schemas: pulumi.Input<pulumi.Input<inputs.ots.SearchIndexSchema>[]>;
    /**
     * The name of the OTS table. If changed, a new table would be created.
     */
    tableName: pulumi.Input<string>;
    /**
     * The index type of the OTS Table. Specifies the retention period of data in the search index. Unit: seconds. Default value: -1.
     * If the retention period exceeds the TTL value, OTS automatically deletes expired data.
     */
    timeToLive?: pulumi.Input<number>;
}
