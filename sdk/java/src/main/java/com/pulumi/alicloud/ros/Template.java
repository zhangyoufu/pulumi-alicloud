// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ros;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ros.TemplateArgs;
import com.pulumi.alicloud.ros.inputs.TemplateState;
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
 * Provides a ROS Template resource.
 * 
 * For information about ROS Template and how to use it, see [What is Template](https://www.alibabacloud.com/help/en/doc-detail/141851.htm).
 * 
 * &gt; **NOTE:** Available in v1.108.0+.
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
 * import com.pulumi.alicloud.ros.Template;
 * import com.pulumi.alicloud.ros.TemplateArgs;
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
 *         var example = new Template("example", TemplateArgs.builder()        
 *             .templateName("example_value")
 *             .templateBody("""
 *     {
 *     	"ROSTemplateFormatVersion": "2015-09-01"
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
 * ROS Template can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ros/template:Template example &lt;template_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ros/template:Template")
public class Template extends com.pulumi.resources.CustomResource {
    /**
     * The description of the template. The description can be up to 256 characters in length.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the template. The description can be up to 256 characters in length.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
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
     * The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You must specify one of the TemplateBody and TemplateURL parameters, but you cannot specify both of them.
     * 
     */
    @Export(name="templateBody", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> templateBody;

    /**
     * @return The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You must specify one of the TemplateBody and TemplateURL parameters, but you cannot specify both of them.
     * 
     */
    public Output<Optional<String>> templateBody() {
        return Codegen.optional(this.templateBody);
    }
    /**
     * The name of the template. The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
     * 
     */
    @Export(name="templateName", refs={String.class}, tree="[0]")
    private Output<String> templateName;

    /**
     * @return The name of the template. The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
     * 
     */
    public Output<String> templateName() {
        return this.templateName;
    }
    /**
     * The template url.
     * 
     */
    @Export(name="templateUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> templateUrl;

    /**
     * @return The template url.
     * 
     */
    public Output<Optional<String>> templateUrl() {
        return Codegen.optional(this.templateUrl);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Template(String name) {
        this(name, TemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Template(String name, TemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Template(String name, TemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ros/template:Template", name, args == null ? TemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Template(String name, Output<String> id, @Nullable TemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ros/template:Template", name, state, makeResourceOptions(options, id));
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
    public static Template get(String name, Output<String> id, @Nullable TemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Template(name, id, state, options);
    }
}
