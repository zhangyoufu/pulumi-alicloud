// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cen.BandwidthPackageAttachmentArgs;
import com.pulumi.alicloud.cen.inputs.BandwidthPackageAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a CEN bandwidth package attachment resource. The resource can be used to bind a bandwidth package to a specified CEN instance.
 * 
 * &gt; **NOTE:** Available since v1.18.0.
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
 * import com.pulumi.alicloud.cen.Instance;
 * import com.pulumi.alicloud.cen.InstanceArgs;
 * import com.pulumi.alicloud.cen.BandwidthPackage;
 * import com.pulumi.alicloud.cen.BandwidthPackageArgs;
 * import com.pulumi.alicloud.cen.BandwidthPackageAttachment;
 * import com.pulumi.alicloud.cen.BandwidthPackageAttachmentArgs;
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
 *         var example = new Instance(&#34;example&#34;, InstanceArgs.builder()        
 *             .cenInstanceName(&#34;tf_example&#34;)
 *             .description(&#34;an example for cen&#34;)
 *             .build());
 * 
 *         var exampleBandwidthPackage = new BandwidthPackage(&#34;exampleBandwidthPackage&#34;, BandwidthPackageArgs.builder()        
 *             .bandwidth(5)
 *             .cenBandwidthPackageName(&#34;tf_example&#34;)
 *             .geographicRegionAId(&#34;China&#34;)
 *             .geographicRegionBId(&#34;China&#34;)
 *             .build());
 * 
 *         var exampleBandwidthPackageAttachment = new BandwidthPackageAttachment(&#34;exampleBandwidthPackageAttachment&#34;, BandwidthPackageAttachmentArgs.builder()        
 *             .instanceId(example.id())
 *             .bandwidthPackageId(exampleBandwidthPackage.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * CEN bandwidth package attachment resource can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cen/bandwidthPackageAttachment:BandwidthPackageAttachment example bwp-abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:cen/bandwidthPackageAttachment:BandwidthPackageAttachment")
public class BandwidthPackageAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the bandwidth package.
     * 
     */
    @Export(name="bandwidthPackageId", refs={String.class}, tree="[0]")
    private Output<String> bandwidthPackageId;

    /**
     * @return The ID of the bandwidth package.
     * 
     */
    public Output<String> bandwidthPackageId() {
        return this.bandwidthPackageId;
    }
    /**
     * The ID of the CEN.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the CEN.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
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
        super("alicloud:cen/bandwidthPackageAttachment:BandwidthPackageAttachment", name, args == null ? BandwidthPackageAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BandwidthPackageAttachment(String name, Output<String> id, @Nullable BandwidthPackageAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/bandwidthPackageAttachment:BandwidthPackageAttachment", name, state, makeResourceOptions(options, id));
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
