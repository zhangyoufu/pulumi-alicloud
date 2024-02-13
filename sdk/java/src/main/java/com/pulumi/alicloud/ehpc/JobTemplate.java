// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ehpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ehpc.JobTemplateArgs;
import com.pulumi.alicloud.ehpc.inputs.JobTemplateState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Ehpc Job Template resource.
 * 
 * For information about Ehpc Job Template and how to use it, see [What is Job Template](https://www.alibabacloud.com/help/product/57664.html).
 * 
 * &gt; **NOTE:** Available in v1.133.0+.
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
 * import com.pulumi.alicloud.ehpc.JobTemplate;
 * import com.pulumi.alicloud.ehpc.JobTemplateArgs;
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
 *         var default_ = new JobTemplate(&#34;default&#34;, JobTemplateArgs.builder()        
 *             .commandLine(&#34;./LammpsTest/lammps.pbs&#34;)
 *             .jobTemplateName(&#34;example_value&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Ehpc Job Template can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ehpc/jobTemplate:JobTemplate example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ehpc/jobTemplate:JobTemplate")
public class JobTemplate extends com.pulumi.resources.CustomResource {
    /**
     * Queue Jobs, Is of the Form: 1-10:2.
     * 
     */
    @Export(name="arrayRequest", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> arrayRequest;

    /**
     * @return Queue Jobs, Is of the Form: 1-10:2.
     * 
     */
    public Output<Optional<String>> arrayRequest() {
        return Codegen.optional(this.arrayRequest);
    }
    /**
     * Job Maximum Run Time.
     * 
     */
    @Export(name="clockTime", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clockTime;

    /**
     * @return Job Maximum Run Time.
     * 
     */
    public Output<Optional<String>> clockTime() {
        return Codegen.optional(this.clockTime);
    }
    /**
     * Job Commands.
     * 
     */
    @Export(name="commandLine", refs={String.class}, tree="[0]")
    private Output<String> commandLine;

    /**
     * @return Job Commands.
     * 
     */
    public Output<String> commandLine() {
        return this.commandLine;
    }
    /**
     * A Single Compute Node Using the GPU Number.Possible Values: 1~20000.
     * 
     */
    @Export(name="gpu", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> gpu;

    /**
     * @return A Single Compute Node Using the GPU Number.Possible Values: 1~20000.
     * 
     */
    public Output<Optional<Integer>> gpu() {
        return Codegen.optional(this.gpu);
    }
    /**
     * A Job Template Name.
     * 
     */
    @Export(name="jobTemplateName", refs={String.class}, tree="[0]")
    private Output<String> jobTemplateName;

    /**
     * @return A Job Template Name.
     * 
     */
    public Output<String> jobTemplateName() {
        return this.jobTemplateName;
    }
    /**
     * A Single Compute Node Maximum Memory.
     * 
     */
    @Export(name="mem", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mem;

    /**
     * @return A Single Compute Node Maximum Memory.
     * 
     */
    public Output<Optional<String>> mem() {
        return Codegen.optional(this.mem);
    }
    /**
     * Submit a Task Is Required for Computing the Number of Data Nodes to Be. Possible Values: 1~5000 .
     * 
     */
    @Export(name="node", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> node;

    /**
     * @return Submit a Task Is Required for Computing the Number of Data Nodes to Be. Possible Values: 1~5000 .
     * 
     */
    public Output<Optional<Integer>> node() {
        return Codegen.optional(this.node);
    }
    /**
     * Job Commands the Directory.
     * 
     */
    @Export(name="packagePath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> packagePath;

    /**
     * @return Job Commands the Directory.
     * 
     */
    public Output<Optional<String>> packagePath() {
        return Codegen.optional(this.packagePath);
    }
    /**
     * The Job Priority.
     * 
     */
    @Export(name="priority", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> priority;

    /**
     * @return The Job Priority.
     * 
     */
    public Output<Optional<Integer>> priority() {
        return Codegen.optional(this.priority);
    }
    /**
     * The Job Queue.
     * 
     */
    @Export(name="queue", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> queue;

    /**
     * @return The Job Queue.
     * 
     */
    public Output<Optional<String>> queue() {
        return Codegen.optional(this.queue);
    }
    /**
     * If the Job Is Support for the Re-Run.
     * 
     */
    @Export(name="reRunable", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> reRunable;

    /**
     * @return If the Job Is Support for the Re-Run.
     * 
     */
    public Output<Boolean> reRunable() {
        return this.reRunable;
    }
    /**
     * The name of the user who performed the job.
     * 
     */
    @Export(name="runasUser", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> runasUser;

    /**
     * @return The name of the user who performed the job.
     * 
     */
    public Output<Optional<String>> runasUser() {
        return Codegen.optional(this.runasUser);
    }
    /**
     * Error Output Path.
     * 
     */
    @Export(name="stderrRedirectPath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stderrRedirectPath;

    /**
     * @return Error Output Path.
     * 
     */
    public Output<Optional<String>> stderrRedirectPath() {
        return Codegen.optional(this.stderrRedirectPath);
    }
    /**
     * Standard Output Path and.
     * 
     */
    @Export(name="stdoutRedirectPath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stdoutRedirectPath;

    /**
     * @return Standard Output Path and.
     * 
     */
    public Output<Optional<String>> stdoutRedirectPath() {
        return Codegen.optional(this.stdoutRedirectPath);
    }
    /**
     * A Single Compute Node Required Number of Tasks. Possible Values: 1~20000 .
     * 
     */
    @Export(name="task", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> task;

    /**
     * @return A Single Compute Node Required Number of Tasks. Possible Values: 1~20000 .
     * 
     */
    public Output<Optional<Integer>> task() {
        return Codegen.optional(this.task);
    }
    /**
     * A Single Task and the Number of Required Threads.
     * 
     */
    @Export(name="thread", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> thread;

    /**
     * @return A Single Task and the Number of Required Threads.
     * 
     */
    public Output<Optional<Integer>> thread() {
        return Codegen.optional(this.thread);
    }
    /**
     * The Job of the Environment Variable.
     * 
     */
    @Export(name="variables", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> variables;

    /**
     * @return The Job of the Environment Variable.
     * 
     */
    public Output<Optional<String>> variables() {
        return Codegen.optional(this.variables);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public JobTemplate(String name) {
        this(name, JobTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public JobTemplate(String name, JobTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public JobTemplate(String name, JobTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ehpc/jobTemplate:JobTemplate", name, args == null ? JobTemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private JobTemplate(String name, Output<String> id, @Nullable JobTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ehpc/jobTemplate:JobTemplate", name, state, makeResourceOptions(options, id));
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
    public static JobTemplate get(String name, Output<String> id, @Nullable JobTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new JobTemplate(name, id, state, options);
    }
}
