// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ga.BandwidthPackageAttachmentArgs;
import com.pulumi.alicloud.ga.inputs.BandwidthPackageAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a Global Accelerator (GA) Bandwidth Package Attachment resource.
 * 
 * For information about Global Accelerator (GA) Bandwidth Package Attachment and how to use it, see [What is Bandwidth Package Attachment](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-bandwidthpackageaddaccelerator).
 * 
 * &gt; **NOTE:** Available since v1.113.0.
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
 * import com.pulumi.alicloud.ga.Accelerator;
 * import com.pulumi.alicloud.ga.AcceleratorArgs;
 * import com.pulumi.alicloud.ga.BandwidthPackage;
 * import com.pulumi.alicloud.ga.BandwidthPackageArgs;
 * import com.pulumi.alicloud.ga.BandwidthPackageAttachment;
 * import com.pulumi.alicloud.ga.BandwidthPackageAttachmentArgs;
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
 *         var default_ = new Accelerator("default", AcceleratorArgs.builder()
 *             .duration(1)
 *             .autoUseCoupon(true)
 *             .spec("1")
 *             .build());
 * 
 *         var defaultBandwidthPackage = new BandwidthPackage("defaultBandwidthPackage", BandwidthPackageArgs.builder()
 *             .bandwidth(100)
 *             .type("Basic")
 *             .bandwidthType("Basic")
 *             .paymentType("PayAsYouGo")
 *             .billingType("PayBy95")
 *             .ratio(30)
 *             .build());
 * 
 *         var defaultBandwidthPackageAttachment = new BandwidthPackageAttachment("defaultBandwidthPackageAttachment", BandwidthPackageAttachmentArgs.builder()
 *             .acceleratorId(default_.id())
 *             .bandwidthPackageId(defaultBandwidthPackage.id())
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
 * Ga Bandwidth Package Attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ga/bandwidthPackageAttachment:BandwidthPackageAttachment example &lt;accelerator_id&gt;:&lt;bandwidth_package_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ga/bandwidthPackageAttachment:BandwidthPackageAttachment")
public class BandwidthPackageAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Global Accelerator instance.
     * 
     */
    @Export(name="acceleratorId", refs={String.class}, tree="[0]")
    private Output<String> acceleratorId;

    /**
     * @return The ID of the Global Accelerator instance.
     * 
     */
    public Output<String> acceleratorId() {
        return this.acceleratorId;
    }
    /**
     * Accelerators bound with current Bandwidth Package.
     * 
     */
    @Export(name="accelerators", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> accelerators;

    /**
     * @return Accelerators bound with current Bandwidth Package.
     * 
     */
    public Output<List<String>> accelerators() {
        return this.accelerators;
    }
    /**
     * The ID of the Bandwidth Package. **NOTE:** From version 1.192.0, `bandwidth_package_id` can be modified.
     * 
     */
    @Export(name="bandwidthPackageId", refs={String.class}, tree="[0]")
    private Output<String> bandwidthPackageId;

    /**
     * @return The ID of the Bandwidth Package. **NOTE:** From version 1.192.0, `bandwidth_package_id` can be modified.
     * 
     */
    public Output<String> bandwidthPackageId() {
        return this.bandwidthPackageId;
    }
    /**
     * State of Bandwidth Package.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return State of Bandwidth Package.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BandwidthPackageAttachment(String name) {
        this(name, BandwidthPackageAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BandwidthPackageAttachment(String name, BandwidthPackageAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BandwidthPackageAttachment(String name, BandwidthPackageAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/bandwidthPackageAttachment:BandwidthPackageAttachment", name, args == null ? BandwidthPackageAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BandwidthPackageAttachment(String name, Output<String> id, @Nullable BandwidthPackageAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/bandwidthPackageAttachment:BandwidthPackageAttachment", name, state, makeResourceOptions(options, id));
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
    public static BandwidthPackageAttachment get(String name, Output<String> id, @Nullable BandwidthPackageAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BandwidthPackageAttachment(name, id, state, options);
    }
}
