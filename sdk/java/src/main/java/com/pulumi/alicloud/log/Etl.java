// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.log.EtlArgs;
import com.pulumi.alicloud.log.inputs.EtlState;
import com.pulumi.alicloud.log.outputs.EtlEtlSink;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The data transformation of the log service is a hosted, highly available, and scalable data processing service,
 * which is widely applicable to scenarios such as data regularization, enrichment, distribution, aggregation, and index reconstruction.
 * [Refer to details](https://www.alibabacloud.com/help/zh/doc-detail/125384.htm).
 * 
 * &gt; **NOTE:** Available in 1.120.0
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.log.Project;
 * import com.pulumi.alicloud.log.ProjectArgs;
 * import com.pulumi.alicloud.log.Store;
 * import com.pulumi.alicloud.log.StoreArgs;
 * import com.pulumi.alicloud.log.Etl;
 * import com.pulumi.alicloud.log.EtlArgs;
 * import com.pulumi.alicloud.log.inputs.EtlEtlSinkArgs;
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
 *         var default_ = new Integer(&#34;default&#34;, IntegerArgs.builder()        
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         var example = new Project(&#34;example&#34;, ProjectArgs.builder()        
 *             .name(String.format(&#34;terraform-example-%s&#34;, default_.result()))
 *             .description(&#34;terraform-example&#34;)
 *             .build());
 * 
 *         var exampleStore = new Store(&#34;exampleStore&#34;, StoreArgs.builder()        
 *             .project(example.name())
 *             .name(&#34;example-store&#34;)
 *             .retentionPeriod(3650)
 *             .shardCount(3)
 *             .autoSplit(true)
 *             .maxSplitShardCount(60)
 *             .appendMeta(true)
 *             .build());
 * 
 *         var example2 = new Store(&#34;example2&#34;, StoreArgs.builder()        
 *             .project(example.name())
 *             .name(&#34;example-store2&#34;)
 *             .retentionPeriod(3650)
 *             .shardCount(3)
 *             .autoSplit(true)
 *             .maxSplitShardCount(60)
 *             .appendMeta(true)
 *             .build());
 * 
 *         var example3 = new Store(&#34;example3&#34;, StoreArgs.builder()        
 *             .project(example.name())
 *             .name(&#34;example-store3&#34;)
 *             .retentionPeriod(3650)
 *             .shardCount(3)
 *             .autoSplit(true)
 *             .maxSplitShardCount(60)
 *             .appendMeta(true)
 *             .build());
 * 
 *         var exampleEtl = new Etl(&#34;exampleEtl&#34;, EtlArgs.builder()        
 *             .etlName(&#34;terraform-example&#34;)
 *             .project(example.name())
 *             .displayName(&#34;terraform-example&#34;)
 *             .description(&#34;terraform-example&#34;)
 *             .accessKeyId(&#34;access_key_id&#34;)
 *             .accessKeySecret(&#34;access_key_secret&#34;)
 *             .script(&#34;e_set(&#39;new&#39;,&#39;key&#39;)&#34;)
 *             .logstore(exampleStore.name())
 *             .etlSinks(            
 *                 EtlEtlSinkArgs.builder()
 *                     .name(&#34;target_name&#34;)
 *                     .accessKeyId(&#34;example2_access_key_id&#34;)
 *                     .accessKeySecret(&#34;example2_access_key_secret&#34;)
 *                     .endpoint(&#34;cn-hangzhou.log.aliyuncs.com&#34;)
 *                     .project(example.name())
 *                     .logstore(example2.name())
 *                     .build(),
 *                 EtlEtlSinkArgs.builder()
 *                     .name(&#34;target_name2&#34;)
 *                     .accessKeyId(&#34;example3_access_key_id&#34;)
 *                     .accessKeySecret(&#34;example3_access_key_secret&#34;)
 *                     .endpoint(&#34;cn-hangzhou.log.aliyuncs.com&#34;)
 *                     .project(example.name())
 *                     .logstore(example3.name())
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Log etl can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:log/etl:Etl example tf-log-project:tf-log-etl-name
 * ```
 * 
 */
@ResourceType(type="alicloud:log/etl:Etl")
public class Etl extends com.pulumi.resources.CustomResource {
    /**
     * Delivery target logstore access key id.
     * 
     */
    @Export(name="accessKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessKeyId;

    /**
     * @return Delivery target logstore access key id.
     * 
     */
    public Output<Optional<String>> accessKeyId() {
        return Codegen.optional(this.accessKeyId);
    }
    /**
     * Delivery target logstore access key secret.
     * 
     */
    @Export(name="accessKeySecret", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessKeySecret;

    /**
     * @return Delivery target logstore access key secret.
     * 
     */
    public Output<Optional<String>> accessKeySecret() {
        return Codegen.optional(this.accessKeySecret);
    }
    /**
     * The etl job create time.
     * 
     */
    @Export(name="createTime", refs={Integer.class}, tree="[0]")
    private Output<Integer> createTime;

    /**
     * @return The etl job create time.
     * 
     */
    public Output<Integer> createTime() {
        return this.createTime;
    }
    /**
     * Description of the log etl job.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the log etl job.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Log service etl job alias.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return Log service etl job alias.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * The name of the log etl job.
     * 
     */
    @Export(name="etlName", refs={String.class}, tree="[0]")
    private Output<String> etlName;

    /**
     * @return The name of the log etl job.
     * 
     */
    public Output<String> etlName() {
        return this.etlName;
    }
    /**
     * Target logstore configuration for delivery after data processing.
     * 
     */
    @Export(name="etlSinks", refs={List.class,EtlEtlSink.class}, tree="[0,1]")
    private Output<List<EtlEtlSink>> etlSinks;

    /**
     * @return Target logstore configuration for delivery after data processing.
     * 
     */
    public Output<List<EtlEtlSink>> etlSinks() {
        return this.etlSinks;
    }
    /**
     * Log service etl type, the default value is `ETL`.
     * 
     */
    @Export(name="etlType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> etlType;

    /**
     * @return Log service etl type, the default value is `ETL`.
     * 
     */
    public Output<Optional<String>> etlType() {
        return Codegen.optional(this.etlType);
    }
    /**
     * The start time of the processing job, if not set the value is 0, indicates to start processing from the oldest data.
     * 
     */
    @Export(name="fromTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> fromTime;

    /**
     * @return The start time of the processing job, if not set the value is 0, indicates to start processing from the oldest data.
     * 
     */
    public Output<Optional<Integer>> fromTime() {
        return Codegen.optional(this.fromTime);
    }
    /**
     * An KMS encrypts access key id used to a log etl job. If the `access_key_id` is filled in, this field will be ignored.
     * 
     */
    @Export(name="kmsEncryptedAccessKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsEncryptedAccessKeyId;

    /**
     * @return An KMS encrypts access key id used to a log etl job. If the `access_key_id` is filled in, this field will be ignored.
     * 
     */
    public Output<Optional<String>> kmsEncryptedAccessKeyId() {
        return Codegen.optional(this.kmsEncryptedAccessKeyId);
    }
    /**
     * An KMS encrypts access key secret used to a log etl job. If the `access_key_secret` is filled in, this field will be ignored.
     * 
     */
    @Export(name="kmsEncryptedAccessKeySecret", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsEncryptedAccessKeySecret;

    /**
     * @return An KMS encrypts access key secret used to a log etl job. If the `access_key_secret` is filled in, this field will be ignored.
     * 
     */
    public Output<Optional<String>> kmsEncryptedAccessKeySecret() {
        return Codegen.optional(this.kmsEncryptedAccessKeySecret);
    }
    /**
     * An KMS encryption context used to decrypt `kms_encrypted_access_key_id` before creating or updating an instance with `kms_encrypted_access_key_id`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set. When it is changed, the instance will reboot to make the change take effect.
     * 
     */
    @Export(name="kmsEncryptionAccessKeyIdContext", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> kmsEncryptionAccessKeyIdContext;

    /**
     * @return An KMS encryption context used to decrypt `kms_encrypted_access_key_id` before creating or updating an instance with `kms_encrypted_access_key_id`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set. When it is changed, the instance will reboot to make the change take effect.
     * 
     */
    public Output<Optional<Map<String,Object>>> kmsEncryptionAccessKeyIdContext() {
        return Codegen.optional(this.kmsEncryptionAccessKeyIdContext);
    }
    /**
     * An KMS encryption context used to decrypt `kms_encrypted_access_key_secret` before creating or updating an instance with `kms_encrypted_access_key_secret`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set. When it is changed, the instance will reboot to make the change take effect.
     * 
     */
    @Export(name="kmsEncryptionAccessKeySecretContext", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> kmsEncryptionAccessKeySecretContext;

    /**
     * @return An KMS encryption context used to decrypt `kms_encrypted_access_key_secret` before creating or updating an instance with `kms_encrypted_access_key_secret`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set. When it is changed, the instance will reboot to make the change take effect.
     * 
     */
    public Output<Optional<Map<String,Object>>> kmsEncryptionAccessKeySecretContext() {
        return Codegen.optional(this.kmsEncryptionAccessKeySecretContext);
    }
    /**
     * ETL job last modified time.
     * 
     */
    @Export(name="lastModifiedTime", refs={Integer.class}, tree="[0]")
    private Output<Integer> lastModifiedTime;

    /**
     * @return ETL job last modified time.
     * 
     */
    public Output<Integer> lastModifiedTime() {
        return this.lastModifiedTime;
    }
    /**
     * Delivery target logstore.
     * 
     */
    @Export(name="logstore", refs={String.class}, tree="[0]")
    private Output<String> logstore;

    /**
     * @return Delivery target logstore.
     * 
     */
    public Output<String> logstore() {
        return this.logstore;
    }
    /**
     * Advanced parameter configuration of processing operations.
     * 
     */
    @Export(name="parameters", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> parameters;

    /**
     * @return Advanced parameter configuration of processing operations.
     * 
     */
    public Output<Optional<Map<String,String>>> parameters() {
        return Codegen.optional(this.parameters);
    }
    /**
     * The project where the target logstore is delivered.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The project where the target logstore is delivered.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * Sts role info under delivery target logstore. `role_arn` and `(access_key_id, access_key_secret)` fill in at most one. If you do not fill in both, then you must fill in `(kms_encrypted_access_key_id, kms_encrypted_access_key_secret, kms_encryption_access_key_id_context, kms_encryption_access_key_secret_context)` to use KMS to get the key pair.
     * 
     */
    @Export(name="roleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> roleArn;

    /**
     * @return Sts role info under delivery target logstore. `role_arn` and `(access_key_id, access_key_secret)` fill in at most one. If you do not fill in both, then you must fill in `(kms_encrypted_access_key_id, kms_encrypted_access_key_secret, kms_encryption_access_key_id_context, kms_encryption_access_key_secret_context)` to use KMS to get the key pair.
     * 
     */
    public Output<Optional<String>> roleArn() {
        return Codegen.optional(this.roleArn);
    }
    /**
     * Job scheduling type, the default value is Resident.
     * 
     */
    @Export(name="schedule", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> schedule;

    /**
     * @return Job scheduling type, the default value is Resident.
     * 
     */
    public Output<Optional<String>> schedule() {
        return Codegen.optional(this.schedule);
    }
    /**
     * Processing operation grammar.
     * 
     */
    @Export(name="script", refs={String.class}, tree="[0]")
    private Output<String> script;

    /**
     * @return Processing operation grammar.
     * 
     */
    public Output<String> script() {
        return this.script;
    }
    /**
     * Log project tags. the default value is RUNNING, Only 4 values are supported: `STARTING`，`RUNNING`，`STOPPING`，`STOPPED`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Log project tags. the default value is RUNNING, Only 4 values are supported: `STARTING`，`RUNNING`，`STOPPING`，`STOPPED`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Deadline of processing job, if not set the value is 0, indicates that new data will be processed continuously.
     * 
     */
    @Export(name="toTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> toTime;

    /**
     * @return Deadline of processing job, if not set the value is 0, indicates that new data will be processed continuously.
     * 
     */
    public Output<Optional<Integer>> toTime() {
        return Codegen.optional(this.toTime);
    }
    /**
     * Log etl job version. the default value is `2`.
     * 
     */
    @Export(name="version", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> version;

    /**
     * @return Log etl job version. the default value is `2`.
     * 
     */
    public Output<Optional<Integer>> version() {
        return Codegen.optional(this.version);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Etl(String name) {
        this(name, EtlArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Etl(String name, EtlArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Etl(String name, EtlArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/etl:Etl", name, args == null ? EtlArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Etl(String name, Output<String> id, @Nullable EtlState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/etl:Etl", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accessKeyId",
                "accessKeySecret"
            ))
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
    public static Etl get(String name, Output<String> id, @Nullable EtlState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Etl(name, id, state, options);
    }
}
