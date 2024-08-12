// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alikafka;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.alikafka.SaslAclArgs;
import com.pulumi.alicloud.alikafka.inputs.SaslAclState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides an ALIKAFKA sasl acl resource, see [What is alikafka sasl acl](https://www.alibabacloud.com/help/en/message-queue-for-apache-kafka/latest/api-alikafka-2019-09-16-createacl).
 * 
 * &gt; **NOTE:** Available since v1.66.0.
 * 
 * &gt; **NOTE:**  Only the following regions support create alikafka sasl user.
 * [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`cn-chengdu`,`cn-heyuan`,`ap-southeast-1`,`ap-southeast-3`,`ap-southeast-5`,`ap-northeast-1`,`eu-central-1`,`eu-west-1`,`us-west-1`,`us-east-1`]
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
 * import com.pulumi.alicloud.alikafka.Topic;
 * import com.pulumi.alicloud.alikafka.TopicArgs;
 * import com.pulumi.alicloud.alikafka.SaslUser;
 * import com.pulumi.alicloud.alikafka.SaslUserArgs;
 * import com.pulumi.alicloud.alikafka.SaslAcl;
 * import com.pulumi.alicloud.alikafka.SaslAclArgs;
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
 *         final var name = config.get("name").orElse("tf_example");
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
 *             .config("{\"enable.acl\":\"true\"}")
 *             .vswitchId(defaultSwitch.id())
 *             .securityGroup(defaultSecurityGroup.id())
 *             .build());
 * 
 *         var defaultTopic = new Topic("defaultTopic", TopicArgs.builder()
 *             .instanceId(defaultInstance.id())
 *             .topic("example-topic")
 *             .remark("topic-remark")
 *             .build());
 * 
 *         var defaultSaslUser = new SaslUser("defaultSaslUser", SaslUserArgs.builder()
 *             .instanceId(defaultInstance.id())
 *             .username(name)
 *             .password("tf_example123")
 *             .build());
 * 
 *         var defaultSaslAcl = new SaslAcl("defaultSaslAcl", SaslAclArgs.builder()
 *             .instanceId(defaultInstance.id())
 *             .username(defaultSaslUser.username())
 *             .aclResourceType("Topic")
 *             .aclResourceName(defaultTopic.topic())
 *             .aclResourcePatternType("LITERAL")
 *             .aclOperationType("Write")
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
 * ALIKAFKA GROUP can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:alikafka/saslAcl:SaslAcl acl alikafka_post-cn-123455abc:username:Topic:test-topic:LITERAL:Write
 * ```
 * 
 */
@ResourceType(type="alicloud:alikafka/saslAcl:SaslAcl")
public class SaslAcl extends com.pulumi.resources.CustomResource {
    /**
     * Operation type for this acl. The operation type can only be &#34;Write&#34; and &#34;Read&#34;.
     * 
     */
    @Export(name="aclOperationType", refs={String.class}, tree="[0]")
    private Output<String> aclOperationType;

    /**
     * @return Operation type for this acl. The operation type can only be &#34;Write&#34; and &#34;Read&#34;.
     * 
     */
    public Output<String> aclOperationType() {
        return this.aclOperationType;
    }
    /**
     * Resource name for this acl. The resource name should be a topic or consumer group name.
     * 
     */
    @Export(name="aclResourceName", refs={String.class}, tree="[0]")
    private Output<String> aclResourceName;

    /**
     * @return Resource name for this acl. The resource name should be a topic or consumer group name.
     * 
     */
    public Output<String> aclResourceName() {
        return this.aclResourceName;
    }
    /**
     * Resource pattern type for this acl. The resource pattern support two types &#34;LITERAL&#34; and &#34;PREFIXED&#34;. &#34;LITERAL&#34;: A literal name defines the full name of a resource. The special wildcard character &#34;*&#34; can be used to represent a resource with any name. &#34;PREFIXED&#34;: A prefixed name defines a prefix for a resource.
     * 
     */
    @Export(name="aclResourcePatternType", refs={String.class}, tree="[0]")
    private Output<String> aclResourcePatternType;

    /**
     * @return Resource pattern type for this acl. The resource pattern support two types &#34;LITERAL&#34; and &#34;PREFIXED&#34;. &#34;LITERAL&#34;: A literal name defines the full name of a resource. The special wildcard character &#34;*&#34; can be used to represent a resource with any name. &#34;PREFIXED&#34;: A prefixed name defines a prefix for a resource.
     * 
     */
    public Output<String> aclResourcePatternType() {
        return this.aclResourcePatternType;
    }
    /**
     * Resource type for this acl. The resource type can only be &#34;Topic&#34; and &#34;Group&#34;.
     * 
     */
    @Export(name="aclResourceType", refs={String.class}, tree="[0]")
    private Output<String> aclResourceType;

    /**
     * @return Resource type for this acl. The resource type can only be &#34;Topic&#34; and &#34;Group&#34;.
     * 
     */
    public Output<String> aclResourceType() {
        return this.aclResourceType;
    }
    /**
     * The host of the acl.
     * 
     */
    @Export(name="host", refs={String.class}, tree="[0]")
    private Output<String> host;

    /**
     * @return The host of the acl.
     * 
     */
    public Output<String> host() {
        return this.host;
    }
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
     * Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SaslAcl(java.lang.String name) {
        this(name, SaslAclArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SaslAcl(java.lang.String name, SaslAclArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SaslAcl(java.lang.String name, SaslAclArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alikafka/saslAcl:SaslAcl", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SaslAcl(java.lang.String name, Output<java.lang.String> id, @Nullable SaslAclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alikafka/saslAcl:SaslAcl", name, state, makeResourceOptions(options, id), false);
    }

    private static SaslAclArgs makeArgs(SaslAclArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SaslAclArgs.Empty : args;
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
    public static SaslAcl get(java.lang.String name, Output<java.lang.String> id, @Nullable SaslAclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SaslAcl(name, id, state, options);
    }
}
