// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Log service data delivery management, this service provides the function of delivering data in logstore to oss product storage.
 * [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/43724.htm).
 *
 * > **NOTE:** Available in 1.121.0+
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const _default = new random.RandomInteger("default", {
 *     max: 99999,
 *     min: 10000,
 * });
 * const exampleProject = new alicloud.log.Project("exampleProject", {
 *     description: "terraform-example",
 *     tags: {
 *         Created: "TF",
 *         For: "example",
 *     },
 * });
 * const exampleStore = new alicloud.log.Store("exampleStore", {
 *     project: exampleProject.name,
 *     retentionPeriod: 3650,
 *     autoSplit: true,
 *     maxSplitShardCount: 60,
 *     appendMeta: true,
 * });
 * const exampleOssShipper = new alicloud.log.OssShipper("exampleOssShipper", {
 *     projectName: exampleProject.name,
 *     logstoreName: exampleStore.name,
 *     shipperName: "terraform-example",
 *     ossBucket: "example_bucket",
 *     ossPrefix: "root",
 *     bufferInterval: 300,
 *     bufferSize: 250,
 *     compressType: "none",
 *     pathFormat: "%Y/%m/%d/%H/%M",
 *     format: "json",
 *     jsonEnableTag: true,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Log oss shipper can be imported using the id or name, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:log/ossShipper:OssShipper example tf-log-project:tf-log-logstore:tf-log-shipper
 * ```
 */
export class OssShipper extends pulumi.CustomResource {
    /**
     * Get an existing OssShipper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OssShipperState, opts?: pulumi.CustomResourceOptions): OssShipper {
        return new OssShipper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:log/ossShipper:OssShipper';

    /**
     * Returns true if the given object is an instance of OssShipper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OssShipper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OssShipper.__pulumiType;
    }

    /**
     * How often is it delivered every interval.
     */
    public readonly bufferInterval!: pulumi.Output<number>;
    /**
     * Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
     */
    public readonly bufferSize!: pulumi.Output<number>;
    /**
     * OSS data storage compression method, support: none, snappy. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
     */
    public readonly compressType!: pulumi.Output<string | undefined>;
    public readonly csvConfigColumns!: pulumi.Output<string[] | undefined>;
    public readonly csvConfigDelimiter!: pulumi.Output<string | undefined>;
    public readonly csvConfigHeader!: pulumi.Output<boolean | undefined>;
    public readonly csvConfigLinefeed!: pulumi.Output<string | undefined>;
    public readonly csvConfigNullidentifier!: pulumi.Output<string | undefined>;
    public readonly csvConfigQuote!: pulumi.Output<string | undefined>;
    /**
     * Storage format, only supports three types: `json`, `parquet`, `csv`.
     * **According to the different format, please select the following parameters**
     * - format = `json`
     * `jsonEnableTag` - (Optional) Whether to deliver the label.
     * - format = `csv`
     * `csvConfigDelimiter` - (Optional) Separator configuration in csv configuration format.
     * `csvConfigColumns` - (Optional) Field configuration in csv configuration format.
     * `csvConfigNullidentifier` - (Optional) Invalid field content.
     * `csvConfigQuote` - (Optional) Escape character under csv configuration.
     * `csvConfigHeader` - (Optional) Indicates whether to write the field name to the CSV file, the default value is `false`.
     * `csvConfigLinefeed` - (Optional) lineFeed in csv configuration.
     * - format = `parquet`
     * `parquetConfig` - (Optional) Configure to use parquet storage format.
     * `name` - (Required) The name of the key.
     * `type` - (Required) Type of configuration name.
     */
    public readonly format!: pulumi.Output<string>;
    public readonly jsonEnableTag!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the log logstore.
     */
    public readonly logstoreName!: pulumi.Output<string>;
    /**
     * The name of the oss bucket.
     */
    public readonly ossBucket!: pulumi.Output<string>;
    /**
     * The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
     */
    public readonly ossPrefix!: pulumi.Output<string | undefined>;
    public readonly parquetConfigs!: pulumi.Output<outputs.log.OssShipperParquetConfig[] | undefined>;
    /**
     * The OSS Bucket directory is dynamically generated according to the creation time of the shipper task, it cannot start with a forward slash `/`, the default value is `%!Y(MISSING)/%!m(MISSING)/%!d(MISSING)/%!H(MISSING)/%!M(MISSING)`.
     */
    public readonly pathFormat!: pulumi.Output<string>;
    /**
     * The name of the log project. It is the only in one Alicloud account.
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * Used for access control, the OSS Bucket owner creates the role mark, such as `acs:ram::13234:role/logrole`
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
     */
    public readonly shipperName!: pulumi.Output<string>;

    /**
     * Create a OssShipper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OssShipperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OssShipperArgs | OssShipperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OssShipperState | undefined;
            resourceInputs["bufferInterval"] = state ? state.bufferInterval : undefined;
            resourceInputs["bufferSize"] = state ? state.bufferSize : undefined;
            resourceInputs["compressType"] = state ? state.compressType : undefined;
            resourceInputs["csvConfigColumns"] = state ? state.csvConfigColumns : undefined;
            resourceInputs["csvConfigDelimiter"] = state ? state.csvConfigDelimiter : undefined;
            resourceInputs["csvConfigHeader"] = state ? state.csvConfigHeader : undefined;
            resourceInputs["csvConfigLinefeed"] = state ? state.csvConfigLinefeed : undefined;
            resourceInputs["csvConfigNullidentifier"] = state ? state.csvConfigNullidentifier : undefined;
            resourceInputs["csvConfigQuote"] = state ? state.csvConfigQuote : undefined;
            resourceInputs["format"] = state ? state.format : undefined;
            resourceInputs["jsonEnableTag"] = state ? state.jsonEnableTag : undefined;
            resourceInputs["logstoreName"] = state ? state.logstoreName : undefined;
            resourceInputs["ossBucket"] = state ? state.ossBucket : undefined;
            resourceInputs["ossPrefix"] = state ? state.ossPrefix : undefined;
            resourceInputs["parquetConfigs"] = state ? state.parquetConfigs : undefined;
            resourceInputs["pathFormat"] = state ? state.pathFormat : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["roleArn"] = state ? state.roleArn : undefined;
            resourceInputs["shipperName"] = state ? state.shipperName : undefined;
        } else {
            const args = argsOrState as OssShipperArgs | undefined;
            if ((!args || args.bufferInterval === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bufferInterval'");
            }
            if ((!args || args.bufferSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bufferSize'");
            }
            if ((!args || args.format === undefined) && !opts.urn) {
                throw new Error("Missing required property 'format'");
            }
            if ((!args || args.logstoreName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logstoreName'");
            }
            if ((!args || args.ossBucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ossBucket'");
            }
            if ((!args || args.pathFormat === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pathFormat'");
            }
            if ((!args || args.projectName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectName'");
            }
            if ((!args || args.shipperName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shipperName'");
            }
            resourceInputs["bufferInterval"] = args ? args.bufferInterval : undefined;
            resourceInputs["bufferSize"] = args ? args.bufferSize : undefined;
            resourceInputs["compressType"] = args ? args.compressType : undefined;
            resourceInputs["csvConfigColumns"] = args ? args.csvConfigColumns : undefined;
            resourceInputs["csvConfigDelimiter"] = args ? args.csvConfigDelimiter : undefined;
            resourceInputs["csvConfigHeader"] = args ? args.csvConfigHeader : undefined;
            resourceInputs["csvConfigLinefeed"] = args ? args.csvConfigLinefeed : undefined;
            resourceInputs["csvConfigNullidentifier"] = args ? args.csvConfigNullidentifier : undefined;
            resourceInputs["csvConfigQuote"] = args ? args.csvConfigQuote : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["jsonEnableTag"] = args ? args.jsonEnableTag : undefined;
            resourceInputs["logstoreName"] = args ? args.logstoreName : undefined;
            resourceInputs["ossBucket"] = args ? args.ossBucket : undefined;
            resourceInputs["ossPrefix"] = args ? args.ossPrefix : undefined;
            resourceInputs["parquetConfigs"] = args ? args.parquetConfigs : undefined;
            resourceInputs["pathFormat"] = args ? args.pathFormat : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["shipperName"] = args ? args.shipperName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OssShipper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OssShipper resources.
 */
export interface OssShipperState {
    /**
     * How often is it delivered every interval.
     */
    bufferInterval?: pulumi.Input<number>;
    /**
     * Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
     */
    bufferSize?: pulumi.Input<number>;
    /**
     * OSS data storage compression method, support: none, snappy. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
     */
    compressType?: pulumi.Input<string>;
    csvConfigColumns?: pulumi.Input<pulumi.Input<string>[]>;
    csvConfigDelimiter?: pulumi.Input<string>;
    csvConfigHeader?: pulumi.Input<boolean>;
    csvConfigLinefeed?: pulumi.Input<string>;
    csvConfigNullidentifier?: pulumi.Input<string>;
    csvConfigQuote?: pulumi.Input<string>;
    /**
     * Storage format, only supports three types: `json`, `parquet`, `csv`.
     * **According to the different format, please select the following parameters**
     * - format = `json`
     * `jsonEnableTag` - (Optional) Whether to deliver the label.
     * - format = `csv`
     * `csvConfigDelimiter` - (Optional) Separator configuration in csv configuration format.
     * `csvConfigColumns` - (Optional) Field configuration in csv configuration format.
     * `csvConfigNullidentifier` - (Optional) Invalid field content.
     * `csvConfigQuote` - (Optional) Escape character under csv configuration.
     * `csvConfigHeader` - (Optional) Indicates whether to write the field name to the CSV file, the default value is `false`.
     * `csvConfigLinefeed` - (Optional) lineFeed in csv configuration.
     * - format = `parquet`
     * `parquetConfig` - (Optional) Configure to use parquet storage format.
     * `name` - (Required) The name of the key.
     * `type` - (Required) Type of configuration name.
     */
    format?: pulumi.Input<string>;
    jsonEnableTag?: pulumi.Input<boolean>;
    /**
     * The name of the log logstore.
     */
    logstoreName?: pulumi.Input<string>;
    /**
     * The name of the oss bucket.
     */
    ossBucket?: pulumi.Input<string>;
    /**
     * The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
     */
    ossPrefix?: pulumi.Input<string>;
    parquetConfigs?: pulumi.Input<pulumi.Input<inputs.log.OssShipperParquetConfig>[]>;
    /**
     * The OSS Bucket directory is dynamically generated according to the creation time of the shipper task, it cannot start with a forward slash `/`, the default value is `%!Y(MISSING)/%!m(MISSING)/%!d(MISSING)/%!H(MISSING)/%!M(MISSING)`.
     */
    pathFormat?: pulumi.Input<string>;
    /**
     * The name of the log project. It is the only in one Alicloud account.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Used for access control, the OSS Bucket owner creates the role mark, such as `acs:ram::13234:role/logrole`
     */
    roleArn?: pulumi.Input<string>;
    /**
     * Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
     */
    shipperName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OssShipper resource.
 */
export interface OssShipperArgs {
    /**
     * How often is it delivered every interval.
     */
    bufferInterval: pulumi.Input<number>;
    /**
     * Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
     */
    bufferSize: pulumi.Input<number>;
    /**
     * OSS data storage compression method, support: none, snappy. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
     */
    compressType?: pulumi.Input<string>;
    csvConfigColumns?: pulumi.Input<pulumi.Input<string>[]>;
    csvConfigDelimiter?: pulumi.Input<string>;
    csvConfigHeader?: pulumi.Input<boolean>;
    csvConfigLinefeed?: pulumi.Input<string>;
    csvConfigNullidentifier?: pulumi.Input<string>;
    csvConfigQuote?: pulumi.Input<string>;
    /**
     * Storage format, only supports three types: `json`, `parquet`, `csv`.
     * **According to the different format, please select the following parameters**
     * - format = `json`
     * `jsonEnableTag` - (Optional) Whether to deliver the label.
     * - format = `csv`
     * `csvConfigDelimiter` - (Optional) Separator configuration in csv configuration format.
     * `csvConfigColumns` - (Optional) Field configuration in csv configuration format.
     * `csvConfigNullidentifier` - (Optional) Invalid field content.
     * `csvConfigQuote` - (Optional) Escape character under csv configuration.
     * `csvConfigHeader` - (Optional) Indicates whether to write the field name to the CSV file, the default value is `false`.
     * `csvConfigLinefeed` - (Optional) lineFeed in csv configuration.
     * - format = `parquet`
     * `parquetConfig` - (Optional) Configure to use parquet storage format.
     * `name` - (Required) The name of the key.
     * `type` - (Required) Type of configuration name.
     */
    format: pulumi.Input<string>;
    jsonEnableTag?: pulumi.Input<boolean>;
    /**
     * The name of the log logstore.
     */
    logstoreName: pulumi.Input<string>;
    /**
     * The name of the oss bucket.
     */
    ossBucket: pulumi.Input<string>;
    /**
     * The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
     */
    ossPrefix?: pulumi.Input<string>;
    parquetConfigs?: pulumi.Input<pulumi.Input<inputs.log.OssShipperParquetConfig>[]>;
    /**
     * The OSS Bucket directory is dynamically generated according to the creation time of the shipper task, it cannot start with a forward slash `/`, the default value is `%!Y(MISSING)/%!m(MISSING)/%!d(MISSING)/%!H(MISSING)/%!M(MISSING)`.
     */
    pathFormat: pulumi.Input<string>;
    /**
     * The name of the log project. It is the only in one Alicloud account.
     */
    projectName: pulumi.Input<string>;
    /**
     * Used for access control, the OSS Bucket owner creates the role mark, such as `acs:ram::13234:role/logrole`
     */
    roleArn?: pulumi.Input<string>;
    /**
     * Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
     */
    shipperName: pulumi.Input<string>;
}
