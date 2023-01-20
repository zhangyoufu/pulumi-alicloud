// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log;

import com.pulumi.alicloud.log.inputs.OssExportConfigColumnArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OssExportArgs extends com.pulumi.resources.ResourceArgs {

    public static final OssExportArgs Empty = new OssExportArgs();

    /**
     * The name of the oss bucket.
     * 
     */
    @Import(name="bucket", required=true)
    private Output<String> bucket;

    /**
     * @return The name of the oss bucket.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }

    /**
     * How often is it delivered every interval.
     * 
     */
    @Import(name="bufferInterval", required=true)
    private Output<Integer> bufferInterval;

    /**
     * @return How often is it delivered every interval.
     * 
     */
    public Output<Integer> bufferInterval() {
        return this.bufferInterval;
    }

    /**
     * Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
     * 
     */
    @Import(name="bufferSize", required=true)
    private Output<Integer> bufferSize;

    /**
     * @return Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
     * 
     */
    public Output<Integer> bufferSize() {
        return this.bufferSize;
    }

    /**
     * OSS data storage compression method, support: `none`, `snappy`, `zstd`, `gzip`. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
     * 
     */
    @Import(name="compressType")
    private @Nullable Output<String> compressType;

    /**
     * @return OSS data storage compression method, support: `none`, `snappy`, `zstd`, `gzip`. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
     * 
     */
    public Optional<Output<String>> compressType() {
        return Optional.ofNullable(this.compressType);
    }

    /**
     * Configure columns when `content_type` is `parquet` or `orc`.
     * 
     */
    @Import(name="configColumns")
    private @Nullable Output<List<OssExportConfigColumnArgs>> configColumns;

    /**
     * @return Configure columns when `content_type` is `parquet` or `orc`.
     * 
     */
    public Optional<Output<List<OssExportConfigColumnArgs>>> configColumns() {
        return Optional.ofNullable(this.configColumns);
    }

    /**
     * Storage format, only supports three types: `json`, `parquet`, `orc`, `csv`.
     * **According to the different format, please select the following parameters**
     * 
     */
    @Import(name="contentType", required=true)
    private Output<String> contentType;

    /**
     * @return Storage format, only supports three types: `json`, `parquet`, `orc`, `csv`.
     * **According to the different format, please select the following parameters**
     * 
     */
    public Output<String> contentType() {
        return this.contentType;
    }

    /**
     * Field configuration in csv content_type.
     * 
     */
    @Import(name="csvConfigColumns")
    private @Nullable Output<List<String>> csvConfigColumns;

    /**
     * @return Field configuration in csv content_type.
     * 
     */
    public Optional<Output<List<String>>> csvConfigColumns() {
        return Optional.ofNullable(this.csvConfigColumns);
    }

    /**
     * Separator configuration in csv content_type.
     * 
     */
    @Import(name="csvConfigDelimiter")
    private @Nullable Output<String> csvConfigDelimiter;

    /**
     * @return Separator configuration in csv content_type.
     * 
     */
    public Optional<Output<String>> csvConfigDelimiter() {
        return Optional.ofNullable(this.csvConfigDelimiter);
    }

    /**
     * escape in csv content_type.
     * 
     */
    @Import(name="csvConfigEscape")
    private @Nullable Output<String> csvConfigEscape;

    /**
     * @return escape in csv content_type.
     * 
     */
    public Optional<Output<String>> csvConfigEscape() {
        return Optional.ofNullable(this.csvConfigEscape);
    }

    /**
     * Indicates whether to write the field name to the CSV file, the default value is `false`.
     * 
     */
    @Import(name="csvConfigHeader")
    private @Nullable Output<Boolean> csvConfigHeader;

    /**
     * @return Indicates whether to write the field name to the CSV file, the default value is `false`.
     * 
     */
    public Optional<Output<Boolean>> csvConfigHeader() {
        return Optional.ofNullable(this.csvConfigHeader);
    }

    /**
     * lineFeed in csv content_type.
     * 
     */
    @Import(name="csvConfigLinefeed")
    private @Nullable Output<String> csvConfigLinefeed;

    /**
     * @return lineFeed in csv content_type.
     * 
     */
    public Optional<Output<String>> csvConfigLinefeed() {
        return Optional.ofNullable(this.csvConfigLinefeed);
    }

    /**
     * Invalid field content in csv content_type.
     * 
     */
    @Import(name="csvConfigNull")
    private @Nullable Output<String> csvConfigNull;

    /**
     * @return Invalid field content in csv content_type.
     * 
     */
    public Optional<Output<String>> csvConfigNull() {
        return Optional.ofNullable(this.csvConfigNull);
    }

    /**
     * Escape character in csv content_type.
     * 
     */
    @Import(name="csvConfigQuote")
    private @Nullable Output<String> csvConfigQuote;

    /**
     * @return Escape character in csv content_type.
     * 
     */
    public Optional<Output<String>> csvConfigQuote() {
        return Optional.ofNullable(this.csvConfigQuote);
    }

    /**
     * The display name for oss export.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The display name for oss export.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
     * 
     */
    @Import(name="exportName", required=true)
    private Output<String> exportName;

    /**
     * @return Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
     * 
     */
    public Output<String> exportName() {
        return this.exportName;
    }

    /**
     * The log from when to export to oss.
     * 
     */
    @Import(name="fromTime")
    private @Nullable Output<Integer> fromTime;

    /**
     * @return The log from when to export to oss.
     * 
     */
    public Optional<Output<Integer>> fromTime() {
        return Optional.ofNullable(this.fromTime);
    }

    /**
     * Whether to deliver the label when `content_type` = `json`.
     * 
     */
    @Import(name="jsonEnableTag")
    private @Nullable Output<Boolean> jsonEnableTag;

    /**
     * @return Whether to deliver the label when `content_type` = `json`.
     * 
     */
    public Optional<Output<Boolean>> jsonEnableTag() {
        return Optional.ofNullable(this.jsonEnableTag);
    }

    /**
     * Used for logstore reading, the role should have log read policy, such as `acs:ram::13234:role/logrole`, if `log_read_role_arn` is not set, `role_arn` is used to read logstore.
     * 
     */
    @Import(name="logReadRoleArn")
    private @Nullable Output<String> logReadRoleArn;

    /**
     * @return Used for logstore reading, the role should have log read policy, such as `acs:ram::13234:role/logrole`, if `log_read_role_arn` is not set, `role_arn` is used to read logstore.
     * 
     */
    public Optional<Output<String>> logReadRoleArn() {
        return Optional.ofNullable(this.logReadRoleArn);
    }

    /**
     * The name of the log logstore.
     * 
     */
    @Import(name="logstoreName", required=true)
    private Output<String> logstoreName;

    /**
     * @return The name of the log logstore.
     * 
     */
    public Output<String> logstoreName() {
        return this.logstoreName;
    }

    /**
     * The OSS Bucket directory is dynamically generated according to the creation time of the export task, it cannot start with a forward slash `/`, the default value is `%Y/%m/%d/%H/%M`.
     * 
     */
    @Import(name="pathFormat", required=true)
    private Output<String> pathFormat;

    /**
     * @return The OSS Bucket directory is dynamically generated according to the creation time of the export task, it cannot start with a forward slash `/`, the default value is `%Y/%m/%d/%H/%M`.
     * 
     */
    public Output<String> pathFormat() {
        return this.pathFormat;
    }

    /**
     * The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
     * 
     */
    @Import(name="prefix")
    private @Nullable Output<String> prefix;

    /**
     * @return The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
     * 
     */
    public Optional<Output<String>> prefix() {
        return Optional.ofNullable(this.prefix);
    }

    /**
     * The name of the log project. It is the only in one Alicloud account.
     * 
     */
    @Import(name="projectName", required=true)
    private Output<String> projectName;

    /**
     * @return The name of the log project. It is the only in one Alicloud account.
     * 
     */
    public Output<String> projectName() {
        return this.projectName;
    }

    /**
     * Used to write to oss bucket, the OSS Bucket owner creates the role mark which has the oss bucket write policy, such as `acs:ram::13234:role/logrole`.
     * 
     */
    @Import(name="roleArn")
    private @Nullable Output<String> roleArn;

    /**
     * @return Used to write to oss bucket, the OSS Bucket owner creates the role mark which has the oss bucket write policy, such as `acs:ram::13234:role/logrole`.
     * 
     */
    public Optional<Output<String>> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    /**
     * The suffix for the objects in which the shipped data is stored.
     * 
     */
    @Import(name="suffix")
    private @Nullable Output<String> suffix;

    /**
     * @return The suffix for the objects in which the shipped data is stored.
     * 
     */
    public Optional<Output<String>> suffix() {
        return Optional.ofNullable(this.suffix);
    }

    /**
     * This time zone that is used to format the time, `+0800` e.g.
     * 
     */
    @Import(name="timeZone", required=true)
    private Output<String> timeZone;

    /**
     * @return This time zone that is used to format the time, `+0800` e.g.
     * 
     */
    public Output<String> timeZone() {
        return this.timeZone;
    }

    private OssExportArgs() {}

    private OssExportArgs(OssExportArgs $) {
        this.bucket = $.bucket;
        this.bufferInterval = $.bufferInterval;
        this.bufferSize = $.bufferSize;
        this.compressType = $.compressType;
        this.configColumns = $.configColumns;
        this.contentType = $.contentType;
        this.csvConfigColumns = $.csvConfigColumns;
        this.csvConfigDelimiter = $.csvConfigDelimiter;
        this.csvConfigEscape = $.csvConfigEscape;
        this.csvConfigHeader = $.csvConfigHeader;
        this.csvConfigLinefeed = $.csvConfigLinefeed;
        this.csvConfigNull = $.csvConfigNull;
        this.csvConfigQuote = $.csvConfigQuote;
        this.displayName = $.displayName;
        this.exportName = $.exportName;
        this.fromTime = $.fromTime;
        this.jsonEnableTag = $.jsonEnableTag;
        this.logReadRoleArn = $.logReadRoleArn;
        this.logstoreName = $.logstoreName;
        this.pathFormat = $.pathFormat;
        this.prefix = $.prefix;
        this.projectName = $.projectName;
        this.roleArn = $.roleArn;
        this.suffix = $.suffix;
        this.timeZone = $.timeZone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OssExportArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OssExportArgs $;

        public Builder() {
            $ = new OssExportArgs();
        }

        public Builder(OssExportArgs defaults) {
            $ = new OssExportArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket The name of the oss bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket The name of the oss bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        /**
         * @param bufferInterval How often is it delivered every interval.
         * 
         * @return builder
         * 
         */
        public Builder bufferInterval(Output<Integer> bufferInterval) {
            $.bufferInterval = bufferInterval;
            return this;
        }

        /**
         * @param bufferInterval How often is it delivered every interval.
         * 
         * @return builder
         * 
         */
        public Builder bufferInterval(Integer bufferInterval) {
            return bufferInterval(Output.of(bufferInterval));
        }

        /**
         * @param bufferSize Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
         * 
         * @return builder
         * 
         */
        public Builder bufferSize(Output<Integer> bufferSize) {
            $.bufferSize = bufferSize;
            return this;
        }

        /**
         * @param bufferSize Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
         * 
         * @return builder
         * 
         */
        public Builder bufferSize(Integer bufferSize) {
            return bufferSize(Output.of(bufferSize));
        }

        /**
         * @param compressType OSS data storage compression method, support: `none`, `snappy`, `zstd`, `gzip`. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
         * 
         * @return builder
         * 
         */
        public Builder compressType(@Nullable Output<String> compressType) {
            $.compressType = compressType;
            return this;
        }

        /**
         * @param compressType OSS data storage compression method, support: `none`, `snappy`, `zstd`, `gzip`. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
         * 
         * @return builder
         * 
         */
        public Builder compressType(String compressType) {
            return compressType(Output.of(compressType));
        }

        /**
         * @param configColumns Configure columns when `content_type` is `parquet` or `orc`.
         * 
         * @return builder
         * 
         */
        public Builder configColumns(@Nullable Output<List<OssExportConfigColumnArgs>> configColumns) {
            $.configColumns = configColumns;
            return this;
        }

        /**
         * @param configColumns Configure columns when `content_type` is `parquet` or `orc`.
         * 
         * @return builder
         * 
         */
        public Builder configColumns(List<OssExportConfigColumnArgs> configColumns) {
            return configColumns(Output.of(configColumns));
        }

        /**
         * @param configColumns Configure columns when `content_type` is `parquet` or `orc`.
         * 
         * @return builder
         * 
         */
        public Builder configColumns(OssExportConfigColumnArgs... configColumns) {
            return configColumns(List.of(configColumns));
        }

        /**
         * @param contentType Storage format, only supports three types: `json`, `parquet`, `orc`, `csv`.
         * **According to the different format, please select the following parameters**
         * 
         * @return builder
         * 
         */
        public Builder contentType(Output<String> contentType) {
            $.contentType = contentType;
            return this;
        }

        /**
         * @param contentType Storage format, only supports three types: `json`, `parquet`, `orc`, `csv`.
         * **According to the different format, please select the following parameters**
         * 
         * @return builder
         * 
         */
        public Builder contentType(String contentType) {
            return contentType(Output.of(contentType));
        }

        /**
         * @param csvConfigColumns Field configuration in csv content_type.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigColumns(@Nullable Output<List<String>> csvConfigColumns) {
            $.csvConfigColumns = csvConfigColumns;
            return this;
        }

        /**
         * @param csvConfigColumns Field configuration in csv content_type.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigColumns(List<String> csvConfigColumns) {
            return csvConfigColumns(Output.of(csvConfigColumns));
        }

        /**
         * @param csvConfigColumns Field configuration in csv content_type.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigColumns(String... csvConfigColumns) {
            return csvConfigColumns(List.of(csvConfigColumns));
        }

        /**
         * @param csvConfigDelimiter Separator configuration in csv content_type.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigDelimiter(@Nullable Output<String> csvConfigDelimiter) {
            $.csvConfigDelimiter = csvConfigDelimiter;
            return this;
        }

        /**
         * @param csvConfigDelimiter Separator configuration in csv content_type.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigDelimiter(String csvConfigDelimiter) {
            return csvConfigDelimiter(Output.of(csvConfigDelimiter));
        }

        /**
         * @param csvConfigEscape escape in csv content_type.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigEscape(@Nullable Output<String> csvConfigEscape) {
            $.csvConfigEscape = csvConfigEscape;
            return this;
        }

        /**
         * @param csvConfigEscape escape in csv content_type.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigEscape(String csvConfigEscape) {
            return csvConfigEscape(Output.of(csvConfigEscape));
        }

        /**
         * @param csvConfigHeader Indicates whether to write the field name to the CSV file, the default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigHeader(@Nullable Output<Boolean> csvConfigHeader) {
            $.csvConfigHeader = csvConfigHeader;
            return this;
        }

        /**
         * @param csvConfigHeader Indicates whether to write the field name to the CSV file, the default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigHeader(Boolean csvConfigHeader) {
            return csvConfigHeader(Output.of(csvConfigHeader));
        }

        /**
         * @param csvConfigLinefeed lineFeed in csv content_type.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigLinefeed(@Nullable Output<String> csvConfigLinefeed) {
            $.csvConfigLinefeed = csvConfigLinefeed;
            return this;
        }

        /**
         * @param csvConfigLinefeed lineFeed in csv content_type.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigLinefeed(String csvConfigLinefeed) {
            return csvConfigLinefeed(Output.of(csvConfigLinefeed));
        }

        /**
         * @param csvConfigNull Invalid field content in csv content_type.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigNull(@Nullable Output<String> csvConfigNull) {
            $.csvConfigNull = csvConfigNull;
            return this;
        }

        /**
         * @param csvConfigNull Invalid field content in csv content_type.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigNull(String csvConfigNull) {
            return csvConfigNull(Output.of(csvConfigNull));
        }

        /**
         * @param csvConfigQuote Escape character in csv content_type.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigQuote(@Nullable Output<String> csvConfigQuote) {
            $.csvConfigQuote = csvConfigQuote;
            return this;
        }

        /**
         * @param csvConfigQuote Escape character in csv content_type.
         * 
         * @return builder
         * 
         */
        public Builder csvConfigQuote(String csvConfigQuote) {
            return csvConfigQuote(Output.of(csvConfigQuote));
        }

        /**
         * @param displayName The display name for oss export.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The display name for oss export.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param exportName Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
         * 
         * @return builder
         * 
         */
        public Builder exportName(Output<String> exportName) {
            $.exportName = exportName;
            return this;
        }

        /**
         * @param exportName Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
         * 
         * @return builder
         * 
         */
        public Builder exportName(String exportName) {
            return exportName(Output.of(exportName));
        }

        /**
         * @param fromTime The log from when to export to oss.
         * 
         * @return builder
         * 
         */
        public Builder fromTime(@Nullable Output<Integer> fromTime) {
            $.fromTime = fromTime;
            return this;
        }

        /**
         * @param fromTime The log from when to export to oss.
         * 
         * @return builder
         * 
         */
        public Builder fromTime(Integer fromTime) {
            return fromTime(Output.of(fromTime));
        }

        /**
         * @param jsonEnableTag Whether to deliver the label when `content_type` = `json`.
         * 
         * @return builder
         * 
         */
        public Builder jsonEnableTag(@Nullable Output<Boolean> jsonEnableTag) {
            $.jsonEnableTag = jsonEnableTag;
            return this;
        }

        /**
         * @param jsonEnableTag Whether to deliver the label when `content_type` = `json`.
         * 
         * @return builder
         * 
         */
        public Builder jsonEnableTag(Boolean jsonEnableTag) {
            return jsonEnableTag(Output.of(jsonEnableTag));
        }

        /**
         * @param logReadRoleArn Used for logstore reading, the role should have log read policy, such as `acs:ram::13234:role/logrole`, if `log_read_role_arn` is not set, `role_arn` is used to read logstore.
         * 
         * @return builder
         * 
         */
        public Builder logReadRoleArn(@Nullable Output<String> logReadRoleArn) {
            $.logReadRoleArn = logReadRoleArn;
            return this;
        }

        /**
         * @param logReadRoleArn Used for logstore reading, the role should have log read policy, such as `acs:ram::13234:role/logrole`, if `log_read_role_arn` is not set, `role_arn` is used to read logstore.
         * 
         * @return builder
         * 
         */
        public Builder logReadRoleArn(String logReadRoleArn) {
            return logReadRoleArn(Output.of(logReadRoleArn));
        }

        /**
         * @param logstoreName The name of the log logstore.
         * 
         * @return builder
         * 
         */
        public Builder logstoreName(Output<String> logstoreName) {
            $.logstoreName = logstoreName;
            return this;
        }

        /**
         * @param logstoreName The name of the log logstore.
         * 
         * @return builder
         * 
         */
        public Builder logstoreName(String logstoreName) {
            return logstoreName(Output.of(logstoreName));
        }

        /**
         * @param pathFormat The OSS Bucket directory is dynamically generated according to the creation time of the export task, it cannot start with a forward slash `/`, the default value is `%Y/%m/%d/%H/%M`.
         * 
         * @return builder
         * 
         */
        public Builder pathFormat(Output<String> pathFormat) {
            $.pathFormat = pathFormat;
            return this;
        }

        /**
         * @param pathFormat The OSS Bucket directory is dynamically generated according to the creation time of the export task, it cannot start with a forward slash `/`, the default value is `%Y/%m/%d/%H/%M`.
         * 
         * @return builder
         * 
         */
        public Builder pathFormat(String pathFormat) {
            return pathFormat(Output.of(pathFormat));
        }

        /**
         * @param prefix The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
         * 
         * @return builder
         * 
         */
        public Builder prefix(@Nullable Output<String> prefix) {
            $.prefix = prefix;
            return this;
        }

        /**
         * @param prefix The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
         * 
         * @return builder
         * 
         */
        public Builder prefix(String prefix) {
            return prefix(Output.of(prefix));
        }

        /**
         * @param projectName The name of the log project. It is the only in one Alicloud account.
         * 
         * @return builder
         * 
         */
        public Builder projectName(Output<String> projectName) {
            $.projectName = projectName;
            return this;
        }

        /**
         * @param projectName The name of the log project. It is the only in one Alicloud account.
         * 
         * @return builder
         * 
         */
        public Builder projectName(String projectName) {
            return projectName(Output.of(projectName));
        }

        /**
         * @param roleArn Used to write to oss bucket, the OSS Bucket owner creates the role mark which has the oss bucket write policy, such as `acs:ram::13234:role/logrole`.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn Used to write to oss bucket, the OSS Bucket owner creates the role mark which has the oss bucket write policy, such as `acs:ram::13234:role/logrole`.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param suffix The suffix for the objects in which the shipped data is stored.
         * 
         * @return builder
         * 
         */
        public Builder suffix(@Nullable Output<String> suffix) {
            $.suffix = suffix;
            return this;
        }

        /**
         * @param suffix The suffix for the objects in which the shipped data is stored.
         * 
         * @return builder
         * 
         */
        public Builder suffix(String suffix) {
            return suffix(Output.of(suffix));
        }

        /**
         * @param timeZone This time zone that is used to format the time, `+0800` e.g.
         * 
         * @return builder
         * 
         */
        public Builder timeZone(Output<String> timeZone) {
            $.timeZone = timeZone;
            return this;
        }

        /**
         * @param timeZone This time zone that is used to format the time, `+0800` e.g.
         * 
         * @return builder
         * 
         */
        public Builder timeZone(String timeZone) {
            return timeZone(Output.of(timeZone));
        }

        public OssExportArgs build() {
            $.bucket = Objects.requireNonNull($.bucket, "expected parameter 'bucket' to be non-null");
            $.bufferInterval = Objects.requireNonNull($.bufferInterval, "expected parameter 'bufferInterval' to be non-null");
            $.bufferSize = Objects.requireNonNull($.bufferSize, "expected parameter 'bufferSize' to be non-null");
            $.contentType = Objects.requireNonNull($.contentType, "expected parameter 'contentType' to be non-null");
            $.exportName = Objects.requireNonNull($.exportName, "expected parameter 'exportName' to be non-null");
            $.logstoreName = Objects.requireNonNull($.logstoreName, "expected parameter 'logstoreName' to be non-null");
            $.pathFormat = Objects.requireNonNull($.pathFormat, "expected parameter 'pathFormat' to be non-null");
            $.projectName = Objects.requireNonNull($.projectName, "expected parameter 'projectName' to be non-null");
            $.timeZone = Objects.requireNonNull($.timeZone, "expected parameter 'timeZone' to be non-null");
            return $;
        }
    }

}
