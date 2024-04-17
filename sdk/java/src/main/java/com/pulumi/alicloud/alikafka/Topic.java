// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alikafka;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.alikafka.TopicArgs;
import com.pulumi.alicloud.alikafka.inputs.TopicState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an ALIKAFKA topic resource, see [What is Alikafka topic ](https://www.alibabacloud.com/help/en/message-queue-for-apache-kafka/latest/api-alikafka-2019-09-16-createtopic).
 * 
 * &gt; **NOTE:** Available since v1.56.0.
 * 
 * &gt; **NOTE:**  Only the following regions support create alikafka topic.
 * [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`cn-chengdu`,`cn-heyuan`,`ap-southeast-1`,`ap-southeast-3`,`ap-southeast-5`,`ap-northeast-1`,`eu-central-1`,`eu-west-1`,`us-west-1`,`us-east-1`]
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.alikafka.Instance;
 * import com.pulumi.alicloud.alikafka.InstanceArgs;
 * import com.pulumi.alicloud.alikafka.Topic;
 * import com.pulumi.alicloud.alikafka.TopicArgs;
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
 *         final var instanceName = config.get(&#34;instanceName&#34;).orElse(&#34;tf-example&#34;);
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         var defaultInteger = new Integer(&#34;defaultInteger&#34;, IntegerArgs.builder()        
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .name(String.format(&#34;%s-%s&#34;, instanceName,defaultInteger.result()))
 *             .partitionNum(&#34;50&#34;)
 *             .diskType(&#34;1&#34;)
 *             .diskSize(&#34;500&#34;)
 *             .deployType(&#34;5&#34;)
 *             .ioMax(&#34;20&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .securityGroup(defaultSecurityGroup.id())
 *             .build());
 * 
 *         var defaultTopic = new Topic(&#34;defaultTopic&#34;, TopicArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .topic(&#34;example-topic&#34;)
 *             .localTopic(&#34;false&#34;)
 *             .compactTopic(&#34;false&#34;)
 *             .partitionNum(&#34;12&#34;)
 *             .remark(&#34;dafault_kafka_topic_remark&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ALIKAFKA TOPIC can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:alikafka/topic:Topic topic alikafka_post-cn-123455abc:topicName
 * ```
 * 
 */
@ResourceType(type="alicloud:alikafka/topic:Topic")
public class Topic extends com.pulumi.resources.CustomResource {
    /**
     * Whether the topic is compactTopic or not. Compact topic must be a localTopic.
     * 
     */
    @Export(name="compactTopic", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> compactTopic;

    /**
     * @return Whether the topic is compactTopic or not. Compact topic must be a localTopic.
     * 
     */
    public Output<Optional<Boolean>> compactTopic() {
        return Codegen.optional(this.compactTopic);
    }
    /**
     * InstanceId of your Kafka resource, the topic will create in this instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return InstanceId of your Kafka resource, the topic will create in this instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Whether the topic is localTopic or not.
     * 
     */
    @Export(name="localTopic", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> localTopic;

    /**
     * @return Whether the topic is localTopic or not.
     * 
     */
    public Output<Optional<Boolean>> localTopic() {
        return Codegen.optional(this.localTopic);
    }
    /**
     * The number of partitions of the topic. The number should between 1 and 48.
     * 
     */
    @Export(name="partitionNum", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> partitionNum;

    /**
     * @return The number of partitions of the topic. The number should between 1 and 48.
     * 
     */
    public Output<Optional<Integer>> partitionNum() {
        return Codegen.optional(this.partitionNum);
    }
    /**
     * This attribute is a concise description of topic. The length cannot exceed 64.
     * 
     */
    @Export(name="remark", refs={String.class}, tree="[0]")
    private Output<String> remark;

    /**
     * @return This attribute is a concise description of topic. The length cannot exceed 64.
     * 
     */
    public Output<String> remark() {
        return this.remark;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Name of the topic. Two topics on a single instance cannot have the same name. The length cannot exceed 249 characters.
     * 
     */
    @Export(name="topic", refs={String.class}, tree="[0]")
    private Output<String> topic;

    /**
     * @return Name of the topic. Two topics on a single instance cannot have the same name. The length cannot exceed 249 characters.
     * 
     */
    public Output<String> topic() {
        return this.topic;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Topic(String name) {
        this(name, TopicArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Topic(String name, TopicArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Topic(String name, TopicArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alikafka/topic:Topic", name, args == null ? TopicArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Topic(String name, Output<String> id, @Nullable TopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alikafka/topic:Topic", name, state, makeResourceOptions(options, id));
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
    public static Topic get(String name, Output<String> id, @Nullable TopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Topic(name, id, state, options);
    }
}
