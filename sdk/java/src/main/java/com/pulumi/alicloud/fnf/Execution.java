// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.fnf;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.fnf.ExecutionArgs;
import com.pulumi.alicloud.fnf.inputs.ExecutionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Serverless Workflow Execution resource.
 * 
 * For information about Serverless Workflow Execution and how to use it, see [What is Execution](https://www.alibabacloud.com/help/en/doc-detail/122628.html).
 * 
 * &gt; **NOTE:** Available since v1.149.0+.
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
 * import com.pulumi.alicloud.ram.Role;
 * import com.pulumi.alicloud.ram.RoleArgs;
 * import com.pulumi.alicloud.fnf.Flow;
 * import com.pulumi.alicloud.fnf.FlowArgs;
 * import com.pulumi.alicloud.fnf.Execution;
 * import com.pulumi.alicloud.fnf.ExecutionArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-example-fnfflow&#34;);
 *         var defaultRole = new Role(&#34;defaultRole&#34;, RoleArgs.builder()        
 *             .document(&#34;&#34;&#34;
 *   {
 *     &#34;Statement&#34;: [
 *       {
 *         &#34;Action&#34;: &#34;sts:AssumeRole&#34;,
 *         &#34;Effect&#34;: &#34;Allow&#34;,
 *         &#34;Principal&#34;: {
 *           &#34;Service&#34;: [
 *             &#34;fnf.aliyuncs.com&#34;
 *           ]
 *         }
 *       }
 *     ],
 *     &#34;Version&#34;: &#34;1&#34;
 *   }
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *         var defaultFlow = new Flow(&#34;defaultFlow&#34;, FlowArgs.builder()        
 *             .definition(&#34;&#34;&#34;
 *   version: v1beta1
 *   type: flow
 *   steps:
 *     - type: wait
 *       name: custom_wait
 *       duration: $.wait
 *             &#34;&#34;&#34;)
 *             .roleArn(defaultRole.arn())
 *             .description(&#34;Test for terraform fnf_flow.&#34;)
 *             .type(&#34;FDL&#34;)
 *             .build());
 * 
 *         var defaultExecution = new Execution(&#34;defaultExecution&#34;, ExecutionArgs.builder()        
 *             .executionName(name)
 *             .flowName(defaultFlow.name())
 *             .input(&#34;{\&#34;wait\&#34;: 600}&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Serverless Workflow Execution can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:fnf/execution:Execution example &lt;flow_name&gt;:&lt;execution_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:fnf/execution:Execution")
public class Execution extends com.pulumi.resources.CustomResource {
    /**
     * The name of the execution.
     * 
     */
    @Export(name="executionName", refs={String.class}, tree="[0]")
    private Output<String> executionName;

    /**
     * @return The name of the execution.
     * 
     */
    public Output<String> executionName() {
        return this.executionName;
    }
    /**
     * The name of the flow.
     * 
     */
    @Export(name="flowName", refs={String.class}, tree="[0]")
    private Output<String> flowName;

    /**
     * @return The name of the flow.
     * 
     */
    public Output<String> flowName() {
        return this.flowName;
    }
    /**
     * The Input information for this execution.
     * 
     */
    @Export(name="input", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> input;

    /**
     * @return The Input information for this execution.
     * 
     */
    public Output<Optional<String>> input() {
        return Codegen.optional(this.input);
    }
    /**
     * The status of the resource. Valid values: `Stopped`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource. Valid values: `Stopped`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Execution(String name) {
        this(name, ExecutionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Execution(String name, ExecutionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Execution(String name, ExecutionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:fnf/execution:Execution", name, args == null ? ExecutionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Execution(String name, Output<String> id, @Nullable ExecutionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:fnf/execution:Execution", name, state, makeResourceOptions(options, id));
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
    public static Execution get(String name, Output<String> id, @Nullable ExecutionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Execution(name, id, state, options);
    }
}
