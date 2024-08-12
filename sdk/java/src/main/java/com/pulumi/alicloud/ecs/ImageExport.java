// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.ImageExportArgs;
import com.pulumi.alicloud.ecs.inputs.ImageExportState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Export a custom image to the OSS bucket in the same region as the custom image.
 * 
 * &gt; **NOTE:** If you create an ECS instance using a mirror image and create a system disk snapshot again, exporting a custom image created from the system disk snapshot is not supported.
 * 
 * &gt; **NOTE:** Support for exporting custom images that include data disk snapshot information in the image. The number of data disks cannot exceed 4 and the maximum capacity of a single data disk cannot exceed 500 GiB.
 * 
 * &gt; **NOTE:** Before exporting the image, you must authorize the cloud server ECS official service account to write OSS permissions through RAM.
 * 
 * &gt; **NOTE:** Available since v1.68.0+.
 * 
 * ## Example Usage
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
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.Instance;
 * import com.pulumi.alicloud.ecs.InstanceArgs;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.ecs.Image;
 * import com.pulumi.alicloud.ecs.ImageArgs;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.ecs.ImageExport;
 * import com.pulumi.alicloud.ecs.ImageExportArgs;
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
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation("Instance")
 *             .build());
 * 
 *         final var defaultGetInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .instanceTypeFamily("ecs.sn1ne")
 *             .build());
 * 
 *         final var defaultGetImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex("^ubuntu_[0-9]+_[0-9]+_x64*")
 *             .owners("system")
 *             .build());
 * 
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()
 *             .vpcName("terraform-example")
 *             .cidrBlock("172.17.3.0/24")
 *             .build());
 * 
 *         var defaultSwitch = new Switch("defaultSwitch", SwitchArgs.builder()
 *             .vswitchName("terraform-example")
 *             .cidrBlock("172.17.3.0/24")
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup("defaultSecurityGroup", SecurityGroupArgs.builder()
 *             .name("terraform-example")
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultInstance = new Instance("defaultInstance", InstanceArgs.builder()
 *             .availabilityZone(default_.zones()[0].id())
 *             .instanceName("terraform-example")
 *             .securityGroups(defaultSecurityGroup.id())
 *             .vswitchId(defaultSwitch.id())
 *             .instanceType(defaultGetInstanceTypes.applyValue(getInstanceTypesResult -> getInstanceTypesResult.ids()[0]))
 *             .imageId(defaultGetImages.applyValue(getImagesResult -> getImagesResult.ids()[0]))
 *             .internetMaxBandwidthOut(10)
 *             .build());
 * 
 *         var defaultInteger = new Integer("defaultInteger", IntegerArgs.builder()
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         var defaultImage = new Image("defaultImage", ImageArgs.builder()
 *             .instanceId(defaultInstance.id())
 *             .imageName(String.format("terraform-example-%s", defaultInteger.result()))
 *             .description("terraform-example")
 *             .build());
 * 
 *         var defaultBucket = new Bucket("defaultBucket", BucketArgs.builder()
 *             .bucket(String.format("example-value-%s", defaultInteger.result()))
 *             .build());
 * 
 *         var defaultImageExport = new ImageExport("defaultImageExport", ImageExportArgs.builder()
 *             .imageId(defaultImage.id())
 *             .ossBucket(defaultBucket.id())
 *             .ossPrefix("ecsExport")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="alicloud:ecs/imageExport:ImageExport")
public class ImageExport extends com.pulumi.resources.CustomResource {
    /**
     * The source image ID.
     * 
     */
    @Export(name="imageId", refs={String.class}, tree="[0]")
    private Output<String> imageId;

    /**
     * @return The source image ID.
     * 
     */
    public Output<String> imageId() {
        return this.imageId;
    }
    /**
     * Save the exported OSS bucket.
     * 
     */
    @Export(name="ossBucket", refs={String.class}, tree="[0]")
    private Output<String> ossBucket;

    /**
     * @return Save the exported OSS bucket.
     * 
     */
    public Output<String> ossBucket() {
        return this.ossBucket;
    }
    /**
     * The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
     * 
     */
    @Export(name="ossPrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ossPrefix;

    /**
     * @return The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
     * 
     */
    public Output<Optional<String>> ossPrefix() {
        return Codegen.optional(this.ossPrefix);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ImageExport(java.lang.String name) {
        this(name, ImageExportArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ImageExport(java.lang.String name, ImageExportArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ImageExport(java.lang.String name, ImageExportArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/imageExport:ImageExport", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ImageExport(java.lang.String name, Output<java.lang.String> id, @Nullable ImageExportState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/imageExport:ImageExport", name, state, makeResourceOptions(options, id), false);
    }

    private static ImageExportArgs makeArgs(ImageExportArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ImageExportArgs.Empty : args;
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
    public static ImageExport get(java.lang.String name, Output<java.lang.String> id, @Nullable ImageExportState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ImageExport(name, id, state, options);
    }
}
