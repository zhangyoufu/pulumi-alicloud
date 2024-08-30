// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.oss.BucketObjectArgs;
import com.pulumi.alicloud.oss.inputs.BucketObjectState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to put a object(content or file) to a oss bucket.
 * 
 * ## Example Usage
 * 
 * ### Uploading a file to a bucket
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.BucketAcl;
 * import com.pulumi.alicloud.oss.BucketAclArgs;
 * import com.pulumi.alicloud.oss.BucketObject;
 * import com.pulumi.alicloud.oss.BucketObjectArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var default_ = new Integer("default", IntegerArgs.builder()
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         var defaultBucket = new Bucket("defaultBucket", BucketArgs.builder()
 *             .bucket(String.format("terraform-example-%s", default_.result()))
 *             .build());
 * 
 *         var defaultBucketAcl = new BucketAcl("defaultBucketAcl", BucketAclArgs.builder()
 *             .bucket(defaultBucket.bucket())
 *             .acl("private")
 *             .build());
 * 
 *         var defaultBucketObject = new BucketObject("defaultBucketObject", BucketObjectArgs.builder()
 *             .bucket(defaultBucket.bucket())
 *             .key("example_key")
 *             .source("./main.tf")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Uploading a content to a bucket
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.BucketAcl;
 * import com.pulumi.alicloud.oss.BucketAclArgs;
 * import com.pulumi.alicloud.oss.BucketObject;
 * import com.pulumi.alicloud.oss.BucketObjectArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var default_ = new Integer("default", IntegerArgs.builder()
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         var defaultBucket = new Bucket("defaultBucket", BucketArgs.builder()
 *             .bucket(String.format("terraform-example-%s", default_.result()))
 *             .build());
 * 
 *         var defaultBucketAcl = new BucketAcl("defaultBucketAcl", BucketAclArgs.builder()
 *             .bucket(defaultBucket.bucket())
 *             .acl("private")
 *             .build());
 * 
 *         var defaultBucketObject = new BucketObject("defaultBucketObject", BucketObjectArgs.builder()
 *             .bucket(defaultBucket.bucket())
 *             .key("example_key")
 *             .content("the content that you want to upload.")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="alicloud:oss/bucketObject:BucketObject")
public class BucketObject extends com.pulumi.resources.CustomResource {
    /**
     * The [canned ACL](https://www.alibabacloud.com/help/doc-detail/52284.htm) to apply. Defaults to &#34;private&#34;.
     * 
     */
    @Export(name="acl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> acl;

    /**
     * @return The [canned ACL](https://www.alibabacloud.com/help/doc-detail/52284.htm) to apply. Defaults to &#34;private&#34;.
     * 
     */
    public Output<Optional<String>> acl() {
        return Codegen.optional(this.acl);
    }
    /**
     * The name of the bucket to put the file in.
     * 
     */
    @Export(name="bucket", refs={String.class}, tree="[0]")
    private Output<String> bucket;

    /**
     * @return The name of the bucket to put the file in.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }
    /**
     * Specifies caching behavior along the request/reply chain. Read [RFC2616 Cache-Control](https://www.ietf.org/rfc/rfc2616.txt) for further details.
     * 
     */
    @Export(name="cacheControl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cacheControl;

    /**
     * @return Specifies caching behavior along the request/reply chain. Read [RFC2616 Cache-Control](https://www.ietf.org/rfc/rfc2616.txt) for further details.
     * 
     */
    public Output<Optional<String>> cacheControl() {
        return Codegen.optional(this.cacheControl);
    }
    /**
     * The literal content being uploaded to the bucket.
     * 
     */
    @Export(name="content", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> content;

    /**
     * @return The literal content being uploaded to the bucket.
     * 
     */
    public Output<Optional<String>> content() {
        return Codegen.optional(this.content);
    }
    /**
     * Specifies presentational information for the object. Read [RFC2616 Content-Disposition](https://www.ietf.org/rfc/rfc2616.txt) for further details.
     * 
     */
    @Export(name="contentDisposition", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> contentDisposition;

    /**
     * @return Specifies presentational information for the object. Read [RFC2616 Content-Disposition](https://www.ietf.org/rfc/rfc2616.txt) for further details.
     * 
     */
    public Output<Optional<String>> contentDisposition() {
        return Codegen.optional(this.contentDisposition);
    }
    /**
     * Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [RFC2616 Content-Encoding](https://www.ietf.org/rfc/rfc2616.txt) for further details.
     * 
     */
    @Export(name="contentEncoding", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> contentEncoding;

    /**
     * @return Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [RFC2616 Content-Encoding](https://www.ietf.org/rfc/rfc2616.txt) for further details.
     * 
     */
    public Output<Optional<String>> contentEncoding() {
        return Codegen.optional(this.contentEncoding);
    }
    /**
     * the content length of request.
     * 
     */
    @Export(name="contentLength", refs={String.class}, tree="[0]")
    private Output<String> contentLength;

    /**
     * @return the content length of request.
     * 
     */
    public Output<String> contentLength() {
        return this.contentLength;
    }
    /**
     * The MD5 value of the content. Read [MD5](https://www.alibabacloud.com/help/doc-detail/31978.htm) for computing method.
     * 
     */
    @Export(name="contentMd5", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> contentMd5;

    /**
     * @return The MD5 value of the content. Read [MD5](https://www.alibabacloud.com/help/doc-detail/31978.htm) for computing method.
     * 
     */
    public Output<Optional<String>> contentMd5() {
        return Codegen.optional(this.contentMd5);
    }
    /**
     * A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
     * 
     */
    @Export(name="contentType", refs={String.class}, tree="[0]")
    private Output<String> contentType;

    /**
     * @return A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
     * 
     */
    public Output<String> contentType() {
        return this.contentType;
    }
    /**
     * the ETag generated for the object (an MD5 sum of the object content).
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return the ETag generated for the object (an MD5 sum of the object content).
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Specifies expire date for the the request/response. Read [RFC2616 Expires](https://www.ietf.org/rfc/rfc2616.txt) for further details.
     * 
     */
    @Export(name="expires", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> expires;

    /**
     * @return Specifies expire date for the the request/response. Read [RFC2616 Expires](https://www.ietf.org/rfc/rfc2616.txt) for further details.
     * 
     */
    public Output<Optional<String>> expires() {
        return Codegen.optional(this.expires);
    }
    /**
     * The name of the object once it is in the bucket.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return The name of the object once it is in the bucket.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * Specifies the primary key managed by KMS. This parameter is valid when the value of `server_side_encryption` is set to KMS.
     * 
     * Either `source` or `content` must be provided to specify the bucket content.
     * These two arguments are mutually-exclusive.
     * 
     */
    @Export(name="kmsKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsKeyId;

    /**
     * @return Specifies the primary key managed by KMS. This parameter is valid when the value of `server_side_encryption` is set to KMS.
     * 
     * Either `source` or `content` must be provided to specify the bucket content.
     * These two arguments are mutually-exclusive.
     * 
     */
    public Output<Optional<String>> kmsKeyId() {
        return Codegen.optional(this.kmsKeyId);
    }
    /**
     * Specifies server-side encryption of the object in OSS. Valid values are `AES256`, `KMS`. Default value is `AES256`.
     * 
     */
    @Export(name="serverSideEncryption", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serverSideEncryption;

    /**
     * @return Specifies server-side encryption of the object in OSS. Valid values are `AES256`, `KMS`. Default value is `AES256`.
     * 
     */
    public Output<Optional<String>> serverSideEncryption() {
        return Codegen.optional(this.serverSideEncryption);
    }
    /**
     * The path to the source file being uploaded to the bucket.
     * 
     */
    @Export(name="source", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> source;

    /**
     * @return The path to the source file being uploaded to the bucket.
     * 
     */
    public Output<Optional<String>> source() {
        return Codegen.optional(this.source);
    }
    /**
     * A unique version ID value for the object, if bucket versioning is enabled.
     * 
     */
    @Export(name="versionId", refs={String.class}, tree="[0]")
    private Output<String> versionId;

    /**
     * @return A unique version ID value for the object, if bucket versioning is enabled.
     * 
     */
    public Output<String> versionId() {
        return this.versionId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BucketObject(java.lang.String name) {
        this(name, BucketObjectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BucketObject(java.lang.String name, BucketObjectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BucketObject(java.lang.String name, BucketObjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oss/bucketObject:BucketObject", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private BucketObject(java.lang.String name, Output<java.lang.String> id, @Nullable BucketObjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oss/bucketObject:BucketObject", name, state, makeResourceOptions(options, id), false);
    }

    private static BucketObjectArgs makeArgs(BucketObjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? BucketObjectArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static BucketObject get(java.lang.String name, Output<java.lang.String> id, @Nullable BucketObjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BucketObject(name, id, state, options);
    }
}
