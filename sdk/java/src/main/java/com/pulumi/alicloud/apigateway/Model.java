// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.apigateway;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.apigateway.ModelArgs;
import com.pulumi.alicloud.apigateway.inputs.ModelState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Api Gateway Model resource.
 * 
 * For information about Api Gateway Model and how to use it, see [What is Model](https://www.alibabacloud.com/help/en/api-gateway/latest/api-cloudapi-2016-07-14-createmodel).
 * 
 * &gt; **NOTE:** Available since v1.187.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.apigateway.Group;
 * import com.pulumi.alicloud.apigateway.GroupArgs;
 * import com.pulumi.alicloud.apigateway.Model;
 * import com.pulumi.alicloud.apigateway.ModelArgs;
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
 *         var defaultGroup = new Group(&#34;defaultGroup&#34;, GroupArgs.builder()        
 *             .description(&#34;example_value&#34;)
 *             .build());
 * 
 *         var defaultModel = new Model(&#34;defaultModel&#34;, ModelArgs.builder()        
 *             .groupId(defaultGroup.id())
 *             .modelName(&#34;example_value&#34;)
 *             .schema(&#34;{\&#34;type\&#34;:\&#34;object\&#34;,\&#34;properties\&#34;:{\&#34;id\&#34;:{\&#34;format\&#34;:\&#34;int64\&#34;,\&#34;maximum\&#34;:100,\&#34;exclusiveMaximum\&#34;:true,\&#34;type\&#34;:\&#34;integer\&#34;},\&#34;name\&#34;:{\&#34;maxLength\&#34;:10,\&#34;type\&#34;:\&#34;string\&#34;}}}&#34;)
 *             .description(&#34;example_value&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Api Gateway Model can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:apigateway/model:Model example &lt;group_id&gt;:&lt;model_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:apigateway/model:Model")
public class Model extends com.pulumi.resources.CustomResource {
    /**
     * The description of the model.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the model.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The group of the model belongs to.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output<String> groupId;

    /**
     * @return The group of the model belongs to.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }
    /**
     * The name of the model.
     * 
     */
    @Export(name="modelName", refs={String.class}, tree="[0]")
    private Output<String> modelName;

    /**
     * @return The name of the model.
     * 
     */
    public Output<String> modelName() {
        return this.modelName;
    }
    /**
     * The schema of the model.
     * 
     */
    @Export(name="schema", refs={String.class}, tree="[0]")
    private Output<String> schema;

    /**
     * @return The schema of the model.
     * 
     */
    public Output<String> schema() {
        return this.schema;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Model(String name) {
        this(name, ModelArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Model(String name, ModelArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Model(String name, ModelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:apigateway/model:Model", name, args == null ? ModelArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Model(String name, Output<String> id, @Nullable ModelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:apigateway/model:Model", name, state, makeResourceOptions(options, id));
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
    public static Model get(String name, Output<String> id, @Nullable ModelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Model(name, id, state, options);
    }
}
