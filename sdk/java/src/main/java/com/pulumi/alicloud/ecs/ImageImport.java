// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.ImageImportArgs;
import com.pulumi.alicloud.ecs.inputs.ImageImportState;
import com.pulumi.alicloud.ecs.outputs.ImageImportDiskDeviceMapping;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Import a copy of your local on-premise file to ECS, and appear as a custom replacement in the corresponding domain.
 * 
 * &gt; **NOTE:** You must upload the image file to the object storage OSS in advance.
 * 
 * &gt; **NOTE:** The region where the image is imported must be the same region as the OSS bucket where the image file is uploaded.
 * 
 * &gt; **NOTE:** Available in 1.69.0+.
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
 * import com.pulumi.alicloud.ecs.ImageImport;
 * import com.pulumi.alicloud.ecs.ImageImportArgs;
 * import com.pulumi.alicloud.ecs.inputs.ImageImportDiskDeviceMappingArgs;
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
 *         var this_ = new ImageImport(&#34;this&#34;, ImageImportArgs.builder()        
 *             .description(&#34;test import image&#34;)
 *             .architecture(&#34;x86_64&#34;)
 *             .imageName(&#34;test-import-image&#34;)
 *             .licenseType(&#34;Auto&#34;)
 *             .platform(&#34;Ubuntu&#34;)
 *             .osType(&#34;linux&#34;)
 *             .diskDeviceMappings(ImageImportDiskDeviceMappingArgs.builder()
 *                 .diskImageSize(5)
 *                 .ossBucket(&#34;testimportimage&#34;)
 *                 .ossObject(&#34;root.img&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Attributes Reference0
 * 
 *  The following attributes are exported:
 * 
 * * `id` - ID of the image.
 * 
 * ## Import
 * 
 * image can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ecs/imageImport:ImageImport default m-uf66871ape***yg1q***
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/imageImport:ImageImport")
public class ImageImport extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x86_64`.
     * 
     */
    @Export(name="architecture", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> architecture;

    /**
     * @return Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x86_64`.
     * 
     */
    public Output<Optional<String>> architecture() {
        return Codegen.optional(this.architecture);
    }
    /**
     * Description of the image. The length is 2 to 256 English or Chinese characters, and cannot begin with http: // and https: //.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the image. The length is 2 to 256 English or Chinese characters, and cannot begin with http: // and https: //.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Description of the system with disks and snapshots under the image.
     * 
     */
    @Export(name="diskDeviceMappings", refs={List.class,ImageImportDiskDeviceMapping.class}, tree="[0,1]")
    private Output<List<ImageImportDiskDeviceMapping>> diskDeviceMappings;

    /**
     * @return Description of the system with disks and snapshots under the image.
     * 
     */
    public Output<List<ImageImportDiskDeviceMapping>> diskDeviceMappings() {
        return this.diskDeviceMappings;
    }
    /**
     * The image name. The length is 2 ~ 128 English or Chinese characters. Must start with a english letter or Chinese, and cannot start with http: // and https: //. Can contain numbers, colons (:), underscores (_), or hyphens (-).
     * 
     */
    @Export(name="imageName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> imageName;

    /**
     * @return The image name. The length is 2 ~ 128 English or Chinese characters. Must start with a english letter or Chinese, and cannot start with http: // and https: //. Can contain numbers, colons (:), underscores (_), or hyphens (-).
     * 
     */
    public Output<Optional<String>> imageName() {
        return Codegen.optional(this.imageName);
    }
    /**
     * The type of the license used to activate the operating system after the image is imported. Default value: `Auto`. Valid values: `Auto`,`Aliyun`,`BYOL`.
     * 
     */
    @Export(name="licenseType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> licenseType;

    /**
     * @return The type of the license used to activate the operating system after the image is imported. Default value: `Auto`. Valid values: `Auto`,`Aliyun`,`BYOL`.
     * 
     */
    public Output<Optional<String>> licenseType() {
        return Codegen.optional(this.licenseType);
    }
    /**
     * Operating system platform type. Valid values: `windows`, Default is `linux`.
     * 
     */
    @Export(name="osType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> osType;

    /**
     * @return Operating system platform type. Valid values: `windows`, Default is `linux`.
     * 
     */
    public Output<Optional<String>> osType() {
        return Codegen.optional(this.osType);
    }
    /**
     * The operating system distribution. Default value: Others Linux.
     * More valid values refer to [ImportImage OpenAPI](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/importimage).
     * **NOTE**: It&#39;s default value is Ubuntu before version 1.197.0.
     * 
     */
    @Export(name="platform", refs={String.class}, tree="[0]")
    private Output<String> platform;

    /**
     * @return The operating system distribution. Default value: Others Linux.
     * More valid values refer to [ImportImage OpenAPI](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/importimage).
     * **NOTE**: It&#39;s default value is Ubuntu before version 1.197.0.
     * 
     */
    public Output<String> platform() {
        return this.platform;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ImageImport(String name) {
        this(name, ImageImportArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ImageImport(String name, ImageImportArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ImageImport(String name, ImageImportArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/imageImport:ImageImport", name, args == null ? ImageImportArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ImageImport(String name, Output<String> id, @Nullable ImageImportState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/imageImport:ImageImport", name, state, makeResourceOptions(options, id));
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
    public static ImageImport get(String name, Output<String> id, @Nullable ImageImportState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ImageImport(name, id, state, options);
    }
}
