// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an independent replication configuration resource for OSS bucket.
 *
 * For information about OSS replication and how to use it, see [What is cross-region replication](https://www.alibabacloud.com/help/doc-detail/31864.html) and [What is same-region replication](https://www.alibabacloud.com/help/doc-detail/254865.html).
 *
 * > **NOTE:** Available since v1.161.0.
 *
 * ## Example Usage
 *
 * Set bucket replication configuration
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
 * const bucketSrc = new alicloud.oss.Bucket("bucketSrc", {bucket: pulumi.interpolate`example-src-${_default.result}`});
 * const bucketDest = new alicloud.oss.Bucket("bucketDest", {bucket: pulumi.interpolate`example-dest-${_default.result}`});
 * const role = new alicloud.ram.Role("role", {
 *     document: `		{
 * 		  "Statement": [
 * 			{
 * 			  "Action": "sts:AssumeRole",
 * 			  "Effect": "Allow",
 * 			  "Principal": {
 * 				"Service": [
 * 				  "oss.aliyuncs.com"
 * 				]
 * 			  }
 * 			}
 * 		  ],
 * 		  "Version": "1"
 * 		}
 * `,
 *     description: "this is a test",
 *     force: true,
 * });
 * const policy = new alicloud.ram.Policy("policy", {
 *     policyName: pulumi.interpolate`example-policy-${_default.result}`,
 *     policyDocument: `		{
 * 		  "Statement": [
 * 			{
 * 			  "Action": [
 * 				"*"
 * 			  ],
 * 			  "Effect": "Allow",
 * 			  "Resource": [
 * 				"*"
 * 			  ]
 * 			}
 * 		  ],
 * 			"Version": "1"
 * 		}
 * `,
 *     description: "this is a policy test",
 *     force: true,
 * });
 * const attach = new alicloud.ram.RolePolicyAttachment("attach", {
 *     policyName: policy.name,
 *     policyType: policy.type,
 *     roleName: role.name,
 * });
 * const key = new alicloud.kms.Key("key", {
 *     description: "Hello KMS",
 *     pendingWindowInDays: 7,
 *     status: "Enabled",
 * });
 * const cross_region_replication = new alicloud.oss.BucketReplication("cross-region-replication", {
 *     bucket: bucketSrc.id,
 *     action: "PUT,DELETE",
 *     historicalObjectReplication: "enabled",
 *     prefixSet: {
 *         prefixes: [
 *             "prefix1/",
 *             "prefix2/",
 *         ],
 *     },
 *     destination: {
 *         bucket: bucketDest.id,
 *         location: bucketDest.location,
 *     },
 *     syncRole: role.name,
 *     encryptionConfiguration: {
 *         replicaKmsKeyId: key.id,
 *     },
 *     sourceSelectionCriteria: {
 *         sseKmsEncryptedObjects: {
 *             status: "Enabled",
 *         },
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ### Timeouts
 *
 * The `timeouts` block allows you to specify timeouts for certain actions:
 *
 * * `delete` - (Defaults to 30 mins) Used when delete a data replication rule (until the data replication task is cleared).
 */
export class BucketReplication extends pulumi.CustomResource {
    /**
     * Get an existing BucketReplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketReplicationState, opts?: pulumi.CustomResourceOptions): BucketReplication {
        return new BucketReplication(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:oss/bucketReplication:BucketReplication';

    /**
     * Returns true if the given object is an instance of BucketReplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketReplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketReplication.__pulumiType;
    }

    /**
     * The operations that can be synchronized to the destination bucket. You can set action to one or more of the following operation types. Valid values: `ALL`(contains PUT, DELETE, and ABORT), `PUT`, `DELETE` and `ABORT`. Defaults to `ALL`.
     */
    public readonly action!: pulumi.Output<string | undefined>;
    /**
     * The name of the bucket.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Specifies the destination for the rule. See `destination` below.
     */
    public readonly destination!: pulumi.Output<outputs.oss.BucketReplicationDestination>;
    /**
     * Specifies the encryption configuration for the objects replicated to the destination bucket. See `encryptionConfiguration` below.
     */
    public readonly encryptionConfiguration!: pulumi.Output<outputs.oss.BucketReplicationEncryptionConfiguration | undefined>;
    /**
     * Specifies whether to replicate historical data from the source bucket to the destination bucket before data replication is enabled. Can be `enabled` or `disabled`. Defaults to `enabled`.
     */
    public readonly historicalObjectReplication!: pulumi.Output<string | undefined>;
    /**
     * The prefixes used to specify the object to replicate. Only objects that match the prefix are replicated to the destination bucket. See `prefixSet` below.
     */
    public readonly prefixSet!: pulumi.Output<outputs.oss.BucketReplicationPrefixSet | undefined>;
    /**
     * Specifies the progress for querying the progress of a data replication task of a bucket.
     */
    public readonly progress!: pulumi.Output<outputs.oss.BucketReplicationProgress>;
    /**
     * The ID of the data replication rule.
     */
    public /*out*/ readonly ruleId!: pulumi.Output<string>;
    /**
     * Specifies other conditions used to filter the source objects to replicate. See `sourceSelectionCriteria` below.
     */
    public readonly sourceSelectionCriteria!: pulumi.Output<outputs.oss.BucketReplicationSourceSelectionCriteria | undefined>;
    /**
     * Specifies whether to replicate objects encrypted by using SSE-KMS. Can be `Enabled` or `Disabled`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the role that you authorize OSS to use to replicate data. If SSE-KMS is specified to encrypt the objects replicated to the destination bucket, it must be specified.
     */
    public readonly syncRole!: pulumi.Output<string | undefined>;

    /**
     * Create a BucketReplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketReplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketReplicationArgs | BucketReplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BucketReplicationState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["destination"] = state ? state.destination : undefined;
            resourceInputs["encryptionConfiguration"] = state ? state.encryptionConfiguration : undefined;
            resourceInputs["historicalObjectReplication"] = state ? state.historicalObjectReplication : undefined;
            resourceInputs["prefixSet"] = state ? state.prefixSet : undefined;
            resourceInputs["progress"] = state ? state.progress : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["sourceSelectionCriteria"] = state ? state.sourceSelectionCriteria : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["syncRole"] = state ? state.syncRole : undefined;
        } else {
            const args = argsOrState as BucketReplicationArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["encryptionConfiguration"] = args ? args.encryptionConfiguration : undefined;
            resourceInputs["historicalObjectReplication"] = args ? args.historicalObjectReplication : undefined;
            resourceInputs["prefixSet"] = args ? args.prefixSet : undefined;
            resourceInputs["progress"] = args ? args.progress : undefined;
            resourceInputs["sourceSelectionCriteria"] = args ? args.sourceSelectionCriteria : undefined;
            resourceInputs["syncRole"] = args ? args.syncRole : undefined;
            resourceInputs["ruleId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BucketReplication.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketReplication resources.
 */
export interface BucketReplicationState {
    /**
     * The operations that can be synchronized to the destination bucket. You can set action to one or more of the following operation types. Valid values: `ALL`(contains PUT, DELETE, and ABORT), `PUT`, `DELETE` and `ABORT`. Defaults to `ALL`.
     */
    action?: pulumi.Input<string>;
    /**
     * The name of the bucket.
     */
    bucket?: pulumi.Input<string>;
    /**
     * Specifies the destination for the rule. See `destination` below.
     */
    destination?: pulumi.Input<inputs.oss.BucketReplicationDestination>;
    /**
     * Specifies the encryption configuration for the objects replicated to the destination bucket. See `encryptionConfiguration` below.
     */
    encryptionConfiguration?: pulumi.Input<inputs.oss.BucketReplicationEncryptionConfiguration>;
    /**
     * Specifies whether to replicate historical data from the source bucket to the destination bucket before data replication is enabled. Can be `enabled` or `disabled`. Defaults to `enabled`.
     */
    historicalObjectReplication?: pulumi.Input<string>;
    /**
     * The prefixes used to specify the object to replicate. Only objects that match the prefix are replicated to the destination bucket. See `prefixSet` below.
     */
    prefixSet?: pulumi.Input<inputs.oss.BucketReplicationPrefixSet>;
    /**
     * Specifies the progress for querying the progress of a data replication task of a bucket.
     */
    progress?: pulumi.Input<inputs.oss.BucketReplicationProgress>;
    /**
     * The ID of the data replication rule.
     */
    ruleId?: pulumi.Input<string>;
    /**
     * Specifies other conditions used to filter the source objects to replicate. See `sourceSelectionCriteria` below.
     */
    sourceSelectionCriteria?: pulumi.Input<inputs.oss.BucketReplicationSourceSelectionCriteria>;
    /**
     * Specifies whether to replicate objects encrypted by using SSE-KMS. Can be `Enabled` or `Disabled`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the role that you authorize OSS to use to replicate data. If SSE-KMS is specified to encrypt the objects replicated to the destination bucket, it must be specified.
     */
    syncRole?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BucketReplication resource.
 */
export interface BucketReplicationArgs {
    /**
     * The operations that can be synchronized to the destination bucket. You can set action to one or more of the following operation types. Valid values: `ALL`(contains PUT, DELETE, and ABORT), `PUT`, `DELETE` and `ABORT`. Defaults to `ALL`.
     */
    action?: pulumi.Input<string>;
    /**
     * The name of the bucket.
     */
    bucket: pulumi.Input<string>;
    /**
     * Specifies the destination for the rule. See `destination` below.
     */
    destination: pulumi.Input<inputs.oss.BucketReplicationDestination>;
    /**
     * Specifies the encryption configuration for the objects replicated to the destination bucket. See `encryptionConfiguration` below.
     */
    encryptionConfiguration?: pulumi.Input<inputs.oss.BucketReplicationEncryptionConfiguration>;
    /**
     * Specifies whether to replicate historical data from the source bucket to the destination bucket before data replication is enabled. Can be `enabled` or `disabled`. Defaults to `enabled`.
     */
    historicalObjectReplication?: pulumi.Input<string>;
    /**
     * The prefixes used to specify the object to replicate. Only objects that match the prefix are replicated to the destination bucket. See `prefixSet` below.
     */
    prefixSet?: pulumi.Input<inputs.oss.BucketReplicationPrefixSet>;
    /**
     * Specifies the progress for querying the progress of a data replication task of a bucket.
     */
    progress?: pulumi.Input<inputs.oss.BucketReplicationProgress>;
    /**
     * Specifies other conditions used to filter the source objects to replicate. See `sourceSelectionCriteria` below.
     */
    sourceSelectionCriteria?: pulumi.Input<inputs.oss.BucketReplicationSourceSelectionCriteria>;
    /**
     * Specifies the role that you authorize OSS to use to replicate data. If SSE-KMS is specified to encrypt the objects replicated to the destination bucket, it must be specified.
     */
    syncRole?: pulumi.Input<string>;
}
