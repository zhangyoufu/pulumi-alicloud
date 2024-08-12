// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rocketmq;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.rocketmq.TopicArgs;
import com.pulumi.alicloud.rocketmq.inputs.TopicState;
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
 * Provides an ONS topic resource.
 * 
 * For more information about how to use it, see [RocketMQ Topic Management API](https://www.alibabacloud.com/help/doc-detail/29591.html).
 * 
 * &gt; **NOTE:** Available in 1.53.0+
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
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.rocketmq.Instance;
 * import com.pulumi.alicloud.rocketmq.InstanceArgs;
 * import com.pulumi.alicloud.rocketmq.Topic;
 * import com.pulumi.alicloud.rocketmq.TopicArgs;
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
 *         final var name = config.get("name").orElse("onsInstanceName");
 *         final var topic = config.get("topic").orElse("onsTopicName");
 *         var default_ = new Integer("default", IntegerArgs.builder()
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var defaultInstance = new Instance("defaultInstance", InstanceArgs.builder()
 *             .name(String.format("%s-%s", name,default_.result()))
 *             .remark("default_ons_instance_remark")
 *             .build());
 * 
 *         var defaultTopic = new Topic("defaultTopic", TopicArgs.builder()
 *             .topicName(topic)
 *             .instanceId(defaultInstance.id())
 *             .messageType(0)
 *             .remark("dafault_ons_topic_remark")
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
 * ONS TOPIC can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:rocketmq/topic:Topic topic MQ_INST_1234567890_Baso1234567:onsTopicDemo
 * ```
 * 
 */
@ResourceType(type="alicloud:rocketmq/topic:Topic")
public class Topic extends com.pulumi.resources.CustomResource {
    /**
     * ID of the ONS Instance that owns the topics.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return ID of the ONS Instance that owns the topics.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The type of the message. Read [Ons Topic Create](https://www.alibabacloud.com/help/doc-detail/29591.html) for further details.
     * 
     */
    @Export(name="messageType", refs={Integer.class}, tree="[0]")
    private Output<Integer> messageType;

    /**
     * @return The type of the message. Read [Ons Topic Create](https://www.alibabacloud.com/help/doc-detail/29591.html) for further details.
     * 
     */
    public Output<Integer> messageType() {
        return this.messageType;
    }
    /**
     * This attribute has been deprecated.
     * 
     * @deprecated
     * Attribute perm has been deprecated and suggest removing it from your template.
     * 
     */
    @Deprecated /* Attribute perm has been deprecated and suggest removing it from your template. */
    @Export(name="perm", refs={Integer.class}, tree="[0]")
    private Output<Integer> perm;

    /**
     * @return This attribute has been deprecated.
     * 
     */
    public Output<Integer> perm() {
        return this.perm;
    }
    /**
     * This attribute is a concise description of topic. The length cannot exceed 128.
     * 
     */
    @Export(name="remark", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> remark;

    /**
     * @return This attribute is a concise description of topic. The length cannot exceed 128.
     * 
     */
    public Output<Optional<String>> remark() {
        return Codegen.optional(this.remark);
    }
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     * &gt; **NOTE:** At least one of `topic_name` and `topic` should be set.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     * &gt; **NOTE:** At least one of `topic_name` and `topic` should be set.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Replaced by `topic_name` after version 1.97.0.
     * 
     * @deprecated
     * Field &#39;topic&#39; has been deprecated from version 1.97.0. Use &#39;topic_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'topic' has been deprecated from version 1.97.0. Use 'topic_name' instead. */
    @Export(name="topic", refs={String.class}, tree="[0]")
    private Output<String> topic;

    /**
     * @return Replaced by `topic_name` after version 1.97.0.
     * 
     */
    public Output<String> topic() {
        return this.topic;
    }
    /**
     * Name of the topic. Two topics on a single instance cannot have the same name and the name cannot start with &#39;GID&#39; or &#39;CID&#39;. The length cannot exceed 64 characters.
     * 
     */
    @Export(name="topicName", refs={String.class}, tree="[0]")
    private Output<String> topicName;

    /**
     * @return Name of the topic. Two topics on a single instance cannot have the same name and the name cannot start with &#39;GID&#39; or &#39;CID&#39;. The length cannot exceed 64 characters.
     * 
     */
    public Output<String> topicName() {
        return this.topicName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Topic(java.lang.String name) {
        this(name, TopicArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Topic(java.lang.String name, TopicArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Topic(java.lang.String name, TopicArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rocketmq/topic:Topic", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Topic(java.lang.String name, Output<java.lang.String> id, @Nullable TopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rocketmq/topic:Topic", name, state, makeResourceOptions(options, id), false);
    }

    private static TopicArgs makeArgs(TopicArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TopicArgs.Empty : args;
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
    public static Topic get(java.lang.String name, Output<java.lang.String> id, @Nullable TopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Topic(name, id, state, options);
    }
}
