// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cddc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cddc.DedicatedHostAccountArgs;
import com.pulumi.alicloud.cddc.inputs.DedicatedHostAccountState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ApsaraDB for MyBase Dedicated Host Account resource.
 * 
 * For information about ApsaraDB for MyBase Dedicated Host Account and how to use it, see [What is Dedicated Host Account](https://www.alibabacloud.com/help/en/doc-detail/196877.html).
 * 
 * &gt; **NOTE:** Available in v1.148.0+.
 * 
 * &gt; **NOTE:** Each Dedicated host can have only one account. Before you create an account for a host, make sure that the existing account is deleted.
 * 
 * ## Import
 * 
 * ApsaraDB for MyBase Dedicated Host Account can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:cddc/dedicatedHostAccount:DedicatedHostAccount example &lt;dedicated_host_id&gt;:&lt;account_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cddc/dedicatedHostAccount:DedicatedHostAccount")
public class DedicatedHostAccount extends com.pulumi.resources.CustomResource {
    /**
     * The name of the Dedicated host account. The account name must be 2 to 16 characters in length, contain lower case letters, digits, and underscore(_). At the same time, the name must start with a letter and end with a letter or number.
     * 
     */
    @Export(name="accountName", type=String.class, parameters={})
    private Output<String> accountName;

    /**
     * @return The name of the Dedicated host account. The account name must be 2 to 16 characters in length, contain lower case letters, digits, and underscore(_). At the same time, the name must start with a letter and end with a letter or number.
     * 
     */
    public Output<String> accountName() {
        return this.accountName;
    }
    /**
     * The password of the Dedicated host account. The account password must be 6 to 32 characters in length, and can contain letters, digits, and special characters `!@#$%^&amp;*()_+-=`.
     * 
     */
    @Export(name="accountPassword", type=String.class, parameters={})
    private Output<String> accountPassword;

    /**
     * @return The password of the Dedicated host account. The account password must be 6 to 32 characters in length, and can contain letters, digits, and special characters `!@#$%^&amp;*()_+-=`.
     * 
     */
    public Output<String> accountPassword() {
        return this.accountPassword;
    }
    /**
     * The type of the Dedicated host account. Valid values: `Admin`, `Normal`.
     * 
     */
    @Export(name="accountType", type=String.class, parameters={})
    private Output</* @Nullable */ String> accountType;

    /**
     * @return The type of the Dedicated host account. Valid values: `Admin`, `Normal`.
     * 
     */
    public Output<Optional<String>> accountType() {
        return Codegen.optional(this.accountType);
    }
    /**
     * The ID of Dedicated the host.
     * 
     */
    @Export(name="dedicatedHostId", type=String.class, parameters={})
    private Output<String> dedicatedHostId;

    /**
     * @return The ID of Dedicated the host.
     * 
     */
    public Output<String> dedicatedHostId() {
        return this.dedicatedHostId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DedicatedHostAccount(String name) {
        this(name, DedicatedHostAccountArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DedicatedHostAccount(String name, DedicatedHostAccountArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DedicatedHostAccount(String name, DedicatedHostAccountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cddc/dedicatedHostAccount:DedicatedHostAccount", name, args == null ? DedicatedHostAccountArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DedicatedHostAccount(String name, Output<String> id, @Nullable DedicatedHostAccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cddc/dedicatedHostAccount:DedicatedHostAccount", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accountPassword"
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
    public static DedicatedHostAccount get(String name, Output<String> id, @Nullable DedicatedHostAccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DedicatedHostAccount(name, id, state, options);
    }
}
