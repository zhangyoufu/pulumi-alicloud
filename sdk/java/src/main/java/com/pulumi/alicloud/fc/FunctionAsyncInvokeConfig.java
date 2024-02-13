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
 *  For the detailed information, please refer to the [developer guide](https://www.alibabacloud.com/help/en/fc/developer-reference/api-fc-open-2021-04-06-putfunctionasyncinvokeconfig).
 * 
 * &gt; **NOTE:** Available since v1.100.0.
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetRegionsArgs;
 * import com.pulumi.random.RandomInteger;
 * import com.pulumi.random.RandomIntegerArgs;
 * import com.pulumi.alicloud.ram.Role;
 * import com.pulumi.alicloud.ram.RoleArgs;
 * import com.pulumi.alicloud.ram.Policy;
 * import com.pulumi.alicloud.ram.PolicyArgs;
 * import com.pulumi.alicloud.ram.RolePolicyAttachment;
 * import com.pulumi.alicloud.ram.RolePolicyAttachmentArgs;
 * import com.pulumi.alicloud.fc.Service;
 * import com.pulumi.alicloud.fc.ServiceArgs;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.BucketObject;
 * import com.pulumi.alicloud.oss.BucketObjectArgs;
 * import com.pulumi.alicloud.fc.Function;
 * import com.pulumi.alicloud.fc.FunctionArgs;
 * import com.pulumi.alicloud.mns.Queue;
 * import com.pulumi.alicloud.mns.Topic;
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
 *         final var defaultAccount = AlicloudFunctions.getAccount();
 * 
 *         final var defaultRegions = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         var defaultRandomInteger = new RandomInteger(&#34;defaultRandomInteger&#34;, RandomIntegerArgs.builder()        
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         var defaultRole = new Role(&#34;defaultRole&#34;, RoleArgs.builder()        
 *             .document(&#34;&#34;&#34;
 * 	{
 * 		&#34;Statement&#34;: [
 * 		  {
 * 			&#34;Action&#34;: &#34;sts:AssumeRole&#34;,
 * 			&#34;Effect&#34;: &#34;Allow&#34;,
 * 			&#34;Principal&#34;: {
 * 			  &#34;Service&#34;: [
 * 				&#34;fc.aliyuncs.com&#34;
 * 			  ]
 * 			}
 * 		  }
 * 		],
 * 		&#34;Version&#34;: &#34;1&#34;
 * 	}
 *             &#34;&#34;&#34;)
 *             .description(&#34;this is a example&#34;)
 *             .force(true)
 *             .build());
 * 
 *         var defaultPolicy = new Policy(&#34;defaultPolicy&#34;, PolicyArgs.builder()        
 *             .policyName(defaultRandomInteger.result().applyValue(result -&gt; String.format(&#34;examplepolicy%s&#34;, result)))
 *             .policyDocument(&#34;&#34;&#34;
 * 	{
 * 		&#34;Version&#34;: &#34;1&#34;,
 * 		&#34;Statement&#34;: [
 * 		  {
 * 			&#34;Action&#34;: &#34;mns:*&#34;,
 * 			&#34;Resource&#34;: &#34;*&#34;,
 * 			&#34;Effect&#34;: &#34;Allow&#34;
 * 		  }
 * 		]
 * 	  }
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *         var defaultRolePolicyAttachment = new RolePolicyAttachment(&#34;defaultRolePolicyAttachment&#34;, RolePolicyAttachmentArgs.builder()        
 *             .roleName(defaultRole.name())
 *             .policyName(defaultPolicy.name())
 *             .policyType(&#34;Custom&#34;)
 *             .build());
 * 
 *         var defaultService = new Service(&#34;defaultService&#34;, ServiceArgs.builder()        
 *             .description(&#34;example-value&#34;)
 *             .role(defaultRole.arn())
 *             .internetAccess(false)
 *             .build());
 * 
 *         var defaultBucket = new Bucket(&#34;defaultBucket&#34;, BucketArgs.builder()        
 *             .bucket(defaultRandomInteger.result().applyValue(result -&gt; String.format(&#34;terraform-example-%s&#34;, result)))
 *             .build());
 * 
 *         var defaultBucketObject = new BucketObject(&#34;defaultBucketObject&#34;, BucketObjectArgs.builder()        
 *             .bucket(defaultBucket.id())
 *             .key(&#34;index.py&#34;)
 *             .content(&#34;&#34;&#34;
 * import logging 
 * def handler(event, context): 
 * logger = logging.getLogger() 
 * logger.info(&#39;hello world&#39;) 
 * return &#39;hello world&#39;            &#34;&#34;&#34;)
 *             .build());
 * 
 *         var defaultFunction = new Function(&#34;defaultFunction&#34;, FunctionArgs.builder()        
 *             .service(defaultService.name())
 *             .description(&#34;example&#34;)
 *             .ossBucket(defaultBucket.id())
 *             .ossKey(defaultBucketObject.key())
 *             .memorySize(&#34;512&#34;)
 *             .runtime(&#34;python3.10&#34;)
 *             .handler(&#34;hello.handler&#34;)
 *             .build());
 * 
 *         var defaultQueue = new Queue(&#34;defaultQueue&#34;);
 * 
 *         var defaultTopic = new Topic(&#34;defaultTopic&#34;);
 * 
 *         var defaultFunctionAsyncInvokeConfig = new FunctionAsyncInvokeConfig(&#34;defaultFunctionAsyncInvokeConfig&#34;, FunctionAsyncInvokeConfigArgs.builder()        
 *             .serviceName(defaultService.name())
 *             .functionName(defaultFunction.name())
 *             .destinationConfig(FunctionAsyncInvokeConfigDestinationConfigArgs.builder()
 *                 .onFailure(FunctionAsyncInvokeConfigDestinationConfigOnFailureArgs.builder()
 *                     .destination(defaultQueue.name().applyValue(name -&gt; String.format(&#34;acs:mns:%s:%s:/queues/%s/messages&#34;, defaultRegions.applyValue(getRegionsResult -&gt; getRegionsResult.regions()[0].id()),defaultAccount.applyValue(getAccountResult -&gt; getAccountResult.id()),name)))
 *                     .build())
 *                 .onSuccess(FunctionAsyncInvokeConfigDestinationConfigOnSuccessArgs.builder()
 *                     .destination(defaultTopic.name().applyValue(name -&gt; String.format(&#34;acs:mns:%s:%s:/topics/%s/messages&#34;, defaultRegions.applyValue(getRegionsResult -&gt; getRegionsResult.regions()[0].id()),defaultAccount.applyValue(getAccountResult -&gt; getAccountResult.id()),name)))
 *                     .build())
 *                 .build())
 *             .maximumEventAgeInSeconds(60)
 *             .maximumRetryAttempts(0)
 *             .statefulInvocation(true)
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
 * $ pulumi import alicloud:fc/functionAsyncInvokeConfig:FunctionAsyncInvokeConfig example my_function
 * ```
 * 
 */
@ResourceType(type="alicloud:fc/functionAsyncInvokeConfig:FunctionAsyncInvokeConfig")
public class FunctionAsyncInvokeConfig extends com.pulumi.resources.CustomResource {
    /**
     * The date this resource was created.
     * 
     */
    @Export(name="createdTime", refs={String.class}, tree="[0]")
    private Output<String> createdTime;

    /**
     * @return The date this resource was created.
     * 
     */
    public Output<String> createdTime() {
        return this.createdTime;
    }
    /**
     * Configuration block with destination configuration. See `destination_config` below.
     * 
     */
    @Export(name="destinationConfig", refs={FunctionAsyncInvokeConfigDestinationConfig.class}, tree="[0]")
    private Output</* @Nullable */ FunctionAsyncInvokeConfigDestinationConfig> destinationConfig;

    /**
     * @return Configuration block with destination configuration. See `destination_config` below.
     * 
     */
    public Output<Optional<FunctionAsyncInvokeConfigDestinationConfig>> destinationConfig() {
        return Codegen.optional(this.destinationConfig);
    }
    /**
     * Name of the Function Compute Function.
     * 
     */
    @Export(name="functionName", refs={String.class}, tree="[0]")
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
    @Export(name="lastModifiedTime", refs={String.class}, tree="[0]")
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
    @Export(name="maximumEventAgeInSeconds", refs={Integer.class}, tree="[0]")
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
    @Export(name="maximumRetryAttempts", refs={Integer.class}, tree="[0]")
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
    @Export(name="qualifier", refs={String.class}, tree="[0]")
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
    @Export(name="serviceName", refs={String.class}, tree="[0]")
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
    @Export(name="statefulInvocation", refs={Boolean.class}, tree="[0]")
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
