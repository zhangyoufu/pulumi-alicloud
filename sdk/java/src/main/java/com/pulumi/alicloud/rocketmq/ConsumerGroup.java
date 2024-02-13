// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rocketmq;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.rocketmq.ConsumerGroupArgs;
import com.pulumi.alicloud.rocketmq.inputs.ConsumerGroupState;
import com.pulumi.alicloud.rocketmq.outputs.ConsumerGroupConsumeRetryPolicy;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a RocketMQ Consumer Group resource.
 * 
 * For information about RocketMQ Consumer Group and how to use it, see [What is Consumer Group](https://www.alibabacloud.com/help/en/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/developer-reference/api-rocketmq-2022-08-01-createconsumergroup).
 * 
 * &gt; **NOTE:** Available since v1.212.0.
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
 * import com.pulumi.alicloud.rocketmq.ConsumerGroup;
 * import com.pulumi.alicloud.rocketmq.ConsumerGroupArgs;
 * import com.pulumi.alicloud.rocketmq.inputs.ConsumerGroupConsumeRetryPolicyArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         var createVpc = new Network(&#34;createVpc&#34;, NetworkArgs.builder()        
 *             .description(&#34;example&#34;)
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .vpcName(name)
 *             .build());
 * 
 *         var createVswitch = new Switch(&#34;createVswitch&#34;, SwitchArgs.builder()        
 *             .description(&#34;example&#34;)
 *             .vpcId(createVpc.id())
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .vswitchName(name)
 *             .build());
 * 
 *         var createInstance = new RocketMQInstance(&#34;createInstance&#34;, RocketMQInstanceArgs.builder()        
 *             .autoRenewPeriod(&#34;1&#34;)
 *             .productInfo(RocketMQInstanceProductInfoArgs.builder()
 *                 .msgProcessSpec(&#34;rmq.p2.4xlarge&#34;)
 *                 .sendReceiveRatio(0.3)
 *                 .messageRetentionTime(&#34;70&#34;)
 *                 .build())
 *             .networkInfo(RocketMQInstanceNetworkInfoArgs.builder()
 *                 .vpcInfo(RocketMQInstanceNetworkInfoVpcInfoArgs.builder()
 *                     .vpcId(createVpc.id())
 *                     .vswitchId(createVswitch.id())
 *                     .build())
 *                 .internetInfo(RocketMQInstanceNetworkInfoInternetInfoArgs.builder()
 *                     .internetSpec(&#34;enable&#34;)
 *                     .flowOutType(&#34;payByBandwidth&#34;)
 *                     .flowOutBandwidth(&#34;30&#34;)
 *                     .build())
 *                 .build())
 *             .period(&#34;1&#34;)
 *             .subSeriesCode(&#34;cluster_ha&#34;)
 *             .remark(&#34;example&#34;)
 *             .instanceName(name)
 *             .serviceCode(&#34;rmq&#34;)
 *             .seriesCode(&#34;professional&#34;)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .periodUnit(&#34;Month&#34;)
 *             .build());
 * 
 *         var defaultConsumerGroup = new ConsumerGroup(&#34;defaultConsumerGroup&#34;, ConsumerGroupArgs.builder()        
 *             .consumerGroupId(name)
 *             .instanceId(createInstance.id())
 *             .consumeRetryPolicy(ConsumerGroupConsumeRetryPolicyArgs.builder()
 *                 .retryPolicy(&#34;DefaultRetryPolicy&#34;)
 *                 .maxRetryTimes(&#34;10&#34;)
 *                 .build())
 *             .deliveryOrderType(&#34;Concurrently&#34;)
 *             .remark(&#34;example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RocketMQ Consumer Group can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:rocketmq/consumerGroup:ConsumerGroup example &lt;instance_id&gt;:&lt;consumer_group_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:rocketmq/consumerGroup:ConsumerGroup")
public class ConsumerGroup extends com.pulumi.resources.CustomResource {
    /**
     * Consumption retry strategy. See `consume_retry_policy` below.
     * 
     */
    @Export(name="consumeRetryPolicy", refs={ConsumerGroupConsumeRetryPolicy.class}, tree="[0]")
    private Output<ConsumerGroupConsumeRetryPolicy> consumeRetryPolicy;

    /**
     * @return Consumption retry strategy. See `consume_retry_policy` below.
     * 
     */
    public Output<ConsumerGroupConsumeRetryPolicy> consumeRetryPolicy() {
        return this.consumeRetryPolicy;
    }
    /**
     * The first ID of the resource.
     * 
     */
    @Export(name="consumerGroupId", refs={String.class}, tree="[0]")
    private Output<String> consumerGroupId;

    /**
     * @return The first ID of the resource.
     * 
     */
    public Output<String> consumerGroupId() {
        return this.consumerGroupId;
    }
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
     * Delivery order.
     * 
     */
    @Export(name="deliveryOrderType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deliveryOrderType;

    /**
     * @return Delivery order.
     * 
     */
    public Output<Optional<String>> deliveryOrderType() {
        return Codegen.optional(this.deliveryOrderType);
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ConsumerGroup(String name) {
        this(name, ConsumerGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ConsumerGroup(String name, ConsumerGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ConsumerGroup(String name, ConsumerGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rocketmq/consumerGroup:ConsumerGroup", name, args == null ? ConsumerGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ConsumerGroup(String name, Output<String> id, @Nullable ConsumerGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rocketmq/consumerGroup:ConsumerGroup", name, state, makeResourceOptions(options, id));
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
    public static ConsumerGroup get(String name, Output<String> id, @Nullable ConsumerGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ConsumerGroup(name, id, state, options);
    }
}
