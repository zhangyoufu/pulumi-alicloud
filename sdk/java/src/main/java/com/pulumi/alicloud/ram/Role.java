// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ram;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ram.RoleArgs;
import com.pulumi.alicloud.ram.inputs.RoleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a RAM Role resource.
 * 
 * &gt; **NOTE:** When you want to destroy this resource forcefully(means remove all the relationships associated with it automatically and then destroy it) without set `force`  with `true` at beginning, you need add `force = true` to configuration file and run `pulumi preview`, then you can delete resource forcefully.
 * 
 * &gt; **NOTE:** Available since v1.0.0+.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.ram.Role;
 * import com.pulumi.alicloud.ram.RoleArgs;
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
 *         var role = new Role(&#34;role&#34;, RoleArgs.builder()        
 *             .description(&#34;this is a role test.&#34;)
 *             .document(&#34;&#34;&#34;
 *   {
 *     &#34;Statement&#34;: [
 *       {
 *         &#34;Action&#34;: &#34;sts:AssumeRole&#34;,
 *         &#34;Effect&#34;: &#34;Allow&#34;,
 *         &#34;Principal&#34;: {
 *           &#34;Service&#34;: [
 *             &#34;apigateway.aliyuncs.com&#34;, 
 *             &#34;ecs.aliyuncs.com&#34;
 *           ]
 *         }
 *       }
 *     ],
 *     &#34;Version&#34;: &#34;1&#34;
 *   }
 *   
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RAM role can be imported using the id or name, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ram/role:Role example my-role
 * ```
 * 
 */
@ResourceType(type="alicloud:ram/role:Role")
public class Role extends com.pulumi.resources.CustomResource {
    /**
     * The role arn.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The role arn.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Authorization strategy of the RAM role. It is required when the `services` and `ram_users` are not specified.
     * 
     */
    @Export(name="document", refs={String.class}, tree="[0]")
    private Output<String> document;

    /**
     * @return Authorization strategy of the RAM role. It is required when the `services` and `ram_users` are not specified.
     * 
     */
    public Output<String> document() {
        return this.document;
    }
    /**
     * This parameter is used for resource destroy. Default value is `false`.
     * 
     */
    @Export(name="force", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> force;

    /**
     * @return This parameter is used for resource destroy. Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> force() {
        return Codegen.optional(this.force);
    }
    /**
     * The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
     * 
     */
    @Export(name="maxSessionDuration", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxSessionDuration;

    /**
     * @return The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
     * 
     */
    public Output<Optional<Integer>> maxSessionDuration() {
        return Codegen.optional(this.maxSessionDuration);
    }
    /**
     * Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;, &#34;_&#34;, and must not begin with a hyphen.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;, &#34;_&#34;, and must not begin with a hyphen.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * (It has been deprecated since version 1.49.0, and use field &#39;document&#39; to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
     * 
     * @deprecated
     * Field &#39;ram_users&#39; has been deprecated from version 1.49.0, and use field &#39;document&#39; to replace.
     * 
     */
    @Deprecated /* Field 'ram_users' has been deprecated from version 1.49.0, and use field 'document' to replace.  */
    @Export(name="ramUsers", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> ramUsers;

    /**
     * @return (It has been deprecated since version 1.49.0, and use field &#39;document&#39; to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
     * 
     */
    public Output<List<String>> ramUsers() {
        return this.ramUsers;
    }
    /**
     * The role ID.
     * 
     */
    @Export(name="roleId", refs={String.class}, tree="[0]")
    private Output<String> roleId;

    /**
     * @return The role ID.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }
    /**
     * (It has been deprecated since version 1.49.0, and use field &#39;document&#39; to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone&#39;s Alicloud account id.
     * 
     * @deprecated
     * Field &#39;services&#39; has been deprecated from version 1.49.0, and use field &#39;document&#39; to replace.
     * 
     */
    @Deprecated /* Field 'services' has been deprecated from version 1.49.0, and use field 'document' to replace.  */
    @Export(name="services", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> services;

    /**
     * @return (It has been deprecated since version 1.49.0, and use field &#39;document&#39; to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone&#39;s Alicloud account id.
     * 
     */
    public Output<List<String>> services() {
        return this.services;
    }
    /**
     * (It has been deprecated since version 1.49.0, and use field &#39;document&#39; to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
     * 
     * @deprecated
     * Field &#39;version&#39; has been deprecated from version 1.49.0, and use field &#39;document&#39; to replace.
     * 
     */
    @Deprecated /* Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.  */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> version;

    /**
     * @return (It has been deprecated since version 1.49.0, and use field &#39;document&#39; to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
     * 
     */
    public Output<Optional<String>> version() {
        return Codegen.optional(this.version);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Role(String name) {
        this(name, RoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Role(String name, @Nullable RoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Role(String name, @Nullable RoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ram/role:Role", name, args == null ? RoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Role(String name, Output<String> id, @Nullable RoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ram/role:Role", name, state, makeResourceOptions(options, id));
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
    public static Role get(String name, Output<String> id, @Nullable RoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Role(name, id, state, options);
    }
}
