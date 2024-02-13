// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.polardb.ParameterGroupArgs;
import com.pulumi.alicloud.polardb.inputs.ParameterGroupState;
import com.pulumi.alicloud.polardb.outputs.ParameterGroupParameter;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a PolarDB Parameter Group resource.
 * 
 * For information about PolarDB Parameter Group and how to use it, see [What is Parameter Group](https://www.alibabacloud.com/help/en/polardb/polardb-for-mysql/user-guide/apply-a-parameter-template).
 * 
 * &gt; **NOTE:** Available in v1.183.0+.
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
 * import com.pulumi.alicloud.polardb.ParameterGroup;
 * import com.pulumi.alicloud.polardb.ParameterGroupArgs;
 * import com.pulumi.alicloud.polardb.inputs.ParameterGroupParameterArgs;
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
 *         var example = new ParameterGroup(&#34;example&#34;, ParameterGroupArgs.builder()        
 *             .dbType(&#34;MySQL&#34;)
 *             .dbVersion(&#34;8.0&#34;)
 *             .description(&#34;example_value&#34;)
 *             .parameters(ParameterGroupParameterArgs.builder()
 *                 .paramName(&#34;wait_timeout&#34;)
 *                 .paramValue(&#34;86400&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * PolarDB Parameter Group can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:polardb/parameterGroup:ParameterGroup example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:polardb/parameterGroup:ParameterGroup")
public class ParameterGroup extends com.pulumi.resources.CustomResource {
    /**
     * The type of the database engine. Only `MySQL` is supported.
     * 
     */
    @Export(name="dbType", refs={String.class}, tree="[0]")
    private Output<String> dbType;

    /**
     * @return The type of the database engine. Only `MySQL` is supported.
     * 
     */
    public Output<String> dbType() {
        return this.dbType;
    }
    /**
     * The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
     * 
     */
    @Export(name="dbVersion", refs={String.class}, tree="[0]")
    private Output<String> dbVersion;

    /**
     * @return The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
     * 
     */
    public Output<String> dbVersion() {
        return this.dbVersion;
    }
    /**
     * The description of the parameter template. It must be 0 to 200 characters in length.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the parameter template. It must be 0 to 200 characters in length.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The parameter template. See the following `Block parameters`.
     * 
     */
    @Export(name="parameters", refs={List.class,ParameterGroupParameter.class}, tree="[0,1]")
    private Output<List<ParameterGroupParameter>> parameters;

    /**
     * @return The parameter template. See the following `Block parameters`.
     * 
     */
    public Output<List<ParameterGroupParameter>> parameters() {
        return this.parameters;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ParameterGroup(String name) {
        this(name, ParameterGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ParameterGroup(String name, ParameterGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ParameterGroup(String name, ParameterGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:polardb/parameterGroup:ParameterGroup", name, args == null ? ParameterGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ParameterGroup(String name, Output<String> id, @Nullable ParameterGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:polardb/parameterGroup:ParameterGroup", name, state, makeResourceOptions(options, id));
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
    public static ParameterGroup get(String name, Output<String> id, @Nullable ParameterGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ParameterGroup(name, id, state, options);
    }
}
