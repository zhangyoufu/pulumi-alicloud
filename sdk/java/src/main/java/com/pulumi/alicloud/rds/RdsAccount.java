// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.rds.RdsAccountArgs;
import com.pulumi.alicloud.rds.inputs.RdsAccountState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a RDS Account resource.
 * 
 * For information about RDS Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/api-rds-2014-08-15-createaccount).
 * 
 * &gt; **NOTE:** Available since v1.120.0.
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
 * import com.pulumi.alicloud.rds.RdsFunctions;
 * import com.pulumi.alicloud.rds.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.rds.inputs.GetInstanceClassesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.rds.Instance;
 * import com.pulumi.alicloud.rds.InstanceArgs;
 * import com.pulumi.alicloud.rds.RdsAccount;
 * import com.pulumi.alicloud.rds.RdsAccountArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf_example&#34;);
 *         final var defaultZones = RdsFunctions.getZones(GetZonesArgs.builder()
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;5.6&#34;)
 *             .build());
 * 
 *         final var defaultInstanceClasses = RdsFunctions.getInstanceClasses(GetInstanceClassesArgs.builder()
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.ids()[0]))
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;5.6&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.ids()[0]))
 *             .vswitchName(name)
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;5.6&#34;)
 *             .instanceType(defaultInstanceClasses.applyValue(getInstanceClassesResult -&gt; getInstanceClassesResult.instanceClasses()[0].instanceClass()))
 *             .instanceStorage(&#34;10&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .instanceName(name)
 *             .build());
 * 
 *         var defaultRdsAccount = new RdsAccount(&#34;defaultRdsAccount&#34;, RdsAccountArgs.builder()        
 *             .dbInstanceId(defaultInstance.id())
 *             .accountName(name)
 *             .accountPassword(&#34;Example1234&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * RDS Account can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:rds/rdsAccount:RdsAccount example &lt;db_instance_id&gt;:&lt;account_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:rds/rdsAccount:RdsAccount")
public class RdsAccount extends com.pulumi.resources.CustomResource {
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     * 
     */
    @Export(name="accountDescription", refs={String.class}, tree="[0]")
    private Output<String> accountDescription;

    /**
     * @return Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     * 
     */
    public Output<String> accountDescription() {
        return this.accountDescription;
    }
    /**
     * Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and end with letters or numbers, The length must be 2-63 characters for PostgreSQL, otherwise the length must be 2-32 characters.
     * 
     */
    @Export(name="accountName", refs={String.class}, tree="[0]")
    private Output<String> accountName;

    /**
     * @return Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and end with letters or numbers, The length must be 2-63 characters for PostgreSQL, otherwise the length must be 2-32 characters.
     * 
     */
    public Output<String> accountName() {
        return this.accountName;
    }
    /**
     * Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
     * 
     */
    @Export(name="accountPassword", refs={String.class}, tree="[0]")
    private Output<String> accountPassword;

    /**
     * @return Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
     * 
     */
    public Output<String> accountPassword() {
        return this.accountPassword;
    }
    /**
     * Privilege type of account. Default to `Normal`.
     * `Normal`: Common privilege.
     * `Super`: High privilege.
     * 
     */
    @Export(name="accountType", refs={String.class}, tree="[0]")
    private Output<String> accountType;

    /**
     * @return Privilege type of account. Default to `Normal`.
     * `Normal`: Common privilege.
     * `Super`: High privilege.
     * 
     */
    public Output<String> accountType() {
        return this.accountType;
    }
    /**
     * The Id of instance in which account belongs.
     * 
     */
    @Export(name="dbInstanceId", refs={String.class}, tree="[0]")
    private Output<String> dbInstanceId;

    /**
     * @return The Id of instance in which account belongs.
     * 
     */
    public Output<String> dbInstanceId() {
        return this.dbInstanceId;
    }
    /**
     * The attribute has been deprecated from 1.120.0 and using `account_description` instead.
     * 
     * @deprecated
     * Field &#39;description&#39; has been deprecated from provider version 1.120.0. New field &#39;account_description&#39; instead.
     * 
     */
    @Deprecated /* Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead. */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The attribute has been deprecated from 1.120.0 and using `account_description` instead.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The attribute has been deprecated from 1.120.0 and using `db_instance_id` instead.
     * 
     * @deprecated
     * Field &#39;instance_id&#39; has been deprecated from provider version 1.120.0. New field &#39;db_instance_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead. */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The attribute has been deprecated from 1.120.0 and using `db_instance_id` instead.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
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
     * The attribute has been deprecated from 1.120.0 and using `account_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.120.0. New field &#39;account_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead. */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The attribute has been deprecated from 1.120.0 and using `account_name` instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The attribute has been deprecated from 1.120.0 and using `account_password` instead.
     * 
     * @deprecated
     * Field &#39;password&#39; has been deprecated from provider version 1.120.0. New field &#39;account_password&#39; instead.
     * 
     */
    @Deprecated /* Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead. */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return The attribute has been deprecated from 1.120.0 and using `account_password` instead.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * Resets permissions flag of the privileged account. Default to `false`. Set it to `true` can resets permissions of the privileged account.
     * 
     */
    @Export(name="resetPermissionFlag", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> resetPermissionFlag;

    /**
     * @return Resets permissions flag of the privileged account. Default to `false`. Set it to `true` can resets permissions of the privileged account.
     * 
     */
    public Output<Optional<Boolean>> resetPermissionFlag() {
        return Codegen.optional(this.resetPermissionFlag);
    }
    /**
     * The status of the resource. Valid values: `Available`, `Unavailable`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource. Valid values: `Available`, `Unavailable`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The attribute has been deprecated from 1.120.0 and using `account_type` instead.
     * 
     * &gt; **NOTE**: Only MySQL engine is supported resets permissions of the privileged account.
     * 
     * @deprecated
     * Field &#39;type&#39; has been deprecated from provider version 1.120.0. New field &#39;account_type&#39; instead.
     * 
     */
    @Deprecated /* Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead. */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The attribute has been deprecated from 1.120.0 and using `account_type` instead.
     * 
     * &gt; **NOTE**: Only MySQL engine is supported resets permissions of the privileged account.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RdsAccount(String name) {
        this(name, RdsAccountArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RdsAccount(String name, @Nullable RdsAccountArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RdsAccount(String name, @Nullable RdsAccountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/rdsAccount:RdsAccount", name, args == null ? RdsAccountArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RdsAccount(String name, Output<String> id, @Nullable RdsAccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/rdsAccount:RdsAccount", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accountPassword",
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
    public static RdsAccount get(String name, Output<String> id, @Nullable RdsAccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RdsAccount(name, id, state, options);
    }
}
