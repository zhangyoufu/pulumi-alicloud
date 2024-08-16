// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz.inputs;

import com.pulumi.alicloud.pvtz.inputs.ZoneUserInfoArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ZoneState extends com.pulumi.resources.ResourceArgs {

    public static final ZoneState Empty = new ZoneState();

    /**
     * Whether the Private Zone is ptr.
     * 
     */
    @Import(name="isPtr")
    private @Nullable Output<Boolean> isPtr;

    /**
     * @return Whether the Private Zone is ptr.
     * 
     */
    public Optional<Output<Boolean>> isPtr() {
        return Optional.ofNullable(this.isPtr);
    }

    /**
     * The language. Valid values: &#34;zh&#34;, &#34;en&#34;, &#34;jp&#34;.
     * 
     */
    @Import(name="lang")
    private @Nullable Output<String> lang;

    /**
     * @return The language. Valid values: &#34;zh&#34;, &#34;en&#34;, &#34;jp&#34;.
     * 
     */
    public Optional<Output<String>> lang() {
        return Optional.ofNullable(this.lang);
    }

    /**
     * The name of the Private Zone. The `name` has been deprecated from provider version 1.107.0. Please use &#39;zone_name&#39; instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from version 1.107.0. Use &#39;zone_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from version 1.107.0. Use 'zone_name' instead. */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Private Zone. The `name` has been deprecated from provider version 1.107.0. Please use &#39;zone_name&#39; instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from version 1.107.0. Use &#39;zone_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from version 1.107.0. Use 'zone_name' instead. */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The recursive DNS proxy. Valid values:
     * - ZONE: indicates that the recursive DNS proxy is disabled.
     * - RECORD: indicates that the recursive DNS proxy is enabled.
     *   Default to &#34;ZONE&#34;.
     * 
     */
    @Import(name="proxyPattern")
    private @Nullable Output<String> proxyPattern;

    /**
     * @return The recursive DNS proxy. Valid values:
     * - ZONE: indicates that the recursive DNS proxy is disabled.
     * - RECORD: indicates that the recursive DNS proxy is enabled.
     *   Default to &#34;ZONE&#34;.
     * 
     */
    public Optional<Output<String>> proxyPattern() {
        return Optional.ofNullable(this.proxyPattern);
    }

    /**
     * The count of the Private Zone Record.
     * 
     */
    @Import(name="recordCount")
    private @Nullable Output<Integer> recordCount;

    /**
     * @return The count of the Private Zone Record.
     * 
     */
    public Optional<Output<Integer>> recordCount() {
        return Optional.ofNullable(this.recordCount);
    }

    /**
     * The remark of the Private Zone.
     * 
     */
    @Import(name="remark")
    private @Nullable Output<String> remark;

    /**
     * @return The remark of the Private Zone.
     * 
     */
    public Optional<Output<String>> remark() {
        return Optional.ofNullable(this.remark);
    }

    /**
     * The Id of resource group which the Private Zone belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The Id of resource group which the Private Zone belongs.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The status of the host synchronization task. Valid values:  `ON`,`OFF`. **NOTE:** You can update the `sync_status` to enable/disable the host synchronization task.
     * 
     */
    @Import(name="syncStatus")
    private @Nullable Output<String> syncStatus;

    /**
     * @return The status of the host synchronization task. Valid values:  `ON`,`OFF`. **NOTE:** You can update the `sync_status` to enable/disable the host synchronization task.
     * 
     */
    public Optional<Output<String>> syncStatus() {
        return Optional.ofNullable(this.syncStatus);
    }

    /**
     * The tags of the Private Zone.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return The tags of the Private Zone.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The IP address of the client.
     * 
     */
    @Import(name="userClientIp")
    private @Nullable Output<String> userClientIp;

    /**
     * @return The IP address of the client.
     * 
     */
    public Optional<Output<String>> userClientIp() {
        return Optional.ofNullable(this.userClientIp);
    }

    /**
     * The user information of the host synchronization task. See `user_info` below.
     * 
     */
    @Import(name="userInfos")
    private @Nullable Output<List<ZoneUserInfoArgs>> userInfos;

    /**
     * @return The user information of the host synchronization task. See `user_info` below.
     * 
     */
    public Optional<Output<List<ZoneUserInfoArgs>>> userInfos() {
        return Optional.ofNullable(this.userInfos);
    }

    /**
     * The zone_name of the Private Zone. The `zone_name` is required when the value of the `name`  is Empty.
     * 
     */
    @Import(name="zoneName")
    private @Nullable Output<String> zoneName;

    /**
     * @return The zone_name of the Private Zone. The `zone_name` is required when the value of the `name`  is Empty.
     * 
     */
    public Optional<Output<String>> zoneName() {
        return Optional.ofNullable(this.zoneName);
    }

    private ZoneState() {}

    private ZoneState(ZoneState $) {
        this.isPtr = $.isPtr;
        this.lang = $.lang;
        this.name = $.name;
        this.proxyPattern = $.proxyPattern;
        this.recordCount = $.recordCount;
        this.remark = $.remark;
        this.resourceGroupId = $.resourceGroupId;
        this.syncStatus = $.syncStatus;
        this.tags = $.tags;
        this.userClientIp = $.userClientIp;
        this.userInfos = $.userInfos;
        this.zoneName = $.zoneName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ZoneState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ZoneState $;

        public Builder() {
            $ = new ZoneState();
        }

        public Builder(ZoneState defaults) {
            $ = new ZoneState(Objects.requireNonNull(defaults));
        }

        /**
         * @param isPtr Whether the Private Zone is ptr.
         * 
         * @return builder
         * 
         */
        public Builder isPtr(@Nullable Output<Boolean> isPtr) {
            $.isPtr = isPtr;
            return this;
        }

        /**
         * @param isPtr Whether the Private Zone is ptr.
         * 
         * @return builder
         * 
         */
        public Builder isPtr(Boolean isPtr) {
            return isPtr(Output.of(isPtr));
        }

        /**
         * @param lang The language. Valid values: &#34;zh&#34;, &#34;en&#34;, &#34;jp&#34;.
         * 
         * @return builder
         * 
         */
        public Builder lang(@Nullable Output<String> lang) {
            $.lang = lang;
            return this;
        }

        /**
         * @param lang The language. Valid values: &#34;zh&#34;, &#34;en&#34;, &#34;jp&#34;.
         * 
         * @return builder
         * 
         */
        public Builder lang(String lang) {
            return lang(Output.of(lang));
        }

        /**
         * @param name The name of the Private Zone. The `name` has been deprecated from provider version 1.107.0. Please use &#39;zone_name&#39; instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from version 1.107.0. Use &#39;zone_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from version 1.107.0. Use 'zone_name' instead. */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Private Zone. The `name` has been deprecated from provider version 1.107.0. Please use &#39;zone_name&#39; instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from version 1.107.0. Use &#39;zone_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from version 1.107.0. Use 'zone_name' instead. */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param proxyPattern The recursive DNS proxy. Valid values:
         * - ZONE: indicates that the recursive DNS proxy is disabled.
         * - RECORD: indicates that the recursive DNS proxy is enabled.
         *   Default to &#34;ZONE&#34;.
         * 
         * @return builder
         * 
         */
        public Builder proxyPattern(@Nullable Output<String> proxyPattern) {
            $.proxyPattern = proxyPattern;
            return this;
        }

        /**
         * @param proxyPattern The recursive DNS proxy. Valid values:
         * - ZONE: indicates that the recursive DNS proxy is disabled.
         * - RECORD: indicates that the recursive DNS proxy is enabled.
         *   Default to &#34;ZONE&#34;.
         * 
         * @return builder
         * 
         */
        public Builder proxyPattern(String proxyPattern) {
            return proxyPattern(Output.of(proxyPattern));
        }

        /**
         * @param recordCount The count of the Private Zone Record.
         * 
         * @return builder
         * 
         */
        public Builder recordCount(@Nullable Output<Integer> recordCount) {
            $.recordCount = recordCount;
            return this;
        }

        /**
         * @param recordCount The count of the Private Zone Record.
         * 
         * @return builder
         * 
         */
        public Builder recordCount(Integer recordCount) {
            return recordCount(Output.of(recordCount));
        }

        /**
         * @param remark The remark of the Private Zone.
         * 
         * @return builder
         * 
         */
        public Builder remark(@Nullable Output<String> remark) {
            $.remark = remark;
            return this;
        }

        /**
         * @param remark The remark of the Private Zone.
         * 
         * @return builder
         * 
         */
        public Builder remark(String remark) {
            return remark(Output.of(remark));
        }

        /**
         * @param resourceGroupId The Id of resource group which the Private Zone belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The Id of resource group which the Private Zone belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param syncStatus The status of the host synchronization task. Valid values:  `ON`,`OFF`. **NOTE:** You can update the `sync_status` to enable/disable the host synchronization task.
         * 
         * @return builder
         * 
         */
        public Builder syncStatus(@Nullable Output<String> syncStatus) {
            $.syncStatus = syncStatus;
            return this;
        }

        /**
         * @param syncStatus The status of the host synchronization task. Valid values:  `ON`,`OFF`. **NOTE:** You can update the `sync_status` to enable/disable the host synchronization task.
         * 
         * @return builder
         * 
         */
        public Builder syncStatus(String syncStatus) {
            return syncStatus(Output.of(syncStatus));
        }

        /**
         * @param tags The tags of the Private Zone.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags of the Private Zone.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param userClientIp The IP address of the client.
         * 
         * @return builder
         * 
         */
        public Builder userClientIp(@Nullable Output<String> userClientIp) {
            $.userClientIp = userClientIp;
            return this;
        }

        /**
         * @param userClientIp The IP address of the client.
         * 
         * @return builder
         * 
         */
        public Builder userClientIp(String userClientIp) {
            return userClientIp(Output.of(userClientIp));
        }

        /**
         * @param userInfos The user information of the host synchronization task. See `user_info` below.
         * 
         * @return builder
         * 
         */
        public Builder userInfos(@Nullable Output<List<ZoneUserInfoArgs>> userInfos) {
            $.userInfos = userInfos;
            return this;
        }

        /**
         * @param userInfos The user information of the host synchronization task. See `user_info` below.
         * 
         * @return builder
         * 
         */
        public Builder userInfos(List<ZoneUserInfoArgs> userInfos) {
            return userInfos(Output.of(userInfos));
        }

        /**
         * @param userInfos The user information of the host synchronization task. See `user_info` below.
         * 
         * @return builder
         * 
         */
        public Builder userInfos(ZoneUserInfoArgs... userInfos) {
            return userInfos(List.of(userInfos));
        }

        /**
         * @param zoneName The zone_name of the Private Zone. The `zone_name` is required when the value of the `name`  is Empty.
         * 
         * @return builder
         * 
         */
        public Builder zoneName(@Nullable Output<String> zoneName) {
            $.zoneName = zoneName;
            return this;
        }

        /**
         * @param zoneName The zone_name of the Private Zone. The `zone_name` is required when the value of the `name`  is Empty.
         * 
         * @return builder
         * 
         */
        public Builder zoneName(String zoneName) {
            return zoneName(Output.of(zoneName));
        }

        public ZoneState build() {
            return $;
        }
    }

}
