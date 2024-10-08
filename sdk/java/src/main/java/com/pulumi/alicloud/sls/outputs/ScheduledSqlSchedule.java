// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sls.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ScheduledSqlSchedule {
    /**
     * @return Cron expression, minimum precision is minutes, 24-hour clock. For example, 0 0/1 **indicates that the check is performed every one hour from 00:00. When type is set to Cron, cronExpression must be set.
     * 
     */
    private @Nullable String cronExpression;
    /**
     * @return Delay time.
     * 
     */
    private @Nullable Integer delay;
    /**
     * @return Time interval, such as 5m, 1H.
     * 
     */
    private @Nullable String interval;
    /**
     * @return Whether to execute the OSS import task immediately after it is created.
     * 
     */
    private @Nullable Boolean runImmediately;
    /**
     * @return Time Zone.
     * 
     */
    private @Nullable String timeZone;
    /**
     * @return Check the frequency type. Log Service checks the query and analysis results based on the frequency you configured. The value is as follows: FixedRate: checks the query and analysis results at fixed intervals. Cron: specifies a time interval through a Cron expression, and checks the query and analysis results at the specified time interval. Weekly: Check the query and analysis results at a fixed point in time on the day of the week. Daily: checks the query and analysis results at a fixed time point every day. Hourly: Check query and analysis results every hour.
     * 
     */
    private @Nullable String type;

    private ScheduledSqlSchedule() {}
    /**
     * @return Cron expression, minimum precision is minutes, 24-hour clock. For example, 0 0/1 **indicates that the check is performed every one hour from 00:00. When type is set to Cron, cronExpression must be set.
     * 
     */
    public Optional<String> cronExpression() {
        return Optional.ofNullable(this.cronExpression);
    }
    /**
     * @return Delay time.
     * 
     */
    public Optional<Integer> delay() {
        return Optional.ofNullable(this.delay);
    }
    /**
     * @return Time interval, such as 5m, 1H.
     * 
     */
    public Optional<String> interval() {
        return Optional.ofNullable(this.interval);
    }
    /**
     * @return Whether to execute the OSS import task immediately after it is created.
     * 
     */
    public Optional<Boolean> runImmediately() {
        return Optional.ofNullable(this.runImmediately);
    }
    /**
     * @return Time Zone.
     * 
     */
    public Optional<String> timeZone() {
        return Optional.ofNullable(this.timeZone);
    }
    /**
     * @return Check the frequency type. Log Service checks the query and analysis results based on the frequency you configured. The value is as follows: FixedRate: checks the query and analysis results at fixed intervals. Cron: specifies a time interval through a Cron expression, and checks the query and analysis results at the specified time interval. Weekly: Check the query and analysis results at a fixed point in time on the day of the week. Daily: checks the query and analysis results at a fixed time point every day. Hourly: Check query and analysis results every hour.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ScheduledSqlSchedule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String cronExpression;
        private @Nullable Integer delay;
        private @Nullable String interval;
        private @Nullable Boolean runImmediately;
        private @Nullable String timeZone;
        private @Nullable String type;
        public Builder() {}
        public Builder(ScheduledSqlSchedule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cronExpression = defaults.cronExpression;
    	      this.delay = defaults.delay;
    	      this.interval = defaults.interval;
    	      this.runImmediately = defaults.runImmediately;
    	      this.timeZone = defaults.timeZone;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder cronExpression(@Nullable String cronExpression) {

            this.cronExpression = cronExpression;
            return this;
        }
        @CustomType.Setter
        public Builder delay(@Nullable Integer delay) {

            this.delay = delay;
            return this;
        }
        @CustomType.Setter
        public Builder interval(@Nullable String interval) {

            this.interval = interval;
            return this;
        }
        @CustomType.Setter
        public Builder runImmediately(@Nullable Boolean runImmediately) {

            this.runImmediately = runImmediately;
            return this;
        }
        @CustomType.Setter
        public Builder timeZone(@Nullable String timeZone) {

            this.timeZone = timeZone;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {

            this.type = type;
            return this;
        }
        public ScheduledSqlSchedule build() {
            final var _resultValue = new ScheduledSqlSchedule();
            _resultValue.cronExpression = cronExpression;
            _resultValue.delay = delay;
            _resultValue.interval = interval;
            _resultValue.runImmediately = runImmediately;
            _resultValue.timeZone = timeZone;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
