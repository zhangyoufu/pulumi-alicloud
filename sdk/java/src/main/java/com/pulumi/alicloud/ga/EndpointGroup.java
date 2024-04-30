// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ga.EndpointGroupArgs;
import com.pulumi.alicloud.ga.inputs.EndpointGroupState;
import com.pulumi.alicloud.ga.outputs.EndpointGroupEndpointConfiguration;
import com.pulumi.alicloud.ga.outputs.EndpointGroupPortOverrides;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.alicloud.ga.Accelerator;
 * import com.pulumi.alicloud.ga.AcceleratorArgs;
 * import com.pulumi.alicloud.ga.BandwidthPackage;
 * import com.pulumi.alicloud.ga.BandwidthPackageArgs;
 * import com.pulumi.alicloud.ga.BandwidthPackageAttachment;
 * import com.pulumi.alicloud.ga.BandwidthPackageAttachmentArgs;
 * import com.pulumi.alicloud.ga.Listener;
 * import com.pulumi.alicloud.ga.ListenerArgs;
 * import com.pulumi.alicloud.ga.inputs.ListenerPortRangeArgs;
 * import com.pulumi.alicloud.ecs.EipAddress;
 * import com.pulumi.alicloud.ecs.EipAddressArgs;
 * import com.pulumi.alicloud.ga.EndpointGroup;
 * import com.pulumi.alicloud.ga.EndpointGroupArgs;
 * import com.pulumi.alicloud.ga.inputs.EndpointGroupEndpointConfigurationArgs;
 * import com.pulumi.codegen.internal.KeyedValue;
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
 *         final var region = config.get(&#34;region&#34;).orElse(&#34;cn-hangzhou&#34;);
 *         var default_ = new Accelerator(&#34;default&#34;, AcceleratorArgs.builder()        
 *             .duration(1)
 *             .autoUseCoupon(true)
 *             .spec(&#34;1&#34;)
 *             .build());
 * 
 *         var defaultBandwidthPackage = new BandwidthPackage(&#34;defaultBandwidthPackage&#34;, BandwidthPackageArgs.builder()        
 *             .bandwidth(100)
 *             .type(&#34;Basic&#34;)
 *             .bandwidthType(&#34;Basic&#34;)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .billingType(&#34;PayBy95&#34;)
 *             .ratio(30)
 *             .build());
 * 
 *         var defaultBandwidthPackageAttachment = new BandwidthPackageAttachment(&#34;defaultBandwidthPackageAttachment&#34;, BandwidthPackageAttachmentArgs.builder()        
 *             .acceleratorId(default_.id())
 *             .bandwidthPackageId(defaultBandwidthPackage.id())
 *             .build());
 * 
 *         var defaultListener = new Listener(&#34;defaultListener&#34;, ListenerArgs.builder()        
 *             .acceleratorId(defaultBandwidthPackageAttachment.acceleratorId())
 *             .portRanges(ListenerPortRangeArgs.builder()
 *                 .fromPort(60)
 *                 .toPort(70)
 *                 .build())
 *             .clientAffinity(&#34;SOURCE_IP&#34;)
 *             .protocol(&#34;UDP&#34;)
 *             .name(&#34;terraform-example&#34;)
 *             .build());
 * 
 *         for (var i = 0; i &lt; 2; i++) {
 *             new EipAddress(&#34;defaultEipAddress-&#34; + i, EipAddressArgs.builder()            
 *                 .bandwidth(&#34;10&#34;)
 *                 .internetChargeType(&#34;PayByBandwidth&#34;)
 *                 .addressName(&#34;terraform-example&#34;)
 *                 .build());
 * 
 *         
 * }
 *         var defaultEndpointGroup = new EndpointGroup(&#34;defaultEndpointGroup&#34;, EndpointGroupArgs.builder()        
 *             .acceleratorId(default_.id())
 *             .endpointConfigurations(            
 *                 EndpointGroupEndpointConfigurationArgs.builder()
 *                     .endpoint(defaultEipAddress[0].ipAddress())
 *                     .type(&#34;PublicIp&#34;)
 *                     .weight(&#34;20&#34;)
 *                     .build(),
 *                 EndpointGroupEndpointConfigurationArgs.builder()
 *                     .endpoint(defaultEipAddress[1].ipAddress())
 *                     .type(&#34;PublicIp&#34;)
 *                     .weight(&#34;20&#34;)
 *                     .build())
 *             .endpointGroupRegion(region)
 *             .listenerId(defaultListener.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Ga Endpoint Group can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ga/endpointGroup:EndpointGroup example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ga/endpointGroup:EndpointGroup")
public class EndpointGroup extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Global Accelerator instance to which the endpoint group will be added.
     * 
     */
    @Export(name="acceleratorId", refs={String.class}, tree="[0]")
    private Output<String> acceleratorId;

    /**
     * @return The ID of the Global Accelerator instance to which the endpoint group will be added.
     * 
     */
    public Output<String> acceleratorId() {
        return this.acceleratorId;
    }
    /**
     * The description of the endpoint group.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the endpoint group.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The endpointConfigurations of the endpoint group. See `endpoint_configurations` below.
     * 
     */
    @Export(name="endpointConfigurations", refs={List.class,EndpointGroupEndpointConfiguration.class}, tree="[0,1]")
    private Output<List<EndpointGroupEndpointConfiguration>> endpointConfigurations;

    /**
     * @return The endpointConfigurations of the endpoint group. See `endpoint_configurations` below.
     * 
     */
    public Output<List<EndpointGroupEndpointConfiguration>> endpointConfigurations() {
        return this.endpointConfigurations;
    }
    /**
     * (Available since v1.213.0) The active endpoint IP addresses of the endpoint group. `endpoint_group_ip_list` will change with the growth of network traffic. You can run `pulumi up` to query the latest CIDR blocks and IP addresses.
     * 
     */
    @Export(name="endpointGroupIpLists", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> endpointGroupIpLists;

    /**
     * @return (Available since v1.213.0) The active endpoint IP addresses of the endpoint group. `endpoint_group_ip_list` will change with the growth of network traffic. You can run `pulumi up` to query the latest CIDR blocks and IP addresses.
     * 
     */
    public Output<List<String>> endpointGroupIpLists() {
        return this.endpointGroupIpLists;
    }
    /**
     * The ID of the region where the endpoint group is deployed.
     * 
     */
    @Export(name="endpointGroupRegion", refs={String.class}, tree="[0]")
    private Output<String> endpointGroupRegion;

    /**
     * @return The ID of the region where the endpoint group is deployed.
     * 
     */
    public Output<String> endpointGroupRegion() {
        return this.endpointGroupRegion;
    }
    /**
     * The endpoint group type. Default value: `default`. Valid values: `default`, `virtual`.
     * &gt; **NOTE:** Currently, only `HTTP` or `HTTPS` protocol listener can directly create a `virtual` Endpoint Group. If it is `TCP` protocol listener, and you want to create a `virtual` Endpoint Group, please ensure that the `default` Endpoint Group has been created.
     * 
     */
    @Export(name="endpointGroupType", refs={String.class}, tree="[0]")
    private Output<String> endpointGroupType;

    /**
     * @return The endpoint group type. Default value: `default`. Valid values: `default`, `virtual`.
     * &gt; **NOTE:** Currently, only `HTTP` or `HTTPS` protocol listener can directly create a `virtual` Endpoint Group. If it is `TCP` protocol listener, and you want to create a `virtual` Endpoint Group, please ensure that the `default` Endpoint Group has been created.
     * 
     */
    public Output<String> endpointGroupType() {
        return this.endpointGroupType;
    }
    /**
     * The endpoint request protocol. Valid values: `HTTP`, `HTTPS`.
     * &gt; **NOTE:** This item is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. For the listening instance of HTTP protocol, the back-end service protocol supports and only supports HTTP.
     * 
     */
    @Export(name="endpointRequestProtocol", refs={String.class}, tree="[0]")
    private Output<String> endpointRequestProtocol;

    /**
     * @return The endpoint request protocol. Valid values: `HTTP`, `HTTPS`.
     * &gt; **NOTE:** This item is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. For the listening instance of HTTP protocol, the back-end service protocol supports and only supports HTTP.
     * 
     */
    public Output<String> endpointRequestProtocol() {
        return this.endpointRequestProtocol;
    }
    /**
     * Specifies whether to enable the health check feature. Valid values:
     * 
     */
    @Export(name="healthCheckEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> healthCheckEnabled;

    /**
     * @return Specifies whether to enable the health check feature. Valid values:
     * 
     */
    public Output<Optional<Boolean>> healthCheckEnabled() {
        return Codegen.optional(this.healthCheckEnabled);
    }
    /**
     * The interval between two consecutive health checks. Unit: seconds.
     * 
     */
    @Export(name="healthCheckIntervalSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> healthCheckIntervalSeconds;

    /**
     * @return The interval between two consecutive health checks. Unit: seconds.
     * 
     */
    public Output<Optional<Integer>> healthCheckIntervalSeconds() {
        return Codegen.optional(this.healthCheckIntervalSeconds);
    }
    /**
     * The path specified as the destination of the targets for health checks.
     * 
     */
    @Export(name="healthCheckPath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> healthCheckPath;

    /**
     * @return The path specified as the destination of the targets for health checks.
     * 
     */
    public Output<Optional<String>> healthCheckPath() {
        return Codegen.optional(this.healthCheckPath);
    }
    /**
     * The port that is used for health checks.
     * 
     */
    @Export(name="healthCheckPort", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> healthCheckPort;

    /**
     * @return The port that is used for health checks.
     * 
     */
    public Output<Optional<Integer>> healthCheckPort() {
        return Codegen.optional(this.healthCheckPort);
    }
    /**
     * The protocol that is used to connect to the targets for health checks. Valid values:
     * 
     */
    @Export(name="healthCheckProtocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> healthCheckProtocol;

    /**
     * @return The protocol that is used to connect to the targets for health checks. Valid values:
     * 
     */
    public Output<Optional<String>> healthCheckProtocol() {
        return Codegen.optional(this.healthCheckProtocol);
    }
    /**
     * The ID of the listener that is associated with the endpoint group.
     * 
     */
    @Export(name="listenerId", refs={String.class}, tree="[0]")
    private Output<String> listenerId;

    /**
     * @return The ID of the listener that is associated with the endpoint group.
     * 
     */
    public Output<String> listenerId() {
        return this.listenerId;
    }
    /**
     * The name of the endpoint group.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the endpoint group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Mapping between listening port and forwarding port of boarding point. See `port_overrides` below.
     * &gt; **NOTE:** Port mapping is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. The listening port in the port map must be consistent with the listening port of the current listening instance.
     * 
     */
    @Export(name="portOverrides", refs={EndpointGroupPortOverrides.class}, tree="[0]")
    private Output</* @Nullable */ EndpointGroupPortOverrides> portOverrides;

    /**
     * @return Mapping between listening port and forwarding port of boarding point. See `port_overrides` below.
     * &gt; **NOTE:** Port mapping is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. The listening port in the port map must be consistent with the listening port of the current listening instance.
     * 
     */
    public Output<Optional<EndpointGroupPortOverrides>> portOverrides() {
        return Codegen.optional(this.portOverrides);
    }
    /**
     * The status of the endpoint group.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the endpoint group.
     * 
     */
    public Output<String> status() {
        return this.status;
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
     * The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value: `3`.
     * 
     */
    @Export(name="thresholdCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> thresholdCount;

    /**
     * @return The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value: `3`.
     * 
     */
    public Output<Integer> thresholdCount() {
        return this.thresholdCount;
    }
    /**
     * The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
     * 
     */
    @Export(name="trafficPercentage", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> trafficPercentage;

    /**
     * @return The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
     * 
     */
    public Output<Optional<Integer>> trafficPercentage() {
        return Codegen.optional(this.trafficPercentage);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EndpointGroup(String name) {
        this(name, EndpointGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EndpointGroup(String name, EndpointGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EndpointGroup(String name, EndpointGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/endpointGroup:EndpointGroup", name, args == null ? EndpointGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EndpointGroup(String name, Output<String> id, @Nullable EndpointGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/endpointGroup:EndpointGroup", name, state, makeResourceOptions(options, id));
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
    public static EndpointGroup get(String name, Output<String> id, @Nullable EndpointGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EndpointGroup(name, id, state, options);
    }
}
