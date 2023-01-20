// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.gpdb.inputs;

import com.pulumi.alicloud.gpdb.inputs.DbInstancePlanPlanConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DbInstancePlanState extends com.pulumi.resources.ResourceArgs {

    public static final DbInstancePlanState Empty = new DbInstancePlanState();

    /**
     * The ID of the Database instance.
     * 
     */
    @Import(name="dbInstanceId")
    private @Nullable Output<String> dbInstanceId;

    /**
     * @return The ID of the Database instance.
     * 
     */
    public Optional<Output<String>> dbInstanceId() {
        return Optional.ofNullable(this.dbInstanceId);
    }

    /**
     * The name of the Plan.
     * 
     */
    @Import(name="dbInstancePlanName")
    private @Nullable Output<String> dbInstancePlanName;

    /**
     * @return The name of the Plan.
     * 
     */
    public Optional<Output<String>> dbInstancePlanName() {
        return Optional.ofNullable(this.dbInstancePlanName);
    }

    /**
     * The plan config. See the following `Block plan_config`.
     * 
     */
    @Import(name="planConfigs")
    private @Nullable Output<List<DbInstancePlanPlanConfigArgs>> planConfigs;

    /**
     * @return The plan config. See the following `Block plan_config`.
     * 
     */
    public Optional<Output<List<DbInstancePlanPlanConfigArgs>>> planConfigs() {
        return Optional.ofNullable(this.planConfigs);
    }

    /**
     * The description of the Plan.
     * 
     */
    @Import(name="planDesc")
    private @Nullable Output<String> planDesc;

    /**
     * @return The description of the Plan.
     * 
     */
    public Optional<Output<String>> planDesc() {
        return Optional.ofNullable(this.planDesc);
    }

    /**
     * The end time of the Plan.
     * 
     */
    @Import(name="planEndDate")
    private @Nullable Output<String> planEndDate;

    /**
     * @return The end time of the Plan.
     * 
     */
    public Optional<Output<String>> planEndDate() {
        return Optional.ofNullable(this.planEndDate);
    }

    /**
     * The ID of DB Instance Plan.
     * 
     */
    @Import(name="planId")
    private @Nullable Output<String> planId;

    /**
     * @return The ID of DB Instance Plan.
     * 
     */
    public Optional<Output<String>> planId() {
        return Optional.ofNullable(this.planId);
    }

    /**
     * Plan scheduling type. Valid values: `Postpone`, `Regular`.
     * 
     */
    @Import(name="planScheduleType")
    private @Nullable Output<String> planScheduleType;

    /**
     * @return Plan scheduling type. Valid values: `Postpone`, `Regular`.
     * 
     */
    public Optional<Output<String>> planScheduleType() {
        return Optional.ofNullable(this.planScheduleType);
    }

    /**
     * The start time of the Plan.
     * 
     */
    @Import(name="planStartDate")
    private @Nullable Output<String> planStartDate;

    /**
     * @return The start time of the Plan.
     * 
     */
    public Optional<Output<String>> planStartDate() {
        return Optional.ofNullable(this.planStartDate);
    }

    /**
     * The type of the Plan. Valid values: `PauseResume`, `Resize`.
     * 
     */
    @Import(name="planType")
    private @Nullable Output<String> planType;

    /**
     * @return The type of the Plan. Valid values: `PauseResume`, `Resize`.
     * 
     */
    public Optional<Output<String>> planType() {
        return Optional.ofNullable(this.planType);
    }

    /**
     * The Status of the Plan. Valid values: `active`, `cancel`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The Status of the Plan. Valid values: `active`, `cancel`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private DbInstancePlanState() {}

    private DbInstancePlanState(DbInstancePlanState $) {
        this.dbInstanceId = $.dbInstanceId;
        this.dbInstancePlanName = $.dbInstancePlanName;
        this.planConfigs = $.planConfigs;
        this.planDesc = $.planDesc;
        this.planEndDate = $.planEndDate;
        this.planId = $.planId;
        this.planScheduleType = $.planScheduleType;
        this.planStartDate = $.planStartDate;
        this.planType = $.planType;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DbInstancePlanState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DbInstancePlanState $;

        public Builder() {
            $ = new DbInstancePlanState();
        }

        public Builder(DbInstancePlanState defaults) {
            $ = new DbInstancePlanState(Objects.requireNonNull(defaults));
        }

        /**
         * @param dbInstanceId The ID of the Database instance.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceId(@Nullable Output<String> dbInstanceId) {
            $.dbInstanceId = dbInstanceId;
            return this;
        }

        /**
         * @param dbInstanceId The ID of the Database instance.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceId(String dbInstanceId) {
            return dbInstanceId(Output.of(dbInstanceId));
        }

        /**
         * @param dbInstancePlanName The name of the Plan.
         * 
         * @return builder
         * 
         */
        public Builder dbInstancePlanName(@Nullable Output<String> dbInstancePlanName) {
            $.dbInstancePlanName = dbInstancePlanName;
            return this;
        }

        /**
         * @param dbInstancePlanName The name of the Plan.
         * 
         * @return builder
         * 
         */
        public Builder dbInstancePlanName(String dbInstancePlanName) {
            return dbInstancePlanName(Output.of(dbInstancePlanName));
        }

        /**
         * @param planConfigs The plan config. See the following `Block plan_config`.
         * 
         * @return builder
         * 
         */
        public Builder planConfigs(@Nullable Output<List<DbInstancePlanPlanConfigArgs>> planConfigs) {
            $.planConfigs = planConfigs;
            return this;
        }

        /**
         * @param planConfigs The plan config. See the following `Block plan_config`.
         * 
         * @return builder
         * 
         */
        public Builder planConfigs(List<DbInstancePlanPlanConfigArgs> planConfigs) {
            return planConfigs(Output.of(planConfigs));
        }

        /**
         * @param planConfigs The plan config. See the following `Block plan_config`.
         * 
         * @return builder
         * 
         */
        public Builder planConfigs(DbInstancePlanPlanConfigArgs... planConfigs) {
            return planConfigs(List.of(planConfigs));
        }

        /**
         * @param planDesc The description of the Plan.
         * 
         * @return builder
         * 
         */
        public Builder planDesc(@Nullable Output<String> planDesc) {
            $.planDesc = planDesc;
            return this;
        }

        /**
         * @param planDesc The description of the Plan.
         * 
         * @return builder
         * 
         */
        public Builder planDesc(String planDesc) {
            return planDesc(Output.of(planDesc));
        }

        /**
         * @param planEndDate The end time of the Plan.
         * 
         * @return builder
         * 
         */
        public Builder planEndDate(@Nullable Output<String> planEndDate) {
            $.planEndDate = planEndDate;
            return this;
        }

        /**
         * @param planEndDate The end time of the Plan.
         * 
         * @return builder
         * 
         */
        public Builder planEndDate(String planEndDate) {
            return planEndDate(Output.of(planEndDate));
        }

        /**
         * @param planId The ID of DB Instance Plan.
         * 
         * @return builder
         * 
         */
        public Builder planId(@Nullable Output<String> planId) {
            $.planId = planId;
            return this;
        }

        /**
         * @param planId The ID of DB Instance Plan.
         * 
         * @return builder
         * 
         */
        public Builder planId(String planId) {
            return planId(Output.of(planId));
        }

        /**
         * @param planScheduleType Plan scheduling type. Valid values: `Postpone`, `Regular`.
         * 
         * @return builder
         * 
         */
        public Builder planScheduleType(@Nullable Output<String> planScheduleType) {
            $.planScheduleType = planScheduleType;
            return this;
        }

        /**
         * @param planScheduleType Plan scheduling type. Valid values: `Postpone`, `Regular`.
         * 
         * @return builder
         * 
         */
        public Builder planScheduleType(String planScheduleType) {
            return planScheduleType(Output.of(planScheduleType));
        }

        /**
         * @param planStartDate The start time of the Plan.
         * 
         * @return builder
         * 
         */
        public Builder planStartDate(@Nullable Output<String> planStartDate) {
            $.planStartDate = planStartDate;
            return this;
        }

        /**
         * @param planStartDate The start time of the Plan.
         * 
         * @return builder
         * 
         */
        public Builder planStartDate(String planStartDate) {
            return planStartDate(Output.of(planStartDate));
        }

        /**
         * @param planType The type of the Plan. Valid values: `PauseResume`, `Resize`.
         * 
         * @return builder
         * 
         */
        public Builder planType(@Nullable Output<String> planType) {
            $.planType = planType;
            return this;
        }

        /**
         * @param planType The type of the Plan. Valid values: `PauseResume`, `Resize`.
         * 
         * @return builder
         * 
         */
        public Builder planType(String planType) {
            return planType(Output.of(planType));
        }

        /**
         * @param status The Status of the Plan. Valid values: `active`, `cancel`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The Status of the Plan. Valid values: `active`, `cancel`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public DbInstancePlanState build() {
            return $;
        }
    }

}
