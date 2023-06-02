// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.ImageSharePermissionArgs;
import com.pulumi.alicloud.ecs.inputs.ImageSharePermissionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manage image sharing permissions. You can share your custom image to other Alibaba Cloud users. The user can use the shared custom image to create ECS instances or replace the system disk of the instance.
 * 
 * &gt; **NOTE:** You can only share your own custom images to other Alibaba Cloud users.
 * 
 * &gt; **NOTE:** Each custom image can be shared with up to 50 Alibaba Cloud accounts. You can submit a ticket to share with more users.
 * 
 * &gt; **NOTE:** After creating an ECS instance using a shared image, once the custom image owner releases the image sharing relationship or deletes the custom image, the instance cannot initialize the system disk.
 * 
 * &gt; **NOTE:** Available in 1.68.0+.
 * 
 * ## Example Usage
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
 * import com.pulumi.alicloud.ecs.Image;
 * import com.pulumi.alicloud.ecs.ImageArgs;
 * import com.pulumi.alicloud.ecs.ImageSharePermission;
 * import com.pulumi.alicloud.ecs.ImageSharePermissionArgs;
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
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;Instance&#34;)
 *             .build());
 * 
 *         final var defaultInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .instanceTypeFamily(&#34;ecs.sn1ne&#34;)
 *             .build());
 * 
 *         final var defaultImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex(&#34;^ubuntu_[0-9]+_[0-9]+_x64*&#34;)
 *             .owners(&#34;system&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;terraform-example&#34;)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(&#34;terraform-example&#34;)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .availabilityZone(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .instanceName(&#34;terraform-example&#34;)
 *             .securityGroups(defaultSecurityGroup.id())
 *             .vswitchId(defaultSwitch.id())
 *             .instanceType(defaultInstanceTypes.applyValue(getInstanceTypesResult -&gt; getInstanceTypesResult.ids()[0]))
 *             .imageId(defaultImages.applyValue(getImagesResult -&gt; getImagesResult.ids()[0]))
 *             .internetMaxBandwidthOut(10)
 *             .build());
 * 
 *         var defaultImage = new Image(&#34;defaultImage&#34;, ImageArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .imageName(&#34;terraform-example&#34;)
 *             .description(&#34;terraform-example&#34;)
 *             .build());
 * 
 *         final var accountId = config.get(&#34;accountId&#34;).orElse(&#34;123456789&#34;);
 *         var defaultImageSharePermission = new ImageSharePermission(&#34;defaultImageSharePermission&#34;, ImageSharePermissionArgs.builder()        
 *             .imageId(defaultImage.id())
 *             .accountId(accountId)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Attributes Reference0
 * 
 *  The following attributes are exported:
 * 
 * * `id` - ID of the image. It formats as `&lt;image_id&gt;:&lt;account_id&gt;`
 * 
 * ## Import
 * 
 * image can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ecs/imageSharePermission:ImageSharePermission default m-uf66yg1q:123456789
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/imageSharePermission:ImageSharePermission")
public class ImageSharePermission extends com.pulumi.resources.CustomResource {
    /**
     * Alibaba Cloud Account ID. It is used to share images.
     * 
     */
    @Export(name="accountId", type=String.class, parameters={})
    private Output<String> accountId;

    /**
     * @return Alibaba Cloud Account ID. It is used to share images.
     * 
     */
    public Output<String> accountId() {
        return this.accountId;
    }
    /**
     * The source image ID.
     * 
     */
    @Export(name="imageId", type=String.class, parameters={})
    private Output<String> imageId;

    /**
     * @return The source image ID.
     * 
     */
    public Output<String> imageId() {
        return this.imageId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ImageSharePermission(String name) {
        this(name, ImageSharePermissionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ImageSharePermission(String name, ImageSharePermissionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ImageSharePermission(String name, ImageSharePermissionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/imageSharePermission:ImageSharePermission", name, args == null ? ImageSharePermissionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ImageSharePermission(String name, Output<String> id, @Nullable ImageSharePermissionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/imageSharePermission:ImageSharePermission", name, state, makeResourceOptions(options, id));
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
    public static ImageSharePermission get(String name, Output<String> id, @Nullable ImageSharePermissionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ImageSharePermission(name, id, state, options);
    }
}
