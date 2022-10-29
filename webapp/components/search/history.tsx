import React, {useEffect, useState} from 'react'
import {ItemType, PaginatedResults, PaginationInfo, RadarrSearchRecord, SonarrSearchRecord} from "../../types/tasks";
import {getSearchHistory, isRadarrSearchRecord} from "../../service/search";
import {OnPageClickParams, Table} from "../table/table";
import {getRelativeDateString} from "../../utils/utils";

export interface SearchHistoryProps {
    type: ItemType
}

function getHeaders(): string[] {
    return [
        "Searched At",
        "Search Task ID",
        "Searched Item",
    ]
}

function getData(results: PaginatedResults<RadarrSearchRecord | SonarrSearchRecord>): string[][] {
    return results.items.map((item: SonarrSearchRecord | RadarrSearchRecord) => {
        if (isRadarrSearchRecord(item)) {
            return [getRelativeDateString(item.CreatedAt), item.task_id.toString(), item.item.name]
        }
        return [getRelativeDateString(item.CreatedAt), item.task_id.toString(), `${item.tv_item.series_title} - S${item.tv_item.season}E${item.tv_item.season_episode_number} ${item.tv_item.name}`]
    })
}

export const SearchHistory: React.FC<SearchHistoryProps> = (props: SearchHistoryProps) => {
    const [searchHistory, setSearchHistory] = useState({} as PaginatedResults<RadarrSearchRecord | SonarrSearchRecord>)
    const onPageClick = async (event: React.MouseEvent<HTMLButtonElement> | null, page : number) => {

        const res = await getSearchHistory(props.type, page)
        if (res !== searchHistory) {
            setSearchHistory(res)
        }
    }

    useEffect(() => {
        onPageClick(null, 0)
    }, [props]);

    return searchHistory && searchHistory.items && (
        <Table headers={getHeaders()} data={getData(searchHistory)} pagination={searchHistory as PaginationInfo}
               paginationCallback={onPageClick} title={`${props.type} Search History`}/>
    );
};

