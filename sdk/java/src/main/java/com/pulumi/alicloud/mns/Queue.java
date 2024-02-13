// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mns;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.mns.QueueArgs;
import com.pulumi.alicloud.mns.inputs.QueueState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.mns.Queue;
 * import com.pulumi.alicloud.mns.QueueArgs;
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
 *         var queue = new Queue(&#34;queue&#34;, QueueArgs.builder()        
 *             .delaySeconds(0)
 *             .maximumMessageSize(65536)
 *             .messageRetentionPeriod(345600)
 *             .pollingWaitSeconds(0)
 *             .visibilityTimeout(30)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * MNS QUEUE can be imported using the id or name, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:mns/queue:Queue queue queuename
 * ```
 * 
 */
@ResourceType(type="alicloud:mns/queue:Queue")
public class Queue extends com.pulumi.resources.CustomResource {
    /**
     * This attribute defines the length of time, in seconds, after which every message sent to the queue is dequeued. Valid value range: 0-604800 seconds, i.e., 0 to 7 days. Default value to 0.
     * 
     */
    @Export(name="delaySeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> delaySeconds;

    /**
     * @return This attribute defines the length of time, in seconds, after which every message sent to the queue is dequeued. Valid value range: 0-604800 seconds, i.e., 0 to 7 days. Default value to 0.
     * 
     */
    public Output<Optional<Integer>> delaySeconds() {
        return Codegen.optional(this.delaySeconds);
    }
    /**
     * This indicates the maximum length, in bytes, of any message body sent to the queue. Valid value range: 1024-65536, i.e., 1K to 64K. Default value to 65536.
     * 
     */
    @Export(name="maximumMessageSize", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maximumMessageSize;

    /**
     * @return This indicates the maximum length, in bytes, of any message body sent to the queue. Valid value range: 1024-65536, i.e., 1K to 64K. Default value to 65536.
     * 
     */
    public Output<Optional<Integer>> maximumMessageSize() {
        return Codegen.optional(this.maximumMessageSize);
    }
    /**
     * Messages are deleted from the queue after a specified length of time, whether they have been activated or not. This attribute defines the viability period, in seconds, for every message in the queue. Valid value range: 60-604800 seconds, i.e., 1 minutes to 7 days. Default value to 345600.
     * 
     */
    @Export(name="messageRetentionPeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> messageRetentionPeriod;

    /**
     * @return Messages are deleted from the queue after a specified length of time, whether they have been activated or not. This attribute defines the viability period, in seconds, for every message in the queue. Valid value range: 60-604800 seconds, i.e., 1 minutes to 7 days. Default value to 345600.
     * 
     */
    public Output<Optional<Integer>> messageRetentionPeriod() {
        return Codegen.optional(this.messageRetentionPeriod);
    }
    /**
     * Two queues on a single account in the same region cannot have the same name. A queue name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters .
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Two queues on a single account in the same region cannot have the same name. A queue name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters .
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Long polling is measured in seconds. When this attribute is set to 0, long polling is disabled. When it is not set to 0, long polling is enabled and message dequeue requests will be processed only when valid messages are received or when long polling times out. Valid value range: 0-30 seconds. Default value to 0.
     * 
     */
    @Export(name="pollingWaitSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> pollingWaitSeconds;

    /**
     * @return Long polling is measured in seconds. When this attribute is set to 0, long polling is disabled. When it is not set to 0, long polling is enabled and message dequeue requests will be processed only when valid messages are received or when long polling times out. Valid value range: 0-30 seconds. Default value to 0.
     * 
     */
    public Output<Optional<Integer>> pollingWaitSeconds() {
        return Codegen.optional(this.pollingWaitSeconds);
    }
    /**
     * The VisibilityTimeout attribute of the queue. A dequeued messages will change from active (visible) status to inactive (invisible) status, and this attribute defines the length of time, in seconds, that messages remain invisible. Messages return to active status after the set period. Valid value range: 1-43200 seconds, i.e., 1 seconds to 12 hours. Default value to 30.
     * 
     */
    @Export(name="visibilityTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> visibilityTimeout;

    /**
     * @return The VisibilityTimeout attribute of the queue. A dequeued messages will change from active (visible) status to inactive (invisible) status, and this attribute defines the length of time, in seconds, that messages remain invisible. Messages return to active status after the set period. Valid value range: 1-43200 seconds, i.e., 1 seconds to 12 hours. Default value to 30.
     * 
     */
    public Output<Optional<Integer>> visibilityTimeout() {
        return Codegen.optional(this.visibilityTimeout);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Queue(String name) {
        this(name, QueueArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Queue(String name, @Nullable QueueArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Queue(String name, @Nullable QueueArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mns/queue:Queue", name, args == null ? QueueArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Queue(String name, Output<String> id, @Nullable QueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mns/queue:Queue", name, state, makeResourceOptions(options, id));
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
    public static Queue get(String name, Output<String> id, @Nullable QueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Queue(name, id, state, options);
    }
}
