// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cen.TransitRouterMulticastDomainArgs;
import com.pulumi.alicloud.cen.inputs.TransitRouterMulticastDomainState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Enterprise Network (CEN) Transit Router Multicast Domain resource.
 * 
 * For information about Cloud Enterprise Network (CEN) Transit Router Multicast Domain and how to use it, see [What is Transit Router Multicast Domain](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createtransitroutermulticastdomain).
 * 
 * &gt; **NOTE:** Available since v1.195.0.
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
 * import com.pulumi.alicloud.cen.Instance;
 * import com.pulumi.alicloud.cen.InstanceArgs;
 * import com.pulumi.alicloud.cen.TransitRouter;
 * import com.pulumi.alicloud.cen.TransitRouterArgs;
 * import com.pulumi.alicloud.cen.TransitRouterMulticastDomain;
 * import com.pulumi.alicloud.cen.TransitRouterMulticastDomainArgs;
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
 *         var example = new Instance("example", InstanceArgs.builder()        
 *             .cenInstanceName("tf_example")
 *             .description("an example for cen")
 *             .build());
 * 
 *         var exampleTransitRouter = new TransitRouter("exampleTransitRouter", TransitRouterArgs.builder()        
 *             .transitRouterName("tf_example")
 *             .cenId(example.id())
 *             .supportMulticast(true)
 *             .build());
 * 
 *         var exampleTransitRouterMulticastDomain = new TransitRouterMulticastDomain("exampleTransitRouterMulticastDomain", TransitRouterMulticastDomainArgs.builder()        
 *             .transitRouterId(exampleTransitRouter.transitRouterId())
 *             .transitRouterMulticastDomainName("tf_example")
 *             .transitRouterMulticastDomainDescription("tf_example")
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
 * Cloud Enterprise Network (CEN) Transit Router Multicast Domain can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cen/transitRouterMulticastDomain:TransitRouterMulticastDomain example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cen/transitRouterMulticastDomain:TransitRouterMulticastDomain")
public class TransitRouterMulticastDomain extends com.pulumi.resources.CustomResource {
    /**
     * The status of the Transit Router Multicast Domain.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the Transit Router Multicast Domain.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The ID of the transit router.
     * 
     */
    @Export(name="transitRouterId", refs={String.class}, tree="[0]")
    private Output<String> transitRouterId;

    /**
     * @return The ID of the transit router.
     * 
     */
    public Output<String> transitRouterId() {
        return this.transitRouterId;
    }
    /**
     * The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs ({@literal @}), underscores (_), and hyphens (-).
     * 
     */
    @Export(name="transitRouterMulticastDomainDescription", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> transitRouterMulticastDomainDescription;

    /**
     * @return The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs ({@literal @}), underscores (_), and hyphens (-).
     * 
     */
    public Output<Optional<String>> transitRouterMulticastDomainDescription() {
        return Codegen.optional(this.transitRouterMulticastDomainDescription);
    }
    /**
     * The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs ({@literal @}), underscores (_), and hyphens (-).
     * 
     */
    @Export(name="transitRouterMulticastDomainName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> transitRouterMulticastDomainName;

    /**
     * @return The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs ({@literal @}), underscores (_), and hyphens (-).
     * 
     */
    public Output<Optional<String>> transitRouterMulticastDomainName() {
        return Codegen.optional(this.transitRouterMulticastDomainName);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TransitRouterMulticastDomain(String name) {
        this(name, TransitRouterMulticastDomainArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TransitRouterMulticastDomain(String name, TransitRouterMulticastDomainArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TransitRouterMulticastDomain(String name, TransitRouterMulticastDomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/transitRouterMulticastDomain:TransitRouterMulticastDomain", name, args == null ? TransitRouterMulticastDomainArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TransitRouterMulticastDomain(String name, Output<String> id, @Nullable TransitRouterMulticastDomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/transitRouterMulticastDomain:TransitRouterMulticastDomain", name, state, makeResourceOptions(options, id));
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
    public static TransitRouterMulticastDomain get(String name, Output<String> id, @Nullable TransitRouterMulticastDomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TransitRouterMulticastDomain(name, id, state, options);
    }
}
