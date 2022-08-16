// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.fc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.fc.FunctionAsyncInvokeConfigArgs;
import com.pulumi.alicloud.fc.inputs.FunctionAsyncInvokeConfigState;
import com.pulumi.alicloud.fc.outputs.FunctionAsyncInvokeConfigDestinationConfig;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an asynchronous invocation configuration for a FC Function or Alias.\
 *  For the detailed information, please refer to the [developer guide](https://www.alibabacloud.com/help/doc-detail/181866.htm).
 * 
 * &gt; **NOTE:** Available in 1.100.0+
 * 
 * ## Example Usage
 * ### Destination Configuration
 * 
 * &gt; **NOTE** Ensure the FC Function RAM Role has necessary permissions for the destination, such as `mns:SendMessage`, `mns:PublishMessage` or `fc:InvokeFunction`, otherwise the API will return a generic error.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.fc.FunctionAsyncInvokeConfig;
 * import com.pulumi.alicloud.fc.FunctionAsyncInvokeConfigArgs;
 * import com.pulumi.alicloud.fc.inputs.FunctionAsyncInvokeConfigDestinationConfigArgs;
 * import com.pulumi.alicloud.fc.inputs.FunctionAsyncInvokeConfigDestinationConfigOnFailureArgs;
 * import com.pulumi.alicloud.fc.inputs.FunctionAsyncInvokeConfigDestinationConfigOnSuccessArgs;
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
 *         var example = new FunctionAsyncInvokeConfig(&#34;example&#34;, FunctionAsyncInvokeConfigArgs.builder()        
 *             .serviceName(alicloud_fc_service.example().name())
 *             .functionName(alicloud_fc_function.example().name())
 *             .destinationConfig(FunctionAsyncInvokeConfigDestinationConfigArgs.builder()
 *                 .onFailure(FunctionAsyncInvokeConfigDestinationConfigOnFailureArgs.builder()
 *                     .destination(the_example_mns_queue_arn)
 *                     .build())
 *                 .onSuccess(FunctionAsyncInvokeConfigDestinationConfigOnSuccessArgs.builder()
 *                     .destination(the_example_mns_topic_arn)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Error Handling Configuration
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.fc.FunctionAsyncInvokeConfig;
 * import com.pulumi.alicloud.fc.FunctionAsyncInvokeConfigArgs;
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
 *         var example = new FunctionAsyncInvokeConfig(&#34;example&#34;, FunctionAsyncInvokeConfigArgs.builder()        
 *             .serviceName(alicloud_fc_service.example().name())
 *             .functionName(alicloud_fc_function.example().name())
 *             .maximumEventAgeInSeconds(60)
 *             .maximumRetryAttempts(0)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Async Job Configuration
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.fc.FunctionAsyncInvokeConfig;
 * import com.pulumi.alicloud.fc.FunctionAsyncInvokeConfigArgs;
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
 *         var example = new FunctionAsyncInvokeConfig(&#34;example&#34;, FunctionAsyncInvokeConfigArgs.builder()        
 *             .serviceName(alicloud_fc_service.example().name())
 *             .functionName(alicloud_fc_function.example().name())
 *             .statefulInvocation(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Configuration for Function Latest Unpublished Version
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.fc.FunctionAsyncInvokeConfig;
 * import com.pulumi.alicloud.fc.FunctionAsyncInvokeConfigArgs;
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
 *         var example = new FunctionAsyncInvokeConfig(&#34;example&#34;, FunctionAsyncInvokeConfigArgs.builder()        
 *             .serviceName(alicloud_fc_service.example().name())
 *             .functionName(alicloud_fc_function.example().name())
 *             .qualifier(&#34;LATEST&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Function Compute Function Async Invoke Configs can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:fc/functionAsyncInvokeConfig:FunctionAsyncInvokeConfig example my_function
 * ```
 * 
 */
@ResourceType(type="alicloud:fc/functionAsyncInvokeConfig:FunctionAsyncInvokeConfig")
public class FunctionAsyncInvokeConfig extends com.pulumi.resources.CustomResource {
    /**
     * The date this resource was created.
     * 
     */
    @Export(name="createdTime", type=String.class, parameters={})
    private Output<String> createdTime;

    /**
     * @return The date this resource was created.
     * 
     */
    public Output<String> createdTime() {
        return this.createdTime;
    }
    /**
     * Configuration block with destination configuration. See below for details.
     * 
     */
    @Export(name="destinationConfig", type=FunctionAsyncInvokeConfigDestinationConfig.class, parameters={})
    private Output</* @Nullable */ FunctionAsyncInvokeConfigDestinationConfig> destinationConfig;

    /**
     * @return Configuration block with destination configuration. See below for details.
     * 
     */
    public Output<Optional<FunctionAsyncInvokeConfigDestinationConfig>> destinationConfig() {
        return Codegen.optional(this.destinationConfig);
    }
    /**
     * Name of the Function Compute Function.
     * 
     */
    @Export(name="functionName", type=String.class, parameters={})
    private Output<String> functionName;

    /**
     * @return Name of the Function Compute Function.
     * 
     */
    public Output<String> functionName() {
        return this.functionName;
    }
    /**
     * The date this resource was last modified.
     * 
     */
    @Export(name="lastModifiedTime", type=String.class, parameters={})
    private Output<String> lastModifiedTime;

    /**
     * @return The date this resource was last modified.
     * 
     */
    public Output<String> lastModifiedTime() {
        return this.lastModifiedTime;
    }
    /**
     * Maximum age of a request that Function Compute sends to a function for processing in seconds. Valid values between 1 and 2592000 (between 60 and 21600 before v1.167.0).
     * 
     */
    @Export(name="maximumEventAgeInSeconds", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> maximumEventAgeInSeconds;

    /**
     * @return Maximum age of a request that Function Compute sends to a function for processing in seconds. Valid values between 1 and 2592000 (between 60 and 21600 before v1.167.0).
     * 
     */
    public Output<Optional<Integer>> maximumEventAgeInSeconds() {
        return Codegen.optional(this.maximumEventAgeInSeconds);
    }
    /**
     * Maximum number of times to retry when the function returns an error. Valid values between 0 and 8 (between 0 and 2 before v1.167.0). Defaults to 2.
     * 
     */
    @Export(name="maximumRetryAttempts", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> maximumRetryAttempts;

    /**
     * @return Maximum number of times to retry when the function returns an error. Valid values between 0 and 8 (between 0 and 2 before v1.167.0). Defaults to 2.
     * 
     */
    public Output<Optional<Integer>> maximumRetryAttempts() {
        return Codegen.optional(this.maximumRetryAttempts);
    }
    /**
     * Function Compute Function published version, `LATEST`, or Function Compute Alias name. The default value is `LATEST`.
     * 
     */
    @Export(name="qualifier", type=String.class, parameters={})
    private Output</* @Nullable */ String> qualifier;

    /**
     * @return Function Compute Function published version, `LATEST`, or Function Compute Alias name. The default value is `LATEST`.
     * 
     */
    public Output<Optional<String>> qualifier() {
        return Codegen.optional(this.qualifier);
    }
    /**
     * Name of the Function Compute Function, omitting any version or alias qualifier.
     * 
     */
    @Export(name="serviceName", type=String.class, parameters={})
    private Output<String> serviceName;

    /**
     * @return Name of the Function Compute Function, omitting any version or alias qualifier.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Function Compute async job configuration(also known as Task Mode). valid values true or false, default `false`
     * 
     */
    @Export(name="statefulInvocation", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> statefulInvocation;

    /**
     * @return Function Compute async job configuration(also known as Task Mode). valid values true or false, default `false`
     * 
     */
    public Output<Optional<Boolean>> statefulInvocation() {
        return Codegen.optional(this.statefulInvocation);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FunctionAsyncInvokeConfig(String name) {
        this(name, FunctionAsyncInvokeConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FunctionAsyncInvokeConfig(String name, FunctionAsyncInvokeConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FunctionAsyncInvokeConfig(String name, FunctionAsyncInvokeConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:fc/functionAsyncInvokeConfig:FunctionAsyncInvokeConfig", name, args == null ? FunctionAsyncInvokeConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FunctionAsyncInvokeConfig(String name, Output<String> id, @Nullable FunctionAsyncInvokeConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:fc/functionAsyncInvokeConfig:FunctionAsyncInvokeConfig", name, state, makeResourceOptions(options, id));
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
    public static FunctionAsyncInvokeConfig get(String name, Output<String> id, @Nullable FunctionAsyncInvokeConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FunctionAsyncInvokeConfig(name, id, state, options);
    }
}
