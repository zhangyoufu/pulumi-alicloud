// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.resourcemanager.SavedQueryArgs;
import com.pulumi.alicloud.resourcemanager.inputs.SavedQueryState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Resource Manager Saved Query resource. ResourceCenter Saved Query.
 * 
 * For information about Resource Manager Saved Query and how to use it, see [What is Saved Query](https://www.alibabacloud.com/help/zh/resource-management/developer-reference/api-resourcecenter-2022-12-01-createsavedquery).
 * 
 * &gt; **NOTE:** Available since v1.212.0.
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
 * import com.pulumi.alicloud.resourcemanager.SavedQuery;
 * import com.pulumi.alicloud.resourcemanager.SavedQueryArgs;
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
 *         var default_ = new SavedQuery("default", SavedQueryArgs.builder()
 *             .description(name)
 *             .expression("select * from resources limit 1;")
 *             .savedQueryName(name)
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
 * Resource Manager Saved Query can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:resourcemanager/savedQuery:SavedQuery example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:resourcemanager/savedQuery:SavedQuery")
public class SavedQuery extends com.pulumi.resources.CustomResource {
    /**
     * The creation time of the resource.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The creation time of the resource.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Query Description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Query Description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Query Expression.
     * 
     */
    @Export(name="expression", refs={String.class}, tree="[0]")
    private Output<String> expression;

    /**
     * @return Query Expression.
     * 
     */
    public Output<String> expression() {
        return this.expression;
    }
    /**
     * The name of the resource.
     * 
     */
    @Export(name="savedQueryName", refs={String.class}, tree="[0]")
    private Output<String> savedQueryName;

    /**
     * @return The name of the resource.
     * 
     */
    public Output<String> savedQueryName() {
        return this.savedQueryName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SavedQuery(java.lang.String name) {
        this(name, SavedQueryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SavedQuery(java.lang.String name, SavedQueryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SavedQuery(java.lang.String name, SavedQueryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:resourcemanager/savedQuery:SavedQuery", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SavedQuery(java.lang.String name, Output<java.lang.String> id, @Nullable SavedQueryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:resourcemanager/savedQuery:SavedQuery", name, state, makeResourceOptions(options, id), false);
    }

    private static SavedQueryArgs makeArgs(SavedQueryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SavedQueryArgs.Empty : args;
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
    public static SavedQuery get(java.lang.String name, Output<java.lang.String> id, @Nullable SavedQueryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SavedQuery(name, id, state, options);
    }
}
