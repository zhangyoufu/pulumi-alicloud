// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.directmail;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.directmail.MailAddressArgs;
import com.pulumi.alicloud.directmail.inputs.MailAddressState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Direct Mail Mail Address resource.
 * 
 * For information about Direct Mail Mail Address and how to use it, see [What is Mail Address](https://www.aliyun.com/product/directmail).
 * 
 * &gt; **NOTE:** Available in v1.134.0+.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.directmail.MailAddress;
 * import com.pulumi.alicloud.directmail.MailAddressArgs;
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
 *         var example = new MailAddress(&#34;example&#34;, MailAddressArgs.builder()        
 *             .accountName(&#34;example_value@email.com&#34;)
 *             .sendtype(&#34;batch&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * &gt; **Note:**
 * A maximum of 10 mailing addresses can be added.
 * Individual users: Up to 10 mailing addresses can be deleted within a month.
 * Enterprise users: Up to 10 mailing addresses can be deleted within a month.
 * 
 * ## Import
 * 
 * Direct Mail Mail Address can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:directmail/mailAddress:MailAddress example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:directmail/mailAddress:MailAddress")
public class MailAddress extends com.pulumi.resources.CustomResource {
    /**
     * The sender address. The email address must be filled in the format of account@domain, and only lowercase letters or numbers can be used.
     * 
     */
    @Export(name="accountName", type=String.class, parameters={})
    private Output<String> accountName;

    /**
     * @return The sender address. The email address must be filled in the format of account@domain, and only lowercase letters or numbers can be used.
     * 
     */
    public Output<String> accountName() {
        return this.accountName;
    }
    /**
     * Account password. The password must be length 10-20 string, contains numbers, uppercase letters, lowercase letters at the same time.
     * 
     */
    @Export(name="password", type=String.class, parameters={})
    private Output</* @Nullable */ String> password;

    /**
     * @return Account password. The password must be length 10-20 string, contains numbers, uppercase letters, lowercase letters at the same time.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * Return address.
     * 
     */
    @Export(name="replyAddress", type=String.class, parameters={})
    private Output</* @Nullable */ String> replyAddress;

    /**
     * @return Return address.
     * 
     */
    public Output<Optional<String>> replyAddress() {
        return Codegen.optional(this.replyAddress);
    }
    /**
     * Account type. Valid values: `batch`, `trigger`.
     * 
     */
    @Export(name="sendtype", type=String.class, parameters={})
    private Output<String> sendtype;

    /**
     * @return Account type. Valid values: `batch`, `trigger`.
     * 
     */
    public Output<String> sendtype() {
        return this.sendtype;
    }
    /**
     * Account Status freeze: 1, normal: 0.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return Account Status freeze: 1, normal: 0.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MailAddress(String name) {
        this(name, MailAddressArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MailAddress(String name, MailAddressArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MailAddress(String name, MailAddressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:directmail/mailAddress:MailAddress", name, args == null ? MailAddressArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MailAddress(String name, Output<String> id, @Nullable MailAddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:directmail/mailAddress:MailAddress", name, state, makeResourceOptions(options, id));
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
    public static MailAddress get(String name, Output<String> id, @Nullable MailAddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MailAddress(name, id, state, options);
    }
}
