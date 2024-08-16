// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sls.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ScheduledSqlScheduledSqlConfiguration {
    /**
     * @return Write Mode.
     * 
     */
    private @Nullable String dataFormat;
    /**
     * @return Target Endpoint.
     * 
     */
    private @Nullable String destEndpoint;
    /**
     * @return Target Logstore.
     * 
     */
    private @Nullable String destLogstore;
    /**
     * @return Target Project.
     * 
     */
    private @Nullable String destProject;
    /**
     * @return Write target role ARN.
     * 
     */
    private @Nullable String destRoleArn;
    /**
     * @return Schedule Start Time.
     * 
     */
    private @Nullable Integer fromTime;
    /**
     * @return SQL time window-start.
     * 
     */
    private @Nullable String fromTimeExpr;
    /**
     * @return Maximum retries.
     * 
     */
    private @Nullable Integer maxRetries;
    /**
     * @return SQL timeout.
     * 
     */
    private @Nullable Integer maxRunTimeInSeconds;
    /**
     * @return Parameter configuration.
     * 
     */
    private @Nullable Map<String,String> parameters;
    /**
     * @return Resource Pool.
     * 
     */
    private @Nullable String resourcePool;
    /**
     * @return Read role ARN.
     * 
     */
    private @Nullable String roleArn;
    /**
     * @return SQL statement.
     * 
     */
    private @Nullable String script;
    /**
     * @return Source Logstore.
     * 
     */
    private @Nullable String sourceLogstore;
    /**
     * @return SQL type.
     * 
     */
    private @Nullable String sqlType;
    /**
     * @return Time at end of schedule.
     * 
     */
    private @Nullable Integer toTime;
    /**
     * @return SQL time window-end.
     * 
     */
    private @Nullable String toTimeExpr;

    private ScheduledSqlScheduledSqlConfiguration() {}
    /**
     * @return Write Mode.
     * 
     */
    public Optional<String> dataFormat() {
        return Optional.ofNullable(this.dataFormat);
    }
    /**
     * @return Target Endpoint.
     * 
     */
    public Optional<String> destEndpoint() {
        return Optional.ofNullable(this.destEndpoint);
    }
    /**
     * @return Target Logstore.
     * 
     */
    public Optional<String> destLogstore() {
        return Optional.ofNullable(this.destLogstore);
    }
    /**
     * @return Target Project.
     * 
     */
    public Optional<String> destProject() {
        return Optional.ofNullable(this.destProject);
    }
    /**
     * @return Write target role ARN.
     * 
     */
    public Optional<String> destRoleArn() {
        return Optional.ofNullable(this.destRoleArn);
    }
    /**
     * @return Schedule Start Time.
     * 
     */
    public Optional<Integer> fromTime() {
        return Optional.ofNullable(this.fromTime);
    }
    /**
     * @return SQL time window-start.
     * 
     */
    public Optional<String> fromTimeExpr() {
        return Optional.ofNullable(this.fromTimeExpr);
    }
    /**
     * @return Maximum retries.
     * 
     */
    public Optional<Integer> maxRetries() {
        return Optional.ofNullable(this.maxRetries);
    }
    /**
     * @return SQL timeout.
     * 
     */
    public Optional<Integer> maxRunTimeInSeconds() {
        return Optional.ofNullable(this.maxRunTimeInSeconds);
    }
    /**
     * @return Parameter configuration.
     * 
     */
    public Map<String,String> parameters() {
        return this.parameters == null ? Map.of() : this.parameters;
    }
    /**
     * @return Resource Pool.
     * 
     */
    public Optional<String> resourcePool() {
        return Optional.ofNullable(this.resourcePool);
    }
    /**
     * @return Read role ARN.
     * 
     */
    public Optional<String> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }
    /**
     * @return SQL statement.
     * 
     */
    public Optional<String> script() {
        return Optional.ofNullable(this.script);
    }
    /**
     * @return Source Logstore.
     * 
     */
    public Optional<String> sourceLogstore() {
        return Optional.ofNullable(this.sourceLogstore);
    }
    /**
     * @return SQL type.
     * 
     */
    public Optional<String> sqlType() {
        return Optional.ofNullable(this.sqlType);
    }
    /**
     * @return Time at end of schedule.
     * 
     */
    public Optional<Integer> toTime() {
        return Optional.ofNullable(this.toTime);
    }
    /**
     * @return SQL time window-end.
     * 
     */
    public Optional<String> toTimeExpr() {
        return Optional.ofNullable(this.toTimeExpr);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ScheduledSqlScheduledSqlConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String dataFormat;
        private @Nullable String destEndpoint;
        private @Nullable String destLogstore;
        private @Nullable String destProject;
        private @Nullable String destRoleArn;
        private @Nullable Integer fromTime;
        private @Nullable String fromTimeExpr;
        private @Nullable Integer maxRetries;
        private @Nullable Integer maxRunTimeInSeconds;
        private @Nullable Map<String,String> parameters;
        private @Nullable String resourcePool;
        private @Nullable String roleArn;
        private @Nullable String script;
        private @Nullable String sourceLogstore;
        private @Nullable String sqlType;
        private @Nullable Integer toTime;
        private @Nullable String toTimeExpr;
        public Builder() {}
        public Builder(ScheduledSqlScheduledSqlConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dataFormat = defaults.dataFormat;
    	      this.destEndpoint = defaults.destEndpoint;
    	      this.destLogstore = defaults.destLogstore;
    	      this.destProject = defaults.destProject;
    	      this.destRoleArn = defaults.destRoleArn;
    	      this.fromTime = defaults.fromTime;
    	      this.fromTimeExpr = defaults.fromTimeExpr;
    	      this.maxRetries = defaults.maxRetries;
    	      this.maxRunTimeInSeconds = defaults.maxRunTimeInSeconds;
    	      this.parameters = defaults.parameters;
    	      this.resourcePool = defaults.resourcePool;
    	      this.roleArn = defaults.roleArn;
    	      this.script = defaults.script;
    	      this.sourceLogstore = defaults.sourceLogstore;
    	      this.sqlType = defaults.sqlType;
    	      this.toTime = defaults.toTime;
    	      this.toTimeExpr = defaults.toTimeExpr;
        }

        @CustomType.Setter
        public Builder dataFormat(@Nullable String dataFormat) {

            this.dataFormat = dataFormat;
            return this;
        }
        @CustomType.Setter
        public Builder destEndpoint(@Nullable String destEndpoint) {

            this.destEndpoint = destEndpoint;
            return this;
        }
        @CustomType.Setter
        public Builder destLogstore(@Nullable String destLogstore) {

            this.destLogstore = destLogstore;
            return this;
        }
        @CustomType.Setter
        public Builder destProject(@Nullable String destProject) {

            this.destProject = destProject;
            return this;
        }
        @CustomType.Setter
        public Builder destRoleArn(@Nullable String destRoleArn) {

            this.destRoleArn = destRoleArn;
            return this;
        }
        @CustomType.Setter
        public Builder fromTime(@Nullable Integer fromTime) {

            this.fromTime = fromTime;
            return this;
        }
        @CustomType.Setter
        public Builder fromTimeExpr(@Nullable String fromTimeExpr) {

            this.fromTimeExpr = fromTimeExpr;
            return this;
        }
        @CustomType.Setter
        public Builder maxRetries(@Nullable Integer maxRetries) {

            this.maxRetries = maxRetries;
            return this;
        }
        @CustomType.Setter
        public Builder maxRunTimeInSeconds(@Nullable Integer maxRunTimeInSeconds) {

            this.maxRunTimeInSeconds = maxRunTimeInSeconds;
            return this;
        }
        @CustomType.Setter
        public Builder parameters(@Nullable Map<String,String> parameters) {

            this.parameters = parameters;
            return this;
        }
        @CustomType.Setter
        public Builder resourcePool(@Nullable String resourcePool) {

            this.resourcePool = resourcePool;
            return this;
        }
        @CustomType.Setter
        public Builder roleArn(@Nullable String roleArn) {

            this.roleArn = roleArn;
            return this;
        }
        @CustomType.Setter
        public Builder script(@Nullable String script) {

            this.script = script;
            return this;
        }
        @CustomType.Setter
        public Builder sourceLogstore(@Nullable String sourceLogstore) {

            this.sourceLogstore = sourceLogstore;
            return this;
        }
        @CustomType.Setter
        public Builder sqlType(@Nullable String sqlType) {

            this.sqlType = sqlType;
            return this;
        }
        @CustomType.Setter
        public Builder toTime(@Nullable Integer toTime) {

            this.toTime = toTime;
            return this;
        }
        @CustomType.Setter
        public Builder toTimeExpr(@Nullable String toTimeExpr) {

            this.toTimeExpr = toTimeExpr;
            return this;
        }
        public ScheduledSqlScheduledSqlConfiguration build() {
            final var _resultValue = new ScheduledSqlScheduledSqlConfiguration();
            _resultValue.dataFormat = dataFormat;
            _resultValue.destEndpoint = destEndpoint;
            _resultValue.destLogstore = destLogstore;
            _resultValue.destProject = destProject;
            _resultValue.destRoleArn = destRoleArn;
            _resultValue.fromTime = fromTime;
            _resultValue.fromTimeExpr = fromTimeExpr;
            _resultValue.maxRetries = maxRetries;
            _resultValue.maxRunTimeInSeconds = maxRunTimeInSeconds;
            _resultValue.parameters = parameters;
            _resultValue.resourcePool = resourcePool;
            _resultValue.roleArn = roleArn;
            _resultValue.script = script;
            _resultValue.sourceLogstore = sourceLogstore;
            _resultValue.sqlType = sqlType;
            _resultValue.toTime = toTime;
            _resultValue.toTimeExpr = toTimeExpr;
            return _resultValue;
        }
    }
}
