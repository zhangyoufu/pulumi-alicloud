// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alikafka;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.alikafka.SaslUserArgs;
import com.pulumi.alicloud.alikafka.inputs.SaslUserState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Alikafka Sasl User resource.
 * 
 * &gt; **NOTE:** Available since v1.66.0.
 * 
 * &gt; **NOTE:**  Only the following regions support create alikafka Sasl User.
 * [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`cn-chengdu`,`cn-heyuan`,`ap-southeast-1`,`ap-southeast-3`,`ap-southeast-5`,`ap-northeast-1`,`eu-central-1`,`eu-west-1`,`us-west-1`,`us-east-1`]
 * 
 * For information about Alikafka Sasl User and how to use it, see [What is Sasl User](https://www.alibabacloud.com/help/en/message-queue-for-apache-kafka/latest/api-alikafka-2019-09-16-createsasluser).
 * 
 * ## Example Usage
 * 
 * Basic Usage
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
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.alikafka.Instance;
 * import com.pulumi.alicloud.alikafka.InstanceArgs;
 * import com.pulumi.alicloud.alikafka.SaslUser;
 * import com.pulumi.alicloud.alikafka.SaslUserArgs;
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
 *         final var name = config.get("name").orElse("tf-example");
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation("VSwitch")
 *             .build());
 * 
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()
 *             .vpcName(name)
 *             .cidrBlock("10.4.0.0/16")
 *             .build());
 * 
 *         var defaultSwitch = new Switch("defaultSwitch", SwitchArgs.builder()
 *             .vswitchName(name)
 *             .cidrBlock("10.4.0.0/24")
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup("defaultSecurityGroup", SecurityGroupArgs.builder()
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultInteger = new Integer("defaultInteger", IntegerArgs.builder()
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var defaultInstance = new Instance("defaultInstance", InstanceArgs.builder()
 *             .name(String.format("%s-%s", name,defaultInteger.result()))
 *             .partitionNum(50)
 *             .diskType("1")
 *             .diskSize("500")
 *             .deployType("5")
 *             .ioMax("20")
 *             .specType("professional")
 *             .serviceVersion("2.2.0")
 *             .vswitchId(defaultSwitch.id())
 *             .securityGroup(defaultSecurityGroup.id())
 *             .config("""
 *   {
 *     "enable.acl": "true"
 *   }
 *             """)
 *             .build());
 * 
 *         var defaultSaslUser = new SaslUser("defaultSaslUser", SaslUserArgs.builder()
 *             .instanceId(defaultInstance.id())
 *             .username(name)
 *             .password("tf_example123")
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
 * Alikafka Sasl User can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:alikafka/saslUser:SaslUser example &lt;instance_id&gt;:&lt;username&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:alikafka/saslUser:SaslUser")
public class SaslUser extends com.pulumi.resources.CustomResource {
    /**
     * ID of the ALIKAFKA Instance that owns the groups.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return ID of the ALIKAFKA Instance that owns the groups.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * An KMS encrypts password used to a db account. You have to specify one of `password` and `kms_encrypted_password` fields.
     * 
     */
    @Export(name="kmsEncryptedPassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsEncryptedPassword;

    /**
     * @return An KMS encrypts password used to a db account. You have to specify one of `password` and `kms_encrypted_password` fields.
     * 
     */
    public Output<Optional<String>> kmsEncryptedPassword() {
        return Codegen.optional(this.kmsEncryptedPassword);
    }
    /**
     * An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a user with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    @Export(name="kmsEncryptionContext", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> kmsEncryptionContext;

    /**
     * @return An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a user with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    public Output<Optional<Map<String,String>>> kmsEncryptionContext() {
        return Codegen.optional(this.kmsEncryptionContext);
    }
    /**
     * The password of the SASL user. It may consist of letters, digits, or underlines, with a length of 1 to 64 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return The password of the SASL user. It may consist of letters, digits, or underlines, with a length of 1 to 64 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * The authentication mechanism. Default value: `plain`. Valid values: `plain`, `scram`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The authentication mechanism. Default value: `plain`. Valid values: `plain`, `scram`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * The name of the SASL user. The length should between `1` to `64` characters. The characters can only contain `a`-`z`, `A`-`Z`, `0`-`9`, `_` and `-`.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return The name of the SASL user. The length should between `1` to `64` characters. The characters can only contain `a`-`z`, `A`-`Z`, `0`-`9`, `_` and `-`.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SaslUser(java.lang.String name) {
        this(name, SaslUserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SaslUser(java.lang.String name, SaslUserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SaslUser(java.lang.String name, SaslUserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alikafka/saslUser:SaslUser", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SaslUser(java.lang.String name, Output<java.lang.String> id, @Nullable SaslUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alikafka/saslUser:SaslUser", name, state, makeResourceOptions(options, id), false);
    }

    private static SaslUserArgs makeArgs(SaslUserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SaslUserArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static SaslUser get(java.lang.String name, Output<java.lang.String> id, @Nullable SaslUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SaslUser(name, id, state, options);
    }
}
