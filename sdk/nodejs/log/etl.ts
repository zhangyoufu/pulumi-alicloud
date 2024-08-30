// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The data transformation of the log service is a hosted, highly available, and scalable data processing service,
 * which is widely applicable to scenarios such as data regularization, enrichment, distribution, aggregation, and index reconstruction.
 * [Refer to details](https://www.alibabacloud.com/help/zh/doc-detail/125384.htm).
 *
 * > **NOTE:** Available in 1.120.0
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
 * const example2 = new alicloud.log.Store("example2", {
 *     projectName: example.projectName,
 *     logstoreName: "example-store2",
 *     retentionPeriod: 3650,
 *     shardCount: 3,
 *     autoSplit: true,
 *     maxSplitShardCount: 60,
 *     appendMeta: true,
 * });
 * const example3 = new alicloud.log.Store("example3", {
 *     projectName: example.projectName,
 *     logstoreName: "example-store3",
 *     retentionPeriod: 3650,
 *     shardCount: 3,
 *     autoSplit: true,
 *     maxSplitShardCount: 60,
 *     appendMeta: true,
 * });
 * const exampleEtl = new alicloud.log.Etl("example", {
 *     etlName: "terraform-example",
 *     project: example.projectName,
 *     displayName: "terraform-example",
 *     description: "terraform-example",
 *     accessKeyId: "access_key_id",
 *     accessKeySecret: "access_key_secret",
 *     script: "e_set('new','key')",
 *     logstore: exampleStore.logstoreName,
 *     etlSinks: [
 *         {
 *             name: "target_name",
 *             accessKeyId: "example2_access_key_id",
 *             accessKeySecret: "example2_access_key_secret",
 *             endpoint: "cn-hangzhou.log.aliyuncs.com",
 *             project: example.projectName,
 *             logstore: example2.logstoreName,
 *         },
 *         {
 *             name: "target_name2",
 *             accessKeyId: "example3_access_key_id",
 *             accessKeySecret: "example3_access_key_secret",
 *             endpoint: "cn-hangzhou.log.aliyuncs.com",
 *             project: example.projectName,
 *             logstore: example3.logstoreName,
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Log etl can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:log/etl:Etl example tf-log-project:tf-log-etl-name
 * ```
 */
export class Etl extends pulumi.CustomResource {
    /**
     * Get an existing Etl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EtlState, opts?: pulumi.CustomResourceOptions): Etl {
        return new Etl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:log/etl:Etl';

    /**
     * Returns true if the given object is an instance of Etl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Etl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Etl.__pulumiType;
    }

    /**
     * Source logstore access key id.
     */
    public readonly accessKeyId!: pulumi.Output<string | undefined>;
    /**
     * Source logstore access key secret.
     */
    public readonly accessKeySecret!: pulumi.Output<string | undefined>;
    /**
     * The etl job create time.
     */
    public readonly createTime!: pulumi.Output<number>;
    /**
     * Description of the log etl job.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Log service etl job alias.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The name of the log etl job.
     */
    public readonly etlName!: pulumi.Output<string>;
    /**
     * Target logstore configuration for delivery after data processing.
     */
    public readonly etlSinks!: pulumi.Output<outputs.log.EtlEtlSink[]>;
    /**
     * Log service etl type, the default value is `ETL`.
     */
    public readonly etlType!: pulumi.Output<string | undefined>;
    /**
     * The start time of the processing job, if not set the value is 0, indicates to start processing from the oldest data.
     */
    public readonly fromTime!: pulumi.Output<number | undefined>;
    /**
     * An KMS encrypts access key id used to a log etl job. If the `accessKeyId` is filled in, this field will be ignored.
     */
    public readonly kmsEncryptedAccessKeyId!: pulumi.Output<string | undefined>;
    /**
     * An KMS encrypts access key secret used to a log etl job. If the `accessKeySecret` is filled in, this field will be ignored.
     */
    public readonly kmsEncryptedAccessKeySecret!: pulumi.Output<string | undefined>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedAccessKeyId` before creating or updating an instance with `kmsEncryptedAccessKeyId`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
     */
    public readonly kmsEncryptionAccessKeyIdContext!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedAccessKeySecret` before creating or updating an instance with `kmsEncryptedAccessKeySecret`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
     */
    public readonly kmsEncryptionAccessKeySecretContext!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * ETL job last modified time.
     */
    public readonly lastModifiedTime!: pulumi.Output<number>;
    /**
     * The source logstore of the processing job.
     */
    public readonly logstore!: pulumi.Output<string>;
    /**
     * Advanced parameter configuration of processing operations.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the project where the etl job is located.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Sts role info under source logstore. `roleArn` and `(access_key_id, access_key_secret)` fill in at most one. If you do not fill in both, then you must fill in `(kms_encrypted_access_key_id, kms_encrypted_access_key_secret, kms_encryption_access_key_id_context, kms_encryption_access_key_secret_context)` to use KMS to get the key pair.
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * Job scheduling type, the default value is Resident.
     */
    public readonly schedule!: pulumi.Output<string | undefined>;
    /**
     * Processing operation grammar.
     */
    public readonly script!: pulumi.Output<string>;
    /**
     * Log project tags. the default value is RUNNING, Only 4 values are supported: `STARTING`，`RUNNING`，`STOPPING`，`STOPPED`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Deadline of processing job, if not set the value is 0, indicates that new data will be processed continuously.
     */
    public readonly toTime!: pulumi.Output<number | undefined>;
    /**
     * Log etl job version. the default value is `2`.
     */
    public readonly version!: pulumi.Output<number | undefined>;

    /**
     * Create a Etl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EtlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EtlArgs | EtlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EtlState | undefined;
            resourceInputs["accessKeyId"] = state ? state.accessKeyId : undefined;
            resourceInputs["accessKeySecret"] = state ? state.accessKeySecret : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["etlName"] = state ? state.etlName : undefined;
            resourceInputs["etlSinks"] = state ? state.etlSinks : undefined;
            resourceInputs["etlType"] = state ? state.etlType : undefined;
            resourceInputs["fromTime"] = state ? state.fromTime : undefined;
            resourceInputs["kmsEncryptedAccessKeyId"] = state ? state.kmsEncryptedAccessKeyId : undefined;
            resourceInputs["kmsEncryptedAccessKeySecret"] = state ? state.kmsEncryptedAccessKeySecret : undefined;
            resourceInputs["kmsEncryptionAccessKeyIdContext"] = state ? state.kmsEncryptionAccessKeyIdContext : undefined;
            resourceInputs["kmsEncryptionAccessKeySecretContext"] = state ? state.kmsEncryptionAccessKeySecretContext : undefined;
            resourceInputs["lastModifiedTime"] = state ? state.lastModifiedTime : undefined;
            resourceInputs["logstore"] = state ? state.logstore : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["roleArn"] = state ? state.roleArn : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
            resourceInputs["script"] = state ? state.script : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["toTime"] = state ? state.toTime : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as EtlArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.etlName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'etlName'");
            }
            if ((!args || args.etlSinks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'etlSinks'");
            }
            if ((!args || args.logstore === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logstore'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.script === undefined) && !opts.urn) {
                throw new Error("Missing required property 'script'");
            }
            resourceInputs["accessKeyId"] = args?.accessKeyId ? pulumi.secret(args.accessKeyId) : undefined;
            resourceInputs["accessKeySecret"] = args?.accessKeySecret ? pulumi.secret(args.accessKeySecret) : undefined;
            resourceInputs["createTime"] = args ? args.createTime : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["etlName"] = args ? args.etlName : undefined;
            resourceInputs["etlSinks"] = args ? args.etlSinks : undefined;
            resourceInputs["etlType"] = args ? args.etlType : undefined;
            resourceInputs["fromTime"] = args ? args.fromTime : undefined;
            resourceInputs["kmsEncryptedAccessKeyId"] = args ? args.kmsEncryptedAccessKeyId : undefined;
            resourceInputs["kmsEncryptedAccessKeySecret"] = args ? args.kmsEncryptedAccessKeySecret : undefined;
            resourceInputs["kmsEncryptionAccessKeyIdContext"] = args ? args.kmsEncryptionAccessKeyIdContext : undefined;
            resourceInputs["kmsEncryptionAccessKeySecretContext"] = args ? args.kmsEncryptionAccessKeySecretContext : undefined;
            resourceInputs["lastModifiedTime"] = args ? args.lastModifiedTime : undefined;
            resourceInputs["logstore"] = args ? args.logstore : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["script"] = args ? args.script : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["toTime"] = args ? args.toTime : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessKeyId", "accessKeySecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Etl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Etl resources.
 */
export interface EtlState {
    /**
     * Source logstore access key id.
     */
    accessKeyId?: pulumi.Input<string>;
    /**
     * Source logstore access key secret.
     */
    accessKeySecret?: pulumi.Input<string>;
    /**
     * The etl job create time.
     */
    createTime?: pulumi.Input<number>;
    /**
     * Description of the log etl job.
     */
    description?: pulumi.Input<string>;
    /**
     * Log service etl job alias.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The name of the log etl job.
     */
    etlName?: pulumi.Input<string>;
    /**
     * Target logstore configuration for delivery after data processing.
     */
    etlSinks?: pulumi.Input<pulumi.Input<inputs.log.EtlEtlSink>[]>;
    /**
     * Log service etl type, the default value is `ETL`.
     */
    etlType?: pulumi.Input<string>;
    /**
     * The start time of the processing job, if not set the value is 0, indicates to start processing from the oldest data.
     */
    fromTime?: pulumi.Input<number>;
    /**
     * An KMS encrypts access key id used to a log etl job. If the `accessKeyId` is filled in, this field will be ignored.
     */
    kmsEncryptedAccessKeyId?: pulumi.Input<string>;
    /**
     * An KMS encrypts access key secret used to a log etl job. If the `accessKeySecret` is filled in, this field will be ignored.
     */
    kmsEncryptedAccessKeySecret?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedAccessKeyId` before creating or updating an instance with `kmsEncryptedAccessKeyId`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
     */
    kmsEncryptionAccessKeyIdContext?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedAccessKeySecret` before creating or updating an instance with `kmsEncryptedAccessKeySecret`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
     */
    kmsEncryptionAccessKeySecretContext?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * ETL job last modified time.
     */
    lastModifiedTime?: pulumi.Input<number>;
    /**
     * The source logstore of the processing job.
     */
    logstore?: pulumi.Input<string>;
    /**
     * Advanced parameter configuration of processing operations.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the project where the etl job is located.
     */
    project?: pulumi.Input<string>;
    /**
     * Sts role info under source logstore. `roleArn` and `(access_key_id, access_key_secret)` fill in at most one. If you do not fill in both, then you must fill in `(kms_encrypted_access_key_id, kms_encrypted_access_key_secret, kms_encryption_access_key_id_context, kms_encryption_access_key_secret_context)` to use KMS to get the key pair.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * Job scheduling type, the default value is Resident.
     */
    schedule?: pulumi.Input<string>;
    /**
     * Processing operation grammar.
     */
    script?: pulumi.Input<string>;
    /**
     * Log project tags. the default value is RUNNING, Only 4 values are supported: `STARTING`，`RUNNING`，`STOPPING`，`STOPPED`.
     */
    status?: pulumi.Input<string>;
    /**
     * Deadline of processing job, if not set the value is 0, indicates that new data will be processed continuously.
     */
    toTime?: pulumi.Input<number>;
    /**
     * Log etl job version. the default value is `2`.
     */
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Etl resource.
 */
export interface EtlArgs {
    /**
     * Source logstore access key id.
     */
    accessKeyId?: pulumi.Input<string>;
    /**
     * Source logstore access key secret.
     */
    accessKeySecret?: pulumi.Input<string>;
    /**
     * The etl job create time.
     */
    createTime?: pulumi.Input<number>;
    /**
     * Description of the log etl job.
     */
    description?: pulumi.Input<string>;
    /**
     * Log service etl job alias.
     */
    displayName: pulumi.Input<string>;
    /**
     * The name of the log etl job.
     */
    etlName: pulumi.Input<string>;
    /**
     * Target logstore configuration for delivery after data processing.
     */
    etlSinks: pulumi.Input<pulumi.Input<inputs.log.EtlEtlSink>[]>;
    /**
     * Log service etl type, the default value is `ETL`.
     */
    etlType?: pulumi.Input<string>;
    /**
     * The start time of the processing job, if not set the value is 0, indicates to start processing from the oldest data.
     */
    fromTime?: pulumi.Input<number>;
    /**
     * An KMS encrypts access key id used to a log etl job. If the `accessKeyId` is filled in, this field will be ignored.
     */
    kmsEncryptedAccessKeyId?: pulumi.Input<string>;
    /**
     * An KMS encrypts access key secret used to a log etl job. If the `accessKeySecret` is filled in, this field will be ignored.
     */
    kmsEncryptedAccessKeySecret?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedAccessKeyId` before creating or updating an instance with `kmsEncryptedAccessKeyId`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
     */
    kmsEncryptionAccessKeyIdContext?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedAccessKeySecret` before creating or updating an instance with `kmsEncryptedAccessKeySecret`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
     */
    kmsEncryptionAccessKeySecretContext?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * ETL job last modified time.
     */
    lastModifiedTime?: pulumi.Input<number>;
    /**
     * The source logstore of the processing job.
     */
    logstore: pulumi.Input<string>;
    /**
     * Advanced parameter configuration of processing operations.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the project where the etl job is located.
     */
    project: pulumi.Input<string>;
    /**
     * Sts role info under source logstore. `roleArn` and `(access_key_id, access_key_secret)` fill in at most one. If you do not fill in both, then you must fill in `(kms_encrypted_access_key_id, kms_encrypted_access_key_secret, kms_encryption_access_key_id_context, kms_encryption_access_key_secret_context)` to use KMS to get the key pair.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * Job scheduling type, the default value is Resident.
     */
    schedule?: pulumi.Input<string>;
    /**
     * Processing operation grammar.
     */
    script: pulumi.Input<string>;
    /**
     * Log project tags. the default value is RUNNING, Only 4 values are supported: `STARTING`，`RUNNING`，`STOPPING`，`STOPPED`.
     */
    status?: pulumi.Input<string>;
    /**
     * Deadline of processing job, if not set the value is 0, indicates that new data will be processed continuously.
     */
    toTime?: pulumi.Input<number>;
    /**
     * Log etl job version. the default value is `2`.
     */
    version?: pulumi.Input<number>;
}
