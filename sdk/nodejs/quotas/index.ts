// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ApplicationInfoArgs, ApplicationInfoState } from "./applicationInfo";
export type ApplicationInfo = import("./applicationInfo").ApplicationInfo;
export const ApplicationInfo: typeof import("./applicationInfo").ApplicationInfo = null as any;
utilities.lazyLoad(exports, ["ApplicationInfo"], () => require("./applicationInfo"));

export { GetApplicationInfosArgs, GetApplicationInfosResult, GetApplicationInfosOutputArgs } from "./getApplicationInfos";
export const getApplicationInfos: typeof import("./getApplicationInfos").getApplicationInfos = null as any;
export const getApplicationInfosOutput: typeof import("./getApplicationInfos").getApplicationInfosOutput = null as any;
utilities.lazyLoad(exports, ["getApplicationInfos","getApplicationInfosOutput"], () => require("./getApplicationInfos"));

export { GetQuotaAlarmsArgs, GetQuotaAlarmsResult, GetQuotaAlarmsOutputArgs } from "./getQuotaAlarms";
export const getQuotaAlarms: typeof import("./getQuotaAlarms").getQuotaAlarms = null as any;
export const getQuotaAlarmsOutput: typeof import("./getQuotaAlarms").getQuotaAlarmsOutput = null as any;
utilities.lazyLoad(exports, ["getQuotaAlarms","getQuotaAlarmsOutput"], () => require("./getQuotaAlarms"));

export { GetQuotaApplicationsArgs, GetQuotaApplicationsResult, GetQuotaApplicationsOutputArgs } from "./getQuotaApplications";
export const getQuotaApplications: typeof import("./getQuotaApplications").getQuotaApplications = null as any;
export const getQuotaApplicationsOutput: typeof import("./getQuotaApplications").getQuotaApplicationsOutput = null as any;
utilities.lazyLoad(exports, ["getQuotaApplications","getQuotaApplicationsOutput"], () => require("./getQuotaApplications"));

export { GetQuotasArgs, GetQuotasResult, GetQuotasOutputArgs } from "./getQuotas";
export const getQuotas: typeof import("./getQuotas").getQuotas = null as any;
export const getQuotasOutput: typeof import("./getQuotas").getQuotasOutput = null as any;
utilities.lazyLoad(exports, ["getQuotas","getQuotasOutput"], () => require("./getQuotas"));

export { QuotaAlarmArgs, QuotaAlarmState } from "./quotaAlarm";
export type QuotaAlarm = import("./quotaAlarm").QuotaAlarm;
export const QuotaAlarm: typeof import("./quotaAlarm").QuotaAlarm = null as any;
utilities.lazyLoad(exports, ["QuotaAlarm"], () => require("./quotaAlarm"));

export { QuotaApplicationArgs, QuotaApplicationState } from "./quotaApplication";
export type QuotaApplication = import("./quotaApplication").QuotaApplication;
export const QuotaApplication: typeof import("./quotaApplication").QuotaApplication = null as any;
utilities.lazyLoad(exports, ["QuotaApplication"], () => require("./quotaApplication"));

export { TemplateQuotaArgs, TemplateQuotaState } from "./templateQuota";
export type TemplateQuota = import("./templateQuota").TemplateQuota;
export const TemplateQuota: typeof import("./templateQuota").TemplateQuota = null as any;
utilities.lazyLoad(exports, ["TemplateQuota"], () => require("./templateQuota"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:quotas/applicationInfo:ApplicationInfo":
                return new ApplicationInfo(name, <any>undefined, { urn })
            case "alicloud:quotas/quotaAlarm:QuotaAlarm":
                return new QuotaAlarm(name, <any>undefined, { urn })
            case "alicloud:quotas/quotaApplication:QuotaApplication":
                return new QuotaApplication(name, <any>undefined, { urn })
            case "alicloud:quotas/templateQuota:TemplateQuota":
                return new TemplateQuota(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "quotas/applicationInfo", _module)
pulumi.runtime.registerResourceModule("alicloud", "quotas/quotaAlarm", _module)
pulumi.runtime.registerResourceModule("alicloud", "quotas/quotaApplication", _module)
pulumi.runtime.registerResourceModule("alicloud", "quotas/templateQuota", _module)
