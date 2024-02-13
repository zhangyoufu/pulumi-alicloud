// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.datahub;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.datahub.SubscriptionArgs;
import com.pulumi.alicloud.datahub.inputs.SubscriptionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The subscription is the basic unit of resource usage in Datahub Service under Publish/Subscribe model. You can manage the relationships between user and topics by using subscriptions. [Refer to details](https://www.alibabacloud.com/help/en/datahub/latest/nerbcz).
 * 
 * &gt; **NOTE:** Available since v1.19.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.datahub.Project;
 * import com.pulumi.alicloud.datahub.ProjectArgs;
 * import com.pulumi.alicloud.datahub.Topic;
 * import com.pulumi.alicloud.datahub.TopicArgs;
 * import com.pulumi.alicloud.datahub.Subscription;
 * import com.pulumi.alicloud.datahub.SubscriptionArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform_example&#34;);
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .comment(&#34;created by terraform&#34;)
 *             .build());
 * 
 *         var exampleTopic = new Topic(&#34;exampleTopic&#34;, TopicArgs.builder()        
 *             .projectName(exampleProject.name())
 *             .recordType(&#34;BLOB&#34;)
 *             .shardCount(3)
 *             .lifeCycle(7)
 *             .comment(&#34;created by terraform&#34;)
 *             .build());
 * 
 *         var exampleSubscription = new Subscription(&#34;exampleSubscription&#34;, SubscriptionArgs.builder()        
 *             .projectName(exampleProject.name())
 *             .topicName(exampleTopic.name())
 *             .comment(&#34;created by terraform&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Datahub subscription can be imported using the ID, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:datahub/subscription:Subscription example tf_datahub_project:tf_datahub_topic:1539073399567UgCzY
 * ```
 * 
 */
@ResourceType(type="alicloud:datahub/subscription:Subscription")
public class Subscription extends com.pulumi.resources.CustomResource {
    /**
     * Comment of the datahub subscription. It cannot be longer than 255 characters.
     * 
     */
    @Export(name="comment", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> comment;

    /**
     * @return Comment of the datahub subscription. It cannot be longer than 255 characters.
     * 
     */
    public Output<Optional<String>> comment() {
        return Codegen.optional(this.comment);
    }
    /**
     * Create time of the datahub subscription. It is a human-readable string rather than 64-bits UTC.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Create time of the datahub subscription. It is a human-readable string rather than 64-bits UTC.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Last modify time of the datahub subscription. It is the same as *create_time* at the beginning. It is also a human-readable string rather than 64-bits UTC.
     * 
     */
    @Export(name="lastModifyTime", refs={String.class}, tree="[0]")
    private Output<String> lastModifyTime;

    /**
     * @return Last modify time of the datahub subscription. It is the same as *create_time* at the beginning. It is also a human-readable string rather than 64-bits UTC.
     * 
     */
    public Output<String> lastModifyTime() {
        return this.lastModifyTime;
    }
    /**
     * The name of the datahub project that the subscription belongs to. Its length is limited to 3-32 and only characters such as letters, digits and &#39;_&#39; are allowed. It is case-insensitive.
     * 
     */
    @Export(name="projectName", refs={String.class}, tree="[0]")
    private Output<String> projectName;

    /**
     * @return The name of the datahub project that the subscription belongs to. Its length is limited to 3-32 and only characters such as letters, digits and &#39;_&#39; are allowed. It is case-insensitive.
     * 
     */
    public Output<String> projectName() {
        return this.projectName;
    }
    /**
     * The identidy of the subscription, generate from server side.
     * 
     */
    @Export(name="subId", refs={String.class}, tree="[0]")
    private Output<String> subId;

    /**
     * @return The identidy of the subscription, generate from server side.
     * 
     */
    public Output<String> subId() {
        return this.subId;
    }
    /**
     * The name of the datahub topic that the subscription belongs to. Its length is limited to 1-128 and only characters such as letters, digits and &#39;_&#39; are allowed. It is case-insensitive.
     * 
     */
    @Export(name="topicName", refs={String.class}, tree="[0]")
    private Output<String> topicName;

    /**
     * @return The name of the datahub topic that the subscription belongs to. Its length is limited to 1-128 and only characters such as letters, digits and &#39;_&#39; are allowed. It is case-insensitive.
     * 
     */
    public Output<String> topicName() {
        return this.topicName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Subscription(String name) {
        this(name, SubscriptionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Subscription(String name, SubscriptionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Subscription(String name, SubscriptionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:datahub/subscription:Subscription", name, args == null ? SubscriptionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Subscription(String name, Output<String> id, @Nullable SubscriptionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:datahub/subscription:Subscription", name, state, makeResourceOptions(options, id));
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
    public static Subscription get(String name, Output<String> id, @Nullable SubscriptionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Subscription(name, id, state, options);
    }
}
