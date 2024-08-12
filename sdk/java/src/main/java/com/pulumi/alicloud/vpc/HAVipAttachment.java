// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.HAVipAttachmentArgs;
import com.pulumi.alicloud.vpc.inputs.HAVipAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.alicloud.ecs.inputs.GetImagesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.vpc.HAVip;
 * import com.pulumi.alicloud.vpc.HAVipArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.Instance;
 * import com.pulumi.alicloud.ecs.InstanceArgs;
 * import com.pulumi.alicloud.vpc.HAVipAttachment;
 * import com.pulumi.alicloud.vpc.HAVipAttachmentArgs;
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
 *         final var example = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(default_.zones()[0].id())
 *             .cpuCoreCount(1)
 *             .memorySize(2)
 *             .build());
 * 
 *         final var exampleGetImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex("^ubuntu_[0-9]+_[0-9]+_x64*")
 *             .owners("system")
 *             .build());
 * 
 *         var exampleNetwork = new Network("exampleNetwork", NetworkArgs.builder()
 *             .vpcName(name)
 *             .cidrBlock("10.4.0.0/16")
 *             .build());
 * 
 *         var exampleSwitch = new Switch("exampleSwitch", SwitchArgs.builder()
 *             .vswitchName(name)
 *             .cidrBlock("10.4.0.0/24")
 *             .vpcId(exampleNetwork.id())
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var exampleHAVip = new HAVip("exampleHAVip", HAVipArgs.builder()
 *             .vswitchId(exampleSwitch.id())
 *             .description(name)
 *             .build());
 * 
 *         var exampleSecurityGroup = new SecurityGroup("exampleSecurityGroup", SecurityGroupArgs.builder()
 *             .name(name)
 *             .description(name)
 *             .vpcId(exampleNetwork.id())
 *             .build());
 * 
 *         var exampleInstance = new Instance("exampleInstance", InstanceArgs.builder()
 *             .availabilityZone(default_.zones()[0].id())
 *             .vswitchId(exampleSwitch.id())
 *             .imageId(exampleGetImages.applyValue(getImagesResult -> getImagesResult.images()[0].id()))
 *             .instanceType(example.applyValue(getInstanceTypesResult -> getInstanceTypesResult.instanceTypes()[0].id()))
 *             .systemDiskCategory("cloud_efficiency")
 *             .internetChargeType("PayByTraffic")
 *             .internetMaxBandwidthOut(5)
 *             .securityGroups(exampleSecurityGroup.id())
 *             .instanceName(name)
 *             .userData("echo 'net.ipv4.ip_forward=1'>> /etc/sysctl.conf")
 *             .build());
 * 
 *         var exampleHAVipAttachment = new HAVipAttachment("exampleHAVipAttachment", HAVipAttachmentArgs.builder()
 *             .havipId(exampleHAVip.id())
 *             .instanceId(exampleInstance.id())
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
 * VPC Ha Vip Attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/hAVipAttachment:HAVipAttachment example &lt;ha_vip_id&gt;:&lt;instance_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/hAVipAttachment:HAVipAttachment")
public class HAVipAttachment extends com.pulumi.resources.CustomResource {
    /**
     * Whether to force the ECS instance or Eni instance bound to AVIP to be unbound. The value is:
     * - **True**: Force unbinding.
     * - **False** (default): unbinding is not forced.
     * &gt; **NOTE:**  If the value of this parameter is **False**, the Master instance bound to HaVip cannot be unbound.
     * 
     */
    @Export(name="force", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> force;

    /**
     * @return Whether to force the ECS instance or Eni instance bound to AVIP to be unbound. The value is:
     * - **True**: Force unbinding.
     * - **False** (default): unbinding is not forced.
     * &gt; **NOTE:**  If the value of this parameter is **False**, the Master instance bound to HaVip cannot be unbound.
     * 
     */
    public Output<Optional<Boolean>> force() {
        return Codegen.optional(this.force);
    }
    /**
     * The ID of the HaVip instance.
     * 
     */
    @Export(name="haVipId", refs={String.class}, tree="[0]")
    private Output<String> haVipId;

    /**
     * @return The ID of the HaVip instance.
     * 
     */
    public Output<String> haVipId() {
        return this.haVipId;
    }
    /**
     * . Field &#39;havip_id&#39; has been deprecated from provider version 1.211.0. New field &#39;ha_vip_id&#39; instead.
     * 
     * @deprecated
     * Field &#39;havip_id&#39; has been deprecated since provider version 1.211.0. New field &#39;ha_vip_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'havip_id' has been deprecated since provider version 1.211.0. New field 'ha_vip_id' instead. */
    @Export(name="havipId", refs={String.class}, tree="[0]")
    private Output<String> havipId;

    /**
     * @return . Field &#39;havip_id&#39; has been deprecated from provider version 1.211.0. New field &#39;ha_vip_id&#39; instead.
     * 
     */
    public Output<String> havipId() {
        return this.havipId;
    }
    /**
     * The ID of the ECS instance bound to the HaVip instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the ECS instance bound to the HaVip instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The type of the instance associated with the VIIP.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    @Export(name="instanceType", refs={String.class}, tree="[0]")
    private Output<String> instanceType;

    /**
     * @return The type of the instance associated with the VIIP.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
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
    public HAVipAttachment(java.lang.String name) {
        this(name, HAVipAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HAVipAttachment(java.lang.String name, HAVipAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HAVipAttachment(java.lang.String name, HAVipAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/hAVipAttachment:HAVipAttachment", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private HAVipAttachment(java.lang.String name, Output<java.lang.String> id, @Nullable HAVipAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/hAVipAttachment:HAVipAttachment", name, state, makeResourceOptions(options, id), false);
    }

    private static HAVipAttachmentArgs makeArgs(HAVipAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? HAVipAttachmentArgs.Empty : args;
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
    public static HAVipAttachment get(java.lang.String name, Output<java.lang.String> id, @Nullable HAVipAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HAVipAttachment(name, id, state, options);
    }
}
