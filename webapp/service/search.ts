import axios, {AxiosResponse} from "axios";
import {ItemType, PaginatedResults, SearchRecord, RadarrSearchRecord, SonarrSearchRecord, Task} from "../types/tasks";

export const getSearchHistory = async (itemType : ItemType, page = 0) : Promise<PaginatedResults<RadarrSearchRecord|SonarrSearchRecord>> => {

    let resp = await axios.get<{}, AxiosResponse<PaginatedResults<RadarrSearchRecord|SonarrSearchRecord>>>(`/api/${itemType}/history`, {params:{page, size:50}})
    return resp.data
}
export function isRadarrSearchRecord(record:  RadarrSearchRecord | SonarrSearchRecord) :record is RadarrSearchRecord{
    if ("item_id" in record){
        return record.item_id > 0
    }
    return false
}