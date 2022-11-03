import React, {useEffect, useState} from 'react'
import {
    ItemType,
    PaginatedResults,
    PaginationInfo,
    RadarrSearchRecord,
    SonarrSearchRecord,
    TaskInfoResponse
} from "../../types/tasks";
import {getSearchHistory, isRadarrSearchRecord, performSearch, performSync} from "../../service/search";
import {getTaskInfo as svcGetTaskInfo} from "../../service/tasks";
import {Table} from "../table/table";
import {getRelativeDateString} from "../../utils/utils";
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';
import {TaskInfoCard} from "../task/infoCard";
import SearchIcon from '@mui/icons-material/Search';
import SyncIcon from '@mui/icons-material/Sync';
import {useInterval} from "usehooks-ts";

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
            return [getRelativeDateString(item.CreatedAt), item.task_id.toString(), item.radarr_item.name]
        }
        return [getRelativeDateString(item.CreatedAt), item.task_id.toString(), `${item.sonarr_item.series_title} - S${item.sonarr_item.season}E${item.sonarr_item.season_episode_number} ${item.sonarr_item.name}`]
    })
}

export const SearchHistory: React.FC<SearchHistoryProps> = (props: SearchHistoryProps) => {
    const [searchHistory, setSearchHistory] = useState({} as PaginatedResults<RadarrSearchRecord | SonarrSearchRecord>)
    const [taskInfo, setTaskInfo] = useState({} as TaskInfoResponse)
    const [page, setPage] = useState(0)
    const onPageClick = async (event: React.MouseEvent<HTMLButtonElement> | null, newPage: number) => {
        if(newPage!= page){
            setPage(newPage)
        }
        const res = await getSearchHistory(props.type, page)
        if (res !== searchHistory) {
            setSearchHistory(res)
        }
    }
    const getTaskInfo = async () => {
        const res = await svcGetTaskInfo()
        if (res !== taskInfo) {

            setTaskInfo(res)
        }
    }
    useInterval(()=>{
        onPageClick(null, page)
        getTaskInfo()
    }, 3000)

    useEffect(() => {
        onPageClick(null, 0)
        getTaskInfo()

    }, [props]);

    return searchHistory && searchHistory.items && (
        <>
            {
                taskInfo &&
                <Box sx={{mt: 4, mb: 4}}>
                    <Grid container spacing={2}>
                        <Grid item xs={12} md={6}>
                            <TaskInfoCard
                                name={`${props.type} Import`}
                                prevRun={taskInfo[`${props.type}_import`].prev_run_at}
                                nextRun={taskInfo[`${props.type}_import`].next_run_at}
                                actionIcon={<SyncIcon/>}
                                actionCallback={() => {
                                    performSync(props.type)
                                    console.log("sync")
                                }}
                            />
                        </Grid>
                        <Grid item xs={12} md={6}>
                            <TaskInfoCard
                                name={`${props.type} Search`}
                                prevRun={taskInfo[`${props.type}_search`].prev_run_at}
                                nextRun={taskInfo[`${props.type}_search`].next_run_at}
                                actionIcon={<SearchIcon/>}
                                actionCallback={() => {
                                    performSearch(props.type)
                                    console.log("search")
                                }}
                            />
                        </Grid>
                    </Grid>
                </Box>
            }
            <Table headers={getHeaders()} data={getData(searchHistory)} pagination={searchHistory as PaginationInfo}
                   paginationCallback={onPageClick} title={`${props.type} Search History`}/>
        </>
    )
        ;
};

