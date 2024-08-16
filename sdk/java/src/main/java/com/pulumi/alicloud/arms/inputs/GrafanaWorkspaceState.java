// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GrafanaWorkspaceState extends com.pulumi.resources.ResourceArgs {

    public static final GrafanaWorkspaceState Empty = new GrafanaWorkspaceState();

    /**
     * The creation time of the resource.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The creation time of the resource.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The version of the grafana.
     * 
     */
    @Import(name="grafanaVersion")
    private @Nullable Output<String> grafanaVersion;

    /**
     * @return The version of the grafana.
     * 
     */
    public Optional<Output<String>> grafanaVersion() {
        return Optional.ofNullable(this.grafanaVersion);
    }

    /**
     * The edition of the grafana.
     * 
     */
    @Import(name="grafanaWorkspaceEdition")
    private @Nullable Output<String> grafanaWorkspaceEdition;

    /**
     * @return The edition of the grafana.
     * 
     */
    public Optional<Output<String>> grafanaWorkspaceEdition() {
        return Optional.ofNullable(this.grafanaWorkspaceEdition);
    }

    /**
     * The name of the resource.
     * 
     */
    @Import(name="grafanaWorkspaceName")
    private @Nullable Output<String> grafanaWorkspaceName;

    /**
     * @return The name of the resource.
     * 
     */
    public Optional<Output<String>> grafanaWorkspaceName() {
        return Optional.ofNullable(this.grafanaWorkspaceName);
    }

    /**
     * The ID of the resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The status of the resource.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The tag of the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return The tag of the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private GrafanaWorkspaceState() {}

    private GrafanaWorkspaceState(GrafanaWorkspaceState $) {
        this.createTime = $.createTime;
        this.description = $.description;
        this.grafanaVersion = $.grafanaVersion;
        this.grafanaWorkspaceEdition = $.grafanaWorkspaceEdition;
        this.grafanaWorkspaceName = $.grafanaWorkspaceName;
        this.resourceGroupId = $.resourceGroupId;
        this.status = $.status;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GrafanaWorkspaceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GrafanaWorkspaceState $;

        public Builder() {
            $ = new GrafanaWorkspaceState();
        }

        public Builder(GrafanaWorkspaceState defaults) {
            $ = new GrafanaWorkspaceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createTime The creation time of the resource.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The creation time of the resource.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param description Description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param grafanaVersion The version of the grafana.
         * 
         * @return builder
         * 
         */
        public Builder grafanaVersion(@Nullable Output<String> grafanaVersion) {
            $.grafanaVersion = grafanaVersion;
            return this;
        }

        /**
         * @param grafanaVersion The version of the grafana.
         * 
         * @return builder
         * 
         */
        public Builder grafanaVersion(String grafanaVersion) {
            return grafanaVersion(Output.of(grafanaVersion));
        }

        /**
         * @param grafanaWorkspaceEdition The edition of the grafana.
         * 
         * @return builder
         * 
         */
        public Builder grafanaWorkspaceEdition(@Nullable Output<String> grafanaWorkspaceEdition) {
            $.grafanaWorkspaceEdition = grafanaWorkspaceEdition;
            return this;
        }

        /**
         * @param grafanaWorkspaceEdition The edition of the grafana.
         * 
         * @return builder
         * 
         */
        public Builder grafanaWorkspaceEdition(String grafanaWorkspaceEdition) {
            return grafanaWorkspaceEdition(Output.of(grafanaWorkspaceEdition));
        }

        /**
         * @param grafanaWorkspaceName The name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder grafanaWorkspaceName(@Nullable Output<String> grafanaWorkspaceName) {
            $.grafanaWorkspaceName = grafanaWorkspaceName;
            return this;
        }

        /**
         * @param grafanaWorkspaceName The name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder grafanaWorkspaceName(String grafanaWorkspaceName) {
            return grafanaWorkspaceName(Output.of(grafanaWorkspaceName));
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags The tag of the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tag of the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public GrafanaWorkspaceState build() {
            return $;
        }
    }

}
