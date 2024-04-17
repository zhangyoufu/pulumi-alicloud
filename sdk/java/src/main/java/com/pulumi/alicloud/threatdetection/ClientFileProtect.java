// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.threatdetection.ClientFileProtectArgs;
import com.pulumi.alicloud.threatdetection.inputs.ClientFileProtectState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Threat Detection Client File Protect resource. Client core file protection event monitoring, including file reading and writing, deletion, and permission change.
 * 
 * For information about Threat Detection Client File Protect and how to use it, see [What is Client File Protect](https://www.alibabacloud.com/help/zh/security-center/developer-reference/api-sas-2018-12-03-createfileprotectrule).
 * 
 * &gt; **NOTE:** Available since v1.212.0.
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
 * import com.pulumi.alicloud.threatdetection.ClientFileProtect;
 * import com.pulumi.alicloud.threatdetection.ClientFileProtectArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         var default_ = new ClientFileProtect(&#34;default&#34;, ClientFileProtectArgs.builder()        
 *             .status(&#34;0&#34;)
 *             .filePaths(&#34;/usr/local&#34;)
 *             .fileOps(&#34;CREATE&#34;)
 *             .ruleAction(&#34;pass&#34;)
 *             .procPaths(&#34;/usr/local&#34;)
 *             .alertLevel(&#34;0&#34;)
 *             .switchId(&#34;FILE_PROTECT_RULE_SWITCH_TYPE_1693474122929&#34;)
 *             .ruleName(&#34;rule_example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Threat Detection Client File Protect can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:threatdetection/clientFileProtect:ClientFileProtect example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:threatdetection/clientFileProtect:ClientFileProtect")
public class ClientFileProtect extends com.pulumi.resources.CustomResource {
    /**
     * 0 no alert 1 info 2 suspicious 3 critical.
     * 
     */
    @Export(name="alertLevel", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> alertLevel;

    /**
     * @return 0 no alert 1 info 2 suspicious 3 critical.
     * 
     */
    public Output<Optional<Integer>> alertLevel() {
        return Codegen.optional(this.alertLevel);
    }
    /**
     * file operation.
     * 
     */
    @Export(name="fileOps", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> fileOps;

    /**
     * @return file operation.
     * 
     */
    public Output<List<String>> fileOps() {
        return this.fileOps;
    }
    /**
     * file path.
     * 
     */
    @Export(name="filePaths", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> filePaths;

    /**
     * @return file path.
     * 
     */
    public Output<List<String>> filePaths() {
        return this.filePaths;
    }
    /**
     * process path.
     * 
     */
    @Export(name="procPaths", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> procPaths;

    /**
     * @return process path.
     * 
     */
    public Output<List<String>> procPaths() {
        return this.procPaths;
    }
    /**
     * rule action, pass or alert.
     * 
     */
    @Export(name="ruleAction", refs={String.class}, tree="[0]")
    private Output<String> ruleAction;

    /**
     * @return rule action, pass or alert.
     * 
     */
    public Output<String> ruleAction() {
        return this.ruleAction;
    }
    /**
     * ruleName.
     * 
     */
    @Export(name="ruleName", refs={String.class}, tree="[0]")
    private Output<String> ruleName;

    /**
     * @return ruleName.
     * 
     */
    public Output<String> ruleName() {
        return this.ruleName;
    }
    /**
     * rule status 0 is disable 1 is enable.
     * 
     */
    @Export(name="status", refs={Integer.class}, tree="[0]")
    private Output<Integer> status;

    /**
     * @return rule status 0 is disable 1 is enable.
     * 
     */
    public Output<Integer> status() {
        return this.status;
    }
    /**
     * switch id.
     * 
     */
    @Export(name="switchId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> switchId;

    /**
     * @return switch id.
     * 
     */
    public Output<Optional<String>> switchId() {
        return Codegen.optional(this.switchId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClientFileProtect(String name) {
        this(name, ClientFileProtectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClientFileProtect(String name, ClientFileProtectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClientFileProtect(String name, ClientFileProtectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/clientFileProtect:ClientFileProtect", name, args == null ? ClientFileProtectArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ClientFileProtect(String name, Output<String> id, @Nullable ClientFileProtectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/clientFileProtect:ClientFileProtect", name, state, makeResourceOptions(options, id));
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
    public static ClientFileProtect get(String name, Output<String> id, @Nullable ClientFileProtectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClientFileProtect(name, id, state, options);
    }
}
