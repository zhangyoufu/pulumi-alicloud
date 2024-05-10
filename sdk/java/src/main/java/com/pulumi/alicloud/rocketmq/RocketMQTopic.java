// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rocketmq;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.rocketmq.RocketMQTopicArgs;
import com.pulumi.alicloud.rocketmq.inputs.RocketMQTopicState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a RocketMQ Topic resource.
 * 
 * For information about RocketMQ Topic and how to use it, see [What is Topic](https://www.alibabacloud.com/help/en/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/developer-reference/api-rocketmq-2022-08-01-createtopic).
 * 
 * &gt; **NOTE:** Available since v1.211.0.
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
 * import com.pulumi.alicloud.rocketmq.RocketMQInstance;
 * import com.pulumi.alicloud.rocketmq.RocketMQInstanceArgs;
 * import com.pulumi.alicloud.rocketmq.inputs.RocketMQInstanceProductInfoArgs;
 * import com.pulumi.alicloud.rocketmq.inputs.RocketMQInstanceNetworkInfoArgs;
 * import com.pulumi.alicloud.rocketmq.inputs.RocketMQInstanceNetworkInfoVpcInfoArgs;
 * import com.pulumi.alicloud.rocketmq.inputs.RocketMQInstanceNetworkInfoInternetInfoArgs;
 * import com.pulumi.alicloud.rocketmq.RocketMQTopic;
 * import com.pulumi.alicloud.rocketmq.RocketMQTopicArgs;
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
 *         final var name = config.get("name").orElse("terraform-example");
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation("VSwitch")
 *             .build());
 * 
 *         var createVpc = new Network("createVpc", NetworkArgs.builder()        
 *             .description("example")
 *             .cidrBlock("172.16.0.0/12")
 *             .vpcName(name)
 *             .build());
 * 
 *         var createVswitch = new Switch("createVswitch", SwitchArgs.builder()        
 *             .description("example")
 *             .vpcId(createVpc.id())
 *             .zoneId(default_.zones()[0].id())
 *             .cidrBlock("172.16.0.0/24")
 *             .vswitchName(name)
 *             .build());
 * 
 *         var createInstance = new RocketMQInstance("createInstance", RocketMQInstanceArgs.builder()        
 *             .autoRenewPeriod("1")
 *             .productInfo(RocketMQInstanceProductInfoArgs.builder()
 *                 .msgProcessSpec("rmq.p2.4xlarge")
 *                 .sendReceiveRatio(0.3)
 *                 .messageRetentionTime("70")
 *                 .build())
 *             .networkInfo(RocketMQInstanceNetworkInfoArgs.builder()
 *                 .vpcInfo(RocketMQInstanceNetworkInfoVpcInfoArgs.builder()
 *                     .vpcId(createVpc.id())
 *                     .vswitchId(createVswitch.id())
 *                     .build())
 *                 .internetInfo(RocketMQInstanceNetworkInfoInternetInfoArgs.builder()
 *                     .internetSpec("enable")
 *                     .flowOutType("payByBandwidth")
 *                     .flowOutBandwidth("30")
 *                     .build())
 *                 .build())
 *             .period("1")
 *             .subSeriesCode("cluster_ha")
 *             .remark("example")
 *             .instanceName(name)
 *             .serviceCode("rmq")
 *             .seriesCode("professional")
 *             .paymentType("PayAsYouGo")
 *             .periodUnit("Month")
 *             .build());
 * 
 *         var defaultRocketMQTopic = new RocketMQTopic("defaultRocketMQTopic", RocketMQTopicArgs.builder()        
 *             .remark("example")
 *             .instanceId(createInstance.id())
 *             .messageType("NORMAL")
 *             .topicName(name)
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
 * RocketMQ Topic can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:rocketmq/rocketMQTopic:RocketMQTopic example &lt;instance_id&gt;:&lt;topic_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:rocketmq/rocketMQTopic:RocketMQTopic")
public class RocketMQTopic extends com.pulumi.resources.CustomResource {
    /**
     * The creation time of the resource.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The creation time of the resource.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Instance ID.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return Instance ID.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Message type.
     * 
     */
    @Export(name="messageType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> messageType;

    /**
     * @return Message type.
     * 
     */
    public Output<Optional<String>> messageType() {
        return Codegen.optional(this.messageType);
    }
    /**
     * Custom remarks.
     * 
     */
    @Export(name="remark", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> remark;

    /**
     * @return Custom remarks.
     * 
     */
    public Output<Optional<String>> remark() {
        return Codegen.optional(this.remark);
    }
    /**
     * The status of the resource.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Topic name and identification.
     * 
     */
    @Export(name="topicName", refs={String.class}, tree="[0]")
    private Output<String> topicName;

    /**
     * @return Topic name and identification.
     * 
     */
    public Output<String> topicName() {
        return this.topicName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RocketMQTopic(String name) {
        this(name, RocketMQTopicArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RocketMQTopic(String name, RocketMQTopicArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RocketMQTopic(String name, RocketMQTopicArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rocketmq/rocketMQTopic:RocketMQTopic", name, args == null ? RocketMQTopicArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RocketMQTopic(String name, Output<String> id, @Nullable RocketMQTopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rocketmq/rocketMQTopic:RocketMQTopic", name, state, makeResourceOptions(options, id));
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
    public static RocketMQTopic get(String name, Output<String> id, @Nullable RocketMQTopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RocketMQTopic(name, id, state, options);
    }
}
