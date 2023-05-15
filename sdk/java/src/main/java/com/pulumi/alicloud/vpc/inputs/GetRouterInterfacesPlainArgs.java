// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRouterInterfacesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRouterInterfacesPlainArgs Empty = new GetRouterInterfacesPlainArgs();

    /**
     * A list of router interface IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of router interface IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string used to filter by router interface name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string used to filter by router interface name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * ID of the peer router interface.
     * 
     */
    @Import(name="oppositeInterfaceId")
    private @Nullable String oppositeInterfaceId;

    /**
     * @return ID of the peer router interface.
     * 
     */
    public Optional<String> oppositeInterfaceId() {
        return Optional.ofNullable(this.oppositeInterfaceId);
    }

    /**
     * Account ID of the owner of the peer router interface.
     * 
     */
    @Import(name="oppositeInterfaceOwnerId")
    private @Nullable String oppositeInterfaceOwnerId;

    /**
     * @return Account ID of the owner of the peer router interface.
     * 
     */
    public Optional<String> oppositeInterfaceOwnerId() {
        return Optional.ofNullable(this.oppositeInterfaceOwnerId);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * Role of the router interface. Valid values are `InitiatingSide` (connection initiator) and
     * `AcceptingSide` (connection receiver). The value of this parameter must be `InitiatingSide` if the `router_type` is set to `VBR`.
     * 
     */
    @Import(name="role")
    private @Nullable String role;

    /**
     * @return Role of the router interface. Valid values are `InitiatingSide` (connection initiator) and
     * `AcceptingSide` (connection receiver). The value of this parameter must be `InitiatingSide` if the `router_type` is set to `VBR`.
     * 
     */
    public Optional<String> role() {
        return Optional.ofNullable(this.role);
    }

    /**
     * ID of the VRouter located in the local region.
     * 
     */
    @Import(name="routerId")
    private @Nullable String routerId;

    /**
     * @return ID of the VRouter located in the local region.
     * 
     */
    public Optional<String> routerId() {
        return Optional.ofNullable(this.routerId);
    }

    /**
     * Router type in the local region. Valid values are `VRouter` and `VBR` (physical connection).
     * 
     */
    @Import(name="routerType")
    private @Nullable String routerType;

    /**
     * @return Router type in the local region. Valid values are `VRouter` and `VBR` (physical connection).
     * 
     */
    public Optional<String> routerType() {
        return Optional.ofNullable(this.routerType);
    }

    /**
     * Specification of the link, such as `Small.1` (10Mb), `Middle.1` (100Mb), `Large.2` (2Gb), ...etc.
     * 
     */
    @Import(name="specification")
    private @Nullable String specification;

    /**
     * @return Specification of the link, such as `Small.1` (10Mb), `Middle.1` (100Mb), `Large.2` (2Gb), ...etc.
     * 
     */
    public Optional<String> specification() {
        return Optional.ofNullable(this.specification);
    }

    /**
     * Expected status. Valid values are `Active`, `Inactive` and `Idle`.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return Expected status. Valid values are `Active`, `Inactive` and `Idle`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    private GetRouterInterfacesPlainArgs() {}

    private GetRouterInterfacesPlainArgs(GetRouterInterfacesPlainArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.oppositeInterfaceId = $.oppositeInterfaceId;
        this.oppositeInterfaceOwnerId = $.oppositeInterfaceOwnerId;
        this.outputFile = $.outputFile;
        this.role = $.role;
        this.routerId = $.routerId;
        this.routerType = $.routerType;
        this.specification = $.specification;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRouterInterfacesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRouterInterfacesPlainArgs $;

        public Builder() {
            $ = new GetRouterInterfacesPlainArgs();
        }

        public Builder(GetRouterInterfacesPlainArgs defaults) {
            $ = new GetRouterInterfacesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of router interface IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of router interface IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string used to filter by router interface name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param oppositeInterfaceId ID of the peer router interface.
         * 
         * @return builder
         * 
         */
        public Builder oppositeInterfaceId(@Nullable String oppositeInterfaceId) {
            $.oppositeInterfaceId = oppositeInterfaceId;
            return this;
        }

        /**
         * @param oppositeInterfaceOwnerId Account ID of the owner of the peer router interface.
         * 
         * @return builder
         * 
         */
        public Builder oppositeInterfaceOwnerId(@Nullable String oppositeInterfaceOwnerId) {
            $.oppositeInterfaceOwnerId = oppositeInterfaceOwnerId;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param role Role of the router interface. Valid values are `InitiatingSide` (connection initiator) and
         * `AcceptingSide` (connection receiver). The value of this parameter must be `InitiatingSide` if the `router_type` is set to `VBR`.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable String role) {
            $.role = role;
            return this;
        }

        /**
         * @param routerId ID of the VRouter located in the local region.
         * 
         * @return builder
         * 
         */
        public Builder routerId(@Nullable String routerId) {
            $.routerId = routerId;
            return this;
        }

        /**
         * @param routerType Router type in the local region. Valid values are `VRouter` and `VBR` (physical connection).
         * 
         * @return builder
         * 
         */
        public Builder routerType(@Nullable String routerType) {
            $.routerType = routerType;
            return this;
        }

        /**
         * @param specification Specification of the link, such as `Small.1` (10Mb), `Middle.1` (100Mb), `Large.2` (2Gb), ...etc.
         * 
         * @return builder
         * 
         */
        public Builder specification(@Nullable String specification) {
            $.specification = specification;
            return this;
        }

        /**
         * @param status Expected status. Valid values are `Active`, `Inactive` and `Idle`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        public GetRouterInterfacesPlainArgs build() {
            return $;
        }
    }

}
