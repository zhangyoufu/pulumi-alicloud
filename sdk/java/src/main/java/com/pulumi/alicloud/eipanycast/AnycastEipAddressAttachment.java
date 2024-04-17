// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eipanycast;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eipanycast.AnycastEipAddressAttachmentArgs;
import com.pulumi.alicloud.eipanycast.inputs.AnycastEipAddressAttachmentState;
import com.pulumi.alicloud.eipanycast.outputs.AnycastEipAddressAttachmentPopLocation;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Eipanycast Anycast Eip Address Attachment resource.
 * 
 * For information about Eipanycast Anycast Eip Address Attachment and how to use it, see [What is Anycast Eip Address Attachment](https://www.alibabacloud.com/help/en/anycast-eip/latest/api-eipanycast-2020-03-09-associateanycasteipaddress).
 * 
 * &gt; **NOTE:** Available since v1.113.0.
 * 
 * &gt; **NOTE:** The following regions support currently while Slb instance support bound.
 * [eu-west-1-gb33-a01,cn-hongkong-am4-c04,ap-southeast-os30-a01,us-west-ot7-a01,ap-south-in73-a01,ap-southeast-my88-a01]
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
 * import com.pulumi.alicloud.slb.SlbFunctions;
 * import com.pulumi.alicloud.slb.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.slb.ApplicationLoadBalancer;
 * import com.pulumi.alicloud.slb.ApplicationLoadBalancerArgs;
 * import com.pulumi.alicloud.eipanycast.AnycastEipAddress;
 * import com.pulumi.alicloud.eipanycast.AnycastEipAddressArgs;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetRegionsArgs;
 * import com.pulumi.alicloud.eipanycast.AnycastEipAddressAttachment;
 * import com.pulumi.alicloud.eipanycast.AnycastEipAddressAttachmentArgs;
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
 *         final var default = SlbFunctions.getZones(GetZonesArgs.builder()
 *             .availableSlbAddressType(&#34;vpc&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;10.0.0.0/8&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock(&#34;10.1.0.0/16&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var defaultApplicationLoadBalancer = new ApplicationLoadBalancer(&#34;defaultApplicationLoadBalancer&#34;, ApplicationLoadBalancerArgs.builder()        
 *             .addressType(&#34;intranet&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .loadBalancerName(name)
 *             .loadBalancerSpec(&#34;slb.s1.small&#34;)
 *             .masterZoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var defaultAnycastEipAddress = new AnycastEipAddress(&#34;defaultAnycastEipAddress&#34;, AnycastEipAddressArgs.builder()        
 *             .anycastEipAddressName(name)
 *             .serviceLocation(&#34;ChineseMainland&#34;)
 *             .build());
 * 
 *         final var defaultGetRegions = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         var defaultAnycastEipAddressAttachment = new AnycastEipAddressAttachment(&#34;defaultAnycastEipAddressAttachment&#34;, AnycastEipAddressAttachmentArgs.builder()        
 *             .bindInstanceId(defaultApplicationLoadBalancer.id())
 *             .bindInstanceType(&#34;SlbInstance&#34;)
 *             .bindInstanceRegionId(defaultGetRegions.applyValue(getRegionsResult -&gt; getRegionsResult.regions()[0].id()))
 *             .anycastId(defaultAnycastEipAddress.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * Multiple Usage
 * 
 * &gt; **NOTE:**  Anycast EIP supports binding cloud resource instances in multiple regions. Only one cloud resource instance is supported as the default origin station, and the rest are normal origin stations. When no access point is specified or an access point is added, the access request is forwarded to the default origin by default.  If you are bound for the first time, the Default value of the binding mode is **Default * *. /li&gt; li&gt; If you are not binding for the first time, you can set the binding mode to **Default**, and the new Default origin will take effect. The original Default origin will be changed to a common origin.
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
 * import com.pulumi.alicloud.ecs.inputs.GetImagesArgs;
 * import com.pulumi.alicloud.ecs.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.Instance;
 * import com.pulumi.alicloud.ecs.InstanceArgs;
 * import com.pulumi.alicloud.eipanycast.AnycastEipAddress;
 * import com.pulumi.alicloud.eipanycast.AnycastEipAddressArgs;
 * import com.pulumi.alicloud.eipanycast.AnycastEipAddressAttachment;
 * import com.pulumi.alicloud.eipanycast.AnycastEipAddressAttachmentArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-example&#34;);
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableDiskCategory(&#34;cloud_efficiency&#34;)
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         final var defaultGetImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex(&#34;^ubuntu_18.*64&#34;)
 *             .mostRecent(true)
 *             .owners(&#34;system&#34;)
 *             .build());
 * 
 *         final var defaultGetInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(default_.zones()[0].id())
 *             .cpuCoreCount(1)
 *             .memorySize(2)
 *             .build());
 * 
 *         var defaultVpc = new Network(&#34;defaultVpc&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;192.168.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultVsw = new Switch(&#34;defaultVsw&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultVpc.id())
 *             .cidrBlock(&#34;192.168.0.0/24&#34;)
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var defaultuBsECI = new SecurityGroup(&#34;defaultuBsECI&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(defaultVpc.id())
 *             .build());
 * 
 *         var default9KDlN7 = new Instance(&#34;default9KDlN7&#34;, InstanceArgs.builder()        
 *             .imageId(defaultGetImages.applyValue(getImagesResult -&gt; getImagesResult.images()[0].id()))
 *             .instanceType(defaultGetInstanceTypes.applyValue(getInstanceTypesResult -&gt; getInstanceTypesResult.instanceTypes()[0].id()))
 *             .instanceName(name)
 *             .securityGroups(defaultuBsECI.id())
 *             .availabilityZone(defaultVsw.zoneId())
 *             .instanceChargeType(&#34;PostPaid&#34;)
 *             .systemDiskCategory(&#34;cloud_efficiency&#34;)
 *             .vswitchId(defaultVsw.id())
 *             .build());
 * 
 *         var defaultXkpFRs = new AnycastEipAddress(&#34;defaultXkpFRs&#34;, AnycastEipAddressArgs.builder()        
 *             .serviceLocation(&#34;ChineseMainland&#34;)
 *             .build());
 * 
 *         var defaultVpc2 = new Network(&#34;defaultVpc2&#34;, NetworkArgs.builder()        
 *             .vpcName(String.format(&#34;%s6&#34;, name))
 *             .cidrBlock(&#34;192.168.0.0/16&#34;)
 *             .build());
 * 
 *         final var default2 = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableDiskCategory(&#34;cloud_efficiency&#34;)
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         final var default2GetImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex(&#34;^ubuntu_18.*64&#34;)
 *             .mostRecent(true)
 *             .owners(&#34;system&#34;)
 *             .build());
 * 
 *         final var default2GetInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(default2.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .cpuCoreCount(1)
 *             .memorySize(2)
 *             .build());
 * 
 *         var defaultdsVsw2 = new Switch(&#34;defaultdsVsw2&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultVpc2.id())
 *             .cidrBlock(&#34;192.168.0.0/24&#34;)
 *             .zoneId(default2.applyValue(getZonesResult -&gt; getZonesResult.zones()[1].id()))
 *             .build());
 * 
 *         var defaultuBsECI2 = new SecurityGroup(&#34;defaultuBsECI2&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(defaultVpc2.id())
 *             .build());
 * 
 *         var defaultEcs2 = new Instance(&#34;defaultEcs2&#34;, InstanceArgs.builder()        
 *             .imageId(default2GetImages.applyValue(getImagesResult -&gt; getImagesResult.images()[0].id()))
 *             .instanceType(default2GetInstanceTypes.applyValue(getInstanceTypesResult -&gt; getInstanceTypesResult.instanceTypes()[0].id()))
 *             .instanceName(name)
 *             .securityGroups(defaultuBsECI2.id())
 *             .availabilityZone(defaultdsVsw2.zoneId())
 *             .instanceChargeType(&#34;PostPaid&#34;)
 *             .systemDiskCategory(&#34;cloud_efficiency&#34;)
 *             .vswitchId(defaultdsVsw2.id())
 *             .build());
 * 
 *         var defaultEfYBJY = new AnycastEipAddressAttachment(&#34;defaultEfYBJY&#34;, AnycastEipAddressAttachmentArgs.builder()        
 *             .bindInstanceId(default9KDlN7.networkInterfaceId())
 *             .bindInstanceType(&#34;NetworkInterface&#34;)
 *             .bindInstanceRegionId(&#34;cn-beijing&#34;)
 *             .anycastId(defaultXkpFRs.id())
 *             .associationMode(&#34;Default&#34;)
 *             .build());
 * 
 *         var normal = new AnycastEipAddressAttachment(&#34;normal&#34;, AnycastEipAddressAttachmentArgs.builder()        
 *             .bindInstanceId(defaultEcs2.networkInterfaceId())
 *             .bindInstanceType(&#34;NetworkInterface&#34;)
 *             .bindInstanceRegionId(&#34;cn-hangzhou&#34;)
 *             .anycastId(defaultEfYBJY.anycastId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Eipanycast Anycast Eip Address Attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment example &lt;anycast_id&gt;:&lt;bind_instance_id&gt;:&lt;bind_instance_region_id&gt;:&lt;bind_instance_type&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment")
public class AnycastEipAddressAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Anycast EIP instance.
     * 
     */
    @Export(name="anycastId", refs={String.class}, tree="[0]")
    private Output<String> anycastId;

    /**
     * @return The ID of the Anycast EIP instance.
     * 
     */
    public Output<String> anycastId() {
        return this.anycastId;
    }
    /**
     * Binding mode, value:
     * - **Default**: The Default mode. The cloud resource instance to be bound is set as the Default origin.
     * - **Normal**: In Normal mode, the cloud resource instance to be bound is set to the common source station.
     * 
     */
    @Export(name="associationMode", refs={String.class}, tree="[0]")
    private Output<String> associationMode;

    /**
     * @return Binding mode, value:
     * - **Default**: The Default mode. The cloud resource instance to be bound is set as the Default origin.
     * - **Normal**: In Normal mode, the cloud resource instance to be bound is set to the common source station.
     * 
     */
    public Output<String> associationMode() {
        return this.associationMode;
    }
    /**
     * The ID of the cloud resource instance to be bound.
     * 
     */
    @Export(name="bindInstanceId", refs={String.class}, tree="[0]")
    private Output<String> bindInstanceId;

    /**
     * @return The ID of the cloud resource instance to be bound.
     * 
     */
    public Output<String> bindInstanceId() {
        return this.bindInstanceId;
    }
    /**
     * The region ID of the cloud resource instance to be bound.You can only bind cloud resource instances in some regions. You can call the describeanystserverregions operation to obtain the region ID of the cloud resource instances that can be bound.
     * 
     */
    @Export(name="bindInstanceRegionId", refs={String.class}, tree="[0]")
    private Output<String> bindInstanceRegionId;

    /**
     * @return The region ID of the cloud resource instance to be bound.You can only bind cloud resource instances in some regions. You can call the describeanystserverregions operation to obtain the region ID of the cloud resource instances that can be bound.
     * 
     */
    public Output<String> bindInstanceRegionId() {
        return this.bindInstanceRegionId;
    }
    /**
     * The type of the cloud resource instance to be bound. Value:
     * - **SlbInstance**: a private network SLB instance.
     * - **NetworkInterface**: ENI.
     * 
     */
    @Export(name="bindInstanceType", refs={String.class}, tree="[0]")
    private Output<String> bindInstanceType;

    /**
     * @return The type of the cloud resource instance to be bound. Value:
     * - **SlbInstance**: a private network SLB instance.
     * - **NetworkInterface**: ENI.
     * 
     */
    public Output<String> bindInstanceType() {
        return this.bindInstanceType;
    }
    /**
     * Binding time.Time is expressed according to ISO8601 standard and UTC time is used. The format is: &#39;YYYY-MM-DDThh:mm:ssZ&#39;.
     * 
     */
    @Export(name="bindTime", refs={String.class}, tree="[0]")
    private Output<String> bindTime;

    /**
     * @return Binding time.Time is expressed according to ISO8601 standard and UTC time is used. The format is: &#39;YYYY-MM-DDThh:mm:ssZ&#39;.
     * 
     */
    public Output<String> bindTime() {
        return this.bindTime;
    }
    /**
     * The access point information of the associated access area when the cloud resource instance is bound.If you are binding for the first time, this parameter does not need to be configured, and the system automatically associates all access areas. See `pop_locations` below.
     * 
     */
    @Export(name="popLocations", refs={List.class,AnycastEipAddressAttachmentPopLocation.class}, tree="[0,1]")
    private Output</* @Nullable */ List<AnycastEipAddressAttachmentPopLocation>> popLocations;

    /**
     * @return The access point information of the associated access area when the cloud resource instance is bound.If you are binding for the first time, this parameter does not need to be configured, and the system automatically associates all access areas. See `pop_locations` below.
     * 
     */
    public Output<Optional<List<AnycastEipAddressAttachmentPopLocation>>> popLocations() {
        return Codegen.optional(this.popLocations);
    }
    /**
     * The secondary private IP address of the elastic network card to be bound.This parameter takes effect only when **BindInstanceType** is set to **NetworkInterface. When you do not enter, this parameter is the primary private IP of the ENI by default.
     * 
     */
    @Export(name="privateIpAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> privateIpAddress;

    /**
     * @return The secondary private IP address of the elastic network card to be bound.This parameter takes effect only when **BindInstanceType** is set to **NetworkInterface. When you do not enter, this parameter is the primary private IP of the ENI by default.
     * 
     */
    public Output<Optional<String>> privateIpAddress() {
        return Codegen.optional(this.privateIpAddress);
    }
    /**
     * The status of the bound cloud resource instance. Value:BINDING: BINDING.Bound: Bound.UNBINDING: UNBINDING.DELETED: DELETED.MODIFYING: being modified.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the bound cloud resource instance. Value:BINDING: BINDING.Bound: Bound.UNBINDING: UNBINDING.DELETED: DELETED.MODIFYING: being modified.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AnycastEipAddressAttachment(String name) {
        this(name, AnycastEipAddressAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AnycastEipAddressAttachment(String name, AnycastEipAddressAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AnycastEipAddressAttachment(String name, AnycastEipAddressAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment", name, args == null ? AnycastEipAddressAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AnycastEipAddressAttachment(String name, Output<String> id, @Nullable AnycastEipAddressAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment", name, state, makeResourceOptions(options, id));
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
    public static AnycastEipAddressAttachment get(String name, Output<String> id, @Nullable AnycastEipAddressAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AnycastEipAddressAttachment(name, id, state, options);
    }
}
