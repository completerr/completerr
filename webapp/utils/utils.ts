import {parseISO, formatDistance} from "date-fns";
import axios, {AxiosInstance} from "axios";

const isoDateFormat = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d*)?(?:[-+]\d{2}:?\d{2}|Z)?$/

function isIsoDateString(value: any): boolean {
    return value && typeof value === "string" && isoDateFormat.test(value);
}

export function handleDates(body: any) {
    if (body === null || body === undefined || typeof body !== "object")
        return body;

    for (const key of Object.keys(body)) {
        const value = body[key];
        if (isIsoDateString(value)) body[key] = parseISO(value);
        else if (typeof value === "object") handleDates(value);
    }
}

export function getAxios(): AxiosInstance {
    const client = axios.create();
    client.interceptors.response.use(originalResponse => {
        handleDates(originalResponse.data);
        return originalResponse;
    });
    return client
}

export function getRelativeDateString(date : Date) : string {
   return formatDistance(date, new Date() , {addSuffix:true})
}