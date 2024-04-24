// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * OSS Bucket Versioning can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:oss/bucketVersioning:BucketVersioning example <id>
 * ```
 */
export class BucketVersioning extends pulumi.CustomResource {
    /**
     * Get an existing BucketVersioning resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketVersioningState, opts?: pulumi.CustomResourceOptions): BucketVersioning {
        return new BucketVersioning(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:oss/bucketVersioning:BucketVersioning';

    /**
     * Returns true if the given object is an instance of BucketVersioning.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketVersioning {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketVersioning.__pulumiType;
    }

    /**
     * The name of the bucket.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * A bucket can be in one of the following versioning states: disabled, enabled, or suspended. By default, versioning is disabled for a bucket. Updating the value from Enabled or Suspended to Disabled will result in errors, because OSS does not support returning buckets to an unversioned state. .
     */
    public readonly status!: pulumi.Output<string>;

    /**
     * Create a BucketVersioning resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketVersioningArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketVersioningArgs | BucketVersioningState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BucketVersioningState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as BucketVersioningArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BucketVersioning.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketVersioning resources.
 */
export interface BucketVersioningState {
    /**
     * The name of the bucket.
     */
    bucket?: pulumi.Input<string>;
    /**
     * A bucket can be in one of the following versioning states: disabled, enabled, or suspended. By default, versioning is disabled for a bucket. Updating the value from Enabled or Suspended to Disabled will result in errors, because OSS does not support returning buckets to an unversioned state. .
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BucketVersioning resource.
 */
export interface BucketVersioningArgs {
    /**
     * The name of the bucket.
     */
    bucket: pulumi.Input<string>;
    /**
     * A bucket can be in one of the following versioning states: disabled, enabled, or suspended. By default, versioning is disabled for a bucket. Updating the value from Enabled or Suspended to Disabled will result in errors, because OSS does not support returning buckets to an unversioned state. .
     */
    status?: pulumi.Input<string>;
}
