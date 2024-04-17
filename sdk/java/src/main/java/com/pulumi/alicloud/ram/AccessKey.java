// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ram;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ram.AccessKeyArgs;
import com.pulumi.alicloud.ram.inputs.AccessKeyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a RAM User access key resource.
 * 
 * &gt; **NOTE:**  You should set the `secret_file` if you want to get the access key.
 * 
 * &gt; **NOTE:**  From version 1.98.0, if not set `pgp_key`, the resource will output the access key secret to field `secret` and please protect your backend state file judiciously
 * 
 * &gt; **NOTE:** Available since v1.0.0+.
 * 
 * ## Example Usage
 * 
 * Output the secret to a file.
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.ram.User;
 * import com.pulumi.alicloud.ram.UserArgs;
 * import com.pulumi.alicloud.ram.AccessKey;
 * import com.pulumi.alicloud.ram.AccessKeyArgs;
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
 *         // Create a new RAM access key for user.
 *         var user = new User(&#34;user&#34;, UserArgs.builder()        
 *             .name(&#34;terraform-example&#34;)
 *             .displayName(&#34;user_display_name&#34;)
 *             .mobile(&#34;86-18688888888&#34;)
 *             .email(&#34;hello.uuu@aaa.com&#34;)
 *             .comments(&#34;yoyoyo&#34;)
 *             .force(true)
 *             .build());
 * 
 *         var ak = new AccessKey(&#34;ak&#34;, AccessKeyArgs.builder()        
 *             .userName(user.name())
 *             .secretFile(&#34;/xxx/xxx/xxx.txt&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * Using `pgp_key` to encrypt the secret.
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.ram.User;
 * import com.pulumi.alicloud.ram.UserArgs;
 * import com.pulumi.alicloud.ram.AccessKey;
 * import com.pulumi.alicloud.ram.AccessKeyArgs;
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
 *         // Create a new RAM access key for user.
 *         var user = new User(&#34;user&#34;, UserArgs.builder()        
 *             .name(&#34;terraform-example&#34;)
 *             .displayName(&#34;user_display_name&#34;)
 *             .mobile(&#34;86-18688888888&#34;)
 *             .email(&#34;hello.uuu@aaa.com&#34;)
 *             .comments(&#34;yoyoyo&#34;)
 *             .force(true)
 *             .build());
 * 
 *         var encrypt = new AccessKey(&#34;encrypt&#34;, AccessKeyArgs.builder()        
 *             .userName(user.name())
 *             .pgpKey(&#34;&#34;&#34;
 * mQENBFXbjPUBCADjNjCUQwfxKL+RR2GA6pv/1K+zJZ8UWIF9S0lk7cVIEfJiprzzwiMwBS5cD0da
 * rGin1FHvIWOZxujA7oW0O2TUuatqI3aAYDTfRYurh6iKLC+VS+F7H+/mhfFvKmgr0Y5kDCF1j0T/
 * 063QZ84IRGucR/X43IY7kAtmxGXH0dYOCzOe5UBX1fTn3mXGe2ImCDWBH7gOViynXmb6XNvXkP0f
 * sF5St9jhO7mbZU9EFkv9O3t3EaURfHopsCVDOlCkFCw5ArY+DUORHRzoMX0PnkyQb5OzibkChzpg
 * 8hQssKeVGpuskTdz5Q7PtdW71jXd4fFVzoNH8fYwRpziD2xNvi6HABEBAAG0EFZhdWx0IFRlc3Qg
 * S2V5IDGJATgEEwECACIFAlXbjPUCGy8GCwkIBwMCBhUIAgkKCwQWAgMBAh4BAheAAAoJEOfLr44B
 * HbeTo+sH/i7bapIgPnZsJ81hmxPj4W12uvunksGJiC7d4hIHsG7kmJRTJfjECi+AuTGeDwBy84TD
 * cRaOB6e79fj65Fg6HgSahDUtKJbGxj/lWzmaBuTzlN3CEe8cMwIPqPT2kajJVdOyrvkyuFOdPFOE
 * A7bdCH0MqgIdM2SdF8t40k/ATfuD2K1ZmumJ508I3gF39jgTnPzD4C8quswrMQ3bzfvKC3klXRlB
 * C0yoArn+0QA3cf2B9T4zJ2qnvgotVbeK/b1OJRNj6Poeo+SsWNc/A5mw7lGScnDgL3yfwCm1gQXa
 * QKfOt5x+7GqhWDw10q+bJpJlI10FfzAnhMF9etSqSeURBRW5AQ0EVduM9QEIAL53hJ5bZJ7oEDCn
 * aY+SCzt9QsAfnFTAnZJQrvkvusJzrTQ088eUQmAjvxkfRqnv981fFwGnh2+I1Ktm698UAZS9Jt8y
 * jak9wWUICKQO5QUt5k8cHwldQXNXVXFa+TpQWQR5yW1a9okjh5o/3d4cBt1yZPUJJyLKY43Wvptb
 * 6EuEsScO2DnRkh5wSMDQ7dTooddJCmaq3LTjOleRFQbu9ij386Do6jzK69mJU56TfdcydkxkWF5N
 * ZLGnED3lq+hQNbe+8UI5tD2oP/3r5tXKgMy1R/XPvR/zbfwvx4FAKFOP01awLq4P3d/2xOkMu4Lu
 * 9p315E87DOleYwxk+FoTqXEAEQEAAYkCPgQYAQIACQUCVduM9QIbLgEpCRDny6+OAR23k8BdIAQZ
 * AQIABgUCVduM9QAKCRAID0JGyHtSGmqYB/4m4rJbbWa7dBJ8VqRU7ZKnNRDR9CVhEGipBmpDGRYu
 * lEimOPzLUX/ZXZmTZzgemeXLBaJJlWnopVUWuAsyjQuZAfdd8nHkGRHG0/DGum0l4sKTta3OPGHN
 * C1z1dAcQ1RCr9bTD3PxjLBczdGqhzw71trkQRBRdtPiUchltPMIyjUHqVJ0xmg0hPqFic0fICsr0
 * YwKoz3h9+QEcZHvsjSZjgydKvfLYcm+4DDMCCqcHuJrbXJKUWmJcXR0y/+HQONGrGJ5xWdO+6eJi
 * oPn2jVMnXCm4EKc7fcLFrz/LKmJ8seXhxjM3EdFtylBGCrx3xdK0f+JDNQaC/rhUb5V2XuX6VwoH
 * /AtY+XsKVYRfNIupLOUcf/srsm3IXT4SXWVomOc9hjGQiJ3rraIbADsc+6bCAr4XNZS7moViAAcI
 * PXFv3m3WfUlnG/om78UjQqyVACRZqqAGmuPq+TSkRUCpt9h+A39LQWkojHqyob3cyLgy6z9Q557O
 * 9uK3lQozbw2gH9zC0RqnePl+rsWIUU/ga16fH6pWc1uJiEBt8UZGypQ/E56/343epmYAe0a87sHx
 * 8iDV+dNtDVKfPRENiLOOc19MmS+phmUyrbHqI91c0pmysYcJZCD3a502X1gpjFbPZcRtiTmGnUKd
 * OIu60YPNE4+h7u2CfYyFPu3AlUaGNMBlvy6PEpU=
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *         ctx.export(&#34;secret&#34;, encrypt.encryptedSecret());
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="alicloud:ram/accessKey:AccessKey")
public class AccessKey extends com.pulumi.resources.CustomResource {
    @Export(name="encryptedSecret", refs={String.class}, tree="[0]")
    private Output<String> encryptedSecret;

    public Output<String> encryptedSecret() {
        return this.encryptedSecret;
    }
    /**
     * The fingerprint of the PGP key used to encrypt the secret
     * 
     */
    @Export(name="keyFingerprint", refs={String.class}, tree="[0]")
    private Output<String> keyFingerprint;

    /**
     * @return The fingerprint of the PGP key used to encrypt the secret
     * 
     */
    public Output<String> keyFingerprint() {
        return this.keyFingerprint;
    }
    /**
     * Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`
     * 
     */
    @Export(name="pgpKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pgpKey;

    /**
     * @return Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`
     * 
     */
    public Output<Optional<String>> pgpKey() {
        return Codegen.optional(this.pgpKey);
    }
    /**
     * (Available since 1.98.0+) - The secret access key. Note that this will be written to the state file.
     * If you use this, please protect your backend state file judiciously.
     * Alternatively, you may supply a `pgp_key` instead, which will prevent the secret from being stored in plaintext,
     * at the cost of preventing the use of the secret key in automation.
     * 
     */
    @Export(name="secret", refs={String.class}, tree="[0]")
    private Output<String> secret;

    /**
     * @return (Available since 1.98.0+) - The secret access key. Note that this will be written to the state file.
     * If you use this, please protect your backend state file judiciously.
     * Alternatively, you may supply a `pgp_key` instead, which will prevent the secret from being stored in plaintext,
     * at the cost of preventing the use of the secret key in automation.
     * 
     */
    public Output<String> secret() {
        return this.secret;
    }
    /**
     * The name of file that can save access key id and access key secret. Strongly suggest you to specified it when you creating access key, otherwise, you wouldn&#39;t get its secret ever.
     * 
     */
    @Export(name="secretFile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> secretFile;

    /**
     * @return The name of file that can save access key id and access key secret. Strongly suggest you to specified it when you creating access key, otherwise, you wouldn&#39;t get its secret ever.
     * 
     */
    public Output<Optional<String>> secretFile() {
        return Codegen.optional(this.secretFile);
    }
    /**
     * Status of access key. It must be `Active` or `Inactive`. Default value is `Active`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> status;

    /**
     * @return Status of access key. It must be `Active` or `Inactive`. Default value is `Active`.
     * 
     */
    public Output<Optional<String>> status() {
        return Codegen.optional(this.status);
    }
    /**
     * Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin with a hyphen.
     * 
     */
    @Export(name="userName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userName;

    /**
     * @return Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin with a hyphen.
     * 
     */
    public Output<Optional<String>> userName() {
        return Codegen.optional(this.userName);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccessKey(String name) {
        this(name, AccessKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccessKey(String name, @Nullable AccessKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccessKey(String name, @Nullable AccessKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ram/accessKey:AccessKey", name, args == null ? AccessKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccessKey(String name, Output<String> id, @Nullable AccessKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ram/accessKey:AccessKey", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secret"
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
    public static AccessKey get(String name, Output<String> id, @Nullable AccessKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccessKey(name, id, state, options);
    }
}
