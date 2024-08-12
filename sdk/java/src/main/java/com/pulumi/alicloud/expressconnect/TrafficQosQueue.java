// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.expressconnect;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.expressconnect.TrafficQosQueueArgs;
import com.pulumi.alicloud.expressconnect.inputs.TrafficQosQueueState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Express Connect Traffic Qos Queue resource. Express Connect Traffic QoS Queue.
 * 
 * For information about Express Connect Traffic Qos Queue and how to use it, see [What is Traffic Qos Queue](https://next.api.alibabacloud.com/document/Vpc/2016-04-28/CreateExpressConnectTrafficQosQueue).
 * 
 * &gt; **NOTE:** Available since v1.224.0.
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
 * import com.pulumi.alicloud.expressconnect.ExpressconnectFunctions;
 * import com.pulumi.alicloud.expressconnect.inputs.GetPhysicalConnectionsArgs;
 * import com.pulumi.alicloud.expressconnect.TrafficQos;
 * import com.pulumi.alicloud.expressconnect.TrafficQosArgs;
 * import com.pulumi.alicloud.expressconnect.TrafficQosAssociation;
 * import com.pulumi.alicloud.expressconnect.TrafficQosAssociationArgs;
 * import com.pulumi.alicloud.expressconnect.TrafficQosQueue;
 * import com.pulumi.alicloud.expressconnect.TrafficQosQueueArgs;
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
 *         final var default = ExpressconnectFunctions.getPhysicalConnections(GetPhysicalConnectionsArgs.builder()
 *             .nameRegex("preserved-NODELETING")
 *             .build());
 * 
 *         var createQos = new TrafficQos("createQos", TrafficQosArgs.builder()
 *             .qosName(name)
 *             .qosDescription("terraform-example")
 *             .build());
 * 
 *         var associateQos = new TrafficQosAssociation("associateQos", TrafficQosAssociationArgs.builder()
 *             .instanceId(default_.ids()[1])
 *             .qosId(createQos.id())
 *             .instanceType("PHYSICALCONNECTION")
 *             .build());
 * 
 *         var createQosQueue = new TrafficQosQueue("createQosQueue", TrafficQosQueueArgs.builder()
 *             .qosId(createQos.id())
 *             .bandwidthPercent("60")
 *             .queueDescription("terraform-example")
 *             .queueName(name)
 *             .queueType("Medium")
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
 * Express Connect Traffic Qos Queue can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:expressconnect/trafficQosQueue:TrafficQosQueue example &lt;qos_id&gt;:&lt;queue_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:expressconnect/trafficQosQueue:TrafficQosQueue")
public class TrafficQosQueue extends com.pulumi.resources.CustomResource {
    /**
     * QoS queue bandwidth percentage.
     * 
     * - When the QoS queue type is **Medium**, this field must be entered. Valid values: 1 to 100.
     * - When the QoS queue type is **Default**, this field is &#34;-&#34;.
     * 
     */
    @Export(name="bandwidthPercent", refs={String.class}, tree="[0]")
    private Output<String> bandwidthPercent;

    /**
     * @return QoS queue bandwidth percentage.
     * 
     * - When the QoS queue type is **Medium**, this field must be entered. Valid values: 1 to 100.
     * - When the QoS queue type is **Default**, this field is &#34;-&#34;.
     * 
     */
    public Output<String> bandwidthPercent() {
        return this.bandwidthPercent;
    }
    /**
     * The QoS policy ID.
     * 
     */
    @Export(name="qosId", refs={String.class}, tree="[0]")
    private Output<String> qosId;

    /**
     * @return The QoS policy ID.
     * 
     */
    public Output<String> qosId() {
        return this.qosId;
    }
    /**
     * The description of the QoS queue.  The length is 0 to 256 characters and cannot start with &#39;http:// &#39;or &#39;https.
     * 
     */
    @Export(name="queueDescription", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> queueDescription;

    /**
     * @return The description of the QoS queue.  The length is 0 to 256 characters and cannot start with &#39;http:// &#39;or &#39;https.
     * 
     */
    public Output<Optional<String>> queueDescription() {
        return Codegen.optional(this.queueDescription);
    }
    /**
     * The QoS queue ID.
     * 
     */
    @Export(name="queueId", refs={String.class}, tree="[0]")
    private Output<String> queueId;

    /**
     * @return The QoS queue ID.
     * 
     */
    public Output<String> queueId() {
        return this.queueId;
    }
    /**
     * The name of the QoS queue.  The length is 0 to 128 characters and cannot start with &#39;http:// &#39;or &#39;https.
     * 
     */
    @Export(name="queueName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> queueName;

    /**
     * @return The name of the QoS queue.  The length is 0 to 128 characters and cannot start with &#39;http:// &#39;or &#39;https.
     * 
     */
    public Output<Optional<String>> queueName() {
        return Codegen.optional(this.queueName);
    }
    /**
     * QoS queue type, value:
     * - **High**: High priority queue.
     * - **Medium**: Normal priority queue.
     * - **Default**: the Default priority queue.
     * &gt; **NOTE:**  Default priority queue cannot be created.
     * 
     */
    @Export(name="queueType", refs={String.class}, tree="[0]")
    private Output<String> queueType;

    /**
     * @return QoS queue type, value:
     * - **High**: High priority queue.
     * - **Medium**: Normal priority queue.
     * - **Default**: the Default priority queue.
     * &gt; **NOTE:**  Default priority queue cannot be created.
     * 
     */
    public Output<String> queueType() {
        return this.queueType;
    }
    /**
     * The status of the resource.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TrafficQosQueue(java.lang.String name) {
        this(name, TrafficQosQueueArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TrafficQosQueue(java.lang.String name, TrafficQosQueueArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TrafficQosQueue(java.lang.String name, TrafficQosQueueArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:expressconnect/trafficQosQueue:TrafficQosQueue", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private TrafficQosQueue(java.lang.String name, Output<java.lang.String> id, @Nullable TrafficQosQueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:expressconnect/trafficQosQueue:TrafficQosQueue", name, state, makeResourceOptions(options, id), false);
    }

    private static TrafficQosQueueArgs makeArgs(TrafficQosQueueArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TrafficQosQueueArgs.Empty : args;
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
    public static TrafficQosQueue get(java.lang.String name, Output<java.lang.String> id, @Nullable TrafficQosQueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TrafficQosQueue(name, id, state, options);
    }
}
