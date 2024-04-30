// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sls.inputs;

import com.pulumi.alicloud.sls.inputs.AlertConfigurationArgs;
import com.pulumi.alicloud.sls.inputs.AlertScheduleArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AlertState extends com.pulumi.resources.ResourceArgs {

    public static final AlertState Empty = new AlertState();

    /**
     * Alert rule ID, unique under Project.
     * 
     */
    @Import(name="alertName")
    private @Nullable Output<String> alertName;

    /**
     * @return Alert rule ID, unique under Project.
     * 
     */
    public Optional<Output<String>> alertName() {
        return Optional.ofNullable(this.alertName);
    }

    /**
     * Detailed configuration of alarm monitoring rules. See `configuration` below.
     * 
     */
    @Import(name="configuration")
    private @Nullable Output<AlertConfigurationArgs> configuration;

    /**
     * @return Detailed configuration of alarm monitoring rules. See `configuration` below.
     * 
     */
    public Optional<Output<AlertConfigurationArgs>> configuration() {
        return Optional.ofNullable(this.configuration);
    }

    /**
     * Alarm rule creation time.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<Integer> createTime;

    /**
     * @return Alarm rule creation time.
     * 
     */
    public Optional<Output<Integer>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Compatible fields, set to empty strings.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Compatible fields, set to empty strings.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Display name of the alarm rule.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Display name of the alarm rule.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Project Name.
     * 
     */
    @Import(name="projectName")
    private @Nullable Output<String> projectName;

    /**
     * @return Project Name.
     * 
     */
    public Optional<Output<String>> projectName() {
        return Optional.ofNullable(this.projectName);
    }

    /**
     * Check the frequency-dependent configuration. See `schedule` below.
     * 
     */
    @Import(name="schedule")
    private @Nullable Output<AlertScheduleArgs> schedule;

    /**
     * @return Check the frequency-dependent configuration. See `schedule` below.
     * 
     */
    public Optional<Output<AlertScheduleArgs>> schedule() {
        return Optional.ofNullable(this.schedule);
    }

    /**
     * Resource attribute field representing alarm status.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Resource attribute field representing alarm status.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private AlertState() {}

    private AlertState(AlertState $) {
        this.alertName = $.alertName;
        this.configuration = $.configuration;
        this.createTime = $.createTime;
        this.description = $.description;
        this.displayName = $.displayName;
        this.projectName = $.projectName;
        this.schedule = $.schedule;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AlertState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AlertState $;

        public Builder() {
            $ = new AlertState();
        }

        public Builder(AlertState defaults) {
            $ = new AlertState(Objects.requireNonNull(defaults));
        }

        /**
         * @param alertName Alert rule ID, unique under Project.
         * 
         * @return builder
         * 
         */
        public Builder alertName(@Nullable Output<String> alertName) {
            $.alertName = alertName;
            return this;
        }

        /**
         * @param alertName Alert rule ID, unique under Project.
         * 
         * @return builder
         * 
         */
        public Builder alertName(String alertName) {
            return alertName(Output.of(alertName));
        }

        /**
         * @param configuration Detailed configuration of alarm monitoring rules. See `configuration` below.
         * 
         * @return builder
         * 
         */
        public Builder configuration(@Nullable Output<AlertConfigurationArgs> configuration) {
            $.configuration = configuration;
            return this;
        }

        /**
         * @param configuration Detailed configuration of alarm monitoring rules. See `configuration` below.
         * 
         * @return builder
         * 
         */
        public Builder configuration(AlertConfigurationArgs configuration) {
            return configuration(Output.of(configuration));
        }

        /**
         * @param createTime Alarm rule creation time.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<Integer> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime Alarm rule creation time.
         * 
         * @return builder
         * 
         */
        public Builder createTime(Integer createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param description Compatible fields, set to empty strings.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Compatible fields, set to empty strings.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param displayName Display name of the alarm rule.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Display name of the alarm rule.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param projectName Project Name.
         * 
         * @return builder
         * 
         */
        public Builder projectName(@Nullable Output<String> projectName) {
            $.projectName = projectName;
            return this;
        }

        /**
         * @param projectName Project Name.
         * 
         * @return builder
         * 
         */
        public Builder projectName(String projectName) {
            return projectName(Output.of(projectName));
        }

        /**
         * @param schedule Check the frequency-dependent configuration. See `schedule` below.
         * 
         * @return builder
         * 
         */
        public Builder schedule(@Nullable Output<AlertScheduleArgs> schedule) {
            $.schedule = schedule;
            return this;
        }

        /**
         * @param schedule Check the frequency-dependent configuration. See `schedule` below.
         * 
         * @return builder
         * 
         */
        public Builder schedule(AlertScheduleArgs schedule) {
            return schedule(Output.of(schedule));
        }

        /**
         * @param status Resource attribute field representing alarm status.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Resource attribute field representing alarm status.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public AlertState build() {
            return $;
        }
    }

}
