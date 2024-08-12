// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ga.AcceleratorSpareIpAttachmentArgs;
import com.pulumi.alicloud.ga.inputs.AcceleratorSpareIpAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Global Accelerator (GA) Accelerator Spare Ip Attachment resource.
 * 
 * For information about Global Accelerator (GA) Accelerator Spare Ip Attachment and how to use it, see [What is Accelerator Spare Ip Attachment](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createspareips).
 * 
 * &gt; **NOTE:** Available since v1.167.0.
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
 * import com.pulumi.alicloud.ga.AcceleratorSpareIpAttachment;
 * import com.pulumi.alicloud.ga.AcceleratorSpareIpAttachmentArgs;
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
 *             .spec("1")
 *             .acceleratorName("terraform-example")
 *             .autoUseCoupon(true)
 *             .description("terraform-example")
 *             .build());
 * 
 *         var defaultBandwidthPackage = new BandwidthPackage("defaultBandwidthPackage", BandwidthPackageArgs.builder()
 *             .bandwidth(100)
 *             .type("Basic")
 *             .bandwidthType("Basic")
 *             .paymentType("PayAsYouGo")
 *             .billingType("PayBy95")
 *             .ratio(30)
 *             .bandwidthPackageName("terraform-example")
 *             .autoPay(true)
 *             .autoUseCoupon(true)
 *             .build());
 * 
 *         var defaultBandwidthPackageAttachment = new BandwidthPackageAttachment("defaultBandwidthPackageAttachment", BandwidthPackageAttachmentArgs.builder()
 *             .acceleratorId(default_.id())
 *             .bandwidthPackageId(defaultBandwidthPackage.id())
 *             .build());
 * 
 *         var defaultAcceleratorSpareIpAttachment = new AcceleratorSpareIpAttachment("defaultAcceleratorSpareIpAttachment", AcceleratorSpareIpAttachmentArgs.builder()
 *             .acceleratorId(defaultBandwidthPackageAttachment.acceleratorId())
 *             .spareIp("127.0.0.1")
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
 * Global Accelerator (GA) Accelerator Spare Ip Attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ga/acceleratorSpareIpAttachment:AcceleratorSpareIpAttachment example &lt;accelerator_id&gt;:&lt;spare_ip&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ga/acceleratorSpareIpAttachment:AcceleratorSpareIpAttachment")
public class AcceleratorSpareIpAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the global acceleration instance.
     * 
     */
    @Export(name="acceleratorId", refs={String.class}, tree="[0]")
    private Output<String> acceleratorId;

    /**
     * @return The ID of the global acceleration instance.
     * 
     */
    public Output<String> acceleratorId() {
        return this.acceleratorId;
    }
    /**
     * The dry run.
     * 
     */
    @Export(name="dryRun", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return The dry run.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The standby IP address of CNAME. When the acceleration area is abnormal, the traffic is switched to the standby IP address.
     * 
     */
    @Export(name="spareIp", refs={String.class}, tree="[0]")
    private Output<String> spareIp;

    /**
     * @return The standby IP address of CNAME. When the acceleration area is abnormal, the traffic is switched to the standby IP address.
     * 
     */
    public Output<String> spareIp() {
        return this.spareIp;
    }
    /**
     * The status of the standby CNAME IP address.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the standby CNAME IP address.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AcceleratorSpareIpAttachment(java.lang.String name) {
        this(name, AcceleratorSpareIpAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AcceleratorSpareIpAttachment(java.lang.String name, AcceleratorSpareIpAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AcceleratorSpareIpAttachment(java.lang.String name, AcceleratorSpareIpAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/acceleratorSpareIpAttachment:AcceleratorSpareIpAttachment", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AcceleratorSpareIpAttachment(java.lang.String name, Output<java.lang.String> id, @Nullable AcceleratorSpareIpAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/acceleratorSpareIpAttachment:AcceleratorSpareIpAttachment", name, state, makeResourceOptions(options, id), false);
    }

    private static AcceleratorSpareIpAttachmentArgs makeArgs(AcceleratorSpareIpAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AcceleratorSpareIpAttachmentArgs.Empty : args;
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
    public static AcceleratorSpareIpAttachment get(java.lang.String name, Output<java.lang.String> id, @Nullable AcceleratorSpareIpAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AcceleratorSpareIpAttachment(name, id, state, options);
    }
}
