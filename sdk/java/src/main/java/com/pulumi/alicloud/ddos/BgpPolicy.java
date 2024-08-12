// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ddos;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ddos.BgpPolicyArgs;
import com.pulumi.alicloud.ddos.inputs.BgpPolicyState;
import com.pulumi.alicloud.ddos.outputs.BgpPolicyContent;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Ddos Bgp Policy resource.
 * 
 * Ddos protection policy.
 * 
 * For information about Ddos Bgp Policy and how to use it, see [What is Policy](https://www.alibabacloud.com/help/en/).
 * 
 * &gt; **NOTE:** Available since v1.226.0.
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
 * import com.pulumi.alicloud.ddos.BgpPolicy;
 * import com.pulumi.alicloud.ddos.BgpPolicyArgs;
 * import com.pulumi.alicloud.ddos.inputs.BgpPolicyContentArgs;
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
 *         final var name = config.get("name").orElse("tf_exampleacc_bgp32594");
 *         final var policyName = config.get("policyName").orElse("example_l4_policy");
 *         var default_ = new BgpPolicy("default", BgpPolicyArgs.builder()
 *             .content(BgpPolicyContentArgs.builder()
 *                 .enableDefense("false")
 *                 .layer4RuleLists(BgpPolicyContentLayer4RuleListArgs.builder()
 *                     .method("hex")
 *                     .match("1")
 *                     .action("1")
 *                     .limited("0")
 *                     .conditionLists(BgpPolicyContentLayer4RuleListConditionListArgs.builder()
 *                         .arg("3C")
 *                         .position("1")
 *                         .depth("2")
 *                         .build())
 *                     .name("11")
 *                     .priority("10")
 *                     .build())
 *                 .build())
 *             .type("l4")
 *             .policyName("tf_exampleacc_bgp32594")
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
 * Ddos Bgp Policy can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ddos/bgpPolicy:BgpPolicy example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ddos/bgpPolicy:BgpPolicy")
public class BgpPolicy extends com.pulumi.resources.CustomResource {
    /**
     * Configuration Content See `content` below.
     * 
     */
    @Export(name="content", refs={BgpPolicyContent.class}, tree="[0]")
    private Output<BgpPolicyContent> content;

    /**
     * @return Configuration Content See `content` below.
     * 
     */
    public Output<BgpPolicyContent> content() {
        return this.content;
    }
    /**
     * The name of the resource
     * 
     */
    @Export(name="policyName", refs={String.class}, tree="[0]")
    private Output<String> policyName;

    /**
     * @return The name of the resource
     * 
     */
    public Output<String> policyName() {
        return this.policyName;
    }
    /**
     * Type
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Type
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BgpPolicy(java.lang.String name) {
        this(name, BgpPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BgpPolicy(java.lang.String name, BgpPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BgpPolicy(java.lang.String name, BgpPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ddos/bgpPolicy:BgpPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private BgpPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable BgpPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ddos/bgpPolicy:BgpPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static BgpPolicyArgs makeArgs(BgpPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? BgpPolicyArgs.Empty : args;
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
    public static BgpPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable BgpPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BgpPolicy(name, id, state, options);
    }
}
