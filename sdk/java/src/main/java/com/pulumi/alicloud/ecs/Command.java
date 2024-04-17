// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.CommandArgs;
import com.pulumi.alicloud.ecs.inputs.CommandState;
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
 * Provides a ECS Command resource.
 * 
 * For information about ECS Command and how to use it, see [What is Command](https://www.alibabacloud.com/help/en/doc-detail/64844.htm).
 * 
 * &gt; **NOTE:** Available in v1.116.0+.
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
 * import com.pulumi.alicloud.ecs.Command;
 * import com.pulumi.alicloud.ecs.CommandArgs;
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
 *         var example = new Command(&#34;example&#34;, CommandArgs.builder()        
 *             .name(&#34;tf-testAcc&#34;)
 *             .commandContent(&#34;bHMK&#34;)
 *             .description(&#34;For Terraform Test&#34;)
 *             .type(&#34;RunShellScript&#34;)
 *             .workingDir(&#34;/root&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ECS Command can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ecs/command:Command example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/command:Command")
public class Command extends com.pulumi.resources.CustomResource {
    /**
     * The Base64-encoded content of the command.
     * 
     */
    @Export(name="commandContent", refs={String.class}, tree="[0]")
    private Output<String> commandContent;

    /**
     * @return The Base64-encoded content of the command.
     * 
     */
    public Output<String> commandContent() {
        return this.commandContent;
    }
    /**
     * The description of command.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of command.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specifies whether to use custom parameters in the command to be created. Default to: false.
     * 
     */
    @Export(name="enableParameter", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableParameter;

    /**
     * @return Specifies whether to use custom parameters in the command to be created. Default to: false.
     * 
     */
    public Output<Optional<Boolean>> enableParameter() {
        return Codegen.optional(this.enableParameter);
    }
    /**
     * The name of the command, which supports all character sets. It can be up to 128 characters in length.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the command, which supports all character sets. It can be up to 128 characters in length.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The timeout period that is specified for the command to be run on ECS instances. Unit: seconds. Default to: `60`.
     * 
     */
    @Export(name="timeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeout;

    /**
     * @return The timeout period that is specified for the command to be run on ECS instances. Unit: seconds. Default to: `60`.
     * 
     */
    public Output<Optional<Integer>> timeout() {
        return Codegen.optional(this.timeout);
    }
    /**
     * The command type. Valid Values: `RunBatScript`, `RunPowerShellScript` and `RunShellScript`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The command type. Valid Values: `RunBatScript`, `RunPowerShellScript` and `RunShellScript`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * The execution path of the command in the ECS instance.
     * 
     */
    @Export(name="workingDir", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> workingDir;

    /**
     * @return The execution path of the command in the ECS instance.
     * 
     */
    public Output<Optional<String>> workingDir() {
        return Codegen.optional(this.workingDir);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Command(String name) {
        this(name, CommandArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Command(String name, CommandArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Command(String name, CommandArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/command:Command", name, args == null ? CommandArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Command(String name, Output<String> id, @Nullable CommandState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/command:Command", name, state, makeResourceOptions(options, id));
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
    public static Command get(String name, Output<String> id, @Nullable CommandState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Command(name, id, state, options);
    }
}
