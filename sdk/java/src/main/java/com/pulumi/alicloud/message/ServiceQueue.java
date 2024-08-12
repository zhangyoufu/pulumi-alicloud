// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.message;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.message.ServiceQueueArgs;
import com.pulumi.alicloud.message.inputs.ServiceQueueState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Message Service Queue resource.
 * 
 * For information about Message Service Queue and how to use it, see [What is Queue](https://www.alibabacloud.com/help/en/message-service/latest/createqueue).
 * 
 * &gt; **NOTE:** Available since v1.188.0.
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
 * import com.pulumi.alicloud.message.ServiceQueue;
 * import com.pulumi.alicloud.message.ServiceQueueArgs;
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
 *         final var name = config.get("name").orElse("terraform-example");
 *         var default_ = new ServiceQueue("default", ServiceQueueArgs.builder()
 *             .delaySeconds("2")
 *             .pollingWaitSeconds("2")
 *             .messageRetentionPeriod("566")
 *             .maximumMessageSize("1123")
 *             .visibilityTimeout("30")
 *             .queueName(name)
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
 * Message Service Queue can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:message/serviceQueue:ServiceQueue example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:message/serviceQueue:ServiceQueue")
public class ServiceQueue extends com.pulumi.resources.CustomResource {
    /**
     * Represents the time when the Queue was created.
     * 
     */
    @Export(name="createTime", refs={Integer.class}, tree="[0]")
    private Output<Integer> createTime;

    /**
     * @return Represents the time when the Queue was created.
     * 
     */
    public Output<Integer> createTime() {
        return this.createTime;
    }
    /**
     * This means that messages sent to the queue can only be consumed after the delay time set by this parameter, in seconds.
     * 
     */
    @Export(name="delaySeconds", refs={Integer.class}, tree="[0]")
    private Output<Integer> delaySeconds;

    /**
     * @return This means that messages sent to the queue can only be consumed after the delay time set by this parameter, in seconds.
     * 
     */
    public Output<Integer> delaySeconds() {
        return this.delaySeconds;
    }
    /**
     * Represents whether the log management function is enabled.
     * 
     */
    @Export(name="loggingEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> loggingEnabled;

    /**
     * @return Represents whether the log management function is enabled.
     * 
     */
    public Output<Optional<Boolean>> loggingEnabled() {
        return Codegen.optional(this.loggingEnabled);
    }
    /**
     * Represents the maximum length of the message body sent to the Queue, in Byte.
     * 
     */
    @Export(name="maximumMessageSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> maximumMessageSize;

    /**
     * @return Represents the maximum length of the message body sent to the Queue, in Byte.
     * 
     */
    public Output<Integer> maximumMessageSize() {
        return this.maximumMessageSize;
    }
    /**
     * Represents the longest life time of the message in the Queue.
     * 
     */
    @Export(name="messageRetentionPeriod", refs={Integer.class}, tree="[0]")
    private Output<Integer> messageRetentionPeriod;

    /**
     * @return Represents the longest life time of the message in the Queue.
     * 
     */
    public Output<Integer> messageRetentionPeriod() {
        return this.messageRetentionPeriod;
    }
    /**
     * The longest waiting time for a Queue request when the number of messages is empty, in seconds.
     * 
     */
    @Export(name="pollingWaitSeconds", refs={Integer.class}, tree="[0]")
    private Output<Integer> pollingWaitSeconds;

    /**
     * @return The longest waiting time for a Queue request when the number of messages is empty, in seconds.
     * 
     */
    public Output<Integer> pollingWaitSeconds() {
        return this.pollingWaitSeconds;
    }
    /**
     * Representative resources.
     * 
     */
    @Export(name="queueName", refs={String.class}, tree="[0]")
    private Output<String> queueName;

    /**
     * @return Representative resources.
     * 
     */
    public Output<String> queueName() {
        return this.queueName;
    }
    /**
     * Represents the duration after the message is removed from the Queue and changed from the Active state to the Inactive state.
     * 
     */
    @Export(name="visibilityTimeout", refs={Integer.class}, tree="[0]")
    private Output<Integer> visibilityTimeout;

    /**
     * @return Represents the duration after the message is removed from the Queue and changed from the Active state to the Inactive state.
     * 
     */
    public Output<Integer> visibilityTimeout() {
        return this.visibilityTimeout;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceQueue(java.lang.String name) {
        this(name, ServiceQueueArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceQueue(java.lang.String name, ServiceQueueArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceQueue(java.lang.String name, ServiceQueueArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:message/serviceQueue:ServiceQueue", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceQueue(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceQueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:message/serviceQueue:ServiceQueue", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceQueueArgs makeArgs(ServiceQueueArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceQueueArgs.Empty : args;
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
    public static ServiceQueue get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceQueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceQueue(name, id, state, options);
    }
}
