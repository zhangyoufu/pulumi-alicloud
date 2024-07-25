// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.directmail.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDomainsDomain {
    /**
     * @return Indicates whether your ownership of the domain is verified.
     * 
     */
    private String cnameAuthStatus;
    /**
     * @return Indicates whether the CNAME record is successfully verified. **Note:** `cname_confirm_status` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String cnameConfirmStatus;
    /**
     * @return The value of the CNAME record. **Note:** `cname_record` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String cnameRecord;
    /**
     * @return The time when the DNS record was created.
     * 
     */
    private String createTime;
    /**
     * @return The default domain name. **Note:** `default_domain` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String defaultDomain;
    /**
     * @return (Available since v1.227.1) The DKIM validation flag. **Note:** `dkim_auth_status` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String dkimAuthStatus;
    /**
     * @return (Available since v1.227.1) The DKIM public key. **Note:** `dkim_public_key` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String dkimPublicKey;
    /**
     * @return (Available since v1.227.1) The DKIM Host Record. **Note:** `dkim_rr` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String dkimRr;
    /**
     * @return (Available since v1.227.1) The DMARC validation flag. **Note:** `dmarc_auth_status` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String dmarcAuthStatus;
    /**
     * @return (Available since v1.227.1) The DMARC Host Record. **Note:** `dmarc_host_record` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String dmarcHostRecord;
    /**
     * @return (Available since v1.227.1) The DMARC record. **Note:** `dmarc_record` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String dmarcRecord;
    /**
     * @return (Available since v1.227.1) The DMARC record value resolved through public DNS. **Note:** `dns_dmarc` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String dnsDmarc;
    /**
     * @return The MX record value resolved through public DNS. **Note:** `dns_mx` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String dnsMx;
    /**
     * @return The SPF record value resolved through public DNS. **Note:** `dns_spf` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String dnsSpf;
    /**
     * @return The TXT record value resolved through public DNS. **Note:** `dns_txt` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String dnsTxt;
    /**
     * @return The ID of the domain name.
     * 
     */
    private String domainId;
    /**
     * @return The domain name.
     * 
     */
    private String domainName;
    /**
     * @return (Available since v1.227.1) The value of the Domain record.
     * 
     */
    private String domainRecord;
    /**
     * @return The type of the domain. **Note:** `domain_type` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String domainType;
    /**
     * @return (Available since v1.227.1) The value of the host record. **Note:** `host_record` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String hostRecord;
    /**
     * @return The status of ICP filing.
     * 
     */
    private String icpStatus;
    /**
     * @return The ID of the Domain.
     * 
     */
    private String id;
    /**
     * @return Indicates whether the MX record is successfully verified.
     * 
     */
    private String mxAuthStatus;
    /**
     * @return The MX verification record provided by the Direct Mail console. **Note:** `mx_record` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String mxRecord;
    /**
     * @return Indicates whether the SPF record is successfully verified.
     * 
     */
    private String spfAuthStatus;
    /**
     * @return The SPF verification record provided by the Direct Mail console. **Note:** `spf_record` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String spfRecord;
    /**
     * @return The status of the domain name. Valid values:
     * 
     */
    private String status;
    /**
     * @return The primary domain name. **Note:** `tl_domain_name` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String tlDomainName;
    /**
     * @return The CNAME verification record provided by the Direct Mail console. **Note:** `tracef_record` takes effect only if `enable_details` is set to `true`.
     * 
     */
    private String tracefRecord;

    private GetDomainsDomain() {}
    /**
     * @return Indicates whether your ownership of the domain is verified.
     * 
     */
    public String cnameAuthStatus() {
        return this.cnameAuthStatus;
    }
    /**
     * @return Indicates whether the CNAME record is successfully verified. **Note:** `cname_confirm_status` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String cnameConfirmStatus() {
        return this.cnameConfirmStatus;
    }
    /**
     * @return The value of the CNAME record. **Note:** `cname_record` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String cnameRecord() {
        return this.cnameRecord;
    }
    /**
     * @return The time when the DNS record was created.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The default domain name. **Note:** `default_domain` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String defaultDomain() {
        return this.defaultDomain;
    }
    /**
     * @return (Available since v1.227.1) The DKIM validation flag. **Note:** `dkim_auth_status` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String dkimAuthStatus() {
        return this.dkimAuthStatus;
    }
    /**
     * @return (Available since v1.227.1) The DKIM public key. **Note:** `dkim_public_key` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String dkimPublicKey() {
        return this.dkimPublicKey;
    }
    /**
     * @return (Available since v1.227.1) The DKIM Host Record. **Note:** `dkim_rr` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String dkimRr() {
        return this.dkimRr;
    }
    /**
     * @return (Available since v1.227.1) The DMARC validation flag. **Note:** `dmarc_auth_status` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String dmarcAuthStatus() {
        return this.dmarcAuthStatus;
    }
    /**
     * @return (Available since v1.227.1) The DMARC Host Record. **Note:** `dmarc_host_record` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String dmarcHostRecord() {
        return this.dmarcHostRecord;
    }
    /**
     * @return (Available since v1.227.1) The DMARC record. **Note:** `dmarc_record` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String dmarcRecord() {
        return this.dmarcRecord;
    }
    /**
     * @return (Available since v1.227.1) The DMARC record value resolved through public DNS. **Note:** `dns_dmarc` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String dnsDmarc() {
        return this.dnsDmarc;
    }
    /**
     * @return The MX record value resolved through public DNS. **Note:** `dns_mx` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String dnsMx() {
        return this.dnsMx;
    }
    /**
     * @return The SPF record value resolved through public DNS. **Note:** `dns_spf` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String dnsSpf() {
        return this.dnsSpf;
    }
    /**
     * @return The TXT record value resolved through public DNS. **Note:** `dns_txt` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String dnsTxt() {
        return this.dnsTxt;
    }
    /**
     * @return The ID of the domain name.
     * 
     */
    public String domainId() {
        return this.domainId;
    }
    /**
     * @return The domain name.
     * 
     */
    public String domainName() {
        return this.domainName;
    }
    /**
     * @return (Available since v1.227.1) The value of the Domain record.
     * 
     */
    public String domainRecord() {
        return this.domainRecord;
    }
    /**
     * @return The type of the domain. **Note:** `domain_type` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String domainType() {
        return this.domainType;
    }
    /**
     * @return (Available since v1.227.1) The value of the host record. **Note:** `host_record` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String hostRecord() {
        return this.hostRecord;
    }
    /**
     * @return The status of ICP filing.
     * 
     */
    public String icpStatus() {
        return this.icpStatus;
    }
    /**
     * @return The ID of the Domain.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Indicates whether the MX record is successfully verified.
     * 
     */
    public String mxAuthStatus() {
        return this.mxAuthStatus;
    }
    /**
     * @return The MX verification record provided by the Direct Mail console. **Note:** `mx_record` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String mxRecord() {
        return this.mxRecord;
    }
    /**
     * @return Indicates whether the SPF record is successfully verified.
     * 
     */
    public String spfAuthStatus() {
        return this.spfAuthStatus;
    }
    /**
     * @return The SPF verification record provided by the Direct Mail console. **Note:** `spf_record` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String spfRecord() {
        return this.spfRecord;
    }
    /**
     * @return The status of the domain name. Valid values:
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The primary domain name. **Note:** `tl_domain_name` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String tlDomainName() {
        return this.tlDomainName;
    }
    /**
     * @return The CNAME verification record provided by the Direct Mail console. **Note:** `tracef_record` takes effect only if `enable_details` is set to `true`.
     * 
     */
    public String tracefRecord() {
        return this.tracefRecord;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDomainsDomain defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cnameAuthStatus;
        private String cnameConfirmStatus;
        private String cnameRecord;
        private String createTime;
        private String defaultDomain;
        private String dkimAuthStatus;
        private String dkimPublicKey;
        private String dkimRr;
        private String dmarcAuthStatus;
        private String dmarcHostRecord;
        private String dmarcRecord;
        private String dnsDmarc;
        private String dnsMx;
        private String dnsSpf;
        private String dnsTxt;
        private String domainId;
        private String domainName;
        private String domainRecord;
        private String domainType;
        private String hostRecord;
        private String icpStatus;
        private String id;
        private String mxAuthStatus;
        private String mxRecord;
        private String spfAuthStatus;
        private String spfRecord;
        private String status;
        private String tlDomainName;
        private String tracefRecord;
        public Builder() {}
        public Builder(GetDomainsDomain defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cnameAuthStatus = defaults.cnameAuthStatus;
    	      this.cnameConfirmStatus = defaults.cnameConfirmStatus;
    	      this.cnameRecord = defaults.cnameRecord;
    	      this.createTime = defaults.createTime;
    	      this.defaultDomain = defaults.defaultDomain;
    	      this.dkimAuthStatus = defaults.dkimAuthStatus;
    	      this.dkimPublicKey = defaults.dkimPublicKey;
    	      this.dkimRr = defaults.dkimRr;
    	      this.dmarcAuthStatus = defaults.dmarcAuthStatus;
    	      this.dmarcHostRecord = defaults.dmarcHostRecord;
    	      this.dmarcRecord = defaults.dmarcRecord;
    	      this.dnsDmarc = defaults.dnsDmarc;
    	      this.dnsMx = defaults.dnsMx;
    	      this.dnsSpf = defaults.dnsSpf;
    	      this.dnsTxt = defaults.dnsTxt;
    	      this.domainId = defaults.domainId;
    	      this.domainName = defaults.domainName;
    	      this.domainRecord = defaults.domainRecord;
    	      this.domainType = defaults.domainType;
    	      this.hostRecord = defaults.hostRecord;
    	      this.icpStatus = defaults.icpStatus;
    	      this.id = defaults.id;
    	      this.mxAuthStatus = defaults.mxAuthStatus;
    	      this.mxRecord = defaults.mxRecord;
    	      this.spfAuthStatus = defaults.spfAuthStatus;
    	      this.spfRecord = defaults.spfRecord;
    	      this.status = defaults.status;
    	      this.tlDomainName = defaults.tlDomainName;
    	      this.tracefRecord = defaults.tracefRecord;
        }

        @CustomType.Setter
        public Builder cnameAuthStatus(String cnameAuthStatus) {
            if (cnameAuthStatus == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "cnameAuthStatus");
            }
            this.cnameAuthStatus = cnameAuthStatus;
            return this;
        }
        @CustomType.Setter
        public Builder cnameConfirmStatus(String cnameConfirmStatus) {
            if (cnameConfirmStatus == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "cnameConfirmStatus");
            }
            this.cnameConfirmStatus = cnameConfirmStatus;
            return this;
        }
        @CustomType.Setter
        public Builder cnameRecord(String cnameRecord) {
            if (cnameRecord == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "cnameRecord");
            }
            this.cnameRecord = cnameRecord;
            return this;
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            if (createTime == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "createTime");
            }
            this.createTime = createTime;
            return this;
        }
        @CustomType.Setter
        public Builder defaultDomain(String defaultDomain) {
            if (defaultDomain == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "defaultDomain");
            }
            this.defaultDomain = defaultDomain;
            return this;
        }
        @CustomType.Setter
        public Builder dkimAuthStatus(String dkimAuthStatus) {
            if (dkimAuthStatus == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "dkimAuthStatus");
            }
            this.dkimAuthStatus = dkimAuthStatus;
            return this;
        }
        @CustomType.Setter
        public Builder dkimPublicKey(String dkimPublicKey) {
            if (dkimPublicKey == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "dkimPublicKey");
            }
            this.dkimPublicKey = dkimPublicKey;
            return this;
        }
        @CustomType.Setter
        public Builder dkimRr(String dkimRr) {
            if (dkimRr == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "dkimRr");
            }
            this.dkimRr = dkimRr;
            return this;
        }
        @CustomType.Setter
        public Builder dmarcAuthStatus(String dmarcAuthStatus) {
            if (dmarcAuthStatus == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "dmarcAuthStatus");
            }
            this.dmarcAuthStatus = dmarcAuthStatus;
            return this;
        }
        @CustomType.Setter
        public Builder dmarcHostRecord(String dmarcHostRecord) {
            if (dmarcHostRecord == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "dmarcHostRecord");
            }
            this.dmarcHostRecord = dmarcHostRecord;
            return this;
        }
        @CustomType.Setter
        public Builder dmarcRecord(String dmarcRecord) {
            if (dmarcRecord == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "dmarcRecord");
            }
            this.dmarcRecord = dmarcRecord;
            return this;
        }
        @CustomType.Setter
        public Builder dnsDmarc(String dnsDmarc) {
            if (dnsDmarc == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "dnsDmarc");
            }
            this.dnsDmarc = dnsDmarc;
            return this;
        }
        @CustomType.Setter
        public Builder dnsMx(String dnsMx) {
            if (dnsMx == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "dnsMx");
            }
            this.dnsMx = dnsMx;
            return this;
        }
        @CustomType.Setter
        public Builder dnsSpf(String dnsSpf) {
            if (dnsSpf == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "dnsSpf");
            }
            this.dnsSpf = dnsSpf;
            return this;
        }
        @CustomType.Setter
        public Builder dnsTxt(String dnsTxt) {
            if (dnsTxt == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "dnsTxt");
            }
            this.dnsTxt = dnsTxt;
            return this;
        }
        @CustomType.Setter
        public Builder domainId(String domainId) {
            if (domainId == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "domainId");
            }
            this.domainId = domainId;
            return this;
        }
        @CustomType.Setter
        public Builder domainName(String domainName) {
            if (domainName == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "domainName");
            }
            this.domainName = domainName;
            return this;
        }
        @CustomType.Setter
        public Builder domainRecord(String domainRecord) {
            if (domainRecord == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "domainRecord");
            }
            this.domainRecord = domainRecord;
            return this;
        }
        @CustomType.Setter
        public Builder domainType(String domainType) {
            if (domainType == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "domainType");
            }
            this.domainType = domainType;
            return this;
        }
        @CustomType.Setter
        public Builder hostRecord(String hostRecord) {
            if (hostRecord == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "hostRecord");
            }
            this.hostRecord = hostRecord;
            return this;
        }
        @CustomType.Setter
        public Builder icpStatus(String icpStatus) {
            if (icpStatus == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "icpStatus");
            }
            this.icpStatus = icpStatus;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder mxAuthStatus(String mxAuthStatus) {
            if (mxAuthStatus == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "mxAuthStatus");
            }
            this.mxAuthStatus = mxAuthStatus;
            return this;
        }
        @CustomType.Setter
        public Builder mxRecord(String mxRecord) {
            if (mxRecord == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "mxRecord");
            }
            this.mxRecord = mxRecord;
            return this;
        }
        @CustomType.Setter
        public Builder spfAuthStatus(String spfAuthStatus) {
            if (spfAuthStatus == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "spfAuthStatus");
            }
            this.spfAuthStatus = spfAuthStatus;
            return this;
        }
        @CustomType.Setter
        public Builder spfRecord(String spfRecord) {
            if (spfRecord == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "spfRecord");
            }
            this.spfRecord = spfRecord;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tlDomainName(String tlDomainName) {
            if (tlDomainName == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "tlDomainName");
            }
            this.tlDomainName = tlDomainName;
            return this;
        }
        @CustomType.Setter
        public Builder tracefRecord(String tracefRecord) {
            if (tracefRecord == null) {
              throw new MissingRequiredPropertyException("GetDomainsDomain", "tracefRecord");
            }
            this.tracefRecord = tracefRecord;
            return this;
        }
        public GetDomainsDomain build() {
            final var _resultValue = new GetDomainsDomain();
            _resultValue.cnameAuthStatus = cnameAuthStatus;
            _resultValue.cnameConfirmStatus = cnameConfirmStatus;
            _resultValue.cnameRecord = cnameRecord;
            _resultValue.createTime = createTime;
            _resultValue.defaultDomain = defaultDomain;
            _resultValue.dkimAuthStatus = dkimAuthStatus;
            _resultValue.dkimPublicKey = dkimPublicKey;
            _resultValue.dkimRr = dkimRr;
            _resultValue.dmarcAuthStatus = dmarcAuthStatus;
            _resultValue.dmarcHostRecord = dmarcHostRecord;
            _resultValue.dmarcRecord = dmarcRecord;
            _resultValue.dnsDmarc = dnsDmarc;
            _resultValue.dnsMx = dnsMx;
            _resultValue.dnsSpf = dnsSpf;
            _resultValue.dnsTxt = dnsTxt;
            _resultValue.domainId = domainId;
            _resultValue.domainName = domainName;
            _resultValue.domainRecord = domainRecord;
            _resultValue.domainType = domainType;
            _resultValue.hostRecord = hostRecord;
            _resultValue.icpStatus = icpStatus;
            _resultValue.id = id;
            _resultValue.mxAuthStatus = mxAuthStatus;
            _resultValue.mxRecord = mxRecord;
            _resultValue.spfAuthStatus = spfAuthStatus;
            _resultValue.spfRecord = spfRecord;
            _resultValue.status = status;
            _resultValue.tlDomainName = tlDomainName;
            _resultValue.tracefRecord = tracefRecord;
            return _resultValue;
        }
    }
}
