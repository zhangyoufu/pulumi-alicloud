// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TemplateArgs extends com.pulumi.resources.ResourceArgs {

    public static final TemplateArgs Empty = new TemplateArgs();

    /**
     * When deleting a template, whether to delete its related executions. Default to `false`.
     * 
     */
    @Import(name="autoDeleteExecutions")
    private @Nullable Output<Boolean> autoDeleteExecutions;

    /**
     * @return When deleting a template, whether to delete its related executions. Default to `false`.
     * 
     */
    public Optional<Output<Boolean>> autoDeleteExecutions() {
        return Optional.ofNullable(this.autoDeleteExecutions);
    }

    /**
     * The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
     * 
     */
    @Import(name="content", required=true)
    private Output<String> content;

    /**
     * @return The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
     * 
     */
    public Output<String> content() {
        return this.content;
    }

    /**
     * The ID of resource group which the template belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of resource group which the template belongs.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
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
     * The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
     * 
     */
    @Import(name="templateName", required=true)
    private Output<String> templateName;

    /**
     * @return The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
     * 
     */
    public Output<String> templateName() {
        return this.templateName;
    }

    /**
     * The name of template version.
     * 
     */
    @Import(name="versionName")
    private @Nullable Output<String> versionName;

    /**
     * @return The name of template version.
     * 
     */
    public Optional<Output<String>> versionName() {
        return Optional.ofNullable(this.versionName);
    }

    private TemplateArgs() {}

    private TemplateArgs(TemplateArgs $) {
        this.autoDeleteExecutions = $.autoDeleteExecutions;
        this.content = $.content;
        this.resourceGroupId = $.resourceGroupId;
        this.tags = $.tags;
        this.templateName = $.templateName;
        this.versionName = $.versionName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TemplateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TemplateArgs $;

        public Builder() {
            $ = new TemplateArgs();
        }

        public Builder(TemplateArgs defaults) {
            $ = new TemplateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoDeleteExecutions When deleting a template, whether to delete its related executions. Default to `false`.
         * 
         * @return builder
         * 
         */
        public Builder autoDeleteExecutions(@Nullable Output<Boolean> autoDeleteExecutions) {
            $.autoDeleteExecutions = autoDeleteExecutions;
            return this;
        }

        /**
         * @param autoDeleteExecutions When deleting a template, whether to delete its related executions. Default to `false`.
         * 
         * @return builder
         * 
         */
        public Builder autoDeleteExecutions(Boolean autoDeleteExecutions) {
            return autoDeleteExecutions(Output.of(autoDeleteExecutions));
        }

        /**
         * @param content The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
         * 
         * @return builder
         * 
         */
        public Builder content(Output<String> content) {
            $.content = content;
            return this;
        }

        /**
         * @param content The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
         * 
         * @return builder
         * 
         */
        public Builder content(String content) {
            return content(Output.of(content));
        }

        /**
         * @param resourceGroupId The ID of resource group which the template belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of resource group which the template belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
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
         * @param templateName The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
         * 
         * @return builder
         * 
         */
        public Builder templateName(Output<String> templateName) {
            $.templateName = templateName;
            return this;
        }

        /**
         * @param templateName The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
         * 
         * @return builder
         * 
         */
        public Builder templateName(String templateName) {
            return templateName(Output.of(templateName));
        }

        /**
         * @param versionName The name of template version.
         * 
         * @return builder
         * 
         */
        public Builder versionName(@Nullable Output<String> versionName) {
            $.versionName = versionName;
            return this;
        }

        /**
         * @param versionName The name of template version.
         * 
         * @return builder
         * 
         */
        public Builder versionName(String versionName) {
            return versionName(Output.of(versionName));
        }

        public TemplateArgs build() {
            if ($.content == null) {
                throw new MissingRequiredPropertyException("TemplateArgs", "content");
            }
            if ($.templateName == null) {
                throw new MissingRequiredPropertyException("TemplateArgs", "templateName");
            }
            return $;
        }
    }

}
