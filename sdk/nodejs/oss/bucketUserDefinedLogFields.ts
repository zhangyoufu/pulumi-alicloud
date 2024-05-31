// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a OSS Bucket User Defined Log Fields resource. Used to personalize the userDefinedLogFields field in the Bucket real-time log.
 *
 * For information about OSS Bucket User Defined Log Fields and how to use it, see [What is Bucket User Defined Log Fields](https://www.alibabacloud.com/help/en/oss/developer-reference/putuserdefinedlogfieldsconfig).
 *
 * > **NOTE:** Available since v1.224.0.
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
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const _default = new random.index.Integer("default", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const createBucket = new alicloud.oss.Bucket("CreateBucket", {
 *     storageClass: "Standard",
 *     bucket: `${name}-${_default.result}`,
 * });
 * const defaultBucketUserDefinedLogFields = new alicloud.oss.BucketUserDefinedLogFields("default", {
 *     bucket: createBucket.bucket,
 *     paramSets: [
 *         "oss-example",
 *         "example-para",
 *         "abc",
 *     ],
 *     headerSets: [
 *         "def",
 *         "example-header",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * OSS Bucket User Defined Log Fields can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:oss/bucketUserDefinedLogFields:BucketUserDefinedLogFields example <id>
 * ```
 */
export class BucketUserDefinedLogFields extends pulumi.CustomResource {
    /**
     * Get an existing BucketUserDefinedLogFields resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketUserDefinedLogFieldsState, opts?: pulumi.CustomResourceOptions): BucketUserDefinedLogFields {
        return new BucketUserDefinedLogFields(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:oss/bucketUserDefinedLogFields:BucketUserDefinedLogFields';

    /**
     * Returns true if the given object is an instance of BucketUserDefinedLogFields.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketUserDefinedLogFields {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketUserDefinedLogFields.__pulumiType;
    }

    /**
     * The name of the bucket.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Container for custom request header configuration information.
     */
    public readonly headerSets!: pulumi.Output<string[] | undefined>;
    /**
     * Container for custom request parameters configuration information.
     */
    public readonly paramSets!: pulumi.Output<string[] | undefined>;

    /**
     * Create a BucketUserDefinedLogFields resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketUserDefinedLogFieldsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketUserDefinedLogFieldsArgs | BucketUserDefinedLogFieldsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BucketUserDefinedLogFieldsState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["headerSets"] = state ? state.headerSets : undefined;
            resourceInputs["paramSets"] = state ? state.paramSets : undefined;
        } else {
            const args = argsOrState as BucketUserDefinedLogFieldsArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["headerSets"] = args ? args.headerSets : undefined;
            resourceInputs["paramSets"] = args ? args.paramSets : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BucketUserDefinedLogFields.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketUserDefinedLogFields resources.
 */
export interface BucketUserDefinedLogFieldsState {
    /**
     * The name of the bucket.
     */
    bucket?: pulumi.Input<string>;
    /**
     * Container for custom request header configuration information.
     */
    headerSets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Container for custom request parameters configuration information.
     */
    paramSets?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a BucketUserDefinedLogFields resource.
 */
export interface BucketUserDefinedLogFieldsArgs {
    /**
     * The name of the bucket.
     */
    bucket: pulumi.Input<string>;
    /**
     * Container for custom request header configuration information.
     */
    headerSets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Container for custom request parameters configuration information.
     */
    paramSets?: pulumi.Input<pulumi.Input<string>[]>;
}
