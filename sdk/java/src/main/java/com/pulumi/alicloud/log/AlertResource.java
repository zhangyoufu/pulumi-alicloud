// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.log.AlertResourceArgs;
import com.pulumi.alicloud.log.inputs.AlertResourceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Using this resource can init SLS Alert resources automatically.
 * 
 * For information about SLS Alert and how to use it, see [SLS Alert Overview](https://www.alibabacloud.com/help/en/doc-detail/209202.html)
 * 
 * &gt; **NOTE:** Available since v1.219.0.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.log.AlertResource;
 * import com.pulumi.alicloud.log.AlertResourceArgs;
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
 *         var exampleUser = new AlertResource(&#34;exampleUser&#34;, AlertResourceArgs.builder()        
 *             .type(&#34;user&#34;)
 *             .lang(&#34;cn&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Log alert resource can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:log/alertResource:AlertResource example alert_resource:project:tf-project
 * ```
 * 
 */
@ResourceType(type="alicloud:log/alertResource:AlertResource")
public class AlertResource extends com.pulumi.resources.CustomResource {
    /**
     * The lang of alert center resource when type is user.
     * 
     */
    @Export(name="lang", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lang;

    /**
     * @return The lang of alert center resource when type is user.
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }
    /**
     * The project of alert resource when type is project.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> project;

    /**
     * @return The project of alert resource when type is project.
     * 
     */
    public Output<Optional<String>> project() {
        return Codegen.optional(this.project);
    }
    /**
     * The type of alert resources, must be user or project, &#39;user&#39; for init aliyuncloud account&#39;s alert center resource, including project named sls-alert-{uid}-{region} and some dashboards; &#39;project&#39; for init project&#39;s alert resource, including logstore named internal-alert-history and alert dashboard.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of alert resources, must be user or project, &#39;user&#39; for init aliyuncloud account&#39;s alert center resource, including project named sls-alert-{uid}-{region} and some dashboards; &#39;project&#39; for init project&#39;s alert resource, including logstore named internal-alert-history and alert dashboard.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AlertResource(String name) {
        this(name, AlertResourceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AlertResource(String name, AlertResourceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AlertResource(String name, AlertResourceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/alertResource:AlertResource", name, args == null ? AlertResourceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AlertResource(String name, Output<String> id, @Nullable AlertResourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/alertResource:AlertResource", name, state, makeResourceOptions(options, id));
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
    public static AlertResource get(String name, Output<String> id, @Nullable AlertResourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AlertResource(name, id, state, options);
    }
}
