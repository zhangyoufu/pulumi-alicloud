// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.log.MachineGroupArgs;
import com.pulumi.alicloud.log.inputs.MachineGroupState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Log Service manages all the ECS instances whose logs need to be collected by using the Logtail client in the form of machine groups.
 *  [Refer to details](https://www.alibabacloud.com/help/doc-detail/28966.htm)
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
 * import com.pulumi.random.RandomInteger;
 * import com.pulumi.random.RandomIntegerArgs;
 * import com.pulumi.alicloud.log.Project;
 * import com.pulumi.alicloud.log.ProjectArgs;
 * import com.pulumi.alicloud.log.MachineGroup;
 * import com.pulumi.alicloud.log.MachineGroupArgs;
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
 *         var default_ = new RandomInteger(&#34;default&#34;, RandomIntegerArgs.builder()        
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .description(&#34;terraform-example&#34;)
 *             .build());
 * 
 *         var exampleMachineGroup = new MachineGroup(&#34;exampleMachineGroup&#34;, MachineGroupArgs.builder()        
 *             .project(exampleProject.name())
 *             .identifyType(&#34;ip&#34;)
 *             .topic(&#34;terraform&#34;)
 *             .identifyLists(            
 *                 &#34;10.0.0.1&#34;,
 *                 &#34;10.0.0.2&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Module Support
 * 
 * You can use the existing sls-logtail module
 * to create logtail config, machine group, install logtail on ECS instances and join instances into machine group one-click.
 * 
 * ## Import
 * 
 * Log machine group can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:log/machineGroup:MachineGroup example tf-log:tf-machine-group
 * ```
 * 
 */
@ResourceType(type="alicloud:log/machineGroup:MachineGroup")
public class MachineGroup extends com.pulumi.resources.CustomResource {
    /**
     * The specific machine identification, which can be an IP address or user-defined identity.
     * 
     */
    @Export(name="identifyLists", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> identifyLists;

    /**
     * @return The specific machine identification, which can be an IP address or user-defined identity.
     * 
     */
    public Output<List<String>> identifyLists() {
        return this.identifyLists;
    }
    /**
     * The machine identification type, including IP and user-defined identity. Valid values are &#34;ip&#34; and &#34;userdefined&#34;. Default to &#34;ip&#34;.
     * 
     */
    @Export(name="identifyType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identifyType;

    /**
     * @return The machine identification type, including IP and user-defined identity. Valid values are &#34;ip&#34; and &#34;userdefined&#34;. Default to &#34;ip&#34;.
     * 
     */
    public Output<Optional<String>> identifyType() {
        return Codegen.optional(this.identifyType);
    }
    /**
     * The machine group name, which is unique in the same project.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The machine group name, which is unique in the same project.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The project name to the machine group belongs.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The project name to the machine group belongs.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The topic of a machine group.
     * 
     */
    @Export(name="topic", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> topic;

    /**
     * @return The topic of a machine group.
     * 
     */
    public Output<Optional<String>> topic() {
        return Codegen.optional(this.topic);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MachineGroup(String name) {
        this(name, MachineGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MachineGroup(String name, MachineGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MachineGroup(String name, MachineGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/machineGroup:MachineGroup", name, args == null ? MachineGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MachineGroup(String name, Output<String> id, @Nullable MachineGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/machineGroup:MachineGroup", name, state, makeResourceOptions(options, id));
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
    public static MachineGroup get(String name, Output<String> id, @Nullable MachineGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MachineGroup(name, id, state, options);
    }
}
