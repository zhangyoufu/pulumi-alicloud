// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.datahub;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.datahub.TopicArgs;
import com.pulumi.alicloud.datahub.inputs.TopicState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The topic is the basic unit of Datahub data source and is used to define one kind of data or stream. It contains a set of subscriptions. You can manage the datahub source of an application by using topics. [Refer to details](https://www.alibabacloud.com/help/en/datahub/latest/nerbcz).
 * 
 * &gt; **NOTE:** Available since v1.19.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.datahub.Project;
 * import com.pulumi.alicloud.datahub.ProjectArgs;
 * import com.pulumi.alicloud.datahub.Topic;
 * import com.pulumi.alicloud.datahub.TopicArgs;
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
 *         final var config = ctx.config();
 *         final var name = config.get("name").orElse("tf_example");
 *         var example = new Project("example", ProjectArgs.builder()        
 *             .name(name)
 *             .comment("created by terraform")
 *             .build());
 * 
 *         var exampleBlob = new Topic("exampleBlob", TopicArgs.builder()        
 *             .name(String.format("%s_blob", name))
 *             .projectName(example.name())
 *             .recordType("BLOB")
 *             .shardCount(3)
 *             .lifeCycle(7)
 *             .comment("created by terraform")
 *             .build());
 * 
 *         var exampleTuple = new Topic("exampleTuple", TopicArgs.builder()        
 *             .name(String.format("%s_tuple", name))
 *             .projectName(example.name())
 *             .recordType("TUPLE")
 *             .recordSchema(Map.ofEntries(
 *                 Map.entry("bigint_field", "BIGINT"),
 *                 Map.entry("timestamp_field", "TIMESTAMP"),
 *                 Map.entry("string_field", "STRING"),
 *                 Map.entry("double_field", "DOUBLE"),
 *                 Map.entry("boolean_field", "BOOLEAN")
 *             ))
 *             .shardCount(3)
 *             .lifeCycle(7)
 *             .comment("created by terraform")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Datahub topic can be imported using the ID, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:datahub/topic:Topic example tf_datahub_project:tf_datahub_topic
 * ```
 * 
 */
@ResourceType(type="alicloud:datahub/topic:Topic")
public class Topic extends com.pulumi.resources.CustomResource {
    /**
     * Comment of the datahub topic. It cannot be longer than 255 characters.
     * 
     * **Notes:** Currently `life_cycle` can not be modified and it will be supported in the next future.
     * 
     */
    @Export(name="comment", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> comment;

    /**
     * @return Comment of the datahub topic. It cannot be longer than 255 characters.
     * 
     * **Notes:** Currently `life_cycle` can not be modified and it will be supported in the next future.
     * 
     */
    public Output<Optional<String>> comment() {
        return Codegen.optional(this.comment);
    }
    /**
     * Create time of the datahub topic. It is a human-readable string rather than 64-bits UTC.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Create time of the datahub topic. It is a human-readable string rather than 64-bits UTC.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Last modify time of the datahub topic. It is the same as *create_time* at the beginning. It is also a human-readable string rather than 64-bits UTC.
     * 
     */
    @Export(name="lastModifyTime", refs={String.class}, tree="[0]")
    private Output<String> lastModifyTime;

    /**
     * @return Last modify time of the datahub topic. It is the same as *create_time* at the beginning. It is also a human-readable string rather than 64-bits UTC.
     * 
     */
    public Output<String> lastModifyTime() {
        return this.lastModifyTime;
    }
    /**
     * How many days this topic lives. The permitted range of values is [1, 7]. The default value is 3.
     * 
     */
    @Export(name="lifeCycle", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> lifeCycle;

    /**
     * @return How many days this topic lives. The permitted range of values is [1, 7]. The default value is 3.
     * 
     */
    public Output<Optional<Integer>> lifeCycle() {
        return Codegen.optional(this.lifeCycle);
    }
    /**
     * The name of the datahub topic. Its length is limited to 1-128 and only characters such as letters, digits and &#39;_&#39; are allowed. It is case-insensitive.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the datahub topic. Its length is limited to 1-128 and only characters such as letters, digits and &#39;_&#39; are allowed. It is case-insensitive.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The name of the datahub project that this topic belongs to. It is case-insensitive.
     * 
     */
    @Export(name="projectName", refs={String.class}, tree="[0]")
    private Output<String> projectName;

    /**
     * @return The name of the datahub project that this topic belongs to. It is case-insensitive.
     * 
     */
    public Output<String> projectName() {
        return this.projectName;
    }
    /**
     * Schema of this topic, required only for TUPLE topic. Supported data types (case-insensitive) are:
     * - BIGINT
     * - STRING
     * - BOOLEAN
     * - DOUBLE
     * - TIMESTAMP
     * 
     */
    @Export(name="recordSchema", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> recordSchema;

    /**
     * @return Schema of this topic, required only for TUPLE topic. Supported data types (case-insensitive) are:
     * - BIGINT
     * - STRING
     * - BOOLEAN
     * - DOUBLE
     * - TIMESTAMP
     * 
     */
    public Output<Optional<Map<String,Object>>> recordSchema() {
        return Codegen.optional(this.recordSchema);
    }
    /**
     * The type of this topic. Its value must be one of {BLOB, TUPLE}. For BLOB topic, data will be organized as binary and encoded by BASE64. For TUPLE topic, data has fixed schema. The default value is &#34;TUPLE&#34; with a schema {STRING}.
     * 
     */
    @Export(name="recordType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> recordType;

    /**
     * @return The type of this topic. Its value must be one of {BLOB, TUPLE}. For BLOB topic, data will be organized as binary and encoded by BASE64. For TUPLE topic, data has fixed schema. The default value is &#34;TUPLE&#34; with a schema {STRING}.
     * 
     */
    public Output<Optional<String>> recordType() {
        return Codegen.optional(this.recordType);
    }
    /**
     * The number of shards this topic contains. The permitted range of values is [1, 10]. The default value is 1.
     * 
     */
    @Export(name="shardCount", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> shardCount;

    /**
     * @return The number of shards this topic contains. The permitted range of values is [1, 10]. The default value is 1.
     * 
     */
    public Output<Optional<Integer>> shardCount() {
        return Codegen.optional(this.shardCount);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Topic(String name) {
        this(name, TopicArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Topic(String name, TopicArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Topic(String name, TopicArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:datahub/topic:Topic", name, args == null ? TopicArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Topic(String name, Output<String> id, @Nullable TopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:datahub/topic:Topic", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Topic get(String name, Output<String> id, @Nullable TopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Topic(name, id, state, options);
    }
}
