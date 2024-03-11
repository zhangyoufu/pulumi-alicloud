// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ens;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ens.SecurityGroupArgs;
import com.pulumi.alicloud.ens.inputs.SecurityGroupState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ENS Security Group resource.
 * 
 * For information about ENS Security Group and how to use it, see [What is Security Group](https://www.alibabacloud.com/help/en/ens/developer-reference/api-createsnapshot).
 * 
 * &gt; **NOTE:** Available since v1.213.0.
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
 * import com.pulumi.alicloud.ens.SecurityGroup;
 * import com.pulumi.alicloud.ens.SecurityGroupArgs;
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
 *         var default_ = new SecurityGroup(&#34;default&#34;, SecurityGroupArgs.builder()        
 *             .description(name)
 *             .securityGroupName(name)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ENS Security Group can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ens/securityGroup:SecurityGroup example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ens/securityGroup:SecurityGroup")
public class SecurityGroup extends com.pulumi.resources.CustomResource {
    /**
     * Security group description informationIt must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Security group description informationIt must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Security group nameThe security group name. The length is 2~128 English or Chinese characters. It must start with an uppercase or lowcase letter or a Chinese character and cannot start with `http://` or `https`. Can contain digits, colons (:), underscores (_), or hyphens (-).
     * 
     */
    @Export(name="securityGroupName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> securityGroupName;

    /**
     * @return Security group nameThe security group name. The length is 2~128 English or Chinese characters. It must start with an uppercase or lowcase letter or a Chinese character and cannot start with `http://` or `https`. Can contain digits, colons (:), underscores (_), or hyphens (-).
     * 
     */
    public Output<Optional<String>> securityGroupName() {
        return Codegen.optional(this.securityGroupName);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecurityGroup(String name) {
        this(name, SecurityGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecurityGroup(String name, @Nullable SecurityGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecurityGroup(String name, @Nullable SecurityGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ens/securityGroup:SecurityGroup", name, args == null ? SecurityGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecurityGroup(String name, Output<String> id, @Nullable SecurityGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ens/securityGroup:SecurityGroup", name, state, makeResourceOptions(options, id));
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
    public static SecurityGroup get(String name, Output<String> id, @Nullable SecurityGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecurityGroup(name, id, state, options);
    }
}
