// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bp;

import com.pulumi.alicloud.bp.inputs.StudioApplicationInstanceArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class StudioApplicationArgs extends com.pulumi.resources.ResourceArgs {

    public static final StudioApplicationArgs Empty = new StudioApplicationArgs();

    /**
     * The name of the application.
     * 
     */
    @Import(name="applicationName", required=true)
    private Output<String> applicationName;

    /**
     * @return The name of the application.
     * 
     */
    public Output<String> applicationName() {
        return this.applicationName;
    }

    /**
     * The id of the area.
     * 
     */
    @Import(name="areaId")
    private @Nullable Output<String> areaId;

    /**
     * @return The id of the area.
     * 
     */
    public Optional<Output<String>> areaId() {
        return Optional.ofNullable(this.areaId);
    }

    /**
     * The configuration of the application.
     * 
     */
    @Import(name="configuration")
    private @Nullable Output<Map<String,String>> configuration;

    /**
     * @return The configuration of the application.
     * 
     */
    public Optional<Output<Map<String,String>>> configuration() {
        return Optional.ofNullable(this.configuration);
    }

    /**
     * The instance list. Support the creation of instances in the existing vpc under the application. See the following `Block instances`.
     * 
     */
    @Import(name="instances")
    private @Nullable Output<List<StudioApplicationInstanceArgs>> instances;

    /**
     * @return The instance list. Support the creation of instances in the existing vpc under the application. See the following `Block instances`.
     * 
     */
    public Optional<Output<List<StudioApplicationInstanceArgs>>> instances() {
        return Optional.ofNullable(this.instances);
    }

    /**
     * The id of the resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The id of the resource group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The id of the template.
     * 
     */
    @Import(name="templateId", required=true)
    private Output<String> templateId;

    /**
     * @return The id of the template.
     * 
     */
    public Output<String> templateId() {
        return this.templateId;
    }

    /**
     * The variables of the application.
     * 
     */
    @Import(name="variables")
    private @Nullable Output<Map<String,String>> variables;

    /**
     * @return The variables of the application.
     * 
     */
    public Optional<Output<Map<String,String>>> variables() {
        return Optional.ofNullable(this.variables);
    }

    private StudioApplicationArgs() {}

    private StudioApplicationArgs(StudioApplicationArgs $) {
        this.applicationName = $.applicationName;
        this.areaId = $.areaId;
        this.configuration = $.configuration;
        this.instances = $.instances;
        this.resourceGroupId = $.resourceGroupId;
        this.templateId = $.templateId;
        this.variables = $.variables;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(StudioApplicationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private StudioApplicationArgs $;

        public Builder() {
            $ = new StudioApplicationArgs();
        }

        public Builder(StudioApplicationArgs defaults) {
            $ = new StudioApplicationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param applicationName The name of the application.
         * 
         * @return builder
         * 
         */
        public Builder applicationName(Output<String> applicationName) {
            $.applicationName = applicationName;
            return this;
        }

        /**
         * @param applicationName The name of the application.
         * 
         * @return builder
         * 
         */
        public Builder applicationName(String applicationName) {
            return applicationName(Output.of(applicationName));
        }

        /**
         * @param areaId The id of the area.
         * 
         * @return builder
         * 
         */
        public Builder areaId(@Nullable Output<String> areaId) {
            $.areaId = areaId;
            return this;
        }

        /**
         * @param areaId The id of the area.
         * 
         * @return builder
         * 
         */
        public Builder areaId(String areaId) {
            return areaId(Output.of(areaId));
        }

        /**
         * @param configuration The configuration of the application.
         * 
         * @return builder
         * 
         */
        public Builder configuration(@Nullable Output<Map<String,String>> configuration) {
            $.configuration = configuration;
            return this;
        }

        /**
         * @param configuration The configuration of the application.
         * 
         * @return builder
         * 
         */
        public Builder configuration(Map<String,String> configuration) {
            return configuration(Output.of(configuration));
        }

        /**
         * @param instances The instance list. Support the creation of instances in the existing vpc under the application. See the following `Block instances`.
         * 
         * @return builder
         * 
         */
        public Builder instances(@Nullable Output<List<StudioApplicationInstanceArgs>> instances) {
            $.instances = instances;
            return this;
        }

        /**
         * @param instances The instance list. Support the creation of instances in the existing vpc under the application. See the following `Block instances`.
         * 
         * @return builder
         * 
         */
        public Builder instances(List<StudioApplicationInstanceArgs> instances) {
            return instances(Output.of(instances));
        }

        /**
         * @param instances The instance list. Support the creation of instances in the existing vpc under the application. See the following `Block instances`.
         * 
         * @return builder
         * 
         */
        public Builder instances(StudioApplicationInstanceArgs... instances) {
            return instances(List.of(instances));
        }

        /**
         * @param resourceGroupId The id of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The id of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param templateId The id of the template.
         * 
         * @return builder
         * 
         */
        public Builder templateId(Output<String> templateId) {
            $.templateId = templateId;
            return this;
        }

        /**
         * @param templateId The id of the template.
         * 
         * @return builder
         * 
         */
        public Builder templateId(String templateId) {
            return templateId(Output.of(templateId));
        }

        /**
         * @param variables The variables of the application.
         * 
         * @return builder
         * 
         */
        public Builder variables(@Nullable Output<Map<String,String>> variables) {
            $.variables = variables;
            return this;
        }

        /**
         * @param variables The variables of the application.
         * 
         * @return builder
         * 
         */
        public Builder variables(Map<String,String> variables) {
            return variables(Output.of(variables));
        }

        public StudioApplicationArgs build() {
            if ($.applicationName == null) {
                throw new MissingRequiredPropertyException("StudioApplicationArgs", "applicationName");
            }
            if ($.templateId == null) {
                throw new MissingRequiredPropertyException("StudioApplicationArgs", "templateId");
            }
            return $;
        }
    }

}
