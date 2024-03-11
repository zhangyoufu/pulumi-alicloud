// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ros;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ros.StackArgs;
import com.pulumi.alicloud.ros.inputs.StackState;
import com.pulumi.alicloud.ros.outputs.StackParameter;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ROS Stack resource.
 * 
 * For information about ROS Stack and how to use it, see [What is Stack](https://www.alibabacloud.com/help/en/doc-detail/132086.htm).
 * 
 * &gt; **NOTE:** Available in v1.106.0+.
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
 * import com.pulumi.alicloud.ros.Stack;
 * import com.pulumi.alicloud.ros.StackArgs;
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
 *         var example = new Stack(&#34;example&#34;, StackArgs.builder()        
 *             .stackName(&#34;tf-testaccstack&#34;)
 *             .stackPolicyBody(&#34;&#34;&#34;
 *     {
 *     	&#34;Statement&#34;: [{
 *     		&#34;Action&#34;: &#34;Update:Delete&#34;,
 *     		&#34;Resource&#34;: &#34;*&#34;,
 *     		&#34;Effect&#34;: &#34;Allow&#34;,
 *     		&#34;Principal&#34;: &#34;*&#34;
 *     	}]
 *     }
 *     
 *             &#34;&#34;&#34;)
 *             .templateBody(&#34;&#34;&#34;
 *     {
 *     	&#34;ROSTemplateFormatVersion&#34;: &#34;2015-09-01&#34;
 *     }
 *     
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ROS Stack can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ros/stack:Stack example &lt;stack_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ros/stack:Stack")
public class Stack extends com.pulumi.resources.CustomResource {
    /**
     * Specifies whether to delete the stack after it is created.
     * 
     */
    @Export(name="createOption", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> createOption;

    /**
     * @return Specifies whether to delete the stack after it is created.
     * 
     */
    public Output<Optional<String>> createOption() {
        return Codegen.optional(this.createOption);
    }
    /**
     * Specifies whether to enable deletion protection on the stack. Valid values: `Disabled`, `Enabled`. Default to: `Disabled`
     * 
     */
    @Export(name="deletionProtection", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deletionProtection;

    /**
     * @return Specifies whether to enable deletion protection on the stack. Valid values: `Disabled`, `Enabled`. Default to: `Disabled`
     * 
     */
    public Output<Optional<String>> deletionProtection() {
        return Codegen.optional(this.deletionProtection);
    }
    /**
     * Specifies whether to disable rollback on stack creation failure. Default to: `false`.
     * 
     */
    @Export(name="disableRollback", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableRollback;

    /**
     * @return Specifies whether to disable rollback on stack creation failure. Default to: `false`.
     * 
     */
    public Output<Optional<Boolean>> disableRollback() {
        return Codegen.optional(this.disableRollback);
    }
    /**
     * The callback URL for receiving stack event N. Only HTTP POST is supported. Maximum value of N: 5.
     * 
     */
    @Export(name="notificationUrls", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> notificationUrls;

    /**
     * @return The callback URL for receiving stack event N. Only HTTP POST is supported. Maximum value of N: 5.
     * 
     */
    public Output<Optional<List<String>>> notificationUrls() {
        return Codegen.optional(this.notificationUrls);
    }
    /**
     * The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
     * 
     */
    @Export(name="parameters", refs={List.class,StackParameter.class}, tree="[0,1]")
    private Output</* @Nullable */ List<StackParameter>> parameters;

    /**
     * @return The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
     * 
     */
    public Output<Optional<List<StackParameter>>> parameters() {
        return Codegen.optional(this.parameters);
    }
    /**
     * The name of the RAM role. ROS assumes the specified RAM role to create the stack and call API operations by using the credentials of the role.
     * 
     */
    @Export(name="ramRoleName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ramRoleName;

    /**
     * @return The name of the RAM role. ROS assumes the specified RAM role to create the stack and call API operations by using the credentials of the role.
     * 
     */
    public Output<Optional<String>> ramRoleName() {
        return Codegen.optional(this.ramRoleName);
    }
    /**
     * Specifies whether to enable replacement update after a resource attribute that does not support modification update is changed. Modification update keeps the physical ID of the resource unchanged. However, the resource is deleted and then recreated, and its physical ID is changed if replacement update is enabled.
     * 
     */
    @Export(name="replacementOption", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> replacementOption;

    /**
     * @return Specifies whether to enable replacement update after a resource attribute that does not support modification update is changed. Modification update keeps the physical ID of the resource unchanged. However, the resource is deleted and then recreated, and its physical ID is changed if replacement update is enabled.
     * 
     */
    public Output<Optional<String>> replacementOption() {
        return Codegen.optional(this.replacementOption);
    }
    /**
     * The retain all resources.
     * 
     */
    @Export(name="retainAllResources", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> retainAllResources;

    /**
     * @return The retain all resources.
     * 
     */
    public Output<Optional<Boolean>> retainAllResources() {
        return Codegen.optional(this.retainAllResources);
    }
    /**
     * Specifies whether to retain the resources in the stack.
     * 
     */
    @Export(name="retainResources", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> retainResources;

    /**
     * @return Specifies whether to retain the resources in the stack.
     * 
     */
    public Output<Optional<List<String>>> retainResources() {
        return Codegen.optional(this.retainResources);
    }
    /**
     * The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
     * 
     */
    @Export(name="stackName", refs={String.class}, tree="[0]")
    private Output<String> stackName;

    /**
     * @return The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
     * 
     */
    public Output<String> stackName() {
        return this.stackName;
    }
    /**
     * The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.
     * 
     */
    @Export(name="stackPolicyBody", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stackPolicyBody;

    /**
     * @return The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.
     * 
     */
    public Output<Optional<String>> stackPolicyBody() {
        return Codegen.optional(this.stackPolicyBody);
    }
    /**
     * The structure that contains the body of the temporary overriding stack policy. The stack policy body must be 1 to 16,384 bytes in length.
     * 
     */
    @Export(name="stackPolicyDuringUpdateBody", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stackPolicyDuringUpdateBody;

    /**
     * @return The structure that contains the body of the temporary overriding stack policy. The stack policy body must be 1 to 16,384 bytes in length.
     * 
     */
    public Output<Optional<String>> stackPolicyDuringUpdateBody() {
        return Codegen.optional(this.stackPolicyDuringUpdateBody);
    }
    /**
     * The URL of the file that contains the temporary overriding stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
     * 
     */
    @Export(name="stackPolicyDuringUpdateUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stackPolicyDuringUpdateUrl;

    /**
     * @return The URL of the file that contains the temporary overriding stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
     * 
     */
    public Output<Optional<String>> stackPolicyDuringUpdateUrl() {
        return Codegen.optional(this.stackPolicyDuringUpdateUrl);
    }
    /**
     * The URL of the file that contains the stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
     * 
     */
    @Export(name="stackPolicyUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stackPolicyUrl;

    /**
     * @return The URL of the file that contains the stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
     * 
     */
    public Output<Optional<String>> stackPolicyUrl() {
        return Codegen.optional(this.stackPolicyUrl);
    }
    /**
     * The status of Stack.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of Stack.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.
     * 
     */
    @Export(name="templateBody", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> templateBody;

    /**
     * @return The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.
     * 
     */
    public Output<Optional<String>> templateBody() {
        return Codegen.optional(this.templateBody);
    }
    /**
     * The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template must be 1 to 524,288 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
     * 
     */
    @Export(name="templateUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> templateUrl;

    /**
     * @return The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template must be 1 to 524,288 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
     * 
     */
    public Output<Optional<String>> templateUrl() {
        return Codegen.optional(this.templateUrl);
    }
    /**
     * The version of the template.
     * 
     */
    @Export(name="templateVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> templateVersion;

    /**
     * @return The version of the template.
     * 
     */
    public Output<Optional<String>> templateVersion() {
        return Codegen.optional(this.templateVersion);
    }
    /**
     * The timeout period that is specified for the stack creation request. Default to: `60`.
     * 
     */
    @Export(name="timeoutInMinutes", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeoutInMinutes;

    /**
     * @return The timeout period that is specified for the stack creation request. Default to: `60`.
     * 
     */
    public Output<Optional<Integer>> timeoutInMinutes() {
        return Codegen.optional(this.timeoutInMinutes);
    }
    /**
     * Specifies whether to use the values that were passed last time for the parameters that you do not specify in the current request.
     * 
     */
    @Export(name="usePreviousParameters", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> usePreviousParameters;

    /**
     * @return Specifies whether to use the values that were passed last time for the parameters that you do not specify in the current request.
     * 
     */
    public Output<Optional<Boolean>> usePreviousParameters() {
        return Codegen.optional(this.usePreviousParameters);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Stack(String name) {
        this(name, StackArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Stack(String name, StackArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Stack(String name, StackArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ros/stack:Stack", name, args == null ? StackArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Stack(String name, Output<String> id, @Nullable StackState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ros/stack:Stack", name, state, makeResourceOptions(options, id));
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
    public static Stack get(String name, Output<String> id, @Nullable StackState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Stack(name, id, state, options);
    }
}
