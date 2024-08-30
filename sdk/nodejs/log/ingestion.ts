// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Log service ingestion, this service provides the function of importing logs of various data sources(OSS, MaxCompute) into logstore.
 * [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
 *
 * > **NOTE:** Available in 1.161.0+
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const _default = new random.index.Integer("default", {
 *     max: 99999,
 *     min: 10000,
 * });
 * const example = new alicloud.log.Project("example", {
 *     projectName: `terraform-example-${_default.result}`,
 *     description: "terraform-example",
 *     tags: {
 *         Created: "TF",
 *         For: "example",
 *     },
 * });
 * const exampleStore = new alicloud.log.Store("example", {
 *     projectName: example.projectName,
 *     logstoreName: "example-store",
 *     retentionPeriod: 3650,
 *     shardCount: 3,
 *     autoSplit: true,
 *     maxSplitShardCount: 60,
 *     appendMeta: true,
 * });
 * const exampleIngestion = new alicloud.log.Ingestion("example", {
 *     project: example.projectName,
 *     logstore: exampleStore.logstoreName,
 *     ingestionName: "terraform-example",
 *     displayName: "terraform-example",
 *     description: "terraform-example",
 *     interval: "30m",
 *     runImmediately: true,
 *     timeZone: "+0800",
 *     source: `        {
 *           "bucket": "bucket_name",
 *           "compressionCodec": "none",
 *           "encoding": "UTF-8",
 *           "endpoint": "oss-cn-hangzhou-internal.aliyuncs.com",
 *           "format": {
 *             "escapeChar": "\\\\",
 *             "fieldDelimiter": ",",
 *             "fieldNames": [],
 *             "firstRowAsHeader": true,
 *             "maxLines": 1,
 *             "quoteChar": "\\"",
 *             "skipLeadingRows": 0,
 *             "timeField": "",
 *             "type": "DelimitedText"
 *           },
 *           "pattern": "",
 *           "prefix": "test-prefix/",
 *           "restoreObjectEnabled": false,
 *           "roleARN": "acs:ram::1049446484210612:role/aliyunlogimportossrole",
 *           "type": "AliyunOSS"
 *         }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * Log ingestion can be imported using the id or name, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:log/ingestion:Ingestion example tf-log-project:tf-log-logstore:ingestion_name
 * ```
 */
export class Ingestion extends pulumi.CustomResource {
    /**
     * Get an existing Ingestion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IngestionState, opts?: pulumi.CustomResourceOptions): Ingestion {
        return new Ingestion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:log/ingestion:Ingestion';

    /**
     * Returns true if the given object is an instance of Ingestion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ingestion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ingestion.__pulumiType;
    }

    /**
     * Ingestion job description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name displayed on the web page.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Ingestion job name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
     */
    public readonly ingestionName!: pulumi.Output<string>;
    /**
     * Task execution interval, support minute `m`, hour `h`, day `d`, for example 30 minutes `30m`.
     */
    public readonly interval!: pulumi.Output<string>;
    /**
     * The name of the target logstore.
     */
    public readonly logstore!: pulumi.Output<string>;
    /**
     * The name of the log project. It is the only in one Alicloud account.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Whether to run the ingestion job immediately, if false, wait for an interval before starting the ingestion.
     */
    public readonly runImmediately!: pulumi.Output<boolean>;
    /**
     * Data source and data format details. [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * Which time zone is the log time imported in, e.g. `+0800`.
     */
    public readonly timeZone!: pulumi.Output<string | undefined>;

    /**
     * Create a Ingestion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IngestionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IngestionArgs | IngestionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IngestionState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["ingestionName"] = state ? state.ingestionName : undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["logstore"] = state ? state.logstore : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["runImmediately"] = state ? state.runImmediately : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["timeZone"] = state ? state.timeZone : undefined;
        } else {
            const args = argsOrState as IngestionArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.ingestionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ingestionName'");
            }
            if ((!args || args.interval === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interval'");
            }
            if ((!args || args.logstore === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logstore'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.runImmediately === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runImmediately'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["ingestionName"] = args ? args.ingestionName : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["logstore"] = args ? args.logstore : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["runImmediately"] = args ? args.runImmediately : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["timeZone"] = args ? args.timeZone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ingestion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ingestion resources.
 */
export interface IngestionState {
    /**
     * Ingestion job description.
     */
    description?: pulumi.Input<string>;
    /**
     * The name displayed on the web page.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Ingestion job name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
     */
    ingestionName?: pulumi.Input<string>;
    /**
     * Task execution interval, support minute `m`, hour `h`, day `d`, for example 30 minutes `30m`.
     */
    interval?: pulumi.Input<string>;
    /**
     * The name of the target logstore.
     */
    logstore?: pulumi.Input<string>;
    /**
     * The name of the log project. It is the only in one Alicloud account.
     */
    project?: pulumi.Input<string>;
    /**
     * Whether to run the ingestion job immediately, if false, wait for an interval before starting the ingestion.
     */
    runImmediately?: pulumi.Input<boolean>;
    /**
     * Data source and data format details. [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
     */
    source?: pulumi.Input<string>;
    /**
     * Which time zone is the log time imported in, e.g. `+0800`.
     */
    timeZone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ingestion resource.
 */
export interface IngestionArgs {
    /**
     * Ingestion job description.
     */
    description?: pulumi.Input<string>;
    /**
     * The name displayed on the web page.
     */
    displayName: pulumi.Input<string>;
    /**
     * Ingestion job name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
     */
    ingestionName: pulumi.Input<string>;
    /**
     * Task execution interval, support minute `m`, hour `h`, day `d`, for example 30 minutes `30m`.
     */
    interval: pulumi.Input<string>;
    /**
     * The name of the target logstore.
     */
    logstore: pulumi.Input<string>;
    /**
     * The name of the log project. It is the only in one Alicloud account.
     */
    project: pulumi.Input<string>;
    /**
     * Whether to run the ingestion job immediately, if false, wait for an interval before starting the ingestion.
     */
    runImmediately: pulumi.Input<boolean>;
    /**
     * Data source and data format details. [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
     */
    source: pulumi.Input<string>;
    /**
     * Which time zone is the log time imported in, e.g. `+0800`.
     */
    timeZone?: pulumi.Input<string>;
}
