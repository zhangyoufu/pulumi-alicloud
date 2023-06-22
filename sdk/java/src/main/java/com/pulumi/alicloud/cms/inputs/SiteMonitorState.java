// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms.inputs;

import com.pulumi.alicloud.cms.inputs.SiteMonitorIspCityArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SiteMonitorState extends com.pulumi.resources.ResourceArgs {

    public static final SiteMonitorState Empty = new SiteMonitorState();

    /**
     * The URL or IP address monitored by the site monitoring task.
     * 
     */
    @Import(name="address")
    private @Nullable Output<String> address;

    /**
     * @return The URL or IP address monitored by the site monitoring task.
     * 
     */
    public Optional<Output<String>> address() {
        return Optional.ofNullable(this.address);
    }

    /**
     * The IDs of existing alert rules to be associated with the site monitoring task.
     * 
     */
    @Import(name="alertIds")
    private @Nullable Output<List<String>> alertIds;

    /**
     * @return The IDs of existing alert rules to be associated with the site monitoring task.
     * 
     */
    public Optional<Output<List<String>>> alertIds() {
        return Optional.ofNullable(this.alertIds);
    }

    /**
     * The time when the site monitoring task was created.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The time when the site monitoring task was created.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * The monitoring interval of the site monitoring task. Unit: minutes. Valid values: `1`, `5`, `15`, `30` and `60`. Default value: `1`. **NOTE:** From version 1.207.0, `interval` can be set to `30`, `60`.
     * 
     */
    @Import(name="interval")
    private @Nullable Output<Integer> interval;

    /**
     * @return The monitoring interval of the site monitoring task. Unit: minutes. Valid values: `1`, `5`, `15`, `30` and `60`. Default value: `1`. **NOTE:** From version 1.207.0, `interval` can be set to `30`, `60`.
     * 
     */
    public Optional<Output<Integer>> interval() {
        return Optional.ofNullable(this.interval);
    }

    /**
     * The detection points in a JSON array. For example, `[{&#34;city&#34;:&#34;546&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;572&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;738&#34;,&#34;isp&#34;:&#34;465&#34;}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring. See `isp_cities` below.
     * 
     */
    @Import(name="ispCities")
    private @Nullable Output<List<SiteMonitorIspCityArgs>> ispCities;

    /**
     * @return The detection points in a JSON array. For example, `[{&#34;city&#34;:&#34;546&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;572&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;738&#34;,&#34;isp&#34;:&#34;465&#34;}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring. See `isp_cities` below.
     * 
     */
    public Optional<Output<List<SiteMonitorIspCityArgs>>> ispCities() {
        return Optional.ofNullable(this.ispCities);
    }

    /**
     * The extended options of the protocol of the site monitoring task. The options vary according to the protocol.
     * 
     */
    @Import(name="optionsJson")
    private @Nullable Output<String> optionsJson;

    /**
     * @return The extended options of the protocol of the site monitoring task. The options vary according to the protocol.
     * 
     */
    public Optional<Output<String>> optionsJson() {
        return Optional.ofNullable(this.optionsJson);
    }

    /**
     * The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
     * 
     */
    @Import(name="taskName")
    private @Nullable Output<String> taskName;

    /**
     * @return The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
     * 
     */
    public Optional<Output<String>> taskName() {
        return Optional.ofNullable(this.taskName);
    }

    /**
     * The status of the site monitoring task.
     * 
     */
    @Import(name="taskState")
    private @Nullable Output<String> taskState;

    /**
     * @return The status of the site monitoring task.
     * 
     */
    public Optional<Output<String>> taskState() {
        return Optional.ofNullable(this.taskState);
    }

    /**
     * The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
     * 
     */
    @Import(name="taskType")
    private @Nullable Output<String> taskType;

    /**
     * @return The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
     * 
     */
    public Optional<Output<String>> taskType() {
        return Optional.ofNullable(this.taskType);
    }

    /**
     * The time when the site monitoring task was updated.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return The time when the site monitoring task was updated.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    private SiteMonitorState() {}

    private SiteMonitorState(SiteMonitorState $) {
        this.address = $.address;
        this.alertIds = $.alertIds;
        this.createTime = $.createTime;
        this.interval = $.interval;
        this.ispCities = $.ispCities;
        this.optionsJson = $.optionsJson;
        this.taskName = $.taskName;
        this.taskState = $.taskState;
        this.taskType = $.taskType;
        this.updateTime = $.updateTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SiteMonitorState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SiteMonitorState $;

        public Builder() {
            $ = new SiteMonitorState();
        }

        public Builder(SiteMonitorState defaults) {
            $ = new SiteMonitorState(Objects.requireNonNull(defaults));
        }

        /**
         * @param address The URL or IP address monitored by the site monitoring task.
         * 
         * @return builder
         * 
         */
        public Builder address(@Nullable Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address The URL or IP address monitored by the site monitoring task.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param alertIds The IDs of existing alert rules to be associated with the site monitoring task.
         * 
         * @return builder
         * 
         */
        public Builder alertIds(@Nullable Output<List<String>> alertIds) {
            $.alertIds = alertIds;
            return this;
        }

        /**
         * @param alertIds The IDs of existing alert rules to be associated with the site monitoring task.
         * 
         * @return builder
         * 
         */
        public Builder alertIds(List<String> alertIds) {
            return alertIds(Output.of(alertIds));
        }

        /**
         * @param alertIds The IDs of existing alert rules to be associated with the site monitoring task.
         * 
         * @return builder
         * 
         */
        public Builder alertIds(String... alertIds) {
            return alertIds(List.of(alertIds));
        }

        /**
         * @param createTime The time when the site monitoring task was created.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The time when the site monitoring task was created.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param interval The monitoring interval of the site monitoring task. Unit: minutes. Valid values: `1`, `5`, `15`, `30` and `60`. Default value: `1`. **NOTE:** From version 1.207.0, `interval` can be set to `30`, `60`.
         * 
         * @return builder
         * 
         */
        public Builder interval(@Nullable Output<Integer> interval) {
            $.interval = interval;
            return this;
        }

        /**
         * @param interval The monitoring interval of the site monitoring task. Unit: minutes. Valid values: `1`, `5`, `15`, `30` and `60`. Default value: `1`. **NOTE:** From version 1.207.0, `interval` can be set to `30`, `60`.
         * 
         * @return builder
         * 
         */
        public Builder interval(Integer interval) {
            return interval(Output.of(interval));
        }

        /**
         * @param ispCities The detection points in a JSON array. For example, `[{&#34;city&#34;:&#34;546&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;572&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;738&#34;,&#34;isp&#34;:&#34;465&#34;}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring. See `isp_cities` below.
         * 
         * @return builder
         * 
         */
        public Builder ispCities(@Nullable Output<List<SiteMonitorIspCityArgs>> ispCities) {
            $.ispCities = ispCities;
            return this;
        }

        /**
         * @param ispCities The detection points in a JSON array. For example, `[{&#34;city&#34;:&#34;546&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;572&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;738&#34;,&#34;isp&#34;:&#34;465&#34;}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring. See `isp_cities` below.
         * 
         * @return builder
         * 
         */
        public Builder ispCities(List<SiteMonitorIspCityArgs> ispCities) {
            return ispCities(Output.of(ispCities));
        }

        /**
         * @param ispCities The detection points in a JSON array. For example, `[{&#34;city&#34;:&#34;546&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;572&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;738&#34;,&#34;isp&#34;:&#34;465&#34;}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring. See `isp_cities` below.
         * 
         * @return builder
         * 
         */
        public Builder ispCities(SiteMonitorIspCityArgs... ispCities) {
            return ispCities(List.of(ispCities));
        }

        /**
         * @param optionsJson The extended options of the protocol of the site monitoring task. The options vary according to the protocol.
         * 
         * @return builder
         * 
         */
        public Builder optionsJson(@Nullable Output<String> optionsJson) {
            $.optionsJson = optionsJson;
            return this;
        }

        /**
         * @param optionsJson The extended options of the protocol of the site monitoring task. The options vary according to the protocol.
         * 
         * @return builder
         * 
         */
        public Builder optionsJson(String optionsJson) {
            return optionsJson(Output.of(optionsJson));
        }

        /**
         * @param taskName The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
         * 
         * @return builder
         * 
         */
        public Builder taskName(@Nullable Output<String> taskName) {
            $.taskName = taskName;
            return this;
        }

        /**
         * @param taskName The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
         * 
         * @return builder
         * 
         */
        public Builder taskName(String taskName) {
            return taskName(Output.of(taskName));
        }

        /**
         * @param taskState The status of the site monitoring task.
         * 
         * @return builder
         * 
         */
        public Builder taskState(@Nullable Output<String> taskState) {
            $.taskState = taskState;
            return this;
        }

        /**
         * @param taskState The status of the site monitoring task.
         * 
         * @return builder
         * 
         */
        public Builder taskState(String taskState) {
            return taskState(Output.of(taskState));
        }

        /**
         * @param taskType The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
         * 
         * @return builder
         * 
         */
        public Builder taskType(@Nullable Output<String> taskType) {
            $.taskType = taskType;
            return this;
        }

        /**
         * @param taskType The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
         * 
         * @return builder
         * 
         */
        public Builder taskType(String taskType) {
            return taskType(Output.of(taskType));
        }

        /**
         * @param updateTime The time when the site monitoring task was updated.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime The time when the site monitoring task was updated.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        public SiteMonitorState build() {
            return $;
        }
    }

}
