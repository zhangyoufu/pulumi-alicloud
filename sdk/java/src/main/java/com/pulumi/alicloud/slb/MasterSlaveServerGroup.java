// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.slb.MasterSlaveServerGroupArgs;
import com.pulumi.alicloud.slb.inputs.MasterSlaveServerGroupState;
import com.pulumi.alicloud.slb.outputs.MasterSlaveServerGroupServer;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A master slave server group contains two ECS instances. The master slave server group can help you to define multiple listening dimension.
 * 
 * &gt; **NOTE:** One ECS instance can be added into multiple master slave server groups.
 * 
 * &gt; **NOTE:** One master slave server group can only add two ECS instances, which are master server and slave server.
 * 
 * &gt; **NOTE:** One master slave server group can be attached with tcp/udp listeners in one load balancer.
 * 
 * &gt; **NOTE:** One Classic and Internet load balancer, its master slave server group can add Classic and VPC ECS instances.
 * 
 * &gt; **NOTE:** One Classic and Intranet load balancer, its master slave server group can only add Classic ECS instances.
 * 
 * &gt; **NOTE:** One VPC load balancer, its master slave server group can only add the same VPC ECS instances.
 * 
 * &gt; **NOTE:** Available in 1.54.0+
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecs.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.ecs.inputs.GetImagesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.Instance;
 * import com.pulumi.alicloud.ecs.InstanceArgs;
 * import com.pulumi.alicloud.slb.ApplicationLoadBalancer;
 * import com.pulumi.alicloud.slb.ApplicationLoadBalancerArgs;
 * import com.pulumi.alicloud.ecs.EcsNetworkInterface;
 * import com.pulumi.alicloud.ecs.EcsNetworkInterfaceArgs;
 * import com.pulumi.alicloud.ecs.EcsNetworkInterfaceAttachment;
 * import com.pulumi.alicloud.ecs.EcsNetworkInterfaceAttachmentArgs;
 * import com.pulumi.alicloud.slb.MasterSlaveServerGroup;
 * import com.pulumi.alicloud.slb.MasterSlaveServerGroupArgs;
 * import com.pulumi.alicloud.slb.inputs.MasterSlaveServerGroupServerArgs;
 * import com.pulumi.alicloud.slb.Listener;
 * import com.pulumi.alicloud.slb.ListenerArgs;
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
 *         final var msServerGroup = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableDiskCategory(&#34;cloud_efficiency&#34;)
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         final var msServerGroupGetInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(msServerGroup.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .eniAmount(2)
 *             .build());
 * 
 *         final var image = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex(&#34;^ubuntu_18.*64&#34;)
 *             .mostRecent(true)
 *             .owners(&#34;system&#34;)
 *             .build());
 * 
 *         final var slbMasterSlaveServerGroup = config.get(&#34;slbMasterSlaveServerGroup&#34;).orElse(&#34;forSlbMasterSlaveServerGroup&#34;);
 *         var main = new Network(&#34;main&#34;, NetworkArgs.builder()        
 *             .vpcName(slbMasterSlaveServerGroup)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var mainSwitch = new Switch(&#34;mainSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(main.id())
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .zoneId(msServerGroup.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .vswitchName(slbMasterSlaveServerGroup)
 *             .build());
 * 
 *         var group = new SecurityGroup(&#34;group&#34;, SecurityGroupArgs.builder()        
 *             .name(slbMasterSlaveServerGroup)
 *             .vpcId(main.id())
 *             .build());
 * 
 *         for (var i = 0; i &lt; 2; i++) {
 *             new Instance(&#34;msServerGroupInstance-&#34; + i, InstanceArgs.builder()            
 *                 .imageId(image.applyValue(getImagesResult -&gt; getImagesResult.images()[0].id()))
 *                 .instanceType(msServerGroupGetInstanceTypes.applyValue(getInstanceTypesResult -&gt; getInstanceTypesResult.instanceTypes()[0].id()))
 *                 .instanceName(slbMasterSlaveServerGroup)
 *                 .securityGroups(group.id())
 *                 .internetChargeType(&#34;PayByTraffic&#34;)
 *                 .internetMaxBandwidthOut(&#34;10&#34;)
 *                 .availabilityZone(msServerGroup.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *                 .instanceChargeType(&#34;PostPaid&#34;)
 *                 .systemDiskCategory(&#34;cloud_efficiency&#34;)
 *                 .vswitchId(mainSwitch.id())
 *                 .build());
 * 
 *         
 * }
 *         var msServerGroupApplicationLoadBalancer = new ApplicationLoadBalancer(&#34;msServerGroupApplicationLoadBalancer&#34;, ApplicationLoadBalancerArgs.builder()        
 *             .loadBalancerName(slbMasterSlaveServerGroup)
 *             .vswitchId(mainSwitch.id())
 *             .loadBalancerSpec(&#34;slb.s2.small&#34;)
 *             .build());
 * 
 *         var msServerGroupEcsNetworkInterface = new EcsNetworkInterface(&#34;msServerGroupEcsNetworkInterface&#34;, EcsNetworkInterfaceArgs.builder()        
 *             .networkInterfaceName(slbMasterSlaveServerGroup)
 *             .vswitchId(mainSwitch.id())
 *             .securityGroupIds(group.id())
 *             .build());
 * 
 *         var msServerGroupEcsNetworkInterfaceAttachment = new EcsNetworkInterfaceAttachment(&#34;msServerGroupEcsNetworkInterfaceAttachment&#34;, EcsNetworkInterfaceAttachmentArgs.builder()        
 *             .instanceId(msServerGroupInstance[0].id())
 *             .networkInterfaceId(msServerGroupEcsNetworkInterface.id())
 *             .build());
 * 
 *         var groupMasterSlaveServerGroup = new MasterSlaveServerGroup(&#34;groupMasterSlaveServerGroup&#34;, MasterSlaveServerGroupArgs.builder()        
 *             .loadBalancerId(msServerGroupApplicationLoadBalancer.id())
 *             .name(slbMasterSlaveServerGroup)
 *             .servers(            
 *                 MasterSlaveServerGroupServerArgs.builder()
 *                     .serverId(msServerGroupInstance[0].id())
 *                     .port(100)
 *                     .weight(100)
 *                     .serverType(&#34;Master&#34;)
 *                     .build(),
 *                 MasterSlaveServerGroupServerArgs.builder()
 *                     .serverId(msServerGroupInstance[1].id())
 *                     .port(100)
 *                     .weight(100)
 *                     .serverType(&#34;Slave&#34;)
 *                     .build())
 *             .build());
 * 
 *         var tcp = new Listener(&#34;tcp&#34;, ListenerArgs.builder()        
 *             .loadBalancerId(msServerGroupApplicationLoadBalancer.id())
 *             .masterSlaveServerGroupId(groupMasterSlaveServerGroup.id())
 *             .frontendPort(&#34;22&#34;)
 *             .protocol(&#34;tcp&#34;)
 *             .bandwidth(&#34;10&#34;)
 *             .healthCheckType(&#34;tcp&#34;)
 *             .persistenceTimeout(3600)
 *             .healthyThreshold(8)
 *             .unhealthyThreshold(8)
 *             .healthCheckTimeout(8)
 *             .healthCheckInterval(5)
 *             .healthCheckHttpCode(&#34;http_2xx&#34;)
 *             .healthCheckConnectPort(20)
 *             .healthCheckUri(&#34;/console&#34;)
 *             .establishedTimeout(600)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Block servers
 * 
 * The servers mapping supports the following:
 * 
 * * `server_ids` - (Required) A list backend server ID (ECS instance ID).
 * * `port` - (Required) The port used by the backend server. Valid value range: [1-65535].
 * * `weight` - (Optional) Weight of the backend server. Valid value range: [0-100]. Default to 100.
 * * `type` - (Optional, Available in 1.51.0+) Type of the backend server. Valid value ecs, eni. Default to eni.
 * * `server_type` - (Optional) The server type of the backend server. Valid value Master, Slave.
 * * `is_backup` - (Removed from v1.63.0) Determine if the server is executing. Valid value 0, 1.
 * 
 * ## Import
 * 
 * Load balancer master slave server group can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup example abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup")
public class MasterSlaveServerGroup extends com.pulumi.resources.CustomResource {
    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     * 
     */
    @Export(name="deleteProtectionValidation", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deleteProtectionValidation;

    /**
     * @return Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     * 
     */
    public Output<Optional<Boolean>> deleteProtectionValidation() {
        return Codegen.optional(this.deleteProtectionValidation);
    }
    /**
     * The Load Balancer ID which is used to launch a new master slave server group.
     * 
     */
    @Export(name="loadBalancerId", refs={String.class}, tree="[0]")
    private Output<String> loadBalancerId;

    /**
     * @return The Load Balancer ID which is used to launch a new master slave server group.
     * 
     */
    public Output<String> loadBalancerId() {
        return this.loadBalancerId;
    }
    /**
     * Name of the master slave server group.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the master slave server group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
     * 
     */
    @Export(name="servers", refs={List.class,MasterSlaveServerGroupServer.class}, tree="[0,1]")
    private Output</* @Nullable */ List<MasterSlaveServerGroupServer>> servers;

    /**
     * @return A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
     * 
     */
    public Output<Optional<List<MasterSlaveServerGroupServer>>> servers() {
        return Codegen.optional(this.servers);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MasterSlaveServerGroup(String name) {
        this(name, MasterSlaveServerGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MasterSlaveServerGroup(String name, MasterSlaveServerGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MasterSlaveServerGroup(String name, MasterSlaveServerGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup", name, args == null ? MasterSlaveServerGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MasterSlaveServerGroup(String name, Output<String> id, @Nullable MasterSlaveServerGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup", name, state, makeResourceOptions(options, id));
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
    public static MasterSlaveServerGroup get(String name, Output<String> id, @Nullable MasterSlaveServerGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MasterSlaveServerGroup(name, id, state, options);
    }
}
