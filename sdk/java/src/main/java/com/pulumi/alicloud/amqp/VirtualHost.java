// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.amqp;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.amqp.VirtualHostArgs;
import com.pulumi.alicloud.amqp.inputs.VirtualHostState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a RabbitMQ (AMQP) Virtual Host resource.
 * 
 * For information about RabbitMQ (AMQP) Virtual Host and how to use it, see [What is Virtual Host](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/createvirtualhost).
 * 
 * &gt; **NOTE:** Available since v1.126.0.
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
 * import com.pulumi.alicloud.amqp.Instance;
 * import com.pulumi.alicloud.amqp.InstanceArgs;
 * import com.pulumi.alicloud.amqp.VirtualHost;
 * import com.pulumi.alicloud.amqp.VirtualHostArgs;
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
 *         var default_ = new Instance("default", InstanceArgs.builder()
 *             .instanceType("professional")
 *             .maxTps(1000)
 *             .queueCapacity(50)
 *             .supportEip(true)
 *             .maxEipTps(128)
 *             .paymentType("Subscription")
 *             .period(1)
 *             .build());
 * 
 *         var defaultVirtualHost = new VirtualHost("defaultVirtualHost", VirtualHostArgs.builder()
 *             .instanceId(default_.id())
 *             .virtualHostName("tf-example")
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
 * RabbitMQ (AMQP) Virtual Host can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:amqp/virtualHost:VirtualHost example &lt;instance_id&gt;:&lt;virtual_host_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:amqp/virtualHost:VirtualHost")
public class VirtualHost extends com.pulumi.resources.CustomResource {
    /**
     * InstanceId.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return InstanceId.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * VirtualHostName.
     * 
     */
    @Export(name="virtualHostName", refs={String.class}, tree="[0]")
    private Output<String> virtualHostName;

    /**
     * @return VirtualHostName.
     * 
     */
    public Output<String> virtualHostName() {
        return this.virtualHostName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VirtualHost(java.lang.String name) {
        this(name, VirtualHostArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VirtualHost(java.lang.String name, VirtualHostArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VirtualHost(java.lang.String name, VirtualHostArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:amqp/virtualHost:VirtualHost", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private VirtualHost(java.lang.String name, Output<java.lang.String> id, @Nullable VirtualHostState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:amqp/virtualHost:VirtualHost", name, state, makeResourceOptions(options, id), false);
    }

    private static VirtualHostArgs makeArgs(VirtualHostArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VirtualHostArgs.Empty : args;
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
    public static VirtualHost get(java.lang.String name, Output<java.lang.String> id, @Nullable VirtualHostState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VirtualHost(name, id, state, options);
    }
}
