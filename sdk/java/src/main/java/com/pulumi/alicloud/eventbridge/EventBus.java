// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eventbridge;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eventbridge.EventBusArgs;
import com.pulumi.alicloud.eventbridge.inputs.EventBusState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Event Bridge Event Bus resource.
 * 
 * For information about Event Bridge Event Bus and how to use it, see [What is Event Bus](https://help.aliyun.com/document_detail/167863.html).
 * 
 * &gt; **NOTE:** Available in v1.129.0+.
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
 * import com.pulumi.alicloud.eventbridge.EventBus;
 * import com.pulumi.alicloud.eventbridge.EventBusArgs;
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
 *         var example = new EventBus(&#34;example&#34;, EventBusArgs.builder()        
 *             .eventBusName(&#34;my-EventBus&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Event Bridge Event Bus can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:eventbridge/eventBus:EventBus example &lt;event_bus_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:eventbridge/eventBus:EventBus")
public class EventBus extends com.pulumi.resources.CustomResource {
    /**
     * The description of event bus.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of event bus.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The name of event bus. The length is limited to 2 ~ 127 characters, which can be composed of letters, numbers or hyphens (-)
     * 
     */
    @Export(name="eventBusName", type=String.class, parameters={})
    private Output<String> eventBusName;

    /**
     * @return The name of event bus. The length is limited to 2 ~ 127 characters, which can be composed of letters, numbers or hyphens (-)
     * 
     */
    public Output<String> eventBusName() {
        return this.eventBusName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EventBus(String name) {
        this(name, EventBusArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EventBus(String name, EventBusArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EventBus(String name, EventBusArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eventbridge/eventBus:EventBus", name, args == null ? EventBusArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EventBus(String name, Output<String> id, @Nullable EventBusState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eventbridge/eventBus:EventBus", name, state, makeResourceOptions(options, id));
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
    public static EventBus get(String name, Output<String> id, @Nullable EventBusState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EventBus(name, id, state, options);
    }
}
