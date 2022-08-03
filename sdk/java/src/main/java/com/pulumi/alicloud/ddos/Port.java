// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ddos;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ddos.PortArgs;
import com.pulumi.alicloud.ddos.inputs.PortState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Anti-DDoS Pro Port resource.
 * 
 * For information about Anti-DDoS Pro Port and how to use it, see [What is Port](https://www.alibabacloud.com/help/en/doc-detail/157482.htm).
 * 
 * &gt; **NOTE:** Available in v1.123.0+.
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
 * import com.pulumi.alicloud.ddos.DdosCooInstance;
 * import com.pulumi.alicloud.ddos.DdosCooInstanceArgs;
 * import com.pulumi.alicloud.ddos.Port;
 * import com.pulumi.alicloud.ddos.PortArgs;
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
 *         var exampleDdosCooInstance = new DdosCooInstance(&#34;exampleDdosCooInstance&#34;, DdosCooInstanceArgs.builder()        
 *             .bandwidth(&#34;30&#34;)
 *             .baseBandwidth(&#34;30&#34;)
 *             .serviceBandwidth(&#34;100&#34;)
 *             .portCount(&#34;50&#34;)
 *             .domainCount(&#34;50&#34;)
 *             .build());
 * 
 *         var examplePort = new Port(&#34;examplePort&#34;, PortArgs.builder()        
 *             .instanceId(exampleDdosCooInstance.id())
 *             .frontendPort(&#34;7001&#34;)
 *             .frontendProtocol(&#34;tcp&#34;)
 *             .realServers(            
 *                 &#34;1.1.1.1&#34;,
 *                 &#34;2.2.2.2&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Anti-DDoS Pro Port can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ddos/port:Port example &lt;instance_id&gt;:&lt;frontend_port&gt;:&lt;frontend_protocol&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ddos/port:Port")
public class Port extends com.pulumi.resources.CustomResource {
    /**
     * The port of the origin server. Valid values: [1~65535].
     * 
     */
    @Export(name="backendPort", type=String.class, parameters={})
    private Output</* @Nullable */ String> backendPort;

    /**
     * @return The port of the origin server. Valid values: [1~65535].
     * 
     */
    public Output<Optional<String>> backendPort() {
        return Codegen.optional(this.backendPort);
    }
    /**
     * The forwarding port. Valid values: [1~65535].
     * 
     */
    @Export(name="frontendPort", type=String.class, parameters={})
    private Output<String> frontendPort;

    /**
     * @return The forwarding port. Valid values: [1~65535].
     * 
     */
    public Output<String> frontendPort() {
        return this.frontendPort;
    }
    /**
     * The forwarding protocol. Valid values `tcp` and `udp`.
     * 
     */
    @Export(name="frontendProtocol", type=String.class, parameters={})
    private Output<String> frontendProtocol;

    /**
     * @return The forwarding protocol. Valid values `tcp` and `udp`.
     * 
     */
    public Output<String> frontendProtocol() {
        return this.frontendProtocol;
    }
    /**
     * The ID of Ddoscoo instance.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return The ID of Ddoscoo instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * List of source IP addresses.
     * 
     */
    @Export(name="realServers", type=List.class, parameters={String.class})
    private Output<List<String>> realServers;

    /**
     * @return List of source IP addresses.
     * 
     */
    public Output<List<String>> realServers() {
        return this.realServers;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Port(String name) {
        this(name, PortArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Port(String name, PortArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Port(String name, PortArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ddos/port:Port", name, args == null ? PortArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Port(String name, Output<String> id, @Nullable PortState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ddos/port:Port", name, state, makeResourceOptions(options, id));
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
    public static Port get(String name, Output<String> id, @Nullable PortState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Port(name, id, state, options);
    }
}
