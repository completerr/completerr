import React, {useEffect, useState} from 'react'
import {ItemType, PaginatedResults, PaginationInfo, RadarrSearchRecord, SonarrSearchRecord} from "../../types/tasks";
import {getSearchHistory, isRadarrSearchRecord} from "../../service/search";
import {OnPageClickParams, Table} from "../table/table";

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
            return [item.CreatedAt.toLocaleString(), item.task_id.toString(), item.item.name]
        }
        return [item.CreatedAt.toLocaleString(), item.task_id.toString(), `${item.tv_item.series_title} - S${item.tv_item.season}E${item.tv_item.season_episode_number} ${item.tv_item.name}`]
    })
}

export const SearchHistory: React.FC<SearchHistoryProps> = (props: SearchHistoryProps) => {
    const [searchHistory, setSearchHistory] = useState({} as PaginatedResults<RadarrSearchRecord | SonarrSearchRecord>)
    const onPageClick = async ({selected = 0}: OnPageClickParams) => {

        const res = await getSearchHistory(props.type, selected)
        if (res !== searchHistory) {
            setSearchHistory(res)
        }
    }

    useEffect(() => {
        onPageClick({selected: 0})
    }, []);

    return searchHistory && searchHistory.items && (
        <Table headers={getHeaders()} data={getData(searchHistory)} pagination={searchHistory as PaginationInfo}
               paginationCallback={onPageClick} title={`${props.type} Search History`}/>
    );
};

