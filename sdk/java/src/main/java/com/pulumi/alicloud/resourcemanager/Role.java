// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.resourcemanager.RoleArgs;
import com.pulumi.alicloud.resourcemanager.inputs.RoleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Resource Manager role resource. Members are resource containers in the resource directory, which can physically isolate resources to form an independent resource grouping unit. You can create members in the resource folder to manage them in a unified manner.
 * For information about Resource Manager role and how to use it, see [What is Resource Manager role](https://www.alibabacloud.com/help/en/doc-detail/111231.htm).
 * 
 * &gt; **NOTE:** Available since v1.82.0.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.resourcemanager.Role;
 * import com.pulumi.alicloud.resourcemanager.RoleArgs;
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
 *         final var name = config.get("name").orElse("tfexample");
 *         final var default = AlicloudFunctions.getAccount();
 * 
 *         var example = new Role("example", RoleArgs.builder()
 *             .roleName(name)
 *             .assumeRolePolicyDocument("""
 *      {
 *           "Statement": [
 *                {
 *                     "Action": "sts:AssumeRole",
 *                     "Effect": "Allow",
 *                     "Principal": {
 *                         "RAM":[
 *                                 "acs:ram::%s:root"
 *                         ]
 *                     }
 *                 }
 *           ],
 *           "Version": "1"
 *      }
 * ", default_.id()))
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
 * Resource Manager can be imported using the id or role_name, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:resourcemanager/role:Role example testrd
 * ```
 * 
 */
@ResourceType(type="alicloud:resourcemanager/role:Role")
public class Role extends com.pulumi.resources.CustomResource {
    /**
     * The resource descriptor of the role.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The resource descriptor of the role.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The content of the permissions strategy that plays a role.
     * 
     */
    @Export(name="assumeRolePolicyDocument", refs={String.class}, tree="[0]")
    private Output<String> assumeRolePolicyDocument;

    /**
     * @return The content of the permissions strategy that plays a role.
     * 
     */
    public Output<String> assumeRolePolicyDocument() {
        return this.assumeRolePolicyDocument;
    }
    /**
     * The description of the Resource Manager role.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the Resource Manager role.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
     * 
     */
    @Export(name="maxSessionDuration", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxSessionDuration;

    /**
     * @return Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
     * 
     */
    public Output<Optional<Integer>> maxSessionDuration() {
        return Codegen.optional(this.maxSessionDuration);
    }
    /**
     * This ID of Resource Manager role. The value is set to `role_name`.
     * 
     */
    @Export(name="roleId", refs={String.class}, tree="[0]")
    private Output<String> roleId;

    /**
     * @return This ID of Resource Manager role. The value is set to `role_name`.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }
    /**
     * Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots &#34;.&#34; and dashes &#34;-&#34;.
     * 
     */
    @Export(name="roleName", refs={String.class}, tree="[0]")
    private Output<String> roleName;

    /**
     * @return Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots &#34;.&#34; and dashes &#34;-&#34;.
     * 
     */
    public Output<String> roleName() {
        return this.roleName;
    }
    /**
     * Role update time.
     * 
     */
    @Export(name="updateDate", refs={String.class}, tree="[0]")
    private Output<String> updateDate;

    /**
     * @return Role update time.
     * 
     */
    public Output<String> updateDate() {
        return this.updateDate;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Role(java.lang.String name) {
        this(name, RoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Role(java.lang.String name, RoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Role(java.lang.String name, RoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:resourcemanager/role:Role", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Role(java.lang.String name, Output<java.lang.String> id, @Nullable RoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:resourcemanager/role:Role", name, state, makeResourceOptions(options, id), false);
    }

    private static RoleArgs makeArgs(RoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RoleArgs.Empty : args;
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
    public static Role get(java.lang.String name, Output<java.lang.String> id, @Nullable RoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Role(name, id, state, options);
    }
}
