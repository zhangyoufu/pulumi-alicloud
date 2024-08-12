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
 * For information about Direct Mail Mail Address and how to use it, see [What is Mail Address](https://www.alibabacloud.com/help/en/directmail/latest/set-up-sender-addresses).
 * 
 * &gt; **NOTE:** Available since v1.134.0.
 * 
 * ## Import
 * 
 * Direct Mail Mail Address can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:directmail/mailAddress:MailAddress example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:directmail/mailAddress:MailAddress")
public class MailAddress extends com.pulumi.resources.CustomResource {
    /**
     * The sender address. The email address must be filled in the format of account{@literal @}domain, and only lowercase letters or numbers can be used.
     * 
     */
    @Export(name="accountName", refs={String.class}, tree="[0]")
    private Output<String> accountName;

    /**
     * @return The sender address. The email address must be filled in the format of account{@literal @}domain, and only lowercase letters or numbers can be used.
     * 
     */
    public Output<String> accountName() {
        return this.accountName;
    }
    /**
     * Account password. The password must be length 10-20 string, contains numbers, uppercase letters, lowercase letters at the same time.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
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
    @Export(name="replyAddress", refs={String.class}, tree="[0]")
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
    @Export(name="sendtype", refs={String.class}, tree="[0]")
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
    @Export(name="status", refs={String.class}, tree="[0]")
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
    public MailAddress(java.lang.String name) {
        this(name, MailAddressArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MailAddress(java.lang.String name, MailAddressArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MailAddress(java.lang.String name, MailAddressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:directmail/mailAddress:MailAddress", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private MailAddress(java.lang.String name, Output<java.lang.String> id, @Nullable MailAddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:directmail/mailAddress:MailAddress", name, state, makeResourceOptions(options, id), false);
    }

    private static MailAddressArgs makeArgs(MailAddressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? MailAddressArgs.Empty : args;
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
    public static MailAddress get(java.lang.String name, Output<java.lang.String> id, @Nullable MailAddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MailAddress(name, id, state, options);
    }
}
