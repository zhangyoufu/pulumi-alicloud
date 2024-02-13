// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.threatdetection.ImageEventOperationArgs;
import com.pulumi.alicloud.threatdetection.inputs.ImageEventOperationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Threat Detection Image Event Operation resource. Image Event Operation.
 * 
 * For information about Threat Detection Image Event Operation and how to use it, see [What is Image Event Operation](https://www.alibabacloud.com/help/zh/security-center/developer-reference/api-sas-2018-12-03-addimageeventoperation).
 * 
 * &gt; **NOTE:** Available since v1.212.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.threatdetection.ImageEventOperation;
 * import com.pulumi.alicloud.threatdetection.ImageEventOperationArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         var default_ = new ImageEventOperation(&#34;default&#34;, ImageEventOperationArgs.builder()        
 *             .conditions(&#34;&#34;&#34;
 * [
 *   {
 *       &#34;condition&#34;:&#34;MD5&#34;,
 *       &#34;type&#34;:&#34;equals&#34;,
 *       &#34;value&#34;:&#34;0083a31cc0083a31ccf7c10367a6e783e&#34;
 *   }
 * ]
 * 
 *             &#34;&#34;&#34;)
 *             .eventKey(&#34;alibabacloud_ak&#34;)
 *             .eventName(&#34;阿里云AK&#34;)
 *             .eventType(&#34;maliciousFile&#34;)
 *             .operationCode(&#34;whitelist&#34;)
 *             .scenarios(&#34;&#34;&#34;
 * {
 *   &#34;type&#34;:&#34;default&#34;,
 *   &#34;value&#34;:&#34;&#34;
 * }
 * 
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Threat Detection Image Event Operation can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:threatdetection/imageEventOperation:ImageEventOperation example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:threatdetection/imageEventOperation:ImageEventOperation")
public class ImageEventOperation extends com.pulumi.resources.CustomResource {
    /**
     * Event Conditions.
     * 
     */
    @Export(name="conditions", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> conditions;

    /**
     * @return Event Conditions.
     * 
     */
    public Output<Optional<String>> conditions() {
        return Codegen.optional(this.conditions);
    }
    /**
     * Image Event Key.
     * 
     */
    @Export(name="eventKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> eventKey;

    /**
     * @return Image Event Key.
     * 
     */
    public Output<Optional<String>> eventKey() {
        return Codegen.optional(this.eventKey);
    }
    /**
     * Image Event Name.
     * 
     */
    @Export(name="eventName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> eventName;

    /**
     * @return Image Event Name.
     * 
     */
    public Output<Optional<String>> eventName() {
        return Codegen.optional(this.eventName);
    }
    /**
     * Image Event Type.
     * 
     */
    @Export(name="eventType", refs={String.class}, tree="[0]")
    private Output<String> eventType;

    /**
     * @return Image Event Type.
     * 
     */
    public Output<String> eventType() {
        return this.eventType;
    }
    /**
     * Event Operation Code.
     * 
     */
    @Export(name="operationCode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> operationCode;

    /**
     * @return Event Operation Code.
     * 
     */
    public Output<Optional<String>> operationCode() {
        return Codegen.optional(this.operationCode);
    }
    /**
     * Event Scenarios.
     * 
     */
    @Export(name="scenarios", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> scenarios;

    /**
     * @return Event Scenarios.
     * 
     */
    public Output<Optional<String>> scenarios() {
        return Codegen.optional(this.scenarios);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ImageEventOperation(String name) {
        this(name, ImageEventOperationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ImageEventOperation(String name, @Nullable ImageEventOperationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ImageEventOperation(String name, @Nullable ImageEventOperationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/imageEventOperation:ImageEventOperation", name, args == null ? ImageEventOperationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ImageEventOperation(String name, Output<String> id, @Nullable ImageEventOperationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/imageEventOperation:ImageEventOperation", name, state, makeResourceOptions(options, id));
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
    public static ImageEventOperation get(String name, Output<String> id, @Nullable ImageEventOperationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ImageEventOperation(name, id, state, options);
    }
}
