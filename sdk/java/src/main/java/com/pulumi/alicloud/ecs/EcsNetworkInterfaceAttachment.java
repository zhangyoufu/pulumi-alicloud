// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.EcsNetworkInterfaceAttachmentArgs;
import com.pulumi.alicloud.ecs.inputs.EcsNetworkInterfaceAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ECS Network Interface Attachment resource.
 * 
 * For information about ECS Network Interface Attachment and how to use it, see [What is Network Interface Attachment](https://www.alibabacloud.com/help/en/doc-detail/58515.htm).
 * 
 * &gt; **NOTE:** Available since v1.123.1.
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
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecs.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.inputs.GetImagesArgs;
 * import com.pulumi.alicloud.ecs.Instance;
 * import com.pulumi.alicloud.ecs.InstanceArgs;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.ecs.EcsNetworkInterface;
 * import com.pulumi.alicloud.ecs.EcsNetworkInterfaceArgs;
 * import com.pulumi.alicloud.ecs.EcsNetworkInterfaceAttachment;
 * import com.pulumi.alicloud.ecs.EcsNetworkInterfaceAttachmentArgs;
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
 *             .availableResourceCreation("Instance")
 *             .build());
 * 
 *         final var defaultGetInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(default_.zones()[0].id())
 *             .eniAmount(3)
 *             .build());
 * 
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()
 *             .vpcName(name)
 *             .cidrBlock("192.168.0.0/24")
 *             .build());
 * 
 *         var defaultSwitch = new Switch("defaultSwitch", SwitchArgs.builder()
 *             .vswitchName(name)
 *             .cidrBlock("192.168.0.0/24")
 *             .zoneId(default_.zones()[0].id())
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup("defaultSecurityGroup", SecurityGroupArgs.builder()
 *             .name(name)
 *             .description("New security group")
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         final var defaultGetImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex("^ubuntu_[0-9]+_[0-9]+_x64*")
 *             .mostRecent(true)
 *             .owners("system")
 *             .build());
 * 
 *         var defaultInstance = new Instance("defaultInstance", InstanceArgs.builder()
 *             .availabilityZone(default_.zones()[0].id())
 *             .instanceName(name)
 *             .hostName("tf-example")
 *             .imageId(defaultGetImages.applyValue(getImagesResult -> getImagesResult.images()[0].id()))
 *             .instanceType(defaultGetInstanceTypes.applyValue(getInstanceTypesResult -> getInstanceTypesResult.instanceTypes()[0].id()))
 *             .securityGroups(defaultSecurityGroup.id())
 *             .vswitchId(defaultSwitch.id())
 *             .build());
 * 
 *         final var defaultGetResourceGroups = ResourcemanagerFunctions.getResourceGroups(GetResourceGroupsArgs.builder()
 *             .status("OK")
 *             .build());
 * 
 *         var defaultEcsNetworkInterface = new EcsNetworkInterface("defaultEcsNetworkInterface", EcsNetworkInterfaceArgs.builder()
 *             .networkInterfaceName(name)
 *             .vswitchId(defaultSwitch.id())
 *             .securityGroupIds(defaultSecurityGroup.id())
 *             .description("Basic example")
 *             .primaryIpAddress("192.168.0.2")
 *             .tags(Map.ofEntries(
 *                 Map.entry("Created", "TF"),
 *                 Map.entry("For", "example")
 *             ))
 *             .resourceGroupId(defaultGetResourceGroups.applyValue(getResourceGroupsResult -> getResourceGroupsResult.ids()[0]))
 *             .build());
 * 
 *         var defaultEcsNetworkInterfaceAttachment = new EcsNetworkInterfaceAttachment("defaultEcsNetworkInterfaceAttachment", EcsNetworkInterfaceAttachmentArgs.builder()
 *             .networkInterfaceId(defaultEcsNetworkInterface.id())
 *             .instanceId(defaultInstance.id())
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
 * ECS Network Interface Attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ecs/ecsNetworkInterfaceAttachment:EcsNetworkInterfaceAttachment example &lt;network_interface_id&gt;:&lt;instance_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/ecsNetworkInterfaceAttachment:EcsNetworkInterfaceAttachment")
public class EcsNetworkInterfaceAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the ECS instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the ECS instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The index of the network card.
     * 
     */
    @Export(name="networkCardIndex", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> networkCardIndex;

    /**
     * @return The index of the network card.
     * 
     */
    public Output<Optional<Integer>> networkCardIndex() {
        return Codegen.optional(this.networkCardIndex);
    }
    /**
     * The ID of the network interface.
     * 
     */
    @Export(name="networkInterfaceId", refs={String.class}, tree="[0]")
    private Output<String> networkInterfaceId;

    /**
     * @return The ID of the network interface.
     * 
     */
    public Output<String> networkInterfaceId() {
        return this.networkInterfaceId;
    }
    /**
     * The ID of the trunk network instance.
     * 
     */
    @Export(name="trunkNetworkInstanceId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trunkNetworkInstanceId;

    /**
     * @return The ID of the trunk network instance.
     * 
     */
    public Output<Optional<String>> trunkNetworkInstanceId() {
        return Codegen.optional(this.trunkNetworkInstanceId);
    }
    /**
     * The wait for network configuration ready.
     * 
     */
    @Export(name="waitForNetworkConfigurationReady", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> waitForNetworkConfigurationReady;

    /**
     * @return The wait for network configuration ready.
     * 
     */
    public Output<Optional<Boolean>> waitForNetworkConfigurationReady() {
        return Codegen.optional(this.waitForNetworkConfigurationReady);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EcsNetworkInterfaceAttachment(java.lang.String name) {
        this(name, EcsNetworkInterfaceAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EcsNetworkInterfaceAttachment(java.lang.String name, EcsNetworkInterfaceAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EcsNetworkInterfaceAttachment(java.lang.String name, EcsNetworkInterfaceAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/ecsNetworkInterfaceAttachment:EcsNetworkInterfaceAttachment", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private EcsNetworkInterfaceAttachment(java.lang.String name, Output<java.lang.String> id, @Nullable EcsNetworkInterfaceAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/ecsNetworkInterfaceAttachment:EcsNetworkInterfaceAttachment", name, state, makeResourceOptions(options, id), false);
    }

    private static EcsNetworkInterfaceAttachmentArgs makeArgs(EcsNetworkInterfaceAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EcsNetworkInterfaceAttachmentArgs.Empty : args;
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
    public static EcsNetworkInterfaceAttachment get(java.lang.String name, Output<java.lang.String> id, @Nullable EcsNetworkInterfaceAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EcsNetworkInterfaceAttachment(name, id, state, options);
    }
}
