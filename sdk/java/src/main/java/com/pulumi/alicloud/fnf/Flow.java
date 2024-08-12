// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.fnf;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.fnf.FlowArgs;
import com.pulumi.alicloud.fnf.inputs.FlowState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Serverless Workflow Flow resource.
 * 
 * For information about Serverless Workflow Flow and how to use it, see [What is Flow](https://www.alibabacloud.com/help/en/doc-detail/123079.htm).
 * 
 * &gt; **NOTE:** Available since v1.105.0+.
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
 * import com.pulumi.alicloud.ram.Role;
 * import com.pulumi.alicloud.ram.RoleArgs;
 * import com.pulumi.alicloud.fnf.Flow;
 * import com.pulumi.alicloud.fnf.FlowArgs;
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
 *         var default_ = new Role("default", RoleArgs.builder()
 *             .name("tf-example-fnfflow")
 *             .document("""
 *   {
 *     "Statement": [
 *       {
 *         "Action": "sts:AssumeRole",
 *         "Effect": "Allow",
 *         "Principal": {
 *           "Service": [
 *             "fnf.aliyuncs.com"
 *           ]
 *         }
 *       }
 *     ],
 *     "Version": "1"
 *   }
 *             """)
 *             .build());
 * 
 *         var example = new Flow("example", FlowArgs.builder()
 *             .definition("""
 *   version: v1beta1
 *   type: flow
 *   steps:
 *     - type: pass
 *       name: helloworld
 *             """)
 *             .roleArn(default_.arn())
 *             .description("Test for terraform fnf_flow.")
 *             .name("tf-example-flow")
 *             .type("FDL")
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
 * Serverless Workflow Flow can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:fnf/flow:Flow example &lt;name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:fnf/flow:Flow")
public class Flow extends com.pulumi.resources.CustomResource {
    /**
     * The definition of the flow. It must comply with the Flow Definition Language (FDL) syntax.
     * 
     */
    @Export(name="definition", refs={String.class}, tree="[0]")
    private Output<String> definition;

    /**
     * @return The definition of the flow. It must comply with the Flow Definition Language (FDL) syntax.
     * 
     */
    public Output<String> definition() {
        return this.definition;
    }
    /**
     * The description of the flow.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the flow.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The unique ID of the flow.
     * 
     */
    @Export(name="flowId", refs={String.class}, tree="[0]")
    private Output<String> flowId;

    /**
     * @return The unique ID of the flow.
     * 
     */
    public Output<String> flowId() {
        return this.flowId;
    }
    /**
     * The time when the flow was last modified.
     * 
     */
    @Export(name="lastModifiedTime", refs={String.class}, tree="[0]")
    private Output<String> lastModifiedTime;

    /**
     * @return The time when the flow was last modified.
     * 
     */
    public Output<String> lastModifiedTime() {
        return this.lastModifiedTime;
    }
    /**
     * The name of the flow. The name must be unique in an Alibaba Cloud account.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the flow. The name must be unique in an Alibaba Cloud account.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ARN of the specified RAM role that Serverless Workflow uses to assume the role when Serverless Workflow executes a flow.
     * 
     */
    @Export(name="roleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> roleArn;

    /**
     * @return The ARN of the specified RAM role that Serverless Workflow uses to assume the role when Serverless Workflow executes a flow.
     * 
     */
    public Output<Optional<String>> roleArn() {
        return Codegen.optional(this.roleArn);
    }
    /**
     * The type of the flow. Valid values are `FDL` or `DEFAULT`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of the flow. Valid values are `FDL` or `DEFAULT`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Flow(java.lang.String name) {
        this(name, FlowArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Flow(java.lang.String name, FlowArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Flow(java.lang.String name, FlowArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:fnf/flow:Flow", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Flow(java.lang.String name, Output<java.lang.String> id, @Nullable FlowState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:fnf/flow:Flow", name, state, makeResourceOptions(options, id), false);
    }

    private static FlowArgs makeArgs(FlowArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? FlowArgs.Empty : args;
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
    public static Flow get(java.lang.String name, Output<java.lang.String> id, @Nullable FlowState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Flow(name, id, state, options);
    }
}
