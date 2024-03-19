// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.adb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.adb.AccountArgs;
import com.pulumi.alicloud.adb.inputs.AccountState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a [ADB](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/latest/api-doc-adb-2019-03-15-api-doc-createaccount) account resource and used to manage databases.
 * 
 * &gt; **NOTE:** Available since v1.71.0.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.adb.AdbFunctions;
 * import com.pulumi.alicloud.adb.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.VpcFunctions;
 * import com.pulumi.alicloud.vpc.inputs.GetNetworksArgs;
 * import com.pulumi.alicloud.vpc.inputs.GetSwitchesArgs;
 * import com.pulumi.alicloud.adb.DBCluster;
 * import com.pulumi.alicloud.adb.DBClusterArgs;
 * import com.pulumi.alicloud.adb.Account;
 * import com.pulumi.alicloud.adb.AccountArgs;
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
 *         final var creation = config.get(&#34;creation&#34;).orElse(&#34;ADB&#34;);
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tfexample&#34;);
 *         final var defaultZones = AdbFunctions.getZones();
 * 
 *         final var defaultNetworks = VpcFunctions.getNetworks(GetNetworksArgs.builder()
 *             .nameRegex(&#34;^default-NODELETING$&#34;)
 *             .build());
 * 
 *         final var defaultSwitches = VpcFunctions.getSwitches(GetSwitchesArgs.builder()
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.ids()[0]))
 *             .build());
 * 
 *         final var vswitchId = defaultSwitches.applyValue(getSwitchesResult -&gt; getSwitchesResult.ids()[0]);
 * 
 *         var cluster = new DBCluster(&#34;cluster&#34;, DBClusterArgs.builder()        
 *             .dbClusterCategory(&#34;MixedStorage&#34;)
 *             .mode(&#34;flexible&#34;)
 *             .computeResource(&#34;8Core32GB&#34;)
 *             .vswitchId(vswitchId)
 *             .description(name)
 *             .build());
 * 
 *         var defaultAccount = new Account(&#34;defaultAccount&#34;, AccountArgs.builder()        
 *             .dbClusterId(cluster.id())
 *             .accountName(name)
 *             .accountPassword(&#34;tf_example123&#34;)
 *             .accountDescription(name)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ADB account can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:adb/account:Account example am-12345:tf_account
 * ```
 * 
 */
@ResourceType(type="alicloud:adb/account:Account")
public class Account extends com.pulumi.resources.CustomResource {
    /**
     * Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     * 
     */
    @Export(name="accountDescription", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accountDescription;

    /**
     * @return Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     * 
     */
    public Output<Optional<String>> accountDescription() {
        return Codegen.optional(this.accountDescription);
    }
    /**
     * Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
     * 
     */
    @Export(name="accountName", refs={String.class}, tree="[0]")
    private Output<String> accountName;

    /**
     * @return Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
     * 
     */
    public Output<String> accountName() {
        return this.accountName;
    }
    /**
     * Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `account_password` and `kms_encrypted_password` fields.
     * 
     */
    @Export(name="accountPassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accountPassword;

    /**
     * @return Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `account_password` and `kms_encrypted_password` fields.
     * 
     */
    public Output<Optional<String>> accountPassword() {
        return Codegen.optional(this.accountPassword);
    }
    /**
     * The Id of cluster in which account belongs.
     * 
     */
    @Export(name="dbClusterId", refs={String.class}, tree="[0]")
    private Output<String> dbClusterId;

    /**
     * @return The Id of cluster in which account belongs.
     * 
     */
    public Output<String> dbClusterId() {
        return this.dbClusterId;
    }
    /**
     * An KMS encrypts password used to a db account. If the `account_password` is filled in, this field will be ignored.
     * 
     */
    @Export(name="kmsEncryptedPassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsEncryptedPassword;

    /**
     * @return An KMS encrypts password used to a db account. If the `account_password` is filled in, this field will be ignored.
     * 
     */
    public Output<Optional<String>> kmsEncryptedPassword() {
        return Codegen.optional(this.kmsEncryptedPassword);
    }
    /**
     * An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    @Export(name="kmsEncryptionContext", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> kmsEncryptionContext;

    /**
     * @return An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    public Output<Optional<Map<String,Object>>> kmsEncryptionContext() {
        return Codegen.optional(this.kmsEncryptionContext);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Account(String name) {
        this(name, AccountArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Account(String name, AccountArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Account(String name, AccountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:adb/account:Account", name, args == null ? AccountArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Account(String name, Output<String> id, @Nullable AccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:adb/account:Account", name, state, makeResourceOptions(options, id));
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
    public static Account get(String name, Output<String> id, @Nullable AccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Account(name, id, state, options);
    }
}
