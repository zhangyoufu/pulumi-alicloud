// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.log.ResourceArgs;
import com.pulumi.alicloud.log.inputs.ResourceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Log resource is a meta store service provided by log service, resource can be used to define meta store&#39;s table structure.
 * 
 * For information about SLS Resource and how to use it, see [Resource management](https://www.alibabacloud.com/help/en/doc-detail/207732.html)
 * 
 * &gt; **NOTE:** Available since v1.162.0. log resource region should be set a main region: cn-heyuan.
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
 * import com.pulumi.alicloud.log.Resource;
 * import com.pulumi.alicloud.log.ResourceArgs;
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
 *         var example = new Resource("example", ResourceArgs.builder()
 *             .type("userdefine")
 *             .name("user.tf.resource")
 *             .description("user tf resource desc")
 *             .extInfo("{}")
 *             .schema("""
 *     {
 *       "schema": [
 *         {
 *           "column": "col1",
 *           "desc": "col1   desc",
 *           "ext_info": {
 *           },
 *           "required": true,
 *           "type": "string"
 *         },
 *         {
 *           "column": "col2",
 *           "desc": "col2   desc",
 *           "ext_info": "optional",
 *           "required": true,
 *           "type": "string"
 *         }
 *       ]
 *     }
 *             """)
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
 * Log resource can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:log/resource:Resource example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:log/resource:Resource")
public class Resource extends com.pulumi.resources.CustomResource {
    /**
     * The meta store&#39;s description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The meta store&#39;s description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The ext info of meta store.
     * 
     */
    @Export(name="extInfo", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> extInfo;

    /**
     * @return The ext info of meta store.
     * 
     */
    public Output<Optional<String>> extInfo() {
        return Codegen.optional(this.extInfo);
    }
    /**
     * The meta store&#39;s name, can be used as table name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The meta store&#39;s name, can be used as table name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The meta store&#39;s schema info, which is json string format, used to define table&#39;s fields.
     * 
     */
    @Export(name="schema", refs={String.class}, tree="[0]")
    private Output<String> schema;

    /**
     * @return The meta store&#39;s schema info, which is json string format, used to define table&#39;s fields.
     * 
     */
    public Output<String> schema() {
        return this.schema;
    }
    /**
     * The meta store&#39;s type, userdefine e.g.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The meta store&#39;s type, userdefine e.g.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Resource(java.lang.String name) {
        this(name, ResourceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Resource(java.lang.String name, ResourceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Resource(java.lang.String name, ResourceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/resource:Resource", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Resource(java.lang.String name, Output<java.lang.String> id, @Nullable ResourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/resource:Resource", name, state, makeResourceOptions(options, id), false);
    }

    private static ResourceArgs makeArgs(ResourceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ResourceArgs.Empty : args;
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
    public static Resource get(java.lang.String name, Output<java.lang.String> id, @Nullable ResourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Resource(name, id, state, options);
    }
}
