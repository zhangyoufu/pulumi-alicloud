// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dns.CustomLineArgs;
import com.pulumi.alicloud.dns.inputs.CustomLineState;
import com.pulumi.alicloud.dns.outputs.CustomLineIpSegmentList;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Alidns Custom Line resource.
 * 
 * For information about Alidns Custom Line and how to use it, see [What is Custom Line](https://www.alibabacloud.com/help/en/doc-detail/145059.html).
 * 
 * &gt; **NOTE:** Available since v1.151.0.
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
 * import com.pulumi.alicloud.dns.CustomLine;
 * import com.pulumi.alicloud.dns.CustomLineArgs;
 * import com.pulumi.alicloud.dns.inputs.CustomLineIpSegmentListArgs;
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
 *         var default_ = new CustomLine("default", CustomLineArgs.builder()        
 *             .customLineName("tf-example")
 *             .domainName("alicloud-provider.com")
 *             .ipSegmentLists(CustomLineIpSegmentListArgs.builder()
 *                 .startIp("192.0.2.123")
 *                 .endIp("192.0.2.125")
 *                 .build())
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
 * Alidns Custom Line can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:dns/customLine:CustomLine example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dns/customLine:CustomLine")
public class CustomLine extends com.pulumi.resources.CustomResource {
    /**
     * The name of the Custom Line.
     * 
     */
    @Export(name="customLineName", refs={String.class}, tree="[0]")
    private Output<String> customLineName;

    /**
     * @return The name of the Custom Line.
     * 
     */
    public Output<String> customLineName() {
        return this.customLineName;
    }
    /**
     * The Domain name.
     * 
     */
    @Export(name="domainName", refs={String.class}, tree="[0]")
    private Output<String> domainName;

    /**
     * @return The Domain name.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }
    /**
     * The IP segment list. See `ip_segment_list` below for details.
     * 
     */
    @Export(name="ipSegmentLists", refs={List.class,CustomLineIpSegmentList.class}, tree="[0,1]")
    private Output<List<CustomLineIpSegmentList>> ipSegmentLists;

    /**
     * @return The IP segment list. See `ip_segment_list` below for details.
     * 
     */
    public Output<List<CustomLineIpSegmentList>> ipSegmentLists() {
        return this.ipSegmentLists;
    }
    /**
     * The lang.
     * 
     */
    @Export(name="lang", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lang;

    /**
     * @return The lang.
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CustomLine(String name) {
        this(name, CustomLineArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CustomLine(String name, CustomLineArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CustomLine(String name, CustomLineArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/customLine:CustomLine", name, args == null ? CustomLineArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CustomLine(String name, Output<String> id, @Nullable CustomLineState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/customLine:CustomLine", name, state, makeResourceOptions(options, id));
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
    public static CustomLine get(String name, Output<String> id, @Nullable CustomLineState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CustomLine(name, id, state, options);
    }
}
