// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetPrometheisPromethei {
    /**
     * @return The token used to access the data source.
     * 
     */
    private String authToken;
    /**
     * @return The ID of the cluster.
     * 
     */
    private String clusterId;
    /**
     * @return The name of the cluster.
     * 
     */
    private String clusterName;
    /**
     * @return The type of the cluster.
     * 
     */
    private String clusterType;
    /**
     * @return The ID of the Grafana workspace.
     * 
     */
    private String grafanaInstanceId;
    /**
     * @return Http api public network address.
     * 
     */
    private String httpApiInterUrl;
    /**
     * @return Http api intranet address.
     * 
     */
    private String httpApiIntraUrl;
    /**
     * @return The ID of the Prometheus.
     * 
     */
    private String id;
    /**
     * @return PushGateway public network Url.
     * 
     */
    private String pushGateWayInterUrl;
    /**
     * @return PushGateway intranet Url.
     * 
     */
    private String pushGateWayIntraUrl;
    /**
     * @return Public Url of remoteRead.
     * 
     */
    private String remoteReadInterUrl;
    /**
     * @return RemoteRead intranet Url.
     * 
     */
    private String remoteReadIntraUrl;
    /**
     * @return RemoteWrite public Url.
     * 
     */
    private String remoteWriteInterUrl;
    /**
     * @return RemoteWrite Intranet Url.
     * 
     */
    private String remoteWriteIntraUrl;
    /**
     * @return The ID of the resource group.
     * 
     */
    private String resourceGroupId;
    /**
     * @return The ID of the security group.
     * 
     */
    private String securityGroupId;
    /**
     * @return The child instance json string of the globalView instance.
     * 
     */
    private String subClustersJson;
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    private Map<String,String> tags;
    /**
     * @return The ID of the VPC.
     * 
     */
    private String vpcId;
    /**
     * @return The ID of the vSwitch.
     * 
     */
    private String vswitchId;

    private GetPrometheisPromethei() {}
    /**
     * @return The token used to access the data source.
     * 
     */
    public String authToken() {
        return this.authToken;
    }
    /**
     * @return The ID of the cluster.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }
    /**
     * @return The name of the cluster.
     * 
     */
    public String clusterName() {
        return this.clusterName;
    }
    /**
     * @return The type of the cluster.
     * 
     */
    public String clusterType() {
        return this.clusterType;
    }
    /**
     * @return The ID of the Grafana workspace.
     * 
     */
    public String grafanaInstanceId() {
        return this.grafanaInstanceId;
    }
    /**
     * @return Http api public network address.
     * 
     */
    public String httpApiInterUrl() {
        return this.httpApiInterUrl;
    }
    /**
     * @return Http api intranet address.
     * 
     */
    public String httpApiIntraUrl() {
        return this.httpApiIntraUrl;
    }
    /**
     * @return The ID of the Prometheus.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return PushGateway public network Url.
     * 
     */
    public String pushGateWayInterUrl() {
        return this.pushGateWayInterUrl;
    }
    /**
     * @return PushGateway intranet Url.
     * 
     */
    public String pushGateWayIntraUrl() {
        return this.pushGateWayIntraUrl;
    }
    /**
     * @return Public Url of remoteRead.
     * 
     */
    public String remoteReadInterUrl() {
        return this.remoteReadInterUrl;
    }
    /**
     * @return RemoteRead intranet Url.
     * 
     */
    public String remoteReadIntraUrl() {
        return this.remoteReadIntraUrl;
    }
    /**
     * @return RemoteWrite public Url.
     * 
     */
    public String remoteWriteInterUrl() {
        return this.remoteWriteInterUrl;
    }
    /**
     * @return RemoteWrite Intranet Url.
     * 
     */
    public String remoteWriteIntraUrl() {
        return this.remoteWriteIntraUrl;
    }
    /**
     * @return The ID of the resource group.
     * 
     */
    public String resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * @return The ID of the security group.
     * 
     */
    public String securityGroupId() {
        return this.securityGroupId;
    }
    /**
     * @return The child instance json string of the globalView instance.
     * 
     */
    public String subClustersJson() {
        return this.subClustersJson;
    }
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Map<String,String> tags() {
        return this.tags;
    }
    /**
     * @return The ID of the VPC.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }
    /**
     * @return The ID of the vSwitch.
     * 
     */
    public String vswitchId() {
        return this.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPrometheisPromethei defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String authToken;
        private String clusterId;
        private String clusterName;
        private String clusterType;
        private String grafanaInstanceId;
        private String httpApiInterUrl;
        private String httpApiIntraUrl;
        private String id;
        private String pushGateWayInterUrl;
        private String pushGateWayIntraUrl;
        private String remoteReadInterUrl;
        private String remoteReadIntraUrl;
        private String remoteWriteInterUrl;
        private String remoteWriteIntraUrl;
        private String resourceGroupId;
        private String securityGroupId;
        private String subClustersJson;
        private Map<String,String> tags;
        private String vpcId;
        private String vswitchId;
        public Builder() {}
        public Builder(GetPrometheisPromethei defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.authToken = defaults.authToken;
    	      this.clusterId = defaults.clusterId;
    	      this.clusterName = defaults.clusterName;
    	      this.clusterType = defaults.clusterType;
    	      this.grafanaInstanceId = defaults.grafanaInstanceId;
    	      this.httpApiInterUrl = defaults.httpApiInterUrl;
    	      this.httpApiIntraUrl = defaults.httpApiIntraUrl;
    	      this.id = defaults.id;
    	      this.pushGateWayInterUrl = defaults.pushGateWayInterUrl;
    	      this.pushGateWayIntraUrl = defaults.pushGateWayIntraUrl;
    	      this.remoteReadInterUrl = defaults.remoteReadInterUrl;
    	      this.remoteReadIntraUrl = defaults.remoteReadIntraUrl;
    	      this.remoteWriteInterUrl = defaults.remoteWriteInterUrl;
    	      this.remoteWriteIntraUrl = defaults.remoteWriteIntraUrl;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.securityGroupId = defaults.securityGroupId;
    	      this.subClustersJson = defaults.subClustersJson;
    	      this.tags = defaults.tags;
    	      this.vpcId = defaults.vpcId;
    	      this.vswitchId = defaults.vswitchId;
        }

        @CustomType.Setter
        public Builder authToken(String authToken) {
            if (authToken == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "authToken");
            }
            this.authToken = authToken;
            return this;
        }
        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder clusterName(String clusterName) {
            if (clusterName == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "clusterName");
            }
            this.clusterName = clusterName;
            return this;
        }
        @CustomType.Setter
        public Builder clusterType(String clusterType) {
            if (clusterType == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "clusterType");
            }
            this.clusterType = clusterType;
            return this;
        }
        @CustomType.Setter
        public Builder grafanaInstanceId(String grafanaInstanceId) {
            if (grafanaInstanceId == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "grafanaInstanceId");
            }
            this.grafanaInstanceId = grafanaInstanceId;
            return this;
        }
        @CustomType.Setter
        public Builder httpApiInterUrl(String httpApiInterUrl) {
            if (httpApiInterUrl == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "httpApiInterUrl");
            }
            this.httpApiInterUrl = httpApiInterUrl;
            return this;
        }
        @CustomType.Setter
        public Builder httpApiIntraUrl(String httpApiIntraUrl) {
            if (httpApiIntraUrl == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "httpApiIntraUrl");
            }
            this.httpApiIntraUrl = httpApiIntraUrl;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder pushGateWayInterUrl(String pushGateWayInterUrl) {
            if (pushGateWayInterUrl == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "pushGateWayInterUrl");
            }
            this.pushGateWayInterUrl = pushGateWayInterUrl;
            return this;
        }
        @CustomType.Setter
        public Builder pushGateWayIntraUrl(String pushGateWayIntraUrl) {
            if (pushGateWayIntraUrl == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "pushGateWayIntraUrl");
            }
            this.pushGateWayIntraUrl = pushGateWayIntraUrl;
            return this;
        }
        @CustomType.Setter
        public Builder remoteReadInterUrl(String remoteReadInterUrl) {
            if (remoteReadInterUrl == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "remoteReadInterUrl");
            }
            this.remoteReadInterUrl = remoteReadInterUrl;
            return this;
        }
        @CustomType.Setter
        public Builder remoteReadIntraUrl(String remoteReadIntraUrl) {
            if (remoteReadIntraUrl == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "remoteReadIntraUrl");
            }
            this.remoteReadIntraUrl = remoteReadIntraUrl;
            return this;
        }
        @CustomType.Setter
        public Builder remoteWriteInterUrl(String remoteWriteInterUrl) {
            if (remoteWriteInterUrl == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "remoteWriteInterUrl");
            }
            this.remoteWriteInterUrl = remoteWriteInterUrl;
            return this;
        }
        @CustomType.Setter
        public Builder remoteWriteIntraUrl(String remoteWriteIntraUrl) {
            if (remoteWriteIntraUrl == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "remoteWriteIntraUrl");
            }
            this.remoteWriteIntraUrl = remoteWriteIntraUrl;
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupId(String resourceGroupId) {
            if (resourceGroupId == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "resourceGroupId");
            }
            this.resourceGroupId = resourceGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder securityGroupId(String securityGroupId) {
            if (securityGroupId == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "securityGroupId");
            }
            this.securityGroupId = securityGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder subClustersJson(String subClustersJson) {
            if (subClustersJson == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "subClustersJson");
            }
            this.subClustersJson = subClustersJson;
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            if (tags == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "tags");
            }
            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            if (vpcId == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "vpcId");
            }
            this.vpcId = vpcId;
            return this;
        }
        @CustomType.Setter
        public Builder vswitchId(String vswitchId) {
            if (vswitchId == null) {
              throw new MissingRequiredPropertyException("GetPrometheisPromethei", "vswitchId");
            }
            this.vswitchId = vswitchId;
            return this;
        }
        public GetPrometheisPromethei build() {
            final var _resultValue = new GetPrometheisPromethei();
            _resultValue.authToken = authToken;
            _resultValue.clusterId = clusterId;
            _resultValue.clusterName = clusterName;
            _resultValue.clusterType = clusterType;
            _resultValue.grafanaInstanceId = grafanaInstanceId;
            _resultValue.httpApiInterUrl = httpApiInterUrl;
            _resultValue.httpApiIntraUrl = httpApiIntraUrl;
            _resultValue.id = id;
            _resultValue.pushGateWayInterUrl = pushGateWayInterUrl;
            _resultValue.pushGateWayIntraUrl = pushGateWayIntraUrl;
            _resultValue.remoteReadInterUrl = remoteReadInterUrl;
            _resultValue.remoteReadIntraUrl = remoteReadIntraUrl;
            _resultValue.remoteWriteInterUrl = remoteWriteInterUrl;
            _resultValue.remoteWriteIntraUrl = remoteWriteIntraUrl;
            _resultValue.resourceGroupId = resourceGroupId;
            _resultValue.securityGroupId = securityGroupId;
            _resultValue.subClustersJson = subClustersJson;
            _resultValue.tags = tags;
            _resultValue.vpcId = vpcId;
            _resultValue.vswitchId = vswitchId;
            return _resultValue;
        }
    }
}
