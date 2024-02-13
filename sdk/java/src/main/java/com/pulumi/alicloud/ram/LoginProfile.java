// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ram;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ram.LoginProfileArgs;
import com.pulumi.alicloud.ram.inputs.LoginProfileState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a RAM User Login Profile resource.
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
 * import com.pulumi.alicloud.ram.User;
 * import com.pulumi.alicloud.ram.UserArgs;
 * import com.pulumi.alicloud.ram.LoginProfile;
 * import com.pulumi.alicloud.ram.LoginProfileArgs;
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
 *         var user = new User(&#34;user&#34;, UserArgs.builder()        
 *             .displayName(&#34;terraform_example&#34;)
 *             .mobile(&#34;86-18688888888&#34;)
 *             .email(&#34;hello.uuu@aaa.com&#34;)
 *             .comments(&#34;terraform_example&#34;)
 *             .force(true)
 *             .build());
 * 
 *         var profile = new LoginProfile(&#34;profile&#34;, LoginProfileArgs.builder()        
 *             .userName(user.name())
 *             .password(&#34;Example_1234&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RAM login profile can be imported using the id or user name, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ram/loginProfile:LoginProfile example my-login
 * ```
 * 
 */
@ResourceType(type="alicloud:ram/loginProfile:LoginProfile")
public class LoginProfile extends com.pulumi.resources.CustomResource {
    /**
     * This parameter indicates whether the MFA needs to be bind when the user first logs in. Default value is `false`.
     * 
     */
    @Export(name="mfaBindRequired", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> mfaBindRequired;

    /**
     * @return This parameter indicates whether the MFA needs to be bind when the user first logs in. Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> mfaBindRequired() {
        return Codegen.optional(this.mfaBindRequired);
    }
    /**
     * Password of the RAM user.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return Password of the RAM user.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * This parameter indicates whether the password needs to be reset when the user first logs in. Default value is `false`.
     * 
     */
    @Export(name="passwordResetRequired", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> passwordResetRequired;

    /**
     * @return This parameter indicates whether the password needs to be reset when the user first logs in. Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> passwordResetRequired() {
        return Codegen.optional(this.passwordResetRequired);
    }
    /**
     * Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin with a hyphen.
     * 
     */
    @Export(name="userName", refs={String.class}, tree="[0]")
    private Output<String> userName;

    /**
     * @return Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin with a hyphen.
     * 
     */
    public Output<String> userName() {
        return this.userName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LoginProfile(String name) {
        this(name, LoginProfileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LoginProfile(String name, LoginProfileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LoginProfile(String name, LoginProfileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ram/loginProfile:LoginProfile", name, args == null ? LoginProfileArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LoginProfile(String name, Output<String> id, @Nullable LoginProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ram/loginProfile:LoginProfile", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
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
    public static LoginProfile get(String name, Output<String> id, @Nullable LoginProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LoginProfile(name, id, state, options);
    }
}
