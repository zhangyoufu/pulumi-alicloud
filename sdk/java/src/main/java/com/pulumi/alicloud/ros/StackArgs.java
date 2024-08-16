// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ros;

import com.pulumi.alicloud.ros.inputs.StackParameterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class StackArgs extends com.pulumi.resources.ResourceArgs {

    public static final StackArgs Empty = new StackArgs();

    /**
     * Specifies whether to delete the stack after it is created.
     * 
     */
    @Import(name="createOption")
    private @Nullable Output<String> createOption;

    /**
     * @return Specifies whether to delete the stack after it is created.
     * 
     */
    public Optional<Output<String>> createOption() {
        return Optional.ofNullable(this.createOption);
    }

    /**
     * Specifies whether to enable deletion protection on the stack. Valid values: `Disabled`, `Enabled`. Default to: `Disabled`
     * 
     */
    @Import(name="deletionProtection")
    private @Nullable Output<String> deletionProtection;

    /**
     * @return Specifies whether to enable deletion protection on the stack. Valid values: `Disabled`, `Enabled`. Default to: `Disabled`
     * 
     */
    public Optional<Output<String>> deletionProtection() {
        return Optional.ofNullable(this.deletionProtection);
    }

    /**
     * Specifies whether to disable rollback on stack creation failure. Default to: `false`.
     * 
     */
    @Import(name="disableRollback")
    private @Nullable Output<Boolean> disableRollback;

    /**
     * @return Specifies whether to disable rollback on stack creation failure. Default to: `false`.
     * 
     */
    public Optional<Output<Boolean>> disableRollback() {
        return Optional.ofNullable(this.disableRollback);
    }

    /**
     * The callback URL for receiving stack event N. Only HTTP POST is supported. Maximum value of N: 5.
     * 
     */
    @Import(name="notificationUrls")
    private @Nullable Output<List<String>> notificationUrls;

    /**
     * @return The callback URL for receiving stack event N. Only HTTP POST is supported. Maximum value of N: 5.
     * 
     */
    public Optional<Output<List<String>>> notificationUrls() {
        return Optional.ofNullable(this.notificationUrls);
    }

    /**
     * The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
     * 
     */
    @Import(name="parameters")
    private @Nullable Output<List<StackParameterArgs>> parameters;

    /**
     * @return The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
     * 
     */
    public Optional<Output<List<StackParameterArgs>>> parameters() {
        return Optional.ofNullable(this.parameters);
    }

    /**
     * The name of the RAM role. ROS assumes the specified RAM role to create the stack and call API operations by using the credentials of the role.
     * 
     */
    @Import(name="ramRoleName")
    private @Nullable Output<String> ramRoleName;

    /**
     * @return The name of the RAM role. ROS assumes the specified RAM role to create the stack and call API operations by using the credentials of the role.
     * 
     */
    public Optional<Output<String>> ramRoleName() {
        return Optional.ofNullable(this.ramRoleName);
    }

    /**
     * Specifies whether to enable replacement update after a resource attribute that does not support modification update is changed. Modification update keeps the physical ID of the resource unchanged. However, the resource is deleted and then recreated, and its physical ID is changed if replacement update is enabled.
     * 
     */
    @Import(name="replacementOption")
    private @Nullable Output<String> replacementOption;

    /**
     * @return Specifies whether to enable replacement update after a resource attribute that does not support modification update is changed. Modification update keeps the physical ID of the resource unchanged. However, the resource is deleted and then recreated, and its physical ID is changed if replacement update is enabled.
     * 
     */
    public Optional<Output<String>> replacementOption() {
        return Optional.ofNullable(this.replacementOption);
    }

    /**
     * The retain all resources.
     * 
     */
    @Import(name="retainAllResources")
    private @Nullable Output<Boolean> retainAllResources;

    /**
     * @return The retain all resources.
     * 
     */
    public Optional<Output<Boolean>> retainAllResources() {
        return Optional.ofNullable(this.retainAllResources);
    }

    /**
     * Specifies whether to retain the resources in the stack.
     * 
     */
    @Import(name="retainResources")
    private @Nullable Output<List<String>> retainResources;

    /**
     * @return Specifies whether to retain the resources in the stack.
     * 
     */
    public Optional<Output<List<String>>> retainResources() {
        return Optional.ofNullable(this.retainResources);
    }

    /**
     * The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
     * 
     */
    @Import(name="stackName", required=true)
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
    @Import(name="stackPolicyBody")
    private @Nullable Output<String> stackPolicyBody;

    /**
     * @return The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.
     * 
     */
    public Optional<Output<String>> stackPolicyBody() {
        return Optional.ofNullable(this.stackPolicyBody);
    }

    /**
     * The structure that contains the body of the temporary overriding stack policy. The stack policy body must be 1 to 16,384 bytes in length.
     * 
     */
    @Import(name="stackPolicyDuringUpdateBody")
    private @Nullable Output<String> stackPolicyDuringUpdateBody;

    /**
     * @return The structure that contains the body of the temporary overriding stack policy. The stack policy body must be 1 to 16,384 bytes in length.
     * 
     */
    public Optional<Output<String>> stackPolicyDuringUpdateBody() {
        return Optional.ofNullable(this.stackPolicyDuringUpdateBody);
    }

    /**
     * The URL of the file that contains the temporary overriding stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
     * 
     */
    @Import(name="stackPolicyDuringUpdateUrl")
    private @Nullable Output<String> stackPolicyDuringUpdateUrl;

    /**
     * @return The URL of the file that contains the temporary overriding stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
     * 
     */
    public Optional<Output<String>> stackPolicyDuringUpdateUrl() {
        return Optional.ofNullable(this.stackPolicyDuringUpdateUrl);
    }

    /**
     * The URL of the file that contains the stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
     * 
     */
    @Import(name="stackPolicyUrl")
    private @Nullable Output<String> stackPolicyUrl;

    /**
     * @return The URL of the file that contains the stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
     * 
     */
    public Optional<Output<String>> stackPolicyUrl() {
        return Optional.ofNullable(this.stackPolicyUrl);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.
     * 
     */
    @Import(name="templateBody")
    private @Nullable Output<String> templateBody;

    /**
     * @return The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.
     * 
     */
    public Optional<Output<String>> templateBody() {
        return Optional.ofNullable(this.templateBody);
    }

    /**
     * The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template must be 1 to 524,288 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
     * 
     */
    @Import(name="templateUrl")
    private @Nullable Output<String> templateUrl;

    /**
     * @return The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template must be 1 to 524,288 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
     * 
     */
    public Optional<Output<String>> templateUrl() {
        return Optional.ofNullable(this.templateUrl);
    }

    /**
     * The version of the template.
     * 
     */
    @Import(name="templateVersion")
    private @Nullable Output<String> templateVersion;

    /**
     * @return The version of the template.
     * 
     */
    public Optional<Output<String>> templateVersion() {
        return Optional.ofNullable(this.templateVersion);
    }

    /**
     * The timeout period that is specified for the stack creation request. Default to: `60`.
     * 
     */
    @Import(name="timeoutInMinutes")
    private @Nullable Output<Integer> timeoutInMinutes;

    /**
     * @return The timeout period that is specified for the stack creation request. Default to: `60`.
     * 
     */
    public Optional<Output<Integer>> timeoutInMinutes() {
        return Optional.ofNullable(this.timeoutInMinutes);
    }

    /**
     * Specifies whether to use the values that were passed last time for the parameters that you do not specify in the current request.
     * 
     */
    @Import(name="usePreviousParameters")
    private @Nullable Output<Boolean> usePreviousParameters;

    /**
     * @return Specifies whether to use the values that were passed last time for the parameters that you do not specify in the current request.
     * 
     */
    public Optional<Output<Boolean>> usePreviousParameters() {
        return Optional.ofNullable(this.usePreviousParameters);
    }

    private StackArgs() {}

    private StackArgs(StackArgs $) {
        this.createOption = $.createOption;
        this.deletionProtection = $.deletionProtection;
        this.disableRollback = $.disableRollback;
        this.notificationUrls = $.notificationUrls;
        this.parameters = $.parameters;
        this.ramRoleName = $.ramRoleName;
        this.replacementOption = $.replacementOption;
        this.retainAllResources = $.retainAllResources;
        this.retainResources = $.retainResources;
        this.stackName = $.stackName;
        this.stackPolicyBody = $.stackPolicyBody;
        this.stackPolicyDuringUpdateBody = $.stackPolicyDuringUpdateBody;
        this.stackPolicyDuringUpdateUrl = $.stackPolicyDuringUpdateUrl;
        this.stackPolicyUrl = $.stackPolicyUrl;
        this.tags = $.tags;
        this.templateBody = $.templateBody;
        this.templateUrl = $.templateUrl;
        this.templateVersion = $.templateVersion;
        this.timeoutInMinutes = $.timeoutInMinutes;
        this.usePreviousParameters = $.usePreviousParameters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(StackArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private StackArgs $;

        public Builder() {
            $ = new StackArgs();
        }

        public Builder(StackArgs defaults) {
            $ = new StackArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param createOption Specifies whether to delete the stack after it is created.
         * 
         * @return builder
         * 
         */
        public Builder createOption(@Nullable Output<String> createOption) {
            $.createOption = createOption;
            return this;
        }

        /**
         * @param createOption Specifies whether to delete the stack after it is created.
         * 
         * @return builder
         * 
         */
        public Builder createOption(String createOption) {
            return createOption(Output.of(createOption));
        }

        /**
         * @param deletionProtection Specifies whether to enable deletion protection on the stack. Valid values: `Disabled`, `Enabled`. Default to: `Disabled`
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(@Nullable Output<String> deletionProtection) {
            $.deletionProtection = deletionProtection;
            return this;
        }

        /**
         * @param deletionProtection Specifies whether to enable deletion protection on the stack. Valid values: `Disabled`, `Enabled`. Default to: `Disabled`
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(String deletionProtection) {
            return deletionProtection(Output.of(deletionProtection));
        }

        /**
         * @param disableRollback Specifies whether to disable rollback on stack creation failure. Default to: `false`.
         * 
         * @return builder
         * 
         */
        public Builder disableRollback(@Nullable Output<Boolean> disableRollback) {
            $.disableRollback = disableRollback;
            return this;
        }

        /**
         * @param disableRollback Specifies whether to disable rollback on stack creation failure. Default to: `false`.
         * 
         * @return builder
         * 
         */
        public Builder disableRollback(Boolean disableRollback) {
            return disableRollback(Output.of(disableRollback));
        }

        /**
         * @param notificationUrls The callback URL for receiving stack event N. Only HTTP POST is supported. Maximum value of N: 5.
         * 
         * @return builder
         * 
         */
        public Builder notificationUrls(@Nullable Output<List<String>> notificationUrls) {
            $.notificationUrls = notificationUrls;
            return this;
        }

        /**
         * @param notificationUrls The callback URL for receiving stack event N. Only HTTP POST is supported. Maximum value of N: 5.
         * 
         * @return builder
         * 
         */
        public Builder notificationUrls(List<String> notificationUrls) {
            return notificationUrls(Output.of(notificationUrls));
        }

        /**
         * @param notificationUrls The callback URL for receiving stack event N. Only HTTP POST is supported. Maximum value of N: 5.
         * 
         * @return builder
         * 
         */
        public Builder notificationUrls(String... notificationUrls) {
            return notificationUrls(List.of(notificationUrls));
        }

        /**
         * @param parameters The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
         * 
         * @return builder
         * 
         */
        public Builder parameters(@Nullable Output<List<StackParameterArgs>> parameters) {
            $.parameters = parameters;
            return this;
        }

        /**
         * @param parameters The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
         * 
         * @return builder
         * 
         */
        public Builder parameters(List<StackParameterArgs> parameters) {
            return parameters(Output.of(parameters));
        }

        /**
         * @param parameters The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
         * 
         * @return builder
         * 
         */
        public Builder parameters(StackParameterArgs... parameters) {
            return parameters(List.of(parameters));
        }

        /**
         * @param ramRoleName The name of the RAM role. ROS assumes the specified RAM role to create the stack and call API operations by using the credentials of the role.
         * 
         * @return builder
         * 
         */
        public Builder ramRoleName(@Nullable Output<String> ramRoleName) {
            $.ramRoleName = ramRoleName;
            return this;
        }

        /**
         * @param ramRoleName The name of the RAM role. ROS assumes the specified RAM role to create the stack and call API operations by using the credentials of the role.
         * 
         * @return builder
         * 
         */
        public Builder ramRoleName(String ramRoleName) {
            return ramRoleName(Output.of(ramRoleName));
        }

        /**
         * @param replacementOption Specifies whether to enable replacement update after a resource attribute that does not support modification update is changed. Modification update keeps the physical ID of the resource unchanged. However, the resource is deleted and then recreated, and its physical ID is changed if replacement update is enabled.
         * 
         * @return builder
         * 
         */
        public Builder replacementOption(@Nullable Output<String> replacementOption) {
            $.replacementOption = replacementOption;
            return this;
        }

        /**
         * @param replacementOption Specifies whether to enable replacement update after a resource attribute that does not support modification update is changed. Modification update keeps the physical ID of the resource unchanged. However, the resource is deleted and then recreated, and its physical ID is changed if replacement update is enabled.
         * 
         * @return builder
         * 
         */
        public Builder replacementOption(String replacementOption) {
            return replacementOption(Output.of(replacementOption));
        }

        /**
         * @param retainAllResources The retain all resources.
         * 
         * @return builder
         * 
         */
        public Builder retainAllResources(@Nullable Output<Boolean> retainAllResources) {
            $.retainAllResources = retainAllResources;
            return this;
        }

        /**
         * @param retainAllResources The retain all resources.
         * 
         * @return builder
         * 
         */
        public Builder retainAllResources(Boolean retainAllResources) {
            return retainAllResources(Output.of(retainAllResources));
        }

        /**
         * @param retainResources Specifies whether to retain the resources in the stack.
         * 
         * @return builder
         * 
         */
        public Builder retainResources(@Nullable Output<List<String>> retainResources) {
            $.retainResources = retainResources;
            return this;
        }

        /**
         * @param retainResources Specifies whether to retain the resources in the stack.
         * 
         * @return builder
         * 
         */
        public Builder retainResources(List<String> retainResources) {
            return retainResources(Output.of(retainResources));
        }

        /**
         * @param retainResources Specifies whether to retain the resources in the stack.
         * 
         * @return builder
         * 
         */
        public Builder retainResources(String... retainResources) {
            return retainResources(List.of(retainResources));
        }

        /**
         * @param stackName The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
         * 
         * @return builder
         * 
         */
        public Builder stackName(Output<String> stackName) {
            $.stackName = stackName;
            return this;
        }

        /**
         * @param stackName The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
         * 
         * @return builder
         * 
         */
        public Builder stackName(String stackName) {
            return stackName(Output.of(stackName));
        }

        /**
         * @param stackPolicyBody The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.
         * 
         * @return builder
         * 
         */
        public Builder stackPolicyBody(@Nullable Output<String> stackPolicyBody) {
            $.stackPolicyBody = stackPolicyBody;
            return this;
        }

        /**
         * @param stackPolicyBody The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.
         * 
         * @return builder
         * 
         */
        public Builder stackPolicyBody(String stackPolicyBody) {
            return stackPolicyBody(Output.of(stackPolicyBody));
        }

        /**
         * @param stackPolicyDuringUpdateBody The structure that contains the body of the temporary overriding stack policy. The stack policy body must be 1 to 16,384 bytes in length.
         * 
         * @return builder
         * 
         */
        public Builder stackPolicyDuringUpdateBody(@Nullable Output<String> stackPolicyDuringUpdateBody) {
            $.stackPolicyDuringUpdateBody = stackPolicyDuringUpdateBody;
            return this;
        }

        /**
         * @param stackPolicyDuringUpdateBody The structure that contains the body of the temporary overriding stack policy. The stack policy body must be 1 to 16,384 bytes in length.
         * 
         * @return builder
         * 
         */
        public Builder stackPolicyDuringUpdateBody(String stackPolicyDuringUpdateBody) {
            return stackPolicyDuringUpdateBody(Output.of(stackPolicyDuringUpdateBody));
        }

        /**
         * @param stackPolicyDuringUpdateUrl The URL of the file that contains the temporary overriding stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
         * 
         * @return builder
         * 
         */
        public Builder stackPolicyDuringUpdateUrl(@Nullable Output<String> stackPolicyDuringUpdateUrl) {
            $.stackPolicyDuringUpdateUrl = stackPolicyDuringUpdateUrl;
            return this;
        }

        /**
         * @param stackPolicyDuringUpdateUrl The URL of the file that contains the temporary overriding stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
         * 
         * @return builder
         * 
         */
        public Builder stackPolicyDuringUpdateUrl(String stackPolicyDuringUpdateUrl) {
            return stackPolicyDuringUpdateUrl(Output.of(stackPolicyDuringUpdateUrl));
        }

        /**
         * @param stackPolicyUrl The URL of the file that contains the stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
         * 
         * @return builder
         * 
         */
        public Builder stackPolicyUrl(@Nullable Output<String> stackPolicyUrl) {
            $.stackPolicyUrl = stackPolicyUrl;
            return this;
        }

        /**
         * @param stackPolicyUrl The URL of the file that contains the stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
         * 
         * @return builder
         * 
         */
        public Builder stackPolicyUrl(String stackPolicyUrl) {
            return stackPolicyUrl(Output.of(stackPolicyUrl));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param templateBody The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.
         * 
         * @return builder
         * 
         */
        public Builder templateBody(@Nullable Output<String> templateBody) {
            $.templateBody = templateBody;
            return this;
        }

        /**
         * @param templateBody The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.
         * 
         * @return builder
         * 
         */
        public Builder templateBody(String templateBody) {
            return templateBody(Output.of(templateBody));
        }

        /**
         * @param templateUrl The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template must be 1 to 524,288 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
         * 
         * @return builder
         * 
         */
        public Builder templateUrl(@Nullable Output<String> templateUrl) {
            $.templateUrl = templateUrl;
            return this;
        }

        /**
         * @param templateUrl The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template must be 1 to 524,288 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
         * 
         * @return builder
         * 
         */
        public Builder templateUrl(String templateUrl) {
            return templateUrl(Output.of(templateUrl));
        }

        /**
         * @param templateVersion The version of the template.
         * 
         * @return builder
         * 
         */
        public Builder templateVersion(@Nullable Output<String> templateVersion) {
            $.templateVersion = templateVersion;
            return this;
        }

        /**
         * @param templateVersion The version of the template.
         * 
         * @return builder
         * 
         */
        public Builder templateVersion(String templateVersion) {
            return templateVersion(Output.of(templateVersion));
        }

        /**
         * @param timeoutInMinutes The timeout period that is specified for the stack creation request. Default to: `60`.
         * 
         * @return builder
         * 
         */
        public Builder timeoutInMinutes(@Nullable Output<Integer> timeoutInMinutes) {
            $.timeoutInMinutes = timeoutInMinutes;
            return this;
        }

        /**
         * @param timeoutInMinutes The timeout period that is specified for the stack creation request. Default to: `60`.
         * 
         * @return builder
         * 
         */
        public Builder timeoutInMinutes(Integer timeoutInMinutes) {
            return timeoutInMinutes(Output.of(timeoutInMinutes));
        }

        /**
         * @param usePreviousParameters Specifies whether to use the values that were passed last time for the parameters that you do not specify in the current request.
         * 
         * @return builder
         * 
         */
        public Builder usePreviousParameters(@Nullable Output<Boolean> usePreviousParameters) {
            $.usePreviousParameters = usePreviousParameters;
            return this;
        }

        /**
         * @param usePreviousParameters Specifies whether to use the values that were passed last time for the parameters that you do not specify in the current request.
         * 
         * @return builder
         * 
         */
        public Builder usePreviousParameters(Boolean usePreviousParameters) {
            return usePreviousParameters(Output.of(usePreviousParameters));
        }

        public StackArgs build() {
            if ($.stackName == null) {
                throw new MissingRequiredPropertyException("StackArgs", "stackName");
            }
            return $;
        }
    }

}
