// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eds.CustomPropertyArgs;
import com.pulumi.alicloud.eds.inputs.CustomPropertyState;
import com.pulumi.alicloud.eds.outputs.CustomPropertyPropertyValue;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ECD Custom Property resource.
 * 
 * For information about ECD Custom Property and how to use it, see [What is Custom Property](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-eds-user-2021-03-08-createproperty-desktop).
 * 
 * &gt; **NOTE:** Available since v1.176.0.
 * 
 * &gt; **NOTE:** Up to 10 different attributes can be created under an alibaba cloud account. Up to 50 different attribute values can be added under an attribute.
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
 * import com.pulumi.alicloud.eds.CustomProperty;
 * import com.pulumi.alicloud.eds.CustomPropertyArgs;
 * import com.pulumi.alicloud.eds.inputs.CustomPropertyPropertyValueArgs;
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
 *         var example = new CustomProperty("example", CustomPropertyArgs.builder()        
 *             .propertyKey("example_key")
 *             .propertyValues(CustomPropertyPropertyValueArgs.builder()
 *                 .propertyValue("example_value")
 *                 .build())
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
 * ECD Custom Property can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:eds/customProperty:CustomProperty example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:eds/customProperty:CustomProperty")
public class CustomProperty extends com.pulumi.resources.CustomResource {
    /**
     * The Custom attribute key.
     * 
     */
    @Export(name="propertyKey", refs={String.class}, tree="[0]")
    private Output<String> propertyKey;

    /**
     * @return The Custom attribute key.
     * 
     */
    public Output<String> propertyKey() {
        return this.propertyKey;
    }
    /**
     * Custom attribute sets the value of. See `property_values` below.
     * 
     */
    @Export(name="propertyValues", refs={List.class,CustomPropertyPropertyValue.class}, tree="[0,1]")
    private Output</* @Nullable */ List<CustomPropertyPropertyValue>> propertyValues;

    /**
     * @return Custom attribute sets the value of. See `property_values` below.
     * 
     */
    public Output<Optional<List<CustomPropertyPropertyValue>>> propertyValues() {
        return Codegen.optional(this.propertyValues);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CustomProperty(String name) {
        this(name, CustomPropertyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CustomProperty(String name, CustomPropertyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CustomProperty(String name, CustomPropertyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eds/customProperty:CustomProperty", name, args == null ? CustomPropertyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CustomProperty(String name, Output<String> id, @Nullable CustomPropertyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eds/customProperty:CustomProperty", name, state, makeResourceOptions(options, id));
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
    public static CustomProperty get(String name, Output<String> id, @Nullable CustomPropertyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CustomProperty(name, id, state, options);
    }
}
