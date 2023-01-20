// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPrometheusAlertRulesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPrometheusAlertRulesPlainArgs Empty = new GetPrometheusAlertRulesPlainArgs();

    /**
     * The ID of the cluster.
     * 
     */
    @Import(name="clusterId", required=true)
    private String clusterId;

    /**
     * @return The ID of the cluster.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }

    /**
     * A list of Prometheus Alert Rule IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Prometheus Alert Rule IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    @Import(name="matchExpressions")
    private @Nullable String matchExpressions;

    public Optional<String> matchExpressions() {
        return Optional.ofNullable(this.matchExpressions);
    }

    /**
     * A regex string to filter results by Prometheus Alert Rule name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Prometheus Alert Rule name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The status of the resource. Valid values: `0`, `1`.
     * 
     */
    @Import(name="status")
    private @Nullable Integer status;

    /**
     * @return The status of the resource. Valid values: `0`, `1`.
     * 
     */
    public Optional<Integer> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The type of the alert rule.
     * 
     */
    @Import(name="type")
    private @Nullable String type;

    /**
     * @return The type of the alert rule.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    private GetPrometheusAlertRulesPlainArgs() {}

    private GetPrometheusAlertRulesPlainArgs(GetPrometheusAlertRulesPlainArgs $) {
        this.clusterId = $.clusterId;
        this.ids = $.ids;
        this.matchExpressions = $.matchExpressions;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.status = $.status;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPrometheusAlertRulesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPrometheusAlertRulesPlainArgs $;

        public Builder() {
            $ = new GetPrometheusAlertRulesPlainArgs();
        }

        public Builder(GetPrometheusAlertRulesPlainArgs defaults) {
            $ = new GetPrometheusAlertRulesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId The ID of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param ids A list of Prometheus Alert Rule IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Prometheus Alert Rule IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        public Builder matchExpressions(@Nullable String matchExpressions) {
            $.matchExpressions = matchExpressions;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Prometheus Alert Rule name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param status The status of the resource. Valid values: `0`, `1`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Integer status) {
            $.status = status;
            return this;
        }

        /**
         * @param type The type of the alert rule.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        public GetPrometheusAlertRulesPlainArgs build() {
            $.clusterId = Objects.requireNonNull($.clusterId, "expected parameter 'clusterId' to be non-null");
            return $;
        }
    }

}
