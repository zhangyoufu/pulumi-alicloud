// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.maxcompute;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.maxcompute.ProjectArgs;
import com.pulumi.alicloud.maxcompute.inputs.ProjectState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The project is the basic unit of operation in maxcompute. It is similar to the concept of Database or Schema in traditional databases, and sets the boundary for maxcompute multi-user isolation and access control. [Refer to details](https://www.alibabacloud.com/help/doc-detail/27818.html).
 * 
 * -&gt;**NOTE:** Available in 1.77.0+.
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
 * import com.pulumi.alicloud.maxcompute.Project;
 * import com.pulumi.alicloud.maxcompute.ProjectArgs;
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
 *         var example = new Project(&#34;example&#34;, ProjectArgs.builder()        
 *             .orderType(&#34;PayAsYouGo&#34;)
 *             .projectName(&#34;tf_maxcompute_project&#34;)
 *             .specificationType(&#34;OdpsStandard&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * MaxCompute project can be imported using the *name* or ID, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:maxcompute/project:Project example tf_maxcompute_project
 * ```
 * 
 */
@ResourceType(type="alicloud:maxcompute/project:Project")
public class Project extends com.pulumi.resources.CustomResource {
    /**
     * It has been deprecated from provider version 1.110.0 and `project_name` instead.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return It has been deprecated from provider version 1.110.0 and `project_name` instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The type of payment, only `PayAsYouGo` supported currently.
     * 
     */
    @Export(name="orderType", type=String.class, parameters={})
    private Output<String> orderType;

    /**
     * @return The type of payment, only `PayAsYouGo` supported currently.
     * 
     */
    public Output<String> orderType() {
        return this.orderType;
    }
    /**
     * The name of the maxcompute project.
     * 
     */
    @Export(name="projectName", type=String.class, parameters={})
    private Output<String> projectName;

    /**
     * @return The name of the maxcompute project.
     * 
     */
    public Output<String> projectName() {
        return this.projectName;
    }
    /**
     * The type of resource Specification, only `OdpsStandard` supported currently.
     * 
     */
    @Export(name="specificationType", type=String.class, parameters={})
    private Output<String> specificationType;

    /**
     * @return The type of resource Specification, only `OdpsStandard` supported currently.
     * 
     */
    public Output<String> specificationType() {
        return this.specificationType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Project(String name) {
        this(name, ProjectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Project(String name, ProjectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Project(String name, ProjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:maxcompute/project:Project", name, args == null ? ProjectArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Project(String name, Output<String> id, @Nullable ProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:maxcompute/project:Project", name, state, makeResourceOptions(options, id));
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
    public static Project get(String name, Output<String> id, @Nullable ProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Project(name, id, state, options);
    }
}
