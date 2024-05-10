// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.TrafficMirrorSessionArgs;
import com.pulumi.alicloud.vpc.inputs.TrafficMirrorSessionState;
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
 * Provides a VPC Traffic Mirror Session resource. Traffic mirroring session.
 * 
 * For information about VPC Traffic Mirror Session and how to use it, see [What is Traffic Mirror Session](https://www.alibabacloud.com/help/en/doc-detail/261364.htm).
 * 
 * &gt; **NOTE:** Available since v1.142.0.
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
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecs.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.inputs.GetImagesArgs;
 * import com.pulumi.alicloud.ecs.Instance;
 * import com.pulumi.alicloud.ecs.InstanceArgs;
 * import com.pulumi.alicloud.ecs.EcsNetworkInterface;
 * import com.pulumi.alicloud.ecs.EcsNetworkInterfaceArgs;
 * import com.pulumi.alicloud.ecs.EcsNetworkInterfaceAttachment;
 * import com.pulumi.alicloud.ecs.EcsNetworkInterfaceAttachmentArgs;
 * import com.pulumi.alicloud.vpc.TrafficMirrorFilter;
 * import com.pulumi.alicloud.vpc.TrafficMirrorFilterArgs;
 * import com.pulumi.alicloud.vpc.TrafficMirrorSession;
 * import com.pulumi.alicloud.vpc.TrafficMirrorSessionArgs;
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
 *         final var name = config.get("name").orElse("tf-example");
 *         final var default = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .instanceTypeFamily("ecs.g7")
 *             .build());
 * 
 *         final var defaultGetZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation("Instance")
 *             .availableInstanceType(default_.instanceTypes()[0].id())
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
 *             .zoneId(defaultGetZones.applyValue(getZonesResult -> getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup("defaultSecurityGroup", SecurityGroupArgs.builder()        
 *             .name(name)
 *             .description(name)
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         final var defaultGetImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex("^ubuntu_[0-9]+_[0-9]+_x64*")
 *             .mostRecent(true)
 *             .owners("system")
 *             .build());
 * 
 *         for (var i = 0; i < 2; i++) {
 *             new Instance("defaultInstance-" + i, InstanceArgs.builder()            
 *                 .availabilityZone(defaultGetZones.applyValue(getZonesResult -> getZonesResult.zones()[0].id()))
 *                 .instanceName(name)
 *                 .hostName(name)
 *                 .imageId(defaultGetImages.applyValue(getImagesResult -> getImagesResult.images()[0].id()))
 *                 .instanceType(default_.instanceTypes()[0].id())
 *                 .securityGroups(defaultSecurityGroup.id())
 *                 .vswitchId(defaultSwitch.id())
 *                 .systemDiskCategory("cloud_essd")
 *                 .build());
 * 
 *         
 * }
 *         for (var i = 0; i < 2; i++) {
 *             new EcsNetworkInterface("defaultEcsNetworkInterface-" + i, EcsNetworkInterfaceArgs.builder()            
 *                 .networkInterfaceName(name)
 *                 .vswitchId(defaultSwitch.id())
 *                 .securityGroupIds(defaultSecurityGroup.id())
 *                 .build());
 * 
 *         
 * }
 *         for (var i = 0; i < 2; i++) {
 *             new EcsNetworkInterfaceAttachment("defaultEcsNetworkInterfaceAttachment-" + i, EcsNetworkInterfaceAttachmentArgs.builder()            
 *                 .instanceId(defaultInstance[range.value()].id())
 *                 .networkInterfaceId(defaultEcsNetworkInterface[range.value()].id())
 *                 .build());
 * 
 *         
 * }
 *         var defaultTrafficMirrorFilter = new TrafficMirrorFilter("defaultTrafficMirrorFilter", TrafficMirrorFilterArgs.builder()        
 *             .trafficMirrorFilterName(name)
 *             .trafficMirrorFilterDescription(name)
 *             .build());
 * 
 *         var defaultTrafficMirrorSession = new TrafficMirrorSession("defaultTrafficMirrorSession", TrafficMirrorSessionArgs.builder()        
 *             .priority(1)
 *             .virtualNetworkId(10)
 *             .trafficMirrorSessionDescription(name)
 *             .trafficMirrorSessionName(name)
 *             .trafficMirrorTargetId(defaultEcsNetworkInterfaceAttachment[0].networkInterfaceId())
 *             .trafficMirrorSourceIds(defaultEcsNetworkInterfaceAttachment[1].networkInterfaceId())
 *             .trafficMirrorFilterId(defaultTrafficMirrorFilter.id())
 *             .trafficMirrorTargetType("NetworkInterface")
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
 * VPC Traffic Mirror Session can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/trafficMirrorSession:TrafficMirrorSession example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/trafficMirrorSession:TrafficMirrorSession")
public class TrafficMirrorSession extends com.pulumi.resources.CustomResource {
    /**
     * Whether to PreCheck only this request, value:
     * - **true**: sends a check request and does not create a mirror session. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
     * - **false** (default): Sends a normal request and directly creates a mirror session after checking.
     * 
     */
    @Export(name="dryRun", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Whether to PreCheck only this request, value:
     * - **true**: sends a check request and does not create a mirror session. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
     * - **false** (default): Sends a normal request and directly creates a mirror session after checking.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * Specifies whether to enable traffic mirror sessions. default to `false`.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Specifies whether to enable traffic mirror sessions. default to `false`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * Maximum Transmission Unit (MTU).
     * 
     */
    @Export(name="packetLength", refs={Integer.class}, tree="[0]")
    private Output<Integer> packetLength;

    /**
     * @return Maximum Transmission Unit (MTU).
     * 
     */
    public Output<Integer> packetLength() {
        return this.packetLength;
    }
    /**
     * The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
     * 
     */
    @Export(name="priority", refs={Integer.class}, tree="[0]")
    private Output<Integer> priority;

    /**
     * @return The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
     * 
     */
    public Output<Integer> priority() {
        return this.priority;
    }
    /**
     * The ID of the resource group.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
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
     * The tags of this resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return The tags of this resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The ID of the filter.
     * 
     */
    @Export(name="trafficMirrorFilterId", refs={String.class}, tree="[0]")
    private Output<String> trafficMirrorFilterId;

    /**
     * @return The ID of the filter.
     * 
     */
    public Output<String> trafficMirrorFilterId() {
        return this.trafficMirrorFilterId;
    }
    /**
     * The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
     * 
     */
    @Export(name="trafficMirrorSessionDescription", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trafficMirrorSessionDescription;

    /**
     * @return The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
     * 
     */
    public Output<Optional<String>> trafficMirrorSessionDescription() {
        return Codegen.optional(this.trafficMirrorSessionDescription);
    }
    /**
     * The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
     * 
     */
    @Export(name="trafficMirrorSessionName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trafficMirrorSessionName;

    /**
     * @return The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
     * 
     */
    public Output<Optional<String>> trafficMirrorSessionName() {
        return Codegen.optional(this.trafficMirrorSessionName);
    }
    /**
     * The ID of the image source instance. Currently, the Eni is supported as the image source. The default value of N is 1, that is, only one mirror source can be added to a mirror session.
     * 
     */
    @Export(name="trafficMirrorSourceIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> trafficMirrorSourceIds;

    /**
     * @return The ID of the image source instance. Currently, the Eni is supported as the image source. The default value of N is 1, that is, only one mirror source can be added to a mirror session.
     * 
     */
    public Output<List<String>> trafficMirrorSourceIds() {
        return this.trafficMirrorSourceIds;
    }
    /**
     * The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
     * 
     */
    @Export(name="trafficMirrorTargetId", refs={String.class}, tree="[0]")
    private Output<String> trafficMirrorTargetId;

    /**
     * @return The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
     * 
     */
    public Output<String> trafficMirrorTargetId() {
        return this.trafficMirrorTargetId;
    }
    /**
     * The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance.
     * 
     */
    @Export(name="trafficMirrorTargetType", refs={String.class}, tree="[0]")
    private Output<String> trafficMirrorTargetType;

    /**
     * @return The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance.
     * 
     */
    public Output<String> trafficMirrorTargetType() {
        return this.trafficMirrorTargetType;
    }
    /**
     * The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
     * 
     */
    @Export(name="virtualNetworkId", refs={Integer.class}, tree="[0]")
    private Output<Integer> virtualNetworkId;

    /**
     * @return The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
     * 
     */
    public Output<Integer> virtualNetworkId() {
        return this.virtualNetworkId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TrafficMirrorSession(String name) {
        this(name, TrafficMirrorSessionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TrafficMirrorSession(String name, TrafficMirrorSessionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TrafficMirrorSession(String name, TrafficMirrorSessionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/trafficMirrorSession:TrafficMirrorSession", name, args == null ? TrafficMirrorSessionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TrafficMirrorSession(String name, Output<String> id, @Nullable TrafficMirrorSessionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/trafficMirrorSession:TrafficMirrorSession", name, state, makeResourceOptions(options, id));
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
    public static TrafficMirrorSession get(String name, Output<String> id, @Nullable TrafficMirrorSessionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TrafficMirrorSession(name, id, state, options);
    }
}
