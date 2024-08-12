// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.alb.LoadBalancerSecurityGroupAttachmentArgs;
import com.pulumi.alicloud.alb.inputs.LoadBalancerSecurityGroupAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a ALB Load Balancer Security Group Attachment resource.
 * 
 * Bind a security group to an application-type Server Load Balancer instance.
 * 
 * For information about ALB Load Balancer Security Group Attachment and how to use it, see [What is Load Balancer Security Group Attachment](https://www.alibabacloud.com/help/en/).
 * 
 * &gt; **NOTE:** Available since v1.226.0.
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
 * import com.pulumi.alicloud.alb.LoadBalancer;
 * import com.pulumi.alicloud.alb.LoadBalancerArgs;
 * import com.pulumi.alicloud.alb.inputs.LoadBalancerLoadBalancerBillingConfigArgs;
 * import com.pulumi.alicloud.alb.inputs.LoadBalancerZoneMappingArgs;
 * import com.pulumi.alicloud.alb.LoadBalancerSecurityGroupAttachment;
 * import com.pulumi.alicloud.alb.LoadBalancerSecurityGroupAttachmentArgs;
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
 *             .cidrBlock("192.168.0.0/16")
 *             .vpcName(name)
 *             .build());
 * 
 *         var createVsw1 = new Switch("createVsw1", SwitchArgs.builder()
 *             .vpcId(createVpc.id())
 *             .zoneId(default_.zones()[0].id())
 *             .cidrBlock("192.168.1.0/24")
 *             .vswitchName(name)
 *             .build());
 * 
 *         var createVsw2 = new Switch("createVsw2", SwitchArgs.builder()
 *             .vpcId(createVpc.id())
 *             .zoneId(default_.zones()[1].id())
 *             .cidrBlock("192.168.2.0/24")
 *             .vswitchName(name)
 *             .build());
 * 
 *         var createSecurityGroup = new SecurityGroup("createSecurityGroup", SecurityGroupArgs.builder()
 *             .name(name)
 *             .vpcId(createVpc.id())
 *             .build());
 * 
 *         var createAlb = new LoadBalancer("createAlb", LoadBalancerArgs.builder()
 *             .loadBalancerName(name)
 *             .loadBalancerEdition("Standard")
 *             .vpcId(createVpc.id())
 *             .loadBalancerBillingConfig(LoadBalancerLoadBalancerBillingConfigArgs.builder()
 *                 .payType("PayAsYouGo")
 *                 .build())
 *             .addressType("Intranet")
 *             .addressAllocatedMode("Fixed")
 *             .zoneMappings(            
 *                 LoadBalancerZoneMappingArgs.builder()
 *                     .vswitchId(createVsw2.id())
 *                     .zoneId(createVsw2.zoneId())
 *                     .build(),
 *                 LoadBalancerZoneMappingArgs.builder()
 *                     .vswitchId(createVsw1.id())
 *                     .zoneId(createVsw1.zoneId())
 *                     .build())
 *             .build());
 * 
 *         var defaultLoadBalancerSecurityGroupAttachment = new LoadBalancerSecurityGroupAttachment("defaultLoadBalancerSecurityGroupAttachment", LoadBalancerSecurityGroupAttachmentArgs.builder()
 *             .securityGroupId(createSecurityGroup.id())
 *             .loadBalancerId(createAlb.id())
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
 * ALB Load Balancer Security Group Attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:alb/loadBalancerSecurityGroupAttachment:LoadBalancerSecurityGroupAttachment example &lt;load_balancer_id&gt;:&lt;security_group_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:alb/loadBalancerSecurityGroupAttachment:LoadBalancerSecurityGroupAttachment")
public class LoadBalancerSecurityGroupAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the load balancing instance.
     * 
     */
    @Export(name="loadBalancerId", refs={String.class}, tree="[0]")
    private Output<String> loadBalancerId;

    /**
     * @return The ID of the load balancing instance.
     * 
     */
    public Output<String> loadBalancerId() {
        return this.loadBalancerId;
    }
    /**
     * Security group ID collection.
     * 
     */
    @Export(name="securityGroupId", refs={String.class}, tree="[0]")
    private Output<String> securityGroupId;

    /**
     * @return Security group ID collection.
     * 
     */
    public Output<String> securityGroupId() {
        return this.securityGroupId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LoadBalancerSecurityGroupAttachment(java.lang.String name) {
        this(name, LoadBalancerSecurityGroupAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LoadBalancerSecurityGroupAttachment(java.lang.String name, LoadBalancerSecurityGroupAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LoadBalancerSecurityGroupAttachment(java.lang.String name, LoadBalancerSecurityGroupAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alb/loadBalancerSecurityGroupAttachment:LoadBalancerSecurityGroupAttachment", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private LoadBalancerSecurityGroupAttachment(java.lang.String name, Output<java.lang.String> id, @Nullable LoadBalancerSecurityGroupAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alb/loadBalancerSecurityGroupAttachment:LoadBalancerSecurityGroupAttachment", name, state, makeResourceOptions(options, id), false);
    }

    private static LoadBalancerSecurityGroupAttachmentArgs makeArgs(LoadBalancerSecurityGroupAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? LoadBalancerSecurityGroupAttachmentArgs.Empty : args;
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
    public static LoadBalancerSecurityGroupAttachment get(java.lang.String name, Output<java.lang.String> id, @Nullable LoadBalancerSecurityGroupAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LoadBalancerSecurityGroupAttachment(name, id, state, options);
    }
}
