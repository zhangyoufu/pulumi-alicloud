// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ram;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ram.AccountPasswordPolicyArgs;
import com.pulumi.alicloud.ram.inputs.AccountPasswordPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * Empty resource sets defaults values for every property.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.ram.AccountPasswordPolicy;
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
 *         var default_ = new AccountPasswordPolicy(&#34;default&#34;);
 * 
 *     }
 * }
 * ```
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.ram.AccountPasswordPolicy;
 * import com.pulumi.alicloud.ram.AccountPasswordPolicyArgs;
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
 *         var corporate = new AccountPasswordPolicy(&#34;corporate&#34;, AccountPasswordPolicyArgs.builder()        
 *             .hardExpiry(true)
 *             .maxLoginAttempts(3)
 *             .maxPasswordAge(12)
 *             .minimumPasswordLength(9)
 *             .passwordReusePrevention(5)
 *             .requireLowercaseCharacters(false)
 *             .requireNumbers(false)
 *             .requireSymbols(false)
 *             .requireUppercaseCharacters(false)
 *             .build());
 * 
 *     }
 * }
 * ```
 * For not specified values sets defaults.
 * 
 * ## Import
 * 
 * RAM account password policy can be imported using the `id`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import alicloud:ram/accountPasswordPolicy:AccountPasswordPolicy example ram-account-password-policy
 * ```
 * 
 */
@ResourceType(type="alicloud:ram/accountPasswordPolicy:AccountPasswordPolicy")
public class AccountPasswordPolicy extends com.pulumi.resources.CustomResource {
    /**
     * Specifies if a password can expire in a hard way. Default to false.
     * 
     */
    @Export(name="hardExpiry", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> hardExpiry;

    /**
     * @return Specifies if a password can expire in a hard way. Default to false.
     * 
     */
    public Output<Optional<Boolean>> hardExpiry() {
        return Codegen.optional(this.hardExpiry);
    }
    /**
     * Maximum logon attempts with an incorrect password within an hour. Valid value range: [0-32]. Default to 5.
     * 
     */
    @Export(name="maxLoginAttempts", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> maxLoginAttempts;

    /**
     * @return Maximum logon attempts with an incorrect password within an hour. Valid value range: [0-32]. Default to 5.
     * 
     */
    public Output<Optional<Integer>> maxLoginAttempts() {
        return Codegen.optional(this.maxLoginAttempts);
    }
    /**
     * The number of days after which password expires. A value of 0 indicates that the password never expires. Valid value range: [0-1095]. Default to 0.
     * 
     */
    @Export(name="maxPasswordAge", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> maxPasswordAge;

    /**
     * @return The number of days after which password expires. A value of 0 indicates that the password never expires. Valid value range: [0-1095]. Default to 0.
     * 
     */
    public Output<Optional<Integer>> maxPasswordAge() {
        return Codegen.optional(this.maxPasswordAge);
    }
    /**
     * Minimal required length of password for a user. Valid value range: [8-32]. Default to 12.
     * 
     */
    @Export(name="minimumPasswordLength", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> minimumPasswordLength;

    /**
     * @return Minimal required length of password for a user. Valid value range: [8-32]. Default to 12.
     * 
     */
    public Output<Optional<Integer>> minimumPasswordLength() {
        return Codegen.optional(this.minimumPasswordLength);
    }
    /**
     * User is not allowed to use the latest number of passwords specified in this parameter. A value of 0 indicates the password history check policy is disabled. Valid value range: [0-24]. Default to 0.
     * 
     */
    @Export(name="passwordReusePrevention", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> passwordReusePrevention;

    /**
     * @return User is not allowed to use the latest number of passwords specified in this parameter. A value of 0 indicates the password history check policy is disabled. Valid value range: [0-24]. Default to 0.
     * 
     */
    public Output<Optional<Integer>> passwordReusePrevention() {
        return Codegen.optional(this.passwordReusePrevention);
    }
    /**
     * Specifies if the occurrence of a lowercase character in the password is mandatory. Default to true.
     * 
     */
    @Export(name="requireLowercaseCharacters", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> requireLowercaseCharacters;

    /**
     * @return Specifies if the occurrence of a lowercase character in the password is mandatory. Default to true.
     * 
     */
    public Output<Optional<Boolean>> requireLowercaseCharacters() {
        return Codegen.optional(this.requireLowercaseCharacters);
    }
    /**
     * Specifies if the occurrence of a number in the password is mandatory. Default to true.
     * 
     */
    @Export(name="requireNumbers", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> requireNumbers;

    /**
     * @return Specifies if the occurrence of a number in the password is mandatory. Default to true.
     * 
     */
    public Output<Optional<Boolean>> requireNumbers() {
        return Codegen.optional(this.requireNumbers);
    }
    /**
     * (Optional Specifies if the occurrence of a special character in the password is mandatory. Default to true.
     * 
     */
    @Export(name="requireSymbols", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> requireSymbols;

    /**
     * @return (Optional Specifies if the occurrence of a special character in the password is mandatory. Default to true.
     * 
     */
    public Output<Optional<Boolean>> requireSymbols() {
        return Codegen.optional(this.requireSymbols);
    }
    /**
     * Specifies if the occurrence of an uppercase character in the password is mandatory. Default to true.
     * 
     */
    @Export(name="requireUppercaseCharacters", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> requireUppercaseCharacters;

    /**
     * @return Specifies if the occurrence of an uppercase character in the password is mandatory. Default to true.
     * 
     */
    public Output<Optional<Boolean>> requireUppercaseCharacters() {
        return Codegen.optional(this.requireUppercaseCharacters);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccountPasswordPolicy(String name) {
        this(name, AccountPasswordPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccountPasswordPolicy(String name, @Nullable AccountPasswordPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccountPasswordPolicy(String name, @Nullable AccountPasswordPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ram/accountPasswordPolicy:AccountPasswordPolicy", name, args == null ? AccountPasswordPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccountPasswordPolicy(String name, Output<String> id, @Nullable AccountPasswordPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ram/accountPasswordPolicy:AccountPasswordPolicy", name, state, makeResourceOptions(options, id));
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
    public static AccountPasswordPolicy get(String name, Output<String> id, @Nullable AccountPasswordPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccountPasswordPolicy(name, id, state, options);
    }
}
